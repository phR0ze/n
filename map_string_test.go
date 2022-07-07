package n

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

// Benchmarks
//--------------------------------------------------------------------------------------------------

// The intent of this benchmark is to determine how long it takes to insert one int into a []interface{}{}
// by converting to []int and inserting than back to []interface{}{}
func BenchmarkStringMap_SetValueToFromInterfaceSlice(t *testing.B) {

	// Create generic slice of ints
	slice := []interface{}{}
	for i := range Range(0, nines6) {
		slice = append(slice, i)
	}

	// Convert to ints
	s := Slice(slice)
	s.Set(1000, 5)

	// Convert back to interface
	result := s.ToInterSlice()
	assert.Equal(t, 5, result[1000])
}

func BenchmarkStringMap_SetValueRefSlice(t *testing.B) {

	// Create generic slice of ints
	slice := []interface{}{}
	for i := range Range(0, nines6) {
		slice = append(slice, i)
	}

	// Set value
	s := NewRefSlice(slice)
	s.Set(1000, 5)

	// Convert back to interface
	result := s.ToInterSlice()
	assert.Equal(t, 5, result[1000])
}

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
		assert.Equal(t, M().Add("k", "v"), NewStringMap(m))
	}

	// StringMap
	{
		m := NewStringMapV()
		m.Set("k", "v")
		assert.Equal(t, m, Map(m))
	}
}

// NewStringMapV
//--------------------------------------------------------------------------------------------------
func ExampleNewStringMapV() {
	fmt.Println(NewStringMapV(map[string]interface{}{"k": "v"}))
	// Output: &[{k v}]
}

func TestNewStringMapV(t *testing.T) {

	// map[string]interface
	{
		m := map[string]interface{}{"k": "v"}
		assert.Equal(t, NewStringMapV(m), NewStringMapV(m))
	}

	// StringMap
	{
		m := NewStringMapV()
		m.Set("k", "v")
		assert.Equal(t, m, Map(m))
	}
}

// Any
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Any() {
	m := NewStringMapV(map[string]interface{}{"k": "v"})
	fmt.Println(m.Any())
	// Output: true
}

func TestStringMap_Any(t *testing.T) {
	// Not empty
	{
		assert.Equal(t, false, NewStringMapV().Any())
		assert.Equal(t, true, NewStringMapV().SetM("1", "one").Any())
	}

	// Specific keys
	{
		assert.Equal(t, false, NewStringMapV().SetM("1", "one").Any("2"))
		assert.Equal(t, true, NewStringMapV().SetM("1", "one").Any("1"))
	}
}

// Clear
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Clear() {
	m := NewStringMapV(map[string]interface{}{"1": "one"})
	fmt.Println(m.Clear())
	// Output: &[]
}

func TestStringMap_Clear(t *testing.T) {

	// nil
	{
		assert.Equal(t, NewStringMapV(), (*StringMap)(nil).Clear())
	}

	// empty
	{
		m := NewStringMapV()
		assert.Equal(t, NewStringMapV(), m.Clear())
		assert.Equal(t, 0, m.Len())
	}

	// one
	{
		m := NewStringMapV(map[string]interface{}{"1": "one"})
		assert.Equal(t, NewStringMapV(), m.Clear())
		assert.Equal(t, 0, m.Len())
	}

	// many
	{
		m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.Equal(t, NewStringMapV(), m.Clear())
		assert.Equal(t, 0, m.Len())
	}
}

// Copy
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Copy() {
	m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two"})
	fmt.Println(m.Copy("1"))
	// Output: &[{1 one}]
}

func TestStringMap_Copy(t *testing.T) {
	// copy all
	{
		m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		copy := m.Copy()
		assert.Equal(t, m, copy)
		copy.Set("1", "foo")
		assert.Equal(t, map[string]interface{}{"1": "one", "2": "two", "3": "three"}, m.O())
		assert.Equal(t, map[string]interface{}{"1": "foo", "2": "two", "3": "three"}, copy.O())
	}

	// Copy specific keys
	{
		m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		copy := m.Copy("2", "3")
		assert.Equal(t, map[string]interface{}{"1": "one", "2": "two", "3": "three"}, m.O())
		assert.Equal(t, map[string]interface{}{"2": "two", "3": "three"}, copy.O())
	}
}

