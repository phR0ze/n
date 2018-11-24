package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------------------------------------------------------------------------------
// IntSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewIntSlice(t *testing.T) {
	assert.NotNil(t, NewIntSlice().Ex())
}

func TestIntSlice(t *testing.T) {
	assert.NotNil(t, IntSlice(nil).Ex())
	assert.NotNil(t, IntSlice([]int{}).Ex())
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
		assert.Equal(t, []int{2}, slice.Ex())
	}
	{
		// Append many
		slice := NewIntSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(2, 4, 6)
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []int{2, 4, 6}, slice.Ex())
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

func TestIntSliceLen(t *testing.T) {
	assert.Equal(t, 0, IntSlice([]int{}).Len())
	assert.Equal(t, 1, IntSlice([]int{1}).Len())
	assert.Equal(t, 2, IntSlice([]int{1, 2}).Len())
}

func TestIntSliceJoin(t *testing.T) {
	slice := NewIntSlice().Append(1, 2, 3)
	assert.Equal(t, "1.2.3", slice.Join(".").Ex())
}

func TestIntSlicePrepend(t *testing.T) {
	slice := NewIntSlice().Prepend(1)
	assert.Equal(t, 1, slice.At(0))

	slice.Prepend(2, 3)
	assert.Equal(t, 2, slice.At(0))
	assert.Equal(t, []int{2, 3, 1}, slice.Ex())
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
		assert.Equal(t, []int{1, 2}, slice.Ex())
	}
	{
		slice := IntSlice([]int{0})
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Ex())
	}
	{
		slice := IntSlice([]int{})
		item, ok := slice.TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Ex())
	}
}

func TestIntSliceTakeFirstCnt(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(2).Ex()
		assert.Equal(t, []int{0, 1}, items)
		assert.Equal(t, []int{2}, slice.Ex())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(3).Ex()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Ex())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(4).Ex()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Ex())
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
		assert.Equal(t, []int{0, 1}, slice.Ex())
	}
	{
		slice := IntSlice([]int{0})
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Ex())
	}
	{
		slice := IntSlice([]int{})
		item, ok := slice.TakeLast()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Ex())
	}
}

func TestIntSliceTakeLastCnt(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(2).Ex()
		assert.Equal(t, []int{1, 2}, items)
		assert.Equal(t, []int{0}, slice.Ex())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(3).Ex()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Ex())
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(4).Ex()
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Ex())
	}
}

func TestIntSliceUniq(t *testing.T) {
	{
		data := IntSlice([]int{}).Uniq().Ex()
		expected := []int{}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 3}).Uniq().Ex()
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 2, 3}).Uniq().Ex()
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
}

//--------------------------------------------------------------------------------------------------
// StrSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewStrSlice(t *testing.T) {
	assert.NotNil(t, NewStrSlice().Ex())
}

func TestStrSlice(t *testing.T) {
	assert.NotNil(t, StrSlice(nil).Ex())
	assert.NotNil(t, StrSlice([]string{}).Ex())
}

func TestStrSliceAny(t *testing.T) {
	assert.False(t, NewStrSlice().Any())
	assert.True(t, NewStrSlice().Append("2").Any())
}

func TestStrSliceAppend(t *testing.T) {
	{
		// Append one
		slice := NewStrSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{"2"}, slice.Ex())
	}
	{
		// Append many
		slice := NewStrSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2", "4", "6")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []string{"2", "4", "6"}, slice.Ex())
	}
}
func TestStrSliceAt(t *testing.T) {
	slice := NewStrSlice().Append("1", "2", "3", "4")
	assert.Equal(t, "4", slice.At(-1))
	assert.Equal(t, "3", slice.At(-2))
	assert.Equal(t, "2", slice.At(-3))
	assert.Equal(t, "1", slice.At(0))
	assert.Equal(t, "2", slice.At(1))
	assert.Equal(t, "3", slice.At(2))
	assert.Equal(t, "4", slice.At(3))
}

