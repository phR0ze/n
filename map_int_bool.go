package n

// IntMapBool implements the Map interface providing a generic way to work with map types
// including convenience methods on par with other rapid development languages.
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

// Any tests if the map is not empty or optionally if it contains any of the given keys.
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

// Set the value with the given key. Returns true if the key was not yet in the map.
func (p *IntMapBool) Set(key, val interface{}) bool {
	if p == nil {
		p = NewIntMapBool()
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
