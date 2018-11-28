package nub

import (
	"errors"
	"reflect"
)

// Queryable provides chainable deferred execution
// and is the heart of the algorithm abstraction layer
type Queryable struct {
	O    interface{}
	Iter func() Iterator

	ref *reflect.Value
}

// Iterator provides a closure to capture the index and reset it
type Iterator func() (item interface{}, ok bool)

// S provides a new empty Queryable slice
func S() *Queryable {
	raw := []interface{}{}
	ref := reflect.ValueOf(raw)
	return &Queryable{
		ref: &ref,
		O:   raw,
		Iter: func() Iterator {
			i := 0
			return func() (item interface{}, ok bool) {
				if ok = i < ref.Len(); ok {
					item = ref.Index(i).Interface()
					i++
				}
				return
			}

		}}
}

// Q provides origination for the Queryable abstraction layer
func Q(obj interface{}) (result *Queryable) {
	if obj != nil {
		ref := reflect.ValueOf(obj)
		result = &Queryable{O: obj, ref: &ref}
		kind := ref.Kind()
		switch kind {

		// Handle slice types
		case reflect.Array, reflect.Slice:
			result.Iter = func() Iterator {
				i := 0
				return func() (item interface{}, ok bool) {
					if ok = i < ref.Len(); ok {
						item = ref.Index(i).Interface()
						i++
					}
					return
				}
			}

		// Handle chan types
		case reflect.Chan:
			panic("TODO: handle reflect.Chan")

		// Handle map types
		case reflect.Map:
			panic("TODO: handle reflect.Map")

		// Handle int types
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// No iterator should exist

		// Handle string types
		case reflect.String:
			result.Iter = func() Iterator {
				i := 0
				return func() (item interface{}, ok bool) {
					if ok = i < ref.Len(); ok {
						item = ref.Index(i).Interface()
						i++
					}
					return
				}
			}

		// Handle unknown types
		default:
			panic("TODO: handle custom types")
		}
	} else {
		result = S()
	}
	return
}

// Append items to the end of the collection and return the queryable
// converting to a collection if necessary
func (q *Queryable) Append(obj interface{}) *Queryable {

	// No existing type return a new queryable
	if q.ref == nil {
		*q = *Q(obj)
		return q
	}

	// Not a collection type create a new queryable
	kind := q.ref.Kind()
	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.Map {
		*q = *S().Append(q.O)
	}

	// Append to the collection type
	ref := reflect.ValueOf(obj)
	switch ref.Kind() {
	case reflect.Map:
		panic("TODO: handle appending to map")
	case reflect.Array, reflect.Slice:
		for i := 0; i < ref.Len(); i++ {
			*q.ref = reflect.Append(*q.ref, ref.Index(i))
		}
	default:
		*q.ref = reflect.Append(*q.ref, ref)
	}
	return q
}

// At returns the item at the given index location. Allows for negative notation
func (q *Queryable) At(i int) *Queryable {
	if q.Iter != nil {
		if i < 0 {
			i = q.ref.Len() + i
		}
		if i >= 0 && i < q.ref.Len() {
			if str, ok := q.O.(string); ok {
				return Q(string(str[i]))
			}
			return Q(q.ref.Index(i).Interface())
		}
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the queryable collection
func (q *Queryable) Clear() *Queryable {
	*q = *S()
	return q
}

// Each iterates over the queryable and executes the given action
func (q *Queryable) Each(action func(interface{})) {
	if q.Iter != nil {
		next := q.Iter()
		for item, ok := next(); ok; item, ok = next() {
			action(item)
		}
	}
}

// Len of the collection type including string
func (q *Queryable) Len() int {
	if q.Iter != nil {
		return q.ref.Len()
	}
	return 1
}

// Set provides a way to set underlying object Queryable is operating on
func (q *Queryable) Set(obj interface{}) *Queryable {
	other := Q(obj)
	q.O = other.O
	q.Iter = other.Iter
	q.ref = other.ref
	return q
}

// Singular is queryable encapsulating a non-collection
func (q *Queryable) Singular() bool {
	_, strType := q.O.(string)
	if q.Iter == nil || strType {
		return true
	}
	return false
}
