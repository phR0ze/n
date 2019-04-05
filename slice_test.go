package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Pair
//--------------------------------------------------------------------------------------------------

func ExampleNSlice_Pair() {
	slice := OldSliceV(1, 2)
	first, second := slice.Pair()
	fmt.Println(first.O(), second.O())
	// Output: 1 2
}

func TestNSlice_Pair(t *testing.T) {

	// int
	{
		// two values
		{
			first, second := OldSliceV(1, 2).Pair()
			assert.Equal(t, &Object{1}, first)
			assert.Equal(t, &Object{2}, second)
		}

		// one value
		{
			first, second := OldSliceV(1).Pair()
			assert.Equal(t, &Object{1}, first)
			assert.Equal(t, &Object{nil}, second)
		}

		// no values
		{
			first, second := OldSliceV().Pair()
			assert.Equal(t, &Object{nil}, first)
			assert.Equal(t, &Object{nil}, second)
		}
	}

	// custom
	{
		// two values
		{
			first, second := OldSlice([]Object{{1}, {2}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, &Object{Object{2}}, second)
		}

		// one value
		{
			first, second := OldSlice([]Object{{1}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, &Object{nil}, second)
		}

		// no values
		{
			first, second := OldSliceV().Pair()
			assert.Equal(t, &Object{nil}, first)
			assert.Equal(t, &Object{nil}, second)
		}
	}
}

// Prepend
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Prepend_Normal(t *testing.B) {
	ints := []int{}
	for i := range Range(0, nines6) {
		ints = append(ints, i)
		copy(ints[1:], ints[1:])
		ints[0] = i
	}
}

func BenchmarkNSlice_Prepend_Optimized(t *testing.B) {
	slice := &NSlice{o: []int{}}
	for i := range Range(0, nines6) {
		slice.Prepend(i)
	}
}

func BenchmarkNSlice_Prepend_Reflect(t *testing.B) {
	slice := &NSlice{o: []Object{}}
	for i := range Range(0, nines6) {
		slice.Prepend(Object{i})
	}
}

func ExampleNSlice_Prepend() {
	slice := OldSliceV(2, 3)
	fmt.Println(slice.Prepend(1).O())
	// Output: [1 2 3]
}

func TestNSlice_Prepend(t *testing.T) {

	// int
	{
		// happy path
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(2), slice.Prepend(2))
			assert.Equal(t, OldSliceV(1, 2), slice.Prepend(1))
			assert.Equal(t, OldSliceV(0, 1, 2), slice.Prepend(0))
		}

		// error cases
		{
			var slice *NSlice
			assert.True(t, slice.Prepend(0).Nil())
			assert.Equal(t, (*NSlice)(nil), slice.Prepend(0))
		}
	}

	// custom
	{
		// prepend
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(Object{2}), slice.Prepend(Object{2}))
			assert.Equal(t, OldSlice([]Object{{1}, {2}}), slice.Prepend(Object{1}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}}), slice.Prepend(Object{0}))
		}

		// error cases
		{
			var slice *NSlice
			assert.True(t, slice.Prepend(Object{0}).Nil())
			assert.Equal(t, (*NSlice)(nil), slice.Prepend(Object{0}))
		}
	}
}

// Set
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Set_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		ints[i] = 0
	}
}

func BenchmarkNSlice_Set_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, 0)
	}
}

func BenchmarkNSlice_Set_Reflect(t *testing.B) {
	slice := OldSlice(rangeInterObject(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, Object{0})
	}
}

func ExampleNSlice_Set() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Set(0, 0).O())
	// Output: [0 2 3]
}

