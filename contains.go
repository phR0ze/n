package nub

import (
	"reflect"
	"strings"
)

// Contains checks if all of the given obj are found.
// When obj is a string and this is a string check will fall back on strings.Contains.
// When obj is a string and this is a string slice, slice will be checked for obj.
// When obj is a non-interable and this is non-iterable a direct check is made.
// When obj is a non-interable and this is slice, slice will be checked for obj.
// When obj is a slice of string and this is a string each string check using strings.Contains.
// When obj is a slice and this is a slice each item will be checked in the slice.
// When obj is a slice and this is a map each item will be checked in the map as a key.
func (q *Queryable) Contains(obj interface{}) bool {
	other := Q(obj)
	if !q.Any() || !other.Any() {
		return false
	}

	// Non iterable type
	if q.TypeSingle() {

		// Both strings - pass through to stings.Contains
		if q.TypeStr() && other.TypeStr() {
			return strings.Contains(q.ref.Interface().(string), obj.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Set([]interface{}{obj})
		}
		next := other.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if str, ok := x.(string); ok {
				if !strings.Contains(q.ref.Interface().(string), str) {
					return false
				}
			} else {
				if q.ref.Interface() != x {
					return false
				}
			}
		}
	} else {
		switch q.ref.Kind() {
		case reflect.Array, reflect.Slice:
			if !other.TypeSingle() {
				next := other.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if !q.Contains(x) {
						return false
					}
				}
			} else {
				next := q.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if x == obj {
						return true
					}
				}
				return false
			}
		case reflect.Map:
			keys := Q(q.ref.MapKeys()).Map(func(x interface{}) interface{} {
				return x.(reflect.Value).Interface()
			})
			if !other.TypeSingle() {
				next := other.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if !keys.Contains(x) {
						return false
					}
				}
			} else {
				if !keys.Contains(obj) {
					return false
				}
			}
		default:
			panic("TODO: implement Contains")
		}
	}
	return true
}

// ContainsAny checks if any of the given obj is found.
// ContainsAny behaves much like Contains only it allows for matching any not all.
func (q *Queryable) ContainsAny(obj interface{}) bool {
	other := Q(obj)

	// Non iterable type
	if q.TypeSingle() {

		// Both strings - pass through to stings.Contains
		if q.TypeStr() && other.TypeStr() {
			return strings.Contains(q.ref.Interface().(string), obj.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Set([]interface{}{obj})
		}
		next := other.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if q.ref.Interface() == x {
				return true
			}
		}
	} else {
		switch q.ref.Kind() {
		case reflect.Array, reflect.Slice:
			if !other.TypeSingle() {
				next := q.Iter()
				for x, ok := next(); ok; x, ok = next() {
					nexty := other.Iter()
					for y, oky := nexty(); oky; y, oky = nexty() {
						if x == y {
							return true
						}
					}
				}
			} else {
				next := q.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if x == obj {
						return true
					}
				}
			}
		case reflect.Map:
			if !other.TypeSingle() {
				for _, key := range q.ref.MapKeys() {
					next := other.Iter()
					for x, ok := next(); ok; x, ok = next() {
						if key.Interface() == x {
							return true
						}
					}
				}
			} else {
				for _, key := range q.ref.MapKeys() {
					if key.Interface() == obj {
						return true
					}
				}
			}
		default:
			panic("TODO: implement Contains")
		}
	}
	return false
}