func TestStrSliceClear(t *testing.T) {
	slice := NewStrSlice().Append("1", "2", "3", "4")
	assert.Equal(t, 4, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
}

func TestStrSliceAnyContain(t *testing.T) {
	assert.True(t, StrSlice([]string{"one", "two", "three"}).AnyContain("thr"))
	assert.False(t, StrSlice([]string{"one", "two", "three"}).AnyContain("2"))
}

func TestStrSliceContains(t *testing.T) {
	assert.True(t, StrSlice([]string{"1", "2", "3"}).Contains("2"))
	assert.False(t, StrSlice([]string{"1", "2", "3"}).Contains("4"))
}

func TestStrSliceContainsAny(t *testing.T) {
	assert.True(t, StrSlice([]string{"1", "2", "3"}).ContainsAny([]string{"2"}))
	assert.False(t, StrSlice([]string{"1", "2", "3"}).ContainsAny([]string{"4"}))
}

func TestStrSliceJoin(t *testing.T) {
	assert.Equal(t, "", StrSlice([]string{}).Join(".").Ex())
	assert.Equal(t, "1", StrSlice([]string{"1"}).Join(".").Ex())
	assert.Equal(t, "1.2", StrSlice([]string{"1", "2"}).Join(".").Ex())
}

func TestStrSliceLen(t *testing.T) {
	assert.Equal(t, 0, StrSlice([]string{}).Len())
	assert.Equal(t, 1, StrSlice([]string{"1"}).Len())
	assert.Equal(t, 2, StrSlice([]string{"1", "2"}).Len())
}

func TestStrSlicePrepend(t *testing.T) {
	slice := NewStrSlice().Prepend("1")
	assert.Equal(t, "1", slice.At(0))

	slice.Prepend("2", "3")
	assert.Equal(t, "2", slice.At(0))
	assert.Equal(t, []string{"2", "3", "1"}, slice.Ex())
}

func TestStrSliceTakeFirst(t *testing.T) {
	{
		slice := StrSlice([]string{"0", "1", "2"})
		results := []string{}
		expected := []string{"0", "1", "2"}
		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
			results = append(results, item)
		}
		assert.Equal(t, expected, results)
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{"1", "2"}, slice.Ex())
	}
	{
		slice := StrSlice([]string{"0"})
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, slice.Ex())
	}
	{
		slice := StrSlice([]string{})
		item, ok := slice.TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, slice.Ex())
	}
}

func TestStrSliceTakeFirstCnt(t *testing.T) {
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeFirstCnt(2).Ex()
		assert.Equal(t, []string{"0", "1"}, items)
		assert.Equal(t, []string{"2"}, slice.Ex())
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeFirstCnt(3).Ex()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Ex())
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeFirstCnt(4).Ex()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Ex())
	}
}

func TestStrSliceTakeLast(t *testing.T) {
	{
		slice := StrSlice([]string{"0", "1", "2"})
		results := []string{}
		expected := []string{"2", "1", "0"}
		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
			results = append(results, item)
		}
		assert.Equal(t, expected, results)
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, "2", item)
		assert.Equal(t, []string{"0", "1"}, slice.Ex())
	}
	{
		slice := StrSlice([]string{"0"})
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, slice.Ex())
	}
	{
		slice := StrSlice([]string{})
		item, ok := slice.TakeLast()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, slice.Ex())
	}
}
func TestStrSliceTakeLastCnt(t *testing.T) {
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeLastCnt(2).Ex()
		assert.Equal(t, []string{"1", "2"}, items)
		assert.Equal(t, []string{"0"}, slice.Ex())
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeLastCnt(3).Ex()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Ex())
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeLastCnt(4).Ex()
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Ex())
	}
}

func TestStrSliceUniq(t *testing.T) {
	{
		data := StrSlice([]string{}).Uniq().Ex()
		expected := []string{}
		assert.Equal(t, expected, data)
	}
	{
		data := StrSlice([]string{"1", "2", "3"}).Uniq().Ex()
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
	{
		data := StrSlice([]string{"1", "2", "2", "3"}).Uniq().Ex()
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
}

//--------------------------------------------------------------------------------------------------
// StrMapSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewStrMapSlice(t *testing.T) {
	assert.NotNil(t, NewStrMapSlice().Ex())
}

func TestStrMapSlice(t *testing.T) {
	assert.NotNil(t, StrMapSlice(nil).Ex())
	assert.NotNil(t, StrMapSlice([]map[string]interface{}{}).Ex())
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
		assert.Equal(t, []map[string]interface{}{{"2": "two"}}, slice.Ex())
	}
	{
		// Append many
		slice := NewStrMapSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(map[string]interface{}{"1": "one"}, map[string]interface{}{"2": "two"})
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []map[string]interface{}{{"1": "one"}, {"2": "two"}}, slice.Ex())
	}
}

