package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Slice function
//--------------------------------------------------------------------------------------------------
func ExampleSlice() {
	slice := Slice([]int{1, 2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_Slice_Constructor(t *testing.T) {

	// arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, []string{"1", "2"}, Slice(array).O())

	// empty
	assert.Equal(t, nil, Slice(nil).O())
	assert.Equal(t, &NSlice{}, Slice(nil))
	assert.Equal(t, []int{}, Slice([]int{}).O())
	assert.Equal(t, []bool{}, Slice([]bool{}).O())
	assert.Equal(t, []string{}, Slice([]string{}).O())
	assert.Equal(t, []NObj{}, Slice([]NObj{}).O())
	assert.Equal(t, nil, Slice([]interface{}{}).O())

	// pointers
	var obj *NObj
	assert.Equal(t, []*NObj{nil}, Slice(obj).O())
	assert.Equal(t, []*NObj{&(NObj{"bob"})}, Slice(&(NObj{"bob"})).O())
	assert.Equal(t, []*NObj{&(NObj{"1"}), &(NObj{"2"})}, Slice([]*NObj{&(NObj{"1"}), &(NObj{"2"})}).O())

	// interface
	assert.Equal(t, nil, Slice([]interface{}{nil}).O())
	assert.Equal(t, []string{""}, Slice([]interface{}{nil, ""}).O())
	assert.Equal(t, []bool{true}, Slice([]interface{}{true}).O())
	assert.Equal(t, []int{1}, Slice([]interface{}{1}).O())
	assert.Equal(t, []string{""}, Slice([]interface{}{""}).O())
	assert.Equal(t, []string{"bob"}, Slice([]interface{}{"bob"}).O())
	assert.Equal(t, []NObj{{nil}}, Slice([]interface{}{NObj{}}).O())

	// singles
	assert.Equal(t, []int{1}, Slice(1).O())
	assert.Equal(t, []bool{true}, Slice(true).O())
	assert.Equal(t, []string{""}, Slice("").O())
	assert.Equal(t, []string{"1"}, Slice("1").O())
	assert.Equal(t, []NObj{{1}}, Slice(NObj{1}).O())
	assert.Equal(t, []NObj{NObj{"bob"}}, Slice(NObj{"bob"}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, Slice(map[string]string{"1": "one"}).O())

	// slices
	assert.Equal(t, []int{1, 2}, Slice([]int{1, 2}).O())
	assert.Equal(t, []bool{true}, Slice([]bool{true}).O())
	assert.Equal(t, []NObj{{"bob"}}, Slice([]NObj{{"bob"}}).O())
	assert.Equal(t, []string{"1", "2"}, Slice([]string{"1", "2"}).O())
	assert.Equal(t, [][]string{{"1"}}, Slice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, Slice([]interface{}{map[string]string{"1": "one"}}).O())
}

// SliceV function
//--------------------------------------------------------------------------------------------------
func ExampleSliceV_empty() {
	slice := SliceV()
	fmt.Println(slice.O())
	// Output: <nil>
}

func ExampleSliceV_variadic() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_SliceV(t *testing.T) {
	var obj *NObj

	// Arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, [][2]string{array}, SliceV(array).O())

	// Test empty values
	assert.True(t, !SliceV().Any())
	assert.Equal(t, 0, SliceV().Len())
	assert.Equal(t, nil, SliceV().O())
	assert.Equal(t, nil, SliceV(nil).O())
	assert.Equal(t, &NSlice{}, SliceV(nil))
	assert.Equal(t, []string{""}, SliceV(nil, "").O())
	assert.Equal(t, []*NObj{nil}, SliceV(nil, obj).O())

	// Test pointers
	assert.Equal(t, []*NObj{nil}, SliceV(obj).O())
	assert.Equal(t, []*NObj{&(NObj{"bob"})}, SliceV(&(NObj{"bob"})).O())
	assert.Equal(t, []*NObj{nil}, SliceV(obj).O())
	assert.Equal(t, []*NObj{&(NObj{"bob"})}, SliceV(&(NObj{"bob"})).O())
	assert.Equal(t, [][]*NObj{{&(NObj{"1"}), &(NObj{"2"})}}, SliceV([]*NObj{&(NObj{"1"}), &(NObj{"2"})}).O())

	// Singles
	assert.Equal(t, []int{1}, SliceV(1).O())
	assert.Equal(t, []string{"1"}, SliceV("1").O())
	assert.Equal(t, []NObj{NObj{"bob"}}, SliceV(NObj{"bob"}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, SliceV(map[string]string{"1": "one"}).O())

	// Multiples
	assert.Equal(t, []int{1, 2}, SliceV(1, 2).O())
	assert.Equal(t, []string{"1", "2"}, SliceV("1", "2").O())
	assert.Equal(t, []NObj{NObj{1}, NObj{2}}, SliceV(NObj{1}, NObj{2}).O())

	// Test slices
	assert.Equal(t, [][]int{{1, 2}}, SliceV([]int{1, 2}).O())
	assert.Equal(t, [][]string{{"1"}}, SliceV([]string{"1"}).O())
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
	assert.Equal(t, []NObj{}, newEmptySlice(NObj{1}).O())

	// Slices
	assert.Equal(t, []int{}, newEmptySlice([]int{1, 2}).O())
	assert.Equal(t, []bool{}, newEmptySlice([]bool{true}).O())
	assert.Equal(t, []string{}, newEmptySlice([]string{"bob"}).O())
	assert.Equal(t, []NObj{}, newEmptySlice([]NObj{{"bob"}}).O())
	assert.Equal(t, [][]string{}, newEmptySlice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []map[string]string{}, newEmptySlice([]interface{}{map[string]string{"1": "one"}}).O())

	// Empty slices
	assert.Equal(t, []int{}, newEmptySlice([]int{}).O())
	assert.Equal(t, []bool{}, newEmptySlice([]bool{}).O())
	assert.Equal(t, []string{}, newEmptySlice([]string{}).O())
	assert.Equal(t, []NObj{}, newEmptySlice([]NObj{}).O())

	// Interface types
	assert.Equal(t, []interface{}{}, newEmptySlice(nil).O())
	assert.Equal(t, []interface{}{}, newEmptySlice([]interface{}{nil}).O())
	assert.Equal(t, []int{}, newEmptySlice([]interface{}{1}).O())
	assert.Equal(t, []int{}, newEmptySlice([]interface{}{interface{}(1)}).O())
	assert.Equal(t, []string{}, newEmptySlice([]interface{}{""}).O())
	assert.Equal(t, []NObj{}, newEmptySlice([]interface{}{NObj{}}).O())
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
	slice := Slice(src)
	for i := range src {
		slice.Any(i)
	}
}

func BenchmarkNSlice_Any_Reflect(t *testing.B) {
	src := rangeNObj(0, nines4)
	slice := Slice(src)
	for _, i := range src {
		slice.Any(i)
	}
}

func ExampleNSlice_Any_empty() {
	slice := SliceV()
	fmt.Println(slice.Any())
	// Output: false
}

func ExampleNSlice_Any_notEmpty() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.Any())
	// Output: true
}

func ExampleNSlice_Any_contains() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.Any(1))
	// Output: true
}

func ExampleNSlice_Any_containsAny() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.Any(0, 1))
	// Output: true
}

