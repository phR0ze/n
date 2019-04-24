package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	// ints
	{
		assert.Equal(t, "10", NewStr(10).A())
	}

	// bytes
	{
		assert.Equal(t, "test", NewStr([]byte{0x74, 0x65, 0x73, 0x74}).A())
	}

	// Object
	{
		assert.Equal(t, "1", NewStr(Object{1}).A())
		assert.Equal(t, "12", NewStr([]Object{{1}, {2}}).A())
	}

	// runes
	{
		assert.Equal(t, "b", NewStr('b').A())
		assert.Equal(t, "test", NewStr([]rune("test")).A())
	}

	// Str
	{
		assert.Equal(t, "test", NewStr(NewStr("test")).A())
	}

	// string
	{
		assert.Equal(t, "test", NewStr("test").A())
		assert.Equal(t, "test1", NewStr([]string{"test", "1"}).O())
	}

	// reflection
	{
		assert.NotEqual(t, "test", NewStr(TestObj{"test"}).A())
	}
}

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

// All
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_All_Go(t *testing.B) {
	// All := func(list []string, x []string) bool {
	// 	for i := range x {
	// 		for j := range list {
	// 			if list[j] == x[i] {
	// 				return true
	// 			}
	// 		}
	// 	}
	// 	return false
	// }

	// // test here
	// src := RangeString(nines4)
	// for _, x := range src {
	// 	All(src, []string{x})
	// }
}

func BenchmarkStr_All_Slice(t *testing.B) {
	// src := RangeString(nines4)
	// slice := NewStr(src)
	// for i := range src {
	// 	slice.All(i)
	// }
}

func ExampleStr_All_contains() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.All("1", "2"))
	// Output: true
}

func TestStr_All(t *testing.T) {

	// empty
	var nilSlice *Str
	assert.False(t, nilSlice.All())
	assert.False(t, NewStrV().All())

	// single
	assert.True(t, NewStrV("2").All("2"))

	// invalid
	assert.False(t, NewStrV("12").Any(TestObj{"2"}))

	assert.True(t, NewStrV("123").All("2"))
	assert.False(t, NewStrV("123").All(4))
	assert.True(t, NewStrV("123").All("2", "3"))
	assert.False(t, NewStrV("123").All(4, 5))

	// Conversion
	assert.True(t, NewStrV("12").All(2))
	assert.True(t, NewStrV("12").All(Object{"2"}))
}

// AllS
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_AllS_Go(t *testing.B) {
	// any := func(list []string, x []string) bool {
	// 	for i := range x {
	// 		for j := range list {
	// 			if list[j] == x[i] {
	// 				return true
	// 			}
	// 		}
	// 	}
	// 	return false
	// }

	// // test here
	// src := RangeString(nines4)
	// for _, x := range src {
	// 	any(src, []string{x})
	// }
}

func BenchmarkStr_AllS_Slice(t *testing.B) {
	// src := RangeString(nines4)
	// slice := NewStr(src)
	// for _, x := range src {
	// 	slice.Any([]string{x})
	// }
}

func ExampleStr_AllS() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.AllS([]string{"1", "2", "3"}))
	// Output: true
}

func TestStr_AllS(t *testing.T) {
	// nil
	{
		var slice *Str
		assert.False(t, slice.AllS([]string{"1"}))
		assert.False(t, NewStrV("1").AllS(nil))
	}

	// byte
	{
		// byte
		assert.True(t, NewStrV("test").AllS(byte(0x74)))

		// []byte
		assert.True(t, NewStrV("test").AllS([]byte{0x74}))
		assert.True(t, NewStrV("test").AllS([]byte{0x74, 0x65}))
		assert.False(t, NewStrV("tbob").AllS([]byte{0x74, 0x65}))

		// *[]byte
		assert.True(t, NewStrV("test").AllS(&[]byte{0x74}))
		assert.True(t, NewStrV("test").AllS(&[]byte{0x74, 0x65}))
		assert.False(t, NewStrV("tbob").AllS(&[]byte{0x74, 0x65}))
	}

	// []int
	{
		assert.True(t, NewStrV("1", "2").AllS([]int{2}))
		assert.True(t, NewStrV("1", "2").AllS([]int{1, 2}))
		assert.False(t, NewStrV("1", "2").AllS([]int{1, 3}))
		assert.False(t, NewStrV("1", "2").AllS([]int{3, 4}))
	}

	// rune
	{
		// rune
		assert.True(t, NewStrV("test").AllS('t'))

		// []rune
		assert.True(t, NewStrV("123").AllS([]rune{'1'}))
		assert.True(t, NewStrV("123").AllS([]rune{'2', '3'}))
		assert.False(t, NewStrV("123").AllS([]rune{'1', '5'}))

		// *[]rune
		assert.True(t, NewStrV("123").AllS(&([]rune{'1'})))
		assert.True(t, NewStrV("123").AllS(&([]rune{'2', '3'})))
		assert.False(t, NewStrV("123").AllS(&([]rune{'1', '5'})))
	}

	// []string
	{
		assert.True(t, NewStrV("123").AllS([]string{"1"}))
		assert.True(t, NewStrV("123").AllS([]string{"2", "3"}))
		assert.False(t, NewStrV("123").AllS([]string{"1", "5"}))

		// *[]string
		assert.True(t, NewStrV("123").AllS(&([]string{"1"})))
		assert.True(t, NewStrV("123").AllS(&([]string{"2", "3"})))
		assert.False(t, NewStrV("123").AllS(&([]string{"1", "5"})))
	}

	// StringSlice
	{
		assert.True(t, NewStrV("123").AllS(NewStringSliceV("1")))
		assert.True(t, NewStrV("123").AllS(NewStringSliceV("2", "3")))
		assert.False(t, NewStrV("123").AllS(NewStringSliceV("1", "5")))
	}

	// Str
	{
		assert.True(t, NewStrV("123").AllS([]Str{Str("1"), Str("2")}))
		assert.True(t, NewStrV("123").AllS(&[]Str{Str("1"), Str("2")}))
		assert.False(t, NewStrV("123").AllS([]Str{Str("1"), Str("5")}))

		// *Str
		assert.True(t, NewStrV("123").AllS([]*Str{A("1"), A("2")}))
		assert.False(t, NewStrV("123").AllS([]*Str{A("1"), A("5")}))
	}

	// invalid types
	assert.False(t, NewStrV("1", "2").AllS(nil))
	assert.False(t, NewStrV("1", "2").AllS((*[]string)(nil)))
	assert.False(t, NewStrV("1", "2").AllS((*Str)(nil)))
}

