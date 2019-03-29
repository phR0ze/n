// Package n provides many Go types with convenience functions reminiscent of Ruby or C#.
//
// n was created to reduce the friction I had adopting Go as my primary language of choice by
// reducing coding verbosity required by Go via code reuse. The Numerable types wrapp various
// Go types to provide this functionality.
//
// Conventions used across n types and pkgs
//
// • In order to deal with Golang's decision to not support function overloading n makes use
// of a variety of prefix/suffix capital letters to indicate different function varieties. The
// function that contains no suffix is known as the base function.
//
// • Function names suffixed with 'E' indicates the function is a corollary to the function
// without the 'E' but returns an Error while the base function does not.
//
// • Function names suffixed with 'V' indicates the function is a corollary to the function
// • function names suffixed with 'S' indicates the function is a corollary to the function
// without the 'S' but either accepts a slice as input or returns a Slice.
//
// • Function names suffixed with 'V' indicates the function is a corollary to the function
// without the 'V' but accepts Variadic input.
//
// • Documentation should be thorough and relied upon for guidance as, for a love of brevity,
// many functions use single capital letters frequently to indicate types. 'O' is being used to
// indicate the interface{} type or to export the underlying Go type as an interface{}. 'S' is
// used to refer to slice types, 'M' refers to map types. 'A' refers to string types and
// combinations may be used to indicate complex types. The documentation will always call out
// what exactly they mean, but the function name may be cryptic until understood.
package n

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

// Comparable provides a means of comparing arbitrary types
// These methods should be implemented using a non-pointer receiver
type Comparable interface {
	Less(other interface{}) bool  // Less returns true if the object is less than the other object
	Equal(other interface{}) bool // Equal returns true if the object is value equal to the other object
}

// O is an alias for interface{} and provides a number of export methods
type O interface{}

// Type provides a simple way to track Numerable types
type Type uint8

// Numerable provides chainable execution and an algorithm abstraction layer for various
// underlying types
type Numerable interface {
	O() interface{}              // O returns the underlying data structure
	Any(obj ...interface{}) bool // Any tests if the numerable is not empty or optionally if it contains any of the given Variadic elements.
	Empty() bool                 // Empty tests if the numerable is empty.
	Len() int                    // Len returns the number of elements in the numerable
	Nil() bool                   // Nil tests if the numerable is nil
}

// indirect dereferences the reflect.Value recursively until its a non-pointer type
func indirect(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}
	return indirect(v.Elem())
}

// Check to see if the given type is an optimized slice type
func optimized(slice interface{}) (result bool) {
	result = true
	switch slice.(type) {
	case []bool:
	case []int:
	case []string:
	default:
		result = false
	}
	return
}

//
//
//
//
//
//
//

