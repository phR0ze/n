package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewFloatSlice
//--------------------------------------------------------------------------------------------------
// func BenchmarkNewFloatSlice_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0.0; i < len(ints); i += 10.0 {
// 		_ = []float64{i, i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9}
// 	}
// }

// func BenchmarkNewFloatSlice_Slice(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i += 10 {
// 		_ = NewFloatSlice([]float64{i, i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9})
// 	}
// }

func ExampleNewFloatSlice() {
	slice := NewFloatSlice([]float64{1, 2, 3})
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_NewFloatSlice(t *testing.T) {

	// array
	var array [2]int
	array[0] = 1
	array[1] = 2
	assert.Equal(t, []float64{1, 2}, NewFloatSlice(array).O())
	assert.Equal(t, []float64{1, 2}, NewFloatSlice(array[:]).O())

	// empty
	assert.Equal(t, []float64{}, NewFloatSlice([]float64{}).O())

	// slice
	assert.Equal(t, []float64{0}, NewFloatSlice([]float64{0}).O())
	assert.Equal(t, []float64{1, 2}, NewFloatSlice([]float64{1, 2}).O())

	// Conversion
	{
		assert.Equal(t, []float64{1}, NewFloatSlice("1").O())
		assert.Equal(t, []float64{1, 2}, NewFloatSlice([]string{"1", "2"}).O())
		assert.Equal(t, []float64{1}, NewFloatSlice(Object{1}).O())
		assert.Equal(t, []float64{1, 2}, NewFloatSlice([]Object{{1}, {2}}).O())
		assert.Equal(t, []float64{1}, NewFloatSlice(true).O())
		assert.Equal(t, []float64{1, 0}, NewFloatSlice([]bool{true, false}).O())
	}
}

// NewFloatSliceV
//--------------------------------------------------------------------------------------------------
// func BenchmarkNewFloatSliceV_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i += 10 {
// 		_ = append([]float64{}, i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9)
// 	}
// }

// func BenchmarkNewFloatSliceV_Slice(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i += 10 {
// 		_ = NewFloatSliceV(i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9)
// 	}
// }

func ExampleNewFloatSliceV_empty() {
	slice := NewFloatSliceV()
	fmt.Println(slice)
	// Output: []
}

func ExampleNewFloatSliceV_variadic() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_NewFloatSliceV(t *testing.T) {

	// empty
	assert.Equal(t, []float64{}, NewFloatSliceV().O())

	// multiples
	assert.Equal(t, []float64{1}, NewFloatSliceV(1).O())
	assert.Equal(t, []float64{1, 2}, NewFloatSliceV(1, 2).O())
	assert.Equal(t, []float64{1, 2}, NewFloatSliceV([]interface{}{1, 2}...).O())

	// Conversion
	{
		assert.Equal(t, []float64{1, 2}, NewFloatSliceV("1", "2").O())
		assert.Equal(t, []float64{1, 2}, NewFloatSliceV(Obj(1.0), Obj(2)).O())
		assert.Equal(t, []float64{1, 0}, NewFloatSliceV(true, false).O())
	}
}

// Any
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Any_Go(t *testing.B) {
// 	any := func(list []float64, x []float64) bool {
// 		for i := range x {
// 			for j := range list {
// 				if list[j] == x[i] {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}

// 	// test here
// 	ints := Range(0, nines4)
// 	for i := range ints {
// 		any(ints, []float64{i})
// 	}
// }

// func BenchmarkFloatSlice_Any_Slice(t *testing.B) {
// 	src := Range(0, nines4)
// 	slice := NewFloatSlice(src)
// 	for i := range src {
// 		slice.Any(i)
// 	}
// }

func ExampleFloatSlice_Any_empty() {
	slice := NewFloatSliceV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleFloatSlice_Any_notEmpty() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleFloatSlice_Any_contains() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Any(1))
	// Output: true
}

func ExampleFloatSlice_Any_containsAny() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Any(0, 1))
	// Output: true
}

func TestFloatSlice_Any(t *testing.T) {

	// empty
	var nilSlice *FloatSlice
	assert.False(t, nilSlice.Any())
	assert.False(t, NewFloatSliceV().Any())

	// single
	assert.True(t, NewFloatSliceV(2).Any())

	// invalid
	assert.False(t, NewFloatSliceV(1, 2).Any(TestObj{2}))

	assert.True(t, NewFloatSliceV(1, 2, 3).Any(2))
	assert.False(t, NewFloatSliceV(1, 2, 3).Any(4))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(4, 3))
	assert.False(t, NewFloatSliceV(1, 2, 3).Any(4, 5))

	// conversion
	assert.True(t, NewFloatSliceV(1, 2).Any(Object{2}))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(int8(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(int16(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(int32(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(int64(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(uint8(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(uint16(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(uint32(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(uint64(2)))
	assert.True(t, NewFloatSliceV(1, 2, 3).Any(uint64(2)))
}

// AnyS
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_AnyS_Go(t *testing.B) {
// 	any := func(list []float64, x []float64) bool {
// 		for i := range x {
// 			for j := range list {
// 				if list[j] == x[i] {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}

// 	// test here
// 	ints := Range(0, nines4)
// 	for i := range ints {
// 		any(ints, []float64{i})
// 	}
// }

// func BenchmarkFloatSlice_AnyS_Slice(t *testing.B) {
// 	src := Range(0, nines4)
// 	slice := NewFloatSlice(src)
// 	for i := range src {
// 		slice.Any([]float64{i})
// 	}
// }

func ExampleFloatSlice_AnyS() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.AnyS([]float64{0, 1}))
	// Output: true
}

func TestFloatSlice_AnyS(t *testing.T) {
	// nil
	{
		var slice *FloatSlice
		assert.False(t, slice.AnyS([]float64{1}))
		assert.False(t, NewFloatSliceV(1.0).AnyS(nil))
	}

	// []float64
	{
		assert.True(t, NewFloatSliceV(1.0, 2.0, 3.0).AnyS([]float64{1}))
		assert.True(t, NewFloatSliceV(1.0, 2.0, 3.0).AnyS([]float64{4, 3}))
		assert.False(t, NewFloatSliceV(1, 2, 3).AnyS([]float64{4, 5}))
	}

	// *[]float64
	{
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(&([]float64{1})))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(&([]float64{4, 3})))
		assert.False(t, NewFloatSliceV(1, 2, 3).AnyS(&([]float64{4, 5})))
	}

	// Object
	{
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS([]Object{{1}, {2}}))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS([]Object{{1}, {4}}))
		assert.False(t, NewFloatSliceV(1, 2, 3).AnyS([]Object{{4}, {5}}))
	}

	// ISlice
	{
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(ISlice(NewFloatSliceV(1))))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(ISlice(NewFloatSliceV(4, 3))))
		assert.False(t, NewFloatSliceV(1, 2, 3).AnyS(ISlice(NewFloatSliceV(4, 5))))
	}

	// FloatSlice
	{
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(*NewFloatSliceV(1)))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(*NewFloatSliceV(4, 3)))
		assert.False(t, NewFloatSliceV(1, 2, 3).AnyS(*NewFloatSliceV(4, 5)))
	}

	// *FloatSlice
	{
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(NewFloatSliceV(1)))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(NewFloatSliceV(4, 3)))
		assert.False(t, NewFloatSliceV(1, 2, 3).AnyS(NewFloatSliceV(4, 5)))
	}

	// invalid types
	assert.False(t, NewFloatSliceV(1, 2).AnyS(nil))
	assert.False(t, NewFloatSliceV(1, 2).AnyS((*[]float64)(nil)))
	assert.False(t, NewFloatSliceV(1, 2).AnyS((*FloatSlice)(nil)))
	assert.False(t, NewFloatSliceV(1, 2).AnyS([]string{"bob"}))

	// Conversion
	{
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(NewFloatSliceV(int64(1))))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(NewFloatSliceV("2")))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(NewFloatSliceV(true)))
		assert.True(t, NewFloatSliceV(1, 2, 3).AnyS(NewFloatSliceV(Str("3"))))
	}
}

// AnyW
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_AnyW_Go(t *testing.B) {
// 	ints := Range(0, nines5)
// 	for i := range ints {
// 		if i == nines4 {
// 			break
// 		}
// 	}
// }

// func BenchmarkFloatSlice_AnyW_Slice(t *testing.B) {
// 	src := Range(0, nines5)
// 	NewFloatSlice(src).AnyW(func(x O) bool {
// 		return ExB(x.(float64) == nines4)
// 	})
// }

func ExampleFloatSlice_AnyW() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.AnyW(func(x O) bool {
		return ExB(x.(float64) == 2)
	}))
	// Output: true
}

func TestFloatSlice_AnyW(t *testing.T) {

	// empty
	var slice *FloatSlice
	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(float64) > 0) }))
	assert.False(t, NewFloatSliceV().AnyW(func(x O) bool { return ExB(x.(float64) > 0) }))

	// single
	assert.True(t, NewFloatSliceV(2).AnyW(func(x O) bool { return ExB(x.(float64) > 0) }))

	assert.True(t, NewFloatSliceV(1, 2).AnyW(func(x O) bool { return ExB(x.(float64) == 2) }))
	assert.True(t, NewFloatSliceV(1, 2, 3).AnyW(func(x O) bool { return ExB(x.(float64) == 4 || x.(float64) == 3) }))
}

// Append
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Append_Go(t *testing.B) {
// 	ints := []float64{}
// 	for _, i := range Range(0, nines6) {
// 		ints = append(ints, i)
// 	}
// }

// func BenchmarkFloatSlice_Append_Slice(t *testing.B) {
// 	slice := NewFloatSliceV()
// 	for _, i := range Range(0, nines6) {
// 		slice.Append(i)
// 	}
// }

func ExampleFloatSlice_Append() {
	slice := NewFloatSliceV(1).Append(2).Append(3)
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Append(t *testing.T) {

	// nil
	{
		var nilSlice *FloatSlice
		assert.Equal(t, NewFloatSliceV(0), nilSlice.Append(0))
		assert.Equal(t, (*FloatSlice)(nil), nilSlice)
	}

	// Append one back to back
	{
		var slice *FloatSlice
		assert.Equal(t, true, slice.Nil())
		slice = NewFloatSliceV()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, false, slice.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		slice.Append(1)
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []float64{1}, slice.O())

		// Second append another which will be 2x at most
		slice.Append(2)
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []float64{1, 2}, slice.O())
		assert.Equal(t, NewFloatSliceV(1, 2), slice)
	}

	// Start with just appending without chaining
	{
		slice := NewFloatSliceV()
		assert.Equal(t, 0, slice.Len())
		slice.Append(1)
		assert.Equal(t, []float64{1}, slice.O())
		slice.Append(2)
		assert.Equal(t, []float64{1, 2}, slice.O())
	}

	// Start with nil not chained
	{
		slice := NewFloatSliceV()
		assert.Equal(t, 0, slice.Len())
		slice.Append(1).Append(2).Append(3)
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []float64{1, 2, 3}, slice.O())
	}

	// Start with nil chained
	{
		slice := NewFloatSliceV().Append(1).Append(2)
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []float64{1, 2}, slice.O())
	}

	// Start with non nil
	{
		slice := NewFloatSliceV(1).Append(2).Append(3)
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []float64{1, 2, 3}, slice.O())
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
	}

	// Use append result directly
	{
		slice := NewFloatSliceV(1)
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []float64{1, 2}, slice.Append(2).O())
		assert.Equal(t, NewFloatSliceV(1, 2), slice)
	}

	// Conversion
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1).Append(Object{2}))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1).Append("2"))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV().Append(true).Append(Char('2')))
	}
}

