package nub

import (
	"reflect"
	"strings"
)

// Contains checks if the given target is found
func (q *Queryable) Contains(obj interface{}) bool {
	if q.Iter != nil {
		switch q.ref.Kind() {
		case reflect.Array, reflect.Slice:
			next := q.Iter()
			for x, ok := next(); ok; x, ok = next() {
				if x == obj {
					return true
				}
			}
		case reflect.String:
			if reflect.ValueOf(obj).Kind() == reflect.String {
				return strings.Contains(q.ref.Interface().(string), obj.(string))
			}
		default:
			panic("TODO: implement Contains")
		}
	} else if q.ref.Interface() == obj {
		return true
	}
	return false
}

// ContainsAny checks if any of the given targets are found
func (q *Queryable) ContainsAny(obj interface{}) bool {
	other := Q(obj)

	// Other is singular so defer to Contains
	if other.Singular() {
		return q.Contains(obj)
	}

	// This is singular
	if q.Singular() {
		next := other.Iter()
		for target, ok := next(); ok; target, ok = next() {
			if q.ref.Interface() == target {
				return true
			}
		}
	}

	// Neither is singular
	switch q.ref.Kind() {
	case reflect.Array, reflect.Slice:
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			oNext := other.Iter()
			for y, oOk := oNext(); oOk; y, oOk = oNext() {
				if x == y {
					return true
				}
			}
		}
	default:
		panic("TODO: implement ContainsAny")
	}
	return false
}
