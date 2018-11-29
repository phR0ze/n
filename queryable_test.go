package n

import (
	"fmt"
	"testing"

	"github.com/ghodss/yaml"
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
	Q(ints).Each(func(item I) {
		fmt.Sprintln(item.(int) + 3)
	})
}

func TestQA(t *testing.T) {
	{
		q := Q("")
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

func TestQI(t *testing.T) {
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
		q := N()
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
			item := x.(KeyVal)
			switch item.Key {
			case "1":
				assert.NotEqual(t, KeyVal{Key: "2", Val: "one"}, item)
				assert.NotEqual(t, KeyVal{Key: "1", Val: "two"}, item)
				assert.Equal(t, KeyVal{Key: "1", Val: "one"}, item)
			case "2":
				assert.Equal(t, KeyVal{Key: "2", Val: "two"}, item)
			case "3":
				assert.Equal(t, KeyVal{Key: "3", Val: "three"}, item)
			}
		}
		assert.Len(t, items, 3)
	}
}
func TestQS(t *testing.T) {
	{
		q := N()
		assert.NotNil(t, q)
		assert.NotNil(t, q.Iter)
		iter := q.Iter()
		assert.NotNil(t, iter)
		x, ok := iter()
		assert.Nil(t, x)
		assert.False(t, ok)
	}
	{
		q := N()
		assert.False(t, q.Any())
		assert.Equal(t, 0, q.Len())
		q2 := q.Append(2)
		assert.Equal(t, 1, q.Len())
		assert.Equal(t, 1, q2.Len())
		assert.True(t, q2.Any())
		assert.Equal(t, q, q2)
		assert.Equal(t, 2, q.At(0).I())
	}
	{
		q := N()
		assert.NotNil(t, q)
		assert.NotNil(t, q.Iter)
		iter := q.Iter()
		assert.NotNil(t, iter)
		x, ok := iter()
		assert.Nil(t, x)
		assert.False(t, ok)
	}
	{
		cnt := []bool{}
		q := Q([]int{1, 2, 3})
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
		q := N()
		assert.False(t, q.Any())
		assert.Equal(t, 0, q.Len())
		q.Append(bob{data: "3"})
		assert.True(t, q.Any())
		assert.Equal(t, 1, q.Len())
		assert.Equal(t, bob{data: "3"}, q.At(0).O())
	}
}