// Delete
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Delete() {
	m := NewStringMapV(map[string]interface{}{"1": "one"})
	fmt.Println(m.Delete("1").O())
	// Output: one
}

func TestStringMap_Delete(t *testing.T) {
	m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
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
	m := NewStringMapV(map[string]interface{}{"1": "one"})
	fmt.Println(m.DeleteM("1"))
	// Output: &[]
}

func TestStringMap_DeleteM(t *testing.T) {
	m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
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

// Dump
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Dump() {
	m := NewStringMapV(map[string]interface{}{"1": "one"})
	fmt.Println(m.Dump())
	// Output: "1": one
}

func TestStringMap_Dump(t *testing.T) {
	{
		m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.Equal(t, `"1": one
"2": two
"3": three
`, m.Dump())
	}
	{
		m := NewStringMapV(map[string]interface{}{"1": map[string]interface{}{"2": "two", "3": "three"}})
		assert.Equal(t, `"1":
  "2": two
  "3": three
`, m.Dump())
	}
}

// Exists
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Exists() {
	m := NewStringMapV(map[string]interface{}{"1": "one"})
	fmt.Println(m.Exists("1"))
	// Output: true
}

func TestStringMap_Exists(t *testing.T) {
	// bool
	{
		m := NewStringMapV(map[string]interface{}{"1": true})

		// Non existant key
		assert.Equal(t, false, m.Exists("0"))

		// exists
		assert.Equal(t, true, m.Exists("1"))
	}

	// string
	{
		// none
		assert.Equal(t, false, NewStringMapV().Exists("0"))

		m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})

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

// Convert to Golang internal types
//--------------------------------------------------------------------------------------------------
func TestToStringMapG(t *testing.T) {
	{
		// single value
		y := yaml.MapSlice{yaml.MapItem{Key: "k", Value: "v"}}
		g := NewStringMapV(y).G()
		assert.Equal(t, map[string]interface{}{"k": "v"}, g)
	}

	{
		// nested MapSlice values
		y := yaml.MapSlice{yaml.MapItem{Key: "k1", Value: yaml.MapSlice{yaml.MapItem{Key: "k2", Value: "v2"}}}}
		g := NewStringMapV(y).G()
		assert.Equal(t, map[string]interface{}{"k1": map[string]interface{}{"k2": "v2"}}, g)
	}

	{
		// nested StringMap values
		y := M().Add("k1", M().Add("k2", "v2"))
		g := NewStringMapV(y).G()
		assert.Equal(t, map[string]interface{}{"k1": map[string]interface{}{"k2": "v2"}}, g)
	}
}

func TestToStringMapGE(t *testing.T) {
	{
		// single value
		y := yaml.MapSlice{yaml.MapItem{Key: "k", Value: "v"}}
		g, err := NewStringMapV(y).GE()
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"k": "v"}, g)
	}

	{
		// nested MapSlice values
		y := yaml.MapSlice{yaml.MapItem{Key: "k1", Value: yaml.MapSlice{yaml.MapItem{Key: "k2", Value: "v2"}}}}
		g, err := NewStringMapV(y).GE()
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"k1": map[string]interface{}{"k2": "v2"}}, g)
	}

	{
		// nested StringMap values
		y := M().Add("k1", M().Add("k2", "v2"))
		g, err := NewStringMapV(y).GE()
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"k1": map[string]interface{}{"k2": "v2"}}, g)
	}
}

// Generic
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Generic() {
	fmt.Println(NewStringMapV().Generic())
	// Output: false
}

func TestStringMap_Generic(t *testing.T) {
	assert.Equal(t, false, NewStringMapV().Generic())
}

// Get
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Get() {
	m := NewStringMapV(map[string]interface{}{"1": "one"})
	fmt.Println(m.Get("1").O())
	// Output: one
}