// AppendV
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_AppendV_Go(t *testing.B) {
// 	ints := []float64{}
// 	ints = append(ints, Range(0, nines6)...)
// }

// func BenchmarkFloatSlice_AppendV_Slice(t *testing.B) {
// 	n := NewFloatSliceV()
// 	new := rangeO(0, nines6)
// 	n.AppendV(new...)
// }

func ExampleFloatSlice_AppendV() {
	slice := NewFloatSliceV(1).AppendV(2, 3)
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_AppendV(t *testing.T) {

	// nil
	{
		var nilSlice *FloatSlice
		assert.Equal(t, NewFloatSliceV(1, 2), nilSlice.AppendV(1, 2))
	}

	// Append many ints
	{
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1).AppendV(2, 3))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4, 5), NewFloatSliceV(1).AppendV(2, 3).AppendV(4, 5))
	}

	// Conversion
	{
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV().AppendV(Object{0}, Object{1}))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV().AppendV("0", "1"))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV().AppendV(false, true))
	}
}

// At
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_At_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := range ints {
// 		assert.IsType(t, 0, ints[i])
// 	}
// }

// func BenchmarkFloatSlice_At_Slice(t *testing.B) {
// 	src := Range(0, nines6)
// 	slice := NewFloatSlice(src)
// 	for _, i := range src {
// 		_, ok := (slice.At(i).O()).(int)
// 		assert.True(t, ok)
// 	}
// }

func ExampleFloatSlice_At() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.At(2))
	// Output: 3
}

func TestFloatSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *FloatSlice
		assert.Equal(t, Obj(nil), nilSlice.At(0))
	}

	// ints
	{
		slice := NewFloatSliceV(1, 2, 3, 4)
		assert.Equal(t, 4.0, slice.At(-1).O())
		assert.Equal(t, 3.0, slice.At(-2).O())
		assert.Equal(t, 2.0, slice.At(-3).O())
		assert.Equal(t, 1.0, slice.At(0).O())
		assert.Equal(t, 2.0, slice.At(1).O())
		assert.Equal(t, 3.0, slice.At(2).O())
		assert.Equal(t, 4.0, slice.At(3).O())
	}

	// index out of bounds
	{
		slice := NewFloatSliceV(1)
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------

func ExampleFloatSlice_Clear() {
	slice := NewFloatSliceV(1).Concat([]float64{2, 3})
	fmt.Println(slice.Clear())
	// Output: []
}

func TestFloatSlice_Clear(t *testing.T) {

	// nil
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.Clear())
		assert.Equal(t, (*FloatSlice)(nil), slice)
	}

	// int
	{
		slice := NewFloatSliceV(1, 2, 3, 4)
		assert.Equal(t, NewFloatSliceV(), slice.Clear())
		assert.Equal(t, NewFloatSliceV(), slice.Clear())
		assert.Equal(t, NewFloatSliceV(), slice)
	}
}

// Concat
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Concat_Go(t *testing.B) {
// 	dest := []float64{}
// 	src := Range(0, nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkFloatSlice_Concat_Slice(t *testing.B) {
// 	dest := NewFloatSliceV()
// 	src := Range(0, nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.Concat(src[j:i])
// 		j = i
// 	}
// }

func ExampleFloatSlice_Concat() {
	slice := NewFloatSliceV(1).Concat([]float64{2, 3})
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Concat(t *testing.T) {

	// nil
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Concat([]float64{1, 2}))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).Concat(nil))
	}

	// []float64
	{
		slice := NewFloatSliceV(1)
		concated := slice.Concat([]float64{2, 3})
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), concated)
	}

	// *[]float64
	{
		slice := NewFloatSliceV(1)
		concated := slice.Concat(&([]float64{2, 3}))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), concated)
	}

	// *FloatSlice
	{
		slice := NewFloatSliceV(1)
		concated := slice.Concat(NewFloatSliceV(2, 3))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), concated)
	}

	// FloatSlice
	{
		slice := NewFloatSliceV(1)
		concated := slice.Concat(*NewFloatSliceV(2, 3))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), concated)
	}

	// ISlice
	{
		slice := NewFloatSliceV(1)
		concated := slice.Concat(ISlice(NewFloatSliceV(2, 3)))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), concated)
	}

	// nils
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).Concat((*[]float64)(nil)))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).Concat((*FloatSlice)(nil)))
	}

	// Conversion
	{
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV().Concat([]Object{{0}, {1}}))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV().Concat([]string{"0", "1"}))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV().Concat([]bool{false, true}))
	}
}

// ConcatM
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_ConcatM_Go(t *testing.B) {
// 	dest := []float64{}
// 	src := Range(0, nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkFloatSlice_ConcatM_Slice(t *testing.B) {
// 	dest := NewFloatSliceV()
// 	src := Range(0, nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.ConcatM(src[j:i])
// 		j = i
// 	}
// }

func ExampleFloatSlice_ConcatM() {
	slice := NewFloatSliceV(1).ConcatM([]float64{2, 3})
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_ConcatM(t *testing.T) {

	// nil
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(1, 2), slice.ConcatM([]float64{1, 2}))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).ConcatM(nil))
	}

	// []float64
	{
		slice := NewFloatSliceV(1)
		concated := slice.ConcatM([]float64{2, 3})
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), concated)
	}

	// *[]float64
	{
		slice := NewFloatSliceV(1)
		concated := slice.ConcatM(&([]float64{2, 3}))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), concated)
	}

	// *FloatSlice
	{
		slice := NewFloatSliceV(1)
		concated := slice.ConcatM(NewFloatSliceV(2, 3))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), concated)
	}

	// FloatSlice
	{
		slice := NewFloatSliceV(1)
		concated := slice.ConcatM(*NewFloatSliceV(2, 3))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), concated)
	}

	// ISlice
	{
		slice := NewFloatSliceV(1)
		concated := slice.ConcatM(ISlice(NewFloatSliceV(2, 3)))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), concated)
	}

	// nils
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).ConcatM((*[]float64)(nil)))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).ConcatM((*FloatSlice)(nil)))
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Copy_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	dst := make([]float64, len(ints), len(ints))
// 	copy(dst, ints)
// }

// func BenchmarkFloatSlice_Copy_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.Copy()
// }

func ExampleFloatSlice_Copy() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Copy())
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Copy(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.Copy(0, -1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(0).Clear().Copy(0, -1))
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := NewFloatSliceV(1, 2, 3)
		result := original.Copy(0, -1)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), original)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), result)
		result.Set(0, 0)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), original)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), result)
	}

	// copy full array
	{
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Copy())
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Copy(0, -1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Copy(0, 1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Copy(0, 5))
		assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1).Copy())
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Copy())
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Copy(0, -1))
		assert.Equal(t, NewFloatSlice([]float64{1, 2, 3}), NewFloatSlice([]float64{1, 2, 3}).Copy())
		assert.Equal(t, NewFloatSlice([]float64{1, 2, 3}), NewFloatSlice([]float64{1, 2, 3}).Copy(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1).Copy(0, 2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Copy(-6, 6))
	}

	// mutually exclusive
	{
		slice := NewFloatSliceV(1, 2, 3, 4)
		assert.Equal(t, NewFloatSliceV(), slice.Copy(2, -3))
		assert.Equal(t, NewFloatSliceV(), slice.Copy(0, -5))
		assert.Equal(t, NewFloatSliceV(), slice.Copy(4, -1))
		assert.Equal(t, NewFloatSliceV(), slice.Copy(6, -1))
		assert.Equal(t, NewFloatSliceV(), slice.Copy(3, -2))
	}

	// singles
	{
		slice := NewFloatSliceV(1, 2, 3, 4)
		assert.Equal(t, NewFloatSliceV(4), slice.Copy(-1, -1))
		assert.Equal(t, NewFloatSliceV(3), slice.Copy(-2, -2))
		assert.Equal(t, NewFloatSliceV(2), slice.Copy(-3, -3))
		assert.Equal(t, NewFloatSliceV(1), slice.Copy(0, 0))
		assert.Equal(t, NewFloatSliceV(1), slice.Copy(-4, -4))
		assert.Equal(t, NewFloatSliceV(2), slice.Copy(1, 1))
		assert.Equal(t, NewFloatSliceV(2), slice.Copy(1, -3))
		assert.Equal(t, NewFloatSliceV(3), slice.Copy(2, 2))
		assert.Equal(t, NewFloatSliceV(3), slice.Copy(2, -2))
		assert.Equal(t, NewFloatSliceV(4), slice.Copy(3, 3))
		assert.Equal(t, NewFloatSliceV(4), slice.Copy(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Copy(1, -1))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Copy(1, 2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Copy(-2, -1))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Copy(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Copy(0, -2))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Copy(-3, -2))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Copy(-3, 1))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Copy(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Copy(1, -2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Copy(-3, -2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Copy(-3, 2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Copy(1, 2))
	}

	// random
	{
		assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1, 2, 3).Copy(0, -3))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Copy(1, 2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Copy(0, 2))
	}
}

// Count
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Count_Go(t *testing.B) {
// 	ints := Range(0, nines5)
// 	for i := range ints {
// 		if i == nines4 {
// 			break
// 		}
// 	}
// }

// func BenchmarkFloatSlice_Count_Slice(t *testing.B) {
// 	src := Range(0, nines5)
// 	NewFloatSlice(src).Count(nines4)
// }

func ExampleFloatSlice_Count() {
	slice := NewFloatSliceV(1, 2, 2)
	fmt.Println(slice.Count(2.0))
	// Output: 2
}

func TestFloatSlice_Count(t *testing.T) {

	// empty
	var slice *FloatSlice
	assert.Equal(t, 0, slice.Count(0))
	assert.Equal(t, 0, NewFloatSliceV().Count(0))

	assert.Equal(t, 1, NewFloatSliceV(2, 3).Count(2))
	assert.Equal(t, 2, NewFloatSliceV(1, 2, 2).Count(2))
	assert.Equal(t, 4, NewFloatSliceV(4, 4, 3, 4, 4).Count(4))
	assert.Equal(t, 3, NewFloatSliceV(3, 2, 3, 3, 5).Count(3))
	assert.Equal(t, 1, NewFloatSliceV(1, 2, 3).Count(3))
}

// CountW
//--------------------------------------------------------------------------------------------------
func BenchmarkFloatSlice_CountW_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if i == nines4 {
			break
		}
	}
}

func BenchmarkFloatSlice_CountW_Slice(t *testing.B) {
	src := Range(0, nines5)
	NewFloatSlice(src).CountW(func(x O) bool {
		return ExB(x.(float64) == nines4)
	})
}

