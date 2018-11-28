package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	assert.Equal(t, 1, Q(1).Int())
}

func TestInts(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Q([]int{1, 2, 3}).Ints())
}

func TestStr(t *testing.T) {
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

func TestStrMap(t *testing.T) {
	{
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.Equal(t, 3, q.Len())
	}
}

func TestStrs(t *testing.T) {
	{
		q := Q([]string{"one"})
		assert.Equal(t, "one", q.At(0).Str())
		assert.Equal(t, []string{"one"}, q.Strs())
	}
	{
		assert.Equal(t, []string{"1", "2", "3"}, Q([]interface{}{"1", "2", "3"}).Strs())
	}
}