func TestStringMap_Get(t *testing.T) {
	m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
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

// // Inject
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Inject() {
// 	fmt.Println(NewStringMapV().Inject(".", map[string]interface{}{"1": "one"}))
// 	// Output: &map[1:one]
// }

// func TestStringMap_Inject(t *testing.T) {

// 	// Move through list
// 	{
// 		// Move through by name and change order
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 				map[string]interface{}{"name": "bob", "order": "3"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar", "order": "4"},
// 				map[string]interface{}{"name": "bob", "order": "3"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2.[name==bar].order", "4").MG())
// 	}
// 	{
// 		// Move through and change name
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "bar"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "blah"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2.[1].name", "blah").MG())
// 	}

// 	// Inject into list
// 	{
// 		// Replace specific map by key value pair
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "bar"},
// 				map[string]interface{}{"name": "bob"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "blah"},
// 				map[string]interface{}{"name": "bob"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2.[name==bar]", map[string]interface{}{"name": "blah"}).MG())
// 	}
// 	{
// 		// Replace element in list
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{"1", "2"},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{"1", "3"},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2.[1]", "3").MG())
// 	}
// 	{
// 		// Replace the whole list as we gave the whole list selector
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []string{"1", "2"},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": "3",
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2.[]", "3").MG())
// 	}

// 	// Simple key indexing
// 	{
// 		// Nesting - merge
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "five",
// 			},
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"4": "four",
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2", b).MG())
// 	}
// 	{
// 		// Nesting - two
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": map[string]interface{}{
// 					"4": "four",
// 				},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2.3", b).MG())
// 	}
// 	{
// 		// Nesting - one
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2", b).MG())
// 	}
// 	{
// 		// Nesting - doesn't exist two
// 		a := map[string]interface{}{
// 			"1": "one",
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": map[string]interface{}{
// 					"4": "four",
// 				},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2.3", b).MG())
// 	}

// 	// Nesting - doesn't exist one
// 	{
// 		a := map[string]interface{}{
// 			"1": "one",
// 		}
// 		b := map[string]interface{}{
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Inject("2", b).MG())
// 	}

// 	// Root injection - override
// 	{
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"2": "two",
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 			"3": "three",
// 		}
// 		assert.Equal(t, expected, M(a).Inject(".", b).MG())
// 	}

// 	// Root injections - empty
// 	{
// 		b := map[string]interface{}{
// 			"foo": "bar",
// 		}
// 		expected := map[string]interface{}{
// 			"foo": "bar",
// 		}
// 		// map[string]interface{}
// 		assert.Equal(t, expected, MV().Inject("", b).MG())
// 		assert.Equal(t, expected, MV().Inject(".", b).MG())

// 		// *StringMap
// 		assert.Equal(t, expected, MV().Inject("", M(b)).MG())
// 		assert.Equal(t, expected, MV().Inject(".", M(b)).MG())

// 		// string
// 		m, err := MV().InjectE("", "foo")
// 		assert.Equal(t, &StringMap{}, m)
// 		assert.Equal(t, "invalid selector for the type of value given, 'string'", err.Error())

// 		// int
// 		m, err = MV().InjectE("", 2)
// 		assert.Equal(t, &StringMap{}, m)
// 		assert.Equal(t, "invalid selector for the type of value given, 'int'", err.Error())
// 	}
// }

// // Keys
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Keys() {
// 	m := NewStringMapV(map[string]interface{}{"1": "one"})
// 	fmt.Println(m.Keys().O())
// 	// Output: [1]
// }

// func TestStringMap_Keys(t *testing.T) {
// 	// nil or empty
// 	{
// 		assert.Equal(t, []string{}, (*StringMap)(nil).Keys().O())
// 		assert.Equal(t, []string{}, NewStringMapV().Keys().O())
// 	}

// 	// many
// 	{
// 		m := NewStringMapV(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
// 		keys := m.Keys().O()
// 		assert.Len(t, keys, 3)
// 		assert.Contains(t, keys, "1")
// 		assert.Contains(t, keys, "2")
// 		assert.Contains(t, keys, "3")
// 	}
// }

// // Len
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Len() {
// 	fmt.Println(NewStringMapV().SetM("k", "v").Len())
// 	// Output: 1
// }

// func TestStringMap_Len(t *testing.T) {
// 	m := NewStringMapV()
// 	assert.Equal(t, 0, m.Len())
// 	assert.Equal(t, 1, m.SetM("1", "one").Len())
// 	assert.Equal(t, 2, m.SetM("2", "two").Len())
// 	assert.Equal(t, 1, m.DeleteM("2").Len())
// 	assert.Equal(t, 0, m.DeleteM("1").Len())
// }

// // Merge
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Merge() {
// 	fmt.Println(NewStringMapV(map[string]interface{}{"1": "two"}).Merge(NewStringMapV(map[string]interface{}{"1": "one"})))
// 	// Output: &map[1:one]
// }

// func TestStringMap_Merge(t *testing.T) {

// 	// Location tests
// 	{
// 		// Nesting - merge advanced
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "five",
// 			},
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 			"5": "five",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "four",
// 				"5": "five",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).Merge(M(b), "2").MG())
// 	}
// 	{
// 		// Nesting - merge
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "five",
// 			},
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "four",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).Merge(M(b), "2").MG())
// 	}
// 	{
// 		// Nesting - two
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": map[string]interface{}{
// 					"4": "four",
// 				},
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).Merge(M(b), "2.3").MG())
// 	}
// 	{
// 		// Nesting - one
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).Merge(M(b), "2").MG())
// 	}
// 	{
// 		// Nesting - doesn't exist two
// 		a := map[string]interface{}{
// 			"1": "one",
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": map[string]interface{}{
// 					"4": "four",
// 				},
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).Merge(M(b), "2.3").MG())
// 	}
// 	{
// 		// Nesting - doesn't exist one
// 		a := map[string]interface{}{
// 			"1": "one",
// 		}
// 		b := map[string]interface{}{
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).Merge(M(b), "2").MG())
// 	}
// 	{
// 		// root indicator
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"2": "two",
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 			"3": "three",
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).Merge(M(b), ".").MG())
// 	}