func ExampleFloatSlice_CountW() {
	slice := NewFloatSliceV(1, 2, 2)
	fmt.Println(slice.CountW(func(x O) bool {
		return ExB(x.(float64) == 2)
	}))
	// Output: 2
}

func TestFloatSlice_CountW(t *testing.T) {

	// empty
	var slice *FloatSlice
	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(float64) > 0) }))
	assert.Equal(t, 0, NewFloatSliceV().CountW(func(x O) bool { return ExB(x.(float64) > 0) }))

	assert.Equal(t, 1, NewFloatSliceV(2, 3).CountW(func(x O) bool { return ExB(x.(float64) > 2) }))
	assert.Equal(t, 1, NewFloatSliceV(1, 2).CountW(func(x O) bool { return ExB(x.(float64) == 2) }))
	assert.Equal(t, 1, NewFloatSliceV(1, 2, 3).CountW(func(x O) bool { return ExB(x.(float64) == 4 || x.(float64) == 3) }))
}

// Drop
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Drop_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 11 {
// 		i := 1
// 		n := 10
// 		if i+n < len(ints) {
// 			ints = append(ints[:i], ints[i+n:]...)
// 		} else {
// 			ints = ints[:i]
// 		}
// 	}
// }

// func BenchmarkFloatSlice_Drop_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 1 {
// 		slice.Drop(1, 10)
// 	}
// }

func ExampleFloatSlice_Drop() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Drop(0, 1))
	// Output: [3.000000]
}

func TestFloatSlice_Drop(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.Drop(0, 1))
	}

	// invalid
	assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(1))
	assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(4, 4))

	// drop 1
	assert.Equal(t, NewFloatSliceV(2, 3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(0, 0))
	assert.Equal(t, NewFloatSliceV(1, 3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(1, 1))
	assert.Equal(t, NewFloatSliceV(1, 2, 4), NewFloatSliceV(1, 2, 3, 4).Drop(2, 2))
	assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3, 4).Drop(3, 3))
	assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3, 4).Drop(-1, -1))
	assert.Equal(t, NewFloatSliceV(1, 2, 4), NewFloatSliceV(1, 2, 3, 4).Drop(-2, -2))
	assert.Equal(t, NewFloatSliceV(1, 3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(-3, -3))
	assert.Equal(t, NewFloatSliceV(2, 3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(-4, -4))

	// drop 2
	assert.Equal(t, NewFloatSliceV(3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(0, 1))
	assert.Equal(t, NewFloatSliceV(1, 4), NewFloatSliceV(1, 2, 3, 4).Drop(1, 2))
	assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3, 4).Drop(2, 3))
	assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3, 4).Drop(-2, -1))
	assert.Equal(t, NewFloatSliceV(1, 4), NewFloatSliceV(1, 2, 3, 4).Drop(-3, -2))
	assert.Equal(t, NewFloatSliceV(3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(-4, -3))

	// drop 3
	assert.Equal(t, NewFloatSliceV(4), NewFloatSliceV(1, 2, 3, 4).Drop(0, 2))
	assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1, 2, 3, 4).Drop(-3, -1))

	// drop everything and beyond
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Drop())
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Drop(0, 3))
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Drop(0, -1))
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Drop(-4, -1))
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Drop(-6, -1))
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Drop(0, 10))

	// move index within bounds
	assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3, 4).Drop(3, 4))
	assert.Equal(t, NewFloatSliceV(2, 3, 4), NewFloatSliceV(1, 2, 3, 4).Drop(-5, 0))
}

// DropAt
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_DropAt_Go(t *testing.B) {
// 	ints := Range(0, nines5)
// 	index := Range(0, nines5)
// 	for i := range index {
// 		if i+1 < len(ints) {
// 			ints = append(ints[:i], ints[i+1:]...)
// 		} else if i >= 0 && i < len(ints) {
// 			ints = ints[:i]
// 		}
// 	}
// }

// func BenchmarkFloatSlice_DropAt_Slice(t *testing.B) {
// 	src := Range(0, nines5)
// 	index := Range(0, nines5)
// 	slice := NewFloatSlice(src)
// 	for i := range index {
// 		slice.DropAt(i)
// 	}
// }

func ExampleFloatSlice_DropAt() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.DropAt(1))
	// Output: [1.000000 3.000000]
}

func TestFloatSlice_DropAt(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.DropAt(0))
	}

	// drop all and more
	{
		slice := NewFloatSliceV(0, 1, 2)
		assert.Equal(t, NewFloatSliceV(0, 1), slice.DropAt(-1))
		assert.Equal(t, NewFloatSliceV(0), slice.DropAt(-1))
		assert.Equal(t, NewFloatSliceV(), slice.DropAt(-1))
		assert.Equal(t, NewFloatSliceV(), slice.DropAt(-1))
	}

	// drop invalid
	assert.Equal(t, NewFloatSliceV(0, 1, 2), NewFloatSliceV(0, 1, 2).DropAt(3))
	assert.Equal(t, NewFloatSliceV(0, 1, 2), NewFloatSliceV(0, 1, 2).DropAt(-4))

	// drop last
	assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1, 2).DropAt(2))
	assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1, 2).DropAt(-1))

	// drop middle
	assert.Equal(t, NewFloatSliceV(0, 2), NewFloatSliceV(0, 1, 2).DropAt(1))
	assert.Equal(t, NewFloatSliceV(0, 2), NewFloatSliceV(0, 1, 2).DropAt(-2))

	// drop first
	assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(0, 1, 2).DropAt(0))
	assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(0, 1, 2).DropAt(-3))
}

// DropFirst
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_DropFirst_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 1 {
// 		ints = ints[1:]
// 	}
// }

// func BenchmarkFloatSlice_DropFirst_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirst()
// 	}
// }

func ExampleFloatSlice_DropFirst() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.DropFirst())
	// Output: [2.000000 3.000000]
}

func TestFloatSlice_DropFirst(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.DropFirst())
	}

	// drop all and beyond
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(2, 3), slice.DropFirst())
		assert.Equal(t, NewFloatSliceV(3), slice.DropFirst())
		assert.Equal(t, NewFloatSliceV(), slice.DropFirst())
		assert.Equal(t, NewFloatSliceV(), slice.DropFirst())
	}
}

// DropFirstN
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_DropFirstN_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 10 {
// 		ints = ints[10:]
// 	}
// }

// func BenchmarkFloatSlice_DropFirstN_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirstN(10)
// 	}
// }

func ExampleFloatSlice_DropFirstN() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.DropFirstN(2))
	// Output: [3.000000]
}

func TestFloatSlice_DropFirstN(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.DropFirstN(1))
	}

	// negative value
	assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).DropFirstN(-1))

	// drop none
	assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).DropFirstN(0))

	// drop 1
	assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).DropFirstN(1))

	// drop 2
	assert.Equal(t, NewFloatSliceV(3), NewFloatSliceV(1, 2, 3).DropFirstN(2))

	// drop 3
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3).DropFirstN(3))

	// drop beyond
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3).DropFirstN(4))
}

// DropLast
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_DropLast_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 1 {
// 		ints = ints[1:]
// 	}
// }

// func BenchmarkFloatSlice_DropLast_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLast()
// 	}
// }

func ExampleFloatSlice_DropLast() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.DropLast())
	// Output: [1.000000 2.000000]
}

func TestFloatSlice_DropLast(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.DropLast())
	}

	// negative value
	assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).DropLastN(-1))

	slice := NewFloatSliceV(1, 2, 3)
	assert.Equal(t, NewFloatSliceV(1, 2), slice.DropLast())
	assert.Equal(t, NewFloatSliceV(1), slice.DropLast())
	assert.Equal(t, NewFloatSliceV(), slice.DropLast())
	assert.Equal(t, NewFloatSliceV(), slice.DropLast())
}

// DropLastN
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_DropLastN_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 10 {
// 		ints = ints[10:]
// 	}
// }

// func BenchmarkFloatSlice_DropLastN_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLastN(10)
// 	}
// }

func ExampleFloatSlice_DropLastN() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.DropLastN(2))
	// Output: [1.000000]
}

func TestFloatSlice_DropLastN(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.DropLastN(1))
	}

	// drop none
	assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).DropLastN(0))

	// drop 1
	assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).DropLastN(1))

	// drop 2
	assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1, 2, 3).DropLastN(2))

	// drop 3
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3).DropLastN(3))

	// drop beyond
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3).DropLastN(4))
}

// DropW
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_DropW_Go(t *testing.B) {
// 	ints := Range(0, nines5)
// 	l := len(ints)
// 	for i := 0; i < l; i++ {
// 		if ints[i]%2 == 0 {
// 			if i+1 < l {
// 				ints = append(ints[:i], ints[i+1:]...)
// 			} else if i >= 0 && i < l {
// 				ints = ints[:i]
// 			}
// 			l--
// 			i--
// 		}
// 	}
// }

// func BenchmarkFloatSlice_DropW_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines5))
// 	slice.DropW(func(x O) bool {
// 		return ExB(x.(float64)%2 == 0)
// 	})
// }

// func ExampleFloatSlice_DropW() {
// 	slice := NewFloatSliceV(1, 2, 3)
// 	fmt.Println(slice.DropW(func(x O) bool {
// 		return ExB(x.(float64)%2 == 0)
// 	}))
// 	// Output: [1 3.000000]
// }

func TestFloatSlice_DropW(t *testing.T) {

	// // drop all odd values
	// {
	// 	slice := NewFloatSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// 	slice.DropW(func(x O) bool {
	// 		return ExB(x.(float64)%2 != 0)
	// 	})
	// 	assert.Equal(t, NewFloatSliceV(2, 4, 6, 8), slice)
	// }

	// // drop all even values
	// {
	// 	slice := NewFloatSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// 	slice.DropW(func(x O) bool {
	// 		return ExB(x.(float64)%2 == 0)
	// 	})
	// 	assert.Equal(t, NewFloatSliceV(1, 3, 5, 7, 9), slice)
	// }
}

// Each
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Each_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_Each_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).Each(func(x O) {
// 		assert.IsType(t, 0, x)
// 	})
// }

func ExampleFloatSlice_Each() {
	NewFloatSliceV(1, 2, 3).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 123
}

func TestFloatSlice_Each(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.Each(func(x O) {})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).Each(func(x O) {
			results = append(results, x.(float64))
		})
		assert.Equal(t, []float64{1, 2, 3}, results)
	}
}

