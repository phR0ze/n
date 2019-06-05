package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewStringMap
//--------------------------------------------------------------------------------------------------
func ExampleNewStringMap() {
	fmt.Println(NewStringMap(map[string]interface{}{"k": "v"}))
	// Output: &map[k:v]
}

func TestNewStringMap(t *testing.T) {

	// map[string]interface
	{
		m := map[string]interface{}{"k": "v"}
		assert.Equal(t, NewStringMap(m), NewStringMap(m))
	}

	// StringMap
	{
		m := NewStringMap()
		m.Set("k", "v")
		assert.Equal(t, m, NewMap(m))
	}
}

// Any
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Any() {
	m := NewStringMap(map[string]interface{}{"k": "v"})
	fmt.Println(m.Any())
	// Output: true
}

func TestStringMap_Any(t *testing.T) {
	// Not empty
	{
		assert.Equal(t, false, NewStringMap().Any())
		assert.Equal(t, true, NewStringMap().SetM("1", "one").Any())
	}

	// Specific keys
	{
		assert.Equal(t, false, NewStringMap().SetM("1", "one").Any("2"))
		assert.Equal(t, true, NewStringMap().SetM("1", "one").Any("1"))
	}
}

// Clear
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Clear() {
	m := NewStringMap(map[string]interface{}{"1": "one"})
	fmt.Println(m.Clear())
	// Output: &map[]
}

func TestStringMap_Clear(t *testing.T) {

	// nil
	{
		assert.Equal(t, NewStringMap(), (*StringMap)(nil).Clear())
	}

	// empty
	{
		m := NewStringMap()
		assert.Equal(t, NewStringMap(), m.Clear())
		assert.Equal(t, 0, m.Len())
	}

	// one
	{
		m := NewStringMap(map[string]interface{}{"1": "one"})
		assert.Equal(t, NewStringMap(), m.Clear())
		assert.Equal(t, 0, m.Len())
	}

	// many
	{
		m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.Equal(t, NewStringMap(), m.Clear())
		assert.Equal(t, 0, m.Len())
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Copy() {
	m := NewStringMap(map[string]interface{}{"1": "one", "2": "two"})
	fmt.Println(m.Copy("1"))
	// Output: &map[1:one]
}

func TestStringMap_Copy(t *testing.T) {
	// copy all
	{
		m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		copy := m.Copy()
		assert.Equal(t, m, copy)
		copy.Set("1", "foo")
		assert.Equal(t, map[string]interface{}{"1": "one", "2": "two", "3": "three"}, m.O())
		assert.Equal(t, map[string]interface{}{"1": "foo", "2": "two", "3": "three"}, copy.O())
	}

	// Copy specific keys
	{
		m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		copy := m.Copy("2", "3")
		assert.Equal(t, map[string]interface{}{"1": "one", "2": "two", "3": "three"}, m.O())
		assert.Equal(t, map[string]interface{}{"2": "two", "3": "three"}, copy.O())
	}
}

// Delete
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Delete() {
	m := NewStringMap(map[string]interface{}{"1": "one"})
	fmt.Println(m.Delete("1").O())
	// Output: one
}

func TestStringMap_Delete(t *testing.T) {
	m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
	assert.Equal(t, 3, m.Len())

	// Non existant key
	assert.Equal(t, nil, m.Delete("0").O())
	assert.Equal(t, 3, m.Len())

	// First key
	assert.Equal(t, "one", m.Delete("1").O())
	assert.Equal(t, 2, m.Len())

	// Second key
	assert.Equal(t, "two", m.Delete("2").O())
	assert.Equal(t, 1, m.Len())

	// Last key
	assert.Equal(t, "three", m.Delete("3").O())
	assert.Equal(t, 0, m.Len())

	// Non existant key
	assert.Equal(t, nil, m.Delete("3").O())
	assert.Equal(t, 0, m.Len())
}

// DeleteM
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_DeleteM() {
	m := NewStringMap(map[string]interface{}{"1": "one"})
	fmt.Println(m.DeleteM("1"))
	// Output: &map[]
}

func TestStringMap_DeleteM(t *testing.T) {
	m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
	assert.Equal(t, 3, m.Len())

	// Non existant key
	assert.Equal(t, 3, m.DeleteM("0").Len())

	// First key
	assert.Equal(t, 2, m.DeleteM("1").Len())

	// Second key
	assert.Equal(t, 1, m.DeleteM("2").Len())

	// Last key
	assert.Equal(t, 0, m.DeleteM("3").Len())

	// Non existant key
	assert.Equal(t, 0, m.DeleteM("3").Len())
}

// Exists
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Exists() {
	m := NewStringMap(map[string]interface{}{"1": "one"})
	fmt.Println(m.Exists("1"))
	// Output: true
}

func TestStringMap_Exists(t *testing.T) {
	// bool
	{
		m := NewStringMap(map[string]interface{}{"1": true})

		// Non existant key
		assert.Equal(t, false, m.Exists("0"))

		// exists
		assert.Equal(t, true, m.Exists("1"))
	}

	// string
	{
		// none
		assert.Equal(t, false, NewStringMap().Exists("0"))

		m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})

		// Non existant key
		assert.Equal(t, false, m.Exists("0"))

		// First key
		assert.Equal(t, true, m.Exists("1"))

		// Second key
		assert.Equal(t, true, m.Exists("2"))

		// Last key
		assert.Equal(t, true, m.Exists("3"))
	}
}