func TestStrMapSliceAt(t *testing.T) {
	slice := NewStrMapSlice()
	assert.Equal(t, 0, slice.Len())
	slice.Append(map[string]interface{}{"1": "one"})
	slice.Append(map[string]interface{}{"2": "two"})
	slice.Append(map[string]interface{}{"3": "three"})

	assert.Equal(t, map[string]interface{}{"3": "three"}, slice.At(-1).Ex())
	assert.Equal(t, map[string]interface{}{"2": "two"}, slice.At(-2).Ex())
	assert.Equal(t, map[string]interface{}{"1": "one"}, slice.At(0).Ex())
	assert.Equal(t, map[string]interface{}{"2": "two"}, slice.At(1).Ex())
	assert.Equal(t, map[string]interface{}{"3": "three"}, slice.At(2).Ex())
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
	assert.Equal(t, map[string]interface{}{"1": "one"}, slice.At(0).Ex())

	slice.Prepend(map[string]interface{}{"2": "two"}, map[string]interface{}{"3": "three"})
	assert.Equal(t, map[string]interface{}{"2": "two"}, slice.At(0).Ex())

	expected := []map[string]interface{}{
		{"2": "two"},
		{"3": "three"},
		{"1": "one"},
	}
	assert.Equal(t, expected, slice.Ex())
}

// func TestStrMapSliceTakeFirst(t *testing.T) {
// 	{
// 		slice := StrMapSlice([]map[string]interface{}{
// 			{"1": "one"},
// 			{"2": "two"},
// 			{"3": "three"},
// 		})
// 		results := []map[string]interface{}{}
// 		expected := []map[string]interface{}{
// 			{"1": "one"},
// 			{"2": "two"},
// 			{"3": "three"},
// 		}
// 		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
// 			results = append(results, item)
// 		}
// 		assert.Equal(t, expected, results)
// 	}
// 	// 	{
// 	// 		slice := StrSlice([]string{"0", "1", "2"})
// 	// 		item, ok := slice.TakeFirst()
// 	// 		assert.True(t, ok)
// 	// 		assert.Equal(t, "0", item)
// 	// 		assert.Equal(t, []string{"1", "2"}, slice.Ex())
// 	// 	}
// 	// 	{
// 	// 		slice := StrSlice([]string{"0"})
// 	// 		item, ok := slice.TakeFirst()
// 	// 		assert.True(t, ok)
// 	// 		assert.Equal(t, "0", item)
// 	// 		assert.Equal(t, []string{}, slice.Ex())
// 	// 	}
// 	// 	{
// 	// 		slice := StrSlice([]string{})
// 	// 		item, ok := slice.TakeFirst()
// 	// 		assert.False(t, ok)
// 	// 		assert.Equal(t, "", item)
// 	// 		assert.Equal(t, []string{}, slice.Ex())
// 	// 	}
// }

// func TestStrTakeFirstCnt(t *testing.T) {
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeFirstCnt(2).Ex()
// 		assert.Equal(t, []string{"0", "1"}, items)
// 		assert.Equal(t, []string{"2"}, slice.Ex())
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeFirstCnt(3).Ex()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Ex())
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeFirstCnt(4).Ex()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Ex())
// 	}
// }

// func TestStrTakeLast(t *testing.T) {
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		results := []string{}
// 		expected := []string{"2", "1", "0"}
// 		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
// 			results = append(results, item)
// 		}
// 		assert.Equal(t, expected, results)
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		item, ok := slice.TakeLast()
// 		assert.True(t, ok)
// 		assert.Equal(t, "2", item)
// 		assert.Equal(t, []string{"0", "1"}, slice.Ex())
// 	}
// 	{
// 		slice := StrSlice([]string{"0"})
// 		item, ok := slice.TakeLast()
// 		assert.True(t, ok)
// 		assert.Equal(t, "0", item)
// 		assert.Equal(t, []string{}, slice.Ex())
// 	}
// 	{
// 		slice := StrSlice([]string{})
// 		item, ok := slice.TakeLast()
// 		assert.False(t, ok)
// 		assert.Equal(t, "", item)
// 		assert.Equal(t, []string{}, slice.Ex())
// 	}
// }
// func TestStrTakeLastCnt(t *testing.T) {
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeLastCnt(2).Ex()
// 		assert.Equal(t, []string{"1", "2"}, items)
// 		assert.Equal(t, []string{"0"}, slice.Ex())
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeLastCnt(3).Ex()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Ex())
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeLastCnt(4).Ex()
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Ex())
// 	}
// }
