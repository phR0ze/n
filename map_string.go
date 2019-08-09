package n

import (
	"strconv"
	"strings"

	"github.com/phR0ze/n/pkg/sys"

	"github.com/pkg/errors"
)

// StringMap implements the Map interface providing a generic way to work with map types
// including convenience methods on par with rapid development languages. This type is
// also specifically designed to handle YAML constructs.
type StringMap map[string]interface{}

// Yaml is an alias to NewStringMap
func Yaml(obj interface{}) *StringMap {
	return ToStringMap(obj)
}

// YamlV is an alias to NewStringMapV
func YamlV(m ...map[string]interface{}) *StringMap {
	return NewStringMapV(m...)
}

// NewStringMap converts the given interface{} into a StringMap
func NewStringMap(obj interface{}) *StringMap {
	return ToStringMap(obj)
}

// NewStringMapV creates a new empty StringMap if nothing given else simply
// casts the given map to StringMap.
func NewStringMapV(m ...map[string]interface{}) *StringMap {
	var new StringMap
	if len(m) == 0 {
		new = StringMap{}
	} else {
		new = *ToStringMap(m[0])
	}
	return &new
}

// Any tests if this Map is not empty or optionally if it contains any of the given variadic keys.
func (p *StringMap) Any(keys ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	if len(keys) == 0 {
		return true
	}
	for i := 0; i < len(keys); i++ {
		key := ToString(keys[i])
		if _, ok := (*p)[key]; ok {
			return true
		}
	}
	return false
}

// Clear modifies this Map to clear out all key-value pairs and returns a reference to this Map.
func (p *StringMap) Clear() Map {
	if p == nil {
		p = NewStringMapV()
	} else if len(*p) > 0 {
		*p = *NewStringMapV()
	}
	return p
}

// Copy returns a new Map with the indicated key-value pairs copied from this Map or all if not given.
func (p *StringMap) Copy(keys ...interface{}) (new Map) {
	val := NewStringMapV()
	if p == nil || len(*p) == 0 {
		return val
	}

	// Copy target keys or all keys
	ks := ToStringSliceG(keys)
	if len(ks) == 0 {
		for k := range *p {
			(*val)[k] = (*p)[k]
		}
	} else {
		for _, k := range ks {
			(*val)[k] = (*p)[k]
		}
	}
	return val
}

// Delete modifies this Map to delete the indicated key-value pair and returns the value from the Map.
func (p *StringMap) Delete(key interface{}) (val *Object) {
	val = &Object{}
	if p == nil {
		return
	}
	k := ToString(key)
	if v, ok := (*p)[k]; ok {
		val.o = v
		delete(*p, k)
	}
	return
}

// DeleteM modifies this Map to delete the indicated key-value pair and returns a reference to this Map rather than the key-value pair.
func (p *StringMap) DeleteM(key interface{}) Map {
	if p == nil {
		return p
	}
	k := ToString(key)
	delete(*p, k)
	return p
}

// Exists checks if the given key exists in this Map.
func (p *StringMap) Exists(key interface{}) bool {
	if p == nil {
		return false
	}
	k := ToString(key)
	if _, ok := (*p)[k]; ok {
		return true
	}
	return false
}

// G returns the underlying data structure as a Go type.
func (p *StringMap) G() map[string]interface{} {
	if p == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}(*p)
}

// Generic returns true if the underlying implementation uses reflection
func (p *StringMap) Generic() bool {
	return false
}

// Get returns the value at the given key location. Returns empty *Object if not found.
func (p *StringMap) Get(key interface{}) (val *Object) {
	val = &Object{}
	if p == nil {
		return
	}
	k := ToString(key)
	if v, ok := (*p)[k]; ok {
		val.o = v
	}
	return
}

// Keys returns all the keys in this Map as a Slice of the key type.
func (p *StringMap) Keys() Slice {
	keys := NewStringSliceV()
	if p != nil {
		for key := range *p {
			*keys = append(*keys, key)
		}
	}
	return keys
}

