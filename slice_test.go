package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Slice function
//--------------------------------------------------------------------------------------------------
func ExampleSlice() {
	slice := OldSlice([]int{1, 2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_Slice_Constructor(t *testing.T) {

	// arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, []string{"1", "2"}, OldSlice(array).O())

	// empty
	assert.Equal(t, nil, OldSlice(nil).O())
	assert.Equal(t, &NSlice{}, OldSlice(nil))
	assert.Equal(t, []int{}, OldSlice([]int{}).O())
	assert.Equal(t, []bool{}, OldSlice([]bool{}).O())
	assert.Equal(t, []string{}, OldSlice([]string{}).O())
	assert.Equal(t, []Object{}, OldSlice([]Object{}).O())
	assert.Equal(t, nil, OldSlice([]interface{}{}).O())

	// pointers
	var obj *Object
	assert.Equal(t, []*Object{nil}, OldSlice(obj).O())
	assert.Equal(t, []*Object{&(Object{"bob"})}, OldSlice(&(Object{"bob"})).O())
	assert.Equal(t, []*Object{&(Object{"1"}), &(Object{"2"})}, OldSlice([]*Object{&(Object{"1"}), &(Object{"2"})}).O())

	// interface
	assert.Equal(t, nil, OldSlice([]interface{}{nil}).O())
	assert.Equal(t, []string{""}, OldSlice([]interface{}{nil, ""}).O())
	assert.Equal(t, []bool{true}, OldSlice([]interface{}{true}).O())
	assert.Equal(t, []int{1}, OldSlice([]interface{}{1}).O())
	assert.Equal(t, []string{""}, OldSlice([]interface{}{""}).O())
	assert.Equal(t, []string{"bob"}, OldSlice([]interface{}{"bob"}).O())
	assert.Equal(t, []Object{{nil}}, OldSlice([]interface{}{Object{}}).O())

	// singles
	assert.Equal(t, []int{1}, OldSlice(1).O())
	assert.Equal(t, []bool{true}, OldSlice(true).O())
	assert.Equal(t, []string{""}, OldSlice("").O())
	assert.Equal(t, []string{"1"}, OldSlice("1").O())
	assert.Equal(t, []Object{{1}}, OldSlice(Object{1}).O())
	assert.Equal(t, []Object{Object{"bob"}}, OldSlice(Object{"bob"}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, OldSlice(map[string]string{"1": "one"}).O())

	// slices
	assert.Equal(t, []int{1, 2}, OldSlice([]int{1, 2}).O())
	assert.Equal(t, []bool{true}, OldSlice([]bool{true}).O())
	assert.Equal(t, []Object{{"bob"}}, OldSlice([]Object{{"bob"}}).O())
	assert.Equal(t, []string{"1", "2"}, OldSlice([]string{"1", "2"}).O())
	assert.Equal(t, [][]string{{"1"}}, OldSlice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, OldSlice([]interface{}{map[string]string{"1": "one"}}).O())
}

// SliceV function
//--------------------------------------------------------------------------------------------------
func ExampleSliceV_empty() {
	slice := OldSliceV()
	fmt.Println(slice.O())
	// Output: <nil>
}

func ExampleSliceV_variadic() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_SliceV(t *testing.T) {
	var obj *Object

	// Arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, [][2]string{array}, OldSliceV(array).O())

	// Test empty values
	assert.True(t, !OldSliceV().Any())
	assert.Equal(t, 0, OldSliceV().Len())
	assert.Equal(t, nil, OldSliceV().O())
	assert.Equal(t, nil, OldSliceV(nil).O())
	assert.Equal(t, &NSlice{}, OldSliceV(nil))
	assert.Equal(t, []string{""}, OldSliceV(nil, "").O())
	assert.Equal(t, []*Object{nil}, OldSliceV(nil, obj).O())

	// Test pointers
	assert.Equal(t, []*Object{nil}, OldSliceV(obj).O())
	assert.Equal(t, []*Object{&(Object{"bob"})}, OldSliceV(&(Object{"bob"})).O())
	assert.Equal(t, []*Object{nil}, OldSliceV(obj).O())
	assert.Equal(t, []*Object{&(Object{"bob"})}, OldSliceV(&(Object{"bob"})).O())
	assert.Equal(t, [][]*Object{{&(Object{"1"}), &(Object{"2"})}}, OldSliceV([]*Object{&(Object{"1"}), &(Object{"2"})}).O())

	// Singles
	assert.Equal(t, []int{1}, OldSliceV(1).O())
	assert.Equal(t, []string{"1"}, OldSliceV("1").O())
	assert.Equal(t, []Object{Object{"bob"}}, OldSliceV(Object{"bob"}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, OldSliceV(map[string]string{"1": "one"}).O())

	// Multiples
	assert.Equal(t, []int{1, 2}, OldSliceV(1, 2).O())
	assert.Equal(t, []string{"1", "2"}, OldSliceV("1", "2").O())
	assert.Equal(t, []Object{Object{1}, Object{2}}, OldSliceV(Object{1}, Object{2}).O())

	// Test slices
	assert.Equal(t, [][]int{{1, 2}}, OldSliceV([]int{1, 2}).O())
	assert.Equal(t, [][]string{{"1"}}, OldSliceV([]string{"1"}).O())
}

func TestNSlice_newEmptySlice(t *testing.T) {

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
func BenchmarkNSlice_Any_Normal(t *testing.B) {
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

func BenchmarkNSlice_Any_Optimized(t *testing.B) {
	src := Range(0, nines4)
	slice := OldSlice(src)
	for i := range src {
		slice.Any(i)
	}
}

func BenchmarkNSlice_Any_Reflect(t *testing.B) {
	src := rangeNObject(0, nines4)
	slice := OldSlice(src)
	for _, i := range src {
		slice.Any(i)
	}
}

func ExampleNSlice_Any_empty() {
	slice := OldSliceV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleNSlice_Any_notEmpty() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleNSlice_Any_contains() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Any(1))
	// Output: true
}

func ExampleNSlice_Any_containsAny() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Any(0, 1))
	// Output: true
}

func TestNSlice_Any(t *testing.T) {
	var nilSlice *NSlice
	assert.False(t, nilSlice.Any())
	assert.False(t, OldSliceV().Any())
	assert.True(t, OldSliceV().Append("2").Any())

	// bool
	assert.True(t, OldSliceV(false, true).Any(true))
	assert.False(t, OldSliceV(true, true).Any(false))
	assert.True(t, OldSliceV(true, true).Any(false, true))
	assert.False(t, OldSliceV(true, true).Any(false, false))

	// int
	assert.True(t, OldSliceV(1, 2, 3).Any(2))
	assert.False(t, OldSliceV(1, 2, 3).Any(4))
	assert.True(t, OldSliceV(1, 2, 3).Any(4, 3))
	assert.False(t, OldSliceV(1, 2, 3).Any(4, 5))

	// int64
	assert.True(t, OldSliceV(int64(1), int64(2), int64(3)).Any(int64(2)))
	assert.False(t, OldSliceV(int64(1), int64(2), int64(3)).Any(int64(4)))
	assert.True(t, OldSliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(2)))
	assert.False(t, OldSliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(5)))

	// string
	assert.True(t, OldSliceV("1", "2", "3").Any("2"))
	assert.False(t, OldSliceV("1", "2", "3").Any("4"))
	assert.True(t, OldSliceV("1", "2", "3").Any("4", "2"))
	assert.False(t, OldSliceV("1", "2", "3").Any("4", "5"))

	// custom
	assert.True(t, OldSliceV(Object{1}, Object{2}).Any(Object{1}))
	assert.False(t, OldSliceV(Object{1}, Object{2}).Any(Object{3}))
	assert.True(t, OldSliceV(Object{1}, Object{2}).Any(Object{4}, Object{2}))
	assert.False(t, OldSliceV(Object{1}, Object{2}).Any(Object{4}, Object{5}))

	// panics need to go as last item as they abort the test method
	defer func() {
		err := recover()
		assert.Equal(t, "can't compare type 'int' with '[]n.Object' elements", err)
	}()
	assert.True(t, OldSliceV(Object{1}, Object{2}).Any(2))
}

// AnyS
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_AnyS_Normal(t *testing.B) {
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

func BenchmarkNSlice_AnyS_Optimized(t *testing.B) {
	src := Range(0, nines4)
	slice := OldSlice(src)
	for i := range src {
		slice.Any([]int{i})
	}
}

func BenchmarkNSlice_AnyS_Reflect(t *testing.B) {
	src := rangeNObject(0, nines4)
	slice := OldSlice(src)
	for _, i := range src {
		slice.Any(Object{i})
	}
}

func ExampleNSlice_AnyS() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.AnyS([]int{0, 1}))
	// Output: true
}

func TestNSlice_AnyS(t *testing.T) {
	var nilSlice *NSlice
	assert.False(t, nilSlice.AnyS([]bool{true}))

	// bool
	assert.True(t, OldSliceV(true, true).AnyS([]bool{true}))
	assert.True(t, OldSliceV(true, true).AnyS([]bool{false, true}))
	assert.False(t, OldSliceV(true, true).AnyS([]bool{false, false}))

	// int
	assert.True(t, OldSliceV(1, 2, 3).AnyS([]int{1}))
	assert.True(t, OldSliceV(1, 2, 3).AnyS([]int{4, 3}))
	assert.False(t, OldSliceV(1, 2, 3).AnyS([]int{4, 5}))

	// int64
	assert.True(t, OldSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(2)}))
	assert.True(t, OldSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(2)}))
	assert.False(t, OldSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(5)}))

	// string
	assert.True(t, OldSliceV("1", "2", "3").AnyS([]string{"2"}))
	assert.True(t, OldSliceV("1", "2", "3").AnyS([]string{"4", "2"}))
	assert.False(t, OldSliceV("1", "2", "3").AnyS([]string{"4", "5"}))

	// custom
	assert.True(t, OldSliceV(Object{1}, Object{2}).AnyS([]Object{{2}}))
	assert.True(t, OldSliceV(Object{1}, Object{2}).AnyS([]Object{{4}, {2}}))
	assert.False(t, OldSliceV(Object{1}, Object{2}).AnyS([]Object{{4}, {5}}))

	// panics need to go as last item as they abort the test method
	defer func() {
		err := recover()
		assert.Equal(t, "can't compare type '[]int' with '[]n.Object' elements", err)
	}()
	assert.True(t, OldSliceV(Object{1}, Object{2}).AnyS([]int{2}))
}

