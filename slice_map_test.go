package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Methods on both pointer and non-pointer
//--------------------------------------------------------------------------------------------------
func TestSliceOfMap_Methods(t *testing.T) {

	// Test conversion directly
	m, err := ToSliceOfMapE([]map[string]interface{}{{"foo": "bar"}})
	assert.Nil(t, err)
	assert.Equal(t, (&SliceOfMap{M().Add("foo", "bar")}).G(), m.G())

	// Test indirectly
	p := NewSliceOfMapV([]map[string]interface{}{{"foo": "bar"}})
	assert.Equal(t, (&SliceOfMap{M().Add("foo", "bar")}).G(), p.G())

	slice := *p
	assert.Equal(t, []*StringMap{}, slice.DropAt(0).O())
	assert.Equal(t, []*StringMap{}, slice.O())
}

// NewSliceOfMap
//--------------------------------------------------------------------------------------------------
// func BenchmarkNewSliceOfMap_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = []string{src[i], src[i] + string(1), src[i] + string(2), src[i] + string(3), src[i] + string(4), src[i] + string(5), src[i] + string(6), src[i] + string(7), src[i] + string(8), src[i] + string(9)}
// 	}
// }

// func BenchmarkNewSliceOfMap_Slice(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = NewSliceOfMap([]string{src[i], src[i] + string(1), src[i] + string(2), src[i] + string(3), src[i] + string(4), src[i] + string(5), src[i] + string(6), src[i] + string(7), src[i] + string(8), src[i] + string(9)})
// 	}
// }

func ExampleNewSliceOfMap() {
	slice := NewSliceOfMap([]map[string]interface{}{{"foo": "bar"}})
	fmt.Println(slice)
	// Output: [&[{foo bar}]]
}

func TestSliceOfMap_NewSliceOfMap(t *testing.T) {

	// empty
	assert.Equal(t, &SliceOfMap{}, NewSliceOfMap([]*StringMap{}))
	assert.Equal(t, []*StringMap{}, NewSliceOfMap([]*StringMap{}).O())

	// slice
	assert.Equal(t, &SliceOfMap{M().Add("foo", "bar")}, NewSliceOfMap([]*StringMap{M().Add("foo", "bar")}))
	assert.Equal(t, []*StringMap{M().Add("foo", "bar")}, NewSliceOfMap([]*StringMap{M().Add("foo", "bar")}).O())

	// Conversion
	assert.Equal(t, []*StringMap{}, NewSliceOfMap("1").O())
	assert.Equal(t, &SliceOfMap{M().Add("foo", "bar")}, NewSliceOfMap(map[string]interface{}{"foo": "bar"}))
	assert.Equal(t, &SliceOfMap{M().Add("foo", "bar")}, NewSliceOfMap(map[string]interface{}{"foo": "bar"}))
	assert.Equal(t, &SliceOfMap{M().Add("foo", "bar")}, NewSliceOfMap([]map[string]string{{"foo": "bar"}}))
}

// NewSliceOfMapV
//--------------------------------------------------------------------------------------------------
// func BenchmarkNewSliceOfMapV_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = append([]string{}, src[i], src[i]+string(1), src[i]+string(2), src[i]+string(3), src[i]+string(4), src[i]+string(5), src[i]+string(6), src[i]+string(7), src[i]+string(8), src[i]+string(9))
// 	}
// }

// func BenchmarkNewSliceOfMapV_Slice(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i += 10 {
// 		_ = NewSliceOfMapV(src[i], src[i]+string(1), src[i]+string(2), src[i]+string(3), src[i]+string(4), src[i]+string(5), src[i]+string(6), src[i]+string(7), src[i]+string(8), src[i]+string(9))
// 	}
// }

func ExampleNewSliceOfMapV_empty() {
	slice := NewSliceOfMapV()
	fmt.Println(slice)
	// Output: []
}

func ExampleNewSliceOfMapV_variadic() {
	slice := NewSliceOfMapV(map[string]interface{}{"foo1": "1"}, map[string]interface{}{"foo2": "2"})
	fmt.Println(slice)
	// Output: [&[{foo1 1}] &[{foo2 2}]]
}

func TestSliceOfMap_NewSliceOfMapV(t *testing.T) {

	// empty
	assert.Equal(t, &SliceOfMap{}, NewSliceOfMapV())
	assert.Equal(t, &SliceOfMap{}, NewSliceOfMapV([]*StringMap{}))
	assert.Equal(t, []*StringMap{}, NewSliceOfMapV([]*StringMap{}).O())

	// multiples
	assert.Equal(t, (&SliceOfMap{M().Add("foo", "bar")}).G(), NewSliceOfMapV([]map[string]interface{}{{"foo": "bar"}}).G())
	assert.Equal(t, (&SliceOfMap{M().Add("foo1", "1"), M().Add("foo2", "2")}).G(), NewSliceOfMapV(map[string]interface{}{"foo1": "1"}, map[string]interface{}{"foo2": "2"}).G())

	// Conversion
	assert.Equal(t, (&SliceOfMap{M().Add("foo", "bar")}).G(), NewSliceOfMapV(M().Add("foo", "bar")).G())
	assert.Equal(t, (&SliceOfMap{M().Add("foo", "bar")}).G(), NewSliceOfMapV(map[string]interface{}{"foo": "bar"}).G())
	assert.Equal(t, (&SliceOfMap{M().Add("foo", "bar")}).G(), NewSliceOfMapV([]map[string]string{{"foo": "bar"}}).G())
}

// Any
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Any_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_Any_Slice(t *testing.B) {
// 	src := RangeString(nines4)
// 	slice := NewSliceOfMap(src)
// 	for i := range src {
// 		slice.Any(i)
// 	}
// }

func ExampleSliceOfMap_Any_empty() {
	slice := NewSliceOfMapV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleSliceOfMap_Any_notEmpty() {
	slice := NewSliceOfMapV("1:", "2:", "3:")
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleSliceOfMap_Any_contains() {
	slice := NewSliceOfMapV("1:", "2:", "3:")
	fmt.Println(slice.Any("1"))
	// Output: true
}

func ExampleSliceOfMap_Any_containsAny() {
	slice := NewSliceOfMapV("1:", "2:", "3:")
	fmt.Println(slice.Any("0", "1"))
	// Output: true
}

func TestSliceOfMap_Any(t *testing.T) {

	// empty
	var nilSlice *SliceOfMap
	assert.False(t, nilSlice.Any())
	assert.False(t, NewSliceOfMapV().Any())

	// single
	assert.True(t, NewSliceOfMapV("2:").Any())

	// invalid
	assert.False(t, NewSliceOfMapV("1:", "2").Any(TestObj{"2"}))

	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any("2"))
	assert.False(t, NewSliceOfMapV("1:", "2:", "3:").Any(4))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(4, "3"))
	assert.False(t, NewSliceOfMapV("1:", "2:", "3:").Any(4, 5))

	// conversion
	assert.True(t, NewSliceOfMapV("1:", "2:").Any(Object{2}))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(int8(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(int16(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(int32('2')))
	assert.False(t, NewSliceOfMapV("1:", "2:", "3:").Any(int32(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(int64(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(uint8('2')))
	assert.False(t, NewSliceOfMapV("1:", "2:", "3:").Any(uint8(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(uint16(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(uint32(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(uint64(2)))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").Any(uint64(2)))
}

// AnyS
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_AnyS_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_AnyS_Slice(t *testing.B) {
// 	src := RangeString(nines4)
// 	slice := NewSliceOfMap(src)
// 	for _, x := range src {
// 		slice.Any([]string{x})
// 	}
// }

func ExampleSliceOfMap_AnyS() {
	slice := NewSliceOfMapV("1:", "2:", "3:")
	fmt.Println(slice.AnyS([]string{"0", "1"}))
	// Output: true
}

func TestSliceOfMap_AnyS(t *testing.T) {
	// nil
	{
		var slice *SliceOfMap
		assert.False(t, slice.AnyS([]string{"1"}))
		assert.False(t, NewSliceOfMapV("1:").AnyS(nil))
	}

	// []string
	{
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS([]string{"1"}))
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS([]string{"4", "3"}))
		assert.False(t, NewSliceOfMapV("1:", "2:", "3:").AnyS([]string{"4", "5"}))
	}

	// *[]string
	{
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(&([]string{"1"})))
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(&([]string{"4", "3"})))
		assert.False(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(&([]string{"4", "5"})))
	}

	// Slice
	{
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(ISlice(NewSliceOfMapV("1:"))))
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(ISlice(NewSliceOfMapV("4:", "3:"))))
		assert.False(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(ISlice(NewSliceOfMapV("4:", "5:"))))
	}

	// SliceOfMap
	{
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(*NewSliceOfMapV("1:")))
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(*NewSliceOfMapV("4:", "3:")))
		assert.False(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(*NewSliceOfMapV("4:", "5:")))
	}

	// *SliceOfMap
	{
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(NewSliceOfMapV("1:")))
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(NewSliceOfMapV("4:", "3:")))
		assert.False(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(NewSliceOfMapV("4:", "5:")))
	}

	// invalid types
	assert.False(t, NewSliceOfMapV("1:", "2").AnyS(nil))
	assert.False(t, NewSliceOfMapV("1:", "2").AnyS((*[]string)(nil)))
	assert.False(t, NewSliceOfMapV("1:", "2").AnyS((*SliceOfMap)(nil)))

	// Conversion
	{
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(int64(1)))
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(2))
		assert.False(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(true))
		assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyS(Char('3')))
	}
}

// AnyW
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_AnyW_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkSliceOfMap_AnyW_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	NewSliceOfMap(src).AnyW(func(x O) bool {
// 		return ExB(x.(string) == string(nines4))
// 	})
// }

// func ExampleSliceOfMap_AnyW() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.AnyW(func(x O) bool {
// 		return ExB(x.(string) == "2")
// 	}))
// 	// Output: true
// }

func TestSliceOfMap_AnyW(t *testing.T) {

	// empty
	var slice *SliceOfMap
	assert.False(t, slice.AnyW(func(x O) bool { return ToStringMap(x).Exists("2") }))
	assert.False(t, NewSliceOfMapV().AnyW(func(x O) bool { return ToStringMap(x).Exists("2") }))

	// single
	assert.True(t, NewSliceOfMapV("2:").AnyW(func(x O) bool { return ToStringMap(x).Exists("2") }))
	assert.True(t, NewSliceOfMapV("1:", "2:").AnyW(func(x O) bool { return ToStringMap(x).Any("2") }))
	assert.True(t, NewSliceOfMapV("1:", "2:", "3:").AnyW(func(x O) bool { return ToStringMap(x).Any("4", "3") }))
}

// Append
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Append_Go(t *testing.B) {
// 	src := []string{}
// 	for _, i := range RangeString(nines6) {
// 		src = append(src, i)
// 	}
// }

// func BenchmarkSliceOfMap_Append_Slice(t *testing.B) {
// 	slice := NewSliceOfMapV()
// 	for _, i := range RangeString(nines6) {
// 		slice.Append(i)
// 	}
// }

// func ExampleSliceOfMap_Append() {
// 	slice := NewSliceOfMapV("1").Append("2").Append("3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

func TestSliceOfMap_Append(t *testing.T) {

	// nil
	{
		var nilSlice *SliceOfMap
		assert.Equal(t, NewSliceOfMapV("0:"), nilSlice.Append("0:"))
		assert.Equal(t, (*SliceOfMap)(nil), nilSlice)
	}

	// Append one back to back
	{
		var slice *SliceOfMap
		assert.Equal(t, true, slice.Nil())
		slice = NewSliceOfMapV()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, false, slice.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		slice.Append("1:")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, (&SliceOfMap{M().Add("1", nil)}).G(), slice.G())

		// Second append another which will be 2x at most
		slice.Append("2:")
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, (&SliceOfMap{M().Add("1", nil), MV("2:")}).G(), slice.G())
		assert.Equal(t, NewSliceOfMapV("1:", "2:").G(), slice.G())
	}

	// Start with just appending without chaining
	{
		slice := NewSliceOfMapV()
		assert.Equal(t, 0, slice.Len())
		slice.Append("1:")
		assert.Equal(t, (&SliceOfMap{M().Add("1", nil)}).G(), slice.G())
		slice.Append("2:")
		assert.Equal(t, (&SliceOfMap{M().Add("1", nil), MV("2:")}).G(), slice.G())
	}

	// Start with nil not chained
	{
		slice := NewSliceOfMapV()
		assert.Equal(t, 0, slice.Len())
		slice.Append("1:").Append("2:").Append("3:")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, (&SliceOfMap{MV("1:"), MV("2:"), MV("3:")}).G(), slice.G())
	}

	// Start with nil chained
	{
		slice := NewSliceOfMapV().Append("1:").Append("2:")
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, (&SliceOfMap{MV("1:"), MV("2:")}).G(), ToSliceOfMap(slice).G())
	}

	// Start with non nil
	{
		slice := NewSliceOfMapV("1:").Append("2:").Append("3:")
		assert.Equal(t, 3, slice.Len())
		assert.Equal(t, (&SliceOfMap{MV("1:"), MV("2:"), MV("3:")}).G(), ToSliceOfMap(slice).G())
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:").G(), ToSliceOfMap(slice).G())
	}

	// Use append result directly
	{
		slice := NewSliceOfMapV("1:")
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, (&SliceOfMap{MV("1:"), MV("2:")}).G(), ToSliceOfMap(slice.Append("2:")).G())
		assert.Equal(t, NewSliceOfMapV("1:", "2:").G(), ToSliceOfMap(slice).G())
	}
}

