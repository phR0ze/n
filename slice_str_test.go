package n

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// // NewStrSlice
// //--------------------------------------------------------------------------------------------------
// func BenchmarkNewStrSlice_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = []string{src[i], src[i] + string(1), src[i] + string(2), src[i] + string(3), src[i] + string(4), src[i] + string(5), src[i] + string(6), src[i] + string(7), src[i] + string(8), src[i] + string(9)}
// 	}
// }

// func BenchmarkNewStrSlice_Slice(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = NewStrSlice([]string{src[i], src[i] + string(1), src[i] + string(2), src[i] + string(3), src[i] + string(4), src[i] + string(5), src[i] + string(6), src[i] + string(7), src[i] + string(8), src[i] + string(9)})
// 	}
// }

// func ExampleNewStrSlice() {
// 	slice := NewStrSlice([]string{"1", "2", "3"})
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_NewStrSlice(t *testing.T) {

// 	// array
// 	var array [2]string
// 	array[0] = "1"
// 	array[1] = "2"
// 	assert.Equal(t, []string{"1", "2"}, NewStrSlice(array[:]).O())

// 	// empty
// 	assert.Equal(t, []string{}, NewStrSlice([]string{}).O())

// 	// slice
// 	assert.Equal(t, []string{"0"}, NewStrSlice([]string{"0"}).O())
// 	assert.Equal(t, []string{"1", "2"}, NewStrSlice([]string{"1", "2"}).O())
// }

// // NewStrSliceV
// //--------------------------------------------------------------------------------------------------
// func BenchmarkNewStrSliceV_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = append([]string{}, src[i], src[i]+string(1), src[i]+string(2), src[i]+string(3), src[i]+string(4), src[i]+string(5), src[i]+string(6), src[i]+string(7), src[i]+string(8), src[i]+string(9))
// 	}
// }

// func BenchmarkNewStrSliceV_Slice(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = NewStrSliceV(src[i], src[i]+string(1), src[i]+string(2), src[i]+string(3), src[i]+string(4), src[i]+string(5), src[i]+string(6), src[i]+string(7), src[i]+string(8), src[i]+string(9))
// 	}
// }

// func ExampleNewStrSliceV_empty() {
// 	slice := NewStrSliceV()
// 	fmt.Println(slice)
// 	// Output: []
// }

// func ExampleNewStrSliceV_variadic() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_NewStrSliceV(t *testing.T) {

// 	// array
// 	var array [2]string
// 	array[0] = "1"
// 	array[1] = "2"
// 	assert.Equal(t, []string{"1", "2"}, NewStrSliceV(array[:]...).O())

// 	// empty
// 	assert.Equal(t, []string{}, NewStrSliceV().O())

// 	// multiples
// 	assert.Equal(t, []string{"1"}, NewStrSliceV("1").O())
// 	assert.Equal(t, []string{"1", "2"}, NewStrSliceV("1", "2").O())
// 	assert.Equal(t, []string{"1", "2"}, NewStrSliceV([]string{"1", "2"}...).O())
// }

// // Any
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Any_Go(t *testing.B) {
// 	any := func(list []string, x []string) bool {
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
// 	src := RangeStr(nines4)
// 	for _, x := range src {
// 		any(src, []string{x})
// 	}
// }

// func BenchmarkStrSlice_Any_Slice(t *testing.B) {
// 	src := RangeStr(nines4)
// 	slice := NewStrSlice(src)
// 	for i := range src {
// 		slice.Any(i)
// 	}
// }

// func ExampleStrSlice_Any_empty() {
// 	slice := NewStrSliceV()
// 	fmt.Println(slice.Any())
// 	// Output: false
// }

// func ExampleStrSlice_Any_notEmpty() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Any())
// 	// Output: true
// }

// func ExampleStrSlice_Any_contains() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Any("1"))
// 	// Output: true
// }

// func ExampleStrSlice_Any_containsAny() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Any("0", "1"))
// 	// Output: true
// }

// func TestStrSlice_Any(t *testing.T) {

// 	// empty
// 	var nilSlice *StrSlice
// 	assert.False(t, nilSlice.Any())
// 	assert.False(t, NewStrSliceV().Any())

// 	// single
// 	assert.True(t, NewStrSliceV("2").Any())

// 	// invalid
// 	assert.False(t, NewStrSliceV("1", "2").Any(Object{"2"}))

// 	assert.True(t, NewStrSliceV("1", "2", "3").Any("2"))
// 	assert.False(t, NewStrSliceV("1", "2", "3").Any(4))
// 	assert.True(t, NewStrSliceV("1", "2", "3").Any(4, "3"))
// 	assert.False(t, NewStrSliceV("1", "2", "3").Any(4, 5))
// }

// // AnyS
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_AnyS_Go(t *testing.B) {
// 	any := func(list []string, x []string) bool {
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
// 	src := RangeStr(nines4)
// 	for _, x := range src {
// 		any(src, []string{x})
// 	}
// }

// func BenchmarkStrSlice_AnyS_Slice(t *testing.B) {
// 	src := RangeStr(nines4)
// 	slice := NewStrSlice(src)
// 	for _, x := range src {
// 		slice.Any([]string{x})
// 	}
// }

// func ExampleStrSlice_AnyS() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.AnyS([]string{"0", "1"}))
// 	// Output: true
// }

// func TestStrSlice_AnyS(t *testing.T) {
// 	// nil
// 	{
// 		var slice *StrSlice
// 		assert.False(t, slice.AnyS([]string{"1"}))
// 		assert.False(t, NewStrSliceV("1").AnyS(nil))
// 	}

// 	// []string
// 	{
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS([]string{"1"}))
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS([]string{"4", "3"}))
// 		assert.False(t, NewStrSliceV("1", "2", "3").AnyS([]string{"4", "5"}))
// 	}

// 	// *[]string
// 	{
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(&([]string{"1"})))
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(&([]string{"4", "3"})))
// 		assert.False(t, NewStrSliceV("1", "2", "3").AnyS(&([]string{"4", "5"})))
// 	}

// 	// Slice
// 	{
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(Slice(NewStrSliceV("1"))))
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(Slice(NewStrSliceV("4", "3"))))
// 		assert.False(t, NewStrSliceV("1", "2", "3").AnyS(Slice(NewStrSliceV("4", "5"))))
// 	}

// 	// StrSlice
// 	{
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(*NewStrSliceV("1")))
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(*NewStrSliceV("4", "3")))
// 		assert.False(t, NewStrSliceV("1", "2", "3").AnyS(*NewStrSliceV("4", "5")))
// 	}

// 	// *StrSlice
// 	{
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(NewStrSliceV("1")))
// 		assert.True(t, NewStrSliceV("1", "2", "3").AnyS(NewStrSliceV("4", "3")))
// 		assert.False(t, NewStrSliceV("1", "2", "3").AnyS(NewStrSliceV("4", "5")))
// 	}

// 	// invalid types
// 	assert.False(t, NewStrSliceV("1", "2").AnyS(nil))
// 	assert.False(t, NewStrSliceV("1", "2").AnyS((*[]string)(nil)))
// 	assert.False(t, NewStrSliceV("1", "2").AnyS((*StrSlice)(nil)))
// 	assert.False(t, NewStrSliceV("1", "2").AnyS([]int{2}))
// }

// // AnyW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_AnyW_Go(t *testing.B) {
// 	src := RangeStr(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkStrSlice_AnyW_Slice(t *testing.B) {
// 	src := RangeStr(nines5)
// 	NewStrSlice(src).AnyW(func(x O) bool {
// 		return ExB(x.(string) == string(nines4))
// 	})
// }

// func ExampleStrSlice_AnyW() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.AnyW(func(x O) bool {
// 		return ExB(x.(string) == "2")
// 	}))
// 	// Output: true
// }

// func TestStrSlice_AnyW(t *testing.T) {

// 	// empty
// 	var slice *StrSlice
// 	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(string) > "0") }))
// 	assert.False(t, NewStrSliceV().AnyW(func(x O) bool { return ExB(x.(string) > "0") }))

// 	// single
// 	assert.True(t, NewStrSliceV("2").AnyW(func(x O) bool { return ExB(x.(string) > "0") }))
// 	assert.True(t, NewStrSliceV("1", "2").AnyW(func(x O) bool { return ExB(x.(string) == "2") }))
// 	assert.True(t, NewStrSliceV("1", "2", "3").AnyW(func(x O) bool { return ExB(x.(string) == "4" || x.(string) == "3") }))
// }

// // Append
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Append_Go(t *testing.B) {
// 	src := []string{}
// 	for _, i := range RangeStr(nines6) {
// 		src = append(src, i)
// 	}
// }

// func BenchmarkStrSlice_Append_Slice(t *testing.B) {
// 	slice := NewStrSliceV()
// 	for _, i := range RangeStr(nines6) {
// 		slice.Append(i)
// 	}
// }

// func ExampleStrSlice_Append() {
// 	slice := NewStrSliceV("1").Append("2").Append("3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Append(t *testing.T) {

// 	// nil
// 	{
// 		var nilSlice *StrSlice
// 		assert.Equal(t, NewStrSliceV("0"), nilSlice.Append("0"))
// 		assert.Equal(t, (*StrSlice)(nil), nilSlice)
// 	}

// 	// Append one back to back
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, true, slice.Nil())
// 		slice = NewStrSliceV()
// 		assert.Equal(t, 0, slice.Len())
// 		assert.Equal(t, false, slice.Nil())

// 		// First append invokes 10x reflect overhead because the slice is nil
// 		slice.Append("1")
// 		assert.Equal(t, 1, slice.Len())
// 		assert.Equal(t, []string{"1"}, slice.O())

