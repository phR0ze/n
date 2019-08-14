package n

import (
	"fmt"
	"sort"
	"strings"
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
func BenchmarkRefSlice_Any_Go(t *testing.B) {
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
func BenchmarkRefSlice_AnyS_Go(t *testing.B) {
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

	// NewRefSliceV
	assert.True(t, NewRefSliceV(1, 2).AnyS(NewRefSliceV(1, 3)))
	assert.False(t, NewRefSliceV(1, 2).AnyS(NewRefSliceV(4, 3)))

	// NewIntSliceV
	assert.True(t, NewRefSliceV(1, 2).AnyS(NewIntSliceV(1, 3)))
	assert.False(t, NewRefSliceV(1, 2).AnyS(NewIntSliceV(4, 3)))

	// nils
	assert.False(t, NewRefSliceV(1, 2).AnyS(nil))
	assert.False(t, NewRefSliceV(1, 2).AnyS((*[]int)(nil)))
	assert.False(t, NewRefSliceV(1, 2).AnyS((*IntSlice)(nil)))
	assert.False(t, NewRefSliceV(1, 2).AnyS((*RefSlice)(nil)))
}

// AnyW
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_AnyW_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if i == nines4 {
			break
		}
	}
}

func BenchmarkRefSlice_AnyW_Optimized(t *testing.B) {
	src := Range(0, nines5)
	NewIntSlice(src).AnyW(func(x O) bool {
		return ExB(x.(int) == nines4)
	})
}

func BenchmarkRefSlice_AnyW_Reflect(t *testing.B) {
	src := Range(0, nines5)
	NewRefSlice(src).AnyW(func(x O) bool {
		return ExB(x.(int) == nines4)
	})
}

func ExampleRefSlice_AnyW() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.AnyW(func(x O) bool {
		return ExB(x.(int) == 2)
	}))
	// Output: true
}

func TestRefSlice_AnyW(t *testing.T) {

	// empty
	var slice *RefSlice
	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(int) > 0) }))
	assert.False(t, NewRefSliceV().AnyW(func(x O) bool { return ExB(x.(int) > 0) }))

	// single
	assert.True(t, NewRefSliceV(2).AnyW(func(x O) bool { return ExB(x.(int) > 0) }))

	assert.True(t, NewRefSliceV(1, 2).AnyW(func(x O) bool { return ExB(x.(int) == 2) }))
	assert.True(t, NewRefSliceV(1, 2).AnyW(func(x O) bool { return ExB(x.(int)%2 == 0) }))
	assert.False(t, NewRefSliceV(2, 4).AnyW(func(x O) bool { return ExB(x.(int)%2 != 0) }))
	assert.True(t, NewRefSliceV(1, 2, 3).AnyW(func(x O) bool { return ExB(x.(int) == 4 || x.(int) == 3) }))
}

// Append
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Append_Go(t *testing.B) {
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
func BenchmarkRefSlice_AppendV_Go(t *testing.B) {
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
func BenchmarkRefSlice_At_Go(t *testing.B) {
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
		assert.Equal(t, Obj(nil), nilSlice.At(0))
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

func TestRefSlice_Clear(t *testing.T) {

	// nil
	{
		var nilSlice *RefSlice
		assert.Equal(t, Obj(nil), nilSlice.At(0))
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

// Concat
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Concat_Go10(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkRefSlice_Concat_Go100(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkRefSlice_Concat_Optimized19(t *testing.B) {
	dest := NewIntSliceV()
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.Concat(src[j:i])
		j = i
	}
}

func BenchmarkRefSlice_Concat_Optimized100(t *testing.B) {
	dest := NewIntSliceV()
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.Concat(src[j:i])
		j = i
	}
}

func BenchmarkRefSlice_Concat_Reflect10(t *testing.B) {
	dest := NewRefSliceV()
	src := rangeInterObject(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.Concat(src[j:i])
		j = i
	}
}

func BenchmarkRefSlice_Concat_Reflect100(t *testing.B) {
	dest := NewRefSliceV()
	src := rangeInterObject(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.Concat(src[j:i])
		j = i
	}
}

func ExampleRefSlice_Concat() {
	slice := NewRefSliceV(1).Concat([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestRefSlice_Concat(t *testing.T) {

	// nil
	{
		var slice *RefSlice
		assert.Equal(t, []int{1, 2}, slice.Concat([]int{1, 2}).O())
		assert.Equal(t, (*RefSlice)(nil), slice)
	}

	// Append many ints
	{
		n := NewRefSliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.Concat([]int{2, 3}).O())
	}

	// Append many strings
	{
		{
			n := NewRefSliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.Concat([]string{"1", "2", "3"}).O())
			assert.Equal(t, 0, n.Len())
		}
		{
			n := NewRefSlice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.Concat([]string{"2", "3"}).O())
			assert.Equal(t, 1, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := NewRefSlice([]Object{{"3"}})
		assert.Equal(t, []Object{{"3"}, {"1"}}, n.Concat([]Object{{"1"}}).O())
		assert.Equal(t, []Object{{"3"}, {"2"}, {"4"}}, n.Concat([]Object{{"2"}, {"4"}}).O())
	}

	// Append to a slice of map
	{
		n := NewRefSliceV(map[string]string{"1": "one"})
		expected := []map[string]string{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, n.Concat([]map[string]string{{"2": "two"}}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewRefSliceV(true)
			assert.Equal(t, []bool{true, false}, n.Concat([]bool{false}).O())
			assert.Equal(t, 1, n.Len())
		}

		// int
		{
			n := NewRefSliceV(0)
			assert.Equal(t, []int{0, 1}, n.Concat([]int{1}).O())
			assert.Equal(t, 1, n.Len())
		}

		// string
		{
			n := NewRefSliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.Concat([]string{"1"}).O())
			assert.Equal(t, 1, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewRefSlice([]Object{{"3"}})
			assert.Equal(t, []Object{{"3"}, {"1"}}, n.Concat([]Object{{"1"}}).O())
			assert.Equal(t, 1, n.Len())
		}
	}

	// nils
	{
		assert.Equal(t, []int{1, 2}, NewIntSliceV(1, 2).Concat((*[]int)(nil)).O())
		assert.Equal(t, []int{1, 2}, NewIntSliceV(1, 2).Concat((*IntSlice)(nil)).O())
	}
}

// ConcatM
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_ConcatM_Go10(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkRefSlice_ConcatM_Go100(t *testing.B) {
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

		// NewIntSliceV
		{
			n := NewRefSliceV(0)
			assert.Equal(t, []int{0, 1}, n.ConcatM(NewIntSliceV(1)).O())
			assert.Equal(t, 2, n.Len())
		}

		// NewRefSliceV
		{
			n := NewRefSliceV(0)
			assert.Equal(t, []int{0, 1}, n.ConcatM(NewRefSliceV(1)).O())
			assert.Equal(t, 2, n.Len())
		}

		// NewRefSlice
		{
			n := NewRefSliceV(0)
			assert.Equal(t, []int{0, 1}, n.ConcatM(NewRefSlice([]int{1})).O())
			assert.Equal(t, 2, n.Len())
		}

		// nils
		{
			assert.Equal(t, []int{1, 2}, NewIntSliceV(1, 2).ConcatM((*[]int)(nil)).O())
			assert.Equal(t, []int{1, 2}, NewIntSliceV(1, 2).ConcatM((*IntSlice)(nil)).O())
		}
	}
}

func TestRefSlice_ConcatM_notASliceType(t *testing.T) {
	slice := NewRefSliceV(1)
	defer func() {
		err := recover()
		assert.Equal(t, "can't concat type 'string' with '[]int'", err)
	}()
	slice.ConcatM("2")
}

func TestRefSlice_ConcatM_wrongType(t *testing.T) {
	slice := NewRefSliceV(1)
	defer func() {
		err := recover()
		assert.Equal(t, "can't concat type '[]string' with '[]int'", err)
	}()
	slice.ConcatM([]string{"2"})
}

// Copy
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Copy_Go(t *testing.B) {
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
		assert.Equal(t, []int{}, slice.Copy(0, -1).O())
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
		assert.Equal(t, []int{}, slice.Copy(2, -3).O())
		assert.Equal(t, []int{}, slice.Copy(0, -5).O())
		assert.Equal(t, []int{}, slice.Copy(4, -1).O())
		assert.Equal(t, []int{}, slice.Copy(6, -1).O())
		assert.Equal(t, []int{}, slice.Copy(3, 2).O())
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

// Count
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Count_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if i == nines4 {
			break
		}
	}
}

func BenchmarkRefSlice_Count_Optimized(t *testing.B) {
	src := Range(0, nines5)
	NewIntSlice(src).Count(nines4)
}

func BenchmarkRefSlice_Count_Reflect(t *testing.B) {
	src := Range(0, nines5)
	NewRefSlice(src).Count(nines4)
}

func ExampleRefSlice_Count() {
	slice := NewRefSliceV(1, 2, 2)
	fmt.Println(slice.Count(2))
	// Output: 2
}

func TestRefSlice_Count(t *testing.T) {

	// empty
	var slice *RefSlice
	assert.Equal(t, 0, slice.Count(0))
	assert.Equal(t, 0, NewRefSliceV().Count(0))

	assert.Equal(t, 1, NewRefSliceV(2, 3).Count(2))
	assert.Equal(t, 2, NewRefSliceV(1, 2, 2).Count(2))
	assert.Equal(t, 4, NewRefSliceV(4, 4, 3, 4, 4).Count(4))
	assert.Equal(t, 3, NewRefSliceV(3, 2, 3, 3, 5).Count(3))
	assert.Equal(t, 1, NewRefSliceV(1, 2, 3).Count(3))
}

// CountW
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_CountW_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if i == nines4 {
			break
		}
	}
}

func BenchmarkRefSlice_CountW_Slice(t *testing.B) {
	src := Range(0, nines5)
	NewRefSlice(src).CountW(func(x O) bool {
		return ExB(x.(int) == nines4)
	})
}

func ExampleRefSlice_CountW() {
	slice := NewRefSliceV(1, 2, 2)
	fmt.Println(slice.CountW(func(x O) bool {
		return ExB(x.(int) == 2)
	}))
	// Output: 2
}

func TestRefSlice_CountW(t *testing.T) {

	// empty
	var slice *RefSlice
	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(int) > 0) }))
	assert.Equal(t, 0, NewRefSliceV().CountW(func(x O) bool { return ExB(x.(int) > 0) }))

	assert.Equal(t, 1, NewRefSliceV(2, 3).CountW(func(x O) bool { return ExB(x.(int) > 2) }))
	assert.Equal(t, 1, NewRefSliceV(1, 2).CountW(func(x O) bool { return ExB(x.(int) == 2) }))
	assert.Equal(t, 2, NewRefSliceV(1, 2, 3, 4, 5).CountW(func(x O) bool { return ExB(x.(int)%2 == 0) }))
	assert.Equal(t, 3, NewRefSliceV(1, 2, 3, 4, 5).CountW(func(x O) bool { return ExB(x.(int)%2 != 0) }))
	assert.Equal(t, 1, NewRefSliceV(1, 2, 3).CountW(func(x O) bool { return ExB(x.(int) == 4 || x.(int) == 3) }))
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

// DropAt
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_DropAt_Go(t *testing.B) {
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

func BenchmarkRefSlice_DropAt_Optimized(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := NewIntSlice(src)
	for i := range index {
		slice.DropAt(i)
	}
}

func BenchmarkRefSlice_DropAt_Reflect(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := NewRefSlice(src)
	for i := range index {
		slice.DropAt(i)
	}
}

func ExampleRefSlice_DropAt() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.DropAt(1))
	// Output: [1 3]
}

func TestRefSlice_DropAt(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, (*RefSlice)(nil), slice.DropAt(0))
	}

	// drop all and more
	{
		slice := NewRefSliceV(0, 1, 2)
		assert.Equal(t, []int{0, 1}, slice.DropAt(-1).O())
		assert.Equal(t, []int{0}, slice.DropAt(-1).O())
		assert.Equal(t, []int{}, slice.DropAt(-1).O())
		assert.Equal(t, []int{}, slice.DropAt(-1).O())
	}

	// drop invalid
	assert.Equal(t, []int{0, 1, 2}, NewRefSliceV(0, 1, 2).DropAt(3).O())
	assert.Equal(t, []int{0, 1, 2}, NewRefSliceV(0, 1, 2).DropAt(-4).O())

	// drop last
	assert.Equal(t, []int{0, 1}, NewRefSliceV(0, 1, 2).DropAt(2).O())
	assert.Equal(t, []int{0, 1}, NewRefSliceV(0, 1, 2).DropAt(-1).O())

	// drop middle
	assert.Equal(t, []int{0, 2}, NewRefSliceV(0, 1, 2).DropAt(1).O())
	assert.Equal(t, []int{0, 2}, NewRefSliceV(0, 1, 2).DropAt(-2).O())

	// drop first
	assert.Equal(t, []int{1, 2}, NewRefSliceV(0, 1, 2).DropAt(0).O())
	assert.Equal(t, []int{1, 2}, NewRefSliceV(0, 1, 2).DropAt(-3).O())
}

