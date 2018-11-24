package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------------------------------------------------------------------------------
// IntSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewIntSlice(t *testing.T) {
	assert.NotNil(t, NewIntSlice().Raw)
}

func TestIntSlice(t *testing.T) {
	assert.NotNil(t, IntSlice(nil).Raw)
	assert.NotNil(t, IntSlice([]int{}).Raw)
}

func TestIntSliceAppend(t *testing.T) {
	{
		// Append one
		slice := NewIntSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(2)
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []int{2}, slice.Raw)
	}
	{
		// Append many
		slice := NewIntSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(2, 4, 6)
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []int{2, 4, 6}, slice.Raw)
	}
}

func TestIntSliceContains(t *testing.T) {
	assert.True(t, IntSlice([]int{1, 2, 3}).Contains(2))
	assert.False(t, IntSlice([]int{1, 2, 3}).Contains(4))
}

func TestIntSliceContainsAny(t *testing.T) {
	assert.True(t, IntSlice([]int{1, 2, 3}).ContainsAny([]int{2}))
	assert.False(t, IntSlice([]int{1, 2, 3}).ContainsAny([]int{4}))
}

func TestIntSliceDistinct(t *testing.T) {
	{
		data := IntSlice([]int{}).Distinct().Raw
		expected := []int{}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 3}).Distinct().Raw
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 2, 3}).Distinct().Raw
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
}

func TestIntSliceLen(t *testing.T) {
	assert.Equal(t, 0, IntSlice([]int{}).Len())
	assert.Equal(t, 1, IntSlice([]int{1}).Len())
	assert.Equal(t, 2, IntSlice([]int{1, 2}).Len())
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
		assert.Equal(t, []int{1, 2}, slice.Raw)
	}
	{
		slice := IntSlice([]int{0})
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Raw)
	}
	{
		slice := IntSlice([]int{})
		item, ok := slice.TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Raw)
	}
}

func TestIntSliceTakeFirstCnt(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(2).Raw
		assert.Equal(t, []int{0, 1}, items)
		assert.Equal(t, []int{2}, slice.Raw)
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(3).Raw
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Raw)
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeFirstCnt(4).Raw
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Raw)
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
		assert.Equal(t, []int{0, 1}, slice.Raw)
	}
	{
		slice := IntSlice([]int{0})
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Raw)
	}
	{
		slice := IntSlice([]int{})
		item, ok := slice.TakeLast()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, slice.Raw)
	}
}

func TestIntSliceTakeLastCnt(t *testing.T) {
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(2).Raw
		assert.Equal(t, []int{1, 2}, items)
		assert.Equal(t, []int{0}, slice.Raw)
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(3).Raw
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Raw)
	}
	{
		slice := IntSlice([]int{0, 1, 2})
		items := slice.TakeLastCnt(4).Raw
		assert.Equal(t, []int{0, 1, 2}, items)
		assert.Equal(t, []int{}, slice.Raw)
	}
}

//--------------------------------------------------------------------------------------------------
// StrSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewStrSlice(t *testing.T) {
	assert.NotNil(t, NewStrSlice().Raw)
}

func TestStrSlice(t *testing.T) {
	assert.NotNil(t, StrSlice(nil).Raw)
	assert.NotNil(t, StrSlice([]string{}).Raw)
}

func TestStrSliceAppend(t *testing.T) {
	{
		// Append one
		slice := NewStrSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{"2"}, slice.Raw)
	}
	{
		// Append many
		slice := NewStrSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append("2", "4", "6")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []string{"2", "4", "6"}, slice.Raw)
	}
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

func TestStrSliceDistinct(t *testing.T) {
	{
		data := StrSlice([]string{}).Distinct().Raw
		expected := []string{}
		assert.Equal(t, expected, data)
	}
	{
		data := StrSlice([]string{"1", "2", "3"}).Distinct().Raw
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
	{
		data := StrSlice([]string{"1", "2", "2", "3"}).Distinct().Raw
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
}

func TestStrSliceJoin(t *testing.T) {
	assert.Equal(t, "", StrSlice([]string{}).Join(".").Raw)
	assert.Equal(t, "1", StrSlice([]string{"1"}).Join(".").Raw)
	assert.Equal(t, "1.2", StrSlice([]string{"1", "2"}).Join(".").Raw)
}

func TestStrSliceLen(t *testing.T) {
	assert.Equal(t, 0, StrSlice([]string{}).Len())
	assert.Equal(t, 1, StrSlice([]string{"1"}).Len())
	assert.Equal(t, 2, StrSlice([]string{"1", "2"}).Len())
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
		assert.Equal(t, []string{"1", "2"}, slice.Raw)
	}
	{
		slice := StrSlice([]string{"0"})
		item, ok := slice.TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, slice.Raw)
	}
	{
		slice := StrSlice([]string{})
		item, ok := slice.TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, slice.Raw)
	}
}