// Any
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Any_Go(t *testing.B) {
	// any := func(list []string, x []string) bool {
	// 	for i := range x {
	// 		for j := range list {
	// 			if list[j] == x[i] {
	// 				return true
	// 			}
	// 		}
	// 	}
	// 	return false
	// }

	// // test here
	// src := RangeString(nines4)
	// for _, x := range src {
	// 	any(src, []string{x})
	// }
}

func BenchmarkStr_Any_Slice(t *testing.B) {
	// src := RangeString(nines4)
	// slice := NewStr(src)
	// for i := range src {
	// 	slice.Any(i)
	// }
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
	assert.False(t, NewStrV("12").Any(TestObj{"2"}))

	assert.True(t, NewStrV("123").Any("2"))
	assert.False(t, NewStrV("123").Any(4))
	assert.True(t, NewStrV("123").Any(4, "3"))
	assert.False(t, NewStrV("123").Any(4, 5))

	// Conversion
	assert.True(t, NewStrV("12").Any(Object{"2"}))
}

// AnyS
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_AnyS_Go(t *testing.B) {
	// any := func(list []string, x []string) bool {
	// 	for i := range x {
	// 		for j := range list {
	// 			if list[j] == x[i] {
	// 				return true
	// 			}
	// 		}
	// 	}
	// 	return false
	// }

	// // test here
	// src := RangeString(nines4)
	// for _, x := range src {
	// 	any(src, []string{x})
	// }
}

func BenchmarkStr_AnyS_Slice(t *testing.B) {
	// src := RangeString(nines4)
	// slice := NewStr(src)
	// for _, x := range src {
	// 	slice.Any([]string{x})
	// }
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

	// byte
	{
		// byte
		assert.True(t, NewStrV("test").AnyS(byte(0x74)))

		// []byte
		assert.True(t, NewStrV("test").AnyS([]byte{0x74}))
		assert.True(t, NewStrV("bobe").AnyS([]byte{0x74, 0x65}))
		assert.False(t, NewStrV("bob").AnyS([]byte{0x74, 0x65}))

		// *[]byte
		assert.True(t, NewStrV("test").AnyS(&[]byte{0x74}))
		assert.True(t, NewStrV("bobe").AnyS(&[]byte{0x74, 0x65}))
		assert.False(t, NewStrV("bob").AnyS(&[]byte{0x74, 0x65}))
	}

	// Char
	{
		assert.True(t, NewStrV("123").AnyS([]Char{'1', '2'}))
		assert.True(t, NewStrV("123").AnyS(&[]Char{'1', '2'}))
		assert.False(t, NewStrV("123").AnyS([]Char{'4', '5'}))
	}

	// ints
	{
		assert.True(t, NewStrV("1", "2").AnyS([]int{2}))
		assert.True(t, NewStrV("1", "2").AnyS([]int8{2}))
		assert.True(t, NewStrV("1", "2").AnyS([]int16{2}))
		assert.True(t, NewStrV("1", "2").AnyS([]int64{2}))
		assert.True(t, NewStrV("1", "2").AnyS([]int{1, 2}))
		assert.True(t, NewStrV("1", "2").AnyS([]int{1, 3}))
		assert.False(t, NewStrV("1", "2").AnyS([]int{3, 4}))
	}

	// rune
	{
		// rune
		assert.True(t, NewStrV("test").AnyS('t'))

		// []rune
		assert.True(t, NewStrV("123").AnyS([]rune{'1'}))
		assert.True(t, NewStrV("123").AnyS([]rune{'4', '3'}))
		assert.False(t, NewStrV("123").AnyS([]rune{'4', '5'}))

		// *[]rune
		assert.True(t, NewStrV("123").AnyS(&([]rune{'1'})))
		assert.True(t, NewStrV("123").AnyS(&([]rune{'4', '3'})))
		assert.False(t, NewStrV("123").AnyS(&([]rune{'4', '5'})))
	}

	// string
	{
		// []string
		assert.True(t, NewStrV("123").AnyS([]string{"1"}))
		assert.True(t, NewStrV("123").AnyS([]string{"4", "3"}))
		assert.False(t, NewStrV("123").AnyS([]string{"4", "5"}))

		// *[]string
		assert.True(t, NewStrV("123").AnyS(&([]string{"1"})))
		assert.True(t, NewStrV("123").AnyS(&([]string{"4", "3"})))
		assert.False(t, NewStrV("123").AnyS(&([]string{"4", "5"})))
	}

	// StringSlice
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

	// uints
	{
		assert.True(t, NewStrV("1", "2").AnyS([]uint{2}))
		assert.True(t, NewStrV("1", "2").AnyS([]uint16{2}))
		assert.True(t, NewStrV("1", "2").AnyS([]uint32{2}))
		assert.True(t, NewStrV("1", "2").AnyS([]uint64{2}))
	}

	// invalid types
	assert.False(t, NewStrV("1", "2").AnyS(nil))
	assert.False(t, NewStrV("1", "2").AnyS((*[]string)(nil)))
	assert.False(t, NewStrV("1", "2").AnyS((*Str)(nil)))
}