// AppendV
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_AppendV_Go(t *testing.B) {
// 	src := []string{}
// 	src = append(src, RangeString(nines6)...)
// }

// func BenchmarkSliceOfMap_AppendV_Slice(t *testing.B) {
// 	n := NewSliceOfMapV()
// 	new := rangeO(0, nines6)
// 	n.AppendV(new...)
// }

func ExampleSliceOfMap_AppendV() {
	slice := NewSliceOfMapV("1:").AppendV("2:", "3:")
	fmt.Println(slice)
	// Output: [&[{1 <nil>}] &[{2 <nil>}] &[{3 <nil>}]]
}

func TestSliceOfMap_AppendV(t *testing.T) {

	// nil
	{
		var nilSlice *SliceOfMap
		assert.Equal(t, NewSliceOfMapV("1:", "2:").G(), ToSliceOfMap(nilSlice.AppendV("1:", "2:")).G())
	}

	// Append many src
	{
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:").G(), ToSliceOfMap(NewSliceOfMapV("1:").AppendV("2:", "3:")).G())
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:", "5:").G(), ToSliceOfMap(NewSliceOfMapV("1:").AppendV("2:", "3:").AppendV("4:", "5:")).G())
	}
}

// At
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_At_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for _, x := range src {
// 		assert.IsType(t, 0, x)
// 	}
// }

// func BenchmarkSliceOfMap_At_Slice(t *testing.B) {
// 	src := RangeString(nines6)
// 	slice := NewSliceOfMap(src)
// 	for i := 0; i < len(src); i++ {
// 		_, ok := (slice.At(i).O()).(string)
// 		assert.True(t, ok)
// 	}
// }

func ExampleSliceOfMap_At() {
	slice := NewSliceOfMapV("1:", "2:", "3:")
	fmt.Println(slice.At(2))
	// Output: &[{3 <nil>}]
}

func TestSliceOfMap_At(t *testing.T) {

	// nil
	{
		var nilSlice *SliceOfMap
		assert.Equal(t, Obj(nil), nilSlice.At(0))
	}

	// src
	{
		slice := NewSliceOfMapV("1:", "2:", "3:", "4:")
		assert.Equal(t, MV("4:").G(), slice.At(-1).ToStringMap().G())
		assert.Equal(t, MV("3:").G(), slice.At(-2).ToStringMap().G())
		assert.Equal(t, MV("2:").G(), slice.At(-3).ToStringMap().G())
		assert.Equal(t, MV("1:").G(), slice.At(0).ToStringMap().G())
		assert.Equal(t, MV("2:").G(), slice.At(1).ToStringMap().G())
		assert.Equal(t, MV("3:").G(), slice.At(2).ToStringMap().G())
		assert.Equal(t, MV("4:").G(), slice.At(3).ToStringMap().G())
	}

	// index out of bounds
	{
		slice := NewSliceOfMapV("1:")
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------
func ExampleSliceOfMap_Clear() {
	slice := NewSliceOfMapV("1:").Concat([]string{"2", "3"})
	fmt.Println(slice.Clear())
	// Output: []
}

func TestSliceOfMap_Clear(t *testing.T) {

	// nil
	{
		var slice *SliceOfMap
		assert.Equal(t, NewSliceOfMapV(), slice.Clear())
		assert.Equal(t, (*SliceOfMap)(nil), slice)
	}

	// int
	{
		slice := NewSliceOfMapV(1, 2, 3, 4)
		assert.Equal(t, NewSliceOfMapV(), slice.Clear())
		assert.Equal(t, NewSliceOfMapV(), slice.Clear())
		assert.Equal(t, NewSliceOfMapV(), slice)
	}
}

// Concat
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Concat_Go(t *testing.B) {
// 	dest := []string{}
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkSliceOfMap_Concat_Slice(t *testing.B) {
// 	dest := NewSliceOfMapV()
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.Concat(src[j:i])
// 		j = i
// 	}
// }

func ExampleSliceOfMap_Concat() {
	slice := NewSliceOfMapV("1:").Concat([]interface{}{"2:", "3:"})
	fmt.Println(slice)
	// Output: [&[{1 <nil>}] &[{2 <nil>}] &[{3 <nil>}]]
}

func TestSliceOfMap_Concat(t *testing.T) {

	// nil
	{
		var slice *SliceOfMap
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), slice.Concat([]interface{}{"1:", "2:"}))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), NewSliceOfMapV("1:", "2:").Concat(nil))
	}

	// []string
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.Concat([]interface{}{"2:", "3:"})
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), slice.Append("2:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:"), concated)
	}

	// *[]string
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.Concat(&([]interface{}{"2:", "3:"}))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), slice.Append("2:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:"), concated)
	}

	// *SliceOfMap
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.Concat(NewSliceOfMapV("2:", "3:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), slice.Append("2:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:"), concated)
	}

	// SliceOfMap
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.Concat(*NewSliceOfMapV("2:", "3:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), slice.Append("2:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:"), concated)
	}

	// Slice
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.Concat(ISlice(NewSliceOfMapV("2:", "3:")))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), slice.Append("2:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:"), concated)
	}

	// nils
	{
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), NewSliceOfMapV("1:", "2:").Concat((*[]interface{})(nil)))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), NewSliceOfMapV("1:", "2:").Concat((*SliceOfMap)(nil)))
	}
}

// ConcatM
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_ConcatM_Go(t *testing.B) {
// 	dest := []string{}
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest = append(dest, (src[j:i])...)
// 		j = i
// 	}
// }

// func BenchmarkSliceOfMap_ConcatM_Slice(t *testing.B) {
// 	dest := NewSliceOfMapV()
// 	src := RangeString(nines6)
// 	j := 0
// 	for i := 10; i < len(src); i += 10 {
// 		dest.ConcatM(src[j:i])
// 		j = i
// 	}
// }

func ExampleSliceOfMap_ConcatM() {
	slice := NewSliceOfMapV("1:").ConcatM([]interface{}{"2:", "3:"})
	fmt.Println(slice)
	// Output: [&[{1 <nil>}] &[{2 <nil>}] &[{3 <nil>}]]
}

func TestSliceOfMap_ConcatM(t *testing.T) {

	// nil
	{
		var slice *SliceOfMap
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), slice.ConcatM([]interface{}{"1:", "2:"}))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), NewSliceOfMapV("1:", "2:").ConcatM(nil))
	}

	// []string
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.ConcatM([]interface{}{"2:", "3:"})
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), slice.Append("4:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), concated)
	}

	// *[]string
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.ConcatM(&([]interface{}{"2:", "3:"}))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), slice.Append("4:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), concated)
	}

	// *SliceOfMap
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.ConcatM(NewSliceOfMapV("2:", "3:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), slice.Append("4:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), concated)
	}

	// SliceOfMap
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.ConcatM(*NewSliceOfMapV("2:", "3:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), slice.Append("4:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), concated)
	}

	// Slice
	{
		slice := NewSliceOfMapV("1:")
		concated := slice.ConcatM(ISlice(NewSliceOfMapV("2:", "3:")))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), slice.Append("4:"))
		assert.Equal(t, NewSliceOfMapV("1:", "2:", "3:", "4:"), concated)
	}

	// nils
	{
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), NewSliceOfMapV("1:", "2:").ConcatM((*[]interface{})(nil)))
		assert.Equal(t, NewSliceOfMapV("1:", "2:"), NewSliceOfMapV("1:", "2:").ConcatM((*SliceOfMap)(nil)))
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Copy_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	dst := make([]string, len(src), len(src))
// 	copy(dst, src)
// }

// func BenchmarkSliceOfMap_Copy_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.Copy()
// }

func ExampleSliceOfMap_Copy() {
	slice := NewSliceOfMapV("1:", "2:", "3:")
	fmt.Println(slice.Copy())
	// Output: [&[{1 <nil>}] &[{2 <nil>}] &[{3 <nil>}]]
}

