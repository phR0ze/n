package nub

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/ghodss/yaml"
)

// KeyValuePair similar to C# for iterator over maps
type KeyValuePair struct {
	Key   interface{}
	Value interface{}
}

// M provides a new empty Queryable map
func M() *Queryable {
	obj := map[interface{}]interface{}{}
	ref := reflect.ValueOf(obj)
	return &Queryable{O: obj, ref: &ref, Iter: mapIter(ref, obj)}
}

func mapIter(ref reflect.Value, obj interface{}) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		keys := ref.MapKeys()
		return func() (item interface{}, ok bool) {
			if ok = i < len; ok {
				item = &KeyValuePair{
					Key:   keys[i].Interface(),
					Value: ref.MapIndex(keys[i]).Interface(),
				}
				i++
			}
			return
		}
	}
}

// StrMap materializes the result into a string to interface map
func (q *Queryable) StrMap() map[string]interface{} {
	result := map[string]interface{}{}
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		pair := x.(*KeyValuePair)
		result[pair.Key.(string)] = pair.Value
	}
	return result
}

// TODO: Need to refactor below here

//--------------------------------------------------------------------------------------------------
// StrMap Nub
//--------------------------------------------------------------------------------------------------
type strMapNub struct {
	raw map[string]interface{}
}

// NewStrMap creates a new nub
func NewStrMap() *strMapNub {
	return &strMapNub{raw: map[string]interface{}{}}
}

// StrMap creates a new nub from the given string map
func StrMap(other map[string]interface{}) *strMapNub {
	return &strMapNub{raw: other}
}

// Load a yaml/json file as a str map
// returns nil on failure of any kind
func Load(target string) (result *strMapNub) {
	if yamlFile, err := ioutil.ReadFile(target); err == nil {
		result = NewStrMap()
		yaml.Unmarshal(yamlFile, &result.raw)
	}
	return result
}

// Add a new key value pair to the map
func (m *strMapNub) Add(key string, value interface{}) *strMapNub {
	switch x := value.(type) {
	case *strMapNub:
		m.raw[key] = x.raw
	default:
		m.raw[key] = value
	}
	return m
}

// Any checks if the map has anything in it
func (m *strMapNub) Any() bool {
	return len(m.raw) > 0
}

// Equals checks if the two maps are equal
func (m *strMapNub) Equals(other *strMapNub) bool {
	return reflect.DeepEqual(m, other)
}

// Len is a pass through to the underlying map
func (m *strMapNub) Len() int {
	return len(m.raw)
}

// M materializes object invoking deferred execution
func (m *strMapNub) M() map[string]interface{} {
	return m.raw
}

// Merge the other maps into this map with the first taking the lowest
// precedence and so on until the last takes the highest
func (m *strMapNub) Merge(items ...map[string]interface{}) *strMapNub {
	for i := range items {
		if items[i] != nil {
			m.raw = mergeMap(m.raw, items[i])
		}
	}
	return m
}

// Merge the other maps into this map with the first taking the lowest
// precedence and so on until the last takes the highest
func (m *strMapNub) MergeNub(items ...*strMapNub) *strMapNub {
	for i := range items {
		if items[i] != nil {
			m.raw = mergeMap(m.raw, items[i].raw)
		}
	}
	return m
}

// Slice returns a slice of interface{} from the given map using the given key
func (m *strMapNub) Slice(key string) (result []interface{}) {
	keys := Str(key).Split(".")
	if k, ok := keys.TakeFirst(); ok {
		if entry, exists := m.raw[k]; exists {
			switch x := entry.(type) {
			case map[string]interface{}:
				result = StrMap(x).Slice(keys.Join(".").M())
			case []map[string]interface{}:
				result = unCastStrMapSlice(x)
			case []interface{}:
				result = x
			}
		}
	}
	return
}

// Str returns a string from the given map using the given key
func (m *strMapNub) Str(key string) *strNub {
	result := NewStr()
	keys := Str(key).Split(".")
	if k, ok := keys.TakeFirst(); ok {
		if entry, exists := m.raw[k]; exists {
			switch v := entry.(type) {
			case string:
				result = Str(v)
			case map[string]interface{}:
				result = StrMap(v).Str(keys.Join(".").M())
			}
		}
	}
	return result
}

// StrMap returns a map of interface from the given map using the given key
func (m *strMapNub) StrMap(key string) *strMapNub {
	result := NewStrMap()

	keys := Str(key).Split(".")
	if k, ok := keys.TakeFirst(); ok {
		if entry, exists := m.raw[k]; exists {
			if v, ok := entry.(map[string]interface{}); ok {
				result.raw = v
				if keys.Len() != 0 {
					result = result.StrMap(keys.Join(".").M())
				}
			}
		}
	}
	return result
}

// StrMapByName returns a map of interface from the given map using the given key
func (m *strMapNub) StrMapByName(key, k, v string) *strMapNub {
	result := NewStrMap()
	slice := m.Slice(key)
	for i := range slice {
		if x, ok := slice[i].(map[string]interface{}); ok {
			if value, exists := x[k]; exists && value == v {
				result.raw = x
				break
			}
		}
	}
	return result
}

// StrMapSlice returns a slice of str map from the given map using the given key
func (m *strMapNub) StrMapSlice(key string) *strMapSliceNub {
	return castStrMapSlice(m.Slice(key))
}

// StrSlice returns a slice of strings from the given map using the given key
func (m *strMapNub) StrSlice(key string) (result []string) {
	items := m.Slice(key)
	for i := range items {
		result = append(result, fmt.Sprint(items[i]))
	}
	return
}

// castStrMapSlice returns a slice of str map from the given interface slice
func castStrMapSlice(items []interface{}) *strMapSliceNub {
	result := NewStrMapSlice()
	for i := range items {
		if x, ok := items[i].(map[string]interface{}); ok {
			result.Append(x)
		}
	}
	return result
}

// Merge b into a and returns the new modified a
// b takes higher precedence and will override a
func mergeMap(a, b map[string]interface{}) map[string]interface{} {
	switch {
	case (a == nil || len(a) == 0) && (b == nil || len(b) == 0):
		return map[string]interface{}{}
	case a == nil || len(a) == 0:
		return b
	case b == nil || len(b) == 0:
		return a
	}

	for k, bv := range b {
		if av, exists := a[k]; !exists {
			a[k] = bv
		} else if bc, ok := bv.(map[string]interface{}); ok {
			if ac, ok := av.(map[string]interface{}); ok {
				a[k] = mergeMap(ac, bc)
			} else {
				a[k] = bv
			}
		} else {
			a[k] = bv
		}
	}

	return a
}

// UnCastStrMapSlice casts the given slice to a slice of interface
func unCastStrMapSlice(items []map[string]interface{}) []interface{} {
	result := []interface{}{}
	for i := range items {
		result = append(result, items[i])
	}

	return result
}
