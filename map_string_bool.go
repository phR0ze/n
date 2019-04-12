package n

// StringMapBool implements the Map interface providing a generic way to work with map types
// including convenience methods on par with other rapid development languages.
type StringMapBool map[string]bool

// NewStringMapBool creates a new empty StringMapBool if nothing given else simply
// casts the given map to StringMapBool.
func NewStringMapBool(m ...map[string]bool) *StringMapBool {
	var new StringMapBool
	if len(m) == 0 {
		new = StringMapBool(map[string]bool{})
	} else {
		new = StringMapBool(m[0])
	}
	return &new
}

// Any tests if the map is not empty or optionally if it contains any of the given keys.
func (p *StringMapBool) Any(keys ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	for i := 0; i < len(keys); i++ {
		if key, ok := keys[i].(string); ok {
			if _, ok := (*p)[key]; ok {
				return true
			}
		}
	}
	return false
}

// Set the value with the given key. Returns true if the key was not yet in the map.
func (p *StringMapBool) Set(key, val interface{}) bool {
	if p == nil {
		p = NewStringMapBool()
	}
	k, okk := key.(string)
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