// EachE
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_EachE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_EachE_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).EachE(func(x O) error {
// 		assert.IsType(t, 0, x)
// 		return nil
// 	})
// }

func ExampleFloatSlice_EachE() {
	NewFloatSliceV(1, 2, 3).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 123
}

func TestFloatSlice_EachE(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.EachE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachE(func(x O) error {
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{1, 2, 3}, results)
	}

	// Break early with error
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachE(func(x O) error {
			if x.(float64) == 3 {
				return Break
			}
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{1, 2}, results)
	}
}

// EachI
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_EachI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_EachI_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).EachI(func(i int, x O) {
// 		assert.IsType(t, 0, x)
// 	})
// }

func ExampleFloatSlice_EachI() {
	NewFloatSliceV(1, 2, 3).EachI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 0:11:22:3
}

func TestFloatSlice_EachI(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.EachI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachI(func(i int, x O) {
			results = append(results, x.(float64))
		})
		assert.Equal(t, []float64{1, 2, 3}, results)
	}
}

// EachIE
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_EachIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_EachIE_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).EachIE(func(i int, x O) error {
// 		assert.IsType(t, 0, x)
// 		return nil
// 	})
// }

func ExampleFloatSlice_EachIE() {
	NewFloatSliceV(1, 2, 3).EachIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 0:11:22:3
}

func TestFloatSlice_EachIE(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.EachIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachIE(func(i int, x O) error {
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{1, 2, 3}, results)
	}

	// Break early with error
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachIE(func(i int, x O) error {
			if i == 2 {
				return Break
			}
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{1, 2}, results)
	}
}

// EachR
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_EachR_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_EachR_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).EachR(func(x O) {
// 		assert.IsType(t, 0, x)
// 	})
// }

func ExampleFloatSlice_EachR() {
	NewFloatSliceV(1, 2, 3).EachR(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 321
}

func TestFloatSlice_EachR(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.EachR(func(x O) {})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachR(func(x O) {
			results = append(results, x.(float64))
		})
		assert.Equal(t, []float64{3, 2, 1}, results)
	}
}

// EachRE
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_EachRE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_EachRE_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).EachRE(func(x O) error {
// 		assert.IsType(t, 0, x)
// 		return nil
// 	})
// }

func ExampleFloatSlice_EachRE() {
	NewFloatSliceV(1, 2, 3).EachRE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 321
}

func TestFloatSlice_EachRE(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.EachRE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachRE(func(x O) error {
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{3, 2, 1}, results)
	}

	// Break early with error
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachRE(func(x O) error {
			if x.(float64) == 1 {
				return Break
			}
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{3, 2}, results)
	}
}

// EachRI
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_EachRI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_EachRI_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).EachRI(func(i int, x O) {
// 		assert.IsType(t, 0, x)
// 	})
// }

func ExampleFloatSlice_EachRI() {
	NewFloatSliceV(1, 2, 3).EachRI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 2:31:20:1
}

func TestFloatSlice_EachRI(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.EachRI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachRI(func(i int, x O) {
			results = append(results, x.(float64))
		})
		assert.Equal(t, []float64{3, 2, 1}, results)
	}
}

// EachRIE
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_EachRIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, 0, x)
// 	}
// 	for i := range Range(0, nines6) {
// 		action(i)
// 	}
// }

// func BenchmarkFloatSlice_EachRIE_Slice(t *testing.B) {
// 	NewFloatSlice(Range(0, nines6)).EachRIE(func(i int, x O) error {
// 		assert.IsType(t, 0, x)
// 		return nil
// 	})
// }

func ExampleFloatSlice_EachRIE() {
	NewFloatSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 2:31:20:1
}

func TestFloatSlice_EachRIE(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		slice.EachRIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{3, 2, 1}, results)
	}

	// Break early with error
	{
		results := []float64{}
		NewFloatSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
			if i == 0 {
				return Break
			}
			results = append(results, x.(float64))
			return nil
		})
		assert.Equal(t, []float64{3, 2}, results)
	}
}

// Empty
//--------------------------------------------------------------------------------------------------
func ExampleFloatSlice_Empty() {
	fmt.Println(NewFloatSliceV().Empty())
	// Output: true
}

func TestFloatSlice_Empty(t *testing.T) {

	// nil or empty
	{
		var nilSlice *FloatSlice
		assert.Equal(t, true, nilSlice.Empty())
	}

	assert.Equal(t, true, NewFloatSliceV().Empty())
	assert.Equal(t, false, NewFloatSliceV(1).Empty())
	assert.Equal(t, false, NewFloatSliceV(1, 2, 3).Empty())
	assert.Equal(t, false, NewFloatSliceV(1).Empty())
	assert.Equal(t, false, NewFloatSlice([]float64{1, 2, 3}).Empty())
}

// First
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_First_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 1 {
// 		_ = ints[0]
// 		ints = ints[1:]
// 	}
// }

// func BenchmarkFloatSlice_First_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.First()
// 		slice.DropFirst()
// 	}
// }

func ExampleFloatSlice_First() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.First())
	// Output: 1
}

func TestFloatSlice_First(t *testing.T) {
	// invalid
	assert.Equal(t, Obj(nil), NewFloatSliceV().First())

	// int
	assert.Equal(t, Obj(2.0), NewFloatSliceV(2, 3).First())
	assert.Equal(t, Obj(3.0), NewFloatSliceV(3, 2).First())
	assert.Equal(t, Obj(1.0), NewFloatSliceV(1, 3, 2).First())
}

// FirstN
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_FirstN_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	_ = ints[0:10]
// }

// func BenchmarkFloatSlice_FirstN_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	slice.FirstN(10)
// }

// func ExampleFloatSlice_FirstN() {
// 	slice := NewFloatSliceV(1, 2, 3)
// 	fmt.Println(slice.FirstN(2))
// 	// Output: [1.000000 2.000000]
// }

func TestFloatSlice_FirstN(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.FirstN(1))
		assert.Equal(t, NewFloatSliceV(), slice.FirstN(-1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewFloatSliceV(1, 2, 3)
		result := original.FirstN(2).Set(0, 0)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), original)
		assert.Equal(t, NewFloatSliceV(0, 2), result)
	}

	// Get none
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3).FirstN(0))

	// slice full array includeing out of bounds
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().FirstN(1))
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().FirstN(10))
	assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).FirstN(10))
	assert.Equal(t, NewFloatSlice([]float64{1, 2, 3}), NewFloatSlice([]float64{1, 2, 3}).FirstN(10))

	// grab a few diff
	assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1, 2, 3).FirstN(1))
	assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).FirstN(2))
}

// G
//--------------------------------------------------------------------------------------------------
func ExampleFloatSlice_G() {
	fmt.Println(NewFloatSliceV(1, 2, 3).G())
	// Output: [1 2 3]
}

func TestFloatSlice_G(t *testing.T) {
	assert.IsType(t, []float64{}, NewFloatSliceV().G())
	assert.IsType(t, []float64{1, 2, 3}, NewFloatSlice([]Object{{1}, {2}, {3}}).G())
}

// Generic
//--------------------------------------------------------------------------------------------------
func ExampleFloatSlice_Generic() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.InterSlice())
	// Output: false
}

// Index
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Index_Go(t *testing.B) {
// 	ints := Range(0, nines5)
// 	for i := range ints {
// 		if ints[i] == nines4 {
// 			break
// 		}
// 	}
// }

// func BenchmarkFloatSlice_Index_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines5))
// 	slice.Index(nines4)
// }

func ExampleFloatSlice_Index() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Index(2))
	// Output: 1
}

func TestFloatSlice_Index(t *testing.T) {

	// empty
	var slice *FloatSlice
	assert.Equal(t, -1, slice.Index(2))
	assert.Equal(t, -1, NewFloatSliceV().Index(1))

	assert.Equal(t, 0, NewFloatSliceV(1, 2, 3).Index(1))
	assert.Equal(t, 1, NewFloatSliceV(1, 2, 3).Index(2))
	assert.Equal(t, 2, NewFloatSliceV(1, 2, 3).Index(3))
	assert.Equal(t, -1, NewFloatSliceV(1, 2, 3).Index(4))
	assert.Equal(t, -1, NewFloatSliceV(1, 2, 3).Index(5))

	// Conversion
	{
		assert.Equal(t, 1, NewFloatSliceV(1, 2, 3).Index(Object{2}))
		assert.Equal(t, 1, NewFloatSliceV(1, 2, 3).Index("2"))
		assert.Equal(t, 0, NewFloatSliceV(1, 2, 3).Index(true))
		assert.Equal(t, 2, NewFloatSliceV(1, 2, 3).Index(Char('3')))
	}
}

// Insert
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Insert_Go(t *testing.B) {
// 	ints := []float64{}
// 	for i := range Range(0, nines6) {
// 		ints = append(ints, i)
// 		copy(ints[1:], ints[1:])
// 		ints[0] = i
// 	}
// }

// func BenchmarkFloatSlice_Insert_Slice(t *testing.B) {
// 	slice := NewFloatSliceV()
// 	for i := range Range(0, nines6) {
// 		slice.Insert(0, i)
// 	}
// }

func ExampleFloatSlice_Insert() {
	slice := NewFloatSliceV(1, 3)
	fmt.Println(slice.Insert(1, 2))
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Insert(t *testing.T) {

	// append
	{
		slice := NewFloatSliceV()
		assert.Equal(t, NewFloatSliceV(0), slice.Insert(-1, 0))
		assert.Equal(t, NewFloatSliceV(0, 1), slice.Insert(-1, 1))
		assert.Equal(t, NewFloatSliceV(0, 1, 2), slice.Insert(-1, 2))

	}

	// [] append
	{
		slice := NewFloatSliceV()
		assert.Equal(t, NewFloatSliceV(0), slice.Insert(-1, []float64{0}))
		assert.Equal(t, NewFloatSliceV(0, 1), slice.Insert(-1, []float64{1}))
		assert.Equal(t, NewFloatSliceV(0, 1, 2), slice.Insert(-1, []float64{2}))
	}

	// prepend
	{
		slice := NewFloatSliceV()
		assert.Equal(t, NewFloatSliceV(2), slice.Insert(0, 2))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Insert(0, 1))
		assert.Equal(t, NewFloatSliceV(0, 1, 2), slice.Insert(0, 0))
	}

	// [] prepend
	{
		slice := NewFloatSliceV()
		assert.Equal(t, NewFloatSliceV(2), slice.Insert(0, []float64{2}))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Insert(0, []float64{1}))
		assert.Equal(t, NewFloatSliceV(0, 1, 2), slice.Insert(0, []float64{0}))
	}

	// middle pos
	{
		slice := NewFloatSliceV(0, 5)
		assert.Equal(t, NewFloatSliceV(0, 1, 5), slice.Insert(1, 1))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 5), slice.Insert(2, 2))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 3, 5), slice.Insert(3, 3))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 3, 4, 5), slice.Insert(4, 4))
	}

	// [] middle pos
	{
		slice := NewFloatSliceV(0, 5)
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 5), slice.Insert(1, []float64{1, 2}))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 3, 4, 5), slice.Insert(3, []float64{3, 4}))
	}

	// middle neg
	{
		slice := NewFloatSliceV(0, 5)
		assert.Equal(t, NewFloatSliceV(0, 1, 5), slice.Insert(-2, 1))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 5), slice.Insert(-2, 2))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 3, 5), slice.Insert(-2, 3))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 3, 4, 5), slice.Insert(-2, 4))
	}

	// [] middle neg
	{
		slice := NewFloatSliceV(0, 5)
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 5), slice.Insert(-2, []float64{1, 2}))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 3, 4, 5), slice.Insert(-2, []float64{3, 4}))
	}

	// error cases
	{
		var slice *FloatSlice
		assert.False(t, slice.Insert(0, 0).Nil())
		assert.Equal(t, NewFloatSliceV(0), slice.Insert(0, 0))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(-10, 1))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(10, 1))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(2, 1))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(-3, 1))
	}

	// [] error cases
	{
		var slice *FloatSlice
		assert.False(t, slice.Insert(0, 0).Nil())
		assert.Equal(t, NewFloatSliceV(0), slice.Insert(0, 0))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(-10, 1))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(10, 1))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(2, 1))
		assert.Equal(t, NewFloatSliceV(0, 1), NewFloatSliceV(0, 1).Insert(-3, 1))
	}

	// Conversion
	{
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 3).Insert(1, Object{2}))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 3).Insert(1, "2"))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(2, 3).Insert(0, true))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2).Insert(-1, Char('3')))
	}

	// [] Conversion
	{
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), NewFloatSliceV(1, 4).Insert(1, []Object{{2}, {3}}))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), NewFloatSliceV(1, 4).Insert(1, []string{"2", "3"}))
		assert.Equal(t, NewFloatSliceV(0, 1, 2, 3), NewFloatSliceV(2, 3).Insert(0, []bool{false, true}))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), NewFloatSliceV(1, 2).Insert(-1, []Char{'3', '4'}))
	}
}

