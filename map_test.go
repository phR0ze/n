package nub

import (
	"testing"

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
