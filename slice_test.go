package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Need to refactor below here

//--------------------------------------------------------------------------------------------------
// IntSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewIntSlice(t *testing.T) {
	assert.NotNil(t, NewIntSlice().S())
}

func TestIntSlice(t *testing.T) {
	assert.NotNil(t, IntSlice(nil).S())
	assert.NotNil(t, IntSlice([]int{}).S())
}

func TestIntSliceAny(t *testing.T) {
	assert.False(t, NewIntSlice().Any())
	assert.True(t, NewIntSlice().Append(2).Any())
}

func TestIntSliceAppend(t *testing.T) {
	{
		// Append one
		slice := NewIntSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(2)
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []int{2}, slice.S())
	}
	{
		// Append many
		slice := NewIntSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(2, 4, 6)
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []int{2, 4, 6}, slice.S())
	}
}

func TestIntSliceAt(t *testing.T) {
	slice := NewIntSlice().Append(1, 2, 3, 4)
	assert.Equal(t, 4, slice.At(-1))
	assert.Equal(t, 3, slice.At(-2))
	assert.Equal(t, 2, slice.At(-3))
	assert.Equal(t, 1, slice.At(0))
	assert.Equal(t, 2, slice.At(1))
	assert.Equal(t, 3, slice.At(2))
	assert.Equal(t, 4, slice.At(3))
}

func TestIntSliceClear(t *testing.T) {
	slice := NewIntSlice().Append(1, 2, 3, 4)
	assert.Equal(t, 4, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
}

func TestIntSliceContains(t *testing.T) {
	assert.True(t, IntSlice([]int{1, 2, 3}).Contains(2))
	assert.False(t, IntSlice([]int{1, 2, 3}).Contains(4))
}

func TestIntSliceContainsAny(t *testing.T) {
	assert.True(t, IntSlice([]int{1, 2, 3}).ContainsAny([]int{2}))
	assert.False(t, IntSlice([]int{1, 2, 3}).ContainsAny([]int{4}))
}

func TestIntSliceDel(t *testing.T) {
	{
		// Pos: delete invalid
		slice := IntSlice([]int{0, 1, 2})
		ok := slice.Del(3)
		assert.False(t, ok)
		assert.Equal(t, []int{0, 1, 2}, slice.S())
	}
	{
		// Pos: delete last
		slice := IntSlice([]int{0, 1, 2})
		ok := slice.Del(2)
		assert.True(t, ok)
		assert.Equal(t, []int{0, 1}, slice.S())
	}
	{
		// Pos: delete middle
		slice := IntSlice([]int{0, 1, 2})
		ok := slice.Del(1)
		assert.True(t, ok)
		assert.Equal(t, []int{0, 2}, slice.S())
	}
	{
		// delete first
		slice := IntSlice([]int{0, 1, 2})
		ok := slice.Del(0)
		assert.True(t, ok)
		assert.Equal(t, []int{1, 2}, slice.S())
	}
	{
		// Neg: delete invalid
		slice := IntSlice([]int{0, 1, 2})
		ok := slice.Del(-4)
		assert.False(t, ok)
		assert.Equal(t, []int{0, 1, 2}, slice.S())
	}
	{
		// Neg: delete last
		slice := IntSlice([]int{0, 1, 2})
		ok := slice.Del(-1)
		assert.True(t, ok)
		assert.Equal(t, []int{0, 1}, slice.S())
	}
	{
		// Neg: delete middle
		slice := IntSlice([]int{0, 1, 2})
		ok := slice.Del(-2)
		assert.True(t, ok)
		assert.Equal(t, []int{0, 2}, slice.S())
	}
}

func TestIntSliceEquals(t *testing.T) {
	{
		slice := NewIntSlice().Append(1, 2, 3)
		target := NewIntSlice().Append(1, 2, 3)
		assert.True(t, slice.Equals(target))
	}
	{
		slice := NewIntSlice().Append(1, 2, 4)
		target := NewIntSlice().Append(1, 2, 3)
		assert.False(t, slice.Equals(target))
	}
	{
		slice := NewIntSlice().Append(1, 2, 3, 4)
		target := NewIntSlice().Append(1, 2, 3)
		assert.False(t, slice.Equals(target))
	}
}

func TestIntSliceLen(t *testing.T) {
	assert.Equal(t, 0, IntSlice([]int{}).Len())
	assert.Equal(t, 1, IntSlice([]int{1}).Len())
	assert.Equal(t, 2, IntSlice([]int{1, 2}).Len())
}

func TestIntSliceJoin(t *testing.T) {
	slice := NewIntSlice().Append(1, 2, 3)
	assert.Equal(t, "1.2.3", slice.Join(".").A())
}

func TestIntSlicePrepend(t *testing.T) {
	slice := NewIntSlice().Prepend(1)
	assert.Equal(t, 1, slice.At(0))

	slice.Prepend(2, 3)
	assert.Equal(t, 2, slice.At(0))
	assert.Equal(t, []int{2, 3, 1}, slice.S())
}

func TestIntSliceSort(t *testing.T) {
	slice := NewIntSlice().Append(3, 1, 2)
	assert.Equal(t, []int{1, 2, 3}, slice.Sort().S())
}

func TestIntSliceTakeFirst(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		results := []int{}
		expected := []int{0, 1, 2}
		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
			results = append(results, item)
		}
		assert.Equal(t, expected, results)
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{1, 2}, slice.S())
	}
	{
		slice := IntSlice([]int{0})
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.S())
	}
	{
		slice := IntSlice([]int{})
		item, ok := slice.TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.S())
	}
}