// AnyW
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_AnyW_Go(t *testing.B) {
	// src := RangeString(nines5)
	// for _, x := range src {
	// 	if x == string(nines4) {
	// 		break
	// 	}
	// }
}

func BenchmarkStr_AnyW_Slice(t *testing.B) {
	// src := RangeString(nines5)
	// NewStr(src).AnyW(func(x O) bool {
	// 	return ExB(x.(string) == string(nines4))
	// })
}

func ExampleStr_AnyW() {
	slice := NewStr("123")
	fmt.Println(slice.AnyW(func(x O) bool {
		return ExB(x.(Char) == '2')
	}))
	// Output: true
}

func TestStr_AnyW(t *testing.T) {

	// empty
	var slice *Str
	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(Char) > '0') }))
	assert.False(t, NewStrV().AnyW(func(x O) bool { return ExB(x.(Char) > '0') }))

	// single
	assert.True(t, NewStr("2").AnyW(func(x O) bool { return ExB(x.(Char) > '0') }))
	assert.True(t, NewStr("12").AnyW(func(x O) bool { return ExB(x.(Char) == '2') }))
	assert.True(t, NewStr("123").AnyW(func(x O) bool { return ExB(x.(Char) == '4' || x.(Char) == '3') }))
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

	// Conversion
	{
		assert.Equal(t, "12", NewStrV(1).Append(Object{2}).A())
		assert.Equal(t, "12", NewStrV(1).Append(2).A())
		assert.Equal(t, "true2", NewStrV().Append(true).Append(Char('2')).A())
	}
}

// AppendV
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_AppendV_Go(t *testing.B) {
	// src := []string{}
	// src = append(src, RangeString(nines6)...)
}

func BenchmarkStr_AppendV_Slice(t *testing.B) {
	// n := NewStrV()
	// new := rangeO(0, nines6)
	// n.AppendV(new...)
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

	// Conversion
	{
		assert.Equal(t, NewStrV("0", "1"), NewStrV().AppendV(Object{0}, Object{1}))
		assert.Equal(t, NewStrV("0", "1"), NewStrV().AppendV(0, 1))
		assert.Equal(t, NewStrV("false", "true"), NewStrV().AppendV(false, true))
	}
}

// Ascii
//--------------------------------------------------------------------------------------------------
func ExampleStr_Ascii() {
	fmt.Println(A("2�gspu�data").Ascii().A())
	// Output: 2 gspu data
}

func TestStr_Ascii(t *testing.T) {
	assert.Equal(t, A("2 gspu data gspm data"), A("2�gspu�data�gspm�data").Ascii())
}

// AsciiA
//--------------------------------------------------------------------------------------------------
func ExampleStr_AsciiA() {
	fmt.Println(A("2�gspu�data").AsciiA())
	// Output: 2 gspu data
}

func TestStr_AsciiA(t *testing.T) {
	assert.Equal(t, "2 gspu data gspm data", A("2�gspu�data�gspm�data").AsciiA())
}

// AsciiOnly
//--------------------------------------------------------------------------------------------------
func ExampleStr_AsciiOnly() {
	fmt.Println(A("foo").AsciiOnly())
	// Output: true
}

func TestStr_AsciiOnly(t *testing.T) {
	assert.Equal(t, true, A("foobar").AsciiOnly())
	assert.Equal(t, false, A("2�gspu�data�gspm�data").AsciiOnly())
}

// At
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_At_Go(t *testing.B) {
	// src := RangeString(nines6)
	// for _, x := range src {
	// 	assert.IsType(t, 0, x)
	// }
}

