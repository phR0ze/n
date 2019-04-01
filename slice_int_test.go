package n

import (
	"fmt"
	"sort"
	"strings"
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
	fmt.Println(slice)
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
	assert.Equal(t, []int{0}, NewIntSlice([]int{0}).O())
	assert.Equal(t, []int{1, 2}, NewIntSlice([]int{1, 2}).O())
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
	fmt.Println(slice)
	// Output: []
}

func ExampleNewIntSliceV_variadic() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice)
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
	assert.False(t, NewIntSliceV(1, 2).Any(Object{2}))

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
	{
		var slice *IntSlice
		assert.False(t, slice.AnyS([]int{1}))
		assert.False(t, NewIntSliceV(1).AnyS(nil))
	}

	// []int
	{
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS([]int{1}))
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS([]int{4, 3}))
		assert.False(t, NewIntSliceV(1, 2, 3).AnyS([]int{4, 5}))
	}

	// *[]int
	{
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(&([]int{1})))
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(&([]int{4, 3})))
		assert.False(t, NewIntSliceV(1, 2, 3).AnyS(&([]int{4, 5})))
	}

	// Slice
	{
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(Slice(NewIntSliceV(1))))
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(Slice(NewIntSliceV(4, 3))))
		assert.False(t, NewIntSliceV(1, 2, 3).AnyS(Slice(NewIntSliceV(4, 5))))
	}

	// IntSlice
	{
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(*NewIntSliceV(1)))
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(*NewIntSliceV(4, 3)))
		assert.False(t, NewIntSliceV(1, 2, 3).AnyS(*NewIntSliceV(4, 5)))
	}

	// *IntSlice
	{
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(NewIntSliceV(1)))
		assert.True(t, NewIntSliceV(1, 2, 3).AnyS(NewIntSliceV(4, 3)))
		assert.False(t, NewIntSliceV(1, 2, 3).AnyS(NewIntSliceV(4, 5)))
	}
	// invalid types
	assert.False(t, NewIntSliceV(1, 2).AnyS([]string{"2"}))
}

// AnyW
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_AnyW_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if i == nines4 {
			break
		}
	}
}

func BenchmarkIntSlice_AnyW_Slice(t *testing.B) {
	src := Range(0, nines5)
	NewIntSlice(src).AnyW(func(x O) bool {
		return ExB(x.(int) == nines4)
	})
}

func ExampleIntSlice_AnyW() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.AnyW(func(x O) bool {
		return ExB(x.(int) == 2)
	}))
	// Output: true
}

func TestIntSlice_AnyW(t *testing.T) {

	// empty
	var slice *IntSlice
	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(int) > 0) }))
	assert.False(t, NewIntSliceV().AnyW(func(x O) bool { return ExB(x.(int) > 0) }))

	// single
	assert.True(t, NewIntSliceV(2).AnyW(func(x O) bool { return ExB(x.(int) > 0) }))

	assert.True(t, NewIntSliceV(1, 2).AnyW(func(x O) bool { return ExB(x.(int) == 2) }))
	assert.True(t, NewIntSliceV(1, 2).AnyW(func(x O) bool { return ExB(x.(int)%2 == 0) }))
	assert.False(t, NewIntSliceV(2, 4).AnyW(func(x O) bool { return ExB(x.(int)%2 != 0) }))
	assert.True(t, NewIntSliceV(1, 2, 3).AnyW(func(x O) bool { return ExB(x.(int) == 4 || x.(int) == 3) }))
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
	fmt.Println(slice)
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
	fmt.Println(slice)
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
	fmt.Println(slice.At(2))
	// Output: 3
}

func TestIntSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *IntSlice
		assert.Equal(t, &Object{nil}, nilSlice.At(0))
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
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------

func ExampleIntSlice_Clear() {
	slice := NewIntSliceV(1).Concat([]int{2, 3})
	fmt.Println(slice.Clear())
	// Output: []
}

func TestIntSlice_Clear(t *testing.T) {

	// nil
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.Clear())
		assert.Equal(t, (*IntSlice)(nil), slice)
	}

	// int
	{
		slice := NewIntSliceV(1, 2, 3, 4)
		assert.Equal(t, NewIntSliceV(), slice.Clear())
		assert.Equal(t, NewIntSliceV(), slice.Clear())
		assert.Equal(t, NewIntSliceV(), slice)
	}
}

// Concat
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Concat_Go(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkIntSlice_Concat_Slice(t *testing.B) {
	dest := NewIntSliceV()
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.Concat(src[j:i])
		j = i
	}
}

func ExampleIntSlice_Concat() {
	slice := NewIntSliceV(1).Concat([]int{2, 3})
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestIntSlice_Concat(t *testing.T) {

	// nil
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(1, 2), slice.Concat([]int{1, 2}))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2).Concat(nil))
	}

	// []int
	{
		slice := NewIntSliceV(1)
		concated := slice.Concat([]int{2, 3})
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), concated)
	}

	// *[]int
	{
		slice := NewIntSliceV(1)
		concated := slice.Concat(&([]int{2, 3}))
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), concated)
	}

	// *IntSlice
	{
		slice := NewIntSliceV(1)
		concated := slice.Concat(NewIntSliceV(2, 3))
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), concated)
	}

	// IntSlice
	{
		slice := NewIntSliceV(1)
		concated := slice.Concat(*NewIntSliceV(2, 3))
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), concated)
	}

	// Slice
	{
		slice := NewIntSliceV(1)
		concated := slice.Concat(Slice(NewIntSliceV(2, 3)))
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), concated)
	}
}

// ConcatM
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_ConcatM_Go(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkIntSlice_ConcatM_Slice(t *testing.B) {
	dest := NewIntSliceV()
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.ConcatM(src[j:i])
		j = i
	}
}

func ExampleIntSlice_ConcatM() {
	slice := NewIntSliceV(1).ConcatM([]int{2, 3})
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestIntSlice_ConcatM(t *testing.T) {

	// nil
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(1, 2), slice.ConcatM([]int{1, 2}))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2).ConcatM(nil))
	}

	// []int
	{
		slice := NewIntSliceV(1)
		concated := slice.ConcatM([]int{2, 3})
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), concated)
	}

	// *[]int
	{
		slice := NewIntSliceV(1)
		concated := slice.ConcatM(&([]int{2, 3}))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), concated)
	}

	// *IntSlice
	{
		slice := NewIntSliceV(1)
		concated := slice.ConcatM(NewIntSliceV(2, 3))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), concated)
	}

	// IntSlice
	{
		slice := NewIntSliceV(1)
		concated := slice.ConcatM(*NewIntSliceV(2, 3))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), concated)
	}

	// Slice
	{
		slice := NewIntSliceV(1)
		concated := slice.ConcatM(Slice(NewIntSliceV(2, 3)))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), concated)
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Copy_Go(t *testing.B) {
	ints := Range(0, nines6)
	dst := make([]int, len(ints), len(ints))
	copy(dst, ints)
}

func BenchmarkIntSlice_Copy_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.Copy()
}

func ExampleIntSlice_Copy() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Copy())
	// Output: [1 2 3]
}