func TestNSlice_Set(t *testing.T) {
	// bool
	{
		assert.Equal(t, []bool{false, true, true}, OldSliceV(true, true, true).Set(0, false).O())
		assert.Equal(t, []bool{true, false, true}, OldSliceV(true, true, true).Set(1, false).O())
		assert.Equal(t, []bool{true, true, false}, OldSliceV(true, true, true).Set(2, false).O())
		assert.Equal(t, []bool{false, true, true}, OldSliceV(true, true, true).Set(-3, false).O())
		assert.Equal(t, []bool{true, false, true}, OldSliceV(true, true, true).Set(-2, false).O())
		assert.Equal(t, []bool{true, true, false}, OldSliceV(true, true, true).Set(-1, false).O())
	}

	// int
	{
		assert.Equal(t, []int{0, 2, 3}, OldSliceV(1, 2, 3).Set(0, 0).O())
		assert.Equal(t, []int{1, 0, 3}, OldSliceV(1, 2, 3).Set(1, 0).O())
		assert.Equal(t, []int{1, 2, 0}, OldSliceV(1, 2, 3).Set(2, 0).O())
		assert.Equal(t, []int{0, 2, 3}, OldSliceV(1, 2, 3).Set(-3, 0).O())
		assert.Equal(t, []int{1, 0, 3}, OldSliceV(1, 2, 3).Set(-2, 0).O())
		assert.Equal(t, []int{1, 2, 0}, OldSliceV(1, 2, 3).Set(-1, 0).O())
	}

	// string
	{
		assert.Equal(t, []string{"0", "2", "3"}, OldSliceV("1", "2", "3").Set(0, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, OldSliceV("1", "2", "3").Set(1, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, OldSliceV("1", "2", "3").Set(2, "0").O())
		assert.Equal(t, []string{"0", "2", "3"}, OldSliceV("1", "2", "3").Set(-3, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, OldSliceV("1", "2", "3").Set(-2, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, OldSliceV("1", "2", "3").Set(-1, "0").O())
	}

	// custom
	{
		assert.Equal(t, []Object{{0}, {2}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(0, Object{0}).O())
		assert.Equal(t, []Object{{1}, {0}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(1, Object{0}).O())
		assert.Equal(t, []Object{{1}, {2}, {0}}, OldSlice([]Object{{1}, {2}, {3}}).Set(2, Object{0}).O())
		assert.Equal(t, []Object{{0}, {2}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(-3, Object{0}).O())
		assert.Equal(t, []Object{{1}, {0}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(-2, Object{0}).O())
		assert.Equal(t, []Object{{1}, {2}, {0}}, OldSlice([]Object{{1}, {2}, {3}}).Set(-1, Object{0}).O())
	}

	// panics need to run as the last test as they abort the test method
	defer func() {
		err := recover()
		assert.Equal(t, "slice assignment is out of bounds", err)
	}()
	OldSliceV(1, 2, 3).Set(5, 1)
}

// Single
//--------------------------------------------------------------------------------------------------

func ExampleNSlice_Single() {
	slice := OldSliceV(1)
	fmt.Println(slice.Single())
	// Output: true
}

func TestNSlice_Single(t *testing.T) {

	// int
	{
		assert.Equal(t, false, OldSliceV().Single())
		assert.Equal(t, true, OldSliceV(1).Single())
		assert.Equal(t, false, OldSliceV(1, 2).Single())
	}

	// custom
	{
		assert.Equal(t, false, OldSliceV().Single())
		assert.Equal(t, true, OldSliceV(Object{1}).Single())
		assert.Equal(t, false, OldSliceV(Object{1}, Object{2}).Single())
	}
}

// Slice
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Slice_Normal(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:len(ints)]
}

func BenchmarkNSlice_Slice_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	slice.Slice(0, -1)
}

func BenchmarkNSlice_Slice_Reflect(t *testing.B) {
	slice := OldSlice(rangeInterObject(0, nines7))
	slice.Slice(0, -1)
}

func ExampleNSlice_Slice() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Slice(1, -1).O())
	// Output: [2 3]
}

func TestNSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, OldSliceV(), nilSlice.Slice(0, -1))
		slice := OldSliceV(0).Clear()
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(0, -1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := OldSliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, []int{0, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2, 3}, result.O())
	}

	// slice full array
	{
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Slice(0, -1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Slice(0, 1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Slice(0, 5))
		assert.Equal(t, OldSliceV(""), OldSliceV("").Slice(0, -1))
		assert.Equal(t, OldSliceV(""), OldSliceV("").Slice(0, 1))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).Slice(0, -1))
		assert.Equal(t, OldSlice([]int{1, 2, 3}), OldSlice([]int{1, 2, 3}).Slice(0, -1))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Slice(0, 2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, OldSliceV("1"), OldSliceV("1").Slice(0, 2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false).Slice(-6, 6))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).Slice(-6, 6))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Slice(-6, 6))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-6, 6))
	}

	// mutually exclusive
	{
		slice := OldSliceV(1, 2, 3, 4)
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(2, -3))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(0, -5))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(4, -1))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(6, -1))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(3, 2))
	}

	// singles
	{
		slice := OldSliceV(1, 2, 3, 4)
		assert.Equal(t, OldSliceV(4), slice.Slice(-1, -1))
		assert.Equal(t, OldSliceV(3), slice.Slice(-2, -2))
		assert.Equal(t, OldSliceV(2), slice.Slice(-3, -3))
		assert.Equal(t, OldSliceV(1), slice.Slice(0, 0))
		assert.Equal(t, OldSliceV(1), slice.Slice(-4, -4))
		assert.Equal(t, OldSliceV(2), slice.Slice(1, 1))
		assert.Equal(t, OldSliceV(2), slice.Slice(1, -3))
		assert.Equal(t, OldSliceV(3), slice.Slice(2, 2))
		assert.Equal(t, OldSliceV(3), slice.Slice(2, -2))
		assert.Equal(t, OldSliceV(4), slice.Slice(3, 3))
		assert.Equal(t, OldSliceV(4), slice.Slice(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(1, -1))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(1, 2))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(-2, -1))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(-2, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(1, -1))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(-2, -1))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(-2, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(1, -1))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(-2, -1))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(-2, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(1, -1))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(1, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-2, -1))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(0, -2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(-3, -2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(-3, 1))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(0, 1))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(0, -2))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(-3, -2))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(-3, 1))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(0, 1))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(0, -2))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(-3, -2))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(-3, 1))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(0, 1))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(0, -2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-3, -2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-3, 1))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(1, -2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(-3, -2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(-3, 2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(1, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(1, -2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(-3, -2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(-3, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(1, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(1, -2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(-3, -2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(-3, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(1, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, -2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, -2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, 2))
	}

	// random
	{
		assert.Equal(t, OldSliceV("1"), OldSliceV("1", "2", "3").Slice(0, -3))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Slice(0, 2))
	}
}

// Sort
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Sort_Normal(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:len(ints)]
}

func BenchmarkNSlice_Sort_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	slice.Slice(0, -1)
}

func BenchmarkNSlice_Sort_Reflect(t *testing.B) {
	slice := OldSlice(rangeInterObject(0, nines7))
	slice.Slice(0, -1)
}

func ExampleNSlice_Sort() {
	slice := OldSliceV(2, 3, 1)
	fmt.Println(slice.Sort().O())
	// Output: [1 2 3]
}

func TestNSlice_Sort(t *testing.T) {

	// empty
	//assert.Equal(t, SliceV(), SliceV().Sort())

	// bool
	//assert.Equal(t, SliceV(false, true, true), SliceV(true, false, true).Sort())

	// int
	assert.Equal(t, OldSliceV(1, 2, 3, 4, 5), OldSliceV(5, 3, 2, 4, 1).Sort())

	// string
	//assert.Equal(t, SliceV("1", "2", "3", "4", "5"), SliceV("5", "3", "2", "4", "1").Sort())

	// custom
	//assert.Equal(t, Slice([]Object{{1}, {2}, {3}, {4}, {5}}), Slice([]Object{{5}, {3}, {2}, {4}, {1}}).Sort())
}

// Take
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Take_Normal(t *testing.B) {
	ints := Range(0, nines5)
	index := Range(0, nines5)
	for i := range index {
		if i+1 < len(ints) {
			ints = append(ints[:i], ints[i+1:]...)
		} else if i >= 0 && i < len(ints) {
			ints = ints[:i]
		}
	}
}

func BenchmarkNSlice_Take_Optimized(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := OldSlice(src)
	for i := range index {
		slice.Take(i)
	}
}

func BenchmarkNSlice_Take_Reflect(t *testing.B) {
	src := rangeInterObject(0, nines5)
	index := Range(0, nines5)
	slice := OldSlice(src)
	for i := range index {
		slice.Take(i)
	}
}

func ExampleNSlice_Take() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Take(2).O())
	// Output: 3
}

func TestNSlice_Take(t *testing.T) {

	// int
	{
		// nil or empty
		{
			var nilSlice *NSlice
			assert.Equal(t, &Object{}, nilSlice.Take(0))
		}

		// Delete all and more
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-1)
			assert.Equal(t, &Object{2}, obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{1}, obj)
			assert.Equal(t, []int{0}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []int{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.Take(-1)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []int{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(3)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(2)
			assert.Equal(t, &Object{2}, obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(1)
			assert.Equal(t, &Object{1}, obj)
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(0)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []int{1, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-4)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-1)
			assert.Equal(t, &Object{2}, obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-2)
			assert.Equal(t, &Object{1}, obj)
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// custom
	{
		// Delete all and more
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []Object{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.Take(-1)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []Object{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(3)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(2)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(0)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []Object{{1}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-4)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-2)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}
}

// // func TestStrSliceTakeFirst(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		results := []string{}
// // 		expected := []string{"0", "1", "2"}
// // 		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
// // 			results = append(results, item)
// // 		}
// // 		assert.Equal(t, expected, results)
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		item, ok := slice.TakeFirst()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{"1", "2"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0")
// // 		item, ok := slice.TakeFirst()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S()
// // 		item, ok := slice.TakeFirst()
// // 		assert.False(t, ok)
// // 		assert.Equal(t, "", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceTakeFirstCnt(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(2).S()
// // 		assert.Equal(t, []string{"0", "1"}, items)
// // 		assert.Equal(t, []string{"2"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(3).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(4).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceTakeLast(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		results := []string{}
// // 		expected := []string{"2", "1", "0"}
// // 		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
// // 			results = append(results, item)
// // 		}
// // 		assert.Equal(t, expected, results)
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		item, ok := slice.TakeLast()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "2", item)
// // 		assert.Equal(t, []string{"0", "1"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0")
// // 		item, ok := slice.TakeLast()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S()
// // 		item, ok := slice.TakeLast()
// // 		assert.False(t, ok)
// // 		assert.Equal(t, "", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }
// // func TestStrSliceTakeLastCnt(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(2).S()
// // 		assert.Equal(t, []string{"1", "2"}, items)
// // 		assert.Equal(t, []string{"0"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(3).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(4).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceUniq(t *testing.T) {
// // 	{
// // 		data := S().Uniq().S()
// // 		expected := []string{}
// // 		assert.Equal(t, expected, data)
// // 	}
// // 	{
// // 		data := S("1", "2", "3").Uniq().S()
// // 		expected := []string{"1", "2", "3"}
// // 		assert.Equal(t, expected, data)
// // 	}
// // 	{
// // 		data := S("1", "2", "2", "3").Uniq().S()
// // 		expected := []string{"1", "2", "3"}
// // 		assert.Equal(t, expected, data)
// // 	}
// // }

// // func TestYamlPair(t *testing.T) {
// // 	{
// // 		k, v := A("foo=bar").Split("=").YamlPair()
// // 		assert.Equal(t, "foo", k)
// // 		assert.Equal(t, "bar", v)
// // 	}
// // 	{
// // 		k, v := A("=bar").Split("=").YamlPair()
// // 		assert.Equal(t, "", k)
// // 		assert.Equal(t, "bar", v)
// // 	}
// // 	{
// // 		k, v := A("bar=").Split("=").YamlPair()
// // 		assert.Equal(t, "bar", k)
// // 		assert.Equal(t, "", v)
// // 	}
// // 	{
// // 		k, v := A("").Split("=").YamlPair()
// // 		assert.Equal(t, "", k)
// // 		assert.Equal(t, nil, v)
// // 	}
// // }
// // func TestYamlKeyVal(t *testing.T) {
// // 	{
// // 		pair := A("foo=bar").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "foo", pair.Key)
// // 		assert.Equal(t, "bar", pair.Val)
// // 	}
// // 	{
// // 		pair := A("=bar").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "", pair.Key)
// // 		assert.Equal(t, "bar", pair.Val)
// // 	}
// // 	{
// // 		pair := A("bar=").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "bar", pair.Key)
// // 		assert.Equal(t, "", pair.Val)
// // 	}
// // 	{
// // 		pair := A("").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "", pair.Key)
// // 		assert.Equal(t, "", pair.Val)
// // 	}
// // }

func TestSlice_absIndex(t *testing.T) {
	//             -4,-3,-2,-1
	//              0, 1, 2, 3
	assert.Equal(t, 3, absIndex(4, -1))
	assert.Equal(t, 2, absIndex(4, -2))
	assert.Equal(t, 1, absIndex(4, -3))
	assert.Equal(t, 0, absIndex(4, -4))

	assert.Equal(t, 0, absIndex(4, 0))
	assert.Equal(t, 1, absIndex(4, 1))
	assert.Equal(t, 2, absIndex(4, 2))
	assert.Equal(t, 3, absIndex(4, 3))

	// out of bounds
	assert.Equal(t, -1, absIndex(4, 4))
	assert.Equal(t, -1, absIndex(4, -5))
}

func TestSlice_absIndices(t *testing.T) {
	len := 4
	// -4,-3,-2,-1
	//  0, 1, 2, 3

	// no indicies given
	{
		i, j, err := absIndices(len)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// one index given
	{
		i, j, err := absIndices(len, 1)
		assert.Equal(t, 0, i)
		assert.Equal(t, -1, j)
		assert.Equal(t, "only one index given", err.Error())
	}

	// end
	{
		i, j, err := absIndices(len, -3, -1)
		assert.Equal(t, 1, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 1, 3)
		assert.Equal(t, 1, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// middle
	{
		i, j, err := absIndices(len, 1, 2)
		assert.Equal(t, 1, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -3, -2)
		assert.Equal(t, 1, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)
	}

	// begining
	{
		i, j, err := absIndices(len, 0, 2)
		assert.Equal(t, 0, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -4, -2)
		assert.Equal(t, 0, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)
	}

	// move within bounds
	{
		i, j, err := absIndices(len, -5, 5)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 0, 5)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -5, -1)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// mutually exclusive
	{
		i, j, err := absIndices(len, -1, -3)
		assert.Equal(t, 3, i)
		assert.Equal(t, 1, j)
		assert.NotNil(t, err)

		i, j, err = absIndices(len, 3, 1)
		assert.Equal(t, 3, i)
		assert.Equal(t, 1, j)
		assert.NotNil(t, err)
	}

	// single
	{
		i, j, err := absIndices(len, 0, 0)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 1, 1)
		assert.Equal(t, 1, i)
		assert.Equal(t, 2, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 3, 3)
		assert.Equal(t, 3, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -1, -1)
		assert.Equal(t, 3, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -2, -2)
		assert.Equal(t, 2, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -4, -4)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
		assert.Nil(t, err)
	}
}