func BenchmarkStr_At_Slice(t *testing.B) {
	// src := RangeString(nines6)
	// slice := NewStr(src)
	// for i := 0; i < len(src); i++ {
	// 	_, ok := (slice.At(i).O()).(string)
	// 	assert.True(t, ok)
	// }
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
		assert.Equal(t, ToChar("4"), slice.At(3).O())
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

// B
//--------------------------------------------------------------------------------------------------
func ExampleStr_B() {
	fmt.Println(A("foobar").B())
	// Output: [102 111 111 98 97 114]
}

func TestStr_B(t *testing.T) {
	// string
	{
		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A("test").B())
	}

	// runes
	{
		assert.Equal(t, "t", string(A('t').B()))
		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A([]rune("test")).B())
	}

	// bytes
	{
		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A([]byte("test")).B())
	}

	// ints
	{
		assert.Equal(t, []byte{0x31, 0x30}, A(10).B())
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

	// Conversion
	{
		assert.Equal(t, NewStrV("0", "1"), NewStrV().Concat([]Object{{0}, {1}}))
		assert.Equal(t, NewStrV("0", "1"), NewStrV().Concat([]int{0, 1}))
		assert.Equal(t, NewStrV("false", "true"), NewStrV().Concat([]bool{false, true}))

		slice := NewStrV(Object{1})
		concated := slice.Concat([]int64{2, 3})
		assert.Equal(t, NewStrV("1", "4"), slice.Append(Char('4')))
		assert.Equal(t, NewStrV("1", "2", "3"), concated)
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

	// Conversion
	{
		slice := NewStrV(Object{1})
		concated := slice.ConcatM([]Object{{2}, {3}})
		assert.Equal(t, NewStrV("1", "2", "3", "4"), slice.Append(Char('4')))
		assert.Equal(t, NewStrV("1", "2", "3", "4"), concated)
	}
}

// Contains
//--------------------------------------------------------------------------------------------------
func ExampleStr_Contains() {
	fmt.Println(A("foobar").Contains("foo"))
	// Output: true
}

func TestStr_Contains(t *testing.T) {
	assert.True(t, A("test").Contains("tes"))
	assert.False(t, A("test").Contains("bob"))
}

// ContainsAny
//--------------------------------------------------------------------------------------------------
func ExampleStr_ContainsAny() {
	fmt.Println(A("foobar").ContainsAny("bob"))
	// Output: true
}

func TestStr_ContainsAny(t *testing.T) {
	assert.True(t, A("test").ContainsAny("tes"))
	assert.False(t, A("test").ContainsAny("bob"))
}

// ContainsRune
//--------------------------------------------------------------------------------------------------
func ExampleStr_ContainsRune() {
	fmt.Println(A("foobar").ContainsRune('b'))
	// Output: true
}

func TestStr_ContainsRune(t *testing.T) {
	assert.True(t, A("test").ContainsRune('t'))
	assert.False(t, A("test").ContainsRune('b'))
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

// Count
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Count_Go(t *testing.B) {
	src := RangeString(nines5)
	for _, x := range src {
		if x == string(nines4) {
			break
		}
	}
}

func BenchmarkStr_Count_Slice(t *testing.B) {
	src := RangeString(nines5)
	NewStr(src).Count(nines4)
}

func ExampleStr_Count() {
	slice := NewStrV("1", "2", "2")
	fmt.Println(slice.Count("2"))
	// Output: 2
}

func TestStr_Count(t *testing.T) {

	// empty
	var slice *Str
	assert.Equal(t, 0, slice.Count(0))
	assert.Equal(t, 0, NewStrV().Count(0))

	assert.Equal(t, 1, NewStr("23").Count("2"))
	assert.Equal(t, 2, NewStr("122").Count("2"))
	assert.Equal(t, 4, NewStr("44344").Count("4"))
	assert.Equal(t, 3, NewStr("32335").Count("3"))
	assert.Equal(t, 1, NewStr("123").Count("3"))
}

// CountW
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_CountW_Go(t *testing.B) {
	src := RangeString(nines5)
	for _, x := range src {
		if x == string(nines4) {
			break
		}
	}
}

func BenchmarkStr_CountW_Slice(t *testing.B) {
	// src := RangeString(nines5)
	// NewStr(src).CountW(func(x O) bool {
	// 	return ExB(x.(Char) == string(nines4))
	// })
}

func ExampleStr_CountW() {
	slice := NewStrV("1", "2", "2")
	fmt.Println(slice.CountW(func(x O) bool {
		return ExB(x.(Char) == '2')
	}))
	// Output: 2
}

func TestStr_CountW(t *testing.T) {

	// empty
	var slice *Str
	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(Char) > '0') }))
	assert.Equal(t, 0, NewStrV().CountW(func(x O) bool { return ExB(x.(Char) > '0') }))

	assert.Equal(t, 1, NewStrV("2", "3").CountW(func(x O) bool { return ExB(x.(Char) > '2') }))
	assert.Equal(t, 1, NewStrV("1", "2").CountW(func(x O) bool { return ExB(x.(Char) == '2') }))
	assert.Equal(t, 1, NewStrV("1", "2", "3").CountW(func(x O) bool { return ExB(x.(Char) == '4' || x.(Char) == '3') }))
}

// Drop
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Drop_Go(t *testing.B) {
	// src := RangeString(nines7)
	// for len(src) > 11 {
	// 	i := 1
	// 	n := 10
	// 	if i+n < len(src) {
	// 		src = append(src[:i], src[i+n:]...)
	// 	} else {
	// 		src = src[:i]
	// 	}
	// }
}

func BenchmarkStr_Drop_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines7))
	// for slice.Len() > 1 {
	// 	slice.Drop(1, 10)
	// }
}

func ExampleStr_Drop() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.Drop(0, 1))
	// Output: 3
}

func TestStr_Drop(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.Drop(0, 1))
	}

	// invalid
	assert.Equal(t, NewStrV("1", "2", "3", "4"), NewStrV("1", "2", "3", "4").Drop(1))
	assert.Equal(t, NewStrV("1", "2", "3", "4"), NewStrV("1", "2", "3", "4").Drop(4, 4))

	// drop 1
	assert.Equal(t, NewStrV("2", "3", "4"), NewStrV("1", "2", "3", "4").Drop(0, 0))
	assert.Equal(t, NewStrV("1", "3", "4"), NewStrV("1", "2", "3", "4").Drop(1, 1))
	assert.Equal(t, NewStrV("1", "2", "4"), NewStrV("1", "2", "3", "4").Drop(2, 2))
	assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3", "4").Drop(3, 3))
	assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3", "4").Drop(-1, -1))
	assert.Equal(t, NewStrV("1", "2", "4"), NewStrV("1", "2", "3", "4").Drop(-2, -2))
	assert.Equal(t, NewStrV("1", "3", "4"), NewStrV("1", "2", "3", "4").Drop(-3, -3))
	assert.Equal(t, NewStrV("2", "3", "4"), NewStrV("1", "2", "3", "4").Drop(-4, -4))

	// drop 2
	assert.Equal(t, NewStrV("3", "4"), NewStrV("1", "2", "3", "4").Drop(0, 1))
	assert.Equal(t, NewStrV("1", "4"), NewStrV("1", "2", "3", "4").Drop(1, 2))
	assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3", "4").Drop(2, 3))
	assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3", "4").Drop(-2, -1))
	assert.Equal(t, NewStrV("1", "4"), NewStrV("1", "2", "3", "4").Drop(-3, -2))
	assert.Equal(t, NewStrV("3", "4"), NewStrV("1", "2", "3", "4").Drop(-4, -3))

	// drop 3
	assert.Equal(t, NewStrV("4"), NewStrV("1", "2", "3", "4").Drop(0, 2))
	assert.Equal(t, NewStrV("1"), NewStrV("1", "2", "3", "4").Drop(-3, -1))

	// drop everything and beyond
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3", "4").Drop())
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3", "4").Drop(0, 3))
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3", "4").Drop(0, -1))
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3", "4").Drop(-4, -1))
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3", "4").Drop(-6, -1))
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3", "4").Drop(0, 10))

	// move index within bounds
	assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3", "4").Drop(3, 4))
	assert.Equal(t, NewStrV("2", "3", "4"), NewStrV("1", "2", "3", "4").Drop(-5, 0))
}