// Join
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Join_Go(t *testing.B) {
// 	ints := Range(0, nines4)
// 	strs := []string{}
// 	for i := 0; i < len(ints); i++ {
// 		strs = append(strs, fmt.Sprintf("%v", ints[i]))
// 	}
// 	strings.Join(strs, ",")
// }

// func BenchmarkFloatSlice_Join_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines4))
// 	slice.Join()
// }

func ExampleFloatSlice_Join() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Join())
	// Output: 1,2,3
}

func TestFloatSlice_Join(t *testing.T) {
	// nil
	{
		var slice *FloatSlice
		assert.Equal(t, Obj(""), slice.Join())
	}

	// empty
	{
		assert.Equal(t, Obj(""), NewFloatSliceV().Join())
	}

	assert.Equal(t, "1,2,3", NewFloatSliceV(1, 2, 3).Join().O())
	assert.Equal(t, "1.2.3", NewFloatSliceV(1, 2, 3).Join(".").O())
}

// Last
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Last_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 1 {
// 		_ = ints[len(ints)-1]
// 		ints = ints[:len(ints)-1]
// 	}
// }

// func BenchmarkFloatSlice_Last_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.Last()
// 		slice.DropLast()
// 	}
// }

func ExampleFloatSlice_Last() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Last())
	// Output: 3
}

func TestFloatSlice_Last(t *testing.T) {
	// invalid
	assert.Equal(t, Obj(nil), NewFloatSliceV().Last())

	// int
	assert.Equal(t, Obj(3.0), NewFloatSliceV(2, 3).Last())
	assert.Equal(t, Obj(2.0), NewFloatSliceV(3, 2).Last())
	assert.Equal(t, Obj(2.0), NewFloatSliceV(1, 3, 2).Last())
}

// LastN
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_LastN_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	_ = ints[0:10]
// }

// func BenchmarkFloatSlice_LastN_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	slice.LastN(10)
// }

func ExampleFloatSlice_LastN() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.LastN(2))
	// Output: [2.000000 3.000000]
}

func TestFloatSlice_LastN(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.LastN(1))
		assert.Equal(t, NewFloatSliceV(), slice.LastN(-1))
	}

	// Get none
	{
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3).LastN(0))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewFloatSliceV(1, 2, 3)
		result := original.LastN(2).Set(0, 0)
		assert.Equal(t, NewFloatSliceV(1, 0, 3), original)
		assert.Equal(t, NewFloatSliceV(0, 3), result)
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().LastN(1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().LastN(10))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).LastN(10))
		assert.Equal(t, NewFloatSlice([]float64{1, 2, 3}), NewFloatSlice([]float64{1, 2, 3}).LastN(10))
	}

	// grab a few diff
	{
		assert.Equal(t, NewFloatSliceV(3), NewFloatSliceV(1, 2, 3).LastN(1))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).LastN(2))
	}
}

// Len
//--------------------------------------------------------------------------------------------------
func ExampleFloatSlice_Len() {
	fmt.Println(NewFloatSliceV(1, 2, 3).Len())
	// Output: 3
}

func TestFloatSlice_Len(t *testing.T) {
	assert.Equal(t, 0, NewFloatSliceV().Len())
	assert.Equal(t, 2, len(*(NewFloatSliceV(1, 2))))
	assert.Equal(t, 2, NewFloatSliceV(1, 2).Len())
}

// Less
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Less_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i++ {
// 		if i+1 < len(ints) {
// 			_ = ints[i] < ints[i+1]
// 		}
// 	}
// }

// func BenchmarkFloatSlice_Less_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Less(i, i+1)
// 		}
// 	}
// }

func ExampleFloatSlice_Less() {
	slice := NewFloatSliceV(2, 3, 1)
	fmt.Println(slice.Less(0, 2))
	// Output: false
}

func TestFloatSlice_Less(t *testing.T) {

	// invalid cases
	{
		var slice *FloatSlice
		assert.False(t, slice.Less(0, 0))

		slice = NewFloatSliceV()
		assert.False(t, slice.Less(0, 0))
		assert.False(t, slice.Less(1, 2))
		assert.False(t, slice.Less(-1, 2))
		assert.False(t, slice.Less(1, -2))
	}

	// valid
	assert.Equal(t, true, NewFloatSliceV(0, 1, 2).Less(0, 1))
	assert.Equal(t, false, NewFloatSliceV(0, 1, 2).Less(1, 0))
	assert.Equal(t, true, NewFloatSliceV(0, 1, 2).Less(1, 2))
}

// Nil
//--------------------------------------------------------------------------------------------------
func ExampleFloatSlice_Nil() {
	var slice *FloatSlice
	fmt.Println(slice.Nil())
	// Output: true
}

func TestFloatSlice_Nil(t *testing.T) {
	var slice *FloatSlice
	assert.True(t, slice.Nil())
	assert.False(t, NewFloatSliceV().Nil())
	assert.False(t, NewFloatSliceV(1, 2, 3).Nil())
}

// O
//--------------------------------------------------------------------------------------------------
func ExampleFloatSlice_O() {
	fmt.Println(NewFloatSliceV(1, 2, 3))
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_O(t *testing.T) {
	assert.Equal(t, []float64{}, (*FloatSlice)(nil).O())
	assert.Equal(t, []float64{}, NewFloatSliceV().O())
	assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3))
}

// Pair
//--------------------------------------------------------------------------------------------------

func ExampleFloatSlice_Pair() {
	slice := NewFloatSliceV(1, 2)
	first, second := slice.Pair()
	fmt.Println(first, second)
	// Output: 1 2
}

func TestFloatSlice_Pair(t *testing.T) {

	// nil
	{
		first, second := (*FloatSlice)(nil).Pair()
		assert.Equal(t, Obj(nil), first)
		assert.Equal(t, Obj(nil), second)
	}

	// two values
	{
		first, second := NewFloatSliceV(1, 2).Pair()
		assert.Equal(t, Obj(1.0), first)
		assert.Equal(t, Obj(2.0), second)
	}

	// one value
	{
		first, second := NewFloatSliceV(1).Pair()
		assert.Equal(t, Obj(1.0), first)
		assert.Equal(t, Obj(nil), second)
	}

	// no values
	{
		first, second := NewFloatSliceV().Pair()
		assert.Equal(t, Obj(nil), first)
		assert.Equal(t, Obj(nil), second)
	}
}

// Pop
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Pop_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 1 {
// 		ints = ints[1:]
// 	}
// }

// func BenchmarkFloatSlice_Pop_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.Pop()
// 	}
// }

func ExampleFloatSlice_Pop() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Pop())
	// Output: 3
}

func TestFloatSlice_Pop(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, Obj(nil), slice.Pop())
	}

	// take all one at a time
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, Obj(3.0), slice.Pop())
		assert.Equal(t, NewFloatSliceV(1, 2), slice)
		assert.Equal(t, Obj(2.0), slice.Pop())
		assert.Equal(t, NewFloatSliceV(1), slice)
		assert.Equal(t, Obj(1.0), slice.Pop())
		assert.Equal(t, NewFloatSliceV(), slice)
		assert.Equal(t, Obj(nil), slice.Pop())
		assert.Equal(t, NewFloatSliceV(), slice)
	}
}

// PopN
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_PopN_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 10 {
// 		ints = ints[10:]
// 	}
// }

// func BenchmarkFloatSlice_PopN_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.PopN(10)
// 	}
// }

func ExampleFloatSlice_PopN() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.PopN(2))
	// Output: [2.000000 3.000000]
}

func TestFloatSlice_PopN(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.PopN(1))
	}

	// take none
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(), slice.PopN(0))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
	}

	// take 1
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(3), slice.PopN(1))
		assert.Equal(t, NewFloatSliceV(1, 2), slice)
	}

	// take 2
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(2, 3), slice.PopN(2))
		assert.Equal(t, NewFloatSliceV(1), slice)
	}

	// take 3
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice.PopN(3))
		assert.Equal(t, NewFloatSliceV(), slice)
	}

	// take beyond
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice.PopN(4))
		assert.Equal(t, NewFloatSliceV(), slice)
	}
}

// Prepend
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Prepend_Go(t *testing.B) {
// 	ints := []float64{}
// 	for i := range Range(0, nines6) {
// 		ints = append(ints, i)
// 		copy(ints[1:], ints[1:])
// 		ints[0] = i
// 	}
// }

// func BenchmarkFloatSlice_Prepend_Slice(t *testing.B) {
// 	slice := NewFloatSliceV()
// 	for i := range Range(0, nines6) {
// 		slice.Prepend(i)
// 	}
// }

func ExampleFloatSlice_Prepend() {
	slice := NewFloatSliceV(2, 3)
	fmt.Println(slice.Prepend(1))
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Prepend(t *testing.T) {

	// happy path
	{
		slice := NewFloatSliceV()
		assert.Equal(t, NewFloatSliceV(2), slice.Prepend(2))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Prepend(1))
		assert.Equal(t, NewFloatSliceV(0, 1, 2), slice.Prepend(0))
	}

	// error cases
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(0), slice.Prepend(0))
	}
}

// Reverse
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Reverse_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
// 		ints[i], ints[j] = ints[j], ints[i]
// 	}
// }

// func BenchmarkFloatSlice_Reverse_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.Reverse()
// }

func ExampleFloatSlice_Reverse() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Reverse())
	// Output: [3.000000 2.000000 1.000000]
}

func TestFloatSlice_Reverse(t *testing.T) {

	// nil
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.Reverse())
	}

	// empty
	{
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Reverse())
	}

	// pos
	{
		slice := NewFloatSliceV(3, 2, 1)
		reversed := slice.Reverse()
		assert.Equal(t, NewFloatSliceV(3, 2, 1, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), reversed)
	}

	// neg
	{
		slice := NewFloatSliceV(2, 3, -2, -3)
		reversed := slice.Reverse()
		assert.Equal(t, NewFloatSliceV(2, 3, -2, -3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(-3, -2, 3, 2), reversed)
	}
}

// ReverseM
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_ReverseM_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
// 		ints[i], ints[j] = ints[j], ints[i]
// 	}
// }

// func BenchmarkFloatSlice_ReverseM_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.ReverseM()
// }

