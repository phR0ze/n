package n

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

// O is an alias for interface{} to reduce verbosity
// i'm using O for Object as I is already taken for Int types
type O interface{}

// Queryable provides chainable deferred execution
// and is the heart of the algorithm abstraction layer
type Queryable struct {
	v    *reflect.Value  // underlying value
	Kind reflect.Kind    // kind of hte underlying value
	Iter func() Iterator // iterator for the underlying value
}

// Iterator provides a closure to capture the index and reset it
type Iterator func() (item O, ok bool)

// KeyVal similar to C# for iterator over maps
type KeyVal struct {
	Key interface{}
	Val interface{}
}

func strIter(ref reflect.Value) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		return func() (item O, ok bool) {
			if ok = i < len; ok {
				item = ref.Index(i).Interface()
				i++
			}
			return
		}
	}
}

func mapIter(ref reflect.Value) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		keys := ref.MapKeys()
		return func() (item O, ok bool) {
			if ok = i < len; ok {
				item = KeyVal{
					Key: keys[i].Interface(),
					Val: ref.MapIndex(keys[i]).Interface(),
				}
				i++
			}
			return
		}
	}
}

func sliceIter(ref reflect.Value) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		return func() (item O, ok bool) {
			if ok = i < len; ok {
				item = ref.Index(i).Interface()
				i++
			}
			return
		}
	}
}

// N provides a new nil Queryable
func N() *Queryable {
	return &Queryable{v: nil, Kind: reflect.Invalid}
}

// Q provides origination for the Queryable abstraction layer
func Q(obj interface{}) *Queryable {
	if obj == nil {
		return N()
	}

	v := reflect.ValueOf(obj)
	q := &Queryable{v: &v, Kind: v.Kind()}
	switch q.Kind {

	// Slice types
	case reflect.Array, reflect.Slice:
		q.Iter = sliceIter(v)

	// Handle map types
	case reflect.Map:
		q.Iter = mapIter(v)

	// Handle string types
	case reflect.String:
		q.Iter = strIter(v)

	// Chan types
	case reflect.Chan:
		panic("TODO: handle reflect.Chan")
	}

	return q
}

// Any checks if the queryable has anything in it
func (q *Queryable) Any() bool {
	if q.v == nil {
		return false
	}
	if q.Iter != nil {
		return q.v.Len() > 0
	}
	return q.v.Interface() != nil
}

// AnyWhere check if any match the given lambda
func (q *Queryable) AnyWhere(lambda func(O) bool) bool {
	if !q.TypeSingle() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if lambda(x) {
				return true
			}
		}
	} else if q.Nil() {
		return false
	} else if lambda(q.v.Interface()) {
		return true
	}
	return false
}

// Append items to the end of the collection and return the queryable
// converting to a collection if necessary
func (q *Queryable) Append(items ...interface{}) *Queryable {
	q.toSlice(items...)
	for i := 0; i < len(items); i++ {
		item := reflect.ValueOf(items[i])
		*q.v = reflect.Append(*q.v, item)
	}
	q.Iter = sliceIter(*q.v)
	return q
}

// At returns the item at the given index location. Allows for negative notation
func (q *Queryable) At(i int) *Queryable {
	if q.TypeIter() {
		if i < 0 {
			i = q.v.Len() + i
		}
		if i >= 0 && i < q.v.Len() {
			if str, ok := q.v.Interface().(string); ok {
				return Q(string(str[i]))
			}
			return Q(q.v.Index(i).Interface())
		}
	}
	return N()
}