// DropFirst
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_DropFirst_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkRefSlice_DropFirst_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func BenchmarkRefSlice_DropFirst_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleRefSlice_DropFirst() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.DropFirst().O())
	// Output: [2 3]
}

func TestRefSlice_DropFirst(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, (*RefSlice)(nil), nilSlice.DropFirst())
	}

	// bool
	{
		slice := NewRefSliceV(true, true, false)
		assert.Equal(t, []bool{true, false}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []bool{false}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []bool{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []bool{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}

	// int
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{2, 3}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []int{3}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []int{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []int{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}

	// string
	{
		slice := NewRefSliceV("1", "2", "3")
		assert.Equal(t, []string{"2", "3"}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []string{"3"}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []string{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}

	// custom
	{
		slice := NewRefSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, []Object{{2}, {3}}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []Object{{3}}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []Object{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []Object{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}
}

// DropFirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_DropFirstN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkRefSlice_DropFirstN_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func BenchmarkRefSlice_DropFirstN_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func ExampleRefSlice_DropFirstN() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.DropFirstN(2).O())
	// Output: [3]
}

func TestRefSlice_DropFirstN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, (*RefSlice)(nil), nilSlice.DropFirstN(1))
	}

	// drop none
	{
		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2, 3}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}, {2}, {3}}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := NewRefSliceV(true, true, false)
			assert.Equal(t, []bool{true, false}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{2, 3}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := NewRefSliceV("1", "2", "3")
			assert.Equal(t, []string{"2", "3"}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{2}, {3}}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := NewRefSliceV(true, false, false)
			assert.Equal(t, []bool{false}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{3}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := NewRefSliceV("1", "2", "3")
			assert.Equal(t, []string{"3"}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{3}}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropFirstN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropFirstN(4).O())
			assert.Equal(t, 0, slice.Len())
		}
	}
}

// DropLast
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_DropLast_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkRefSlice_DropLast_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func BenchmarkRefSlice_DropLast_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func ExampleRefSlice_DropLast() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.DropLast().O())
	// Output: [1 2]
}

func TestRefSlice_DropLast(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, (*RefSlice)(nil), nilSlice.DropLast())
	}

	// bool
	{
		slice := NewRefSliceV(true, true, false)
		assert.Equal(t, []bool{true, true}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []bool{true}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []bool{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []bool{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}

	// int
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1, 2}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []int{1}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []int{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []int{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}

	// string
	{
		slice := NewRefSliceV("1", "2", "3")
		assert.Equal(t, []string{"1", "2"}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []string{"1"}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []string{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []string{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}

	// custom
	{
		slice := NewRefSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, []Object{{1}, {2}}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []Object{{1}}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []Object{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []Object{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}
}

// DropLastN
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_DropLastN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkRefSlice_DropLastN_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func BenchmarkRefSlice_DropLastN_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func ExampleRefSlice_DropLastN() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.DropLastN(2).O())
	// Output: [1]
}

func TestRefSlice_DropLastN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, (*RefSlice)(nil), nilSlice.DropLastN(1))
	}

	// drop none
	{
		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2, 3}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}, {2}, {3}}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := NewRefSliceV(true, true, false)
			assert.Equal(t, []bool{true, true}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := NewRefSliceV("1", "2", "3")
			assert.Equal(t, []string{"1", "2"}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}, {2}}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := NewRefSliceV(true, false, false)
			assert.Equal(t, []bool{true}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{1}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := NewRefSliceV("1", "2", "3")
			assert.Equal(t, []string{"1"}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := NewRefSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropLastN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewRefSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropLastN(4).O())
			assert.Equal(t, 0, slice.Len())
		}
	}
}

