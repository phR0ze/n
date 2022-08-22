package n

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Methods on both pointer and non-pointer
// --------------------------------------------------------------------------------------------------
func TestStringSlice_Methods(t *testing.T) {
	pointer := NewStringSliceV("0")
	assert.Equal(t, []string{"0"}, pointer.O())

	slice := *pointer
	assert.Equal(t, []string{}, slice.DropAt(0).O())
	assert.Equal(t, []string{}, slice.O())
}

// NewStringSlice
// --------------------------------------------------------------------------------------------------
func BenchmarkNewStringSlice_Go(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i += 10 {
		_ = []string{src[i], src[i] + fmt.Sprint(1), src[i] + fmt.Sprint(2), src[i] + fmt.Sprint(3), src[i] + fmt.Sprint(4), src[i] + fmt.Sprint(5), src[i] + fmt.Sprint(6), src[i] + fmt.Sprint(7), src[i] + fmt.Sprint(8), src[i] + fmt.Sprint(9)}
	}
}

func BenchmarkNewStringSlice_Slice(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i += 10 {
		_ = NewStringSlice([]string{src[i], src[i] + fmt.Sprint(1), src[i] + fmt.Sprint(2), src[i] + fmt.Sprint(3), src[i] + fmt.Sprint(4), src[i] + fmt.Sprint(5), src[i] + fmt.Sprint(6), src[i] + fmt.Sprint(7), src[i] + fmt.Sprint(8), src[i] + fmt.Sprint(9)})
	}
}

func ExampleNewStringSlice() {
	slice := NewStringSlice([]string{"1", "2", "3"})
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_NewStringSlice(t *testing.T) {

	// array
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, []string{"1", "2"}, NewStringSlice(array).O())
	assert.Equal(t, []string{"1", "2"}, NewStringSlice(array[:]).O())

	// empty
	assert.Equal(t, []string{}, NewStringSlice([]string{}).O())

	// slice
	assert.Equal(t, []string{"0"}, NewStringSlice([]string{"0"}).O())
	assert.Equal(t, []string{"1", "2"}, NewStringSlice([]string{"1", "2"}).O())

	// Conversion
	{
		assert.Equal(t, []string{"1"}, NewStringSlice("1").O())
		assert.Equal(t, []string{"1", "2"}, NewStringSlice([]string{"1", "2"}).O())
		assert.Equal(t, []string{"1"}, NewStringSlice(Object{1}).O())
		assert.Equal(t, []string{"1", "2"}, NewStringSlice([]Object{{1}, {2}}).O())
		assert.Equal(t, []string{"true"}, NewStringSlice(true).O())
		assert.Equal(t, []string{"true", "false"}, NewStringSlice([]bool{true, false}).O())
	}
}

// NewStringSliceV
// --------------------------------------------------------------------------------------------------
func BenchmarkNewStringSliceV_Go(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i += 10 {
		_ = append([]string{}, src[i], src[i]+fmt.Sprint(1), src[i]+fmt.Sprint(2), src[i]+fmt.Sprint(3), src[i]+fmt.Sprint(4), src[i]+fmt.Sprint(5), src[i]+fmt.Sprint(6), src[i]+fmt.Sprint(7), src[i]+fmt.Sprint(8), src[i]+fmt.Sprint(9))
	}
}

func BenchmarkNewStringSliceV_Slice(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i += 10 {
		_ = NewStringSliceV(src[i], src[i]+fmt.Sprint(1), src[i]+fmt.Sprint(2), src[i]+fmt.Sprint(3), src[i]+fmt.Sprint(4), src[i]+fmt.Sprint(5), src[i]+fmt.Sprint(6), src[i]+fmt.Sprint(7), src[i]+fmt.Sprint(8), src[i]+fmt.Sprint(9))
	}
}

func ExampleNewStringSliceV_empty() {
	slice := NewStringSliceV()
	fmt.Println(slice)
	// Output: []
}

func ExampleNewStringSliceV_variadic() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_NewStringSliceV(t *testing.T) {

	// empty
	assert.Equal(t, []string{}, NewStringSliceV().O())

	// multiples
	assert.Equal(t, []string{"1"}, NewStringSliceV("1").O())
	assert.Equal(t, []string{"1", "2"}, NewStringSliceV("1", "2").O())
	assert.Equal(t, []string{"1", "2"}, NewStringSliceV([]interface{}{"1", "2"}...).O())

	// Conversion
	{
		assert.Equal(t, []string{"1", "2"}, NewStringSliceV("1", "2").O())
		assert.Equal(t, []string{"1", "2"}, NewStringSliceV(Obj(1), Obj(2)).O())
		assert.Equal(t, []string{"true", "false"}, NewStringSliceV(true, false).O())
	}
}

// All
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_All_empty() {
	slice := NewStringSliceV()
	fmt.Println(slice.All())
	// Output: false
}

func ExampleStringSlice_All_notEmpty() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.All())
	// Output: true
}

func ExampleStringSlice_All_contains() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.All("1"))
	// Output: true
}

func ExampleStringSlice_All_containsAll() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.All("3", "1"))
	// Output: true
}

func TestStringSlice_All(t *testing.T) {

	// empty
	var nilSlice *StringSlice
	assert.False(t, nilSlice.All())
	assert.False(t, NewStringSliceV().All())

	// single
	assert.True(t, NewStringSliceV("2").All())

	// invalid
	assert.False(t, NewStringSliceV("1", "2").All(TestObj{"2"}))

	assert.True(t, NewStringSliceV("1", "2", "3").All("2"))
	assert.False(t, NewStringSliceV("1", "2", "3").All(4))
	assert.True(t, NewStringSliceV("1", "2", "3").All(2, "3"))
	assert.False(t, NewStringSliceV("1", "2", "3").All(4, 5))

	// conversion
	assert.True(t, NewStringSliceV("1", "2").All(Object{2}))
	assert.True(t, NewStringSliceV("1", "2", "3").All(int8(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(int16(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(int32('2')))
	assert.False(t, NewStringSliceV("1", "2", "3").All(int32(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(int64(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(uint8('2')))
	assert.False(t, NewStringSliceV("1", "2", "3").All(uint8(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(uint16(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(uint32(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(uint64(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").All(uint64(2)))
}

// AllS
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_AllS() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.AllS([]string{"2", "1"}))
	// Output: true
}

func TestStringSlice_AllS(t *testing.T) {
	// nil
	{
		var slice *StringSlice
		assert.False(t, slice.AllS([]string{"1"}))
		assert.True(t, NewStringSliceV("1").AllS(nil))
	}

	// []string
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AllS([]string{"1"}))
		assert.True(t, NewStringSliceV("1", "2", "3").AllS([]string{"2", "3"}))
		assert.False(t, NewStringSliceV("1", "2", "3").AllS([]string{"2", "5"}))
	}

	// *[]string
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(&([]string{"1"})))
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(&([]string{"2", "3"})))
		assert.False(t, NewStringSliceV("1", "2", "3").AllS(&([]string{"2", "5"})))
	}

	// Slice
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(ISlice(NewStringSliceV("1"))))
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(ISlice(NewStringSliceV("2", "3"))))
		assert.False(t, NewStringSliceV("1", "2", "3").AllS(ISlice(NewStringSliceV("2", "5"))))
	}

	// StringSlice
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(*NewStringSliceV("1")))
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(*NewStringSliceV("2", "3")))
		assert.False(t, NewStringSliceV("1", "2", "3").AllS(*NewStringSliceV("2", "5")))
	}

	// *StringSlice
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(NewStringSliceV("1")))
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(NewStringSliceV("2", "3")))
		assert.False(t, NewStringSliceV("1", "2", "3").AllS(NewStringSliceV("2", "5")))
	}

	// nothing to find to found
	assert.True(t, NewStringSliceV("1", "2").AllS(nil))
	assert.True(t, NewStringSliceV("1", "2").AllS((*[]string)(nil)))
	assert.True(t, NewStringSliceV("1", "2").AllS((*StringSlice)(nil)))

	// Conversion
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(NewStringSliceV(int64(1))))
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(NewStringSliceV(2)))
		assert.False(t, NewStringSliceV("1", "2", "3").AllS(NewStringSliceV(true)))
		assert.True(t, NewStringSliceV("1", "2", "3").AllS(NewStringSliceV(Char('3'))))
	}
}

// Any
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Any_Go(t *testing.B) {
	any := func(list []string, x []string) bool {
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
	src := RangeString(nines4)
	for _, x := range src {
		any(src, []string{x})
	}
}

func BenchmarkStringSlice_Any_Slice(t *testing.B) {
	src := RangeString(nines4)
	slice := NewStringSlice(src)
	for i := range src {
		slice.Any(i)
	}
}

func ExampleStringSlice_Any_empty() {
	slice := NewStringSliceV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleStringSlice_Any_notEmpty() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleStringSlice_Any_contains() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Any("1"))
	// Output: true
}

func ExampleStringSlice_Any_containsAny() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Any("0", "1"))
	// Output: true
}

func TestStringSlice_Any(t *testing.T) {

	// empty
	var nilSlice *StringSlice
	assert.False(t, nilSlice.Any())
	assert.False(t, NewStringSliceV().Any())

	// single
	assert.True(t, NewStringSliceV("2").Any())

	// invalid
	assert.False(t, NewStringSliceV("1", "2").Any(TestObj{"2"}))

	assert.True(t, NewStringSliceV("1", "2", "3").Any("2"))
	assert.False(t, NewStringSliceV("1", "2", "3").Any(4))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(4, "3"))
	assert.False(t, NewStringSliceV("1", "2", "3").Any(4, 5))

	// conversion
	assert.True(t, NewStringSliceV("1", "2").Any(Object{2}))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(int8(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(int16(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(int32('2')))
	assert.False(t, NewStringSliceV("1", "2", "3").Any(int32(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(int64(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(uint8('2')))
	assert.False(t, NewStringSliceV("1", "2", "3").Any(uint8(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(uint16(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(uint32(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(uint64(2)))
	assert.True(t, NewStringSliceV("1", "2", "3").Any(uint64(2)))
}

// AnyS
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_AnyS_Go(t *testing.B) {
	any := func(list []string, x []string) bool {
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
	src := RangeString(nines4)
	for _, x := range src {
		any(src, []string{x})
	}
}

func BenchmarkStringSlice_AnyS_Slice(t *testing.B) {
	src := RangeString(nines4)
	slice := NewStringSlice(src)
	for _, x := range src {
		slice.Any([]string{x})
	}
}

func ExampleStringSlice_AnyS() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.AnyS([]string{"0", "1"}))
	// Output: true
}

func TestStringSlice_AnyS(t *testing.T) {
	// nil
	{
		var slice *StringSlice
		assert.False(t, slice.AnyS([]string{"1"}))
		assert.False(t, NewStringSliceV("1").AnyS(nil))
	}

	// []string
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS([]string{"1"}))
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS([]string{"4", "3"}))
		assert.False(t, NewStringSliceV("1", "2", "3").AnyS([]string{"4", "5"}))
	}

	// *[]string
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(&([]string{"1"})))
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(&([]string{"4", "3"})))
		assert.False(t, NewStringSliceV("1", "2", "3").AnyS(&([]string{"4", "5"})))
	}

	// Slice
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(ISlice(NewStringSliceV("1"))))
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(ISlice(NewStringSliceV("4", "3"))))
		assert.False(t, NewStringSliceV("1", "2", "3").AnyS(ISlice(NewStringSliceV("4", "5"))))
	}

	// StringSlice
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(*NewStringSliceV("1")))
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(*NewStringSliceV("4", "3")))
		assert.False(t, NewStringSliceV("1", "2", "3").AnyS(*NewStringSliceV("4", "5")))
	}

	// *StringSlice
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(NewStringSliceV("1")))
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(NewStringSliceV("4", "3")))
		assert.False(t, NewStringSliceV("1", "2", "3").AnyS(NewStringSliceV("4", "5")))
	}

	// invalid types
	assert.False(t, NewStringSliceV("1", "2").AnyS(nil))
	assert.False(t, NewStringSliceV("1", "2").AnyS((*[]string)(nil)))
	assert.False(t, NewStringSliceV("1", "2").AnyS((*StringSlice)(nil)))

	// Conversion
	{
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(NewStringSliceV(int64(1))))
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(NewStringSliceV(2)))
		assert.False(t, NewStringSliceV("1", "2", "3").AnyS(NewStringSliceV(true)))
		assert.True(t, NewStringSliceV("1", "2", "3").AnyS(NewStringSliceV(Char('3'))))
	}
}