func TestIntSliceTakeFirstCnt(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(2).S()
		assert.Equal(t, []int{0, 1}, items)
		assert.Equal(t, []int{2}, slice.S())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(3).S()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.S())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(4).S()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.S())
	}
}

func TestIntSliceTakeLast(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		results := []int{}
		expected := []int{2, 1, 0}
		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
			results = append(results, item)
		}
		assert.Equal(t, expected, results)
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, 2, item)
		assert.Equal(t, []int{0, 1}, slice.S())
	}
	{
		slice := IntSlice([]int{0})
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.S())
	}
	{
		slice := IntSlice([]int{})
		item, ok := slice.TakeLast()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.S())
	}
}

func TestIntSliceTakeLastCnt(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(2).S()
		assert.Equal(t, []int{1, 2}, items)
		assert.Equal(t, []int{0}, slice.S())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(3).S()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.S())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(4).S()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.S())
	}
}

func TestIntSliceUniq(t *testing.T) {
	{
		data := IntSlice([]int{}).Uniq().S()
		expected := []int{}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 3}).Uniq().S()
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 2, 3}).Uniq().S()
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
}

//--------------------------------------------------------------------------------------------------
// StrSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewStrSlice(t *testing.T) {
	assert.NotNil(t, S().S())
}

func TestStrSlice(t *testing.T) {
	assert.NotNil(t, S().S())
	assert.NotNil(t, S([]string{}...).S())
}

func TestStrSliceAny(t *testing.T) {
	assert.False(t, S().Any())
	assert.True(t, S().Append("2").Any())
}

func TestStrSliceAppend(t *testing.T) {
	{
		// Append one
		slice := S()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{"2"}, slice.S())
	}
	{
		// Append many
		slice := S()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2", "4", "6")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []string{"2", "4", "6"}, slice.S())
	}
}
func TestStrSliceAt(t *testing.T) {
	slice := S().Append("1", "2", "3", "4")
	assert.Equal(t, "4", slice.At(-1))
	assert.Equal(t, "3", slice.At(-2))
	assert.Equal(t, "2", slice.At(-3))
	assert.Equal(t, "1", slice.At(0))
	assert.Equal(t, "2", slice.At(1))
	assert.Equal(t, "3", slice.At(2))
	assert.Equal(t, "4", slice.At(3))
}

