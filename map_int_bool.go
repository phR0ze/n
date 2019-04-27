package n

// IntMapBool implements the Map interface providing a generic way to work with map types
// including convenience methods on par with rapid development languages.
type IntMapBool map[int]bool

// NewIntMapBool creates a new empty IntMapBool if nothing given else simply
// casts the given map to IntMapBool.
func NewIntMapBool(m ...map[int]bool) *IntMapBool {
	var new IntMapBool
	if len(m) == 0 {
		new = IntMapBool(map[int]bool{})
	} else {
		new = IntMapBool(m[0])
	}
	return &new
}

// Any tests if this Map is not empty or optionally if it contains any of the given variadic keys.
func (p *IntMapBool) Any(keys ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	for i := 0; i < len(keys); i++ {
		if key, ok := keys[i].(int); ok {
			if _, ok := (*p)[key]; ok {
				return true
			}
		}
	}
	return false
}

// Len returns the number of elements in this Map.
func (p *IntMapBool) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// Set the value for the given key to the given val. Returns true if the key did not yet exists in this Map.
func (p *IntMapBool) Set(key, val interface{}) bool {
	if p == nil {
		return false
	}
	k, okk := key.(int)
	v, okv := val.(bool)
	if !okk || !okv {
		return false
	}
	if _, ok := (*p)[k]; !ok {
		(*p)[k] = v
		return true
	}
	return false
}