// AnyW
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_AnyW_Go(t *testing.B) {
	src := RangeString(nines5)
	for _, x := range src {
		if x == fmt.Sprint(nines4) {
			break
		}
	}
}

func BenchmarkStringSlice_AnyW_Slice(t *testing.B) {
	src := RangeString(nines5)
	NewStringSlice(src).AnyW(func(x O) bool {
		return ExB(x.(string) == fmt.Sprint(nines4))
	})
}

func ExampleStringSlice_AnyW() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.AnyW(func(x O) bool {
		return ExB(x.(string) == "2")
	}))
	// Output: true
}

func TestStringSlice_AnyW(t *testing.T) {

	// empty
	var slice *StringSlice
	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(string) > "0") }))
	assert.False(t, NewStringSliceV().AnyW(func(x O) bool { return ExB(x.(string) > "0") }))

	// single
	assert.True(t, NewStringSliceV("2").AnyW(func(x O) bool { return ExB(x.(string) > "0") }))
	assert.True(t, NewStringSliceV("1", "2").AnyW(func(x O) bool { return ExB(x.(string) == "2") }))
	assert.True(t, NewStringSliceV("1", "2", "3").AnyW(func(x O) bool { return ExB(x.(string) == "4" || x.(string) == "3") }))
}

// Append
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Append_Go(t *testing.B) {
	src := []string{}
	for _, i := range RangeString(nines6) {
		src = append(src, i)
	}
}

func BenchmarkStringSlice_Append_Slice(t *testing.B) {
	slice := NewStringSliceV()
	for _, i := range RangeString(nines6) {
		slice.Append(i)
	}
}

func ExampleStringSlice_Append() {
	slice := NewStringSliceV("1").Append("2").Append("3")
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_Append(t *testing.T) {

	// nil
	{
		var nilSlice *StringSlice
		assert.Equal(t, NewStringSliceV("0"), nilSlice.Append("0"))
		assert.Equal(t, (*StringSlice)(nil), nilSlice)
	}

	// Append one back to back
	{
		var slice *StringSlice
		assert.Equal(t, true, slice.Nil())
		slice = NewStringSliceV()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, false, slice.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		slice.Append("1")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{"1"}, slice.O())

		// Second append another which will be 2x at most
		slice.Append("2")
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []string{"1", "2"}, slice.O())
		assert.Equal(t, NewStringSliceV("1", "2"), slice)
	}

	// Start with just appending without chaining
	{
		slice := NewStringSliceV()
		assert.Equal(t, 0, slice.Len())
		slice.Append("1")
		assert.Equal(t, []string{"1"}, slice.O())
		slice.Append("2")
		assert.Equal(t, []string{"1", "2"}, slice.O())
	}

	// Start with nil not chained
	{
		slice := NewStringSliceV()
		assert.Equal(t, 0, slice.Len())
		slice.Append("1").Append("2").Append("3")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
	}

	// Start with nil chained
	{
		slice := NewStringSliceV().Append("1").Append("2")
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []string{"1", "2"}, slice.O())
	}

	// Start with non nil
	{
		slice := NewStringSliceV("1").Append("2").Append("3")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
	}

	// Use append result directly
	{
		slice := NewStringSliceV("1")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{"1", "2"}, slice.Append("2").O())
		assert.Equal(t, NewStringSliceV("1", "2"), slice)
	}

	// Conversion
	{
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV(1).Append(Object{2}))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV(1).Append("2"))
		assert.Equal(t, NewStringSliceV("true", "2"), NewStringSliceV().Append(true).Append(Char('2')))
	}
}

// AppendV
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_AppendV_Go(t *testing.B) {
	src := []string{}
	src = append(src, RangeString(nines6)...)
}

func BenchmarkStringSlice_AppendV_Slice(t *testing.B) {
	n := NewStringSliceV()
	new := rangeO(0, nines6)
	n.AppendV(new...)
}

func ExampleStringSlice_AppendV() {
	slice := NewStringSliceV("1").AppendV("2", "3")
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_AppendV(t *testing.T) {

	// nil
	{
		var nilSlice *StringSlice
		assert.Equal(t, NewStringSliceV("1", "2"), nilSlice.AppendV("1", "2"))
	}

	// Append many src
	{
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1").AppendV("2", "3"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4", "5"), NewStringSliceV("1").AppendV("2", "3").AppendV("4", "5"))
	}

	// Conversion
	{
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV().AppendV(Object{0}, Object{1}))
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV().AppendV("0", "1"))
		assert.Equal(t, NewStringSliceV("false", "true"), NewStringSliceV().AppendV(false, true))
	}
}

// At
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_At_Go(t *testing.B) {
	src := RangeString(nines6)
	for _, x := range src {
		assert.IsType(t, 0, x)
	}
}

func BenchmarkStringSlice_At_Slice(t *testing.B) {
	src := RangeString(nines6)
	slice := NewStringSlice(src)
	for i := 0; i < len(src); i++ {
		_, ok := (slice.At(i).O()).(string)
		assert.True(t, ok)
	}
}

func ExampleStringSlice_At() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.At(2))
	// Output: 3
}

func TestStringSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *StringSlice
		assert.Equal(t, Obj(nil), nilSlice.At(0))
	}

	// src
	{
		slice := NewStringSliceV("1", "2", "3", "4")
		assert.Equal(t, "4", slice.At(-1).O())
		assert.Equal(t, "3", slice.At(-2).O())
		assert.Equal(t, "2", slice.At(-3).O())
		assert.Equal(t, "1", slice.At(0).O())
		assert.Equal(t, "2", slice.At(1).O())
		assert.Equal(t, "3", slice.At(2).O())
		assert.Equal(t, "4", slice.At(3).O())
	}

	// index out of bounds
	{
		slice := NewStringSliceV("1")
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_Clear() {
	slice := NewStringSliceV("1").Concat([]string{"2", "3"})
	fmt.Println(slice.Clear())
	// Output: []
}

func TestStringSlice_Clear(t *testing.T) {

	// nil
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.Clear())
		assert.Equal(t, (*StringSlice)(nil), slice)
	}

	// int
	{
		slice := NewStringSliceV("1", "2", "3", "4")
		assert.Equal(t, NewStringSliceV(), slice.Clear())
		assert.Equal(t, NewStringSliceV(), slice.Clear())
		assert.Equal(t, NewStringSliceV(), slice)
	}
}

// Concat
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Concat_Go(t *testing.B) {
	dest := []string{}
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkStringSlice_Concat_Slice(t *testing.B) {
	dest := NewStringSliceV()
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.Concat(src[j:i])
		j = i
	}
}

func ExampleStringSlice_Concat() {
	slice := NewStringSliceV("1").Concat([]string{"2", "3"})
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_Concat(t *testing.T) {

	// nil
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Concat([]string{"1", "2"}))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").Concat(nil))
	}

	// []string
	{
		slice := NewStringSliceV("1")
		concated := slice.Concat([]string{"2", "3"})
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), concated)
	}

	// *[]string
	{
		slice := NewStringSliceV("1")
		concated := slice.Concat(&([]string{"2", "3"}))
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), concated)
	}

	// *StringSlice
	{
		slice := NewStringSliceV("1")
		concated := slice.Concat(NewStringSliceV("2", "3"))
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), concated)
	}

	// StringSlice
	{
		slice := NewStringSliceV("1")
		concated := slice.Concat(*NewStringSliceV("2", "3"))
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), concated)
	}

	// Slice
	{
		slice := NewStringSliceV("1")
		concated := slice.Concat(ISlice(NewStringSliceV("2", "3")))
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), concated)
	}

	// nils
	{
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").Concat((*[]string)(nil)))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").Concat((*StringSlice)(nil)))
	}

	// Conversion
	{
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV().Concat([]Object{{0}, {1}}))
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV().Concat([]string{"0", "1"}))
		assert.Equal(t, NewStringSliceV("false", "true"), NewStringSliceV().Concat([]bool{false, true}))

		slice := NewStringSliceV(Object{1})
		concated := slice.Concat([]int64{2, 3})
		assert.Equal(t, NewStringSliceV("1", "4"), slice.Append(Char('4')))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), concated)
	}
}

// ConcatM
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_ConcatM_Go(t *testing.B) {
	dest := []string{}
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkStringSlice_ConcatM_Slice(t *testing.B) {
	dest := NewStringSliceV()
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.ConcatM(src[j:i])
		j = i
	}
}

func ExampleStringSlice_ConcatM() {
	slice := NewStringSliceV("1").ConcatM([]string{"2", "3"})
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_ConcatM(t *testing.T) {

	// nil
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV("1", "2"), slice.ConcatM([]string{"1", "2"}))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").ConcatM(nil))
	}

	// []string
	{
		slice := NewStringSliceV("1")
		concated := slice.ConcatM([]string{"2", "3"})
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), concated)
	}

	// *[]string
	{
		slice := NewStringSliceV("1")
		concated := slice.ConcatM(&([]string{"2", "3"}))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), concated)
	}

	// *StringSlice
	{
		slice := NewStringSliceV("1")
		concated := slice.ConcatM(NewStringSliceV("2", "3"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), concated)
	}

	// StringSlice
	{
		slice := NewStringSliceV("1")
		concated := slice.ConcatM(*NewStringSliceV("2", "3"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), concated)
	}

	// Slice
	{
		slice := NewStringSliceV("1")
		concated := slice.ConcatM(ISlice(NewStringSliceV("2", "3")))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), concated)
	}

	// nils
	{
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").ConcatM((*[]string)(nil)))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").ConcatM((*StringSlice)(nil)))
	}

	// Conversion
	{
		slice := NewStringSliceV(Object{1})
		concated := slice.ConcatM([]Object{{2}, {3}})
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append(Char('4')))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), concated)
	}
}

// Copy
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Copy_Go(t *testing.B) {
	src := RangeString(nines6)
	dst := make([]string, len(src), len(src))
	copy(dst, src)
}

func BenchmarkStringSlice_Copy_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.Copy()
}

func ExampleStringSlice_Copy() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Copy())
	// Output: [1 2 3]
}

