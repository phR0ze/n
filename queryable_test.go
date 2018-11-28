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
	for x, ok := next(); ok; x, ok = next() {
		fmt.Sprintln(x.(int) + 2)
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

func TestIQ(t *testing.T) {
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

func TestM(t *testing.T) {
	q := M()
	assert.NotNil(t, q)
	assert.NotNil(t, q.Iter)
	iter := q.Iter()
	assert.NotNil(t, iter)
	x, ok := iter()
	assert.Nil(t, x)
	assert.False(t, ok)
}

func TestMQ(t *testing.T) {
	{
		items := []interface{}{}
		q := Q(map[string]string{"1": "one", "2": "two", "3": "three"})
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			items = append(items, x)
			item := x.(*KeyVal)
			switch item.Key {
			case "1":
				assert.NotEqual(t, &KeyVal{Key: "2", Val: "one"}, item)
				assert.NotEqual(t, &KeyVal{Key: "1", Val: "two"}, item)
				assert.Equal(t, &KeyVal{Key: "1", Val: "one"}, item)
			case "2":
				assert.Equal(t, &KeyVal{Key: "2", Val: "two"}, item)
			case "3":
				assert.Equal(t, &KeyVal{Key: "3", Val: "three"}, item)
			}
		}
		assert.Len(t, items, 3)
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
		assert.Equal(t, 3, q.Append(1, 2, 3).Len())
	}
	{
		// Append many strings
		q := S()
		assert.Equal(t, 0, q.Len())
		assert.Equal(t, 3, q.Append("1", "2", "3").Len())
	}
}

func TestAt(t *testing.T) {
	q := Q([]int{1, 2, 3, 4})
	assert.Equal(t, 4, q.At(-1).Int())
	assert.Equal(t, 3, q.At(-2).Int())
	assert.Equal(t, 2, q.At(-3).Int())
	assert.Equal(t, 1, q.At(0).Int())
	assert.Equal(t, 2, q.At(1).Int())
	assert.Equal(t, 3, q.At(2).Int())
	assert.Equal(t, 4, q.At(3).Int())
}

func TestClear(t *testing.T) {
	q := Q([]int{1, 2, 3})
	assert.True(t, q.Any())
	assert.Equal(t, 3, q.Len())
	q.Clear()
	assert.False(t, q.Any())
	assert.Equal(t, 0, q.Len())
}

func TestJoin(t *testing.T) {
	{
		q := A()
		assert.Equal(t, "", q.Join(".").Str())
	}
	{
		q := S()
		assert.Equal(t, "", q.Join(".").Str())
	}
	{
		q := M()
		assert.Equal(t, "", q.Join(".").Str())
	}
	{
		q := Q("test")
		assert.Equal(t, "test", q.Join(".").Str())
	}
	{
		q := S().Append("1", "2", "3")
		assert.Equal(t, 3, q.Len())
		//joined := q.Join(".")
		//assert.Equal(t, 5, joined.Len())
		//assert.Equal(t, "1.2.3", q.Join(".").Str())
	}
}

func TestLen(t *testing.T) {
	{
		// Strings
		q := Q("test")
		assert.Equal(t, 4, q.Len())
	}
}

func TestMLen(t *testing.T) {
	{
		q := M()
		assert.Equal(t, 0, q.Len())
	}
	{
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.Equal(t, 3, q.Len())
	}
}

func TestSet(t *testing.T) {
	{
		cnt := []bool{}
		q := S()
		q.Set([]int{1, 2, 3})
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			cnt = append(cnt, true)
			switch len(cnt) {
			case 1:
				assert.Equal(t, 1, x)
			case 2:
				assert.Equal(t, 2, x)
			case 3:
				assert.Equal(t, 3, x)
			}
		}
		assert.Len(t, cnt, 3)
	}
}
