// +build darwin

package main

import (
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

func isTTY(f *os.File) bool {
	_, err := getTerminalIOs(f)
	return err == nil
}

func terminalSize(f *os.File) (int, int, error) {
	var dims [4]uint16
	err := ioctl(f, unix.TIOCGWINSZ, unsafe.Pointer(&dims))
	return int(dims[1]), int(dims[0]), err
}

func getTerminalIOs(f *os.File) (*unix.Termios, error) {
	termios := new(unix.Termios)
	err := ioctl(f, unix.TIOCGETA, unsafe.Pointer(termios))
	return termios, err
}

func setTerminalIOs(f *os.File, termios *unix.Termios) error {
	return ioctl(f, unix.TIOCSETA, unsafe.Pointer(termios))
}

func ioctl(f *os.File, typ uintptr, arg unsafe.Pointer) error {
	_, _, err := unix.Syscall(unix.SYS_IOCTL, f.Fd(), typ, uintptr(arg))
	if err != 0 {
		return err
	}

	return nil
}
