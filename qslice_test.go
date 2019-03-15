package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQSlice_Slice(t *testing.T) {
	assert.Equal(t, []int{1}, Slice(1).O())
	assert.Equal(t, []interface{}{}, Slice(nil).O())
	assert.Equal(t, []string{"1"}, Slice("1").O())
	assert.Equal(t, []map[string]string{map[string]string{"1": "one"}}, Slice(map[string]string{"1": "one"}).O())
	assert.Equal(t, []string{"1", "2"}, Slice([]string{"1", "2"}).O())
}

func TestQSlice_Slicef(t *testing.T) {
	assert.Equal(t, []int{1}, Slicef(1).O())
	assert.Equal(t, []string{"1"}, Slicef("1").O())
	assert.Equal(t, []interface{}{}, Slice(nil).O())
	assert.Equal(t, []interface{}{}, Slicef(nil).O())
	assert.Equal(t, []interface{}{}, Slicef().O())
	assert.False(t, Slicef().Nil())
	assert.Equal(t, 0, Slicef().Len())
	assert.Equal(t, []string{"1", "2"}, Slicef("1", "2").O())
	assert.Equal(t, [][]string{{"1"}}, Slicef([]string{"1"}).O())
	assert.Equal(t, []map[string]string{map[string]string{"1": "one"}}, Slicef(map[string]string{"1": "one"}).O())
}

// func TestStrSliceAny(t *testing.T) {
// 	assert.False(t, S().Any())
// 	assert.True(t, S().Append("2").Any())
// }

func TestQSlice_Append(t *testing.T) {
	{
		// Append one
		slice := Slicef()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{"2"}, slice.O())
	}
	{
		// Append many
		slice := Slicef()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2", "4", "6")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []string{"2", "4", "6"}, slice.O())
	}
}

// // func TestStrSliceAt(t *testing.T) {
// 	{
// 		slice := S().Append("1", "2", "3", "4")
// 		assert.Equal(t, "4", slice.At(-1))
// 		assert.Equal(t, "3", slice.At(-2))
// 		assert.Equal(t, "2", slice.At(-3))
// 		assert.Equal(t, "1", slice.At(0))
// 		assert.Equal(t, "2", slice.At(1))
// 		assert.Equal(t, "3", slice.At(2))
// 		assert.Equal(t, "4", slice.At(3))
// 	}
// 	{
// 		slice := S().Append("1")
// 		assert.Equal(t, "1", slice.At(-1))
// 	}
// }

// func TestStrSliceClear(t *testing.T) {
// 	slice := S().Append("1", "2", "3", "4")
// 	assert.Equal(t, 4, slice.Len())
// 	slice.Clear()
// 	assert.Equal(t, 0, slice.Len())
// 	slice.Clear()
// 	assert.Equal(t, 0, slice.Len())
// }

// func TestStrSliceAnyContain(t *testing.T) {
// 	assert.True(t, S("one", "two", "three").AnyContain("thr"))
// 	assert.False(t, S("one", "two", "three").AnyContain("2"))
// }

// func TestStrSliceContains(t *testing.T) {
// 	assert.True(t, S("1", "2", "3").Contains("2"))
// 	assert.False(t, S("1", "2", "3").Contains("4"))
// }

// func TestStrSliceContainsAny(t *testing.T) {
// 	assert.True(t, S("1", "2", "3").ContainsAny([]string{"2"}))
// 	assert.False(t, S("1", "2", "3").ContainsAny([]string{"4"}))
// }

