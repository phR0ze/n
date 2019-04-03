package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewRefSliceV function
//--------------------------------------------------------------------------------------------------
func ExampleNewRefSliceV_empty() {
	slice := OldSliceV()
	fmt.Println(slice.O())
	// Output: <nil>
}

func ExampleNewRefSliceV_variadic() {
	slice := OldSliceV(1, 2, 3)
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
		assert.True(t, !OldSliceV().Any())
		assert.Equal(t, 0, OldSliceV().Len())
		assert.Equal(t, nil, OldSliceV().O())
		assert.Equal(t, nil, OldSliceV(nil).O())
		assert.Equal(t, &NSlice{}, OldSliceV(nil))
		assert.Equal(t, []string{""}, OldSliceV(nil, "").O())
		assert.Equal(t, []*Object{nil}, OldSliceV(nil, obj).O())
	}

	// Test pointers
	{
		assert.Equal(t, []*Object{nil}, OldSliceV(obj).O())
		assert.Equal(t, []*Object{&(Object{"bob"})}, OldSliceV(&(Object{"bob"})).O())
		assert.Equal(t, []*Object{nil}, OldSliceV(obj).O())
		assert.Equal(t, []*Object{&(Object{"bob"})}, OldSliceV(&(Object{"bob"})).O())
		assert.Equal(t, [][]*Object{{&(Object{"1"}), &(Object{"2"})}}, OldSliceV([]*Object{&(Object{"1"}), &(Object{"2"})}).O())
	}

	// Singles
	{
		assert.Equal(t, []int{1}, OldSliceV(1).O())
		assert.Equal(t, []string{"1"}, OldSliceV("1").O())
		assert.Equal(t, []Object{Object{"bob"}}, OldSliceV(Object{"bob"}).O())
		assert.Equal(t, []map[string]string{{"1": "one"}}, OldSliceV(map[string]string{"1": "one"}).O())
	}

	// Multiples
	{
		assert.Equal(t, []int{1, 2}, OldSliceV(1, 2).O())
		assert.Equal(t, []string{"1", "2"}, OldSliceV("1", "2").O())
		assert.Equal(t, []Object{Object{1}, Object{2}}, OldSliceV(Object{1}, Object{2}).O())
	}

	// Test slices
	{
		assert.Equal(t, [][]int{{1, 2}}, OldSliceV([]int{1, 2}).O())
		assert.Equal(t, [][]string{{"1"}}, OldSliceV([]string{"1"}).O())
	}
}
