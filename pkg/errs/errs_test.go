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
		assert.False(t, TmplTagsInvalidError == NewTmplEndTagNotFoundError("foo", []byte("1")))
	}
	{
		// Valid test case
		err := NewTmplEndTagNotFoundError("foo", []byte("1"))
		assert.True(t, TmplEndTagNotFoundError(err))
	}
}

func TestTmplTagsInvalid(t *testing.T) {

	// Invalid cast
	{
		err := errors.New("foo bar")
		assert.False(t, err == TmplTagsInvalidError)
	}

	// Valid test case
	{
		assert.True(t, TmplTagsInvalidError == TmplTagsInvalidError)
	}
}

func TestTmplVarsNotFound(t *testing.T) {

	// Invalid cast
	{
		err := errors.New("foo bar")
		assert.False(t, err == TmplVarsNotFoundError)
	}

	// Valid test case
	{
		assert.True(t, TmplVarsNotFoundError == TmplVarsNotFoundError)
	}
}
