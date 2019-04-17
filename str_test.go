package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A
//--------------------------------------------------------------------------------------------------
func ExampleA() {
	str := A("test")
	fmt.Println(str)
	// Output: test
}

func TestStr_A(t *testing.T) {
	assert.Equal(t, "test", A("test").A())
}

// NewStr
//--------------------------------------------------------------------------------------------------
func ExampleNewStr() {
	str := NewStr("test")
	fmt.Println(str)
	// Output: test
}

func TestStr_NewStr(t *testing.T) {
	// nil
	{
		assert.Equal(t, "", NewStr(nil).A())
	}

	// Str
	{
		assert.Equal(t, "test", A(A("test")).A())
	}

	// string
	{
		assert.Equal(t, "test", A("test").A())
		assert.Equal(t, "test1", A([]string{"test", "1"}).O())
	}

	// runes
	{
		assert.Equal(t, "b", A('b').A())
		assert.Equal(t, "test", A([]rune("test")).A())
	}

	// bytes
	{
		assert.Equal(t, "test", A([]byte{0x74, 0x65, 0x73, 0x74}).A())
	}

	// ints
	{
		assert.Equal(t, "10", A(10).A())
	}
}

// Any
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Any_Go(t *testing.B) {
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

func BenchmarkStr_Any_Slice(t *testing.B) {
	src := RangeString(nines4)
	slice := NewStr(src)
	for i := range src {
		slice.Any(i)
	}
}

func ExampleStr_Any_empty() {
	slice := NewStrV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleStr_Any_notEmpty() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleStr_Any_contains() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.Any("1"))
	// Output: true
}

func ExampleStr_Any_containsAny() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.Any("0", "1"))
	// Output: true
}

func TestStr_Any(t *testing.T) {

	// empty
	var nilSlice *Str
	assert.False(t, nilSlice.Any())
	assert.False(t, NewStrV().Any())

	// single
	assert.True(t, NewStrV("2").Any())

	// invalid
	assert.False(t, NewStrV("12").Any(Object{"2"}))

	assert.True(t, NewStrV("123").Any("2"))
	assert.False(t, NewStrV("123").Any(4))
	assert.True(t, NewStrV("123").Any(4, "3"))
	assert.False(t, NewStrV("123").Any(4, 5))
}

// AnyS
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_AnyS_Go(t *testing.B) {
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

func BenchmarkStr_AnyS_Slice(t *testing.B) {
	src := RangeString(nines4)
	slice := NewStr(src)
	for _, x := range src {
		slice.Any([]string{x})
	}
}

func ExampleStr_AnyS() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.AnyS([]string{"0", "1"}))
	// Output: true
}

func TestStr_AnyS(t *testing.T) {
	// nil
	{
		var slice *Str
		assert.False(t, slice.AnyS([]string{"1"}))
		assert.False(t, NewStrV("1").AnyS(nil))
	}

	// []string
	{
		assert.True(t, NewStrV("123").AnyS([]string{"1"}))
		assert.True(t, NewStrV("123").AnyS([]string{"4", "3"}))
		assert.False(t, NewStrV("123").AnyS([]string{"4", "5"}))
	}

	// *[]string
	{
		assert.True(t, NewStrV("123").AnyS(&([]string{"1"})))
		assert.True(t, NewStrV("123").AnyS(&([]string{"4", "3"})))
		assert.False(t, NewStrV("123").AnyS(&([]string{"4", "5"})))
	}

	// Slice
	{
		assert.True(t, NewStrV("123").AnyS(NewStringSliceV("1")))
		assert.True(t, NewStrV("123").AnyS(NewStringSliceV("4", "3")))
		assert.False(t, NewStrV("123").AnyS(NewStringSliceV("4", "5")))
	}

	// Str
	{
		assert.True(t, NewStrV("123").AnyS([]Str{Str("1"), Str("2")}))
		assert.True(t, NewStrV("123").AnyS(&[]Str{Str("1"), Str("2")}))
		assert.False(t, NewStrV("123").AnyS([]Str{Str("4"), Str("5")}))
	}

	// invalid types
	assert.False(t, NewStrV("1", "2").AnyS(nil))
	assert.False(t, NewStrV("1", "2").AnyS((*[]string)(nil)))
	assert.False(t, NewStrV("1", "2").AnyS((*Str)(nil)))
	assert.False(t, NewStrV("1", "2").AnyS([]int{2}))
}

// AnyW
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_AnyW_Go(t *testing.B) {
	src := RangeString(nines5)
	for _, x := range src {
		if x == string(nines4) {
			break
		}
	}
}

func BenchmarkStr_AnyW_Slice(t *testing.B) {
	src := RangeString(nines5)
	NewStr(src).AnyW(func(x O) bool {
		return ExB(x.(string) == string(nines4))
	})
}