func TestIntSlice_Copy(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.Copy(0, -1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV(0).Clear().Copy(0, -1))
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := NewIntSliceV(1, 2, 3)
		result := original.Copy(0, -1)
		assert.Equal(t, NewIntSliceV(1, 2, 3), original)
		assert.Equal(t, NewIntSliceV(1, 2, 3), result)
		result.Set(0, 0)
		assert.Equal(t, NewIntSliceV(1, 2, 3), original)
		assert.Equal(t, NewIntSliceV(0, 2, 3), result)
	}

	// copy full array
	{
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Copy())
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Copy(0, -1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Copy(0, 1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Copy(0, 5))
		assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1).Copy())
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Copy())
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Copy(0, -1))
		assert.Equal(t, NewIntSlice([]int{1, 2, 3}), NewIntSlice([]int{1, 2, 3}).Copy())
		assert.Equal(t, NewIntSlice([]int{1, 2, 3}), NewIntSlice([]int{1, 2, 3}).Copy(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1).Copy(0, 2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Copy(-6, 6))
	}

	// mutually exclusive
	{
		slice := NewIntSliceV(1, 2, 3, 4)
		assert.Equal(t, NewIntSliceV(), slice.Copy(2, -3))
		assert.Equal(t, NewIntSliceV(), slice.Copy(0, -5))
		assert.Equal(t, NewIntSliceV(), slice.Copy(4, -1))
		assert.Equal(t, NewIntSliceV(), slice.Copy(6, -1))
		assert.Equal(t, NewIntSliceV(), slice.Copy(3, -2))
	}

	// singles
	{
		slice := NewIntSliceV(1, 2, 3, 4)
		assert.Equal(t, NewIntSliceV(4), slice.Copy(-1, -1))
		assert.Equal(t, NewIntSliceV(3), slice.Copy(-2, -2))
		assert.Equal(t, NewIntSliceV(2), slice.Copy(-3, -3))
		assert.Equal(t, NewIntSliceV(1), slice.Copy(0, 0))
		assert.Equal(t, NewIntSliceV(1), slice.Copy(-4, -4))
		assert.Equal(t, NewIntSliceV(2), slice.Copy(1, 1))
		assert.Equal(t, NewIntSliceV(2), slice.Copy(1, -3))
		assert.Equal(t, NewIntSliceV(3), slice.Copy(2, 2))
		assert.Equal(t, NewIntSliceV(3), slice.Copy(2, -2))
		assert.Equal(t, NewIntSliceV(4), slice.Copy(3, 3))
		assert.Equal(t, NewIntSliceV(4), slice.Copy(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Copy(1, -1))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Copy(1, 2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Copy(-2, -1))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Copy(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Copy(0, -2))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Copy(-3, -2))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Copy(-3, 1))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Copy(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Copy(1, -2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Copy(-3, -2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Copy(-3, 2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Copy(1, 2))
	}

	// random
	{
		assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1, 2, 3).Copy(0, -3))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Copy(1, 2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Copy(0, 2))
	}
}

// Count
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Count_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if i == nines4 {
			break
		}
	}
}

func BenchmarkIntSlice_Count_Slice(t *testing.B) {
	src := Range(0, nines5)
	NewIntSlice(src).Count(nines4)
}

func ExampleIntSlice_Count() {
	slice := NewIntSliceV(1, 2, 2)
	fmt.Println(slice.Count(2))
	// Output: 2
}

func TestIntSlice_Count(t *testing.T) {

	// empty
	var slice *IntSlice
	assert.Equal(t, 0, slice.Count(0))
	assert.Equal(t, 0, NewIntSliceV().Count(0))

	assert.Equal(t, 1, NewIntSliceV(2, 3).Count(2))
	assert.Equal(t, 2, NewIntSliceV(1, 2, 2).Count(2))
	assert.Equal(t, 4, NewIntSliceV(4, 4, 3, 4, 4).Count(4))
	assert.Equal(t, 3, NewIntSliceV(3, 2, 3, 3, 5).Count(3))
	assert.Equal(t, 1, NewIntSliceV(1, 2, 3).Count(3))
}

// CountW
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_CountW_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if i == nines4 {
			break
		}
	}
}

func BenchmarkIntSlice_CountW_Slice(t *testing.B) {
	src := Range(0, nines5)
	NewIntSlice(src).CountW(func(x O) bool {
		return ExB(x.(int) == nines4)
	})
}

func ExampleIntSlice_CountW() {
	slice := NewIntSliceV(1, 2, 2)
	fmt.Println(slice.CountW(func(x O) bool {
		return ExB(x.(int) == 2)
	}))
	// Output: 2
}

func TestIntSlice_CountW(t *testing.T) {

	// empty
	var slice *IntSlice
	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(int) > 0) }))
	assert.Equal(t, 0, NewIntSliceV().CountW(func(x O) bool { return ExB(x.(int) > 0) }))

	assert.Equal(t, 1, NewIntSliceV(2, 3).CountW(func(x O) bool { return ExB(x.(int) > 2) }))
	assert.Equal(t, 1, NewIntSliceV(1, 2).CountW(func(x O) bool { return ExB(x.(int) == 2) }))
	assert.Equal(t, 2, NewIntSliceV(1, 2, 3, 4, 5).CountW(func(x O) bool { return ExB(x.(int)%2 == 0) }))
	assert.Equal(t, 3, NewIntSliceV(1, 2, 3, 4, 5).CountW(func(x O) bool { return ExB(x.(int)%2 != 0) }))
	assert.Equal(t, 1, NewIntSliceV(1, 2, 3).CountW(func(x O) bool { return ExB(x.(int) == 4 || x.(int) == 3) }))
}

// Drop
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Drop_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 11 {
		i := 1
		n := 10
		if i+n < len(ints) {
			ints = append(ints[:i], ints[i+n:]...)
		} else {
			ints = ints[:i]
		}
	}
}

func BenchmarkIntSlice_Drop_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 1 {
		slice.Drop(1, 10)
	}
}

func ExampleIntSlice_Drop() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Drop(0, 1))
	// Output: [3]
}

func TestIntSlice_Drop(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, (*IntSlice)(nil), slice.Drop(0, 1))
	}

	// invalid
	assert.Equal(t, NewIntSliceV(1, 2, 3, 4), NewIntSliceV(1, 2, 3, 4).Drop(1))
	assert.Equal(t, NewIntSliceV(1, 2, 3, 4), NewIntSliceV(1, 2, 3, 4).Drop(4, 4))

	// drop 1
	assert.Equal(t, NewIntSliceV(2, 3, 4), NewIntSliceV(1, 2, 3, 4).Drop(0, 0))
	assert.Equal(t, NewIntSliceV(1, 3, 4), NewIntSliceV(1, 2, 3, 4).Drop(1, 1))
	assert.Equal(t, NewIntSliceV(1, 2, 4), NewIntSliceV(1, 2, 3, 4).Drop(2, 2))
	assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3, 4).Drop(3, 3))
	assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3, 4).Drop(-1, -1))
	assert.Equal(t, NewIntSliceV(1, 2, 4), NewIntSliceV(1, 2, 3, 4).Drop(-2, -2))
	assert.Equal(t, NewIntSliceV(1, 3, 4), NewIntSliceV(1, 2, 3, 4).Drop(-3, -3))
	assert.Equal(t, NewIntSliceV(2, 3, 4), NewIntSliceV(1, 2, 3, 4).Drop(-4, -4))

	// drop 2
	assert.Equal(t, NewIntSliceV(3, 4), NewIntSliceV(1, 2, 3, 4).Drop(0, 1))
	assert.Equal(t, NewIntSliceV(1, 4), NewIntSliceV(1, 2, 3, 4).Drop(1, 2))
	assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3, 4).Drop(2, 3))
	assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3, 4).Drop(-2, -1))
	assert.Equal(t, NewIntSliceV(1, 4), NewIntSliceV(1, 2, 3, 4).Drop(-3, -2))
	assert.Equal(t, NewIntSliceV(3, 4), NewIntSliceV(1, 2, 3, 4).Drop(-4, -3))

	// drop 3
	assert.Equal(t, NewIntSliceV(4), NewIntSliceV(1, 2, 3, 4).Drop(0, 2))
	assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1, 2, 3, 4).Drop(-3, -1))

	// drop everything and beyond
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Drop())
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Drop(0, 3))
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Drop(0, -1))
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Drop(-4, -1))
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Drop(-6, -1))
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Drop(0, 10))

	// move index within bounds
	assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3, 4).Drop(3, 4))
	assert.Equal(t, NewIntSliceV(2, 3, 4), NewIntSliceV(1, 2, 3, 4).Drop(-5, 0))
}

