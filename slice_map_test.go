package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Methods on both pointer and non-pointer
//--------------------------------------------------------------------------------------------------
func TestMapSlice_Methods(t *testing.T) {

	// Test conversion directly
	m, err := ToMapSliceE([]map[string]interface{}{{"foo": "bar"}})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]interface{}{{"foo": "bar"}}, m.O())

	// Test indirectly
	p := NewMapSliceV([]map[string]interface{}{{"foo": "bar"}})
	assert.Equal(t, []map[string]interface{}{{"foo": "bar"}}, p.O())

	slice := *p
	assert.Equal(t, []map[string]interface{}{}, slice.DropAt(0).O())
	assert.Equal(t, []map[string]interface{}{}, slice.O())
}

// NewMapSlice
//--------------------------------------------------------------------------------------------------
// func BenchmarkNewMapSlice_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = []string{src[i], src[i] + string(1), src[i] + string(2), src[i] + string(3), src[i] + string(4), src[i] + string(5), src[i] + string(6), src[i] + string(7), src[i] + string(8), src[i] + string(9)}
// 	}
// }

// func BenchmarkNewMapSlice_Slice(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = NewMapSlice([]string{src[i], src[i] + string(1), src[i] + string(2), src[i] + string(3), src[i] + string(4), src[i] + string(5), src[i] + string(6), src[i] + string(7), src[i] + string(8), src[i] + string(9)})
// 	}
// }

func ExampleNewMapSlice() {
	slice := NewMapSlice([]map[string]interface{}{{"foo": "bar"}})
	fmt.Println(slice)
	// Output: [&map[foo:bar]]
}

func TestMapSlice_NewMapSlice(t *testing.T) {

	// empty
	assert.Equal(t, &MapSlice{}, NewMapSlice([]map[string]interface{}{}))
	assert.Equal(t, []map[string]interface{}{}, NewMapSlice([]map[string]interface{}{}).O())

	// slice
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSlice([]map[string]interface{}{{"foo": "bar"}}))
	assert.Equal(t, []map[string]interface{}{{"foo": "bar"}}, NewMapSlice([]map[string]interface{}{{"foo": "bar"}}).O())

	// Conversion
	assert.Equal(t, []map[string]interface{}{}, NewMapSlice("1").O())
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSlice(map[string]string{"foo": "bar"}))
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSlice(map[string]interface{}{"foo": "bar"}))
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSlice([]map[string]string{{"foo": "bar"}}))
}

// NewMapSliceV
//--------------------------------------------------------------------------------------------------
// func BenchmarkNewMapSliceV_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = append([]string{}, src[i], src[i]+string(1), src[i]+string(2), src[i]+string(3), src[i]+string(4), src[i]+string(5), src[i]+string(6), src[i]+string(7), src[i]+string(8), src[i]+string(9))
// 	}
// }

// func BenchmarkNewMapSliceV_Slice(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = NewMapSliceV(src[i], src[i]+string(1), src[i]+string(2), src[i]+string(3), src[i]+string(4), src[i]+string(5), src[i]+string(6), src[i]+string(7), src[i]+string(8), src[i]+string(9))
// 	}
// }

func ExampleNewMapSliceV_empty() {
	slice := NewMapSliceV()
	fmt.Println(slice)
	// Output: []
}

func ExampleNewMapSliceV_variadic() {
	slice := NewMapSliceV(map[string]interface{}{"foo1": "1"}, map[string]interface{}{"foo2": "2"})
	fmt.Println(slice)
	// Output: [&map[foo1:1] &map[foo2:2]]
}

func TestMapSlice_NewMapSliceV(t *testing.T) {

	// empty
	assert.Equal(t, &MapSlice{}, NewMapSliceV())
	assert.Equal(t, &MapSlice{}, NewMapSliceV([]map[string]interface{}{}))
	assert.Equal(t, []map[string]interface{}{}, NewMapSliceV([]map[string]interface{}{}).O())

	// multiples
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSliceV([]map[string]interface{}{{"foo": "bar"}}))
	assert.Equal(t, &MapSlice{{"foo1": "1"}, {"foo2": "2"}}, NewMapSliceV(map[string]interface{}{"foo1": "1"}, map[string]interface{}{"foo2": "2"}))

	// Conversion
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSliceV(map[string]string{"foo": "bar"}))
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSliceV(map[string]interface{}{"foo": "bar"}))
	assert.Equal(t, &MapSlice{{"foo": "bar"}}, NewMapSliceV([]map[string]string{{"foo": "bar"}}))
}

// // Any
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Any_Go(t *testing.B) {
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
// 	src := RangeString(nines4)
// 	for _, x := range src {
// 		any(src, []string{x})
// 	}
// }

// func BenchmarkMapSlice_Any_Slice(t *testing.B) {
// 	src := RangeString(nines4)
// 	slice := NewMapSlice(src)
// 	for i := range src {
// 		slice.Any(i)
// 	}
// }

// func ExampleMapSlice_Any_empty() {
// 	slice := NewMapSliceV()
// 	fmt.Println(slice.Any())
// 	// Output: false
// }

// func ExampleMapSlice_Any_notEmpty() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Any())
// 	// Output: true
// }

// func ExampleMapSlice_Any_contains() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Any("1"))
// 	// Output: true
// }

// func ExampleMapSlice_Any_containsAny() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Any("0", "1"))
// 	// Output: true
// }

// func TestMapSlice_Any(t *testing.T) {

// 	// empty
// 	var nilSlice *MapSlice
// 	assert.False(t, nilSlice.Any())
// 	assert.False(t, NewMapSliceV().Any())

// 	// single
// 	assert.True(t, NewMapSliceV("2").Any())

// 	// invalid
// 	assert.False(t, NewMapSliceV("1", "2").Any(TestObj{"2"}))

// 	assert.True(t, NewMapSliceV("1", "2", "3").Any("2"))
// 	assert.False(t, NewMapSliceV("1", "2", "3").Any(4))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(4, "3"))
// 	assert.False(t, NewMapSliceV("1", "2", "3").Any(4, 5))

// 	// conversion
// 	assert.True(t, NewMapSliceV("1", "2").Any(Object{2}))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(int8(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(int16(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(int32('2')))
// 	assert.False(t, NewMapSliceV("1", "2", "3").Any(int32(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(int64(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(uint8('2')))
// 	assert.False(t, NewMapSliceV("1", "2", "3").Any(uint8(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(uint16(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(uint32(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(uint64(2)))
// 	assert.True(t, NewMapSliceV("1", "2", "3").Any(uint64(2)))
// }

// // AnyS
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_AnyS_Go(t *testing.B) {
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
// 	src := RangeString(nines4)
// 	for _, x := range src {
// 		any(src, []string{x})
// 	}
// }

// func BenchmarkMapSlice_AnyS_Slice(t *testing.B) {
// 	src := RangeString(nines4)
// 	slice := NewMapSlice(src)
// 	for _, x := range src {
// 		slice.Any([]string{x})
// 	}
// }

// func ExampleMapSlice_AnyS() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.AnyS([]string{"0", "1"}))
// 	// Output: true
// }

// func TestMapSlice_AnyS(t *testing.T) {
// 	// nil
// 	{
// 		var slice *MapSlice
// 		assert.False(t, slice.AnyS([]string{"1"}))
// 		assert.False(t, NewMapSliceV("1").AnyS(nil))
// 	}

// 	// []string
// 	{
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS([]string{"1"}))
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS([]string{"4", "3"}))
// 		assert.False(t, NewMapSliceV("1", "2", "3").AnyS([]string{"4", "5"}))
// 	}

// 	// *[]string
// 	{
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(&([]string{"1"})))
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(&([]string{"4", "3"})))
// 		assert.False(t, NewMapSliceV("1", "2", "3").AnyS(&([]string{"4", "5"})))
// 	}

// 	// Slice
// 	{
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(Slice(NewMapSliceV("1"))))
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(Slice(NewMapSliceV("4", "3"))))
// 		assert.False(t, NewMapSliceV("1", "2", "3").AnyS(Slice(NewMapSliceV("4", "5"))))
// 	}

// 	// MapSlice
// 	{
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(*NewMapSliceV("1")))
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(*NewMapSliceV("4", "3")))
// 		assert.False(t, NewMapSliceV("1", "2", "3").AnyS(*NewMapSliceV("4", "5")))
// 	}

// 	// *MapSlice
// 	{
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(NewMapSliceV("1")))
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(NewMapSliceV("4", "3")))
// 		assert.False(t, NewMapSliceV("1", "2", "3").AnyS(NewMapSliceV("4", "5")))
// 	}

// 	// invalid types
// 	assert.False(t, NewMapSliceV("1", "2").AnyS(nil))
// 	assert.False(t, NewMapSliceV("1", "2").AnyS((*[]string)(nil)))
// 	assert.False(t, NewMapSliceV("1", "2").AnyS((*MapSlice)(nil)))

// 	// Conversion
// 	{
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(NewMapSliceV(int64(1))))
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(NewMapSliceV(2)))
// 		assert.False(t, NewMapSliceV("1", "2", "3").AnyS(NewMapSliceV(true)))
// 		assert.True(t, NewMapSliceV("1", "2", "3").AnyS(NewMapSliceV(Char('3'))))
// 	}
// }

// // AnyW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_AnyW_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkMapSlice_AnyW_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	NewMapSlice(src).AnyW(func(x O) bool {
// 		return ExB(x.(string) == string(nines4))
// 	})
// }

// func ExampleMapSlice_AnyW() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.AnyW(func(x O) bool {
// 		return ExB(x.(string) == "2")
// 	}))
// 	// Output: true
// }

// func TestMapSlice_AnyW(t *testing.T) {

// 	// empty
// 	var slice *MapSlice
// 	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(string) > "0") }))
// 	assert.False(t, NewMapSliceV().AnyW(func(x O) bool { return ExB(x.(string) > "0") }))

// 	// single
// 	assert.True(t, NewMapSliceV("2").AnyW(func(x O) bool { return ExB(x.(string) > "0") }))
// 	assert.True(t, NewMapSliceV("1", "2").AnyW(func(x O) bool { return ExB(x.(string) == "2") }))
// 	assert.True(t, NewMapSliceV("1", "2", "3").AnyW(func(x O) bool { return ExB(x.(string) == "4" || x.(string) == "3") }))
// }

