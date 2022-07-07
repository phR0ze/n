package n

import (
	"github.com/phR0ze/n/pkg/enc/json"
	yaml_enc "github.com/phR0ze/n/pkg/enc/yaml"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// StringMap implements the IMap interface providing a generic way to work with map types
// including convenience methods on par with rapid development languages. This type is
// also specifically designed to handle ordered YAML constructs to work with YAML files
// with minimal structural changes i.e. no mass sorting changes.
type StringMap yaml.MapSlice

//type StringMap map[string]interface{}

// M is an alias to NewStringMap
func M(obj interface{}) *StringMap {
	return ToStringMap(obj)
}

// MV is an alias to NewStringMapV
func MV(m ...map[string]interface{}) *StringMap {
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
	ks := ToStrs(keys)
	if len(ks) == 0 {
		return true
	}
	for _, k := range ks {
		if p.Exists(k) {
			return true
		}
	}
	return false
}

// Clear modifies this Map to clear out all key-value pairs and returns a reference to this Map.
func (p *StringMap) Clear() IMap {
	if p == nil {
		p = NewStringMapV()
	} else if len(*p) > 0 {
		*p = *NewStringMapV()
	}
	return p
}

// Copy returns a new Map with the indicated key-value pairs copied from this Map or all if not given.
func (p *StringMap) Copy(keys ...interface{}) (new IMap) {
	val := NewStringMapV()
	if p == nil || len(*p) == 0 {
		return val
	}

	// Copy target keys or all keys
	ks := ToStrs(keys)
	if len(ks) == 0 {
		(*val) = append(*val, (*p)...)
	} else {
		for _, k := range ks {
			for i := 0; i < len(*p); i++ {
				if k == ToString((*p)[i].Key) {
					(*val) = append(*val, (*p)[i])
					break
				}
			}
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
	for i := 0; i < len(*p); i++ {
		if k == ToString((*p)[i].Key) {
			val.o = (*p)[i].Value
			if i+1 < len(*p) {
				*p = append((*p)[:i], (*p)[i+1:]...)
			} else {
				*p = (*p)[:i]
			}
			break
		}
	}
	return
}

// DeleteM modifies this Map to delete the indicated key-value pair and returns a reference to this Map rather than the key-value pair.
func (p *StringMap) DeleteM(key interface{}) IMap {
	p.Delete(key)
	return p
}

// Dump convert the StringMap into a pretty printed yaml string
func (p *StringMap) Dump() (pretty string) {
	if p == nil {
		return
	}
	if yml, err := yaml.Marshal(yaml.MapSlice(*p)); err == nil {
		pretty = string(yml)
	}
	return
}

// Exists checks if the given key exists in this Map.
func (p *StringMap) Exists(key interface{}) bool {
	return !p.Get(key).Nil()
}

// G returns the underlying data structure as a Go type.
func (p *StringMap) G() map[string]interface{} {
	if p == nil {
		return map[string]interface{}{}
	}
	// TODO: fix return map[string]interface{}(*p)
	return map[string]interface{}{}
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
	for _, x := range *p {
		if k == ToString(x.Key) {
			val.o = x.Value
			break
		}
	}
	return
}

// Inject sets the value for the given key location, using jq type selectors. Returns a reference to this Map.
func (p *StringMap) Inject(key string, val interface{}) IMap {
	m, _ := p.InjectE(key, val)
	return m
}

// InjectE sets the value for the given key location, using jq type selectors. Returns a reference to this Map.
func (p *StringMap) InjectE(key string, val interface{}) (m IMap, err error) {
	if p == nil {
		p = NewStringMapV()
	}
	m = p

	// Process keys from left to right
	var keys *StringSlice
	if keys, err = KeysFromSelector(key); err != nil {
		return
	}

	// Inject at root as no keys were given
	if !keys.Any() {
		if x, e := ToStringMapE(val); e == nil {
			for i := 0; i < len(*x); i++ {
				for j := 0; j < len(*p); j++ {
					if ToString((*x)[i].Key) == ToString((*p)[j].Key) {
						(*p)[j].Value = (*x)[i].Value
						break
					}
				}
			}
		} else {
			err = errors.Errorf("invalid selector for the type of value given, '%T'", val)
			return
		}
	}

	// Inject at given selector location
	obj := &Object{o: p}
	for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
		key := ko.ToStr()

		// All array selector is handled on previous loop
		if key.A() == "[]" {
			break
		}

		switch o := obj.o.(type) {

		// Identifier Index: .foo, .foo.bar
		case map[string]interface{}, *StringMap:
			x := ToStringMap(o)

			// Continue to drill or update and done
			create := false
			if keys.Any() && keys.First().A() != "[]" {
				if v := x.Get(key); !v.Nil() {
					if YAMLCont(v.O()) {
						obj.o = v.O()
					} else {
						create = true
					}
				} else {
					create = true
				}
			} else {
				x.Set(key.A(), val)
			}

			// Doesn't exist so create and drill further
			if create {
				x.Set(key.A(), map[string]interface{}{})
				obj.o = x.Get(key).O()
			}

		// Array Index/Iterator: .[2], .[-1], .[], .[key==val]
		case []interface{}:
			x := ToInterSlice(o)

			// Get array selectors
			var i int
			var k, v string
			if i, k, v, err = IdxFromSelector(key.A(), len(o)); err != nil {
				obj.o = nil
				return
			}

			// Select by key==value, e.g. .[k==v]
			if k != "" && v != "" {
				idx := -1
				x.EachIE(func(i int, y O) error {
					if hit := ToStringMap(y).Get(k); !hit.Nil() && hit.A() == v {
						idx = i
						return Break
					}
					return nil
				})
				i = idx // reuse code below for indexing
			}

			// Index in if the value is a valid integer, e.g. .[2], .[-1]
			// -1 indicates all should be selected.
			if i != -1 {
				if keys.Any() {
					obj.o = (*x)[i]
				} else {
					x.Set(i, val)
				}
			}
		}
	}
	if err != nil {
		m = nil
	}
	return
}

// Keys returns all the keys in this Map as a ISlice of the key type.
func (p *StringMap) Keys() ISlice {
	keys := NewStringSliceV()
	if p != nil {
		for i := 0; i < len(*p); i++ {
			*keys = append(*keys, ToString((*p)[i].Key))
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

// M is an alias to ToStringMap
func (p *StringMap) M() (m *StringMap) {
	return ToStringMap(p)
}

// MG is an alias ToStringMapG
func (p *StringMap) MG() (m map[string]interface{}) {
	return p.O().(map[string]interface{})
}

// Merge modifies this Map by overriding its values at location with the given map where they both exist and returns a reference to this Map.
// Converting all string maps into *StringMap instances.
func (p *StringMap) Merge(m IMap, location ...string) IMap {
	// p2 := p

	// // 1. Handle location if given
	// key := ""
	// if len(location) > 0 {
	// 	key = location[0]
	// 	var val interface{}
	// 	val = *p

	// 	// Process keys from left to right
	// 	keys, err := KeysFromSelector(key)
	// 	if err == nil {
	// 		for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
	// 			key := ko.ToString()
	// 			m := ToStringMap(val)

	// 			// Set a new map as the value if not a map
	// 			v := m.Get(key)
	// 			if !v.Nil() {
	// 				if !ToStringMap(v.O()).Any() {
	// 					m.Set(key, map[string]interface{}{})
	// 				}
	// 			} else {
	// 				m.Set(key, map[string]interface{}{})
	// 			}
	// 			val = v.O()
	// 		}
	// 	}
	// 	p2 = ToStringMap(val)
	// }

	// // 2. Merge at location
	// x, err := ToStringMapE(m)
	// switch {
	// case p2 == nil && (err != nil || m == nil):
	// 	return NewStringMapV()
	// case p2 == nil:
	// 	return x
	// case err != nil || m == nil:
	// 	return p2
	// }

	// for _, o := range *x {
	// 	var av, bv interface{}

	// 	// Ensure b value is n type
	// 	if val, ok := o.(map[string]interface{}); ok {
	// 		bv = ToStringMap(val)
	// 	} else {
	// 		bv = v
	// 	}

	// 	// a doesn't have the key so just set b's value
	// 	if val, exists := (*p2)[k]; !exists {
	// 		(*p2)[k] = bv
	// 	} else {
	// 		if _val, ok := val.(map[string]interface{}); ok {
	// 			av = ToStringMap(_val)
	// 		} else {
	// 			av = val
	// 		}

	// 		if bc, ok := bv.(*StringMap); ok {
	// 			if ac, ok := av.(*StringMap); ok {
	// 				// a and b both contain the key and are both submaps so recurse
	// 				(*p2)[k] = ac.Merge(bc)
	// 			} else {
	// 				// a is not a map so just override with b
	// 				(*p2)[k] = bv
	// 			}
	// 		} else {
	// 			// b is not a map so just override a, no need to recurse
	// 			(*p2)[k] = bv
	// 		}
	// 	}
	// }

	return p
}

// MergeG modifies this Map by overriding its values with the given map location where they both exist and returns the Go type
func (p *StringMap) MergeG(m IMap, location ...string) map[string]interface{} {
	// p2 := p

	// // 1. Handle location if given
	// key := ""
	// if len(location) > 0 {
	// 	key = location[0]
	// 	var val interface{}
	// 	val = *p

	// 	// Process keys from left to right
	// 	keys, err := KeysFromSelector(key)
	// 	if err == nil {
	// 		for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
	// 			key := ko.ToString()
	// 			m := ToStringMap(val)

	// 			// Set a new map as the value if not a map
	// 			if v, ok := (*m)[key]; ok {
	// 				if !ToStringMap(v).Any() {
	// 					m.Set(key, map[string]interface{}{})
	// 				}
	// 			} else {
	// 				m.Set(key, map[string]interface{}{})
	// 			}
	// 			val = (*m)[key]
	// 		}
	// 	}
	// 	p2 = ToStringMap(val)
	// }

	// // 2. Merge at location
	// x, err := ToStringMapE(m)
	// switch {
	// case p2 == nil && (err != nil || m == nil):
	// 	return NewStringMapV().G()
	// case p2 == nil:
	// 	return x.G()
	// case err != nil || m == nil:
	// 	return p2.G()
	// }

	// // Call type specific function helper
	// *p2 = MergeStringMap(*p2, *x)
	return p.G()
}

// O returns the underlying data structure as is.
func (p *StringMap) O() interface{} {
	// if p == nil {
	// 	return map[string]interface{}{}
	// }
	return map[string]interface{}{}
}

// Query returns the value at the given key location, using jq type selectors. Returns empty *Object if not found.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) Query(key string) (val *Object) {
	val, _ = p.QueryE(key)
	return val
}

// QueryE returns the value at the given key location, using jq type selectors. Returns empty *Object if not found.
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
			val.o = m.Get(key).O()

		// Array Index/Iterator: .[2], .[-1], .[], .[key==val]
		case []interface{}:

			// Get array selectors
			var i int
			var k, v string
			if i, k, v, err = IdxFromSelector(key.A(), len(x)); err != nil {
				val.o = nil
				return
			}

			// Select by key==value, e.g. .[k==v]
			if k != "" && v != "" {
				m := Slice(x).Select(func(x O) bool {
					return ToStringMap(x).Get(k).A() == v
				})
				if m.Any() {
					val.o = m.First().o
				}
			}

			// Index in if the value is a valid integer, e.g. .[2], .[-1]
			// -1 indicates all should be selected.
			if i != -1 {
				if val.o = Slice(x).At(i).o; val.Nil() {
					err = errors.Errorf("invalid array index %v", i)
					val.o = nil
					return
				}
			}
		}
	}
	return
}

// Remove modifies this Map to delete the given key location, using jq type selectors
// and returns a reference to this Map rather than the deleted value.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) Remove(key string) IMap {
	_, _ = p.RemoveE(key)
	return p
}