func TestSliceOfMap_Copy(t *testing.T) {

	// nil or empty
	{
		var slice *SliceOfMap
		assert.Equal(t, NewSliceOfMapV(), slice.Copy(0, -1))
		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV(0).Clear().Copy(0, -1))
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := NewSliceOfMapV(1, 2, 3)
		result := original.Copy(0, -1)
		assert.Equal(t, NewSliceOfMapV(1, 2, 3).G(), original.G())
		assert.Equal(t, NewSliceOfMapV(1, 2, 3).G(), ToSliceOfMap(result).G())
		result.Set(0, 0)
		assert.Equal(t, NewSliceOfMapV(1, 2, 3).G(), original.G())
		assert.Equal(t, NewSliceOfMapV(0, 2, 3).G(), ToSliceOfMap(result).G())
	}

	// copy full array
	{
		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Copy())
		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Copy(0, -1))
		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Copy(0, 1))
		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Copy(0, 5))
		assert.Equal(t, NewSliceOfMapV(1), NewSliceOfMapV(1).Copy())
		assert.Equal(t, NewSliceOfMapV(1, 2, 3), NewSliceOfMapV(1, 2, 3).Copy())
		assert.Equal(t, NewSliceOfMapV(1, 2, 3), NewSliceOfMapV(1, 2, 3).Copy(0, -1))
		assert.Equal(t, NewSliceOfMap([]interface{}{1, 2, 3}), NewSliceOfMap([]interface{}{1, 2, 3}).Copy())
		assert.Equal(t, NewSliceOfMap([]interface{}{1, 2, 3}), NewSliceOfMap([]interface{}{1, 2, 3}).Copy(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, NewSliceOfMapV(1), NewSliceOfMapV(1).Copy(0, 2))
		assert.Equal(t, NewSliceOfMapV(1, 2, 3), NewSliceOfMapV(1, 2, 3).Copy(-6, 6))
	}

	// mutually exclusive
	{
		slice := NewSliceOfMapV(1, 2, 3, 4)
		assert.Equal(t, NewSliceOfMapV(), slice.Copy(2, -3))
		assert.Equal(t, NewSliceOfMapV(), slice.Copy(0, -5))
		assert.Equal(t, NewSliceOfMapV(), slice.Copy(4, -1))
		assert.Equal(t, NewSliceOfMapV(), slice.Copy(6, -1))
		assert.Equal(t, NewSliceOfMapV(), slice.Copy(3, -2))
	}

	// singles
	{
		slice := NewSliceOfMapV(1, 2, 3, 4)
		assert.Equal(t, NewSliceOfMapV(4), slice.Copy(-1, -1))
		assert.Equal(t, NewSliceOfMapV(3), slice.Copy(-2, -2))
		assert.Equal(t, NewSliceOfMapV(2), slice.Copy(-3, -3))
		assert.Equal(t, NewSliceOfMapV(1), slice.Copy(0, 0))
		assert.Equal(t, NewSliceOfMapV(1), slice.Copy(-4, -4))
		assert.Equal(t, NewSliceOfMapV(2), slice.Copy(1, 1))
		assert.Equal(t, NewSliceOfMapV(2), slice.Copy(1, -3))
		assert.Equal(t, NewSliceOfMapV(3), slice.Copy(2, 2))
		assert.Equal(t, NewSliceOfMapV(3), slice.Copy(2, -2))
		assert.Equal(t, NewSliceOfMapV(4), slice.Copy(3, 3))
		assert.Equal(t, NewSliceOfMapV(4), slice.Copy(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3).Copy(1, -1))
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3).Copy(1, 2))
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3).Copy(-2, -1))
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3).Copy(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, NewSliceOfMapV(1, 2), NewSliceOfMapV(1, 2, 3).Copy(0, -2))
		assert.Equal(t, NewSliceOfMapV(1, 2), NewSliceOfMapV(1, 2, 3).Copy(-3, -2))
		assert.Equal(t, NewSliceOfMapV(1, 2), NewSliceOfMapV(1, 2, 3).Copy(-3, 1))
		assert.Equal(t, NewSliceOfMapV(1, 2), NewSliceOfMapV(1, 2, 3).Copy(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3, 4).Copy(1, -2))
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3, 4).Copy(-3, -2))
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3, 4).Copy(-3, 2))
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3, 4).Copy(1, 2))
	}

	// random
	{
		assert.Equal(t, NewSliceOfMapV(1), NewSliceOfMapV(1, 2, 3).Copy(0, -3))
		assert.Equal(t, NewSliceOfMapV(2, 3), NewSliceOfMapV(1, 2, 3).Copy(1, 2))
		assert.Equal(t, NewSliceOfMapV(1, 2, 3), NewSliceOfMapV(1, 2, 3).Copy(0, 2))
	}
}

// // Count
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Count_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkSliceOfMap_Count_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	NewSliceOfMap(src).Count(nines4)
// }

// func ExampleSliceOfMap_Count() {
// 	slice := NewSliceOfMapV("1", "2", "2")
// 	fmt.Println(slice.Count("2"))
// 	// Output: 2
// }

// func TestSliceOfMap_Count(t *testing.T) {

// 	// empty
// 	var slice *SliceOfMap
// 	assert.Equal(t, 0, slice.Count(0))
// 	assert.Equal(t, 0, NewSliceOfMapV().Count(0))

// 	assert.Equal(t, 1, NewSliceOfMapV("2", "3").Count("2"))
// 	assert.Equal(t, 2, NewSliceOfMapV("1", "2", "2").Count("2"))
// 	assert.Equal(t, 4, NewSliceOfMapV("4", "4", "3", "4", "4").Count("4"))
// 	assert.Equal(t, 3, NewSliceOfMapV("3", "2", "3", "3", "5").Count("3"))
// 	assert.Equal(t, 1, NewSliceOfMapV("1", "2", "3").Count("3"))
// }

// // CountW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_CountW_Go(t *testing.B) {
// 	src := RangeString(nines5)
// 	for _, x := range src {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkSliceOfMap_CountW_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	NewSliceOfMap(src).CountW(func(x O) bool {
// 		return ExB(x.(string) == string(nines4))
// 	})
// }

// func ExampleSliceOfMap_CountW() {
// 	slice := NewSliceOfMapV("1", "2", "2")
// 	fmt.Println(slice.CountW(func(x O) bool {
// 		return ExB(x.(string) == "2")
// 	}))
// 	// Output: 2
// }

// func TestSliceOfMap_CountW(t *testing.T) {

// 	// empty
// 	var slice *SliceOfMap
// 	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(string) > "0") }))
// 	assert.Equal(t, 0, NewSliceOfMapV().CountW(func(x O) bool { return ExB(x.(string) > "0") }))

// 	assert.Equal(t, 1, NewSliceOfMapV("2", "3").CountW(func(x O) bool { return ExB(x.(string) > "2") }))
// 	assert.Equal(t, 1, NewSliceOfMapV("1", "2").CountW(func(x O) bool { return ExB(x.(string) == "2") }))
// 	assert.Equal(t, 1, NewSliceOfMapV("1", "2", "3").CountW(func(x O) bool { return ExB(x.(string) == "4" || x.(string) == "3") }))
// }

// // Drop
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Drop_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_Drop_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 1 {
// 		slice.Drop(1, 10)
// 	}
// }

// func ExampleSliceOfMap_Drop() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Drop(0, 1))
// 	// Output: [3]
// }

// func TestSliceOfMap_Drop(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.Drop(0, 1))
// 	}

// 	// invalid
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(1))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(4, 4))

// 	// drop 1
// 	assert.Equal(t, NewSliceOfMapV("2", "3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(0, 0))
// 	assert.Equal(t, NewSliceOfMapV("1", "3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(1, 1))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(2, 2))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3", "4").Drop(3, 3))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3", "4").Drop(-1, -1))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(-2, -2))
// 	assert.Equal(t, NewSliceOfMapV("1", "3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(-3, -3))
// 	assert.Equal(t, NewSliceOfMapV("2", "3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(-4, -4))

// 	// drop 2
// 	assert.Equal(t, NewSliceOfMapV("3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(0, 1))
// 	assert.Equal(t, NewSliceOfMapV("1", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(1, 2))
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3", "4").Drop(2, 3))
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3", "4").Drop(-2, -1))
// 	assert.Equal(t, NewSliceOfMapV("1", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(-3, -2))
// 	assert.Equal(t, NewSliceOfMapV("3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(-4, -3))

// 	// drop 3
// 	assert.Equal(t, NewSliceOfMapV("4"), NewSliceOfMapV("1", "2", "3", "4").Drop(0, 2))
// 	assert.Equal(t, NewSliceOfMapV("1"), NewSliceOfMapV("1", "2", "3", "4").Drop(-3, -1))

// 	// drop everything and beyond
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Drop())
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Drop(0, 3))
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Drop(0, -1))
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Drop(-4, -1))
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Drop(-6, -1))
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Drop(0, 10))

// 	// move index within bounds
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3", "4").Drop(3, 4))
// 	assert.Equal(t, NewSliceOfMapV("2", "3", "4"), NewSliceOfMapV("1", "2", "3", "4").Drop(-5, 0))
// }

// // DropAt
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_DropAt_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_DropAt_Slice(t *testing.B) {
// 	index := Range(0, nines5)
// 	slice := NewSliceOfMap(RangeString(nines5))
// 	for i := range index {
// 		slice.DropAt(i)
// 	}
// }

// func ExampleSliceOfMap_DropAt() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.DropAt(1))
// 	// Output: [1 3]
// }

// func TestSliceOfMap_DropAt(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.DropAt(0))
// 	}

// 	// drop all and more
// 	{
// 		slice := NewSliceOfMapV("0", "1", "2")
// 		assert.Equal(t, NewSliceOfMapV("0", "1"), slice.DropAt(-1))
// 		assert.Equal(t, NewSliceOfMapV("0"), slice.DropAt(-1))
// 		assert.Equal(t, NewSliceOfMapV(), slice.DropAt(-1))
// 		assert.Equal(t, NewSliceOfMapV(), slice.DropAt(-1))
// 	}

// 	// drop invalid
// 	assert.Equal(t, NewSliceOfMapV("0", "1", "2"), NewSliceOfMapV("0", "1", "2").DropAt(3))
// 	assert.Equal(t, NewSliceOfMapV("0", "1", "2"), NewSliceOfMapV("0", "1", "2").DropAt(-4))

// 	// drop last
// 	assert.Equal(t, NewSliceOfMapV("0", "1"), NewSliceOfMapV("0", "1", "2").DropAt(2))
// 	assert.Equal(t, NewSliceOfMapV("0", "1"), NewSliceOfMapV("0", "1", "2").DropAt(-1))

// 	// drop middle
// 	assert.Equal(t, NewSliceOfMapV("0", "2"), NewSliceOfMapV("0", "1", "2").DropAt(1))
// 	assert.Equal(t, NewSliceOfMapV("0", "2"), NewSliceOfMapV("0", "1", "2").DropAt(-2))

// 	// drop first
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("0", "1", "2").DropAt(0))
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("0", "1", "2").DropAt(-3))
// }

// // DropFirst
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_DropFirst_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkSliceOfMap_DropFirst_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirst()
// 	}
// }

// func ExampleSliceOfMap_DropFirst() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.DropFirst())
// 	// Output: [2 3]
// }

// func TestSliceOfMap_DropFirst(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.DropFirst())
// 	}

// 	// drop all and beyond
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), slice.DropFirst())
// 		assert.Equal(t, NewSliceOfMapV("3"), slice.DropFirst())
// 		assert.Equal(t, NewSliceOfMapV(), slice.DropFirst())
// 		assert.Equal(t, NewSliceOfMapV(), slice.DropFirst())
// 	}
// }

// // DropFirstN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_DropFirstN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkSliceOfMap_DropFirstN_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropFirstN(10)
// 	}
// }

// func ExampleSliceOfMap_DropFirstN() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.DropFirstN(2))
// 	// Output: [3]
// }

// func TestSliceOfMap_DropFirstN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.DropFirstN(1))
// 	}

// 	// negative value
// 	assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").DropFirstN(-1))

// 	// drop none
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").DropFirstN(0))

// 	// drop 1
// 	assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").DropFirstN(1))

// 	// drop 2
// 	assert.Equal(t, NewSliceOfMapV("3"), NewSliceOfMapV("1", "2", "3").DropFirstN(2))

// 	// drop 3
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3").DropFirstN(3))

// 	// drop beyond
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3").DropFirstN(4))
// }

// // DropLast
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_DropLast_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkSliceOfMap_DropLast_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLast()
// 	}
// }

// func ExampleSliceOfMap_DropLast() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.DropLast())
// 	// Output: [1 2]
// }

// func TestSliceOfMap_DropLast(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.DropLast())
// 	}

