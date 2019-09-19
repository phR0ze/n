package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmplEndTagNotFound(t *testing.T) {
	{
		// Invalid cast
		err := errors.New("foo bar")
		assert.False(t, TmplEndTagNotFoundError(err))
	}
	{
		// Invalid errs type
		err := NewTmplVarsNotFoundError()
		assert.False(t, TmplEndTagNotFoundError(err))
	}
	{
		// Valid test case
		err := NewTmplEndTagNotFoundError("foo", []byte("1"))
		assert.True(t, TmplEndTagNotFoundError(err))
	}
}

func TestTmplTagsInvalid(t *testing.T) {
	{
		// Invalid cast
		err := errors.New("foo bar")
		assert.False(t, TmplTagsInvalidError(err))
	}
	{
		// Invalid errs type
		err := NewTmplVarsNotFoundError()
		assert.False(t, TmplTagsInvalidError(err))
	}
	{
		// Valid test case
		err := NewTmplTagsInvalidError()
		assert.True(t, TmplTagsInvalidError(err))
	}
}

func TestTmplVarsNotFound(t *testing.T) {
	{
		// Invalid cast
		err := errors.New("foo bar")
		assert.False(t, TmplVarsNotFoundError(err))
	}
	{
		// Invalid errs type
		err := NewTmplTagsInvalidError()
		assert.False(t, TmplVarsNotFoundError(err))
	}
	{
		// Valid test case
		err := NewTmplVarsNotFoundError()
		assert.True(t, TmplVarsNotFoundError(err))
	}
}