func ExampleStr_AnyW() {
	slice := NewStr("123")
	fmt.Println(slice.AnyW(func(x O) bool {
		return ExB(x.(rune) == '2')
	}))
	// Output: true
}

func TestStr_AnyW(t *testing.T) {

	// empty
	var slice *Str
	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(rune) > '0') }))
	assert.False(t, NewStrV().AnyW(func(x O) bool { return ExB(x.(rune) > '0') }))

	// single
	assert.True(t, NewStr("2").AnyW(func(x O) bool { return ExB(x.(rune) > '0') }))
	assert.True(t, NewStr("12").AnyW(func(x O) bool { return ExB(x.(rune) == '2') }))
	assert.True(t, NewStr("123").AnyW(func(x O) bool { return ExB(x.(rune) == '4' || x.(rune) == '3') }))
}

// Append
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Append_Go(t *testing.B) {
	src := []string{}
	for _, i := range RangeString(nines6) {
		src = append(src, i)
	}
}

func BenchmarkStr_Append_Slice(t *testing.B) {
	slice := NewStrV()
	for _, i := range RangeString(nines6) {
		slice.Append(i)
	}
}

func ExampleStr_Append() {
	slice := NewStrV("1").Append("2").Append("3")
	fmt.Println(slice)
	// Output: 123
}

func TestStr_Append(t *testing.T) {

	// nil
	{
		var nilSlice *Str
		assert.Equal(t, NewStr("0"), nilSlice.Append("0"))
		assert.Equal(t, (*Str)(nil), nilSlice)
	}

	// Append one back to back
	{
		var slice *Str
		assert.Equal(t, true, slice.Nil())
		slice = NewStrV()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, false, slice.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		slice.Append("1")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, "1", slice.O())

		// Second append another which will be 2x at most
		slice.Append("2")
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, "12", slice.O())
		assert.Equal(t, NewStrV("12"), slice)
	}

	// Start with just appending without chaining
	{
		slice := NewStrV()
		assert.Equal(t, 0, slice.Len())
		slice.Append("1")
		assert.Equal(t, "1", slice.O())
		slice.Append("2")
		assert.Equal(t, "12", slice.O())
	}

	// Start with nil not chained
	{
		slice := NewStrV()
		assert.Equal(t, 0, slice.Len())
		slice.Append("1").Append("2").Append("3")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, "123", slice.O())
	}

	// Start with nil chained
	{
		slice := NewStrV().Append("1").Append("2")
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, "12", slice.O())
	}

	// Start with non nil
	{
		slice := NewStrV("1").Append("2").Append("3")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, "123", slice.O())
		assert.Equal(t, NewStrV("123"), slice)
	}

	// Use append result directly
	{
		slice := NewStrV("1")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, "12", slice.Append("2").O())
		assert.Equal(t, NewStrV("12"), slice)
	}

	// other types
	{
		assert.Equal(t, "1test4", NewStrV("1").Append([]byte{0x74, 0x65, 0x73, 0x74}).Append("4").O())
		assert.Equal(t, "1234", NewStrV("1").Append([]rune{'2', '3'}).Append("4").O())
	}
}

// AppendV
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_AppendV_Go(t *testing.B) {
	src := []string{}
	src = append(src, RangeString(nines6)...)
}

func BenchmarkStr_AppendV_Slice(t *testing.B) {
	n := NewStrV()
	new := rangeO(0, nines6)
	n.AppendV(new...)
}

func ExampleStr_AppendV() {
	slice := NewStrV("1").AppendV("2", "3")
	fmt.Println(slice)
	// Output: 123
}

func TestStr_AppendV(t *testing.T) {

	// nil
	{
		var nilSlice *Str
		assert.Equal(t, NewStrV("12"), nilSlice.AppendV("1", "2"))
	}

	// other types
	{
		assert.Equal(t, NewStrV("1test4"), NewStrV("1").AppendV([]byte{0x74, 0x65, 0x73, 0x74}).AppendV("4"))
		assert.Equal(t, NewStrV("1234"), NewStrV("1").AppendV([]rune{'2', '3'}).AppendV("4"))
	}

	// many
	{
		assert.Equal(t, NewStrV("123"), NewStrV("1").AppendV("2", "3"))
		assert.Equal(t, NewStrV("12345"), NewStrV("1").AppendV("2", "3").AppendV("4", "5"))
	}
}

// At
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_At_Go(t *testing.B) {
	src := RangeString(nines6)
	for _, x := range src {
		assert.IsType(t, 0, x)
	}
}

func BenchmarkStr_At_Slice(t *testing.B) {
	src := RangeString(nines6)
	slice := NewStr(src)
	for i := 0; i < len(src); i++ {
		_, ok := (slice.At(i).O()).(string)
		assert.True(t, ok)
	}
}

func ExampleStr_At() {
	fmt.Println(NewStrV("123").At(2).A())
	// Output: 3
}

