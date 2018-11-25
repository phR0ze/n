package nub

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestStrMapMerge(t *testing.T) {
	{
		strMap := NewStrMap()
		assert.Equal(t, map[string]interface{}{}, strMap.Merge(nil).M())
	}
	{
		data := map[string]interface{}{"1": "one"}
		expected := map[string]interface{}{"1": "one"}
		assert.Equal(t, expected, NewStrMap().Merge(data).M())
	}
	{
		strMap := StrMap(map[string]interface{}{
			"1": "one", "2": "three", "3": "four",
		})
		data := []map[string]interface{}{
			{"2": "two"},
			{"3": "three"},
			{"4": "four"},
		}
		expected := map[string]interface{}{
			"1": "one", "2": "two", "3": "three", "4": "four",
		}
		assert.Equal(t, expected, strMap.Merge(data...).M())
	}
}
func TestStrMapStrMap(t *testing.T) {
	{
		// Manual: Not a valid str map should return empty
		strMap := StrMap(map[string]interface{}{
			"test1": "foo",
		})
		expected := map[string]interface{}{}
		assert.Equal(t, expected, strMap.StrMap("test1").M())
	}
	{
		// Unmarshal: Not a valid str map should return nil
		raw := `test1: "foobar"`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(raw), &data)
		expected := map[string]interface{}{}

		assert.Equal(t, expected, StrMap(data).StrMap("test1").M())
	}
	{
		// Manual: valid nested str map
		strMap := StrMap(map[string]interface{}{
			"test1": map[string]interface{}{
				"test2": "foobar",
			},
		})
		expected := map[string]interface{}{
			"test2": "foobar",
		}
		assert.Equal(t, expected, strMap.StrMap("test1").M())
	}
	{
		// Unmarshal: valid nested JQ
		raw := `test1: 
  test2: foobar`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(raw), &data)
		expected := map[string]interface{}{
			"test2": "foobar",
		}
		assert.Equal(t, expected, StrMap(data).StrMap("test1").M())
	}
	{
		strMap := StrMap(map[string]interface{}{
			"test1": map[string]interface{}{
				"test2": map[string]interface{}{
					"test3": "foobar",
				},
			},
		})
		expected := map[string]interface{}{
			"test3": "foobar",
		}
		assert.Equal(t, expected, strMap.StrMap("test1.test2").M())
	}
}
func TestMergeMap(t *testing.T) {
	{
		assert.Equal(t, map[string]interface{}{}, mergeMap(nil, nil))
	}
	{
		a := map[string]interface{}{}
		b := map[string]interface{}{}
		assert.Equal(t, map[string]interface{}{}, mergeMap(a, b))
	}
	{
		a := map[string]interface{}{}
		b := map[string]interface{}{"1": "one"}
		assert.Equal(t, b, mergeMap(a, b))
	}
	{
		a := map[string]interface{}{"1": "one"}
		b := map[string]interface{}{}
		assert.Equal(t, a, mergeMap(a, b))
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
		assert.Equal(t, expected, mergeMap(a, b))
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
		assert.Equal(t, expected, mergeMap(a, b))
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
		assert.Equal(t, expected, mergeMap(a, b))
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
		assert.Equal(t, expected, mergeMap(a, b))
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
		assert.Equal(t, expected, mergeMap(a, b))
	}
}
