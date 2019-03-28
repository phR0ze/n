package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewIntSlice
//--------------------------------------------------------------------------------------------------
func BenchmarkNewIntSlice_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i += 10 {
		_ = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9}
	}
}

func BenchmarkNewIntSlice_Slice(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i += 10 {
		_ = NewIntSlice([]int{i, i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9})
	}
}

func ExampleNewIntSlice() {
	slice := NewIntSlice([]int{1, 2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestIntSlice_NewIntSlice(t *testing.T) {

	// array
	var array [2]int
	array[0] = 1
	array[1] = 2
	assert.Equal(t, []int{1, 2}, NewIntSlice(array[:]).O())

	// empty
	assert.Equal(t, []int{}, NewIntSlice([]int{}).O())

	// slice
	assert.Equal(t, []int{0}, NewSlice([]int{0}).O())
	assert.Equal(t, []int{1, 2}, NewSlice([]int{1, 2}).O())
}

// NewIntSliceV
//--------------------------------------------------------------------------------------------------
func BenchmarkNewIntSliceV_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i += 10 {
		_ = append([]int{}, i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9)
	}
}

func BenchmarkNewIntSliceV_Slice(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i += 10 {
		_ = NewIntSliceV(i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9)
	}
}

func ExampleNewIntSliceV_empty() {
	slice := NewIntSliceV()
	fmt.Println(slice.O())
	// Output: []
}

func ExampleNewIntSliceV_variadic() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestIntSlice_NewIntSliceV(t *testing.T) {

	// array
	var array [2]int
	array[0] = 1
	array[1] = 2
	assert.Equal(t, []int{1, 2}, NewIntSliceV(array[:]...).O())

	// empty
	assert.Equal(t, []int{}, NewIntSliceV().O())

	// multiples
	assert.Equal(t, []int{1}, NewIntSliceV(1).O())
	assert.Equal(t, []int{1, 2}, NewIntSliceV(1, 2).O())
	assert.Equal(t, []int{1, 2}, NewIntSliceV([]int{1, 2}...).O())
}

// Any
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Any_Go(t *testing.B) {
	any := func(list []int, x []int) bool {
		for i := range x {
			for j := range list {
				if list[j] == x[i] {
					return true
				}
			}
		}
		return false
	}

	// test here
	ints := Range(0, nines4)
	for i := range ints {
		any(ints, []int{i})
	}
}

func BenchmarkIntSlice_Any_Slice(t *testing.B) {
	src := Range(0, nines4)
	slice := NewIntSlice(src)
	for i := range src {
		slice.Any(i)
	}
}

func ExampleIntSlice_Any_empty() {
	slice := NewIntSliceV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleIntSlice_Any_notEmpty() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleIntSlice_Any_contains() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Any(1))
	// Output: true
}

func ExampleIntSlice_Any_containsAny() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Any(0, 1))
	// Output: true
}

func TestIntSlice_Any(t *testing.T) {

	// empty
	var nilSlice *IntSlice
	assert.False(t, nilSlice.Any())
	assert.False(t, NewIntSliceV().Any())

	// single
	assert.True(t, NewIntSliceV(2).Any())

	// invalid
	assert.False(t, NewIntSliceV(1, 2).Any(NObj{2}))

	assert.True(t, NewIntSliceV(1, 2, 3).Any(2))
	assert.False(t, NewIntSliceV(1, 2, 3).Any(4))
	assert.True(t, NewIntSliceV(1, 2, 3).Any(4, 3))
	assert.False(t, NewIntSliceV(1, 2, 3).Any(4, 5))
}

// AnyS
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_AnyS_Go(t *testing.B) {
	any := func(list []int, x []int) bool {
		for i := range x {
			for j := range list {
				if list[j] == x[i] {
					return true
				}
			}
		}
		return false
	}

	// test here
	ints := Range(0, nines4)
	for i := range ints {
		any(ints, []int{i})
	}
}