// Append
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Append_Normal(t *testing.B) {
	ints := []int{}
	for _, i := range Range(0, nines6) {
		ints = append(ints, i)
	}
}

func BenchmarkNSlice_Append_Optimized(t *testing.B) {
	n := &NSlice{o: []int{}}
	for _, i := range Range(0, nines6) {
		n.Append(i)
	}
}

func BenchmarkNSlice_Append_Reflect(t *testing.B) {
	n := &NSlice{o: []Object{}}
	for _, i := range Range(0, nines6) {
		n.Append(Object{i})
	}
}

func ExampleNSlice_Append() {
	slice := OldSliceV(1).Append(2).Append(3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_Append_Reflect(t *testing.T) {

	// Use a custom type to invoke reflection
	n := OldSliceV(Object{"1"})
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
		assert.Equal(t, "reflect.Set: value of type int is not assignable to type n.Object", err)
	}()
	n.Append(2)
}

func TestNSlice_Append(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, OldSliceV(0), nilSlice.Append(0))
		assert.Equal(t, (*NSlice)(nil), nilSlice)
	}

	// Append one back to back
	{
		n := OldSliceV()
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
		n := OldSliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1)
		assert.Equal(t, []int{1}, n.O())
		n.Append(2)
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with nil not chained
	{
		n := OldSliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Start with nil chained
	{
		n := OldSliceV().Append(1).Append(2)
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with non nil
	{
		n := OldSliceV(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Use append result directly
	{
		n := OldSliceV(1)
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []int{1, 2}, n.Append(2).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := OldSliceV(true)
			assert.Equal(t, []bool{true, false}, n.Append(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := OldSliceV(0)
			assert.Equal(t, []int{0, 1}, n.Append(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := OldSliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.Append("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := OldSlice([]Object{{"3"}})
			assert.Equal(t, []Object{{"3"}, {"1"}}, n.Append(Object{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}
}

func TestNSlice_Append_boolTypeError(t *testing.T) {
	n := OldSliceV(true)
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'string' to '[]bool'", err)
	}()
	n.Append("2")
}

func TestNSlice_Append_intTypeError(t *testing.T) {
	n := OldSliceV(1)
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'string' to '[]int'", err)
	}()
	n.Append("2")
}

func TestNSlice_Append_stringTypeError(t *testing.T) {
	n := OldSliceV("1")
	defer func() {
		err := recover()
		assert.Equal(t, "can't append type 'int' to '[]string'", err)
	}()
	n.Append(2)
}

func TestNSlice_Append_customTypeError(t *testing.T) {
	n := OldSliceV(Object{1})
	defer func() {
		err := recover()
		assert.Equal(t, "reflect.Set: value of type int is not assignable to type n.Object", err)
	}()
	n.Append(2)
}

// AppendV
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_AppendV_Normal(t *testing.B) {
	ints := []int{}
	ints = append(ints, Range(0, nines6)...)
}

func BenchmarkNSlice_AppendV_Optimized(t *testing.B) {
	n := &NSlice{o: []int{}}
	new := rangeO(0, nines6)
	n.AppendV(new...)
}

func BenchmarkNSlice_AppendV_Reflect(t *testing.B) {
	n := &NSlice{o: []Object{}}
	new := rangeNObject(0, nines6)
	n.AppendV(new...)
}

func ExampleNSlice_AppendV() {
	slice := OldSliceV(1).AppendV(2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_AppendV(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, (*NSlice)(nil), nilSlice.AppendV(0))
		assert.Equal(t, (*NSlice)(nil), nilSlice)
	}

	// Append many ints
	{
		n := OldSliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.AppendV(2, 3).O())
	}

	// Append many strings
	{
		{
			n := OldSliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendV("1", "2", "3").O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := OldSlice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendV("2", "3").O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := OldSlice([]Object{{"3"}})
		assert.Equal(t, []Object{{"3"}, {"1"}}, n.AppendV(Object{"1"}).O())
		assert.Equal(t, []Object{{"3"}, {"1"}, {"2"}, {"4"}}, n.AppendV(Object{"2"}, Object{"4"}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := OldSliceV(true)
			assert.Equal(t, []bool{true, false}, n.AppendV(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := OldSliceV(0)
			assert.Equal(t, []int{0, 1}, n.AppendV(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := OldSliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.AppendV("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := OldSlice([]Object{{"3"}})
			assert.Equal(t, []Object{{"3"}, {"1"}}, n.AppendV(Object{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}

	// Append to a slice of map
	{
		n := OldSliceV(map[string]string{"1": "one"})
		expected := []map[string]string{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, n.AppendV(map[string]string{"2": "two"}).O())
	}
}

// AppendS
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_AppendS_Normal10(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkNSlice_AppendS_Normal100(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkNSlice_AppendS_Optimized19(t *testing.B) {
	dest := &NSlice{o: []int{}}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func BenchmarkNSlice_AppendS_Optimized100(t *testing.B) {
	dest := &NSlice{o: []int{}}
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func BenchmarkNSlice_AppendS_Reflect10(t *testing.B) {
	dest := &NSlice{o: []Object{}}
	src := rangeNObject(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func BenchmarkNSlice_AppendS_Reflect100(t *testing.B) {
	dest := &NSlice{o: []Object{}}
	src := rangeNObject(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func ExampleNSlice_AppendS() {
	slice := OldSliceV(1).AppendS([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_AppendS(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, OldSliceV(1, 2), nilSlice.AppendS([]int{1, 2}))
		assert.Equal(t, (*NSlice)(nil), nilSlice)
	}

	// Append many ints
	{
		n := OldSliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.AppendS([]int{2, 3}).O())
	}

	// Append many strings
	{
		{
			n := OldSliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendS([]string{"1", "2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := OldSlice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendS([]string{"2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := OldSlice([]Object{{"3"}})
		assert.Equal(t, []Object{{"3"}, {"1"}}, n.AppendS([]Object{{"1"}}).O())
		assert.Equal(t, []Object{{"3"}, {"1"}, {"2"}, {"4"}}, n.AppendS([]Object{{"2"}, {"4"}}).O())
	}

	// Append to a slice of map
	{
		n := OldSliceV(map[string]string{"1": "one"})
		expected := []map[string]string{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, n.AppendS([]map[string]string{{"2": "two"}}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := OldSliceV(true)
			assert.Equal(t, []bool{true, false}, n.AppendS([]bool{false}).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := OldSliceV(0)
			assert.Equal(t, []int{0, 1}, n.AppendS([]int{1}).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := OldSliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.AppendS([]string{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := OldSlice([]Object{{"3"}})
			assert.Equal(t, []Object{{"3"}, {"1"}}, n.AppendS([]Object{{"1"}}).O())
			assert.Equal(t, 2, n.Len())
		}
	}
}

// At
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_At_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := range ints {
		assert.IsType(t, 0, ints[i])
	}
}

func BenchmarkNSlice_At_Optimized(t *testing.B) {
	src := Range(0, nines6)
	slice := OldSlice(src)
	for _, i := range src {
		_, ok := (slice.At(i).O()).(int)
		assert.True(t, ok)
	}
}

func BenchmarkNSlice_At_Reflect(t *testing.B) {
	src := rangeNObject(0, nines6)
	slice := OldSlice(src)
	for i := range src {
		_, ok := (slice.At(i).O()).(Object)
		assert.True(t, ok)
	}
}

func ExampleNSlice_At() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.At(2).O())
	// Output: 3
}

func TestNSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, &Object{nil}, nilSlice.At(0))
	}

	// strings
	{
		slice := OldSliceV("1", "2", "3", "4")
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
		slice := OldSliceV("1")
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------

func ExampleNSlice_Clear() {
	slice := OldSliceV(1).AppendS([]int{2, 3})
	fmt.Println(slice.Clear().O())
	// Output: []
}

func TestQSlice_Clear(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, &Object{nil}, nilSlice.At(0))
	}

	// bool
	{
		slice := OldSliceV(true, false)
		assert.Equal(t, 2, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []bool{}}, slice)
	}

	// int
	{
		slice := OldSliceV(1, 2, 3, 4)
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []int{}}, slice)
	}

	// string
	{
		slice := OldSliceV("1", "2", "3", "4")
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []string{}}, slice)
	}

	// custom
	{
		slice := OldSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, 3, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []Object{}}, slice)
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Copy_Normal(t *testing.B) {
	ints := Range(0, nines6)
	dst := make([]int, len(ints), len(ints))
	copy(dst, ints)
}

func BenchmarkNSlice_Copy_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines6))
	slice.Copy()
}

func BenchmarkNSlice_Copy_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines6))
	slice.Copy()
}

func ExampleNSlice_Copy() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Copy().O())
	// Output: [1 2 3]
}

func TestNSlice_Copy(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, OldSliceV(), nilSlice.Copy(0, -1))
		slice := OldSliceV(0).Clear()
		assert.Equal(t, &NSlice{o: []int{}}, slice.Copy(0, -1))
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := OldSliceV(1, 2, 3)
		result := original.Copy(0, -1)
		assert.Equal(t, []int{1, 2, 3}, original.O())
		assert.Equal(t, []int{1, 2, 3}, result.O())
		result.Set(0, 0)
		assert.Equal(t, []int{1, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2, 3}, result.O())
	}

	// copy full array
	{
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Copy())
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Copy(0, -1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Copy(0, 1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Copy(0, 5))
		assert.Equal(t, OldSliceV(""), OldSliceV("").Copy())
		assert.Equal(t, OldSliceV(""), OldSliceV("").Copy(0, -1))
		assert.Equal(t, OldSliceV(""), OldSliceV("").Copy(0, 1))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).Copy())
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).Copy(0, -1))
		assert.Equal(t, OldSlice([]int{1, 2, 3}), OldSlice([]int{1, 2, 3}).Copy())
		assert.Equal(t, OldSlice([]int{1, 2, 3}), OldSlice([]int{1, 2, 3}).Copy(0, -1))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Copy())
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Copy(0, 2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Copy())
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, OldSliceV("1"), OldSliceV("1").Copy(0, 2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false).Copy(-6, 6))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).Copy(-6, 6))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Copy(-6, 6))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(-6, 6))
	}

	// mutually exclusive
	{
		slice := OldSliceV(1, 2, 3, 4)
		assert.Equal(t, &NSlice{o: []int{}}, slice.Copy(2, -3))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Copy(0, -5))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Copy(4, -1))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Copy(6, -1))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Copy(3, 2))
	}

	// singles
	{
		slice := OldSliceV(1, 2, 3, 4)
		assert.Equal(t, OldSliceV(4), slice.Copy(-1, -1))
		assert.Equal(t, OldSliceV(3), slice.Copy(-2, -2))
		assert.Equal(t, OldSliceV(2), slice.Copy(-3, -3))
		assert.Equal(t, OldSliceV(1), slice.Copy(0, 0))
		assert.Equal(t, OldSliceV(1), slice.Copy(-4, -4))
		assert.Equal(t, OldSliceV(2), slice.Copy(1, 1))
		assert.Equal(t, OldSliceV(2), slice.Copy(1, -3))
		assert.Equal(t, OldSliceV(3), slice.Copy(2, 2))
		assert.Equal(t, OldSliceV(3), slice.Copy(2, -2))
		assert.Equal(t, OldSliceV(4), slice.Copy(3, 3))
		assert.Equal(t, OldSliceV(4), slice.Copy(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Copy(1, -1))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Copy(1, 2))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Copy(-2, -1))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Copy(-2, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Copy(1, -1))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Copy(1, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Copy(-2, -1))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Copy(-2, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Copy(1, -1))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Copy(1, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Copy(-2, -1))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Copy(-2, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(1, -1))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(1, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(-2, -1))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Copy(0, -2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Copy(-3, -2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Copy(-3, 1))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Copy(0, 1))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Copy(0, -2))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Copy(-3, -2))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Copy(-3, 1))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Copy(0, 1))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Copy(0, -2))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Copy(-3, -2))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Copy(-3, 1))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Copy(0, 1))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(0, -2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(-3, -2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(-3, 1))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Copy(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Copy(1, -2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Copy(-3, -2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Copy(-3, 2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Copy(1, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Copy(1, -2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Copy(-3, -2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Copy(-3, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Copy(1, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Copy(1, -2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Copy(-3, -2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Copy(-3, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Copy(1, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Copy(1, -2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Copy(-3, -2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Copy(-3, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Copy(1, 2))
	}

	// random
	{
		assert.Equal(t, OldSliceV("1"), OldSliceV("1", "2", "3").Copy(0, -3))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Copy(1, 2))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Copy(0, 2))
	}
}

// Drop
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Drop_Normal(t *testing.B) {
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

func BenchmarkNSlice_Drop_Optimized(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := OldSlice(src)
	for i := range index {
		slice.Drop(i)
	}
}

func BenchmarkNSlice_Drop_Reflect(t *testing.B) {
	src := rangeNObject(0, nines5)
	index := Range(0, nines5)
	slice := OldSlice(src)
	for i := range index {
		slice.Drop(i)
	}
}

func ExampleNSlice_Drop() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Drop(1).O())
	// Output: [1 3]
}

func TestNSlice_Drop(t *testing.T) {

	// int
	{
		// nil or empty
		{
			var nilSlice *NSlice
			assert.Equal(t, (*NSlice)(nil), nilSlice.Drop(0))
		}

		// drop all and more
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(0, 1), slice.Drop(-1))
			assert.Equal(t, OldSliceV(0), slice.Drop(-1))
			assert.Equal(t, &NSlice{o: []int{}}, slice.Drop(-1))
			assert.Equal(t, &NSlice{o: []int{}}, slice.Drop(-1))
		}

		// Pos: drop invalid
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(0, 1, 2), slice.Drop(3))
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: drop last
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(0, 1), slice.Drop(2))
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: drop middle
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(0, 2), slice.Drop(1))
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos drop first
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(1, 2), slice.Drop(0))
			assert.Equal(t, []int{1, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: drop invalid
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(0, 1, 2), slice.Drop(-4))
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: drop last
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(0, 1), slice.Drop(-1))
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: drop middle
		{
			slice := OldSliceV(0, 1, 2)
			assert.Equal(t, OldSliceV(0, 2), slice.Drop(-2))
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// custom
	{

		// drop all and more
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{0}, {1}}), slice.Drop(-1))
			assert.Equal(t, OldSlice([]Object{{0}}), slice.Drop(-1))
			assert.Equal(t, OldSlice([]Object{}), slice.Drop(-1))
			assert.Equal(t, OldSlice([]Object{}), slice.Drop(-1))
		}

		// Pos: drop invalid
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}}), slice.Drop(3))
		}

		// Pos: drop last
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{0}, {1}}), slice.Drop(2))
		}

		// Pos: drop middle
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{0}, {2}}), slice.Drop(1))
		}

		// Pos drop first
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{1}, {2}}), slice.Drop(0))
		}

		// Neg: drop invalid
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}}), slice.Drop(-4))
		}

		// Neg: drop last
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{0}, {1}}), slice.Drop(-1))
		}

		// Neg: drop middle
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			assert.Equal(t, OldSlice([]Object{{0}, {2}}), slice.Drop(-2))
		}
	}
}