// RemoveE modifies this Map to delete the given key location, using jq type selectors
// and returns a reference to this Map rather than the deleted value.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) RemoveE(key string) (m IMap, err error) {
	if p == nil {
		p = NewStringMapV()
	}
	m = p

	// Process keys from left to right
	var keys *StringSlice
	if keys, err = KeysFromSelector(key); err != nil {
		return
	}

	// Inject at given selector location
	pk := ""
	obj, pobj := interface{}(p), interface{}(p)
	for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
		key := ko.ToStr()

		switch o := obj.(type) {

		// Identifier Index: .foo, .foo.bar
		case map[string]interface{}, *StringMap:
			x := ToStringMap(o)

			// Continue to drill or remove and done
			done := false
			if keys.Any() {
				if v := x.Get(key); !v.Nil() {
					if YAMLCont(v.O()) {
						pobj = obj
						pk = key.A()
						obj = v.O()
					} else {
						done = true
					}
				} else {
					done = true
				}
			} else {
				x.Delete(key.A())
			}

			// Key doesn't exist or is invalid type
			if done {
				return
			}

		// Array Index/Iterator: .[2], .[-1], .[], .[key==val]
		case []interface{}:
			x := ToInterSlice(o)

			// Get array selectors
			var i int
			var k, v string
			if i, k, v, err = IdxFromSelector(key.A(), len(o)); err != nil {
				obj = nil
				return
			}

			// Select by key==value, e.g. .[k==v]
			if k != "" && v != "" {
				idx := -1
				x.EachIE(func(i int, y O) error {
					if hit := ToStringMap(y).Get(k); !hit.Nil() && hit.A() == v {
						idx = i
						return Break
					}
					return nil
				})
				if idx == -1 {
					return
				}
				i = idx // reuse code below for indexing
			}

			// Index in if the value is a valid integer, e.g. .[2], .[-1]
			// -1 indicates all should be selected.
			if i == -1 && keys.Any() {
				childKeys := keys.Copy()
				for _, elem := range *x {
					if _, err = ToStringMap(elem).RemoveE(childKeys.Join(".").A()); err != nil {
						return
					}
				}
				keys.Clear()
			} else {
				if keys.Any() {
					pobj = obj
					pk = key.A()
					obj = (*x)[i]
				} else {
					ToStringMap(pobj).Set(pk, x.DropAt(i).O())
				}
			}
		}
	}
	if err != nil {
		m = nil
	}
	return
}