// Each
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Each_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_Each_Optimized(t *testing.B) {
	NewIntSlice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkRefSlice_Each_Reflect(t *testing.B) {
	NewRefSlice(rangeInterObject(0, nines6)).Each(func(x O) {
		assert.IsType(t, Object{}, x)
	})
}

func ExampleRefSlice_Each() {
	NewRefSliceV(1, 2, 3).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 123
}

func TestRefSlice_Each(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		nilSlice.Each(func(x O) {})
	}

	// int
	{
		NewRefSliceV(1, 2, 3).Each(func(x O) {
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

	// string
	{
		NewRefSliceV("1", "2", "3").Each(func(x O) {
			switch x {
			case "1":
				assert.Equal(t, "1", x)
			case "2":
				assert.Equal(t, "2", x)
			case "3":
				assert.Equal(t, "3", x)
			}
		})
	}

	// custom
	{
		NewRefSlice([]Object{{1}, {2}, {3}}).Each(func(x O) {
			switch x {
			case Object{1}:
				assert.Equal(t, Object{1}, x)
			case Object{2}:
				assert.Equal(t, Object{2}, x)
			case Object{3}:
				assert.Equal(t, Object{3}, x)
			}
		})
	}
}

// EachE
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_EachE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_EachE_Optimized(t *testing.B) {
	NewIntSlice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkRefSlice_EachE_Reflect(t *testing.B) {
	NewRefSlice(rangeInterObject(0, nines6)).Each(func(x O) {
		assert.IsType(t, Object{}, x)
	})
}

func ExampleRefSlice_EachE() {
	NewRefSliceV(1, 2, 3).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 123
}

func TestRefSlice_EachE(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		nilSlice.EachE(func(x O) error {
			return nil
		})
	}

	// int
	{
		NewRefSliceV(1, 2, 3).EachE(func(x O) error {
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

	// string
	{
		NewRefSliceV("1", "2", "3").EachE(func(x O) error {
			switch x {
			case "1":
				assert.Equal(t, "1", x)
			case "2":
				assert.Equal(t, "2", x)
			case "3":
				assert.Equal(t, "3", x)
			}
			return nil
		})
	}

	// custom
	{
		NewRefSlice([]Object{{1}, {2}, {3}}).EachE(func(x O) error {
			switch x {
			case Object{1}:
				assert.Equal(t, Object{1}, x)
			case Object{2}:
				assert.Equal(t, Object{2}, x)
			case Object{3}:
				assert.Equal(t, Object{3}, x)
			}
			return nil
		})
	}
}

// EachI
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_EachI_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_EachI_Optimized(t *testing.B) {
	NewIntSlice(Range(0, nines6)).EachI(func(i int, x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkRefSlice_EachI_Reflect(t *testing.B) {
	NewRefSlice(Range(0, nines6)).EachI(func(i int, x O) {
		assert.IsType(t, 0, x)
	})
}

func ExampleRefSlice_EachI() {
	NewRefSliceV(1, 2, 3).EachI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 0:11:22:3
}

func TestRefSlice_EachI(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		slice.EachI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachI(func(i int, x O) {
			results = append(results, x.(int))
		})
		assert.Equal(t, []int{1, 2, 3}, results)
	}
}

// EachIE
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_EachIE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_EachIE_Optimized(t *testing.B) {
	NewIntSlice(Range(0, nines6)).EachIE(func(i int, x O) error {
		assert.IsType(t, 0, x)
		return nil
	})
}

func BenchmarkRefSlice_EachIE_Reflect(t *testing.B) {
	NewRefSlice(Range(0, nines6)).EachIE(func(i int, x O) error {
		assert.IsType(t, 0, x)
		return nil
	})
}

func ExampleRefSlice_EachIE() {
	NewRefSliceV(1, 2, 3).EachIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 0:11:22:3
}

func TestRefSlice_EachIE(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		slice.EachIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachIE(func(i int, x O) error {
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{1, 2, 3}, results)
	}

	// Break early with error
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachIE(func(i int, x O) error {
			if i == 2 {
				return Break
			}
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{1, 2}, results)
	}
}

// EachR
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_EachR_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_EachR_Optimized(t *testing.B) {
	NewIntSlice(Range(0, nines6)).EachR(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkRefSlice_EachR_Reflect(t *testing.B) {
	NewRefSlice(Range(0, nines6)).EachR(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func ExampleRefSlice_EachR() {
	NewRefSliceV(1, 2, 3).EachR(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 321
}

func TestRefSlice_EachR(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		slice.EachR(func(x O) {})
	}

	// Loop through
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachR(func(x O) {
			results = append(results, x.(int))
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}
}

// EachRE
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_EachRE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_EachRE_Optimized(t *testing.B) {
	NewIntSlice(Range(0, nines6)).EachRE(func(x O) error {
		assert.IsType(t, 0, x)
		return nil
	})
}

func BenchmarkRefSlice_EachRE_Reflect(t *testing.B) {
	NewRefSlice(Range(0, nines6)).EachRE(func(x O) error {
		assert.IsType(t, 0, x)
		return nil
	})
}

func ExampleRefSlice_EachRE() {
	NewRefSliceV(1, 2, 3).EachRE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 321
}

func TestRefSlice_EachRE(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		slice.EachRE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachRE(func(x O) error {
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}

	// Break early with error
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachRE(func(x O) error {
			if x.(int) == 1 {
				return Break
			}
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{3, 2}, results)
	}
}

// EachRI
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_EachRI_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_EachRI_Reflect(t *testing.B) {
	NewIntSlice(Range(0, nines6)).EachRI(func(i int, x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkRefSlice_EachRI_Optimized(t *testing.B) {
	NewRefSlice(Range(0, nines6)).EachRI(func(i int, x O) {
		assert.IsType(t, 0, x)
	})
}

func ExampleRefSlice_EachRI() {
	NewRefSliceV(1, 2, 3).EachRI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 2:31:20:1
}

func TestRefSlice_EachRI(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		slice.EachRI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachRI(func(i int, x O) {
			results = append(results, x.(int))
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}
}

// EachRIE
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_EachRIE_Go(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkRefSlice_EachRIE_Reflect(t *testing.B) {
	NewRefSlice(Range(0, nines6)).EachRIE(func(i int, x O) error {
		assert.IsType(t, 0, x)
		return nil
	})
}

func BenchmarkRefSlice_EachRIE_Optimized(t *testing.B) {
	NewRefSlice(Range(0, nines6)).EachRIE(func(i int, x O) error {
		assert.IsType(t, 0, x)
		return nil
	})
}

func ExampleRefSlice_EachRIE() {
	NewRefSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 2:31:20:1
}

func TestRefSlice_EachRIE(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		slice.EachRIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}

	// Break early with error
	{
		results := []int{}
		NewRefSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
			if i == 0 {
				return Break
			}
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{3, 2}, results)
	}
}

// Empty
//--------------------------------------------------------------------------------------------------
func ExampleRefSlice_Empty() {
	fmt.Println(NewRefSliceV().Empty())
	// Output: true
}

func TestRefSlice_Empty(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, true, nilSlice.Empty())
	}

	assert.Equal(t, true, NewRefSliceV().Empty())
	assert.Equal(t, false, NewRefSliceV(1).Empty())
	assert.Equal(t, false, NewRefSliceV(1, 2, 3).Empty())
	assert.Equal(t, false, NewRefSlice(1).Empty())
	assert.Equal(t, false, NewRefSlice([]int{1, 2, 3}).Empty())
}

// First
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_First_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkRefSlice_First_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.First()
	}
}

func BenchmarkRefSlice_First_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	for slice.Len() > 0 {
		slice.First()
	}
}

func ExampleRefSlice_First() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.First().O())
	// Output: 1
}

func TestRefSlice_First(t *testing.T) {
	// invalid
	{
		assert.Equal(t, Obj(nil), NewRefSliceV().First())
	}

	// bool
	{
		assert.Equal(t, &Object{true}, NewRefSliceV(true, false).First())
		assert.Equal(t, &Object{false}, NewRefSliceV(false, true).First())
	}

	// int
	{
		assert.Equal(t, Obj(2), NewRefSliceV(2, 3).First())
		assert.Equal(t, Obj(3), NewRefSliceV(3, 2).First())
		assert.Equal(t, Obj(1), NewRefSliceV(1, 3, 2).First())
	}

	// string
	{
		assert.Equal(t, &Object{"2"}, NewRefSliceV("2", "3").First())
		assert.Equal(t, &Object{"3"}, NewRefSliceV("3", "2").First())
		assert.Equal(t, &Object{"1"}, NewRefSliceV("1", "3", "2").First())
	}

	// custom
	{
		assert.Equal(t, &Object{Object{2}}, NewRefSlice([]Object{{2}, {3}}).First())
		assert.Equal(t, &Object{Object{3}}, NewRefSlice([]Object{{3}, {2}}).First())
		assert.Equal(t, &Object{Object{1}}, NewRefSlice([]Object{{1}, {3}, {2}}).First())
	}
}

// FirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_FirstN_Go(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:10]
}

func BenchmarkRefSlice_FirstN_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	slice.FirstN(10)
}

func BenchmarkRefSlice_FirstN_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	slice.FirstN(10)
}

func ExampleRefSlice_FirstN() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.FirstN(2).O())
	// Output: [1 2]
}

func TestRefSlice_FirstN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, nil, nilSlice.FirstN(1).O())
		slice := NewRefSliceV(0).Clear()
		assert.Equal(t, []int{}, slice.FirstN(-1).O())
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewRefSliceV(1, 2, 3)
		result := original.FirstN(2).Set(0, 0)
		assert.Equal(t, []int{0, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2}, result.O())
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, nil, NewRefSliceV().FirstN(1).O())
		assert.Equal(t, nil, NewRefSliceV().FirstN(10).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").FirstN(1).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").FirstN(10).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3).FirstN(10).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSlice([]int{1, 2, 3}).FirstN(10).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").FirstN(10).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).FirstN(10).O())
	}

	// grab a few diff
	{
		assert.Equal(t, []bool{true}, NewRefSliceV(true, false, true).FirstN(1).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).FirstN(2).O())
		assert.Equal(t, []int{1}, NewRefSliceV(1, 2, 3).FirstN(1).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).FirstN(2).O())
		assert.Equal(t, []string{"1"}, NewRefSliceV("1", "2", "3").FirstN(1).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").FirstN(2).O())
		assert.Equal(t, []Object{{1}}, NewRefSlice([]Object{{1}, {2}, {3}}).FirstN(1).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).FirstN(2).O())
	}
}

// Generic
//--------------------------------------------------------------------------------------------------
func ExampleRefSlice_Generic() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Generic())
	// Output: true
}

// Index
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Index_Go(t *testing.B) {
	ints := Range(0, nines5)
	for i := range ints {
		if ints[i] == nines4 {
			break
		}
	}
}

func BenchmarkRefSlice_Index_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines5))
	slice.Index(nines4)
}

func BenchmarkRefSlice_Index_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines5))
	slice.Index(nines4)
}

func ExampleRefSlice_Index() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Index(2))
	// Output: 1
}

func TestRefSlice_Index(t *testing.T) {

	// empty
	var slice *RefSlice
	assert.Equal(t, -1, slice.Index(2))
	assert.Equal(t, -1, NewRefSliceV().Index(1))

	assert.Equal(t, 0, NewRefSliceV(1, 2, 3).Index(1))
	assert.Equal(t, 1, NewRefSliceV(1, 2, 3).Index(2))
	assert.Equal(t, 2, NewRefSliceV(1, 2, 3).Index(3))
	assert.Equal(t, -1, NewRefSliceV(1, 2, 3).Index(4))
	assert.Equal(t, -1, NewRefSliceV(1, 2, 3).Index(5))
}

// Insert
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Insert_Go(t *testing.B) {
	ints := []int{}
	for i := range Range(0, nines6) {
		ints = append(ints, i)
		copy(ints[1:], ints[1:])
		ints[0] = i
	}
}

func BenchmarkRefSlice_Insert_Optimized(t *testing.B) {
	slice := NewIntSliceV()
	for i := range Range(0, nines6) {
		slice.Insert(0, i)
	}
}

func BenchmarkRefSlice_Insert_Reflect(t *testing.B) {
	slice := NewRefSliceV()
	for i := range Range(0, nines6) {
		slice.Insert(0, Object{i})
	}
}

func ExampleRefSlice_Insert() {
	slice := NewRefSliceV(1, 3)
	fmt.Println(slice.Insert(1, 2).O())
	// Output: [1 2 3]
}

func TestRefSlice_Insert(t *testing.T) {

	// int
	{
		// append
		{
			slice := NewRefSliceV()
			assert.Equal(t, []int{0}, slice.Insert(-1, 0).O())
			assert.Equal(t, []int{0, 1}, slice.Insert(-1, 1).O())
			assert.Equal(t, []int{0, 1, 2}, slice.Insert(-1, 2).O())
		}

		// prepend
		{
			slice := NewRefSliceV()
			assert.Equal(t, []int{2}, slice.Insert(0, 2).O())
			assert.Equal(t, []int{1, 2}, slice.Insert(0, 1).O())
			assert.Equal(t, []int{0, 1, 2}, slice.Insert(0, 0).O())
		}

		// middle pos
		{
			slice := NewRefSliceV(0, 5)
			assert.Equal(t, []int{0, 1, 5}, slice.Insert(1, 1).O())
			assert.Equal(t, []int{0, 1, 2, 5}, slice.Insert(2, 2).O())
			assert.Equal(t, []int{0, 1, 2, 3, 5}, slice.Insert(3, 3).O())
			assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, slice.Insert(4, 4).O())
		}

		// middle neg
		{
			slice := NewRefSliceV(0, 5)
			assert.Equal(t, []int{0, 1, 5}, slice.Insert(-2, 1).O())
			assert.Equal(t, []int{0, 1, 2, 5}, slice.Insert(-2, 2).O())
			assert.Equal(t, []int{0, 1, 2, 3, 5}, slice.Insert(-2, 3).O())
			assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, slice.Insert(-2, 4).O())
		}

		// error cases
		{
			var slice *RefSlice
			assert.False(t, slice.Insert(0, 0).Nil())
			assert.Equal(t, []int{0}, slice.Insert(0, 0).O())
			assert.Equal(t, []int{0, 1}, NewRefSliceV(0, 1).Insert(-10, 1).O())
			assert.Equal(t, []int{0, 1}, NewRefSliceV(0, 1).Insert(10, 1).O())
			assert.Equal(t, []int{0, 1}, NewRefSliceV(0, 1).Insert(2, 1).O())
			assert.Equal(t, []int{0, 1}, NewRefSliceV(0, 1).Insert(-3, 1).O())
		}
	}

	// custom
	{
		// append
		{
			slice := NewRefSliceV()
			assert.Equal(t, []Object{{0}}, slice.Insert(-1, Object{0}).O())
			assert.Equal(t, []Object{{0}, {1}}, slice.Insert(-1, Object{1}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.Insert(-1, Object{2}).O())
		}

		// prepend
		{
			slice := NewRefSliceV()
			assert.Equal(t, []Object{{2}}, slice.Insert(0, Object{2}).O())
			assert.Equal(t, []Object{{1}, {2}}, slice.Insert(0, Object{1}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.Insert(0, Object{0}).O())
		}

		// middle pos
		{
			slice := NewRefSlice([]Object{{0}, {5}})
			assert.Equal(t, []Object{{0}, {1}, {5}}, slice.Insert(1, Object{1}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}, {5}}, slice.Insert(2, Object{2}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}, {3}, {5}}, slice.Insert(3, Object{3}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}, {3}, {4}, {5}}, slice.Insert(4, Object{4}).O())
		}

		// middle neg
		{
			slice := NewRefSlice([]Object{{0}, {5}})
			assert.Equal(t, []Object{{0}, {1}, {5}}, slice.Insert(-2, Object{1}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}, {5}}, slice.Insert(-2, Object{2}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}, {3}, {5}}, slice.Insert(-2, Object{3}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}, {3}, {4}, {5}}, slice.Insert(-2, Object{4}).O())
		}

		// error cases
		{
			var slice *RefSlice
			assert.False(t, slice.Insert(0, Object{0}).Nil())
			assert.Equal(t, []Object{{0}}, slice.Insert(0, Object{0}).O())
			assert.Equal(t, []Object{{0}, {1}}, NewRefSlice([]Object{{0}, {1}}).Insert(-10, 1).O())
			assert.Equal(t, []Object{{0}, {1}}, NewRefSlice([]Object{{0}, {1}}).Insert(10, 1).O())
			assert.Equal(t, []Object{{0}, {1}}, NewRefSlice([]Object{{0}, {1}}).Insert(2, 1).O())
			assert.Equal(t, []Object{{0}, {1}}, NewRefSlice([]Object{{0}, {1}}).Insert(-3, 1).O())
		}
	}
}

