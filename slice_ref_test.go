package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewRefSlice function
//--------------------------------------------------------------------------------------------------
func ExampleNewRefSlice() {
	slice := NewRefSlice([]int{1, 2, 3})
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestRefSlice_NewRefSlice(t *testing.T) {

	// arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, []string{"1", "2"}, NewRefSlice(array).O())

	// empty
	assert.Equal(t, nil, NewRefSlice(nil).O())
	assert.Equal(t, &RefSlice{}, NewRefSlice(nil))
	assert.Equal(t, []int{}, NewRefSlice([]int{}).O())
	assert.Equal(t, []bool{}, NewRefSlice([]bool{}).O())
	assert.Equal(t, []string{}, NewRefSlice([]string{}).O())
	assert.Equal(t, []Object{}, NewRefSlice([]Object{}).O())
	assert.Equal(t, nil, NewRefSlice([]interface{}{}).O())

	// pointers
	var obj *Object
	assert.Equal(t, []*Object{nil}, NewRefSlice(obj).O())
	assert.Equal(t, []*Object{&(Object{"bob"})}, NewRefSlice(&(Object{"bob"})).O())
	assert.Equal(t, []*Object{&(Object{"1"}), &(Object{"2"})}, NewRefSlice([]*Object{&(Object{"1"}), &(Object{"2"})}).O())

	// interface
	assert.Equal(t, nil, NewRefSlice([]interface{}{nil}).O())
	assert.Equal(t, []string{""}, NewRefSlice([]interface{}{nil, ""}).O())
	assert.Equal(t, []bool{true}, NewRefSlice([]interface{}{true}).O())
	assert.Equal(t, []int{1}, NewRefSlice([]interface{}{1}).O())
	assert.Equal(t, []string{""}, NewRefSlice([]interface{}{""}).O())
	assert.Equal(t, []string{"bob"}, NewRefSlice([]interface{}{"bob"}).O())
	assert.Equal(t, []Object{{nil}}, NewRefSlice([]interface{}{Object{}}).O())

	// singles
	assert.Equal(t, []int{1}, NewRefSlice(1).O())
	assert.Equal(t, []bool{true}, NewRefSlice(true).O())
	assert.Equal(t, []string{""}, NewRefSlice("").O())
	assert.Equal(t, []string{"1"}, NewRefSlice("1").O())
	assert.Equal(t, []Object{{1}}, NewRefSlice(Object{1}).O())
	assert.Equal(t, []Object{Object{"bob"}}, NewRefSlice(Object{"bob"}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, NewRefSlice(map[string]string{"1": "one"}).O())

	// slices
	assert.Equal(t, []int{1, 2}, NewRefSlice([]int{1, 2}).O())
	assert.Equal(t, []bool{true}, NewRefSlice([]bool{true}).O())
	assert.Equal(t, []Object{{"bob"}}, NewRefSlice([]Object{{"bob"}}).O())
	assert.Equal(t, []string{"1", "2"}, NewRefSlice([]string{"1", "2"}).O())
	assert.Equal(t, [][]string{{"1"}}, NewRefSlice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, NewRefSlice([]interface{}{map[string]string{"1": "one"}}).O())
}

// NewRefSliceV function
//--------------------------------------------------------------------------------------------------
func ExampleNewRefSliceV_empty() {
	slice := NewRefSliceV()
	fmt.Println(slice.O())
	// Output: <nil>
}

func ExampleNewRefSliceV_variadic() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestRefSlice_NewRefSliceV(t *testing.T) {
	var obj *Object

	// Arrays
	{
		var array [2]string
		array[0] = "1"
		array[1] = "2"
		assert.Equal(t, [][2]string{array}, NewRefSliceV(array).O())
	}

	// Test empty values
	{
		assert.True(t, !NewRefSliceV().Any())
		assert.Equal(t, 0, NewRefSliceV().Len())
		assert.Equal(t, nil, NewRefSliceV().O())
		assert.Equal(t, nil, NewRefSliceV(nil).O())
		assert.Equal(t, &RefSlice{}, NewRefSliceV(nil))
		assert.Equal(t, []string{""}, NewRefSliceV(nil, "").O())
		assert.Equal(t, []*Object{nil}, NewRefSliceV(nil, obj).O())
	}

	// Test pointers
	{
		assert.Equal(t, []*Object{nil}, NewRefSliceV(obj).O())
		assert.Equal(t, []*Object{&(Object{"bob"})}, NewRefSliceV(&(Object{"bob"})).O())
		assert.Equal(t, []*Object{nil}, NewRefSliceV(obj).O())
		assert.Equal(t, []*Object{&(Object{"bob"})}, NewRefSliceV(&(Object{"bob"})).O())
		assert.Equal(t, [][]*Object{{&(Object{"1"}), &(Object{"2"})}}, NewRefSliceV([]*Object{&(Object{"1"}), &(Object{"2"})}).O())
	}

	// Singles
	{
		assert.Equal(t, []int{1}, NewRefSliceV(1).O())
		assert.Equal(t, []string{"1"}, NewRefSliceV("1").O())
		assert.Equal(t, []Object{Object{"bob"}}, NewRefSliceV(Object{"bob"}).O())
		assert.Equal(t, []map[string]string{{"1": "one"}}, NewRefSliceV(map[string]string{"1": "one"}).O())
	}

	// Multiples
	{
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2").O())
		assert.Equal(t, []Object{Object{1}, Object{2}}, NewRefSliceV(Object{1}, Object{2}).O())
	}

	// Test slices
	{
		assert.Equal(t, [][]int{{1, 2}}, NewRefSliceV([]int{1, 2}).O())
		assert.Equal(t, [][]string{{"1"}}, NewRefSliceV([]string{"1"}).O())
	}
}

func TestRefSlice_newEmptySlice(t *testing.T) {

	// Array
	var array [2]string
	array[0] = "1"
	assert.Equal(t, []string{}, newEmptySlice(array).O())

	// Singles
	assert.Equal(t, []int{}, newEmptySlice(1).O())
	assert.Equal(t, []bool{}, newEmptySlice(true).O())
	assert.Equal(t, []string{}, newEmptySlice("").O())
	assert.Equal(t, []string{}, newEmptySlice("bob").O())
	assert.Equal(t, []Object{}, newEmptySlice(Object{1}).O())

	// Slices
	assert.Equal(t, []int{}, newEmptySlice([]int{1, 2}).O())
	assert.Equal(t, []bool{}, newEmptySlice([]bool{true}).O())
	assert.Equal(t, []string{}, newEmptySlice([]string{"bob"}).O())
	assert.Equal(t, []Object{}, newEmptySlice([]Object{{"bob"}}).O())
	assert.Equal(t, [][]string{}, newEmptySlice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []map[string]string{}, newEmptySlice([]interface{}{map[string]string{"1": "one"}}).O())

	// Empty slices
	assert.Equal(t, []int{}, newEmptySlice([]int{}).O())
	assert.Equal(t, []bool{}, newEmptySlice([]bool{}).O())
	assert.Equal(t, []string{}, newEmptySlice([]string{}).O())
	assert.Equal(t, []Object{}, newEmptySlice([]Object{}).O())

	// Interface types
	assert.Equal(t, []interface{}{}, newEmptySlice(nil).O())
	assert.Equal(t, []interface{}{}, newEmptySlice([]interface{}{nil}).O())
	assert.Equal(t, []int{}, newEmptySlice([]interface{}{1}).O())
	assert.Equal(t, []int{}, newEmptySlice([]interface{}{interface{}(1)}).O())
	assert.Equal(t, []string{}, newEmptySlice([]interface{}{""}).O())
	assert.Equal(t, []Object{}, newEmptySlice([]interface{}{Object{}}).O())
}

// Any
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Any_Normal(t *testing.B) {
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

func BenchmarkRefSlice_Any_Optimized(t *testing.B) {
	src := Range(0, nines4)
	slice := NewIntSlice(src)
	for i := range src {
		slice.Any(i)
	}
}

func BenchmarkRefSlice_Any_Reflect(t *testing.B) {
	src := rangeInterObject(0, nines4)
	slice := NewRefSlice(src)
	for _, i := range src {
		slice.Any(i)
	}
}

func ExampleRefSlice_Any_empty() {
	slice := NewRefSliceV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleRefSlice_Any_notEmpty() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleRefSlice_Any_contains() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Any(1))
	// Output: true
}

func ExampleRefSlice_Any_containsAny() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Any(0, 1))
	// Output: true
}

