package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	{
		// Empty slice
		q := S()
		assert.False(t, q.Contains(1))
	}
	{
		// []int
		q := Q([]int{})
		assert.False(t, q.Contains(1))
	}
	{
		// []int
		q := Q([]int{1, 2, 3})
		assert.False(t, q.Contains(0))
		assert.True(t, q.Contains(2))
	}
	{
		// int
		q := Q(2)
		assert.True(t, q.Contains(2))
	}
	{
		// empty []string
		q := Q([]string{})
		assert.False(t, q.Contains(""))
	}
	{
		// string
		q := Q("testing")
		assert.False(t, q.Contains("bob"))
		assert.True(t, q.Contains("test"))
	}
	{
		// full []string
		q := Q([]string{"1", "2", "3"})
		assert.False(t, q.Contains(""))
		assert.True(t, q.Contains("3"))
	}
	{
		// Custom type
		q := Q([]bob{{data: "3"}})
		assert.False(t, q.Contains(""))
		assert.False(t, q.Contains(bob{data: "2"}))
		assert.True(t, q.Contains(bob{data: "3"}))
	}
}

func TestContainsAny(t *testing.T) {
	{
		q := S()
		assert.False(t, q.ContainsAny(nil))
	}
	{
		q := S()
		assert.False(t, q.ContainsAny([]int{}))
	}
	{
		q := S()
		assert.False(t, q.ContainsAny([]string{}))
	}
	{
		q := Q([]int{1, 2, 3})
		assert.False(t, q.ContainsAny([]string{}))
		assert.True(t, q.ContainsAny(2))
		assert.True(t, q.ContainsAny([]int{0, 3}))
		assert.False(t, q.ContainsAny("2"))
	}
	{
		q := Q([]string{"1", "2", "3"})
		assert.False(t, q.ContainsAny([]int{}))
		assert.False(t, q.ContainsAny(2))
		assert.True(t, q.ContainsAny([]string{"0", "3"}))
		assert.True(t, q.ContainsAny("2"))
	}
}