// OldNumerable provides chainable deferred execution
// and is the heart of the algorithm abstraction layer
type OldNumerable struct {
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

// Nil provides a new nil Numerable
func Nil() *OldNumerable {
	return &OldNumerable{v: nil, Kind: reflect.Invalid}
}

// Q provides origination for the Numerable abstraction layer
func Q(obj interface{}) *OldNumerable {
	v := reflect.ValueOf(obj)
	q := &OldNumerable{v: &v, Kind: v.Kind()}

Switch:
	switch q.Kind {

	// Slice types
	case reflect.Array, reflect.Slice:
		if q.v.IsNil() {
			*q.v = reflect.MakeSlice(q.v.Type(), 0, 10)
		}
		q.Iter = sliceIter(v)

	// Handle map types
	case reflect.Map:
		if q.v.IsNil() {
			*q.v = reflect.MakeMap(q.v.Type())
		}
		q.Iter = mapIter(v)

	// Handle string types
	case reflect.String:
		q.Iter = strIter(v)

	// Chan types
	case reflect.Chan:
		panic("TODO: handle reflect.Chan")

	// Pointer types
	case reflect.Ptr:
		nv := reflect.Indirect(reflect.ValueOf(obj))
		nq := &OldNumerable{v: &nv, Kind: nv.Kind()}
		switch nq.Kind {
		case reflect.Array, reflect.Slice, reflect.Map, reflect.String, reflect.Chan:
			v = nv
			q = nq
			goto Switch
		}
	}

	return q
}

// Any checks if the numerable has anything in it
func (q *OldNumerable) Any() bool {
	if q.v == nil {
		return false
	}
	if q.Iter != nil {
		return q.v.Len() > 0
	}
	return q.v.Interface() != nil
}

// AnyWhere check if any match the given lambda
func (q *OldNumerable) AnyWhere(lambda func(O) bool) bool {
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

// Append modifies the underlying type, converting it to a slice as needed,
// then appending the given items to the underlying collection.
// Returns the numerable for chaining.
func (q *OldNumerable) Append(items ...interface{}) *OldNumerable {
	if q.TypeMap() {
		panic("Append doesn't support map types")
	}
	q.toSlice(items...)
	for i := 0; i < len(items); i++ {
		item := reflect.ValueOf(items[i])
		*q.v = reflect.Append(*q.v, item)
	}
	q.Iter = sliceIter(*q.v)
	return q
}

// At returns the item at the given index location. Allows for negative notation
func (q *OldNumerable) At(i int) *OldNumerable {
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
	return Nil()
}

// Clear the underlying collection in the numerable
func (q *OldNumerable) Clear() *OldNumerable {
	switch q.Kind {
	case reflect.Array, reflect.Slice:
		*q.v = reflect.MakeSlice(q.v.Type(), 0, 10)
		q.Iter = sliceIter(*q.v)
	case reflect.Map:
		*q.v = reflect.MakeMap(q.v.Type())
		q.Iter = mapIter(*q.v)
	case reflect.String:
		*q.v = reflect.ValueOf("")
		q.Iter = strIter(*q.v)
	default:
		panic("unhandled type")
	}
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
func (q *OldNumerable) Contains(obj interface{}) bool {
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
func (q *OldNumerable) ContainsAny(obj interface{}) bool {
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
func (q *OldNumerable) Copy(obj interface{}) *OldNumerable {
	var other *OldNumerable
	if x, ok := obj.(*OldNumerable); ok {
		other = x
	} else {
		other = Q(obj)
	}
	q.Kind = other.Kind
	q.Iter = other.Iter
	q.v = other.v
	return q
}

// Delete all items that match the given item for slices or the key value
// pair for maps or matching rune for strings. Returns true if something was deleted.
func (q *OldNumerable) Delete(obj interface{}) (ok bool) {
	switch q.Kind {
	case reflect.Array, reflect.Slice:
		//*q.v = reflect.MakeSlice(q.v.Type(), 0, 10)
		//q.Iter = sliceIter(*q.v)
	case reflect.Map:
		key := reflect.ValueOf(obj)
		if val := q.v.MapIndex(key); val != (reflect.Value{}) {
			ok = true
			q.v.SetMapIndex(reflect.ValueOf(obj), reflect.Value{})
		}
	case reflect.String:
		//*q.v = reflect.ValueOf("")
		//q.Iter = strIter(*q.v)
	default:
		panic("unhandled type")
	}
	return
}

// DeleteAt deletes the item at the given index location. Allows for negative notation.
// Returns the deleted element Numerable or Nil Numerable if missing.
func (q *OldNumerable) DeleteAt(i int) (item *OldNumerable) {
	if q.TypeIter() && !q.TypeMap() {
		if i < 0 {
			i = q.v.Len() + i
		}
		if i >= 0 && i < q.v.Len() {
			switch x := q.v.Interface().(type) {

			// for strings delete at the rune level
			case string:
				item = Q(string(x[i]))
				if i+1 < len(x) {
					*q.v = reflect.ValueOf(string(append([]rune(x[:i]), []rune(x[i+1:])...)))
				} else {
					*q.v = reflect.ValueOf(x[:i])
				}

			// delete object from iterable
			default:
				item = Q(q.v.Index(i).Interface())
				if i+1 < q.v.Len() {
					*q.v = reflect.AppendSlice(q.v.Slice(0, i), q.v.Slice(i+1, q.v.Len()))
				} else {
					*q.v = q.v.Slice(0, i)
				}
			}

			q.Iter = sliceIter(*q.v)
			return item
		}
	}
	if item == nil {
		item = Nil()
	}
	return
}

// Each iterates over the numerable and executes the given action
func (q *OldNumerable) Each(action func(O)) {
	if q.TypeIter() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			action(x)
		}
	}
}

// EachE iterates over the numerable and executes the given action
// Abort early and return error if non nil
func (q *OldNumerable) EachE(action func(O) error) (err error) {
	if q.TypeIter() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if err = action(x); err != nil {
				return err
			}
		}
	}
	return
}

// Find returns a new numerable containing the first item which matches the given lambda.
// Returns nil if not found.
func (q *OldNumerable) Find(lambda func(O) bool) (result *OldNumerable) {
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		if lambda(x) {
			result = Q(x)
			return
		}
	}
	return
}