func TestStringSlice_Copy(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.Copy(0, -1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV("0").Clear().Copy(0, -1))
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := NewStringSliceV("1", "2", "3")
		result := original.Copy(0, -1)
		assert.Equal(t, NewStringSliceV("1", "2", "3"), original)
		assert.Equal(t, NewStringSliceV("1", "2", "3"), result)
		result.Set(0, "0")
		assert.Equal(t, NewStringSliceV("1", "2", "3"), original)
		assert.Equal(t, NewStringSliceV("0", "2", "3"), result)
	}

	// copy full array
	{
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Copy())
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Copy(0, -1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Copy(0, 1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Copy(0, 5))
		assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1").Copy())
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Copy())
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Copy(0, -1))
		assert.Equal(t, NewStringSlice([]string{"1", "2", "3"}), NewStringSlice([]string{"1", "2", "3"}).Copy())
		assert.Equal(t, NewStringSlice([]string{"1", "2", "3"}), NewStringSlice([]string{"1", "2", "3"}).Copy(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1").Copy(0, 2))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Copy(-6, 6))
	}

	// mutually exclusive
	{
		slice := NewStringSliceV("1", "2", "3", "4")
		assert.Equal(t, NewStringSliceV(), slice.Copy(2, -3))
		assert.Equal(t, NewStringSliceV(), slice.Copy(0, -5))
		assert.Equal(t, NewStringSliceV(), slice.Copy(4, -1))
		assert.Equal(t, NewStringSliceV(), slice.Copy(6, -1))
		assert.Equal(t, NewStringSliceV(), slice.Copy(3, -2))
	}

	// singles
	{
		slice := NewStringSliceV("1", "2", "3", "4")
		assert.Equal(t, NewStringSliceV("4"), slice.Copy(-1, -1))
		assert.Equal(t, NewStringSliceV("3"), slice.Copy(-2, -2))
		assert.Equal(t, NewStringSliceV("2"), slice.Copy(-3, -3))
		assert.Equal(t, NewStringSliceV("1"), slice.Copy(0, 0))
		assert.Equal(t, NewStringSliceV("1"), slice.Copy(-4, -4))
		assert.Equal(t, NewStringSliceV("2"), slice.Copy(1, 1))
		assert.Equal(t, NewStringSliceV("2"), slice.Copy(1, -3))
		assert.Equal(t, NewStringSliceV("3"), slice.Copy(2, 2))
		assert.Equal(t, NewStringSliceV("3"), slice.Copy(2, -2))
		assert.Equal(t, NewStringSliceV("4"), slice.Copy(3, 3))
		assert.Equal(t, NewStringSliceV("4"), slice.Copy(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Copy(1, -1))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Copy(1, 2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Copy(-2, -1))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Copy(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Copy(0, -2))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Copy(-3, -2))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Copy(-3, 1))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Copy(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Copy(1, -2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Copy(-3, -2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Copy(-3, 2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Copy(1, 2))
	}

	// random
	{
		assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1", "2", "3").Copy(0, -3))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Copy(1, 2))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Copy(0, 2))
	}
}

// Count
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Count_Go(t *testing.B) {
	src := RangeString(nines5)
	for _, x := range src {
		if x == fmt.Sprint(nines4) {
			break
		}
	}
}

func BenchmarkStringSlice_Count_Slice(t *testing.B) {
	src := RangeString(nines5)
	NewStringSlice(src).Count(nines4)
}

func ExampleStringSlice_Count() {
	slice := NewStringSliceV("1", "2", "2")
	fmt.Println(slice.Count("2"))
	// Output: 2
}

func TestStringSlice_Count(t *testing.T) {

	// empty
	var slice *StringSlice
	assert.Equal(t, 0, slice.Count(0))
	assert.Equal(t, 0, NewStringSliceV().Count(0))

	assert.Equal(t, 1, NewStringSliceV("2", "3").Count("2"))
	assert.Equal(t, 2, NewStringSliceV("1", "2", "2").Count("2"))
	assert.Equal(t, 4, NewStringSliceV("4", "4", "3", "4", "4").Count("4"))
	assert.Equal(t, 3, NewStringSliceV("3", "2", "3", "3", "5").Count("3"))
	assert.Equal(t, 1, NewStringSliceV("1", "2", "3").Count("3"))
}

// CountW
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_CountW_Go(t *testing.B) {
	src := RangeString(nines5)
	for _, x := range src {
		if x == fmt.Sprint(nines4) {
			break
		}
	}
}

func BenchmarkStringSlice_CountW_Slice(t *testing.B) {
	src := RangeString(nines5)
	NewStringSlice(src).CountW(func(x O) bool {
		return ExB(x.(string) == fmt.Sprint(nines4))
	})
}

func ExampleStringSlice_CountW() {
	slice := NewStringSliceV("1", "2", "2")
	fmt.Println(slice.CountW(func(x O) bool {
		return ExB(x.(string) == "2")
	}))
	// Output: 2
}

func TestStringSlice_CountW(t *testing.T) {

	// empty
	var slice *StringSlice
	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(string) > "0") }))
	assert.Equal(t, 0, NewStringSliceV().CountW(func(x O) bool { return ExB(x.(string) > "0") }))

	assert.Equal(t, 1, NewStringSliceV("2", "3").CountW(func(x O) bool { return ExB(x.(string) > "2") }))
	assert.Equal(t, 1, NewStringSliceV("1", "2").CountW(func(x O) bool { return ExB(x.(string) == "2") }))
	assert.Equal(t, 1, NewStringSliceV("1", "2", "3").CountW(func(x O) bool { return ExB(x.(string) == "4" || x.(string) == "3") }))
}

// Drop
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Drop_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 11 {
		i := 1
		n := 10
		if i+n < len(src) {
			src = append(src[:i], src[i+n:]...)
		} else {
			src = src[:i]
		}
	}
}

func BenchmarkStringSlice_Drop_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 1 {
		slice.Drop(1, 10)
	}
}

func ExampleStringSlice_Drop() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Drop(0, 1))
	// Output: [3]
}

func TestStringSlice_Drop(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.Drop(0, 1))
	}

	// invalid
	assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(1))
	assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(4, 4))

	// drop 1
	assert.Equal(t, NewStringSliceV("2", "3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(0, 0))
	assert.Equal(t, NewStringSliceV("1", "3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(1, 1))
	assert.Equal(t, NewStringSliceV("1", "2", "4"), NewStringSliceV("1", "2", "3", "4").Drop(2, 2))
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3", "4").Drop(3, 3))
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3", "4").Drop(-1, -1))
	assert.Equal(t, NewStringSliceV("1", "2", "4"), NewStringSliceV("1", "2", "3", "4").Drop(-2, -2))
	assert.Equal(t, NewStringSliceV("1", "3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(-3, -3))
	assert.Equal(t, NewStringSliceV("2", "3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(-4, -4))

	// drop 2
	assert.Equal(t, NewStringSliceV("3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(0, 1))
	assert.Equal(t, NewStringSliceV("1", "4"), NewStringSliceV("1", "2", "3", "4").Drop(1, 2))
	assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3", "4").Drop(2, 3))
	assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3", "4").Drop(-2, -1))
	assert.Equal(t, NewStringSliceV("1", "4"), NewStringSliceV("1", "2", "3", "4").Drop(-3, -2))
	assert.Equal(t, NewStringSliceV("3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(-4, -3))

	// drop 3
	assert.Equal(t, NewStringSliceV("4"), NewStringSliceV("1", "2", "3", "4").Drop(0, 2))
	assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1", "2", "3", "4").Drop(-3, -1))

	// drop everything and beyond
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Drop())
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Drop(0, 3))
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Drop(0, -1))
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Drop(-4, -1))
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Drop(-6, -1))
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Drop(0, 10))

	// move index within bounds
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3", "4").Drop(3, 4))
	assert.Equal(t, NewStringSliceV("2", "3", "4"), NewStringSliceV("1", "2", "3", "4").Drop(-5, 0))
}

// DropAt
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_DropAt_Go(t *testing.B) {
	src := RangeString(nines5)
	index := Range(0, nines5)
	for i := range index {
		if i+1 < len(src) {
			src = append(src[:i], src[i+1:]...)
		} else if i >= 0 && i < len(src) {
			src = src[:i]
		}
	}
}

func BenchmarkStringSlice_DropAt_Slice(t *testing.B) {
	index := Range(0, nines5)
	slice := NewStringSlice(RangeString(nines5))
	for i := range index {
		slice.DropAt(i)
	}
}

func ExampleStringSlice_DropAt() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.DropAt(1))
	// Output: [1 3]
}

func TestStringSlice_DropAt(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.DropAt(0))
	}

	// drop all and more
	{
		slice := NewStringSliceV("0", "1", "2")
		assert.Equal(t, NewStringSliceV("0", "1"), slice.DropAt(-1))
		assert.Equal(t, NewStringSliceV("0"), slice.DropAt(-1))
		assert.Equal(t, NewStringSliceV(), slice.DropAt(-1))
		assert.Equal(t, NewStringSliceV(), slice.DropAt(-1))
	}

	// drop invalid
	assert.Equal(t, NewStringSliceV("0", "1", "2"), NewStringSliceV("0", "1", "2").DropAt(3))
	assert.Equal(t, NewStringSliceV("0", "1", "2"), NewStringSliceV("0", "1", "2").DropAt(-4))

	// drop last
	assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV("0", "1", "2").DropAt(2))
	assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV("0", "1", "2").DropAt(-1))

	// drop middle
	assert.Equal(t, NewStringSliceV("0", "2"), NewStringSliceV("0", "1", "2").DropAt(1))
	assert.Equal(t, NewStringSliceV("0", "2"), NewStringSliceV("0", "1", "2").DropAt(-2))

	// drop first
	assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("0", "1", "2").DropAt(0))
	assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("0", "1", "2").DropAt(-3))
}

// DropFirst
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_DropFirst_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 1 {
		src = src[1:]
	}
}

func BenchmarkStringSlice_DropFirst_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleStringSlice_DropFirst() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.DropFirst())
	// Output: [2 3]
}

func TestStringSlice_DropFirst(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.DropFirst())
	}

	// drop all and beyond
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("2", "3"), slice.DropFirst())
		assert.Equal(t, NewStringSliceV("3"), slice.DropFirst())
		assert.Equal(t, NewStringSliceV(), slice.DropFirst())
		assert.Equal(t, NewStringSliceV(), slice.DropFirst())
	}
}

// DropFirstN
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_DropFirstN_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 10 {
		src = src[10:]
	}
}

func BenchmarkStringSlice_DropFirstN_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func ExampleStringSlice_DropFirstN() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.DropFirstN(2))
	// Output: [3]
}

func TestStringSlice_DropFirstN(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.DropFirstN(1))
	}

	// negative value
	assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").DropFirstN(-1))

	// drop none
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").DropFirstN(0))

	// drop 1
	assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").DropFirstN(1))

	// drop 2
	assert.Equal(t, NewStringSliceV("3"), NewStringSliceV("1", "2", "3").DropFirstN(2))

	// drop 3
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3").DropFirstN(3))

	// drop beyond
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3").DropFirstN(4))
}

func TestStringSlice_DropFirstW(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.DropFirstW(func(o O) bool { return true }))
	}

	// drop none
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").DropFirstW(func(o O) bool { return Obj(o).A() == "4" }))

	// drop 1
	assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").DropFirstW(func(o O) bool { return Obj(o).A() == "1" }))

	// drop 2
	assert.Equal(t, NewStringSliceV("3"), NewStringSliceV("1", "2", "3").DropFirstW(func(o O) bool { return Obj(o).A() == "1" || Obj(o).A() == "2" }))

	// drop 3
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3").DropFirstW(func(o O) bool { return Obj(o).A() == "1" || Obj(o).A() == "2" || Obj(o).A() == "3" }))
}

// DropLast
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_DropLast_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 1 {
		src = src[1:]
	}
}

func BenchmarkStringSlice_DropLast_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func ExampleStringSlice_DropLast() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.DropLast())
	// Output: [1 2]
}

func TestStringSlice_DropLast(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.DropLast())
	}

	// negative value
	assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").DropLastN(-1))

	slice := NewStringSliceV("1", "2", "3")
	assert.Equal(t, NewStringSliceV("1", "2"), slice.DropLast())
	assert.Equal(t, NewStringSliceV("1"), slice.DropLast())
	assert.Equal(t, NewStringSliceV(), slice.DropLast())
	assert.Equal(t, NewStringSliceV(), slice.DropLast())
}

// DropLastN
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_DropLastN_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 10 {
		src = src[10:]
	}
}

func BenchmarkStringSlice_DropLastN_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func ExampleStringSlice_DropLastN() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.DropLastN(2))
	// Output: [1]
}