func TestNSlice_Any(t *testing.T) {
	var nilSlice *NSlice
	assert.False(t, nilSlice.Any())
	assert.False(t, SliceV().Any())
	assert.True(t, SliceV().Append("2").Any())

	// bool
	assert.True(t, SliceV(false, true).Any(true))
	assert.False(t, SliceV(true, true).Any(false))
	assert.True(t, SliceV(true, true).Any(false, true))
	assert.False(t, SliceV(true, true).Any(false, false))

	// int
	assert.True(t, SliceV(1, 2, 3).Any(2))
	assert.False(t, SliceV(1, 2, 3).Any(4))
	assert.True(t, SliceV(1, 2, 3).Any(4, 3))
	assert.False(t, SliceV(1, 2, 3).Any(4, 5))

	// int64
	assert.True(t, SliceV(int64(1), int64(2), int64(3)).Any(int64(2)))
	assert.False(t, SliceV(int64(1), int64(2), int64(3)).Any(int64(4)))
	assert.True(t, SliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(2)))
	assert.False(t, SliceV(int64(1), int64(2), int64(3)).Any(int64(4), int64(5)))

	// string
	assert.True(t, SliceV("1", "2", "3").Any("2"))
	assert.False(t, SliceV("1", "2", "3").Any("4"))
	assert.True(t, SliceV("1", "2", "3").Any("4", "2"))
	assert.False(t, SliceV("1", "2", "3").Any("4", "5"))

	// custom
	assert.True(t, SliceV(NObj{1}, NObj{2}).Any(NObj{1}))
	assert.False(t, SliceV(NObj{1}, NObj{2}).Any(NObj{3}))
	assert.True(t, SliceV(NObj{1}, NObj{2}).Any(NObj{4}, NObj{2}))
	assert.False(t, SliceV(NObj{1}, NObj{2}).Any(NObj{4}, NObj{5}))

	// panics need to go as last item as they abort the test method
	defer func() {
		err := recover()
		assert.Equal(t, "can't compare type 'int' with '[]n.NObj' elements", err)
	}()
	assert.True(t, SliceV(NObj{1}, NObj{2}).Any(2))
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
	slice := Slice(src)
	for i := range src {
		slice.Any([]int{i})
	}
}

func BenchmarkNSlice_AnyS_Reflect(t *testing.B) {
	src := rangeNObj(0, nines4)
	slice := Slice(src)
	for _, i := range src {
		slice.Any([]NObj{i})
	}
}

func ExampleNSlice_AnyS() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.AnyS([]int{0, 1}))
	// Output: true
}

func TestNSlice_AnyS(t *testing.T) {
	var nilSlice *NSlice
	assert.False(t, nilSlice.AnyS([]bool{true}))

	// bool
	assert.True(t, SliceV(true, true).AnyS([]bool{true}))
	assert.True(t, SliceV(true, true).AnyS([]bool{false, true}))
	assert.False(t, SliceV(true, true).AnyS([]bool{false, false}))

	// int
	assert.True(t, SliceV(1, 2, 3).AnyS([]int{1}))
	assert.True(t, SliceV(1, 2, 3).AnyS([]int{4, 3}))
	assert.False(t, SliceV(1, 2, 3).AnyS([]int{4, 5}))

	// int64
	assert.True(t, SliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(2)}))
	assert.True(t, SliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(2)}))
	assert.False(t, SliceV(int64(1), int64(2), int64(3)).AnyS([]int64{int64(4), int64(5)}))

	// string
	assert.True(t, SliceV("1", "2", "3").AnyS([]string{"2"}))
	assert.True(t, SliceV("1", "2", "3").AnyS([]string{"4", "2"}))
	assert.False(t, SliceV("1", "2", "3").AnyS([]string{"4", "5"}))

	// custom
	assert.True(t, SliceV(NObj{1}, NObj{2}).AnyS([]NObj{{2}}))
	assert.True(t, SliceV(NObj{1}, NObj{2}).AnyS([]NObj{{4}, {2}}))
	assert.False(t, SliceV(NObj{1}, NObj{2}).AnyS([]NObj{{4}, {5}}))

	// panics need to go as last item as they abort the test method
	defer func() {
		err := recover()
		assert.Equal(t, "can't compare type '[]int' with '[]n.NObj' elements", err)
	}()
	assert.True(t, SliceV(NObj{1}, NObj{2}).AnyS([]int{2}))
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
	n := &NSlice{o: []NObj{}}
	for _, i := range Range(0, nines6) {
		n.Append(NObj{i})
	}
}

