package nub

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const benchMarkSize = 9999999

func BenchmarkClosureIterator(t *testing.B) {
	ints := make([]int, benchMarkSize)
	for i := range ints {
		ints[i] = i
	}
	q := Q(ints)
	next := q.Iter()
	for item, ok := next(); ok; item, ok = next() {
		fmt.Sprintln(item.(int) + 2)
	}
}

func BenchmarkArrayIterator(t *testing.B) {
	ints := make([]int, benchMarkSize)
	for i := range ints {
		ints[i] = i
	}
	for _, item := range ints {
		fmt.Sprintln(item + 1)
	}
}

func BenchmarkEach(t *testing.B) {
	ints := make([]int, benchMarkSize)
	for i := range ints {
		ints[i] = i
	}
	Q(ints).Each(func(item interface{}) {
		fmt.Sprintln(item.(int) + 3)
	})
}

func TestN(t *testing.T) {
	q := S()
	assert.NotNil(t, q)
	assert.NotNil(t, q.Iter)
	iter := q.Iter()
	assert.NotNil(t, iter)
	item, ok := iter()
	assert.Nil(t, item)
	assert.False(t, ok)
}

func TestQ(t *testing.T) {
	{
		q := Q(nil)
		assert.NotNil(t, q)
		assert.NotNil(t, q.Iter)
		iter := q.Iter()
		assert.NotNil(t, iter)
		item, ok := iter()
		assert.Nil(t, item)
		assert.False(t, ok)
	}
	{
		cnt := []bool{}
		q := Q([]int{1, 2, 3})
		next := q.Iter()
		for item, ok := next(); ok; item, ok = next() {
			cnt = append(cnt, true)
			switch len(cnt) {
			case 1:
				assert.Equal(t, 1, item)
			case 2:
				assert.Equal(t, 2, item)
			case 3:
				assert.Equal(t, 3, item)
			}
		}
		assert.Len(t, cnt, 3)
	}
}

func TestSliceInput(t *testing.T) {
	q := S()
	assert.False(t, q.Any())
	assert.Equal(t, 0, q.Len())
	q2 := q.Append(2)
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, 1, q2.Len())
	assert.True(t, q2.Any())
	assert.Equal(t, q, q2)
	assert.Equal(t, 2, q.At(0).Int())
}

func TestIntInput(t *testing.T) {
	q := Q(5)
	assert.True(t, q.Any())
	assert.Equal(t, 1, q.Len())
	q2 := q.Append(2)
	assert.True(t, q2.Any())
	assert.Equal(t, q, q2)
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, 2, q2.Len())
	assert.Equal(t, 5, q.At(0).Int())
	assert.Equal(t, 2, q.At(1).Int())
}

func TestStrInput(t *testing.T) {
	q := Q("one")
	assert.True(t, q.Any())
	assert.Equal(t, 3, q.Len())
	assert.Equal(t, "o", q.At(0).Str())
	assert.Equal(t, 2, q.Append("four").Len())
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, "one", q.At(0).Str())
	assert.Equal(t, "four", q.At(1).Str())
}

func TestSet(t *testing.T) {
	{
		cnt := []bool{}
		q := S()
		q.Set([]int{1, 2, 3})
		next := q.Iter()
		for item, ok := next(); ok; item, ok = next() {
			cnt = append(cnt, true)
			switch len(cnt) {
			case 1:
				assert.Equal(t, 1, item)
			case 2:
				assert.Equal(t, 2, item)
			case 3:
				assert.Equal(t, 3, item)
			}
		}
		assert.Len(t, cnt, 3)
	}
}

func TestEach(t *testing.T) {
	{
		cnt := []bool{}
		q := Q([]int{1, 2, 3})
		q.Each(func(item interface{}) {
			cnt = append(cnt, true)
			switch len(cnt) {
			case 1:
				assert.Equal(t, 1, item)
			case 2:
				assert.Equal(t, 2, item)
			case 3:
				assert.Equal(t, 3, item)
			}
		})
		assert.Len(t, cnt, 3)

		// Check iterator again making sure it reset
		cnt = []bool{}
		q.Each(func(item interface{}) {
			cnt = append(cnt, true)
			switch len(cnt) {
			case 1:
				assert.Equal(t, 1, item)
			case 2:
				assert.Equal(t, 2, item)
			case 3:
				assert.Equal(t, 3, item)
			}
		})
	}
}

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
	}
	{
		// Append many ints
		q := S()
		assert.Equal(t, 0, q.Len())
		assert.Equal(t, 3, q.Append([]int{1, 2, 3}).Len())
	}
	{
		// Append many strings
		q := S()
		assert.Equal(t, 0, q.Len())
		assert.Equal(t, 3, q.Append([]string{"1", "2", "3"}).Len())
	}
}

func TestAt(t *testing.T) {
	q := Q([]int{1, 2, 3, 4})
	assert.Equal(t, 4, q.At(-1).Int())
	assert.Equal(t, 3, q.At(-2).Int())
	assert.Equal(t, 2, q.At(-3).Int())
	assert.Equal(t, 1, q.At(0).Int())
	assert.Equal(t, 2, q.At(1).Int())
	assert.Equal(t, 3, q.At(2).O.(int))
	assert.Equal(t, 4, q.At(3).Int())
}

func TestInt(t *testing.T) {
	assert.Equal(t, 1, Q(1).Int())
}

func TestInts(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Q([]int{1, 2, 3}).Ints())
}

func TestStr(t *testing.T) {
	assert.Equal(t, "1", Q("1").Str())
}

func TestStrs(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, Q([]interface{}{"1", "2", "3"}).Strs())
}
