package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Map
//--------------------------------------------------------------------------------------------------
func ExampleMap() {
	fmt.Println(Map(map[string]interface{}{"k": "v"}))
	// Output: &map[k:v]
}

func TestMap(t *testing.T) {

	// map[string]bool
	{
		assert.Equal(t, &StringMap{"1": true}, Map(map[string]bool{"1": true}))
	}

	// float64: map[string]interface{}
	{
		assert.Equal(t, &StringMap{"foo": map[string]interface{}{"bar": float64(1)}}, Map("foo:\n bar: 1\n"))
	}

	// string: map[string]interface{}
	{
		m := map[string]interface{}{"k": "v"}
		assert.Equal(t, NewStringMapV(m), Map(m))
	}

	// StringMap
	{
		m := NewStringMapV()
		m.Set("k", "v")
		assert.Equal(t, m, Map(m))
	}
}

// MergeStringMap
//--------------------------------------------------------------------------------------------------
func ExampleMergeStringMap() {
	fmt.Println(MergeStringMap(map[string]interface{}{"1": "two"}, map[string]interface{}{"1": "one"}))
	// Output: map[1:one]
}

func TestMergeStringMap(t *testing.T) {

	// Location tests
	{
		// Nesting - merge advanced
		a := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": "three",
				"4": "five",
			},
		}
		b := map[string]interface{}{
			"4": "four",
			"5": "five",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": "three",
				"4": "four",
				"5": "five",
			},
		}
		assert.Equal(t, expected, MergeStringMap(a, b, "2"))
	}
	{
		// Nesting - merge
		a := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": "three",
				"4": "five",
			},
		}
		b := map[string]interface{}{
			"4": "four",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": "three",
				"4": "four",
			},
		}
		assert.Equal(t, expected, MergeStringMap(a, b, "2"))
	}
	{
		// Nesting - two
		a := map[string]interface{}{
			"1": "one",
			"2": "2",
		}
		b := map[string]interface{}{
			"4": "four",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": map[string]interface{}{
					"4": "four",
				},
			},
		}
		assert.Equal(t, expected, MergeStringMap(a, b, "2.3"))
	}
	{
		// Nesting - one
		a := map[string]interface{}{
			"1": "one",
			"2": "2",
		}
		b := map[string]interface{}{
			"3": "three",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": "three",
			},
		}
		assert.Equal(t, expected, MergeStringMap(a, b, "2"))
	}
	{
		// Nesting - doesn't exist two
		a := map[string]interface{}{
			"1": "one",
		}
		b := map[string]interface{}{
			"4": "four",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": map[string]interface{}{
					"4": "four",
				},
			},
		}
		assert.Equal(t, expected, MergeStringMap(a, b, "2.3"))
	}
	{
		// Nesting - doesn't exist one
		a := map[string]interface{}{
			"1": "one",
		}
		b := map[string]interface{}{
			"3": "three",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"3": "three",
			},
		}
		assert.Equal(t, expected, MergeStringMap(a, b, "2"))
	}
	{
		// root indicator
		a := map[string]interface{}{
			"1": "one",
			"2": "2",
		}
		b := map[string]interface{}{
			"2": "two",
			"3": "three",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		}
		assert.Equal(t, expected, MergeStringMap(a, b, "."))
	}

	// Regular tests
	{
		assert.Equal(t, map[string]interface{}{}, MergeStringMap(nil, nil))
	}
	{
		a := map[string]interface{}{}
		b := map[string]interface{}{}
		assert.Equal(t, map[string]interface{}{}, MergeStringMap(a, b))
	}
	{
		a := map[string]interface{}{}
		b := map[string]interface{}{"1": "one"}
		assert.Equal(t, b, MergeStringMap(a, b))
	}
	{
		a := map[string]interface{}{"1": "one"}
		b := map[string]interface{}{}
		assert.Equal(t, a, MergeStringMap(a, b))
	}
	{
		a := map[string]interface{}{
			"1": "one",
		}
		b := map[string]interface{}{
			"2": "two",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
		}
		assert.Equal(t, expected, MergeStringMap(a, b))
	}
	{
		// Override string in a with string in b
		a := map[string]interface{}{
			"1": "one",
			"2": "2",
		}
		b := map[string]interface{}{
			"2": "two",
			"3": "three",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		}
		assert.Equal(t, expected, MergeStringMap(a, b))
	}
	{
		// Override string in a with map from b
		a := map[string]interface{}{
			"1": "one",
			"2": "two",
		}
		b := map[string]interface{}{
			"2": map[string]interface{}{"foo": "bar"},
			"3": "three",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{"foo": "bar"},
			"3": "three",
		}
		assert.Equal(t, expected, MergeStringMap(a, b))
	}
	{
		// Override map in a with string from b
		a := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{"foo": "bar"},
		}
		b := map[string]interface{}{
			"2": "two",
			"3": "three",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		}
		assert.Equal(t, expected, MergeStringMap(a, b))
	}
	{
		// Override sub map string in a with sub map string from b
		a := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{"foo": "bar1"},
		}
		b := map[string]interface{}{
			"2": map[string]interface{}{
				"foo":  "bar2",
				"foo2": "bar2",
			},
			"3": "three",
		}
		expected := map[string]interface{}{
			"1": "one",
			"2": map[string]interface{}{
				"foo":  "bar2",
				"foo2": "bar2",
			},
			"3": "three",
		}
		assert.Equal(t, expected, MergeStringMap(a, b))
	}
}

