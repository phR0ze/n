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
