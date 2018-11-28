package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeFirst(t *testing.T) {
	{
		// empty
		x, ok := S().TakeFirst()
		assert.False(t, ok)
		assert.Nil(t, x)
	}
	{
		// non-iterable
		x, ok := Q(1).TakeFirst()
		assert.False(t, ok)
		assert.Nil(t, x)
	}
	{
		// ints
		q := Q([]int{1, 2, 3})
		x, ok := q.TakeFirst()
		assert.True(t, ok)
		assert.NotNil(t, x)
		assert.Equal(t, 1, x.(int))
		assert.Equal(t, []int{2, 3}, q.O())
	}
	{
		// strings
		q := S().Append("1", "2", "3")
		x, ok := q.TakeFirst()
		assert.True(t, ok)
		assert.NotNil(t, x)
		assert.Equal(t, "1", x.(string))
		assert.Equal(t, []string{"2", "3"}, q.Strs())
	}
	{
		// maps
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.True(t, q.Any())
		x, ok := q.TakeFirst()
		assert.False(t, ok)
		assert.Nil(t, x)
	}
	{
		// bobs
		q := S().Append(bob{data: "3"})
		assert.True(t, q.Any())
		x, ok := q.TakeFirst()
		assert.True(t, ok)
		assert.NotNil(t, x)
		assert.Equal(t, bob{data: "3"}, x.(bob))
		assert.False(t, q.Any())
	}
}