// 		// Second append another which will be 2x at most
// 		slice.Append("2")
// 		assert.Equal(t, 2, slice.Len())
// 		assert.Equal(t, []string{"1", "2"}, slice.O())
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 	}

// 	// Start with just appending without chaining
// 	{
// 		slice := NewStrSliceV()
// 		assert.Equal(t, 0, slice.Len())
// 		slice.Append("1")
// 		assert.Equal(t, []string{"1"}, slice.O())
// 		slice.Append("2")
// 		assert.Equal(t, []string{"1", "2"}, slice.O())
// 	}

// 	// Start with nil not chained
// 	{
// 		slice := NewStrSliceV()
// 		assert.Equal(t, 0, slice.Len())
// 		slice.Append("1").Append("2").Append("3")
// 		assert.Equal(t, 3, slice.Len())
// 		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
// 	}

// 	// Start with nil chained
// 	{
// 		slice := NewStrSliceV().Append("1").Append("2")
// 		assert.Equal(t, 2, slice.Len())
// 		assert.Equal(t, []string{"1", "2"}, slice.O())
// 	}

// 	// Start with non nil
// 	{
// 		slice := NewStrSliceV("1").Append("2").Append("3")
// 		assert.Equal(t, 3, slice.Len())
// 		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 	}

// 	// Use append result directly
// 	{
// 		slice := NewStrSliceV("1")
// 		assert.Equal(t, 1, slice.Len())
// 		assert.Equal(t, []string{"1", "2"}, slice.Append("2").O())
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 	}
// }

// // AppendV
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_AppendV_Go(t *testing.B) {
// 	src := []string{}
// 	src = append(src, RangeStr(nines6)...)
// }

// func BenchmarkStrSlice_AppendV_Slice(t *testing.B) {
// 	n := NewStrSliceV()
// 	new := rangeO(0, nines6)
// 	n.AppendV(new...)
// }

// func ExampleStrSlice_AppendV() {
// 	slice := NewStrSliceV("1").AppendV("2", "3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_AppendV(t *testing.T) {

// 	// nil
// 	{
// 		var nilSlice *StrSlice
// 		assert.Equal(t, NewStrSliceV("1", "2"), nilSlice.AppendV("1", "2"))
// 	}

// 	// Append many src
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1").AppendV("2", "3"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4", "5"), NewStrSliceV("1").AppendV("2", "3").AppendV("4", "5"))
// 	}
// }

// // At
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_At_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for _, x := range src {
// 		assert.IsType(t, 0, x)
// 	}
// }

// func BenchmarkStrSlice_At_Slice(t *testing.B) {
// 	src := RangeStr(nines6)
// 	slice := NewStrSlice(src)
// 	for i := 0; i < len(src); i++ {
// 		_, ok := (slice.At(i).O()).(string)
// 		assert.True(t, ok)
// 	}
// }

// func ExampleStrSlice_At() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.At(2))
// 	// Output: 3
// }

// func TestStrSlice_At(t *testing.T) {

// 	// nil
// 	{
// 		var nilSlice *StrSlice
// 		assert.Equal(t, Obj(nil), nilSlice.At(0))
// 	}

// 	// src
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4")
// 		assert.Equal(t, "4", slice.At(-1).O())
// 		assert.Equal(t, "3", slice.At(-2).O())
// 		assert.Equal(t, "2", slice.At(-3).O())
// 		assert.Equal(t, "1", slice.At(0).O())
// 		assert.Equal(t, "2", slice.At(1).O())
// 		assert.Equal(t, "3", slice.At(2).O())
// 		assert.Equal(t, "4", slice.At(3).O())
// 	}

// 	// index out of bounds
// 	{
// 		slice := NewStrSliceV("1")
// 		assert.Equal(t, &Object{}, slice.At(3))
// 		assert.Equal(t, nil, slice.At(3).O())
// 		assert.Equal(t, &Object{}, slice.At(-3))
// 		assert.Equal(t, nil, slice.At(-3).O())
// 	}
// }

// // Clear
// //--------------------------------------------------------------------------------------------------
// func ExampleStrSlice_Clear() {
// 	slice := NewStrSliceV("1").Concat([]string{"2", "3"})
// 	fmt.Println(slice.Clear())
// 	// Output: []
// }

// func TestStrSlice_Clear(t *testing.T) {

// 	// nil
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.Clear())
// 		assert.Equal(t, (*StrSlice)(nil), slice)
// 	}

// 	// int
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewStrSliceV(), slice.Clear())
// 		assert.Equal(t, NewStrSliceV(), slice.Clear())
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}
// }

// // Concat
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Concat_Go(t *testing.B) {
// 	dest := []string{}
// 	src := RangeStr(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkStrSlice_Concat_Slice(t *testing.B) {
// 	dest := NewStrSliceV()
// 	src := RangeStr(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.Concat(src[j:i])
// 		j = i
// 	}
// }

// func ExampleStrSlice_Concat() {
// 	slice := NewStrSliceV("1").Concat([]string{"2", "3"})
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Concat(t *testing.T) {

// 	// nil
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Concat([]string{"1", "2"}))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").Concat(nil))
// 	}

// 	// []string
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.Concat([]string{"2", "3"})
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), concated)
// 	}

// 	// *[]string
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.Concat(&([]string{"2", "3"}))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), concated)
// 	}

// 	// *StrSlice
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.Concat(NewStrSliceV("2", "3"))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), concated)
// 	}

// 	// StrSlice
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.Concat(*NewStrSliceV("2", "3"))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), concated)
// 	}

// 	// Slice
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.Concat(Slice(NewStrSliceV("2", "3")))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), concated)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").Concat((*[]string)(nil)))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").Concat((*StrSlice)(nil)))
// 	}
// }

// // ConcatM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_ConcatM_Go(t *testing.B) {
// 	dest := []string{}
// 	src := RangeStr(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkStrSlice_ConcatM_Slice(t *testing.B) {
// 	dest := NewStrSliceV()
// 	src := RangeStr(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.ConcatM(src[j:i])
// 		j = i
// 	}
// }

// func ExampleStrSlice_ConcatM() {
// 	slice := NewStrSliceV("1").ConcatM([]string{"2", "3"})
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_ConcatM(t *testing.T) {

// 	// nil
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.ConcatM([]string{"1", "2"}))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").ConcatM(nil))
// 	}

// 	// []string
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.ConcatM([]string{"2", "3"})
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// *[]string
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.ConcatM(&([]string{"2", "3"}))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// *StrSlice
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.ConcatM(NewStrSliceV("2", "3"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// StrSlice
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.ConcatM(*NewStrSliceV("2", "3"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// Slice
// 	{
// 		slice := NewStrSliceV("1")
// 		concated := slice.ConcatM(Slice(NewStrSliceV("2", "3")))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").ConcatM((*[]string)(nil)))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").ConcatM((*StrSlice)(nil)))
// 	}
// }

// // Copy
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Copy_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	dst := make([]string, len(src), len(src))
// 	copy(dst, src)
// }

// func BenchmarkStrSlice_Copy_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.Copy()
// }

// func ExampleStrSlice_Copy() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Copy())
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Copy(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.Copy(0, -1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV("0").Clear().Copy(0, -1))
// 	}

// 	// Test that the original is NOT modified when the slice is modified
// 	{
// 		original := NewStrSliceV("1", "2", "3")
// 		result := original.Copy(0, -1)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), original)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), result)
// 		result.Set(0, "0")
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), original)
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), result)
// 	}

// 	// copy full array
// 	{
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Copy())
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Copy(0, -1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Copy(0, 1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Copy(0, 5))
// 		assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1").Copy())
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Copy())
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Copy(0, -1))
// 		assert.Equal(t, NewStrSlice([]string{"1", "2", "3"}), NewStrSlice([]string{"1", "2", "3"}).Copy())
// 		assert.Equal(t, NewStrSlice([]string{"1", "2", "3"}), NewStrSlice([]string{"1", "2", "3"}).Copy(0, -1))
// 	}

// 	// out of bounds should be moved in
// 	{
// 		assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1").Copy(0, 2))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Copy(-6, 6))
// 	}

// 	// mutually exclusive
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewStrSliceV(), slice.Copy(2, -3))
// 		assert.Equal(t, NewStrSliceV(), slice.Copy(0, -5))
// 		assert.Equal(t, NewStrSliceV(), slice.Copy(4, -1))
// 		assert.Equal(t, NewStrSliceV(), slice.Copy(6, -1))
// 		assert.Equal(t, NewStrSliceV(), slice.Copy(3, -2))
// 	}

// 	// singles
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewStrSliceV("4"), slice.Copy(-1, -1))
// 		assert.Equal(t, NewStrSliceV("3"), slice.Copy(-2, -2))
// 		assert.Equal(t, NewStrSliceV("2"), slice.Copy(-3, -3))
// 		assert.Equal(t, NewStrSliceV("1"), slice.Copy(0, 0))
// 		assert.Equal(t, NewStrSliceV("1"), slice.Copy(-4, -4))
// 		assert.Equal(t, NewStrSliceV("2"), slice.Copy(1, 1))
// 		assert.Equal(t, NewStrSliceV("2"), slice.Copy(1, -3))
// 		assert.Equal(t, NewStrSliceV("3"), slice.Copy(2, 2))
// 		assert.Equal(t, NewStrSliceV("3"), slice.Copy(2, -2))
// 		assert.Equal(t, NewStrSliceV("4"), slice.Copy(3, 3))
// 		assert.Equal(t, NewStrSliceV("4"), slice.Copy(3, -1))
// 	}

// 	// grab all but first
// 	{
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Copy(1, -1))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Copy(1, 2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Copy(-2, -1))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Copy(-2, 2))
// 	}

// 	// grab all but last
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Copy(0, -2))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Copy(-3, -2))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Copy(-3, 1))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Copy(0, 1))
// 	}

