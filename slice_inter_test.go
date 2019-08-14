package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewInterSlice function
//--------------------------------------------------------------------------------------------------
func ExampleNewInterSlice() {
	slice := NewInterSlice([]int{1, 2, 3}).O()
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestInterSlice_NewInterSlice(t *testing.T) {

	// arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, []interface{}{"1", "2"}, NewInterSlice(array).O())

	// empty
	assert.Equal(t, []interface{}{}, NewInterSlice(nil).O())
	assert.Equal(t, &InterSlice{}, NewInterSlice(nil))
	assert.Equal(t, []interface{}{}, NewInterSlice([]int{}).O())
	assert.Equal(t, []interface{}{}, NewInterSlice([]bool{}).O())
	assert.Equal(t, []interface{}{}, NewInterSlice([]string{}).O())
	assert.Equal(t, []interface{}{}, NewInterSlice([]Object{}).O())
	assert.Equal(t, []interface{}{}, NewInterSlice([]interface{}{}).O())

	// pointers
	var obj *Object
	assert.Equal(t, []interface{}{(*Object)(nil)}, NewInterSlice(obj).O())
	assert.Equal(t, []interface{}{&(Object{"bob"})}, NewInterSlice(&(Object{"bob"})).O())
	assert.Equal(t, []interface{}{&(Object{"1"}), &(Object{"2"})}, NewInterSlice([]*Object{&(Object{"1"}), &(Object{"2"})}).O())

	// interface
	assert.Equal(t, []interface{}{nil}, NewInterSlice([]interface{}{nil}).O())
	assert.Equal(t, []interface{}{nil, ""}, NewInterSlice([]interface{}{nil, ""}).O())
	assert.Equal(t, []interface{}{true}, NewInterSlice([]interface{}{true}).O())
	assert.Equal(t, []interface{}{1}, NewInterSlice([]interface{}{1}).O())
	assert.Equal(t, []interface{}{""}, NewInterSlice([]interface{}{""}).O())
	assert.Equal(t, []interface{}{"bob"}, NewInterSlice([]interface{}{"bob"}).O())
	assert.Equal(t, []interface{}{Object{}}, NewInterSlice([]interface{}{Object{}}).O())

	// singles
	assert.Equal(t, []interface{}{1}, NewInterSlice(1).O())
	assert.Equal(t, []interface{}{true}, NewInterSlice(true).O())
	assert.Equal(t, []interface{}{""}, NewInterSlice("").O())
	assert.Equal(t, []interface{}{"1"}, NewInterSlice("1").O())
	assert.Equal(t, []interface{}{Object{1}}, NewInterSlice(Object{1}).O())
	assert.Equal(t, []interface{}{Object{"bob"}}, NewInterSlice(Object{"bob"}).O())
	assert.Equal(t, []interface{}{map[string]string{"1": "one"}}, NewInterSlice(map[string]string{"1": "one"}).O())

	// slices
	assert.Equal(t, []interface{}{1, 2}, NewInterSlice([]int{1, 2}).O())
	assert.Equal(t, []interface{}{true}, NewInterSlice([]bool{true}).O())
	assert.Equal(t, []interface{}{Object{"bob"}}, NewInterSlice([]Object{{"bob"}}).O())
	assert.Equal(t, []interface{}{"1", "2"}, NewInterSlice([]string{"1", "2"}).O())
	assert.Equal(t, []interface{}{[]string{"1"}}, NewInterSlice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []interface{}{map[string]string{"1": "one"}}, NewInterSlice([]interface{}{map[string]string{"1": "one"}}).O())
}

// NewInterSliceV function
//--------------------------------------------------------------------------------------------------
func ExampleNewInterSliceV_empty() {
	slice := NewInterSliceV()
	fmt.Println(slice.O())
	// Output: []
}

func ExampleNewInterSliceV_variadic() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestInterSlice_NewInterSliceV(t *testing.T) {
	var obj *Object

	// Arrays
	{
		var array [2]string
		array[0] = "1"
		array[1] = "2"
		assert.Equal(t, []interface{}{array}, NewInterSliceV(array).O())
	}

	// Test empty values
	{
		assert.True(t, !NewInterSliceV().Any())
		assert.Equal(t, 0, NewInterSliceV().Len())
		assert.Equal(t, []interface{}{}, NewInterSliceV().O())
		assert.Equal(t, []interface{}{nil}, NewInterSliceV(nil).O())
		assert.Equal(t, &InterSlice{nil}, NewInterSliceV(nil))
		assert.Equal(t, []interface{}{nil, ""}, NewInterSliceV(nil, "").O())
		assert.Equal(t, []interface{}{nil, Object{}}, NewInterSliceV(nil, Object{}).O())
	}

	// Test pointers
	{
		assert.Equal(t, []interface{}{(*Object)(nil)}, NewInterSliceV(obj).O())
		assert.Equal(t, []interface{}{&(Object{"bob"})}, NewInterSliceV(&(Object{"bob"})).O())
		assert.Equal(t, []interface{}{[]*Object{&Object{"1"}, &Object{"2"}}}, NewInterSliceV([]*Object{&(Object{"1"}), &(Object{"2"})}).O())
	}

	// Singles
	{
		assert.Equal(t, []interface{}{1}, NewInterSliceV(1).O())
		assert.Equal(t, []interface{}{"1"}, NewInterSliceV("1").O())
		assert.Equal(t, []interface{}{Object{"bob"}}, NewInterSliceV(Object{"bob"}).O())
		assert.Equal(t, []interface{}{map[string]string{"1": "one"}}, NewInterSliceV(map[string]string{"1": "one"}).O())
	}

	// Multiples
	{
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2").O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSliceV(Object{1}, Object{2}).O())
	}

	// Test slices
	{
		assert.Equal(t, []interface{}{[]int{1, 2}}, NewInterSliceV([]int{1, 2}).O())
		assert.Equal(t, []interface{}{[]string{"1"}}, NewInterSliceV([]string{"1"}).O())
	}
}

// Any
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Any_empty() {
	slice := NewInterSliceV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleInterSlice_Any_notEmpty() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleInterSlice_Any_contains() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Any(1))
	// Output: true
}

func ExampleInterSlice_Any_containsAny() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Any(0, 1))
	// Output: true
}

func TestInterSlice_Any(t *testing.T) {
	var slice *InterSlice
	assert.False(t, slice.Any())
	assert.False(t, NewInterSliceV().Any())
	assert.True(t, NewInterSliceV().Append("2").Any())

	// bool
	assert.True(t, NewInterSliceV(false, true).Any(true))
	assert.False(t, NewInterSliceV(true, true).Any(false))
	assert.True(t, NewInterSliceV(true, true).Any(false, true))
	assert.False(t, NewInterSliceV(true, true).Any(false, false))

	// int
	assert.True(t, NewInterSliceV(1, 2, 3).Any(2))
	assert.False(t, NewInterSliceV(1, 2, 3).Any(4))
	assert.True(t, NewInterSliceV(1, 2, 3).Any(4, 3))
	assert.False(t, NewInterSliceV(1, 2, 3).Any(4, 5))

	// int64
	assert.True(t, NewInterSliceV(int64(1), int64(2), int64(3)).Any(int64(2)))
	assert.False(t, NewInterSliceV(int64(1), int64(2), int64(3)).Any(int64(4)))
	assert.True(t, NewInterSliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(2)))
	assert.False(t, NewInterSliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(5)))

	// string
	assert.True(t, NewInterSliceV("1", "2", "3").Any("2"))
	assert.False(t, NewInterSliceV("1", "2", "3").Any("4"))
	assert.True(t, NewInterSliceV("1", "2", "3").Any("4", "2"))
	assert.False(t, NewInterSliceV("1", "2", "3").Any("4", "5"))

	// custom
	assert.True(t, NewInterSliceV(Object{1}, Object{2}).Any(Object{1}))
	assert.False(t, NewInterSliceV(Object{1}, Object{2}).Any(Object{3}))
	assert.True(t, NewInterSliceV(Object{1}, Object{2}).Any(Object{4}, Object{2}))
	assert.False(t, NewInterSliceV(Object{1}, Object{2}).Any(Object{4}, Object{5}))
}

// AnyS
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_AnyS() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.AnyS([]int{0, 1}))
	// Output: true
}

func TestInterSlice_AnyS(t *testing.T) {
	var nilSlice *InterSlice
	assert.False(t, nilSlice.AnyS([]bool{true}))

	// bool
	assert.True(t, NewInterSliceV(true, true).AnyS([]bool{true}))
	assert.True(t, NewInterSliceV(true, true).AnyS([]bool{false, true}))
	assert.False(t, NewInterSliceV(true, true).AnyS([]bool{false, false}))

	// int
	assert.True(t, NewInterSliceV(1, 2, 3).AnyS([]int{1}))
	assert.True(t, NewInterSliceV(1, 2, 3).AnyS([]int{4, 3}))
	assert.False(t, NewInterSliceV(1, 2, 3).AnyS([]int{4, 5}))

	// int64
	assert.True(t, NewInterSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(2)}))
	assert.True(t, NewInterSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(2)}))
	assert.False(t, NewInterSliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(5)}))

	// string
	assert.True(t, NewInterSliceV("1", "2", "3").AnyS([]string{"2"}))
	assert.True(t, NewInterSliceV("1", "2", "3").AnyS([]string{"4", "2"}))
	assert.False(t, NewInterSliceV("1", "2", "3").AnyS([]string{"4", "5"}))

	// custom
	assert.True(t, NewInterSliceV(Object{1}, Object{2}).AnyS([]Object{{2}}))
	assert.True(t, NewInterSliceV(Object{1}, Object{2}).AnyS([]Object{{4}, {2}}))
	assert.False(t, NewInterSliceV(Object{1}, Object{2}).AnyS([]Object{{4}, {5}}))

	// NewInterSliceV
	assert.True(t, NewInterSliceV(1, 2).AnyS(NewInterSliceV(1, 3)))
	assert.False(t, NewInterSliceV(1, 2).AnyS(NewInterSliceV(4, 3)))

	// NewIntSliceV
	assert.True(t, NewInterSliceV(1, 2).AnyS(NewIntSliceV(1, 3)))
	assert.False(t, NewInterSliceV(1, 2).AnyS(NewIntSliceV(4, 3)))

	// nils
	assert.False(t, NewInterSliceV(1, 2).AnyS(nil))
	assert.False(t, NewInterSliceV(1, 2).AnyS((*[]int)(nil)))
	assert.False(t, NewInterSliceV(1, 2).AnyS((*IntSlice)(nil)))
	assert.False(t, NewInterSliceV(1, 2).AnyS((*InterSlice)(nil)))
}

// AnyW
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_AnyW() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.AnyW(func(x O) bool {
		return ExB(x.(int) == 2)
	}))
	// Output: true
}

func TestInterSlice_AnyW(t *testing.T) {

	// empty
	var slice *InterSlice
	assert.False(t, slice.AnyW(func(x O) bool { return ExB(x.(int) > 0) }))
	assert.False(t, NewInterSliceV().AnyW(func(x O) bool { return ExB(x.(int) > 0) }))

	// single
	assert.True(t, NewInterSliceV(2).AnyW(func(x O) bool { return ExB(x.(int) > 0) }))

	assert.True(t, NewInterSliceV(1, 2).AnyW(func(x O) bool { return ExB(x.(int) == 2) }))
	assert.True(t, NewInterSliceV(1, 2).AnyW(func(x O) bool { return ExB(x.(int)%2 == 0) }))
	assert.False(t, NewInterSliceV(2, 4).AnyW(func(x O) bool { return ExB(x.(int)%2 != 0) }))
	assert.True(t, NewInterSliceV(1, 2, 3).AnyW(func(x O) bool { return ExB(x.(int) == 4 || x.(int) == 3) }))
}