// DropFirst
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_DropFirst_Normal(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkNSlice_DropFirst_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func BenchmarkNSlice_DropFirst_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleNSlice_DropFirst() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.DropFirst().O())
	// Output: [2 3]
}

func TestNSlice_DropFirst(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, (*NSlice)(nil), nilSlice.DropFirst())
	}

	// bool
	{
		slice := OldSliceV(true, true, false)
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
		slice := OldSliceV(1, 2, 3)
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
		slice := OldSliceV("1", "2", "3")
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
		slice := OldSlice([]Object{{1}, {2}, {3}})
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
func BenchmarkNSlice_DropFirstN_Normal(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkNSlice_DropFirstN_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func BenchmarkNSlice_DropFirstN_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func ExampleNSlice_DropFirstN() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.DropFirstN(2).O())
	// Output: [3]
}

func TestNSlice_DropFirstN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, (*NSlice)(nil), nilSlice.DropFirstN(1))
	}

	// drop none
	{
		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2, 3}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}, {2}, {3}}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := OldSliceV(true, true, false)
			assert.Equal(t, []bool{true, false}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{2, 3}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := OldSliceV("1", "2", "3")
			assert.Equal(t, []string{"2", "3"}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{2}, {3}}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := OldSliceV(true, false, false)
			assert.Equal(t, []bool{false}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{3}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := OldSliceV("1", "2", "3")
			assert.Equal(t, []string{"3"}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{3}}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropFirstN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropFirstN(4).O())
			assert.Equal(t, 0, slice.Len())
		}
	}
}