// // Append
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Append_Go(t *testing.B) {
// 	src := []string{}
// 	for _, i := range RangeString(nines6) {
// 		src = append(src, i)
// 	}
// }

// func BenchmarkMapSlice_Append_Slice(t *testing.B) {
// 	slice := NewMapSliceV()
// 	for _, i := range RangeString(nines6) {
// 		slice.Append(i)
// 	}
// }

// func ExampleMapSlice_Append() {
// 	slice := NewMapSliceV("1").Append("2").Append("3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Append(t *testing.T) {

// 	// nil
// 	{
// 		var nilSlice *MapSlice
// 		assert.Equal(t, NewMapSliceV("0"), nilSlice.Append("0"))
// 		assert.Equal(t, (*MapSlice)(nil), nilSlice)
// 	}

// 	// Append one back to back
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, true, slice.Nil())
// 		slice = NewMapSliceV()
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
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 	}

// 	// Start with just appending without chaining
// 	{
// 		slice := NewMapSliceV()
// 		assert.Equal(t, 0, slice.Len())
// 		slice.Append("1")
// 		assert.Equal(t, []string{"1"}, slice.O())
// 		slice.Append("2")
// 		assert.Equal(t, []string{"1", "2"}, slice.O())
// 	}

// 	// Start with nil not chained
// 	{
// 		slice := NewMapSliceV()
// 		assert.Equal(t, 0, slice.Len())
// 		slice.Append("1").Append("2").Append("3")
// 		assert.Equal(t, 3, slice.Len())
// 		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
// 	}

// 	// Start with nil chained
// 	{
// 		slice := NewMapSliceV().Append("1").Append("2")
// 		assert.Equal(t, 2, slice.Len())
// 		assert.Equal(t, []string{"1", "2"}, slice.O())
// 	}

// 	// Start with non nil
// 	{
// 		slice := NewMapSliceV("1").Append("2").Append("3")
// 		assert.Equal(t, 3, slice.Len())
// 		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 	}

// 	// Use append result directly
// 	{
// 		slice := NewMapSliceV("1")
// 		assert.Equal(t, 1, slice.Len())
// 		assert.Equal(t, []string{"1", "2"}, slice.Append("2").O())
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 	}

// 	// Conversion
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV(1).Append(Object{2}))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV(1).Append("2"))
// 		assert.Equal(t, NewMapSliceV("true", "2"), NewMapSliceV().Append(true).Append(Char('2')))
// 	}
// }

// // AppendV
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_AppendV_Go(t *testing.B) {
// 	src := []string{}
// 	src = append(src, RangeString(nines6)...)
// }

// func BenchmarkMapSlice_AppendV_Slice(t *testing.B) {
// 	n := NewMapSliceV()
// 	new := rangeO(0, nines6)
// 	n.AppendV(new...)
// }

// func ExampleMapSlice_AppendV() {
// 	slice := NewMapSliceV("1").AppendV("2", "3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestMapSlice_AppendV(t *testing.T) {

// 	// nil
// 	{
// 		var nilSlice *MapSlice
// 		assert.Equal(t, NewMapSliceV("1", "2"), nilSlice.AppendV("1", "2"))
// 	}

// 	// Append many src
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1").AppendV("2", "3"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4", "5"), NewMapSliceV("1").AppendV("2", "3").AppendV("4", "5"))
// 	}

// 	// Conversion
// 	{
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV().AppendV(Object{0}, Object{1}))
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV().AppendV("0", "1"))
// 		assert.Equal(t, NewMapSliceV("false", "true"), NewMapSliceV().AppendV(false, true))
// 	}
// }

// // At
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_At_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for _, x := range src {
// 		assert.IsType(t, 0, x)
// 	}
// }

// func BenchmarkMapSlice_At_Slice(t *testing.B) {
// 	src := RangeString(nines6)
// 	slice := NewMapSlice(src)
// 	for i := 0; i < len(src); i++ {
// 		_, ok := (slice.At(i).O()).(string)
// 		assert.True(t, ok)
// 	}
// }

// func ExampleMapSlice_At() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.At(2))
// 	// Output: 3
// }

// func TestMapSlice_At(t *testing.T) {

// 	// nil
// 	{
// 		var nilSlice *MapSlice
// 		assert.Equal(t, Obj(nil), nilSlice.At(0))
// 	}

// 	// src
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4")
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
// 		slice := NewMapSliceV("1")
// 		assert.Equal(t, &Object{}, slice.At(3))
// 		assert.Equal(t, nil, slice.At(3).O())
// 		assert.Equal(t, &Object{}, slice.At(-3))
// 		assert.Equal(t, nil, slice.At(-3).O())
// 	}
// }

// // Clear
// //--------------------------------------------------------------------------------------------------
// func ExampleMapSlice_Clear() {
// 	slice := NewMapSliceV("1").Concat([]string{"2", "3"})
// 	fmt.Println(slice.Clear())
// 	// Output: []
// }

// func TestMapSlice_Clear(t *testing.T) {

// 	// nil
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.Clear())
// 		assert.Equal(t, (*MapSlice)(nil), slice)
// 	}

// 	// int
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewMapSliceV(), slice.Clear())
// 		assert.Equal(t, NewMapSliceV(), slice.Clear())
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}
// }

// // Concat
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Concat_Go(t *testing.B) {
// 	dest := []string{}
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkMapSlice_Concat_Slice(t *testing.B) {
// 	dest := NewMapSliceV()
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.Concat(src[j:i])
// 		j = i
// 	}
// }

// func ExampleMapSlice_Concat() {
// 	slice := NewMapSliceV("1").Concat([]string{"2", "3"})
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Concat(t *testing.T) {

// 	// nil
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Concat([]string{"1", "2"}))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").Concat(nil))
// 	}

// 	// []string
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.Concat([]string{"2", "3"})
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), concated)
// 	}

// 	// *[]string
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.Concat(&([]string{"2", "3"}))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), concated)
// 	}

// 	// *MapSlice
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.Concat(NewMapSliceV("2", "3"))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), concated)
// 	}

// 	// MapSlice
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.Concat(*NewMapSliceV("2", "3"))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), concated)
// 	}

// 	// Slice
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.Concat(Slice(NewMapSliceV("2", "3")))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), concated)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").Concat((*[]string)(nil)))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").Concat((*MapSlice)(nil)))
// 	}

// 	// Conversion
// 	{
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV().Concat([]Object{{0}, {1}}))
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV().Concat([]string{"0", "1"}))
// 		assert.Equal(t, NewMapSliceV("false", "true"), NewMapSliceV().Concat([]bool{false, true}))

// 		slice := NewMapSliceV(Object{1})
// 		concated := slice.Concat([]int64{2, 3})
// 		assert.Equal(t, NewMapSliceV("1", "4"), slice.Append(Char('4')))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), concated)
// 	}
// }

// // ConcatM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_ConcatM_Go(t *testing.B) {
// 	dest := []string{}
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkMapSlice_ConcatM_Slice(t *testing.B) {
// 	dest := NewMapSliceV()
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.ConcatM(src[j:i])
// 		j = i
// 	}
// }

// func ExampleMapSlice_ConcatM() {
// 	slice := NewMapSliceV("1").ConcatM([]string{"2", "3"})
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestMapSlice_ConcatM(t *testing.T) {

// 	// nil
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.ConcatM([]string{"1", "2"}))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").ConcatM(nil))
// 	}

// 	// []string
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.ConcatM([]string{"2", "3"})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// *[]string
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.ConcatM(&([]string{"2", "3"}))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// *MapSlice
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.ConcatM(NewMapSliceV("2", "3"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// MapSlice
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.ConcatM(*NewMapSliceV("2", "3"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// Slice
// 	{
// 		slice := NewMapSliceV("1")
// 		concated := slice.ConcatM(Slice(NewMapSliceV("2", "3")))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), concated)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").ConcatM((*[]string)(nil)))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").ConcatM((*MapSlice)(nil)))
// 	}

// 	// Conversion
// 	{
// 		slice := NewMapSliceV(Object{1})
// 		concated := slice.ConcatM([]Object{{2}, {3}})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append(Char('4')))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), concated)
// 	}
// }

// // Copy
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Copy_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	dst := make([]string, len(src), len(src))
// 	copy(dst, src)
// }

// func BenchmarkMapSlice_Copy_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.Copy()
// }

// func ExampleMapSlice_Copy() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Copy())
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Copy(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.Copy(0, -1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV("0").Clear().Copy(0, -1))
// 	}

// 	// Test that the original is NOT modified when the slice is modified
// 	{
// 		original := NewMapSliceV("1", "2", "3")
// 		result := original.Copy(0, -1)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), original)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), result)
// 		result.Set(0, "0")
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), original)
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), result)
// 	}

// 	// copy full array
// 	{
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Copy())
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Copy(0, -1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Copy(0, 1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Copy(0, 5))
// 		assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1").Copy())
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Copy())
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Copy(0, -1))
// 		assert.Equal(t, NewMapSlice([]string{"1", "2", "3"}), NewMapSlice([]string{"1", "2", "3"}).Copy())
// 		assert.Equal(t, NewMapSlice([]string{"1", "2", "3"}), NewMapSlice([]string{"1", "2", "3"}).Copy(0, -1))
// 	}

// 	// out of bounds should be moved in
// 	{
// 		assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1").Copy(0, 2))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Copy(-6, 6))
// 	}

// 	// mutually exclusive
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewMapSliceV(), slice.Copy(2, -3))
// 		assert.Equal(t, NewMapSliceV(), slice.Copy(0, -5))
// 		assert.Equal(t, NewMapSliceV(), slice.Copy(4, -1))
// 		assert.Equal(t, NewMapSliceV(), slice.Copy(6, -1))
// 		assert.Equal(t, NewMapSliceV(), slice.Copy(3, -2))
// 	}

// 	// singles
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewMapSliceV("4"), slice.Copy(-1, -1))
// 		assert.Equal(t, NewMapSliceV("3"), slice.Copy(-2, -2))
// 		assert.Equal(t, NewMapSliceV("2"), slice.Copy(-3, -3))
// 		assert.Equal(t, NewMapSliceV("1"), slice.Copy(0, 0))
// 		assert.Equal(t, NewMapSliceV("1"), slice.Copy(-4, -4))
// 		assert.Equal(t, NewMapSliceV("2"), slice.Copy(1, 1))
// 		assert.Equal(t, NewMapSliceV("2"), slice.Copy(1, -3))
// 		assert.Equal(t, NewMapSliceV("3"), slice.Copy(2, 2))
// 		assert.Equal(t, NewMapSliceV("3"), slice.Copy(2, -2))
// 		assert.Equal(t, NewMapSliceV("4"), slice.Copy(3, 3))
// 		assert.Equal(t, NewMapSliceV("4"), slice.Copy(3, -1))
// 	}