func TestStrSliceClear(t *testing.T) {
	slice := S().Append("1", "2", "3", "4")
	assert.Equal(t, 4, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
}

func TestStrSliceAnyContain(t *testing.T) {
	assert.True(t, S("one", "two", "three").AnyContain("thr"))
	assert.False(t, S("one", "two", "three").AnyContain("2"))
}

func TestStrSliceContains(t *testing.T) {
	assert.True(t, S("1", "2", "3").Contains("2"))
	assert.False(t, S("1", "2", "3").Contains("4"))
}

func TestStrSliceContainsAny(t *testing.T) {
	assert.True(t, S("1", "2", "3").ContainsAny([]string{"2"}))
	assert.False(t, S("1", "2", "3").ContainsAny([]string{"4"}))
}

func TestStrSliceDel(t *testing.T) {
	{
		// Pos: delete invalid
		slice := S("0", "1", "2")
		ok := slice.Del(3)
		assert.False(t, ok)
		assert.Equal(t, []string{"0", "1", "2"}, slice.S())
	}
	{
		// Pos: delete last
		slice := S("0", "1", "2")
		ok := slice.Del(2)
		assert.True(t, ok)
		assert.Equal(t, []string{"0", "1"}, slice.S())
	}
	{
		// Pos: delete middle
		slice := S("0", "1", "2")
		ok := slice.Del(1)
		assert.True(t, ok)
		assert.Equal(t, []string{"0", "2"}, slice.S())
	}
	{
		// delete first
		slice := S("0", "1", "2")
		ok := slice.Del(0)
		assert.True(t, ok)
		assert.Equal(t, []string{"1", "2"}, slice.S())
	}
	{
		// Neg: delete invalid
		slice := S("0", "1", "2")
		ok := slice.Del(-4)
		assert.False(t, ok)
		assert.Equal(t, []string{"0", "1", "2"}, slice.S())
	}
	{
		// Neg: delete last
		slice := S("0", "1", "2")
		ok := slice.Del(-1)
		assert.True(t, ok)
		assert.Equal(t, []string{"0", "1"}, slice.S())
	}
	{
		// Neg: delete middle
		slice := S("0", "1", "2")
		ok := slice.Del(-2)
		assert.True(t, ok)
		assert.Equal(t, []string{"0", "2"}, slice.S())
	}
}

func TestStrSliceEquals(t *testing.T) {
	{
		slice := S().Append("1", "2", "3")
		target := S().Append("1", "2", "3")
		assert.True(t, slice.Equals(target))
	}
	{
		slice := S().Append("1", "2", "4")
		target := S().Append("1", "2", "3")
		assert.False(t, slice.Equals(target))
	}
	{
		slice := S().Append("1", "2", "3", "4")
		target := S().Append("1", "2", "3")
		assert.False(t, slice.Equals(target))
	}
}

func TestStrSliceJoin(t *testing.T) {
	assert.Equal(t, "", S().Join(".").A())
	assert.Equal(t, "1", S("1").Join(".").A())
	assert.Equal(t, "1.2", S("1", "2").Join(".").A())
}

func TestStrSliceLen(t *testing.T) {
	assert.Equal(t, 0, S().Len())
	assert.Equal(t, 1, S("1").Len())
	assert.Equal(t, 2, S("1", "2").Len())
}

func TestStrSliceSort(t *testing.T) {
	slice := S().Append("b", "d", "a")
	assert.Equal(t, []string{"a", "b", "d"}, slice.Sort().S())
}

func TestStrSlicePrepend(t *testing.T) {
	slice := S().Prepend("1")
	assert.Equal(t, "1", slice.At(0))

	slice.Prepend("2", "3")
	assert.Equal(t, "2", slice.At(0))
	assert.Equal(t, []string{"2", "3", "1"}, slice.S())
}

func TestStrSliceTakeFirst(t *testing.T) {
	{
		slice := S("0", "1", "2")
		results := []string{}
		expected := []string{"0", "1", "2"}
		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
			results = append(results, item)
		}
		assert.Equal(t, expected, results)
	}
	{
		slice := S("0", "1", "2")
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{"1", "2"}, slice.S())
	}
	{
		slice := S("0")
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, slice.S())
	}
	{
		slice := S()
		item, ok := slice.TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, slice.S())
	}
}

func TestStrSliceTakeFirstCnt(t *testing.T) {
	{
		slice := S("0", "1", "2")
		items := slice.TakeFirstCnt(2).S()
		assert.Equal(t, []string{"0", "1"}, items)
		assert.Equal(t, []string{"2"}, slice.S())
	}
	{
		slice := S("0", "1", "2")
		items := slice.TakeFirstCnt(3).S()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.S())
	}
	{
		slice := S("0", "1", "2")
		items := slice.TakeFirstCnt(4).S()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.S())
	}
}

