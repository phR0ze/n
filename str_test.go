package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	assert.Equal(t, "test", Str("test").M())
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

func TestStrHasAnyPrefix(t *testing.T) {
	assert.True(t, Str("test").HasAnyPrefix([]string{"tes"}))
	assert.True(t, Str("test").HasAnyPrefix([]string{"bob", "tes"}))
	assert.False(t, Str("test").HasAnyPrefix([]string{"bob"}))
}

func TestStrHasAnySuffix(t *testing.T) {
	assert.True(t, Str("test").HasAnySuffix([]string{"est"}))
	assert.True(t, Str("test").HasAnySuffix([]string{"bob", "est"}))
	assert.False(t, Str("test").HasAnySuffix([]string{"bob"}))
}

func TestStrHasPrefix(t *testing.T) {
	assert.True(t, Str("test").HasPrefix("tes"))
}

func TestStrHasSuffix(t *testing.T) {
	assert.True(t, Str("test").HasSuffix("est"))
}

func TestStrSplit(t *testing.T) {
	assert.Equal(t, []string{"1", "2"}, Str("1.2").Split(".").M())
}