// 	// negative value
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3").DropLastN(-1))

// 	slice := NewSliceOfMapV("1", "2", "3")
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), slice.DropLast())
// 	assert.Equal(t, NewSliceOfMapV("1"), slice.DropLast())
// 	assert.Equal(t, NewSliceOfMapV(), slice.DropLast())
// 	assert.Equal(t, NewSliceOfMapV(), slice.DropLast())
// }

// // DropLastN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_DropLastN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkSliceOfMap_DropLastN_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.DropLastN(10)
// 	}
// }

// func ExampleSliceOfMap_DropLastN() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.DropLastN(2))
// 	// Output: [1]
// }

// func TestSliceOfMap_DropLastN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.DropLastN(1))
// 	}

// 	// drop none
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").DropLastN(0))

// 	// drop 1
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3").DropLastN(1))

// 	// drop 2
// 	assert.Equal(t, NewSliceOfMapV("1"), NewSliceOfMapV("1", "2", "3").DropLastN(2))

// 	// drop 3
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3").DropLastN(3))

// 	// drop beyond
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3").DropLastN(4))
// }

// // DropW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_DropW_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_DropW_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines5))
// 	slice.DropW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	})
// }

// func ExampleSliceOfMap_DropW() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.DropW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	}))
// 	// Output: [1 3]
// }

// func TestSliceOfMap_DropW(t *testing.T) {

// 	// drop all odd values
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		slice.DropW(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 != 0)
// 		})
// 		assert.Equal(t, NewSliceOfMapV("2", "4", "6", "8"), slice)
// 	}

// 	// drop all even values
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		slice.DropW(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 == 0)
// 		})
// 		assert.Equal(t, NewSliceOfMapV("1", "3", "5", "7", "9"), slice)
// 	}
// }

// Each
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Each_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_Each_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).Each(func(x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

func ExampleSliceOfMap_Each() {
	NewSliceOfMapV([]map[string]interface{}{{"foo": "bar"}}).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: &[{foo bar}]
}

func TestSliceOfMap_Each(t *testing.T) {

	// nil or empty
	{
		var slice *SliceOfMap
		slice.Each(func(x O) {})
	}

	// Loop through
	{
		results := []string{}
		NewSliceOfMapV([]map[string]interface{}{{"foo": "1"}, {"foo": "2"}}).Each(func(x O) {
			results = append(results, ToStringMap(x).Get("foo").A())
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// EachE
//--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_EachE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "0", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_EachE_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).EachE(func(x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

func ExampleSliceOfMap_EachE() {
	NewSliceOfMapV([]map[string]interface{}{{"foo": "bar"}}).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: &[{foo bar}]
}

func TestSliceOfMap_EachE(t *testing.T) {

	// nil or empty
	{
		var slice *SliceOfMap
		slice.EachE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []string{}
		NewSliceOfMapV([]map[string]interface{}{{"foo": "1"}, {"foo": "2"}, {"foo": "3"}}).EachE(func(x O) error {
			results = append(results, ToStringMap(x).Get("foo").A())
			return nil
		})
		assert.Equal(t, []string{"1", "2", "3"}, results)
	}

	// Break early with error
	{
		results := []string{}
		NewSliceOfMapV([]map[string]interface{}{{"foo": "1"}, {"foo": "2"}, {"foo": "3"}}).EachE(func(x O) error {
			if ToStringMap(x).Get("foo").A() == "3" {
				return Break
			}
			results = append(results, ToStringMap(x).Get("foo").A())
			return nil
		})
		assert.Equal(t, []string{"1", "2"}, results)
	}
}

// // EachI
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_EachI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_EachI_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).EachI(func(i int, x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleSliceOfMap_EachI() {
// 	NewSliceOfMapV("1", "2", "3").EachI(func(i int, x O) {
// 		fmt.Printf("%v:%v", i, x)
// 	})
// 	// Output: 0:11:22:3
// }

// func TestSliceOfMap_EachI(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		slice.EachI(func(i int, x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachI(func(i int, x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}
// }

// // EachIE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_EachIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_EachIE_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).EachIE(func(i int, x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleSliceOfMap_EachIE() {
// 	NewSliceOfMapV("1", "2", "3").EachIE(func(i int, x O) error {
// 		fmt.Printf("%v:%v", i, x)
// 		return nil
// 	})
// 	// Output: 0:11:22:3
// }

// func TestSliceOfMap_EachIE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		slice.EachIE(func(i int, x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachIE(func(i int, x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"1", "2", "3"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachIE(func(i int, x O) error {
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
// func BenchmarkSliceOfMap_EachR_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_EachR_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).EachR(func(x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleSliceOfMap_EachR() {
// 	NewSliceOfMapV("1", "2", "3").EachR(func(x O) {
// 		fmt.Printf("%v", x)
// 	})
// 	// Output: 321
// }

// func TestSliceOfMap_EachR(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		slice.EachR(func(x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachR(func(x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}
// }

// // EachRE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_EachRE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_EachRE_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).EachRE(func(x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleSliceOfMap_EachRE() {
// 	NewSliceOfMapV("1", "2", "3").EachRE(func(x O) error {
// 		fmt.Printf("%v", x)
// 		return nil
// 	})
// 	// Output: 321
// }

// func TestSliceOfMap_EachRE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		slice.EachRE(func(x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachRE(func(x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachRE(func(x O) error {
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
// func BenchmarkSliceOfMap_EachRI_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_EachRI_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).EachRI(func(i int, x O) {
// 		assert.IsType(t, "", x)
// 	})
// }

// func ExampleSliceOfMap_EachRI() {
// 	NewSliceOfMapV("1", "2", "3").EachRI(func(i int, x O) {
// 		fmt.Printf("%v:%v", i, x)
// 	})
// 	// Output: 2:31:20:1
// }

// func TestSliceOfMap_EachRI(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		slice.EachRI(func(i int, x O) {})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachRI(func(i int, x O) {
// 			results = append(results, x.(string))
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}
// }

// // EachRIE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_EachRIE_Go(t *testing.B) {
// 	action := func(x interface{}) {
// 		assert.IsType(t, "", x)
// 	}
// 	for _, x := range RangeString(nines6) {
// 		action(x)
// 	}
// }

// func BenchmarkSliceOfMap_EachRIE_Slice(t *testing.B) {
// 	NewSliceOfMap(RangeString(nines6)).EachRIE(func(i int, x O) error {
// 		assert.IsType(t, "", x)
// 		return nil
// 	})
// }

// func ExampleSliceOfMap_EachRIE() {
// 	NewSliceOfMapV("1", "2", "3").EachRIE(func(i int, x O) error {
// 		fmt.Printf("%v:%v", i, x)
// 		return nil
// 	})
// 	// Output: 2:31:20:1
// }

// func TestSliceOfMap_EachRIE(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		slice.EachRIE(func(i int, x O) error {
// 			return nil
// 		})
// 	}

// 	// Loop through
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachRIE(func(i int, x O) error {
// 			results = append(results, x.(string))
// 			return nil
// 		})
// 		assert.Equal(t, []string{"3", "2", "1"}, results)
// 	}

// 	// Break early with error
// 	{
// 		results := []string{}
// 		NewSliceOfMapV("1", "2", "3").EachRIE(func(i int, x O) error {
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
// func ExampleSliceOfMap_Empty() {
// 	fmt.Println(NewSliceOfMapV().Empty())
// 	// Output: true
// }

// func TestSliceOfMap_Empty(t *testing.T) {

// 	// nil or empty
// 	{
// 		var nilSlice *SliceOfMap
// 		assert.Equal(t, true, nilSlice.Empty())
// 	}

// 	assert.Equal(t, true, NewSliceOfMapV().Empty())
// 	assert.Equal(t, false, NewSliceOfMapV("1").Empty())
// 	assert.Equal(t, false, NewSliceOfMapV("1", "2", "3").Empty())
// 	assert.Equal(t, false, NewSliceOfMapV("1").Empty())
// 	assert.Equal(t, false, NewSliceOfMap([]string{"1", "2", "3"}).Empty())
// }

// // First
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_First_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		_ = src[0]
// 		src = src[1:]
// 	}
// }

// func BenchmarkSliceOfMap_First_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.First()
// 		slice.DropFirst()
// 	}
// }

// func ExampleSliceOfMap_First() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.First())
// 	// Output: 1
// }

// func TestSliceOfMap_First(t *testing.T) {
// 	// invalid
// 	assert.Equal(t, Obj(nil), NewSliceOfMapV().First())

// 	// int
// 	assert.Equal(t, Obj("2"), NewSliceOfMapV("2", "3").First())
// 	assert.Equal(t, Obj("3"), NewSliceOfMapV("3", "2").First())
// 	assert.Equal(t, Obj("1"), NewSliceOfMapV("1", "3", "2").First())
// }

// // FirstN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_FirstN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	_ = src[0:10]
// }

// func BenchmarkSliceOfMap_FirstN_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	slice.FirstN(10)
// }

// func ExampleSliceOfMap_FirstN() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.FirstN(2))
// 	// Output: [1 2]
// }

// func TestSliceOfMap_FirstN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.FirstN(1))
// 		assert.Equal(t, NewSliceOfMapV(), slice.FirstN(-1))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewSliceOfMapV("1", "2", "3")
// 		result := original.FirstN(2).Set(0, "0")
// 		assert.Equal(t, NewSliceOfMapV("0", "2", "3"), original)
// 		assert.Equal(t, NewSliceOfMapV("0", "2"), result)
// 	}

// 	// Get none
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3").FirstN(0))

// 	// slice full array includeing out of bounds
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().FirstN(1))
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().FirstN(10))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").FirstN(10))
// 	assert.Equal(t, NewSliceOfMap([]string{"1", "2", "3"}), NewSliceOfMap([]string{"1", "2", "3"}).FirstN(10))

// 	// grab a few diff
// 	assert.Equal(t, NewSliceOfMapV("1"), NewSliceOfMapV("1", "2", "3").FirstN(1))
// 	assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3").FirstN(2))
// }

// // G
// //--------------------------------------------------------------------------------------------------
// func ExampleSliceOfMap_G() {
// 	fmt.Println(NewSliceOfMapV("1", "2", "3").G())
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_G(t *testing.T) {
// 	assert.IsType(t, []string{}, NewSliceOfMapV().G())
// 	assert.IsType(t, []string{"1", "2", "3"}, NewSliceOfMapV("1", "2", "3").G())
// }

// // Generic
// //--------------------------------------------------------------------------------------------------
// func ExampleSliceOfMap_Generic() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Generic())
// 	// Output: false
// }

// // Index
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Index_Go(t *testing.B) {
// 	for _, x := range RangeString(nines5) {
// 		if x == string(nines4) {
// 			break
// 		}
// 	}
// }

// func BenchmarkSliceOfMap_Index_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines5))
// 	slice.Index(nines4)
// }

// func ExampleSliceOfMap_Index() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Index("2"))
// 	// Output: 1
// }

// func TestSliceOfMap_Index(t *testing.T) {

// 	// empty
// 	var slice *SliceOfMap
// 	assert.Equal(t, -1, slice.Index("2"))
// 	assert.Equal(t, -1, NewSliceOfMapV().Index("1"))

// 	assert.Equal(t, 0, NewSliceOfMapV("1", "2", "3").Index("1"))
// 	assert.Equal(t, 1, NewSliceOfMapV("1", "2", "3").Index("2"))
// 	assert.Equal(t, 2, NewSliceOfMapV("1", "2", "3").Index("3"))
// 	assert.Equal(t, -1, NewSliceOfMapV("1", "2", "3").Index("4"))
// 	assert.Equal(t, -1, NewSliceOfMapV("1", "2", "3").Index("5"))

// 	// Conversion
// 	{
// 		assert.Equal(t, 1, NewSliceOfMapV("1", "2", "3").Index(Object{2}))
// 		assert.Equal(t, 1, NewSliceOfMapV("1", "2", "3").Index("2"))
// 		assert.Equal(t, -1, NewSliceOfMapV("1", "2", "3").Index(true))
// 		assert.Equal(t, 2, NewSliceOfMapV("1", "2", "3").Index(Char('3')))
// 	}
// }

// // Insert
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Insert_Go(t *testing.B) {
// 	src := []string{}
// 	for _, x := range RangeString(nines6) {
// 		src = append(src, x)
// 		copy(src[1:], src[1:])
// 		src[0] = x
// 	}
// }

// func BenchmarkSliceOfMap_Insert_Slice(t *testing.B) {
// 	slice := NewSliceOfMapV()
// 	for x := range RangeString(nines6) {
// 		slice.Insert(0, x)
// 	}
// }

// func ExampleSliceOfMap_Insert() {
// 	slice := NewSliceOfMapV("1", "3")
// 	fmt.Println(slice.Insert(1, "2"))
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_Insert(t *testing.T) {

// 	// append
// 	{
// 		slice := NewSliceOfMapV()
// 		assert.Equal(t, NewSliceOfMapV("0"), slice.Insert(-1, "0"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1"), slice.Insert(-1, "1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2"), slice.Insert(-1, "2"))
// 	}

// 	// [] append
// 	{
// 		slice := NewSliceOfMapV()
// 		assert.Equal(t, NewSliceOfMapV("0"), slice.Insert(-1, []string{"0"}))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2"), slice.Insert(-1, []string{"1", "2"}))
// 	}

// 	// prepend
// 	{
// 		slice := NewSliceOfMapV()
// 		assert.Equal(t, NewSliceOfMapV("2"), slice.Insert(0, "2"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Insert(0, "1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2"), slice.Insert(0, "0"))
// 	}

// 	// [] prepend
// 	{
// 		slice := NewSliceOfMapV()
// 		assert.Equal(t, NewSliceOfMapV("2"), slice.Insert(0, []string{"2"}))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2"), slice.Insert(0, []string{"0", "1"}))
// 	}

// 	// middle pos
// 	{
// 		slice := NewSliceOfMapV("0", "5")
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "5"), slice.Insert(1, "1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "5"), slice.Insert(2, "2"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "3", "5"), slice.Insert(3, "3"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "3", "4", "5"), slice.Insert(4, "4"))
// 	}

// 	// [] middle pos
// 	{
// 		slice := NewSliceOfMapV("0", "5")
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "5"), slice.Insert(1, []string{"1", "2"}))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "3", "4", "5"), slice.Insert(3, []string{"3", "4"}))
// 	}

// 	// middle neg
// 	{
// 		slice := NewSliceOfMapV("0", "5")
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "5"), slice.Insert(-2, "1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "5"), slice.Insert(-2, "2"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "3", "5"), slice.Insert(-2, "3"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2", "3", "4", "5"), slice.Insert(-2, "4"))
// 	}

// 	// [] middle neg
// 	{
// 		slice := NewSliceOfMapV(0, 5)
// 		assert.Equal(t, NewSliceOfMapV(0, 1, 2, 5), slice.Insert(-2, []string{"1", "2"}))
// 		assert.Equal(t, NewSliceOfMapV(0, "1", "2", "3", 4, 5), slice.Insert(-2, []int{3, 4}))
// 	}

// 	// error cases
// 	{
// 		var slice *SliceOfMap
// 		assert.False(t, slice.Insert(0, 0).Nil())
// 		assert.Equal(t, NewSliceOfMapV("0"), slice.Insert(0, "0"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1"), NewSliceOfMapV("0", "1").Insert(-10, "1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1"), NewSliceOfMapV("0", "1").Insert(10, "1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1"), NewSliceOfMapV("0", "1").Insert(2, "1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1"), NewSliceOfMapV("0", "1").Insert(-3, "1"))
// 	}

// 	// [] error cases
// 	{
// 		var slice *SliceOfMap
// 		assert.False(t, slice.Insert(0, 0).Nil())
// 		assert.Equal(t, NewSliceOfMapV(0), slice.Insert(0, 0))
// 		assert.Equal(t, NewSliceOfMapV(0, 1), NewSliceOfMapV(0, 1).Insert(-10, 1))
// 		assert.Equal(t, NewSliceOfMapV(0, 1), NewSliceOfMapV(0, 1).Insert(10, 1))
// 		assert.Equal(t, NewSliceOfMapV(0, 1), NewSliceOfMapV(0, 1).Insert(2, 1))
// 		assert.Equal(t, NewSliceOfMapV(0, 1), NewSliceOfMapV(0, 1).Insert(-3, 1))
// 	}

// 	// Conversion
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV(1, 3).Insert(1, Object{2}))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV(1, 3).Insert(1, "2"))
// 		assert.Equal(t, NewSliceOfMapV(true, "2", "3"), NewSliceOfMapV(2, 3).Insert(0, true))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV(1, 2).Insert(-1, Char('3')))
// 	}

// 	// [] Conversion
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", 4), NewSliceOfMapV(1, 4).Insert(1, []Object{{2}, {3}}))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", 4), NewSliceOfMapV(1, 4).Insert(1, []string{"2", "3"}))
// 		assert.Equal(t, NewSliceOfMapV(false, true, "2", "3"), NewSliceOfMapV(2, 3).Insert(0, []bool{false, true}))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", 4), NewSliceOfMapV(1, 2).Insert(-1, []Char{'3', '4'}))
// 	}
// }

// // Join
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Join_Go(t *testing.B) {
// 	src := RangeString(nines4)
// 	strings.Join(src, ",")
// }

// func BenchmarkSliceOfMap_Join_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines4))
// 	slice.Join()
// }

// func ExampleSliceOfMap_Join() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Join())
// 	// Output: 1,2,3
// }

// func TestSliceOfMap_Join(t *testing.T) {
// 	// nil
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, Obj(""), slice.Join())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, Obj(""), NewSliceOfMapV().Join())
// 	}

