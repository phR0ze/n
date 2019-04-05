package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	OldSlice(rangeInterObject(0, nines6)).Each(func(x O) {
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
	OldSlice(rangeInterObject(0, nines6)).Each(func(x O) {
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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	slice := OldSlice(rangeInterObject(0, nines6))
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
	slice := OldSlice(rangeInterObject(0, nines6))
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
	slice := OldSlice(rangeInterObject(0, nines7))
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
	slice := OldSlice(rangeInterObject(0, nines6))
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