// Append
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Append() {
	slice := NewInterSliceV(1).Append(2).Append(3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestInterSlice_Append(t *testing.T) {

	// nil
	{
		var slice *InterSlice
		assert.Equal(t, []interface{}{0}, slice.Append(0).O())
		assert.Equal(t, (*InterSlice)(nil), slice)
	}

	// Append one back to back
	{
		n := NewInterSliceV()
		assert.Equal(t, 0, n.Len())
		assert.Equal(t, false, n.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		n.Append("1")
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []interface{}{"1"}, n.O())

		// Second append another which will be 2x at most
		n.Append("2")
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []interface{}{"1", "2"}, n.O())
	}

	// Start with just appending without chaining
	{
		n := NewInterSliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1)
		assert.Equal(t, []interface{}{1}, n.O())
		n.Append(2)
		assert.Equal(t, []interface{}{1, 2}, n.O())
	}

	// Start with nil not chained
	{
		n := NewInterSliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []interface{}{1, 2, 3}, n.O())
	}

	// Start with nil chained
	{
		n := NewInterSliceV().Append(1).Append(2)
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []interface{}{1, 2}, n.O())
	}

	// Start with non nil
	{
		n := NewInterSliceV(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []interface{}{1, 2, 3}, n.O())
	}

	// Use append result directly
	{
		n := NewInterSliceV(1)
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []interface{}{1, 2}, n.Append(2).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewInterSliceV(true)
			assert.Equal(t, []interface{}{true, false}, n.Append(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := NewInterSliceV(0)
			assert.Equal(t, []interface{}{0, 1}, n.Append(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := NewInterSliceV("0")
			assert.Equal(t, []interface{}{"0", "1"}, n.Append("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewInterSlice([]Object{{"3"}})
			assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}}, n.Append(Object{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}
}

func TestInterSlice_Append_MixedTypes(t *testing.T) {
	n := NewInterSliceV(true)
	assert.Equal(t, []interface{}{true}, n.G())

	n.Append("2")
	assert.Equal(t, []interface{}{true, "2"}, n.G())

	n.Append(3)
	assert.Equal(t, []interface{}{true, "2", 3}, n.G())

	n.Append(4.1)
	assert.Equal(t, []interface{}{true, "2", 3, 4.1}, n.G())
}

// AppendV
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_AppendV() {
	slice := NewInterSliceV(1).AppendV(2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestInterSlice_AppendV(t *testing.T) {

	// nil
	{
		var slice *InterSlice
		assert.Equal(t, []interface{}{0}, slice.AppendV(0).O())
		assert.Equal(t, (*InterSlice)(nil), slice)
	}

	// Append many ints
	{
		slice := NewInterSliceV(1)
		assert.Equal(t, []interface{}{1, 2, 3}, slice.AppendV(2, 3).O())
	}

	// Append many strings
	{
		{
			slice := NewInterSliceV()
			assert.Equal(t, 0, slice.Len())
			assert.Equal(t, []interface{}{"1", "2", "3"}, slice.AppendV("1", "2", "3").O())
			assert.Equal(t, 3, slice.Len())
		}
		{
			slice := NewInterSlice([]string{"1"})
			assert.Equal(t, 1, slice.Len())
			assert.Equal(t, []interface{}{"1", "2", "3"}, slice.AppendV("2", "3").O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// Append to a slice of custom type
	{
		slice := NewInterSlice([]Object{{"3"}})
		assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}}, slice.AppendV(Object{"1"}).O())
		assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}, Object{"2"}, Object{"4"}}, slice.AppendV(Object{"2"}, Object{"4"}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewInterSliceV(true)
			assert.Equal(t, []interface{}{true, false}, n.AppendV(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := NewInterSliceV(0)
			assert.Equal(t, []interface{}{0, 1}, n.AppendV(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := NewInterSliceV("0")
			assert.Equal(t, []interface{}{"0", "1"}, n.AppendV("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewInterSlice([]Object{{"3"}})
			assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}}, n.AppendV(Object{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}

	// Append to a slice of map
	{
		n := NewInterSliceV(map[string]string{"1": "one"})
		expected := []interface{}{
			map[string]string{"1": "one"},
			map[string]string{"2": "two"},
		}
		assert.Equal(t, expected, n.AppendV(map[string]string{"2": "two"}).O())
	}
}

// At
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_At() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.At(2).O())
	// Output: 3
}

func TestInterSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *InterSlice
		assert.Equal(t, Obj(nil), nilSlice.At(0))
	}

	// strings
	{
		slice := NewInterSliceV("1", "2", "3", "4")
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
		slice := NewInterSliceV("1")
		assert.Equal(t, &Object{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &Object{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Clear() {
	slice := NewInterSliceV(1).ConcatM([]int{2, 3})
	fmt.Println(slice.Clear().O())
	// Output: []
}

func TestInterSlice_Clear(t *testing.T) {

	// nil
	{
		var nilSlice *InterSlice
		assert.Equal(t, Obj(nil), nilSlice.At(0))
	}

	// bool
	{
		slice := NewInterSliceV(true, false)
		assert.Equal(t, 2, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.O())
	}

	// int
	{
		slice := NewInterSliceV(1, 2, 3, 4)
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.O())
	}

	// string
	{
		slice := NewInterSliceV("1", "2", "3", "4")
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.O())
	}

	// custom
	{
		slice := NewInterSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, 3, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.O())
	}
}

// Concat
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Concat() {
	slice := NewInterSliceV(1).Concat([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestInterSlice_Concat(t *testing.T) {

	// nil
	{
		var slice *InterSlice
		assert.Equal(t, []interface{}{1, 2}, slice.Concat([]int{1, 2}).O())
		assert.Equal(t, (*InterSlice)(nil), slice)
	}

	// Append many ints
	{
		n := NewInterSliceV(1)
		assert.Equal(t, []interface{}{1, 2, 3}, n.Concat([]int{2, 3}).O())
	}

	// Append many strings
	{
		{
			n := NewInterSliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []interface{}{"1", "2", "3"}, n.Concat([]string{"1", "2", "3"}).O())
			assert.Equal(t, 0, n.Len())
		}
		{
			n := NewInterSlice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []interface{}{"1", "2", "3"}, n.Concat([]string{"2", "3"}).O())
			assert.Equal(t, 1, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := NewInterSlice([]Object{{"3"}})
		assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}}, n.Concat([]Object{{"1"}}).O())
		assert.Equal(t, []interface{}{Object{"3"}, Object{"2"}, Object{"4"}}, n.Concat([]Object{{"2"}, {"4"}}).O())
	}

	// Append to a slice of map
	{
		n := NewInterSliceV(map[string]string{"1": "one"})
		expected := []interface{}{
			map[string]string{"1": "one"},
			map[string]string{"2": "two"},
		}
		assert.Equal(t, expected, n.Concat([]map[string]string{{"2": "two"}}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewInterSliceV(true)
			assert.Equal(t, []interface{}{true, false}, n.Concat([]bool{false}).O())
			assert.Equal(t, 1, n.Len())
		}

		// int
		{
			n := NewInterSliceV(0)
			assert.Equal(t, []interface{}{0, 1}, n.Concat([]int{1}).O())
			assert.Equal(t, 1, n.Len())
		}

		// string
		{
			n := NewInterSliceV("0")
			assert.Equal(t, []interface{}{"0", "1"}, n.Concat([]string{"1"}).O())
			assert.Equal(t, 1, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewInterSlice([]Object{{"3"}})
			assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}}, n.Concat([]Object{{"1"}}).O())
			assert.Equal(t, 1, n.Len())
		}
	}

	// nils
	{
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2).Concat((*[]int)(nil)).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2).Concat((*IntSlice)(nil)).O())
	}
}

// ConcatM
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_ConcatM() {
	slice := NewInterSliceV(1).ConcatM([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestInterSlice_ConcatM(t *testing.T) {

	// nil
	{
		var slice *InterSlice
		assert.Equal(t, []interface{}{1, 2}, slice.ConcatM([]int{1, 2}).O())
		assert.Equal(t, (*InterSlice)(nil), slice)
	}

	// Append many ints
	{
		n := NewInterSliceV(1)
		assert.Equal(t, []interface{}{1, 2, 3}, n.ConcatM([]int{2, 3}).O())
	}

	// Append many strings
	{
		{
			n := NewInterSliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []interface{}{"1", "2", "3"}, n.ConcatM([]string{"1", "2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := NewInterSlice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []interface{}{"1", "2", "3"}, n.ConcatM([]string{"2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := NewInterSlice([]Object{{"3"}})
		assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}}, n.ConcatM([]Object{{"1"}}).O())
		assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}, Object{"2"}, Object{"4"}}, n.ConcatM([]Object{{"2"}, {"4"}}).O())
	}

	// Append to a slice of map
	{
		n := NewInterSliceV(map[string]string{"1": "one"})
		expected := []interface{}{
			map[string]string{"1": "one"},
			map[string]string{"2": "two"},
		}
		assert.Equal(t, expected, n.ConcatM([]map[string]string{{"2": "two"}}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := NewInterSliceV(true)
			assert.Equal(t, []interface{}{true, false}, n.ConcatM([]bool{false}).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := NewInterSliceV(0)
			assert.Equal(t, []interface{}{0, 1}, n.ConcatM([]int{1}).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := NewInterSliceV("0")
			assert.Equal(t, []interface{}{"0", "1"}, n.ConcatM([]string{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := NewInterSlice([]Object{{"3"}})
			assert.Equal(t, []interface{}{Object{"3"}, Object{"1"}}, n.ConcatM([]Object{{"1"}}).O())
			assert.Equal(t, 2, n.Len())
		}

		// NewIntSliceV
		{
			n := NewInterSliceV(0)
			assert.Equal(t, []interface{}{0, 1}, n.ConcatM(NewIntSliceV(1)).O())
			assert.Equal(t, 2, n.Len())
		}

		// NewInterSliceV
		{
			n := NewInterSliceV(0)
			assert.Equal(t, []interface{}{0, 1}, n.ConcatM(NewInterSliceV(1)).O())
			assert.Equal(t, 2, n.Len())
		}

		// NewInterSlice
		{
			n := NewInterSliceV(0)
			assert.Equal(t, []interface{}{0, 1}, n.ConcatM(NewInterSlice([]int{1})).O())
			assert.Equal(t, 2, n.Len())
		}

		// nils
		{
			assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2).ConcatM((*[]int)(nil)).O())
			assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2).ConcatM((*IntSlice)(nil)).O())
		}
	}
}

func TestInterSlice_ConcatM_notASliceType(t *testing.T) {
	slice := NewInterSliceV(1)
	slice.ConcatM("2")
	assert.Equal(t, []interface{}{1, "2"}, slice.G())
}

// Copy
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Copy() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Copy().O())
	// Output: [1 2 3]
}

func TestInterSlice_Copy(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, NewInterSliceV(), nilSlice.Copy(0, -1))
		slice := NewInterSliceV(0).Clear()
		assert.Equal(t, []interface{}{}, slice.Copy(0, -1).O())
	}

	// Test that the original is NOT modified when the slice is modified
	{
		original := NewInterSliceV(1, 2, 3)
		result := original.Copy(0, -1)
		assert.Equal(t, []interface{}{1, 2, 3}, original.O())
		assert.Equal(t, []interface{}{1, 2, 3}, result.O())
		result.Set(0, 0)
		assert.Equal(t, []interface{}{1, 2, 3}, original.O())
		assert.Equal(t, []interface{}{0, 2, 3}, result.O())
	}

	// copy full array
	{
		assert.Equal(t, []interface{}{}, NewInterSliceV().Copy().O())
		assert.Equal(t, []interface{}{}, NewInterSliceV().Copy(0, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV().Copy(0, 1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV().Copy(0, 5).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").Copy().O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").Copy(0, -1).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").Copy(0, 1).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3).Copy().O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3).Copy(0, -1).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSlice([]int{1, 2, 3}).Copy().O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSlice([]int{1, 2, 3}).Copy(0, -1).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").Copy().O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").Copy(0, 2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy().O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(0, -1).O())
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, []interface{}{"1"}, NewInterSliceV("1").Copy(0, 2).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false).Copy(-6, 6).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3).Copy(-6, 6).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").Copy(-6, 6).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(-6, 6).O())
	}

	// mutually exclusive
	{
		slice := NewInterSliceV(1, 2, 3, 4)
		assert.Equal(t, []interface{}{}, slice.Copy(2, -3).O())
		assert.Equal(t, []interface{}{}, slice.Copy(0, -5).O())
		assert.Equal(t, []interface{}{}, slice.Copy(4, -1).O())
		assert.Equal(t, []interface{}{}, slice.Copy(6, -1).O())
		assert.Equal(t, []interface{}{}, slice.Copy(3, 2).O())
	}

	// singles
	{
		slice := NewInterSliceV(1, 2, 3, 4)
		assert.Equal(t, []interface{}{4}, slice.Copy(-1, -1).O())
		assert.Equal(t, []interface{}{3}, slice.Copy(-2, -2).O())
		assert.Equal(t, []interface{}{2}, slice.Copy(-3, -3).O())
		assert.Equal(t, []interface{}{1}, slice.Copy(0, 0).O())
		assert.Equal(t, []interface{}{1}, slice.Copy(-4, -4).O())
		assert.Equal(t, []interface{}{2}, slice.Copy(1, 1).O())
		assert.Equal(t, []interface{}{2}, slice.Copy(1, -3).O())
		assert.Equal(t, []interface{}{3}, slice.Copy(2, 2).O())
		assert.Equal(t, []interface{}{3}, slice.Copy(2, -2).O())
		assert.Equal(t, []interface{}{4}, slice.Copy(3, 3).O())
		assert.Equal(t, []interface{}{4}, slice.Copy(3, -1).O())
	}

	// grab all but first
	{
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Copy(1, -1).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Copy(1, 2).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Copy(-2, -1).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Copy(-2, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Copy(1, -1).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Copy(1, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Copy(-2, -1).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Copy(-2, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Copy(1, -1).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Copy(1, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Copy(-2, -1).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Copy(-2, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(1, -1).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(1, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(-2, -1).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(-2, 2).O())
	}

	// grab all but last
	{
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Copy(0, -2).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Copy(-3, -2).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Copy(-3, 1).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Copy(0, 1).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Copy(0, -2).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Copy(-3, -2).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Copy(-3, 1).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Copy(0, 1).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Copy(0, -2).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Copy(-3, -2).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Copy(-3, 1).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Copy(0, 1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(0, -2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(-3, -2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(-3, 1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Copy(0, 1).O())
	}

	// grab middle
	{
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Copy(1, -2).O())
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Copy(-3, -2).O())
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Copy(-3, 2).O())
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Copy(1, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Copy(1, -2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Copy(-3, -2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Copy(-3, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Copy(1, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Copy(1, -2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Copy(-3, -2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Copy(-3, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Copy(1, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Copy(1, -2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Copy(-3, -2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Copy(-3, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Copy(1, 2).O())
	}

	// random
	{
		assert.Equal(t, []interface{}{"1"}, NewInterSliceV("1", "2", "3").Copy(0, -3).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Copy(1, 2).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").Copy(0, 2).O())
	}
}

// Count
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Count() {
	slice := NewInterSliceV(1, 2, 2)
	fmt.Println(slice.Count(2))
	// Output: 2
}

func TestInterSlice_Count(t *testing.T) {

	// empty
	var slice *InterSlice
	assert.Equal(t, 0, slice.Count(0))
	assert.Equal(t, 0, NewInterSliceV().Count(0))

	assert.Equal(t, 1, NewInterSliceV(2, 3).Count(2))
	assert.Equal(t, 2, NewInterSliceV(1, 2, 2).Count(2))
	assert.Equal(t, 4, NewInterSliceV(4, 4, 3, 4, 4).Count(4))
	assert.Equal(t, 3, NewInterSliceV(3, 2, 3, 3, 5).Count(3))
	assert.Equal(t, 1, NewInterSliceV(1, 2, 3).Count(3))
}

// CountW
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_CountW() {
	slice := NewInterSliceV(1, 2, 2)
	fmt.Println(slice.CountW(func(x O) bool {
		return ExB(x.(int) == 2)
	}))
	// Output: 2
}

func TestInterSlice_CountW(t *testing.T) {

	// empty
	var slice *InterSlice
	assert.Equal(t, 0, slice.CountW(func(x O) bool { return ExB(x.(int) > 0) }))
	assert.Equal(t, 0, NewInterSliceV().CountW(func(x O) bool { return ExB(x.(int) > 0) }))

	assert.Equal(t, 1, NewInterSliceV(2, 3).CountW(func(x O) bool { return ExB(x.(int) > 2) }))
	assert.Equal(t, 1, NewInterSliceV(1, 2).CountW(func(x O) bool { return ExB(x.(int) == 2) }))
	assert.Equal(t, 2, NewInterSliceV(1, 2, 3, 4, 5).CountW(func(x O) bool { return ExB(x.(int)%2 == 0) }))
	assert.Equal(t, 3, NewInterSliceV(1, 2, 3, 4, 5).CountW(func(x O) bool { return ExB(x.(int)%2 != 0) }))
	assert.Equal(t, 1, NewInterSliceV(1, 2, 3).CountW(func(x O) bool { return ExB(x.(int) == 4 || x.(int) == 3) }))
}

// Drop
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Drop() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Drop(1, 1).O())
	// Output: [1 3]
}

func TestInterSlice_Drop(t *testing.T) {
	// nil or empty
	{
		var slice *InterSlice
		assert.Equal(t, (*InterSlice)(nil), slice.Drop(0, 1))
	}

	// int
	{
		// invalid
		assert.Equal(t, []interface{}{1, 2, 3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(1).O())
		assert.Equal(t, []interface{}{1, 2, 3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(4, 4).O())

		// drop 1
		assert.Equal(t, []interface{}{2, 3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(0, 0).O())
		assert.Equal(t, []interface{}{1, 3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(1, 1).O())
		assert.Equal(t, []interface{}{1, 2, 4}, NewInterSliceV(1, 2, 3, 4).Drop(2, 2).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3, 4).Drop(3, 3).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3, 4).Drop(-1, -1).O())
		assert.Equal(t, []interface{}{1, 2, 4}, NewInterSliceV(1, 2, 3, 4).Drop(-2, -2).O())
		assert.Equal(t, []interface{}{1, 3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(-3, -3).O())
		assert.Equal(t, []interface{}{2, 3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(-4, -4).O())

		// drop 2
		assert.Equal(t, []interface{}{3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(0, 1).O())
		assert.Equal(t, []interface{}{1, 4}, NewInterSliceV(1, 2, 3, 4).Drop(1, 2).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3, 4).Drop(2, 3).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3, 4).Drop(-2, -1).O())
		assert.Equal(t, []interface{}{1, 4}, NewInterSliceV(1, 2, 3, 4).Drop(-3, -2).O())
		assert.Equal(t, []interface{}{3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(-4, -3).O())

		// drop 3
		assert.Equal(t, []interface{}{4}, NewInterSliceV(1, 2, 3, 4).Drop(0, 2).O())
		assert.Equal(t, []interface{}{1}, NewInterSliceV(1, 2, 3, 4).Drop(-3, -1).O())

		// drop everything and beyond
		assert.Equal(t, []interface{}{}, NewInterSliceV(1, 2, 3, 4).Drop().O())
		assert.Equal(t, []interface{}{}, NewInterSliceV(1, 2, 3, 4).Drop(0, 3).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV(1, 2, 3, 4).Drop(0, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV(1, 2, 3, 4).Drop(-4, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV(1, 2, 3, 4).Drop(-6, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV(1, 2, 3, 4).Drop(0, 10).O())

		// move index within bounds
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3, 4).Drop(3, 4).O())
		assert.Equal(t, []interface{}{2, 3, 4}, NewInterSliceV(1, 2, 3, 4).Drop(-5, 0).O())
	}

	// int
	{
		// invalid
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(4, 4).O())

		// drop {1}
		assert.Equal(t, []interface{}{Object{2}, Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 0).O())
		assert.Equal(t, []interface{}{Object{1}, Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(1, 1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(2, 2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(3, 3).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-1, -1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-2, -2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-3, -3).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-4, -4).O())

		// drop {2}
		assert.Equal(t, []interface{}{Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(1, 2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(2, 3).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-2, -1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-3, -2).O())
		assert.Equal(t, []interface{}{Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-4, -3).O())

		// drop {3}
		assert.Equal(t, []interface{}{Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 2).O())
		assert.Equal(t, []interface{}{Object{1}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-3, -1).O())

		// drop everything and beyond
		assert.Equal(t, []interface{}{}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop().O())
		assert.Equal(t, []interface{}{}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 3).O())
		assert.Equal(t, []interface{}{}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-4, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-6, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(0, 10).O())

		// move index within bounds
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(3, 4).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}, Object{4}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Drop(-5, 0).O())
	}
}

// DropAt
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_DropAt() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.DropAt(1).O())
	// Output: [1 3]
}

func TestInterSlice_DropAt(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		assert.Equal(t, (*InterSlice)(nil), slice.DropAt(0))
	}

	// drop all and more
	{
		slice := NewInterSliceV(0, 1, 2)
		assert.Equal(t, []interface{}{0, 1}, slice.DropAt(-1).O())
		assert.Equal(t, []interface{}{0}, slice.DropAt(-1).O())
		assert.Equal(t, []interface{}{}, slice.DropAt(-1).O())
		assert.Equal(t, []interface{}{}, slice.DropAt(-1).O())
	}

	// drop invalid
	assert.Equal(t, []interface{}{0, 1, 2}, NewInterSliceV(0, 1, 2).DropAt(3).O())
	assert.Equal(t, []interface{}{0, 1, 2}, NewInterSliceV(0, 1, 2).DropAt(-4).O())

	// drop last
	assert.Equal(t, []interface{}{0, 1}, NewInterSliceV(0, 1, 2).DropAt(2).O())
	assert.Equal(t, []interface{}{0, 1}, NewInterSliceV(0, 1, 2).DropAt(-1).O())

	// drop middle
	assert.Equal(t, []interface{}{0, 2}, NewInterSliceV(0, 1, 2).DropAt(1).O())
	assert.Equal(t, []interface{}{0, 2}, NewInterSliceV(0, 1, 2).DropAt(-2).O())

	// drop first
	assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(0, 1, 2).DropAt(0).O())
	assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(0, 1, 2).DropAt(-3).O())
}

// DropFirst
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_DropFirst() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.DropFirst().O())
	// Output: [2 3]
}

func TestInterSlice_DropFirst(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, (*InterSlice)(nil), nilSlice.DropFirst())
	}

	// bool
	{
		slice := NewInterSliceV(true, true, false)
		assert.Equal(t, []interface{}{true, false}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{false}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}

	// int
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{2, 3}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{3}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}

	// string
	{
		slice := NewInterSliceV("1", "2", "3")
		assert.Equal(t, []interface{}{"2", "3"}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{"3"}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}

	// custom
	{
		slice := NewInterSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{Object{3}}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}
}

// DropFirstN
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_DropFirstN() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.DropFirstN(2).O())
	// Output: [3]
}

func TestInterSlice_DropFirstN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, (*InterSlice)(nil), nilSlice.DropFirstN(1))
	}

	// drop none
	{
		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{1, 2, 3}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := NewInterSliceV(true, true, false)
			assert.Equal(t, []interface{}{true, false}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{2, 3}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := NewInterSliceV("1", "2", "3")
			assert.Equal(t, []interface{}{"2", "3"}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{Object{2}, Object{3}}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := NewInterSliceV(true, false, false)
			assert.Equal(t, []interface{}{false}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{3}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := NewInterSliceV("1", "2", "3")
			assert.Equal(t, []interface{}{"3"}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{Object{3}}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{}, slice.DropFirstN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{}, slice.DropFirstN(4).O())
			assert.Equal(t, 0, slice.Len())
		}
	}
}

// DropLast
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_DropLast() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.DropLast().O())
	// Output: [1 2]
}

func TestInterSlice_DropLast(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, (*InterSlice)(nil), nilSlice.DropLast())
	}

	// bool
	{
		slice := NewInterSliceV(true, true, false)
		assert.Equal(t, []interface{}{true, true}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{true}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}

	// int
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1, 2}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{1}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}

	// string
	{
		slice := NewInterSliceV("1", "2", "3")
		assert.Equal(t, []interface{}{"1", "2"}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{"1"}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}

	// custom
	{
		slice := NewInterSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []interface{}{Object{1}}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []interface{}{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
	}
}

// DropLastN
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_DropLastN() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.DropLastN(2).O())
	// Output: [1]
}

func TestInterSlice_DropLastN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, (*InterSlice)(nil), nilSlice.DropLastN(1))
	}

	// drop none
	{
		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{1, 2, 3}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := NewInterSliceV(true, true, false)
			assert.Equal(t, []interface{}{true, true}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{1, 2}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := NewInterSliceV("1", "2", "3")
			assert.Equal(t, []interface{}{"1", "2"}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{Object{1}, Object{2}}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := NewInterSliceV(true, false, false)
			assert.Equal(t, []interface{}{true}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{1}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := NewInterSliceV("1", "2", "3")
			assert.Equal(t, []interface{}{"1"}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{Object{1}}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := NewInterSliceV(1, 2, 3)
			assert.Equal(t, []interface{}{}, slice.DropLastN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := NewInterSlice([]Object{{1}, {2}, {3}})
			assert.Equal(t, []interface{}{}, slice.DropLastN(4).O())
			assert.Equal(t, 0, slice.Len())
		}
	}
}

// DropW
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_DropW() {
	slice := NewInterSliceV("1", "2", "3")
	fmt.Println(slice.DropW(func(x O) bool {
		return ExB(Obj(x).ToInt()%2 == 0)
	}).O())
	// Output: [1 3]
}

func TestInterSlice_DropW(t *testing.T) {

	// drop all odd values
	{
		slice := NewInterSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		slice.DropW(func(x O) bool {
			return ExB(Obj(x).ToInt()%2 != 0)
		})
		assert.Equal(t, NewInterSliceV("2", "4", "6", "8"), slice)
	}

	// drop all even values
	{
		slice := NewInterSliceV("1", "2", "3", "4", "5", "6", "7", "8", "9")
		slice.DropW(func(x O) bool {
			return ExB(Obj(x).ToInt()%2 == 0)
		})
		assert.Equal(t, NewInterSliceV("1", "3", "5", "7", "9"), slice)
	}
}

// Each
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Each() {
	NewInterSliceV(1, 2, 3).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 123
}

func TestInterSlice_Each(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		nilSlice.Each(func(x O) {})
	}

	// int
	{
		NewInterSliceV(1, 2, 3).Each(func(x O) {
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
		NewInterSliceV("1", "2", "3").Each(func(x O) {
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
		NewInterSlice([]Object{{1}, {2}, {3}}).Each(func(x O) {
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
func ExampleInterSlice_EachE() {
	NewInterSliceV(1, 2, 3).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 123
}

func TestInterSlice_EachE(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		nilSlice.EachE(func(x O) error {
			return nil
		})
	}

	// int
	{
		NewInterSliceV(1, 2, 3).EachE(func(x O) error {
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
		NewInterSliceV("1", "2", "3").EachE(func(x O) error {
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
		NewInterSlice([]Object{{1}, {2}, {3}}).EachE(func(x O) error {
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
func ExampleInterSlice_EachI() {
	NewInterSliceV(1, 2, 3).EachI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 0:11:22:3
}

func TestInterSlice_EachI(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		slice.EachI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachI(func(i int, x O) {
			results = append(results, x.(int))
		})
		assert.Equal(t, []int{1, 2, 3}, results)
	}
}

// EachIE
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_EachIE() {
	NewInterSliceV(1, 2, 3).EachIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 0:11:22:3
}

func TestInterSlice_EachIE(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		slice.EachIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachIE(func(i int, x O) error {
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{1, 2, 3}, results)
	}

	// Break early with error
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachIE(func(i int, x O) error {
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
func ExampleInterSlice_EachR() {
	NewInterSliceV(1, 2, 3).EachR(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 321
}

func TestInterSlice_EachR(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		slice.EachR(func(x O) {})
	}

	// Loop through
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachR(func(x O) {
			results = append(results, x.(int))
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}
}

// EachRE
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_EachRE() {
	NewInterSliceV(1, 2, 3).EachRE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 321
}

func TestInterSlice_EachRE(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		slice.EachRE(func(x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachRE(func(x O) error {
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}

	// Break early with error
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachRE(func(x O) error {
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
func ExampleInterSlice_EachRI() {
	NewInterSliceV(1, 2, 3).EachRI(func(i int, x O) {
		fmt.Printf("%v:%v", i, x)
	})
	// Output: 2:31:20:1
}

func TestInterSlice_EachRI(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		slice.EachRI(func(i int, x O) {})
	}

	// Loop through
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachRI(func(i int, x O) {
			results = append(results, x.(int))
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}
}

// EachRIE
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_EachRIE() {
	NewInterSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
		fmt.Printf("%v:%v", i, x)
		return nil
	})
	// Output: 2:31:20:1
}

func TestInterSlice_EachRIE(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		slice.EachRIE(func(i int, x O) error {
			return nil
		})
	}

	// Loop through
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
			results = append(results, x.(int))
			return nil
		})
		assert.Equal(t, []int{3, 2, 1}, results)
	}

	// Break early with error
	{
		results := []int{}
		NewInterSliceV(1, 2, 3).EachRIE(func(i int, x O) error {
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
func ExampleInterSlice_Empty() {
	fmt.Println(NewInterSliceV().Empty())
	// Output: true
}

func TestInterSlice_Empty(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, true, nilSlice.Empty())
	}

	assert.Equal(t, true, NewInterSliceV().Empty())
	assert.Equal(t, false, NewInterSliceV(1).Empty())
	assert.Equal(t, false, NewInterSliceV(1, 2, 3).Empty())
	assert.Equal(t, false, NewInterSlice(1).Empty())
	assert.Equal(t, false, NewInterSlice([]int{1, 2, 3}).Empty())
}

// First
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_First() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.First().O())
	// Output: 1
}

func TestInterSlice_First(t *testing.T) {
	// invalid
	{
		assert.Equal(t, Obj(nil), NewInterSliceV().First())
	}

	// bool
	{
		assert.Equal(t, &Object{true}, NewInterSliceV(true, false).First())
		assert.Equal(t, &Object{false}, NewInterSliceV(false, true).First())
	}

	// int
	{
		assert.Equal(t, Obj(2), NewInterSliceV(2, 3).First())
		assert.Equal(t, Obj(3), NewInterSliceV(3, 2).First())
		assert.Equal(t, Obj(1), NewInterSliceV(1, 3, 2).First())
	}

	// string
	{
		assert.Equal(t, &Object{"2"}, NewInterSliceV("2", "3").First())
		assert.Equal(t, &Object{"3"}, NewInterSliceV("3", "2").First())
		assert.Equal(t, &Object{"1"}, NewInterSliceV("1", "3", "2").First())
	}

	// custom
	{
		assert.Equal(t, &Object{Object{2}}, NewInterSlice([]Object{{2}, {3}}).First())
		assert.Equal(t, &Object{Object{3}}, NewInterSlice([]Object{{3}, {2}}).First())
		assert.Equal(t, &Object{Object{1}}, NewInterSlice([]Object{{1}, {3}, {2}}).First())
	}
}

// FirstN
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_FirstN() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.FirstN(2).O())
	// Output: [1 2]
}

func TestInterSlice_FirstN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, []interface{}{}, nilSlice.FirstN(1).O())
		slice := NewInterSliceV(0).Clear()
		assert.Equal(t, []interface{}{}, slice.FirstN(-1).O())
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewInterSliceV(1, 2, 3)
		result := original.FirstN(2).Set(0, 0)
		assert.Equal(t, []interface{}{0, 2, 3}, original.O())
		assert.Equal(t, []interface{}{0, 2}, result.O())
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, []interface{}{}, NewInterSliceV().FirstN(1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV().FirstN(10).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").FirstN(1).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").FirstN(10).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3).FirstN(10).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSlice([]int{1, 2, 3}).FirstN(10).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").FirstN(10).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).FirstN(10).O())
	}

	// grab a few diff
	{
		assert.Equal(t, []interface{}{true}, NewInterSliceV(true, false, true).FirstN(1).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).FirstN(2).O())
		assert.Equal(t, []interface{}{1}, NewInterSliceV(1, 2, 3).FirstN(1).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).FirstN(2).O())
		assert.Equal(t, []interface{}{"1"}, NewInterSliceV("1", "2", "3").FirstN(1).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").FirstN(2).O())
		assert.Equal(t, []interface{}{Object{1}}, NewInterSlice([]Object{{1}, {2}, {3}}).FirstN(1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).FirstN(2).O())
	}
}

// InterSlice
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_InterSlice() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.InterSlice())
	// Output: true
}

// Index
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Index() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Index(2))
	// Output: 1
}

func TestInterSlice_Index(t *testing.T) {

	// empty
	var slice *InterSlice
	assert.Equal(t, -1, slice.Index(2))
	assert.Equal(t, -1, NewInterSliceV().Index(1))

	assert.Equal(t, 0, NewInterSliceV(1, 2, 3).Index(1))
	assert.Equal(t, 1, NewInterSliceV(1, 2, 3).Index(2))
	assert.Equal(t, 2, NewInterSliceV(1, 2, 3).Index(3))
	assert.Equal(t, -1, NewInterSliceV(1, 2, 3).Index(4))
	assert.Equal(t, -1, NewInterSliceV(1, 2, 3).Index(5))
}

// Insert
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Insert() {
	slice := NewInterSliceV(1, 3)
	fmt.Println(slice.Insert(1, 2).O())
	// Output: [1 2 3]
}

func TestInterSlice_Insert(t *testing.T) {

	// int
	{
		// append
		{
			slice := NewInterSliceV()
			assert.Equal(t, []interface{}{0}, slice.Insert(-1, 0).O())
			assert.Equal(t, []interface{}{0, 1}, slice.Insert(-1, 1).O())
			assert.Equal(t, []interface{}{0, 1, 2}, slice.Insert(-1, 2).O())
		}

		// prepend
		{
			slice := NewInterSliceV()
			assert.Equal(t, []interface{}{2}, slice.Insert(0, 2).O())
			assert.Equal(t, []interface{}{1, 2}, slice.Insert(0, 1).O())
			assert.Equal(t, []interface{}{0, 1, 2}, slice.Insert(0, 0).O())
		}

		// middle pos
		{
			slice := NewInterSliceV(0, 5)
			assert.Equal(t, []interface{}{0, 1, 5}, slice.Insert(1, 1).O())
			assert.Equal(t, []interface{}{0, 1, 2, 5}, slice.Insert(2, 2).O())
			assert.Equal(t, []interface{}{0, 1, 2, 3, 5}, slice.Insert(3, 3).O())
			assert.Equal(t, []interface{}{0, 1, 2, 3, 4, 5}, slice.Insert(4, 4).O())
		}

		// middle neg
		{
			slice := NewInterSliceV(0, 5)
			assert.Equal(t, []interface{}{0, 1, 5}, slice.Insert(-2, 1).O())
			assert.Equal(t, []interface{}{0, 1, 2, 5}, slice.Insert(-2, 2).O())
			assert.Equal(t, []interface{}{0, 1, 2, 3, 5}, slice.Insert(-2, 3).O())
			assert.Equal(t, []interface{}{0, 1, 2, 3, 4, 5}, slice.Insert(-2, 4).O())
		}

		// error cases
		{
			var slice *InterSlice
			assert.False(t, slice.Insert(0, 0).Nil())
			assert.Equal(t, []interface{}{0}, slice.Insert(0, 0).O())
			assert.Equal(t, []interface{}{0, 1}, NewInterSliceV(0, 1).Insert(-10, 1).O())
			assert.Equal(t, []interface{}{0, 1}, NewInterSliceV(0, 1).Insert(10, 1).O())
			assert.Equal(t, []interface{}{0, 1}, NewInterSliceV(0, 1).Insert(2, 1).O())
			assert.Equal(t, []interface{}{0, 1}, NewInterSliceV(0, 1).Insert(-3, 1).O())
		}
	}

	// custom
	{
		// append
		{
			slice := NewInterSliceV()
			assert.Equal(t, []interface{}{Object{0}}, slice.Insert(-1, Object{0}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, slice.Insert(-1, Object{1}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}}, slice.Insert(-1, Object{2}).O())
		}

		// prepend
		{
			slice := NewInterSliceV()
			assert.Equal(t, []interface{}{Object{2}}, slice.Insert(0, Object{2}).O())
			assert.Equal(t, []interface{}{Object{1}, Object{2}}, slice.Insert(0, Object{1}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}}, slice.Insert(0, Object{0}).O())
		}

		// middle pos
		{
			slice := NewInterSlice([]Object{{0}, {5}})
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{5}}, slice.Insert(1, Object{1}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}, Object{5}}, slice.Insert(2, Object{2}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}, Object{3}, Object{5}}, slice.Insert(3, Object{3}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}, Object{3}, Object{4}, Object{5}}, slice.Insert(4, Object{4}).O())
		}

		// middle neg
		{
			slice := NewInterSlice([]Object{{0}, {5}})
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{5}}, slice.Insert(-2, Object{1}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}, Object{5}}, slice.Insert(-2, Object{2}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}, Object{3}, Object{5}}, slice.Insert(-2, Object{3}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}, Object{3}, Object{4}, Object{5}}, slice.Insert(-2, Object{4}).O())
		}

		// error cases
		{
			var slice *InterSlice
			assert.False(t, slice.Insert(0, Object{0}).Nil())
			assert.Equal(t, []interface{}{Object{0}}, slice.Insert(0, Object{0}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, NewInterSlice([]Object{{0}, {1}}).Insert(-10, 1).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, NewInterSlice([]Object{{0}, {1}}).Insert(10, 1).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, NewInterSlice([]Object{{0}, {1}}).Insert(2, 1).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, NewInterSlice([]Object{{0}, {1}}).Insert(-3, 1).O())
		}
	}
}

func TestInterSlice_Insert_MixedType(t *testing.T) {
	slice := NewInterSliceV(1)
	slice.Insert(0, "2")
	assert.Equal(t, []interface{}{"2", 1}, slice.G())
}

// Join
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Join() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Join())
	// Output: 1,2,3
}