// 	assert.Equal(t, "1,2,3", NewSliceOfMapV("1", "2", "3").Join().O())
// 	assert.Equal(t, "1.2.3", NewSliceOfMapV("1", "2", "3").Join(".").O())
// }

// // Last
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Last_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		_ = src[len(src)-1]
// 		src = src[:len(src)-1]
// 	}
// }

// func BenchmarkSliceOfMap_Last_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.Last()
// 		slice.DropLast()
// 	}
// }

// func ExampleSliceOfMap_Last() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Last())
// 	// Output: 3
// }

// func TestSliceOfMap_Last(t *testing.T) {
// 	// invalid
// 	assert.Equal(t, Obj(nil), NewSliceOfMapV().Last())

// 	// int
// 	assert.Equal(t, Obj("3"), NewSliceOfMapV("2", "3").Last())
// 	assert.Equal(t, Obj("2"), NewSliceOfMapV("3", "2").Last())
// 	assert.Equal(t, Obj("2"), NewSliceOfMapV("1", "3", "2").Last())
// }

// // LastN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_LastN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	_ = src[0:10]
// }

// func BenchmarkSliceOfMap_LastN_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	slice.LastN(10)
// }

// func ExampleSliceOfMap_LastN() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.LastN(2))
// 	// Output: [2 3]
// }

// func TestSliceOfMap_LastN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.LastN(1))
// 		assert.Equal(t, NewSliceOfMapV(), slice.LastN(-1))
// 	}

// 	// Get none
// 	{
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3").LastN(0))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewSliceOfMapV("1", "2", "3")
// 		result := original.LastN(2).Set(0, "0")
// 		assert.Equal(t, NewSliceOfMapV("1", "0", "3"), original)
// 		assert.Equal(t, NewSliceOfMapV("0", "3"), result)
// 	}

// 	// slice full array includeing out of bounds
// 	{
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().LastN(1))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().LastN(10))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").LastN(10))
// 		assert.Equal(t, NewSliceOfMap([]string{"1", "2", "3"}), NewSliceOfMap([]string{"1", "2", "3"}).LastN(10))
// 	}

// 	// grab a few diff
// 	{
// 		assert.Equal(t, NewSliceOfMapV("3"), NewSliceOfMapV("1", "2", "3").LastN(1))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").LastN(2))
// 	}
// }

// // Len
// //--------------------------------------------------------------------------------------------------
// func ExampleSliceOfMap_Len() {
// 	fmt.Println(NewSliceOfMapV("1", "2", "3").Len())
// 	// Output: 3
// }

// func TestSliceOfMap_Len(t *testing.T) {
// 	assert.Equal(t, 0, NewSliceOfMapV().Len())
// 	assert.Equal(t, 2, len(*(NewSliceOfMapV("1", "2"))))
// 	assert.Equal(t, 2, NewSliceOfMapV("1", "2").Len())
// }

// // Less
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Less_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if i+1 < len(src) {
// 			_ = src[i] < src[i+1]
// 		}
// 	}
// }

// func BenchmarkSliceOfMap_Less_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Less(i, i+1)
// 		}
// 	}
// }

// func ExampleSliceOfMap_Less() {
// 	slice := NewSliceOfMapV("2", "3", "1")
// 	fmt.Println(slice.Less(0, 2))
// 	// Output: false
// }

// func TestSliceOfMap_Less(t *testing.T) {

// 	// invalid cases
// 	{
// 		var slice *SliceOfMap
// 		assert.False(t, slice.Less(0, 0))

// 		slice = NewSliceOfMapV()
// 		assert.False(t, slice.Less(0, 0))
// 		assert.False(t, slice.Less(1, 2))
// 		assert.False(t, slice.Less(-1, 2))
// 		assert.False(t, slice.Less(1, -2))
// 	}

// 	// valid
// 	assert.Equal(t, true, NewSliceOfMapV("0", "1", "2").Less(0, 1))
// 	assert.Equal(t, false, NewSliceOfMapV("0", "1", "2").Less(1, 0))
// 	assert.Equal(t, true, NewSliceOfMapV("0", "1", "2").Less(1, 2))
// }

// // Nil
// //--------------------------------------------------------------------------------------------------
// func ExampleSliceOfMap_Nil() {
// 	var slice *SliceOfMap
// 	fmt.Println(slice.Nil())
// 	// Output: true
// }

// func TestSliceOfMap_Nil(t *testing.T) {
// 	var slice *SliceOfMap
// 	assert.True(t, slice.Nil())
// 	assert.False(t, NewSliceOfMapV().Nil())
// 	assert.False(t, NewSliceOfMapV("1", "2", "3").Nil())
// }

// // O
// //--------------------------------------------------------------------------------------------------
// func ExampleSliceOfMap_O() {
// 	fmt.Println(NewSliceOfMapV("1", "2", "3"))
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_O(t *testing.T) {
// 	assert.Equal(t, []string{}, (*SliceOfMap)(nil).O())
// 	assert.Equal(t, []string{}, NewSliceOfMapV().O())
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3"))
// }

// // Pair
// //--------------------------------------------------------------------------------------------------

// func ExampleSliceOfMap_Pair() {
// 	slice := NewSliceOfMapV("1", "2")
// 	first, second := slice.Pair()
// 	fmt.Println(first, second)
// 	// Output: 1 2
// }

// func TestSliceOfMap_Pair(t *testing.T) {

// 	// nil
// 	{
// 		first, second := (*SliceOfMap)(nil).Pair()
// 		assert.Equal(t, Obj(nil), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}

// 	// two values
// 	{
// 		first, second := NewSliceOfMapV("1", "2").Pair()
// 		assert.Equal(t, Obj("1"), first)
// 		assert.Equal(t, Obj("2"), second)
// 	}

// 	// one value
// 	{
// 		first, second := NewSliceOfMapV("1").Pair()
// 		assert.Equal(t, Obj("1"), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}

// 	// no values
// 	{
// 		first, second := NewSliceOfMapV().Pair()
// 		assert.Equal(t, Obj(nil), first)
// 		assert.Equal(t, Obj(nil), second)
// 	}
// }

// // Pop
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Pop_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkSliceOfMap_Pop_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.Pop()
// 	}
// }

// func ExampleSliceOfMap_Pop() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Pop())
// 	// Output: 3
// }

// func TestSliceOfMap_Pop(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, Obj(nil), slice.Pop())
// 	}

// 	// take all one at a time
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, Obj("3"), slice.Pop())
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice)
// 		assert.Equal(t, Obj("2"), slice.Pop())
// 		assert.Equal(t, NewSliceOfMapV("1"), slice)
// 		assert.Equal(t, Obj("1"), slice.Pop())
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 		assert.Equal(t, Obj(nil), slice.Pop())
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}
// }

