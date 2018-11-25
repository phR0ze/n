package nub

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

// M materializes object invoking deferred execution
func (m *strMapNub) M() map[string]interface{} {
	return m.raw
}

// Merge the other maps into this map with the first taking the lowest
// precedence and so on until the last takes the highest
func (m *strMapNub) Merge(other ...map[string]interface{}) *strMapNub {
	slice := StrMapSlice(other)
	for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
		m.raw = mergeMap(m.raw, item.raw)
	}
	return m
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