// DropLast
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_DropLast_Normal(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkNSlice_DropLast_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func BenchmarkNSlice_DropLast_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func ExampleNSlice_DropLast() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.DropLast().O())
	// Output: [1 2]
}

func TestNSlice_DropLast(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, (*NSlice)(nil), nilSlice.DropLast())
	}

	// bool
	{
		slice := OldSliceV(true, true, false)
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
		slice := OldSliceV(1, 2, 3)
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
		slice := OldSliceV("1", "2", "3")
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
		slice := OldSlice([]Object{{1}, {2}, {3}})
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
func BenchmarkNSlice_DropLastN_Normal(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 10 {
		ints = ints[10:]
	}
}

func BenchmarkNSlice_DropLastN_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func BenchmarkNSlice_DropLastN_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func ExampleNSlice_DropLastN() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.DropLastN(2).O())
	// Output: [1]
}

func TestNSlice_DropLastN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, (*NSlice)(nil), nilSlice.DropLastN(1))
	}

	// drop none
	{
		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2, 3}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}, {2}, {3}}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := OldSliceV(true, true, false)
			assert.Equal(t, []bool{true, true}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := OldSliceV("1", "2", "3")
			assert.Equal(t, []string{"1", "2"}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}, {2}}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := OldSliceV(true, false, false)
			assert.Equal(t, []bool{true}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{1}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := OldSliceV("1", "2", "3")
			assert.Equal(t, []string{"1"}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{{1}}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := OldSliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropLastN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := OldSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []Object{}, slice.DropLastN(4).O())
			assert.Equal(t, 0, slice.Len())
		}
	}
}

// Each
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Each_Normal(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkNSlice_Each_Optimized(t *testing.B) {
	OldSlice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkNSlice_Each_Reflect(t *testing.B) {
	OldSlice(rangeNObject(0, nines6)).Each(func(x O) {
		assert.IsType(t, Object{}, x)
	})
}

func ExampleNSlice_Each() {
	OldSliceV(1, 2, 3).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 123
}

func TestNSlice_Each(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		nilSlice.Each(func(x O) {})
	}

	// int
	{
		OldSliceV(1, 2, 3).Each(func(x O) {
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
		OldSliceV("1", "2", "3").Each(func(x O) {
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
		OldSlice([]Object{{1}, {2}, {3}}).Each(func(x O) {
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
func BenchmarkNSlice_EachE_Normal(t *testing.B) {
	action := func(x interface{}) {
		assert.IsType(t, 0, x)
	}
	for i := range Range(0, nines6) {
		action(i)
	}
}

func BenchmarkNSlice_EachE_Optimized(t *testing.B) {
	OldSlice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkNSlice_EachE_Reflect(t *testing.B) {
	OldSlice(rangeNObject(0, nines6)).Each(func(x O) {
		assert.IsType(t, Object{}, x)
	})
}

func ExampleNSlice_EachE() {
	OldSliceV(1, 2, 3).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 123
}

func TestNSlice_EachE(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		nilSlice.EachE(func(x O) error {
			return nil
		})
	}

	// int
	{
		OldSliceV(1, 2, 3).EachE(func(x O) error {
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
		OldSliceV("1", "2", "3").EachE(func(x O) error {
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
		OldSlice([]Object{{1}, {2}, {3}}).EachE(func(x O) error {
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

// Empty
//--------------------------------------------------------------------------------------------------
func ExampleNSlice_Empty() {
	fmt.Println(OldSliceV().Empty())
	// Output: true
}

func TestNSlice_Empty(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, true, nilSlice.Empty())
	}

	assert.Equal(t, true, OldSliceV().Empty())
	assert.Equal(t, false, OldSliceV(1).Empty())
	assert.Equal(t, false, OldSliceV(1, 2, 3).Empty())
	assert.Equal(t, false, OldSlice(1).Empty())
	assert.Equal(t, false, OldSlice([]int{1, 2, 3}).Empty())
}

// First
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_First_Normal(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkNSlice_First_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.First()
	}
}

func BenchmarkNSlice_First_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	for slice.Len() > 0 {
		slice.First()
	}
}

func ExampleNSlice_First() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.First().O())
	// Output: 1
}

func TestNSlice_First(t *testing.T) {
	// invalid
	{
		assert.Equal(t, &Object{nil}, OldSliceV().First())
	}

	// bool
	{
		assert.Equal(t, &Object{true}, OldSliceV(true, false).First())
		assert.Equal(t, &Object{false}, OldSliceV(false, true).First())
	}

	// int
	{
		assert.Equal(t, &Object{2}, OldSliceV(2, 3).First())
		assert.Equal(t, &Object{3}, OldSliceV(3, 2).First())
		assert.Equal(t, &Object{1}, OldSliceV(1, 3, 2).First())
	}

	// string
	{
		assert.Equal(t, &Object{"2"}, OldSliceV("2", "3").First())
		assert.Equal(t, &Object{"3"}, OldSliceV("3", "2").First())
		assert.Equal(t, &Object{"1"}, OldSliceV("1", "3", "2").First())
	}

	// custom
	{
		assert.Equal(t, &Object{Object{2}}, OldSlice([]Object{{2}, {3}}).First())
		assert.Equal(t, &Object{Object{3}}, OldSlice([]Object{{3}, {2}}).First())
		assert.Equal(t, &Object{Object{1}}, OldSlice([]Object{{1}, {3}, {2}}).First())
	}
}

// FirstN
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_FirstN_Normal(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:10]
}

func BenchmarkNSlice_FirstN_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	slice.FirstN(10)
}

func BenchmarkNSlice_FirstN_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	slice.FirstN(10)
}

func ExampleNSlice_FirstN() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.FirstN(2).O())
	// Output: [1 2]
}

func TestNSlice_FirstN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, OldSliceV(), nilSlice.FirstN(1))
		slice := OldSliceV(0).Clear()
		assert.Equal(t, &NSlice{o: []int{}}, slice.FirstN(-1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := OldSliceV(1, 2, 3)
		result := original.FirstN(2).Set(0, 0)
		assert.Equal(t, []int{0, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2}, result.O())
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().FirstN(1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().FirstN(10))
		assert.Equal(t, OldSliceV(""), OldSliceV("").FirstN(1))
		assert.Equal(t, OldSliceV(""), OldSliceV("").FirstN(10))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).FirstN(10))
		assert.Equal(t, OldSlice([]int{1, 2, 3}), OldSlice([]int{1, 2, 3}).FirstN(10))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").FirstN(10))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).FirstN(10))
	}

	// grab a few diff
	{
		assert.Equal(t, OldSliceV(true), OldSliceV(true, false, true).FirstN(1))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).FirstN(2))
		assert.Equal(t, OldSliceV(1), OldSliceV(1, 2, 3).FirstN(1))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).FirstN(2))
		assert.Equal(t, OldSliceV("1"), OldSliceV("1", "2", "3").FirstN(1))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").FirstN(2))
		assert.Equal(t, OldSlice([]Object{{1}}), OldSlice([]Object{{1}, {2}, {3}}).FirstN(1))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).FirstN(2))
	}
}

