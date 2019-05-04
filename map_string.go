package n

// StringMap implements the Map interface providing a generic way to work with map types
// including convenience methods on par with rapid development languages.
type StringMap map[string]interface{}

// NewStringMap creates a new empty StringMap if nothing given else simply
// casts the given map to StringMap.
func NewStringMap(m ...map[string]interface{}) *StringMap {
	var new StringMap
	if len(m) == 0 {
		new = StringMap(map[string]interface{}{})
	} else {
		new = StringMap(m[0])
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

// Delete modifies this Map to delete the indicated key-value pair and returns the value from the Map.
func (p *StringMap) Delete(key interface{}) (obj *Object) {
	obj = &Object{}
	if p == nil {
		return
	}
	k := ToString(key)
	if val, ok := (*p)[k]; ok {
		obj.o = val
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

// Generic returns true if the underlying implementation uses reflection
func (p *StringMap) Generic() bool {
	return false
}

// Len returns the number of elements in this Map.
func (p *StringMap) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
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
		p = NewStringMap()
	}
	p.Set(key, val)
	return p
}
