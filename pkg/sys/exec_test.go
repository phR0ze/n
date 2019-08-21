package sys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecOut(t *testing.T) {
	// Invalid command
	{
		result, err := ExecOut("blah")
		expected := ""
		assert.Equal(t, expected, result)
		assert.Equal(t, `exec: "blah": executable file not found in $PATH`, err.Error())
	}

	// Valid listing of a directory
	{
		result, err := ExecOut("ls -1 ../net")
		assert.Nil(t, err)
		expected := "agent\nmech\nnet.go\nnet_test.go\n"
		assert.Equal(t, expected, result)
	}
}