// func TestStrSliceDel(t *testing.T) {
// 	{
// 		// Pos: delete invalid
// 		slice := S("0", "1", "2")
// 		ok := slice.Del(3)
// 		assert.False(t, ok)
// 		assert.Equal(t, []string{"0", "1", "2"}, slice.S())
// 	}
// 	{
// 		// Pos: delete last
// 		slice := S("0", "1", "2")
// 		ok := slice.Del(2)
// 		assert.True(t, ok)
// 		assert.Equal(t, []string{"0", "1"}, slice.S())
// 	}
// 	{
// 		// Pos: delete middle
// 		slice := S("0", "1", "2")
// 		ok := slice.Del(1)
// 		assert.True(t, ok)
// 		assert.Equal(t, []string{"0", "2"}, slice.S())
// 	}
// 	{
// 		// delete first
// 		slice := S("0", "1", "2")
// 		ok := slice.Del(0)
// 		assert.True(t, ok)
// 		assert.Equal(t, []string{"1", "2"}, slice.S())
// 	}
// 	{
// 		// Neg: delete invalid
// 		slice := S("0", "1", "2")
// 		ok := slice.Del(-4)
// 		assert.False(t, ok)
// 		assert.Equal(t, []string{"0", "1", "2"}, slice.S())
// 	}
// 	{
// 		// Neg: delete last
// 		slice := S("0", "1", "2")
// 		ok := slice.Del(-1)
// 		assert.True(t, ok)
// 		assert.Equal(t, []string{"0", "1"}, slice.S())
// 	}
// 	{
// 		// Neg: delete middle
// 		slice := S("0", "1", "2")
// 		ok := slice.Del(-2)
// 		assert.True(t, ok)
// 		assert.Equal(t, []string{"0", "2"}, slice.S())
// 	}
// }

// func TestStrSliceDrop(t *testing.T) {
// 	{
// 		slice := S().Append("1", "2", "3").Drop(3)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// 	{
// 		slice := S().Append("1", "2", "3").Drop(5)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// 	{
// 		slice := S().Drop(3)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// 	{
// 		slice := S().Append("1", "2", "3").Drop(1)
// 		assert.Equal(t, []string{"2", "3"}, slice.S())
// 	}
// 	{
// 		slice := S().Append("1", "2", "3").Drop(2)
// 		assert.Equal(t, []string{"3"}, slice.S())
// 	}
// 	{
// 		slice := S().Append("1", "2", "3").Drop(0)
// 		assert.Equal(t, []string{"1", "2", "3"}, slice.S())
// 	}
// }

// func TestStrSliceEquals(t *testing.T) {
// 	{
// 		slice := S().Append("1", "2", "3")
// 		target := S().Append("1", "2", "3")
// 		assert.True(t, slice.Equals(target))
// 	}
// 	{
// 		slice := S().Append("1", "2", "4")
// 		target := S().Append("1", "2", "3")
// 		assert.False(t, slice.Equals(target))
// 	}
// 	{
// 		slice := S().Append("1", "2", "3", "4")
// 		target := S().Append("1", "2", "3")
// 		assert.False(t, slice.Equals(target))
// 	}
// }

// func TestStrSliceFirst(t *testing.T) {
// 	assert.Equal(t, A(""), S().First())
// 	assert.Equal(t, A("1"), S("1").First())
// 	assert.Equal(t, A("1"), S("1", "2").First())
// 	assert.Equal(t, "foo", A("foo::").Split("::").First().A())
// 	{
// 		// Test that the original slice wasn't modified
// 		q := S("1")
// 		assert.Equal(t, []string{"1"}, q.S())
// 		assert.Equal(t, A("1"), q.First())
// 		assert.Equal(t, []string{"1"}, q.S())
// 	}
// }

// func TestStrSliceJoin(t *testing.T) {
// 	assert.Equal(t, "", S().Join(".").A())
// 	assert.Equal(t, "1", S("1").Join(".").A())
// 	assert.Equal(t, "1.2", S("1", "2").Join(".").A())
// }

// func TestStrSliceLen(t *testing.T) {
// 	assert.Equal(t, 0, S().Len())
// 	assert.Equal(t, 1, S("1").Len())
// 	assert.Equal(t, 2, S("1", "2").Len())
// }

// func TestStrSliceLast(t *testing.T) {
// 	assert.Equal(t, A(""), S().Last())
// 	assert.Equal(t, A("1"), S("1").Last())
// 	assert.Equal(t, A("2"), S("1", "2").Last())
// 	assert.Equal(t, "foo", A("::foo").Split("::").Last().A())
// 	{
// 		// Test that the original slice wasn't modified
// 		q := S("1")
// 		assert.Equal(t, []string{"1"}, q.S())
// 		assert.Equal(t, A("1"), q.Last())
// 		assert.Equal(t, []string{"1"}, q.S())
// 	}
// }

// func TestStrSlicePrepend(t *testing.T) {
// 	slice := S().Prepend("1")
// 	assert.Equal(t, "1", slice.At(0))

