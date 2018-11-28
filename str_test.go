package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	q := A()
	assert.NotNil(t, q)
	assert.NotNil(t, q.Iter)
	iter := q.Iter()
	assert.NotNil(t, iter)
	x, ok := iter()
	assert.Nil(t, x)
	assert.False(t, ok)
}

func TestAQ(t *testing.T) {
	q := Q("one")
	assert.True(t, q.Any())
	assert.Equal(t, 3, q.Len())
	assert.Equal(t, "o", q.At(0).Str())
	assert.Equal(t, 2, q.Append("four").Len())
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, "one", q.At(0).Str())
	assert.Equal(t, "four", q.At(1).Str())
}

func TestAStr(t *testing.T) {
	{
		q := Q("one")
		assert.Equal(t, "o", q.At(0).Str())
	}
	{
		q := Q([]string{"one"})
		assert.Equal(t, "one", q.At(0).Str())
		assert.Equal(t, []string{"one"}, q.Strs())
	}
	{
		assert.Equal(t, "1", Q("1").Str())
	}
}

func TestAStrs(t *testing.T) {
	{
		q := Q([]string{"one"})
		assert.Equal(t, "one", q.At(0).Str())
		assert.Equal(t, []string{"one"}, q.Strs())
	}
	{
		assert.Equal(t, []string{"1", "2", "3"}, Q([]interface{}{"1", "2", "3"}).Strs())
	}
}

func TestALen(t *testing.T) {
	{
		q := Q("test")
		assert.Equal(t, 4, q.Len())
	}
}

// TODO: Need to refactor below here
func TestStrStr(t *testing.T) {
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