func TestRefSlice_Insert_wrongType(t *testing.T) {
	slice := NewRefSliceV(1)
	defer func() {
		err := recover()
		assert.Equal(t, "can't insert type 'string' into '[]int'", err)
	}()
	slice.Insert(0, "2")
}

// Join
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Join_Go(t *testing.B) {
	ints := Range(0, nines4)
	strs := []string{}
	for i := 0; i < len(ints); i++ {
		strs = append(strs, fmt.Sprintf("%v", ints[i]))
	}
	strings.Join(strs, ",")
}

func BenchmarkRefSlice_Join_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines4))
	slice.Join()
}

func BenchmarkRefSlice_Join_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines4))
	slice.Join()
}

func ExampleRefSlice_Join() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Join())
	// Output: 1,2,3
}

func TestRefSlice_Join(t *testing.T) {
	// nil
	{
		var slice *RefSlice
		assert.Equal(t, Obj(""), slice.Join())
	}

	// empty
	{
		assert.Equal(t, Obj(""), NewRefSliceV().Join())
	}

	assert.Equal(t, "1,2,3", NewRefSliceV(1, 2, 3).Join().O())
	assert.Equal(t, "1.2.3", NewRefSliceV(1, 2, 3).Join(".").O())
}

// Last
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Last_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		_ = ints[len(ints)-1]
		ints = ints[:len(ints)-1]
	}
}

func BenchmarkRefSlice_Last_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Last()
		slice.DropLast()
	}
}

func BenchmarkRefSlice_Last_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Last()
		slice.DropLast()
	}
}
func ExampleRefSlice_Last() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Last())
	// Output: 3
}

func TestRefSlice_Last(t *testing.T) {

	// invalid
	{
		assert.Equal(t, Obj(nil), NewRefSliceV().Last())
	}

	// int
	{
		assert.Equal(t, Obj(3), NewRefSliceV(2, 3).Last())
		assert.Equal(t, Obj(2), NewRefSliceV(3, 2).Last())
		assert.Equal(t, Obj(2), NewRefSliceV(1, 3, 2).Last())
	}

	// object
	{
		assert.Equal(t, &Object{Object{3}}, NewRefSlice([]Object{{2}, {3}}).Last())
		assert.Equal(t, &Object{Object{2}}, NewRefSlice([]Object{{3}, {2}}).Last())
		assert.Equal(t, &Object{Object{2}}, NewRefSlice([]Object{{1}, {3}, {2}}).Last())
	}
}

// LastN
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_LastN_Go(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:10]
}

func BenchmarkRefSlice_LastN_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	slice.LastN(10)
}

func BenchmarkRefSlice_LastN_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	slice.LastN(10)
}

func ExampleRefSlice_LastN() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.LastN(2).O())
	// Output: [2 3]
}

func TestRefSlice_LastN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, nil, nilSlice.LastN(1).O())
		slice := NewRefSliceV(0).Clear()
		assert.Equal(t, []int{}, slice.LastN(-1).O())
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewRefSliceV(1, 2, 3)
		result := original.LastN(2).Set(0, 0)
		assert.Equal(t, []int{1, 0, 3}, original.O())
		assert.Equal(t, []int{0, 3}, result.O())
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, nil, NewRefSliceV().LastN(1).O())
		assert.Equal(t, nil, NewRefSliceV().LastN(10).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").LastN(1).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").LastN(10).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3).LastN(10).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSlice([]int{1, 2, 3}).LastN(10).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").LastN(10).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).LastN(10).O())
	}

	// grab a few diff
	{
		assert.Equal(t, []bool{false}, NewRefSliceV(true, true, false).LastN(1).O())
		assert.Equal(t, []bool{false}, NewRefSliceV(true, true, false).LastN(-1).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).LastN(2).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).LastN(-2).O())
		assert.Equal(t, []int{3}, NewRefSliceV(1, 2, 3).LastN(1).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).LastN(2).O())
		assert.Equal(t, []string{"3"}, NewRefSliceV("1", "2", "3").LastN(1).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").LastN(2).O())
		assert.Equal(t, []Object{{3}}, NewRefSlice([]Object{{1}, {2}, {3}}).LastN(1).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).LastN(2).O())
	}
}

// Len
//--------------------------------------------------------------------------------------------------
func TestRefSlice_Len(t *testing.T) {
	assert.Equal(t, 0, NewRefSliceV().Len())
	assert.Equal(t, 1, NewRefSliceV().Append("2").Len())
}

// Less
//--------------------------------------------------------------------------------------------------
// func BenchmarkRefSlice_Less_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i++ {
// 		if i+1 < len(ints) {
// 			_ = ints[i] < ints[i+1]
// 		}
// 	}
// }

// func BenchmarkRefSlice_Less_Optimized(t *testing.B) {
// 	slice := NewIntSlice(Range(0, nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Less(i, i+1)
// 		}
// 	}
// }

// func BenchmarkRefSlice_Less_Reflect(t *testing.B) {
// 	slice := NewRefSlice(rangeInterObject(0, nines6))
// 	for i := 0; i < slice.Len(); i++ {
// 		if i+1 < slice.Len() {
// 			slice.Less(i, i+1)
// 		}
// 	}
// }

func ExampleRefSlice_Less() {
	slice := NewRefSliceV(2, 3, 1)
	fmt.Println(slice.Sort().O())
	// Output: [1 2 3]
}

func TestRefSlice_Less(t *testing.T) {

	// invalid cases
	{
		var slice *RefSlice
		assert.False(t, slice.Less(0, 0))
		slice = NewRefSliceV()
		assert.False(t, slice.Less(0, 0))
		assert.False(t, slice.Less(1, 2))
		assert.False(t, slice.Less(-1, 2))
		assert.False(t, slice.Less(1, -2))
	}

	// // bool
	// {
	// 	assert.Equal(t, false, NewRefSliceV(true, false, true).Less(0, 1))
	// 	assert.Equal(t, true, NewRefSliceV(true, false, true).Less(1, 0))
	// }

	// int
	{
		assert.Equal(t, true, NewRefSliceV(0, 1, 2).Less(0, 1))
		assert.Equal(t, false, NewRefSliceV(0, 1, 2).Less(1, 0))
		assert.Equal(t, true, NewRefSliceV(0, 1, 2).Less(1, 2))
	}

	// // string
	// {
	// 	assert.Equal(t, true, NewRefSliceV("0", "1", "2").Less(0, 1))
	// 	assert.Equal(t, false, NewRefSliceV("0", "1", "2").Less(1, 0))
	// 	assert.Equal(t, true, NewRefSliceV("0", "1", "2").Less(1, 2))
	// }

	// // custom
	// {
	// 	assert.Equal(t, true, NewRefSlice([]Object{{0}, {1}, {2}}).Less(0, 1))
	// 	assert.Equal(t, false, NewRefSlice([]Object{{0}, {1}, {2}}).Less(1, 0))
	// 	assert.Equal(t, true, NewRefSlice([]Object{{0}, {1}, {2}}).Less(1, 2))
	// }
}

// Nil
//--------------------------------------------------------------------------------------------------
func TestRefSlice_Nil(t *testing.T) {
	assert.True(t, NewRefSliceV().Nil())
	var q *RefSlice
	assert.True(t, q.Nil())
	assert.False(t, NewRefSliceV().Append("2").Nil())
}

// O
//--------------------------------------------------------------------------------------------------
func TestRefSlice_O(t *testing.T) {
	assert.Nil(t, NewRefSliceV().O())
	assert.Len(t, NewRefSliceV().Append("2").O(), 1)
}

// Pair
//--------------------------------------------------------------------------------------------------

func ExampleRefSlice_Pair() {
	slice := NewRefSliceV(1, 2)
	first, second := slice.Pair()
	fmt.Println(first.O(), second.O())
	// Output: 1 2
}

func TestRefSlice_Pair(t *testing.T) {

	// nil
	{
		first, second := (*RefSlice)(nil).Pair()
		assert.Equal(t, Obj(nil), first)
		assert.Equal(t, Obj(nil), second)
	}

	// int
	{
		// two values
		{
			first, second := NewRefSliceV(1, 2).Pair()
			assert.Equal(t, Obj(1), first)
			assert.Equal(t, Obj(2), second)
		}

		// one value
		{
			first, second := NewRefSliceV(1).Pair()
			assert.Equal(t, Obj(1), first)
			assert.Equal(t, Obj(nil), second)
		}

		// no values
		{
			first, second := NewRefSliceV().Pair()
			assert.Equal(t, Obj(nil), first)
			assert.Equal(t, Obj(nil), second)
		}
	}

	// custom
	{
		// two values
		{
			first, second := NewRefSlice([]Object{{1}, {2}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, &Object{Object{2}}, second)
		}

		// one value
		{
			first, second := NewRefSlice([]Object{{1}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, Obj(nil), second)
		}

		// no values
		{
			first, second := NewRefSliceV().Pair()
			assert.Equal(t, Obj(nil), first)
			assert.Equal(t, Obj(nil), second)
		}
	}
}

// Pop
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Pop_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkRefSlice_Pop_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Pop()
	}
}