// Generic
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Generic() {
	fmt.Println(NewStringMap().Generic())
	// Output: false
}

func TestStringMap_Generic(t *testing.T) {
	assert.Equal(t, false, NewStringMap().Generic())
}

// Get
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Get() {
	m := NewStringMap(map[string]interface{}{"1": "one"})
	fmt.Println(m.Get("1").O())
	// Output: one
}

func TestStringMap_Get(t *testing.T) {
	m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
	assert.Equal(t, 3, m.Len())

	// Non existant key
	assert.Equal(t, nil, m.Get("0").O())
	assert.Equal(t, 3, m.Len())

	// First key
	assert.Equal(t, "one", m.Get("1").O())
	assert.Equal(t, 3, m.Len())

	// Second key
	assert.Equal(t, "two", m.Get("2").O())
	assert.Equal(t, 3, m.Len())

	// Last key
	assert.Equal(t, "three", m.Get("3").O())
	assert.Equal(t, 3, m.Len())

	// Try a key again to make sure its still there
	assert.Equal(t, "three", m.Get("3").O())
	assert.Equal(t, 3, m.Len())
}

// Keys
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Keys() {
	m := NewStringMap(map[string]interface{}{"1": "one"})
	fmt.Println(m.Keys().O())
	// Output: [1]
}