// 	// grab middle
// 	{
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Copy(1, -2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Copy(-3, -2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Copy(-3, 2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Copy(1, 2))
// 	}

// 	// random
// 	{
// 		assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1", "2", "3").Copy(0, -3))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Copy(1, 2))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Copy(0, 2))
// 	}
// }

// // Count
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Count_Go(t *testing.B) {
// 	src := RangeStr(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkStrSlice_Count_Slice(t *testing.B) {
// 	src := RangeStr(nines5)
// 	NewStrSlice(src).Count(nines4)
// }

// func ExampleStrSlice_Count() {
// 	slice := NewStrSliceV("1", "2", "2")
// 	fmt.Println(slice.Count("2"))
// 	// Output: 2
// }

// func TestStrSlice_Count(t *testing.T) {

// 	// empty
// 	var slice *StrSlice
// 	assert.Equal(t, 0, slice.Count(0))
// 	assert.Equal(t, 0, NewStrSliceV().Count(0))

// 	assert.Equal(t, 1, NewStrSliceV("2", "3").Count("2"))
// 	assert.Equal(t, 2, NewStrSliceV("1", "2", "2").Count("2"))
// 	assert.Equal(t, 4, NewStrSliceV("4", "4", "3", "4", "4").Count("4"))
// 	assert.Equal(t, 3, NewStrSliceV("3", "2", "3", "3", "5").Count("3"))
// 	assert.Equal(t, 1, NewStrSliceV("1", "2", "3").Count("3"))
// }

// // CountW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_CountW_Go(t *testing.B) {
// 	src := RangeStr(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkStrSlice_CountW_Slice(t *testing.B) {
// 	src := RangeStr(nines5)
// 	NewStrSlice(src).CountW(func(x O) bool {
// 		return ExB(x.(string) == string(nines4))
// 	})
// }

// func ExampleStrSlice_CountW() {
// 	slice := NewStrSliceV("1", "2", "2")
// 	fmt.Println(slice.CountW(func(x O) bool {
// 		return ExB(x.(string) == "2")
// 	}))
// 	// Output: 2
// }

// func TestStrSlice_CountW(t *testing.T) {

// 	// empty
// 	var slice *StrSlice
// 	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(string) > "0") }))
// 	assert.Equal(t, 0, NewStrSliceV().CountW(func(x O) bool { return ExB(x.(string) > "0") }))

// 	assert.Equal(t, 1, NewStrSliceV("2", "3").CountW(func(x O) bool { return ExB(x.(string) > "2") }))
// 	assert.Equal(t, 1, NewStrSliceV("1", "2").CountW(func(x O) bool { return ExB(x.(string) == "2") }))
// 	assert.Equal(t, 1, NewStrSliceV("1", "2", "3").CountW(func(x O) bool { return ExB(x.(string) == "4" || x.(string) == "3") }))
// }

// // Drop
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Drop_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 11 {
// 		i := 1
// 		n := 10
// 		if i+n < len(src) {
// 			src = append(src[:i], src[i+n:]...)
// 		} else {
// 			src = src[:i]
// 		}
// 	}
// }

// func BenchmarkStrSlice_Drop_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 1 {
// 		slice.Drop(1, 10)
// 	}
// }

// func ExampleStrSlice_Drop() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Drop(0, 1))
// 	// Output: [3]
// }

// func TestStrSlice_Drop(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.Drop(0, 1))
// 	}

// 	// invalid
// 	assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(1))
// 	assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(4, 4))

// 	// drop 1
// 	assert.Equal(t, NewStrSliceV("2", "3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(0, 0))
// 	assert.Equal(t, NewStrSliceV("1", "3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(1, 1))
// 	assert.Equal(t, NewStrSliceV("1", "2", "4"), NewStrSliceV("1", "2", "3", "4").Drop(2, 2))
// 	assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3", "4").Drop(3, 3))
// 	assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3", "4").Drop(-1, -1))
// 	assert.Equal(t, NewStrSliceV("1", "2", "4"), NewStrSliceV("1", "2", "3", "4").Drop(-2, -2))
// 	assert.Equal(t, NewStrSliceV("1", "3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(-3, -3))
// 	assert.Equal(t, NewStrSliceV("2", "3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(-4, -4))

// 	// drop 2
// 	assert.Equal(t, NewStrSliceV("3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(0, 1))
// 	assert.Equal(t, NewStrSliceV("1", "4"), NewStrSliceV("1", "2", "3", "4").Drop(1, 2))
// 	assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3", "4").Drop(2, 3))
// 	assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3", "4").Drop(-2, -1))
// 	assert.Equal(t, NewStrSliceV("1", "4"), NewStrSliceV("1", "2", "3", "4").Drop(-3, -2))
// 	assert.Equal(t, NewStrSliceV("3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(-4, -3))

// 	// drop 3
// 	assert.Equal(t, NewStrSliceV("4"), NewStrSliceV("1", "2", "3", "4").Drop(0, 2))
// 	assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1", "2", "3", "4").Drop(-3, -1))

// 	// drop everything and beyond
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Drop())
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Drop(0, 3))
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Drop(0, -1))
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Drop(-4, -1))
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Drop(-6, -1))
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Drop(0, 10))

// 	// move index within bounds
// 	assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3", "4").Drop(3, 4))
// 	assert.Equal(t, NewStrSliceV("2", "3", "4"), NewStrSliceV("1", "2", "3", "4").Drop(-5, 0))
// }

// // DropAt
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_DropAt_Go(t *testing.B) {
// 	src := RangeStr(nines5)
// 	index := Range(0, nines5)
// 	for i := range index {
// 		if i+1 < len(src) {
// 			src = append(src[:i], src[i+1:]...)
// 		} else if i >= 0 && i < len(src) {
// 			src = src[:i]
// 		}
// 	}
// }

// func BenchmarkStrSlice_DropAt_Slice(t *testing.B) {
// 	index := Range(0, nines5)
// 	slice := NewStrSlice(RangeStr(nines5))
// 	for i := range index {
// 		slice.DropAt(i)
// 	}
// }

// func ExampleStrSlice_DropAt() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.DropAt(1))
// 	// Output: [1 3]
// }

// func TestStrSlice_DropAt(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.DropAt(0))
// 	}

// 	// drop all and more
// 	{
// 		slice := NewStrSliceV("0", "1", "2")
// 		assert.Equal(t, NewStrSliceV("0", "1"), slice.DropAt(-1))
// 		assert.Equal(t, NewStrSliceV("0"), slice.DropAt(-1))
// 		assert.Equal(t, NewStrSliceV(), slice.DropAt(-1))
// 		assert.Equal(t, NewStrSliceV(), slice.DropAt(-1))
// 	}

// 	// drop invalid
// 	assert.Equal(t, NewStrSliceV("0", "1", "2"), NewStrSliceV("0", "1", "2").DropAt(3))
// 	assert.Equal(t, NewStrSliceV("0", "1", "2"), NewStrSliceV("0", "1", "2").DropAt(-4))

// 	// drop last
// 	assert.Equal(t, NewStrSliceV("0", "1"), NewStrSliceV("0", "1", "2").DropAt(2))
// 	assert.Equal(t, NewStrSliceV("0", "1"), NewStrSliceV("0", "1", "2").DropAt(-1))

// 	// drop middle
// 	assert.Equal(t, NewStrSliceV("0", "2"), NewStrSliceV("0", "1", "2").DropAt(1))
// 	assert.Equal(t, NewStrSliceV("0", "2"), NewStrSliceV("0", "1", "2").DropAt(-2))

// 	// drop first
// 	assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("0", "1", "2").DropAt(0))
// 	assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("0", "1", "2").DropAt(-3))
// }

// // DropFirst
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_DropFirst_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkStrSlice_DropFirst_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirst()
// 	}
// }

// func ExampleStrSlice_DropFirst() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.DropFirst())
// 	// Output: [2 3]
// }

// func TestStrSlice_DropFirst(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.DropFirst())
// 	}

// 	// drop all and beyond
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("2", "3"), slice.DropFirst())
// 		assert.Equal(t, NewStrSliceV("3"), slice.DropFirst())
// 		assert.Equal(t, NewStrSliceV(), slice.DropFirst())
// 		assert.Equal(t, NewStrSliceV(), slice.DropFirst())
// 	}
// }

// // DropFirstN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_DropFirstN_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkStrSlice_DropFirstN_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirstN(10)
// 	}
// }

// func ExampleStrSlice_DropFirstN() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.DropFirstN(2))
// 	// Output: [3]
// }

// func TestStrSlice_DropFirstN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.DropFirstN(1))
// 	}

// 	// negative value
// 	assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").DropFirstN(-1))

// 	// drop none
// 	assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").DropFirstN(0))

// 	// drop 1
// 	assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").DropFirstN(1))

// 	// drop 2
// 	assert.Equal(t, NewStrSliceV("3"), NewStrSliceV("1", "2", "3").DropFirstN(2))

// 	// drop 3
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3").DropFirstN(3))

// 	// drop beyond
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3").DropFirstN(4))
// }

// // DropLast
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_DropLast_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkStrSlice_DropLast_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLast()
// 	}
// }

// func ExampleStrSlice_DropLast() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.DropLast())
// 	// Output: [1 2]
// }

// func TestStrSlice_DropLast(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.DropLast())
// 	}

// 	// negative value
// 	assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").DropLastN(-1))

// 	slice := NewStrSliceV("1", "2", "3")
// 	assert.Equal(t, NewStrSliceV("1", "2"), slice.DropLast())
// 	assert.Equal(t, NewStrSliceV("1"), slice.DropLast())
// 	assert.Equal(t, NewStrSliceV(), slice.DropLast())
// 	assert.Equal(t, NewStrSliceV(), slice.DropLast())
// }