// // PopN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_PopN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkSliceOfMap_PopN_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.PopN(10)
// 	}
// }

// func ExampleSliceOfMap_PopN() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.PopN(2))
// 	// Output: [2 3]
// }

// func TestSliceOfMap_PopN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.PopN(1))
// 	}

// 	// take none
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV(), slice.PopN(0))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 	}

// 	// take 1
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("3"), slice.PopN(1))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice)
// 	}

// 	// take 2
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), slice.PopN(2))
// 		assert.Equal(t, NewSliceOfMapV("1"), slice)
// 	}

// 	// take 3
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice.PopN(3))
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}

// 	// take beyond
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice.PopN(4))
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}
// }

// // Prepend
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Prepend_Go(t *testing.B) {
// 	src := []string{}
// 	for _, x := range RangeString(nines6) {
// 		src = append(src, x)
// 		copy(src[1:], src[1:])
// 		src[0] = x
// 	}
// }

// func BenchmarkSliceOfMap_Prepend_Slice(t *testing.B) {
// 	slice := NewSliceOfMapV()
// 	for _, x := range RangeString(nines6) {
// 		slice.Prepend(x)
// 	}
// }

// func ExampleSliceOfMap_Prepend() {
// 	slice := NewSliceOfMapV("2", "3")
// 	fmt.Println(slice.Prepend("1"))
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_Prepend(t *testing.T) {

// 	// happy path
// 	{
// 		slice := NewSliceOfMapV()
// 		assert.Equal(t, NewSliceOfMapV("2"), slice.Prepend("2"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Prepend("1"))
// 		assert.Equal(t, NewSliceOfMapV("0", "1", "2"), slice.Prepend("0"))
// 	}

// 	// error cases
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV("0"), slice.Prepend("0"))
// 	}
// }

// // Reverse
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Reverse_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
// 		src[i], src[j] = src[j], src[i]
// 	}
// }

// func BenchmarkSliceOfMap_Reverse_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.Reverse()
// }

// func ExampleSliceOfMap_Reverse() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Reverse())
// 	// Output: [3 2 1]
// }

// func TestSliceOfMap_Reverse(t *testing.T) {

// 	// nil
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.Reverse())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Reverse())
// 	}

// 	// pos
// 	{
// 		slice := NewSliceOfMapV("3", "2", "1")
// 		reversed := slice.Reverse()
// 		assert.Equal(t, NewSliceOfMapV("3", "2", "1", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), reversed)
// 	}

// 	// neg
// 	{
// 		slice := NewSliceOfMapV("2", "3", "-2", "-3")
// 		reversed := slice.Reverse()
// 		assert.Equal(t, NewSliceOfMapV("2", "3", "-2", "-3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("-3", "-2", "3", "2"), reversed)
// 	}
// }

// // ReverseM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_ReverseM_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
// 		src[i], src[j] = src[j], src[i]
// 	}
// }

// func BenchmarkSliceOfMap_ReverseM_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.ReverseM()
// }

// func ExampleSliceOfMap_ReverseM() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.ReverseM())
// 	// Output: [3 2 1]
// }

// func TestSliceOfMap_ReverseM(t *testing.T) {

// 	// nil
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.ReverseM())
// 	}

// 	// empty
// 	{
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().ReverseM())
// 	}

// 	// pos
// 	{
// 		slice := NewSliceOfMapV("3", "2", "1")
// 		reversed := slice.ReverseM()
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), reversed)
// 	}

// 	// neg
// 	{
// 		slice := NewSliceOfMapV("2", "3", "-2", "-3")
// 		reversed := slice.ReverseM()
// 		assert.Equal(t, NewSliceOfMapV("-3", "-2", "3", "2", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("-3", "-2", "3", "2", "4"), reversed)
// 	}
// }

// // Select
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Select_Go(t *testing.B) {
// 	even := []string{}
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if Obj(src[i]).ToInt()%2 == 0 {
// 			even = append(even, src[i])
// 		}
// 	}
// }

// func BenchmarkSliceOfMap_Select_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.Select(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	})
// }

// func ExampleSliceOfMap_Select() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Select(func(x O) bool {
// 		return ExB(x.(string) == "2" || x.(string) == "3")
// 	}))
// 	// Output: [2 3]
// }

// func TestSliceOfMap_Select(t *testing.T) {

// 	// Select all odd values
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.Select(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 != 0)
// 		})
// 		slice.DropFirst()
// 		assert.Equal(t, NewSliceOfMapV("2", "3", "4", "5", "6", "7", "8", "9"), slice)
// 		assert.Equal(t, NewSliceOfMapV("1", "3", "5", "7", "9"), new)
// 	}

// 	// Select all even values
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.Select(func(x O) bool {
// 			return ExB(Obj(x).ToInt()%2 == 0)
// 		})
// 		slice.DropAt(1)
// 		assert.Equal(t, NewSliceOfMapV("1", "3", "4", "5", "6", "7", "8", "9"), slice)
// 		assert.Equal(t, NewSliceOfMapV("2", "4", "6", "8"), new)
// 	}
// }

// // Set
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Set_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		src[i] = "0"
// 	}
// }

// func BenchmarkSliceOfMap_Set_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.Set(i, "0")
// 	}
// }

// func ExampleSliceOfMap_Set() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Set(0, "0"))
// 	// Output: [0 2 3]
// }

// func TestSliceOfMap_Set(t *testing.T) {
// 	assert.Equal(t, NewSliceOfMapV("0", "2", "3"), NewSliceOfMapV("1", "2", "3").Set(0, "0"))
// 	assert.Equal(t, NewSliceOfMapV("1", "0", "3"), NewSliceOfMapV("1", "2", "3").Set(1, "0"))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "0"), NewSliceOfMapV("1", "2", "3").Set(2, "0"))
// 	assert.Equal(t, NewSliceOfMapV("0", "2", "3"), NewSliceOfMapV("1", "2", "3").Set(-3, "0"))
// 	assert.Equal(t, NewSliceOfMapV("1", "0", "3"), NewSliceOfMapV("1", "2", "3").Set(-2, "0"))
// 	assert.Equal(t, NewSliceOfMapV("1", "2", "0"), NewSliceOfMapV("1", "2", "3").Set(-1, "0"))

// 	// Test out of bounds
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").Set(5, "1"))
// 	}

// 	// Test wrong type
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").Set(5, "1"))
// 	}

// 	// Conversion
// 	{
// 		assert.Equal(t, NewSliceOfMapV(0, 2, 0), NewSliceOfMapV(0, 0, 0).Set(1, Object{2}))
// 		assert.Equal(t, NewSliceOfMapV(0, 2, 0), NewSliceOfMapV(0, 0, 0).Set(1, "2"))
// 		assert.Equal(t, NewSliceOfMapV(true, 0, 0), NewSliceOfMapV(0, 0, 0).Set(0, true))
// 		assert.Equal(t, NewSliceOfMapV(0, 0, 3), NewSliceOfMapV(0, 0, 0).Set(-1, Char('3')))
// 	}
// }

// // SetE
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_SetE_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		src[i] = "0"
// 	}
// }

// func BenchmarkSliceOfMap_SetE_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		slice.SetE(i, "0")
// 	}
// }

// func ExampleSliceOfMap_SetE() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.SetE(0, "0"))
// 	// Output: [0 2 3] <nil>
// }

// func TestSliceOfMap_SetE(t *testing.T) {

// 	// pos - begining
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		result, err := slice.SetE(0, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("0", "2", "3"), slice)
// 		assert.Equal(t, NewSliceOfMapV("0", "2", "3"), result)

// 		// multiple
// 		result, err = slice.SetE(0, []string{"4", "5", "6"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), slice)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), result)

// 		// multiple over
// 		result, err = slice.SetE(0, []string{"4", "5", "6", "7"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), slice)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), result)
// 	}

// 	// pos - middle
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		result, err := slice.SetE(1, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("1", "0", "3"), slice)
// 		assert.Equal(t, NewSliceOfMapV("1", "0", "3"), result)
// 	}

// 	// pos - end
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		result, err := slice.SetE(2, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "0"), slice)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "0"), result)
// 	}

// 	// neg - begining
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		result, err := slice.SetE(-3, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("0", "2", "3"), slice)
// 		assert.Equal(t, NewSliceOfMapV("0", "2", "3"), result)

// 		// multiple
// 		result, err = slice.SetE(-3, []string{"4", "5", "6"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), slice)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), result)

// 		// multiple over
// 		result, err = slice.SetE(-3, []string{"4", "5", "6", "7"})
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), slice)
// 		assert.Equal(t, NewSliceOfMapV("4", "5", "6"), result)
// 	}

// 	// neg - middle
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		result, err := slice.SetE(-2, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("1", "0", "3"), slice)
// 		assert.Equal(t, NewSliceOfMapV("1", "0", "3"), result)
// 	}

// 	// neg - end
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		result, err := slice.SetE(-1, "0")
// 		assert.Nil(t, err)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "0"), slice)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "0"), result)
// 	}

// 	// Test out of bounds
// 	{
// 		slice, err := NewSliceOfMapV("1", "2", "3").SetE(5, "1")
// 		assert.NotNil(t, slice)
// 		assert.NotNil(t, err)
// 	}

// 	// Test wrong type
// 	{
// 		slice, err := NewSliceOfMapV("1", "2", "3").SetE(5, "1")
// 		assert.NotNil(t, slice)
// 		assert.NotNil(t, err)
// 	}
// }