func TestStr_At(t *testing.T) {

	// nil
	{
		var nilSlice *Str
		assert.Equal(t, Obj(nil), nilSlice.At(0))
	}

	// src
	{
		slice := NewStrV("1", "2", "3", "4")
		assert.Equal(t, "4", slice.At(-1).A())
		assert.Equal(t, "3", slice.At(-2).A())
		assert.Equal(t, "2", slice.At(-3).A())
		assert.Equal(t, "1", slice.At(0).A())
		assert.Equal(t, "2", slice.At(1).A())
		assert.Equal(t, "3", slice.At(2).A())
		assert.Equal(t, "4", slice.At(3).A())
		assert.Equal(t, Char(52), slice.At(3).O())
	}

	// index out of bounds
	{
		slice := NewStrV("1")
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------
func ExampleStr_Clear() {
	slice := NewStrV("1").Concat([]string{"2", "3"})
	fmt.Println(slice.Clear())
	// Output:
}

func TestStr_Clear(t *testing.T) {

	// nil
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.Clear())
		assert.Equal(t, (*Str)(nil), slice)
	}

	// int
	{
		slice := NewStrV("1", "2", "3", "4")
		assert.Equal(t, "1234", slice.A())
		assert.Equal(t, NewStrV(), slice.Clear())
		assert.Equal(t, NewStrV(), slice.Clear())
		assert.Equal(t, NewStrV(), slice)
		assert.Equal(t, "", slice.A())
	}
}

// Concat
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Concat_Go(t *testing.B) {
	dest := []string{}
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkStr_Concat_Slice(t *testing.B) {
	dest := NewStrV()
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.Concat(src[j:i])
		j = i
	}
}

func ExampleStr_Concat() {
	slice := NewStrV("1").Concat([]string{"2", "3"})
	fmt.Println(slice)
	// Output: 123
}

func TestStr_Concat(t *testing.T) {

	// nil
	{
		var slice *Str
		assert.Equal(t, "12", slice.Concat([]string{"1", "2"}).O())
		assert.Equal(t, "12", NewStrV("1", "2").Concat(nil).O())
	}

	// []string
	{
		slice := NewStrV("1")
		concated := slice.Concat([]string{"2", "3"})
		assert.Equal(t, "12", slice.Append("2").O())
		assert.Equal(t, "123", concated.O())
	}

	// *[]string
	{
		slice := NewStrV("1")
		concated := slice.Concat(&([]string{"2", "3"}))
		assert.Equal(t, "12", slice.Append("2").O())
		assert.Equal(t, "123", concated.O())
	}

	// *Str
	{
		slice := NewStrV("1")
		concated := slice.Concat(NewStrV("2", "3"))
		assert.Equal(t, "12", slice.Append("2").O())
		assert.Equal(t, "123", concated.O())
	}

	// Str
	{
		slice := NewStrV("1")
		concated := slice.Concat(*NewStrV("2", "3"))
		assert.Equal(t, "12", slice.Append("2").O())
		assert.Equal(t, "123", concated.O())
	}

	// Slice
	{
		slice := NewStrV("1")
		concated := slice.Concat(Slice(NewStrV("2", "3")))
		assert.Equal(t, "12", slice.Append("2").O())
		assert.Equal(t, "123", concated.O())
	}

	// nils
	{
		assert.Equal(t, "12", NewStrV("1", "2").Concat((*[]string)(nil)).O())
		assert.Equal(t, "12", NewStrV("1", "2").Concat((*Str)(nil)).O())
	}
}