// // DropLastN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_DropLastN_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkStrSlice_DropLastN_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLastN(10)
// 	}
// }

// func ExampleStrSlice_DropLastN() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.DropLastN(2))
// 	// Output: [1]
// }

// func TestStrSlice_DropLastN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.DropLastN(1))
// 	}

// 	// drop none
// 	assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").DropLastN(0))

// 	// drop 1
// 	assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").DropLastN(1))

// 	// drop 2
// 	assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1", "2", "3").DropLastN(2))

// 	// drop 3
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3").DropLastN(3))

// 	// drop beyond
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3").DropLastN(4))
// }

// // DropW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_DropW_Go(t *testing.B) {
// 	src := RangeStr(nines5)
// 	l := len(src)
// 	for i := 0; i < l; i++ {
// 		if Obj(src[i]).ToInt()%2 == 0 {
// 			if i+1 < l {
// 				src = append(src[:i], src[i+1:]...)
// 			} else if i >= 0 && i < l {
// 				src = src[:i]
// 			}
// 			l--
// 			i--
// 		}
// 	}
// }

// func BenchmarkStrSlice_DropW_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines5))
// 	slice.DropW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	})
// }

// func ExampleStrSlice_DropW() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.DropW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	}))
// 	// Output: [1 3]
// }

// func TestStrSlice_DropW(t *testing.T) {

// 	// drop all odd values
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		slice.DropW(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 != 0)
// 		})
// 		assert.Equal(t, NewStrSliceV("2", "4", "6", "8"), slice)
// 	}

// 	// drop all even values
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		slice.DropW(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 == 0)
// 		})
// 		assert.Equal(t, NewStrSliceV("1", "3", "5", "7", "9"), slice)
// 	}
// }