// 	// nil or empty
// 	{
// 		assert.Equal(t, NewStringMapV(), (*StringMap)(nil).Merge(nil))
// 		assert.Equal(t, NewStringMapV(), NewStringMapV().Merge(nil))
// 		assert.Equal(t, map[string]interface{}{}, NewStringMapV().Merge(nil).O())
// 		assert.Equal(t, NewStringMapV(), NewStringMapV().Merge(NewStringMapV()))
// 		assert.Equal(t, map[string]interface{}{}, NewStringMapV().Merge(NewStringMapV()).O())
// 	}
// 	{
// 		a := NewStringMapV(map[string]interface{}{})
// 		b := NewStringMapV(map[string]interface{}{"1": "one"})
// 		assert.Equal(t, b, a.Merge(b))
// 	}
// 	{
// 		a := NewStringMapV(map[string]interface{}{"1": "one"})
// 		b := NewStringMapV(map[string]interface{}{})
// 		assert.Equal(t, a, a.Merge(b))
// 	}
// 	{
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": "two",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 		})
// 		assert.Equal(t, expected, a.Merge(b))
// 	}
// 	{
// 		// Override string in a with string in b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": "two",
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected, a.Merge(b))
// 	}
// 	{
// 		// Override string in a with map from b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": NewStringMapV(map[string]interface{}{"foo": "bar"}),
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": NewStringMapV(map[string]interface{}{"foo": "bar"}),
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected, a.Merge(b))
// 	}
// 	{
// 		// Override map in a with string from b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": NewStringMapV(map[string]interface{}{"foo": "bar"}),
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": "two",
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected, a.Merge(b))
// 	}
// 	{
// 		// Override sub map string in a with sub map string from b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": NewStringMapV(map[string]interface{}{"foo": "bar1"}),
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": NewStringMapV(map[string]interface{}{
// 				"foo":  "bar2",
// 				"foo2": "bar2",
// 			}),
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": NewStringMapV(map[string]interface{}{
// 				"foo":  "bar2",
// 				"foo2": "bar2",
// 			}),
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected, a.Merge(b))
// 	}
// }

// // MergeG
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_MergeG() {
// 	fmt.Println(NewStringMapV(map[string]interface{}{"1": "two"}).MergeG(NewStringMapV(map[string]interface{}{"1": "one"})))
// 	// Output: map[1:one]
// }

// func TestStringMap_MergeG(t *testing.T) {

// 	// Location tests
// 	{
// 		// Nesting - merge advanced
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "five",
// 			},
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 			"5": "five",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "four",
// 				"5": "five",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).MergeG(M(b), "2"))
// 	}
// 	{
// 		// Nesting - merge
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "five",
// 			},
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 				"4": "four",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).MergeG(M(b), "2"))
// 	}
// 	{
// 		// Nesting - two
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": map[string]interface{}{
// 					"4": "four",
// 				},
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).MergeG(M(b), "2.3"))
// 	}
// 	{
// 		// Nesting - one
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).MergeG(M(b), "2"))
// 	}
// 	{
// 		// Nesting - doesn't exist two
// 		a := map[string]interface{}{
// 			"1": "one",
// 		}
// 		b := map[string]interface{}{
// 			"4": "four",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": map[string]interface{}{
// 					"4": "four",
// 				},
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).MergeG(M(b), "2.3"))
// 	}
// 	{
// 		// Nesting - doesn't exist one
// 		a := map[string]interface{}{
// 			"1": "one",
// 		}
// 		b := map[string]interface{}{
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"3": "three",
// 			},
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).MergeG(M(b), "2"))
// 	}
// 	{
// 		// root indicator
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		}
// 		b := map[string]interface{}{
// 			"2": "two",
// 			"3": "three",
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 			"3": "three",
// 		}
// 		assert.Equal(t, expected, NewStringMap(a).MergeG(M(b), "."))
// 	}