func TestStrSliceTakeLast(t *testing.T) {
	{
		slice := S("0", "1", "2")
		results := []string{}
		expected := []string{"2", "1", "0"}
		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
			results = append(results, item)
		}
		assert.Equal(t, expected, results)
	}
	{
		slice := S("0", "1", "2")
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, "2", item)
		assert.Equal(t, []string{"0", "1"}, slice.S())
	}
	{
		slice := S("0")
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, slice.S())
	}
	{
		slice := S()
		item, ok := slice.TakeLast()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, slice.S())
	}
}
func TestStrSliceTakeLastCnt(t *testing.T) {
	{
		slice := S("0", "1", "2")
		items := slice.TakeLastCnt(2).S()
		assert.Equal(t, []string{"1", "2"}, items)
		assert.Equal(t, []string{"0"}, slice.S())
	}
	{
		slice := S("0", "1", "2")
		items := slice.TakeLastCnt(3).S()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.S())
	}
	{
		slice := S("0", "1", "2")
		items := slice.TakeLastCnt(4).S()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.S())
	}
}

func TestStrSliceUniq(t *testing.T) {
	{
		data := S().Uniq().S()
		expected := []string{}
		assert.Equal(t, expected, data)
	}
	{
		data := S("1", "2", "3").Uniq().S()
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
	{
		data := S("1", "2", "2", "3").Uniq().S()
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
}

//--------------------------------------------------------------------------------------------------
// StrMapSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewStrMapSlice(t *testing.T) {
	assert.NotNil(t, NewStrMapSlice().S())
}

func TestStrMapSlice(t *testing.T) {
	assert.NotNil(t, StrMapSlice(nil).S())
	assert.NotNil(t, StrMapSlice([]map[string]interface{}{}).S())
}
func TestStrMapSliceAny(t *testing.T) {
	assert.False(t, NewStrMapSlice().Any())
	assert.True(t, NewStrMapSlice().Append(map[string]interface{}{"1": "one"}).Any())
}

func TestStrMapSliceAppend(t *testing.T) {
	{
		// Append one
		slice := NewStrMapSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(map[string]interface{}{"2": "two"})
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []map[string]interface{}{{"2": "two"}}, slice.S())
	}
	{
		// Append many
		slice := NewStrMapSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(map[string]interface{}{"1": "one"}, map[string]interface{}{"2": "two"})
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []map[string]interface{}{{"1": "one"}, {"2": "two"}}, slice.S())
	}
}

func TestStrMapSliceAt(t *testing.T) {
	slice := NewStrMapSlice()
	assert.Equal(t, 0, slice.Len())
	slice.Append(map[string]interface{}{"1": "one"})
	slice.Append(map[string]interface{}{"2": "two"})
	slice.Append(map[string]interface{}{"3": "three"})

	assert.Equal(t, map[string]interface{}{"3": "three"}, slice.At(-1).M())
	assert.Equal(t, map[string]interface{}{"2": "two"}, slice.At(-2).M())
	assert.Equal(t, map[string]interface{}{"1": "one"}, slice.At(0).M())
	assert.Equal(t, map[string]interface{}{"2": "two"}, slice.At(1).M())
	assert.Equal(t, map[string]interface{}{"3": "three"}, slice.At(2).M())
}

func TestStrMapSliceClear(t *testing.T) {
	slice := NewStrMapSlice()
	slice.Append(map[string]interface{}{"1": "one"})
	slice.Append(map[string]interface{}{"2": "two"})
	slice.Append(map[string]interface{}{"3": "three"})

	assert.Equal(t, 3, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
}

func TestStrMapSliceContains(t *testing.T) {
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.True(t, StrMapSlice(raw).Contains("1"))
	}
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.False(t, StrMapSlice(raw).Contains("4"))
	}
}

func TestStrMapSliceContainsAny(t *testing.T) {
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.True(t, StrMapSlice(raw).ContainsAny([]string{"4", "1"}))
	}
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.False(t, StrMapSlice(raw).ContainsAny([]string{}))
	}
}

