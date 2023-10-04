package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var ss syscallCounter

	ss = ss.init()

	fmt.Printf("Run %v\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait() // cmd is paused
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	var regs syscall.PtraceRegs
	pid := cmd.Process.Pid
	exit := true

	for {
		if err = syscall.PtraceGetRegs(pid, &regs); err != nil {
			if err.Error() == "no such process" {
				break
			}
			panic(err)
		}
		if exit {

			// https://man7.org/linux/man-pages/man2/syscall.2.html
			//   Arch/ABI    arg1  arg2  arg3  arg4  arg5  arg6  arg7   Notes
			//   ────────────────────────────────────────────────────────────
			//   x86-64      rdi   rsi   rdx   r10   r8    r9    -
			//
			//   Arch/ABI    Instruction       System  Ret  Ret  Error  Notes
			//                                 call #  val  val2
			//   ────────────────────────────────────────────────────────────
			//   x86-64      syscall           rax     rax  rdx  -      5

			f := ""
			if regs.Orig_rax == syscall.SYS_OPEN {
				fmt.Printf("%d %x\n", regs.Rdi, regs.Rdi)
				path := readPtraceText(pid, uintptr(regs.Rdi))
				fd := int(regs.Rax)
				f = fmt.Sprintf(`("%s") => %d`, path, fd)
			}

			name := ss.getName(regs.Orig_rax)
			fmt.Printf("%s %s\n", name, f)
			ss.inc(regs.Orig_rax)
		}

		err = syscall.PtraceSyscall(pid, 0) // wait for next syscall to begin or exit
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		exit = !exit
	}

	ss.print()
}

func readPtraceText(pid int, addr uintptr) string {
	s := ""
	buf := []byte{1}
	for i := addr; buf[0] != byte(0); i++ {
		if c, err := syscall.PtracePeekText(pid, i, buf); err != nil {
			panic(err)
		} else if c == 0 {
			break
		}
		s += string(buf)
	}
	return s
}
