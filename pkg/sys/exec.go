package sys

import (
	"fmt"
	"os/exec"
)

// ExecOut executes the given command and get the output as a string
func ExecOut(str string, a ...interface{}) (out string, err error) {
	cmd := fmt.Sprintf(str, a...)
	pieces := SplitCmd(cmd)
	if len(pieces) == 0 {
		return
	}

	// Execute the command
	p := exec.Command(pieces[0], pieces[1:]...)
	var output []byte
	output, err = p.CombinedOutput()
	out = string(output)

	return
}
