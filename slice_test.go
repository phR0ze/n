package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//--------------------------------------------------------------------------------------------------
// IntSlice tests
//--------------------------------------------------------------------------------------------------

func TestIntContains(t *testing.T) {
	assert.True(t, IntSlice([]int{1, 2, 3}).Contains(2))
	assert.False(t, IntSlice([]int{1, 2, 3}).Contains(4))
}

func TestIntContainsAny(t *testing.T) {
	assert.True(t, IntSlice([]int{1, 2, 3}).ContainsAny([]int{2}))
	assert.False(t, IntSlice([]int{1, 2, 3}).ContainsAny([]int{4}))
}

func TestIntDistinct(t *testing.T) {
	{
		data := IntSlice([]int{}).Distinct()
		expected := []int{}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 3}).Distinct()
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
	{
		data := IntSlice([]int{1, 2, 2, 3}).Distinct()
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, data)
	}
}

func TestIntLen(t *testing.T) {
	assert.Equal(t, 0, IntSlice([]int{}).Len())
	assert.Equal(t, 1, IntSlice([]int{1}).Len())
	assert.Equal(t, 2, IntSlice([]int{1, 2}).Len())
}

func TestIntTakeFirst(t *testing.T) {
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

func TestIntTakeLast(t *testing.T) {
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

//--------------------------------------------------------------------------------------------------
// StrSlice tests
//--------------------------------------------------------------------------------------------------

func TestStrContains(t *testing.T) {
	assert.True(t, StrSlice([]string{"1", "2", "3"}).Contains("2"))
	assert.False(t, StrSlice([]string{"1", "2", "3"}).Contains("4"))
}

func TestStrContainsAny(t *testing.T) {
	assert.True(t, StrSlice([]string{"1", "2", "3"}).ContainsAny([]string{"2"}))
	assert.False(t, StrSlice([]string{"1", "2", "3"}).ContainsAny([]string{"4"}))
}

func TestStrDistinct(t *testing.T) {
	{
		data := StrSlice([]string{}).Distinct()
		expected := []string{}
		assert.Equal(t, expected, data)
	}
	{
		data := StrSlice([]string{"1", "2", "3"}).Distinct()
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
	{
		data := StrSlice([]string{"1", "2", "2", "3"}).Distinct()
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, data)
	}
}

func TestStrJoin(t *testing.T) {
	assert.Equal(t, "", StrSlice([]string{}).Join("."))
	assert.Equal(t, "1", StrSlice([]string{"1"}).Join("."))
	assert.Equal(t, "1.2", StrSlice([]string{"1", "2"}).Join("."))
}

func TestStrLen(t *testing.T) {
	assert.Equal(t, 0, StrSlice([]string{}).Len())
	assert.Equal(t, 1, StrSlice([]string{"1"}).Len())
	assert.Equal(t, 2, StrSlice([]string{"1", "2"}).Len())
}

func TestStrTakeFirst(t *testing.T) {
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

func TestStrTakeLast(t *testing.T) {
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