func ExampleNSlice_Append() {
	slice := SliceV(1).Append(2).Append(3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_Append_Reflect(t *testing.T) {

	// Use a custom type to invoke reflection
	n := SliceV(NObj{"1"})
	assert.Equal(t, 1, n.Len())
	assert.Equal(t, false, n.Nil())
	assert.Equal(t, []NObj{{"1"}}, n.O())

	// Append another to it
	n.Append(NObj{"2"})
	assert.Equal(t, 2, n.Len())
	assert.Equal(t, []NObj{{"1"}, {"2"}}, n.O())

	// Given an invalid type which will abort the function so put at end
	defer func() {
		err := recover()
		assert.Equal(t, "reflect.Set: value of type int is not assignable to type n.NObj", err)
	}()
	n.Append(2)
}

func TestNSlice_Append(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, SliceV(0), nilSlice.Append(0))
		assert.Equal(t, (*NSlice)(nil), nilSlice)
	}

	// Append one back to back
	{
		n := SliceV()
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
		n := SliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1)
		assert.Equal(t, []int{1}, n.O())
		n.Append(2)
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with nil not chained
	{
		n := SliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Start with nil chained
	{
		n := SliceV().Append(1).Append(2)
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with non nil
	{
		n := SliceV(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Use append result directly
	{
		n := SliceV(1)
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []int{1, 2}, n.Append(2).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := SliceV(true)
			assert.Equal(t, []bool{true, false}, n.Append(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := SliceV(0)
			assert.Equal(t, []int{0, 1}, n.Append(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := SliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.Append("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := Slice([]NObj{{"3"}})
			assert.Equal(t, []NObj{{"3"}, {"1"}}, n.Append(NObj{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}

	// Invalid type which will abort the function so put at end
	{
		n := SliceV("1")
		defer func() {
			err := recover()
			assert.Equal(t, "can't insert type 'int' into '[]string'", err)
		}()
		n.Append(2)
	}
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
	n := &NSlice{o: []NObj{}}
	new := rangeNObjO(0, nines6)
	n.AppendV(new...)
}

func ExampleNSlice_AppendV() {
	slice := SliceV(1).AppendV(2, 3)
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
		n := SliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.AppendV(2, 3).O())
	}

	// Append many strings
	{
		{
			n := SliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendV("1", "2", "3").O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := Slice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendV("2", "3").O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := Slice([]NObj{{"3"}})
		assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendV(NObj{"1"}).O())
		assert.Equal(t, []NObj{{"3"}, {"1"}, {"2"}, {"4"}}, n.AppendV(NObj{"2"}, NObj{"4"}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := SliceV(true)
			assert.Equal(t, []bool{true, false}, n.AppendV(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := SliceV(0)
			assert.Equal(t, []int{0, 1}, n.AppendV(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := SliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.AppendV("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := Slice([]NObj{{"3"}})
			assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendV(NObj{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}

	// Append to a slice of map
	{
		n := SliceV(map[string]string{"1": "one"})
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
	dest := &NSlice{o: []NObj{}}
	src := rangeNObj(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func BenchmarkNSlice_AppendS_Reflect100(t *testing.B) {
	dest := &NSlice{o: []NObj{}}
	src := rangeNObj(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func ExampleNSlice_AppendS() {
	slice := SliceV(1).AppendS([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_AppendS(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, SliceV(1, 2), nilSlice.AppendS([]int{1, 2}))
		assert.Equal(t, (*NSlice)(nil), nilSlice)
	}

	// Append many ints
	{
		n := SliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.AppendS([]int{2, 3}).O())
	}

	// Append many strings
	{
		{
			n := SliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendS([]string{"1", "2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := Slice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendS([]string{"2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := Slice([]NObj{{"3"}})
		assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendS([]NObj{{"1"}}).O())
		assert.Equal(t, []NObj{{"3"}, {"1"}, {"2"}, {"4"}}, n.AppendS([]NObj{{"2"}, {"4"}}).O())
	}

	// Append to a slice of map
	{
		n := SliceV(map[string]string{"1": "one"})
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
			n := SliceV(true)
			assert.Equal(t, []bool{true, false}, n.AppendS([]bool{false}).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := SliceV(0)
			assert.Equal(t, []int{0, 1}, n.AppendS([]int{1}).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := SliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.AppendS([]string{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := Slice([]NObj{{"3"}})
			assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendS([]NObj{{"1"}}).O())
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
	slice := Slice(src)
	for _, i := range src {
		_, ok := (slice.At(i).O()).(int)
		assert.True(t, ok)
	}
}

func BenchmarkNSlice_At_Reflect(t *testing.B) {
	src := rangeNObj(0, nines6)
	slice := Slice(src)
	for i := range src {
		_, ok := (slice.At(i).O()).(NObj)
		assert.True(t, ok)
	}
}

func ExampleNSlice_At() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.At(2).O())
	// Output: 3
}

func TestNSlice_indexAbs(t *testing.T) {
	//             -4,-3,-2,-1
	//              0, 1, 2, 3
	slice := SliceV(0, 0, 0, 0)
	assert.Equal(t, 3, slice.absIndex(-1))
	assert.Equal(t, 2, slice.absIndex(-2))
	assert.Equal(t, 1, slice.absIndex(-3))
	assert.Equal(t, 0, slice.absIndex(-4))

	assert.Equal(t, 0, slice.absIndex(0))
	assert.Equal(t, 1, slice.absIndex(1))
	assert.Equal(t, 2, slice.absIndex(2))
	assert.Equal(t, 3, slice.absIndex(3))

	// out of bounds
	assert.Equal(t, -1, slice.absIndex(4))
	assert.Equal(t, -1, slice.absIndex(-5))
}

func TestNSlice_At(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, &NObj{nil}, nilSlice.At(0))
	}

	// strings
	{
		slice := SliceV("1", "2", "3", "4")
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
		slice := SliceV("1")
		assert.Equal(t, &NObj{}, slice.At(3))
		assert.Equal(t, nil, slice.At(3).O())
		assert.Equal(t, &NObj{}, slice.At(-3))
		assert.Equal(t, nil, slice.At(-3).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------

func ExampleNSlice_Clear() {
	slice := SliceV(1).AppendS([]int{2, 3})
	fmt.Println(slice.Clear().O())
	// Output: []
}

func TestQSlice_Clear(t *testing.T) {

	// nil
	{
		var nilSlice *NSlice
		assert.Equal(t, &NObj{nil}, nilSlice.At(0))
	}

	// bool
	{
		slice := SliceV(true, false)
		assert.Equal(t, 2, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []bool{}}, slice)
	}

	// int
	{
		slice := SliceV(1, 2, 3, 4)
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []int{}}, slice)
	}

	// string
	{
		slice := SliceV("1", "2", "3", "4")
		assert.Equal(t, 4, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []string{}}, slice)
	}

	// custom
	{
		slice := Slice([]NObj{{1}, {2}, {3}})
		assert.Equal(t, 3, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		slice.Clear()
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, &NSlice{o: []NObj{}}, slice)
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Copy_Normal(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkNSlice_Copy_Optimized(t *testing.B) {
	slice := Slice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func BenchmarkNSlice_Copy_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleNSlice_Copy() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.DropFirst().O())
	// Output: [2 3]
}

func TestNSlice_Copy(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, SliceV(), nilSlice.Copy(0, -1))
		slice := SliceV(0).Clear()
		assert.Equal(t, &NSlice{o: []int{}}, slice.Copy(0, -1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := SliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, []int{0, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2, 3}, result.O())
	}

	// slice full array
	{
		assert.Equal(t, SliceV(), SliceV().Slice(0, -1))
		assert.Equal(t, SliceV(), SliceV().Slice(0, 1))
		assert.Equal(t, SliceV(), SliceV().Slice(0, 5))
		assert.Equal(t, SliceV(""), SliceV("").Slice(0, -1))
		assert.Equal(t, SliceV(""), SliceV("").Slice(0, 1))
		assert.Equal(t, SliceV(1, 2, 3), SliceV(1, 2, 3).Slice(0, -1))
		assert.Equal(t, Slice([]int{1, 2, 3}), Slice([]int{1, 2, 3}).Slice(0, -1))
		assert.Equal(t, SliceV("1", "2", "3"), SliceV("1", "2", "3").Slice(0, 2))
		assert.Equal(t, Slice([]NObj{{1}, {2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, SliceV("1"), SliceV("1").Slice(0, 2))
		assert.Equal(t, SliceV(true, false), SliceV(true, false).Slice(-6, 6))
		assert.Equal(t, SliceV(1, 2, 3), SliceV(1, 2, 3).Slice(-6, 6))
		assert.Equal(t, SliceV("1", "2", "3"), SliceV("1", "2", "3").Slice(-6, 6))
		assert.Equal(t, Slice([]NObj{{1}, {2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-6, 6))
	}

	// mutually exclusive
	{
		slice := SliceV(1, 2, 3, 4)
		assert.Equal(t, &NSlice{}, slice.Slice(2, -3))
		assert.Equal(t, &NSlice{}, slice.Slice(0, -5))
		assert.Equal(t, &NSlice{}, slice.Slice(4, -1))
		assert.Equal(t, &NSlice{}, slice.Slice(6, -1))
		assert.Equal(t, &NSlice{}, slice.Slice(3, 2))
	}

	// singles
	{
		slice := SliceV(1, 2, 3, 4)
		assert.Equal(t, SliceV(4), slice.Slice(-1, -1))
		assert.Equal(t, SliceV(3), slice.Slice(-2, -2))
		assert.Equal(t, SliceV(2), slice.Slice(-3, -3))
		assert.Equal(t, SliceV(1), slice.Slice(0, 0))
		assert.Equal(t, SliceV(1), slice.Slice(-4, -4))
		assert.Equal(t, SliceV(2), slice.Slice(1, 1))
		assert.Equal(t, SliceV(2), slice.Slice(1, -3))
		assert.Equal(t, SliceV(3), slice.Slice(2, 2))
		assert.Equal(t, SliceV(3), slice.Slice(2, -2))
		assert.Equal(t, SliceV(4), slice.Slice(3, 3))
		assert.Equal(t, SliceV(4), slice.Slice(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(1, -1))
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(1, 2))
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(-2, -1))
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(-2, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(1, -1))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(-2, -1))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(-2, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(1, -1))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(-2, -1))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(-2, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(1, -1))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(1, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-2, -1))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(0, -2))
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(-3, -2))
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(-3, 1))
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(0, 1))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(0, -2))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(-3, -2))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(-3, 1))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(0, 1))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(0, -2))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(-3, -2))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(-3, 1))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(0, 1))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(0, -2))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-3, -2))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-3, 1))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(1, -2))
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(-3, -2))
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(-3, 2))
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(1, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(1, -2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(-3, -2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(-3, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(1, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(1, -2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(-3, -2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(-3, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(1, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(1, -2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(-3, -2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(-3, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(1, 2))
	}

	// random
	{
		assert.Equal(t, SliceV("1"), SliceV("1", "2", "3").Slice(0, -3))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, SliceV("1", "2", "3"), SliceV("1", "2", "3").Slice(0, 2))
	}
}

// DeleteAt
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_DeleteAt_Normal(t *testing.B) {
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

func BenchmarkNSlice_DeleteAt_Optimized(t *testing.B) {
	src := Range(0, nines5)
	index := Range(0, nines5)
	slice := Slice(src)
	for i := range index {
		slice.DeleteAt(i)
	}
}

func BenchmarkNSlice_DeleteAt_Reflect(t *testing.B) {
	src := rangeNObj(0, nines5)
	index := Range(0, nines5)
	slice := Slice(src)
	for i := range index {
		slice.DeleteAt(i)
	}
}

func ExampleNSlice_DeleteAt() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.At(2).O())
	// Output: 3
}

func TestNSlice_DeleteAt(t *testing.T) {

	// Delete all and more
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(-1)
		assert.Equal(t, &NObj{2}, obj)
		assert.Equal(t, []int{0, 1}, slice.O())
		assert.Equal(t, 2, slice.Len())

		obj = slice.DeleteAt(-1)
		assert.Equal(t, &NObj{1}, obj)
		assert.Equal(t, []int{0}, slice.O())
		assert.Equal(t, 1, slice.Len())

		obj = slice.DeleteAt(-1)
		assert.Equal(t, &NObj{0}, obj)
		assert.Equal(t, []int{}, slice.O())
		assert.Equal(t, 0, slice.Len())

		// delete nothing
		obj = slice.DeleteAt(-1)
		assert.Equal(t, &NObj{nil}, obj)
		assert.Equal(t, []int{}, slice.O())
		assert.Equal(t, 0, slice.Len())
	}

	// Pos: delete invalid
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(3)
		assert.Equal(t, &NObj{nil}, obj)
		assert.Equal(t, []int{0, 1, 2}, slice.O())
		assert.Equal(t, 3, slice.Len())
	}

	// Pos: delete last
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(2)
		assert.Equal(t, &NObj{2}, obj)
		assert.Equal(t, []int{0, 1}, slice.O())
		assert.Equal(t, 2, slice.Len())
	}

	// Pos: delete middle
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(1)
		assert.Equal(t, &NObj{1}, obj)
		assert.Equal(t, []int{0, 2}, slice.O())
		assert.Equal(t, 2, slice.Len())
	}

	// Pos delete first
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(0)
		assert.Equal(t, &NObj{0}, obj)
		assert.Equal(t, []int{1, 2}, slice.O())
		assert.Equal(t, 2, slice.Len())
	}

	// Neg: delete invalid
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(-4)
		assert.Equal(t, &NObj{nil}, obj)
		assert.Equal(t, []int{0, 1, 2}, slice.O())
		assert.Equal(t, 3, slice.Len())
	}

	// Neg: delete last
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(-1)
		assert.Equal(t, &NObj{2}, obj)
		assert.Equal(t, []int{0, 1}, slice.O())
		assert.Equal(t, 2, slice.Len())
	}

	// Neg: delete middle
	{
		slice := SliceV(0, 1, 2)
		obj := slice.DeleteAt(-2)
		assert.Equal(t, &NObj{1}, obj)
		assert.Equal(t, []int{0, 2}, slice.O())
		assert.Equal(t, 2, slice.Len())
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
	slice := Slice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func BenchmarkNSlice_DropFirst_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleNSlice_DropFirst() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.DropFirst().O())
	// Output: [2 3]
}

func TestNSlice_DropFirst(t *testing.T) {
	// bool
	{
		slice := SliceV(true, true, false)
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
		slice := SliceV(1, 2, 3)
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
		slice := SliceV("1", "2", "3")
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
		slice := Slice([]NObj{{1}, {2}, {3}})
		assert.Equal(t, []NObj{{2}, {3}}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []NObj{{3}}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []NObj{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []NObj{}, slice.DropFirst().O())
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
	slice := Slice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func BenchmarkNSlice_DropFirstN_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirstN(10)
	}
}

func ExampleNSlice_DropFirstN() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.DropFirstN(2).O())
	// Output: [3]
}

func TestNSlice_DropFirstN(t *testing.T) {

	// drop none
	{
		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2, 3}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{{1}, {2}, {3}}, slice.DropFirstN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := SliceV(true, true, false)
			assert.Equal(t, []bool{true, false}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{2, 3}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := SliceV("1", "2", "3")
			assert.Equal(t, []string{"2", "3"}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{{2}, {3}}, slice.DropFirstN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := SliceV(true, false, false)
			assert.Equal(t, []bool{false}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{3}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := SliceV("1", "2", "3")
			assert.Equal(t, []string{"3"}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{{3}}, slice.DropFirstN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{}, slice.DropFirstN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropFirstN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{}, slice.DropFirstN(4).O())
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
	slice := Slice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func BenchmarkNSlice_DropLast_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines7))
	for slice.Len() > 0 {
		slice.DropLast()
	}
}

func ExampleNSlice_DropLast() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.DropLast().O())
	// Output: [1 2]
}

func TestNSlice_DropLast(t *testing.T) {
	// bool
	{
		slice := SliceV(true, true, false)
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
		slice := SliceV(1, 2, 3)
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
		slice := SliceV("1", "2", "3")
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
		slice := Slice([]NObj{{1}, {2}, {3}})
		assert.Equal(t, []NObj{{1}, {2}}, slice.DropLast().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []NObj{{1}}, slice.DropLast().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []NObj{}, slice.DropLast().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []NObj{}, slice.DropLast().O())
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
	slice := Slice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func BenchmarkNSlice_DropLastN_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines7))
	for slice.Len() > 0 {
		slice.DropLastN(10)
	}
}

func ExampleNSlice_DropLastN() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.DropLastN(2).O())
	// Output: [1]
}

func TestNSlice_DropLastN(t *testing.T) {

	// drop none
	{
		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2, 3}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{{1}, {2}, {3}}, slice.DropLastN(0).O())
			assert.Equal(t, 3, slice.Len())
		}
	}

	// drop 1
	{
		// bool
		{
			slice := SliceV(true, true, false)
			assert.Equal(t, []bool{true, true}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{1, 2}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// string
		{
			slice := SliceV("1", "2", "3")
			assert.Equal(t, []string{"1", "2"}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{{1}, {2}}, slice.DropLastN(1).O())
			assert.Equal(t, 2, slice.Len())
		}
	}

	// drop 2
	{
		// bool
		{
			slice := SliceV(true, false, false)
			assert.Equal(t, []bool{true}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{1}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// string
		{
			slice := SliceV("1", "2", "3")
			assert.Equal(t, []string{"1"}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{{1}}, slice.DropLastN(2).O())
			assert.Equal(t, 1, slice.Len())
		}
	}

	// drop 3
	{
		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{}, slice.DropLastN(3).O())
			assert.Equal(t, 0, slice.Len())
		}
	}

	// drop beyond
	{
		// int
		{
			slice := SliceV(1, 2, 3)
			assert.Equal(t, []int{}, slice.DropLastN(4).O())
			assert.Equal(t, 0, slice.Len())
		}

		// custom
		{
			slice := Slice([]NObj{{1}, {2}, {3}})
			assert.Equal(t, []NObj{}, slice.DropLastN(4).O())
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
	Slice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkNSlice_Each_Reflect(t *testing.B) {
	Slice(rangeNObj(0, nines6)).Each(func(x O) {
		assert.IsType(t, NObj{}, x)
	})
}

func ExampleNSlice_Each() {
	SliceV(1, 2, 3).Each(func(x O) {
		fmt.Printf("%v", x)
	})
	// Output: 123
}

func TestNSlice_Each(t *testing.T) {
	// int
	{
		SliceV(1, 2, 3).Each(func(x O) {
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
		SliceV("1", "2", "3").Each(func(x O) {
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
		Slice([]NObj{{1}, {2}, {3}}).Each(func(x O) {
			switch x {
			case NObj{1}:
				assert.Equal(t, NObj{1}, x)
			case NObj{2}:
				assert.Equal(t, NObj{2}, x)
			case NObj{3}:
				assert.Equal(t, NObj{3}, x)
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
	Slice(Range(0, nines6)).Each(func(x O) {
		assert.IsType(t, 0, x)
	})
}

func BenchmarkNSlice_EachE_Reflect(t *testing.B) {
	Slice(rangeNObj(0, nines6)).Each(func(x O) {
		assert.IsType(t, NObj{}, x)
	})
}

func ExampleNSlice_EachE() {
	SliceV(1, 2, 3).EachE(func(x O) error {
		fmt.Printf("%v", x)
		return nil
	})
	// Output: 123
}

func TestNSlice_EachE(t *testing.T) {
	// int
	{
		SliceV(1, 2, 3).EachE(func(x O) error {
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
		SliceV("1", "2", "3").EachE(func(x O) error {
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
		Slice([]NObj{{1}, {2}, {3}}).EachE(func(x O) error {
			switch x {
			case NObj{1}:
				assert.Equal(t, NObj{1}, x)
			case NObj{2}:
				assert.Equal(t, NObj{2}, x)
			case NObj{3}:
				assert.Equal(t, NObj{3}, x)
			}
			return nil
		})
	}
}

// Empty
//--------------------------------------------------------------------------------------------------
func ExampleNSlice_Empty() {
	fmt.Println(SliceV().Empty())
	// Output: true
}

func TestNSlice_Empty(t *testing.T) {
	assert.Equal(t, true, SliceV().Empty())
	assert.Equal(t, false, SliceV(1).Empty())
	assert.Equal(t, false, SliceV(1, 2, 3).Empty())
	assert.Equal(t, false, Slice(1).Empty())
	assert.Equal(t, false, Slice([]int{1, 2, 3}).Empty())
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
	slice := Slice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func BenchmarkNSlice_First_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleNSlice_First() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.DropFirst().O())
	// Output: [2 3]
}

func TestNSlice_First(t *testing.T) {
	// bool
	{
		slice := SliceV(true, true, false)
		assert.Equal(t, &NObj{true}, slice.First())
		assert.Equal(t, 3, slice.Len())
	}

	// int
	{
		slice := SliceV(1, 2, 3)
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
		slice := SliceV("1", "2", "3")
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
		slice := Slice([]NObj{{1}, {2}, {3}})
		assert.Equal(t, []NObj{{2}, {3}}, slice.DropFirst().O())
		assert.Equal(t, 2, slice.Len())
		assert.Equal(t, []NObj{{3}}, slice.DropFirst().O())
		assert.Equal(t, 1, slice.Len())
		assert.Equal(t, []NObj{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
		assert.Equal(t, []NObj{}, slice.DropFirst().O())
		assert.Equal(t, 0, slice.Len())
	}
}

// // func TestStrSliceFirst(t *testing.T) {
// // 	assert.Equal(t, A(""), S().First())
// // 	assert.Equal(t, A("1"), S("1").First())
// // 	assert.Equal(t, A("1"), S("1", "2").First())
// // 	assert.Equal(t, "foo", A("foo::").Split("::").First().A())
// // 	{
// // 		// Test that the original slice wasn't modified
// // 		q := S("1")
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 		assert.Equal(t, A("1"), q.First())
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 	}
// // }

// // func TestStrSliceJoin(t *testing.T) {
// // 	assert.Equal(t, "", S().Join(".").A())
// // 	assert.Equal(t, "1", S("1").Join(".").A())
// // 	assert.Equal(t, "1.2", S("1", "2").Join(".").A())
// // }

// // func TestStrSliceLen(t *testing.T) {
// // 	assert.Equal(t, 0, S().Len())
// // 	assert.Equal(t, 1, S("1").Len())
// // 	assert.Equal(t, 2, S("1", "2").Len())
// // }

// // func TestStrSliceLast(t *testing.T) {
// // 	assert.Equal(t, A(""), S().Last())
// // 	assert.Equal(t, A("1"), S("1").Last())
// // 	assert.Equal(t, A("2"), S("1", "2").Last())
// // 	assert.Equal(t, "foo", A("::foo").Split("::").Last().A())
// // 	{
// // 		// Test that the original slice wasn't modified
// // 		q := S("1")
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 		assert.Equal(t, A("1"), q.Last())
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 	}
// // }

// Len
//--------------------------------------------------------------------------------------------------
func TestNSlice_Len(t *testing.T) {
	assert.Equal(t, 0, SliceV().Len())
	assert.Equal(t, 1, SliceV().Append("2").Len())
}

// Nil
//--------------------------------------------------------------------------------------------------
func TestNSlice_Nil(t *testing.T) {
	assert.True(t, SliceV().Nil())
	var q *NSlice
	assert.True(t, q.Nil())
	assert.False(t, SliceV().Append("2").Nil())
}

// O
//--------------------------------------------------------------------------------------------------
func TestNSlice_O(t *testing.T) {
	assert.Nil(t, SliceV().O())
	assert.Len(t, SliceV().Append("2").O(), 1)
}

// // func TestStrSlicePrepend(t *testing.T) {
// // 	slice := S().Prepend("1")
// // 	assert.Equal(t, "1", slice.At(0))

// // 	slice.Prepend("2", "3")
// // 	assert.Equal(t, "2", slice.At(0))
// // 	assert.Equal(t, []string{"2", "3", "1"}, slice.S())
// // }

// // func TestStrSliceSort(t *testing.T) {
// // 	slice := S().Append("b", "d", "a")
// // 	assert.Equal(t, []string{"a", "b", "d"}, slice.Sort().S())
// // }

// Set
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Set_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := 0; i < len(ints); i++ {
		ints[i] = 0
	}
}

func BenchmarkNSlice_Set_Optimized(t *testing.B) {
	slice := Slice(Range(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, 0)
	}
}

func BenchmarkNSlice_Set_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines6))
	for i := 0; i < slice.Len(); i++ {
		slice.Set(i, NObj{0})
	}
}

func ExampleNSlice_Set() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.Set(0, 0).O())
	// Output: [0 2 3]
}

func TestNSlice_Set(t *testing.T) {
	// bool
	{
		assert.Equal(t, []bool{false, true, true}, SliceV(true, true, true).Set(0, false).O())
		assert.Equal(t, []bool{true, false, true}, SliceV(true, true, true).Set(1, false).O())
		assert.Equal(t, []bool{true, true, false}, SliceV(true, true, true).Set(2, false).O())
		assert.Equal(t, []bool{false, true, true}, SliceV(true, true, true).Set(-3, false).O())
		assert.Equal(t, []bool{true, false, true}, SliceV(true, true, true).Set(-2, false).O())
		assert.Equal(t, []bool{true, true, false}, SliceV(true, true, true).Set(-1, false).O())
	}

	// int
	{
		assert.Equal(t, []int{0, 2, 3}, SliceV(1, 2, 3).Set(0, 0).O())
		assert.Equal(t, []int{1, 0, 3}, SliceV(1, 2, 3).Set(1, 0).O())
		assert.Equal(t, []int{1, 2, 0}, SliceV(1, 2, 3).Set(2, 0).O())
		assert.Equal(t, []int{0, 2, 3}, SliceV(1, 2, 3).Set(-3, 0).O())
		assert.Equal(t, []int{1, 0, 3}, SliceV(1, 2, 3).Set(-2, 0).O())
		assert.Equal(t, []int{1, 2, 0}, SliceV(1, 2, 3).Set(-1, 0).O())
	}

	// string
	{
		assert.Equal(t, []string{"0", "2", "3"}, SliceV("1", "2", "3").Set(0, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, SliceV("1", "2", "3").Set(1, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, SliceV("1", "2", "3").Set(2, "0").O())
		assert.Equal(t, []string{"0", "2", "3"}, SliceV("1", "2", "3").Set(-3, "0").O())
		assert.Equal(t, []string{"1", "0", "3"}, SliceV("1", "2", "3").Set(-2, "0").O())
		assert.Equal(t, []string{"1", "2", "0"}, SliceV("1", "2", "3").Set(-1, "0").O())
	}

	// custom
	{
		assert.Equal(t, []NObj{{0}, {2}, {3}}, Slice([]NObj{{1}, {2}, {3}}).Set(0, NObj{0}).O())
		assert.Equal(t, []NObj{{1}, {0}, {3}}, Slice([]NObj{{1}, {2}, {3}}).Set(1, NObj{0}).O())
		assert.Equal(t, []NObj{{1}, {2}, {0}}, Slice([]NObj{{1}, {2}, {3}}).Set(2, NObj{0}).O())
		assert.Equal(t, []NObj{{0}, {2}, {3}}, Slice([]NObj{{1}, {2}, {3}}).Set(-3, NObj{0}).O())
		assert.Equal(t, []NObj{{1}, {0}, {3}}, Slice([]NObj{{1}, {2}, {3}}).Set(-2, NObj{0}).O())
		assert.Equal(t, []NObj{{1}, {2}, {0}}, Slice([]NObj{{1}, {2}, {3}}).Set(-1, NObj{0}).O())
	}

	// panics need to run as the last test as they abort the test method
	defer func() {
		err := recover()
		assert.Equal(t, "slice assignment is out of bounds", err)
	}()
	SliceV(1, 2, 3).Set(5, 1)
}

// Slice
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Slice_Normal(t *testing.B) {
	ints := Range(0, nines7)
	for len(ints) > 1 {
		ints = ints[1:]
	}
}

func BenchmarkNSlice_Slice_Optimized(t *testing.B) {
	slice := Slice(Range(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func BenchmarkNSlice_Slice_Reflect(t *testing.B) {
	slice := Slice(rangeNObj(0, nines7))
	for slice.Len() > 0 {
		slice.DropFirst()
	}
}

func ExampleNSlice_Slice() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.DropFirst().O())
	// Output: [2 3]
}

func TestNSlice_Slice(t *testing.T) {

	// nil or empty
	{
		var nilSlice *NSlice
		assert.Equal(t, SliceV(), nilSlice.Slice(0, -1))
		slice := SliceV(0).Clear()
		assert.Equal(t, &NSlice{o: []int{}}, slice.Slice(0, -1))
	}

	// Test that the original is modified when the slice is modified
	{
		original := SliceV(1, 2, 3)
		result := original.Slice(0, -1).Set(0, 0)
		assert.Equal(t, []int{0, 2, 3}, original.O())
		assert.Equal(t, []int{0, 2, 3}, result.O())
	}

	// slice full array
	{
		assert.Equal(t, SliceV(), SliceV().Slice(0, -1))
		assert.Equal(t, SliceV(), SliceV().Slice(0, 1))
		assert.Equal(t, SliceV(), SliceV().Slice(0, 5))
		assert.Equal(t, SliceV(""), SliceV("").Slice(0, -1))
		assert.Equal(t, SliceV(""), SliceV("").Slice(0, 1))
		assert.Equal(t, SliceV(1, 2, 3), SliceV(1, 2, 3).Slice(0, -1))
		assert.Equal(t, Slice([]int{1, 2, 3}), Slice([]int{1, 2, 3}).Slice(0, -1))
		assert.Equal(t, SliceV("1", "2", "3"), SliceV("1", "2", "3").Slice(0, 2))
		assert.Equal(t, Slice([]NObj{{1}, {2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(0, -1))
	}

	// out of bounds should be moved in
	{
		assert.Equal(t, SliceV("1"), SliceV("1").Slice(0, 2))
		assert.Equal(t, SliceV(true, false), SliceV(true, false).Slice(-6, 6))
		assert.Equal(t, SliceV(1, 2, 3), SliceV(1, 2, 3).Slice(-6, 6))
		assert.Equal(t, SliceV("1", "2", "3"), SliceV("1", "2", "3").Slice(-6, 6))
		assert.Equal(t, Slice([]NObj{{1}, {2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-6, 6))
	}

	// mutually exclusive
	{
		slice := SliceV(1, 2, 3, 4)
		assert.Equal(t, &NSlice{}, slice.Slice(2, -3))
		assert.Equal(t, &NSlice{}, slice.Slice(0, -5))
		assert.Equal(t, &NSlice{}, slice.Slice(4, -1))
		assert.Equal(t, &NSlice{}, slice.Slice(6, -1))
		assert.Equal(t, &NSlice{}, slice.Slice(3, 2))
	}

	// singles
	{
		slice := SliceV(1, 2, 3, 4)
		assert.Equal(t, SliceV(4), slice.Slice(-1, -1))
		assert.Equal(t, SliceV(3), slice.Slice(-2, -2))
		assert.Equal(t, SliceV(2), slice.Slice(-3, -3))
		assert.Equal(t, SliceV(1), slice.Slice(0, 0))
		assert.Equal(t, SliceV(1), slice.Slice(-4, -4))
		assert.Equal(t, SliceV(2), slice.Slice(1, 1))
		assert.Equal(t, SliceV(2), slice.Slice(1, -3))
		assert.Equal(t, SliceV(3), slice.Slice(2, 2))
		assert.Equal(t, SliceV(3), slice.Slice(2, -2))
		assert.Equal(t, SliceV(4), slice.Slice(3, 3))
		assert.Equal(t, SliceV(4), slice.Slice(3, -1))
	}

	// grab all but first
	{
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(1, -1))
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(1, 2))
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(-2, -1))
		assert.Equal(t, SliceV(false, true), SliceV(true, false, true).Slice(-2, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(1, -1))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(1, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(-2, -1))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3).Slice(-2, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(1, -1))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(-2, -1))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(-2, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(1, -1))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(1, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-2, -1))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-2, 2))
	}

	// grab all but last
	{
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(0, -2))
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(-3, -2))
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(-3, 1))
		assert.Equal(t, SliceV(true, false), SliceV(true, false, true).Slice(0, 1))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(0, -2))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(-3, -2))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(-3, 1))
		assert.Equal(t, SliceV(1, 2), SliceV(1, 2, 3).Slice(0, 1))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(0, -2))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(-3, -2))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(-3, 1))
		assert.Equal(t, SliceV("1", "2"), SliceV("1", "2", "3").Slice(0, 1))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(0, -2))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-3, -2))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(-3, 1))
		assert.Equal(t, Slice([]NObj{{1}, {2}}), Slice([]NObj{{1}, {2}, {3}}).Slice(0, 1))
	}

	// grab middle
	{
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(1, -2))
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(-3, -2))
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(-3, 2))
		assert.Equal(t, SliceV(true, true), SliceV(false, true, true, false).Slice(1, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(1, -2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(-3, -2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(-3, 2))
		assert.Equal(t, SliceV(2, 3), SliceV(1, 2, 3, 4).Slice(1, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(1, -2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(-3, -2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(-3, 2))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3", "4").Slice(1, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(1, -2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(-3, -2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(-3, 2))
		assert.Equal(t, Slice([]NObj{{2}, {3}}), Slice([]NObj{{1}, {2}, {3}, {4}}).Slice(1, 2))
	}

	// random
	{
		assert.Equal(t, SliceV("1"), SliceV("1", "2", "3").Slice(0, -3))
		assert.Equal(t, SliceV("2", "3"), SliceV("1", "2", "3").Slice(1, 2))
		assert.Equal(t, SliceV("1", "2", "3"), SliceV("1", "2", "3").Slice(0, 2))
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