func TestStrSliceTakeFirstCnt(t *testing.T) {
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeFirstCnt(2).Raw
		assert.Equal(t, []string{"0", "1"}, items)
		assert.Equal(t, []string{"2"}, slice.Raw)
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeFirstCnt(3).Raw
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Raw)
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeFirstCnt(4).Raw
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Raw)
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
		assert.Equal(t, []string{"0", "1"}, slice.Raw)
	}
	{
		slice := StrSlice([]string{"0"})
		item, ok := slice.TakeLast()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, slice.Raw)
	}
	{
		slice := StrSlice([]string{})
		item, ok := slice.TakeLast()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, slice.Raw)
	}
}
func TestStrSliceTakeLastCnt(t *testing.T) {
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeLastCnt(2).Raw
		assert.Equal(t, []string{"1", "2"}, items)
		assert.Equal(t, []string{"0"}, slice.Raw)
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeLastCnt(3).Raw
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Raw)
	}
	{
		slice := StrSlice([]string{"0", "1", "2"})
		items := slice.TakeLastCnt(4).Raw
		assert.Equal(t, []string{"0", "1", "2"}, items)
		assert.Equal(t, []string{}, slice.Raw)
	}
}

//--------------------------------------------------------------------------------------------------
// StrMapSlice tests
//--------------------------------------------------------------------------------------------------
func TestNewStrMapSlice(t *testing.T) {
	assert.NotNil(t, NewStrMapSlice().Raw)
}

func TestStrMapSlice(t *testing.T) {
	assert.NotNil(t, StrMapSlice(nil).Raw)
	assert.NotNil(t, StrMapSlice([]map[string]interface{}{}).Raw)
}

func TestStrMapSliceAppend(t *testing.T) {
	{
		// Append one
		slice := NewStrMapSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(map[string]interface{}{"2": "two"})
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []map[string]interface{}{{"2": "two"}}, slice.Raw)
	}
	{
		// Append many
		slice := NewStrMapSlice()
		assert.Equal(t, 0, slice.Len())
		slice.Append(map[string]interface{}{"1": "one"}, map[string]interface{}{"2": "two"})
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []map[string]interface{}{{"1": "one"}, {"2": "two"}}, slice.Raw)
	}
}
func TestStrMapSliceContainsKey(t *testing.T) {
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.True(t, StrMapSlice(raw).ContainsKey("1"))
	}
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.False(t, StrMapSlice(raw).ContainsKey("4"))
	}
}

func TestStrMapSliceContainsAny(t *testing.T) {
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.True(t, StrMapSlice(raw).ContainsAnyKey([]string{"4", "1"}))
	}
	{
		raw := []map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
			{"3": "three"},
		}
		assert.False(t, StrMapSlice(raw).ContainsAnyKey([]string{}))
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
// 	// 		assert.Equal(t, []string{"1", "2"}, slice.Raw)
// 	// 	}
// 	// 	{
// 	// 		slice := StrSlice([]string{"0"})
// 	// 		item, ok := slice.TakeFirst()
// 	// 		assert.True(t, ok)
// 	// 		assert.Equal(t, "0", item)
// 	// 		assert.Equal(t, []string{}, slice.Raw)
// 	// 	}
// 	// 	{
// 	// 		slice := StrSlice([]string{})
// 	// 		item, ok := slice.TakeFirst()
// 	// 		assert.False(t, ok)
// 	// 		assert.Equal(t, "", item)
// 	// 		assert.Equal(t, []string{}, slice.Raw)
// 	// 	}
// }

// func TestStrTakeFirstCnt(t *testing.T) {
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeFirstCnt(2).Raw
// 		assert.Equal(t, []string{"0", "1"}, items)
// 		assert.Equal(t, []string{"2"}, slice.Raw)
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeFirstCnt(3).Raw
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Raw)
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeFirstCnt(4).Raw
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Raw)
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
// 		assert.Equal(t, []string{"0", "1"}, slice.Raw)
// 	}
// 	{
// 		slice := StrSlice([]string{"0"})
// 		item, ok := slice.TakeLast()
// 		assert.True(t, ok)
// 		assert.Equal(t, "0", item)
// 		assert.Equal(t, []string{}, slice.Raw)
// 	}
// 	{
// 		slice := StrSlice([]string{})
// 		item, ok := slice.TakeLast()
// 		assert.False(t, ok)
// 		assert.Equal(t, "", item)
// 		assert.Equal(t, []string{}, slice.Raw)
// 	}
// }
// func TestStrTakeLastCnt(t *testing.T) {
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeLastCnt(2).Raw
// 		assert.Equal(t, []string{"1", "2"}, items)
// 		assert.Equal(t, []string{"0"}, slice.Raw)
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeLastCnt(3).Raw
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Raw)
// 	}
// 	{
// 		slice := StrSlice([]string{"0", "1", "2"})
// 		items := slice.TakeLastCnt(4).Raw
// 		assert.Equal(t, []string{"0", "1", "2"}, items)
// 		assert.Equal(t, []string{}, slice.Raw)
// 	}
// }
