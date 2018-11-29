package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	{
		// empty []int
		assert.False(t, Q([]int{}).Any())

		// empty []interface{}
		assert.False(t, S().Any())
	}
	{
		// int
		assert.True(t, Q(1).Any())
	}
	{
		// string
		assert.True(t, Q("test").Any())
	}
	{
		// map
		assert.False(t, M().Any())
		assert.False(t, Q(map[int]interface{}{}).Any())
		assert.True(t, Q(map[int]interface{}{1: "one"}).Any())
	}
	{
		// empty []bob
		q := Q([]bob{})
		assert.False(t, q.Any())
	}
	{
		// []bob
		q := Q([]bob{{data: "3"}})
		assert.True(t, q.Any())
	}
	{
		assert.False(t, S().Any())
		assert.True(t, S().Append(1).Any())
		assert.False(t, Q([]int{}).Any())
		assert.True(t, Q([]int{1}).Any())
	}
}

func TestAnyWhere(t *testing.T) {
	{
		// string
		assert.True(t, Q("test").AnyWhere(func(x interface{}) bool {
			return x == "test"
		}))
	}
	{
		// int slice
		q := Q([]int{1, 2, 3})
		exists := q.AnyWhere(func(item interface{}) bool {
			return item.(int) == 5
		})
		assert.False(t, exists)
		exists = q.AnyWhere(func(item interface{}) bool {
			return item.(int) == 2
		})
		assert.True(t, exists)
	}
	{
		// empty map
		q := M()
		assert.False(t, q.AnyWhere(func(x interface{}) bool {
			return x == 3
		}))
	}
	{
		// str map
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.False(t, q.AnyWhere(func(x interface{}) bool { return x == 3 }))
		assert.True(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(KeyVal)).Key == "3"
		}))
		assert.True(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(KeyVal)).Val == "two"
		}))
	}
	{
		// []bob
		q := Q([]bob{{data: "3"}, {data: "4"}})
		assert.True(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(bob)).data == "3"
		}))
		assert.False(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(bob)).data == "5"
		}))
	}
	{
		q := S()
		assert.False(t, q.AnyWhere(func(x interface{}) bool {
			return x == 3
		}))
	}
	{
		q := Q([]string{"1", "2", "3"})
		assert.False(t, q.AnyWhere(func(x interface{}) bool { return x == 3 }))
		assert.True(t, q.AnyWhere(func(x interface{}) bool { return x == "3" }))
	}
}