func TestStrMapSliceDel(t *testing.T) {
	{
		// Pos: delete invalid
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		slice.Append(map[string]interface{}{"3": "three"})
		ok := slice.Del(3)
		assert.False(t, ok)
		expected := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.Equal(t, expected, slice.S())
	}
	{
		// Pos: delete last
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		slice.Append(map[string]interface{}{"3": "three"})
		ok := slice.Del(2)
		assert.True(t, ok)
		expected := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, slice.S())
	}
	{
		// Pos: delete middle
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		slice.Append(map[string]interface{}{"3": "three"})
		ok := slice.Del(1)
		assert.True(t, ok)
		expected := []map[string]interface{}{
			{"1": "one"},
			{"3": "three"},
		}
		assert.Equal(t, expected, slice.S())
	}
	{
		// delete first
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		slice.Append(map[string]interface{}{"3": "three"})
		ok := slice.Del(0)
		assert.True(t, ok)
		expected := []map[string]interface{}{
			{"2": "two"},
			{"3": "three"},
		}
		assert.Equal(t, expected, slice.S())
	}
	{
		// Neg: delete invalid
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		slice.Append(map[string]interface{}{"3": "three"})
		ok := slice.Del(-4)
		assert.False(t, ok)
		expected := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.Equal(t, expected, slice.S())
	}
	{
		// Neg: delete last
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		slice.Append(map[string]interface{}{"3": "three"})
		ok := slice.Del(-1)
		assert.True(t, ok)
		expected := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, slice.S())
	}
	{
		// Neg: delete middle
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		slice.Append(map[string]interface{}{"3": "three"})
		ok := slice.Del(-2)
		assert.True(t, ok)
		expected := []map[string]interface{}{
			{"1": "one"},
			{"3": "three"},
		}
		assert.Equal(t, expected, slice.S())
	}
}

func TestStrMapSliceEquals(t *testing.T) {
	{
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		other := NewStrMapSlice()
		other.Append(map[string]interface{}{"1": "one"})
		other.Append(map[string]interface{}{"2": "two"})
		assert.True(t, slice.Equals(other))
	}
	{
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		other := NewStrMapSlice()
		other.Append(map[string]interface{}{"1": "one"})
		other.Append(map[string]interface{}{"2": "three"})
		assert.False(t, slice.Equals(other))
	}
	{
		slice := NewStrMapSlice()
		slice.Append(map[string]interface{}{"1": "one"})
		slice.Append(map[string]interface{}{"2": "two"})
		other := NewStrMapSlice()
		other.Append(map[string]interface{}{"1": "one"})
		other.Append(map[string]interface{}{"2": "two"})
		other.Append(map[string]interface{}{"3": "three"})
		assert.False(t, slice.Equals(other))
	}
}

func TestStrMapSliceLen(t *testing.T) {
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.Equal(t, 3, StrMapSlice(raw).Len())
	}
	{
		raw := []map[string]interface{}{
			{"3": "three"},
		}
		assert.Equal(t, 1, StrMapSlice(raw).Len())
	}
	{
		raw := []map[string]interface{}{}
		assert.Equal(t, 0, StrMapSlice(raw).Len())
	}
	{
		assert.Equal(t, 0, StrMapSlice(nil).Len())
	}
}

func TestStrMapSlicePrepend(t *testing.T) {
	slice := NewStrMapSlice()
	slice.Prepend(map[string]interface{}{"1": "one"})
	assert.Equal(t, map[string]interface{}{"1": "one"}, slice.At(0).M())

	slice.Prepend(map[string]interface{}{"2": "two"}, map[string]interface{}{"3": "three"})
	assert.Equal(t, map[string]interface{}{"2": "two"}, slice.At(0).M())

	expected := []map[string]interface{}{
		{"2": "two"},
		{"3": "three"},
		{"1": "one"},
	}
	assert.Equal(t, expected, slice.S())
}

func TestStrMapSliceTakeFirst(t *testing.T) {
	{
		// Take interator
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		results := []map[string]interface{}{}
		expected := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
			results = append(results, item.M())
		}
		assert.Equal(t, expected, results)
	}
	{
		// TakeFirst one
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		expectedRemainder := []map[string]interface{}{
			{"2": "two"},
			{"3": "three"},
		}
		expectedResult := map[string]interface{}{"1": "one"}
		result, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, expectedResult, result.M())
		assert.Equal(t, expectedRemainder, slice.S())
	}
	{
		// TakeFirst when empty
		slice := NewStrMapSlice()
		result, ok := slice.TakeFirst()
		assert.False(t, ok)
		assert.Nil(t, result)
	}
}

