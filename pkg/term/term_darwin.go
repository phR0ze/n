// +build darwin

package term

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	ioctlReadTermios  = unix.TIOCGETA
	ioctlWriteTermios = unix.TIOCSETA
)

// IsTTY simply checks if we are working with a TTY on os.Stdout
func IsTTY() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, os.Stdout.Fd(), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

// IsTTYP simply checks if we are working with a TTY File Descriptor
func IsTTYP(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
