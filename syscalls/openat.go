//go:build !(darwin && amd64) && !(darwin && arm64)

package syscalls

import "syscall"

func init() {
	syscallNames[syscall.SYS_OPENAT] = "OPENAT"
}