// // Shift
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Shift_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 1 {
// 		src = src[1:]
// 	}
// }

// func BenchmarkSliceOfMap_Shift_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.Shift()
// 	}
// }

// func ExampleSliceOfMap_Shift() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Shift())
// 	// Output: 1
// }

// func TestSliceOfMap_Shift(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, Obj(nil), slice.Shift())
// 	}

// 	// take all and beyond
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, Obj("1"), slice.Shift())
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), slice)
// 		assert.Equal(t, Obj("2"), slice.Shift())
// 		assert.Equal(t, NewSliceOfMapV("3"), slice)
// 		assert.Equal(t, Obj("3"), slice.Shift())
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 		assert.Equal(t, Obj(nil), slice.Shift())
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}
// }

// // ShiftN
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_ShiftN_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	for len(src) > 10 {
// 		src = src[10:]
// 	}
// }

// func BenchmarkSliceOfMap_ShiftN_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 0 {
// 		slice.ShiftN(10)
// 	}
// }

// func ExampleSliceOfMap_ShiftN() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.ShiftN(2))
// 	// Output: [1 2]
// }

// func TestSliceOfMap_ShiftN(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.ShiftN(1))
// 	}

// 	// negative value
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("1"), slice.ShiftN(-1))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), slice)
// 	}

// 	// take none
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV(), slice.ShiftN(0))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 	}

// 	// take 1
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("1"), slice.ShiftN(1))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), slice)
// 	}

// 	// take 2
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.ShiftN(2))
// 		assert.Equal(t, NewSliceOfMapV("3"), slice)
// 	}

// 	// take 3
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice.ShiftN(3))
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}

// 	// take beyond
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice.ShiftN(4))
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}
// }

// // Single
// //--------------------------------------------------------------------------------------------------

// func ExampleSliceOfMap_Single() {
// 	slice := NewSliceOfMapV("1")
// 	fmt.Println(slice.Single())
// 	// Output: true
// }

// func TestSliceOfMap_Single(t *testing.T) {

// 	assert.Equal(t, false, NewSliceOfMapV().Single())
// 	assert.Equal(t, true, NewSliceOfMapV("1").Single())
// 	assert.Equal(t, false, NewSliceOfMapV("1", "2").Single())
// }

// // Slice
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Slice_Go(t *testing.B) {
// 	src := RangeString(nines7)
// 	_ = src[0:len(src)]
// }

// func BenchmarkSliceOfMap_Slice_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	slice.Slice(0, -1)
// }

// func ExampleSliceOfMap_Slice() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Slice(1, -1))
// 	// Output: [2 3]
// }

// func TestSliceOfMap_Slice(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.Slice(0, -1))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Slice(0, -1))
// 	}

// 	// Test that the original is modified when the slice is modified
// 	{
// 		original := NewSliceOfMapV("1", "2", "3")
// 		result := original.Slice(0, -1).Set(0, "0")
// 		assert.Equal(t, NewSliceOfMapV("0", "2", "3"), original)
// 		assert.Equal(t, NewSliceOfMapV("0", "2", "3"), result)
// 	}

// 	// slice full array
// 	{
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Slice(0, -1))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Slice(0, 1))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Slice(0, 5))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").Slice(0, -1))
// 		assert.Equal(t, NewSliceOfMap([]string{"1", "2", "3"}), NewSliceOfMap([]string{"1", "2", "3"}).Slice(0, -1))
// 	}

// 	// out of bounds should be moved in
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1"), NewSliceOfMapV("1").Slice(0, 2))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").Slice(-6, 6))
// 	}

// 	// mutually exclusive
// 	{
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Slice(2, -3))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Slice(0, -5))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Slice(4, -1))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Slice(6, -1))
// 		assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV("1", "2", "3", "4").Slice(3, 2))
// 	}

// 	// singles
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4")
// 		assert.Equal(t, NewSliceOfMapV("4"), slice.Slice(-1, -1))
// 		assert.Equal(t, NewSliceOfMapV("3"), slice.Slice(-2, -2))
// 		assert.Equal(t, NewSliceOfMapV("2"), slice.Slice(-3, -3))
// 		assert.Equal(t, NewSliceOfMapV("1"), slice.Slice(0, 0))
// 		assert.Equal(t, NewSliceOfMapV("1"), slice.Slice(-4, -4))
// 		assert.Equal(t, NewSliceOfMapV("2"), slice.Slice(1, 1))
// 		assert.Equal(t, NewSliceOfMapV("2"), slice.Slice(1, -3))
// 		assert.Equal(t, NewSliceOfMapV("3"), slice.Slice(2, 2))
// 		assert.Equal(t, NewSliceOfMapV("3"), slice.Slice(2, -2))
// 		assert.Equal(t, NewSliceOfMapV("4"), slice.Slice(3, 3))
// 		assert.Equal(t, NewSliceOfMapV("4"), slice.Slice(3, -1))
// 	}

// 	// grab all but first
// 	{
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").Slice(1, -1))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").Slice(1, 2))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").Slice(-2, -1))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").Slice(-2, 2))
// 	}

// 	// grab all but last
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3").Slice(0, -2))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3").Slice(-3, -2))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3").Slice(-3, 1))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2", "3").Slice(0, 1))
// 	}

// 	// grab middle
// 	{
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3", "4").Slice(1, -2))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3", "4").Slice(-3, -2))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3", "4").Slice(-3, 2))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3", "4").Slice(1, 2))
// 	}

// 	// random
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1"), NewSliceOfMapV("1", "2", "3").Slice(0, -3))
// 		assert.Equal(t, NewSliceOfMapV("2", "3"), NewSliceOfMapV("1", "2", "3").Slice(1, 2))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), NewSliceOfMapV("1", "2", "3").Slice(0, 2))
// 	}
// }

// // Sort
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Sort_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.SliceOfMap(src))
// }

// func BenchmarkSliceOfMap_Sort_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.Sort()
// }

// func ExampleSliceOfMap_Sort() {
// 	slice := NewSliceOfMapV("2", "3", "1")
// 	fmt.Println(slice.Sort())
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_Sort(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().Sort())

// 	// pos
// 	{
// 		slice := NewSliceOfMapV("5", "3", "2", "4", "1")
// 		sorted := slice.Sort()
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("5", "3", "2", "4", "1"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewSliceOfMapV("5", "3", "-2", "4", "-1")
// 		sorted := slice.Sort()
// 		assert.Equal(t, NewSliceOfMapV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("5", "3", "-2", "4", "-1"), slice)
// 	}
// }

// // SortM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_SortM_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.SliceOfMap(src))
// }

// func BenchmarkSliceOfMap_SortM_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.SortM()
// }

// func ExampleSliceOfMap_SortM() {
// 	slice := NewSliceOfMapV("2", "3", "1")
// 	fmt.Println(slice.SortM())
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_SortM(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().SortM())

// 	// pos
// 	{
// 		slice := NewSliceOfMapV("5", "3", "2", "4", "1")
// 		sorted := slice.SortM()
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4", "5", "6"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewSliceOfMapV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortM()
// 		assert.Equal(t, NewSliceOfMapV("-1", "-2", "3", "4", "5", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("-1", "-2", "3", "4", "5", "6"), slice)
// 	}
// }

// // SortReverse
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_SortReverse_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.Reverse(sort.SliceOfMap(src)))
// }

// func BenchmarkSliceOfMap_SortReverse_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.SortReverse()
// }

// func ExampleSliceOfMap_SortReverse() {
// 	slice := NewSliceOfMapV("2", "3", "1")
// 	fmt.Println(slice.SortReverse())
// 	// Output: [3 2 1]
// }

// func TestSliceOfMap_SortReverse(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().SortReverse())

// 	// pos
// 	{
// 		slice := NewSliceOfMapV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverse()
// 		assert.Equal(t, NewSliceOfMapV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("5", "3", "2", "4", "1"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewSliceOfMapV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverse()
// 		assert.Equal(t, NewSliceOfMapV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("5", "3", "-2", "4", "-1"), slice)
// 	}
// }

// // SortReverseM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_SortReverseM_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	sort.Sort(sort.Reverse(sort.SliceOfMap(src)))
// }

// func BenchmarkSliceOfMap_SortReverseM_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	slice.SortReverseM()
// }

// func ExampleSliceOfMap_SortReverseM() {
// 	slice := NewSliceOfMapV("2", "3", "1")
// 	fmt.Println(slice.SortReverseM())
// 	// Output: [3 2 1]
// }

// func TestSliceOfMap_SortReverseM(t *testing.T) {

// 	// empty
// 	assert.Equal(t, NewSliceOfMapV(), NewSliceOfMapV().SortReverse())

// 	// pos
// 	{
// 		slice := NewSliceOfMapV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, NewSliceOfMapV("5", "4", "3", "2", "1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("5", "4", "3", "2", "1", "6"), slice)
// 	}

// 	// neg
// 	{
// 		slice := NewSliceOfMapV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, NewSliceOfMapV("5", "4", "3", "-2", "-1", "6"), sorted.Append("6"))
// 		assert.Equal(t, NewSliceOfMapV("5", "4", "3", "-2", "-1", "6"), slice)
// 	}
// }

// // String
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_String_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	_ = fmt.Sprintf("%v", src)
// }

// func BenchmarkSliceOfMap_String_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	_ = slice.String()
// }

// func ExampleSliceOfMap_String() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_String(t *testing.T) {
// 	// nil
// 	assert.Equal(t, "[]", (*SliceOfMap)(nil).String())

// 	// empty
// 	assert.Equal(t, "[]", NewSliceOfMapV().String())

// 	// pos
// 	{
// 		slice := NewSliceOfMapV("5", "3", "2", "4", "1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, "[5 4 3 2 1 6]", sorted.Append("6").String())
// 		assert.Equal(t, "[5 4 3 2 1 6]", slice.String())
// 	}

// 	// neg
// 	{
// 		slice := NewSliceOfMapV("5", "3", "-2", "4", "-1")
// 		sorted := slice.SortReverseM()
// 		assert.Equal(t, "[5 4 3 -2 -1 6]", sorted.Append("6").String())
// 		assert.Equal(t, "[5 4 3 -2 -1 6]", slice.String())
// 	}
// }

// // Swap
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Swap_Go(t *testing.B) {
// 	src := RangeString(nines6)
// 	for i := 0; i < len(src); i++ {
// 		if i+1 < len(src) {
// 			src[i], src[i+1] = src[i+1], src[i]
// 		}
// 	}
// }

// func BenchmarkSliceOfMap_Swap_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Swap(i, i+1)
// 		}
// 	}
// }

// func ExampleSliceOfMap_Swap() {
// 	slice := NewSliceOfMapV("2", "3", "1")
// 	slice.Swap(0, 2)
// 	slice.Swap(1, 2)
// 	fmt.Println(slice)
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_Swap(t *testing.T) {

// 	// invalid cases
// 	{
// 		var slice *SliceOfMap
// 		slice.Swap(0, 0)
// 		assert.Equal(t, (*SliceOfMap)(nil), slice)

// 		slice = NewSliceOfMapV()
// 		slice.Swap(0, 0)
// 		assert.Equal(t, NewSliceOfMapV(), slice)

// 		slice.Swap(1, 2)
// 		assert.Equal(t, NewSliceOfMapV(), slice)

// 		slice.Swap(-1, 2)
// 		assert.Equal(t, NewSliceOfMapV(), slice)

// 		slice.Swap(1, -2)
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}

// 	// normal
// 	{
// 		slice := NewSliceOfMapV("0", "1", "2")
// 		slice.Swap(0, 1)
// 		assert.Equal(t, NewSliceOfMapV("1", "0", "2"), slice)
// 	}
// }

// // Take
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Take_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_Take_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines7))
// 	for slice.Len() > 1 {
// 		slice.Take(1, 10)
// 	}
// }