func TestStringMap_Keys(t *testing.T) {
	// nil or empty
	{
		assert.Equal(t, []string{}, (*StringMap)(nil).Keys().O())
		assert.Equal(t, []string{}, NewStringMap().Keys().O())
	}

	// many
	{
		m := NewStringMap(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		keys := m.Keys().O()
		assert.Len(t, keys, 3)
		assert.Contains(t, keys, "1")
		assert.Contains(t, keys, "2")
		assert.Contains(t, keys, "3")
	}
}

// Len
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Len() {
	fmt.Println(NewStringMap().SetM("k", "v").Len())
	// Output: 1
}

func TestStringMap_Len(t *testing.T) {
	m := NewStringMap()
	assert.Equal(t, 0, m.Len())
	assert.Equal(t, 1, m.SetM("1", "one").Len())
	assert.Equal(t, 2, m.SetM("2", "two").Len())
	assert.Equal(t, 1, m.DeleteM("2").Len())
	assert.Equal(t, 0, m.DeleteM("1").Len())
}

// Merge
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Merge() {
	fmt.Println(NewStringMap(map[string]interface{}{"1": "two"}).Merge(NewStringMap(map[string]interface{}{"1": "one"})))
	// Output: &map[1:one]
}

func TestStringMap_Merge(t *testing.T) {
	// nil or empty
	{
		assert.Equal(t, NewStringMap(), (*StringMap)(nil).Merge(nil))
		assert.Equal(t, NewStringMap(), NewStringMap().Merge(nil))
		assert.Equal(t, map[string]interface{}{}, NewStringMap().Merge(nil).O())
		assert.Equal(t, NewStringMap(), NewStringMap().Merge(NewStringMap()))
		assert.Equal(t, map[string]interface{}{}, NewStringMap().Merge(NewStringMap()).O())
	}
	{
		a := NewStringMap(map[string]interface{}{})
		b := NewStringMap(map[string]interface{}{"1": "one"})
		assert.Equal(t, b, a.Merge(b))
	}
	{
		a := NewStringMap(map[string]interface{}{"1": "one"})
		b := NewStringMap(map[string]interface{}{})
		assert.Equal(t, a, a.Merge(b))
	}
	{
		a := NewStringMap(map[string]interface{}{
			"1": "one",
		})
		b := NewStringMap(map[string]interface{}{
			"2": "two",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
		})
		assert.Equal(t, expected, a.Merge(b))
	}
	{
		// Override string in a with string in b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "2",
		})
		b := NewStringMap(map[string]interface{}{
			"2": "two",
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		})
		assert.Equal(t, expected, a.Merge(b))
	}
	{
		// Override string in a with map from b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
		})
		b := NewStringMap(map[string]interface{}{
			"2": NewStringMap(map[string]interface{}{"foo": "bar"}),
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": NewStringMap(map[string]interface{}{"foo": "bar"}),
			"3": "three",
		})
		assert.Equal(t, expected, a.Merge(b))
	}
	{
		// Override map in a with string from b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": NewStringMap(map[string]interface{}{"foo": "bar"}),
		})
		b := NewStringMap(map[string]interface{}{
			"2": "two",
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		})
		assert.Equal(t, expected, a.Merge(b))
	}
	{
		// Override sub map string in a with sub map string from b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": NewStringMap(map[string]interface{}{"foo": "bar1"}),
		})
		b := NewStringMap(map[string]interface{}{
			"2": NewStringMap(map[string]interface{}{
				"foo":  "bar2",
				"foo2": "bar2",
			}),
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": NewStringMap(map[string]interface{}{
				"foo":  "bar2",
				"foo2": "bar2",
			}),
			"3": "three",
		})
		assert.Equal(t, expected, a.Merge(b))
	}
}

// MergeG
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_MergeG() {
	fmt.Println(NewStringMap(map[string]interface{}{"1": "two"}).MergeG(NewStringMap(map[string]interface{}{"1": "one"})))
	// Output: map[1:one]
}

func TestStringMap_MergeG(t *testing.T) {
	// nil or empty
	{
		assert.Equal(t, map[string]interface{}{}, NewStringMap().MergeG(nil))
		assert.Equal(t, map[string]interface{}{}, NewStringMap().MergeG(NewStringMap()))
	}
	{
		a := NewStringMap(map[string]interface{}{})
		b := NewStringMap(map[string]interface{}{"1": "one"})
		assert.Equal(t, b.G(), a.MergeG(b))
	}
	{
		a := NewStringMap(map[string]interface{}{"1": "one"})
		b := NewStringMap(map[string]interface{}{})
		assert.Equal(t, a.G(), a.MergeG(b))
	}
	{
		a := NewStringMap(map[string]interface{}{
			"1": "one",
		})
		b := NewStringMap(map[string]interface{}{
			"2": "two",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
		})
		assert.Equal(t, expected.G(), a.MergeG(b))
	}
	{
		// Override string in a with string in b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "2",
		})
		b := NewStringMap(map[string]interface{}{
			"2": "two",
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		})
		assert.Equal(t, expected.G(), a.MergeG(b))
	}
	{
		// Override string in a with map from b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
		})
		b := NewStringMap(map[string]interface{}{
			"2": NewStringMap(map[string]interface{}{"foo": "bar"}),
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{"foo": "bar"},
			"3": "three",
		})
		assert.Equal(t, expected.G(), a.MergeG(b))
	}
	{
		// Override map in a with string from b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": NewStringMap(map[string]interface{}{"foo": "bar"}),
		})
		b := NewStringMap(map[string]interface{}{
			"2": "two",
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		})
		assert.Equal(t, expected.G(), a.MergeG(b))
	}
	{
		// Override sub map string in a with sub map string from b
		a := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": NewStringMap(map[string]interface{}{"foo": "bar1"}),
		})
		b := NewStringMap(map[string]interface{}{
			"2": NewStringMap(map[string]interface{}{
				"foo":  "bar2",
				"foo2": "bar2",
			}),
			"3": "three",
		})
		expected := NewStringMap(map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"foo":  "bar2",
				"foo2": "bar2",
			},
			"3": "three",
		})
		assert.Equal(t, expected.G(), a.MergeG(b))
		assert.Equal(t, map[string]interface{}{"foo": "bar2", "foo2": "bar2"}, expected.G()["2"])
	}
}

