package sys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitCmd(t *testing.T) {
	{
		cmd := ` arg1 arg2 '   hello    world' "  another hello world   " `
		expected := []string{"arg1", "arg2", "'   hello    world'", `"  another hello world   "`}
		assert.Equal(t, expected, SplitCmd(cmd))
	}

	// Multiple types of quotes
	{
		cmd := `bash -c 'ls -la ${DIR}' | exec "$FOO"`
		expected := []string{"bash", "-c", "'ls -la ${DIR}'", "|", "exec", `"$FOO"`}
		assert.Equal(t, expected, SplitCmd(cmd))
	}

	// single value with spaces with quotes
	{
		cmd := " '  foo' "
		expected := []string{"'  foo'"}
		assert.Equal(t, expected, SplitCmd(cmd))
	}

	// single value with spaces
	{
		cmd := "   foo "
		expected := []string{"foo"}
		assert.Equal(t, expected, SplitCmd(cmd))
	}

	// single value
	{
		cmd := "foo"
		expected := []string{"foo"}
		assert.Equal(t, expected, SplitCmd(cmd))
	}
}