// 	// grab all but first
// 	{
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Copy(1, -1))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Copy(1, 2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Copy(-2, -1))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Copy(-2, 2))
// 	}

// 	// grab all but last
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Copy(0, -2))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Copy(-3, -2))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Copy(-3, 1))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Copy(0, 1))
// 	}

// 	// grab middle
// 	{
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Copy(1, -2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Copy(-3, -2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Copy(-3, 2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Copy(1, 2))
// 	}

// 	// random
// 	{
// 		assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1", "2", "3").Copy(0, -3))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Copy(1, 2))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Copy(0, 2))
// 	}
// }

// // Count
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Count_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkMapSlice_Count_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	NewMapSlice(src).Count(nines4)
// }

// func ExampleMapSlice_Count() {
// 	slice := NewMapSliceV("1", "2", "2")
// 	fmt.Println(slice.Count("2"))
// 	// Output: 2
// }

// func TestMapSlice_Count(t *testing.T) {

// 	// empty
// 	var slice *MapSlice
// 	assert.Equal(t, 0, slice.Count(0))
// 	assert.Equal(t, 0, NewMapSliceV().Count(0))

// 	assert.Equal(t, 1, NewMapSliceV("2", "3").Count("2"))
// 	assert.Equal(t, 2, NewMapSliceV("1", "2", "2").Count("2"))
// 	assert.Equal(t, 4, NewMapSliceV("4", "4", "3", "4", "4").Count("4"))
// 	assert.Equal(t, 3, NewMapSliceV("3", "2", "3", "3", "5").Count("3"))
// 	assert.Equal(t, 1, NewMapSliceV("1", "2", "3").Count("3"))
// }

// // CountW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_CountW_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkMapSlice_CountW_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	NewMapSlice(src).CountW(func(x O) bool {
// 		return ExB(x.(string) == string(nines4))
// 	})
// }

// func ExampleMapSlice_CountW() {
// 	slice := NewMapSliceV("1", "2", "2")
// 	fmt.Println(slice.CountW(func(x O) bool {
// 		return ExB(x.(string) == "2")
// 	}))
// 	// Output: 2
// }

// func TestMapSlice_CountW(t *testing.T) {

// 	// empty
// 	var slice *MapSlice
// 	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(string) > "0") }))
// 	assert.Equal(t, 0, NewMapSliceV().CountW(func(x O) bool { return ExB(x.(string) > "0") }))

// 	assert.Equal(t, 1, NewMapSliceV("2", "3").CountW(func(x O) bool { return ExB(x.(string) > "2") }))
// 	assert.Equal(t, 1, NewMapSliceV("1", "2").CountW(func(x O) bool { return ExB(x.(string) == "2") }))
// 	assert.Equal(t, 1, NewMapSliceV("1", "2", "3").CountW(func(x O) bool { return ExB(x.(string) == "4" || x.(string) == "3") }))
// }

// // Drop
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Drop_Go(t *testing.B) {
// 	src := RangeString(nines7)
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

// func BenchmarkMapSlice_Drop_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 1 {
// 		slice.Drop(1, 10)
// 	}
// }

// func ExampleMapSlice_Drop() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Drop(0, 1))
// 	// Output: [3]
// }

// func TestMapSlice_Drop(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.Drop(0, 1))
// 	}

// 	// invalid
// 	assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(1))
// 	assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(4, 4))

// 	// drop 1
// 	assert.Equal(t, NewMapSliceV("2", "3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(0, 0))
// 	assert.Equal(t, NewMapSliceV("1", "3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(1, 1))
// 	assert.Equal(t, NewMapSliceV("1", "2", "4"), NewMapSliceV("1", "2", "3", "4").Drop(2, 2))
// 	assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3", "4").Drop(3, 3))
// 	assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3", "4").Drop(-1, -1))
// 	assert.Equal(t, NewMapSliceV("1", "2", "4"), NewMapSliceV("1", "2", "3", "4").Drop(-2, -2))
// 	assert.Equal(t, NewMapSliceV("1", "3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(-3, -3))
// 	assert.Equal(t, NewMapSliceV("2", "3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(-4, -4))

// 	// drop 2
// 	assert.Equal(t, NewMapSliceV("3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(0, 1))
// 	assert.Equal(t, NewMapSliceV("1", "4"), NewMapSliceV("1", "2", "3", "4").Drop(1, 2))
// 	assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3", "4").Drop(2, 3))
// 	assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3", "4").Drop(-2, -1))
// 	assert.Equal(t, NewMapSliceV("1", "4"), NewMapSliceV("1", "2", "3", "4").Drop(-3, -2))
// 	assert.Equal(t, NewMapSliceV("3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(-4, -3))

// 	// drop 3
// 	assert.Equal(t, NewMapSliceV("4"), NewMapSliceV("1", "2", "3", "4").Drop(0, 2))
// 	assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1", "2", "3", "4").Drop(-3, -1))

// 	// drop everything and beyond
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Drop())
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Drop(0, 3))
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Drop(0, -1))
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Drop(-4, -1))
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Drop(-6, -1))
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Drop(0, 10))

// 	// move index within bounds
// 	assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3", "4").Drop(3, 4))
// 	assert.Equal(t, NewMapSliceV("2", "3", "4"), NewMapSliceV("1", "2", "3", "4").Drop(-5, 0))
// }

// // DropAt
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_DropAt_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	index := Range(0, nines5)
// 	for i := range index {
// 		if i+1 < len(src) {
// 			src = append(src[:i], src[i+1:]...)
// 		} else if i >= 0 && i < len(src) {
// 			src = src[:i]
// 		}
// 	}
// }

// func BenchmarkMapSlice_DropAt_Slice(t *testing.B) {
// 	index := Range(0, nines5)
// 	slice := NewMapSlice(RangeString(nines5))
// 	for i := range index {
// 		slice.DropAt(i)
// 	}
// }

// func ExampleMapSlice_DropAt() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.DropAt(1))
// 	// Output: [1 3]
// }

// func TestMapSlice_DropAt(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.DropAt(0))
// 	}

// 	// drop all and more
// 	{
// 		slice := NewMapSliceV("0", "1", "2")
// 		assert.Equal(t, NewMapSliceV("0", "1"), slice.DropAt(-1))
// 		assert.Equal(t, NewMapSliceV("0"), slice.DropAt(-1))
// 		assert.Equal(t, NewMapSliceV(), slice.DropAt(-1))
// 		assert.Equal(t, NewMapSliceV(), slice.DropAt(-1))
// 	}

// 	// drop invalid
// 	assert.Equal(t, NewMapSliceV("0", "1", "2"), NewMapSliceV("0", "1", "2").DropAt(3))
// 	assert.Equal(t, NewMapSliceV("0", "1", "2"), NewMapSliceV("0", "1", "2").DropAt(-4))

// 	// drop last
// 	assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV("0", "1", "2").DropAt(2))
// 	assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV("0", "1", "2").DropAt(-1))

// 	// drop middle
// 	assert.Equal(t, NewMapSliceV("0", "2"), NewMapSliceV("0", "1", "2").DropAt(1))
// 	assert.Equal(t, NewMapSliceV("0", "2"), NewMapSliceV("0", "1", "2").DropAt(-2))

// 	// drop first
// 	assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("0", "1", "2").DropAt(0))
// 	assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("0", "1", "2").DropAt(-3))
// }

// // DropFirst
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_DropFirst_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkMapSlice_DropFirst_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirst()
// 	}
// }

// func ExampleMapSlice_DropFirst() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.DropFirst())
// 	// Output: [2 3]
// }

// func TestMapSlice_DropFirst(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.DropFirst())
// 	}

// 	// drop all and beyond
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("2", "3"), slice.DropFirst())
// 		assert.Equal(t, NewMapSliceV("3"), slice.DropFirst())
// 		assert.Equal(t, NewMapSliceV(), slice.DropFirst())
// 		assert.Equal(t, NewMapSliceV(), slice.DropFirst())
// 	}
// }

// // DropFirstN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_DropFirstN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkMapSlice_DropFirstN_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirstN(10)
// 	}
// }

// func ExampleMapSlice_DropFirstN() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.DropFirstN(2))
// 	// Output: [3]
// }

// func TestMapSlice_DropFirstN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.DropFirstN(1))
// 	}

// 	// negative value
// 	assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").DropFirstN(-1))

// 	// drop none
// 	assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").DropFirstN(0))

// 	// drop 1
// 	assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").DropFirstN(1))

// 	// drop 2
// 	assert.Equal(t, NewMapSliceV("3"), NewMapSliceV("1", "2", "3").DropFirstN(2))

// 	// drop 3
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3").DropFirstN(3))

// 	// drop beyond
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3").DropFirstN(4))
// }

// // DropLast
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_DropLast_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkMapSlice_DropLast_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLast()
// 	}
// }

// func ExampleMapSlice_DropLast() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.DropLast())
// 	// Output: [1 2]
// }

// func TestMapSlice_DropLast(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.DropLast())
// 	}

// 	// negative value
// 	assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").DropLastN(-1))

// 	slice := NewMapSliceV("1", "2", "3")
// 	assert.Equal(t, NewMapSliceV("1", "2"), slice.DropLast())
// 	assert.Equal(t, NewMapSliceV("1"), slice.DropLast())
// 	assert.Equal(t, NewMapSliceV(), slice.DropLast())
// 	assert.Equal(t, NewMapSliceV(), slice.DropLast())
// }

// // DropLastN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_DropLastN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkMapSlice_DropLastN_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLastN(10)
// 	}
// }

// func ExampleMapSlice_DropLastN() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.DropLastN(2))
// 	// Output: [1]
// }

// func TestMapSlice_DropLastN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.DropLastN(1))
// 	}

// 	// drop none
// 	assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").DropLastN(0))

// 	// drop 1
// 	assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").DropLastN(1))

