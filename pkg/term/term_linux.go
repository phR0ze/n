// +build linux

package term

import (
	"os"

	"golang.org/x/sys/unix"
)

const (
	ioctlReadTermios  = unix.TCGETS
	ioctlWriteTermios = unix.TCSETS
)

// IsTTY simply checks if we are working with a TTY on os.Stdout
func IsTTY() bool {
	_, err := unix.IoctlGetTermios(int(os.Stdout.Fd()), unix.TCGETS)
	return err == nil
}

// IsTTYP simply checks if we are working with a TTY File Descriptor
func IsTTYP(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	return err == nil
}
