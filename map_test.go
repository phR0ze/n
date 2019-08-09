package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewMap
//--------------------------------------------------------------------------------------------------
func ExampleNewMap() {
	fmt.Println(NewMap(map[string]interface{}{"k": "v"}))
	// Output: &map[k:v]
}

func TestNewMap(t *testing.T) {

	// []byte
	{
		assert.Equal(t, &StringMap{"foo": map[string]interface{}{"bar": float64(1)}}, NewMap([]byte("foo:\n bar: 1\n")))
	}

	// string
	{
		assert.Equal(t, &StringMap{"foo": map[string]interface{}{"bar": float64(1)}}, NewMap("foo:\n bar: 1\n"))
	}

	// string interface
	{
		m := map[string]interface{}{"k": "v"}
		assert.Equal(t, NewStringMapV(m), NewMap(m))
	}

	// StringMap
	{
		m := NewStringMapV()
		m.Set("k", "v")
		assert.Equal(t, m, NewMap(m))
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