// ConcatM
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_ConcatM_Go(t *testing.B) {
	dest := []string{}
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkStr_ConcatM_Slice(t *testing.B) {
	dest := NewStrV()
	src := RangeString(nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.ConcatM(src[j:i])
		j = i
	}
}

func ExampleStr_ConcatM() {
	slice := NewStrV("1").ConcatM([]string{"2", "3"})
	fmt.Println(slice)
	// Output: 123
}

func TestStr_ConcatM(t *testing.T) {

	// nil
	{
		var slice *Str
		assert.Equal(t, NewStrV("12"), slice.ConcatM([]string{"1", "2"}))
		assert.Equal(t, NewStrV("12"), NewStrV("1", "2").ConcatM(nil))
	}

	// []string
	{
		slice := NewStrV("1")
		concated := slice.ConcatM([]string{"2", "3"})
		assert.Equal(t, NewStrV("1234"), slice.Append("4"))
		assert.Equal(t, NewStrV("1234"), concated)
	}

	// *[]string
	{
		slice := NewStrV("1")
		concated := slice.ConcatM(&([]string{"2", "3"}))
		assert.Equal(t, NewStrV("1234"), slice.Append("4"))
		assert.Equal(t, NewStrV("1234"), concated)
	}

	// *Str
	{
		slice := NewStrV("1")
		concated := slice.ConcatM(NewStrV("2", "3"))
		assert.Equal(t, NewStrV("1234"), slice.Append("4"))
		assert.Equal(t, NewStrV("1234"), concated)
	}

	// Str
	{
		slice := NewStrV("1")
		concated := slice.ConcatM(*NewStrV("2", "3"))
		assert.Equal(t, NewStrV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStrV("1", "2", "3", "4"), concated)
	}

	// Slice
	{
		slice := NewStrV("1")
		concated := slice.ConcatM(Slice(NewStrV("2", "3")))
		assert.Equal(t, NewStrV("1", "2", "3", "4"), slice.Append("4"))
		assert.Equal(t, NewStrV("1", "2", "3", "4"), concated)
	}

	// nils
	{
		assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2").ConcatM((*[]string)(nil)))
		assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2").ConcatM((*Str)(nil)))
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Copy_Go(t *testing.B) {
	src := RangeString(nines6)
	dst := make([]string, len(src), len(src))
	copy(dst, src)
}

func BenchmarkStr_Copy_Slice(t *testing.B) {
	slice := NewStr(RangeString(nines6))
	slice.Copy()
}

func ExampleStr_Copy() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.Copy())
	// Output: 123
}

func TestStr_Copy(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.Copy(0, -1))
		assert.Equal(t, NewStrV(), NewStrV("0").Clear().Copy(0, -1))
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := NewStrV("1", "2", "3")
		result := original.Copy(0, -1)
		assert.Equal(t, NewStrV("1", "2", "3"), original)
		assert.Equal(t, NewStrV("1", "2", "3"), result)
		result.Set(0, "0")
		assert.Equal(t, NewStrV("1", "2", "3"), original)
		assert.Equal(t, NewStrV("0", "2", "3"), result)
	}

	// copy full array
	{
		assert.Equal(t, NewStrV(), NewStrV().Copy())
		assert.Equal(t, NewStrV(), NewStrV().Copy(0, -1))
		assert.Equal(t, NewStrV(), NewStrV().Copy(0, 1))
		assert.Equal(t, NewStrV(), NewStrV().Copy(0, 5))
		assert.Equal(t, NewStrV("1"), NewStrV("1").Copy())
		assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3").Copy())
		assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3").Copy(0, -1))
		assert.Equal(t, "123", NewStr([]string{"1", "2", "3"}).Copy().A())
		assert.Equal(t, "123", NewStr([]string{"1", "2", "3"}).Copy(0, -1).A())
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewStrV("1"), NewStrV("1").Copy(0, 2))
		assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3").Copy(-6, 6))
	}

	// mutually exclusive
	{
		slice := NewStrV("1", "2", "3", "4")
		assert.Equal(t, NewStrV(), slice.Copy(2, -3))
		assert.Equal(t, NewStrV(), slice.Copy(0, -5))
		assert.Equal(t, NewStrV(), slice.Copy(4, -1))
		assert.Equal(t, NewStrV(), slice.Copy(6, -1))
		assert.Equal(t, NewStrV(), slice.Copy(3, -2))
	}

	// singles
	{
		slice := NewStrV("1", "2", "3", "4")
		assert.Equal(t, NewStrV("4"), slice.Copy(-1, -1))
		assert.Equal(t, NewStrV("3"), slice.Copy(-2, -2))
		assert.Equal(t, NewStrV("2"), slice.Copy(-3, -3))
		assert.Equal(t, NewStrV("1"), slice.Copy(0, 0))
		assert.Equal(t, NewStrV("1"), slice.Copy(-4, -4))
		assert.Equal(t, NewStrV("2"), slice.Copy(1, 1))
		assert.Equal(t, NewStrV("2"), slice.Copy(1, -3))
		assert.Equal(t, NewStrV("3"), slice.Copy(2, 2))
		assert.Equal(t, NewStrV("3"), slice.Copy(2, -2))
		assert.Equal(t, NewStrV("4"), slice.Copy(3, 3))
		assert.Equal(t, NewStrV("4"), slice.Copy(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3").Copy(1, -1))
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3").Copy(1, 2))
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3").Copy(-2, -1))
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3").Copy(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3").Copy(0, -2))
		assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3").Copy(-3, -2))
		assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3").Copy(-3, 1))
		assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3").Copy(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3", "4").Copy(1, -2))
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3", "4").Copy(-3, -2))
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3", "4").Copy(-3, 2))
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3", "4").Copy(1, 2))
	}

	// random
	{
		assert.Equal(t, NewStrV("1"), NewStrV("1", "2", "3").Copy(0, -3))
		assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3").Copy(1, 2))
		assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3").Copy(0, 2))
	}
}