// 	// drop 2
// 	assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1", "2", "3").DropLastN(2))

// 	// drop 3
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3").DropLastN(3))

// 	// drop beyond
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3").DropLastN(4))
// }

// // DropW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_DropW_Go(t *testing.B) {
// 	src := RangeString(nines5)
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

// func BenchmarkMapSlice_DropW_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines5))
// 	slice.DropW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	})
// }

// func ExampleMapSlice_DropW() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.DropW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	}))
// 	// Output: [1 3]
// }

// func TestMapSlice_DropW(t *testing.T) {

// 	// drop all odd values
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		slice.DropW(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 != 0)
// 		})
// 		assert.Equal(t, NewMapSliceV("2", "4", "6", "8"), slice)
// 	}

// 	// drop all even values
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		slice.DropW(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 == 0)
// 		})
// 		assert.Equal(t, NewMapSliceV("1", "3", "5", "7", "9"), slice)
// 	}
// }

// Each
//--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Each_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_Each_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).Each(func(x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

func ExampleMapSlice_Each() {
	NewMapSliceV([]map[string]interface{}{{"foo": "bar"}}).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: map[foo:bar]
}

func TestMapSlice_Each(t *testing.T) {

	// nil or empty
	{
		var slice *MapSlice
		slice.Each(func(x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewMapSliceV([]map[string]interface{}{{"foo": "1"}, {"foo": "2"}}).Each(func(x O) {
			results = append(results, (x.(map[string]interface{}))["foo"].(string))
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// EachE
//--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_EachE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "0", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_EachE_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).EachE(func(x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

func ExampleMapSlice_EachE() {
	NewMapSliceV([]map[string]interface{}{{"foo": "bar"}}).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: map[foo:bar]
}

func TestMapSlice_EachE(t *testing.T) {

	// nil or empty
	{
		var slice *MapSlice
		slice.EachE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewMapSliceV([]map[string]interface{}{{"foo": "1"}, {"foo": "2"}, {"foo": "3"}}).EachE(func(x O) error {
			results = append(results, (x.(map[string]interface{}))["foo"].(string))
			return nil
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewMapSliceV([]map[string]interface{}{{"foo": "1"}, {"foo": "2"}, {"foo": "3"}}).EachE(func(x O) error {
			if (x.(map[string]interface{}))["foo"].(string) == "3" {
				return Break
			}
			results = append(results, (x.(map[string]interface{}))["foo"].(string))
			return nil
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// // EachI
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_EachI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_EachI_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).EachI(func(i int, x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleMapSlice_EachI() {
// 	NewMapSliceV("1", "2", "3").EachI(func(i int, x O) {
// 		fmt.Printf("%v:%v", i, x)
// 	})
// 	// Output: 0:11:22:3
// }

// func TestMapSlice_EachI(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		slice.EachI(func(i int, x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachI(func(i int, x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}
// }

// // EachIE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_EachIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_EachIE_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).EachIE(func(i int, x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleMapSlice_EachIE() {
// 	NewMapSliceV("1", "2", "3").EachIE(func(i int, x O) error {
// 		fmt.Printf("%v:%v", i, x)
// 		return nil
// 	})
// 	// Output: 0:11:22:3
// }

// func TestMapSlice_EachIE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		slice.EachIE(func(i int, x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachIE(func(i int, x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachIE(func(i int, x O) error {
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
// func BenchmarkMapSlice_EachR_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_EachR_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).EachR(func(x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleMapSlice_EachR() {
// 	NewMapSliceV("1", "2", "3").EachR(func(x O) {
// 		fmt.Printf("%v", x)
// 	})
// 	// Output: 321
// }

// func TestMapSlice_EachR(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		slice.EachR(func(x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachR(func(x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}
// }

// // EachRE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_EachRE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_EachRE_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).EachRE(func(x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleMapSlice_EachRE() {
// 	NewMapSliceV("1", "2", "3").EachRE(func(x O) error {
// 		fmt.Printf("%v", x)
// 		return nil
// 	})
// 	// Output: 321
// }

// func TestMapSlice_EachRE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		slice.EachRE(func(x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachRE(func(x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachRE(func(x O) error {
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
// func BenchmarkMapSlice_EachRI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_EachRI_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).EachRI(func(i int, x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleMapSlice_EachRI() {
// 	NewMapSliceV("1", "2", "3").EachRI(func(i int, x O) {
// 		fmt.Printf("%v:%v", i, x)
// 	})
// 	// Output: 2:31:20:1
// }

// func TestMapSlice_EachRI(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		slice.EachRI(func(i int, x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachRI(func(i int, x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}
// }

// // EachRIE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_EachRIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkMapSlice_EachRIE_Slice(t *testing.B) {
// 	NewMapSlice(RangeString(nines6)).EachRIE(func(i int, x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleMapSlice_EachRIE() {
// 	NewMapSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
// 		fmt.Printf("%v:%v", i, x)
// 		return nil
// 	})
// 	// Output: 2:31:20:1
// }

// func TestMapSlice_EachRIE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		slice.EachRIE(func(i int, x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewMapSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
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
// func ExampleMapSlice_Empty() {
// 	fmt.Println(NewMapSliceV().Empty())
// 	// Output: true
// }

// func TestMapSlice_Empty(t *testing.T) {

// 	// nil or empty
// 	{
// 		var nilSlice *MapSlice
// 		assert.Equal(t, true, nilSlice.Empty())
// 	}

// 	assert.Equal(t, true, NewMapSliceV().Empty())
// 	assert.Equal(t, false, NewMapSliceV("1").Empty())
// 	assert.Equal(t, false, NewMapSliceV("1", "2", "3").Empty())
// 	assert.Equal(t, false, NewMapSliceV("1").Empty())
// 	assert.Equal(t, false, NewMapSlice([]string{"1", "2", "3"}).Empty())
// }

// // First
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_First_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		_ = src[0]
// 		src = src[1:]
// 	}
// }

// func BenchmarkMapSlice_First_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.First()
// 		slice.DropFirst()
// 	}
// }

// func ExampleMapSlice_First() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.First())
// 	// Output: 1
// }

// func TestMapSlice_First(t *testing.T) {
// 	// invalid
// 	assert.Equal(t, Obj(nil), NewMapSliceV().First())

// 	// int
// 	assert.Equal(t, Obj("2"), NewMapSliceV("2", "3").First())
// 	assert.Equal(t, Obj("3"), NewMapSliceV("3", "2").First())
// 	assert.Equal(t, Obj("1"), NewMapSliceV("1", "3", "2").First())
// }

// // FirstN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_FirstN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	_ = src[0:10]
// }

// func BenchmarkMapSlice_FirstN_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	slice.FirstN(10)
// }

// func ExampleMapSlice_FirstN() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.FirstN(2))
// 	// Output: [1 2]
// }

// func TestMapSlice_FirstN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.FirstN(1))
// 		assert.Equal(t, NewMapSliceV(), slice.FirstN(-1))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewMapSliceV("1", "2", "3")
// 		result := original.FirstN(2).Set(0, "0")
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), original)
// 		assert.Equal(t, NewMapSliceV("0", "2"), result)
// 	}

// 	// Get none
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3").FirstN(0))

// 	// slice full array includeing out of bounds
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV().FirstN(1))
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV().FirstN(10))
// 	assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").FirstN(10))
// 	assert.Equal(t, NewMapSlice([]string{"1", "2", "3"}), NewMapSlice([]string{"1", "2", "3"}).FirstN(10))

// 	// grab a few diff
// 	assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1", "2", "3").FirstN(1))
// 	assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").FirstN(2))
// }

// // G
// //--------------------------------------------------------------------------------------------------
// func ExampleMapSlice_G() {
// 	fmt.Println(NewMapSliceV("1", "2", "3").G())
// 	// Output: [1 2 3]
// }

// func TestMapSlice_G(t *testing.T) {
// 	assert.IsType(t, []string{}, NewMapSliceV().G())
// 	assert.IsType(t, []string{"1", "2", "3"}, NewMapSliceV("1", "2", "3").G())
// }

// // Generic
// //--------------------------------------------------------------------------------------------------
// func ExampleMapSlice_Generic() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Generic())
// 	// Output: false
// }

// // Index
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Index_Go(t *testing.B) {
// 	for _, x := range RangeString(nines5) {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkMapSlice_Index_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines5))
// 	slice.Index(nines4)
// }

// func ExampleMapSlice_Index() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Index("2"))
// 	// Output: 1
// }

// func TestMapSlice_Index(t *testing.T) {

// 	// empty
// 	var slice *MapSlice
// 	assert.Equal(t, -1, slice.Index("2"))
// 	assert.Equal(t, -1, NewMapSliceV().Index("1"))

// 	assert.Equal(t, 0, NewMapSliceV("1", "2", "3").Index("1"))
// 	assert.Equal(t, 1, NewMapSliceV("1", "2", "3").Index("2"))
// 	assert.Equal(t, 2, NewMapSliceV("1", "2", "3").Index("3"))
// 	assert.Equal(t, -1, NewMapSliceV("1", "2", "3").Index("4"))
// 	assert.Equal(t, -1, NewMapSliceV("1", "2", "3").Index("5"))

// 	// Conversion
// 	{
// 		assert.Equal(t, 1, NewMapSliceV("1", "2", "3").Index(Object{2}))
// 		assert.Equal(t, 1, NewMapSliceV("1", "2", "3").Index("2"))
// 		assert.Equal(t, -1, NewMapSliceV("1", "2", "3").Index(true))
// 		assert.Equal(t, 2, NewMapSliceV("1", "2", "3").Index(Char('3')))
// 	}
// }

// // Insert
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Insert_Go(t *testing.B) {
// 	src := []string{}
// 	for _, x := range RangeString(nines6) {
// 		src = append(src, x)
// 		copy(src[1:], src[1:])
// 		src[0] = x
// 	}
// }

// func BenchmarkMapSlice_Insert_Slice(t *testing.B) {
// 	slice := NewMapSliceV()
// 	for x := range RangeString(nines6) {
// 		slice.Insert(0, x)
// 	}
// }

// func ExampleMapSlice_Insert() {
// 	slice := NewMapSliceV("1", "3")
// 	fmt.Println(slice.Insert(1, "2"))
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Insert(t *testing.T) {

// 	// append
// 	{
// 		slice := NewMapSliceV()
// 		assert.Equal(t, NewMapSliceV("0"), slice.Insert(-1, "0"))
// 		assert.Equal(t, NewMapSliceV("0", "1"), slice.Insert(-1, "1"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2"), slice.Insert(-1, "2"))
// 	}

// 	// [] append
// 	{
// 		slice := NewMapSliceV()
// 		assert.Equal(t, NewMapSliceV("0"), slice.Insert(-1, []string{"0"}))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2"), slice.Insert(-1, []string{"1", "2"}))
// 	}

// 	// prepend
// 	{
// 		slice := NewMapSliceV()
// 		assert.Equal(t, NewMapSliceV("2"), slice.Insert(0, "2"))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Insert(0, "1"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2"), slice.Insert(0, "0"))
// 	}

// 	// [] prepend
// 	{
// 		slice := NewMapSliceV()
// 		assert.Equal(t, NewMapSliceV("2"), slice.Insert(0, []string{"2"}))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2"), slice.Insert(0, []string{"0", "1"}))
// 	}

// 	// middle pos
// 	{
// 		slice := NewMapSliceV("0", "5")
// 		assert.Equal(t, NewMapSliceV("0", "1", "5"), slice.Insert(1, "1"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "5"), slice.Insert(2, "2"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "3", "5"), slice.Insert(3, "3"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(4, "4"))
// 	}

// 	// [] middle pos
// 	{
// 		slice := NewMapSliceV("0", "5")
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "5"), slice.Insert(1, []string{"1", "2"}))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(3, []string{"3", "4"}))
// 	}

// 	// middle neg
// 	{
// 		slice := NewMapSliceV("0", "5")
// 		assert.Equal(t, NewMapSliceV("0", "1", "5"), slice.Insert(-2, "1"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "5"), slice.Insert(-2, "2"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "3", "5"), slice.Insert(-2, "3"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(-2, "4"))
// 	}

// 	// [] middle neg
// 	{
// 		slice := NewMapSliceV(0, 5)
// 		assert.Equal(t, NewMapSliceV(0, 1, 2, 5), slice.Insert(-2, []string{"1", "2"}))
// 		assert.Equal(t, NewMapSliceV(0, "1", "2", "3", 4, 5), slice.Insert(-2, []int{3, 4}))
// 	}

// 	// error cases
// 	{
// 		var slice *MapSlice
// 		assert.False(t, slice.Insert(0, 0).Nil())
// 		assert.Equal(t, NewMapSliceV("0"), slice.Insert(0, "0"))
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV("0", "1").Insert(-10, "1"))
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV("0", "1").Insert(10, "1"))
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV("0", "1").Insert(2, "1"))
// 		assert.Equal(t, NewMapSliceV("0", "1"), NewMapSliceV("0", "1").Insert(-3, "1"))
// 	}

// 	// [] error cases
// 	{
// 		var slice *MapSlice
// 		assert.False(t, slice.Insert(0, 0).Nil())
// 		assert.Equal(t, NewMapSliceV(0), slice.Insert(0, 0))
// 		assert.Equal(t, NewMapSliceV(0, 1), NewMapSliceV(0, 1).Insert(-10, 1))
// 		assert.Equal(t, NewMapSliceV(0, 1), NewMapSliceV(0, 1).Insert(10, 1))
// 		assert.Equal(t, NewMapSliceV(0, 1), NewMapSliceV(0, 1).Insert(2, 1))
// 		assert.Equal(t, NewMapSliceV(0, 1), NewMapSliceV(0, 1).Insert(-3, 1))
// 	}

// 	// Conversion
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV(1, 3).Insert(1, Object{2}))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV(1, 3).Insert(1, "2"))
// 		assert.Equal(t, NewMapSliceV(true, "2", "3"), NewMapSliceV(2, 3).Insert(0, true))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV(1, 2).Insert(-1, Char('3')))
// 	}

// 	// [] Conversion
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", 4), NewMapSliceV(1, 4).Insert(1, []Object{{2}, {3}}))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", 4), NewMapSliceV(1, 4).Insert(1, []string{"2", "3"}))
// 		assert.Equal(t, NewMapSliceV(false, true, "2", "3"), NewMapSliceV(2, 3).Insert(0, []bool{false, true}))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", 4), NewMapSliceV(1, 2).Insert(-1, []Char{'3', '4'}))
// 	}
// }

// // Join
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Join_Go(t *testing.B) {
// 	src := RangeString(nines4)
// 	strings.Join(src, ",")
// }

// func BenchmarkMapSlice_Join_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines4))
// 	slice.Join()
// }

// func ExampleMapSlice_Join() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Join())
// 	// Output: 1,2,3
// }

// func TestMapSlice_Join(t *testing.T) {
// 	// nil
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, Obj(""), slice.Join())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, Obj(""), NewMapSliceV().Join())
// 	}

// 	assert.Equal(t, "1,2,3", NewMapSliceV("1", "2", "3").Join().O())
// 	assert.Equal(t, "1.2.3", NewMapSliceV("1", "2", "3").Join(".").O())
// }

// // Last
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Last_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		_ = src[len(src)-1]
// 		src = src[:len(src)-1]
// 	}
// }

// func BenchmarkMapSlice_Last_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.Last()
// 		slice.DropLast()
// 	}
// }

// func ExampleMapSlice_Last() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Last())
// 	// Output: 3
// }

// func TestMapSlice_Last(t *testing.T) {
// 	// invalid
// 	assert.Equal(t, Obj(nil), NewMapSliceV().Last())

// 	// int
// 	assert.Equal(t, Obj("3"), NewMapSliceV("2", "3").Last())
// 	assert.Equal(t, Obj("2"), NewMapSliceV("3", "2").Last())
// 	assert.Equal(t, Obj("2"), NewMapSliceV("1", "3", "2").Last())
// }

// // LastN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_LastN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	_ = src[0:10]
// }

// func BenchmarkMapSlice_LastN_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	slice.LastN(10)
// }

// func ExampleMapSlice_LastN() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.LastN(2))
// 	// Output: [2 3]
// }

// func TestMapSlice_LastN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.LastN(1))
// 		assert.Equal(t, NewMapSliceV(), slice.LastN(-1))
// 	}

// 	// Get none
// 	{
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3").LastN(0))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewMapSliceV("1", "2", "3")
// 		result := original.LastN(2).Set(0, "0")
// 		assert.Equal(t, NewMapSliceV("1", "0", "3"), original)
// 		assert.Equal(t, NewMapSliceV("0", "3"), result)
// 	}

// 	// slice full array includeing out of bounds
// 	{
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().LastN(1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().LastN(10))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").LastN(10))
// 		assert.Equal(t, NewMapSlice([]string{"1", "2", "3"}), NewMapSlice([]string{"1", "2", "3"}).LastN(10))
// 	}

// 	// grab a few diff
// 	{
// 		assert.Equal(t, NewMapSliceV("3"), NewMapSliceV("1", "2", "3").LastN(1))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").LastN(2))
// 	}
// }

// // Len
// //--------------------------------------------------------------------------------------------------
// func ExampleMapSlice_Len() {
// 	fmt.Println(NewMapSliceV("1", "2", "3").Len())
// 	// Output: 3
// }

// func TestMapSlice_Len(t *testing.T) {
// 	assert.Equal(t, 0, NewMapSliceV().Len())
// 	assert.Equal(t, 2, len(*(NewMapSliceV("1", "2"))))
// 	assert.Equal(t, 2, NewMapSliceV("1", "2").Len())
// }

// // Less
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Less_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if i+1 < len(src) {
// 			_ = src[i] < src[i+1]
// 		}
// 	}
// }

// func BenchmarkMapSlice_Less_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Less(i, i+1)
// 		}
// 	}
// }

// func ExampleMapSlice_Less() {
// 	slice := NewMapSliceV("2", "3", "1")
// 	fmt.Println(slice.Less(0, 2))
// 	// Output: false
// }

// func TestMapSlice_Less(t *testing.T) {

// 	// invalid cases
// 	{
// 		var slice *MapSlice
// 		assert.False(t, slice.Less(0, 0))

// 		slice = NewMapSliceV()
// 		assert.False(t, slice.Less(0, 0))
// 		assert.False(t, slice.Less(1, 2))
// 		assert.False(t, slice.Less(-1, 2))
// 		assert.False(t, slice.Less(1, -2))
// 	}

// 	// valid
// 	assert.Equal(t, true, NewMapSliceV("0", "1", "2").Less(0, 1))
// 	assert.Equal(t, false, NewMapSliceV("0", "1", "2").Less(1, 0))
// 	assert.Equal(t, true, NewMapSliceV("0", "1", "2").Less(1, 2))
// }

// // Nil
// //--------------------------------------------------------------------------------------------------
// func ExampleMapSlice_Nil() {
// 	var slice *MapSlice
// 	fmt.Println(slice.Nil())
// 	// Output: true
// }

// func TestMapSlice_Nil(t *testing.T) {
// 	var slice *MapSlice
// 	assert.True(t, slice.Nil())
// 	assert.False(t, NewMapSliceV().Nil())
// 	assert.False(t, NewMapSliceV("1", "2", "3").Nil())
// }

// // O
// //--------------------------------------------------------------------------------------------------
// func ExampleMapSlice_O() {
// 	fmt.Println(NewMapSliceV("1", "2", "3"))
// 	// Output: [1 2 3]
// }

// func TestMapSlice_O(t *testing.T) {
// 	assert.Equal(t, []string{}, (*MapSlice)(nil).O())
// 	assert.Equal(t, []string{}, NewMapSliceV().O())
// 	assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3"))
// }

// // Pair
// //--------------------------------------------------------------------------------------------------

// func ExampleMapSlice_Pair() {
// 	slice := NewMapSliceV("1", "2")
// 	first, second := slice.Pair()
// 	fmt.Println(first, second)
// 	// Output: 1 2
// }

// func TestMapSlice_Pair(t *testing.T) {

// 	// nil
// 	{
// 		first, second := (*MapSlice)(nil).Pair()
// 		assert.Equal(t, Obj(nil), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}

// 	// two values
// 	{
// 		first, second := NewMapSliceV("1", "2").Pair()
// 		assert.Equal(t, Obj("1"), first)
// 		assert.Equal(t, Obj("2"), second)
// 	}

// 	// one value
// 	{
// 		first, second := NewMapSliceV("1").Pair()
// 		assert.Equal(t, Obj("1"), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}

// 	// no values
// 	{
// 		first, second := NewMapSliceV().Pair()
// 		assert.Equal(t, Obj(nil), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}
// }

// // Pop
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Pop_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkMapSlice_Pop_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.Pop()
// 	}
// }

// func ExampleMapSlice_Pop() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Pop())
// 	// Output: 3
// }

// func TestMapSlice_Pop(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, Obj(nil), slice.Pop())
// 	}

// 	// take all one at a time
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, Obj("3"), slice.Pop())
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 		assert.Equal(t, Obj("2"), slice.Pop())
// 		assert.Equal(t, NewMapSliceV("1"), slice)
// 		assert.Equal(t, Obj("1"), slice.Pop())
// 		assert.Equal(t, NewMapSliceV(), slice)
// 		assert.Equal(t, Obj(nil), slice.Pop())
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}
// }

// // PopN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_PopN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkMapSlice_PopN_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.PopN(10)
// 	}
// }

// func ExampleMapSlice_PopN() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.PopN(2))
// 	// Output: [2 3]
// }

// func TestMapSlice_PopN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.PopN(1))
// 	}

// 	// take none
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV(), slice.PopN(0))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 	}