func BenchmarkRefSlice_Pop_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Pop()
	}
}

func ExampleRefSlice_Pop() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Pop())
	// Output: 3
}

func TestRefSlice_Pop(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, Obj(nil), slice.Pop())
	}

	// take all one at a time
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, Obj(3), slice.Pop())
		assert.Equal(t, []int{1, 2}, slice.O())
		assert.Equal(t, Obj(2), slice.Pop())
		assert.Equal(t, []int{1}, slice.O())
		assert.Equal(t, Obj(1), slice.Pop())
		assert.Equal(t, []int{}, slice.O())
		assert.Equal(t, Obj(nil), slice.Pop())
		assert.Equal(t, []int{}, slice.O())
	}
}

// PopN
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_PopN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkRefSlice_PopN_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.PopN(10)
	}
}

func BenchmarkRefSlice_PopN_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.PopN(10)
	}
}

func ExampleRefSlice_PopN() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.PopN(2))
	// Output: [2 3]
}

func TestRefSlice_PopN(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, NewRefSliceV(), slice.PopN(1))
	}

	// take none
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{}, slice.PopN(0).O())
		assert.Equal(t, []int{1, 2, 3}, slice.O())
	}

	// take 1
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{3}, slice.PopN(1).O())
		assert.Equal(t, []int{1, 2}, slice.O())
	}

	// take 2
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{2, 3}, slice.PopN(2).O())
		assert.Equal(t, []int{1}, slice.O())
	}

	// take 3
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1, 2, 3}, slice.PopN(3).O())
		assert.Equal(t, []int{}, slice.O())
	}

	// take beyond
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1, 2, 3}, slice.PopN(4).O())
		assert.Equal(t, []int{}, slice.O())
	}
}

// Prepend
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Prepend_Go(t *testing.B) {
	ints := []int{}
	for i := range Range(0, nines6) {
		ints = append(ints, i)
		copy(ints[1:], ints[1:])
		ints[0] = i
	}
}

func BenchmarkRefSlice_Prepend_Optimized(t *testing.B) {
	slice := NewIntSliceV()
	for i := range Range(0, nines6) {
		slice.Prepend(i)
	}
}

func BenchmarkRefSlice_Prepend_Reflect(t *testing.B) {
	slice := NewRefSliceV()
	for i := range Range(0, nines6) {
		slice.Prepend(Object{i})
	}
}

func ExampleRefSlice_Prepend() {
	slice := NewRefSliceV(2, 3)
	fmt.Println(slice.Prepend(1).O())
	// Output: [1 2 3]
}

func TestRefSlice_Prepend(t *testing.T) {

	// int
	{
		// happy path
		{
			slice := NewRefSliceV()
			assert.Equal(t, []int{2}, slice.Prepend(2).O())
			assert.Equal(t, []int{1, 2}, slice.Prepend(1).O())
			assert.Equal(t, []int{0, 1, 2}, slice.Prepend(0).O())
		}

		// error cases
		{
			var slice *RefSlice
			assert.False(t, slice.Prepend(0).Nil())
			assert.Equal(t, []int{0}, slice.Prepend(0).O())
		}
	}

	// custom
	{
		// prepend
		{
			slice := NewRefSliceV()
			assert.Equal(t, []Object{{2}}, slice.Prepend(Object{2}).O())
			assert.Equal(t, []Object{{1}, {2}}, slice.Prepend(Object{1}).O())
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.Prepend(Object{0}).O())
		}

		// error cases
		{
			var slice *RefSlice
			assert.False(t, slice.Prepend(Object{0}).Nil())
			assert.Equal(t, []Object{{0}}, slice.Prepend(Object{0}).O())
		}
	}
}

// Reverse
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Reverse_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

func BenchmarkRefSlice_Reverse_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.Reverse()
}

func BenchmarkRefSlice_Reverse_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	slice.Reverse()
}

func ExampleRefSlice_Reverse() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Reverse())
	// Output: [3 2 1]
}

func TestRefSlice_Reverse(t *testing.T) {

	// nil
	{
		var slice *RefSlice
		assert.Equal(t, NewRefSliceV(), slice.Reverse())
	}

	// empty
	{
		assert.Equal(t, nil, NewRefSliceV().Reverse().O())
	}

	// pos
	{
		slice := NewRefSliceV(3, 2, 1)
		reversed := slice.Reverse()
		assert.Equal(t, []int{3, 2, 1, 4}, slice.Append(4).O())
		assert.Equal(t, []int{1, 2, 3}, reversed.O())
	}

	// neg
	{
		slice := NewRefSliceV(2, 3, -2, -3)
		reversed := slice.Reverse()
		assert.Equal(t, []int{2, 3, -2, -3, 4}, slice.Append(4).O())
		assert.Equal(t, []int{-3, -2, 3, 2}, reversed.O())
	}
}

// ReverseM
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_ReverseM_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

func BenchmarkRefSlice_ReverseM_Slice(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.ReverseM()
}

func BenchmarkRefSlice_ReverseM_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	slice.ReverseM()
}

func ExampleRefSlice_ReverseM() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.ReverseM())
	// Output: [3 2 1]
}

func TestRefSlice_ReverseM(t *testing.T) {

	// nil
	{
		var slice *RefSlice
		assert.Equal(t, nil, slice.ReverseM().O())
	}

	// empty
	{
		assert.Equal(t, nil, NewRefSliceV().ReverseM().O())
	}

	// pos
	{
		slice := NewRefSliceV(3, 2, 1)
		reversed := slice.ReverseM()
		assert.Equal(t, []int{1, 2, 3, 4}, slice.Append(4).O())
		assert.Equal(t, []int{1, 2, 3, 4}, reversed.O())
	}

	// neg
	{
		slice := NewRefSliceV(2, 3, -2, -3)
		reversed := slice.ReverseM()
		assert.Equal(t, []int{-3, -2, 3, 2, 4}, slice.Append(4).O())
		assert.Equal(t, []int{-3, -2, 3, 2, 4}, reversed.O())
	}
}

// Select
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Select_Go(t *testing.B) {
	even := []int{}
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		if ints[i]%2 == 0 {
			even = append(even, ints[i])
		}
	}
}

func BenchmarkRefSlice_Select_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	slice.Select(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	})
}

func BenchmarkRefSlice_Select_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	slice.Select(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	})
}

func ExampleRefSlice_Select() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Select(func(x O) bool {
		return ExB(x.(int) == 2 || x.(int) == 3)
	}))
	// Output: [2 3]
}

func TestRefSlice_Select(t *testing.T) {

	// Select all odd values
	{
		slice := NewRefSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.Select(func(x O) bool {
			return ExB(x.(int)%2 != 0)
		})
		slice.DropFirst()
		assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9}, slice.O())
		assert.Equal(t, []int{1, 3, 5, 7, 9}, new.O())
	}

	// Select all even values
	{
		slice := NewRefSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.Select(func(x O) bool {
			return ExB(x.(int)%2 == 0)
		})
		slice.DropAt(1)
		assert.Equal(t, []int{1, 3, 4, 5, 6, 7, 8, 9}, slice.O())
		assert.Equal(t, []int{2, 4, 6, 8}, new.O())
	}
}

// Set
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Set_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		ints[i] = 0
	}
}

func BenchmarkRefSlice_Set_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, 0)
	}
}

func BenchmarkRefSlice_Set_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, Object{0})
	}
}

func ExampleRefSlice_Set() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Set(0, 0).O())
	// Output: [0 2 3]
}

func TestRefSlice_Set(t *testing.T) {
	// bool
	{
		assert.Equal(t, []bool{false, true, true}, NewRefSliceV(true, true, true).Set(0, false).O())
		assert.Equal(t, []bool{true, false, true}, NewRefSliceV(true, true, true).Set(1, false).O())
		assert.Equal(t, []bool{true, true, false}, NewRefSliceV(true, true, true).Set(2, false).O())
		assert.Equal(t, []bool{false, true, true}, NewRefSliceV(true, true, true).Set(-3, false).O())
		assert.Equal(t, []bool{true, false, true}, NewRefSliceV(true, true, true).Set(-2, false).O())
		assert.Equal(t, []bool{true, true, false}, NewRefSliceV(true, true, true).Set(-1, false).O())
	}

	// int
	{
		assert.Equal(t, []int{0, 2, 3}, NewRefSliceV(1, 2, 3).Set(0, 0).O())
		assert.Equal(t, []int{1, 0, 3}, NewRefSliceV(1, 2, 3).Set(1, 0).O())
		assert.Equal(t, []int{1, 2, 0}, NewRefSliceV(1, 2, 3).Set(2, 0).O())
		assert.Equal(t, []int{0, 2, 3}, NewRefSliceV(1, 2, 3).Set(-3, 0).O())
		assert.Equal(t, []int{1, 0, 3}, NewRefSliceV(1, 2, 3).Set(-2, 0).O())
		assert.Equal(t, []int{1, 2, 0}, NewRefSliceV(1, 2, 3).Set(-1, 0).O())
	}

	// string
	{
		assert.Equal(t, []string{"0", "2", "3"}, NewRefSliceV("1", "2", "3").Set(0, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, NewRefSliceV("1", "2", "3").Set(1, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, NewRefSliceV("1", "2", "3").Set(2, "0").O())
		assert.Equal(t, []string{"0", "2", "3"}, NewRefSliceV("1", "2", "3").Set(-3, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, NewRefSliceV("1", "2", "3").Set(-2, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, NewRefSliceV("1", "2", "3").Set(-1, "0").O())
	}

	// custom
	{
		assert.Equal(t, []Object{{0}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Set(0, Object{0}).O())
		assert.Equal(t, []Object{{1}, {0}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Set(1, Object{0}).O())
		assert.Equal(t, []Object{{1}, {2}, {0}}, NewRefSlice([]Object{{1}, {2}, {3}}).Set(2, Object{0}).O())
		assert.Equal(t, []Object{{0}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Set(-3, Object{0}).O())
		assert.Equal(t, []Object{{1}, {0}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Set(-2, Object{0}).O())
		assert.Equal(t, []Object{{1}, {2}, {0}}, NewRefSlice([]Object{{1}, {2}, {3}}).Set(-1, Object{0}).O())
	}
}

// SetE
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_SetE_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		ints[i] = 0
	}
}

func BenchmarkRefSlice_SetE_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.SetE(i, 0)
	}
}

func BenchmarkRefSlice_SetE_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.SetE(i, Object{0})
	}
}

func ExampleRefSlice_SetE() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.SetE(0, 0))
	// Output: [0 2 3] <nil>
}