// DropAt
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_DropAt_Go(t *testing.B) {
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

func BenchmarkIntSlice_DropAt_Slice(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := NewIntSlice(src)
	for i := range index {
		slice.DropAt(i)
	}
}

func ExampleIntSlice_DropAt() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.DropAt(1))
	// Output: [1 3]
}

func TestIntSlice_DropAt(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, (*IntSlice)(nil), slice.DropAt(0))
	}

	// drop all and more
	{
		slice := NewIntSliceV(0, 1, 2)
		assert.Equal(t, NewIntSliceV(0, 1), slice.DropAt(-1))
		assert.Equal(t, NewIntSliceV(0), slice.DropAt(-1))
		assert.Equal(t, NewIntSliceV(), slice.DropAt(-1))
		assert.Equal(t, NewIntSliceV(), slice.DropAt(-1))
	}

	// drop invalid
	assert.Equal(t, NewIntSliceV(0, 1, 2), NewIntSliceV(0, 1, 2).DropAt(3))
	assert.Equal(t, NewIntSliceV(0, 1, 2), NewIntSliceV(0, 1, 2).DropAt(-4))

	// drop last
	assert.Equal(t, NewIntSliceV(0, 1), NewIntSliceV(0, 1, 2).DropAt(2))
	assert.Equal(t, NewIntSliceV(0, 1), NewIntSliceV(0, 1, 2).DropAt(-1))

	// drop middle
	assert.Equal(t, NewIntSliceV(0, 2), NewIntSliceV(0, 1, 2).DropAt(1))
	assert.Equal(t, NewIntSliceV(0, 2), NewIntSliceV(0, 1, 2).DropAt(-2))

	// drop first
	assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(0, 1, 2).DropAt(0))
	assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(0, 1, 2).DropAt(-3))
}

// DropFirst
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_DropFirst_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkIntSlice_DropFirst_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleIntSlice_DropFirst() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.DropFirst())
	// Output: [2 3]
}

func TestIntSlice_DropFirst(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, (*IntSlice)(nil), slice.DropFirst())
	}

	// drop all and beyond
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(2, 3), slice.DropFirst())
		assert.Equal(t, NewIntSliceV(3), slice.DropFirst())
		assert.Equal(t, NewIntSliceV(), slice.DropFirst())
		assert.Equal(t, NewIntSliceV(), slice.DropFirst())
	}
}

// DropFirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_DropFirstN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkIntSlice_DropFirstN_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func ExampleIntSlice_DropFirstN() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.DropFirstN(2))
	// Output: [3]
}

func TestIntSlice_DropFirstN(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, (*IntSlice)(nil), slice.DropFirstN(1))
	}

	// negative value
	assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).DropFirstN(-1))

	// drop none
	assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).DropFirstN(0))

	// drop 1
	assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).DropFirstN(1))

	// drop 2
	assert.Equal(t, NewIntSliceV(3), NewIntSliceV(1, 2, 3).DropFirstN(2))

	// drop 3
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3).DropFirstN(3))

	// drop beyond
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3).DropFirstN(4))
}

// DropLast
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_DropLast_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkIntSlice_DropLast_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func ExampleIntSlice_DropLast() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.DropLast())
	// Output: [1 2]
}

func TestIntSlice_DropLast(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, (*IntSlice)(nil), slice.DropLast())
	}

	// negative value
	assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).DropLastN(-1))

	slice := NewIntSliceV(1, 2, 3)
	assert.Equal(t, NewIntSliceV(1, 2), slice.DropLast())
	assert.Equal(t, NewIntSliceV(1), slice.DropLast())
	assert.Equal(t, NewIntSliceV(), slice.DropLast())
	assert.Equal(t, NewIntSliceV(), slice.DropLast())
}

// DropLastN
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_DropLastN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkIntSlice_DropLastN_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func ExampleIntSlice_DropLastN() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.DropLastN(2))
	// Output: [1]
}

func TestIntSlice_DropLastN(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, (*IntSlice)(nil), slice.DropLastN(1))
	}

	// drop none
	assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).DropLastN(0))

	// drop 1
	assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).DropLastN(1))

	// drop 2
	assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1, 2, 3).DropLastN(2))

	// drop 3
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3).DropLastN(3))

	// drop beyond
	assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3).DropLastN(4))
}

// DropW
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_DropW_Go(t *testing.B) {
	ints := Range(0, nines5)
	l := len(ints)
	for i := 0; i < l; i++ {
		if ints[i]%2 == 0 {
			if i+1 < l {
				ints = append(ints[:i], ints[i+1:]...)
			} else if i >= 0 && i < l {
				ints = ints[:i]
			}
			l--
			i--
		}
	}
}

func BenchmarkIntSlice_DropW_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines5))
	slice.DropW(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	})
}

func ExampleIntSlice_DropW() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.DropW(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	}))
	// Output: [1 3]
}

func TestIntSlice_DropW(t *testing.T) {

	// drop all odd values
	{
		slice := NewIntSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		slice.DropW(func(x O) bool {
			return ExB(x.(int)%2 != 0)
		})
		assert.Equal(t, NewIntSliceV(2, 4, 6, 8), slice)
	}

	// drop all even values
	{
		slice := NewIntSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		slice.DropW(func(x O) bool {
			return ExB(x.(int)%2 == 0)
		})
		assert.Equal(t, NewIntSliceV(1, 3, 5, 7, 9), slice)
	}
}

// Each
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Each_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkIntSlice_Each_Slice(t *testing.B) {
	NewIntSlice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func ExampleIntSlice_Each() {
	NewIntSliceV(1, 2, 3).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 123
}

func TestIntSlice_Each(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		slice.Each(func(x O) {})
	}

	NewIntSliceV(1, 2, 3).Each(func(x O) {
		switch x {
		case 1:
			assert.Equal(t, 1, x)
		case 2:
			assert.Equal(t, 2, x)
		case 3:
			assert.Equal(t, 3, x)
		}
	})
}

// EachE
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_EachE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkIntSlice_EachE_Slice(t *testing.B) {
	NewIntSlice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func ExampleIntSlice_EachE() {
	NewIntSliceV(1, 2, 3).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 123
}

func TestIntSlice_EachE(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		slice.EachE(func(x O) error {
			return nil
		})
	}

	NewIntSliceV(1, 2, 3).EachE(func(x O) error {
		switch x {
		case 1:
			assert.Equal(t, 1, x)
		case 2:
			assert.Equal(t, 2, x)
		case 3:
			assert.Equal(t, 3, x)
		}
		return nil
	})
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

// First
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_First_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		_ = ints[0]
		ints = ints[1:]
	}
}

func BenchmarkIntSlice_First_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.First()
		slice.DropFirst()
	}
}

func ExampleIntSlice_First() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.First())
	// Output: 1
}