func TestRefSlice_Any(t *testing.T) {
	var slice *RefSlice
	assert.False(t, slice.Any())
	assert.False(t, NewRefSliceV().Any())
	assert.True(t, NewRefSliceV().Append("2").Any())

	// bool
	assert.True(t, NewRefSliceV(false, true).Any(true))
	assert.False(t, NewRefSliceV(true, true).Any(false))
	assert.True(t, NewRefSliceV(true, true).Any(false, true))
	assert.False(t, NewRefSliceV(true, true).Any(false, false))

	// int
	assert.True(t, NewRefSliceV(1, 2, 3).Any(2))
	assert.False(t, NewRefSliceV(1, 2, 3).Any(4))
	assert.True(t, NewRefSliceV(1, 2, 3).Any(4, 3))
	assert.False(t, NewRefSliceV(1, 2, 3).Any(4, 5))

	// int64
	assert.True(t, NewRefSliceV(int64(1), int64(2), int64(3)).Any(int64(2)))
	assert.False(t, NewRefSliceV(int64(1), int64(2), int64(3)).Any(int64(4)))
	assert.True(t, NewRefSliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(2)))
	assert.False(t, NewRefSliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(5)))

	// string
	assert.True(t, NewRefSliceV("1", "2", "3").Any("2"))
	assert.False(t, NewRefSliceV("1", "2", "3").Any("4"))
	assert.True(t, NewRefSliceV("1", "2", "3").Any("4", "2"))
	assert.False(t, NewRefSliceV("1", "2", "3").Any("4", "5"))

	// custom
	assert.True(t, NewRefSliceV(Object{1}, Object{2}).Any(Object{1}))
	assert.False(t, NewRefSliceV(Object{1}, Object{2}).Any(Object{3}))
	assert.True(t, NewRefSliceV(Object{1}, Object{2}).Any(Object{4}, Object{2}))
	assert.False(t, NewRefSliceV(Object{1}, Object{2}).Any(Object{4}, Object{5}))
}

