package n

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Integer struct{ Value int }

func RangeInteger(min, max int) []Integer {
	result := make([]Integer, max-min+1)
	for i := range result {
		result[i] = Integer{min + i}
	}
	return result
}

func BenchmarkSlice_CustomNative(t *testing.B) {
	seedData := RangeInteger(0, 999999)

	// Select the actual values out of the custom object
	lambda := func(x Integer) int {
		return x.Value
	}

	// Because we are assuming the developer isn't reusing anything we know the types
	// and can use them directly
	ints := []int{}
	for i := range seedData {
		ints = append(ints, lambda(seedData[i]))
	}

	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}

func BenchmarkSlice_PureReflection(t *testing.B) {
	seedData := RangeInteger(0, 999999)

	// Select the actual values out of the custom object
	lambda := func(x interface{}) interface{} {
		return x.(Integer).Value
	}

	// Use reflection to determine the Kind this would model creating a reference
	// to a new type e.g. NewRefSlice(seed). Because we used reflection we don't
	// need to convert the seed data we can just use the given slice object the
	// way it is.
	var v2 reflect.Value
	v1 := reflect.ValueOf(seedData)
	k1 := v1.Kind()

	// We do need to validate that is a slice type before working with it
	if k1 == reflect.Slice {
		for i := 0; i < v1.Len(); i++ {
			result := lambda(v1.Index(i).Interface())

			// We need to create a new slice based on the result type to store all the results
			if !v2.IsValid() {
				typ := reflect.SliceOf(reflect.TypeOf(result))
				v2 = reflect.MakeSlice(typ, 0, v1.Len())
			}

			// Appending the native type to a reflect.Value of slice is more complicated as we
			// now have to convert the result type into a reflect.Value before appending
			// type asssert the native type.
			v2 = reflect.Append(v2, reflect.ValueOf(result))
		}
	}

	// Now convert the results into a native type
	// Because we created the native slice type as v2 we can simply get its interface and cast
	ints := v2.Interface().([]int)
	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}

func BenchmarkSlice_SliceOfInterface(t *testing.B) {
	seedData := RangeInteger(0, 999999)

	// Select the actual values out of the custom object
	lambda := func(x interface{}) interface{} {
		return x.(Integer).Value
	}

	// Use reflection to determine the Kind this would model creating a reference
	// to a new type e.g. NewSlice(seed). Once we've done that we need to convert
	// it into a []interface{}
	v1 := reflect.ValueOf(seedData)
	k1 := v1.Kind()

	// We do need to validate that is a slice type before working with it
	g1 := []interface{}{}
	if k1 == reflect.Slice {
		for i := 0; i < v1.Len(); i++ {
			g1 = append(g1, v1.Index(i).Interface())
		}
	}

	// Now iterate and execute the lambda and create the new result
	results := []interface{}{}
	for i := range g1 {
		results = append(results, lambda(g1[i]))
	}

	// Now convert the results into a native type
	ints := []int{}
	for i := range results {
		ints = append(ints, results[i].(int))
	}
	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}

func BenchmarkSlice_IntSlice(t *testing.B) {
	ints := NewSlice(Range(0, 999999)).Map(func(x O) O {
		return x.(int) + 1
	}).ToInts()
	assert.Equal(t, 3, ints[2])
	assert.Equal(t, 100000, ints[99999])
}

func BenchmarkSlice_RefSlice(t *testing.B) {
	ints := NewSlice(RangeInteger(0, 999999)).Map(func(x O) O {
		return x.(Integer).Value
	}).ToInts()
	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}

func TestSlice_NewSlice(t *testing.T) {

	// float
	{
		assert.Equal(t, NewFloatSliceV(3), NewSlice([]float32{3}))
		assert.Equal(t, NewFloatSliceV(3), NewSlice([]float64{3}))
	}

	// int
	{
		assert.Equal(t, NewIntSliceV(3), NewSlice([]int{3}))
		assert.Equal(t, NewIntSliceV(3), NewSlice(&([]int{3})))
		assert.Equal(t, NewIntSliceV(3), NewSlice([]int8{3}))
		assert.Equal(t, NewIntSliceV(3), NewSlice([]int16{3}))
		assert.Equal(t, NewIntSliceV(3), NewSlice([]int64{3}))
		assert.Equal(t, NewIntSliceV(3), NewSlice([]uint16{3}))
		assert.Equal(t, NewIntSliceV(3), NewSlice([]uint32{3}))
		assert.Equal(t, NewIntSliceV(3), NewSlice([]uint64{3}))
	}

	// string
	{
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewSlice([]string{"1", "2", "3"}))
		assert.Equal(t, NewStringSliceV("1", "2", "3"), NewSlice(&([]string{"1", "2", "3"})))
		assert.Equal(t, NewStringSliceV("3"), NewSlice([][]byte{{'3'}}))
		assert.Equal(t, NewStringSliceV("3"), NewSlice([][]rune{{'3'}}))
	}

	// Str
	{
		assert.Equal(t, NewStrV("3"), NewSlice([]rune{'3'}))
		assert.Equal(t, NewStrV("3"), NewSlice([]byte{'3'}))
		assert.Equal(t, NewStrV("3"), NewSlice([]Char{'3'}))
	}
}

func TestSlice_NewSliceV(t *testing.T) {

	// int
	{
		assert.Equal(t, NewFloatSliceV(3), NewSliceV(float32(3.0)))
		assert.Equal(t, NewFloatSliceV(3), NewSliceV(float64(3.0)))
		assert.Equal(t, NewIntSliceV(3), NewSliceV(3))
		assert.Equal(t, NewIntSliceV(3), NewSliceV(int8(3)))
		assert.Equal(t, NewIntSliceV(3), NewSliceV(int16(3)))
		assert.Equal(t, NewIntSliceV(3), NewSliceV(int64(3)))
		assert.Equal(t, NewIntSliceV(3), NewSliceV(uint16(3)))
		assert.Equal(t, NewIntSliceV(3), NewSliceV(uint32(3)))
		assert.Equal(t, NewIntSliceV(3), NewSliceV(uint64(3)))
	}

	// string
	{
		assert.Equal(t, NewStringSliceV("3"), NewSliceV("3"))
		assert.Equal(t, NewStringSliceV("3"), NewSliceV([]byte{'3'}))
		assert.Equal(t, NewStringSliceV("3"), NewSliceV([]rune{'3'}))
		assert.Equal(t, NewStringSliceV("3"), NewSliceV(Str{'3'}))
	}

	// Str
	{
		assert.Equal(t, NewStrV("3"), NewSliceV('3'))
		assert.Equal(t, NewStrV("3"), NewSliceV(byte('3')))
		assert.Equal(t, NewStrV("3"), NewSliceV(Char('3')))
	}
}

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
