package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	assert.Equal(t, "test", Str("test").Raw)
}
