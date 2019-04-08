package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	src := rangeInterObject(0, nines5)
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