// AnyS
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_AnyS_Normal(t *testing.B) {
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

func BenchmarkRefSlice_AnyS_Optimized(t *testing.B) {
	src := Range(0, nines4)
	slice := NewIntSlice(src)
	for i := range src {
		slice.Any([]int{i})
	}
}

func BenchmarkRefSlice_AnyS_Reflect(t *testing.B) {
	src := rangeInterObject(0, nines4)
	slice := NewRefSlice(src)
	for _, i := range src {
		slice.Any(Object{i})
	}
}

func ExampleRefSlice_AnyS() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.AnyS([]int{0, 1}))
	// Output: true
}

func TestRefSlice_AnyS(t *testing.T) {
	var nilSlice *RefSlice
	assert.False(t, nilSlice.AnyS([]bool{true}))

	// bool
	assert.True(t, NewRefSliceV(true, true).AnyS([]bool{true}))
	assert.True(t, NewRefSliceV(true, true).AnyS([]bool{false, true}))
	assert.False(t, NewRefSliceV(true, true).AnyS([]bool{false, false}))

	// int
	assert.True(t, NewRefSliceV(1, 2, 3).AnyS([]int{1}))
	assert.True(t, NewRefSliceV(1, 2, 3).AnyS([]int{4, 3}))
	assert.False(t, NewRefSliceV(1, 2, 3).AnyS([]int{4, 5}))

	// int64
	assert.True(t, NewRefSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(2)}))
	assert.True(t, NewRefSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(2)}))
	assert.False(t, NewRefSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(5)}))

	// string
	assert.True(t, NewRefSliceV("1", "2", "3").AnyS([]string{"2"}))
	assert.True(t, NewRefSliceV("1", "2", "3").AnyS([]string{"4", "2"}))
	assert.False(t, NewRefSliceV("1", "2", "3").AnyS([]string{"4", "5"}))

	// custom
	assert.True(t, NewRefSliceV(Object{1}, Object{2}).AnyS([]Object{{2}}))
	assert.True(t, NewRefSliceV(Object{1}, Object{2}).AnyS([]Object{{4}, {2}}))
	assert.False(t, NewRefSliceV(Object{1}, Object{2}).AnyS([]Object{{4}, {5}}))
}