// Len returns the number of elements in this Map.
func (p *StringMap) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// Merge modifies this Map by overriding its values at location with the given map where they both exist and returns a reference to this Map.
// Converting all string maps into *StringMap instances.
func (p *StringMap) Merge(m Map, location ...string) Map {
	p2 := p

	// 1. Handle location if given
	key := ""
	if len(location) > 0 {
		key = location[0]
		var val interface{}
		val = *p

		// Process keys from left to right
		keys, err := KeysFromSelector(key)
		if err == nil {
			for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
				key := ko.ToString()
				m := ToStringMap(val)

				// Set a new map as the value if not a map
				if v, ok := (*m)[key]; ok {
					if !ToStringMap(v).Any() {
						m.Set(key, map[string]interface{}{})
					}
				} else {
					m.Set(key, map[string]interface{}{})
				}
				val = (*m)[key]
			}
		}
		p2 = ToStringMap(val)
	}

	// 2. Merge at location
	x, err := ToStringMapE(m)
	switch {
	case p2 == nil && (err != nil || m == nil):
		return NewStringMapV()
	case p2 == nil:
		return x
	case err != nil || m == nil:
		return p2
	}

	for k, v := range *x {
		var av, bv interface{}

		// Ensure b value is n type
		if val, ok := v.(map[string]interface{}); ok {
			bv = ToStringMap(val)
		} else {
			bv = v
		}

		// a doesn't have the key so just set b's value
		if val, exists := (*p2)[k]; !exists {
			(*p2)[k] = bv
		} else {
			if _val, ok := val.(map[string]interface{}); ok {
				av = ToStringMap(_val)
			} else {
				av = val
			}

			if bc, ok := bv.(*StringMap); ok {
				if ac, ok := av.(*StringMap); ok {
					// a and b both contain the key and are both submaps so recurse
					(*p2)[k] = ac.Merge(bc)
				} else {
					// a is not a map so just override with b
					(*p2)[k] = bv
				}
			} else {
				// b is not a map so just override a, no need to recurse
				(*p2)[k] = bv
			}
		}
	}

	return p
}

// MergeG modifies this Map by overriding its values with the given map location where they both exist and returns the Go type
func (p *StringMap) MergeG(m Map, location ...string) map[string]interface{} {
	p2 := p

	// 1. Handle location if given
	key := ""
	if len(location) > 0 {
		key = location[0]
		var val interface{}
		val = *p

		// Process keys from left to right
		keys, err := KeysFromSelector(key)
		if err == nil {
			for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
				key := ko.ToString()
				m := ToStringMap(val)

				// Set a new map as the value if not a map
				if v, ok := (*m)[key]; ok {
					if !ToStringMap(v).Any() {
						m.Set(key, map[string]interface{}{})
					}
				} else {
					m.Set(key, map[string]interface{}{})
				}
				val = (*m)[key]
			}
		}
		p2 = ToStringMap(val)
	}

	// 2. Merge at location
	x, err := ToStringMapE(m)
	switch {
	case p2 == nil && (err != nil || m == nil):
		return NewStringMapV().G()
	case p2 == nil:
		return x.G()
	case err != nil || m == nil:
		return p2.G()
	}

	// Call type specific function helper
	*p2 = MergeStringMap(*p2, *x)
	return p.G()
}

// O returns the underlying data structure as is.
func (p *StringMap) O() interface{} {
	if p == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}(*p)
}

// Query returns the value at the given key location, using a jq type selectors. Returns empty *Object if not found.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) Query(key string) (val *Object) {
	val, _ = p.QueryE(key)
	return val
}

// QueryE returns the value at the given key location, using a jq type selectors. Returns empty *Object if not found.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) QueryE(key string) (val *Object, err error) {
	if p == nil || len(*p) == 0 {
		err = errors.Errorf("failed to query empty map")
		return
	}

	// Default object is self for identity case: .
	val = &Object{o: p}
	if p == nil {
		return
	}

	// Process keys from left to right
	var keys *StringSlice
	if keys, err = KeysFromSelector(key); err != nil {
		return
	}
	for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
		key := ko.ToStr()

		switch x := val.o.(type) {

		// Identifier Index: .foo, .foo.bar
		case map[string]interface{}, *StringMap:
			m := ToStringMap(x)
			val.o = (*m)[key.A()]

		// Array Index/Iterator: .[2], .[-1], .[], .[key==val]
		case []interface{}:

			// Empty list so return nil i.e. failed
			if len(x) == 0 {
				err = errors.Errorf("array is empty")
				val.o = nil
				return
			}

			// Continue if list is not empty
			if key.First().A() == "[" && key.Last().A() == "]" {

				// Trim off the indexer/selector brackets and check the indexer
				idx := key.TrimPrefix("[").TrimSuffix("]").A()
				if idx != "" {
					pieces := strings.Split(idx, "==")
					i, e := strconv.Atoi(idx)
					switch {

					// Select by key==value, e.g. .[k==v]
					case len(pieces) == 2:
						k, v := pieces[0], pieces[1]
						m := NewSlice(x).Select(func(x O) bool {
							return ToStringMap(x).Get(k).A() == v
						})
						if m.Any() {
							val.o = m.First().o
						}

					// Index in if the value is a valid integer, e.g. .[2], .[-1]
					case e == nil:
						if val.o = NewSlice(x).At(i).o; val.Nil() {
							err = errors.Errorf("invalid array index %v", i)
							val.o = nil
							return
						}
					}

					// Fall through to return all array elements
				}
			}
		}
	}
	return
}