func TestStrMapSliceTakeFirstCnt(t *testing.T) {
	{
		// TakeFirst two
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		expectedRemainder := []map[string]interface{}{
			{"3": "three"},
		}
		expectedResult := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
		}
		result := slice.TakeFirstCnt(2)
		assert.Equal(t, expectedResult, result.S())
		assert.Equal(t, expectedRemainder, slice.S())
	}
	{
		// TakeFirst four
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		expectedRemainder := []map[string]interface{}{}
		expectedResult := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		result := slice.TakeFirstCnt(4)
		assert.Equal(t, expectedResult, result.S())
		assert.Equal(t, expectedRemainder, slice.S())
	}
	{
		// TakeFirst when empty
		slice := NewStrMapSlice()
		result := slice.TakeFirstCnt(3)
		assert.Equal(t, 0, result.Len())
	}
}

func TestStrMapSliceTakeLast(t *testing.T) {
	{
		// Take interator
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		results := []map[string]interface{}{}
		expected := []map[string]interface{}{
			{"3": "three"},
			{"2": "two"},
			{"1": "one"},
		}
		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
			results = append(results, item.M())
		}
		assert.Equal(t, expected, results)
	}
	{
		// TakeLast one
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		expectedRemainder := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
		}
		expectedResult := map[string]interface{}{"3": "three"}
		result, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, expectedResult, result.M())
		assert.Equal(t, expectedRemainder, slice.S())
	}
	{
		// TakeFirst when empty
		slice := NewStrMapSlice()
		result, ok := slice.TakeLast()
		assert.False(t, ok)
		assert.Nil(t, result)
	}
}

func TestStrMapSliceTakeLastCnt(t *testing.T) {
	{
		// TakeLast two
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		expectedRemainder := []map[string]interface{}{
			{"1": "one"},
		}
		expectedResult := []map[string]interface{}{
			{"2": "two"},
			{"3": "three"},
		}
		result := slice.TakeLastCnt(2)
		assert.Equal(t, expectedResult, result.S())
		assert.Equal(t, expectedRemainder, slice.S())
	}
	{
		// TakeLast four
		slice := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		})
		expectedRemainder := []map[string]interface{}{}
		expectedResult := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		result := slice.TakeLastCnt(4)
		assert.Equal(t, expectedResult, result.S())
		assert.Equal(t, expectedRemainder, slice.S())
	}
	{
		// TakeFirst when empty
		slice := NewStrMapSlice()
		result := slice.TakeLastCnt(3)
		assert.Equal(t, 0, result.Len())
	}
}

func TestYAMLPair(t *testing.T) {
	{
		k, v := A("foo=bar").Split("=").YAMLPair()
		assert.Equal(t, "foo", k)
		assert.Equal(t, "bar", v)
	}
	{
		k, v := A("=bar").Split("=").YAMLPair()
		assert.Equal(t, "", k)
		assert.Equal(t, "bar", v)
	}
	{
		k, v := A("bar=").Split("=").YAMLPair()
		assert.Equal(t, "bar", k)
		assert.Equal(t, "", v)
	}
	{
		k, v := A("").Split("=").YAMLPair()
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
	}
}
func TestYAMLKeyVal(t *testing.T) {
	{
		pair := A("foo=bar").Split("=").YAMLKeyVal()
		assert.Equal(t, "foo", pair.Key)
		assert.Equal(t, "bar", pair.Val)
	}
	{
		pair := A("=bar").Split("=").YAMLKeyVal()
		assert.Equal(t, "", pair.Key)
		assert.Equal(t, "bar", pair.Val)
	}
	{
		pair := A("bar=").Split("=").YAMLKeyVal()
		assert.Equal(t, "bar", pair.Key)
		assert.Equal(t, "", pair.Val)
	}
	{
		pair := A("").Split("=").YAMLKeyVal()
		assert.Equal(t, "", pair.Key)
		assert.Equal(t, "", pair.Val)
	}
}