func TestRefSlice_SetE(t *testing.T) {

	// Error cases
	{
		// out of bounds
		slice, err := NewRefSliceV("1", "2", "3").SetE(5, "0")
		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
		assert.Equal(t, "slice assignment is out of bounds", err.Error())

		// wrong type
		slice, err = NewRefSliceV("1", "2", "3").SetE(0, 1)
		assert.Equal(t, []string{"1", "2", "3"}, slice.O())
		assert.Equal(t, "can't set type 'int' in '[]string'", err.Error())
	}

	// string
	{
		slice, err := NewRefSliceV("1", "2", "3").SetE(0, "0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"0", "2", "3"}, slice.O())

		slice, err = NewRefSliceV("1", "2", "3").SetE(1, "0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "0", "3"}, slice.O())

		slice, err = NewRefSliceV("1", "2", "3").SetE(2, "0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "2", "0"}, slice.O())

		slice, err = NewRefSliceV("1", "2", "3").SetE(-3, "0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"0", "2", "3"}, slice.O())

		slice, err = NewRefSliceV("1", "2", "3").SetE(-2, "0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "0", "3"}, slice.O())

		slice, err = NewRefSliceV("1", "2", "3").SetE(-1, "0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "2", "0"}, slice.O())
	}

	// custom
	{
		slice, err := NewRefSlice([]Object{{1}, {2}, {3}}).SetE(0, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []Object{{0}, {2}, {3}}, slice.O())

		slice, err = NewRefSlice([]Object{{1}, {2}, {3}}).SetE(1, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []Object{{1}, {0}, {3}}, slice.O())

		slice, err = NewRefSlice([]Object{{1}, {2}, {3}}).SetE(2, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []Object{{1}, {2}, {0}}, slice.O())

		slice, err = NewRefSlice([]Object{{1}, {2}, {3}}).SetE(-3, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []Object{{0}, {2}, {3}}, slice.O())

		slice, err = NewRefSlice([]Object{{1}, {2}, {3}}).SetE(-2, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []Object{{1}, {0}, {3}}, slice.O())

		slice, err = NewRefSlice([]Object{{1}, {2}, {3}}).SetE(-1, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []Object{{1}, {2}, {0}}, slice.O())
	}
}

// Shift
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Shift_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkRefSlice_Shift_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Shift()
	}
}

func BenchmarkRefSlice_Shift_Relfect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.Shift()
	}
}

func ExampleRefSlice_Shift() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Shift().O())
	// Output: 1
}

func TestRefSlice_Shift(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, Obj(nil), slice.Shift())
	}

	// take all and beyond
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, Obj(1), slice.Shift())
		assert.Equal(t, []int{2, 3}, slice.O())
		assert.Equal(t, Obj(2), slice.Shift())
		assert.Equal(t, []int{3}, slice.O())
		assert.Equal(t, Obj(3), slice.Shift())
		assert.Equal(t, []int{}, slice.O())
		assert.Equal(t, Obj(nil), slice.Shift())
		assert.Equal(t, []int{}, slice.O())
	}

	// generic: take all and beyond
	{
		slice := NewRefSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, &Object{Object{1}}, slice.Shift())
		assert.Equal(t, []Object{{2}, {3}}, slice.O())
		assert.Equal(t, &Object{Object{2}}, slice.Shift())
		assert.Equal(t, []Object{{3}}, slice.O())
		assert.Equal(t, &Object{Object{3}}, slice.Shift())
		assert.Equal(t, []Object{}, slice.O())
		assert.Equal(t, Obj(nil), slice.Shift())
		assert.Equal(t, []Object{}, slice.O())
	}
}

// ShiftN
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_ShiftN_Go(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkRefSlice_ShiftN_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.ShiftN(10)
	}
}

func BenchmarkRefSlice_ShiftN_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.ShiftN(10)
	}
}

func ExampleRefSlice_ShiftN() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.ShiftN(2).O())
	// Output: [1 2]
}

func TestRefSlice_ShiftN(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, NewRefSliceV(), slice.ShiftN(1))
	}

	// negative value
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1}, slice.ShiftN(-1).O())
		assert.Equal(t, []int{2, 3}, slice.O())
	}

	// take none
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{}, slice.ShiftN(0).O())
		assert.Equal(t, []int{1, 2, 3}, slice.O())
	}

	// take 1
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1}, slice.ShiftN(1).O())
		assert.Equal(t, []int{2, 3}, slice.O())
	}

	// take 2
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1, 2}, slice.ShiftN(2).O())
		assert.Equal(t, []int{3}, slice.O())
	}

	// take 3
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1, 2, 3}, slice.ShiftN(3).O())
		assert.Equal(t, []int{}, slice.O())
	}

	// take beyond
	{
		slice := NewRefSliceV(1, 2, 3)
		assert.Equal(t, []int{1, 2, 3}, slice.ShiftN(4).O())
		assert.Equal(t, []int{}, slice.O())
	}
}

// Single
//--------------------------------------------------------------------------------------------------

func ExampleRefSlice_Single() {
	slice := NewRefSliceV(1)
	fmt.Println(slice.Single())
	// Output: true
}

func TestRefSlice_Single(t *testing.T) {

	// int
	{
		assert.Equal(t, false, NewRefSliceV().Single())
		assert.Equal(t, true, NewRefSliceV(1).Single())
		assert.Equal(t, false, NewRefSliceV(1, 2).Single())
	}

	// custom
	{
		assert.Equal(t, false, NewRefSliceV().Single())
		assert.Equal(t, true, NewRefSliceV(Object{1}).Single())
		assert.Equal(t, false, NewRefSliceV(Object{1}, Object{2}).Single())
	}
}

// Slice
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Slice_Go(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:len(ints)]
}

func BenchmarkRefSlice_Slice_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	slice.Slice(0, -1)
}

func BenchmarkRefSlice_Slice_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines7))
	slice.Slice(0, -1)
}

func ExampleRefSlice_Slice() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Slice(1, -1).O())
	// Output: [2 3]
}

func TestRefSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var nilSlice *RefSlice
		assert.Equal(t, NewRefSliceV(), nilSlice.Slice(0, -1))
		slice := NewRefSliceV(0).Clear()
		assert.Equal(t, []int{}, slice.Slice(0, -1).O())
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewRefSliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, []int{0, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2, 3}, result.O())
	}

	// slice full array
	{
		assert.Equal(t, nil, NewRefSliceV().Slice(0, -1).O())
		assert.Equal(t, nil, NewRefSliceV().Slice(0, 1).O())
		assert.Equal(t, nil, NewRefSliceV().Slice(0, 5).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").Slice(0, -1).O())
		assert.Equal(t, []string{""}, NewRefSliceV("").Slice(0, 1).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3).Slice(0, -1).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSlice([]int{1, 2, 3}).Slice(0, -1).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").Slice(0, 2).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(0, -1).O())
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, []string{"1"}, NewRefSliceV("1").Slice(0, 2).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false).Slice(-6, 6).O())
		assert.Equal(t, []int{1, 2, 3}, NewRefSliceV(1, 2, 3).Slice(-6, 6).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").Slice(-6, 6).O())
		assert.Equal(t, []Object{{1}, {2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(-6, 6).O())
	}

	// mutually exclusive
	{
		slice := NewRefSliceV(1, 2, 3, 4)
		assert.Equal(t, []int{}, slice.Slice(2, -3).O())
		assert.Equal(t, []int{}, slice.Slice(0, -5).O())
		assert.Equal(t, []int{}, slice.Slice(4, -1).O())
		assert.Equal(t, []int{}, slice.Slice(6, -1).O())
		assert.Equal(t, []int{}, slice.Slice(3, 2).O())
	}

	// singles
	{
		slice := NewRefSliceV(1, 2, 3, 4)
		assert.Equal(t, []int{4}, slice.Slice(-1, -1).O())
		assert.Equal(t, []int{3}, slice.Slice(-2, -2).O())
		assert.Equal(t, []int{2}, slice.Slice(-3, -3).O())
		assert.Equal(t, []int{1}, slice.Slice(0, 0).O())
		assert.Equal(t, []int{1}, slice.Slice(-4, -4).O())
		assert.Equal(t, []int{2}, slice.Slice(1, 1).O())
		assert.Equal(t, []int{2}, slice.Slice(1, -3).O())
		assert.Equal(t, []int{3}, slice.Slice(2, 2).O())
		assert.Equal(t, []int{3}, slice.Slice(2, -2).O())
		assert.Equal(t, []int{4}, slice.Slice(3, 3).O())
		assert.Equal(t, []int{4}, slice.Slice(3, -1).O())
	}

	// grab all but first
	{
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Slice(1, -1).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Slice(1, 2).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Slice(-2, -1).O())
		assert.Equal(t, []bool{false, true}, NewRefSliceV(true, false, true).Slice(-2, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Slice(1, -1).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Slice(1, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Slice(-2, -1).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3).Slice(-2, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Slice(1, -1).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Slice(1, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Slice(-2, -1).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Slice(-2, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(1, -1).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(1, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(-2, -1).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(-2, 2).O())
	}

	// grab all but last
	{
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Slice(0, -2).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Slice(-3, -2).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Slice(-3, 1).O())
		assert.Equal(t, []bool{true, false}, NewRefSliceV(true, false, true).Slice(0, 1).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Slice(0, -2).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Slice(-3, -2).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Slice(-3, 1).O())
		assert.Equal(t, []int{1, 2}, NewRefSliceV(1, 2, 3).Slice(0, 1).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Slice(0, -2).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Slice(-3, -2).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Slice(-3, 1).O())
		assert.Equal(t, []string{"1", "2"}, NewRefSliceV("1", "2", "3").Slice(0, 1).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(0, -2).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(-3, -2).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(-3, 1).O())
		assert.Equal(t, []Object{{1}, {2}}, NewRefSlice([]Object{{1}, {2}, {3}}).Slice(0, 1).O())
	}

	// grab middle
	{
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Slice(1, -2).O())
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Slice(-3, -2).O())
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Slice(-3, 2).O())
		assert.Equal(t, []bool{true, true}, NewRefSliceV(false, true, true, false).Slice(1, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Slice(1, -2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Slice(-3, -2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Slice(-3, 2).O())
		assert.Equal(t, []int{2, 3}, NewRefSliceV(1, 2, 3, 4).Slice(1, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Slice(1, -2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Slice(-3, -2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Slice(-3, 2).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3", "4").Slice(1, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, -2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, -2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, 2).O())
		assert.Equal(t, []Object{{2}, {3}}, NewRefSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, 2).O())
	}

	// random
	{
		assert.Equal(t, []string{"1"}, NewRefSliceV("1", "2", "3").Slice(0, -3).O())
		assert.Equal(t, []string{"2", "3"}, NewRefSliceV("1", "2", "3").Slice(1, 2).O())
		assert.Equal(t, []string{"1", "2", "3"}, NewRefSliceV("1", "2", "3").Slice(0, 2).O())
	}
}

// String
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_String_Go(t *testing.B) {
	ints := Range(0, nines6)
	_ = fmt.Sprintf("%v", ints)
}

func BenchmarkRefSlice_String_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	_ = slice.String()
}

func BenchmarkRefSlice_String_Optimized(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	_ = slice.String()
}

func ExampleRefSlice_String() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestRefSlice_String(t *testing.T) {

	assert.Equal(t, "[]", (*RefSlice)(nil).String())
	assert.Equal(t, "[]", NewRefSliceV().String())
	assert.Equal(t, "[5 4 3 2 1]", NewRefSliceV(5, 4, 3, 2, 1).String())
	assert.Equal(t, "[5 4 3 2 1 6]", NewRefSliceV(5, 4, 3, 2, 1).Append(6).String())
}

// Sort
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Sort_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.IntSlice(ints))
}

func BenchmarkRefSlice_Sort_Slice(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	slice.Sort()
}

func ExampleRefSlice_Sort() {
	slice := NewRefSliceV(2, 3, 1)
	fmt.Println(slice.Sort())
	// Output: [1 2 3]
}

func TestRefSlice_Sort(t *testing.T) {

	// empty
	assert.Equal(t, NewRefSliceV(), NewRefSliceV().Sort())

	// pos
	{
		slice := NewRefSliceV(5, 3, 2, 4, 1)
		sorted := slice.Sort()
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{5, 3, 2, 4, 1}, slice.O())
	}

	// neg
	{
		slice := NewRefSliceV(5, 3, -2, 4, -1)
		sorted := slice.Sort()
		assert.Equal(t, []int{-2, -1, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{5, 3, -2, 4, -1}, slice.O())
	}
}

// SortM
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_SortM_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.IntSlice(ints))
}

func BenchmarkRefSlice_SortM_Slice(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	slice.SortM()
}

func ExampleRefSlice_SortM() {
	slice := NewRefSliceV(2, 3, 1)
	fmt.Println(slice.SortM())
	// Output: [1 2 3]
}

func TestRefSlice_SortM(t *testing.T) {

	// empty
	assert.Equal(t, NewRefSliceV(), NewRefSliceV().SortM())

	// pos
	{
		slice := NewRefSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortM()
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, slice.O())
	}

	// neg
	{
		slice := NewRefSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortM()
		assert.Equal(t, []int{-2, -1, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{-2, -1, 3, 4, 5, 6}, slice.O())
	}
}

// SortReverse
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_SortReverse_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}

func BenchmarkRefSlice_SortReverse_Slice(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	slice.SortReverse()
}

func ExampleRefSlice_SortReverse() {
	slice := NewRefSliceV(2, 3, 1)
	fmt.Println(slice.SortReverse())
	// Output: [3 2 1]
}

func TestRefSlice_SortReverse(t *testing.T) {

	// empty
	assert.Equal(t, NewRefSliceV(), NewRefSliceV().SortReverse())

	// pos
	{
		slice := NewRefSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverse()
		assert.Equal(t, []int{5, 4, 3, 2, 1, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{5, 3, 2, 4, 1}, slice.O())
	}

	// neg
	{
		slice := NewRefSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverse()
		assert.Equal(t, []int{5, 4, 3, -1, -2, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{5, 3, -2, 4, -1}, slice.O())
	}
}

// SortReverseM
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_SortReverseM_Go(t *testing.B) {
	ints := Range(0, nines6)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}

func BenchmarkRefSlice_SortReverseM_Slice(t *testing.B) {
	slice := NewRefSlice(Range(0, nines6))
	slice.SortReverseM()
}

func ExampleRefSlice_SortReverseM() {
	slice := NewRefSliceV(2, 3, 1)
	fmt.Println(slice.SortReverseM())
	// Output: [3 2 1]
}

func TestRefSlice_SortReverseM(t *testing.T) {

	// empty
	assert.Equal(t, NewRefSliceV(), NewRefSliceV().SortReverse())

	// pos
	{
		slice := NewRefSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverseM()
		assert.Equal(t, []int{5, 4, 3, 2, 1, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{5, 4, 3, 2, 1, 6}, slice.O())
	}

	// neg
	{
		slice := NewRefSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverseM()
		assert.Equal(t, []int{5, 4, 3, -1, -2, 6}, sorted.Append(6).O())
		assert.Equal(t, []int{5, 4, 3, -1, -2, 6}, slice.O())
	}
}

// Swap
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Swap_Go(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		if i+1 < len(ints) {
			ints[i], ints[i+1] = ints[i+1], ints[i]
		}
	}
}

func BenchmarkRefSlice_Swap_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Swap(i, i+1)
		}
	}
}

func BenchmarkRefSlice_Swap_Reflect(t *testing.B) {
	slice := NewRefSlice(rangeInterObject(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Swap(i, i+1)
		}
	}
}

func ExampleRefSlice_Swap() {
	slice := NewRefSliceV(2, 3, 1)
	slice.Swap(0, 2)
	slice.Swap(1, 2)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestRefSlice_Swap(t *testing.T) {

	// invalid cases
	{
		var slice *RefSlice
		slice.Swap(0, 0)
		assert.Equal(t, (*RefSlice)(nil), slice)

		slice = NewRefSliceV()
		slice.Swap(0, 0)
		assert.Equal(t, NewRefSliceV(), slice)

		slice.Swap(1, 2)
		assert.Equal(t, NewRefSliceV(), slice)

		slice.Swap(-1, 2)
		assert.Equal(t, NewRefSliceV(), slice)

		slice.Swap(1, -2)
		assert.Equal(t, NewRefSliceV(), slice)
	}

	// bool
	{
		slice := NewRefSliceV(true, false, true)
		slice.Swap(0, 1)
		assert.Equal(t, []bool{false, true, true}, slice.O())
	}

	// int
	{
		slice := NewRefSliceV(0, 1, 2)
		slice.Swap(0, 1)
		assert.Equal(t, []int{1, 0, 2}, slice.O())
	}

	// string
	{
		slice := NewRefSliceV("0", "1", "2")
		slice.Swap(0, 1)
		assert.Equal(t, []string{"1", "0", "2"}, slice.O())
	}

	// custom
	{
		slice := NewRefSlice([]Object{{0}, {1}, {2}})
		slice.Swap(0, 1)
		assert.Equal(t, []Object{{1}, {0}, {2}}, slice.O())
	}
}

// Take
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Take_Go(t *testing.B) {
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

func BenchmarkRefSlice_Take_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines7))
	for slice.Len() > 1 {
		slice.Take(1, 10)
	}
}

func BenchmarkRefSlice_Take_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines7))
	for slice.Len() > 1 {
		slice.Take(1, 10)
	}
}

func ExampleRefSlice_Take() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.Take(0, 1))
	// Output: [1 2]
}

func TestRefSlice_Take(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, NewRefSliceV(), slice.Take(0, 1))
	}

	// invalid
	{
		slice := NewRefSliceV(1, 2, 3, 4)
		assert.Equal(t, []int{}, slice.Take(1).O())
		assert.Equal(t, []int{1, 2, 3, 4}, slice.O())
		assert.Equal(t, []int{}, slice.Take(4, 4).O())
		assert.Equal(t, []int{1, 2, 3, 4}, slice.O())
	}

	// take 1
	{
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1}, slice.Take(0, 0).O())
			assert.Equal(t, []int{2, 3, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{2}, slice.Take(1, 1).O())
			assert.Equal(t, []int{1, 3, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{3}, slice.Take(2, 2).O())
			assert.Equal(t, []int{1, 2, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{4}, slice.Take(3, 3).O())
			assert.Equal(t, []int{1, 2, 3}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{4}, slice.Take(-1, -1).O())
			assert.Equal(t, []int{1, 2, 3}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{3}, slice.Take(-2, -2).O())
			assert.Equal(t, []int{1, 2, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{2}, slice.Take(-3, -3).O())
			assert.Equal(t, []int{1, 3, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1}, slice.Take(-4, -4).O())
			assert.Equal(t, []int{2, 3, 4}, slice.O())
		}
	}

	// take 2
	{
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2}, slice.Take(0, 1).O())
			assert.Equal(t, []int{3, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{2, 3}, slice.Take(1, 2).O())
			assert.Equal(t, []int{1, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{3, 4}, slice.Take(2, 3).O())
			assert.Equal(t, []int{1, 2}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{3, 4}, slice.Take(-2, -1).O())
			assert.Equal(t, []int{1, 2}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{2, 3}, slice.Take(-3, -2).O())
			assert.Equal(t, []int{1, 4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2}, slice.Take(-4, -3).O())
			assert.Equal(t, []int{3, 4}, slice.O())
		}
	}

	// take 3
	{
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2, 3}, slice.Take(0, 2).O())
			assert.Equal(t, []int{4}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{2, 3, 4}, slice.Take(-3, -1).O())
			assert.Equal(t, []int{1}, slice.O())
		}
	}

	// take everything and beyond
	{
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2, 3, 4}, slice.Take().O())
			assert.Equal(t, []int{}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2, 3, 4}, slice.Take(0, 3).O())
			assert.Equal(t, []int{}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2, 3, 4}, slice.Take(0, -1).O())
			assert.Equal(t, []int{}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2, 3, 4}, slice.Take(-4, -1).O())
			assert.Equal(t, []int{}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2, 3, 4}, slice.Take(-6, -1).O())
			assert.Equal(t, []int{}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1, 2, 3, 4}, slice.Take(0, 10).O())
			assert.Equal(t, []int{}, slice.O())
		}
	}

	// move index within bounds
	{
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{4}, slice.Take(3, 4).O())
			assert.Equal(t, []int{1, 2, 3}, slice.O())
		}
		{
			slice := NewRefSliceV(1, 2, 3, 4)
			assert.Equal(t, []int{1}, slice.Take(-5, 0).O())
			assert.Equal(t, []int{2, 3, 4}, slice.O())
		}
	}
}