func TestIntSlice_First(t *testing.T) {
	// invalid
	assert.Equal(t, &Object{nil}, NewIntSliceV().First())

	// int
	assert.Equal(t, &Object{2}, NewIntSliceV(2, 3).First())
	assert.Equal(t, &Object{3}, NewIntSliceV(3, 2).First())
	assert.Equal(t, &Object{1}, NewIntSliceV(1, 3, 2).First())
}

// FirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_FirstN_Go(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:10]
}

func BenchmarkIntSlice_FirstN_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	slice.FirstN(10)
}

func ExampleIntSlice_FirstN() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.FirstN(2))
	// Output: [1 2]
}

func TestIntSlice_FirstN(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.FirstN(1))
		assert.Equal(t, NewIntSliceV(), slice.FirstN(-1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewIntSliceV(1, 2, 3)
		result := original.FirstN(2).Set(0, 0)
		assert.Equal(t, NewIntSliceV(0, 2, 3), original)
		assert.Equal(t, NewIntSliceV(0, 2), result)
	}

	// slice full array includeing out of bounds
	assert.Equal(t, NewIntSliceV(), NewIntSliceV().FirstN(1))
	assert.Equal(t, NewIntSliceV(), NewIntSliceV().FirstN(10))
	assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).FirstN(10))
	assert.Equal(t, NewIntSlice([]int{1, 2, 3}), NewIntSlice([]int{1, 2, 3}).FirstN(10))

	// grab a few diff
	assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1, 2, 3).FirstN(1))
	assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).FirstN(2))
}

// Index
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Index_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if ints[i] == nines4 {
			break
		}
	}
}

func BenchmarkIntSlice_Index_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines5))
	slice.Index(nines4)
}

func ExampleIntSlice_Index() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Index(2))
	// Output: 1
}

func TestIntSlice_Index(t *testing.T) {

	// empty
	var slice *IntSlice
	assert.Equal(t, -1, slice.Index(2))
	assert.Equal(t, -1, NewIntSliceV().Index(1))

	assert.Equal(t, 0, NewIntSliceV(1, 2, 3).Index(1))
	assert.Equal(t, 1, NewIntSliceV(1, 2, 3).Index(2))
	assert.Equal(t, 2, NewIntSliceV(1, 2, 3).Index(3))
	assert.Equal(t, -1, NewIntSliceV(1, 2, 3).Index(4))
	assert.Equal(t, -1, NewIntSliceV(1, 2, 3).Index(5))
}

// Insert
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Insert_Go(t *testing.B) {
	ints := []int{}
	for i := range Range(0, nines6) {
		ints = append(ints, i)
		copy(ints[1:], ints[1:])
		ints[0] = i
	}
}

func BenchmarkIntSlice_Insert_Slice(t *testing.B) {
	slice := NewIntSliceV()
	for i := range Range(0, nines6) {
		slice.Insert(0, i)
	}
}

func ExampleIntSlice_Insert() {
	slice := NewIntSliceV(1, 3)
	fmt.Println(slice.Insert(1, 2))
	// Output: [1 2 3]
}

func TestIntSlice_Insert(t *testing.T) {

	// append
	{
		slice := NewIntSliceV()
		assert.Equal(t, NewIntSliceV(0), slice.Insert(-1, 0))
		assert.Equal(t, NewIntSliceV(0, 1), slice.Insert(-1, 1))
		assert.Equal(t, NewIntSliceV(0, 1, 2), slice.Insert(-1, 2))
	}

	// prepend
	{
		slice := NewIntSliceV()
		assert.Equal(t, NewIntSliceV(2), slice.Insert(0, 2))
		assert.Equal(t, NewIntSliceV(1, 2), slice.Insert(0, 1))
		assert.Equal(t, NewIntSliceV(0, 1, 2), slice.Insert(0, 0))
	}

	// middle pos
	{
		slice := NewIntSliceV(0, 5)
		assert.Equal(t, NewIntSliceV(0, 1, 5), slice.Insert(1, 1))
		assert.Equal(t, NewIntSliceV(0, 1, 2, 5), slice.Insert(2, 2))
		assert.Equal(t, NewIntSliceV(0, 1, 2, 3, 5), slice.Insert(3, 3))
		assert.Equal(t, NewIntSliceV(0, 1, 2, 3, 4, 5), slice.Insert(4, 4))
	}

	// middle neg
	{
		slice := NewIntSliceV(0, 5)
		assert.Equal(t, NewIntSliceV(0, 1, 5), slice.Insert(-2, 1))
		assert.Equal(t, NewIntSliceV(0, 1, 2, 5), slice.Insert(-2, 2))
		assert.Equal(t, NewIntSliceV(0, 1, 2, 3, 5), slice.Insert(-2, 3))
		assert.Equal(t, NewIntSliceV(0, 1, 2, 3, 4, 5), slice.Insert(-2, 4))
	}

	// error cases
	{
		var slice *NSlice
		assert.True(t, slice.Insert(0, 0).Nil())
		assert.Equal(t, (*NSlice)(nil), slice.Insert(0, 0))
		assert.Equal(t, NewIntSliceV(0, 1), NewIntSliceV(0, 1).Insert(-10, 1))
		assert.Equal(t, NewIntSliceV(0, 1), NewIntSliceV(0, 1).Insert(10, 1))
		assert.Equal(t, NewIntSliceV(0, 1), NewIntSliceV(0, 1).Insert(2, 1))
		assert.Equal(t, NewIntSliceV(0, 1), NewIntSliceV(0, 1).Insert(-3, 1))
	}
}

// Join
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Join_Go(t *testing.B) {
	ints := Range(0, nines4)
	strs := []string{}
	for i := 0; i < len(ints); i++ {
		strs = append(strs, fmt.Sprintf("%v", ints[i]))
	}
	strings.Join(strs, ",")
}

func BenchmarkIntSlice_Join_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines4))
	slice.Join()
}

func ExampleIntSlice_Join() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Join())
	// Output: 1,2,3
}

func TestIntSlice_Join(t *testing.T) {
	// nil
	{
		var slice *IntSlice
		assert.Equal(t, &Object{""}, slice.Join())
	}

	// empty
	{
		assert.Equal(t, &Object{""}, NewIntSliceV().Join())
	}

	assert.Equal(t, "1,2,3", NewIntSliceV(1, 2, 3).Join().O())
	assert.Equal(t, "1.2.3", NewIntSliceV(1, 2, 3).Join(".").O())
}

// Last
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Last_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		_ = ints[len(ints)-1]
		ints = ints[:len(ints)-1]
	}
}

func BenchmarkIntSlice_Last_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Last()
		slice.DropLast()
	}
}

func ExampleIntSlice_Last() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Last())
	// Output: 3
}

func TestIntSlice_Last(t *testing.T) {
	// invalid
	assert.Equal(t, &Object{nil}, NewIntSliceV().Last())

	// int
	assert.Equal(t, &Object{3}, NewIntSliceV(2, 3).Last())
	assert.Equal(t, &Object{2}, NewIntSliceV(3, 2).Last())
	assert.Equal(t, &Object{2}, NewIntSliceV(1, 3, 2).Last())
}

// LastN
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_LastN_Go(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:10]
}

func BenchmarkIntSlice_LastN_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	slice.LastN(10)
}

func ExampleIntSlice_LastN() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.LastN(2))
	// Output: [2 3]
}

func TestIntSlice_LastN(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.LastN(1))
		assert.Equal(t, NewIntSliceV(), slice.LastN(-1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewIntSliceV(1, 2, 3)
		result := original.LastN(2).Set(0, 0)
		assert.Equal(t, NewIntSliceV(1, 0, 3), original)
		assert.Equal(t, NewIntSliceV(0, 3), result)
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().LastN(1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().LastN(10))
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).LastN(10))
		assert.Equal(t, NewIntSlice([]int{1, 2, 3}), NewIntSlice([]int{1, 2, 3}).LastN(10))
	}

	// grab a few diff
	{
		assert.Equal(t, NewIntSliceV(3), NewIntSliceV(1, 2, 3).LastN(1))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).LastN(2))
	}
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