// Insert
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Insert_Normal(t *testing.B) {
	ints := []int{}
	for i := range Range(0, nines6) {
		ints = append(ints, i)
		copy(ints[1:], ints[1:])
		ints[0] = i
	}
}

func BenchmarkNSlice_Insert_Optimized(t *testing.B) {
	slice := &NSlice{o: []int{}}
	for i := range Range(0, nines6) {
		slice.Insert(0, i)
	}
}

func BenchmarkNSlice_Insert_Reflect(t *testing.B) {
	slice := &NSlice{o: []Object{}}
	for i := range Range(0, nines6) {
		slice.Insert(0, Object{i})
	}
}

func ExampleNSlice_Insert() {
	slice := OldSliceV(1, 3)
	fmt.Println(slice.Insert(1, 2).O())
	// Output: [1 2 3]
}

func TestNSlice_Insert(t *testing.T) {

	// int
	{
		// append
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(0), slice.Insert(-1, 0))
			assert.Equal(t, OldSliceV(0, 1), slice.Insert(-1, 1))
			assert.Equal(t, OldSliceV(0, 1, 2), slice.Insert(-1, 2))
		}

		// prepend
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(2), slice.Insert(0, 2))
			assert.Equal(t, OldSliceV(1, 2), slice.Insert(0, 1))
			assert.Equal(t, OldSliceV(0, 1, 2), slice.Insert(0, 0))
		}

		// middle pos
		{
			slice := OldSliceV(0, 5)
			assert.Equal(t, OldSliceV(0, 1, 5), slice.Insert(1, 1))
			assert.Equal(t, OldSliceV(0, 1, 2, 5), slice.Insert(2, 2))
			assert.Equal(t, OldSliceV(0, 1, 2, 3, 5), slice.Insert(3, 3))
			assert.Equal(t, OldSliceV(0, 1, 2, 3, 4, 5), slice.Insert(4, 4))
		}

		// middle neg
		{
			slice := OldSliceV(0, 5)
			assert.Equal(t, OldSliceV(0, 1, 5), slice.Insert(-2, 1))
			assert.Equal(t, OldSliceV(0, 1, 2, 5), slice.Insert(-2, 2))
			assert.Equal(t, OldSliceV(0, 1, 2, 3, 5), slice.Insert(-2, 3))
			assert.Equal(t, OldSliceV(0, 1, 2, 3, 4, 5), slice.Insert(-2, 4))
		}

		// error cases
		{
			var slice *NSlice
			assert.True(t, slice.Insert(0, 0).Nil())
			assert.Equal(t, (*NSlice)(nil), slice.Insert(0, 0))
			assert.Equal(t, OldSliceV(0, 1), OldSliceV(0, 1).Insert(-10, 1))
			assert.Equal(t, OldSliceV(0, 1), OldSliceV(0, 1).Insert(10, 1))
			assert.Equal(t, OldSliceV(0, 1), OldSliceV(0, 1).Insert(2, 1))
			assert.Equal(t, OldSliceV(0, 1), OldSliceV(0, 1).Insert(-3, 1))
		}
	}

	// custom
	{
		// append
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(0), slice.Insert(-1, 0))
			assert.Equal(t, OldSliceV(0, 1), slice.Insert(-1, 1))
			assert.Equal(t, OldSliceV(0, 1, 2), slice.Insert(-1, 2))
		}

		// prepend
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(Object{2}), slice.Insert(0, Object{2}))
			assert.Equal(t, OldSlice([]Object{{1}, {2}}), slice.Insert(0, Object{1}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}}), slice.Insert(0, Object{0}))
		}

		// middle pos
		{
			slice := OldSlice([]Object{{0}, {5}})
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {5}}), slice.Insert(1, Object{1}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}, {5}}), slice.Insert(2, Object{2}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}, {3}, {5}}), slice.Insert(3, Object{3}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}, {3}, {4}, {5}}), slice.Insert(4, Object{4}))
		}

		// middle neg
		{
			slice := OldSlice([]Object{{0}, {5}})
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {5}}), slice.Insert(-2, Object{1}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}, {5}}), slice.Insert(-2, Object{2}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}, {3}, {5}}), slice.Insert(-2, Object{3}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}, {3}, {4}, {5}}), slice.Insert(-2, Object{4}))
		}

		// error cases
		{
			var slice *NSlice
			assert.True(t, slice.Insert(0, Object{0}).Nil())
			assert.Equal(t, (*NSlice)(nil), slice.Insert(0, Object{0}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}}), OldSlice([]Object{{0}, {1}}).Insert(-10, 1))
			assert.Equal(t, OldSlice([]Object{{0}, {1}}), OldSlice([]Object{{0}, {1}}).Insert(10, 1))
			assert.Equal(t, OldSlice([]Object{{0}, {1}}), OldSlice([]Object{{0}, {1}}).Insert(2, 1))
			assert.Equal(t, OldSlice([]Object{{0}, {1}}), OldSlice([]Object{{0}, {1}}).Insert(-3, 1))
		}
	}
}

// // func TestStrSliceJoin(t *testing.T) {
// // 	assert.Equal(t, "", S().Join(".").A())
// // 	assert.Equal(t, "1", S("1").Join(".").A())
// // 	assert.Equal(t, "1.2", S("1", "2").Join(".").A())
// // }

// LastN
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_LastN_Normal(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:10]
}

func BenchmarkNSlice_LastN_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	slice.LastN(10)
}

func BenchmarkNSlice_LastN_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	slice.LastN(10)
}

func ExampleNSlice_LastN() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.LastN(2).O())
	// Output: [2 3]
}

func TestNSlice_LastN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, OldSliceV(), nilSlice.LastN(1))
		slice := OldSliceV(0).Clear()
		assert.Equal(t, &NSlice{o: []int{}}, slice.LastN(-1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := OldSliceV(1, 2, 3)
		result := original.LastN(2).Set(0, 0)
		assert.Equal(t, []int{1, 0, 3}, original.O())
		assert.Equal(t, []int{0, 3}, result.O())
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().LastN(1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().LastN(10))
		assert.Equal(t, OldSliceV(""), OldSliceV("").LastN(1))
		assert.Equal(t, OldSliceV(""), OldSliceV("").LastN(10))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).LastN(10))
		assert.Equal(t, OldSlice([]int{1, 2, 3}), OldSlice([]int{1, 2, 3}).LastN(10))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").LastN(10))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).LastN(10))
	}

	// grab a few diff
	{
		assert.Equal(t, OldSliceV(false), OldSliceV(true, true, false).LastN(1))
		assert.Equal(t, OldSliceV(false), OldSliceV(true, true, false).LastN(-1))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).LastN(2))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).LastN(-2))
		assert.Equal(t, OldSliceV(3), OldSliceV(1, 2, 3).LastN(1))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).LastN(2))
		assert.Equal(t, OldSliceV("3"), OldSliceV("1", "2", "3").LastN(1))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").LastN(2))
		assert.Equal(t, OldSlice([]Object{{3}}), OldSlice([]Object{{1}, {2}, {3}}).LastN(1))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).LastN(2))
	}
}

// Len
//--------------------------------------------------------------------------------------------------
func TestNSlice_Len(t *testing.T) {
	assert.Equal(t, 0, OldSliceV().Len())
	assert.Equal(t, 1, OldSliceV().Append("2").Len())
}

// Less
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Less_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		if i+1 < len(ints) {
			_ = ints[i] < ints[i+1]
		}
	}
}

