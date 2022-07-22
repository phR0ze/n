package n

import (
	"github.com/phR0ze/n/pkg/enc/json"
	yaml_enc "github.com/phR0ze/n/pkg/enc/yaml"
	yaml "github.com/phR0ze/yaml/v2"
	"github.com/pkg/errors"
)

// StringMap implements the IMap interface providing a generic way to work with map types
// including convenience methods on par with rapid development languages. This type is
// also specifically designed to handle ordered YAML constructs to work with YAML files
// with minimal structural changes i.e. no mass sorting changes.
type StringMap yaml.MapSlice

// M is an alias to NewStringMap
func M() *StringMap {
	return &StringMap{}
}

// MV is an alias to NewStringMapV
func MV(m ...interface{}) *StringMap {
	return NewStringMapV(m...)
}

// NewStringMap converts the given interface{} into a StringMap
func NewStringMap(obj interface{}) *StringMap {
	return ToStringMap(obj)
}

// NewStringMapV creates a new empty StringMap if nothing given else
// converts the given value into a StringMap.
func NewStringMapV(m ...interface{}) *StringMap {
	var new *StringMap
	if len(m) == 0 {
		new = &StringMap{}
	} else {
		new = ToStringMap(m[0])
	}
	return new
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

// Add the value to the map if it doesn't exist or update value if it does
func (p *StringMap) Add(key, val interface{}) *StringMap {
	p.Set(key, val)
	return p
}

// At gets the key value pair for the given index location
func (p *StringMap) At(i int) (key string, val *Object) {
	val = &Object{}
	if p == nil {
		return
	}
	if i = absIndex(len(*p), i); i == -1 {
		return
	}
	key = ToString((*p)[i].Key)
	val.o = (*p)[i].Value
	return
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
	if p == nil {
		return false
	}
	k := ToString(key)
	for i := 0; i < len(*p); i++ {
		if k == ToString((*p)[i].Key) {
			return true
		}
	}
	return false
}

// G returns the underlying data structure as a Go type.
func (p *StringMap) G() map[string]interface{} {
	val, _ := p.GE()
	return val
}

// GE returns the underlying data structure as a Go type
func (p *StringMap) GE() (val map[string]interface{}, err error) {
	val = map[string]interface{}{}
	if p == nil {
		return
	}
	for _, o := range *p {
		k := ToString(o.Key)
		v := DeReference(o.Value)
		switch x := v.(type) {
		case StringMap, yaml.MapSlice:
			var m map[string]interface{}
			if m, err = ToStringMap(x).GE(); err != nil {
				return
			}
			val[k] = m
		case []interface{}:
			for i := 0; i < len(x); i++ {
				switch x2 := x[i].(type) {
				case StringMap, yaml.MapSlice:
					var m map[string]interface{}
					if m, err = ToStringMap(x2).GE(); err != nil {
						return
					}
					x[i] = m
				}
			}
			val[k] = x
		default:
			val[k] = v
		}
	}
	return
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
	for i := 0; i < len(*p); i++ {
		if k == ToString((*p)[i].Key) {
			val.o = (*p)[i].Value
			break
		}
	}
	return
}

// Update sets the value for the given selector, using jq type selectors. Returns a reference to this Map.
func (p *StringMap) Update(selector string, val interface{}) IMap {
	m, _ := p.UpdateE(selector, val)
	return m
}

// UpdateE sets the value for the given selector, using jq type selectors. Returns a reference to this Map.
func (p *StringMap) UpdateE(selector string, val interface{}) (m IMap, err error) {
	if p == nil {
		p = NewStringMapV()
	}
	m = p
	val = convertValue(val)

	// Process keys from left to right
	var keys *StringSlice
	if keys, err = KeysFromSelector(selector); err != nil {
		return
	}

	// Merge at root as no keys were given
	if !keys.Any() {
		if x, e := ToStringMapE(val); e == nil {
			for i := 0; i < len(*x); i++ {
				found := false
				for j := 0; j < len(*p); j++ {
					if ToString((*x)[i].Key) == ToString((*p)[j].Key) {
						(*p)[j].Value = (*x)[i].Value
						found = true
						break
					}
				}
				if !found {
					p.Set((*x)[i].Key, (*x)[i].Value)
				}
			}
		} else {
			err = errors.Errorf("invalid selector for the type of value given, '%T'", val)
			return
		}
	}

	var pk interface{}
	var m1, m2 *StringMap
	cx, px := interface{}(p), interface{}(p)
	for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
		key := ko.ToStr()

		switch x := cx.(type) {

		// Identifier Index: .foo, .foo.bar
		case map[string]interface{}, *StringMap, StringMap, yaml.MapSlice:
			if m1, err = ToStringMapE(x); err != nil {
				return
			}

			// Continue to drill or update and done
			addOrReplace := false
			if keys.Any() {
				if v := m1.Get(key); !v.Nil() {
					if YAMLCont(v.O()) {
						cx = v.O()
						pk, px = key, m1
					} else {
						addOrReplace = true
					}
				} else {
					addOrReplace = true
				}
			} else {
				m1.Set(key, val)
				if px != nil && pk != nil {
					if v, ok := px.([]interface{}); ok {
						v[ToInt(pk)] = yaml.MapSlice(*m1)
					} else {
						if m2, err = ToStringMapE(px); err != nil {
							return
						}
						m2.Set(pk, m1)
					}
				}
			}
			if addOrReplace {
				m1.Set(key, M())
				cx = m1.Get(key).O()
				pk, px = key, m1
			}

		// Array Index/Iterator: .[2], .[-1], .[], .[key==val]
		case []interface{}:

			// Get array selectors
			var i int
			var k, v string
			if i, k, v, err = IdxFromSelector(key.A(), len(x)); err != nil {
				err = errors.Errorf("invalid array index selector %v", key.A())
				cx = nil
				return
			}

			// Select single element by key==value or index, e.g. .[k==v], [i]
			if (k != "" && v != "") || i != -1 {

				// Determine index
				if k != "" && v != "" {
					for j := 0; j < len(x); j++ {
						if m1, err = ToStringMapE(x[j]); err != nil {
							return
						}
						if m1.Get(k).A() == v {
							i = j
							break
						}
					}
				}

				// Move through or update index
				if keys.Any() {
					cx = x[i]
					pk, px = i, x
				} else {
					x[i] = val
					if px != nil && pk != nil {
						if v, ok := px.([]interface{}); ok {
							v[ToInt(pk)] = x
						} else {
							if m2, err = ToStringMapE(px); err != nil {
								return
							}
							m2.Set(pk, x)
						}
					}
				}
			}

			// Select all elements by .[] and translated to a -1 at this level
			if i == -1 {
				if keys.Any() {
					for j := 0; j < len(x); j++ {
						var o IMap
						if o, err = ToStringMap(x[j]).UpdateE(keys.Join(".").A(), val); err != nil {
							return
						}
						if m1, err = ToStringMapE(o); err != nil {
							return
						}
						x[j] = yaml.MapSlice(*m1)
					}
					keys.Clear()
				} else {
					for j := 0; j < len(x); j++ {
						x[j] = val
					}
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
	return p.G()
}

// Merge modifies this Map by overriding its values at selector with the given map
// where they both exist and returns a reference to this Map. Converting all string
// maps into *StringMap instances.
// Note: this function is unable to traverse through lists
func (p *StringMap) Merge(m IMap, selector ...string) IMap {
	x1 := p

	// 1. Select target if given
	var pk1, px1 interface{}
	if len(selector) > 0 {
		key := selector[0]
		cx1 := interface{}(*p)

		// Process keys from left to right
		keys, err := KeysFromSelector(key)
		if err == nil {
			for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
				key := ko.ToString()
				m := ToStringMap(cx1)

				// Set a new map as the value if not a map
				v := m.Get(key)
				if v.Nil() || !YAMLMap(v.O()) {
					m.Set(key, M())
					v = m.Get(key)
					if pk1 != nil && px1 != nil {
						ToStringMap(px1).Set(pk1, m)
					}
				}
				pk1, px1 = key, cx1
				cx1 = v.O()
			}
		}
		x1 = ToStringMap(cx1)
	}

	// 2. Merge at selection or root
	x2, err := ToStringMapE(m)
	switch {
	case x1 == nil && (err != nil || m == nil):
		return M()
	case x1 == nil:
		return x2
	case err != nil || m == nil:
		return x1
	}

	for _, o := range *x2 {
		k := o.Key
		v2 := DeReference(o.Value)

		// x1 doesn't have the key so just set x2's value
		v1 := x1.Get(k)
		if v1.Nil() {
			x1.Set(k, v2)
		} else {
			if YAMLMap(v1.o) && YAMLMap(v2) {
				// x1 and x2 both contain the key and are both map types so recurse
				x1.Set(k, ToStringMap(v1.o).Merge(ToStringMap(v2)))
			} else {
				// x1 or x2 is not a map so just override
				x1.Set(k, v2)
			}
		}
	}
	if len(selector) > 0 {
		p.Update(selector[0], x1)
	}

	return p
}

// Merge modifies this Map by overriding its values at selector with the given map
// where they both exist and returns a reference to this Map. Converting all string
// maps into *StringMap instances.
// Note: this function is unable to traverse through lists
func (p *StringMap) MergeG(m IMap, selector ...string) map[string]interface{} {
	return p.Merge(m, selector...).MG()
}

// O returns the underlying data structure as is.
func (p *StringMap) O() interface{} {
	return p.G()
}

// Query returns the value for the given selector, using jq type selectors. Returns empty *Object if not found.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) Query(selector string) (val *Object) {
	val, _ = p.QueryE(selector)
	return val
}

// QueryE returns the value for the given selector, using jq type selectors. Returns empty *Object if not found.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) QueryE(selector string) (val *Object, err error) {
	if p == nil || len(*p) == 0 {
		err = errors.Errorf("failed to query empty map")
		return
	}

	// Default object is self for identity case: .
	val = &Object{o: p}

	// Process keys from left to right
	var keys *StringSlice
	if keys, err = KeysFromSelector(selector); err != nil {
		return
	}
	for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
		key := ko.ToStr()

		switch x := val.o.(type) {

		// Identifier Index: .foo, .foo.bar
		case map[string]interface{}, *StringMap, StringMap, yaml.MapSlice:
			m := ToStringMap(x)
			val.o = m.Get(key).O()

		// Array Index/Iterator: .[2], .[-1], .[], .[key==val]
		case []interface{}:

			// Get array selectors
			var i int
			var k, v string
			if i, k, v, err = IdxFromSelector(key.A(), len(x)); err != nil {
				err = errors.Errorf("invalid array index selector %v", key.A())
				val.o = nil
				return
			}

			// Select by key==value, e.g. .[k==v]
			if k != "" && v != "" {
				for j := 0; j < len(x); j++ {
					m := ToStringMap(x[j])
					if m.Get(k).A() == v {
						val.o = m
					}
				}
			}

			// Index in if the value is a valid integer
			if i != -1 {
				val.o = x[i]
			}
		}
	}
	return
}