// // Each
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Each_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_Each_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).Each(func(x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleStrSlice_Each() {
// 	NewStrSliceV("1", "2", "3").Each(func(x O) {
// 		fmt.Printf("%v", x)
// 	})
// 	// Output: 123
// }

// func TestStrSlice_Each(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.Each(func(x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").Each(func(x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}
// }

// // EachE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_EachE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "0", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_EachE_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).EachE(func(x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleStrSlice_EachE() {
// 	NewStrSliceV("1", "2", "3").EachE(func(x O) error {
// 		fmt.Printf("%v", x)
// 		return nil
// 	})
// 	// Output: 123
// }

// func TestStrSlice_EachE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.EachE(func(x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachE(func(x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachE(func(x O) error {
// 			if x.(string) == "3" {
// 				return ErrBreak
// 			}
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"1", "2"}, results)
// 	}
// }

// // EachI
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_EachI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_EachI_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).EachI(func(i int, x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleStrSlice_EachI() {
// 	NewStrSliceV("1", "2", "3").EachI(func(i int, x O) {
// 		fmt.Printf("%v:%v", i, x)
// 	})
// 	// Output: 0:11:22:3
// }

// func TestStrSlice_EachI(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.EachI(func(i int, x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachI(func(i int, x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}
// }

// // EachIE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_EachIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_EachIE_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).EachIE(func(i int, x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleStrSlice_EachIE() {
// 	NewStrSliceV("1", "2", "3").EachIE(func(i int, x O) error {
// 		fmt.Printf("%v:%v", i, x)
// 		return nil
// 	})
// 	// Output: 0:11:22:3
// }

// func TestStrSlice_EachIE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.EachIE(func(i int, x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachIE(func(i int, x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachIE(func(i int, x O) error {
// 			if i == 2 {
// 				return ErrBreak
// 			}
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"1", "2"}, results)
// 	}
// }

// // EachR
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_EachR_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_EachR_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).EachR(func(x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleStrSlice_EachR() {
// 	NewStrSliceV("1", "2", "3").EachR(func(x O) {
// 		fmt.Printf("%v", x)
// 	})
// 	// Output: 321
// }

// func TestStrSlice_EachR(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.EachR(func(x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachR(func(x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}
// }

// // EachRE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_EachRE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_EachRE_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).EachRE(func(x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleStrSlice_EachRE() {
// 	NewStrSliceV("1", "2", "3").EachRE(func(x O) error {
// 		fmt.Printf("%v", x)
// 		return nil
// 	})
// 	// Output: 321
// }

// func TestStrSlice_EachRE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.EachRE(func(x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachRE(func(x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachRE(func(x O) error {
// 			if x.(string) == "1" {
// 				return ErrBreak
// 			}
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2"}, results)
// 	}
// }

// // EachRI
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_EachRI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_EachRI_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).EachRI(func(i int, x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleStrSlice_EachRI() {
// 	NewStrSliceV("1", "2", "3").EachRI(func(i int, x O) {
// 		fmt.Printf("%v:%v", i, x)
// 	})
// 	// Output: 2:31:20:1
// }

// func TestStrSlice_EachRI(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.EachRI(func(i int, x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachRI(func(i int, x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}
// }

// // EachRIE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_EachRIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeStr(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkStrSlice_EachRIE_Slice(t *testing.B) {
// 	NewStrSlice(RangeStr(nines6)).EachRIE(func(i int, x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleStrSlice_EachRIE() {
// 	NewStrSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
// 		fmt.Printf("%v:%v", i, x)
// 		return nil
// 	})
// 	// Output: 2:31:20:1
// }

// func TestStrSlice_EachRIE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		slice.EachRIE(func(i int, x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewStrSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
// 			if i == 0 {
// 				return ErrBreak
// 			}
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2"}, results)
// 	}
// }

// // Empty
// //--------------------------------------------------------------------------------------------------
// func ExampleStrSlice_Empty() {
// 	fmt.Println(NewStrSliceV().Empty())
// 	// Output: true
// }

// func TestStrSlice_Empty(t *testing.T) {

// 	// nil or empty
// 	{
// 		var nilSlice *StrSlice
// 		assert.Equal(t, true, nilSlice.Empty())
// 	}

// 	assert.Equal(t, true, NewStrSliceV().Empty())
// 	assert.Equal(t, false, NewStrSliceV("1").Empty())
// 	assert.Equal(t, false, NewStrSliceV("1", "2", "3").Empty())
// 	assert.Equal(t, false, NewStrSliceV("1").Empty())
// 	assert.Equal(t, false, NewStrSlice([]string{"1", "2", "3"}).Empty())
// }

// // First
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_First_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 1 {
// 		_ = src[0]
// 		src = src[1:]
// 	}
// }

// func BenchmarkStrSlice_First_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.First()
// 		slice.DropFirst()
// 	}
// }

// func ExampleStrSlice_First() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.First())
// 	// Output: 1
// }

// func TestStrSlice_First(t *testing.T) {
// 	// invalid
// 	assert.Equal(t, Obj(nil), NewStrSliceV().First())

// 	// int
// 	assert.Equal(t, Obj("2"), NewStrSliceV("2", "3").First())
// 	assert.Equal(t, Obj("3"), NewStrSliceV("3", "2").First())
// 	assert.Equal(t, Obj("1"), NewStrSliceV("1", "3", "2").First())
// }

// // FirstN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_FirstN_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	_ = src[0:10]
// }

// func BenchmarkStrSlice_FirstN_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	slice.FirstN(10)
// }

// func ExampleStrSlice_FirstN() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.FirstN(2))
// 	// Output: [1 2]
// }

// func TestStrSlice_FirstN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.FirstN(1))
// 		assert.Equal(t, NewStrSliceV(), slice.FirstN(-1))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewStrSliceV("1", "2", "3")
// 		result := original.FirstN(2).Set(0, "0")
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), original)
// 		assert.Equal(t, NewStrSliceV("0", "2"), result)
// 	}

// 	// Get none
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3").FirstN(0))

// 	// slice full array includeing out of bounds
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV().FirstN(1))
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV().FirstN(10))
// 	assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").FirstN(10))
// 	assert.Equal(t, NewStrSlice([]string{"1", "2", "3"}), NewStrSlice([]string{"1", "2", "3"}).FirstN(10))

// 	// grab a few diff
// 	assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1", "2", "3").FirstN(1))
// 	assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").FirstN(2))
// }

// // Generic
// //--------------------------------------------------------------------------------------------------
// func ExampleStrSlice_Generic() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Generic())
// 	// Output: false
// }

// // Index
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Index_Go(t *testing.B) {
// 	for _, x := range RangeStr(nines5) {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkStrSlice_Index_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines5))
// 	slice.Index(nines4)
// }

// func ExampleStrSlice_Index() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Index("2"))
// 	// Output: 1
// }

// func TestStrSlice_Index(t *testing.T) {

// 	// empty
// 	var slice *StrSlice
// 	assert.Equal(t, -1, slice.Index("2"))
// 	assert.Equal(t, -1, NewStrSliceV().Index("1"))

// 	assert.Equal(t, 0, NewStrSliceV("1", "2", "3").Index("1"))
// 	assert.Equal(t, 1, NewStrSliceV("1", "2", "3").Index("2"))
// 	assert.Equal(t, 2, NewStrSliceV("1", "2", "3").Index("3"))
// 	assert.Equal(t, -1, NewStrSliceV("1", "2", "3").Index("4"))
// 	assert.Equal(t, -1, NewStrSliceV("1", "2", "3").Index("5"))
// }

// // Insert
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Insert_Go(t *testing.B) {
// 	src := []string{}
// 	for _, x := range RangeStr(nines6) {
// 		src = append(src, x)
// 		copy(src[1:], src[1:])
// 		src[0] = x
// 	}
// }

// func BenchmarkStrSlice_Insert_Slice(t *testing.B) {
// 	slice := NewStrSliceV()
// 	for x := range RangeStr(nines6) {
// 		slice.Insert(0, x)
// 	}
// }

// func ExampleStrSlice_Insert() {
// 	slice := NewStrSliceV("1", "3")
// 	fmt.Println(slice.Insert(1, "2"))
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Insert(t *testing.T) {

// 	// append
// 	{
// 		slice := NewStrSliceV()
// 		assert.Equal(t, NewStrSliceV("0"), slice.Insert(-1, "0"))
// 		assert.Equal(t, NewStrSliceV("0", "1"), slice.Insert(-1, "1"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2"), slice.Insert(-1, "2"))
// 	}

// 	// prepend
// 	{
// 		slice := NewStrSliceV()
// 		assert.Equal(t, NewStrSliceV("2"), slice.Insert(0, "2"))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Insert(0, "1"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2"), slice.Insert(0, "0"))
// 	}

// 	// middle pos
// 	{
// 		slice := NewStrSliceV("0", "5")
// 		assert.Equal(t, NewStrSliceV("0", "1", "5"), slice.Insert(1, "1"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2", "5"), slice.Insert(2, "2"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2", "3", "5"), slice.Insert(3, "3"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(4, "4"))
// 	}

// 	// middle neg
// 	{
// 		slice := NewStrSliceV("0", "5")
// 		assert.Equal(t, NewStrSliceV("0", "1", "5"), slice.Insert(-2, "1"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2", "5"), slice.Insert(-2, "2"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2", "3", "5"), slice.Insert(-2, "3"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(-2, "4"))
// 	}

// 	// error cases
// 	{
// 		var slice *StrSlice
// 		assert.False(t, slice.Insert(0, 0).Nil())
// 		assert.Equal(t, NewStrSliceV("0"), slice.Insert(0, "0"))
// 		assert.Equal(t, NewStrSliceV("0", "1"), NewStrSliceV("0", "1").Insert(-10, "1"))
// 		assert.Equal(t, NewStrSliceV("0", "1"), NewStrSliceV("0", "1").Insert(10, "1"))
// 		assert.Equal(t, NewStrSliceV("0", "1"), NewStrSliceV("0", "1").Insert(2, "1"))
// 		assert.Equal(t, NewStrSliceV("0", "1"), NewStrSliceV("0", "1").Insert(-3, "1"))
// 	}
// }

// // Join
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Join_Go(t *testing.B) {
// 	src := RangeStr(nines4)
// 	strings.Join(src, ",")
// }

// func BenchmarkStrSlice_Join_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines4))
// 	slice.Join()
// }

// func ExampleStrSlice_Join() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Join())
// 	// Output: 1,2,3
// }

// func TestStrSlice_Join(t *testing.T) {
// 	// nil
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, Obj(""), slice.Join())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, Obj(""), NewStrSliceV().Join())
// 	}

// 	assert.Equal(t, "1,2,3", NewStrSliceV("1", "2", "3").Join().O())
// 	assert.Equal(t, "1.2.3", NewStrSliceV("1", "2", "3").Join(".").O())
// }

// // Last
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Last_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 1 {
// 		_ = src[len(src)-1]
// 		src = src[:len(src)-1]
// 	}
// }

// func BenchmarkStrSlice_Last_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.Last()
// 		slice.DropLast()
// 	}
// }

// func ExampleStrSlice_Last() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Last())
// 	// Output: 3
// }

// func TestStrSlice_Last(t *testing.T) {
// 	// invalid
// 	assert.Equal(t, Obj(nil), NewStrSliceV().Last())

// 	// int
// 	assert.Equal(t, Obj("3"), NewStrSliceV("2", "3").Last())
// 	assert.Equal(t, Obj("2"), NewStrSliceV("3", "2").Last())
// 	assert.Equal(t, Obj("2"), NewStrSliceV("1", "3", "2").Last())
// }

// // LastN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_LastN_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	_ = src[0:10]
// }

// func BenchmarkStrSlice_LastN_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	slice.LastN(10)
// }

// func ExampleStrSlice_LastN() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.LastN(2))
// 	// Output: [2 3]
// }

// func TestStrSlice_LastN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.LastN(1))
// 		assert.Equal(t, NewStrSliceV(), slice.LastN(-1))
// 	}

// 	// Get none
// 	{
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3").LastN(0))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewStrSliceV("1", "2", "3")
// 		result := original.LastN(2).Set(0, "0")
// 		assert.Equal(t, NewStrSliceV("1", "0", "3"), original)
// 		assert.Equal(t, NewStrSliceV("0", "3"), result)
// 	}

// 	// slice full array includeing out of bounds
// 	{
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().LastN(1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().LastN(10))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").LastN(10))
// 		assert.Equal(t, NewStrSlice([]string{"1", "2", "3"}), NewStrSlice([]string{"1", "2", "3"}).LastN(10))
// 	}

// 	// grab a few diff
// 	{
// 		assert.Equal(t, NewStrSliceV("3"), NewStrSliceV("1", "2", "3").LastN(1))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").LastN(2))
// 	}
// }

// // Len
// //--------------------------------------------------------------------------------------------------
// func ExampleStrSlice_Len() {
// 	fmt.Println(NewStrSliceV("1", "2", "3").Len())
// 	// Output: 3
// }

// func TestStrSlice_Len(t *testing.T) {
// 	assert.Equal(t, 0, NewStrSliceV().Len())
// 	assert.Equal(t, 2, len(*(NewStrSliceV("1", "2"))))
// 	assert.Equal(t, 2, NewStrSliceV("1", "2").Len())
// }

// // Less
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Less_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if i+1 < len(src) {
// 			_ = src[i] < src[i+1]
// 		}
// 	}
// }

// func BenchmarkStrSlice_Less_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Less(i, i+1)
// 		}
// 	}
// }

// func ExampleStrSlice_Less() {
// 	slice := NewStrSliceV("2", "3", "1")
// 	fmt.Println(slice.Less(0, 2))
// 	// Output: false
// }

// func TestStrSlice_Less(t *testing.T) {

// 	// invalid cases
// 	{
// 		var slice *StrSlice
// 		assert.False(t, slice.Less(0, 0))

// 		slice = NewStrSliceV()
// 		assert.False(t, slice.Less(0, 0))
// 		assert.False(t, slice.Less(1, 2))
// 		assert.False(t, slice.Less(-1, 2))
// 		assert.False(t, slice.Less(1, -2))
// 	}

// 	// valid
// 	assert.Equal(t, true, NewStrSliceV("0", "1", "2").Less(0, 1))
// 	assert.Equal(t, false, NewStrSliceV("0", "1", "2").Less(1, 0))
// 	assert.Equal(t, true, NewStrSliceV("0", "1", "2").Less(1, 2))
// }

// // Nil
// //--------------------------------------------------------------------------------------------------
// func ExampleStrSlice_Nil() {
// 	var slice *StrSlice
// 	fmt.Println(slice.Nil())
// 	// Output: true
// }

// func TestStrSlice_Nil(t *testing.T) {
// 	var slice *StrSlice
// 	assert.True(t, slice.Nil())
// 	assert.False(t, NewStrSliceV().Nil())
// 	assert.False(t, NewStrSliceV("1", "2", "3").Nil())
// }

// // O
// //--------------------------------------------------------------------------------------------------
// func ExampleStrSlice_O() {
// 	fmt.Println(NewStrSliceV("1", "2", "3"))
// 	// Output: [1 2 3]
// }

// func TestStrSlice_O(t *testing.T) {
// 	assert.Equal(t, []string{}, NewStrSliceV().O())
// 	assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3"))
// }

// // Pair
// //--------------------------------------------------------------------------------------------------

// func ExampleStrSlice_Pair() {
// 	slice := NewStrSliceV("1", "2")
// 	first, second := slice.Pair()
// 	fmt.Println(first, second)
// 	// Output: 1 2
// }

// func TestStrSlice_Pair(t *testing.T) {

// 	// nil
// 	{
// 		first, second := (*StrSlice)(nil).Pair()
// 		assert.Equal(t, Obj(nil), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}

// 	// two values
// 	{
// 		first, second := NewStrSliceV("1", "2").Pair()
// 		assert.Equal(t, Obj("1"), first)
// 		assert.Equal(t, Obj("2"), second)
// 	}

// 	// one value
// 	{
// 		first, second := NewStrSliceV("1").Pair()
// 		assert.Equal(t, Obj("1"), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}

// 	// no values
// 	{
// 		first, second := NewStrSliceV().Pair()
// 		assert.Equal(t, Obj(nil), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}
// }

// // Pop
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Pop_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkStrSlice_Pop_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.Pop()
// 	}
// }

// func ExampleStrSlice_Pop() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Pop())
// 	// Output: 3
// }

// func TestStrSlice_Pop(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, Obj(nil), slice.Pop())
// 	}

// 	// take all one at a time
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, Obj("3"), slice.Pop())
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 		assert.Equal(t, Obj("2"), slice.Pop())
// 		assert.Equal(t, NewStrSliceV("1"), slice)
// 		assert.Equal(t, Obj("1"), slice.Pop())
// 		assert.Equal(t, NewStrSliceV(), slice)
// 		assert.Equal(t, Obj(nil), slice.Pop())
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}
// }

// // PopN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_PopN_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkStrSlice_PopN_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.PopN(10)
// 	}
// }

// func ExampleStrSlice_PopN() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.PopN(2))
// 	// Output: [2 3]
// }

// func TestStrSlice_PopN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.PopN(1))
// 	}

// 	// take none
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV(), slice.PopN(0))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 	}

// 	// take 1
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("3"), slice.PopN(1))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 	}

// 	// take 2
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("2", "3"), slice.PopN(2))
// 		assert.Equal(t, NewStrSliceV("1"), slice)
// 	}

