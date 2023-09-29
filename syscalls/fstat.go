//go:build !(linux && loong64)

package syscalls

import "syscall"

func init() {
	syscallNames[syscall.SYS_FSTAT] = "FSTAT"
}