// DropAt
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_DropAt_Go(t *testing.B) {
	// src := RangeString(nines5)
	// index := Range(0, nines5)
	// for i := range index {
	// 	if i+1 < len(src) {
	// 		src = append(src[:i], src[i+1:]...)
	// 	} else if i >= 0 && i < len(src) {
	// 		src = src[:i]
	// 	}
	// }
}

func BenchmarkStr_DropAt_Slice(t *testing.B) {
	// index := Range(0, nines5)
	// slice := NewStr(RangeString(nines5))
	// for i := range index {
	// 	slice.DropAt(i)
	// }
}

func ExampleStr_DropAt() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.DropAt(1))
	// Output: 13
}

func TestStr_DropAt(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.DropAt(0))
	}

	// drop all and more
	{
		slice := NewStrV("0", "1", "2")
		assert.Equal(t, NewStrV("0", "1"), slice.DropAt(-1))
		assert.Equal(t, NewStrV("0"), slice.DropAt(-1))
		assert.Equal(t, NewStrV(), slice.DropAt(-1))
		assert.Equal(t, NewStrV(), slice.DropAt(-1))
	}

	// drop invalid
	assert.Equal(t, NewStrV("0", "1", "2"), NewStrV("0", "1", "2").DropAt(3))
	assert.Equal(t, NewStrV("0", "1", "2"), NewStrV("0", "1", "2").DropAt(-4))

	// drop last
	assert.Equal(t, NewStrV("0", "1"), NewStrV("0", "1", "2").DropAt(2))
	assert.Equal(t, NewStrV("0", "1"), NewStrV("0", "1", "2").DropAt(-1))

	// drop middle
	assert.Equal(t, NewStrV("0", "2"), NewStrV("0", "1", "2").DropAt(1))
	assert.Equal(t, NewStrV("0", "2"), NewStrV("0", "1", "2").DropAt(-2))

	// drop first
	assert.Equal(t, NewStrV("1", "2"), NewStrV("0", "1", "2").DropAt(0))
	assert.Equal(t, NewStrV("1", "2"), NewStrV("0", "1", "2").DropAt(-3))
}

// DropFirst
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_DropFirst_Go(t *testing.B) {
	// src := RangeString(nines7)
	// for len(src) > 1 {
	// 	src = src[1:]
	// }
}

func BenchmarkStr_DropFirst_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines7))
	// for slice.Len() > 0 {
	// 	slice.DropFirst()
	// }
}

func ExampleStr_DropFirst() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.DropFirst())
	// Output: 23
}

func TestStr_DropFirst(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.DropFirst())
	}

	// drop all and beyond
	{
		slice := NewStrV("1", "2", "3")
		assert.Equal(t, NewStrV("2", "3"), slice.DropFirst())
		assert.Equal(t, NewStrV("3"), slice.DropFirst())
		assert.Equal(t, NewStrV(), slice.DropFirst())
		assert.Equal(t, NewStrV(), slice.DropFirst())
	}
}

// DropFirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_DropFirstN_Go(t *testing.B) {
	// src := RangeString(nines7)
	// for len(src) > 10 {
	// 	src = src[10:]
	// }
}

func BenchmarkStr_DropFirstN_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines7))
	// for slice.Len() > 0 {
	// 	slice.DropFirstN(10)
	// }
}

func ExampleStr_DropFirstN() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.DropFirstN(2))
	// Output: 3
}

func TestStr_DropFirstN(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.DropFirstN(1))
	}

	// negative value
	assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3").DropFirstN(-1))

	// drop none
	assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3").DropFirstN(0))

	// drop 1
	assert.Equal(t, NewStrV("2", "3"), NewStrV("1", "2", "3").DropFirstN(1))

	// drop 2
	assert.Equal(t, NewStrV("3"), NewStrV("1", "2", "3").DropFirstN(2))

	// drop 3
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3").DropFirstN(3))

	// drop beyond
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3").DropFirstN(4))
}

// DropLast
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_DropLast_Go(t *testing.B) {
	// src := RangeString(nines7)
	// for len(src) > 1 {
	// 	src = src[1:]
	// }
}

func BenchmarkStr_DropLast_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines7))
	// for slice.Len() > 0 {
	// 	slice.DropLast()
	// }
}

func ExampleStr_DropLast() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.DropLast())
	// Output: 12
}

func TestStr_DropLast(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.DropLast())
	}

	// negative value
	assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3").DropLastN(-1))

	slice := NewStrV("1", "2", "3")
	assert.Equal(t, NewStrV("1", "2"), slice.DropLast())
	assert.Equal(t, NewStrV("1"), slice.DropLast())
	assert.Equal(t, NewStrV(), slice.DropLast())
	assert.Equal(t, NewStrV(), slice.DropLast())
}

// DropLastN
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_DropLastN_Go(t *testing.B) {
	// src := RangeString(nines7)
	// for len(src) > 10 {
	// 	src = src[10:]
	// }
}

func BenchmarkStr_DropLastN_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines7))
	// for slice.Len() > 0 {
	// 	slice.DropLastN(10)
	// }
}

func ExampleStr_DropLastN() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.DropLastN(2))
	// Output: 1
}