// 	// take 1
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("3"), slice.PopN(1))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 	}

// 	// take 2
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("2", "3"), slice.PopN(2))
// 		assert.Equal(t, NewMapSliceV("1"), slice)
// 	}

// 	// take 3
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice.PopN(3))
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}

// 	// take beyond
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice.PopN(4))
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}
// }

// // Prepend
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Prepend_Go(t *testing.B) {
// 	src := []string{}
// 	for _, x := range RangeString(nines6) {
// 		src = append(src, x)
// 		copy(src[1:], src[1:])
// 		src[0] = x
// 	}
// }

// func BenchmarkMapSlice_Prepend_Slice(t *testing.B) {
// 	slice := NewMapSliceV()
// 	for _, x := range RangeString(nines6) {
// 		slice.Prepend(x)
// 	}
// }

// func ExampleMapSlice_Prepend() {
// 	slice := NewMapSliceV("2", "3")
// 	fmt.Println(slice.Prepend("1"))
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Prepend(t *testing.T) {

// 	// happy path
// 	{
// 		slice := NewMapSliceV()
// 		assert.Equal(t, NewMapSliceV("2"), slice.Prepend("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Prepend("1"))
// 		assert.Equal(t, NewMapSliceV("0", "1", "2"), slice.Prepend("0"))
// 	}

// 	// error cases
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV("0"), slice.Prepend("0"))
// 	}
// }

// // Reverse
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Reverse_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
// 		src[i], src[j] = src[j], src[i]
// 	}
// }

// func BenchmarkMapSlice_Reverse_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.Reverse()
// }

// func ExampleMapSlice_Reverse() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Reverse())
// 	// Output: [3 2 1]
// }

// func TestMapSlice_Reverse(t *testing.T) {

// 	// nil
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.Reverse())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Reverse())
// 	}