func BenchmarkIntSlice_AnyS_Slice(t *testing.B) {
	src := Range(0, nines4)
	slice := NewIntSlice(src)
	for i := range src {
		slice.Any([]int{i})
	}
}

func ExampleIntSlice_AnyS() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.AnyS([]int{0, 1}))
	// Output: true
}

func TestIntSlice_AnyS(t *testing.T) {
	// nil
	var nilSlice *NSlice
	assert.False(t, nilSlice.AnyS([]int{1}))

	// int
	assert.True(t, NewIntSliceV(1, 2, 3).AnyS([]int{1}))
	assert.True(t, NewIntSliceV(1, 2, 3).AnyS([]int{4, 3}))
	assert.False(t, NewIntSliceV(1, 2, 3).AnyS([]int{4, 5}))

	// invalid
	assert.False(t, NewIntSliceV(1, 2).AnyS([]string{"2"}))
}

// Append
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Append_Go(t *testing.B) {
	ints := []int{}
	for _, i := range Range(0, nines6) {
		ints = append(ints, i)
	}
}

func BenchmarkIntSlice_Append_Slice(t *testing.B) {
	slice := NewIntSliceV()
	for _, i := range Range(0, nines6) {
		slice.Append(i)
	}
}

func ExampleIntSlice_Append() {
	slice := NewIntSliceV(1).Append(2).Append(3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestIntSlice_Append(t *testing.T) {

	// nil
	{
		var nilSlice *IntSlice
		assert.Equal(t, NewIntSliceV(0), nilSlice.Append(0))
		assert.Equal(t, (*IntSlice)(nil), nilSlice)
	}

	// Append one back to back
	{
		var slice *IntSlice
		assert.Equal(t, true, slice.Nil())
		slice = NewIntSliceV()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, false, slice.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		slice.Append(1)
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []int{1}, slice.O())

		// Second append another which will be 2x at most
		slice.Append(2)
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []int{1, 2}, slice.O())
		assert.Equal(t, NewIntSliceV(1, 2), slice)
	}

	// Start with just appending without chaining
	{
		slice := NewIntSliceV()
		assert.Equal(t, 0, slice.Len())
		slice.Append(1)
		assert.Equal(t, []int{1}, slice.O())
		slice.Append(2)
		assert.Equal(t, []int{1, 2}, slice.O())
	}

	// Start with nil not chained
	{
		slice := NewIntSliceV()
		assert.Equal(t, 0, slice.Len())
		slice.Append(1).Append(2).Append(3)
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []int{1, 2, 3}, slice.O())
	}

	// Start with nil chained
	{
		slice := NewIntSliceV().Append(1).Append(2)
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []int{1, 2}, slice.O())
	}

	// Start with non nil
	{
		slice := NewIntSliceV(1).Append(2).Append(3)
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []int{1, 2, 3}, slice.O())
		assert.Equal(t, NewIntSliceV(1, 2, 3), slice)
	}

	// Use append result directly
	{
		slice := NewIntSliceV(1)
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []int{1, 2}, slice.Append(2).O())
		assert.Equal(t, NewIntSliceV(1, 2), slice)
	}
}