func TestStringSlice_DropLastN(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.DropLastN(1))
	}

	// drop none
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").DropLastN(0))

	// drop 1
	assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").DropLastN(1))

	// drop 2
	assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1", "2", "3").DropLastN(2))

	// drop 3
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3").DropLastN(3))

	// drop beyond
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3").DropLastN(4))
}

// DropW
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_DropW_Go(t *testing.B) {
	src := RangeString(nines5)
	l := len(src)
	for i := 0; i < l; i++ {
		if Obj(src[i]).ToInt()%2 == 0 {
			if i+1 < l {
				src = append(src[:i], src[i+1:]...)
			} else if i >= 0 && i < l {
				src = src[:i]
			}
			l--
			i--
		}
	}
}

func BenchmarkStringSlice_DropW_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines5))
	slice.DropW(func(x O) bool {
		return ExB(Obj(x).ToInt()%2 == 0)
	})
}

func ExampleStringSlice_DropW() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.DropW(func(x O) bool {
		return ExB(Obj(x).ToInt()%2 == 0)
	}))
	// Output: [1 3]
}

func TestStringSlice_DropW(t *testing.T) {

	// drop all odd values
	{
		slice := NewStringSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		slice.DropW(func(x O) bool {
			return ExB(Obj(x).ToInt()%2 != 0)
		})
		assert.Equal(t, NewStringSliceV("2", "4", "6", "8"), slice)
	}

	// drop all even values
	{
		slice := NewStringSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		slice.DropW(func(x O) bool {
			return ExB(Obj(x).ToInt()%2 == 0)
		})
		assert.Equal(t, NewStringSliceV("1", "3", "5", "7", "9"), slice)
	}
}

// Each
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Each_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_Each_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).Each(func(x O) {
		assert.IsType(t, "", x)
	})
}

func ExampleStringSlice_Each() {
	NewStringSliceV("1", "2", "3").Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 123
}

func TestStringSlice_Each(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.Each(func(x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").Each(func(x O) {
			results = append(results, x.(string))
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}
}

// EachE
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_EachE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "0", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_EachE_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).EachE(func(x O) error {
		assert.IsType(t, "", x)
		return nil
	})
}

func ExampleStringSlice_EachE() {
	NewStringSliceV("1", "2", "3").EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 123
}

func TestStringSlice_EachE(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.EachE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachE(func(x O) error {
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachE(func(x O) error {
			if x.(string) == "3" {
				return Break
			}
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// EachI
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_EachI_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_EachI_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).EachI(func(i int, x O) {
		assert.IsType(t, "", x)
	})
}

func ExampleStringSlice_EachI() {
	NewStringSliceV("1", "2", "3").EachI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 0:11:22:3
}

func TestStringSlice_EachI(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.EachI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachI(func(i int, x O) {
			results = append(results, x.(string))
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}
}

// EachIE
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_EachIE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_EachIE_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).EachIE(func(i int, x O) error {
		assert.IsType(t, "", x)
		return nil
	})
}

func ExampleStringSlice_EachIE() {
	NewStringSliceV("1", "2", "3").EachIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 0:11:22:3
}

func TestStringSlice_EachIE(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.EachIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachIE(func(i int, x O) error {
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachIE(func(i int, x O) error {
			if i == 2 {
				return Break
			}
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// EachR
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_EachR_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_EachR_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).EachR(func(x O) {
		assert.IsType(t, "", x)
	})
}

func ExampleStringSlice_EachR() {
	NewStringSliceV("1", "2", "3").EachR(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 321
}

func TestStringSlice_EachR(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.EachR(func(x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachR(func(x O) {
			results = append(results, x.(string))
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}
}

// EachRE
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_EachRE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_EachRE_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).EachRE(func(x O) error {
		assert.IsType(t, "", x)
		return nil
	})
}

func ExampleStringSlice_EachRE() {
	NewStringSliceV("1", "2", "3").EachRE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 321
}

func TestStringSlice_EachRE(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.EachRE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachRE(func(x O) error {
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachRE(func(x O) error {
			if x.(string) == "1" {
				return Break
			}
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"3", "2"}, results)
	}
}

// EachRI
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_EachRI_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_EachRI_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).EachRI(func(i int, x O) {
		assert.IsType(t, "", x)
	})
}

func ExampleStringSlice_EachRI() {
	NewStringSliceV("1", "2", "3").EachRI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 2:31:20:1
}

func TestStringSlice_EachRI(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.EachRI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachRI(func(i int, x O) {
			results = append(results, x.(string))
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}
}

// EachRIE
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_EachRIE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, "", x)
	}
	for _, x := range RangeString(nines6) {
		action(x)
	}
}

func BenchmarkStringSlice_EachRIE_Slice(t *testing.B) {
	NewStringSlice(RangeString(nines6)).EachRIE(func(i int, x O) error {
		assert.IsType(t, "", x)
		return nil
	})
}

func ExampleStringSlice_EachRIE() {
	NewStringSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 2:31:20:1
}

func TestStringSlice_EachRIE(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		slice.EachRIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStringSliceV("1", "2", "3").EachRIE(func(i int, x O) error {
			if i == 0 {
				return Break
			}
			results = append(results, x.(string))
			return nil
		})
		assert.Equal(t, []string{"3", "2"}, results)
	}
}

// Empty
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_Empty() {
	fmt.Println(NewStringSliceV().Empty())
	// Output: true
}

func TestStringSlice_Empty(t *testing.T) {

	// nil or empty
	{
		var nilSlice *StringSlice
		assert.Equal(t, true, nilSlice.Empty())
	}

	assert.Equal(t, true, NewStringSliceV().Empty())
	assert.Equal(t, false, NewStringSliceV("1").Empty())
	assert.Equal(t, false, NewStringSliceV("1", "2", "3").Empty())
	assert.Equal(t, false, NewStringSliceV("1").Empty())
	assert.Equal(t, false, NewStringSlice([]string{"1", "2", "3"}).Empty())
}

// First
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_First_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 1 {
		_ = src[0]
		src = src[1:]
	}
}

func BenchmarkStringSlice_First_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.First()
		slice.DropFirst()
	}
}

func ExampleStringSlice_First() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.First())
	// Output: 1
}

func TestStringSlice_First(t *testing.T) {
	// invalid
	assert.Equal(t, Obj(nil), NewStringSliceV().First())

	// int
	assert.Equal(t, Obj("2"), NewStringSliceV("2", "3").First())
	assert.Equal(t, Obj("3"), NewStringSliceV("3", "2").First())
	assert.Equal(t, Obj("1"), NewStringSliceV("1", "3", "2").First())
}

// FirstN
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_FirstN_Go(t *testing.B) {
	src := RangeString(nines7)
	_ = src[0:10]
}

func BenchmarkStringSlice_FirstN_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	slice.FirstN(10)
}

func ExampleStringSlice_FirstN() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.FirstN(2))
	// Output: [1 2]
}

func TestStringSlice_FirstN(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.FirstN(1))
		assert.Equal(t, NewStringSliceV(), slice.FirstN(-1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewStringSliceV("1", "2", "3")
		result := original.FirstN(2).Set(0, "0")
		assert.Equal(t, NewStringSliceV("0", "2", "3"), original)
		assert.Equal(t, NewStringSliceV("0", "2"), result)
	}

	// Get none
	assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3").FirstN(0))

	// slice full array includeing out of bounds
	assert.Equal(t, NewStringSliceV(), NewStringSliceV().FirstN(1))
	assert.Equal(t, NewStringSliceV(), NewStringSliceV().FirstN(10))
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").FirstN(10))
	assert.Equal(t, NewStringSlice([]string{"1", "2", "3"}), NewStringSlice([]string{"1", "2", "3"}).FirstN(10))

	// grab a few diff
	assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1", "2", "3").FirstN(1))
	assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").FirstN(2))
}

// FirstW
// --------------------------------------------------------------------------------------------------
func TestStringSlice_FirstW(t *testing.T) {

	// empty
	assert.Equal(t, Obj(nil), NewStringSliceV().FirstW(func(o O) bool { return A(o).G() == "1" }))

	// item doesn't exist
	assert.Equal(t, Obj(nil), NewStringSliceV("2").FirstW(func(o O) bool { return A(o).G() == "1" }))

	// item exists
	assert.Equal(t, Obj("1"), NewStringSliceV("2", "1").FirstW(func(o O) bool { return A(o).G() == "1" }))
}

// G
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_G() {
	fmt.Println(NewStringSliceV("1", "2", "3").G())
	// Output: [1 2 3]
}

func TestStringSlice_G(t *testing.T) {
	assert.IsType(t, []string{}, NewStringSliceV().G())
	assert.IsType(t, []string{"1", "2", "3"}, NewStringSliceV("1", "2", "3").G())
}

// Generic
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_Generic() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.InterSlice())
	// Output: false
}

// Index
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Index_Go(t *testing.B) {
	for _, x := range RangeString(nines5) {
		if x == fmt.Sprint(nines4) {
			break
		}
	}
}

func BenchmarkStringSlice_Index_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines5))
	slice.Index(nines4)
}

func ExampleStringSlice_Index() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Index("2"))
	// Output: 1
}

func TestStringSlice_Index(t *testing.T) {

	// empty
	var slice *StringSlice
	assert.Equal(t, -1, slice.Index("2"))
	assert.Equal(t, -1, NewStringSliceV().Index("1"))

	assert.Equal(t, 0, NewStringSliceV("1", "2", "3").Index("1"))
	assert.Equal(t, 1, NewStringSliceV("1", "2", "3").Index("2"))
	assert.Equal(t, 2, NewStringSliceV("1", "2", "3").Index("3"))
	assert.Equal(t, -1, NewStringSliceV("1", "2", "3").Index("4"))
	assert.Equal(t, -1, NewStringSliceV("1", "2", "3").Index("5"))

	// Conversion
	{
		assert.Equal(t, 1, NewStringSliceV("1", "2", "3").Index(Object{2}))
		assert.Equal(t, 1, NewStringSliceV("1", "2", "3").Index("2"))
		assert.Equal(t, -1, NewStringSliceV("1", "2", "3").Index(true))
		assert.Equal(t, 2, NewStringSliceV("1", "2", "3").Index(Char('3')))
	}
}

// Insert
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Insert_Go(t *testing.B) {
	src := []string{}
	for _, x := range RangeString(nines6) {
		src = append(src, x)
		copy(src[1:], src[1:])
		src[0] = x
	}
}

func BenchmarkStringSlice_Insert_Slice(t *testing.B) {
	slice := NewStringSliceV()
	for x := range RangeString(nines6) {
		slice.Insert(0, x)
	}
}

func ExampleStringSlice_Insert() {
	slice := NewStringSliceV("1", "3")
	fmt.Println(slice.Insert(1, "2"))
	// Output: [1 2 3]
}