// Remove modifies this Map to delete the given key location, using jq type selectors
// and returns a reference to this Map rather than the deleted value.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) Remove(selector string) IMap {
	_, _ = p.RemoveE(selector)
	return p
}

// RemoveE modifies this Map to delete the given key location, using jq type selectors
// and returns a reference to this Map rather than the deleted value.
// see dot notation from https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
func (p *StringMap) RemoveE(selector string) (m IMap, err error) {
	if p == nil {
		p = NewStringMapV()
	}
	m = p

	// Process keys from left to right
	var keys *StringSlice
	if keys, err = KeysFromSelector(selector); err != nil {
		return
	}
	var pk interface{}
	var m1, m2 *StringMap
	cx, px := interface{}(p), interface{}(p)
	for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
		key := ko.ToStr()

		switch x := cx.(type) {

		// Identifier Index: .foo, .foo.bar
		case map[string]interface{}, *StringMap, StringMap, yaml.MapSlice:
			if m1, err = ToStringMapE(x); err != nil {
				return
			}

			// Continue to drill or remove and done
			if keys.Any() {
				cx = m1.Get(key).O()
				pk = key
				px = m1
			} else {
				m1.DeleteM(key)
				if px != nil && pk != nil {
					if v, ok := px.([]interface{}); ok {
						v[ToInt(pk)] = yaml.MapSlice(*m1)
					} else {
						if m2, err = ToStringMapE(px); err != nil {
							return
						}
						m2.Set(pk, m1)
					}
				}
			}

		// Array Index/Iterator: .[2], .[-1], .[], .[key==val]
		case []interface{}:

			// Get array selectors
			var i int
			var k, v string
			if i, k, v, err = IdxFromSelector(key.A(), len(x)); err != nil {
				err = errors.Errorf("invalid array index selector %v", key.A())
				cx = nil
				return
			}

			// Select single element by key==value or index, e.g. .[k==v], [i]
			if (k != "" && v != "") || i != -1 {

				// Determine index
				if k != "" && v != "" {
					for j := 0; j < len(x); j++ {
						if m1, err = ToStringMapE(x[j]); err != nil {
							return
						}
						if m1.Get(k).A() == v {
							i = j
							break
						}
					}
				}

				// Move through or remove index
				if keys.Any() {
					cx = x[i]
					pk = i
					px = x
				} else {
					if i+1 < len(x) {
						x = append(x[:i], x[i+1:]...)
					} else {
						x = x[:i]
					}
					if px != nil && pk != nil {
						if v, ok := px.([]interface{}); ok {
							v[ToInt(pk)] = x
						} else {
							if m2, err = ToStringMapE(px); err != nil {
								return
							}
							m2.Set(pk, x)
						}
					}
				}
			}

			// Select all elements by .[] and translated to a -1 at this level
			if i == -1 {
				if keys.Any() {
					for j := 0; j < len(x); j++ {
						var o IMap
						if o, err = ToStringMap(x[j]).RemoveE(keys.Join(".").A()); err != nil {
							return
						}
						if m1, err = ToStringMapE(o); err != nil {
							return
						}
						x[j] = yaml.MapSlice(*m1)
					}
					keys.Clear()
				} else {
					if px != nil && pk != nil {
						if v, ok := px.([]interface{}); ok {
							v[ToInt(pk)] = []interface{}{}
						} else {
							if m2, err = ToStringMapE(px); err != nil {
								return
							}
							m2.Set(pk, []interface{}{})
						}
					}
				}
			}
		}
	}
	if err != nil {
		m = nil
	}
	return
}