// 	// pos
// 	{
// 		slice := NewMapSliceV("3", "2", "1")
// 		reversed := slice.Reverse()
// 		assert.Equal(t, NewMapSliceV("3", "2", "1", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), reversed)
// 	}

// 	// neg
// 	{
// 		slice := NewMapSliceV("2", "3", "-2", "-3")
// 		reversed := slice.Reverse()
// 		assert.Equal(t, NewMapSliceV("2", "3", "-2", "-3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("-3", "-2", "3", "2"), reversed)
// 	}
// }

// // ReverseM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_ReverseM_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
// 		src[i], src[j] = src[j], src[i]
// 	}
// }

// func BenchmarkMapSlice_ReverseM_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.ReverseM()
// }

// func ExampleMapSlice_ReverseM() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.ReverseM())
// 	// Output: [3 2 1]
// }

// func TestMapSlice_ReverseM(t *testing.T) {

// 	// nil
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.ReverseM())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().ReverseM())
// 	}

// 	// pos
// 	{
// 		slice := NewMapSliceV("3", "2", "1")
// 		reversed := slice.ReverseM()
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), reversed)
// 	}

// 	// neg
// 	{
// 		slice := NewMapSliceV("2", "3", "-2", "-3")
// 		reversed := slice.ReverseM()
// 		assert.Equal(t, NewMapSliceV("-3", "-2", "3", "2", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("-3", "-2", "3", "2", "4"), reversed)
// 	}
// }

// // Select
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Select_Go(t *testing.B) {
// 	even := []string{}
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if Obj(src[i]).ToInt()%2 == 0 {
// 			even = append(even, src[i])
// 		}
// 	}
// }

// func BenchmarkMapSlice_Select_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.Select(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	})
// }

// func ExampleMapSlice_Select() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Select(func(x O) bool {
// 		return ExB(x.(string) == "2" || x.(string) == "3")
// 	}))
// 	// Output: [2 3]
// }

// func TestMapSlice_Select(t *testing.T) {

// 	// Select all odd values
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.Select(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 != 0)
// 		})
// 		slice.DropFirst()
// 		assert.Equal(t, NewMapSliceV("2", "3", "4", "5", "6", "7", "8", "9"), slice)
// 		assert.Equal(t, NewMapSliceV("1", "3", "5", "7", "9"), new)
// 	}

// 	// Select all even values
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.Select(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 == 0)
// 		})
// 		slice.DropAt(1)
// 		assert.Equal(t, NewMapSliceV("1", "3", "4", "5", "6", "7", "8", "9"), slice)
// 		assert.Equal(t, NewMapSliceV("2", "4", "6", "8"), new)
// 	}
// }

// // Set
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Set_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		src[i] = "0"
// 	}
// }

// func BenchmarkMapSlice_Set_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.Set(i, "0")
// 	}
// }

// func ExampleMapSlice_Set() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Set(0, "0"))
// 	// Output: [0 2 3]
// }

// func TestMapSlice_Set(t *testing.T) {
// 	assert.Equal(t, NewMapSliceV("0", "2", "3"), NewMapSliceV("1", "2", "3").Set(0, "0"))
// 	assert.Equal(t, NewMapSliceV("1", "0", "3"), NewMapSliceV("1", "2", "3").Set(1, "0"))
// 	assert.Equal(t, NewMapSliceV("1", "2", "0"), NewMapSliceV("1", "2", "3").Set(2, "0"))
// 	assert.Equal(t, NewMapSliceV("0", "2", "3"), NewMapSliceV("1", "2", "3").Set(-3, "0"))
// 	assert.Equal(t, NewMapSliceV("1", "0", "3"), NewMapSliceV("1", "2", "3").Set(-2, "0"))
// 	assert.Equal(t, NewMapSliceV("1", "2", "0"), NewMapSliceV("1", "2", "3").Set(-1, "0"))

// 	// Test out of bounds
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Set(5, "1"))
// 	}

// 	// Test wrong type
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Set(5, "1"))
// 	}

// 	// Conversion
// 	{
// 		assert.Equal(t, NewMapSliceV(0, 2, 0), NewMapSliceV(0, 0, 0).Set(1, Object{2}))
// 		assert.Equal(t, NewMapSliceV(0, 2, 0), NewMapSliceV(0, 0, 0).Set(1, "2"))
// 		assert.Equal(t, NewMapSliceV(true, 0, 0), NewMapSliceV(0, 0, 0).Set(0, true))
// 		assert.Equal(t, NewMapSliceV(0, 0, 3), NewMapSliceV(0, 0, 0).Set(-1, Char('3')))
// 	}
// }

// // SetE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_SetE_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		src[i] = "0"
// 	}
// }

// func BenchmarkMapSlice_SetE_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.SetE(i, "0")
// 	}
// }

// func ExampleMapSlice_SetE() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.SetE(0, "0"))
// 	// Output: [0 2 3] <nil>
// }

// func TestMapSlice_SetE(t *testing.T) {

// 	// pos - begining
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		result, err := slice.SetE(0, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), slice)
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), result)

// 		// multiple
// 		result, err = slice.SetE(0, []string{"4", "5", "6"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), slice)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), result)

// 		// multiple over
// 		result, err = slice.SetE(0, []string{"4", "5", "6", "7"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), slice)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), result)
// 	}

// 	// pos - middle
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		result, err := slice.SetE(1, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("1", "0", "3"), slice)
// 		assert.Equal(t, NewMapSliceV("1", "0", "3"), result)
// 	}

// 	// pos - end
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		result, err := slice.SetE(2, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("1", "2", "0"), slice)
// 		assert.Equal(t, NewMapSliceV("1", "2", "0"), result)
// 	}

// 	// neg - begining
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		result, err := slice.SetE(-3, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), slice)
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), result)

// 		// multiple
// 		result, err = slice.SetE(-3, []string{"4", "5", "6"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), slice)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), result)

// 		// multiple over
// 		result, err = slice.SetE(-3, []string{"4", "5", "6", "7"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), slice)
// 		assert.Equal(t, NewMapSliceV("4", "5", "6"), result)
// 	}

// 	// neg - middle
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		result, err := slice.SetE(-2, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("1", "0", "3"), slice)
// 		assert.Equal(t, NewMapSliceV("1", "0", "3"), result)
// 	}

// 	// neg - end
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		result, err := slice.SetE(-1, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewMapSliceV("1", "2", "0"), slice)
// 		assert.Equal(t, NewMapSliceV("1", "2", "0"), result)
// 	}

// 	// Test out of bounds
// 	{
// 		slice, err := NewMapSliceV("1", "2", "3").SetE(5, "1")
// 		assert.NotNil(t, slice)
// 		assert.NotNil(t, err)
// 	}

// 	// Test wrong type
// 	{
// 		slice, err := NewMapSliceV("1", "2", "3").SetE(5, "1")
// 		assert.NotNil(t, slice)
// 		assert.NotNil(t, err)
// 	}
// }

// // Shift
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Shift_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkMapSlice_Shift_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.Shift()
// 	}
// }

// func ExampleMapSlice_Shift() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Shift())
// 	// Output: 1
// }

// func TestMapSlice_Shift(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, Obj(nil), slice.Shift())
// 	}