// O
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_O() {
	fmt.Println(NewStringMap(map[string]interface{}{"1": "one"}).O())
	// Output: map[1:one]
}

func TestStringMap_O(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, (*StringMap)(nil).O())
	assert.Equal(t, map[string]interface{}{}, NewStringMap().O())
	assert.Equal(t, map[string]interface{}{"1": "one"}, NewStringMap(map[string]interface{}{"1": "one"}).O())
}

// Set
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Set() {
	fmt.Println(NewStringMap().Set("key", "value"))
	// Output: true
}

func TestStringMap_Set(t *testing.T) {
	// bool
	{
		m := NewStringMap()
		assert.Equal(t, true, m.Set("test", false))
		assert.Equal(t, false, m.Set("test", true))
		assert.Equal(t, true, (*m)["test"].(bool))
	}

	// string
	{
		m := NewStringMap()
		assert.Equal(t, true, m.Set("1", "one"))
		assert.Equal(t, false, m.Set("1", "two"))
		assert.Equal(t, "two", (*m)["1"].(string))
	}
}

// SetM
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_SetM() {
	fmt.Println(NewStringMap().SetM("k", "v"))
	// Output: &map[k:v]
}

func TestStringMap_SetM(t *testing.T) {
	// bool
	{
		m := NewStringMap()
		assert.Equal(t, NewStringMap(map[string]interface{}{"test": false}), m.SetM("test", false))
		assert.Equal(t, NewStringMap(map[string]interface{}{"test": true}), m.SetM("test", true))
		assert.Equal(t, true, (*m)["test"].(bool))
	}

	// string
	{
		m := NewStringMap()
		assert.Equal(t, NewStringMap(map[string]interface{}{"1": "one"}), m.SetM("1", "one"))
		assert.Equal(t, NewStringMap(map[string]interface{}{"1": "two"}), m.SetM("1", "two"))
		assert.Equal(t, "two", (*m)["1"].(string))
	}
}

// Yaml
//--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Yaml() {
// 	m := NewStringMap(map[string]interface{}{"1": map[string]interface{}{"2": "two"}})
// 	fmt.Println(m.Yaml("1.2").O())
// 	// Output: two
// }

