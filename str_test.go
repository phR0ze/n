package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrA(t *testing.T) {
	assert.Equal(t, "test", A("test").A())
}

func TestStrQ(t *testing.T) {
	assert.Equal(t, "test", A("test").Q().A())
}

func TestStrContains(t *testing.T) {
	assert.True(t, A("test").Contains("tes"))
	assert.False(t, A("test").Contains("bob"))
}

func TestStrContainsAny(t *testing.T) {
	assert.True(t, A("test").ContainsAny("tes"))
	assert.True(t, A("test").ContainsAny("f", "t"))
	assert.False(t, A("test").ContainsAny("f", "b"))
}

func TestStrHasAnyPrefix(t *testing.T) {
	assert.True(t, A("test").HasAnyPrefix("tes"))
	assert.True(t, A("test").HasAnyPrefix("bob", "tes"))
	assert.False(t, A("test").HasAnyPrefix("bob"))
}

func TestStrHasAnySuffix(t *testing.T) {
	assert.True(t, A("test").HasAnySuffix("est"))
	assert.True(t, A("test").HasAnySuffix("bob", "est"))
	assert.False(t, A("test").HasAnySuffix("bob"))
}

func TestStrHasPrefix(t *testing.T) {
	assert.True(t, A("test").HasPrefix("tes"))
}

func TestStrHasSuffix(t *testing.T) {
	assert.True(t, A("test").HasSuffix("est"))
}

func TestStrSplit(t *testing.T) {
	assert.Equal(t, []string{"1", "2"}, A("1.2").Split(".").S())
}
func TestStrTrimPrefix(t *testing.T) {
	assert.Equal(t, "test]", A("[test]").TrimPrefix("[").A())
}

func TestStrTrimSuffix(t *testing.T) {
	assert.Equal(t, "[test", A("[test]").TrimSuffix("]").A())
}

func TestYAMLType(t *testing.T) {
	{
		// string
		assert.Equal(t, "test", A("\"test\"").YAMLType())
		assert.Equal(t, "test", A("'test'").YAMLType())
		assert.Equal(t, "1", A("\"1\"").YAMLType())
		assert.Equal(t, "1", A("'1'").YAMLType())
	}
	{
		// int
		assert.Equal(t, 1, A("1").YAMLType())
		assert.Equal(t, 0, A("0").YAMLType())
		assert.Equal(t, 25, A("25").YAMLType())
	}
	{
		// bool
		assert.Equal(t, true, A("true").YAMLType())
		assert.Equal(t, false, A("false").YAMLType())
	}
	{
		// default
		assert.Equal(t, "True", A("True").YAMLType())
	}
}
