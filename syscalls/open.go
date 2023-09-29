//go:build !(linux && riscv64) && !(linux && arm64) && !(linux && loong64)

package syscalls

import "syscall"

func init() {
	syscallNames[syscall.SYS_OPEN] = "OPEN"
}
