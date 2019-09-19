package sys

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Cache if bash exists to avoid wasting overhead checking
var bashExists bool
var bashExistsTried bool

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

// Shell executes the given command using a shell and returns the output as a string.
// Supports string interpolation like fmt.Sprintf
func Shell(str string, a ...interface{}) (out string, err error) {
	cmd := "bash"
	if !bashExistsTried {
		bashExists = ExecExists("bash")
		bashExistsTried = true
	}
	if !bashExists {
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
