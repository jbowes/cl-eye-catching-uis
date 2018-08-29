package main

import (
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

var kernel32 = windows.NewLazySystemDLL("kernel32.dll")

var (
	getConsoleMode = kernel32.NewProc("GetConsoleMode")
	setConsoleMode = kernel32.NewProc("SetConsoleMode")
)

func initOutput() {
	var st uint32
	r, _, _ := getConsoleMode.Call(os.Stdout.Fd(), uintptr(unsafe.Pointer(&st)))
	if r == 0 {
		return // not a windows console, or some other error ;)
	}

	st |= 0x0004 // ENABLE_VIRTUAL_TERMINAL_PROCESSING
	setConsoleMode.Call(os.Stdout.Fd(), uintptr(st))
}
