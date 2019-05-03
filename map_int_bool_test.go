package n

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// // NewIntMapBool
// //--------------------------------------------------------------------------------------------------
// func BenchmarkNewIntMap_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i += 10 {
// 		_ = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9}
// 	}
// }

// func BenchmarkNewIntSlice_Slice(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i += 10 {
// 		_ = NewIntSlice([]int{i, i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9})
// 	}
// }

// func ExampleNewIntSlice() {
// 	slice := NewIntSlice([]int{1, 2, 3})
// 	fmt.Println(slice.O())
// 	// Output: [1 2 3]
// }

// func TestIntSlice_NewIntSlice(t *testing.T) {

// 	// array
// 	var array [2]int
// 	array[0] = 1
// 	array[1] = 2
// 	assert.Equal(t, []int{1, 2}, NewIntSlice(array[:]).O())

// 	// empty
// 	assert.Equal(t, []int{}, NewIntSlice([]int{}).O())

// 	// slice
// 	assert.Equal(t, []int{0}, NewIntSlice([]int{0}).O())
// 	assert.Equal(t, []int{1, 2}, NewIntSlice([]int{1, 2}).O())
// }

// // NewIntSliceV
// //--------------------------------------------------------------------------------------------------
// func BenchmarkNewIntSliceV_Go(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i += 10 {
// 		_ = append([]int{}, i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9)
// 	}
// }

// func BenchmarkNewIntSliceV_Slice(t *testing.B) {
// 	ints := Range(0, nines6)
// 	for i := 0; i < len(ints); i += 10 {
// 		_ = NewIntSliceV(i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9)
// 	}
// }

// func ExampleNewIntSliceV_empty() {
// 	slice := NewIntSliceV()
// 	fmt.Println(slice.O())
// 	// Output: []
// }

// func ExampleNewIntSliceV_variadic() {
// 	slice := NewIntSliceV(1, 2, 3)
// 	fmt.Println(slice.O())
// 	// Output: [1 2 3]
// }

// func TestIntSlice_NewIntSliceV(t *testing.T) {

// 	// array
// 	var array [2]int
// 	array[0] = 1
// 	array[1] = 2
// 	assert.Equal(t, []int{1, 2}, NewIntSliceV(array[:]...).O())

// 	// empty
// 	assert.Equal(t, []int{}, NewIntSliceV().O())

// 	// multiples
// 	assert.Equal(t, []int{1}, NewIntSliceV(1).O())
// 	assert.Equal(t, []int{1, 2}, NewIntSliceV(1, 2).O())
// 	assert.Equal(t, []int{1, 2}, NewIntSliceV([]int{1, 2}...).O())
// }

// // Any
// //--------------------------------------------------------------------------------------------------
// func BenchmarkIntSlice_Any_Go(t *testing.B) {
// 	any := func(list []int, x []int) bool {
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
// 	ints := Range(0, nines4)
// 	for i := range ints {
// 		any(ints, []int{i})
// 	}
// }

// func BenchmarkIntSlice_Any_Slice(t *testing.B) {
// 	src := Range(0, nines4)
// 	slice := NewIntSlice(src)
// 	for i := range src {
// 		slice.Any(i)
// 	}
// }

// func ExampleIntSlice_Any_empty() {
// 	slice := NewIntSliceV()
// 	fmt.Println(slice.Any())
// 	// Output: false
// }

// func ExampleIntSlice_Any_notEmpty() {
// 	slice := NewIntSliceV(1, 2, 3)
// 	fmt.Println(slice.Any())
// 	// Output: true
// }

// func ExampleIntSlice_Any_contains() {
// 	slice := NewIntSliceV(1, 2, 3)
// 	fmt.Println(slice.Any(1))
// 	// Output: true
// }

// func ExampleIntSlice_Any_containsAny() {
// 	slice := NewIntSliceV(1, 2, 3)
// 	fmt.Println(slice.Any(0, 1))
// 	// Output: true
// }

// func TestIntSlice_Any(t *testing.T) {

// 	// empty
// 	var nilSlice *IntSlice
// 	assert.False(t, nilSlice.Any())
// 	assert.False(t, NewIntSliceV().Any())

// 	// single
// 	assert.True(t, NewIntSliceV(2).Any())

// 	// invalid
// 	assert.False(t, NewIntSliceV(1, 2).Any(Object{2}))

// 	assert.True(t, NewIntSliceV(1, 2, 3).Any(2))
// 	assert.False(t, NewIntSliceV(1, 2, 3).Any(4))
// 	assert.True(t, NewIntSliceV(1, 2, 3).Any(4, 3))
// 	assert.False(t, NewIntSliceV(1, 2, 3).Any(4, 5))
// }
