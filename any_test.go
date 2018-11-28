package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {

	// Test empty queryable
	assert.False(t, S().Any())

	// Test empty collection object
	assert.False(t, Q([]int{}).Any())

	// Test value object
	assert.True(t, Q(1).Any())

	// Test string object
	assert.True(t, Q("2").Any())
}

func TestAAny(t *testing.T) {
	assert.False(t, A().Any())
	assert.True(t, Q("test").Any())
}

func TestMAny(t *testing.T) {
	assert.False(t, M().Any())
	assert.False(t, Q(map[int]interface{}{}).Any())
	assert.True(t, Q(map[int]interface{}{1: "one"}).Any())
}

func TestIAnyWhere(t *testing.T) {
	{
		q := Q([]int{1, 2, 3})
		exists := q.AnyWhere(func(item interface{}) bool {
			return item.(int) == 5
		})
		assert.False(t, exists)
	}
	{
		q := Q([]int{1, 2, 3})
		exists := q.AnyWhere(func(item interface{}) bool {
			return item.(int) == 2
		})
		assert.True(t, exists)
	}
}

func TestMAnyWhere(t *testing.T) {
	{
		q := M()
		assert.False(t, q.AnyWhere(func(x interface{}) bool {
			return x == 3
		}))
	}
	{
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.False(t, q.AnyWhere(func(x interface{}) bool { return x == 3 }))
		assert.True(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(*KeyVal)).Key == "3"
		}))
		assert.True(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(*KeyVal)).Val == "two"
		}))
	}
}
