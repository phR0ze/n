package opt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultDebugOpt(t *testing.T) {
	opts := []*Opt{}
	assert.False(t, GetDebugOpt(opts))
	DefaultDebugOpt(&opts, true)
	assert.True(t, GetDebugOpt(opts))
}
