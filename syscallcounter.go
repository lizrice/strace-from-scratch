package main

import (
	"fmt"
	"os"
	"strace/syscalls"
	"text/tabwriter"
)

type syscallCounter []int

const maxSyscalls = 303

func (s syscallCounter) init() syscallCounter {
	s = make(syscallCounter, maxSyscalls)
	return s
}

func (s syscallCounter) inc(syscallID int) error {
	if syscallID > maxSyscalls {
		return fmt.Errorf("invalid syscall ID (%x)", syscallID)
	}

	s[syscallID]++
	return nil
}

func (s syscallCounter) print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for k, v := range s {
		if v > 0 {
			name := s.getName(k)
			fmt.Fprintf(w, "%d\t%s\n", v, name)
		}
	}
	w.Flush()
}

func (s syscallCounter) getName(syscallID int) string {
	name := syscalls.GetName(syscallID)
	return name
}