func TestStringSlice_Insert(t *testing.T) {

	// append
	{
		slice := NewStringSliceV()
		assert.Equal(t, NewStringSliceV("0"), slice.Insert(-1, "0"))
		assert.Equal(t, NewStringSliceV("0", "1"), slice.Insert(-1, "1"))
		assert.Equal(t, NewStringSliceV("0", "1", "2"), slice.Insert(-1, "2"))
	}

	// [] append
	{
		slice := NewStringSliceV()
		assert.Equal(t, NewStringSliceV("0"), slice.Insert(-1, []string{"0"}))
		assert.Equal(t, NewStringSliceV("0", "1", "2"), slice.Insert(-1, []string{"1", "2"}))
	}

	// prepend
	{
		slice := NewStringSliceV()
		assert.Equal(t, NewStringSliceV("2"), slice.Insert(0, "2"))
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Insert(0, "1"))
		assert.Equal(t, NewStringSliceV("0", "1", "2"), slice.Insert(0, "0"))
	}

	// [] prepend
	{
		slice := NewStringSliceV()
		assert.Equal(t, NewStringSliceV("2"), slice.Insert(0, []string{"2"}))
		assert.Equal(t, NewStringSliceV("0", "1", "2"), slice.Insert(0, []string{"0", "1"}))
	}

	// middle pos
	{
		slice := NewStringSliceV("0", "5")
		assert.Equal(t, NewStringSliceV("0", "1", "5"), slice.Insert(1, "1"))
		assert.Equal(t, NewStringSliceV("0", "1", "2", "5"), slice.Insert(2, "2"))
		assert.Equal(t, NewStringSliceV("0", "1", "2", "3", "5"), slice.Insert(3, "3"))
		assert.Equal(t, NewStringSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(4, "4"))
	}

	// [] middle pos
	{
		slice := NewStringSliceV("0", "5")
		assert.Equal(t, NewStringSliceV("0", "1", "2", "5"), slice.Insert(1, []string{"1", "2"}))
		assert.Equal(t, NewStringSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(3, []string{"3", "4"}))
	}

	// middle neg
	{
		slice := NewStringSliceV("0", "5")
		assert.Equal(t, NewStringSliceV("0", "1", "5"), slice.Insert(-2, "1"))
		assert.Equal(t, NewStringSliceV("0", "1", "2", "5"), slice.Insert(-2, "2"))
		assert.Equal(t, NewStringSliceV("0", "1", "2", "3", "5"), slice.Insert(-2, "3"))
		assert.Equal(t, NewStringSliceV("0", "1", "2", "3", "4", "5"), slice.Insert(-2, "4"))
	}

	// [] middle neg
	{
		slice := NewStringSliceV(0, 5)
		assert.Equal(t, NewStringSliceV(0, 1, 2, 5), slice.Insert(-2, []string{"1", "2"}))
		assert.Equal(t, NewStringSliceV(0, "1", "2", "3", 4, 5), slice.Insert(-2, []int{3, 4}))
	}

	// error cases
	{
		var slice *StringSlice
		assert.False(t, slice.Insert(0, 0).Nil())
		assert.Equal(t, NewStringSliceV("0"), slice.Insert(0, "0"))
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV("0", "1").Insert(-10, "1"))
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV("0", "1").Insert(10, "1"))
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV("0", "1").Insert(2, "1"))
		assert.Equal(t, NewStringSliceV("0", "1"), NewStringSliceV("0", "1").Insert(-3, "1"))
	}

	// [] error cases
	{
		var slice *StringSlice
		assert.False(t, slice.Insert(0, 0).Nil())
		assert.Equal(t, NewStringSliceV(0), slice.Insert(0, 0))
		assert.Equal(t, NewStringSliceV(0, 1), NewStringSliceV(0, 1).Insert(-10, 1))
		assert.Equal(t, NewStringSliceV(0, 1), NewStringSliceV(0, 1).Insert(10, 1))
		assert.Equal(t, NewStringSliceV(0, 1), NewStringSliceV(0, 1).Insert(2, 1))
		assert.Equal(t, NewStringSliceV(0, 1), NewStringSliceV(0, 1).Insert(-3, 1))
	}

	// Conversion
	{
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV(1, 3).Insert(1, Object{2}))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV(1, 3).Insert(1, "2"))
		assert.Equal(t, NewStringSliceV(true, "2", "3"), NewStringSliceV(2, 3).Insert(0, true))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV(1, 2).Insert(-1, Char('3')))
	}

	// [] Conversion
	{
		assert.Equal(t, NewStringSliceV("1", "2", "3", 4), NewStringSliceV(1, 4).Insert(1, []Object{{2}, {3}}))
		assert.Equal(t, NewStringSliceV("1", "2", "3", 4), NewStringSliceV(1, 4).Insert(1, []string{"2", "3"}))
		assert.Equal(t, NewStringSliceV(false, true, "2", "3"), NewStringSliceV(2, 3).Insert(0, []bool{false, true}))
		assert.Equal(t, NewStringSliceV("1", "2", "3", 4), NewStringSliceV(1, 2).Insert(-1, []Char{'3', '4'}))
	}
}

// Join
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Join_Go(t *testing.B) {
	src := RangeString(nines4)
	strings.Join(src, ",")
}

func BenchmarkStringSlice_Join_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines4))
	slice.Join()
}

func ExampleStringSlice_Join() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Join())
	// Output: 1,2,3
}

func TestStringSlice_Join(t *testing.T) {
	// nil
	{
		var slice *StringSlice
		assert.Equal(t, Obj(""), slice.Join())
	}

	// empty
	{
		assert.Equal(t, Obj(""), NewStringSliceV().Join())
	}

	assert.Equal(t, "1,2,3", NewStringSliceV("1", "2", "3").Join().O())
	assert.Equal(t, "1.2.3", NewStringSliceV("1", "2", "3").Join(".").O())
}

// Last
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Last_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 1 {
		_ = src[len(src)-1]
		src = src[:len(src)-1]
	}
}

func BenchmarkStringSlice_Last_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.Last()
		slice.DropLast()
	}
}

func ExampleStringSlice_Last() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Last())
	// Output: 3
}

func TestStringSlice_Last(t *testing.T) {
	// invalid
	assert.Equal(t, Obj(nil), NewStringSliceV().Last())

	// int
	assert.Equal(t, Obj("3"), NewStringSliceV("2", "3").Last())
	assert.Equal(t, Obj("2"), NewStringSliceV("3", "2").Last())
	assert.Equal(t, Obj("2"), NewStringSliceV("1", "3", "2").Last())
}

// LastN
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_LastN_Go(t *testing.B) {
	src := RangeString(nines7)
	_ = src[0:10]
}

func BenchmarkStringSlice_LastN_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	slice.LastN(10)
}

func ExampleStringSlice_LastN() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.LastN(2))
	// Output: [2 3]
}

func TestStringSlice_LastN(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.LastN(1))
		assert.Equal(t, NewStringSliceV(), slice.LastN(-1))
	}

	// Get none
	{
		assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3").LastN(0))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewStringSliceV("1", "2", "3")
		result := original.LastN(2).Set(0, "0")
		assert.Equal(t, NewStringSliceV("1", "0", "3"), original)
		assert.Equal(t, NewStringSliceV("0", "3"), result)
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().LastN(1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().LastN(10))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").LastN(10))
		assert.Equal(t, NewStringSlice([]string{"1", "2", "3"}), NewStringSlice([]string{"1", "2", "3"}).LastN(10))
	}

	// grab a few diff
	{
		assert.Equal(t, NewStringSliceV("3"), NewStringSliceV("1", "2", "3").LastN(1))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").LastN(2))
	}
}

// Len
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_Len() {
	fmt.Println(NewStringSliceV("1", "2", "3").Len())
	// Output: 3
}

func TestStringSlice_Len(t *testing.T) {
	assert.Equal(t, 0, NewStringSliceV().Len())
	assert.Equal(t, 2, len(*(NewStringSliceV("1", "2"))))
	assert.Equal(t, 2, NewStringSliceV("1", "2").Len())
}

// Less
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Less_Go(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i++ {
		if i+1 < len(src) {
			_ = src[i] < src[i+1]
		}
	}
}

func BenchmarkStringSlice_Less_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Less(i, i+1)
		}
	}
}

func ExampleStringSlice_Less() {
	slice := NewStringSliceV("2", "3", "1")
	fmt.Println(slice.Less(0, 2))
	// Output: false
}

func TestStringSlice_Less(t *testing.T) {

	// invalid cases
	{
		var slice *StringSlice
		assert.False(t, slice.Less(0, 0))

		slice = NewStringSliceV()
		assert.False(t, slice.Less(0, 0))
		assert.False(t, slice.Less(1, 2))
		assert.False(t, slice.Less(-1, 2))
		assert.False(t, slice.Less(1, -2))
	}

	// valid
	assert.Equal(t, true, NewStringSliceV("0", "1", "2").Less(0, 1))
	assert.Equal(t, false, NewStringSliceV("0", "1", "2").Less(1, 0))
	assert.Equal(t, true, NewStringSliceV("0", "1", "2").Less(1, 2))
}

// Map
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_Map() {
	slice := NewStringSliceV("1", "2", "3")
	slice = slice.Map(func(x O) O {
		return ToStr(ToInt(x.(string)) + 1).A()
	}).S()
	fmt.Println(slice.G())
	// Output: [2 3 4]
}

func TestStringSlice_Map(t *testing.T) {
	// int result
	{
		slice := SV("1", "2", "3")
		new := slice.Map(func(x O) O {
			return ToInt(x.(string)) + 1
		})
		assert.Equal(t, []int{2, 3, 4}, new.O())
	}

	// string
	{
		slice := SV("1", "2", "3")
		slice = slice.Map(func(x O) O {
			return ToStr(ToInt(x.(string)) + 1).A()
		}).S()
		assert.Equal(t, []string{"2", "3", "4"}, slice.G())
		assert.Equal(t, S([]string{"2", "3", "4"}), slice)
	}
}

// Nil
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_Nil() {
	var slice *StringSlice
	fmt.Println(slice.Nil())
	// Output: true
}

func TestStringSlice_Nil(t *testing.T) {
	var slice *StringSlice
	assert.True(t, slice.Nil())
	assert.False(t, NewStringSliceV().Nil())
	assert.False(t, NewStringSliceV("1", "2", "3").Nil())
}

// O
// --------------------------------------------------------------------------------------------------
func ExampleStringSlice_O() {
	fmt.Println(NewStringSliceV("1", "2", "3"))
	// Output: [1 2 3]
}

func TestStringSlice_O(t *testing.T) {
	assert.Equal(t, []string{}, (*StringSlice)(nil).O())
	assert.Equal(t, []string{}, NewStringSliceV().O())
	assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3"))
}

// Pair
//--------------------------------------------------------------------------------------------------

func ExampleStringSlice_Pair() {
	slice := NewStringSliceV("1", "2")
	first, second := slice.Pair()
	fmt.Println(first, second)
	// Output: 1 2
}

func TestStringSlice_Pair(t *testing.T) {

	// nil
	{
		first, second := (*StringSlice)(nil).Pair()
		assert.Equal(t, Obj(nil), first)
		assert.Equal(t, Obj(nil), second)
	}

	// two values
	{
		first, second := NewStringSliceV("1", "2").Pair()
		assert.Equal(t, Obj("1"), first)
		assert.Equal(t, Obj("2"), second)
	}

	// one value
	{
		first, second := NewStringSliceV("1").Pair()
		assert.Equal(t, Obj("1"), first)
		assert.Equal(t, Obj(nil), second)
	}

	// no values
	{
		first, second := NewStringSliceV().Pair()
		assert.Equal(t, Obj(nil), first)
		assert.Equal(t, Obj(nil), second)
	}
}

// Pop
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Pop_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 1 {
		src = src[1:]
	}
}

func BenchmarkStringSlice_Pop_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.Pop()
	}
}

func ExampleStringSlice_Pop() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Pop())
	// Output: 3
}

func TestStringSlice_Pop(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, Obj(nil), slice.Pop())
	}

	// take all one at a time
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, Obj("3"), slice.Pop())
		assert.Equal(t, NewStringSliceV("1", "2"), slice)
		assert.Equal(t, Obj("2"), slice.Pop())
		assert.Equal(t, NewStringSliceV("1"), slice)
		assert.Equal(t, Obj("1"), slice.Pop())
		assert.Equal(t, NewStringSliceV(), slice)
		assert.Equal(t, Obj(nil), slice.Pop())
		assert.Equal(t, NewStringSliceV(), slice)
	}
}

// PopN
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_PopN_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 10 {
		src = src[10:]
	}
}

func BenchmarkStringSlice_PopN_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.PopN(10)
	}
}

func ExampleStringSlice_PopN() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.PopN(2))
	// Output: [2 3]
}