// func ExampleSliceOfMap_Take() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.Take(0, 1))
// 	// Output: [1 2]
// }

// func TestSliceOfMap_Take(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.Take(0, 1))
// 	}

// 	// invalid
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4")
// 		assert.Equal(t, NewSliceOfMapV(), slice.Take(1))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice)
// 		assert.Equal(t, NewSliceOfMapV(), slice.Take(4, 4))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice)
// 	}

// 	// take 1
// 	{
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1"), slice.Take(0, 0))
// 			assert.Equal(t, NewSliceOfMapV("2", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("2"), slice.Take(1, 1))
// 			assert.Equal(t, NewSliceOfMapV("1", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("3"), slice.Take(2, 2))
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("4"), slice.Take(3, 3))
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("4"), slice.Take(-1, -1))
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("3"), slice.Take(-2, -2))
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("2"), slice.Take(-3, -3))
// 			assert.Equal(t, NewSliceOfMapV("1", "3", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1"), slice.Take(-4, -4))
// 			assert.Equal(t, NewSliceOfMapV("2", "3", "4"), slice)
// 		}
// 	}

// 	// take 2
// 	{
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Take(0, 1))
// 			assert.Equal(t, NewSliceOfMapV("3", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("2", "3"), slice.Take(1, 2))
// 			assert.Equal(t, NewSliceOfMapV("1", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("3", "4"), slice.Take(2, 3))
// 			assert.Equal(t, NewSliceOfMapV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("3", "4"), slice.Take(-2, -1))
// 			assert.Equal(t, NewSliceOfMapV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("2", "3"), slice.Take(-3, -2))
// 			assert.Equal(t, NewSliceOfMapV("1", "4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Take(-4, -3))
// 			assert.Equal(t, NewSliceOfMapV("3", "4"), slice)
// 		}
// 	}

// 	// take 3
// 	{
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice.Take(0, 2))
// 			assert.Equal(t, NewSliceOfMapV("4"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("2", "3", "4"), slice.Take(-3, -1))
// 			assert.Equal(t, NewSliceOfMapV("1"), slice)
// 		}
// 	}

// 	// take everything and beyond
// 	{
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Take())
// 			assert.Equal(t, NewSliceOfMapV(), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Take(0, 3))
// 			assert.Equal(t, NewSliceOfMapV(), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Take(0, -1))
// 			assert.Equal(t, NewSliceOfMapV(), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Take(-4, -1))
// 			assert.Equal(t, NewSliceOfMapV(), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Take(-6, -1))
// 			assert.Equal(t, NewSliceOfMapV(), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Take(0, 10))
// 			assert.Equal(t, NewSliceOfMapV(), slice)
// 		}
// 	}

// 	// move index within bounds
// 	{
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("4"), slice.Take(3, 4))
// 			assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("1", "2", "3", "4")
// 			assert.Equal(t, NewSliceOfMapV("1"), slice.Take(-5, 0))
// 			assert.Equal(t, NewSliceOfMapV("2", "3", "4"), slice)
// 		}
// 	}
// }

// // TakeAt
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_TakeAt_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_TakeAt_Slice(t *testing.B) {
// 	src := RangeString(nines5)
// 	index := RangeString(nines5)
// 	slice := NewSliceOfMap(src)
// 	for i := range index {
// 		slice.TakeAt(i)
// 	}
// }

// func ExampleSliceOfMap_TakeAt() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.TakeAt(1))
// 	// Output: 2
// }

// func TestSliceOfMap_TakeAt(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, Obj(nil), slice.TakeAt(0))
// 	}

// 	// all and more
// 	{
// 		slice := NewSliceOfMapV("0", "1", "2")
// 		assert.Equal(t, Obj("2"), slice.TakeAt(-1))
// 		assert.Equal(t, NewSliceOfMapV("0", "1"), slice)
// 		assert.Equal(t, Obj("1"), slice.TakeAt(-1))
// 		assert.Equal(t, NewSliceOfMapV("0"), slice)
// 		assert.Equal(t, Obj("0"), slice.TakeAt(-1))
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 		assert.Equal(t, Obj(nil), slice.TakeAt(-1))
// 		assert.Equal(t, NewSliceOfMapV(), slice)
// 	}

// 	// take invalid
// 	{
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj(nil), slice.TakeAt(3))
// 			assert.Equal(t, NewSliceOfMapV("0", "1", "2"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj(nil), slice.TakeAt(-4))
// 			assert.Equal(t, NewSliceOfMapV("0", "1", "2"), slice)
// 		}
// 	}

// 	// take last
// 	{
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj("2"), slice.TakeAt(2))
// 			assert.Equal(t, NewSliceOfMapV("0", "1"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj("2"), slice.TakeAt(-1))
// 			assert.Equal(t, NewSliceOfMapV("0", "1"), slice)
// 		}
// 	}

// 	// take middle
// 	{
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj("1"), slice.TakeAt(1))
// 			assert.Equal(t, NewSliceOfMapV("0", "2"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj("1"), slice.TakeAt(-2))
// 			assert.Equal(t, NewSliceOfMapV("0", "2"), slice)
// 		}
// 	}

// 	// take first
// 	{
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj("0"), slice.TakeAt(0))
// 			assert.Equal(t, NewSliceOfMapV("1", "2"), slice)
// 		}
// 		{
// 			slice := NewSliceOfMapV("0", "1", "2")
// 			assert.Equal(t, Obj("0"), slice.TakeAt(-3))
// 			assert.Equal(t, NewSliceOfMapV("1", "2"), slice)
// 		}
// 	}
// }

// // TakeW
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_TakeW_Go(t *testing.B) {
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

// func BenchmarkSliceOfMap_TakeW_Slice(t *testing.B) {
// 	slice := NewSliceOfMap(RangeString(nines5))
// 	slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
// }

// func ExampleSliceOfMap_TakeW() {
// 	slice := NewSliceOfMapV("1", "2", "3")
// 	fmt.Println(slice.TakeW(func(x O) bool {
// 		return ExB(Obj(x).ToInt()%2 == 0)
// 	}))
// 	// Output: [2]
// }

// func TestSliceOfMap_TakeW(t *testing.T) {

// 	// take all odd values
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 != 0) })
// 		assert.Equal(t, NewSliceOfMapV("2", "4", "6", "8"), slice)
// 		assert.Equal(t, NewSliceOfMapV("1", "3", "5", "7", "9"), new)
// 	}

// 	// take all even values
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3", "4", "5", "6", "7", "8", "9")
// 		new := slice.TakeW(func(x O) bool { return ExB(Obj(x).ToInt()%2 == 0) })
// 		assert.Equal(t, NewSliceOfMapV("1", "3", "5", "7", "9"), slice)
// 		assert.Equal(t, NewSliceOfMapV("2", "4", "6", "8"), new)
// 	}
// }

// // Union
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Union_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkSliceOfMap_Union_Slice(t *testing.B) {
// 	// slice := NewSliceOfMap(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleSliceOfMap_Union() {
// 	slice := NewSliceOfMapV("1", "2")
// 	fmt.Println(slice.Union([]string{"2", "3"}))
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_Union(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Union(NewSliceOfMapV("1", "2")))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Union([]string{"1", "2"}))
// 	}

// 	// size of one
// 	{
// 		slice := NewSliceOfMapV("1")
// 		union := slice.Union([]string{"1", "2", "3"})
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), union)
// 		assert.Equal(t, NewSliceOfMapV("1"), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewSliceOfMapV("1", "1")
// 		union := slice.Union(NewSliceOfMapV("2", "3"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), union)
// 		assert.Equal(t, NewSliceOfMapV("1", "1"), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "2", "3", "3")
// 		union := slice.Union([]string{"1", "2", "3"})
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), union)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "2", "3", "3"), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		union := slice.Union([]string{"4", "5"})
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4", "5"), union)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2").Union((*[]string)(nil)))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2").Union((*SliceOfMap)(nil)))
// 	}
// }

// // UnionM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_UnionM_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkSliceOfMap_UnionM_Slice(t *testing.B) {
// 	// slice := NewSliceOfMap(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleSliceOfMap_UnionM() {
// 	slice := NewSliceOfMapV("1", "2")
// 	fmt.Println(slice.UnionM([]string{"2", "3"}))
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_UnionM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.UnionM(NewSliceOfMapV("1", "2")))
// 		assert.Equal(t, (*SliceOfMap)(nil), slice)
// 	}

// 	// size of one
// 	{
// 		slice := NewSliceOfMapV("1")
// 		union := slice.UnionM([]string{"1", "2", "3"})
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), union)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewSliceOfMapV("1", "1")
// 		union := slice.UnionM(NewSliceOfMapV("2", "3"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), union)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "2", "3", "3")
// 		union := slice.UnionM([]string{"1", "2", "3"})
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), union)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		union := slice.UnionM([]string{"4", "5"})
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4", "5"), union)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4", "5"), slice)
// 	}

// 	// nils
// 	{
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2").UnionM((*[]string)(nil)))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), NewSliceOfMapV("1", "2").UnionM((*SliceOfMap)(nil)))
// 	}
// }

// // Uniq
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_Uniq_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkSliceOfMap_Uniq_Slice(t *testing.B) {
// 	// slice := NewSliceOfMap(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleSliceOfMap_Uniq() {
// 	slice := NewSliceOfMapV("1", "2", "3", "3")
// 	fmt.Println(slice.Uniq())
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_Uniq(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, NewSliceOfMapV(), slice.Uniq())
// 	}

// 	// size of one
// 	{
// 		slice := NewSliceOfMapV("1")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewSliceOfMapV("1"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewSliceOfMapV("1"), uniq)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewSliceOfMapV("1", "1")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewSliceOfMapV("1"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewSliceOfMapV("1"), uniq)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "2", "3", "3")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "2", "3", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), uniq)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		uniq := slice.Uniq()
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), uniq)
// 	}
// }

// // UniqM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkSliceOfMap_UniqM_Go(t *testing.B) {
// 	// src := RangeStr(nines7)
// 	// for len(src) > 10 {
// 	// 	src = src[10:]
// 	// }
// }

// func BenchmarkSliceOfMap_UniqM_Slice(t *testing.B) {
// 	// slice := NewSliceOfMap(RangeStr(nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleSliceOfMap_UniqM() {
// 	slice := NewSliceOfMapV("1", "2", "3", "3")
// 	fmt.Println(slice.UniqM())
// 	// Output: [1 2 3]
// }

// func TestSliceOfMap_UniqM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *SliceOfMap
// 		assert.Equal(t, (*SliceOfMap)(nil), slice.UniqM())
// 	}

// 	// size of one
// 	{
// 		slice := NewSliceOfMapV("1")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewSliceOfMapV("1"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), uniq)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewSliceOfMapV("1", "1")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewSliceOfMapV("1"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), slice.Append("2"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2"), uniq)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "2", "3", "3")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), uniq)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewSliceOfMapV("1", "2", "3")
// 		uniq := slice.UniqM()
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3"), uniq)
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), slice.Append("4"))
// 		assert.Equal(t, NewSliceOfMapV("1", "2", "3", "4"), uniq)
// 	}
// }

// func RangeString(size int) (new []string) {
// 	for _, x := range Range(0, size) {
// 		new = append(new, string(x))
// 	}
// 	return
// }
