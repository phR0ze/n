package nub

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const benchMarkSize = 9999999

type bob struct {
	data string
}

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

func TestQA(t *testing.T) {
	{
		q := A()
		assert.NotNil(t, q)
		assert.NotNil(t, q.Iter)
		iter := q.Iter()
		assert.NotNil(t, iter)
		x, ok := iter()
		assert.Nil(t, x)
		assert.False(t, ok)
	}
	{
		q := Q("one")
		assert.True(t, q.Any())
		assert.Equal(t, 3, q.Len())
		assert.Equal(t, "o", q.At(0).A())
		assert.Equal(t, 2, q.Append("four").Len())
		assert.Equal(t, 2, q.Len())
		assert.Equal(t, "one", q.At(0).A())
		assert.Equal(t, "four", q.At(1).A())
	}
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
	assert.Equal(t, 5, q.At(0).I())
	assert.Equal(t, 2, q.At(1).I())
}

func TestQM(t *testing.T) {
	{
		q := M()
		assert.NotNil(t, q)
		assert.NotNil(t, q.Iter)
		iter := q.Iter()
		assert.NotNil(t, iter)
		x, ok := iter()
		assert.Nil(t, x)
		assert.False(t, ok)
	}
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

func TestCustomQ(t *testing.T) {
	{
		// []bob
		q := Q([]bob{})
		assert.False(t, q.Any())
	}
	{
		// []bob
		q := Q([]bob{{data: "3"}})
		assert.True(t, q.Any())
		assert.Equal(t, bob{data: "3"}, q.At(0).O())
	}
	{
		// []bob
		q := S()
		assert.False(t, q.Any())
		assert.Equal(t, 0, q.Len())
		q.Append(bob{data: "3"})
		assert.True(t, q.Any())
		assert.Equal(t, 1, q.Len())
		assert.Equal(t, bob{data: "3"}, q.At(0).O())
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
	{
		// String
		q := Q("test")
		assert.Equal(t, "t", q.At(0).A())
		assert.Equal(t, "e", q.At(1).A())
		assert.Equal(t, "s", q.At(2).A())
		assert.Equal(t, "t", q.At(3).A())
		assert.Equal(t, "t", q.At(-1).A())
		assert.Equal(t, "s", q.At(-2).A())
		assert.Equal(t, "e", q.At(-3).A())
		assert.Equal(t, "t", q.At(-4).A())
	}
	{
		// []int
		q := Q([]int{1, 2, 3, 4})
		assert.Equal(t, 4, q.At(-1).I())
		assert.Equal(t, 3, q.At(-2).I())
		assert.Equal(t, 2, q.At(-3).I())
		assert.Equal(t, 1, q.At(0).I())
		assert.Equal(t, 2, q.At(1).I())
		assert.Equal(t, 3, q.At(2).I())
		assert.Equal(t, 4, q.At(3).I())
	}
}

func TestClear(t *testing.T) {
	{
		// Empty
		q := A()
		assert.False(t, q.Any())
		assert.Equal(t, 0, q.Clear().Len())
		assert.False(t, q.Any())
	}
	{
		// String
		q := Q("test")
		assert.True(t, q.Any())
		assert.Equal(t, "test", q.A())
		assert.Equal(t, 4, q.Len())
		assert.Equal(t, 0, q.Clear().Len())
		assert.False(t, q.Any())
	}
	{
		// []int
		q := Q([]int{1, 2, 3})
		assert.True(t, q.Any())
		assert.Equal(t, 3, q.Len())
		q.Clear()
		assert.False(t, q.Any())
		assert.Equal(t, 0, q.Len())
	}
	{
		// map[string]interface
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.True(t, q.Any())
		assert.Equal(t, 3, q.Len())
		q.Clear()
		assert.False(t, q.Any())
		assert.Equal(t, 0, q.Len())
	}
}

func TestEach(t *testing.T) {
	{
		// []int
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
	{
		// String
		q := Q("test")
		cnt := []bool{}
		q.Each(func(x interface{}) {
			cnt = append(cnt, true)
			item := string(x.(uint8))
			switch len(cnt) {
			case 1:
				assert.Equal(t, "t", item)
			case 2:
				assert.Equal(t, "e", item)
			case 3:
				assert.Equal(t, "s", item)
			case 4:
				assert.Equal(t, "t", item)
			}
		})
	}
	{
		// maps
		items := []interface{}{}
		q := Q(map[string]string{"1": "one", "2": "two", "3": "three"})
		q.Each(func(x interface{}) {
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
		})
		assert.Len(t, items, 3)
	}
}

func TestJoin(t *testing.T) {
	{
		q := A()
		assert.Equal(t, "", q.Join(".").A())
	}
	{
		q := S()
		assert.Equal(t, "", q.Join(".").A())
	}
	{
		q := M()
		assert.Equal(t, "", q.Join(".").A())
	}
	{
		q := Q("test")
		assert.Equal(t, "test", q.Join(".").A())
	}
	{
		q := Q(bob{data: "3"})
		assert.Equal(t, "", q.Join(".").A())
	}
	{
		q := S().Append("1", "2", "3")
		assert.Equal(t, 3, q.Len())
		joined := q.Join(".")
		assert.Equal(t, 5, joined.Len())
		assert.Equal(t, "1.2.3", q.Join(".").A())
	}
	{
		q := Q([]string{"1", "2", "3"})
		assert.Equal(t, 3, q.Len())
		joined := q.Join(".")
		assert.Equal(t, 5, joined.Len())
		assert.Equal(t, "1.2.3", q.Join(".").A())
	}
	{
		q := S().Append(1, 2, 3)
		assert.Equal(t, 3, q.Len())
		joined := q.Join(".")
		assert.Equal(t, 5, joined.Len())
		assert.Equal(t, "1.2.3", q.Join(".").A())
	}
	{
		q := Q([]int{1, 2, 3})
		assert.Equal(t, 3, q.Len())
		joined := q.Join(".")
		assert.Equal(t, 5, joined.Len())
		assert.Equal(t, "1.2.3", q.Join(".").A())
	}
}

func TestLen(t *testing.T) {
	{
		// Strings
		q := Q("test")
		assert.Equal(t, 4, q.Len())
	}
	{
		// Maps
		{
			q := M()
			assert.Equal(t, 0, q.Len())
		}
		{
			q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
			assert.Equal(t, 3, q.Len())
		}
	}
	{
		// Slices
		{
			q := S()
			assert.Equal(t, 0, q.Len())
		}
		{
			q := Q([]int{1, 2, 3})
			assert.Equal(t, 3, q.Len())
		}
	}
	{
		// Custom type
		{
			q := Q([]bob{{data: "3"}})
			assert.Equal(t, 1, q.Len())
		}
	}
}

func TestSet(t *testing.T) {
	{
		// Strings
		{
			q := A()
			assert.False(t, q.Any())
			assert.True(t, q.Set("test").Any())
			assert.Equal(t, "test", q.A())
		}
		{
			q := Q("test")
			assert.Equal(t, "test", q.A())
			assert.Equal(t, "foo", q.Set("foo").A())
			assert.Equal(t, "foo", q.A())
		}
	}
	{
		// Maps
		{
			q := M()
			assert.False(t, q.Any())
			data := map[string]interface{}{"1": "one"}
			assert.True(t, q.Set(data).Any())
			assert.Equal(t, data, q.M())
		}
		{
			data1 := map[string]interface{}{"1": "one"}
			data2 := map[string]interface{}{"1": "two"}
			q := Q(data1)
			assert.True(t, q.Any())
			assert.Equal(t, data1, q.M())
			assert.True(t, q.Set(data2).Any())
			assert.Equal(t, data2, q.M())
		}
	}
	{
		// custom type
		q := S()
		assert.False(t, q.Any())
		data := []bob{{data: "3"}}
		assert.True(t, q.Set(data).Any())
		assert.Equal(t, data[0], q.At(0).O())
	}
	{
		// []int
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

func TestSplit(t *testing.T) {
	{
		q := S()
		assert.Equal(t, []string{}, q.Split(".").Strs())
	}
	{
		q := Q([]int{1, 2, 3})
		assert.Equal(t, []string{}, q.Split(".").Strs())
	}
	{
		q := Q([]string{"1", "2", "3"})
		assert.Equal(t, []string{}, q.Split(".").Strs())
	}
	{
		q := Q(1)
		assert.Equal(t, []string{}, q.Split(".").Strs())
	}
	{
		q := M()
		assert.Equal(t, []string{}, q.Split(".").Strs())
	}
	{
		q := Q(map[string]interface{}{"1": "one"})
		assert.Equal(t, []string{}, q.Split(".").Strs())
	}
	{
		q := Q([]bob{{data: "2"}})
		assert.Equal(t, []string{}, q.Split(".").Strs())
	}
	{
		q := Q("1.2.3")
		assert.Equal(t, []string{"1", "2", "3"}, q.Split(".").Strs())
	}
	{
		q := Q("test1,test2")
		assert.Equal(t, []string{"test1", "test2"}, q.Split(",").Strs())
	}
}