// 	// take all and beyond
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, Obj("1"), slice.Shift())
// 		assert.Equal(t, NewMapSliceV("2", "3"), slice)
// 		assert.Equal(t, Obj("2"), slice.Shift())
// 		assert.Equal(t, NewMapSliceV("3"), slice)
// 		assert.Equal(t, Obj("3"), slice.Shift())
// 		assert.Equal(t, NewMapSliceV(), slice)
// 		assert.Equal(t, Obj(nil), slice.Shift())
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}
// }

// // ShiftN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_ShiftN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkMapSlice_ShiftN_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.ShiftN(10)
// 	}
// }

// func ExampleMapSlice_ShiftN() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.ShiftN(2))
// 	// Output: [1 2]
// }

// func TestMapSlice_ShiftN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.ShiftN(1))
// 	}

// 	// negative value
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("1"), slice.ShiftN(-1))
// 		assert.Equal(t, NewMapSliceV("2", "3"), slice)
// 	}

// 	// take none
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV(), slice.ShiftN(0))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 	}

// 	// take 1
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("1"), slice.ShiftN(1))
// 		assert.Equal(t, NewMapSliceV("2", "3"), slice)
// 	}

// 	// take 2
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.ShiftN(2))
// 		assert.Equal(t, NewMapSliceV("3"), slice)
// 	}

// 	// take 3
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice.ShiftN(3))
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}

// 	// take beyond
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice.ShiftN(4))
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}
// }

// // Single
// //--------------------------------------------------------------------------------------------------

// func ExampleMapSlice_Single() {
// 	slice := NewMapSliceV("1")
// 	fmt.Println(slice.Single())
// 	// Output: true
// }

// func TestMapSlice_Single(t *testing.T) {

// 	assert.Equal(t, false, NewMapSliceV().Single())
// 	assert.Equal(t, true, NewMapSliceV("1").Single())
// 	assert.Equal(t, false, NewMapSliceV("1", "2").Single())
// }

// // Slice
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Slice_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	_ = src[0:len(src)]
// }

// func BenchmarkMapSlice_Slice_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	slice.Slice(0, -1)
// }

// func ExampleMapSlice_Slice() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Slice(1, -1))
// 	// Output: [2 3]
// }

// func TestMapSlice_Slice(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.Slice(0, -1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Slice(0, -1))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewMapSliceV("1", "2", "3")
// 		result := original.Slice(0, -1).Set(0, "0")
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), original)
// 		assert.Equal(t, NewMapSliceV("0", "2", "3"), result)
// 	}

// 	// slice full array
// 	{
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Slice(0, -1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Slice(0, 1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV().Slice(0, 5))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Slice(0, -1))
// 		assert.Equal(t, NewMapSlice([]string{"1", "2", "3"}), NewMapSlice([]string{"1", "2", "3"}).Slice(0, -1))
// 	}

// 	// out of bounds should be moved in
// 	{
// 		assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1").Slice(0, 2))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Slice(-6, 6))
// 	}

// 	// mutually exclusive
// 	{
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Slice(2, -3))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Slice(0, -5))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Slice(4, -1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Slice(6, -1))
// 		assert.Equal(t, NewMapSliceV(), NewMapSliceV("1", "2", "3", "4").Slice(3, 2))
// 	}

// 	// singles
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewMapSliceV("4"), slice.Slice(-1, -1))
// 		assert.Equal(t, NewMapSliceV("3"), slice.Slice(-2, -2))
// 		assert.Equal(t, NewMapSliceV("2"), slice.Slice(-3, -3))
// 		assert.Equal(t, NewMapSliceV("1"), slice.Slice(0, 0))
// 		assert.Equal(t, NewMapSliceV("1"), slice.Slice(-4, -4))
// 		assert.Equal(t, NewMapSliceV("2"), slice.Slice(1, 1))
// 		assert.Equal(t, NewMapSliceV("2"), slice.Slice(1, -3))
// 		assert.Equal(t, NewMapSliceV("3"), slice.Slice(2, 2))
// 		assert.Equal(t, NewMapSliceV("3"), slice.Slice(2, -2))
// 		assert.Equal(t, NewMapSliceV("4"), slice.Slice(3, 3))
// 		assert.Equal(t, NewMapSliceV("4"), slice.Slice(3, -1))
// 	}

// 	// grab all but first
// 	{
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Slice(1, -1))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Slice(1, 2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Slice(-2, -1))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Slice(-2, 2))
// 	}

// 	// grab all but last
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Slice(0, -2))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Slice(-3, -2))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Slice(-3, 1))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2", "3").Slice(0, 1))
// 	}

// 	// grab middle
// 	{
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Slice(1, -2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Slice(-3, -2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Slice(-3, 2))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3", "4").Slice(1, 2))
// 	}

// 	// random
// 	{
// 		assert.Equal(t, NewMapSliceV("1"), NewMapSliceV("1", "2", "3").Slice(0, -3))
// 		assert.Equal(t, NewMapSliceV("2", "3"), NewMapSliceV("1", "2", "3").Slice(1, 2))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), NewMapSliceV("1", "2", "3").Slice(0, 2))
// 	}
// }

// // Sort
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Sort_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.MapSlice(src))
// }

// func BenchmarkMapSlice_Sort_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.Sort()
// }

// func ExampleMapSlice_Sort() {
// 	slice := NewMapSliceV("2", "3", "1")
// 	fmt.Println(slice.Sort())
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Sort(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV().Sort())

// 	// pos
// 	{
// 		slice := NewMapSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.Sort()
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("5", "3", "2", "4", "1"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewMapSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.Sort()
// 		assert.Equal(t, NewMapSliceV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("5", "3", "-2", "4", "-1"), slice)
// 	}
// }

// // SortM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_SortM_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.MapSlice(src))
// }

// func BenchmarkMapSlice_SortM_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.SortM()
// }

// func ExampleMapSlice_SortM() {
// 	slice := NewMapSliceV("2", "3", "1")
// 	fmt.Println(slice.SortM())
// 	// Output: [1 2 3]
// }

// func TestMapSlice_SortM(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV().SortM())

// 	// pos
// 	{
// 		slice := NewMapSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortM()
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4", "5", "6"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewMapSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortM()
// 		assert.Equal(t, NewMapSliceV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("-1", "-2", "3", "4", "5", "6"), slice)
// 	}
// }

// // SortReverse
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_SortReverse_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.Reverse(sort.MapSlice(src)))
// }

// func BenchmarkMapSlice_SortReverse_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.SortReverse()
// }

// func ExampleMapSlice_SortReverse() {
// 	slice := NewMapSliceV("2", "3", "1")
// 	fmt.Println(slice.SortReverse())
// 	// Output: [3 2 1]
// }

// func TestMapSlice_SortReverse(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV().SortReverse())

// 	// pos
// 	{
// 		slice := NewMapSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverse()
// 		assert.Equal(t, NewMapSliceV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("5", "3", "2", "4", "1"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewMapSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverse()
// 		assert.Equal(t, NewMapSliceV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("5", "3", "-2", "4", "-1"), slice)
// 	}
// }

// // SortReverseM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_SortReverseM_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.Reverse(sort.MapSlice(src)))
// }

// func BenchmarkMapSlice_SortReverseM_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	slice.SortReverseM()
// }

// func ExampleMapSlice_SortReverseM() {
// 	slice := NewMapSliceV("2", "3", "1")
// 	fmt.Println(slice.SortReverseM())
// 	// Output: [3 2 1]
// }

// func TestMapSlice_SortReverseM(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewMapSliceV(), NewMapSliceV().SortReverse())

// 	// pos
// 	{
// 		slice := NewMapSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, NewMapSliceV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("5", "4", "3", "2", "1", "6"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewMapSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, NewMapSliceV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewMapSliceV("5", "4", "3", "-2", "-1", "6"), slice)
// 	}
// }

// // String
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_String_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	_ = fmt.Sprintf("%v", src)
// }

// func BenchmarkMapSlice_String_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	_ = slice.String()
// }

// func ExampleMapSlice_String() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestMapSlice_String(t *testing.T) {
// 	// nil
// 	assert.Equal(t, "[]", (*MapSlice)(nil).String())

// 	// empty
// 	assert.Equal(t, "[]", NewMapSliceV().String())

// 	// pos
// 	{
// 		slice := NewMapSliceV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, "[5 4 3 2 1 6]", sorted.Append("6").String())
// 		assert.Equal(t, "[5 4 3 2 1 6]", slice.String())
// 	}

// 	// neg
// 	{
// 		slice := NewMapSliceV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, "[5 4 3 -2 -1 6]", sorted.Append("6").String())
// 		assert.Equal(t, "[5 4 3 -2 -1 6]", slice.String())
// 	}
// }

