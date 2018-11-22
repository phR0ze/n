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

func TestIntTakeFirst(t *testing.T) {
	{
		data := []int{0, 1, 2}
		data, item, ok := IntSlice(data).TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{1, 2}, data)
	}
	{
		data := []int{0}
		data, item, ok := IntSlice(data).TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, data)
	}
	{
		data := []int{}
		data, item, ok := IntSlice(data).TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, data)
	}
}

func TestIntTakeLast(t *testing.T) {
	{
		data := []int{0, 1, 2}
		data, item, ok := IntSlice(data).TakeLast()
		assert.True(t, ok)
		assert.Equal(t, 2, item)
		assert.Equal(t, []int{0, 1}, data)
	}
	{
		data := []int{0}
		data, item, ok := IntSlice(data).TakeLast()
		assert.True(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, data)
	}
	{
		data := []int{}
		data, item, ok := IntSlice(data).TakeLast()
		assert.False(t, ok)
		assert.Equal(t, 0, item)
		assert.Equal(t, []int{}, data)
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

func TestStrTakeFirst(t *testing.T) {
	{
		data := []string{"0", "1", "2"}
		data, item, ok := StrSlice(data).TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{"1", "2"}, data)
	}
	{
		data := []string{"0"}
		data, item, ok := StrSlice(data).TakeFirst()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, data)
	}
	{
		data := []string{}
		data, item, ok := StrSlice(data).TakeFirst()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, data)
	}
}

func TestStrTakeLast(t *testing.T) {
	{
		data := []string{"0", "1", "2"}
		data, item, ok := StrSlice(data).TakeLast()
		assert.True(t, ok)
		assert.Equal(t, "2", item)
		assert.Equal(t, []string{"0", "1"}, data)
	}
	{
		data := []string{"0"}
		data, item, ok := StrSlice(data).TakeLast()
		assert.True(t, ok)
		assert.Equal(t, "0", item)
		assert.Equal(t, []string{}, data)
	}
	{
		data := []string{}
		data, item, ok := StrSlice(data).TakeLast()
		assert.False(t, ok)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{}, data)
	}
}