// Set the value for the given key to the given val. Returns true if the key did not yet exist in this Map.
func (p *StringMap) Set(key, val interface{}) (new bool) {
	if p == nil {
		return
	}
	new = true
	k := ToString(key)
	for i := 0; i < len(*p); i++ {
		if k == ToString((*p)[i].Key) {
			new = false
			(*p)[i].Value = convertValue(val)
			break
		}
	}
	if new {
		*p = append(*p, yaml.MapItem{Key: k, Value: convertValue(val)})
	}
	return
}

// Convert known types
func convertValue(in interface{}) (out interface{}) {
	v1 := DeReference(in)
	switch x1 := v1.(type) {
	case StringMap, map[string]interface{}, map[interface{}]interface{}:
		out = yaml.MapSlice(*ToStringMap(x1))
	case []interface{}:
		for j := 0; j < len(x1); j++ {
			v2 := DeReference(x1[j])
			switch x2 := v2.(type) {
			case StringMap, map[string]interface{}, map[interface{}]interface{}:
				x1[j] = yaml.MapSlice(*ToStringMap(x2))
			}
		}
		out = x1
	default:
		out = in
	}
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
	_data, err := yaml.Marshal(yaml.MapSlice(*p))
	if err != nil {
		return
	}
	data = string(_data)
	return
}

// YAMLE converts the Map into a YAML string
func (p *StringMap) YAMLE() (data string, err error) {
	var _data []byte
	if _data, err = yaml.Marshal(yaml.MapSlice(*p)); err != nil {
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
	return yaml_enc.WriteYAML(filename, yaml.MapSlice(*p))
}