// Less
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Less_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		if i+1 < len(ints) {
			_ = ints[i] < ints[i+1]
		}
	}
}

func BenchmarkIntSlice_Less_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Less(i, i+1)
		}
	}
}

func ExampleIntSlice_Less() {
	slice := NewIntSliceV(2, 3, 1)
	fmt.Println(slice.Less(0, 2))
	// Output: false
}

func TestIntSlice_Less(t *testing.T) {

	// invalid cases
	{
		var slice *IntSlice
		assert.False(t, slice.Less(0, 0))

		slice = NewIntSliceV()
		assert.False(t, slice.Less(0, 0))
		assert.False(t, slice.Less(1, 2))
		assert.False(t, slice.Less(-1, 2))
		assert.False(t, slice.Less(1, -2))
	}

	// valid
	assert.Equal(t, true, NewIntSliceV(0, 1, 2).Less(0, 1))
	assert.Equal(t, false, NewIntSliceV(0, 1, 2).Less(1, 0))
	assert.Equal(t, true, NewIntSliceV(0, 1, 2).Less(1, 2))
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
	fmt.Println(NewIntSliceV(1, 2, 3))
	// Output: [1 2 3]
}

func TestIntSlice_O(t *testing.T) {
	assert.Equal(t, []int{}, NewIntSliceV().O())
	assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3))
}

// Pair
//--------------------------------------------------------------------------------------------------

func ExampleIntSlice_Pair() {
	slice := NewIntSliceV(1, 2)
	first, second := slice.Pair()
	fmt.Println(first, second)
	// Output: 1 2
}

func TestIntSlice_Pair(t *testing.T) {

	// two values
	{
		first, second := NewIntSliceV(1, 2).Pair()
		assert.Equal(t, &Object{1}, first)
		assert.Equal(t, &Object{2}, second)
	}

	// one value
	{
		first, second := NewIntSliceV(1).Pair()
		assert.Equal(t, &Object{1}, first)
		assert.Equal(t, &Object{nil}, second)
	}

	// no values
	{
		first, second := NewIntSliceV().Pair()
		assert.Equal(t, &Object{nil}, first)
		assert.Equal(t, &Object{nil}, second)
	}
}

// Pop
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Pop_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkIntSlice_Pop_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Pop()
	}
}

func ExampleIntSlice_Pop() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Pop())
	// Output: 3
}

func TestIntSlice_Pop(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, &Object{nil}, slice.Pop())
	}

	// take all one at a time
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, &Object{3}, slice.Pop())
		assert.Equal(t, NewIntSliceV(1, 2), slice)
		assert.Equal(t, &Object{2}, slice.Pop())
		assert.Equal(t, NewIntSliceV(1), slice)
		assert.Equal(t, &Object{1}, slice.Pop())
		assert.Equal(t, NewIntSliceV(), slice)
		assert.Equal(t, &Object{nil}, slice.Pop())
		assert.Equal(t, NewIntSliceV(), slice)
	}
}

// PopN
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_PopN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkIntSlice_PopN_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.PopN(10)
	}
}

func ExampleIntSlice_PopN() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.PopN(2))
	// Output: [2 3]
}

func TestIntSlice_PopN(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.PopN(1))
	}

	// take none
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(), slice.PopN(0))
		assert.Equal(t, NewIntSliceV(1, 2, 3), slice)
	}

	// take 1
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(3), slice.PopN(1))
		assert.Equal(t, NewIntSliceV(1, 2), slice)
	}

	// take 2
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(2, 3), slice.PopN(2))
		assert.Equal(t, NewIntSliceV(1), slice)
	}

	// take 3
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(1, 2, 3), slice.PopN(3))
		assert.Equal(t, NewIntSliceV(), slice)
	}

	// take beyond
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(1, 2, 3), slice.PopN(4))
		assert.Equal(t, NewIntSliceV(), slice)
	}
}

// Prepend
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Prepend_Go(t *testing.B) {
	ints := []int{}
	for i := range Range(0, nines6) {
		ints = append(ints, i)
		copy(ints[1:], ints[1:])
		ints[0] = i
	}
}

func BenchmarkIntSlice_Prepend_Slice(t *testing.B) {
	slice := NewIntSliceV()
	for i := range Range(0, nines6) {
		slice.Prepend(i)
	}
}

func ExampleIntSlice_Prepend() {
	slice := NewIntSliceV(2, 3)
	fmt.Println(slice.Prepend(1))
	// Output: [1 2 3]
}

func TestIntSlice_Prepend(t *testing.T) {

	// happy path
	{
		slice := NewIntSliceV()
		assert.Equal(t, NewIntSliceV(2), slice.Prepend(2))
		assert.Equal(t, NewIntSliceV(1, 2), slice.Prepend(1))
		assert.Equal(t, NewIntSliceV(0, 1, 2), slice.Prepend(0))
	}

	// error cases
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(0), slice.Prepend(0))
	}
}

// Reverse
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Reverse_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

func BenchmarkIntSlice_Reverse_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.Reverse()
}

func ExampleIntSlice_Reverse() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Reverse())
	// Output: [3 2 1]
}

func TestIntSlice_Reverse(t *testing.T) {

	// nil
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.Reverse())
	}

	// empty
	{
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Reverse())
	}

	// pos
	{
		slice := NewIntSliceV(3, 2, 1)
		reversed := slice.Reverse()
		assert.Equal(t, NewIntSliceV(3, 2, 1, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3), reversed)
	}

	// neg
	{
		slice := NewIntSliceV(2, 3, -2, -3)
		reversed := slice.Reverse()
		assert.Equal(t, NewIntSliceV(2, 3, -2, -3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(-3, -2, 3, 2), reversed)
	}
}

// ReverseM
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_ReverseM_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

func BenchmarkIntSlice_ReverseM_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.ReverseM()
}

func ExampleIntSlice_ReverseM() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.ReverseM())
	// Output: [3 2 1]
}

func TestIntSlice_ReverseM(t *testing.T) {

	// nil
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.Reverse())
	}

	// empty
	{
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Reverse())
	}

	// pos
	{
		slice := NewIntSliceV(3, 2, 1)
		reversed := slice.ReverseM()
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), reversed)
	}

	// neg
	{
		slice := NewIntSliceV(2, 3, -2, -3)
		reversed := slice.ReverseM()
		assert.Equal(t, NewIntSliceV(-3, -2, 3, 2, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(-3, -2, 3, 2, 4), reversed)
	}
}

// Select
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Select_Go(t *testing.B) {
	even := []int{}
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		if ints[i]%2 == 0 {
			even = append(even, ints[i])
		}
	}
}

func BenchmarkIntSlice_Select_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.Select(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	})

}

func ExampleIntSlice_Select() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Select(func(x O) bool {
		return ExB(x.(int) == 2 || x.(int) == 3)
	}))
	// Output: [2 3]
}

func TestIntSlice_Select(t *testing.T) {

	// Select all odd values
	{
		slice := NewIntSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.Select(func(x O) bool {
			return ExB(x.(int)%2 != 0)
		})
		slice.DropFirst()
		assert.Equal(t, NewIntSliceV(2, 3, 4, 5, 6, 7, 8, 9), slice)
		assert.Equal(t, NewIntSliceV(1, 3, 5, 7, 9), new)
	}

	// Select all even values
	{
		slice := NewIntSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.Select(func(x O) bool {
			return ExB(x.(int)%2 == 0)
		})
		slice.DropAt(1)
		assert.Equal(t, NewIntSliceV(1, 3, 4, 5, 6, 7, 8, 9), slice)
		assert.Equal(t, NewIntSliceV(2, 4, 6, 8), new)
	}
}

