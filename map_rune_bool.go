package n

// RuneMapBool implements the Map interface providing a generic way to work with map types
// including convenience methods on par with rapid development languages.
type RuneMapBool map[rune]bool

// NewRuneMapBool creates a new empty RuneMapBool if nothing given else simply
// casts the given map to RuneMapBool.
func NewRuneMapBool(m ...map[rune]bool) *RuneMapBool {
	var new RuneMapBool
	if len(m) == 0 {
		new = RuneMapBool(map[rune]bool{})
	} else {
		new = RuneMapBool(m[0])
	}
	return &new
}

// Any tests if this Map is not empty or optionally if it contains any of the given variadic keys.
func (p *RuneMapBool) Any(keys ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	for i := 0; i < len(keys); i++ {
		if key, ok := keys[i].(rune); ok {
			if _, ok := (*p)[key]; ok {
				return true
			}
		}
	}
	return false
}

// Len returns the number of elements in this Map.
func (p *RuneMapBool) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// Set the value for the given key to the given val. Returns true if the key did not yet exists in this Map.
func (p *RuneMapBool) Set(key, val interface{}) bool {
	if p == nil {
		return false
	}
	k, okk := key.(rune)
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