func BenchmarkNSlice_Less_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Less(i, i+1)
		}
	}
}

func BenchmarkNSlice_Less_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Less(i, i+1)
		}
	}
}

func ExampleNSlice_Less() {
	slice := OldSliceV(2, 3, 1)
	fmt.Println(slice.Sort().O())
	// Output: [1 2 3]
}

func TestNSlice_Less(t *testing.T) {

	// invalid cases
	{
		var slice *NSlice
		assert.False(t, slice.Less(0, 0))

		slice = OldSliceV()
		assert.False(t, slice.Less(0, 0))
		assert.False(t, slice.Less(1, 2))
		assert.False(t, slice.Less(-1, 2))
		assert.False(t, slice.Less(1, -2))
	}

	// bool
	{
		assert.Equal(t, false, OldSliceV(true, false, true).Less(0, 1))
		assert.Equal(t, true, OldSliceV(true, false, true).Less(1, 0))
	}

	// int
	{
		assert.Equal(t, true, OldSliceV(0, 1, 2).Less(0, 1))
		assert.Equal(t, false, OldSliceV(0, 1, 2).Less(1, 0))
		assert.Equal(t, true, OldSliceV(0, 1, 2).Less(1, 2))
	}

	// string
	{
		assert.Equal(t, true, OldSliceV("0", "1", "2").Less(0, 1))
		assert.Equal(t, false, OldSliceV("0", "1", "2").Less(1, 0))
		assert.Equal(t, true, OldSliceV("0", "1", "2").Less(1, 2))
	}

	// custom
	{
		assert.Equal(t, true, OldSlice([]Object{{0}, {1}, {2}}).Less(0, 1))
		assert.Equal(t, false, OldSlice([]Object{{0}, {1}, {2}}).Less(1, 0))
		assert.Equal(t, true, OldSlice([]Object{{0}, {1}, {2}}).Less(1, 2))
	}
}

// Nil
//--------------------------------------------------------------------------------------------------
func TestNSlice_Nil(t *testing.T) {
	assert.True(t, OldSliceV().Nil())
	var q *NSlice
	assert.True(t, q.Nil())
	assert.False(t, OldSliceV().Append("2").Nil())
}

// O
//--------------------------------------------------------------------------------------------------
func TestNSlice_O(t *testing.T) {
	assert.Nil(t, OldSliceV().O())
	assert.Len(t, OldSliceV().Append("2").O(), 1)
}

// Pair
//--------------------------------------------------------------------------------------------------

func ExampleNSlice_Pair() {
	slice := OldSliceV(1, 2)
	first, second := slice.Pair()
	fmt.Println(first.O(), second.O())
	// Output: 1 2
}

func TestNSlice_Pair(t *testing.T) {

	// int
	{
		// two values
		{
			first, second := OldSliceV(1, 2).Pair()
			assert.Equal(t, &Object{1}, first)
			assert.Equal(t, &Object{2}, second)
		}

		// one value
		{
			first, second := OldSliceV(1).Pair()
			assert.Equal(t, &Object{1}, first)
			assert.Equal(t, &Object{nil}, second)
		}

		// no values
		{
			first, second := OldSliceV().Pair()
			assert.Equal(t, &Object{nil}, first)
			assert.Equal(t, &Object{nil}, second)
		}
	}

	// custom
	{
		// two values
		{
			first, second := OldSlice([]Object{{1}, {2}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, &Object{Object{2}}, second)
		}

		// one value
		{
			first, second := OldSlice([]Object{{1}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, &Object{nil}, second)
		}

		// no values
		{
			first, second := OldSliceV().Pair()
			assert.Equal(t, &Object{nil}, first)
			assert.Equal(t, &Object{nil}, second)
		}
	}
}

// Prepend
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Prepend_Normal(t *testing.B) {
	ints := []int{}
	for i := range Range(0, nines6) {
		ints = append(ints, i)
		copy(ints[1:], ints[1:])
		ints[0] = i
	}
}

func BenchmarkNSlice_Prepend_Optimized(t *testing.B) {
	slice := &NSlice{o: []int{}}
	for i := range Range(0, nines6) {
		slice.Prepend(i)
	}
}

func BenchmarkNSlice_Prepend_Reflect(t *testing.B) {
	slice := &NSlice{o: []Object{}}
	for i := range Range(0, nines6) {
		slice.Prepend(Object{i})
	}
}

func ExampleNSlice_Prepend() {
	slice := OldSliceV(2, 3)
	fmt.Println(slice.Prepend(1).O())
	// Output: [1 2 3]
}

func TestNSlice_Prepend(t *testing.T) {

	// int
	{
		// happy path
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(2), slice.Prepend(2))
			assert.Equal(t, OldSliceV(1, 2), slice.Prepend(1))
			assert.Equal(t, OldSliceV(0, 1, 2), slice.Prepend(0))
		}

		// error cases
		{
			var slice *NSlice
			assert.True(t, slice.Prepend(0).Nil())
			assert.Equal(t, (*NSlice)(nil), slice.Prepend(0))
		}
	}

	// custom
	{
		// prepend
		{
			slice := OldSliceV()
			assert.Equal(t, OldSliceV(Object{2}), slice.Prepend(Object{2}))
			assert.Equal(t, OldSlice([]Object{{1}, {2}}), slice.Prepend(Object{1}))
			assert.Equal(t, OldSlice([]Object{{0}, {1}, {2}}), slice.Prepend(Object{0}))
		}

		// error cases
		{
			var slice *NSlice
			assert.True(t, slice.Prepend(Object{0}).Nil())
			assert.Equal(t, (*NSlice)(nil), slice.Prepend(Object{0}))
		}
	}
}

// Set
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Set_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		ints[i] = 0
	}
}

func BenchmarkNSlice_Set_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, 0)
	}
}

func BenchmarkNSlice_Set_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, Object{0})
	}
}

func ExampleNSlice_Set() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Set(0, 0).O())
	// Output: [0 2 3]
}