func TestStringMap_Yaml(t *testing.T) {

	// dot notation
	{
		// Identity: .
		assert.Equal(t, &StringMap{"one": "1"}, NewStringMap(map[string]interface{}{"one": "1"}).Query(``).ToStringMap())
		assert.Equal(t, &StringMap{"one": "1"}, NewStringMap(map[string]interface{}{"one": "1"}).Query(`.`).ToStringMap())

		// Object Identifier-Index: .foo, .foo.bar
		assert.Equal(t, "1", NewStringMap(map[string]interface{}{"one": "1"}).Query(`one`).ToString())
		assert.Equal(t, "1", NewStringMap(map[string]interface{}{"one": "1"}).Query(`.one`).ToString())
		assert.Equal(t, "1", NewStringMap(map[string]interface{}{"one": "1"}).Query(`."one"`).ToString())
		assert.Equal(t, "2", NewStringMap(map[string]interface{}{"one": map[string]interface{}{"two": "2"}}).Query(`one.two`).ToString())
		assert.Equal(t, "3", NewStringMap(map[string]interface{}{"one": map[string]interface{}{"two": map[string]interface{}{"three": "3"}}}).Query(`one.two.three`).ToString())
		assert.Equal(t, "foo", NewStringMap(map[string]interface{}{"one": map[string]interface{}{"two.three": "foo"}}).Query(`one."two.three"`).ToString())

		// Array Index: .[2], .[-1]
		assert.Equal(t, 0, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[0]`).ToInt())
		assert.Equal(t, 1, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[1]`).ToInt())
		assert.Equal(t, 2, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[2]`).ToInt())
		assert.Equal(t, 2, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[-1]`).ToInt())
		assert.Equal(t, 1, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[-2]`).ToInt())
		assert.Equal(t, 0, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[-3]`).ToInt())
		assert.Equal(t, float32(2.2), ToStringMap("foo:\n  - 0.5\n  - 1.1\n  - 2.2\n").Query(`foo.[-1]`).ToFloat32())
		assert.Equal(t, float64(2.2), ToStringMap("foo:\n  - 0.5\n  - 1.1\n  - 2.2\n").Query(`foo.[-1]`).ToFloat64())
		assert.Equal(t, "0", NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[0]`).ToString())
		assert.Equal(t, "1", NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[1]`).ToString())
		assert.Equal(t, "2", NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[2]`).ToString())
		assert.Equal(t, "2", NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-1]`).ToString())
		assert.Equal(t, "1", NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-2]`).ToString())
		assert.Equal(t, "0", NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-3]`).ToString())
		assert.Equal(t, "", NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-5]`).ToString())

		// Array Interator: .[]
		assert.Equal(t, []int{0, 1, 2}, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[]`).ToIntSliceG())
		assert.Equal(t, []int{0, 1, 2}, NewStringMap(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[]`).ToIntSliceG())

		// Array Slice: .[1:]
	}

	// no dot notation
	{
		// floats
		assert.Equal(t, float32(1.2), NewStringMap(map[string]interface{}{"1": 1.2}).Query("1").ToFloat32())
		assert.Equal(t, float64(1.2), NewStringMap(map[string]interface{}{"1": 1.2}).Query("1").ToFloat64())

		// ints
		assert.Equal(t, int(1), NewStringMap(map[string]interface{}{"1": 1}).Query("1").ToInt())
		assert.Equal(t, int8(1), NewStringMap(map[string]interface{}{"1": 1}).Query("1").ToInt8())
		assert.Equal(t, int16(1), NewStringMap(map[string]interface{}{"1": 1}).Query("1").ToInt16())
		assert.Equal(t, int32(1), NewStringMap(map[string]interface{}{"1": 1}).Query("1").ToInt32())
		assert.Equal(t, int64(1), NewStringMap(map[string]interface{}{"1": 1}).Query("1").ToInt64())
		assert.Equal(t, uint(1), NewStringMap(map[string]interface{}{"1": 1}).Query("1").ToUint())

		// string
		assert.Equal(t, "one", NewStringMap(map[string]interface{}{"1": "one"}).Query("1").ToString())

		// maps
		assert.Equal(t, &StringMap{"2": "two"}, NewStringMap(map[string]interface{}{"1": map[string]interface{}{"2": "two"}}).Query("1").ToStringMap())
		assert.Equal(t, map[string]interface{}{"2": "two"}, NewStringMap(map[string]interface{}{"1": map[string]interface{}{"2": "two"}}).Query("1").ToStringMapG())

		// yaml
		assert.Equal(t, "1", ToStringMap("one: 1").Query("one").ToString())
		assert.Equal(t, []int{1, 2}, ToStringMap("foo: \n  - 1\n  - 2").Query("foo").ToIntSliceG())
		assert.Equal(t, &IntSlice{1, 2}, ToStringMap("foo: \n  - 1\n  - 2").Query("foo").ToIntSlice())
	}
}