func ExampleFloatSlice_ReverseM() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.ReverseM())
	// Output: [3.000000 2.000000 1.000000]
}

func TestFloatSlice_ReverseM(t *testing.T) {

	// nil
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.ReverseM())
	}

	// empty
	{
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().ReverseM())
	}

	// pos
	{
		slice := NewFloatSliceV(3, 2, 1)
		reversed := slice.ReverseM()
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), reversed)
	}

	// neg
	{
		slice := NewFloatSliceV(2, 3, -2, -3)
		reversed := slice.ReverseM()
		assert.Equal(t, NewFloatSliceV(-3, -2, 3, 2, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(-3, -2, 3, 2, 4), reversed)
	}
}

// Select
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Select_Go(t *testing.B) {
// 	even := []float64{}
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i++ {
// 		if ints[i]%2 == 0 {
// 			even = append(even, ints[i])
// 		}
// 	}
// }

// func BenchmarkFloatSlice_Select_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.Select(func(x O) bool {
// 		return ExB(x.(float64)%2 == 0)
// 	})
// }

func ExampleFloatSlice_Select() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Select(func(x O) bool {
		return ExB(x.(float64) == 2 || x.(float64) == 3)
	}))
	// Output: [2.000000 3.000000]
}

func TestFloatSlice_Select(t *testing.T) {

	// // Select all odd values
	// {
	// 	slice := NewFloatSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// 	new := slice.Select(func(x O) bool {
	// 		return ExB(x.(float64)%2 != 0)
	// 	})
	// 	slice.DropFirst()
	// 	assert.Equal(t, NewFloatSliceV(2, 3, 4, 5, 6, 7, 8, 9), slice)
	// 	assert.Equal(t, NewFloatSliceV(1, 3, 5, 7, 9), new)
	// }

	// // Select all even values
	// {
	// 	slice := NewFloatSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// 	new := slice.Select(func(x O) bool {
	// 		return ExB(x.(float64)%2 == 0)
	// 	})
	// 	slice.DropAt(1)
	// 	assert.Equal(t, NewFloatSliceV(1, 3, 4, 5, 6, 7, 8, 9), slice)
	// 	assert.Equal(t, NewFloatSliceV(2, 4, 6, 8), new)
	// }
}

// Set
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Set_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i++ {
// 		ints[i] = 0
// 	}
// }

// func BenchmarkFloatSlice_Set_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.Set(i, 0)
// 	}
// }

func ExampleFloatSlice_Set() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Set(0, 0))
	// Output: [0.000000 2.000000 3.000000]
}

func TestFloatSlice_Set(t *testing.T) {
	assert.Equal(t, NewFloatSliceV(0, 2, 3), NewFloatSliceV(1, 2, 3).Set(0, 0))
	assert.Equal(t, NewFloatSliceV(1, 0, 3), NewFloatSliceV(1, 2, 3).Set(1, 0))
	assert.Equal(t, NewFloatSliceV(1, 2, 0), NewFloatSliceV(1, 2, 3).Set(2, 0))
	assert.Equal(t, NewFloatSliceV(0, 2, 3), NewFloatSliceV(1, 2, 3).Set(-3, 0))
	assert.Equal(t, NewFloatSliceV(1, 0, 3), NewFloatSliceV(1, 2, 3).Set(-2, 0))
	assert.Equal(t, NewFloatSliceV(1, 2, 0), NewFloatSliceV(1, 2, 3).Set(-1, 0))

	// Test out of bounds
	{
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Set(5, 1))
	}

	// Test wrong type
	{
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Set(5, "1"))
	}

	// Conversion
	{
		assert.Equal(t, NewFloatSliceV(0, 2, 0), NewFloatSliceV(0, 0, 0).Set(1, Object{2}))
		assert.Equal(t, NewFloatSliceV(0, 2, 0), NewFloatSliceV(0, 0, 0).Set(1, "2"))
		assert.Equal(t, NewFloatSliceV(1, 0, 0), NewFloatSliceV(0, 0, 0).Set(0, true))
		assert.Equal(t, NewFloatSliceV(0, 0, 3), NewFloatSliceV(0, 0, 0).Set(-1, Char('3')))
	}
}

// SetE
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_SetE_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i++ {
// 		ints[i] = 0
// 	}
// }

// func BenchmarkFloatSlice_SetE_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.SetE(i, 0)
// 	}
// }

func ExampleFloatSlice_SetE() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.SetE(0, 0))
	// Output: [0.000000 2.000000 3.000000] <nil>
}

func TestFloatSlice_SetE(t *testing.T) {

	// pos - begining
	{
		slice := NewFloatSliceV(1, 2, 3)
		result, err := slice.SetE(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), slice)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), result)

		// multiple
		result, err = slice.SetE(0, []float64{4, 5, 6})
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(4, 5, 6), slice)
		assert.Equal(t, NewFloatSliceV(4, 5, 6), result)

		// multiple over
		result, err = slice.SetE(0, []float64{1, 2, 3, 4})
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), result)
	}

	// pos - middle
	{
		slice := NewFloatSliceV(1, 2, 3)
		result, err := slice.SetE(1, 0)
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(1, 0, 3), slice)
		assert.Equal(t, NewFloatSliceV(1, 0, 3), result)
	}

	// pos - end
	{
		slice := NewFloatSliceV(1, 2, 3)
		result, err := slice.SetE(2, 0)
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(1, 2, 0), slice)
		assert.Equal(t, NewFloatSliceV(1, 2, 0), result)
	}

	// neg - begining
	{
		slice := NewFloatSliceV(1, 2, 3)
		result, err := slice.SetE(-3, 0)
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), slice)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), result)

		// multiple
		result, err = slice.SetE(-3, []float64{4, 5, 6})
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(4, 5, 6), slice)
		assert.Equal(t, NewFloatSliceV(4, 5, 6), result)

		// multiple over
		result, err = slice.SetE(-3, []float64{1, 2, 3, 4})
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), result)
	}

	// neg - middle
	{
		slice := NewFloatSliceV(1, 2, 3)
		result, err := slice.SetE(-2, 0)
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(1, 0, 3), slice)
		assert.Equal(t, NewFloatSliceV(1, 0, 3), result)
	}

	// neg - end
	{
		slice := NewFloatSliceV(1, 2, 3)
		result, err := slice.SetE(-1, 0)
		assert.Nil(t, err)
		assert.Equal(t, NewFloatSliceV(1, 2, 0), slice)
		assert.Equal(t, NewFloatSliceV(1, 2, 0), result)
	}

	// Test out of bounds
	{
		slice, err := NewFloatSliceV(1, 2, 3).SetE(5, 1)
		assert.NotNil(t, slice)
		assert.NotNil(t, err)
	}

	// Test wrong type
	{
		slice, err := NewFloatSliceV(1, 2, 3).SetE(5, "1")
		assert.NotNil(t, slice)
		assert.NotNil(t, err)
	}
}

// Shift
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Shift_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 1 {
// 		ints = ints[1:]
// 	}
// }

// func BenchmarkFloatSlice_Shift_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.Shift()
// 	}
// }

func ExampleFloatSlice_Shift() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Shift())
	// Output: 1
}

func TestFloatSlice_Shift(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, Obj(nil), slice.Shift())
	}

	// take all and beyond
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, Obj(1.0), slice.Shift())
		assert.Equal(t, NewFloatSliceV(2, 3), slice)
		assert.Equal(t, Obj(2.0), slice.Shift())
		assert.Equal(t, NewFloatSliceV(3), slice)
		assert.Equal(t, Obj(3.0), slice.Shift())
		assert.Equal(t, NewFloatSliceV(), slice)
		assert.Equal(t, Obj(nil), slice.Shift())
		assert.Equal(t, NewFloatSliceV(), slice)
	}
}

// ShiftN
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_ShiftN_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 10 {
// 		ints = ints[10:]
// 	}
// }

// func BenchmarkFloatSlice_ShiftN_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 0 {
// 		slice.ShiftN(10)
// 	}
// }

func ExampleFloatSlice_ShiftN() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.ShiftN(2))
	// Output: [1.000000 2.000000]
}

func TestFloatSlice_ShiftN(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.ShiftN(1))
	}

	// negative value
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(1), slice.ShiftN(-1))
		assert.Equal(t, NewFloatSliceV(2, 3), slice)
	}

	// take none
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(), slice.ShiftN(0))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
	}

	// take 1
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(1), slice.ShiftN(1))
		assert.Equal(t, NewFloatSliceV(2, 3), slice)
	}

	// take 2
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(1, 2), slice.ShiftN(2))
		assert.Equal(t, NewFloatSliceV(3), slice)
	}

	// take 3
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice.ShiftN(3))
		assert.Equal(t, NewFloatSliceV(), slice)
	}

	// take beyond
	{
		slice := NewFloatSliceV(1, 2, 3)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice.ShiftN(4))
		assert.Equal(t, NewFloatSliceV(), slice)
	}
}

// Single
//--------------------------------------------------------------------------------------------------

// func ExampleFloatSlice_Single() {
// 	slice := NewFloatSliceV(1)
// 	fmt.Println(slice.Single())
// 	// Output: true
// }

// func TestFloatSlice_Single(t *testing.T) {

// 	assert.Equal(t, false, NewFloatSliceV().Single())
// 	assert.Equal(t, true, NewFloatSliceV(1).Single())
// 	assert.Equal(t, false, NewFloatSliceV(1, 2).Single())
// }

// Slice
//--------------------------------------------------------------------------------------------------
func BenchmarkFloatSlice_Slice_Go(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:len(ints)]
}

func BenchmarkFloatSlice_Slice_Slice(t *testing.B) {
	slice := NewFloatSlice(Range(0, nines7))
	slice.Slice(0, -1)
}

func ExampleFloatSlice_Slice() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Slice(1, -1))
	// Output: [2.000000 3.000000]
}

func TestFloatSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.Slice(0, -1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Slice(0, -1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewFloatSliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), original)
		assert.Equal(t, NewFloatSliceV(0, 2, 3), result)
	}

	// slice full array
	{
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Slice(0, -1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Slice(0, 1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Slice(0, 5))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Slice(0, -1))
		assert.Equal(t, NewFloatSlice([]float64{1, 2, 3}), NewFloatSlice([]float64{1, 2, 3}).Slice(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1).Slice(0, 2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Slice(-6, 6))
	}

	// mutually exclusive
	{
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Slice(2, -3))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Slice(0, -5))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Slice(4, -1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Slice(6, -1))
		assert.Equal(t, NewFloatSliceV(), NewFloatSliceV(1, 2, 3, 4).Slice(3, 2))
	}

	// singles
	{
		slice := NewFloatSliceV(1, 2, 3, 4)
		assert.Equal(t, NewFloatSliceV(4), slice.Slice(-1, -1))
		assert.Equal(t, NewFloatSliceV(3), slice.Slice(-2, -2))
		assert.Equal(t, NewFloatSliceV(2), slice.Slice(-3, -3))
		assert.Equal(t, NewFloatSliceV(1), slice.Slice(0, 0))
		assert.Equal(t, NewFloatSliceV(1), slice.Slice(-4, -4))
		assert.Equal(t, NewFloatSliceV(2), slice.Slice(1, 1))
		assert.Equal(t, NewFloatSliceV(2), slice.Slice(1, -3))
		assert.Equal(t, NewFloatSliceV(3), slice.Slice(2, 2))
		assert.Equal(t, NewFloatSliceV(3), slice.Slice(2, -2))
		assert.Equal(t, NewFloatSliceV(4), slice.Slice(3, 3))
		assert.Equal(t, NewFloatSliceV(4), slice.Slice(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Slice(1, -1))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Slice(-2, -1))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Slice(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Slice(0, -2))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Slice(-3, -2))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Slice(-3, 1))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2, 3).Slice(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Slice(1, -2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Slice(-3, -2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Slice(-3, 2))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3, 4).Slice(1, 2))
	}

	// random
	{
		assert.Equal(t, NewFloatSliceV(1), NewFloatSliceV(1, 2, 3).Slice(0, -3))
		assert.Equal(t, NewFloatSliceV(2, 3), NewFloatSliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2, 3).Slice(0, 2))
	}
}