// 	// take 3
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice.PopN(3))
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}

// 	// take beyond
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice.PopN(4))
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}
// }

// // Prepend
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Prepend_Go(t *testing.B) {
// 	src := []string{}
// 	for _, x := range RangeStr(nines6) {
// 		src = append(src, x)
// 		copy(src[1:], src[1:])
// 		src[0] = x
// 	}
// }

// func BenchmarkStrSlice_Prepend_Slice(t *testing.B) {
// 	slice := NewStrSliceV()
// 	for _, x := range RangeStr(nines6) {
// 		slice.Prepend(x)
// 	}
// }

// func ExampleStrSlice_Prepend() {
// 	slice := NewStrSliceV("2", "3")
// 	fmt.Println(slice.Prepend("1"))
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Prepend(t *testing.T) {

// 	// happy path
// 	{
// 		slice := NewStrSliceV()
// 		assert.Equal(t, NewStrSliceV("2"), slice.Prepend("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Prepend("1"))
// 		assert.Equal(t, NewStrSliceV("0", "1", "2"), slice.Prepend("0"))
// 	}

// 	// error cases
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV("0"), slice.Prepend("0"))
// 	}
// }

// // Reverse
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Reverse_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
// 		src[i], src[j] = src[j], src[i]
// 	}
// }

// func BenchmarkStrSlice_Reverse_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.Reverse()
// }

// func ExampleStrSlice_Reverse() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Reverse())
// 	// Output: [3 2 1]
// }

// func TestStrSlice_Reverse(t *testing.T) {

// 	// nil
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.Reverse())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Reverse())
// 	}

// 	// pos
// 	{
// 		slice := NewStrSliceV("3", "2", "1")
// 		reversed := slice.Reverse()
// 		assert.Equal(t, NewStrSliceV("3", "2", "1", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), reversed)
// 	}

// 	// neg
// 	{
// 		slice := NewStrSliceV("2", "3", "-2", "-3")
// 		reversed := slice.Reverse()
// 		assert.Equal(t, NewStrSliceV("2", "3", "-2", "-3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("-3", "-2", "3", "2"), reversed)
// 	}
// }

// // ReverseM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_ReverseM_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
// 		src[i], src[j] = src[j], src[i]
// 	}
// }

// func BenchmarkStrSlice_ReverseM_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.ReverseM()
// }

// func ExampleStrSlice_ReverseM() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.ReverseM())
// 	// Output: [3 2 1]
// }

// func TestStrSlice_ReverseM(t *testing.T) {

// 	// nil
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.ReverseM())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().ReverseM())
// 	}

// 	// pos
// 	{
// 		slice := NewStrSliceV("3", "2", "1")
// 		reversed := slice.ReverseM()
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), reversed)
// 	}

// 	// neg
// 	{
// 		slice := NewStrSliceV("2", "3", "-2", "-3")
// 		reversed := slice.ReverseM()
// 		assert.Equal(t, NewStrSliceV("-3", "-2", "3", "2", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("-3", "-2", "3", "2", "4"), reversed)
// 	}
// }

// // Select
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Select_Go(t *testing.B) {
// 	even := []string{}
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if Obj(src[i]).ToInt()%2 == 0 {
// 			even = append(even, src[i])
// 		}
// 	}
// }

// func BenchmarkStrSlice_Select_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.Select(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	})
// }

// func ExampleStrSlice_Select() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Select(func(x O) bool {
// 		return ExB(x.(string) == "2" || x.(string) == "3")
// 	}))
// 	// Output: [2 3]
// }

// func TestStrSlice_Select(t *testing.T) {

// 	// Select all odd values
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.Select(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 != 0)
// 		})
// 		slice.DropFirst()
// 		assert.Equal(t, NewStrSliceV("2", "3", "4", "5", "6", "7", "8", "9"), slice)
// 		assert.Equal(t, NewStrSliceV("1", "3", "5", "7", "9"), new)
// 	}

// 	// Select all even values
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.Select(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 == 0)
// 		})
// 		slice.DropAt(1)
// 		assert.Equal(t, NewStrSliceV("1", "3", "4", "5", "6", "7", "8", "9"), slice)
// 		assert.Equal(t, NewStrSliceV("2", "4", "6", "8"), new)
// 	}
// }

// // Set
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Set_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i++ {
// 		src[i] = "0"
// 	}
// }

// func BenchmarkStrSlice_Set_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.Set(i, "0")
// 	}
// }

// func ExampleStrSlice_Set() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Set(0, "0"))
// 	// Output: [0 2 3]
// }

// func TestStrSlice_Set(t *testing.T) {
// 	assert.Equal(t, NewStrSliceV("0", "2", "3"), NewStrSliceV("1", "2", "3").Set(0, "0"))
// 	assert.Equal(t, NewStrSliceV("1", "0", "3"), NewStrSliceV("1", "2", "3").Set(1, "0"))
// 	assert.Equal(t, NewStrSliceV("1", "2", "0"), NewStrSliceV("1", "2", "3").Set(2, "0"))
// 	assert.Equal(t, NewStrSliceV("0", "2", "3"), NewStrSliceV("1", "2", "3").Set(-3, "0"))
// 	assert.Equal(t, NewStrSliceV("1", "0", "3"), NewStrSliceV("1", "2", "3").Set(-2, "0"))
// 	assert.Equal(t, NewStrSliceV("1", "2", "0"), NewStrSliceV("1", "2", "3").Set(-1, "0"))

// 	// Test out of bounds
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Set(5, "1"))
// 	}

// 	// Test wrong type
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Set(5, "1"))
// 	}
// }

// // SetE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_SetE_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i++ {
// 		src[i] = "0"
// 	}
// }

// func BenchmarkStrSlice_SetE_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.SetE(i, "0")
// 	}
// }

// func ExampleStrSlice_SetE() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.SetE(0, "0"))
// 	// Output: [0 2 3] <nil>
// }

// func TestStrSlice_SetE(t *testing.T) {
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		result, err := slice.SetE(0, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), slice)
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), result)
// 	}
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		result, err := slice.SetE(1, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewStrSliceV("1", "0", "3"), slice)
// 		assert.Equal(t, NewStrSliceV("1", "0", "3"), result)
// 	}
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		result, err := slice.SetE(2, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewStrSliceV("1", "2", "0"), slice)
// 		assert.Equal(t, NewStrSliceV("1", "2", "0"), result)
// 	}
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		result, err := slice.SetE(-3, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), slice)
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), result)
// 	}
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		result, err := slice.SetE(-2, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewStrSliceV("1", "0", "3"), slice)
// 		assert.Equal(t, NewStrSliceV("1", "0", "3"), result)
// 	}
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		result, err := slice.SetE(-1, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewStrSliceV("1", "2", "0"), slice)
// 		assert.Equal(t, NewStrSliceV("1", "2", "0"), result)
// 	}

// 	// Test out of bounds
// 	{
// 		slice, err := NewStrSliceV("1", "2", "3").SetE(5, "1")
// 		assert.NotNil(t, slice)
// 		assert.NotNil(t, err)
// 	}

// 	// Test wrong type
// 	{
// 		slice, err := NewStrSliceV("1", "2", "3").SetE(5, "1")
// 		assert.NotNil(t, slice)
// 		assert.NotNil(t, err)
// 	}
// }

// // Shift
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Shift_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkStrSlice_Shift_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.Shift()
// 	}
// }

// func ExampleStrSlice_Shift() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Shift())
// 	// Output: 1
// }

// func TestStrSlice_Shift(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, Obj(nil), slice.Shift())
// 	}

// 	// take all and beyond
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, Obj("1"), slice.Shift())
// 		assert.Equal(t, NewStrSliceV("2", "3"), slice)
// 		assert.Equal(t, Obj("2"), slice.Shift())
// 		assert.Equal(t, NewStrSliceV("3"), slice)
// 		assert.Equal(t, Obj("3"), slice.Shift())
// 		assert.Equal(t, NewStrSliceV(), slice)
// 		assert.Equal(t, Obj(nil), slice.Shift())
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}
// }

// // ShiftN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_ShiftN_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkStrSlice_ShiftN_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 0 {
// 		slice.ShiftN(10)
// 	}
// }

// func ExampleStrSlice_ShiftN() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.ShiftN(2))
// 	// Output: [1 2]
// }

// func TestStrSlice_ShiftN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.ShiftN(1))
// 	}

// 	// negative value
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("1"), slice.ShiftN(-1))
// 		assert.Equal(t, NewStrSliceV("2", "3"), slice)
// 	}

// 	// take none
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV(), slice.ShiftN(0))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 	}

// 	// take 1
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("1"), slice.ShiftN(1))
// 		assert.Equal(t, NewStrSliceV("2", "3"), slice)
// 	}

// 	// take 2
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.ShiftN(2))
// 		assert.Equal(t, NewStrSliceV("3"), slice)
// 	}

// 	// take 3
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice.ShiftN(3))
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}

// 	// take beyond
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice.ShiftN(4))
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}
// }

// // Single
// //--------------------------------------------------------------------------------------------------

// func ExampleStrSlice_Single() {
// 	slice := NewStrSliceV("1")
// 	fmt.Println(slice.Single())
// 	// Output: true
// }

// func TestStrSlice_Single(t *testing.T) {

// 	assert.Equal(t, false, NewStrSliceV().Single())
// 	assert.Equal(t, true, NewStrSliceV("1").Single())
// 	assert.Equal(t, false, NewStrSliceV("1", "2").Single())
// }

// // Slice
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Slice_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	_ = src[0:len(src)]
// }

// func BenchmarkStrSlice_Slice_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	slice.Slice(0, -1)
// }

// func ExampleStrSlice_Slice() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Slice(1, -1))
// 	// Output: [2 3]
// }

