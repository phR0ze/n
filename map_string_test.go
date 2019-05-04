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

	// string interface
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
