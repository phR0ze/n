package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	{
		// Append to valuetype
		q := Q(2)
		assert.Equal(t, 1, q.Len())
		assert.Equal(t, 2, q.Append(1).Len())
	}
	{
		// Append one
		q := S()
		assert.Equal(t, 0, q.Len())
		assert.Equal(t, 1, q.Append(2).Len())
		assert.Equal(t, 2, q.Append(3).Len())
	}
	{
		// Append many ints
		q := Q([]int{1})
		assert.Equal(t, 1, q.Len())
		assert.Equal(t, 3, q.Append(2, 3).Len())
		assert.Equal(t, []int{1, 2, 3}, q.O())
	}
	{
		// Append many strings
		{
			q := S()
			assert.Equal(t, 0, q.Len())
			assert.Equal(t, 3, q.Append("1", "2", "3").Len())
		}
		{
			q := Q([]string{"1", "2"})
			assert.Equal(t, 2, q.Len())
			assert.Equal(t, 4, q.Append("3", "4").Len())
		}
	}
}
