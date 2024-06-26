package sys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

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
		assert.Equal(t, "invalid empty command", err.Error())
	}

	// Invalid command
	{
		result, err := ExecOut("blah")
		expected := ""
		assert.Equal(t, expected, result)
		assert.Equal(t, `failed to execute system command: exec: "blah": executable file not found in $PATH`, err.Error())
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
		assert.True(t, path == "/bin/sh" || path == "/usr/bin/sh" || path == "/run/current-system/sw/bin/sh")
	}

	// sad
	{
		assert.Equal(t, "", ExecPath("footmp"))
	}
}

func TestKernelInfo(t *testing.T) {
	k, err := KernelInfo()
	assert.Nil(t, err)
	assert.NotNil(t, k)
}

func TestShell(t *testing.T) {
	resetTest()

	// Echo output
	{
		assert.False(t, Exists(tmpfile))
		result, err := Shell("echo 'test' > %s", tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, "", result)
		result, err = ReadString(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, "test\n", result)
		assert.Nil(t, Remove(tmpfile))
		assert.False(t, Exists(tmpfile))

		result, err = ExecOut("echo test > %s", tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, "test > ../../test/temp/.tmp\n", result)
		assert.False(t, Exists(tmpfile))
	}

	// Empty command
	{
		result, err := Shell("")
		assert.Equal(t, "", result)
		assert.Equal(t, "invalid empty command", err.Error())
	}

	// Valid listing of a directory
	{
		result, err := Shell("ls -1 ../net")
		assert.Nil(t, err)
		expected := "agent\nmech\nnet.go\nnet_test.go\n"
		assert.Equal(t, expected, result)
	}
}

func resetTest() {
	RemoveAll(tmpDir)
	MkdirP(tmpDir)
}