func TestInterSlice_Join(t *testing.T) {
	// nil
	{
		var slice *InterSlice
		assert.Equal(t, Obj(""), slice.Join())
	}

	// empty
	{
		assert.Equal(t, Obj(""), NewInterSliceV().Join())
	}

	assert.Equal(t, "1,2,3", NewInterSliceV(1, 2, 3).Join().O())
	assert.Equal(t, "1.2.3", NewInterSliceV(1, 2, 3).Join(".").O())
}

// Last
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Last() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Last())
	// Output: 3
}

func TestInterSlice_Last(t *testing.T) {

	// invalid
	{
		assert.Equal(t, Obj(nil), NewInterSliceV().Last())
	}

	// int
	{
		assert.Equal(t, Obj(3), NewInterSliceV(2, 3).Last())
		assert.Equal(t, Obj(2), NewInterSliceV(3, 2).Last())
		assert.Equal(t, Obj(2), NewInterSliceV(1, 3, 2).Last())
	}

	// object
	{
		assert.Equal(t, &Object{Object{3}}, NewInterSlice([]Object{{2}, {3}}).Last())
		assert.Equal(t, &Object{Object{2}}, NewInterSlice([]Object{{3}, {2}}).Last())
		assert.Equal(t, &Object{Object{2}}, NewInterSlice([]Object{{1}, {3}, {2}}).Last())
	}
}