// func TestStrSlice_Slice(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.Slice(0, -1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Slice(0, -1))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewStrSliceV("1", "2", "3")
// 		result := original.Slice(0, -1).Set(0, "0")
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), original)
// 		assert.Equal(t, NewStrSliceV("0", "2", "3"), result)
// 	}

// 	// slice full array
// 	{
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Slice(0, -1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Slice(0, 1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV().Slice(0, 5))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Slice(0, -1))
// 		assert.Equal(t, NewStrSlice([]string{"1", "2", "3"}), NewStrSlice([]string{"1", "2", "3"}).Slice(0, -1))
// 	}

// 	// out of bounds should be moved in
// 	{
// 		assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1").Slice(0, 2))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Slice(-6, 6))
// 	}

// 	// mutually exclusive
// 	{
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Slice(2, -3))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Slice(0, -5))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Slice(4, -1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Slice(6, -1))
// 		assert.Equal(t, NewStrSliceV(), NewStrSliceV("1", "2", "3", "4").Slice(3, 2))
// 	}

// 	// singles
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewStrSliceV("4"), slice.Slice(-1, -1))
// 		assert.Equal(t, NewStrSliceV("3"), slice.Slice(-2, -2))
// 		assert.Equal(t, NewStrSliceV("2"), slice.Slice(-3, -3))
// 		assert.Equal(t, NewStrSliceV("1"), slice.Slice(0, 0))
// 		assert.Equal(t, NewStrSliceV("1"), slice.Slice(-4, -4))
// 		assert.Equal(t, NewStrSliceV("2"), slice.Slice(1, 1))
// 		assert.Equal(t, NewStrSliceV("2"), slice.Slice(1, -3))
// 		assert.Equal(t, NewStrSliceV("3"), slice.Slice(2, 2))
// 		assert.Equal(t, NewStrSliceV("3"), slice.Slice(2, -2))
// 		assert.Equal(t, NewStrSliceV("4"), slice.Slice(3, 3))
// 		assert.Equal(t, NewStrSliceV("4"), slice.Slice(3, -1))
// 	}

// 	// grab all but first
// 	{
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Slice(1, -1))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Slice(1, 2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Slice(-2, -1))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Slice(-2, 2))
// 	}

// 	// grab all but last
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Slice(0, -2))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Slice(-3, -2))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Slice(-3, 1))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2", "3").Slice(0, 1))
// 	}

// 	// grab middle
// 	{
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Slice(1, -2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Slice(-3, -2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Slice(-3, 2))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3", "4").Slice(1, 2))
// 	}

// 	// random
// 	{
// 		assert.Equal(t, NewStrSliceV("1"), NewStrSliceV("1", "2", "3").Slice(0, -3))
// 		assert.Equal(t, NewStrSliceV("2", "3"), NewStrSliceV("1", "2", "3").Slice(1, 2))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), NewStrSliceV("1", "2", "3").Slice(0, 2))
// 	}
// }

// // Sort
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Sort_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	sort.Sort(sort.StrSlice(src))
// }

// func BenchmarkStrSlice_Sort_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.Sort()
// }

// func ExampleStrSlice_Sort() {
// 	slice := NewStrSliceV("2", "3", "1")
// 	fmt.Println(slice.Sort())
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Sort(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV().Sort())

// 	// pos
// 	{
// 		slice := NewStrSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.Sort()
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("5", "3", "2", "4", "1"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewStrSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.Sort()
// 		assert.Equal(t, NewStrSliceV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("5", "3", "-2", "4", "-1"), slice)
// 	}
// }

// // SortM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_SortM_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	sort.Sort(sort.StrSlice(src))
// }

// func BenchmarkStrSlice_SortM_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.SortM()
// }

// func ExampleStrSlice_SortM() {
// 	slice := NewStrSliceV("2", "3", "1")
// 	fmt.Println(slice.SortM())
// 	// Output: [1 2 3]
// }

// func TestStrSlice_SortM(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV().SortM())

// 	// pos
// 	{
// 		slice := NewStrSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortM()
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4", "5", "6"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewStrSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortM()
// 		assert.Equal(t, NewStrSliceV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("-1", "-2", "3", "4", "5", "6"), slice)
// 	}
// }

// // SortReverse
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_SortReverse_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	sort.Sort(sort.Reverse(sort.StrSlice(src)))
// }

// func BenchmarkStrSlice_SortReverse_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.SortReverse()
// }

// func ExampleStrSlice_SortReverse() {
// 	slice := NewStrSliceV("2", "3", "1")
// 	fmt.Println(slice.SortReverse())
// 	// Output: [3 2 1]
// }

// func TestStrSlice_SortReverse(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV().SortReverse())

// 	// pos
// 	{
// 		slice := NewStrSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverse()
// 		assert.Equal(t, NewStrSliceV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("5", "3", "2", "4", "1"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewStrSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverse()
// 		assert.Equal(t, NewStrSliceV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("5", "3", "-2", "4", "-1"), slice)
// 	}
// }

// // SortReverseM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_SortReverseM_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	sort.Sort(sort.Reverse(sort.StrSlice(src)))
// }

// func BenchmarkStrSlice_SortReverseM_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	slice.SortReverseM()
// }

// func ExampleStrSlice_SortReverseM() {
// 	slice := NewStrSliceV("2", "3", "1")
// 	fmt.Println(slice.SortReverseM())
// 	// Output: [3 2 1]
// }

// func TestStrSlice_SortReverseM(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewStrSliceV(), NewStrSliceV().SortReverse())

// 	// pos
// 	{
// 		slice := NewStrSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, NewStrSliceV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("5", "4", "3", "2", "1", "6"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewStrSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, NewStrSliceV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewStrSliceV("5", "4", "3", "-2", "-1", "6"), slice)
// 	}
// }

// // String
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_String_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	_ = fmt.Sprintf("%v", src)
// }

// func BenchmarkStrSlice_String_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	_ = slice.String()
// }

// func ExampleStrSlice_String() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_String(t *testing.T) {
// 	// nil
// 	assert.Equal(t, "[]", (*StrSlice)(nil).String())

// 	// empty
// 	assert.Equal(t, "[]", NewStrSliceV().String())

// 	// pos
// 	{
// 		slice := NewStrSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, "[5 4 3 2 1 6]", sorted.Append("6").String())
// 		assert.Equal(t, "[5 4 3 2 1 6]", slice.String())
// 	}

// 	// neg
// 	{
// 		slice := NewStrSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, "[5 4 3 -2 -1 6]", sorted.Append("6").String())
// 		assert.Equal(t, "[5 4 3 -2 -1 6]", slice.String())
// 	}
// }

// // Swap
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Swap_Go(t *testing.B) {
// 	src := RangeStr(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if i+1 < len(src) {
// 			src[i], src[i+1] = src[i+1], src[i]
// 		}
// 	}
// }

// func BenchmarkStrSlice_Swap_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Swap(i, i+1)
// 		}
// 	}
// }

// func ExampleStrSlice_Swap() {
// 	slice := NewStrSliceV("2", "3", "1")
// 	slice.Swap(0, 2)
// 	slice.Swap(1, 2)
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Swap(t *testing.T) {

// 	// invalid cases
// 	{
// 		var slice *StrSlice
// 		slice.Swap(0, 0)
// 		assert.Equal(t, (*StrSlice)(nil), slice)

// 		slice = NewStrSliceV()
// 		slice.Swap(0, 0)
// 		assert.Equal(t, NewStrSliceV(), slice)

// 		slice.Swap(1, 2)
// 		assert.Equal(t, NewStrSliceV(), slice)

// 		slice.Swap(-1, 2)
// 		assert.Equal(t, NewStrSliceV(), slice)

// 		slice.Swap(1, -2)
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}

// 	// normal
// 	{
// 		slice := NewStrSliceV("0", "1", "2")
// 		slice.Swap(0, 1)
// 		assert.Equal(t, NewStrSliceV("1", "0", "2"), slice)
// 	}
// }

// // Take
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Take_Go(t *testing.B) {
// 	src := RangeStr(nines7)
// 	for len(src) > 11 {
// 		i := 1
// 		n := 10
// 		if i+n < len(src) {
// 			src = append(src[:i], src[i+n:]...)
// 		} else {
// 			src = src[:i]
// 		}
// 	}
// }

// func BenchmarkStrSlice_Take_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines7))
// 	for slice.Len() > 1 {
// 		slice.Take(1, 10)
// 	}
// }

// func ExampleStrSlice_Take() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.Take(0, 1))
// 	// Output: [1 2]
// }

// func TestStrSlice_Take(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.Take(0, 1))
// 	}

// 	// invalid
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewStrSliceV(), slice.Take(1))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice)
// 		assert.Equal(t, NewStrSliceV(), slice.Take(4, 4))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice)
// 	}

// 	// take 1
// 	{
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1"), slice.Take(0, 0))
// 			assert.Equal(t, NewStrSliceV("2", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("2"), slice.Take(1, 1))
// 			assert.Equal(t, NewStrSliceV("1", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("3"), slice.Take(2, 2))
// 			assert.Equal(t, NewStrSliceV("1", "2", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("4"), slice.Take(3, 3))
// 			assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("4"), slice.Take(-1, -1))
// 			assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("3"), slice.Take(-2, -2))
// 			assert.Equal(t, NewStrSliceV("1", "2", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("2"), slice.Take(-3, -3))
// 			assert.Equal(t, NewStrSliceV("1", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1"), slice.Take(-4, -4))
// 			assert.Equal(t, NewStrSliceV("2", "3", "4"), slice)
// 		}
// 	}

// 	// take 2
// 	{
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2"), slice.Take(0, 1))
// 			assert.Equal(t, NewStrSliceV("3", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("2", "3"), slice.Take(1, 2))
// 			assert.Equal(t, NewStrSliceV("1", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("3", "4"), slice.Take(2, 3))
// 			assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("3", "4"), slice.Take(-2, -1))
// 			assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("2", "3"), slice.Take(-3, -2))
// 			assert.Equal(t, NewStrSliceV("1", "4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2"), slice.Take(-4, -3))
// 			assert.Equal(t, NewStrSliceV("3", "4"), slice)
// 		}
// 	}

// 	// take 3
// 	{
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2", "3"), slice.Take(0, 2))
// 			assert.Equal(t, NewStrSliceV("4"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("2", "3", "4"), slice.Take(-3, -1))
// 			assert.Equal(t, NewStrSliceV("1"), slice)
// 		}
// 	}

// 	// take everything and beyond
// 	{
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Take())
// 			assert.Equal(t, NewStrSliceV(), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Take(0, 3))
// 			assert.Equal(t, NewStrSliceV(), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Take(0, -1))
// 			assert.Equal(t, NewStrSliceV(), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Take(-4, -1))
// 			assert.Equal(t, NewStrSliceV(), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Take(-6, -1))
// 			assert.Equal(t, NewStrSliceV(), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Take(0, 10))
// 			assert.Equal(t, NewStrSliceV(), slice)
// 		}
// 	}

// 	// move index within bounds
// 	{
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("4"), slice.Take(3, 4))
// 			assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewStrSliceV("1"), slice.Take(-5, 0))
// 			assert.Equal(t, NewStrSliceV("2", "3", "4"), slice)
// 		}
// 	}
// }