// First returns the first item as numerable
// returns a nil numerable when index out of bounds
func (q *OldNumerable) First() (result *OldNumerable) {
	if q.Len() > 0 {
		return q.At(0)
	}
	return Nil()
}

// Flatten returns a new slice that is one-dimensional flattening.
// That is, for every item that is a slice, extract its items into the new slice.
func (q *OldNumerable) Flatten() (result *OldNumerable) {
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
func (q *OldNumerable) Insert(i int, items ...interface{}) *OldNumerable {
	q.toSlice(items...)
	if i < 0 {
		i = q.v.Len() + i
	}
	if i >= 0 && i < q.v.Len() && q.v.Len() > 0 && len(items) > 0 {

		// Create a new slice
		typ := q.v.Index(0).Type()
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
func (q *OldNumerable) Join(delim string) *OldNumerable {
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

// Last returns the last item as numerable
// returns a nil numerable when index out of bounds
func (q *OldNumerable) Last() (result *OldNumerable) {
	if q.Len() > 0 {
		return q.At(-1)
	}
	return Nil()
}

// Len of the collection type including string
func (q *OldNumerable) Len() int {
	if q.TypeIter() {
		return q.v.Len()
	} else if q.Nil() {
		return 0
	}
	return 1
}

// Map manipulates the numerable data into a new form
func (q *OldNumerable) Map(sel func(O) O) (result *OldNumerable) {
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		obj := sel(x)

		// Drill into numerables
		if s, ok := obj.(*OldNumerable); ok {
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

// MapF manipulates the numerable data into a new form then flattens
func (q *OldNumerable) MapF(sel func(O) O) (result *OldNumerable) {
	result = q.Map(sel).Flatten()
	return
}

// MapMany manipulates the numerable data from two sources in a cross join
func (q *OldNumerable) MapMany(sel func(O) O) (result *OldNumerable) {
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

// Nil tests if the numerable is a nil numerable
func (q *OldNumerable) Nil() bool {
	if q.v == nil || q.Kind == reflect.Invalid {
		return true
	}
	return false
}

// Select returns a new numerable containing all items which match the given lambda
func (q *OldNumerable) Select(lambda func(O) bool) (result *OldNumerable) {
	result = q.newSlice()
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		if lambda(x) {
			result.Append(x)
		}
	}
	return result
}

// Set the item at the given index to the given item
func (q *OldNumerable) Set(i int, item interface{}) *OldNumerable {
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

// // Split the string into a slice on delimiter
// func (q *Numerable) Split(delim string) *NSlice {
// 	if q.TypeStr() {
// 		return A(q.v.Interface().(string)).Split(delim)
// 	}
// 	return S()
// }

// TypeIter checks if the numerable is iterable
func (q *OldNumerable) TypeIter() bool {
	if q.Iter != nil {
		return true
	}
	return false
}

// TypeMap checks if the numerable is reflect.Map
func (q *OldNumerable) TypeMap() bool {
	return q.Kind == reflect.Map
}

// TypeSlice checks if the numerable is reflect.Array or reflect.Slice
func (q *OldNumerable) TypeSlice() bool {
	return q.Kind == reflect.Array || q.Kind == reflect.Slice
}

// TypeStr checks if the numerable is encapsulating a string
func (q *OldNumerable) TypeStr() bool {
	return q.Kind == reflect.String
}

// TypeSingle checks if the numerable is ecapuslating a string or is not iterable
func (q *OldNumerable) TypeSingle() bool {
	if !q.TypeIter() || q.TypeStr() || q.Nil() {
		return true
	}
	return false
}

// Convert the single type into a slice type
func (q *OldNumerable) toSlice(items ...interface{}) {
	if q.TypeSingle() {
		nq := q.newSlice(items...)
		if !q.Nil() {
			*nq.v = reflect.Append(*nq.v, *q.v)
		}
		*q = *nq
	}
}

// Create a new slice of the inner type
func (q *OldNumerable) newSlice(items ...interface{}) *OldNumerable {
	var typ reflect.Type
	switch {
	case len(items) > 0:
		typ = reflect.SliceOf(reflect.TypeOf(items[0]))
	case q.Nil():
		typ = reflect.TypeOf([]interface{}{})
	case q.TypeSingle():
		typ = reflect.SliceOf(q.v.Type())
	case q.TypeMap():
		typ = reflect.SliceOf(reflect.TypeOf(KeyVal{}))
	default:
		if q.Any() {
			typ = reflect.SliceOf(q.v.Index(0).Type())
		} else {
			typ = q.v.Type()
		}
	}
	return Q(reflect.MakeSlice(typ, 0, 10).Interface())
}
