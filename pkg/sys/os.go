// Package sys provides os level helper functions for interacting with the system
package sys

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"unsafe"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

// Cache if bash exists to avoid wasting overhead checking
var (
	gBashExists      bool
	gBashExistsTried bool
	gRXKernelVersion = regexp.MustCompile(`\d+\.\d+\.\d+`)
)

// Kernel information
type Kernel struct {
	Arch    string // x86_64
	Release string // 5.3.1-arch1-1-ARCH
	Version string // 5.3.1
}

// // Capture the output of a function as a string
// func Capture(f func()) string {
// 	var buf bytes.Buffer // create a buffer to store your data
// 	stdout := os.Stdout  // save off original stdout to replace
// 	r, w, _ := os.Pipe() // create a pipe to be able to read what the function writes
// 	os.Stdout = w        // pass the pipe write to stdout for the function to write to

// 	// call the function
// 	f()

// 	// Clean up
// 	os.Stdout = stdout // restore stdout to its original
// 	return bufout.String()
// }

// Darwin returns true if the OS is OSX
func Darwin() (result bool) {
	if runtime.GOOS == "darwin" {
		result = true
	}
	return
}

// Linux returns true if the OS is Linux
func Linux() (result bool) {
	if runtime.GOOS == "linux" {
		result = true
	}
	return
}

// Windows returns true if the OS is Windows
func Windows() (result bool) {
	if runtime.GOOS == "windows" {
		result = true
	}
	return
}

// ExecExists checks if the given executable exists on the PATH
func ExecExists(target string) (ok bool) {
	if path := ExecPath(target); path != "" {
		ok = true
	}
	return
}

// ExecOut executes the given command and returns the output as a string.
func ExecOut(str string, a ...interface{}) (out string, err error) {

	// Parse command
	cmd := fmt.Sprintf(str, a...)
	pieces := SplitCmd(cmd)
	if len(pieces) == 0 {
		err = errors.Errorf("invalid empty command")
		return
	}

	p := exec.Command(pieces[0], pieces[1:]...)
	var output []byte
	output, err = p.CombinedOutput()
	out = string(output)
	if err != nil {
		err = errors.Wrap(err, "failed to execute system command")
	}

	return
}

// ExecPath wraps exec.LookPath to find the path of an executable using the PATH environment variable.
// Returns an empty string if not found.
func ExecPath(target string) (path string) {
	var err error
	if path, err = exec.LookPath(target); err != nil {
		path = ""
	}
	return
}

// KernelInfo ...
func KernelInfo() (kernel *Kernel, err error) {
	kernel = &Kernel{}

	// Get the system info and convert it to sanity
	var uname unix.Utsname
	if err = unix.Uname(&uname); err != nil {
		return
	}
	kernel.Release = strings.TrimRight(string((*[65]byte)(unsafe.Pointer(&uname.Release))[:]), "\000")
	kernel.Arch = strings.TrimRight(string((*[65]byte)(unsafe.Pointer(&uname.Machine))[:]), "\000")

	// Parse out the kernel version
	if match := gRXKernelVersion.FindString(kernel.Release); match != "" {
		kernel.Version = match
	}

	return
}

// Shell executes the given command using a shell and returns the output as a string.
// Supports string interpolation like fmt.Sprintf
func Shell(str string, a ...interface{}) (out string, err error) {
	cmd := "bash"
	if !gBashExistsTried {
		gBashExists = ExecExists("bash")
		gBashExistsTried = true
	}
	if !gBashExists {
		cmd = "sh"
	}

	// Check command validity
	if strings.TrimSpace(str) == "" {
		err = errors.Errorf("invalid empty command")
		return
	}

	// Execute the command
	p := exec.Command(cmd, "-c", fmt.Sprintf(str, a...))
	var output []byte
	output, err = p.CombinedOutput()
	out = string(output)
	if err != nil {
		err = errors.Wrap(err, "failed to execute system command")
	}

	return
}