// Sort
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Sort_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	sort.Sort(sort.FloatSlice(ints))
// }

// func BenchmarkFloatSlice_Sort_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.Sort()
// }

func ExampleFloatSlice_Sort() {
	slice := NewFloatSliceV(2, 3, 1)
	fmt.Println(slice.Sort())
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Sort(t *testing.T) {

	// empty
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().Sort())

	// pos
	{
		slice := NewFloatSliceV(5, 3, 2, 4, 1)
		sorted := slice.Sort()
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(5, 3, 2, 4, 1), slice)
	}

	// neg
	{
		slice := NewFloatSliceV(5, 3, -2, 4, -1)
		sorted := slice.Sort()
		assert.Equal(t, NewFloatSliceV(-2, -1, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(5, 3, -2, 4, -1), slice)
	}
}

// SortM
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_SortM_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	sort.Sort(sort.FloatSlice(ints))
// }

// func BenchmarkFloatSlice_SortM_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.SortM()
// }

func ExampleFloatSlice_SortM() {
	slice := NewFloatSliceV(2, 3, 1)
	fmt.Println(slice.SortM())
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_SortM(t *testing.T) {

	// empty
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().SortM())

	// pos
	{
		slice := NewFloatSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortM()
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4, 5, 6), slice)
	}

	// neg
	{
		slice := NewFloatSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortM()
		assert.Equal(t, NewFloatSliceV(-2, -1, 3, 4, 5, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(-2, -1, 3, 4, 5, 6), slice)
	}
}

// SortReverse
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_SortReverse_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	sort.Sort(sort.Reverse(sort.FloatSlice(ints)))
// }

// func BenchmarkFloatSlice_SortReverse_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.SortReverse()
// }

func ExampleFloatSlice_SortReverse() {
	slice := NewFloatSliceV(2, 3, 1)
	fmt.Println(slice.SortReverse())
	// Output: [3.000000 2.000000 1.000000]
}

func TestFloatSlice_SortReverse(t *testing.T) {

	// empty
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().SortReverse())

	// pos
	{
		slice := NewFloatSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverse()
		assert.Equal(t, NewFloatSliceV(5, 4, 3, 2, 1, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(5, 3, 2, 4, 1), slice)
	}

	// neg
	{
		slice := NewFloatSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverse()
		assert.Equal(t, NewFloatSliceV(5, 4, 3, -1, -2, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(5, 3, -2, 4, -1), slice)
	}
}

// SortReverseM
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_SortReverseM_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	sort.Sort(sort.Reverse(sort.FloatSlice(ints)))
// }

// func BenchmarkFloatSlice_SortReverseM_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	slice.SortReverseM()
// }

func ExampleFloatSlice_SortReverseM() {
	slice := NewFloatSliceV(2, 3, 1)
	fmt.Println(slice.SortReverseM())
	// Output: [3.000000 2.000000 1.000000]
}

func TestFloatSlice_SortReverseM(t *testing.T) {

	// empty
	assert.Equal(t, NewFloatSliceV(), NewFloatSliceV().SortReverse())

	// pos
	{
		slice := NewFloatSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverseM()
		assert.Equal(t, NewFloatSliceV(5, 4, 3, 2, 1, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(5, 4, 3, 2, 1, 6), slice)
	}

	// neg
	{
		slice := NewFloatSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverseM()
		assert.Equal(t, NewFloatSliceV(5, 4, 3, -1, -2, 6), sorted.Append(6))
		assert.Equal(t, NewFloatSliceV(5, 4, 3, -1, -2, 6), slice)
	}
}

// String
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_String_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	_ = fmt.Sprintf("%v", ints)
// }

// func BenchmarkFloatSlice_String_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	_ = slice.String()
// }

func ExampleFloatSlice_String() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_String(t *testing.T) {

	// nil
	assert.Equal(t, "[]", (*FloatSlice)(nil).String())

	// empty
	assert.Equal(t, "[]", NewFloatSliceV().String())

	// pos
	{
		slice := NewFloatSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverseM()
		assert.Equal(t, "[5.000000 4.000000 3.000000 2.000000 1.000000 6.000000]", sorted.Append(6).String())
		assert.Equal(t, "[5.000000 4.000000 3.000000 2.000000 1.000000 6.000000]", slice.String())
	}

	// neg
	{
		slice := NewFloatSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverseM()
		assert.Equal(t, "[5.000000 4.000000 3.000000 -1.000000 -2.000000 6.000000]", sorted.Append(6).String())
		assert.Equal(t, "[5.000000 4.000000 3.000000 -1.000000 -2.000000 6.000000]", slice.String())
	}
}

// Swap
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Swap_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i++ {
// 		if i+1 < len(ints) {
// 			ints[i], ints[i+1] = ints[i+1], ints[i]
// 		}
// 	}
// }

// func BenchmarkFloatSlice_Swap_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Swap(i, i+1)
// 		}
// 	}
// }

func ExampleFloatSlice_Swap() {
	slice := NewFloatSliceV(2, 3, 1)
	slice.Swap(0, 2)
	slice.Swap(1, 2)
	fmt.Println(slice)
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Swap(t *testing.T) {

	// invalid cases
	{
		var slice *FloatSlice
		slice.Swap(0, 0)
		assert.Equal(t, (*FloatSlice)(nil), slice)

		slice = NewFloatSliceV()
		slice.Swap(0, 0)
		assert.Equal(t, NewFloatSliceV(), slice)

		slice.Swap(1, 2)
		assert.Equal(t, NewFloatSliceV(), slice)

		slice.Swap(-1, 2)
		assert.Equal(t, NewFloatSliceV(), slice)

		slice.Swap(1, -2)
		assert.Equal(t, NewFloatSliceV(), slice)
	}

	// normal
	{
		slice := NewFloatSliceV(0, 1, 2)
		slice.Swap(0, 1)
		assert.Equal(t, NewFloatSliceV(1, 0, 2), slice)
	}
}

// Take
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Take_Go(t *testing.B) {
// 	ints := Range(0, nines7)
// 	for len(ints) > 11 {
// 		i := 1
// 		n := 10
// 		if i+n < len(ints) {
// 			ints = append(ints[:i], ints[i+n:]...)
// 		} else {
// 			ints = ints[:i]
// 		}
// 	}
// }

// func BenchmarkFloatSlice_Take_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines7))
// 	for slice.Len() > 1 {
// 		slice.Take(1, 10)
// 	}
// }

func ExampleFloatSlice_Take() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.Take(0, 1))
	// Output: [1.000000 2.000000]
}

func TestFloatSlice_Take(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.Take(0, 1))
	}

	// invalid
	{
		slice := NewFloatSliceV(1, 2, 3, 4)
		assert.Equal(t, NewFloatSliceV(), slice.Take(1))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice)
		assert.Equal(t, NewFloatSliceV(), slice.Take(4, 4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice)
	}

	// take 1
	{
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1), slice.Take(0, 0))
			assert.Equal(t, NewFloatSliceV(2, 3, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(2), slice.Take(1, 1))
			assert.Equal(t, NewFloatSliceV(1, 3, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(3), slice.Take(2, 2))
			assert.Equal(t, NewFloatSliceV(1, 2, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(4), slice.Take(3, 3))
			assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(4), slice.Take(-1, -1))
			assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(3), slice.Take(-2, -2))
			assert.Equal(t, NewFloatSliceV(1, 2, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(2), slice.Take(-3, -3))
			assert.Equal(t, NewFloatSliceV(1, 3, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1), slice.Take(-4, -4))
			assert.Equal(t, NewFloatSliceV(2, 3, 4), slice)
		}
	}

	// take 2
	{
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2), slice.Take(0, 1))
			assert.Equal(t, NewFloatSliceV(3, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(2, 3), slice.Take(1, 2))
			assert.Equal(t, NewFloatSliceV(1, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(3, 4), slice.Take(2, 3))
			assert.Equal(t, NewFloatSliceV(1, 2), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(3, 4), slice.Take(-2, -1))
			assert.Equal(t, NewFloatSliceV(1, 2), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(2, 3), slice.Take(-3, -2))
			assert.Equal(t, NewFloatSliceV(1, 4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2), slice.Take(-4, -3))
			assert.Equal(t, NewFloatSliceV(3, 4), slice)
		}
	}

	// take 3
	{
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2, 3), slice.Take(0, 2))
			assert.Equal(t, NewFloatSliceV(4), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(2, 3, 4), slice.Take(-3, -1))
			assert.Equal(t, NewFloatSliceV(1), slice)
		}
	}

	// take everything and beyond
	{
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Take())
			assert.Equal(t, NewFloatSliceV(), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Take(0, 3))
			assert.Equal(t, NewFloatSliceV(), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Take(0, -1))
			assert.Equal(t, NewFloatSliceV(), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Take(-4, -1))
			assert.Equal(t, NewFloatSliceV(), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Take(-6, -1))
			assert.Equal(t, NewFloatSliceV(), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Take(0, 10))
			assert.Equal(t, NewFloatSliceV(), slice)
		}
	}

	// move index within bounds
	{
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(4), slice.Take(3, 4))
			assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
		}
		{
			slice := NewFloatSliceV(1, 2, 3, 4)
			assert.Equal(t, NewFloatSliceV(1), slice.Take(-5, 0))
			assert.Equal(t, NewFloatSliceV(2, 3, 4), slice)
		}
	}
}