func TestStringSlice_PopN(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.PopN(1))
	}

	// take none
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV(), slice.PopN(0))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
	}

	// take 1
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("3"), slice.PopN(1))
		assert.Equal(t, NewStringSliceV("1", "2"), slice)
	}

	// take 2
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("2", "3"), slice.PopN(2))
		assert.Equal(t, NewStringSliceV("1"), slice)
	}

	// take 3
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice.PopN(3))
		assert.Equal(t, NewStringSliceV(), slice)
	}

	// take beyond
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice.PopN(4))
		assert.Equal(t, NewStringSliceV(), slice)
	}
}

// Prepend
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Prepend_Go(t *testing.B) {
	src := []string{}
	for _, x := range RangeString(nines6) {
		src = append(src, x)
		copy(src[1:], src[1:])
		src[0] = x
	}
}

func BenchmarkStringSlice_Prepend_Slice(t *testing.B) {
	slice := NewStringSliceV()
	for _, x := range RangeString(nines6) {
		slice.Prepend(x)
	}
}

func ExampleStringSlice_Prepend() {
	slice := NewStringSliceV("2", "3")
	fmt.Println(slice.Prepend("1"))
	// Output: [1 2 3]
}

func TestStringSlice_Prepend(t *testing.T) {

	// happy path
	{
		slice := NewStringSliceV()
		assert.Equal(t, NewStringSliceV("2"), slice.Prepend("2"))
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Prepend("1"))
		assert.Equal(t, NewStringSliceV("0", "1", "2"), slice.Prepend("0"))
	}

	// error cases
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV("0"), slice.Prepend("0"))
	}
}

// Reverse
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Reverse_Go(t *testing.B) {
	src := RangeString(nines6)
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}

func BenchmarkStringSlice_Reverse_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.Reverse()
}

func ExampleStringSlice_Reverse() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Reverse())
	// Output: [3 2 1]
}

func TestStringSlice_Reverse(t *testing.T) {

	// nil
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.Reverse())
	}

	// empty
	{
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Reverse())
	}

	// pos
	{
		slice := NewStringSliceV("3", "2", "1")
		reversed := slice.Reverse()
		assert.Equal(t, NewStringSliceV("3", "2", "1", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), reversed)
	}

	// neg
	{
		slice := NewStringSliceV("2", "3", "-2", "-3")
		reversed := slice.Reverse()
		assert.Equal(t, NewStringSliceV("2", "3", "-2", "-3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("-3", "-2", "3", "2"), reversed)
	}
}

// ReverseM
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_ReverseM_Go(t *testing.B) {
	src := RangeString(nines6)
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}

func BenchmarkStringSlice_ReverseM_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.ReverseM()
}

func ExampleStringSlice_ReverseM() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.ReverseM())
	// Output: [3 2 1]
}

func TestStringSlice_ReverseM(t *testing.T) {

	// nil
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.ReverseM())
	}

	// empty
	{
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().ReverseM())
	}

	// pos
	{
		slice := NewStringSliceV("3", "2", "1")
		reversed := slice.ReverseM()
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), reversed)
	}

	// neg
	{
		slice := NewStringSliceV("2", "3", "-2", "-3")
		reversed := slice.ReverseM()
		assert.Equal(t, NewStringSliceV("-3", "-2", "3", "2", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("-3", "-2", "3", "2", "4"), reversed)
	}
}

// Select
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Select_Go(t *testing.B) {
	even := []string{}
	src := RangeString(nines6)
	for i := 0; i < len(src); i++ {
		if Obj(src[i]).ToInt()%2 == 0 {
			even = append(even, src[i])
		}
	}
}

func BenchmarkStringSlice_Select_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.Select(func(x O) bool {
		return ExB(Obj(x).ToInt()%2 == 0)
	})
}

func ExampleStringSlice_Select() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Select(func(x O) bool {
		return ExB(x.(string) == "2" || x.(string) == "3")
	}))
	// Output: [2 3]
}

func TestStringSlice_Select(t *testing.T) {

	// Select all odd values
	{
		slice := NewStringSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		new := slice.Select(func(x O) bool {
			return ExB(Obj(x).ToInt()%2 != 0)
		})
		slice.DropFirst()
		assert.Equal(t, NewStringSliceV("2", "3", "4", "5", "6", "7", "8", "9"), slice)
		assert.Equal(t, NewStringSliceV("1", "3", "5", "7", "9"), new)
	}

	// Select all even values
	{
		slice := NewStringSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		new := slice.Select(func(x O) bool {
			return ExB(Obj(x).ToInt()%2 == 0)
		})
		slice.DropAt(1)
		assert.Equal(t, NewStringSliceV("1", "3", "4", "5", "6", "7", "8", "9"), slice)
		assert.Equal(t, NewStringSliceV("2", "4", "6", "8"), new)
	}
}

// Set
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Set_Go(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i++ {
		src[i] = "0"
	}
}

func BenchmarkStringSlice_Set_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, "0")
	}
}

func ExampleStringSlice_Set() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Set(0, "0"))
	// Output: [0 2 3]
}

func TestStringSlice_Set(t *testing.T) {
	assert.Equal(t, NewStringSliceV("0", "2", "3"), NewStringSliceV("1", "2", "3").Set(0, "0"))
	assert.Equal(t, NewStringSliceV("1", "0", "3"), NewStringSliceV("1", "2", "3").Set(1, "0"))
	assert.Equal(t, NewStringSliceV("1", "2", "0"), NewStringSliceV("1", "2", "3").Set(2, "0"))
	assert.Equal(t, NewStringSliceV("0", "2", "3"), NewStringSliceV("1", "2", "3").Set(-3, "0"))
	assert.Equal(t, NewStringSliceV("1", "0", "3"), NewStringSliceV("1", "2", "3").Set(-2, "0"))
	assert.Equal(t, NewStringSliceV("1", "2", "0"), NewStringSliceV("1", "2", "3").Set(-1, "0"))

	// Test out of bounds
	{
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Set(5, "1"))
	}

	// Test wrong type
	{
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Set(5, "1"))
	}

	// Conversion
	{
		assert.Equal(t, NewStringSliceV(0, 2, 0), NewStringSliceV(0, 0, 0).Set(1, Object{2}))
		assert.Equal(t, NewStringSliceV(0, 2, 0), NewStringSliceV(0, 0, 0).Set(1, "2"))
		assert.Equal(t, NewStringSliceV(true, 0, 0), NewStringSliceV(0, 0, 0).Set(0, true))
		assert.Equal(t, NewStringSliceV(0, 0, 3), NewStringSliceV(0, 0, 0).Set(-1, Char('3')))
	}
}

// SetE
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_SetE_Go(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i++ {
		src[i] = "0"
	}
}

func BenchmarkStringSlice_SetE_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.SetE(i, "0")
	}
}

func ExampleStringSlice_SetE() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.SetE(0, "0"))
	// Output: [0 2 3] <nil>
}

func TestStringSlice_SetE(t *testing.T) {

	// pos - begining
	{
		slice := NewStringSliceV("1", "2", "3")
		result, err := slice.SetE(0, "0")
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("0", "2", "3"), slice)
		assert.Equal(t, NewStringSliceV("0", "2", "3"), result)

		// multiple
		result, err = slice.SetE(0, []string{"4", "5", "6"})
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), slice)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), result)

		// multiple over
		result, err = slice.SetE(0, []string{"4", "5", "6", "7"})
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), slice)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), result)
	}

	// pos - middle
	{
		slice := NewStringSliceV("1", "2", "3")
		result, err := slice.SetE(1, "0")
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("1", "0", "3"), slice)
		assert.Equal(t, NewStringSliceV("1", "0", "3"), result)
	}

	// pos - end
	{
		slice := NewStringSliceV("1", "2", "3")
		result, err := slice.SetE(2, "0")
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("1", "2", "0"), slice)
		assert.Equal(t, NewStringSliceV("1", "2", "0"), result)
	}

	// neg - begining
	{
		slice := NewStringSliceV("1", "2", "3")
		result, err := slice.SetE(-3, "0")
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("0", "2", "3"), slice)
		assert.Equal(t, NewStringSliceV("0", "2", "3"), result)

		// multiple
		result, err = slice.SetE(-3, []string{"4", "5", "6"})
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), slice)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), result)

		// multiple over
		result, err = slice.SetE(-3, []string{"4", "5", "6", "7"})
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), slice)
		assert.Equal(t, NewStringSliceV("4", "5", "6"), result)
	}

	// neg - middle
	{
		slice := NewStringSliceV("1", "2", "3")
		result, err := slice.SetE(-2, "0")
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("1", "0", "3"), slice)
		assert.Equal(t, NewStringSliceV("1", "0", "3"), result)
	}

	// neg - end
	{
		slice := NewStringSliceV("1", "2", "3")
		result, err := slice.SetE(-1, "0")
		assert.Nil(t, err)
		assert.Equal(t, NewStringSliceV("1", "2", "0"), slice)
		assert.Equal(t, NewStringSliceV("1", "2", "0"), result)
	}

	// Test out of bounds
	{
		slice, err := NewStringSliceV("1", "2", "3").SetE(5, "1")
		assert.NotNil(t, slice)
		assert.NotNil(t, err)
	}

	// Test wrong type
	{
		slice, err := NewStringSliceV("1", "2", "3").SetE(5, "1")
		assert.NotNil(t, slice)
		assert.NotNil(t, err)
	}
}

// Shift
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Shift_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 1 {
		src = src[1:]
	}
}

func BenchmarkStringSlice_Shift_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.Shift()
	}
}

func ExampleStringSlice_Shift() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Shift())
	// Output: 1
}

func TestStringSlice_Shift(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, Obj(nil), slice.Shift())
	}

	// take all and beyond
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, Obj("1"), slice.Shift())
		assert.Equal(t, NewStringSliceV("2", "3"), slice)
		assert.Equal(t, Obj("2"), slice.Shift())
		assert.Equal(t, NewStringSliceV("3"), slice)
		assert.Equal(t, Obj("3"), slice.Shift())
		assert.Equal(t, NewStringSliceV(), slice)
		assert.Equal(t, Obj(nil), slice.Shift())
		assert.Equal(t, NewStringSliceV(), slice)
	}
}

// ShiftN
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_ShiftN_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 10 {
		src = src[10:]
	}
}

func BenchmarkStringSlice_ShiftN_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 0 {
		slice.ShiftN(10)
	}
}

func ExampleStringSlice_ShiftN() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.ShiftN(2))
	// Output: [1 2]
}

func TestStringSlice_ShiftN(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.ShiftN(1))
	}

	// negative value
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("1"), slice.ShiftN(-1))
		assert.Equal(t, NewStringSliceV("2", "3"), slice)
	}

	// take none
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV(), slice.ShiftN(0))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
	}

	// take 1
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("1"), slice.ShiftN(1))
		assert.Equal(t, NewStringSliceV("2", "3"), slice)
	}

	// take 2
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("1", "2"), slice.ShiftN(2))
		assert.Equal(t, NewStringSliceV("3"), slice)
	}

	// take 3
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice.ShiftN(3))
		assert.Equal(t, NewStringSliceV(), slice)
	}

	// take beyond
	{
		slice := NewStringSliceV("1", "2", "3")
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice.ShiftN(4))
		assert.Equal(t, NewStringSliceV(), slice)
	}
}

// Single
//--------------------------------------------------------------------------------------------------

func ExampleStringSlice_Single() {
	slice := NewStringSliceV("1")
	fmt.Println(slice.Single())
	// Output: true
}

func TestStringSlice_Single(t *testing.T) {

	assert.Equal(t, false, NewStringSliceV().Single())
	assert.Equal(t, true, NewStringSliceV("1").Single())
	assert.Equal(t, false, NewStringSliceV("1", "2").Single())
}

// Slice
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Slice_Go(t *testing.B) {
	src := RangeString(nines7)
	_ = src[0:len(src)]
}