// // TakeAt
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_TakeAt_Go(t *testing.B) {
// 	src := RangeStr(nines5)
// 	index := RangeStr(nines5)
// 	for i := range index {
// 		if i+1 < len(src) {
// 			src = append(src[:i], src[i+1:]...)
// 		} else if i >= 0 && i < len(src) {
// 			src = src[:i]
// 		}
// 	}
// }

// func BenchmarkStrSlice_TakeAt_Slice(t *testing.B) {
// 	src := RangeStr(nines5)
// 	index := RangeStr(nines5)
// 	slice := NewStrSlice(src)
// 	for i := range index {
// 		slice.TakeAt(i)
// 	}
// }

// func ExampleStrSlice_TakeAt() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.TakeAt(1))
// 	// Output: 2
// }

// func TestStrSlice_TakeAt(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, Obj(nil), slice.TakeAt(0))
// 	}

// 	// all and more
// 	{
// 		slice := NewStrSliceV("0", "1", "2")
// 		assert.Equal(t, Obj("2"), slice.TakeAt(-1))
// 		assert.Equal(t, NewStrSliceV("0", "1"), slice)
// 		assert.Equal(t, Obj("1"), slice.TakeAt(-1))
// 		assert.Equal(t, NewStrSliceV("0"), slice)
// 		assert.Equal(t, Obj("0"), slice.TakeAt(-1))
// 		assert.Equal(t, NewStrSliceV(), slice)
// 		assert.Equal(t, Obj(nil), slice.TakeAt(-1))
// 		assert.Equal(t, NewStrSliceV(), slice)
// 	}

// 	// take invalid
// 	{
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj(nil), slice.TakeAt(3))
// 			assert.Equal(t, NewStrSliceV("0", "1", "2"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj(nil), slice.TakeAt(-4))
// 			assert.Equal(t, NewStrSliceV("0", "1", "2"), slice)
// 		}
// 	}

// 	// take last
// 	{
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("2"), slice.TakeAt(2))
// 			assert.Equal(t, NewStrSliceV("0", "1"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("2"), slice.TakeAt(-1))
// 			assert.Equal(t, NewStrSliceV("0", "1"), slice)
// 		}
// 	}

// 	// take middle
// 	{
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("1"), slice.TakeAt(1))
// 			assert.Equal(t, NewStrSliceV("0", "2"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("1"), slice.TakeAt(-2))
// 			assert.Equal(t, NewStrSliceV("0", "2"), slice)
// 		}
// 	}

// 	// take first
// 	{
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("0"), slice.TakeAt(0))
// 			assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewStrSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("0"), slice.TakeAt(-3))
// 			assert.Equal(t, NewStrSliceV("1", "2"), slice)
// 		}
// 	}
// }

// // TakeW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_TakeW_Go(t *testing.B) {
// 	new := []string{}
// 	src := RangeStr(nines5)
// 	l := len(src)
// 	for i := 0; i < l; i++ {
// 		if Obj(src[i]).ToInt()%2 == 0 {
// 			new = append(new, src[i])
// 			if i+1 < l {
// 				src = append(src[:i], src[i+1:]...)
// 			} else if i >= 0 && i < l {
// 				src = src[:i]
// 			}
// 			l--
// 			i--
// 		}
// 	}
// }

// func BenchmarkStrSlice_TakeW_Slice(t *testing.B) {
// 	slice := NewStrSlice(RangeStr(nines5))
// 	slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
// }

// func ExampleStrSlice_TakeW() {
// 	slice := NewStrSliceV("1", "2", "3")
// 	fmt.Println(slice.TakeW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	}))
// 	// Output: [2]
// }

// func TestStrSlice_TakeW(t *testing.T) {

// 	// take all odd values
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 != 0) })
// 		assert.Equal(t, NewStrSliceV("2", "4", "6", "8"), slice)
// 		assert.Equal(t, NewStrSliceV("1", "3", "5", "7", "9"), new)
// 	}

// 	// take all even values
// 	{
// 		slice := NewStrSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
// 		assert.Equal(t, NewStrSliceV("1", "3", "5", "7", "9"), slice)
// 		assert.Equal(t, NewStrSliceV("2", "4", "6", "8"), new)
// 	}
// }

// // Union
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Union_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkStrSlice_Union_Slice(t *testing.B) {
// 	// slice := NewStrSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleStrSlice_Union() {
// 	slice := NewStrSliceV("1", "2")
// 	fmt.Println(slice.Union([]string{"2", "3"}))
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Union(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Union(NewStrSliceV("1", "2")))
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Union([]string{"1", "2"}))
// 	}

// 	// size of one
// 	{
// 		slice := NewStrSliceV("1")
// 		union := slice.Union([]string{"1", "2", "3"})
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewStrSliceV("1"), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewStrSliceV("1", "1")
// 		union := slice.Union(NewStrSliceV("2", "3"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewStrSliceV("1", "1"), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "2", "3", "3")
// 		union := slice.Union([]string{"1", "2", "3"})
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewStrSliceV("1", "2", "2", "3", "3"), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		union := slice.Union([]string{"4", "5"})
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4", "5"), union)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").Union((*[]string)(nil)))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").Union((*StrSlice)(nil)))
// 	}
// }

// // UnionM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_UnionM_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkStrSlice_UnionM_Slice(t *testing.B) {
// 	// slice := NewStrSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleStrSlice_UnionM() {
// 	slice := NewStrSliceV("1", "2")
// 	fmt.Println(slice.UnionM([]string{"2", "3"}))
// 	// Output: [1 2 3]
// }

// func TestStrSlice_UnionM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.UnionM(NewStrSliceV("1", "2")))
// 		assert.Equal(t, (*StrSlice)(nil), slice)
// 	}

// 	// size of one
// 	{
// 		slice := NewStrSliceV("1")
// 		union := slice.UnionM([]string{"1", "2", "3"})
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewStrSliceV("1", "1")
// 		union := slice.UnionM(NewStrSliceV("2", "3"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "2", "3", "3")
// 		union := slice.UnionM([]string{"1", "2", "3"})
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		union := slice.UnionM([]string{"4", "5"})
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4", "5"), union)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4", "5"), slice)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").UnionM((*[]string)(nil)))
// 		assert.Equal(t, NewStrSliceV("1", "2"), NewStrSliceV("1", "2").UnionM((*StrSlice)(nil)))
// 	}
// }

// // Uniq
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_Uniq_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkStrSlice_Uniq_Slice(t *testing.B) {
// 	// slice := NewStrSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleStrSlice_Uniq() {
// 	slice := NewStrSliceV("1", "2", "3", "3")
// 	fmt.Println(slice.Uniq())
// 	// Output: [1 2 3]
// }

// func TestStrSlice_Uniq(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, NewStrSliceV(), slice.Uniq())
// 	}

// 	// size of one
// 	{
// 		slice := NewStrSliceV("1")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewStrSliceV("1"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1"), uniq)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewStrSliceV("1", "1")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewStrSliceV("1"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1"), uniq)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "2", "3", "3")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "2", "2", "3", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), uniq)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), uniq)
// 	}
// }

// // UniqM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStrSlice_UniqM_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkStrSlice_UniqM_Slice(t *testing.B) {
// 	// slice := NewStrSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleStrSlice_UniqM() {
// 	slice := NewStrSliceV("1", "2", "3", "3")
// 	fmt.Println(slice.UniqM())
// 	// Output: [1 2 3]
// }

// func TestStrSlice_UniqM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *StrSlice
// 		assert.Equal(t, (*StrSlice)(nil), slice.UniqM())
// 	}

// 	// size of one
// 	{
// 		slice := NewStrSliceV("1")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewStrSliceV("1"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2"), uniq)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewStrSliceV("1", "1")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewStrSliceV("1"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewStrSliceV("1", "2"), uniq)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "2", "3", "3")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), uniq)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewStrSliceV("1", "2", "3")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewStrSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewStrSliceV("1", "2", "3", "4"), uniq)
// 	}
// }

// func RangeStr(size int) (new []Str) {
// 	for _, x := range Range(0, size) {
// 		new = append(new, string(x))
// 	}
// 	return
// }
