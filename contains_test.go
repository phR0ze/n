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
		// map
		data := map[string]interface{}{"1": "one", "2": "two", "3": "three"}
		q := Q(data)
		assert.True(t, q.Contains("1"))
	}
	{
		// Custom type
		q := Q([]bob{{data: "3"}})
		assert.False(t, q.Contains(""))
		assert.False(t, q.Contains(bob{data: "2"}))
		assert.True(t, q.Contains(bob{data: "3"}))
	}
	{
		q := Q([]int{1, 2, 3})
		assert.False(t, q.Contains([]string{}))
		assert.True(t, q.Contains(2))
		assert.False(t, q.Contains([]int{0, 3}))
		assert.True(t, q.Contains([]int{1, 3}))
		assert.True(t, q.Contains([]int{2, 3}))
		assert.False(t, q.Contains([]int{4, 5}))
		assert.False(t, q.Contains("2"))
	}
	{
		q := Q([]string{"1", "2", "3"})
		assert.False(t, q.Contains([]int{}))
		assert.False(t, q.Contains(2))
		assert.False(t, q.Contains([]string{"0", "3"}))
		assert.True(t, q.Contains([]string{"1", "3"}))
		assert.True(t, q.Contains([]string{"2", "3"}))
		assert.True(t, q.Contains("2"))
	}
	{
		assert.True(t, Q("test").Contains("tes"))
		assert.False(t, Q("test").Contains([]string{"foo", "test"}))
		assert.True(t, Q("test").Contains([]string{"tes", "test"}))
		assert.True(t, Q([]string{"foo", "test"}).Contains("test"))
	}
	{
		// map
		data := map[string]interface{}{"1": "one", "2": "two", "3": "three"}
		q := Q(data)
		assert.True(t, q.Contains("1"))
		assert.False(t, q.Contains("4"))
		assert.False(t, q.Contains([]string{"4", "2"}))
		assert.True(t, q.Contains([]string{"3", "2"}))
	}
}

func TestContainsAny(t *testing.T) {
	{
		// Empty slice
		q := S()
		assert.False(t, q.ContainsAny(1))
	}
	{
		// []int
		q := Q([]int{})
		assert.False(t, q.ContainsAny(1))
	}
	{
		// []int
		q := Q([]int{1, 2, 3})
		assert.False(t, q.ContainsAny(0))
		assert.True(t, q.ContainsAny(2))
	}
	{
		// int
		q := Q(2)
		assert.True(t, q.ContainsAny(2))
	}
	{
		// empty []string
		q := Q([]string{})
		assert.False(t, q.ContainsAny(""))
	}
	{
		// string
		q := Q("testing")
		assert.False(t, q.ContainsAny("bob"))
		assert.True(t, q.ContainsAny("test"))
	}
	{
		// full []string
		q := Q([]string{"1", "2", "3"})
		assert.False(t, q.ContainsAny(""))
		assert.True(t, q.ContainsAny("3"))
	}
	{
		// map
		data := map[string]interface{}{"1": "one", "2": "two", "3": "three"}
		q := Q(data)
		assert.True(t, q.ContainsAny("1"))
	}
	{
		// Custom type
		q := Q([]bob{{data: "3"}})
		assert.False(t, q.ContainsAny(""))
		assert.False(t, q.ContainsAny(bob{data: "2"}))
		assert.True(t, q.ContainsAny(bob{data: "3"}))
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
	{
		assert.True(t, Q("test").ContainsAny("tes"))
		assert.True(t, Q("test").ContainsAny([]string{"foo", "test"}))
		assert.True(t, Q([]string{"foo", "test"}).ContainsAny("test"))
	}
	{
		// map
		data := map[string]interface{}{"1": "one", "2": "two", "3": "three"}
		q := Q(data)
		assert.True(t, q.ContainsAny("1"))
		assert.False(t, q.ContainsAny("4"))
		assert.True(t, q.ContainsAny([]string{"4", "2"}))
	}
}