func BenchmarkStringSlice_Slice_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	slice.Slice(0, -1)
}

func ExampleStringSlice_Slice() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Slice(1, -1))
	// Output: [2 3]
}

func TestStringSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.Slice(0, -1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Slice(0, -1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewStringSliceV("1", "2", "3")
		result := original.Slice(0, -1).Set(0, "0")
		assert.Equal(t, NewStringSliceV("0", "2", "3"), original)
		assert.Equal(t, NewStringSliceV("0", "2", "3"), result)
	}

	// slice full array
	{
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Slice(0, -1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Slice(0, 1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV().Slice(0, 5))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Slice(0, -1))
		assert.Equal(t, NewStringSlice([]string{"1", "2", "3"}), NewStringSlice([]string{"1", "2", "3"}).Slice(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1").Slice(0, 2))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Slice(-6, 6))
	}

	// mutually exclusive
	{
		assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Slice(2, -3))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Slice(0, -5))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Slice(4, -1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Slice(6, -1))
		assert.Equal(t, NewStringSliceV(), NewStringSliceV("1", "2", "3", "4").Slice(3, 2))
	}

	// singles
	{
		slice := NewStringSliceV("1", "2", "3", "4")
		assert.Equal(t, NewStringSliceV("4"), slice.Slice(-1, -1))
		assert.Equal(t, NewStringSliceV("3"), slice.Slice(-2, -2))
		assert.Equal(t, NewStringSliceV("2"), slice.Slice(-3, -3))
		assert.Equal(t, NewStringSliceV("1"), slice.Slice(0, 0))
		assert.Equal(t, NewStringSliceV("1"), slice.Slice(-4, -4))
		assert.Equal(t, NewStringSliceV("2"), slice.Slice(1, 1))
		assert.Equal(t, NewStringSliceV("2"), slice.Slice(1, -3))
		assert.Equal(t, NewStringSliceV("3"), slice.Slice(2, 2))
		assert.Equal(t, NewStringSliceV("3"), slice.Slice(2, -2))
		assert.Equal(t, NewStringSliceV("4"), slice.Slice(3, 3))
		assert.Equal(t, NewStringSliceV("4"), slice.Slice(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Slice(1, -1))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Slice(-2, -1))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Slice(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Slice(0, -2))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Slice(-3, -2))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Slice(-3, 1))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2", "3").Slice(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Slice(1, -2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Slice(-3, -2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Slice(-3, 2))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3", "4").Slice(1, 2))
	}

	// random
	{
		assert.Equal(t, NewStringSliceV("1"), NewStringSliceV("1", "2", "3").Slice(0, -3))
		assert.Equal(t, NewStringSliceV("2", "3"), NewStringSliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewStringSliceV("1", "2", "3").Slice(0, 2))
	}
}

// Sort
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Sort_Go(t *testing.B) {
	src := RangeString(nines6)
	sort.Sort(sort.StringSlice(src))
}

func BenchmarkStringSlice_Sort_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.Sort()
}

func ExampleStringSlice_Sort() {
	slice := NewStringSliceV("2", "3", "1")
	fmt.Println(slice.Sort())
	// Output: [1 2 3]
}

func TestStringSlice_Sort(t *testing.T) {

	// empty
	assert.Equal(t, NewStringSliceV(), NewStringSliceV().Sort())

	// pos
	{
		slice := NewStringSliceV("5", "3", "2", "4", "1")
		sorted := slice.Sort()
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("5", "3", "2", "4", "1"), slice)
	}

	// neg
	{
		slice := NewStringSliceV("5", "3", "-2", "4", "-1")
		sorted := slice.Sort()
		assert.Equal(t, NewStringSliceV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("5", "3", "-2", "4", "-1"), slice)
	}
}

// SortM
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_SortM_Go(t *testing.B) {
	src := RangeString(nines6)
	sort.Sort(sort.StringSlice(src))
}

func BenchmarkStringSlice_SortM_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.SortM()
}

func ExampleStringSlice_SortM() {
	slice := NewStringSliceV("2", "3", "1")
	fmt.Println(slice.SortM())
	// Output: [1 2 3]
}

func TestStringSlice_SortM(t *testing.T) {

	// empty
	assert.Equal(t, NewStringSliceV(), NewStringSliceV().SortM())

	// pos
	{
		slice := NewStringSliceV("5", "3", "2", "4", "1")
		sorted := slice.SortM()
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4", "5", "6"), slice)
	}

	// neg
	{
		slice := NewStringSliceV("5", "3", "-2", "4", "-1")
		sorted := slice.SortM()
		assert.Equal(t, NewStringSliceV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("-1", "-2", "3", "4", "5", "6"), slice)
	}
}

// SortReverse
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_SortReverse_Go(t *testing.B) {
	src := RangeString(nines6)
	sort.Sort(sort.Reverse(sort.StringSlice(src)))
}

func BenchmarkStringSlice_SortReverse_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.SortReverse()
}

func ExampleStringSlice_SortReverse() {
	slice := NewStringSliceV("2", "3", "1")
	fmt.Println(slice.SortReverse())
	// Output: [3 2 1]
}

func TestStringSlice_SortReverse(t *testing.T) {

	// empty
	assert.Equal(t, NewStringSliceV(), NewStringSliceV().SortReverse())

	// pos
	{
		slice := NewStringSliceV("5", "3", "2", "4", "1")
		sorted := slice.SortReverse()
		assert.Equal(t, NewStringSliceV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("5", "3", "2", "4", "1"), slice)
	}

	// neg
	{
		slice := NewStringSliceV("5", "3", "-2", "4", "-1")
		sorted := slice.SortReverse()
		assert.Equal(t, NewStringSliceV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("5", "3", "-2", "4", "-1"), slice)
	}
}

// SortReverseM
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_SortReverseM_Go(t *testing.B) {
	src := RangeString(nines6)
	sort.Sort(sort.Reverse(sort.StringSlice(src)))
}

func BenchmarkStringSlice_SortReverseM_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	slice.SortReverseM()
}

func ExampleStringSlice_SortReverseM() {
	slice := NewStringSliceV("2", "3", "1")
	fmt.Println(slice.SortReverseM())
	// Output: [3 2 1]
}

func TestStringSlice_SortReverseM(t *testing.T) {

	// empty
	assert.Equal(t, NewStringSliceV(), NewStringSliceV().SortReverse())

	// pos
	{
		slice := NewStringSliceV("5", "3", "2", "4", "1")
		sorted := slice.SortReverseM()
		assert.Equal(t, NewStringSliceV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("5", "4", "3", "2", "1", "6"), slice)
	}

	// neg
	{
		slice := NewStringSliceV("5", "3", "-2", "4", "-1")
		sorted := slice.SortReverseM()
		assert.Equal(t, NewStringSliceV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
		assert.Equal(t, NewStringSliceV("5", "4", "3", "-2", "-1", "6"), slice)
	}
}

// String
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_String_Go(t *testing.B) {
	src := RangeString(nines6)
	_ = fmt.Sprintf("%v", src)
}

func BenchmarkStringSlice_String_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	_ = slice.String()
}

func ExampleStringSlice_String() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_String(t *testing.T) {
	// nil
	assert.Equal(t, "[]", (*StringSlice)(nil).String())

	// empty
	assert.Equal(t, "[]", NewStringSliceV().String())

	// pos
	{
		slice := NewStringSliceV("5", "3", "2", "4", "1")
		sorted := slice.SortReverseM()
		assert.Equal(t, "[5 4 3 2 1 6]", sorted.Append("6").String())
		assert.Equal(t, "[5 4 3 2 1 6]", slice.String())
	}

	// neg
	{
		slice := NewStringSliceV("5", "3", "-2", "4", "-1")
		sorted := slice.SortReverseM()
		assert.Equal(t, "[5 4 3 -2 -1 6]", sorted.Append("6").String())
		assert.Equal(t, "[5 4 3 -2 -1 6]", slice.String())
	}
}

// Swap
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Swap_Go(t *testing.B) {
	src := RangeString(nines6)
	for i := 0; i < len(src); i++ {
		if i+1 < len(src) {
			src[i], src[i+1] = src[i+1], src[i]
		}
	}
}

func BenchmarkStringSlice_Swap_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Swap(i, i+1)
		}
	}
}

func ExampleStringSlice_Swap() {
	slice := NewStringSliceV("2", "3", "1")
	slice.Swap(0, 2)
	slice.Swap(1, 2)
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestStringSlice_Swap(t *testing.T) {

	// invalid cases
	{
		var slice *StringSlice
		slice.Swap(0, 0)
		assert.Equal(t, (*StringSlice)(nil), slice)

		slice = NewStringSliceV()
		slice.Swap(0, 0)
		assert.Equal(t, NewStringSliceV(), slice)

		slice.Swap(1, 2)
		assert.Equal(t, NewStringSliceV(), slice)

		slice.Swap(-1, 2)
		assert.Equal(t, NewStringSliceV(), slice)

		slice.Swap(1, -2)
		assert.Equal(t, NewStringSliceV(), slice)
	}

	// normal
	{
		slice := NewStringSliceV("0", "1", "2")
		slice.Swap(0, 1)
		assert.Equal(t, NewStringSliceV("1", "0", "2"), slice)
	}
}

// Take
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Take_Go(t *testing.B) {
	src := RangeString(nines7)
	for len(src) > 11 {
		i := 1
		n := 10
		if i+n < len(src) {
			src = append(src[:i], src[i+n:]...)
		} else {
			src = src[:i]
		}
	}
}

func BenchmarkStringSlice_Take_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines7))
	for slice.Len() > 1 {
		slice.Take(1, 10)
	}
}

func ExampleStringSlice_Take() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.Take(0, 1))
	// Output: [1 2]
}

func TestStringSlice_Take(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.Take(0, 1))
	}

	// invalid
	{
		slice := NewStringSliceV("1", "2", "3", "4")
		assert.Equal(t, NewStringSliceV(), slice.Take(1))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice)
		assert.Equal(t, NewStringSliceV(), slice.Take(4, 4))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice)
	}

	// take 1
	{
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1"), slice.Take(0, 0))
			assert.Equal(t, NewStringSliceV("2", "3", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("2"), slice.Take(1, 1))
			assert.Equal(t, NewStringSliceV("1", "3", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("3"), slice.Take(2, 2))
			assert.Equal(t, NewStringSliceV("1", "2", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("4"), slice.Take(3, 3))
			assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("4"), slice.Take(-1, -1))
			assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("3"), slice.Take(-2, -2))
			assert.Equal(t, NewStringSliceV("1", "2", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("2"), slice.Take(-3, -3))
			assert.Equal(t, NewStringSliceV("1", "3", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1"), slice.Take(-4, -4))
			assert.Equal(t, NewStringSliceV("2", "3", "4"), slice)
		}
	}

	// take 2
	{
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2"), slice.Take(0, 1))
			assert.Equal(t, NewStringSliceV("3", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("2", "3"), slice.Take(1, 2))
			assert.Equal(t, NewStringSliceV("1", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("3", "4"), slice.Take(2, 3))
			assert.Equal(t, NewStringSliceV("1", "2"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("3", "4"), slice.Take(-2, -1))
			assert.Equal(t, NewStringSliceV("1", "2"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("2", "3"), slice.Take(-3, -2))
			assert.Equal(t, NewStringSliceV("1", "4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2"), slice.Take(-4, -3))
			assert.Equal(t, NewStringSliceV("3", "4"), slice)
		}
	}

	// take 3
	{
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2", "3"), slice.Take(0, 2))
			assert.Equal(t, NewStringSliceV("4"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("2", "3", "4"), slice.Take(-3, -1))
			assert.Equal(t, NewStringSliceV("1"), slice)
		}
	}

	// take everything and beyond
	{
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Take())
			assert.Equal(t, NewStringSliceV(), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Take(0, 3))
			assert.Equal(t, NewStringSliceV(), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Take(0, -1))
			assert.Equal(t, NewStringSliceV(), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Take(-4, -1))
			assert.Equal(t, NewStringSliceV(), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Take(-6, -1))
			assert.Equal(t, NewStringSliceV(), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Take(0, 10))
			assert.Equal(t, NewStringSliceV(), slice)
		}
	}

	// move index within bounds
	{
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("4"), slice.Take(3, 4))
			assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
		}
		{
			slice := NewStringSliceV("1", "2", "3", "4")
			assert.Equal(t, NewStringSliceV("1"), slice.Take(-5, 0))
			assert.Equal(t, NewStringSliceV("2", "3", "4"), slice)
		}
	}
}

// TakeAt
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_TakeAt_Go(t *testing.B) {
	src := RangeString(nines5)
	index := RangeString(nines5)
	for i := range index {
		if i+1 < len(src) {
			src = append(src[:i], src[i+1:]...)
		} else if i >= 0 && i < len(src) {
			src = src[:i]
		}
	}
}

func BenchmarkStringSlice_TakeAt_Slice(t *testing.B) {
	src := RangeString(nines5)
	index := RangeString(nines5)
	slice := NewStringSlice(src)
	for i := range index {
		slice.TakeAt(i)
	}
}

func ExampleStringSlice_TakeAt() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.TakeAt(1))
	// Output: 2
}