// Remove modifies this Map to delete the given key location, using jq type selectors
// and returns a reference to this Map rather than the deleted value.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) Remove(key string) Map {
	_, _ = p.RemoveE(key)
	return p
}

// RemoveE modifies this Map to delete the given key location, using jq type selectors
// and returns a reference to this Map rather than the deleted value.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) RemoveE(key string) (m Map, err error) {
	m = p
	if p == nil || len(*p) == 0 {
		return
	}
	val := &Object{o: p}

	// Process keys from left to right
	var keys *StringSlice
	if keys, err = KeysFromSelector(key); err != nil {
		return
	}
	for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
		key := ko.ToStr()

		switch x := val.o.(type) {

		// Identifier Index: .foo, .foo.bar
		case map[string]interface{}, *StringMap:
			m := ToStringMap(x)
			if keys.Any() {
				val.o = (*m)[key.A()]
			} else {
				delete(*m, key.A())
			}

			// // Array Index/Iterator: .[2], .[-1], .[], .[key==val]
			// case []interface{}:

			// 	// Empty list so return nil i.e. failed
			// 	if len(x) == 0 {
			// 		err = errors.Errorf("array is empty")
			// 		val.o = nil
			// 		return
			// 	}

			// 	// Continue if list is not empty
			// 	if key.First().A() == "[" && key.Last().A() == "]" {

			// 		// Trim off the indexer/selector brackets and check the indexer
			// 		idx := key.TrimPrefix("[").TrimSuffix("]").A()
			// 		if idx != "" {
			// 			pieces := strings.Split(idx, "==")
			// 			i, e := strconv.Atoi(idx)
			// 			switch {

			// 			// Select by key==value, e.g. .[k==v]
			// 			case len(pieces) == 2:
			// 				k, v := pieces[0], pieces[1]
			// 				m := NewSlice(x).Select(func(x O) bool {
			// 					return ToStringMap(x).Get(k).A() == v
			// 				})
			// 				if m.Any() {
			// 					val.o = m.First().o
			// 				}

			// 			// Index in if the value is a valid integer, e.g. .[2], .[-1]
			// 			case e == nil:
			// 				if val.o = NewSlice(x).At(i).o; val.Nil() {
			// 					err = errors.Errorf("invalid array index %v", i)
			// 					val.o = nil
			// 					return
			// 				}
			// 			}

			// 			// Fall through to return all array elements
			// 		}
			// 	}
		}
	}
	return
}

// Set the value for the given key to the given val. Returns true if the key did not yet exists in this Map.
func (p *StringMap) Set(key, val interface{}) (new bool) {
	if p == nil {
		return
	}
	k := ToString(key)
	if _, ok := (*p)[k]; !ok {
		new = true
	}
	(*p)[k] = val
	return
}

// SetM the value for the given key to the given val creating map if necessary. Returns a reference to this Map.
func (p *StringMap) SetM(key, val interface{}) Map {
	if p == nil {
		p = NewStringMapV()
	}
	p.Set(key, val)
	return p
}

// ToStringMap converts the map to a *StringMap
func (p *StringMap) ToStringMap() (m *StringMap) {
	if p == nil {
		return NewStringMapV()
	}
	return p
}

// ToStringMapG converts the map to a google type map[string]interface{}
func (p *StringMap) ToStringMapG() (m map[string]interface{}) {
	return p.G()
}

// ToYaml is an alias to ToStringMap
func (p *StringMap) ToYaml() (m *StringMap) {
	if p == nil {
		return NewStringMapV()
	}
	return p
}

// ToYamlG is an alias to ToStringMapG
func (p *StringMap) ToYamlG() (m map[string]interface{}) {
	return p.G()
}

// WriteYaml converts the *StringMap into a map[string]interface{} then calls
// sys.WriteYaml on it to write it out to disk.
func (p *StringMap) WriteYaml(filename string) (err error) {
	return sys.WriteYaml(filename, p.G())
}
