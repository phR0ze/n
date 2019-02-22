package sys

import (
	"testing"

	"github.com/phR0ze/n/pkg/opt"
	"github.com/stretchr/testify/assert"
)

func TestDefaultFollowOpt(t *testing.T) {
	opts := []*opt.Opt{}
	assert.False(t, followOpt(opts))
	defaultFollowOpt(&opts, true)
	assert.True(t, followOpt(opts))
}