// Clear the queryable collection
func (q *Queryable) Clear() *Queryable {
	*q = *N()
	return q
}

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
			return strings.Contains(q.v.Interface().(string), obj.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Copy([]interface{}{obj})
		}
		next := other.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if str, ok := x.(string); ok {
				if !strings.Contains(q.v.Interface().(string), str) {
					return false
				}
			} else {
				if q.v.Interface() != x {
					return false
				}
			}
		}
	} else {
		switch q.Kind {
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
			keys := Q(q.v.MapKeys()).Map(func(x O) O {
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
	if q.Nil() {
		return false
	}

	// Non iterable type
	if q.TypeSingle() {

		// Both strings - pass through to stings.Contains
		if q.TypeStr() && other.TypeStr() {
			return strings.Contains(q.v.Interface().(string), obj.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Copy([]interface{}{obj})
		}
		next := other.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if q.v.Interface() == x {
				return true
			}
		}
	} else {
		switch q.Kind {
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
				for _, key := range q.v.MapKeys() {
					next := other.Iter()
					for x, ok := next(); ok; x, ok = next() {
						if key.Interface() == x {
							return true
						}
					}
				}
			} else {
				for _, key := range q.v.MapKeys() {
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

// Copy given obj into this one and reset types
func (q *Queryable) Copy(obj interface{}) *Queryable {
	other := Q(obj)
	q.Kind = other.Kind
	q.Iter = other.Iter
	q.v = other.v
	return q
}

// Each iterates over the queryable and executes the given action
func (q *Queryable) Each(action func(O)) {
	if q.TypeIter() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			action(x)
		}
	}
}

// First returns the first item as queryable
// returns a nil queryable when index out of bounds
func (q *Queryable) First() (result *Queryable) {
	if q.Len() > 0 {
		return q.At(0)
	}
	return N()
}

// Flatten returns a new slice that is one-dimensional flattening.
// That is, for every item that is a slice, extract its items into the new slice.
func (q *Queryable) Flatten() (result *Queryable) {
	if q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {

			// Create new slice of the inner type
			if result == nil {
				result = Q(reflect.MakeSlice(reflect.TypeOf(x), 0, 10).Interface())
			}

			// Add sub slice's elements to new slice
			Q(x).Each(func(y O) {
				result.Append(y)
			})
		}
	} else {
		panic("TODO: implement Flatten() for maps")
	}
	if result == nil {
		result = q
	}
	return
}

// Insert the item at the given index, negative notation supported
func (q *Queryable) Insert(i int, items ...interface{}) *Queryable {
	q.toSlice(items...)
	if i < 0 {
		i = q.v.Len() + i
	}
	if i >= 0 && i < q.v.Len() && len(items) > 0 {

		// Create a new slice
		typ := reflect.TypeOf(items[0])
		slice := reflect.MakeSlice(reflect.SliceOf(typ), 0, 10)

		// Append those before i
		for _, j := range Range(0, i-1) {
			slice = reflect.Append(slice, q.v.Index(j))
		}

		// Append new items
		for j := 0; j < len(items); j++ {
			slice = reflect.Append(slice, reflect.ValueOf(items[j]))
		}

		// Append those after
		for _, j := range Range(i, q.Len()-1) {
			slice = reflect.Append(slice, q.v.Index(j))
		}

		*q = *Q(slice.Interface())
		q.Iter = sliceIter(*q.v)
	} else {
		q.Append(items...)
	}
	return q
}

// Join slice items as string with given delimeter
func (q *Queryable) Join(delim string) *Queryable {
	var joined bytes.Buffer
	if q.TypeStr() {
		joined.WriteString(q.v.Interface().(string))
	} else if q.TypeIter() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			switch x := x.(type) {
			case string:
				joined.WriteString(x)
				joined.WriteString(delim)
			case int:
				joined.WriteString(strconv.Itoa(x))
				joined.WriteString(delim)
			}
		}
	}
	return Q(strings.TrimSuffix(joined.String(), delim))
}

// Last returns the last item as queryable
// returns a nil queryable when index out of bounds
func (q *Queryable) Last() (result *Queryable) {
	if q.Len() > 0 {
		return q.At(-1)
	}
	return N()
}

// Len of the collection type including string
func (q *Queryable) Len() int {
	if q.TypeIter() {
		return q.v.Len()
	} else if q.Nil() {
		return 0
	}
	return 1
}

// Map manipulates the queryable data into a new form
func (q *Queryable) Map(sel func(O) O) (result *Queryable) {
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		obj := sel(x)

		// Drill into queryables
		if s, ok := obj.(*Queryable); ok {
			obj = s.v.Interface()
		}

		// Create new slice of the return type of sel
		if result == nil {
			typ := reflect.TypeOf(obj)
			result = Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
		}
		result.Append(obj)
	}
	if result == nil {
		result = Q([]interface{}{})
	}
	return
}

// MapF manipulates the queryable data into a new form then flattens
func (q *Queryable) MapF(sel func(O) O) (result *Queryable) {
	result = q.Map(sel).Flatten()
	return
}

// MapMany manipulates the queryable data from two sources in a cross join
func (q *Queryable) MapMany(sel func(O) O) (result *Queryable) {
	// next := q.Iter()
	// for x, ok := next(); ok; x, ok = next() {
	// 	s := sel(x)

	// 	// Create new slice of the return type of sel
	// 	if result == nil {
	// 		typ := reflect.TypeOf(s)
	// 		result = Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
	// 	}
	// 	result.Append(s)
	// }
	// return result
	return
}

// Nil tests if the queryable is a nil queryable
func (q *Queryable) Nil() bool {
	if q.v == nil {
		return true
	}
	return false
}

// Set the item at the given index to the given item
func (q *Queryable) Set(i int, item interface{}) *Queryable {
	if q.TypeIter() && !q.TypeStr() {
		if i < 0 {
			i = q.v.Len() + i
		}
		if i >= 0 && i < q.v.Len() {
			v := reflect.ValueOf(item)
			q.v.Index(i).Set(v)
		}
	}
	return q
}

// Split the string into a slice on delimiter
func (q *Queryable) Split(delim string) *strSliceN {
	if q.TypeStr() {
		return A(q.v.Interface().(string)).Split(delim)
	}
	return S()
}

// TypeIter checks if the queryable is iterable
func (q *Queryable) TypeIter() bool {
	if q.Iter != nil {
		return true
	}
	return false
}

// TypeMap checks if the queryable is reflect.Map
func (q *Queryable) TypeMap() bool {
	return q.Kind == reflect.Map
}

// TypeSlice checks if the queryable is reflect.Array or reflect.Slice
func (q *Queryable) TypeSlice() bool {
	return q.Kind == reflect.Array || q.Kind == reflect.Slice
}

// TypeStr checks if the queryable is encapsulating a string
func (q *Queryable) TypeStr() bool {
	return q.Kind == reflect.String
}

// TypeSingle checks if the queryable is ecapuslating a string or is not iterable
func (q *Queryable) TypeSingle() bool {
	if !q.TypeIter() || q.TypeStr() || q.Nil() {
		return true
	}
	return false
}

// Convert the single type into a slice type
func (q *Queryable) toSlice(items ...interface{}) {
	if q.TypeSingle() && len(items) > 0 {
		typ := reflect.TypeOf(items[0])
		qSlice := Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
		if !q.Nil() {
			qSlice.Append(q.v.Interface())
		}
		*q = *qSlice
	}
}
