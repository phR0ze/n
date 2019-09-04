package sys

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

// ExecExists checks if the given executable exists on the PATH
func ExecExists(target string) (ok bool) {
	if path := ExecPath(target); path != "" {
		ok = true
	}
	return
}

// ExecOut executes the given command and returns the output as a string.
// Supports string interpolation like fmt.Sprintf
func ExecOut(str string, a ...interface{}) (out string, err error) {
	cmd := fmt.Sprintf(str, a...)
	pieces := SplitCmd(cmd)
	if len(pieces) == 0 {
		err = errors.Errorf("Failed to execute empty command")
		return
	}

	// Execute the command
	p := exec.Command(pieces[0], pieces[1:]...)
	var output []byte
	output, err = p.CombinedOutput()
	out = string(output)

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