// 	slice.Prepend("2", "3")
// 	assert.Equal(t, "2", slice.At(0))
// 	assert.Equal(t, []string{"2", "3", "1"}, slice.S())
// }

// func TestStrSliceSort(t *testing.T) {
// 	slice := S().Append("b", "d", "a")
// 	assert.Equal(t, []string{"a", "b", "d"}, slice.Sort().S())
// }

// func TestStrSliceSlice(t *testing.T) {
// 	assert.Equal(t, S(), S().Slice(0, -1))
// 	assert.Equal(t, S(""), S("").Slice(0, -1))
// 	assert.Equal(t, S("1", "2", "3"), S("1", "2", "3").Slice(0, -1))
// 	assert.Equal(t, S("1", "2"), S("1", "2", "3").Slice(0, -2))
// 	assert.Equal(t, S("1"), S("1", "2", "3").Slice(0, -3))
// 	assert.Equal(t, S(), S("1", "2", "3").Slice(0, -4))
// 	assert.Equal(t, S("2", "3"), S("1", "2", "3").Slice(1, -1))
// 	assert.Equal(t, S("3"), S("1", "2", "3").Slice(2, -1))
// 	assert.Equal(t, S(), S("1", "2", "3").Slice(3, -1))
// 	assert.Equal(t, S(), S("1", "2", "3").Slice(5, -1))
// 	assert.Equal(t, S("2", "3"), S("1", "2", "3").Slice(1, 2))
// 	assert.Equal(t, S(), S("1", "2", "3").Slice(3, 2))
// 	{
// 		// old FirstCnt ops
// 		assert.Equal(t, S(), S().Slice(0, 2))
// 		assert.Equal(t, S("1"), S("1").Slice(0, 2))
// 		assert.Equal(t, S("1", "2"), S("1", "2").Slice(0, 2))
// 		assert.Equal(t, S("1", "2", "3"), S("1", "2", "3").Slice(0, 2))
// 		assert.Equal(t, S("", "foo", "bar"), A("/foo/bar/one").Split("/").Slice(0, 2))
// 		assert.Equal(t, A("/foo/bar"), A("/foo/bar/one").Split("/").Slice(0, 2).Join("/"))
// 		{
// 			// Test that the original slice wasn't modified
// 			q := S("1")
// 			assert.Equal(t, []string{"1"}, q.S())
// 			assert.Equal(t, S("1"), q.Slice(0, 1))
// 			assert.Equal(t, []string{"1"}, q.S())
// 		}
// 	}
// 	{
// 		// old LastCnt(2) tests
// 		assert.Equal(t, S(), S().Slice(-3, -1))
// 		assert.Equal(t, S("1"), S("1").Slice(-2, -1))
// 		assert.Equal(t, S("1", "2"), S("1", "2").Slice(-2, -1))
// 		assert.Equal(t, S("2", "3"), S("1", "2", "3").Slice(-2, -1))
// 		assert.Equal(t, S("bar", "one"), A("/foo/bar/one").Split("/").Slice(-2, -1))
// 		assert.Equal(t, A("bar/one"), A("/foo/bar/one").Split("/").Slice(-2, -1).Join("/"))
// 		{
// 			// Test that the original slice wasn't modified
// 			q := S("1")
// 			assert.Equal(t, []string{"1"}, q.S())
// 			assert.Equal(t, S("1"), q.Slice(-2, -1))
// 			assert.Equal(t, []string{"1"}, q.S())
// 		}
// 	}
// }

// func TestStrSliceTakeFirst(t *testing.T) {
// 	{
// 		slice := S("0", "1", "2")
// 		results := []string{}
// 		expected := []string{"0", "1", "2"}
// 		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
// 			results = append(results, item)
// 		}
// 		assert.Equal(t, expected, results)
// 	}
// 	{
// 		slice := S("0", "1", "2")
// 		item, ok := slice.TakeFirst()
// 		assert.True(t, ok)
// 		assert.Equal(t, "0", item)
// 		assert.Equal(t, []string{"1", "2"}, slice.S())
// 	}
// 	{
// 		slice := S("0")
// 		item, ok := slice.TakeFirst()
// 		assert.True(t, ok)
// 		assert.Equal(t, "0", item)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// 	{
// 		slice := S()
// 		item, ok := slice.TakeFirst()
// 		assert.False(t, ok)
// 		assert.Equal(t, "", item)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// }