// AppendS
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_AppendS_Go(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkIntSlice_AppendS_Slice(t *testing.B) {
	dest := NewIntSliceV()
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func ExampleIntSlice_AppendS() {
	slice := NewIntSliceV(1).AppendS([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestIntSlice_AppendS(t *testing.T) {

	// nil
	{
		var nilSlice *IntSlice
		assert.Equal(t, NewIntSliceV(1, 2), nilSlice.AppendS([]int{1, 2}))
	}

	// Append many ints
	{
		assert.Equal(t, []int{1, 2, 3}, NewIntSliceV(1).AppendS([]int{2, 3}).O())
		assert.Equal(t, []int{1, 2, 3}, NewIntSlice([]int{1}).AppendS([]int{2, 3}).O())
	}
}

// AppendV
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_AppendV_Go(t *testing.B) {
	ints := []int{}
	ints = append(ints, Range(0, nines6)...)
}

func BenchmarkIntSlice_AppendV_Slice(t *testing.B) {
	n := NewIntSliceV()
	new := rangeO(0, nines6)
	n.AppendV(new...)
}

func ExampleIntSlice_AppendV() {
	slice := NewIntSliceV(1).AppendV(2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestIntSlice_AppendV(t *testing.T) {

	// nil
	{
		var nilSlice *IntSlice
		assert.Equal(t, NewIntSliceV(1, 2), nilSlice.AppendV(1, 2))
	}

	// Append many ints
	{
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1).AppendV(2, 3))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4, 5), NewIntSliceV(1).AppendV(2, 3).AppendV(4, 5))
	}
}

// At
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_At_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := range ints {
		assert.IsType(t, 0, ints[i])
	}
}

func BenchmarkIntSlice_At_Slice(t *testing.B) {
	src := Range(0, nines6)
	slice := NewIntSlice(src)
	for _, i := range src {
		_, ok := (slice.At(i).O()).(int)
		assert.True(t, ok)
	}
}

func ExampleIntSlice_At() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.At(2).O())
	// Output: 3
}

func TestIntSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *IntSlice
		assert.Equal(t, &NObj{nil}, nilSlice.At(0))
	}

	// ints
	{
		slice := NewIntSliceV(1, 2, 3, 4)
		assert.Equal(t, 4, slice.At(-1).O())
		assert.Equal(t, 3, slice.At(-2).O())
		assert.Equal(t, 2, slice.At(-3).O())
		assert.Equal(t, 1, slice.At(0).O())
		assert.Equal(t, 2, slice.At(1).O())
		assert.Equal(t, 3, slice.At(2).O())
		assert.Equal(t, 4, slice.At(3).O())
	}

	// index out of bounds
	{
		slice := NewIntSliceV(1)
		assert.Equal(t, &NObj{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &NObj{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Empty
//--------------------------------------------------------------------------------------------------
func ExampleIntSlice_Empty() {
	fmt.Println(NewIntSliceV().Empty())
	// Output: true
}

func TestIntSlice_Empty(t *testing.T) {

	// nil or empty
	{
		var nilSlice *IntSlice
		assert.Equal(t, true, nilSlice.Empty())
	}

	assert.Equal(t, true, NewIntSliceV().Empty())
	assert.Equal(t, false, NewIntSliceV(1).Empty())
	assert.Equal(t, false, NewIntSliceV(1, 2, 3).Empty())
	assert.Equal(t, false, NewIntSliceV(1).Empty())
	assert.Equal(t, false, NewIntSlice([]int{1, 2, 3}).Empty())
}

// Len
//--------------------------------------------------------------------------------------------------
func ExampleIntSlice_Len() {
	fmt.Println(NewIntSliceV(1, 2, 3).Len())
	// Output: 3
}

func TestIntSlice_Len(t *testing.T) {
	assert.Equal(t, 0, NewIntSliceV().Len())
	assert.Equal(t, 2, len(*(NewIntSliceV(1, 2))))
	assert.Equal(t, 2, NewIntSliceV(1, 2).Len())
}

// Nil
//--------------------------------------------------------------------------------------------------
func ExampleIntSlice_Nil() {
	var slice *IntSlice
	fmt.Println(slice.Nil())
	// Output: true
}

func TestIntSlice_Nil(t *testing.T) {
	var slice *IntSlice
	assert.True(t, slice.Nil())
	assert.False(t, NewIntSliceV().Nil())
	assert.False(t, NewIntSliceV(1, 2, 3).Nil())
}

// O
//--------------------------------------------------------------------------------------------------
func ExampleIntSlice_O() {
	fmt.Println(NewIntSliceV(1, 2, 3).O())
	// Output: [1 2 3]
}

func TestIntSlice_O(t *testing.T) {
	assert.Equal(t, []int{}, NewIntSliceV().O())
	assert.Equal(t, []int{1, 2, 3}, NewIntSliceV(1, 2, 3).O())
}