// Set
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Set_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		ints[i] = 0
	}
}

func BenchmarkIntSlice_Set_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, 0)
	}
}

func ExampleIntSlice_Set() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Set(0, 0))
	// Output: [0 2 3]
}

func TestIntSlice_Set(t *testing.T) {
	assert.Equal(t, NewIntSliceV(0, 2, 3), NewIntSliceV(1, 2, 3).Set(0, 0))
	assert.Equal(t, NewIntSliceV(1, 0, 3), NewIntSliceV(1, 2, 3).Set(1, 0))
	assert.Equal(t, NewIntSliceV(1, 2, 0), NewIntSliceV(1, 2, 3).Set(2, 0))
	assert.Equal(t, NewIntSliceV(0, 2, 3), NewIntSliceV(1, 2, 3).Set(-3, 0))
	assert.Equal(t, NewIntSliceV(1, 0, 3), NewIntSliceV(1, 2, 3).Set(-2, 0))
	assert.Equal(t, NewIntSliceV(1, 2, 0), NewIntSliceV(1, 2, 3).Set(-1, 0))

	// Test out of bounds
	{
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Set(5, 1))
		slice, err := NewIntSliceV(1, 2, 3).SetE(5, 1)
		assert.NotNil(t, slice)
		assert.NotNil(t, err)
	}

	// Test wrong type
	{
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Set(5, "1"))
		slice, err := NewIntSliceV(1, 2, 3).SetE(5, "1")
		assert.NotNil(t, slice)
		assert.NotNil(t, err)
	}
}

// Single
//--------------------------------------------------------------------------------------------------

func ExampleIntSlice_Single() {
	slice := NewIntSliceV(1)
	fmt.Println(slice.Single())
	// Output: true
}

func TestIntSlice_Single(t *testing.T) {

	assert.Equal(t, false, NewIntSliceV().Single())
	assert.Equal(t, true, NewIntSliceV(1).Single())
	assert.Equal(t, false, NewIntSliceV(1, 2).Single())
}

// Slice
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Slice_Go(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:len(ints)]
}

func BenchmarkIntSlice_Slice_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	slice.Slice(0, -1)
}

func ExampleIntSlice_Slice() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Slice(1, -1))
	// Output: [2 3]
}

func TestIntSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.Slice(0, -1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Slice(0, -1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewIntSliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, NewIntSliceV(0, 2, 3), original)
		assert.Equal(t, NewIntSliceV(0, 2, 3), result)
	}

	// slice full array
	{
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Slice(0, -1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Slice(0, 1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV().Slice(0, 5))
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Slice(0, -1))
		assert.Equal(t, NewIntSlice([]int{1, 2, 3}), NewIntSlice([]int{1, 2, 3}).Slice(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1).Slice(0, 2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Slice(-6, 6))
	}

	// mutually exclusive
	{
		assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Slice(2, -3))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Slice(0, -5))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Slice(4, -1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Slice(6, -1))
		assert.Equal(t, NewIntSliceV(), NewIntSliceV(1, 2, 3, 4).Slice(3, 2))
	}

	// singles
	{
		slice := NewIntSliceV(1, 2, 3, 4)
		assert.Equal(t, NewIntSliceV(4), slice.Slice(-1, -1))
		assert.Equal(t, NewIntSliceV(3), slice.Slice(-2, -2))
		assert.Equal(t, NewIntSliceV(2), slice.Slice(-3, -3))
		assert.Equal(t, NewIntSliceV(1), slice.Slice(0, 0))
		assert.Equal(t, NewIntSliceV(1), slice.Slice(-4, -4))
		assert.Equal(t, NewIntSliceV(2), slice.Slice(1, 1))
		assert.Equal(t, NewIntSliceV(2), slice.Slice(1, -3))
		assert.Equal(t, NewIntSliceV(3), slice.Slice(2, 2))
		assert.Equal(t, NewIntSliceV(3), slice.Slice(2, -2))
		assert.Equal(t, NewIntSliceV(4), slice.Slice(3, 3))
		assert.Equal(t, NewIntSliceV(4), slice.Slice(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Slice(1, -1))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Slice(-2, -1))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Slice(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Slice(0, -2))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Slice(-3, -2))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Slice(-3, 1))
		assert.Equal(t, NewIntSliceV(1, 2), NewIntSliceV(1, 2, 3).Slice(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Slice(1, -2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Slice(-3, -2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Slice(-3, 2))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3, 4).Slice(1, 2))
	}

	// random
	{
		assert.Equal(t, NewIntSliceV(1), NewIntSliceV(1, 2, 3).Slice(0, -3))
		assert.Equal(t, NewIntSliceV(2, 3), NewIntSliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, NewIntSliceV(1, 2, 3), NewIntSliceV(1, 2, 3).Slice(0, 2))
	}
}

// Sort
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Sort_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.IntSlice(ints))
}

func BenchmarkIntSlice_Sort_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.Sort()
}

func ExampleIntSlice_Sort() {
	slice := NewIntSliceV(2, 3, 1)
	fmt.Println(slice.Sort())
	// Output: [1 2 3]
}

func TestIntSlice_Sort(t *testing.T) {

	// empty
	assert.Equal(t, NewIntSliceV(), NewIntSliceV().Sort())

	// pos
	{
		slice := NewIntSliceV(5, 3, 2, 4, 1)
		sorted := slice.Sort()
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 3, 2, 4, 1), slice)
	}

	// neg
	{
		slice := NewIntSliceV(5, 3, -2, 4, -1)
		sorted := slice.Sort()
		assert.Equal(t, NewIntSliceV(-2, -1, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 3, -2, 4, -1), slice)
	}
}

// SortM
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_SortM_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.IntSlice(ints))
}

func BenchmarkIntSlice_SortM_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.SortM()
}

func ExampleIntSlice_SortM() {
	slice := NewIntSliceV(2, 3, 1)
	fmt.Println(slice.SortM())
	// Output: [1 2 3]
}

func TestIntSlice_SortM(t *testing.T) {

	// empty
	assert.Equal(t, NewIntSliceV(), NewIntSliceV().SortM())

	// pos
	{
		slice := NewIntSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortM()
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4, 5, 6), slice)
	}

	// neg
	{
		slice := NewIntSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortM()
		assert.Equal(t, NewIntSliceV(-2, -1, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(-2, -1, 3, 4, 5, 6), slice)
	}
}

// SortReverse
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_SortReverse_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}

func BenchmarkIntSlice_SortReverse_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.SortReverse()
}

func ExampleIntSlice_SortReverse() {
	slice := NewIntSliceV(2, 3, 1)
	fmt.Println(slice.SortReverse())
	// Output: [3 2 1]
}

func TestIntSlice_SortReverse(t *testing.T) {

	// empty
	assert.Equal(t, NewIntSliceV(), NewIntSliceV().SortReverse())

	// pos
	{
		slice := NewIntSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverse()
		assert.Equal(t, NewIntSliceV(5, 4, 3, 2, 1, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 3, 2, 4, 1), slice)
	}

	// neg
	{
		slice := NewIntSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverse()
		assert.Equal(t, NewIntSliceV(5, 4, 3, -1, -2, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 3, -2, 4, -1), slice)
	}
}