// LastN
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_LastN() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.LastN(2).O())
	// Output: [2 3]
}

func TestInterSlice_LastN(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, []interface{}{}, nilSlice.LastN(1).O())
		slice := NewInterSliceV(0).Clear()
		assert.Equal(t, []interface{}{}, slice.LastN(-1).O())
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewInterSliceV(1, 2, 3)
		result := original.LastN(2).Set(0, 0)
		assert.Equal(t, []interface{}{1, 0, 3}, original.O())
		assert.Equal(t, []interface{}{0, 3}, result.O())
	}

	// slice full array includeing out of bounds
	{
		assert.Equal(t, []interface{}{}, NewInterSliceV().LastN(1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV().LastN(10).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").LastN(1).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").LastN(10).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3).LastN(10).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSlice([]int{1, 2, 3}).LastN(10).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").LastN(10).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).LastN(10).O())
	}

	// grab a few diff
	{
		assert.Equal(t, []interface{}{false}, NewInterSliceV(true, true, false).LastN(1).O())
		assert.Equal(t, []interface{}{false}, NewInterSliceV(true, true, false).LastN(-1).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).LastN(2).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).LastN(-2).O())
		assert.Equal(t, []interface{}{3}, NewInterSliceV(1, 2, 3).LastN(1).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).LastN(2).O())
		assert.Equal(t, []interface{}{"3"}, NewInterSliceV("1", "2", "3").LastN(1).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").LastN(2).O())
		assert.Equal(t, []interface{}{Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).LastN(1).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).LastN(2).O())
	}
}