// 	// nil or empty
// 	{
// 		assert.Equal(t, map[string]interface{}{}, NewStringMapV().MergeG(nil))
// 		assert.Equal(t, map[string]interface{}{}, NewStringMapV().MergeG(NewStringMapV()))
// 	}
// 	{
// 		a := NewStringMapV(map[string]interface{}{})
// 		b := NewStringMapV(map[string]interface{}{"1": "one"})
// 		assert.Equal(t, b.G(), a.MergeG(b))
// 	}
// 	{
// 		a := NewStringMapV(map[string]interface{}{"1": "one"})
// 		b := NewStringMapV(map[string]interface{}{})
// 		assert.Equal(t, a.G(), a.MergeG(b))
// 	}
// 	{
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": "two",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 		})
// 		assert.Equal(t, expected.G(), a.MergeG(b))
// 	}
// 	{
// 		// Override string in a with string in b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "2",
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": "two",
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected.G(), a.MergeG(b))
// 	}
// 	{
// 		// Override string in a with map from b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": NewStringMapV(map[string]interface{}{"foo": "bar"}),
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{"foo": "bar"},
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected.G(), a.MergeG(b))
// 	}
// 	{
// 		// Override map in a with string from b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": NewStringMapV(map[string]interface{}{"foo": "bar"}),
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": "two",
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": "two",
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected.G(), a.MergeG(b))
// 	}
// 	{
// 		// Override sub map string in a with sub map string from b
// 		a := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": NewStringMapV(map[string]interface{}{"foo": "bar1"}),
// 		})
// 		b := NewStringMapV(map[string]interface{}{
// 			"2": NewStringMapV(map[string]interface{}{
// 				"foo":  "bar2",
// 				"foo2": "bar2",
// 			}),
// 			"3": "three",
// 		})
// 		expected := NewStringMapV(map[string]interface{}{
// 			"1": "one",
// 			"2": map[string]interface{}{
// 				"foo":  "bar2",
// 				"foo2": "bar2",
// 			},
// 			"3": "three",
// 		})
// 		assert.Equal(t, expected.G(), a.MergeG(b))
// 		assert.Equal(t, map[string]interface{}{"foo": "bar2", "foo2": "bar2"}, expected.G()["2"])
// 	}
// }

