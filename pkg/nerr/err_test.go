package nerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidTmplTag(t *testing.T) {
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
