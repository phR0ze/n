package nerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmplEndTagNotFound(t *testing.T) {
	{
		// Invalid cast
		err := errors.New("foo bar")
		assert.False(t, IsTmplEndTagNotFound(err))
	}
	{
		// Invalid nerr type
		err := NewTmplVarsNotFound()
		assert.False(t, IsTmplEndTagNotFound(err))
	}
	{
		// Valid test case
		err := NewTmplEndTagNotFound("foo", []byte("1"))
		assert.True(t, IsTmplEndTagNotFound(err))
	}
}

func TestTmplTagsInvalid(t *testing.T) {
	{
		// Invalid cast
		err := errors.New("foo bar")
		assert.False(t, IsTmplTagsInvalid(err))
	}
	{
		// Invalid nerr type
		err := NewTmplVarsNotFound()
		assert.False(t, IsTmplTagsInvalid(err))
	}
	{
		// Valid test case
		err := NewTmplTagsInvalid()
		assert.True(t, IsTmplTagsInvalid(err))
	}
}

func TestTmplVarsNotFound(t *testing.T) {
	{
		// Invalid cast
		err := errors.New("foo bar")
		assert.False(t, IsTmplVarsNotFound(err))
	}
	{
		// Invalid nerr type
		err := NewTmplTagsInvalid()
		assert.False(t, IsTmplVarsNotFound(err))
	}
	{
		// Valid test case
		err := NewTmplVarsNotFound()
		assert.True(t, IsTmplVarsNotFound(err))
	}
}