// TakeAt
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_TakeAt_Go(t *testing.B) {
// 	ints := Range(0, nines5)
// 	index := Range(0, nines5)
// 	for i := range index {
// 		if i+1 < len(ints) {
// 			ints = append(ints[:i], ints[i+1:]...)
// 		} else if i >= 0 && i < len(ints) {
// 			ints = ints[:i]
// 		}
// 	}
// }

// func BenchmarkFloatSlice_TakeAt_Slice(t *testing.B) {
// 	src := Range(0, nines5)
// 	index := Range(0, nines5)
// 	slice := NewFloatSlice(src)
// 	for i := range index {
// 		slice.TakeAt(i)
// 	}
// }

func ExampleFloatSlice_TakeAt() {
	slice := NewFloatSliceV(1, 2, 3)
	fmt.Println(slice.TakeAt(1))
	// Output: 2
}

func TestFloatSlice_TakeAt(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, Obj(nil), slice.TakeAt(0))
	}

	// all and more
	{
		slice := NewFloatSliceV(0, 1, 2)
		assert.Equal(t, Obj(2.0), slice.TakeAt(-1))
		assert.Equal(t, NewFloatSliceV(0, 1), slice)
		assert.Equal(t, Obj(1.0), slice.TakeAt(-1))
		assert.Equal(t, NewFloatSliceV(0), slice)
		assert.Equal(t, Obj(0.0), slice.TakeAt(-1))
		assert.Equal(t, NewFloatSliceV(), slice)
		assert.Equal(t, Obj(nil), slice.TakeAt(-1))
		assert.Equal(t, NewFloatSliceV(), slice)
	}

	// take invalid
	{
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(nil), slice.TakeAt(3))
			assert.Equal(t, NewFloatSliceV(0, 1, 2), slice)
		}
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(nil), slice.TakeAt(-4))
			assert.Equal(t, NewFloatSliceV(0, 1, 2), slice)
		}
	}

	// take last
	{
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(2.0), slice.TakeAt(2))
			assert.Equal(t, NewFloatSliceV(0, 1), slice)
		}
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(2.0), slice.TakeAt(-1))
			assert.Equal(t, NewFloatSliceV(0, 1), slice)
		}
	}

	// take middle
	{
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(1.0), slice.TakeAt(1))
			assert.Equal(t, NewFloatSliceV(0, 2), slice)
		}
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(1.0), slice.TakeAt(-2))
			assert.Equal(t, NewFloatSliceV(0, 2), slice)
		}
	}

	// take first
	{
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(0.0), slice.TakeAt(0))
			assert.Equal(t, NewFloatSliceV(1, 2), slice)
		}
		{
			slice := NewFloatSliceV(0, 1, 2)
			assert.Equal(t, Obj(0.0), slice.TakeAt(-3))
			assert.Equal(t, NewFloatSliceV(1, 2), slice)
		}
	}
}

// TakeW
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_TakeW_Go(t *testing.B) {
// 	new := []float64{}
// 	ints := Range(0, nines5)
// 	l := len(ints)
// 	for i := 0; i < l; i++ {
// 		if ints[i]%2 == 0 {
// 			new = append(new, ints[i])
// 			if i+1 < l {
// 				ints = append(ints[:i], ints[i+1:]...)
// 			} else if i >= 0 && i < l {
// 				ints = ints[:i]
// 			}
// 			l--
// 			i--
// 		}
// 	}
// }

// func BenchmarkFloatSlice_TakeW_Slice(t *testing.B) {
// 	slice := NewFloatSlice(Range(0, nines5))
// 	slice.TakeW(func(x O) bool { return ExB(x.(float64)%2 == 0) })
// // }

// func ExampleFloatSlice_TakeW() {
// 	slice := NewFloatSliceV(1, 2, 3)
// 	fmt.Println(slice.TakeW(func(x O) bool {
// 		return ExB(x.(float64)%2 == 0)
// 	}))
// 	// Output: [2.000000]
// }

// func TestFloatSlice_TakeW(t *testing.T) {

// 	// take all odd values
// 	{
// 		slice := NewFloatSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
// 		new := slice.TakeW(func(x O) bool { return ExB(x.(float64)%2 != 0) })
// 		assert.Equal(t, NewFloatSliceV(2, 4, 6, 8), slice)
// 		assert.Equal(t, NewFloatSliceV(1, 3, 5, 7, 9), new)
// 	}

// 	// take all even values
// 	{
// 		slice := NewFloatSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
// 		new := slice.TakeW(func(x O) bool { return ExB(x.(float64)%2 == 0) })
// 		assert.Equal(t, NewFloatSliceV(1, 3, 5, 7, 9), slice)
// 		assert.Equal(t, NewFloatSliceV(2, 4, 6, 8), new)
// 	}
// }

// Union
//--------------------------------------------------------------------------------------------------
// func BenchmarkFloatSlice_Union_Go(t *testing.B) {
// 	// ints := Range(0, nines7)
// 	// for len(ints) > 10 {
// 	// 	ints = ints[10:]
// 	// }
// }

// func BenchmarkFloatSlice_Union_Slice(t *testing.B) {
// 	// slice := NewFloatSlice(Range(0, nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

func ExampleFloatSlice_Union() {
	slice := NewFloatSliceV(1, 2)
	fmt.Println(slice.Union([]float64{2, 3}))
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Union(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Union(NewFloatSliceV(1, 2)))
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Union([]float64{1, 2}))
	}

	// size of one
	{
		slice := NewFloatSliceV(1)
		union := slice.Union([]float64{1, 2, 3})
		assert.Equal(t, NewFloatSliceV(1, 2, 3), union)
		assert.Equal(t, NewFloatSliceV(1), slice)
	}

	// one duplicate
	{
		slice := NewFloatSliceV(1, 1)
		union := slice.Union(NewFloatSliceV(2, 3))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), union)
		assert.Equal(t, NewFloatSliceV(1, 1), slice)
	}

	// multiple duplicates
	{
		slice := NewFloatSliceV(1, 2, 2, 3, 3)
		union := slice.Union([]float64{1, 2, 3})
		assert.Equal(t, NewFloatSliceV(1, 2, 3), union)
		assert.Equal(t, NewFloatSliceV(1, 2, 2, 3, 3), slice)
	}

	// no duplicates
	{
		slice := NewFloatSliceV(1, 2, 3)
		union := slice.Union([]float64{4, 5})
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4, 5), union)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
	}

	// nils
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).Union((*[]float64)(nil)))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).Union((*FloatSlice)(nil)))
	}

	// Conversion
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).Union([]Object{{1}, {2}}))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2).Union([]string{"2", "3"}))
		assert.Equal(t, NewFloatSliceV(1, 2, 0), NewFloatSliceV(1, 2).Union([]bool{true, false}))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), NewFloatSliceV(1, 2).Union([]Char{'3', '2'}))
	}
}

// UnionM
//--------------------------------------------------------------------------------------------------
func BenchmarkFloatSlice_UnionM_Go(t *testing.B) {
	// ints := Range(0, nines7)
	// for len(ints) > 10 {
	// 	ints = ints[10:]
	// }
}

func BenchmarkFloatSlice_UnionM_Slice(t *testing.B) {
	// slice := NewFloatSlice(Range(0, nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleFloatSlice_UnionM() {
	slice := NewFloatSliceV(1, 2)
	fmt.Println(slice.UnionM([]float64{2, 3}))
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_UnionM(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(1, 2), slice.UnionM(NewFloatSliceV(1, 2)))
		assert.Equal(t, (*FloatSlice)(nil), slice)
	}

	// size of one
	{
		slice := NewFloatSliceV(1)
		union := slice.UnionM([]float64{1, 2, 3})
		assert.Equal(t, NewFloatSliceV(1, 2, 3), union)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
	}

	// one duplicate
	{
		slice := NewFloatSliceV(1, 1)
		union := slice.UnionM(NewFloatSliceV(2, 3))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), union)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
	}

	// multiple duplicates
	{
		slice := NewFloatSliceV(1, 2, 2, 3, 3)
		union := slice.UnionM([]float64{1, 2, 3})
		assert.Equal(t, NewFloatSliceV(1, 2, 3), union)
		assert.Equal(t, NewFloatSliceV(1, 2, 3), slice)
	}

	// no duplicates
	{
		slice := NewFloatSliceV(1, 2, 3)
		union := slice.UnionM([]float64{4, 5})
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4, 5), union)
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4, 5), slice)
	}

	// nils
	{
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).UnionM((*[]float64)(nil)))
		assert.Equal(t, NewFloatSliceV(1, 2), NewFloatSliceV(1, 2).UnionM((*FloatSlice)(nil)))
	}
}

// Uniq
//--------------------------------------------------------------------------------------------------
func BenchmarkFloatSlice_Uniq_Go(t *testing.B) {
	// ints := Range(0, nines7)
	// for len(ints) > 10 {
	// 	ints = ints[10:]
	// }
}

func BenchmarkFloatSlice_Uniq_Slice(t *testing.B) {
	// slice := NewFloatSlice(Range(0, nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleFloatSlice_Uniq() {
	slice := NewFloatSliceV(1, 2, 3, 3)
	fmt.Println(slice.Uniq())
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_Uniq(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, NewFloatSliceV(), slice.Uniq())
	}

	// size of one
	{
		slice := NewFloatSliceV(1)
		uniq := slice.Uniq()
		assert.Equal(t, NewFloatSliceV(1), uniq)
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1), uniq)
	}

	// one duplicate
	{
		slice := NewFloatSliceV(1, 1)
		uniq := slice.Uniq()
		assert.Equal(t, NewFloatSliceV(1), uniq)
		assert.Equal(t, NewFloatSliceV(1, 1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1), uniq)
	}

	// multiple duplicates
	{
		slice := NewFloatSliceV(1, 2, 2, 3, 3)
		uniq := slice.Uniq()
		assert.Equal(t, NewFloatSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewFloatSliceV(1, 2, 2, 3, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), uniq)
	}

	// no duplicates
	{
		slice := NewFloatSliceV(1, 2, 3)
		uniq := slice.Uniq()
		assert.Equal(t, NewFloatSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3), uniq)
	}
}

// UniqM
//--------------------------------------------------------------------------------------------------
func BenchmarkFloatSlice_UniqM_Go(t *testing.B) {
	// ints := Range(0, nines7)
	// for len(ints) > 10 {
	// 	ints = ints[10:]
	// }
}

func BenchmarkFloatSlice_UniqM_Slice(t *testing.B) {
	// slice := NewFloatSlice(Range(0, nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleFloatSlice_UniqM() {
	slice := NewFloatSliceV(1, 2, 3, 3)
	fmt.Println(slice.UniqM())
	// Output: [1.000000 2.000000 3.000000]
}

func TestFloatSlice_UniqM(t *testing.T) {

	// nil or empty
	{
		var slice *FloatSlice
		assert.Equal(t, (*FloatSlice)(nil), slice.UniqM())
	}

	// size of one
	{
		slice := NewFloatSliceV(1)
		uniq := slice.UniqM()
		assert.Equal(t, NewFloatSliceV(1), uniq)
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1, 2), uniq)
	}

	// one duplicate
	{
		slice := NewFloatSliceV(1, 1)
		uniq := slice.UniqM()
		assert.Equal(t, NewFloatSliceV(1), uniq)
		assert.Equal(t, NewFloatSliceV(1, 2), slice.Append(2))
		assert.Equal(t, NewFloatSliceV(1, 2), uniq)
	}

	// multiple duplicates
	{
		slice := NewFloatSliceV(1, 2, 2, 3, 3)
		uniq := slice.UniqM()
		assert.Equal(t, NewFloatSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), uniq)
	}

	// no duplicates
	{
		slice := NewFloatSliceV(1, 2, 3)
		uniq := slice.UniqM()
		assert.Equal(t, NewFloatSliceV(1, 2, 3), uniq)
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), slice.Append(4))
		assert.Equal(t, NewFloatSliceV(1, 2, 3, 4), uniq)
	}
}