// // All
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_All() {
// 	fmt.Println(A("foobar").All([]string{"foo"}))
// 	// Output: true
// }

// func TestStr_All(t *testing.T) {
// 	assert.True(t, A("test").All([]string{"tes", "est"}))
// 	assert.False(t, A("test").All([]string{"bob", "est"}))
// }

// // AllV
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_AllV() {
// 	fmt.Println(A("foobar").AllV("foo"))
// 	// Output: true
// }

// func TestStr_AllV(t *testing.T) {
// 	assert.True(t, A("test").AllV("tes", "est"))
// 	assert.False(t, A("test").AllV("bob", "est"))
// }

// // Any
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_Any() {
// 	fmt.Println(A("foobar").Any([]string{"foo"}))
// 	// Output: true
// }

// func TestStr_Any(t *testing.T) {
// 	assert.True(t, A("test").Any([]string{"tes", "est"}))
// 	assert.True(t, A("test").Any([]string{"bob", "est"}))
// 	assert.False(t, A("test").Any([]string{"bob", "foo"}))
// }

// // AnyV
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_AnyV() {
// 	fmt.Println(A("foobar").AnyV("foo"))
// 	// Output: true
// }

// func TestStr_AnyV(t *testing.T) {
// 	assert.True(t, A("test").AnyV("tes", "est"))
// 	assert.True(t, A("test").AnyV("bob", "est"))
// 	assert.False(t, A("test").AnyV("bob", "foo"))
// }

// // Ascii
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_Ascii() {
// 	fmt.Println(A("2�gspu�data").Ascii().A())
// 	// Output: 2 gspu data
// }

// func TestStr_Ascii(t *testing.T) {
// 	assert.Equal(t, A("2 gspu data gspm data"), A("2�gspu�data�gspm�data").Ascii())
// }

// // AsciiA
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_AsciiA() {
// 	fmt.Println(A("2�gspu�data").AsciiA())
// 	// Output: 2 gspu data
// }

// func TestStr_AsciiA(t *testing.T) {
// 	assert.Equal(t, "2 gspu data gspm data", A("2�gspu�data�gspm�data").AsciiA())
// }

// // AsciiOnly
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_AsciiOnly() {
// 	fmt.Println(A("foo").AsciiOnly())
// 	// Output: true
// }

// func TestStr_AsciiOnly(t *testing.T) {
// 	assert.Equal(t, true, A("foobar").AsciiOnly())
// 	assert.Equal(t, false, A("2�gspu�data�gspm�data").AsciiOnly())
// }

// // At
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_At() {
// 	fmt.Println(A("foobar").At(-1))
// 	// Output: 114
// }

// func TestStr_At(t *testing.T) {
// 	q := A("test")
// 	assert.Equal(t, 't', q.At(0))
// 	assert.Equal(t, 'e', q.At(1))
// 	assert.Equal(t, 's', q.At(2))
// 	assert.Equal(t, 't', q.At(3))
// 	assert.Equal(t, 't', q.At(-1))
// 	assert.Equal(t, 's', q.At(-2))
// 	assert.Equal(t, 'e', q.At(-3))
// 	assert.Equal(t, 't', q.At(-4))
// 	assert.Equal(t, rune(0), q.At(5))
// }

// // AtE
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_AtE() {
// 	fmt.Println(A("foobar").AtE(-1))
// 	// Output: 114 <nil>
// }

// func TestStr_AtE(t *testing.T) {
// 	q := A("test")
// 	r, err := q.AtE(0)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 't', r)

// 	r, err = q.AtE(1)
// 	assert.Equal(t, 'e', r)
// 	assert.Nil(t, err)

// 	r, err = q.AtE(2)
// 	assert.Equal(t, 's', r)
// 	assert.Nil(t, err)

// 	r, err = q.AtE(5)
// 	assert.Equal(t, rune(0), r)
// 	assert.Equal(t, "index out of Str bounds", err.Error())

// 	// nil
// 	{
// 		r, err := (*Str)(nil).AtE(2)
// 		assert.Equal(t, rune(0), r)
// 		assert.Equal(t, "Str is nil", err.Error())
// 	}
// }

// // AtA
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_AtA() {
// 	fmt.Println(A("foobar").AtA(-1))
// 	// Output: r
// }

// func TestStr_AtA(t *testing.T) {
// 	q := A("test")
// 	assert.Equal(t, "t", q.AtA(0))
// 	assert.Equal(t, "e", q.AtA(1))
// 	assert.Equal(t, "s", q.AtA(2))
// 	assert.Equal(t, "t", q.AtA(3))
// 	assert.Equal(t, "t", q.AtA(-1))
// 	assert.Equal(t, "s", q.AtA(-2))
// 	assert.Equal(t, "e", q.AtA(-3))
// 	assert.Equal(t, "t", q.AtA(-4))
// 	assert.Equal(t, "", q.AtA(5))
// }

// // AtAE
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_AtAE() {
// 	fmt.Println(A("foobar").AtAE(-1))
// 	// Output: r <nil>
// }

// func TestStr_AtAE(t *testing.T) {
// 	q := A("test")
// 	r, err := q.AtAE(0)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "t", r)

// 	r, err = q.AtAE(1)
// 	assert.Equal(t, "e", r)
// 	assert.Nil(t, err)

// 	r, err = q.AtAE(2)
// 	assert.Equal(t, "s", r)
// 	assert.Nil(t, err)

// 	r, err = q.AtAE(5)
// 	assert.Equal(t, "", r)
// 	assert.Equal(t, "index out of Str bounds", err.Error())

// 	// nil
// 	{
// 		r, err := (*Str)(nil).AtAE(2)
// 		assert.Equal(t, "", r)
// 		assert.Equal(t, "Str is nil", err.Error())
// 	}
// }

// // B
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_B() {
// 	fmt.Println(A("foobar").B())
// 	// Output: [102 111 111 98 97 114]
// }

// func TestStr_B(t *testing.T) {
// 	// string
// 	{
// 		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A("test").B())
// 	}

// 	// runes
// 	{
// 		assert.Equal(t, []byte{0x74}, A('t').B())
// 		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A([]rune("test")).B())
// 	}

// 	// bytes
// 	{
// 		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A([]byte("test")).B())
// 	}

// 	// ints
// 	{
// 		assert.Equal(t, []byte{0x31, 0x30}, A(10).B())
// 	}
// }

// // Clear
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_Clear() {
// 	fmt.Println(A("foobar").Clear())
// 	// Output:
// }

// func TestStr_Clear(t *testing.T) {
// 	assert.Equal(t, A(""), (*Str)(nil).Clear())
// 	assert.Equal(t, A(""), A("test").Clear())
// }

// // Contains
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_Contains() {
// 	fmt.Println(A("foobar").Contains("foo"))
// 	// Output: true
// }

// func TestStr_Contains(t *testing.T) {
// 	assert.True(t, A("test").Contains("tes"))
// 	assert.False(t, A("test").Contains("bob"))
// }

// // ContainsAny
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_ContainsAny() {
// 	fmt.Println(A("foobar").ContainsAny("bob"))
// 	// Output: true
// }

// func TestStr_ContainsAny(t *testing.T) {
// 	assert.True(t, A("test").ContainsAny("tes"))
// 	assert.False(t, A("test").ContainsAny("bob"))
// }

// // ContainsRune
// //--------------------------------------------------------------------------------------------------
// func ExampleStr_ContainsRune() {
// 	fmt.Println(A("foobar").ContainsRune('b'))
// 	// Output: true
// }

// func TestStr_ContainsRune(t *testing.T) {
// 	assert.True(t, A("test").ContainsRune('t'))
// 	assert.False(t, A("test").ContainsRune('b'))
// }

// // Copy
// //--------------------------------------------------------------------------------------------------
// func BenchmarkStr_Copy_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	dst := make([]string, len(src), len(src))
// 	copy(dst, src)
// }

// func BenchmarkStr_Copy_Str(t *testing.B) {
// 	slice := NewStr(RangeString(nines6))
// 	slice.Copy()
// }

// func ExampleStr_Copy() {
// 	str := A("0123456789")
// 	fmt.Println(str.Copy(1, 5))
// 	// Output: 12345
// }

// func TestStr_Copy(t *testing.T) {

// 	// nil or empty
// 	{
// 		var str *Str
// 		assert.Equal(t, A(""), str.Copy(0, -1))
// 		assert.Equal(t, A(""), A("0").Clear().Copy(0, -1))
// 	}

// 	// Test that the original is NOT modified when the slice is modified
// 	{
// 		original := A("123")
// 		result := original.Copy(0, -1)
// 		assert.Equal(t, A("123"), original)
// 		assert.Equal(t, A("123"), result)
// 		result.Set(0, "0")
// 		assert.Equal(t, A("1", "2", "3"), original)
// 		assert.Equal(t, A("0", "2", "3"), result)
// 	}

// 	// copy full array
// 	{
// 		assert.Equal(t, A(), A().Copy())
// 		assert.Equal(t, A(), A().Copy(0, -1))
// 		assert.Equal(t, A(), A().Copy(0, 1))
// 		assert.Equal(t, A(), A().Copy(0, 5))
// 		assert.Equal(t, A("1"), A("1").Copy())
// 		assert.Equal(t, A("1", "2", "3"), A("1", "2", "3").Copy())
// 		assert.Equal(t, A("1", "2", "3"), A("1", "2", "3").Copy(0, -1))
// 		assert.Equal(t, NewStr([]string{"1", "2", "3"}), NewStr([]string{"1", "2", "3"}).Copy())
// 		assert.Equal(t, NewStr([]string{"1", "2", "3"}), NewStr([]string{"1", "2", "3"}).Copy(0, -1))
// 	}

// 	// out of bounds should be moved in
// 	{
// 		assert.Equal(t, A("1"), A("1").Copy(0, 2))
// 		assert.Equal(t, A("1", "2", "3"), A("1", "2", "3").Copy(-6, 6))
// 	}

// 	// mutually exclusive
// 	{
// 		slice := A("1", "2", "3", "4")
// 		assert.Equal(t, A(), slice.Copy(2, -3))
// 		assert.Equal(t, A(), slice.Copy(0, -5))
// 		assert.Equal(t, A(), slice.Copy(4, -1))
// 		assert.Equal(t, A(), slice.Copy(6, -1))
// 		assert.Equal(t, A(), slice.Copy(3, -2))
// 	}

// 	// singles
// 	{
// 		slice := A("1", "2", "3", "4")
// 		assert.Equal(t, A("4"), slice.Copy(-1, -1))
// 		assert.Equal(t, A("3"), slice.Copy(-2, -2))
// 		assert.Equal(t, A("2"), slice.Copy(-3, -3))
// 		assert.Equal(t, A("1"), slice.Copy(0, 0))
// 		assert.Equal(t, A("1"), slice.Copy(-4, -4))
// 		assert.Equal(t, A("2"), slice.Copy(1, 1))
// 		assert.Equal(t, A("2"), slice.Copy(1, -3))
// 		assert.Equal(t, A("3"), slice.Copy(2, 2))
// 		assert.Equal(t, A("3"), slice.Copy(2, -2))
// 		assert.Equal(t, A("4"), slice.Copy(3, 3))
// 		assert.Equal(t, A("4"), slice.Copy(3, -1))
// 	}

// 	// grab all but first
// 	{
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3").Copy(1, -1))
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3").Copy(1, 2))
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3").Copy(-2, -1))
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3").Copy(-2, 2))
// 	}

// 	// grab all but last
// 	{
// 		assert.Equal(t, A("1", "2"), A("1", "2", "3").Copy(0, -2))
// 		assert.Equal(t, A("1", "2"), A("1", "2", "3").Copy(-3, -2))
// 		assert.Equal(t, A("1", "2"), A("1", "2", "3").Copy(-3, 1))
// 		assert.Equal(t, A("1", "2"), A("1", "2", "3").Copy(0, 1))
// 	}

// 	// grab middle
// 	{
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3", "4").Copy(1, -2))
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3", "4").Copy(-3, -2))
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3", "4").Copy(-3, 2))
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3", "4").Copy(1, 2))
// 	}

// 	// random
// 	{
// 		assert.Equal(t, A("1"), A("1", "2", "3").Copy(0, -3))
// 		assert.Equal(t, A("2", "3"), A("1", "2", "3").Copy(1, 2))
// 		assert.Equal(t, A("1", "2", "3"), A("1", "2", "3").Copy(0, 2))
// 	}
// }

// // func TestStr_Empty(t *testing.T) {
// // 	var empty *Str
// // 	assert.True(t, A("").Empty())
// // 	assert.True(t, empty.Empty())
// // 	assert.True(t, A("  ").Empty())
// // 	assert.True(t, A("\n").Empty())
// // 	assert.True(t, A("\t").Empty())
// // }

// // func TestStr_HasAnyPrefix(t *testing.T) {
// // 	assert.True(t, A("test").HasAnyPrefix("tes"))
// // 	assert.True(t, A("test").HasAnyPrefix("bob", "tes"))
// // 	assert.False(t, A("test").HasAnyPrefix("bob"))
// // }

// // func TestStr_HasAnySuffix(t *testing.T) {
// // 	assert.True(t, A("test").HasAnySuffix("est"))
// // 	assert.True(t, A("test").HasAnySuffix("bob", "est"))
// // 	assert.False(t, A("test").HasAnySuffix("bob"))
// // }

// // func TestStr_HasPrefix(t *testing.T) {
// // 	assert.True(t, A("test").HasPrefix("tes"))
// // }

// // func TestStr_HasSuffix(t *testing.T) {
// // 	assert.True(t, A("test").HasSuffix("est"))
// // }

// // func TestStr_Len(t *testing.T) {
// // 	assert.Equal(t, 0, A("").Len())
// // 	assert.Equal(t, 4, A("test").Len())
// // }

// // func TestStr_Replace(t *testing.T) {
// // 	assert.Equal(t, "tfoo", A("test").Replace("est", "foo").A())
// // 	assert.Equal(t, "foost", A("test").Replace("te", "foo").A())
// // 	assert.Equal(t, "foostfoo", A("testte").Replace("te", "foo").A())
// // }

// // func TestStr_SpaceLeft(t *testing.T) {
// // 	assert.Equal(t, "", A("").SpaceLeft().A())
// // 	assert.Equal(t, "  ", A("  bob").SpaceLeft().A())
// // 	assert.Equal(t, "\n", A("\nbob").SpaceLeft().A())
// // 	assert.Equal(t, " \t ", A(" \t bob").SpaceLeft().A())
// // }

// // func TestStr_SpaceRight(t *testing.T) {
// // 	assert.Equal(t, "", A("").SpaceRight().A())
// // 	assert.Equal(t, "  ", A("bob  ").SpaceRight().A())
// // 	assert.Equal(t, "\n", A("bob\n").SpaceRight().A())
// // 	assert.Equal(t, " \t ", A("bob \t ").SpaceRight().A())
// // }

// // // func TestStr_Split(t *testing.T) {
// // // 	assert.Equal(t, []string{""}, A("").Split(".").O())
// // // 	assert.Equal(t, []string{"1", "2"}, A("1.2").Split(".").O())
// // // 	assert.Equal(t, []string{"1", "2"}, A("1.2").Split(".").S())
// // // }

// // // func TestStr_SplitOn(t *testing.T) {
// // // 	{
// // // 		first, second := A("").SplitOn(":")
// // // 		assert.Equal(t, "", first)
// // // 		assert.Equal(t, "", second)
// // // 	}
// // // 	{
// // // 		first, second := A("foo").SplitOn(":")
// // // 		assert.Equal(t, "foo", first)
// // // 		assert.Equal(t, "", second)
// // // 	}
// // // 	{
// // // 		first, second := A("foo:").SplitOn(":")
// // // 		assert.Equal(t, "foo:", first)
// // // 		assert.Equal(t, "", second)
// // // 	}
// // // 	{
// // // 		first, second := A(":foo").SplitOn(":")
// // // 		assert.Equal(t, ":", first)
// // // 		assert.Equal(t, "foo", second)
// // // 	}
// // // 	{
// // // 		first, second := A("foo: bar").SplitOn(":")
// // // 		assert.Equal(t, "foo:", first)
// // // 		assert.Equal(t, " bar", second)
// // // 	}
// // // 	{
// // // 		first, second := A("foo: bar:frodo").SplitOn(":")
// // // 		assert.Equal(t, "foo:", first)
// // // 		assert.Equal(t, " bar:frodo", second)
// // // 	}
// // // }

// // // func TestStrSpaceLeft(t *testing.T) {
// // // 	assert.Equal(t, "", A("").SpaceLeft())
// // // 	assert.Equal(t, "", A("bob").SpaceLeft())
// // // 	assert.Equal(t, "  ", A("  bob").SpaceLeft())
// // // 	assert.Equal(t, "    ", A("    bob").SpaceLeft())
// // // 	assert.Equal(t, "\n", A("\nbob").SpaceLeft())
// // // 	assert.Equal(t, "\t", A("\tbob").SpaceLeft())
// // // }

// // // func TestStrTrimPrefix(t *testing.T) {
// // // 	assert.Equal(t, "test]", A("[test]").TrimPrefix("[").A())
// // // }

// // // func TestStrTrimSpace(t *testing.T) {
// // // 	{
// // // 		//Left
// // // 		assert.Equal(t, "bob", A("bob").TrimSpaceLeft().A())
// // // 		assert.Equal(t, "bob", A("  bob").TrimSpaceLeft().A())
// // // 		assert.Equal(t, "bob  ", A("  bob  ").TrimSpaceLeft().A())
// // // 		assert.Equal(t, 3, A("  bob").TrimSpaceLeft().Len())
// // // 	}
// // // 	{
// // // 		// Right
// // // 		assert.Equal(t, "bob", A("bob").TrimSpaceRight().A())
// // // 		assert.Equal(t, "bob", A("bob  ").TrimSpaceRight().A())
// // // 		assert.Equal(t, "  bob", A("  bob  ").TrimSpaceRight().A())
// // // 		assert.Equal(t, 3, A("bob  ").TrimSpaceRight().Len())
// // // 	}
// // // }

// // // func TestStrTrimSuffix(t *testing.T) {
// // // 	assert.Equal(t, "[test", A("[test]").TrimSuffix("]").A())
// // // }

// // // func TestYamlType(t *testing.T) {
// // // 	{
// // // 		// string
// // // 		assert.Equal(t, "test", A("\"test\"").YamlType())
// // // 		assert.Equal(t, "test", A("'test'").YamlType())
// // // 		assert.Equal(t, "1", A("\"1\"").YamlType())
// // // 		assert.Equal(t, "1", A("'1'").YamlType())
// // // 	}
// // // 	{
// // // 		// int
// // // 		assert.Equal(t, 1.0, A("1").YamlType())
// // // 		assert.Equal(t, 0.0, A("0").YamlType())
// // // 		assert.Equal(t, 25.0, A("25").YamlType())
// // // 	}
// // // 	{
// // // 		// bool
// // // 		assert.Equal(t, true, A("true").YamlType())
// // // 		assert.Equal(t, false, A("false").YamlType())
// // // 	}
// // // 	{
// // // 		// default
// // // 		assert.Equal(t, "True", A("True").YamlType())
// // // 	}
// // // }
