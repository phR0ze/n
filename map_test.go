package nub

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {
	q := M()
	assert.NotNil(t, q)
	assert.NotNil(t, q.Iter)
	iter := q.Iter()
	assert.NotNil(t, iter)
	x, ok := iter()
	assert.Nil(t, x)
	assert.False(t, ok)
}

func TestMQ(t *testing.T) {
	{
		items := []interface{}{}
		q := Q(map[string]string{"1": "one", "2": "two", "3": "three"})
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			items = append(items, x)
			item := x.(*KeyVal)
			switch item.Key {
			case "1":
				assert.NotEqual(t, &KeyVal{Key: "2", Val: "one"}, item)
				assert.NotEqual(t, &KeyVal{Key: "1", Val: "two"}, item)
				assert.Equal(t, &KeyVal{Key: "1", Val: "one"}, item)
			case "2":
				assert.Equal(t, &KeyVal{Key: "2", Val: "two"}, item)
			case "3":
				assert.Equal(t, &KeyVal{Key: "3", Val: "three"}, item)
			}
		}
		assert.Len(t, items, 3)
	}
}

func TestMStrMap(t *testing.T) {
	{
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.Equal(t, 3, q.Len())
	}
}

func TestMAny(t *testing.T) {
	assert.False(t, M().Any())
	assert.False(t, Q(map[int]interface{}{}).Any())
	assert.True(t, Q(map[int]interface{}{1: "one"}).Any())
}

func TestMAnyWhere(t *testing.T) {
	{
		q := M()
		assert.False(t, q.AnyWhere(func(x interface{}) bool {
			return x == 3
		}))
	}
	{
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.False(t, q.AnyWhere(func(x interface{}) bool { return x == 3 }))
		assert.True(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(*KeyVal)).Key == "3"
		}))
		assert.True(t, q.AnyWhere(func(x interface{}) bool {
			return (x.(*KeyVal)).Val == "two"
		}))
	}
}

func TestMLen(t *testing.T) {
	{
		q := M()
		assert.Equal(t, 0, q.Len())
	}
	{
		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
		assert.Equal(t, 3, q.Len())
	}
}

// TODO: Need to refactor below here

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

func TestStrMapMergeNub(t *testing.T) {
	{
		strMap := NewStrMap()
		assert.Equal(t, map[string]interface{}{}, strMap.MergeNub().M())
	}
	{
		strMap := NewStrMap()
		assert.Equal(t, map[string]interface{}{}, strMap.MergeNub(nil).M())
	}
	{
		data := StrMap(map[string]interface{}{"1": "one"})
		expected := map[string]interface{}{"1": "one"}
		assert.Equal(t, expected, NewStrMap().MergeNub(data).M())
	}
	{
		strMap := StrMap(map[string]interface{}{
			"1": "one", "2": "three", "3": "four",
		})
		data1 := StrMap(map[string]interface{}{"2": "two"})
		data2 := StrMap(map[string]interface{}{"3": "three"})
		data3 := StrMap(map[string]interface{}{"4": "four"})
		expected := map[string]interface{}{
			"1": "one", "2": "two", "3": "three", "4": "four",
		}
		assert.Equal(t, expected, strMap.MergeNub(data1, data2, data3).M())
	}
}

func TestStrMapSlice2(t *testing.T) {
	{
		data := map[string]interface{}{
			"test1": "foobar",
		}
		var expected []interface{}
		assert.Equal(t, expected, StrMap(data).Slice("test1"))
	}
	{
		data := map[string]interface{}{
			"test1": []interface{}{"foobar"},
		}
		expected := []interface{}{"foobar"}
		assert.Equal(t, expected, StrMap(data).Slice("test1"))
	}
	{
		data := map[string]interface{}{
			"test1": map[string]interface{}{
				"test2": []interface{}{"foobar"},
			},
		}
		expected := []interface{}{"foobar"}
		assert.Equal(t, expected, StrMap(data).Slice("test1.test2"))
	}
}

func TestStrMapStr(t *testing.T) {
	{
		target := NewStrMap().Add("test1", "foobar")
		assert.Equal(t, "foobar", target.Str("test1").M())
	}
	{
		target := NewStrMap().Add("test1", NewStrMap().Add("test2", "foo2"))
		assert.Equal(t, "foo2", target.Str("test1.test2").M())
	}
	{
		target := NewStrMap().Add("test1", NewStrMap().Add("tes2", "foo2"))
		assert.Equal(t, "", target.Str("test1.test2").M())
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

func TestStrMapStrMapByName(t *testing.T) {
	{
		data := map[string]interface{}{
			"foo": "one",
			"releases": []interface{}{
				map[string]interface{}{"name": "foo1"},
				map[string]interface{}{"name": "foo2"},
				map[string]interface{}{"name": "foo3"},
			},
		}
		expected := map[string]interface{}{"name": "foo2"}
		result := StrMap(data).StrMapByName("releases", "name", "foo2")
		assert.Equal(t, expected, result.M())
	}
	{
		rawYAMl := `releases:
- name: common
  chart: recurly/common:latest
- name: environment-services
  chart: recurly/environment-services:latest
  import-values: [tld]`

		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAMl), &data)
		expected := map[string]interface{}{
			"name":          "environment-services",
			"chart":         "recurly/environment-services:latest",
			"import-values": []interface{}{"tld"},
		}
		target := StrMap(data).StrMapByName("releases", "name", "environment-services").M()
		assert.Equal(t, expected, target)
	}
}

func TestStrMapStrMapSlice(t *testing.T) {
	{
		slice := StrMap(map[string]interface{}{
			"test1": []map[string]interface{}{
				{"1": interface{}("one")},
				{"2": interface{}("two")},
			},
		})
		expected := NewStrMapSlice()
		assert.Equal(t, expected, slice.StrMapSlice("1"))
	}
	{
		rawYAML := `test1:
  - 1: one
  - 2: two
`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
		})
		assert.Equal(t, expected, StrMap(data).StrMapSlice("test1"))
	}
	{
		slice := StrMap(map[string]interface{}{
			"test1": []map[string]interface{}{
				{"1": interface{}("one")},
				{"2": interface{}("two")},
			},
		})
		expected := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
		})
		assert.Equal(t, expected, slice.StrMapSlice("test1"))
	}
	{
		rawYAML := `test1:
  test2:
    - 1: one
    - 2: two
`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := StrMapSlice([]map[string]interface{}{
			{"1": "one"},
			{"2": "two"},
		})
		assert.Equal(t, expected, StrMap(data).StrMapSlice("test1.test2"))
	}
}

func TestStrSlice2(t *testing.T) {
	{
		data := map[string]interface{}{
			"test1": "foobar",
		}
		var expected []string
		assert.Equal(t, expected, StrMap(data).StrSlice("test1"))
	}
	{
		data := map[string]interface{}{
			"test1": []interface{}{"foobar"},
		}
		expected := []string{"foobar"}
		assert.Equal(t, expected, StrMap(data).StrSlice("test1"))
	}
	{
		data := map[string]interface{}{
			"test1": map[string]interface{}{
				"test2": []interface{}{"foobar"},
			},
		}
		expected := []string{"foobar"}
		assert.Equal(t, expected, StrMap(data).StrSlice("test1.test2"))
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