func TestNSlice_Set(t *testing.T) {
	// bool
	{
		assert.Equal(t, []bool{false, true, true}, OldSliceV(true, true, true).Set(0, false).O())
		assert.Equal(t, []bool{true, false, true}, OldSliceV(true, true, true).Set(1, false).O())
		assert.Equal(t, []bool{true, true, false}, OldSliceV(true, true, true).Set(2, false).O())
		assert.Equal(t, []bool{false, true, true}, OldSliceV(true, true, true).Set(-3, false).O())
		assert.Equal(t, []bool{true, false, true}, OldSliceV(true, true, true).Set(-2, false).O())
		assert.Equal(t, []bool{true, true, false}, OldSliceV(true, true, true).Set(-1, false).O())
	}

	// int
	{
		assert.Equal(t, []int{0, 2, 3}, OldSliceV(1, 2, 3).Set(0, 0).O())
		assert.Equal(t, []int{1, 0, 3}, OldSliceV(1, 2, 3).Set(1, 0).O())
		assert.Equal(t, []int{1, 2, 0}, OldSliceV(1, 2, 3).Set(2, 0).O())
		assert.Equal(t, []int{0, 2, 3}, OldSliceV(1, 2, 3).Set(-3, 0).O())
		assert.Equal(t, []int{1, 0, 3}, OldSliceV(1, 2, 3).Set(-2, 0).O())
		assert.Equal(t, []int{1, 2, 0}, OldSliceV(1, 2, 3).Set(-1, 0).O())
	}

	// string
	{
		assert.Equal(t, []string{"0", "2", "3"}, OldSliceV("1", "2", "3").Set(0, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, OldSliceV("1", "2", "3").Set(1, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, OldSliceV("1", "2", "3").Set(2, "0").O())
		assert.Equal(t, []string{"0", "2", "3"}, OldSliceV("1", "2", "3").Set(-3, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, OldSliceV("1", "2", "3").Set(-2, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, OldSliceV("1", "2", "3").Set(-1, "0").O())
	}

	// custom
	{
		assert.Equal(t, []Object{{0}, {2}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(0, Object{0}).O())
		assert.Equal(t, []Object{{1}, {0}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(1, Object{0}).O())
		assert.Equal(t, []Object{{1}, {2}, {0}}, OldSlice([]Object{{1}, {2}, {3}}).Set(2, Object{0}).O())
		assert.Equal(t, []Object{{0}, {2}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(-3, Object{0}).O())
		assert.Equal(t, []Object{{1}, {0}, {3}}, OldSlice([]Object{{1}, {2}, {3}}).Set(-2, Object{0}).O())
		assert.Equal(t, []Object{{1}, {2}, {0}}, OldSlice([]Object{{1}, {2}, {3}}).Set(-1, Object{0}).O())
	}

	// panics need to run as the last test as they abort the test method
	defer func() {
		err := recover()
		assert.Equal(t, "slice assignment is out of bounds", err)
	}()
	OldSliceV(1, 2, 3).Set(5, 1)
}

// Single
//--------------------------------------------------------------------------------------------------

func ExampleNSlice_Single() {
	slice := OldSliceV(1)
	fmt.Println(slice.Single())
	// Output: true
}

func TestNSlice_Single(t *testing.T) {

	// int
	{
		assert.Equal(t, false, OldSliceV().Single())
		assert.Equal(t, true, OldSliceV(1).Single())
		assert.Equal(t, false, OldSliceV(1, 2).Single())
	}

	// custom
	{
		assert.Equal(t, false, OldSliceV().Single())
		assert.Equal(t, true, OldSliceV(Object{1}).Single())
		assert.Equal(t, false, OldSliceV(Object{1}, Object{2}).Single())
	}
}

// Slice
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Slice_Normal(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:len(ints)]
}

func BenchmarkNSlice_Slice_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	slice.Slice(0, -1)
}

func BenchmarkNSlice_Slice_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	slice.Slice(0, -1)
}

func ExampleNSlice_Slice() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Slice(1, -1).O())
	// Output: [2 3]
}

func TestNSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, OldSliceV(), nilSlice.Slice(0, -1))
		slice := OldSliceV(0).Clear()
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(0, -1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := OldSliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, []int{0, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2, 3}, result.O())
	}

	// slice full array
	{
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Slice(0, -1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Slice(0, 1))
		assert.Equal(t, &NSlice{o: []interface{}{}}, OldSliceV().Slice(0, 5))
		assert.Equal(t, OldSliceV(""), OldSliceV("").Slice(0, -1))
		assert.Equal(t, OldSliceV(""), OldSliceV("").Slice(0, 1))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).Slice(0, -1))
		assert.Equal(t, OldSlice([]int{1, 2, 3}), OldSlice([]int{1, 2, 3}).Slice(0, -1))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Slice(0, 2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, OldSliceV("1"), OldSliceV("1").Slice(0, 2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false).Slice(-6, 6))
		assert.Equal(t, OldSliceV(1, 2, 3), OldSliceV(1, 2, 3).Slice(-6, 6))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Slice(-6, 6))
		assert.Equal(t, OldSlice([]Object{{1}, {2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-6, 6))
	}

	// mutually exclusive
	{
		slice := OldSliceV(1, 2, 3, 4)
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(2, -3))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(0, -5))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(4, -1))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(6, -1))
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(3, 2))
	}

	// singles
	{
		slice := OldSliceV(1, 2, 3, 4)
		assert.Equal(t, OldSliceV(4), slice.Slice(-1, -1))
		assert.Equal(t, OldSliceV(3), slice.Slice(-2, -2))
		assert.Equal(t, OldSliceV(2), slice.Slice(-3, -3))
		assert.Equal(t, OldSliceV(1), slice.Slice(0, 0))
		assert.Equal(t, OldSliceV(1), slice.Slice(-4, -4))
		assert.Equal(t, OldSliceV(2), slice.Slice(1, 1))
		assert.Equal(t, OldSliceV(2), slice.Slice(1, -3))
		assert.Equal(t, OldSliceV(3), slice.Slice(2, 2))
		assert.Equal(t, OldSliceV(3), slice.Slice(2, -2))
		assert.Equal(t, OldSliceV(4), slice.Slice(3, 3))
		assert.Equal(t, OldSliceV(4), slice.Slice(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(1, -1))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(1, 2))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(-2, -1))
		assert.Equal(t, OldSliceV(false, true), OldSliceV(true, false, true).Slice(-2, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(1, -1))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(-2, -1))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3).Slice(-2, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(1, -1))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(-2, -1))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(-2, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(1, -1))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(1, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-2, -1))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(0, -2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(-3, -2))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(-3, 1))
		assert.Equal(t, OldSliceV(true, false), OldSliceV(true, false, true).Slice(0, 1))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(0, -2))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(-3, -2))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(-3, 1))
		assert.Equal(t, OldSliceV(1, 2), OldSliceV(1, 2, 3).Slice(0, 1))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(0, -2))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(-3, -2))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(-3, 1))
		assert.Equal(t, OldSliceV("1", "2"), OldSliceV("1", "2", "3").Slice(0, 1))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(0, -2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-3, -2))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(-3, 1))
		assert.Equal(t, OldSlice([]Object{{1}, {2}}), OldSlice([]Object{{1}, {2}, {3}}).Slice(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(1, -2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(-3, -2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(-3, 2))
		assert.Equal(t, OldSliceV(true, true), OldSliceV(false, true, true, false).Slice(1, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(1, -2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(-3, -2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(-3, 2))
		assert.Equal(t, OldSliceV(2, 3), OldSliceV(1, 2, 3, 4).Slice(1, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(1, -2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(-3, -2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(-3, 2))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3", "4").Slice(1, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, -2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, -2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, 2))
		assert.Equal(t, OldSlice([]Object{{2}, {3}}), OldSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, 2))
	}

	// random
	{
		assert.Equal(t, OldSliceV("1"), OldSliceV("1", "2", "3").Slice(0, -3))
		assert.Equal(t, OldSliceV("2", "3"), OldSliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, OldSliceV("1", "2", "3"), OldSliceV("1", "2", "3").Slice(0, 2))
	}
}

// Sort
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Sort_Normal(t *testing.B) {
	ints := Range(0, nines7)
	_ = ints[0:len(ints)]
}

func BenchmarkNSlice_Sort_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines7))
	slice.Slice(0, -1)
}

func BenchmarkNSlice_Sort_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines7))
	slice.Slice(0, -1)
}

func ExampleNSlice_Sort() {
	slice := OldSliceV(2, 3, 1)
	fmt.Println(slice.Sort().O())
	// Output: [1 2 3]
}

func TestNSlice_Sort(t *testing.T) {

	// empty
	//assert.Equal(t, SliceV(), SliceV().Sort())

	// bool
	//assert.Equal(t, SliceV(false, true, true), SliceV(true, false, true).Sort())

	// int
	assert.Equal(t, OldSliceV(1, 2, 3, 4, 5), OldSliceV(5, 3, 2, 4, 1).Sort())

	// string
	//assert.Equal(t, SliceV("1", "2", "3", "4", "5"), SliceV("5", "3", "2", "4", "1").Sort())

	// custom
	//assert.Equal(t, Slice([]Object{{1}, {2}, {3}, {4}, {5}}), Slice([]Object{{5}, {3}, {2}, {4}, {1}}).Sort())
}

// Swap
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Swap_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		if i+1 < len(ints) {
			ints[i], ints[i+1] = ints[i+1], ints[i]
		}
	}
}

func BenchmarkNSlice_Swap_Optimized(t *testing.B) {
	slice := OldSlice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Swap(i, i+1)
		}
	}
}