func TestStr_DropLastN(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.DropLastN(1))
	}

	// drop none
	assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3").DropLastN(0))

	// drop 1
	assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3").DropLastN(1))

	// drop 2
	assert.Equal(t, NewStrV("1"), NewStrV("1", "2", "3").DropLastN(2))

	// drop 3
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3").DropLastN(3))

	// drop beyond
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3").DropLastN(4))
}

// DropW
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_DropW_Go(t *testing.B) {
	// src := RangeString(nines5)
	// l := len(src)
	// for i := 0; i < l; i++ {
	// 	if Obj(src[i]).ToInt()%2 == 0 {
	// 		if i+1 < l {
	// 			src = append(src[:i], src[i+1:]...)
	// 		} else if i >= 0 && i < l {
	// 			src = src[:i]
	// 		}
	// 		l--
	// 		i--
	// 	}
	// }
}

func BenchmarkStr_DropW_Slice(t *testing.B) {
	slice := NewStr(RangeString(nines5))
	slice.DropW(func(x O) bool {
		return ExB(Obj(x).ToInt()%2 == 0)
	})
}

func ExampleStr_DropW() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.DropW(func(x O) bool {
		return ExB(Obj(x).ToInt()%2 == 0)
	}))
	// Output: 13
}

func TestStr_DropW(t *testing.T) {

	// drop all odd values
	{
		slice := NewStrV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		slice.DropW(func(x O) bool {
			return ExB(ToInt(x)%2 != 0)
		})
		assert.Equal(t, "2468", slice.A())
	}

	// drop all even values
	{
		slice := NewStrV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		slice.DropW(func(x O) bool {
			return ExB(ToInt(x)%2 == 0)
		})
		assert.Equal(t, "13579", slice.A())
	}
}

// Each
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Each_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_Each_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).Each(func(x O) {
	// 	assert.IsType(t, "", x)
	// })
}

func ExampleStr_Each() {
	NewStrV("1", "2", "3").Each(func(x O) {
		fmt.Printf("%v", x.(*Char).A())
	})
	// Output: 123
}

func TestStr_Each(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.Each(func(x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").Each(func(x O) {
			results = append(results, x.(*Char).A())
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}
}

// EachE
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_EachE_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "0", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_EachE_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).EachE(func(x O) error {
	// 	assert.IsType(t, "", x)
	// 	return nil
	// })
}

func ExampleStr_EachE() {
	NewStrV("1", "2", "3").EachE(func(x O) error {
		fmt.Printf("%v", x.(*Char))
		return nil
	})
	// Output: 123
}

func TestStr_EachE(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.EachE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").EachE(func(x O) error {
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStrV("1", "2", "3").EachE(func(x O) error {
			if x.(*Char).G() == '3' {
				return ErrBreak
			}
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// EachI
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_EachI_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_EachI_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).EachI(func(i int, x O) {
	// 	assert.IsType(t, "", x)
	// })
}

func ExampleStr_EachI() {
	NewStrV("1", "2", "3").EachI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x.(*Char))
	})
	// Output: 0:11:22:3
}

func TestStr_EachI(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.EachI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").EachI(func(i int, x O) {
			results = append(results, x.(*Char).A())
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}
}

// EachIE
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_EachIE_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_EachIE_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).EachIE(func(i int, x O) error {
	// 	assert.IsType(t, "", x)
	// 	return nil
	// })
}

func ExampleStr_EachIE() {
	NewStrV("1", "2", "3").EachIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x.(*Char))
		return nil
	})
	// Output: 0:11:22:3
}

func TestStr_EachIE(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.EachIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").EachIE(func(i int, x O) error {
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStrV("1", "2", "3").EachIE(func(i int, x O) error {
			if i == 2 {
				return ErrBreak
			}
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// EachR
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_EachR_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_EachR_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).EachR(func(x O) {
	// 	assert.IsType(t, "", x)
	// })
}

func ExampleStr_EachR() {
	NewStrV("1", "2", "3").EachR(func(x O) {
		fmt.Printf("%v", x.(*Char))
	})
	// Output: 321
}

func TestStr_EachR(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.EachR(func(x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").EachR(func(x O) {
			results = append(results, x.(*Char).A())
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}
}

// EachRE
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_EachRE_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_EachRE_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).EachRE(func(x O) error {
	// 	assert.IsType(t, "", x)
	// 	return nil
	// })
}

func ExampleStr_EachRE() {
	NewStrV("1", "2", "3").EachRE(func(x O) error {
		fmt.Printf("%v", x.(*Char))
		return nil
	})
	// Output: 321
}

func TestStr_EachRE(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.EachRE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").EachRE(func(x O) error {
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStrV("1", "2", "3").EachRE(func(x O) error {
			if x.(*Char).A() == "1" {
				return ErrBreak
			}
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"3", "2"}, results)
	}
}

// EachRI
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_EachRI_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_EachRI_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).EachRI(func(i int, x O) {
	// 	assert.IsType(t, "", x)
	// })
}

func ExampleStr_EachRI() {
	NewStrV("1", "2", "3").EachRI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x.(*Char))
	})
	// Output: 2:31:20:1
}

func TestStr_EachRI(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.EachRI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").EachRI(func(i int, x O) {
			results = append(results, x.(*Char).A())
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}
}

// EachRIE
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_EachRIE_Go(t *testing.B) {
	// action := func(x interface{}) {
	// 	assert.IsType(t, "", x)
	// }
	// for _, x := range RangeString(nines6) {
	// 	action(x)
	// }
}

func BenchmarkStr_EachRIE_Slice(t *testing.B) {
	// NewStr(RangeString(nines6)).EachRIE(func(i int, x O) error {
	// 	assert.IsType(t, "", x)
	// 	return nil
	// })
}

func ExampleStr_EachRIE() {
	NewStrV("1", "2", "3").EachRIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x.(*Char))
		return nil
	})
	// Output: 2:31:20:1
}