// Set the value for the given key to the given val. Returns true if the key did not yet exists in this Map.
func (p *StringMap) Set(key, val interface{}) (new bool) {
	if p == nil {
		return
	}
	k := ToString(key)
	for i := 0; i < len(*p); i++ {
		if k == ToString((*p)[i].Key) {
			new = false
			(*p)[i].Value = val
			break
		}
	}
	new = true
	*p = append(*p, yaml.MapItem{k, val})
	return
}

// SetM the value for the given key to the given val creating map if necessary. Returns a reference to this Map.
func (p *StringMap) SetM(key, val interface{}) IMap {
	if p == nil {
		p = NewStringMapV()
	}
	p.Set(key, val)
	return p
}

// ToStringMap converts the map to a *StringMap
func (p *StringMap) ToStringMap() (m *StringMap) {
	return ToStringMap(p)
}

// ToStringMapG converts the map to a Golang map[string]interface{}
func (p *StringMap) ToStringMapG() (m map[string]interface{}) {
	return p.O().(map[string]interface{})
}

// YAML converts the Map into a YAML string
func (p *StringMap) YAML() (data string) {
	_data, err := yaml.Marshal(p.G())
	if err != nil {
		return
	}
	data = string(_data)
	return
}

// YAMLE converts the Map into a YAML string
func (p *StringMap) YAMLE() (data string, err error) {
	var _data []byte
	if _data, err = yaml.Marshal(p.G()); err != nil {
		err = errors.Wrapf(err, "failed to marshal map[string]interface{}")
		return
	}
	data = string(_data)
	return
}

// WriteJSON converts the *StringMap into a map[string]interface{} then calls
// json.WriteJSON on it to write it out to disk.
func (p *StringMap) WriteJSON(filename string) (err error) {
	return json.WriteJSON(filename, p.G())
}

// WriteYAML converts the *StringMap into a map[string]interface{} then calls
// yaml.WriteYAML on it to write it out to disk.
func (p *StringMap) WriteYAML(filename string) (err error) {
	return yaml_enc.WriteYAML(filename, p.G())
}
