// +build !windows

package main

import (
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func makeRaw() {
	var tios unix.Termios

	// This attempts to replicate the behaviour documented for cfmakeraw in
	// the termios(3) manpage.
	tios.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	tios.Oflag &^= syscall.OPOST
	tios.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	tios.Cflag &^= syscall.CSIZE | syscall.PARENB
	tios.Cflag |= syscall.CS8

	tios.Cc[syscall.VMIN] = 1
	tios.Cc[syscall.VTIME] = 0

	setTerminalIOs(os.Stdin, &tios)
}