// Len
//--------------------------------------------------------------------------------------------------
func TestInterSlice_Len(t *testing.T) {
	assert.Equal(t, 0, NewInterSliceV().Len())
	assert.Equal(t, 1, NewInterSliceV().Append("2").Len())
}

// Less
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Less() {
	slice := NewInterSliceV(2, 3, 1)
	fmt.Println(slice.Sort().O())
	// Output: [1 2 3]
}

func TestInterSlice_Less(t *testing.T) {

	// invalid cases
	{
		var slice *InterSlice
		assert.False(t, slice.Less(0, 0))
		slice = NewInterSliceV()
		assert.False(t, slice.Less(0, 0))
		assert.False(t, slice.Less(1, 2))
		assert.False(t, slice.Less(-1, 2))
		assert.False(t, slice.Less(1, -2))
	}

	// // bool
	// {
	// 	assert.Equal(t, false, NewInterSliceV(true, false, true).Less(0, 1))
	// 	assert.Equal(t, true, NewInterSliceV(true, false, true).Less(1, 0))
	// }

	// int
	{
		assert.Equal(t, true, NewInterSliceV(0, 1, 2).Less(0, 1))
		assert.Equal(t, false, NewInterSliceV(0, 1, 2).Less(1, 0))
		assert.Equal(t, true, NewInterSliceV(0, 1, 2).Less(1, 2))
	}

	// string
	{
		assert.Equal(t, true, NewInterSliceV("0", "1", "2").Less(0, 1))
		assert.Equal(t, false, NewInterSliceV("0", "1", "2").Less(1, 0))
		assert.Equal(t, true, NewInterSliceV("0", "1", "2").Less(1, 2))
	}

	// // custom
	// {
	// 	assert.Equal(t, true, NewInterSlice([]Object{{0}, {1}, {2}}).Less(0, 1))
	// 	assert.Equal(t, false, NewInterSlice([]Object{{0}, {1}, {2}}).Less(1, 0))
	// 	assert.Equal(t, true, NewInterSlice([]Object{{0}, {1}, {2}}).Less(1, 2))
	// }
}

// Nil
//--------------------------------------------------------------------------------------------------
func TestInterSlice_Nil(t *testing.T) {
	assert.False(t, NewInterSliceV().Nil())
	var q *InterSlice
	assert.True(t, q.Nil())
	assert.False(t, NewInterSliceV().Append("2").Nil())
}

// O
//--------------------------------------------------------------------------------------------------
func TestInterSlice_O(t *testing.T) {
	assert.Equal(t, []interface{}{}, NewInterSliceV().O())
	assert.Len(t, NewInterSliceV().Append("2").O(), 1)
}

// Pair
//--------------------------------------------------------------------------------------------------

func ExampleInterSlice_Pair() {
	slice := NewInterSliceV(1, 2)
	first, second := slice.Pair()
	fmt.Println(first.O(), second.O())
	// Output: 1 2
}

func TestInterSlice_Pair(t *testing.T) {

	// nil
	{
		first, second := (*InterSlice)(nil).Pair()
		assert.Equal(t, Obj(nil), first)
		assert.Equal(t, Obj(nil), second)
	}

	// int
	{
		// two values
		{
			first, second := NewInterSliceV(1, 2).Pair()
			assert.Equal(t, Obj(1), first)
			assert.Equal(t, Obj(2), second)
		}

		// one value
		{
			first, second := NewInterSliceV(1).Pair()
			assert.Equal(t, Obj(1), first)
			assert.Equal(t, Obj(nil), second)
		}

		// no values
		{
			first, second := NewInterSliceV().Pair()
			assert.Equal(t, Obj(nil), first)
			assert.Equal(t, Obj(nil), second)
		}
	}

	// custom
	{
		// two values
		{
			first, second := NewInterSlice([]Object{{1}, {2}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, &Object{Object{2}}, second)
		}

		// one value
		{
			first, second := NewInterSlice([]Object{{1}}).Pair()
			assert.Equal(t, &Object{Object{1}}, first)
			assert.Equal(t, Obj(nil), second)
		}

		// no values
		{
			first, second := NewInterSliceV().Pair()
			assert.Equal(t, Obj(nil), first)
			assert.Equal(t, Obj(nil), second)
		}
	}
}

// Pop
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Pop() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Pop())
	// Output: 3
}

func TestInterSlice_Pop(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		assert.Equal(t, Obj(nil), slice.Pop())
	}

	// take all one at a time
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, Obj(3), slice.Pop())
		assert.Equal(t, []interface{}{1, 2}, slice.O())
		assert.Equal(t, Obj(2), slice.Pop())
		assert.Equal(t, []interface{}{1}, slice.O())
		assert.Equal(t, Obj(1), slice.Pop())
		assert.Equal(t, []interface{}{}, slice.O())
		assert.Equal(t, Obj(nil), slice.Pop())
		assert.Equal(t, []interface{}{}, slice.O())
	}
}

// PopN
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_PopN() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.PopN(2))
	// Output: [2 3]
}

func TestInterSlice_PopN(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		assert.Equal(t, NewInterSliceV(), slice.PopN(1))
	}

	// take none
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{}, slice.PopN(0).O())
		assert.Equal(t, []interface{}{1, 2, 3}, slice.O())
	}

	// take 1
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{3}, slice.PopN(1).O())
		assert.Equal(t, []interface{}{1, 2}, slice.O())
	}

	// take 2
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{2, 3}, slice.PopN(2).O())
		assert.Equal(t, []interface{}{1}, slice.O())
	}

	// take 3
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1, 2, 3}, slice.PopN(3).O())
		assert.Equal(t, []interface{}{}, slice.O())
	}

	// take beyond
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1, 2, 3}, slice.PopN(4).O())
		assert.Equal(t, []interface{}{}, slice.O())
	}
}

// Prepend
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Prepend() {
	slice := NewInterSliceV(2, 3)
	fmt.Println(slice.Prepend(1).O())
	// Output: [1 2 3]
}

func TestInterSlice_Prepend(t *testing.T) {

	// int
	{
		// happy path
		{
			slice := NewInterSliceV()
			assert.Equal(t, []interface{}{2}, slice.Prepend(2).O())
			assert.Equal(t, []interface{}{1, 2}, slice.Prepend(1).O())
			assert.Equal(t, []interface{}{0, 1, 2}, slice.Prepend(0).O())
		}

		// error cases
		{
			var slice *InterSlice
			assert.False(t, slice.Prepend(0).Nil())
			assert.Equal(t, []interface{}{0}, slice.Prepend(0).O())
		}
	}

	// custom
	{
		// prepend
		{
			slice := NewInterSliceV()
			assert.Equal(t, []interface{}{Object{2}}, slice.Prepend(Object{2}).O())
			assert.Equal(t, []interface{}{Object{1}, Object{2}}, slice.Prepend(Object{1}).O())
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}}, slice.Prepend(Object{0}).O())
		}

		// error cases
		{
			var slice *InterSlice
			assert.False(t, slice.Prepend(Object{0}).Nil())
			assert.Equal(t, []interface{}{Object{0}}, slice.Prepend(Object{0}).O())
		}
	}
}

// Reverse
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Reverse() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Reverse())
	// Output: [3 2 1]
}

func TestInterSlice_Reverse(t *testing.T) {

	// nil
	{
		var slice *InterSlice
		assert.Equal(t, NewInterSliceV(), slice.Reverse())
	}

	// empty
	{
		assert.Equal(t, []interface{}{}, NewInterSliceV().Reverse().O())
	}

	// pos
	{
		slice := NewInterSliceV(3, 2, 1)
		reversed := slice.Reverse()
		assert.Equal(t, []interface{}{3, 2, 1, 4}, slice.Append(4).O())
		assert.Equal(t, []interface{}{1, 2, 3}, reversed.O())
	}

	// neg
	{
		slice := NewInterSliceV(2, 3, -2, -3)
		reversed := slice.Reverse()
		assert.Equal(t, []interface{}{2, 3, -2, -3, 4}, slice.Append(4).O())
		assert.Equal(t, []interface{}{-3, -2, 3, 2}, reversed.O())
	}
}

// ReverseM
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_ReverseM() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.ReverseM())
	// Output: [3 2 1]
}

