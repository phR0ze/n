package nub

import (
	"reflect"
	"strings"
)

// Contains checks if all of the given targets are found
func (q *Queryable) Contains(targets interface{}) bool {
	other := Q(targets)
	if !q.Any() || !other.Any() {
		return false
	}

	// Non iterable type
	if q.TypeSingle() {

		// Both strings - pass through to stings.Contains
		if q.TypeStr() && other.TypeStr() {
			return strings.Contains(q.ref.Interface().(string), targets.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Set([]interface{}{targets})
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
					if x == targets {
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
				if !keys.Contains(targets) {
					return false
				}
			}
		default:
			panic("TODO: implement Contains")
		}
	}
	return true
}

// ContainsAny checks if any of the given targets is found
func (q *Queryable) ContainsAny(targets interface{}) bool {
	other := Q(targets)

	// Non iterable type
	if q.TypeSingle() {

		// Both strings - pass through to stings.Contains
		if q.TypeStr() && other.TypeStr() {
			return strings.Contains(q.ref.Interface().(string), targets.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Set([]interface{}{targets})
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
					if x == targets {
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
					if key.Interface() == targets {
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