// TakeAt
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_TakeAt_Go(t *testing.B) {
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

func BenchmarkRefSlice_TakeAt_Optimized(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := NewIntSlice(src)
	for i := range index {
		slice.TakeAt(i)
	}
}

func BenchmarkRefSlice_TakeAt_Reflect(t *testing.B) {
	src := rangeInterObject(0, nines5)
	index := Range(0, nines5)
	slice := NewRefSlice(src)
	for i := range index {
		slice.TakeAt(i)
	}
}

func ExampleRefSlice_TakeAt() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.TakeAt(2).O())
	// Output: 3
}

func TestRefSlice_TakeAt(t *testing.T) {

	// int
	{
		// nil or empty
		{
			var nilSlice *RefSlice
			assert.Equal(t, &Object{}, nilSlice.TakeAt(0))
		}

		// Delete all and more
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(-1)
			assert.Equal(t, Obj(2), obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, Obj(1), obj)
			assert.Equal(t, []int{0}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []int{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.TakeAt(-1)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []int{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(3)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(2)
			assert.Equal(t, Obj(2), obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(1)
			assert.Equal(t, Obj(1), obj)
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(0)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []int{1, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(-4)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(-1)
			assert.Equal(t, Obj(2), obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := NewRefSliceV(0, 1, 2)
			obj := slice.TakeAt(-2)
			assert.Equal(t, Obj(1), obj)
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// custom
	{
		// Delete all and more
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []Object{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.TakeAt(-1)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []Object{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(3)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(2)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(0)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []Object{{1}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-4)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := NewRefSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-2)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}
}

// TakeW
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_TakeW_Go(t *testing.B) {
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

func BenchmarkRefSlice_TakeW_Optimized(t *testing.B) {
	slice := NewIntSlice(Range(0, nines5))
	slice.TakeW(func(x O) bool { return ExB(x.(int)%2 == 0) })
}

func BenchmarkRefSlice_TakeW_Reflect(t *testing.B) {
	slice := NewRefSlice(Range(0, nines5))
	slice.TakeW(func(x O) bool { return ExB(x.(int)%2 == 0) })
}

func ExampleRefSlice_TakeW() {
	slice := NewRefSliceV(1, 2, 3)
	fmt.Println(slice.TakeW(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	}))
	// Output: [2]
}

func TestRefSlice_TakeW(t *testing.T) {

	// take all odd values
	{
		slice := NewRefSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.TakeW(func(x O) bool { return ExB(x.(int)%2 != 0) })
		assert.Equal(t, []int{2, 4, 6, 8}, slice.O())
		assert.Equal(t, []int{1, 3, 5, 7, 9}, new.O())
	}

	// take all even values
	{
		slice := NewRefSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.TakeW(func(x O) bool { return ExB(x.(int)%2 == 0) })
		assert.Equal(t, []int{1, 3, 5, 7, 9}, slice.O())
		assert.Equal(t, []int{2, 4, 6, 8}, new.O())
	}
}

// // Union
// //--------------------------------------------------------------------------------------------------
// func BenchmarkRefSlice_Union_Go(t *testing.B) {
// 	// ints := Range(0, nines7)
// 	// for len(ints) > 10 {
// 	// 	ints = ints[10:]
// 	// }
// }

// func BenchmarkRefSlice_Union_Slice(t *testing.B) {
// 	// slice := NewRefSlice(Range(0, nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleRefSlice_Union() {
// 	slice := NewRefSliceV(1, 2)
// 	fmt.Println(slice.Union([]int{2, 3}))
// 	// Output: [1 2 3]
// }

// func TestRefSlice_Union(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *RefSlice
// 		assert.Equal(t, NewRefSliceV(1, 2), slice.Union(NewRefSliceV(1, 2)))
// 		assert.Equal(t, NewRefSliceV(1, 2), slice.Union([]int{1, 2}))
// 	}

// 	// size of one
// 	{
// 		slice := NewRefSliceV(1)
// 		union := slice.Union([]int{1, 2, 3})
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), union)
// 		assert.Equal(t, NewRefSliceV(1), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewRefSliceV(1, 1)
// 		union := slice.Union(NewRefSliceV(2, 3))
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), union)
// 		assert.Equal(t, NewRefSliceV(1, 1), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewRefSliceV(1, 2, 2, 3, 3)
// 		union := slice.Union([]int{1, 2, 3})
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), union)
// 		assert.Equal(t, NewRefSliceV(1, 2, 2, 3, 3), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewRefSliceV(1, 2, 3)
// 		union := slice.Union([]int{4, 5})
// 		assert.Equal(t, NewRefSliceV(1, 2, 3, 4, 5), union)
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), slice)
// 	}
// }

// // UnionM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkRefSlice_UnionM_Go(t *testing.B) {
// 	// ints := Range(0, nines7)
// 	// for len(ints) > 10 {
// 	// 	ints = ints[10:]
// 	// }
// }

// func BenchmarkRefSlice_UnionM_Slice(t *testing.B) {
// 	// slice := NewRefSlice(Range(0, nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleRefSlice_UnionM() {
// 	slice := NewRefSliceV(1, 2)
// 	fmt.Println(slice.UnionM([]int{2, 3}))
// 	// Output: [1 2 3]
// }

// func TestRefSlice_UnionM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *RefSlice
// 		assert.Equal(t, NewRefSliceV(1, 2), slice.UnionM(NewRefSliceV(1, 2)))
// 		assert.Equal(t, (*RefSlice)(nil), slice)
// 	}

// 	// size of one
// 	{
// 		slice := NewRefSliceV(1)
// 		union := slice.UnionM([]int{1, 2, 3})
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), union)
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), slice)
// 	}

// 	// one duplicate
// 	{
// 		slice := NewRefSliceV(1, 1)
// 		union := slice.UnionM(NewRefSliceV(2, 3))
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), union)
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), slice)
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewRefSliceV(1, 2, 2, 3, 3)
// 		union := slice.UnionM([]int{1, 2, 3})
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), union)
// 		assert.Equal(t, NewRefSliceV(1, 2, 3), slice)
// 	}

// 	// no duplicates
// 	{
// 		slice := NewRefSliceV(1, 2, 3)
// 		union := slice.UnionM([]int{4, 5})
// 		assert.Equal(t, NewRefSliceV(1, 2, 3, 4, 5), union)
// 		assert.Equal(t, NewRefSliceV(1, 2, 3, 4, 5), slice)
// 	}
// }

// Uniq
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_Uniq_Go(t *testing.B) {
	// ints := Range(0, nines7)
	// for len(ints) > 10 {
	// 	ints = ints[10:]
	// }
}

func BenchmarkRefSlice_Uniq_Slice(t *testing.B) {
	// slice := NewRefSlice(Range(0, nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleRefSlice_Uniq() {
	slice := NewRefSliceV(1, 2, 3, 3)
	fmt.Println(slice.Uniq())
	// Output: [1 2 3]
}

func TestRefSlice_Uniq(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, NewRefSliceV(), slice.Uniq())
	}

	// size of one
	{
		slice := NewRefSliceV(1)
		uniq := slice.Uniq()
		assert.Equal(t, []int{1}, uniq.O())
		assert.Equal(t, []int{1, 2}, slice.Append(2).O())
		assert.Equal(t, []int{1}, uniq.O())
	}

	// one duplicate
	{
		slice := NewRefSliceV(1, 1)
		uniq := slice.Uniq()
		assert.Equal(t, []int{1}, uniq.O())
		assert.Equal(t, []int{1, 1, 2}, slice.Append(2).O())
		assert.Equal(t, []int{1}, uniq.O())
	}

	// multiple duplicates
	{
		slice := NewRefSliceV(1, 2, 2, 3, 3)
		uniq := slice.Uniq()
		assert.Equal(t, []int{1, 2, 3}, uniq.O())
		assert.Equal(t, []int{1, 2, 2, 3, 3, 4}, slice.Append(4).O())
		assert.Equal(t, []int{1, 2, 3}, uniq.O())
	}

	// no duplicates
	{
		slice := NewRefSliceV(1, 2, 3)
		uniq := slice.Uniq()
		assert.Equal(t, []int{1, 2, 3}, uniq.O())
		assert.Equal(t, []int{1, 2, 3, 4}, slice.Append(4).O())
		assert.Equal(t, []int{1, 2, 3}, uniq.O())
	}
}

// UniqM
//--------------------------------------------------------------------------------------------------
func BenchmarkRefSlice_UniqM_Go(t *testing.B) {
	// ints := Range(0, nines7)
	// for len(ints) > 10 {
	// 	ints = ints[10:]
	// }
}

func BenchmarkRefSlice_UniqM_Slice(t *testing.B) {
	// slice := NewRefSlice(Range(0, nines7))
	// for slice.Len() > 0 {
	// 	slice.PopN(10)
	// }
}

func ExampleRefSlice_UniqM() {
	slice := NewRefSliceV(1, 2, 3, 3)
	fmt.Println(slice.UniqM())
	// Output: [1 2 3]
}

func TestRefSlice_UniqM(t *testing.T) {

	// nil or empty
	{
		var slice *RefSlice
		assert.Equal(t, (*RefSlice)(nil), slice.UniqM())
	}

	// size of one
	{
		slice := NewRefSliceV(1)
		uniq := slice.UniqM()
		assert.Equal(t, []int{1}, uniq.O())
		assert.Equal(t, []int{1, 2}, slice.Append(2).O())
		assert.Equal(t, []int{1, 2}, uniq.O())
	}

	// one duplicate
	{
		slice := NewRefSliceV(1, 1)
		uniq := slice.UniqM()
		assert.Equal(t, []int{1}, uniq.O())
		assert.Equal(t, []int{1, 2}, slice.Append(2).O())
		assert.Equal(t, []int{1, 2}, uniq.O())
	}

	// multiple duplicates
	{
		slice := NewRefSliceV(1, 2, 2, 3, 3)
		uniq := slice.UniqM()
		assert.Equal(t, []int{1, 2, 3}, uniq.O())
		assert.Equal(t, []int{1, 2, 3, 4}, slice.Append(4).O())
		assert.Equal(t, []int{1, 2, 3, 4}, uniq.O())
	}

	// no duplicates
	{
		slice := NewRefSliceV(1, 2, 3)
		uniq := slice.UniqM()
		assert.Equal(t, []int{1, 2, 3}, uniq.O())
		assert.Equal(t, []int{1, 2, 3, 4}, slice.Append(4).O())
		assert.Equal(t, []int{1, 2, 3, 4}, uniq.O())
	}
}