func TestInterSlice_ReverseM(t *testing.T) {

	// nil
	{
		var slice *InterSlice
		assert.Equal(t, []interface{}{}, slice.ReverseM().O())
	}

	// empty
	{
		assert.Equal(t, []interface{}{}, NewInterSliceV().ReverseM().O())
	}

	// pos
	{
		slice := NewInterSliceV(3, 2, 1)
		reversed := slice.ReverseM()
		assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.Append(4).O())
		assert.Equal(t, []interface{}{1, 2, 3, 4}, reversed.O())
	}

	// neg
	{
		slice := NewInterSliceV(2, 3, -2, -3)
		reversed := slice.ReverseM()
		assert.Equal(t, []interface{}{-3, -2, 3, 2, 4}, slice.Append(4).O())
		assert.Equal(t, []interface{}{-3, -2, 3, 2, 4}, reversed.O())
	}
}

// Select
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Select() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Select(func(x O) bool {
		return ExB(x.(int) == 2 || x.(int) == 3)
	}))
	// Output: [2 3]
}

func TestInterSlice_Select(t *testing.T) {

	// Select all odd values
	{
		slice := NewInterSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.Select(func(x O) bool {
			return ExB(x.(int)%2 != 0)
		})
		slice.DropFirst()
		assert.Equal(t, []interface{}{2, 3, 4, 5, 6, 7, 8, 9}, slice.O())
		assert.Equal(t, []interface{}{1, 3, 5, 7, 9}, new.O())
	}

	// Select all even values
	{
		slice := NewInterSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.Select(func(x O) bool {
			return ExB(x.(int)%2 == 0)
		})
		slice.DropAt(1)
		assert.Equal(t, []interface{}{1, 3, 4, 5, 6, 7, 8, 9}, slice.O())
		assert.Equal(t, []interface{}{2, 4, 6, 8}, new.O())
	}
}

// Set
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Set() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Set(0, 0).O())
	// Output: [0 2 3]
}

func TestInterSlice_Set(t *testing.T) {
	// bool
	{
		assert.Equal(t, []interface{}{false, true, true}, NewInterSliceV(true, true, true).Set(0, false).O())
		assert.Equal(t, []interface{}{true, false, true}, NewInterSliceV(true, true, true).Set(1, false).O())
		assert.Equal(t, []interface{}{true, true, false}, NewInterSliceV(true, true, true).Set(2, false).O())
		assert.Equal(t, []interface{}{false, true, true}, NewInterSliceV(true, true, true).Set(-3, false).O())
		assert.Equal(t, []interface{}{true, false, true}, NewInterSliceV(true, true, true).Set(-2, false).O())
		assert.Equal(t, []interface{}{true, true, false}, NewInterSliceV(true, true, true).Set(-1, false).O())
	}

	// int
	{
		assert.Equal(t, []interface{}{0, 2, 3}, NewInterSliceV(1, 2, 3).Set(0, 0).O())
		assert.Equal(t, []interface{}{1, 0, 3}, NewInterSliceV(1, 2, 3).Set(1, 0).O())
		assert.Equal(t, []interface{}{1, 2, 0}, NewInterSliceV(1, 2, 3).Set(2, 0).O())
		assert.Equal(t, []interface{}{0, 2, 3}, NewInterSliceV(1, 2, 3).Set(-3, 0).O())
		assert.Equal(t, []interface{}{1, 0, 3}, NewInterSliceV(1, 2, 3).Set(-2, 0).O())
		assert.Equal(t, []interface{}{1, 2, 0}, NewInterSliceV(1, 2, 3).Set(-1, 0).O())
	}

	// string
	{
		assert.Equal(t, []interface{}{"0", "2", "3"}, NewInterSliceV("1", "2", "3").Set(0, "0").O())
		assert.Equal(t, []interface{}{"1", "0", "3"}, NewInterSliceV("1", "2", "3").Set(1, "0").O())
		assert.Equal(t, []interface{}{"1", "2", "0"}, NewInterSliceV("1", "2", "3").Set(2, "0").O())
		assert.Equal(t, []interface{}{"0", "2", "3"}, NewInterSliceV("1", "2", "3").Set(-3, "0").O())
		assert.Equal(t, []interface{}{"1", "0", "3"}, NewInterSliceV("1", "2", "3").Set(-2, "0").O())
		assert.Equal(t, []interface{}{"1", "2", "0"}, NewInterSliceV("1", "2", "3").Set(-1, "0").O())
	}

	// custom
	{
		assert.Equal(t, []interface{}{Object{0}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Set(0, Object{0}).O())
		assert.Equal(t, []interface{}{Object{1}, Object{0}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Set(1, Object{0}).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{0}}, NewInterSlice([]Object{{1}, {2}, {3}}).Set(2, Object{0}).O())
		assert.Equal(t, []interface{}{Object{0}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Set(-3, Object{0}).O())
		assert.Equal(t, []interface{}{Object{1}, Object{0}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Set(-2, Object{0}).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{0}}, NewInterSlice([]Object{{1}, {2}, {3}}).Set(-1, Object{0}).O())
	}
}

// SetE
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_SetE() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.SetE(0, 0))
	// Output: [0 2 3] <nil>
}

func TestInterSlice_SetE(t *testing.T) {

	// Error cases
	{
		// out of bounds
		slice, err := NewInterSliceV("1", "2", "3").SetE(5, "0")
		assert.Equal(t, []interface{}{"1", "2", "3"}, slice.O())
		assert.Equal(t, "slice assignment is out of bounds", err.Error())

		// mixed types
		slice, err = NewInterSliceV("1", "2", "3").SetE(0, 1)
		assert.Equal(t, []interface{}{1, "2", "3"}, slice.O())
	}

	// string
	{
		slice, err := NewInterSliceV("1", "2", "3").SetE(0, "0")
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{"0", "2", "3"}, slice.O())

		slice, err = NewInterSliceV("1", "2", "3").SetE(1, "0")
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{"1", "0", "3"}, slice.O())

		slice, err = NewInterSliceV("1", "2", "3").SetE(2, "0")
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{"1", "2", "0"}, slice.O())

		slice, err = NewInterSliceV("1", "2", "3").SetE(-3, "0")
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{"0", "2", "3"}, slice.O())

		slice, err = NewInterSliceV("1", "2", "3").SetE(-2, "0")
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{"1", "0", "3"}, slice.O())

		slice, err = NewInterSliceV("1", "2", "3").SetE(-1, "0")
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{"1", "2", "0"}, slice.O())
	}

	// custom
	{
		slice, err := NewInterSlice([]Object{{1}, {2}, {3}}).SetE(0, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{Object{0}, Object{2}, Object{3}}, slice.O())

		slice, err = NewInterSlice([]Object{{1}, {2}, {3}}).SetE(1, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{Object{1}, Object{0}, Object{3}}, slice.O())

		slice, err = NewInterSlice([]Object{{1}, {2}, {3}}).SetE(2, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{0}}, slice.O())

		slice, err = NewInterSlice([]Object{{1}, {2}, {3}}).SetE(-3, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{Object{0}, Object{2}, Object{3}}, slice.O())

		slice, err = NewInterSlice([]Object{{1}, {2}, {3}}).SetE(-2, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{Object{1}, Object{0}, Object{3}}, slice.O())

		slice, err = NewInterSlice([]Object{{1}, {2}, {3}}).SetE(-1, Object{0})
		assert.Nil(t, err)
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{0}}, slice.O())
	}
}

// Shift
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Shift() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Shift().O())
	// Output: 1
}

func TestInterSlice_Shift(t *testing.T) {

	// nil or empty
	{
		var slice *IntSlice
		assert.Equal(t, Obj(nil), slice.Shift())
	}

	// take all and beyond
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, Obj(1), slice.Shift())
		assert.Equal(t, []interface{}{2, 3}, slice.O())
		assert.Equal(t, Obj(2), slice.Shift())
		assert.Equal(t, []interface{}{3}, slice.O())
		assert.Equal(t, Obj(3), slice.Shift())
		assert.Equal(t, []interface{}{}, slice.O())
		assert.Equal(t, Obj(nil), slice.Shift())
		assert.Equal(t, []interface{}{}, slice.O())
	}

	// generic: take all and beyond
	{
		slice := NewInterSlice([]Object{{1}, {2}, {3}})
		assert.Equal(t, &Object{Object{1}}, slice.Shift())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, slice.O())
		assert.Equal(t, &Object{Object{2}}, slice.Shift())
		assert.Equal(t, []interface{}{Object{3}}, slice.O())
		assert.Equal(t, &Object{Object{3}}, slice.Shift())
		assert.Equal(t, []interface{}{}, slice.O())
		assert.Equal(t, Obj(nil), slice.Shift())
		assert.Equal(t, []interface{}{}, slice.O())
	}
}

// ShiftN
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_ShiftN() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.ShiftN(2).O())
	// Output: [1 2]
}

func TestInterSlice_ShiftN(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		assert.Equal(t, NewInterSliceV(), slice.ShiftN(1))
	}

	// negative value
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1}, slice.ShiftN(-1).O())
		assert.Equal(t, []interface{}{2, 3}, slice.O())
	}

	// take none
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{}, slice.ShiftN(0).O())
		assert.Equal(t, []interface{}{1, 2, 3}, slice.O())
	}

	// take 1
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1}, slice.ShiftN(1).O())
		assert.Equal(t, []interface{}{2, 3}, slice.O())
	}

	// take 2
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1, 2}, slice.ShiftN(2).O())
		assert.Equal(t, []interface{}{3}, slice.O())
	}

	// take 3
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1, 2, 3}, slice.ShiftN(3).O())
		assert.Equal(t, []interface{}{}, slice.O())
	}

	// take beyond
	{
		slice := NewInterSliceV(1, 2, 3)
		assert.Equal(t, []interface{}{1, 2, 3}, slice.ShiftN(4).O())
		assert.Equal(t, []interface{}{}, slice.O())
	}
}

// Single
//--------------------------------------------------------------------------------------------------

func ExampleInterSlice_Single() {
	slice := NewInterSliceV(1)
	fmt.Println(slice.Single())
	// Output: true
}

func TestInterSlice_Single(t *testing.T) {

	// int
	{
		assert.Equal(t, false, NewInterSliceV().Single())
		assert.Equal(t, true, NewInterSliceV(1).Single())
		assert.Equal(t, false, NewInterSliceV(1, 2).Single())
	}

	// custom
	{
		assert.Equal(t, false, NewInterSliceV().Single())
		assert.Equal(t, true, NewInterSliceV(Object{1}).Single())
		assert.Equal(t, false, NewInterSliceV(Object{1}, Object{2}).Single())
	}
}

// Slice
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Slice() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Slice(1, -1).O())
	// Output: [2 3]
}

func TestInterSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var nilSlice *InterSlice
		assert.Equal(t, NewInterSliceV(), nilSlice.Slice(0, -1))
		slice := NewInterSliceV(0).Clear()
		assert.Equal(t, []interface{}{}, slice.Slice(0, -1).O())
	}

	// Test that the original is modified when the slice is modified
	{
		original := NewInterSliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, []interface{}{0, 2, 3}, original.O())
		assert.Equal(t, []interface{}{0, 2, 3}, result.O())
	}

	// slice full array
	{
		assert.Equal(t, []interface{}{}, NewInterSliceV().Slice(0, -1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV().Slice(0, 1).O())
		assert.Equal(t, []interface{}{}, NewInterSliceV().Slice(0, 5).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").Slice(0, -1).O())
		assert.Equal(t, []interface{}{""}, NewInterSliceV("").Slice(0, 1).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3).Slice(0, -1).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSlice([]int{1, 2, 3}).Slice(0, -1).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").Slice(0, 2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(0, -1).O())
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, []interface{}{"1"}, NewInterSliceV("1").Slice(0, 2).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false).Slice(-6, 6).O())
		assert.Equal(t, []interface{}{1, 2, 3}, NewInterSliceV(1, 2, 3).Slice(-6, 6).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").Slice(-6, 6).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(-6, 6).O())
	}

	// mutually exclusive
	{
		slice := NewInterSliceV(1, 2, 3, 4)
		assert.Equal(t, []interface{}{}, slice.Slice(2, -3).O())
		assert.Equal(t, []interface{}{}, slice.Slice(0, -5).O())
		assert.Equal(t, []interface{}{}, slice.Slice(4, -1).O())
		assert.Equal(t, []interface{}{}, slice.Slice(6, -1).O())
		assert.Equal(t, []interface{}{}, slice.Slice(3, 2).O())
	}

	// singles
	{
		slice := NewInterSliceV(1, 2, 3, 4)
		assert.Equal(t, []interface{}{4}, slice.Slice(-1, -1).O())
		assert.Equal(t, []interface{}{3}, slice.Slice(-2, -2).O())
		assert.Equal(t, []interface{}{2}, slice.Slice(-3, -3).O())
		assert.Equal(t, []interface{}{1}, slice.Slice(0, 0).O())
		assert.Equal(t, []interface{}{1}, slice.Slice(-4, -4).O())
		assert.Equal(t, []interface{}{2}, slice.Slice(1, 1).O())
		assert.Equal(t, []interface{}{2}, slice.Slice(1, -3).O())
		assert.Equal(t, []interface{}{3}, slice.Slice(2, 2).O())
		assert.Equal(t, []interface{}{3}, slice.Slice(2, -2).O())
		assert.Equal(t, []interface{}{4}, slice.Slice(3, 3).O())
		assert.Equal(t, []interface{}{4}, slice.Slice(3, -1).O())
	}

	// grab all but first
	{
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Slice(1, -1).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Slice(1, 2).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Slice(-2, -1).O())
		assert.Equal(t, []interface{}{false, true}, NewInterSliceV(true, false, true).Slice(-2, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Slice(1, -1).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Slice(1, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Slice(-2, -1).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3).Slice(-2, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Slice(1, -1).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Slice(1, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Slice(-2, -1).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Slice(-2, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(1, -1).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(1, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(-2, -1).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(-2, 2).O())
	}

	// grab all but last
	{
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Slice(0, -2).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Slice(-3, -2).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Slice(-3, 1).O())
		assert.Equal(t, []interface{}{true, false}, NewInterSliceV(true, false, true).Slice(0, 1).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Slice(0, -2).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Slice(-3, -2).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Slice(-3, 1).O())
		assert.Equal(t, []interface{}{1, 2}, NewInterSliceV(1, 2, 3).Slice(0, 1).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Slice(0, -2).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Slice(-3, -2).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Slice(-3, 1).O())
		assert.Equal(t, []interface{}{"1", "2"}, NewInterSliceV("1", "2", "3").Slice(0, 1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(0, -2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(-3, -2).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(-3, 1).O())
		assert.Equal(t, []interface{}{Object{1}, Object{2}}, NewInterSlice([]Object{{1}, {2}, {3}}).Slice(0, 1).O())
	}

	// grab middle
	{
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Slice(1, -2).O())
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Slice(-3, -2).O())
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Slice(-3, 2).O())
		assert.Equal(t, []interface{}{true, true}, NewInterSliceV(false, true, true, false).Slice(1, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Slice(1, -2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Slice(-3, -2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Slice(-3, 2).O())
		assert.Equal(t, []interface{}{2, 3}, NewInterSliceV(1, 2, 3, 4).Slice(1, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Slice(1, -2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Slice(-3, -2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Slice(-3, 2).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3", "4").Slice(1, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, -2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, -2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Slice(-3, 2).O())
		assert.Equal(t, []interface{}{Object{2}, Object{3}}, NewInterSlice([]Object{{1}, {2}, {3}, {4}}).Slice(1, 2).O())
	}

	// random
	{
		assert.Equal(t, []interface{}{"1"}, NewInterSliceV("1", "2", "3").Slice(0, -3).O())
		assert.Equal(t, []interface{}{"2", "3"}, NewInterSliceV("1", "2", "3").Slice(1, 2).O())
		assert.Equal(t, []interface{}{"1", "2", "3"}, NewInterSliceV("1", "2", "3").Slice(0, 2).O())
	}
}

// String
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_String() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice)
	// Output: [1 2 3]
}

func TestInterSlice_String(t *testing.T) {

	assert.Equal(t, "[]", (*InterSlice)(nil).String())
	assert.Equal(t, "[]", NewInterSliceV().String())
	assert.Equal(t, "[5 4 3 2 1]", NewInterSliceV(5, 4, 3, 2, 1).String())
	assert.Equal(t, "[5 4 3 2 1 6]", NewInterSliceV(5, 4, 3, 2, 1).Append(6).String())
}

// Sort
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Sort() {
	slice := NewInterSliceV(2, 3, 1)
	fmt.Println(slice.Sort())
	// Output: [1 2 3]
}

func TestInterSlice_Sort(t *testing.T) {

	// empty
	assert.Equal(t, NewInterSliceV(), NewInterSliceV().Sort())

	// pos
	{
		slice := NewInterSliceV(5, 3, 2, 4, 1)
		sorted := slice.Sort()
		assert.Equal(t, []interface{}{1, 2, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{5, 3, 2, 4, 1}, slice.O())
	}

	// neg
	{
		slice := NewInterSliceV(5, 3, -2, 4, -1)
		sorted := slice.Sort()
		assert.Equal(t, []interface{}{-2, -1, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{5, 3, -2, 4, -1}, slice.O())
	}
}

// SortM
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_SortM() {
	slice := NewInterSliceV(2, 3, 1)
	fmt.Println(slice.SortM())
	// Output: [1 2 3]
}

func TestInterSlice_SortM(t *testing.T) {

	// empty
	assert.Equal(t, NewInterSliceV(), NewInterSliceV().SortM())

	// pos
	{
		slice := NewInterSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortM()
		assert.Equal(t, []interface{}{1, 2, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{1, 2, 3, 4, 5, 6}, slice.O())
	}

	// neg
	{
		slice := NewInterSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortM()
		assert.Equal(t, []interface{}{-2, -1, 3, 4, 5, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{-2, -1, 3, 4, 5, 6}, slice.O())
	}
}

// SortReverse
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_SortReverse() {
	slice := NewInterSliceV(2, 3, 1)
	fmt.Println(slice.SortReverse())
	// Output: [3 2 1]
}

func TestInterSlice_SortReverse(t *testing.T) {

	// empty
	assert.Equal(t, NewInterSliceV(), NewInterSliceV().SortReverse())

	// pos
	{
		slice := NewInterSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverse()
		assert.Equal(t, []interface{}{5, 4, 3, 2, 1, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{5, 3, 2, 4, 1}, slice.O())
	}

	// neg
	{
		slice := NewInterSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverse()
		assert.Equal(t, []interface{}{5, 4, 3, -1, -2, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{5, 3, -2, 4, -1}, slice.O())
	}
}

// SortReverseM
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_SortReverseM() {
	slice := NewInterSliceV(2, 3, 1)
	fmt.Println(slice.SortReverseM())
	// Output: [3 2 1]
}

func TestInterSlice_SortReverseM(t *testing.T) {

	// empty
	assert.Equal(t, NewInterSliceV(), NewInterSliceV().SortReverse())

	// pos
	{
		slice := NewInterSliceV(5, 3, 2, 4, 1)
		sorted := slice.SortReverseM()
		assert.Equal(t, []interface{}{5, 4, 3, 2, 1, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{5, 4, 3, 2, 1, 6}, slice.O())
	}

	// neg
	{
		slice := NewInterSliceV(5, 3, -2, 4, -1)
		sorted := slice.SortReverseM()
		assert.Equal(t, []interface{}{5, 4, 3, -1, -2, 6}, sorted.Append(6).O())
		assert.Equal(t, []interface{}{5, 4, 3, -1, -2, 6}, slice.O())
	}
}

// Swap
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Swap() {
	slice := NewInterSliceV(2, 3, 1)
	slice.Swap(0, 2)
	slice.Swap(1, 2)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestInterSlice_Swap(t *testing.T) {

	// invalid cases
	{
		var slice *InterSlice
		slice.Swap(0, 0)
		assert.Equal(t, (*InterSlice)(nil), slice)

		slice = NewInterSliceV()
		slice.Swap(0, 0)
		assert.Equal(t, NewInterSliceV(), slice)

		slice.Swap(1, 2)
		assert.Equal(t, NewInterSliceV(), slice)

		slice.Swap(-1, 2)
		assert.Equal(t, NewInterSliceV(), slice)

		slice.Swap(1, -2)
		assert.Equal(t, NewInterSliceV(), slice)
	}

	// bool
	{
		slice := NewInterSliceV(true, false, true)
		slice.Swap(0, 1)
		assert.Equal(t, []interface{}{false, true, true}, slice.O())
	}

	// int
	{
		slice := NewInterSliceV(0, 1, 2)
		slice.Swap(0, 1)
		assert.Equal(t, []interface{}{1, 0, 2}, slice.O())
	}

	// string
	{
		slice := NewInterSliceV("0", "1", "2")
		slice.Swap(0, 1)
		assert.Equal(t, []interface{}{"1", "0", "2"}, slice.O())
	}

	// custom
	{
		slice := NewInterSlice([]Object{{0}, {1}, {2}})
		slice.Swap(0, 1)
		assert.Equal(t, []interface{}{Object{1}, Object{0}, Object{2}}, slice.O())
	}
}

// Take
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_Take() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.Take(0, 1))
	// Output: [1 2]
}

func TestInterSlice_Take(t *testing.T) {

	// nil or empty
	{
		var slice *InterSlice
		assert.Equal(t, NewInterSliceV(), slice.Take(0, 1))
	}

	// invalid
	{
		slice := NewInterSliceV(1, 2, 3, 4)
		assert.Equal(t, []interface{}{}, slice.Take(1).O())
		assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.O())
		assert.Equal(t, []interface{}{}, slice.Take(4, 4).O())
		assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.O())
	}

	// take 1
	{
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1}, slice.Take(0, 0).O())
			assert.Equal(t, []interface{}{2, 3, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{2}, slice.Take(1, 1).O())
			assert.Equal(t, []interface{}{1, 3, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{3}, slice.Take(2, 2).O())
			assert.Equal(t, []interface{}{1, 2, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{4}, slice.Take(3, 3).O())
			assert.Equal(t, []interface{}{1, 2, 3}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{4}, slice.Take(-1, -1).O())
			assert.Equal(t, []interface{}{1, 2, 3}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{3}, slice.Take(-2, -2).O())
			assert.Equal(t, []interface{}{1, 2, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{2}, slice.Take(-3, -3).O())
			assert.Equal(t, []interface{}{1, 3, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1}, slice.Take(-4, -4).O())
			assert.Equal(t, []interface{}{2, 3, 4}, slice.O())
		}
	}

	// take 2
	{
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2}, slice.Take(0, 1).O())
			assert.Equal(t, []interface{}{3, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{2, 3}, slice.Take(1, 2).O())
			assert.Equal(t, []interface{}{1, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{3, 4}, slice.Take(2, 3).O())
			assert.Equal(t, []interface{}{1, 2}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{3, 4}, slice.Take(-2, -1).O())
			assert.Equal(t, []interface{}{1, 2}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{2, 3}, slice.Take(-3, -2).O())
			assert.Equal(t, []interface{}{1, 4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2}, slice.Take(-4, -3).O())
			assert.Equal(t, []interface{}{3, 4}, slice.O())
		}
	}

	// take 3
	{
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2, 3}, slice.Take(0, 2).O())
			assert.Equal(t, []interface{}{4}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{2, 3, 4}, slice.Take(-3, -1).O())
			assert.Equal(t, []interface{}{1}, slice.O())
		}
	}

	// take everything and beyond
	{
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.Take().O())
			assert.Equal(t, []interface{}{}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.Take(0, 3).O())
			assert.Equal(t, []interface{}{}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.Take(0, -1).O())
			assert.Equal(t, []interface{}{}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.Take(-4, -1).O())
			assert.Equal(t, []interface{}{}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.Take(-6, -1).O())
			assert.Equal(t, []interface{}{}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1, 2, 3, 4}, slice.Take(0, 10).O())
			assert.Equal(t, []interface{}{}, slice.O())
		}
	}

	// move index within bounds
	{
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{4}, slice.Take(3, 4).O())
			assert.Equal(t, []interface{}{1, 2, 3}, slice.O())
		}
		{
			slice := NewInterSliceV(1, 2, 3, 4)
			assert.Equal(t, []interface{}{1}, slice.Take(-5, 0).O())
			assert.Equal(t, []interface{}{2, 3, 4}, slice.O())
		}
	}
}

// TakeAt
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_TakeAt() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.TakeAt(2).O())
	// Output: 3
}

func TestInterSlice_TakeAt(t *testing.T) {

	// int
	{
		// nil or empty
		{
			var nilSlice *InterSlice
			assert.Equal(t, &Object{}, nilSlice.TakeAt(0))
		}

		// Delete all and more
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(-1)
			assert.Equal(t, Obj(2), obj)
			assert.Equal(t, []interface{}{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, Obj(1), obj)
			assert.Equal(t, []interface{}{0}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []interface{}{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.TakeAt(-1)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []interface{}{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(3)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []interface{}{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(2)
			assert.Equal(t, Obj(2), obj)
			assert.Equal(t, []interface{}{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(1)
			assert.Equal(t, Obj(1), obj)
			assert.Equal(t, []interface{}{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(0)
			assert.Equal(t, &Object{0}, obj)
			assert.Equal(t, []interface{}{1, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(-4)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []interface{}{0, 1, 2}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(-1)
			assert.Equal(t, Obj(2), obj)
			assert.Equal(t, []interface{}{0, 1}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := NewInterSliceV(0, 1, 2)
			obj := slice.TakeAt(-2)
			assert.Equal(t, Obj(1), obj)
			assert.Equal(t, []interface{}{0, 2}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// custom
	{
		// Delete all and more
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, slice.O())
			assert.Equal(t, 2, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []interface{}{Object{0}}, slice.O())
			assert.Equal(t, 1, slice.Len())

			obj = slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []interface{}{}, slice.O())
			assert.Equal(t, 0, slice.Len())

			// delete nothing
			obj = slice.TakeAt(-1)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []interface{}{}, slice.O())
			assert.Equal(t, 0, slice.Len())
		}

		// Pos: delete invalid
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(3)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Pos: delete last
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(2)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos: delete middle
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(1)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []interface{}{Object{0}, Object{2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Pos delete first
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(0)
			assert.Equal(t, &Object{Object{0}}, obj)
			assert.Equal(t, []interface{}{Object{1}, Object{2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete invalid
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-4)
			assert.Equal(t, Obj(nil), obj)
			assert.Equal(t, []interface{}{Object{0}, Object{1}, Object{2}}, slice.O())
			assert.Equal(t, 3, slice.Len())
		}

		// Neg: delete last
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-1)
			assert.Equal(t, &Object{Object{2}}, obj)
			assert.Equal(t, []interface{}{Object{0}, Object{1}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}

		// Neg: delete middle
		{
			slice := NewInterSlice([]Object{{0}, {1}, {2}})
			obj := slice.TakeAt(-2)
			assert.Equal(t, &Object{Object{1}}, obj)
			assert.Equal(t, []interface{}{Object{0}, Object{2}}, slice.O())
			assert.Equal(t, 2, slice.Len())
		}
	}
}

// TakeW
//--------------------------------------------------------------------------------------------------
func ExampleInterSlice_TakeW() {
	slice := NewInterSliceV(1, 2, 3)
	fmt.Println(slice.TakeW(func(x O) bool {
		return ExB(x.(int)%2 == 0)
	}))
	// Output: [2]
}

func TestInterSlice_TakeW(t *testing.T) {

	// take all odd values
	{
		slice := NewInterSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.TakeW(func(x O) bool { return ExB(x.(int)%2 != 0) })
		assert.Equal(t, []interface{}{2, 4, 6, 8}, slice.O())
		assert.Equal(t, []interface{}{1, 3, 5, 7, 9}, new.O())
	}

	// take all even values
	{
		slice := NewInterSliceV(1, 2, 3, 4, 5, 6, 7, 8, 9)
		new := slice.TakeW(func(x O) bool { return ExB(x.(int)%2 == 0) })
		assert.Equal(t, []interface{}{1, 3, 5, 7, 9}, slice.O())
		assert.Equal(t, []interface{}{2, 4, 6, 8}, new.O())
	}
}

// // // Union
// // //--------------------------------------------------------------------------------------------------
// // func BenchmarkInterSlice_Union_Go(t *testing.B) {
// // 	// ints := Range(0, nines7)
// // 	// for len(ints) > 10 {
// // 	// 	ints = ints[10:]
// // 	// }
// // }

// // func BenchmarkInterSlice_Union_Slice(t *testing.B) {
// // 	// slice := NewInterSlice(Range(0, nines7))
// // 	// for slice.Len() > 0 {
// // 	// 	slice.PopN(10)
// // 	// }
// // }

// // func ExampleInterSlice_Union() {
// // 	slice := NewInterSliceV(1, 2)
// // 	fmt.Println(slice.Union([]int{2, 3}))
// // 	// Output: [1 2 3]
// // }

// // func TestInterSlice_Union(t *testing.T) {

// // 	// nil or empty
// // 	{
// // 		var slice *InterSlice
// // 		assert.Equal(t, NewInterSliceV(1, 2), slice.Union(NewInterSliceV(1, 2)))
// // 		assert.Equal(t, NewInterSliceV(1, 2), slice.Union([]int{1, 2}))
// // 	}

// // 	// size of one
// // 	{
// // 		slice := NewInterSliceV(1)
// // 		union := slice.Union([]int{1, 2, 3})
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), union)
// // 		assert.Equal(t, NewInterSliceV(1), slice)
// // 	}

// // 	// one duplicate
// // 	{
// // 		slice := NewInterSliceV(1, 1)
// // 		union := slice.Union(NewInterSliceV(2, 3))
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), union)
// // 		assert.Equal(t, NewInterSliceV(1, 1), slice)
// // 	}

// // 	// multiple duplicates
// // 	{
// // 		slice := NewInterSliceV(1, 2, 2, 3, 3)
// // 		union := slice.Union([]int{1, 2, 3})
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), union)
// // 		assert.Equal(t, NewInterSliceV(1, 2, 2, 3, 3), slice)
// // 	}

// // 	// no duplicates
// // 	{
// // 		slice := NewInterSliceV(1, 2, 3)
// // 		union := slice.Union([]int{4, 5})
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3, 4, 5), union)
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), slice)
// // 	}
// // }

// // // UnionM
// // //--------------------------------------------------------------------------------------------------
// // func BenchmarkInterSlice_UnionM_Go(t *testing.B) {
// // 	// ints := Range(0, nines7)
// // 	// for len(ints) > 10 {
// // 	// 	ints = ints[10:]
// // 	// }
// // }

// // func BenchmarkInterSlice_UnionM_Slice(t *testing.B) {
// // 	// slice := NewInterSlice(Range(0, nines7))
// // 	// for slice.Len() > 0 {
// // 	// 	slice.PopN(10)
// // 	// }
// // }

// // func ExampleInterSlice_UnionM() {
// // 	slice := NewInterSliceV(1, 2)
// // 	fmt.Println(slice.UnionM([]int{2, 3}))
// // 	// Output: [1 2 3]
// // }

// // func TestInterSlice_UnionM(t *testing.T) {

// // 	// nil or empty
// // 	{
// // 		var slice *InterSlice
// // 		assert.Equal(t, NewInterSliceV(1, 2), slice.UnionM(NewInterSliceV(1, 2)))
// // 		assert.Equal(t, (*InterSlice)(nil), slice)
// // 	}

// // 	// size of one
// // 	{
// // 		slice := NewInterSliceV(1)
// // 		union := slice.UnionM([]int{1, 2, 3})
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), union)
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), slice)
// // 	}