// SortReverseM
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_SortReverseM_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}

func BenchmarkIntSlice_SortReverseM_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.SortReverseM()
}

func ExampleIntSlice_SortReverseM() {
	slice := NewIntSliceV(2, 3, 1)
	fmt.Println(slice.SortReverseM())
	// Output: [3 2 1]
}

func TestIntSlice_SortReverseM(t *testing.T) {

	// empty
	assert.Equal(t, NewIntSliceV(), NewIntSliceV().SortReverse())

	// pos
	{
		slice := NewIntSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverseM()
		assert.Equal(t, NewIntSliceV(5, 4, 3, 2, 1, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 4, 3, 2, 1, 6), slice)
	}

	// neg
	{
		slice := NewIntSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverseM()
		assert.Equal(t, NewIntSliceV(5, 4, 3, -1, -2, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 4, 3, -1, -2, 6), slice)
	}
}

// String
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_String_Go(t *testing.B) {
	ints := Range(0, nines6)
	_ = fmt.Sprintf("%v", ints)
}

func BenchmarkIntSlice_Stringj_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	_ = slice.String()
}

func ExampleIntSlice_String() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestIntSlice_String(t *testing.T) {

	// empty
	assert.Equal(t, NewIntSliceV(), NewIntSliceV().SortReverse())

	// pos
	{
		slice := NewIntSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverseM()
		assert.Equal(t, NewIntSliceV(5, 4, 3, 2, 1, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 4, 3, 2, 1, 6), slice)
	}

	// neg
	{
		slice := NewIntSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverseM()
		assert.Equal(t, NewIntSliceV(5, 4, 3, -1, -2, 6), sorted.Append(6))
		assert.Equal(t, NewIntSliceV(5, 4, 3, -1, -2, 6), slice)
	}
}

// Swap
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Swap_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		if i+1 < len(ints) {
			ints[i], ints[i+1] = ints[i+1], ints[i]
		}
	}
}

func BenchmarkIntSlice_Swap_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Swap(i, i+1)
		}
	}
}

func ExampleIntSlice_Swap() {
	slice := NewIntSliceV(2, 3, 1)
	slice.Swap(0, 2)
	slice.Swap(1, 2)
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestIntSlice_Swap(t *testing.T) {

	// invalid cases
	{
		var slice *IntSlice
		slice.Swap(0, 0)
		assert.Equal(t, (*IntSlice)(nil), slice)

		slice = NewIntSliceV()
		slice.Swap(0, 0)
		assert.Equal(t, NewIntSliceV(), slice)

		slice.Swap(1, 2)
		assert.Equal(t, NewIntSliceV(), slice)

		slice.Swap(-1, 2)
		assert.Equal(t, NewIntSliceV(), slice)

		slice.Swap(1, -2)
		assert.Equal(t, NewIntSliceV(), slice)
	}

	// normal
	{
		slice := NewIntSliceV(0, 1, 2)
		slice.Swap(0, 1)
		assert.Equal(t, NewIntSliceV(1, 0, 2), slice)
	}
}

// Take
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Take_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 11 {
		i := 1
		n := 10
		if i+n < len(ints) {
			ints = append(ints[:i], ints[i+n:]...)
		} else {
			ints = ints[:i]
		}
	}
}

func BenchmarkIntSlice_Take_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 1 {
		slice.Take(1, 10)
	}
}

func ExampleIntSlice_Take() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Take(0, 1))
	// Output: [1 2]
}

func TestIntSlice_Take(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.Take(0, 1))
	}

	// invalid
	{
		slice := NewIntSliceV(1, 2, 3, 4)
		assert.Equal(t, NewIntSliceV(), slice.Take(1))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice)
		assert.Equal(t, NewIntSliceV(), slice.Take(4, 4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice)
	}

	// take 1
	{
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1), slice.Take(0, 0))
			assert.Equal(t, NewIntSliceV(2, 3, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(2), slice.Take(1, 1))
			assert.Equal(t, NewIntSliceV(1, 3, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(3), slice.Take(2, 2))
			assert.Equal(t, NewIntSliceV(1, 2, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(4), slice.Take(3, 3))
			assert.Equal(t, NewIntSliceV(1, 2, 3), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(4), slice.Take(-1, -1))
			assert.Equal(t, NewIntSliceV(1, 2, 3), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(3), slice.Take(-2, -2))
			assert.Equal(t, NewIntSliceV(1, 2, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(2), slice.Take(-3, -3))
			assert.Equal(t, NewIntSliceV(1, 3, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1), slice.Take(-4, -4))
			assert.Equal(t, NewIntSliceV(2, 3, 4), slice)
		}
	}

	// take 2
	{
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2), slice.Take(0, 1))
			assert.Equal(t, NewIntSliceV(3, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(2, 3), slice.Take(1, 2))
			assert.Equal(t, NewIntSliceV(1, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(3, 4), slice.Take(2, 3))
			assert.Equal(t, NewIntSliceV(1, 2), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(3, 4), slice.Take(-2, -1))
			assert.Equal(t, NewIntSliceV(1, 2), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(2, 3), slice.Take(-3, -2))
			assert.Equal(t, NewIntSliceV(1, 4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2), slice.Take(-4, -3))
			assert.Equal(t, NewIntSliceV(3, 4), slice)
		}
	}

	// take 3
	{
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2, 3), slice.Take(0, 2))
			assert.Equal(t, NewIntSliceV(4), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(2, 3, 4), slice.Take(-3, -1))
			assert.Equal(t, NewIntSliceV(1), slice)
		}
	}

	// take everything and beyond
	{
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Take())
			assert.Equal(t, NewIntSliceV(), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Take(0, 3))
			assert.Equal(t, NewIntSliceV(), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Take(0, -1))
			assert.Equal(t, NewIntSliceV(), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Take(-4, -1))
			assert.Equal(t, NewIntSliceV(), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Take(-6, -1))
			assert.Equal(t, NewIntSliceV(), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Take(0, 10))
			assert.Equal(t, NewIntSliceV(), slice)
		}
	}

	// move index within bounds
	{
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(4), slice.Take(3, 4))
			assert.Equal(t, NewIntSliceV(1, 2, 3), slice)
		}
		{
			slice := NewIntSliceV(1, 2, 3, 4)
			assert.Equal(t, NewIntSliceV(1), slice.Take(-5, 0))
			assert.Equal(t, NewIntSliceV(2, 3, 4), slice)
		}
	}
}

// TakeAt
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_TakeAt_Go(t *testing.B) {
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

func BenchmarkIntSlice_TakeAt_Slice(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := NewIntSlice(src)
	for i := range index {
		slice.TakeAt(i)
	}
}

func ExampleIntSlice_TakeAt() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.TakeAt(1))
	// Output: 2
}

