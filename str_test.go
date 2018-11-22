package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	assert.Equal(t, "test", Str("test").Raw)
}

func TestStrContains(t *testing.T) {
	assert.True(t, Str("test").Contains("tes"))
	assert.False(t, Str("test").Contains("bob"))
}

func TestStrContainsAny(t *testing.T) {
	assert.True(t, Str("test").ContainsAny([]string{"tes"}))
	assert.True(t, Str("test").ContainsAny([]string{"f", "t"}))
	assert.False(t, Str("test").ContainsAny([]string{"f", "b"}))
}

func TestStrSplit(t *testing.T) {
	assert.Equal(t, []string{"1", "2"}, Str("1.2").Split(".").Raw)
}