// // O
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_O() {
// 	fmt.Println(NewStringMapV(map[string]interface{}{"1": "one"}).O())
// 	// Output: map[1:one]
// }

// func TestStringMap_O(t *testing.T) {
// 	assert.Equal(t, map[string]interface{}{}, (*StringMap)(nil).O())
// 	assert.Equal(t, map[string]interface{}{}, NewStringMapV().O())
// 	assert.Equal(t, map[string]interface{}{"1": "one"}, NewStringMapV(map[string]interface{}{"1": "one"}).O())
// }

// // Query
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Query() {
// 	fmt.Println(ToStringMap("foo:\n  bar: 1\n").Query("foo.bar"))
// 	// Output: 1
// }

// func TestStringMap_Query(t *testing.T) {

// 	// malformed yaml
// 	{
// 		_, err := NewStringMap("foo:\n\t- 1\n").QueryE("foo.[0]")
// 		assert.Equal(t, "failed to query empty map", err.Error())
// 	}

// 	// dot notation
// 	{
// 		// Identity: .
// 		assert.Equal(t, &StringMap{"one": "1"}, NewStringMapV(map[string]interface{}{"one": "1"}).Query(``).ToStringMap())
// 		assert.Equal(t, &StringMap{"one": "1"}, NewStringMapV(map[string]interface{}{"one": "1"}).Query(`.`).ToStringMap())

// 		// Object Identifier-Index: .foo, .foo.bar
// 		assert.Equal(t, "1", NewStringMapV(map[string]interface{}{"one": "1"}).Query(`one`).ToString())
// 		assert.Equal(t, "1", NewStringMapV(map[string]interface{}{"one": "1"}).Query(`.one`).ToString())
// 		assert.Equal(t, "1", NewStringMapV(map[string]interface{}{"one": "1"}).Query(`."one"`).ToString())
// 		assert.Equal(t, "2", NewStringMapV(map[string]interface{}{"one": map[string]interface{}{"two": "2"}}).Query(`one.two`).ToString())
// 		assert.Equal(t, "3", NewStringMapV(map[string]interface{}{"one": map[string]interface{}{"two": map[string]interface{}{"three": "3"}}}).Query(`one.two.three`).ToString())
// 		assert.Equal(t, "foo", NewStringMapV(map[string]interface{}{"one": map[string]interface{}{"two.three": "foo"}}).Query(`one."two.three"`).ToString())

// 		// Array Index: empty array
// 		assert.True(t, ToStringMap("foo: []\n").Query(`foo.[0]`).Nil())

// 		// Array Index: move through index
// 		assert.Equal(t, 1.1, ToStringMap("items:\n  - item: 1\n    val: 1.1\n  - item: 2\n    val: 2.2\n  - item: 3\n    val: 3.3\n").Query("items.[0].val").ToFloat64())

// 		// Array Index: .[2], .[-1]
// 		assert.Equal(t, 0, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[0]`).ToInt())
// 		assert.Equal(t, 1, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[1]`).ToInt())
// 		assert.Equal(t, 2, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[2]`).ToInt())
// 		assert.Equal(t, 2, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[-1]`).ToInt())
// 		assert.Equal(t, 1, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[-2]`).ToInt())
// 		assert.Equal(t, 0, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[-3]`).ToInt())
// 		assert.Equal(t, float32(2.2), ToStringMap("foo:\n  - 0.5\n  - 1.1\n  - 2.2\n").Query(`foo.[-1]`).ToFloat32())
// 		assert.Equal(t, float64(2.2), ToStringMap("foo:\n  - 0.5\n  - 1.1\n  - 2.2\n").Query(`foo.[-1]`).ToFloat64())
// 		assert.Equal(t, "0", NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[0]`).ToString())
// 		assert.Equal(t, "1", NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[1]`).ToString())
// 		assert.Equal(t, "2", NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[2]`).ToString())
// 		assert.Equal(t, "2", NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-1]`).ToString())
// 		assert.Equal(t, "1", NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-2]`).ToString())
// 		assert.Equal(t, "0", NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-3]`).ToString())
// 		assert.Equal(t, "", NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[-5]`).ToString())

// 		// Array Interator: .[]
// 		assert.Equal(t, []int{0, 1, 2}, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[]`).ToIntSliceG())
// 		assert.Equal(t, []int{0, 1, 2}, NewStringMapV(map[string]interface{}{"foo": []interface{}{float64(0), float64(1), float64(2)}}).Query(`foo.[]`).ToIntSliceG())

// 		// Array Slice: .[1:]

// 		// Array element selection based on element value
// 		yml := `one:
//   - name: foo
//     val:
//       - 3.3
//       - 1.1
//       - 2.2
//   - name: bar
//     val: 3
// `
// 		assert.Equal(t, 3, NewStringMap(yml).Query(`one.[name==bar].val`).ToInt())
// 		assert.Equal(t, 2, NewStringMap(yml).Query(`one.[name==foo].val.[-1]`).ToInt())
// 		assert.Equal(t, 1.1, NewStringMap(yml).Query(`one.[name==foo].val.[-2]`).ToFloat64())
// 		assert.Equal(t, "3.3", NewStringMap(yml).Query(`one.[name==foo].val.[-3]`).ToString())
// 	}

// 	// no dot notation
// 	{
// 		// floats
// 		assert.Equal(t, float32(1.2), NewStringMapV(map[string]interface{}{"1": 1.2}).Query("1").ToFloat32())
// 		assert.Equal(t, float64(1.2), NewStringMapV(map[string]interface{}{"1": 1.2}).Query("1").ToFloat64())

// 		// ints
// 		assert.Equal(t, int(1), NewStringMapV(map[string]interface{}{"1": 1}).Query("1").ToInt())
// 		assert.Equal(t, int8(1), NewStringMapV(map[string]interface{}{"1": 1}).Query("1").ToInt8())
// 		assert.Equal(t, int16(1), NewStringMapV(map[string]interface{}{"1": 1}).Query("1").ToInt16())
// 		assert.Equal(t, int32(1), NewStringMapV(map[string]interface{}{"1": 1}).Query("1").ToInt32())
// 		assert.Equal(t, int64(1), NewStringMapV(map[string]interface{}{"1": 1}).Query("1").ToInt64())
// 		assert.Equal(t, uint(1), NewStringMapV(map[string]interface{}{"1": 1}).Query("1").ToUint())

// 		// string
// 		assert.Equal(t, "one", NewStringMapV(map[string]interface{}{"1": "one"}).Query("1").ToString())

// 		// maps
// 		assert.Equal(t, &StringMap{"2": "two"}, NewStringMapV(map[string]interface{}{"1": map[string]interface{}{"2": "two"}}).Query("1").ToStringMap())
// 		assert.Equal(t, map[string]interface{}{"2": "two"}, NewStringMapV(map[string]interface{}{"1": map[string]interface{}{"2": "two"}}).Query("1").ToStringMapG())

// 		// yaml
// 		assert.Equal(t, "1", ToStringMap("one: 1").Query("one").ToString())
// 		assert.Equal(t, []int{1, 2}, ToStringMap("foo: \n  - 1\n  - 2").Query("foo").ToIntSliceG())
// 		assert.Equal(t, &IntSlice{1, 2}, ToStringMap("foo: \n  - 1\n  - 2").Query("foo").ToIntSlice())
// 	}
// }

// // Remove
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Remove() {
// 	fmt.Println(ToStringMap("foo:\n  bar: 1\n  foo2: 2\n").Remove("foo.bar"))
// 	// Output: &map[foo:map[foo2:2]]
// }

// func TestStringMap_Remove(t *testing.T) {

// 	// Move through list
// 	{
// 		// Remove the order property on all maps in the list
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 				map[string]interface{}{"name": "bob", "order": "3"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "bar"},
// 				map[string]interface{}{"name": "bob"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2.[].order").MG())
// 	}
// 	{
// 		// Remove the first item's order property
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2.[0].order").MG())
// 	}
// 	{
// 		// Remove the first item's order property - neg notation
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2.[-2].order").MG())
// 	}
// 	{
// 		// Remove the second item's order property - neg
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2.[-1].order").MG())
// 	}
// 	{
// 		// Remove the second item's order property
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar", "order": "2"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo", "order": "1"},
// 				map[string]interface{}{"name": "bar"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2.[1].order").MG())
// 	}

// 	// Remove from list
// 	{
// 		// Remove list item by map's key value
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "bar"},
// 				map[string]interface{}{"name": "bob"},
// 			},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{
// 				map[string]interface{}{"name": "foo"},
// 				map[string]interface{}{"name": "bob"},
// 			},
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2.[name==bar]").MG())
// 	}
// 	{
// 		// Replace element in list
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{"1", "2"},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 			"2": []interface{}{"2"},
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2.[0]").MG())
// 	}
// 	{
// 		// Remove the whole list by key
// 		a := map[string]interface{}{
// 			"1": "one",
// 			"2": []string{"1", "2"},
// 		}
// 		expected := map[string]interface{}{
// 			"1": "one",
// 		}
// 		assert.Equal(t, expected, M(a).Remove("2").MG())
// 	}

// 	// Identity: .
// 	assert.Equal(t, &StringMap{"one": "1"}, NewStringMapV(map[string]interface{}{"one": "1"}).Remove(``))
// 	assert.Equal(t, &StringMap{"one": "1"}, NewStringMapV(map[string]interface{}{"one": "1"}).Remove(`.`))

// 	// Object Identifier-Index: .foo, .foo.bar
// 	assert.Equal(t, &StringMap{}, NewStringMapV(map[string]interface{}{"one": "1"}).Remove(`one`))
// 	assert.Equal(t, &StringMap{}, NewStringMapV(map[string]interface{}{"one": "1"}).Remove(`.one`))
// 	assert.Equal(t, &StringMap{}, NewStringMapV(map[string]interface{}{"one": "1"}).Remove(`."one"`))
// 	assert.Equal(t, &StringMap{"one": map[string]interface{}{}}, NewStringMapV(map[string]interface{}{"one": map[string]interface{}{"two": "2"}}).Remove(`one.two`))
// 	assert.Equal(t, &StringMap{"one": map[string]interface{}{"two": map[string]interface{}{}}}, NewStringMapV(map[string]interface{}{"one": map[string]interface{}{"two": map[string]interface{}{"three": "3"}}}).Remove(`one.two.three`))
// 	assert.Equal(t, &StringMap{"one": map[string]interface{}{}}, NewStringMapV(map[string]interface{}{"one": map[string]interface{}{"two.three": "foo"}}).Remove(`one."two.three"`))
// }

// // Set
// //--------------------------------------------------------------------------------------------------
// func ExampleStringMap_Set() {
// 	fmt.Println(NewStringMapV().Set("key", "value"))
// 	// Output: true
// }

// func TestStringMap_Set(t *testing.T) {
// 	// bool
// 	{
// 		m := NewStringMapV()
// 		assert.Equal(t, true, m.Set("test", false))
// 		assert.Equal(t, false, m.Set("test", true))
// 		assert.Equal(t, true, (*m)["test"].(bool))
// 	}

// 	// string
// 	{
// 		m := NewStringMapV()
// 		assert.Equal(t, true, m.Set("1", "one"))
// 		assert.Equal(t, false, m.Set("1", "two"))
// 		assert.Equal(t, "two", (*m)["1"].(string))
// 	}
// }

// SetM
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_SetM() {
	fmt.Println(NewStringMapV().SetM("k", "v"))
	// Output: &[{k v}]
}

func TestStringMap_SetM(t *testing.T) {
	// bool
	{
		m := NewStringMapV()
		assert.Equal(t, NewStringMapV(map[string]interface{}{"test": false}), m.SetM("test", false))
		assert.Equal(t, false, m.Get("test").ToBool())
		assert.Equal(t, NewStringMapV(map[string]interface{}{"test": true}), m.SetM("test", true))
		assert.Equal(t, true, m.Get("test").ToBool())
	}

	// string
	{
		m := NewStringMapV()
		assert.Equal(t, NewStringMapV(map[string]interface{}{"1": "one"}), m.SetM("1", "one"))
		assert.Equal(t, "one", m.Get("1").A())
		assert.Equal(t, NewStringMapV(map[string]interface{}{"1": "two"}), m.SetM("1", "two"))
		assert.Equal(t, "two", m.Get("1").A())
	}
}

// WriteJSON
//--------------------------------------------------------------------------------------------------
func TestWriteJSON(t *testing.T) {
	clearTmpDir()

	// Convert yaml string into a data structure
	m1 := NewStringMap(map[string]interface{}{"1": "one"})

	// Write out the data structure as json to disk
	err := m1.WriteJSON(tmpFile)
	assert.Nil(t, err)

	// Read the file back into memory and compare data structure
	m2, err := LoadJSONE(tmpFile)
	assert.Nil(t, err)

	assert.Equal(t, m1, m2)
}

// WriteYAML
//--------------------------------------------------------------------------------------------------
func TestWriteYAML(t *testing.T) {
	clearTmpDir()

	{
		// Convert yaml string into a data structure
		m1 := NewStringMap(map[string]interface{}{"1": "one"})

		// Write out the data structure as yaml to disk
		err := m1.WriteYAML(tmpFile)
		assert.Nil(t, err)

		// Read the file back into memory and compare data structure
		m2, err := LoadYAMLE(tmpFile)
		assert.Nil(t, err)

		assert.Equal(t, m1, m2)
	}

	{
		// Write out the data structure as yaml to disk
		data1 := "b: b1\na: a1\n"
		err := MV(data1).WriteYAML(tmpFile)
		assert.NoError(t, err)

		// Read the file back into memory and compare raw string
		var buffer []byte
		buffer, err = ioutil.ReadFile(tmpFile)
		assert.Nil(t, err)
		data2 := string(buffer)
		assert.Equal(t, data1, data2)

		// Now compare data structures
		m, err := LoadYAMLE(tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, M().Add("b", "b1").Add("a", "a1"), m)
	}
}