func TestIntSlice_TakeAt(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, &Object{nil}, slice.TakeAt(0))
	}

	// all and more
	{
		slice := NewIntSliceV(0, 1, 2)
		assert.Equal(t, &Object{2}, slice.TakeAt(-1))
		assert.Equal(t, NewIntSliceV(0, 1), slice)
		assert.Equal(t, &Object{1}, slice.TakeAt(-1))
		assert.Equal(t, NewIntSliceV(0), slice)
		assert.Equal(t, &Object{0}, slice.TakeAt(-1))
		assert.Equal(t, NewIntSliceV(), slice)
		assert.Equal(t, &Object{nil}, slice.TakeAt(-1))
		assert.Equal(t, NewIntSliceV(), slice)
	}

	// take invalid
	{
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{nil}, slice.TakeAt(3))
			assert.Equal(t, NewIntSliceV(0, 1, 2), slice)
		}
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{nil}, slice.TakeAt(-4))
			assert.Equal(t, NewIntSliceV(0, 1, 2), slice)
		}
	}

	// take last
	{
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{2}, slice.TakeAt(2))
			assert.Equal(t, NewIntSliceV(0, 1), slice)
		}
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{2}, slice.TakeAt(-1))
			assert.Equal(t, NewIntSliceV(0, 1), slice)
		}
	}

	// take middle
	{
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{1}, slice.TakeAt(1))
			assert.Equal(t, NewIntSliceV(0, 2), slice)
		}
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{1}, slice.TakeAt(-2))
			assert.Equal(t, NewIntSliceV(0, 2), slice)
		}
	}

	// take first
	{
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{0}, slice.TakeAt(0))
			assert.Equal(t, NewIntSliceV(1, 2), slice)
		}
		{
			slice := NewIntSliceV(0, 1, 2)
			assert.Equal(t, &Object{0}, slice.TakeAt(-3))
			assert.Equal(t, NewIntSliceV(1, 2), slice)
		}
	}
}

// TakeFirst
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_TakeFirst_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkIntSlice_TakeFirst_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Shift()
	}
}

func ExampleIntSlice_TakeFirst() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.Shift())
	// Output: 1
}

func TestIntSlice_TakeFirst(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, &Object{nil}, slice.Shift())
	}

	// take all and beyond
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, &Object{1}, slice.Shift())
		assert.Equal(t, NewIntSliceV(2, 3), slice)
		assert.Equal(t, &Object{2}, slice.Shift())
		assert.Equal(t, NewIntSliceV(3), slice)
		assert.Equal(t, &Object{3}, slice.Shift())
		assert.Equal(t, NewIntSliceV(), slice)
		assert.Equal(t, &Object{nil}, slice.Shift())
		assert.Equal(t, NewIntSliceV(), slice)
	}
}

// TakeFirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_TakeFirstN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkIntSlice_TakeFirstN_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.ShiftN(10)
	}
}

func ExampleIntSlice_TakeFirstN() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.ShiftN(2))
	// Output: [1 2]
}

func TestIntSlice_TakeFirstN(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.ShiftN(1))
	}

	// negative value
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(1), slice.ShiftN(-1))
		assert.Equal(t, NewIntSliceV(2, 3), slice)
	}

	// take none
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(), slice.ShiftN(0))
		assert.Equal(t, NewIntSliceV(1, 2, 3), slice)
	}

	// take 1
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(1), slice.ShiftN(1))
		assert.Equal(t, NewIntSliceV(2, 3), slice)
	}

	// take 2
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(1, 2), slice.ShiftN(2))
		assert.Equal(t, NewIntSliceV(3), slice)
	}

	// take 3
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(1, 2, 3), slice.ShiftN(3))
		assert.Equal(t, NewIntSliceV(), slice)
	}

	// take beyond
	{
		slice := NewIntSliceV(1, 2, 3)
		assert.Equal(t, NewIntSliceV(1, 2, 3), slice.ShiftN(4))
		assert.Equal(t, NewIntSliceV(), slice)
	}
}

// TakeW
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_TakeW_Go(t *testing.B) {
	new := []int{}
	ints := Range(0, nines5)
	l := len(ints)
	for i := 0; i < l; i++ {
		if ints[i]%2 == 0 {
			new = append(new, ints[i])
			if i+1 < l {
				ints = append(ints[:i], ints[i+1:]...)
			} else if i >= 0 && i < l {
				ints = ints[:i]
			}
			l--
			i--
		}
	}
}

func BenchmarkIntSlice_TakeW_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines5))
	slice.TakeW(func(x O) bool { return ExB(x.(int)%2 == 0) })
}

func ExampleIntSlice_TakeW() {
	slice := NewIntSliceV(1, 2, 3)
	fmt.Println(slice.TakeW(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	}))
	// Output: [2]
}

func TestIntSlice_TakeW(t *testing.T) {

	// drop all odd values
	{
		slice := NewIntSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.TakeW(func(x O) bool { return ExB(x.(int)%2 != 0) })
		assert.Equal(t, NewIntSliceV(2, 4, 6, 8), slice)
		assert.Equal(t, NewIntSliceV(1, 3, 5, 7, 9), new)
	}

	// drop all even values
	{
		slice := NewIntSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.TakeW(func(x O) bool { return ExB(x.(int)%2 == 0) })
		assert.Equal(t, NewIntSliceV(1, 3, 5, 7, 9), slice)
		assert.Equal(t, NewIntSliceV(2, 4, 6, 8), new)
	}
}

// Uniq
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_Uniq_Go(t *testing.B) {
	// ints := Range(0, nines7)
	// for len(ints) > 10 {
	// 	ints = ints[10:]
	// }
}

func BenchmarkIntSlice_Uniq_Slice(t *testing.B) {
	// slice := NewIntSlice(Range(0, nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleIntSlice_Uniq() {
	slice := NewIntSliceV(1, 2, 3, 3)
	fmt.Println(slice.Uniq())
	// Output: [1 2 3]
}

func TestIntSlice_Uniq(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, NewIntSliceV(), slice.Uniq())
	}

	// size of one
	{
		slice := NewIntSliceV(1)
		uniq := slice.Uniq()
		assert.Equal(t, NewIntSliceV(1), uniq)
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1), uniq)
	}

	// one duplicate
	{
		slice := NewIntSliceV(1, 1)
		uniq := slice.Uniq()
		assert.Equal(t, NewIntSliceV(1), uniq)
		assert.Equal(t, NewIntSliceV(1, 1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1), uniq)
	}

	// multiple duplicates
	{
		slice := NewIntSliceV(1, 2, 2, 3, 3)
		uniq := slice.Uniq()
		assert.Equal(t, NewIntSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewIntSliceV(1, 2, 2, 3, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3), uniq)
	}

	// no duplicates
	{
		slice := NewIntSliceV(1, 2, 3)
		uniq := slice.Uniq()
		assert.Equal(t, NewIntSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3), uniq)
	}
}

// UniqM
//--------------------------------------------------------------------------------------------------
func BenchmarkIntSlice_UniqM_Go(t *testing.B) {
	// ints := Range(0, nines7)
	// for len(ints) > 10 {
	// 	ints = ints[10:]
	// }
}

func BenchmarkIntSlice_UniqM_Slice(t *testing.B) {
	// slice := NewIntSlice(Range(0, nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleIntSlice_UniqM() {
	slice := NewIntSliceV(1, 2, 3, 3)
	fmt.Println(slice.UniqM())
	// Output: [1 2 3]
}

func TestIntSlice_UniqM(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, (*IntSlice)(nil), slice.UniqM())
	}

	// size of one
	{
		slice := NewIntSliceV(1)
		uniq := slice.UniqM()
		assert.Equal(t, NewIntSliceV(1), uniq)
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1, 2), uniq)
	}

	// one duplicate
	{
		slice := NewIntSliceV(1, 1)
		uniq := slice.UniqM()
		assert.Equal(t, NewIntSliceV(1), uniq)
		assert.Equal(t, NewIntSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewIntSliceV(1, 2), uniq)
	}

	// multiple duplicates
	{
		slice := NewIntSliceV(1, 2, 2, 3, 3)
		uniq := slice.UniqM()
		assert.Equal(t, NewIntSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), uniq)
	}

	// no duplicates
	{
		slice := NewIntSliceV(1, 2, 3)
		uniq := slice.UniqM()
		assert.Equal(t, NewIntSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewIntSliceV(1, 2, 3, 4), uniq)
	}
}