// // Swap
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Swap_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if i+1 < len(src) {
// 			src[i], src[i+1] = src[i+1], src[i]
// 		}
// 	}
// }

// func BenchmarkMapSlice_Swap_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Swap(i, i+1)
// 		}
// 	}
// }

// func ExampleMapSlice_Swap() {
// 	slice := NewMapSliceV("2", "3", "1")
// 	slice.Swap(0, 2)
// 	slice.Swap(1, 2)
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Swap(t *testing.T) {

// 	// invalid cases
// 	{
// 		var slice *MapSlice
// 		slice.Swap(0, 0)
// 		assert.Equal(t, (*MapSlice)(nil), slice)

// 		slice = NewMapSliceV()
// 		slice.Swap(0, 0)
// 		assert.Equal(t, NewMapSliceV(), slice)

// 		slice.Swap(1, 2)
// 		assert.Equal(t, NewMapSliceV(), slice)

// 		slice.Swap(-1, 2)
// 		assert.Equal(t, NewMapSliceV(), slice)

// 		slice.Swap(1, -2)
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}

// 	// normal
// 	{
// 		slice := NewMapSliceV("0", "1", "2")
// 		slice.Swap(0, 1)
// 		assert.Equal(t, NewMapSliceV("1", "0", "2"), slice)
// 	}
// }

// // Take
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Take_Go(t *testing.B) {
// 	src := RangeString(nines7)
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

// func BenchmarkMapSlice_Take_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines7))
// 	for slice.Len() > 1 {
// 		slice.Take(1, 10)
// 	}
// }

// func ExampleMapSlice_Take() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.Take(0, 1))
// 	// Output: [1 2]
// }

// func TestMapSlice_Take(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.Take(0, 1))
// 	}

// 	// invalid
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4")
// 		assert.Equal(t, NewMapSliceV(), slice.Take(1))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice)
// 		assert.Equal(t, NewMapSliceV(), slice.Take(4, 4))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice)
// 	}

// 	// take 1
// 	{
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1"), slice.Take(0, 0))
// 			assert.Equal(t, NewMapSliceV("2", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("2"), slice.Take(1, 1))
// 			assert.Equal(t, NewMapSliceV("1", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("3"), slice.Take(2, 2))
// 			assert.Equal(t, NewMapSliceV("1", "2", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("4"), slice.Take(3, 3))
// 			assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("4"), slice.Take(-1, -1))
// 			assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("3"), slice.Take(-2, -2))
// 			assert.Equal(t, NewMapSliceV("1", "2", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("2"), slice.Take(-3, -3))
// 			assert.Equal(t, NewMapSliceV("1", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1"), slice.Take(-4, -4))
// 			assert.Equal(t, NewMapSliceV("2", "3", "4"), slice)
// 		}
// 	}

// 	// take 2
// 	{
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2"), slice.Take(0, 1))
// 			assert.Equal(t, NewMapSliceV("3", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("2", "3"), slice.Take(1, 2))
// 			assert.Equal(t, NewMapSliceV("1", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("3", "4"), slice.Take(2, 3))
// 			assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("3", "4"), slice.Take(-2, -1))
// 			assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("2", "3"), slice.Take(-3, -2))
// 			assert.Equal(t, NewMapSliceV("1", "4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2"), slice.Take(-4, -3))
// 			assert.Equal(t, NewMapSliceV("3", "4"), slice)
// 		}
// 	}

// 	// take 3
// 	{
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2", "3"), slice.Take(0, 2))
// 			assert.Equal(t, NewMapSliceV("4"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("2", "3", "4"), slice.Take(-3, -1))
// 			assert.Equal(t, NewMapSliceV("1"), slice)
// 		}
// 	}

// 	// take everything and beyond
// 	{
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Take())
// 			assert.Equal(t, NewMapSliceV(), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Take(0, 3))
// 			assert.Equal(t, NewMapSliceV(), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Take(0, -1))
// 			assert.Equal(t, NewMapSliceV(), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Take(-4, -1))
// 			assert.Equal(t, NewMapSliceV(), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Take(-6, -1))
// 			assert.Equal(t, NewMapSliceV(), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Take(0, 10))
// 			assert.Equal(t, NewMapSliceV(), slice)
// 		}
// 	}

// 	// move index within bounds
// 	{
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("4"), slice.Take(3, 4))
// 			assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("1", "2", "3", "4")
// 			assert.Equal(t, NewMapSliceV("1"), slice.Take(-5, 0))
// 			assert.Equal(t, NewMapSliceV("2", "3", "4"), slice)
// 		}
// 	}
// }

// // TakeAt
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_TakeAt_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	index := RangeString(nines5)
// 	for i := range index {
// 		if i+1 < len(src) {
// 			src = append(src[:i], src[i+1:]...)
// 		} else if i >= 0 && i < len(src) {
// 			src = src[:i]
// 		}
// 	}
// }

// func BenchmarkMapSlice_TakeAt_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	index := RangeString(nines5)
// 	slice := NewMapSlice(src)
// 	for i := range index {
// 		slice.TakeAt(i)
// 	}
// }

// func ExampleMapSlice_TakeAt() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.TakeAt(1))
// 	// Output: 2
// }

// func TestMapSlice_TakeAt(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, Obj(nil), slice.TakeAt(0))
// 	}

// 	// all and more
// 	{
// 		slice := NewMapSliceV("0", "1", "2")
// 		assert.Equal(t, Obj("2"), slice.TakeAt(-1))
// 		assert.Equal(t, NewMapSliceV("0", "1"), slice)
// 		assert.Equal(t, Obj("1"), slice.TakeAt(-1))
// 		assert.Equal(t, NewMapSliceV("0"), slice)
// 		assert.Equal(t, Obj("0"), slice.TakeAt(-1))
// 		assert.Equal(t, NewMapSliceV(), slice)
// 		assert.Equal(t, Obj(nil), slice.TakeAt(-1))
// 		assert.Equal(t, NewMapSliceV(), slice)
// 	}

// 	// take invalid
// 	{
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj(nil), slice.TakeAt(3))
// 			assert.Equal(t, NewMapSliceV("0", "1", "2"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj(nil), slice.TakeAt(-4))
// 			assert.Equal(t, NewMapSliceV("0", "1", "2"), slice)
// 		}
// 	}

// 	// take last
// 	{
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("2"), slice.TakeAt(2))
// 			assert.Equal(t, NewMapSliceV("0", "1"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("2"), slice.TakeAt(-1))
// 			assert.Equal(t, NewMapSliceV("0", "1"), slice)
// 		}
// 	}

// 	// take middle
// 	{
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("1"), slice.TakeAt(1))
// 			assert.Equal(t, NewMapSliceV("0", "2"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("1"), slice.TakeAt(-2))
// 			assert.Equal(t, NewMapSliceV("0", "2"), slice)
// 		}
// 	}

// 	// take first
// 	{
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("0"), slice.TakeAt(0))
// 			assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewMapSliceV("0", "1", "2")
// 			assert.Equal(t, Obj("0"), slice.TakeAt(-3))
// 			assert.Equal(t, NewMapSliceV("1", "2"), slice)
// 		}
// 	}
// }

// // TakeW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_TakeW_Go(t *testing.B) {
// 	new := []string{}
// 	src := RangeString(nines5)
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

// func BenchmarkMapSlice_TakeW_Slice(t *testing.B) {
// 	slice := NewMapSlice(RangeString(nines5))
// 	slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
// }

// func ExampleMapSlice_TakeW() {
// 	slice := NewMapSliceV("1", "2", "3")
// 	fmt.Println(slice.TakeW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	}))
// 	// Output: [2]
// }

// func TestMapSlice_TakeW(t *testing.T) {

// 	// take all odd values
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 != 0) })
// 		assert.Equal(t, NewMapSliceV("2", "4", "6", "8"), slice)
// 		assert.Equal(t, NewMapSliceV("1", "3", "5", "7", "9"), new)
// 	}

// 	// take all even values
// 	{
// 		slice := NewMapSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
// 		assert.Equal(t, NewMapSliceV("1", "3", "5", "7", "9"), slice)
// 		assert.Equal(t, NewMapSliceV("2", "4", "6", "8"), new)
// 	}
// }

// // Union
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Union_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkMapSlice_Union_Slice(t *testing.B) {
// 	// slice := NewMapSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleMapSlice_Union() {
// 	slice := NewMapSliceV("1", "2")
// 	fmt.Println(slice.Union([]string{"2", "3"}))
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Union(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Union(NewMapSliceV("1", "2")))
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Union([]string{"1", "2"}))
// 	}

// 	// size of one
// 	{
// 		slice := NewMapSliceV("1")
// 		union := slice.Union([]string{"1", "2", "3"})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewMapSliceV("1"), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewMapSliceV("1", "1")
// 		union := slice.Union(NewMapSliceV("2", "3"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewMapSliceV("1", "1"), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "2", "3", "3")
// 		union := slice.Union([]string{"1", "2", "3"})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewMapSliceV("1", "2", "2", "3", "3"), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		union := slice.Union([]string{"4", "5"})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4", "5"), union)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").Union((*[]string)(nil)))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").Union((*MapSlice)(nil)))
// 	}
// }

// // UnionM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_UnionM_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkMapSlice_UnionM_Slice(t *testing.B) {
// 	// slice := NewMapSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleMapSlice_UnionM() {
// 	slice := NewMapSliceV("1", "2")
// 	fmt.Println(slice.UnionM([]string{"2", "3"}))
// 	// Output: [1 2 3]
// }

// func TestMapSlice_UnionM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.UnionM(NewMapSliceV("1", "2")))
// 		assert.Equal(t, (*MapSlice)(nil), slice)
// 	}

// 	// size of one
// 	{
// 		slice := NewMapSliceV("1")
// 		union := slice.UnionM([]string{"1", "2", "3"})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewMapSliceV("1", "1")
// 		union := slice.UnionM(NewMapSliceV("2", "3"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "2", "3", "3")
// 		union := slice.UnionM([]string{"1", "2", "3"})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), union)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		union := slice.UnionM([]string{"4", "5"})
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4", "5"), union)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4", "5"), slice)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").UnionM((*[]string)(nil)))
// 		assert.Equal(t, NewMapSliceV("1", "2"), NewMapSliceV("1", "2").UnionM((*MapSlice)(nil)))
// 	}
// }

// // Uniq
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_Uniq_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkMapSlice_Uniq_Slice(t *testing.B) {
// 	// slice := NewMapSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleMapSlice_Uniq() {
// 	slice := NewMapSliceV("1", "2", "3", "3")
// 	fmt.Println(slice.Uniq())
// 	// Output: [1 2 3]
// }

// func TestMapSlice_Uniq(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, NewMapSliceV(), slice.Uniq())
// 	}

// 	// size of one
// 	{
// 		slice := NewMapSliceV("1")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewMapSliceV("1"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1"), uniq)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewMapSliceV("1", "1")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewMapSliceV("1"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1"), uniq)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "2", "3", "3")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "2", "2", "3", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), uniq)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), uniq)
// 	}
// }

// // UniqM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkMapSlice_UniqM_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkMapSlice_UniqM_Slice(t *testing.B) {
// 	// slice := NewMapSlice(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleMapSlice_UniqM() {
// 	slice := NewMapSliceV("1", "2", "3", "3")
// 	fmt.Println(slice.UniqM())
// 	// Output: [1 2 3]
// }

// func TestMapSlice_UniqM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *MapSlice
// 		assert.Equal(t, (*MapSlice)(nil), slice.UniqM())
// 	}

// 	// size of one
// 	{
// 		slice := NewMapSliceV("1")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewMapSliceV("1"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2"), uniq)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewMapSliceV("1", "1")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewMapSliceV("1"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewMapSliceV("1", "2"), uniq)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "2", "3", "3")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), uniq)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewMapSliceV("1", "2", "3")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewMapSliceV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewMapSliceV("1", "2", "3", "4"), uniq)
// 	}
// }

// func RangeString(size int) (new []string) {
// 	for _, x := range Range(0, size) {
// 		new = append(new, string(x))
// 	}
// 	return
// }