// IdxFromSelector
//--------------------------------------------------------------------------------------------------
func TestStringMap_IdxFromSelector(t *testing.T) {

	assert.Equal(t, []int{0, 1, 2}, ToStringMap("foo:\n  - 0\n  - 1\n  - 2\n").Query(`foo.[]`).ToIntSliceG())

	// Full slice
	{
		// Valid selector
		i, k, v, err := IdxFromSelector("[foo==bar]", 3)
		assert.Equal(t, -1, i)
		assert.Equal(t, "foo", k)
		assert.Equal(t, "bar", v)
		assert.Nil(t, err)

		// Invalid selector
		i, k, v, err = IdxFromSelector("[foo=bar]", 3)
		assert.Equal(t, -1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Equal(t, "invalid array index selector foo=bar", err.Error())

		// Valid index neg
		i, k, v, err = IdxFromSelector("[-1]", 3)
		assert.Equal(t, 2, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Nil(t, err)

		// Valid index pos
		i, k, v, err = IdxFromSelector("[1]", 3)
		assert.Equal(t, 1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Nil(t, err)

		// Invalid index
		i, k, v, err = IdxFromSelector("[3]", 3)
		assert.Equal(t, -1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Equal(t, "invalid array index 3", err.Error())

		// Invalid index
		i, k, v, err = IdxFromSelector("[-4]", 3)
		assert.Equal(t, -1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Equal(t, "invalid array index -4", err.Error())
	}
	// Empty slice
	{
		i, k, v, err := IdxFromSelector("[foo==bar]", 0)
		assert.Equal(t, -1, i)
		assert.Equal(t, "foo", k)
		assert.Equal(t, "bar", v)
		assert.Nil(t, err)

		// Invalid selector
		i, k, v, err = IdxFromSelector("[foo=bar]", 0)
		assert.Equal(t, -1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Equal(t, "invalid array index selector foo=bar", err.Error())

		i, k, v, err = IdxFromSelector("[]", 0)
		assert.Equal(t, -1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Nil(t, err)

		// Invalid index
		i, k, v, err = IdxFromSelector("[1]", 0)
		assert.Equal(t, -1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Equal(t, "invalid array index 1", err.Error())

		// Invalid index
		i, k, v, err = IdxFromSelector("[-2]", 0)
		assert.Equal(t, -1, i)
		assert.Equal(t, "", k)
		assert.Equal(t, "", v)
		assert.Equal(t, "invalid array index -2", err.Error())
	}
}

// KeysFromSelector
//--------------------------------------------------------------------------------------------------
func TestStringMap_KeysFromSelector(t *testing.T) {

	// Empty
	keys, err := KeysFromSelector("")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{}, keys)

	// Identity
	keys, err = KeysFromSelector(".")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{}, keys)

	// Object Identifier-Index: .foo, .foo.bar
	keys, err = KeysFromSelector("foo")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo"}, keys)

	keys, err = KeysFromSelector(".foo")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo"}, keys)

	keys, err = KeysFromSelector(`."foo"`)
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo"}, keys)

	keys, err = KeysFromSelector(".foo.bar")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo", "bar"}, keys)

	keys, err = KeysFromSelector("foo.bar")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo", "bar"}, keys)

	keys, err = KeysFromSelector("foo1.bar.foo2")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo1", "bar", "foo2"}, keys)

	keys, err = KeysFromSelector(`foo1."bar.foo2"`)
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo1", "bar.foo2"}, keys)

	// Array Index: .[], .[0], .[-1]
	keys, err = KeysFromSelector("foo.[]")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo", "[]"}, keys)

	keys, err = KeysFromSelector("foo.[0]")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo", "[0]"}, keys)

	keys, err = KeysFromSelector("foo.[-1]")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo", "[-1]"}, keys)

	// Array Index: move through index
	keys, err = KeysFromSelector("foo.[0].val")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"foo", "[0]", "val"}, keys)

	// Array element selection based on element value
	keys, err = KeysFromSelector("one.[name==bar].val")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"one", "[name==bar]", "val"}, keys)

	keys, err = KeysFromSelector("one.[name==bar].val.[-2]")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"one", "[name==bar]", "val", "[-2]"}, keys)

	// no dot notation
	keys, err = KeysFromSelector("one")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"one"}, keys)

	keys, err = KeysFromSelector("1")
	assert.Nil(t, err)
	assert.Equal(t, &StringSlice{"1"}, keys)
}
