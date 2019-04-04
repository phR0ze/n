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
	slice := NewRefSlice(src)
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
