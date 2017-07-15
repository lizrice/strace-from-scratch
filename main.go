package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var regs syscall.PtraceRegs
	var ss syscallCounter

	ss = ss.init()

	fmt.Printf("Run %v\n", os.Args[1:])

	// Uncommenting this will cause the open syscall to return with Operation Not Permitted error
	// disallow("open")

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	exit := true

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}

			// Uncomment to show each syscall as it's called
			// name := ss.getName(regs.Orig_rax)
			// fmt.Printf("%s\n", name)
			ss.inc(regs.Orig_rax)
		}

		err = syscall.PtraceSyscall(pid, 0)
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