// func TestStrSliceTakeFirstCnt(t *testing.T) {
// 	{
// 		slice := S("0", "1", "2")
// 		items := slice.TakeFirstCnt(2).S()
// 		assert.Equal(t, []string{"0", "1"}, items)
// 		assert.Equal(t, []string{"2"}, slice.S())
// 	}
// 	{
// 		slice := S("0", "1", "2")
// 		items := slice.TakeFirstCnt(3).S()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// 	{
// 		slice := S("0", "1", "2")
// 		items := slice.TakeFirstCnt(4).S()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// }

// func TestStrSliceTakeLast(t *testing.T) {
// 	{
// 		slice := S("0", "1", "2")
// 		results := []string{}
// 		expected := []string{"2", "1", "0"}
// 		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
// 			results = append(results, item)
// 		}
// 		assert.Equal(t, expected, results)
// 	}
// 	{
// 		slice := S("0", "1", "2")
// 		item, ok := slice.TakeLast()
// 		assert.True(t, ok)
// 		assert.Equal(t, "2", item)
// 		assert.Equal(t, []string{"0", "1"}, slice.S())
// 	}
// 	{
// 		slice := S("0")
// 		item, ok := slice.TakeLast()
// 		assert.True(t, ok)
// 		assert.Equal(t, "0", item)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// 	{
// 		slice := S()
// 		item, ok := slice.TakeLast()
// 		assert.False(t, ok)
// 		assert.Equal(t, "", item)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// }
// func TestStrSliceTakeLastCnt(t *testing.T) {
// 	{
// 		slice := S("0", "1", "2")
// 		items := slice.TakeLastCnt(2).S()
// 		assert.Equal(t, []string{"1", "2"}, items)
// 		assert.Equal(t, []string{"0"}, slice.S())
// 	}
// 	{
// 		slice := S("0", "1", "2")
// 		items := slice.TakeLastCnt(3).S()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// 	{
// 		slice := S("0", "1", "2")
// 		items := slice.TakeLastCnt(4).S()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.S())
// 	}
// }

// func TestStrSliceUniq(t *testing.T) {
// 	{
// 		data := S().Uniq().S()
// 		expected := []string{}
// 		assert.Equal(t, expected, data)
// 	}
// 	{
// 		data := S("1", "2", "3").Uniq().S()
// 		expected := []string{"1", "2", "3"}
// 		assert.Equal(t, expected, data)
// 	}
// 	{
// 		data := S("1", "2", "2", "3").Uniq().S()
// 		expected := []string{"1", "2", "3"}
// 		assert.Equal(t, expected, data)
// 	}
// }

// func TestYamlPair(t *testing.T) {
// 	{
// 		k, v := A("foo=bar").Split("=").YamlPair()
// 		assert.Equal(t, "foo", k)
// 		assert.Equal(t, "bar", v)
// 	}
// 	{
// 		k, v := A("=bar").Split("=").YamlPair()
// 		assert.Equal(t, "", k)
// 		assert.Equal(t, "bar", v)
// 	}
// 	{
// 		k, v := A("bar=").Split("=").YamlPair()
// 		assert.Equal(t, "bar", k)
// 		assert.Equal(t, "", v)
// 	}
// 	{
// 		k, v := A("").Split("=").YamlPair()
// 		assert.Equal(t, "", k)
// 		assert.Equal(t, nil, v)
// 	}
// }
// func TestYamlKeyVal(t *testing.T) {
// 	{
// 		pair := A("foo=bar").Split("=").YamlKeyVal()
// 		assert.Equal(t, "foo", pair.Key)
// 		assert.Equal(t, "bar", pair.Val)
// 	}
// 	{
// 		pair := A("=bar").Split("=").YamlKeyVal()
// 		assert.Equal(t, "", pair.Key)
// 		assert.Equal(t, "bar", pair.Val)
// 	}
// 	{
// 		pair := A("bar=").Split("=").YamlKeyVal()
// 		assert.Equal(t, "bar", pair.Key)
// 		assert.Equal(t, "", pair.Val)
// 	}
// 	{
// 		pair := A("").Split("=").YamlKeyVal()
// 		assert.Equal(t, "", pair.Key)
// 		assert.Equal(t, "", pair.Val)
// 	}
// }
