package sys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecExists(t *testing.T) {
	// happy
	{
		assert.True(t, ExecExists("sh"))
	}

	// sad
	{
		assert.False(t, ExecExists("footmp"))
	}
}

func TestExecOut(t *testing.T) {

	// Empty command
	{
		result, err := ExecOut("")
		assert.Equal(t, "", result)
		assert.Equal(t, "Failed to execute empty command", err.Error())
	}

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

func TestExecPath(t *testing.T) {
	// happy
	{
		path := ExecPath("sh")
		assert.True(t, path == "/bin/sh" || path == "/usr/bin/sh")
	}

	// sad
	{
		assert.Equal(t, "", ExecPath("footmp"))
	}
}