// Append
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Append_Normal(t *testing.B) {
	ints := []int{}
	for _, i := range Range(0, nines6) {
		ints = append(ints, i)
	}
}

func BenchmarkRefSlice_Append_Optimized(t *testing.B) {
	slice := NewIntSliceV()
	for _, i := range Range(0, nines6) {
		slice.Append(i)
	}
}

func BenchmarkRefSlice_Append_Reflect(t *testing.B) {
	n := NewRefSlice([]Object{})
	for _, i := range Range(0, nines6) {
		n.Append(Object{i})
	}
}

func ExampleRefSlice_Append() {
	slice := NewRefSliceV(1).Append(2).Append(3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestRefSlice_Append_Reflect(t *testing.T) {

	// Use a custom type to invoke reflection
	n := NewRefSliceV(Object{"1"})
	assert.Equal(t, 1, n.Len())
	assert.Equal(t, false, n.Nil())
	assert.Equal(t, []Object{{"1"}}, n.O())

	// Append another to it
	n.Append(Object{"2"})
	assert.Equal(t, 2, n.Len())
	assert.Equal(t, []Object{{"1"}, {"2"}}, n.O())

	// Given an invalid type which will abort the function so put at end
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'int' to '[]n.Object'", err)
	}()
	n.Append(2)
}

func TestRefSlice_Append(t *testing.T) {

	// nil
	{
		var slice *RefSlice
		assert.Equal(t, []int{0}, slice.Append(0).O())
		assert.Equal(t, (*RefSlice)(nil), slice)
	}

	// Append one back to back
	{
		n := NewRefSliceV()
		assert.Equal(t, 0, n.Len())
		assert.Equal(t, true, n.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		n.Append("1")
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []string{"1"}, n.O())

		// Second append another which will be 2x at most
		n.Append("2")
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []string{"1", "2"}, n.O())
	}

	// Start with just appending without chaining
	{
		n := NewRefSliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1)
		assert.Equal(t, []int{1}, n.O())
		n.Append(2)
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with nil not chained
	{
		n := NewRefSliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Start with nil chained
	{
		n := NewRefSliceV().Append(1).Append(2)
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with non nil
	{
		n := NewRefSliceV(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Use append result directly
	{
		n := NewRefSliceV(1)
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []int{1, 2}, n.Append(2).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewRefSliceV(true)
			assert.Equal(t, []bool{true, false}, n.Append(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := NewRefSliceV(0)
			assert.Equal(t, []int{0, 1}, n.Append(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := NewRefSliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.Append("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewRefSlice([]Object{{"3"}})
			assert.Equal(t, []Object{{"3"}, {"1"}}, n.Append(Object{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}
}

func TestRefSlice_Append_boolTypeError(t *testing.T) {
	n := NewRefSliceV(true)
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'string' to '[]bool'", err)
	}()
	n.Append("2")
}

func TestRefSlice_Append_intTypeError(t *testing.T) {
	n := NewRefSliceV(1)
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'string' to '[]int'", err)
	}()
	n.Append("2")
}

func TestRefSlice_Append_stringTypeError(t *testing.T) {
	n := NewRefSliceV("1")
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'int' to '[]string'", err)
	}()
	n.Append(2)
}

func TestRefSlice_Append_customTypeError(t *testing.T) {
	n := NewRefSliceV(Object{1})
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'int' to '[]n.Object'", err)
	}()
	n.Append(2)
}

// AppendV
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_AppendV_Normal(t *testing.B) {
	ints := []int{}
	ints = append(ints, Range(0, nines6)...)
}

func BenchmarkRefSlice_AppendV_Optimized(t *testing.B) {
	slice := NewIntSliceV()
	new := rangeO(0, nines6)
	slice.AppendV(new...)
}

func BenchmarkRefSlice_AppendV_Reflect(t *testing.B) {
	slice := NewRefSliceV()
	new := rangeInterObject(0, nines6)
	slice.AppendV(new...)
}

func ExampleRefSlice_AppendV() {
	slice := NewRefSliceV(1).AppendV(2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestRefSlice_AppendV(t *testing.T) {

	// nil
	{
		var slice *RefSlice
		assert.Equal(t, []int{0}, slice.AppendV(0).O())
		assert.Equal(t, (*RefSlice)(nil), slice)
	}

	// Append many ints
	{
		slice := NewRefSliceV(1)
		assert.Equal(t, []int{1, 2, 3}, slice.AppendV(2, 3).O())
	}

	// Append many strings
	{
		{
			slice := NewRefSliceV()
			assert.Equal(t, 0, slice.Len())
			assert.Equal(t, []string{"1", "2", "3"}, slice.AppendV("1", "2", "3").O())
			assert.Equal(t, 3, slice.Len())
		}
		{
			slice := NewRefSlice([]string{"1"})
			assert.Equal(t, 1, slice.Len())
			assert.Equal(t, []string{"1", "2", "3"}, slice.AppendV("2", "3").O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// Append to a slice of custom type
	{
		slice := NewRefSlice([]Object{{"3"}})
		assert.Equal(t, []Object{{"3"}, {"1"}}, slice.AppendV(Object{"1"}).O())
		assert.Equal(t, []Object{{"3"}, {"1"}, {"2"}, {"4"}}, slice.AppendV(Object{"2"}, Object{"4"}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewRefSliceV(true)
			assert.Equal(t, []bool{true, false}, n.AppendV(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := NewRefSliceV(0)
			assert.Equal(t, []int{0, 1}, n.AppendV(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := NewRefSliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.AppendV("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewRefSlice([]Object{{"3"}})
			assert.Equal(t, []Object{{"3"}, {"1"}}, n.AppendV(Object{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}

	// Append to a slice of map
	{
		n := NewRefSliceV(map[string]string{"1": "one"})
		expected := []map[string]string{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, n.AppendV(map[string]string{"2": "two"}).O())
	}
}

// At
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_At_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := range ints {
		assert.IsType(t, 0, ints[i])
	}
}

func BenchmarkRefSlice_At_Optimized(t *testing.B) {
	src := Range(0, nines6)
	slice := NewIntSlice(src)
	for _, i := range src {
		_, ok := (slice.At(i).O()).(int)
		assert.True(t, ok)
	}
}

func BenchmarkRefSlice_At_Reflect(t *testing.B) {
	src := rangeInterObject(0, nines6)
	slice := NewRefSlice(src)
	for i := range src {
		_, ok := (slice.At(i).O()).(Object)
		assert.True(t, ok)
	}
}

func ExampleRefSlice_At() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.At(2).O())
	// Output: 3
}

func TestRefSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *RefSlice
		assert.Equal(t, &Object{nil}, nilSlice.At(0))
	}

	// strings
	{
		slice := NewRefSliceV("1", "2", "3", "4")
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
		slice := NewRefSliceV("1")
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------

func ExampleRefSlice_Clear() {
	slice := NewRefSliceV(1).ConcatM([]int{2, 3})
	fmt.Println(slice.Clear().O())
	// Output: []
}

func TestQSlice_Clear(t *testing.T) {

	// nil
	{
		var nilSlice *RefSlice
		assert.Equal(t, &Object{nil}, nilSlice.At(0))
	}

	// bool
	{
		slice := NewRefSliceV(true, false)
		assert.Equal(t, 2, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []bool{}, slice.O())
	}

	// int
	{
		slice := NewRefSliceV(1, 2, 3, 4)
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []int{}, slice.O())
	}

	// string
	{
		slice := NewRefSliceV("1", "2", "3", "4")
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []string{}, slice.O())
	}

	// custom
	{
		slice := NewRefSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, 3, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []Object{}, slice.O())
	}
}

// ConcatM
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_ConcatM_Normal10(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkRefSlice_ConcatM_Normal100(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkRefSlice_ConcatM_Optimized19(t *testing.B) {
	dest := NewIntSliceV()
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.Concat(src[j:i])
		j = i
	}
}

func BenchmarkRefSlice_ConcatM_Optimized100(t *testing.B) {
	dest := NewIntSliceV()
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.ConcatM(src[j:i])
		j = i
	}
}

func BenchmarkRefSlice_ConcatM_Reflect10(t *testing.B) {
	dest := NewRefSliceV()
	src := rangeInterObject(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.ConcatM(src[j:i])
		j = i
	}
}

func BenchmarkRefSlice_ConcatM_Reflect100(t *testing.B) {
	dest := NewRefSliceV()
	src := rangeInterObject(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.ConcatM(src[j:i])
		j = i
	}
}

func ExampleRefSlice_ConcatM() {
	slice := NewRefSliceV(1).ConcatM([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestRefSlice_ConcatM(t *testing.T) {

	// nil
	{
		var slice *RefSlice
		assert.Equal(t, []int{1, 2}, slice.ConcatM([]int{1, 2}).O())
		assert.Equal(t, (*RefSlice)(nil), slice)
	}

	// Append many ints
	{
		n := NewRefSliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.ConcatM([]int{2, 3}).O())
	}

	// Append many strings
	{
		{
			n := NewRefSliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.ConcatM([]string{"1", "2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := NewRefSlice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.ConcatM([]string{"2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := NewRefSlice([]Object{{"3"}})
		assert.Equal(t, []Object{{"3"}, {"1"}}, n.ConcatM([]Object{{"1"}}).O())
		assert.Equal(t, []Object{{"3"}, {"1"}, {"2"}, {"4"}}, n.ConcatM([]Object{{"2"}, {"4"}}).O())
	}

	// Append to a slice of map
	{
		n := NewRefSliceV(map[string]string{"1": "one"})
		expected := []map[string]string{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, n.ConcatM([]map[string]string{{"2": "two"}}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewRefSliceV(true)
			assert.Equal(t, []bool{true, false}, n.ConcatM([]bool{false}).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := NewRefSliceV(0)
			assert.Equal(t, []int{0, 1}, n.ConcatM([]int{1}).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := NewRefSliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.ConcatM([]string{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewRefSlice([]Object{{"3"}})
			assert.Equal(t, []Object{{"3"}, {"1"}}, n.ConcatM([]Object{{"1"}}).O())
			assert.Equal(t, 2, n.Len())
		}
	}
}

func TestRefSlice_ConcatM_notASliceType(t *testing.T) {
	slice := NewRefSliceV(1)
	defer func() {
		err := recover()
		assert.Equal(t, "can't concat non slice type 'string' with '[]int'", err)
	}()
	slice.ConcatM("2")
}

func TestRefSlice_ConcatM_wrongType(t *testing.T) {
	slice := NewRefSliceV(1)
	defer func() {
		err := recover()
		assert.Equal(t, "can't concat type 'string' with '[]int'", err)
	}()
	slice.ConcatM([]string{"2"})
}

// Copy
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Copy_Normal(t *testing.B) {
	ints := Range(0, nines6)
	dst := make([]int, len(ints), len(ints))
	copy(dst, ints)
}

func BenchmarkRefSlice_Copy_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.Copy()
}

func BenchmarkRefSlice_Copy_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines6))
	slice.Copy()
}

func ExampleRefSlice_Copy() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Copy().O())
	// Output: [1 2 3]
}

func TestRefSlice_Copy(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, NewRefSliceV(), nilSlice.Copy(0, -1))
		slice := NewRefSliceV(0).Clear()
		assert.Equal(t, nil, slice.Copy(0, -1).O())
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := NewRefSliceV(1, 2, 3)
		result := original.Copy(0, -1)
		assert.Equal(t, []int{1, 2, 3}, original.O())
		assert.Equal(t, []int{1, 2, 3}, result.O())
		result.Set(0, 0)
		assert.Equal(t, []int{1, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2, 3}, result.O())
	}

	// copy full array
	{
		assert.Equal(t, nil, NewRefSliceV().Copy().O())
		assert.Equal(t, nil, NewRefSliceV().Copy(0, -1).O())
		assert.Equal(t, nil, NewRefSliceV().Copy(0, 1).O())
		assert.Equal(t, nil, NewRefSliceV().Copy(0, 5).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").Copy().O())
		assert.Equal(t, []string{""}, NewRefSliceV("").Copy(0, -1).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").Copy(0, 1).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3).Copy().O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3).Copy(0, -1).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSlice([]int{1, 2, 3}).Copy().O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSlice([]int{1, 2, 3}).Copy(0, -1).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").Copy().O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").Copy(0, 2).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy().O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(0, -1).O())
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, []string{"1"}, NewRefSliceV("1").Copy(0, 2).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false).Copy(-6, 6).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3).Copy(-6, 6).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").Copy(-6, 6).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(-6, 6).O())
	}

	// mutually exclusive
	{
		slice := NewRefSliceV(1, 2, 3, 4)
		assert.Equal(t, nil, slice.Copy(2, -3).O())
		assert.Equal(t, nil, slice.Copy(0, -5).O())
		assert.Equal(t, nil, slice.Copy(4, -1).O())
		assert.Equal(t, nil, slice.Copy(6, -1).O())
		assert.Equal(t, nil, slice.Copy(3, 2).O())
	}

	// singles
	{
		slice := NewRefSliceV(1, 2, 3, 4)
		assert.Equal(t, []int{4}, slice.Copy(-1, -1).O())
		assert.Equal(t, []int{3}, slice.Copy(-2, -2).O())
		assert.Equal(t, []int{2}, slice.Copy(-3, -3).O())
		assert.Equal(t, []int{1}, slice.Copy(0, 0).O())
		assert.Equal(t, []int{1}, slice.Copy(-4, -4).O())
		assert.Equal(t, []int{2}, slice.Copy(1, 1).O())
		assert.Equal(t, []int{2}, slice.Copy(1, -3).O())
		assert.Equal(t, []int{3}, slice.Copy(2, 2).O())
		assert.Equal(t, []int{3}, slice.Copy(2, -2).O())
		assert.Equal(t, []int{4}, slice.Copy(3, 3).O())
		assert.Equal(t, []int{4}, slice.Copy(3, -1).O())
	}

	// grab all but first
	{
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Copy(1, -1).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Copy(1, 2).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Copy(-2, -1).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Copy(-2, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Copy(1, -1).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Copy(1, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Copy(-2, -1).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Copy(-2, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Copy(1, -1).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Copy(1, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Copy(-2, -1).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Copy(-2, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(1, -1).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(1, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(-2, -1).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(-2, 2).O())
	}

	// grab all but last
	{
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Copy(0, -2).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Copy(-3, -2).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Copy(-3, 1).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Copy(0, 1).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Copy(0, -2).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Copy(-3, -2).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Copy(-3, 1).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Copy(0, 1).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Copy(0, -2).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Copy(-3, -2).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Copy(-3, 1).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Copy(0, 1).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(0, -2).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(-3, -2).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(-3, 1).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Copy(0, 1).O())
	}

	// grab middle
	{
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Copy(1, -2).O())
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Copy(-3, -2).O())
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Copy(-3, 2).O())
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Copy(1, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Copy(1, -2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Copy(-3, -2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Copy(-3, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Copy(1, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Copy(1, -2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Copy(-3, -2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Copy(-3, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Copy(1, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Copy(1, -2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Copy(-3, -2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Copy(-3, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Copy(1, 2).O())
	}

	// random
	{
		assert.Equal(t, []string{"1"}, NewRefSliceV("1", "2", "3").Copy(0, -3).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Copy(1, 2).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").Copy(0, 2).O())
	}
}

// Drop
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Drop_Go(t *testing.B) {
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

func BenchmarkRefSlice_Drop_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 1 {
		slice.Drop(1, 10)
	}
}

func BenchmarkRefSlice_Drop_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines5))
	for slice.Len() > 1 {
		slice.Drop(1, 10)
	}
}

func ExampleRefSlice_Drop() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Drop(1, 1).O())
	// Output: [1 3]
}

func TestRefSlice_Drop(t *testing.T) {
	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, (*RefSlice)(nil), slice.Drop(0, 1))
	}

	// int
	{
		// invalid
		assert.Equal(t, []int{1, 2, 3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(1).O())
		assert.Equal(t, []int{1, 2, 3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(4, 4).O())

		// drop 1
		assert.Equal(t, []int{2, 3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(0, 0).O())
		assert.Equal(t, []int{1, 3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(1, 1).O())
		assert.Equal(t, []int{1, 2, 4}, NewRefSliceV(1, 2, 3, 4).Drop(2, 2).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3, 4).Drop(3, 3).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3, 4).Drop(-1, -1).O())
		assert.Equal(t, []int{1, 2, 4}, NewRefSliceV(1, 2, 3, 4).Drop(-2, -2).O())
		assert.Equal(t, []int{1, 3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(-3, -3).O())
		assert.Equal(t, []int{2, 3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(-4, -4).O())

		// drop 2
		assert.Equal(t, []int{3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(0, 1).O())
		assert.Equal(t, []int{1, 4}, NewRefSliceV(1, 2, 3, 4).Drop(1, 2).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3, 4).Drop(2, 3).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3, 4).Drop(-2, -1).O())
		assert.Equal(t, []int{1, 4}, NewRefSliceV(1, 2, 3, 4).Drop(-3, -2).O())
		assert.Equal(t, []int{3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(-4, -3).O())

		// drop 3
		assert.Equal(t, []int{4}, NewRefSliceV(1, 2, 3, 4).Drop(0, 2).O())
		assert.Equal(t, []int{1}, NewRefSliceV(1, 2, 3, 4).Drop(-3, -1).O())

		// drop everything and beyond
		assert.Equal(t, []int{}, NewRefSliceV(1, 2, 3, 4).Drop().O())
		assert.Equal(t, []int{}, NewRefSliceV(1, 2, 3, 4).Drop(0, 3).O())
		assert.Equal(t, []int{}, NewRefSliceV(1, 2, 3, 4).Drop(0, -1).O())
		assert.Equal(t, []int{}, NewRefSliceV(1, 2, 3, 4).Drop(-4, -1).O())
		assert.Equal(t, []int{}, NewRefSliceV(1, 2, 3, 4).Drop(-6, -1).O())
		assert.Equal(t, []int{}, NewRefSliceV(1, 2, 3, 4).Drop(0, 10).O())

		// move index within bounds
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3, 4).Drop(3, 4).O())
		assert.Equal(t, []int{2, 3, 4}, NewRefSliceV(1, 2, 3, 4).Drop(-5, 0).O())
	}

	// int
	{
		// invalid
		assert.Equal(t, []Object{{1}, {2}, {3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(1).O())
		assert.Equal(t, []Object{{1}, {2}, {3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(4, 4).O())

		// drop {1}
		assert.Equal(t, []Object{{2}, {3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 0).O())
		assert.Equal(t, []Object{{1}, {3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(1, 1).O())
		assert.Equal(t, []Object{{1}, {2}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(2, 2).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(3, 3).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-1, -1).O())
		assert.Equal(t, []Object{{1}, {2}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-2, -2).O())
		assert.Equal(t, []Object{{1}, {3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-3, -3).O())
		assert.Equal(t, []Object{{2}, {3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-4, -4).O())

		// drop {2}
		assert.Equal(t, []Object{{3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 1).O())
		assert.Equal(t, []Object{{1}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(1, 2).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(2, 3).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-2, -1).O())
		assert.Equal(t, []Object{{1}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-3, -2).O())
		assert.Equal(t, []Object{{3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-4, -3).O())

		// drop {3}
		assert.Equal(t, []Object{{4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 2).O())
		assert.Equal(t, []Object{{1}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-3, -1).O())

		// drop everything and beyond
		assert.Equal(t, []Object{}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop().O())
		assert.Equal(t, []Object{}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 3).O())
		assert.Equal(t, []Object{}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, -1).O())
		assert.Equal(t, []Object{}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-4, -1).O())
		assert.Equal(t, []Object{}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-6, -1).O())
		assert.Equal(t, []Object{}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 10).O())

		// move index within bounds
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(3, 4).O())
		assert.Equal(t, []Object{{2}, {3}, {4}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-5, 0).O())
	}
}