// // 	// one duplicate
// // 	{
// // 		slice := NewInterSliceV(1, 1)
// // 		union := slice.UnionM(NewInterSliceV(2, 3))
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), union)
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), slice)
// // 	}

// // 	// multiple duplicates
// // 	{
// // 		slice := NewInterSliceV(1, 2, 2, 3, 3)
// // 		union := slice.UnionM([]int{1, 2, 3})
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), union)
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3), slice)
// // 	}

// // 	// no duplicates
// // 	{
// // 		slice := NewInterSliceV(1, 2, 3)
// // 		union := slice.UnionM([]int{4, 5})
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3, 4, 5), union)
// // 		assert.Equal(t, NewInterSliceV(1, 2, 3, 4, 5), slice)
// // 	}
// // }

// // Uniq
// //--------------------------------------------------------------------------------------------------
// func BenchmarkInterSlice_Uniq_Go(t *testing.B) {
// 	// ints := Range(0, nines7)
// 	// for len(ints) > 10 {
// 	// 	ints = ints[10:]
// 	// }
// }

// func BenchmarkInterSlice_Uniq_Slice(t *testing.B) {
// 	// slice := NewInterSlice(Range(0, nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleInterSlice_Uniq() {
// 	slice := NewInterSliceV(1, 2, 3, 3)
// 	fmt.Println(slice.Uniq())
// 	// Output: [1 2 3]
// }

// func TestInterSlice_Uniq(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *InterSlice
// 		assert.Equal(t, NewInterSliceV(), slice.Uniq())
// 	}

// 	// size of one
// 	{
// 		slice := NewInterSliceV(1)
// 		uniq := slice.Uniq()
// 		assert.Equal(t, []int{1}, uniq.O())
// 		assert.Equal(t, []int{1, 2}, slice.Append(2).O())
// 		assert.Equal(t, []int{1}, uniq.O())
// 	}

// 	// one duplicate
// 	{
// 		slice := NewInterSliceV(1, 1)
// 		uniq := slice.Uniq()
// 		assert.Equal(t, []int{1}, uniq.O())
// 		assert.Equal(t, []int{1, 1, 2}, slice.Append(2).O())
// 		assert.Equal(t, []int{1}, uniq.O())
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewInterSliceV(1, 2, 2, 3, 3)
// 		uniq := slice.Uniq()
// 		assert.Equal(t, []int{1, 2, 3}, uniq.O())
// 		assert.Equal(t, []int{1, 2, 2, 3, 3, 4}, slice.Append(4).O())
// 		assert.Equal(t, []int{1, 2, 3}, uniq.O())
// 	}

// 	// no duplicates
// 	{
// 		slice := NewInterSliceV(1, 2, 3)
// 		uniq := slice.Uniq()
// 		assert.Equal(t, []int{1, 2, 3}, uniq.O())
// 		assert.Equal(t, []int{1, 2, 3, 4}, slice.Append(4).O())
// 		assert.Equal(t, []int{1, 2, 3}, uniq.O())
// 	}
// }

// // UniqM
// //--------------------------------------------------------------------------------------------------
// func BenchmarkInterSlice_UniqM_Go(t *testing.B) {
// 	// ints := Range(0, nines7)
// 	// for len(ints) > 10 {
// 	// 	ints = ints[10:]
// 	// }
// }

// func BenchmarkInterSlice_UniqM_Slice(t *testing.B) {
// 	// slice := NewInterSlice(Range(0, nines7))
// 	// for slice.Len() > 0 {
// 	// 	slice.PopN(10)
// 	// }
// }

// func ExampleInterSlice_UniqM() {
// 	slice := NewInterSliceV(1, 2, 3, 3)
// 	fmt.Println(slice.UniqM())
// 	// Output: [1 2 3]
// }

// func TestInterSlice_UniqM(t *testing.T) {

// 	// nil or empty
// 	{
// 		var slice *InterSlice
// 		assert.Equal(t, (*InterSlice)(nil), slice.UniqM())
// 	}

// 	// size of one
// 	{
// 		slice := NewInterSliceV(1)
// 		uniq := slice.UniqM()
// 		assert.Equal(t, []int{1}, uniq.O())
// 		assert.Equal(t, []int{1, 2}, slice.Append(2).O())
// 		assert.Equal(t, []int{1, 2}, uniq.O())
// 	}

// 	// one duplicate
// 	{
// 		slice := NewInterSliceV(1, 1)
// 		uniq := slice.UniqM()
// 		assert.Equal(t, []int{1}, uniq.O())
// 		assert.Equal(t, []int{1, 2}, slice.Append(2).O())
// 		assert.Equal(t, []int{1, 2}, uniq.O())
// 	}

// 	// multiple duplicates
// 	{
// 		slice := NewInterSliceV(1, 2, 2, 3, 3)
// 		uniq := slice.UniqM()
// 		assert.Equal(t, []int{1, 2, 3}, uniq.O())
// 		assert.Equal(t, []int{1, 2, 3, 4}, slice.Append(4).O())
// 		assert.Equal(t, []int{1, 2, 3, 4}, uniq.O())
// 	}

// 	// no duplicates
// 	{
// 		slice := NewInterSliceV(1, 2, 3)
// 		uniq := slice.UniqM()
// 		assert.Equal(t, []int{1, 2, 3}, uniq.O())
// 		assert.Equal(t, []int{1, 2, 3, 4}, slice.Append(4).O())
// 		assert.Equal(t, []int{1, 2, 3, 4}, uniq.O())
// 	}
// }