func TestStr_EachRIE(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		slice.EachRIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewStrV("1", "2", "3").EachRIE(func(i int, x O) error {
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"3", "2", "1"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewStrV("1", "2", "3").EachRIE(func(i int, x O) error {
			if i == 0 {
				return ErrBreak
			}
			results = append(results, x.(*Char).A())
			return nil
		})
		assert.Equal(t, []string{"3", "2"}, results)
	}
}

// Empty
//--------------------------------------------------------------------------------------------------
func ExampleStr_Empty() {
	fmt.Println(NewStrV().Empty())
	// Output: true
}

func TestStr_Empty(t *testing.T) {

	// nil or empty
	{
		var nilSlice *Str
		assert.Equal(t, true, nilSlice.Empty())
	}

	assert.Equal(t, true, NewStrV().Empty())
	assert.Equal(t, false, NewStrV("1").Empty())
	assert.Equal(t, false, NewStrV("1", "2", "3").Empty())
	assert.Equal(t, false, NewStrV("1").Empty())
	assert.Equal(t, false, NewStr([]string{"1", "2", "3"}).Empty())
}

// First
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_First_Go(t *testing.B) {
	// src := RangeString(nines7)
	// for len(src) > 1 {
	// 	_ = src[0]
	// 	src = src[1:]
	// }
}

func BenchmarkStr_First_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines7))
	// for slice.Len() > 0 {
	// 	slice.First()
	// 	slice.DropFirst()
	// }
}

func ExampleStr_First() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.First())
	// Output: 1
}

func TestStr_First(t *testing.T) {
	// invalid
	assert.Equal(t, Obj(nil), NewStrV().First())

	// int
	assert.Equal(t, ToChar("2"), NewStrV("2", "3").First().O())
	assert.Equal(t, ToChar("3"), NewStrV("3", "2").First().O())
	assert.Equal(t, ToChar("1"), NewStrV("1", "3", "2").First().O())
}

// FirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_FirstN_Go(t *testing.B) {
	// src := RangeString(nines7)
	// _ = src[0:10]
}

func BenchmarkStr_FirstN_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines7))
	// slice.FirstN(10)
}

func ExampleStr_FirstN() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.FirstN(2))
	// Output: 12
}

func TestStr_FirstN(t *testing.T) {

	// nil or empty
	{
		var slice *Str
		assert.Equal(t, NewStrV(), slice.FirstN(1))
		assert.Equal(t, NewStrV(), slice.FirstN(-1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewStrV("1", "2", "3")
		result := original.FirstN(2).Set(0, "0")
		assert.Equal(t, NewStrV("0", "2", "3"), original)
		assert.Equal(t, NewStrV("0", "2"), result)
	}

	// Get none
	assert.Equal(t, NewStrV(), NewStrV("1", "2", "3").FirstN(0))

	// slice full array includeing out of bounds
	assert.Equal(t, NewStrV(), NewStrV().FirstN(1))
	assert.Equal(t, NewStrV(), NewStrV().FirstN(10))
	assert.Equal(t, NewStrV("1", "2", "3"), NewStrV("1", "2", "3").FirstN(10))
	assert.Equal(t, NewStr([]string{"1", "2", "3"}), NewStr([]string{"1", "2", "3"}).FirstN(10))

	// grab a few diff
	assert.Equal(t, NewStrV("1"), NewStrV("1", "2", "3").FirstN(1))
	assert.Equal(t, NewStrV("1", "2"), NewStrV("1", "2", "3").FirstN(2))
}

// G
//--------------------------------------------------------------------------------------------------
func ExampleStr_G() {
	fmt.Println(NewStrV("1", "2", "3").G())
	// Output: 123
}

func TestStr_G(t *testing.T) {
	assert.IsType(t, "", NewStrV().G())
	assert.IsType(t, "123", NewStrV("1", "2", "3").G())
}

// Generic
//--------------------------------------------------------------------------------------------------
func ExampleStr_Generic() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.Generic())
	// Output: false
}

// Index
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Index_Go(t *testing.B) {
	// for _, x := range RangeString(nines5) {
	// 	if x == string(nines4) {
	// 		break
	// 	}
	// }
}

func BenchmarkStr_Index_Slice(t *testing.B) {
	// slice := NewStr(RangeString(nines5))
	// slice.Index(nines4)
}

func ExampleStr_Index() {
	slice := NewStrV("1", "2", "3")
	fmt.Println(slice.Index("2"))
	// Output: 1
}

func TestStr_Index(t *testing.T) {

	// empty
	var slice *Str
	assert.Equal(t, -1, slice.Index("2"))
	assert.Equal(t, -1, NewStrV().Index("1"))

	assert.Equal(t, 0, NewStrV("1", "2", "3").Index("1"))
	assert.Equal(t, 1, NewStrV("1", "2", "3").Index("2"))
	assert.Equal(t, 2, NewStrV("1", "2", "3").Index("3"))
	assert.Equal(t, -1, NewStrV("1", "2", "3").Index("4"))
	assert.Equal(t, -1, NewStrV("1", "2", "3").Index("5"))

	// Conversion
	{
		assert.Equal(t, 1, NewStrV("1", "2", "3").Index(Object{2}))
		assert.Equal(t, 1, NewStrV("1", "2", "3").Index("2"))
		assert.Equal(t, 0, NewStrV("1", "2", "3").Index(true))
		assert.Equal(t, -1, NewStrV("1", "2", "3").Index(false))
		assert.Equal(t, 2, NewStrV("1", "2", "3").Index(Char('3')))
	}
}

// Insert
//--------------------------------------------------------------------------------------------------
func BenchmarkStr_Insert_Go(t *testing.B) {
	src := []string{}
	for _, x := range RangeString(nines6) {
		src = append(src, x)
		copy(src[1:], src[1:])
		src[0] = x
	}
}