func TestStringSlice_TakeAt(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, Obj(nil), slice.TakeAt(0))
	}

	// all and more
	{
		slice := NewStringSliceV("0", "1", "2")
		assert.Equal(t, Obj("2"), slice.TakeAt(-1))
		assert.Equal(t, NewStringSliceV("0", "1"), slice)
		assert.Equal(t, Obj("1"), slice.TakeAt(-1))
		assert.Equal(t, NewStringSliceV("0"), slice)
		assert.Equal(t, Obj("0"), slice.TakeAt(-1))
		assert.Equal(t, NewStringSliceV(), slice)
		assert.Equal(t, Obj(nil), slice.TakeAt(-1))
		assert.Equal(t, NewStringSliceV(), slice)
	}

	// take invalid
	{
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj(nil), slice.TakeAt(3))
			assert.Equal(t, NewStringSliceV("0", "1", "2"), slice)
		}
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj(nil), slice.TakeAt(-4))
			assert.Equal(t, NewStringSliceV("0", "1", "2"), slice)
		}
	}

	// take last
	{
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj("2"), slice.TakeAt(2))
			assert.Equal(t, NewStringSliceV("0", "1"), slice)
		}
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj("2"), slice.TakeAt(-1))
			assert.Equal(t, NewStringSliceV("0", "1"), slice)
		}
	}

	// take middle
	{
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj("1"), slice.TakeAt(1))
			assert.Equal(t, NewStringSliceV("0", "2"), slice)
		}
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj("1"), slice.TakeAt(-2))
			assert.Equal(t, NewStringSliceV("0", "2"), slice)
		}
	}

	// take first
	{
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj("0"), slice.TakeAt(0))
			assert.Equal(t, NewStringSliceV("1", "2"), slice)
		}
		{
			slice := NewStringSliceV("0", "1", "2")
			assert.Equal(t, Obj("0"), slice.TakeAt(-3))
			assert.Equal(t, NewStringSliceV("1", "2"), slice)
		}
	}
}

// TakeW
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_TakeW_Go(t *testing.B) {
	new := []string{}
	src := RangeString(nines5)
	l := len(src)
	for i := 0; i < l; i++ {
		if Obj(src[i]).ToInt()%2 == 0 {
			new = append(new, src[i])
			if i+1 < l {
				src = append(src[:i], src[i+1:]...)
			} else if i >= 0 && i < l {
				src = src[:i]
			}
			l--
			i--
		}
	}
}

func BenchmarkStringSlice_TakeW_Slice(t *testing.B) {
	slice := NewStringSlice(RangeString(nines5))
	slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
}

func ExampleStringSlice_TakeW() {
	slice := NewStringSliceV("1", "2", "3")
	fmt.Println(slice.TakeW(func(x O) bool {
		return ExB(Obj(x).ToInt()%2 == 0)
	}))
	// Output: [2]
}

func TestStringSlice_TakeW(t *testing.T) {

	// take all odd values
	{
		slice := NewStringSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 != 0) })
		assert.Equal(t, NewStringSliceV("2", "4", "6", "8"), slice)
		assert.Equal(t, NewStringSliceV("1", "3", "5", "7", "9"), new)
	}

	// take all even values
	{
		slice := NewStringSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
		assert.Equal(t, NewStringSliceV("1", "3", "5", "7", "9"), slice)
		assert.Equal(t, NewStringSliceV("2", "4", "6", "8"), new)
	}
}

// Union
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Union_Go(t *testing.B) {
	// src := RangeStr(nines7)
	// for len(src) > 10 {
	// 	src = src[10:]
	// }
}

func BenchmarkStringSlice_Union_Slice(t *testing.B) {
	// slice := NewStringSlice(RangeStr(nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleStringSlice_Union() {
	slice := NewStringSliceV("1", "2")
	fmt.Println(slice.Union([]string{"2", "3"}))
	// Output: [1 2 3]
}

func TestStringSlice_Union(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Union(NewStringSliceV("1", "2")))
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Union([]string{"1", "2"}))
	}

	// size of one
	{
		slice := NewStringSliceV("1")
		union := slice.Union([]string{"1", "2", "3"})
		assert.Equal(t, NewStringSliceV("1", "2", "3"), union)
		assert.Equal(t, NewStringSliceV("1"), slice)
	}

	// one duplicate
	{
		slice := NewStringSliceV("1", "1")
		union := slice.Union(NewStringSliceV("2", "3"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), union)
		assert.Equal(t, NewStringSliceV("1", "1"), slice)
	}

	// multiple duplicates
	{
		slice := NewStringSliceV("1", "2", "2", "3", "3")
		union := slice.Union([]string{"1", "2", "3"})
		assert.Equal(t, NewStringSliceV("1", "2", "3"), union)
		assert.Equal(t, NewStringSliceV("1", "2", "2", "3", "3"), slice)
	}

	// no duplicates
	{
		slice := NewStringSliceV("1", "2", "3")
		union := slice.Union([]string{"4", "5"})
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4", "5"), union)
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
	}

	// nils
	{
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").Union((*[]string)(nil)))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").Union((*StringSlice)(nil)))
	}
}

// UnionM
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_UnionM_Go(t *testing.B) {
	// src := RangeStr(nines7)
	// for len(src) > 10 {
	// 	src = src[10:]
	// }
}

func BenchmarkStringSlice_UnionM_Slice(t *testing.B) {
	// slice := NewStringSlice(RangeStr(nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleStringSlice_UnionM() {
	slice := NewStringSliceV("1", "2")
	fmt.Println(slice.UnionM([]string{"2", "3"}))
	// Output: [1 2 3]
}

func TestStringSlice_UnionM(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV("1", "2"), slice.UnionM(NewStringSliceV("1", "2")))
		assert.Equal(t, (*StringSlice)(nil), slice)
	}

	// size of one
	{
		slice := NewStringSliceV("1")
		union := slice.UnionM([]string{"1", "2", "3"})
		assert.Equal(t, NewStringSliceV("1", "2", "3"), union)
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
	}

	// one duplicate
	{
		slice := NewStringSliceV("1", "1")
		union := slice.UnionM(NewStringSliceV("2", "3"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), union)
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
	}

	// multiple duplicates
	{
		slice := NewStringSliceV("1", "2", "2", "3", "3")
		union := slice.UnionM([]string{"1", "2", "3"})
		assert.Equal(t, NewStringSliceV("1", "2", "3"), union)
		assert.Equal(t, NewStringSliceV("1", "2", "3"), slice)
	}

	// no duplicates
	{
		slice := NewStringSliceV("1", "2", "3")
		union := slice.UnionM([]string{"4", "5"})
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4", "5"), union)
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4", "5"), slice)
	}

	// nils
	{
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").UnionM((*[]string)(nil)))
		assert.Equal(t, NewStringSliceV("1", "2"), NewStringSliceV("1", "2").UnionM((*StringSlice)(nil)))
	}
}

// Uniq
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_Uniq_Go(t *testing.B) {
	// src := RangeStr(nines7)
	// for len(src) > 10 {
	// 	src = src[10:]
	// }
}

func BenchmarkStringSlice_Uniq_Slice(t *testing.B) {
	// slice := NewStringSlice(RangeStr(nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleStringSlice_Uniq() {
	slice := NewStringSliceV("1", "2", "3", "3")
	fmt.Println(slice.Uniq())
	// Output: [1 2 3]
}

func TestStringSlice_Uniq(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, NewStringSliceV(), slice.Uniq())
	}

	// size of one
	{
		slice := NewStringSliceV("1")
		uniq := slice.Uniq()
		assert.Equal(t, NewStringSliceV("1"), uniq)
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1"), uniq)
	}

	// one duplicate
	{
		slice := NewStringSliceV("1", "1")
		uniq := slice.Uniq()
		assert.Equal(t, NewStringSliceV("1"), uniq)
		assert.Equal(t, NewStringSliceV("1", "1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1"), uniq)
	}

	// multiple duplicates
	{
		slice := NewStringSliceV("1", "2", "2", "3", "3")
		uniq := slice.Uniq()
		assert.Equal(t, NewStringSliceV("1", "2", "3"), uniq)
		assert.Equal(t, NewStringSliceV("1", "2", "2", "3", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), uniq)
	}

	// no duplicates
	{
		slice := NewStringSliceV("1", "2", "3")
		uniq := slice.Uniq()
		assert.Equal(t, NewStringSliceV("1", "2", "3"), uniq)
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), uniq)
	}
}

// UniqM
// --------------------------------------------------------------------------------------------------
func BenchmarkStringSlice_UniqM_Go(t *testing.B) {
	// src := RangeStr(nines7)
	// for len(src) > 10 {
	// 	src = src[10:]
	// }
}

func BenchmarkStringSlice_UniqM_Slice(t *testing.B) {
	// slice := NewStringSlice(RangeStr(nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleStringSlice_UniqM() {
	slice := NewStringSliceV("1", "2", "3", "3")
	fmt.Println(slice.UniqM())
	// Output: [1 2 3]
}

func TestStringSlice_UniqM(t *testing.T) {

	// nil or empty
	{
		var slice *StringSlice
		assert.Equal(t, (*StringSlice)(nil), slice.UniqM())
	}

	// size of one
	{
		slice := NewStringSliceV("1")
		uniq := slice.UniqM()
		assert.Equal(t, NewStringSliceV("1"), uniq)
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1", "2"), uniq)
	}

	// one duplicate
	{
		slice := NewStringSliceV("1", "1")
		uniq := slice.UniqM()
		assert.Equal(t, NewStringSliceV("1"), uniq)
		assert.Equal(t, NewStringSliceV("1", "2"), slice.Append("2"))
		assert.Equal(t, NewStringSliceV("1", "2"), uniq)
	}

	// multiple duplicates
	{
		slice := NewStringSliceV("1", "2", "2", "3", "3")
		uniq := slice.UniqM()
		assert.Equal(t, NewStringSliceV("1", "2", "3"), uniq)
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), uniq)
	}

	// no duplicates
	{
		slice := NewStringSliceV("1", "2", "3")
		uniq := slice.UniqM()
		assert.Equal(t, NewStringSliceV("1", "2", "3"), uniq)
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStringSliceV("1", "2", "3", "4"), uniq)
	}
}

func RangeString(size int) (new []string) {
	for _, x := range Range(0, size) {
		new = append(new, fmt.Sprint(x))
	}
	return
}
