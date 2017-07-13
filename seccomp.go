package main

import (
	"syscall"

	sec "github.com/seccomp/libseccomp-golang"
)

func disallow(sc string) {
	id, err := sec.GetSyscallFromName(sc)
	if err != nil {
		panic(err)
	}

	filter, _ := sec.NewFilter(sec.ActAllow)
	filter.AddRule(id, sec.ActErrno.SetReturnCode(int16(syscall.EPERM)))
	filter.Load()
}