func BenchmarkNSlice_Swap_Reflect(t *testing.B) {
	slice := OldSlice(rangeNObject(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		if i+1 < slice.Len() {
			slice.Swap(i, i+1)
		}
	}
}

func ExampleNSlice_Swap() {
	slice := OldSliceV(2, 3, 1)
	slice.Swap(0, 2)
	slice.Swap(1, 2)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_Swap(t *testing.T) {

	// invalid cases
	{
		var slice *NSlice
		slice.Swap(0, 0)
		assert.Equal(t, (*NSlice)(nil), slice)

		slice = OldSliceV()
		slice.Swap(0, 0)
		assert.Equal(t, OldSliceV(), slice)

		slice.Swap(1, 2)
		assert.Equal(t, OldSliceV(), slice)

		slice.Swap(-1, 2)
		assert.Equal(t, OldSliceV(), slice)

		slice.Swap(1, -2)
		assert.Equal(t, OldSliceV(), slice)
	}

	// bool
	{
		slice := OldSliceV(true, false, true)
		slice.Swap(0, 1)
		assert.Equal(t, OldSliceV(false, true, true), slice)
	}

	// int
	{
		slice := OldSliceV(0, 1, 2)
		slice.Swap(0, 1)
		assert.Equal(t, OldSliceV(1, 0, 2), slice)
	}

	// string
	{
		slice := OldSliceV("0", "1", "2")
		slice.Swap(0, 1)
		assert.Equal(t, OldSliceV("1", "0", "2"), slice)
	}

	// custom
	{
		slice := OldSlice([]Object{{0}, {1}, {2}})
		slice.Swap(0, 1)
		assert.Equal(t, OldSlice([]Object{{1}, {0}, {2}}), slice)
	}
}

// Take
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Take_Normal(t *testing.B) {
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

func BenchmarkNSlice_Take_Optimized(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := OldSlice(src)
	for i := range index {
		slice.Take(i)
	}
}

func BenchmarkNSlice_Take_Reflect(t *testing.B) {
	src := rangeNObject(0, nines5)
	index := Range(0, nines5)
	slice := OldSlice(src)
	for i := range index {
		slice.Take(i)
	}
}

func ExampleNSlice_Take() {
	slice := OldSliceV(1, 2, 3)
	fmt.Println(slice.Take(2).O())
	// Output: 3
}

func TestNSlice_Take(t *testing.T) {

	// int
	{
		// nil or empty
		{
			var nilSlice *NSlice
			assert.Equal(t, &Object{}, nilSlice.Take(0))
		}

		// Delete all and more
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-1)
			assert.Equal(t, &Object{2}, obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{1}, obj)
			assert.Equal(t, []int{0}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []int{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.Take(-1)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []int{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(3)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(2)
			assert.Equal(t, &Object{2}, obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(1)
			assert.Equal(t, &Object{1}, obj)
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(0)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []int{1, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-4)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []int{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-1)
			assert.Equal(t, &Object{2}, obj)
			assert.Equal(t, []int{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := OldSliceV(0, 1, 2)
			obj := slice.Take(-2)
			assert.Equal(t, &Object{1}, obj)
			assert.Equal(t, []int{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// custom
	{
		// Delete all and more
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.Take(-1)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []Object{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.Take(-1)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []Object{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(3)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(2)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(0)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []Object{{1}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-4)
			assert.Equal(t, &Object{nil}, obj)
			assert.Equal(t, []Object{{0}, {1}, {2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []Object{{0}, {1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := OldSlice([]Object{{0}, {1}, {2}})
			obj := slice.Take(-2)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []Object{{0}, {2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}
}

// // func TestStrSliceTakeFirst(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		results := []string{}
// // 		expected := []string{"0", "1", "2"}
// // 		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
// // 			results = append(results, item)
// // 		}
// // 		assert.Equal(t, expected, results)
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		item, ok := slice.TakeFirst()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{"1", "2"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0")
// // 		item, ok := slice.TakeFirst()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S()
// // 		item, ok := slice.TakeFirst()
// // 		assert.False(t, ok)
// // 		assert.Equal(t, "", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceTakeFirstCnt(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(2).S()
// // 		assert.Equal(t, []string{"0", "1"}, items)
// // 		assert.Equal(t, []string{"2"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(3).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(4).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceTakeLast(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		results := []string{}
// // 		expected := []string{"2", "1", "0"}
// // 		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
// // 			results = append(results, item)
// // 		}
// // 		assert.Equal(t, expected, results)
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		item, ok := slice.TakeLast()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "2", item)
// // 		assert.Equal(t, []string{"0", "1"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0")
// // 		item, ok := slice.TakeLast()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S()
// // 		item, ok := slice.TakeLast()
// // 		assert.False(t, ok)
// // 		assert.Equal(t, "", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }
// // func TestStrSliceTakeLastCnt(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(2).S()
// // 		assert.Equal(t, []string{"1", "2"}, items)
// // 		assert.Equal(t, []string{"0"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(3).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(4).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceUniq(t *testing.T) {
// // 	{
// // 		data := S().Uniq().S()
// // 		expected := []string{}
// // 		assert.Equal(t, expected, data)
// // 	}
// // 	{
// // 		data := S("1", "2", "3").Uniq().S()
// // 		expected := []string{"1", "2", "3"}
// // 		assert.Equal(t, expected, data)
// // 	}
// // 	{
// // 		data := S("1", "2", "2", "3").Uniq().S()
// // 		expected := []string{"1", "2", "3"}
// // 		assert.Equal(t, expected, data)
// // 	}
// // }

// // func TestYamlPair(t *testing.T) {
// // 	{
// // 		k, v := A("foo=bar").Split("=").YamlPair()
// // 		assert.Equal(t, "foo", k)
// // 		assert.Equal(t, "bar", v)
// // 	}
// // 	{
// // 		k, v := A("=bar").Split("=").YamlPair()
// // 		assert.Equal(t, "", k)
// // 		assert.Equal(t, "bar", v)
// // 	}
// // 	{
// // 		k, v := A("bar=").Split("=").YamlPair()
// // 		assert.Equal(t, "bar", k)
// // 		assert.Equal(t, "", v)
// // 	}
// // 	{
// // 		k, v := A("").Split("=").YamlPair()
// // 		assert.Equal(t, "", k)
// // 		assert.Equal(t, nil, v)
// // 	}
// // }
// // func TestYamlKeyVal(t *testing.T) {
// // 	{
// // 		pair := A("foo=bar").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "foo", pair.Key)
// // 		assert.Equal(t, "bar", pair.Val)
// // 	}
// // 	{
// // 		pair := A("=bar").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "", pair.Key)
// // 		assert.Equal(t, "bar", pair.Val)
// // 	}
// // 	{
// // 		pair := A("bar=").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "bar", pair.Key)
// // 		assert.Equal(t, "", pair.Val)
// // 	}
// // 	{
// // 		pair := A("").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "", pair.Key)
// // 		assert.Equal(t, "", pair.Val)
// // 	}
// // }

func TestSlice_absIndex(t *testing.T) {
	//             -4,-3,-2,-1
	//              0, 1, 2, 3
	assert.Equal(t, 3, absIndex(4, -1))
	assert.Equal(t, 2, absIndex(4, -2))
	assert.Equal(t, 1, absIndex(4, -3))
	assert.Equal(t, 0, absIndex(4, -4))

	assert.Equal(t, 0, absIndex(4, 0))
	assert.Equal(t, 1, absIndex(4, 1))
	assert.Equal(t, 2, absIndex(4, 2))
	assert.Equal(t, 3, absIndex(4, 3))

	// out of bounds
	assert.Equal(t, -1, absIndex(4, 4))
	assert.Equal(t, -1, absIndex(4, -5))
}

func TestSlice_absIndices(t *testing.T) {
	len := 4
	// -4,-3,-2,-1
	//  0, 1, 2, 3

	// no indicies given
	{
		i, j, err := absIndices(len)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// one index given
	{
		i, j, err := absIndices(len, 1)
		assert.Equal(t, 0, i)
		assert.Equal(t, -1, j)
		assert.Equal(t, "only one index given", err.Error())
	}

	// end
	{
		i, j, err := absIndices(len, -3, -1)
		assert.Equal(t, 1, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 1, 3)
		assert.Equal(t, 1, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// middle
	{
		i, j, err := absIndices(len, 1, 2)
		assert.Equal(t, 1, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -3, -2)
		assert.Equal(t, 1, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)
	}

	// begining
	{
		i, j, err := absIndices(len, 0, 2)
		assert.Equal(t, 0, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -4, -2)
		assert.Equal(t, 0, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)
	}

	// move within bounds
	{
		i, j, err := absIndices(len, -5, 5)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 0, 5)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -5, -1)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// mutually exclusive
	{
		i, j, err := absIndices(len, -1, -3)
		assert.Equal(t, 3, i)
		assert.Equal(t, 1, j)
		assert.NotNil(t, err)

		i, j, err = absIndices(len, 3, 1)
		assert.Equal(t, 3, i)
		assert.Equal(t, 1, j)
		assert.NotNil(t, err)
	}

	// single
	{
		i, j, err := absIndices(len, 0, 0)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 1, 1)
		assert.Equal(t, 1, i)
		assert.Equal(t, 2, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 3, 3)
		assert.Equal(t, 3, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -1, -1)
		assert.Equal(t, 3, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -2, -2)
		assert.Equal(t, 2, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -4, -4)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
		assert.Nil(t, err)
	}
}