func BenchmarkStr_Insert_Slice(t *testing.B) {
	slice := NewStrV()
	for x := range RangeString(nines6) {
		slice.Insert(0, x)
	}
}

func ExampleStr_Insert() {
	slice := NewStrV("1", "3")
	fmt.Println(slice.Insert(1, "2"))
	// Output: 123
}

func TestStr_Insert(t *testing.T) {

	// append
	{
		slice := NewStrV()
		assert.Equal(t, NewStrV("0"), slice.Insert(-1, "0"))
		assert.Equal(t, NewStrV("0", "1"), slice.Insert(-1, "1"))
		assert.Equal(t, NewStrV("0", "1", "2"), slice.Insert(-1, "2"))
	}

	// [] append
	{
		slice := NewStrV()
		assert.Equal(t, NewStrV("0"), slice.Insert(-1, []string{"0"}))
		assert.Equal(t, NewStrV("0", "1", "2"), slice.Insert(-1, []string{"1", "2"}))
	}

	// prepend
	{
		slice := NewStrV()
		assert.Equal(t, NewStrV("2"), slice.Insert(0, "2"))
		assert.Equal(t, NewStrV("1", "2"), slice.Insert(0, "1"))
		assert.Equal(t, NewStrV("0", "1", "2"), slice.Insert(0, "0"))
	}

	// [] prepend
	{
		slice := NewStrV()
		assert.Equal(t, NewStrV("2"), slice.Insert(0, []string{"2"}))
		assert.Equal(t, NewStrV("0", "1", "2"), slice.Insert(0, []string{"0", "1"}))
	}

	// middle pos
	{
		slice := NewStrV("0", "5")
		assert.Equal(t, NewStrV("0", "1", "5"), slice.Insert(1, "1"))
		assert.Equal(t, NewStrV("0", "1", "2", "5"), slice.Insert(2, "2"))
		assert.Equal(t, NewStrV("0", "1", "2", "3", "5"), slice.Insert(3, "3"))
		assert.Equal(t, NewStrV("0", "1", "2", "3", "4", "5"), slice.Insert(4, "4"))
	}

	// [] middle pos
	{
		slice := NewStrV("0", "5")
		assert.Equal(t, NewStrV("0", "1", "2", "5"), slice.Insert(1, []string{"1", "2"}))
		assert.Equal(t, NewStrV("0", "1", "2", "3", "4", "5"), slice.Insert(3, []string{"3", "4"}))
	}

	// middle neg
	{
		slice := NewStrV("0", "5")
		assert.Equal(t, NewStrV("0", "1", "5"), slice.Insert(-2, "1"))
		assert.Equal(t, NewStrV("0", "1", "2", "5"), slice.Insert(-2, "2"))
		assert.Equal(t, NewStrV("0", "1", "2", "3", "5"), slice.Insert(-2, "3"))
		assert.Equal(t, NewStrV("0", "1", "2", "3", "4", "5"), slice.Insert(-2, "4"))
	}

	// [] middle neg
	{
		slice := NewStrV(0, 5)
		assert.Equal(t, NewStrV(0, 1, 2, 5), slice.Insert(-2, []string{"1", "2"}))
		assert.Equal(t, NewStrV(0, "1", "2", "3", 4, 5), slice.Insert(-2, []int{3, 4}))
	}

	// error cases
	{
		var slice *Str
		assert.False(t, slice.Insert(0, 0).Nil())
		assert.Equal(t, NewStrV("0"), slice.Insert(0, "0"))
		assert.Equal(t, NewStrV("0", "1"), NewStrV("0", "1").Insert(-10, "1"))
		assert.Equal(t, NewStrV("0", "1"), NewStrV("0", "1").Insert(10, "1"))
		assert.Equal(t, NewStrV("0", "1"), NewStrV("0", "1").Insert(2, "1"))
		assert.Equal(t, NewStrV("0", "1"), NewStrV("0", "1").Insert(-3, "1"))
	}

	// [] error cases
	{
		var slice *Str
		assert.False(t, slice.Insert(0, 0).Nil())
		assert.Equal(t, NewStrV(0), slice.Insert(0, 0))
		assert.Equal(t, NewStrV(0, 1), NewStrV(0, 1).Insert(-10, 1))
		assert.Equal(t, NewStrV(0, 1), NewStrV(0, 1).Insert(10, 1))
		assert.Equal(t, NewStrV(0, 1), NewStrV(0, 1).Insert(2, 1))
		assert.Equal(t, NewStrV(0, 1), NewStrV(0, 1).Insert(-3, 1))
	}

	// Conversion
	{
		assert.Equal(t, NewStrV("1", "2", "3"), NewStrV(1, 3).Insert(1, Object{2}))
		assert.Equal(t, NewStrV("1", "2", "3"), NewStrV(1, 3).Insert(1, "2"))
		assert.Equal(t, NewStrV(true, "2", "3"), NewStrV(2, 3).Insert(0, true))
		assert.Equal(t, NewStrV("1", "2", "3"), NewStrV(1, 2).Insert(-1, Char('3')))
	}

	// [] Conversion
	{
		assert.Equal(t, NewStrV("1", "2", "3", 4), NewStrV(1, 4).Insert(1, []Object{{2}, {3}}))
		assert.Equal(t, NewStrV("1", "2", "3", 4), NewStrV(1, 4).Insert(1, []string{"2", "3"}))
		assert.Equal(t, NewStrV(false, true, "2", "3"), NewStrV(2, 3).Insert(0, []bool{false, true}))
		assert.Equal(t, NewStrV("1", "2", "3", 4), NewStrV(1, 2).Insert(-1, []Char{'3', '4'}))
	}
}

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