func TestAny(t *testing.T) {
	{
		// empty []int
		assert.False(t, Q([]int{}).Any())

		// empty []interface{}
		assert.False(t, N().Any())
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
		assert.False(t, N().Any())
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
		assert.False(t, N().Any())
		assert.True(t, N().Append(1).Any())
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
		q := N()
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
		q := N()
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

func TestAppend(t *testing.T) {
	{
		// Append to valuetype
		q := Q(2)
		assert.Equal(t, 1, q.Len())
		assert.Equal(t, 2, q.Append(1).Len())
	}
	{
		// Append one
		q := N()
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
			q := N()
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
		q := Q("")
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

func TestContains(t *testing.T) {
	{
		// Empty slice
		q := N()
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
		q := N()
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

func TestEach(t *testing.T) {
	{
		// []int
		cnt := []bool{}
		q := Q([]int{1, 2, 3})
		q.Each(func(item I) {
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
		q.Each(func(item I) {
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
		q.Each(func(x I) {
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
		q.Each(func(x I) {
			items = append(items, x)
			item := x.(KeyVal)
			switch item.Key {
			case "1":
				assert.NotEqual(t, KeyVal{Key: "2", Val: "one"}, item)
				assert.NotEqual(t, KeyVal{Key: "1", Val: "two"}, item)
				assert.Equal(t, KeyVal{Key: "1", Val: "one"}, item)
			case "2":
				assert.Equal(t, KeyVal{Key: "2", Val: "two"}, item)
			case "3":
				assert.Equal(t, KeyVal{Key: "3", Val: "three"}, item)
			}
		})
		assert.Len(t, items, 3)
	}
}

func TestJoin(t *testing.T) {
	{
		q := Q("")
		assert.Equal(t, "", q.Join(".").A())
	}
	{
		q := N()
		assert.Equal(t, "", q.Join(".").A())
	}
	{
		q := N()
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
		q := N().Append("1", "2", "3")
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
		q := N().Append(1, 2, 3)
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
			q := N()
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
			q := N()
			assert.Equal(t, 0, q.Len())
		}
		{
			q := Q([]int{1, 2, 3})
			assert.Equal(t, 3, q.Len())
		}
		{
			q := Q([]string{"1", "2", "3"})
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

func TestMap(t *testing.T) {
	{
		// Get map keys
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		keys := q.Map(func(x interface{}) interface{} {
			return (x.(KeyVal)).Key
		})
		expected := Q([]string{"1", "2", "3"})
		assert.Equal(t, 3, keys.Len())
		assert.True(t, expected.Contains(keys.At(0).A()))
		assert.True(t, expected.Contains(keys.At(1).A()))
		assert.True(t, expected.Contains(keys.At(2).A()))
	}
	{
		// Get map values
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		vals := q.Map(func(x interface{}) interface{} {
			return (x.(KeyVal)).Val
		})
		expected := Q([]string{"one", "two", "three"})
		assert.Equal(t, 3, vals.Len())
		assert.True(t, expected.Contains(vals.At(0).A()))
		assert.True(t, expected.Contains(vals.At(1).A()))
		assert.True(t, expected.Contains(vals.At(2).A()))
	}
}

func TestSet(t *testing.T) {
	{
		// Strings
		{
			q := Q("")
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
			q := N()
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
		q := N()
		assert.False(t, q.Any())
		data := []bob{{data: "3"}}
		assert.True(t, q.Set(data).Any())
		assert.Equal(t, data[0], q.At(0).O())
	}
	{
		// []int
		cnt := []bool{}
		q := N()
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
		q := N()
		assert.Equal(t, []string{}, q.Split(".").S())
	}
	{
		q := Q([]int{1, 2, 3})
		assert.Equal(t, []string{}, q.Split(".").S())
	}
	{
		q := Q([]string{"1", "2", "3"})
		assert.Equal(t, []string{}, q.Split(".").S())
	}
	{
		q := Q(1)
		assert.Equal(t, []string{}, q.Split(".").S())
	}
	{
		q := N()
		assert.Equal(t, []string{}, q.Split(".").S())
	}
	{
		q := Q(map[string]interface{}{"1": "one"})
		assert.Equal(t, []string{}, q.Split(".").S())
	}
	{
		q := Q([]bob{{data: "2"}})
		assert.Equal(t, []string{}, q.Split(".").S())
	}
	{
		q := Q("1.2.3")
		assert.Equal(t, []string{"1", "2", "3"}, q.Split(".").S())
	}
	{
		q := Q("test1,test2")
		assert.Equal(t, []string{"test1", "test2"}, q.Split(",").S())
	}
}

func TestYAML(t *testing.T) {
	{
		// Get string from map
		rawYAML := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "two", q.YAML("1.2").A())
	}
	{
		// Get string from nested map
		rawYAML := `1:
  2:
    3: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "three", q.YAML("1.2.3").A())
	}
	{
		// Get map from map
		rawYAML := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := map[string]interface{}{"2": "two"}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.YAML("1").M())
	}
	{
		// Get map from map from map
		rawYAML := `1:
  2:
    3: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := map[string]interface{}{"3": "three"}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.YAML("1.2").M())
	}
	{
		// Get slice from map
		rawYAML := `foo:
  - 1
  - 2
  - 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, []string{"1", "2", "3"}, q.YAML("foo").Strs())
	}
	{
		// Select map from slice from map
		rawYAML := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := map[string]interface{}{"name": 2.0}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.YAML("foo.[name:2]").M())
	}
}
