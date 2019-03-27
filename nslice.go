package n

import (
	"fmt"
	"reflect"
)

// NSlice provides a generic way to work with slice types providing convenience methods
// on par with other rapid development languages.
//
// Implements the Numerable interface.
type NSlice struct {
	o   interface{} // underlying slice object
	len int         // slice length
}

// Slice creates a new NSlice by simply storing slice 'obj' directly to avoid using reflection
// processing at a 10x overhead savings. Non slice 'obj' are encapsulated in a new slice of
// that type using reflection, thus incurring the standard 10x overhead.
//
// Return value n *NSlice will never be nil but n.Nil() may be true as nil or empty []interface{}
// values are ignored to avoid internally using a []interface{}. The internal type will be
// set later with the given type when an n.AppendX method is called.
//
// Cost: ~0x - 10x
func Slice(obj interface{}) (n *NSlice) {
	n = &NSlice{}
	v := reflect.ValueOf(obj)

	k := v.Kind()
	x, interfaceSliceType := obj.([]interface{})
	switch {

	// Return the NSlice.Nil
	case k == reflect.Invalid:

	// Iterate over array and append
	case k == reflect.Array:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			n.Append(item)
		}

	// Convert []interface to slice of elem type
	case interfaceSliceType:
		n = SliceV(x...)

	// Slice of distinct type can be used directly
	case k == reflect.Slice:
		n.o = obj
		n.len = v.Len()

	// Append single items
	default:
		n.Append(obj)
	}
	return
}

// SliceV creates a new NSlice encapsulating the given variadic elements in a new slice of
// that type. Thi call incurs the full 10x reflection overhead. For large slice params use
// the Slice() func instead.
//
// Cost: ~10x
func SliceV(items ...interface{}) (n *NSlice) {
	n = &NSlice{}

	// Return NSlice.Nil if nothing given
	if len(items) == 0 {
		return
	}

	// Create new slice from the type of the first non Invalid element
	var slice *reflect.Value
	for i := 0; i < len(items); i++ {

		// Create target slice from first Valid element
		if slice == nil && reflect.ValueOf(items[i]).IsValid() {
			typ := reflect.SliceOf(reflect.TypeOf(items[i]))
			v := reflect.MakeSlice(typ, 0, 10)
			slice = &v
		}

		// Append item to slice
		if slice != nil {
			item := reflect.ValueOf(items[i])
			*slice = reflect.Append(*slice, item)
		}
	}
	if slice != nil {
		n.o = slice.Interface()
		n.len = slice.Len()
	}
	return
}

// create a new empty slice of the given type or element type if a slice/array
// want to return a new *NSlice so that we can use this in the AppendX functions
// to defer creating an underlying slice type until we have an actual type to work with.
func newEmptySlice(items interface{}) (n *NSlice) {
	n = &NSlice{}
	v := reflect.ValueOf(items)
	typ := reflect.TypeOf([]interface{}{})

	k := v.Kind()
	switch k {

	// Use a new generic slice for nils
	case reflect.Invalid:

	// Use the element type of slice/arrays
	case reflect.Slice, reflect.Array:

		// Use slice type if not generic
		if _, ok := items.([]interface{}); !ok {
			typ = reflect.SliceOf(v.Type().Elem())
		} else {
			// For generics try to find actual element type
			if v.Len() != 0 {
				elem := v.Index(0).Interface()
				if elem != nil {
					typ = reflect.SliceOf(reflect.TypeOf(elem))
				}
			}
		}
	default:
		typ = reflect.SliceOf(v.Type())
	}

	// Create new slice with type of the element
	slice := reflect.MakeSlice(typ, 0, 10)
	n.o = slice.Interface()
	return
}

// Any tests if the numerable is not empty or optionally if it contains
// any of the given Variadic elements.
//
// Cost: ~0x - 10x
//
// Optimized types: bool, int, string
func (s *NSlice) Any(obj ...interface{}) bool {

	// No elements
	if s.Nil() || s.len == 0 {
		return false
	}

	// Elements and not looking for anything
	if len(obj) == 0 {
		return true
	}

	// Looking for something specific
	ok := false
	var typ reflect.Type
	switch slice := s.o.(type) {
	case []bool:
		var x bool
		for i := range obj {
			if x, ok = obj[i].(bool); !ok {
				typ = reflect.TypeOf(obj[i])
				break
			} else {
				for j := range slice {
					if slice[j] == x {
						return true
					}
				}
			}
		}
	case []int:
		var x int
		for i := range obj {
			if x, ok = obj[i].(int); !ok {
				typ = reflect.TypeOf(obj[i])
				break
			} else {
				for j := range slice {
					if slice[j] == x {
						return true
					}
				}
			}
		}
	case []string:
		var x string
		for i := range obj {
			if x, ok = obj[i].(string); !ok {
				typ = reflect.TypeOf(obj[i])
				break
			} else {
				for j := range slice {
					if slice[j] == x {
						return true
					}
				}
			}
		}
	default:
		v := reflect.ValueOf(s.o)
		for i := 0; i < len(obj); i++ {
			x := reflect.ValueOf(obj[i])
			typ = x.Type()
			if v.Type().Elem() != typ {
				break
			} else {
				ok = true
				for j := 0; j < v.Len(); j++ {
					if v.Index(j).Interface() == x.Interface() {
						return true
					}
				}
			}
		}
	}
	if !ok {
		panic(fmt.Sprintf("can't compare type '%v' with '%T' elements", typ, s.o))
	}
	return false
}

// AnyS checks if any of the given 'obj' elements are contained in this Numerable
//
// 'obj' must be a slice type
//
// Cost: ~0x - 10x
//
// Optimized types: bool, int, string
func (s *NSlice) AnyS(obj interface{}) (result bool) {
	if s.Nil() {
		return
	}
	ok := false
	switch slice := s.o.(type) {
	case []bool:
		var x []bool
		if x, ok = obj.([]bool); ok {
			for i := range x {
				for j := range slice {
					if slice[j] == x[i] {
						return true
					}
				}
			}
		}
	case []int:
		var x []int
		if x, ok = obj.([]int); ok {
			for i := range x {
				for j := range slice {
					if slice[j] == x[i] {
						return true
					}
				}
			}
		}
	case []string:
		var x []string
		if x, ok = obj.([]string); ok {
			for i := range x {
				for j := range slice {
					if slice[j] == x[i] {
						return true
					}
				}
			}
		}
	default:
		v := reflect.ValueOf(s.o)
		x := reflect.ValueOf(obj)
		if v.Type() == x.Type() {
			ok = true

			for i := 0; i < x.Len(); i++ {
				for j := 0; j < v.Len(); j++ {
					if v.Index(j).Interface() == x.Index(i).Interface() {
						return true
					}
				}
			}
		}
	}
	if !ok {
		panic(fmt.Sprintf("can't compare type '%T' with '%T' elements", obj, s.o))
	}
	return
}

// Append an item to the end of the NSlice and returns the NSlice for chaining. Avoids the 10x
// reflection overhead cost by type asserting common types. Types not optimized in this way incur
// the full 10x reflection overhead cost.
//
// Cost: ~4x - 10x
//
// Optimized types: bool, int, string
func (s *NSlice) Append(item interface{}) *NSlice {
	if s.Nil() {
		if s == nil {
			s = newEmptySlice(item)
		} else {
			*s = *(newEmptySlice(item))
		}
	}
	ok := false
	switch slice := s.o.(type) {
	case []bool:
		var x bool
		if x, ok = item.(bool); ok {
			s.o = append(slice, x)
		}
	case []int:
		var x int
		if x, ok = item.(int); ok {
			s.o = append(slice, x)
		}
	case []string:
		var x string
		if x, ok = item.(string); ok {
			s.o = append(slice, x)
		}
	default:
		ok = true
		v := reflect.ValueOf(s.o)
		item := reflect.ValueOf(item)
		s.o = reflect.Append(v, item).Interface()
	}
	if !ok {
		panic(fmt.Sprintf("can't insert type '%T' into '%T'", item, s.o))
	}
	s.len++
	return s
}

// AppendV appends the variadic items to the end of the NSlice and returns the NSlice for chaining.
// Avoids the 10x reflection overhead cost by type asserting common types. Types not optimized in
// this way incur the full 10x reflection overhead cost.
//
// Cost: ~6x - 10x
//
// Optimized types: bool, int, string
func (s *NSlice) AppendV(items ...interface{}) *NSlice {
	for _, item := range items {
		s.Append(item)
	}
	return s
}

// AppendS appends the given slice using variadic expansion and returns NSlice for chaining. Avoids
// the 10x reflection overhead cost by type asserting common types. Types not optimized in this
// way incur the full 10x reflection overhead cost. However when appending larger slices fewer times
// the cost reduces down to 2x.
//
// Cost: ~0x - 2x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) AppendS(items interface{}) *NSlice {
	if s.Nil() {
		if s == nil {
			s = newEmptySlice(items)
		} else {
			*s = *(newEmptySlice(items))
		}
	}
	ok := false
	switch slice := s.o.(type) {
	case []bool:
		var x []bool
		if x, ok = items.([]bool); ok {
			s.o = append(slice, x...)
			s.len += len(x)
		}
	case []int:
		var x []int
		if x, ok = items.([]int); ok {
			s.o = append(slice, x...)
			s.len += len(x)
		}
	case []string:
		var x []string
		if x, ok = items.([]string); ok {
			s.o = append(slice, x...)
			s.len += len(x)
		}
	default:
		ok = true
		v := reflect.ValueOf(s.o)
		x := reflect.ValueOf(items)
		s.o = reflect.AppendSlice(v, x).Interface()
		s.len += x.Len()
	}
	if !ok {
		panic(fmt.Sprintf("can't concat type '%T' with '%T'", items, s.o))
	}
	return s
}

// get the absolute value for the pos/neg index.
// return of -1 indicates out of bounds
func (s *NSlice) absIndex(i int) (abs int) {
	if i < 0 {
		abs = s.len + i
	} else {
		abs = i
	}
	if abs < 0 || abs >= s.len {
		abs = -1
	}
	return
}

// At returns the item at the given index location. Allows for negative notation.
// Cost for reflection in this case doesn't seem to add much.
//
// Cost: ~0x - 2x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) At(i int) (obj *NObj) {
	obj = &NObj{}
	if s.Nil() {
		return
	}
	if i = s.absIndex(i); i == -1 {
		return
	}

	switch slice := s.o.(type) {
	case []bool:
		obj.o = slice[i]
	case []int:
		obj.o = slice[i]
	case []string:
		obj.o = slice[i]
	default:
		obj.o = reflect.ValueOf(s.o).Index(i).Interface()
	}
	return
}

// Clear the underlying slice.
//
// Cost: constant
func (s *NSlice) Clear() *NSlice {
	if !s.Nil() {
		*s = *(newEmptySlice(s.o))
	}
	return s
}

// Copy performs a deep copy such that modifications to the copy will not affect
// the original. Expects nothing, in which case everything is copied, or two
// indices i and j, in which case positive and negative notation is supported and
// uses an inclusive behavior such that Slice(0, -1) includes index -1 as opposed
// to Go's exclusive  behavior. Out of bounds indices will be moved within bounds.
//
// An empty NSlice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. SliceV(1,2,3).Copy() == [1,2,3] && SliceV(1,2,3).Copy(1,2) == [2,3]
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) Copy(indices ...int) (result *NSlice) {
	if s == nil {
		result = &NSlice{}
		return
	}
	result = newEmptySlice(s.o)
	if s.len == 0 || len(indices) == 1 {
		return
	}

	// Get indices
	i, j := 0, s.len-1
	if len(indices) == 2 {
		i = indices[0]
		j = indices[1]
	}

	// Convert to postive notation
	if i < 0 {
		i = s.len + i
	}
	if j < 0 {
		j = s.len + j
	}

	// Start can't be past end else nothing to get
	if i > j {
		return
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= s.len {
		j = s.len - 1
	}

	// Go has an exclusive behavior by default and we want inclusive
	// so offsetting the end by one
	j++

	result.len = j - i
	switch slice := s.o.(type) {
	case []bool:
		x := make([]bool, result.len, result.len)
		copy(x, slice[i:j])
		result.o = x
	case []int:
		x := make([]int, result.len, result.len)
		copy(x, slice[i:j])
		result.o = x
	case []string:
		x := make([]string, result.len, result.len)
		copy(x, slice[i:j])
		result.o = x
	default:
		v := reflect.ValueOf(s.o)
		x := reflect.MakeSlice(v.Type(), result.len, result.len)
		reflect.Copy(x, v.Slice(i, j))
		result.o = x.Interface()
	}
	return
}

// DeleteAt deletes the item at the given index location. Allows for negative notation.
// Returns the deleted item as a NObj which will be NObj.Nil() true if it didn't exist.
// Cost for reflection in this case doesn't seem to add much.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) DeleteAt(i int) (obj *NObj) {

	// Get the item and check out-of-bounds
	obj = s.At(i)
	if obj.Nil() {
		return
	}
	i = s.absIndex(i) // don't need bounds check as At call handles this

	// Delete the item
	switch slice := s.o.(type) {
	case []bool:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		s.o = slice
	case []int:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		s.o = slice
	case []string:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		s.o = slice
	default:
		v := reflect.ValueOf(s.o)
		if i+1 < v.Len() {
			v = reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len()))
		} else {
			v = v.Slice(0, i)
		}
		s.o = slice
	}
	s.len--
	return
}

// DropFirst deletes the first element and returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) DropFirst() *NSlice {
	n := 1
	if s.Nil() {
		return s
	} else if s.len >= n {
		switch slice := s.o.(type) {
		case []bool:
			slice = slice[n:]
			s.o = slice
		case []int:
			slice = slice[n:]
			s.o = slice
		case []string:
			slice = slice[n:]
			s.o = slice
		default:
			v := reflect.ValueOf(s.o)
			s.o = v.Slice(n, v.Len()).Interface()
		}
		s.len -= n
	} else {
		*s = *(newEmptySlice(s.o))
	}
	return s
}

// DropFirstN deletes first n elements and returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) DropFirstN(n int) *NSlice {
	if n == 0 {
		return s
	}
	if s.Nil() {
		return s
	} else if s.len >= n {
		switch slice := s.o.(type) {
		case []bool:
			slice = slice[n:]
			s.o = slice
		case []int:
			slice = slice[n:]
			s.o = slice
		case []string:
			slice = slice[n:]
			s.o = slice
		default:
			v := reflect.ValueOf(s.o)
			s.o = v.Slice(n, v.Len()).Interface()
		}
		s.len -= n
	} else {
		*s = *(newEmptySlice(s.o))
	}
	return s
}

// DropLast deletes last element returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) DropLast() *NSlice {
	n := 1
	if s.Nil() {
		return s
	} else if s.len >= n {
		switch slice := s.o.(type) {
		case []bool:
			slice = slice[:len(slice)-n]
			s.o = slice
		case []int:
			slice = slice[:len(slice)-n]
			s.o = slice
		case []string:
			slice = slice[:len(slice)-n]
			s.o = slice
		default:
			v := reflect.ValueOf(s.o)
			s.o = v.Slice(0, v.Len()-n).Interface()
		}
		s.len -= n
	} else {
		*s = *(newEmptySlice(s.o))
	}
	return s
}

// DropLastN deletes last n elements and returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) DropLastN(n int) *NSlice {
	if n == 0 {
		return s
	}
	if s.Nil() {
		return s
	} else if s.len >= n {
		switch slice := s.o.(type) {
		case []bool:
			slice = slice[:len(slice)-n]
			s.o = slice
		case []int:
			slice = slice[:len(slice)-n]
			s.o = slice
		case []string:
			slice = slice[:len(slice)-n]
			s.o = slice
		default:
			v := reflect.ValueOf(s.o)
			s.o = v.Slice(0, v.Len()-n).Interface()
		}
		s.len -= n
	} else {
		*s = *(newEmptySlice(s.o))
	}
	return s
}

// Each calls the given function once for each element in the numerable, passing that element in
// as a parameter. Returns a reference to the numerable
//
// Cost: ~0
//
// Optimized types: []bool, []int, []string
func (s *NSlice) Each(action func(O)) *NSlice {
	if s.Nil() {
		return s
	}
	switch slice := s.o.(type) {
	case []bool:
		for i := 0; i < len(slice); i++ {
			action(slice[i])
		}
	case []int:
		for i := 0; i < len(slice); i++ {
			action(slice[i])
		}
	case []string:
		for i := 0; i < len(slice); i++ {
			action(slice[i])
		}
	default:
		v := reflect.ValueOf(s.o)
		for i := 0; i < v.Len(); i++ {
			action(v.Index(i).Interface())
		}
	}
	return s
}

// EachE calls the given function once for each element in the numerable, passing that element in
// as a parameter. Returns a reference to the numerable and any error from the user function.
//
// Cost: ~0
//
// Optimized types: []bool, []int, []string
func (s *NSlice) EachE(action func(O) error) (*NSlice, error) {
	var err error
	if s.Nil() {
		return s, err
	}
	switch slice := s.o.(type) {
	case []bool:
		for i := 0; i < len(slice); i++ {
			if err = action(slice[i]); err != nil {
				return s, err
			}
		}
	case []int:
		for i := 0; i < len(slice); i++ {
			if err = action(slice[i]); err != nil {
				return s, err
			}
		}
	case []string:
		for i := 0; i < len(slice); i++ {
			if err = action(slice[i]); err != nil {
				return s, err
			}
		}
	default:
		v := reflect.ValueOf(s.o)
		for i := 0; i < v.Len(); i++ {
			if err = action(v.Index(i).Interface()); err != nil {
				return s, err
			}
		}
	}
	return s, err
}

// Empty tests if the numerable is empty.
//
// Cost: ~0x
func (s *NSlice) Empty() bool {
	if s.Nil() || s.len == 0 {
		return true
	}
	return false
}

// First returns the first element in the slice as NObj which will be NObj.Nil true if
// there are no elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) First() (obj *NObj) {
	return s.At(0)
}

// FirstN returns the first n elements in the slice as a NSlice. Best effort is used such
// that as many as can be will be returned up until the request is satisfied.
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) FirstN(n int) (result *NSlice) {
	j := n - 1
	if n < 0 {
		j = (n * -1) - 1
	}
	return s.Slice(0, j)
}

// Last returns the last element in the slice as NObj which will be NObj.Nil true if
// there are no elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) Last() *NObj {
	return s.At(-1)
}

// LastN returns the last n elements in the slice as a NSlice. Best effort is used such
// that as many as can be will be returned up until the request is satisfied.
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) LastN(n int) *NSlice {
	i := n * -1
	if n < 0 {
		i = n
	}
	return s.Slice(i, -1)
}

// // Join the underlying slice with the given delim
// func (s *strSliceN) Join(delim string) *strN {
// 	return A(strings.Join(s.v, delim))
// }

// Len returns the number of elements in the numerable
func (s *NSlice) Len() int {
	if s.Nil() {
		return 0
	}
	return s.len
}

// // // Map manipulates the slice into a new form
// // func (s *strSliceN) Map(sel func(string) O) (result *Numerable) {
// // 	for i := range s.v {
// // 		obj := sel(s.v[i])

// // 		// Drill into numerables
// // 		if s, ok := obj.(*Numerable); ok {
// // 			obj = s.v.Interface()
// // 		}

// // 		// Create new slice of the return type of sel
// // 		if result == nil {
// // 			typ := reflect.TypeOf(obj)
// // 			result = Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
// // 		}
// // 		result.Append(obj)
// // 	}
// // 	if result == nil {
// // 		result = Q([]interface{}{})
// // 	}
// // 	return
// // }

// // // MapF manipulates the numerable data into a new form then flattens
// // func (s *strSliceN) MapF(sel func(string) O) (result *Numerable) {
// // 	result = s.Map(sel).Flatten()
// // 	return
// // }

// // // Pair simply returns the first and second slice items
// // func (s *strSliceN) Pair() (first, second string) {
// // 	if s.Len() > 0 {
// // 		first = s.v[0]
// // 	}
// // 	if s.Len() > 1 {
// // 		second = s.v[1]
// // 	}
// // 	return
// // }

// Nil tests if the numerable is nil
func (s *NSlice) Nil() bool {
	if s == nil || s.o == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (s *NSlice) O() interface{} {
	if s.Nil() {
		return nil
	}
	return s.o
}

// // // Prepend items to the begining of the slice and return slice
// // func (s *strSliceN) Prepend(items ...string) *strSliceN {
// // 	items = append(items, s.v...)
// // 	s.v = items
// // 	return s
// // }

// // // Single simple report true if there is only one item
// // func (s *strSliceN) Single() (result bool) {
// // 	return s.Len() == 1
// // }

// Set the item at the given index location to the given item. Allows for negative notation.
// Returns the slice for chaining. Cost for reflection in this case doesn't seem to add much.
//
// Cost: ~1x - 10x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) Set(i int, obj interface{}) *NSlice {
	if s.Nil() {
		return s
	}
	if i = s.absIndex(i); i == -1 {
		panic("slice assignment is out of bounds")
	}

	var ok bool
	switch slice := s.o.(type) {
	case []bool:
		var x bool
		if x, ok = obj.(bool); ok {
			slice[i] = x
		}
	case []int:
		var x int
		if x, ok = obj.(int); ok {
			slice[i] = x
		}
	case []string:
		var x string
		if x, ok = obj.(string); ok {
			slice[i] = x
		}
	default:
		ok = true
		v := reflect.ValueOf(s.o)
		item := v.Index(i)
		item.Set(reflect.ValueOf(obj))
	}
	if !ok {
		panic(fmt.Sprintf("can't insert type '%T' into '%T'", obj, s.o))
	}
	return s
}

// Slice provides a Ruby like slice function for NSlice allowing for positive and negative notation.
// Slice uses an inclusive behavior such that Slice(0, -1) includes index -1 as opposed to Go's exclusive
// behavior. Out of bounds indices will be moved within bounds.
//
// An empty NSlice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. SliceV(1,2,3).Slice(0, -1) == [1,2,3] && SliceV(1,2,3).Slice(1,2) == [2,3]
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (s *NSlice) Slice(i, j int) (result *NSlice) {
	if s == nil {
		result = &NSlice{}
		return
	}
	result = newEmptySlice(s.o)
	if s.len == 0 {
		return
	}

	// Convert to postive notation
	if i < 0 {
		i = s.len + i
	}
	if j < 0 {
		j = s.len + j
	}

	// Start can't be past end else nothing to get
	if i > j {
		return
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= s.len {
		j = s.len - 1
	}

	// Go has an exclusive behavior by default and we want inclusive
	// so offsetting the end by one
	j++

	switch slice := s.o.(type) {
	case []bool:
		result.o = slice[i:j]
	case []int:
		result.o = slice[i:j]
	case []string:
		result.o = slice[i:j]
	default:
		v := reflect.ValueOf(s.o)
		result.o = v.Slice(i, j).Interface()
	}
	result.len = j - i
	return
}

// // Sort the underlying slice
// func (s *strSliceN) Sort() *strSliceN {
// 	sort.Strings(s.v)
// 	return s
// }

// // // TakeFirst updates the underlying slice and returns the item and status
// // func (s *strSliceN) TakeFirst() (string, bool) {
// // 	if len(s.v) > 0 {
// // 		item := s.v[0]
// // 		s.v = s.v[1:]
// // 		return item, true
// // 	}
// // 	return "", false
// // }

// // // TakeFirstCnt updates the underlying slice and returns the items
// // func (s *strSliceN) TakeFirstCnt(cnt int) (result *strSliceN) {
// // 	if cnt > 0 {
// // 		if len(s.v) >= cnt {
// // 			result = S(s.v[:cnt]...)
// // 			s.v = s.v[cnt:]
// // 		} else {
// // 			result = S(s.v...)
// // 			s.v = []string{}
// // 		}
// // 	}
// // 	return
// // }

// // // TakeLast updates the underlying slice and returns the item and status
// // func (s *strSliceN) TakeLast() (string, bool) {
// // 	if len(s.v) > 0 {
// // 		item := s.v[len(s.v)-1]
// // 		s.v = s.v[:len(s.v)-1]
// // 		return item, true
// // 	}
// // 	return "", false
// // }

// // // TakeLastCnt updates the underlying slice and returns a new nub
// // func (s *strSliceN) TakeLastCnt(cnt int) (result *strSliceN) {
// // 	if cnt > 0 {
// // 		if len(s.v) >= cnt {
// // 			i := len(s.v) - cnt
// // 			result = S(s.v[i:]...)
// // 			s.v = s.v[:i]
// // 		} else {
// // 			result = S(s.v...)
// // 			s.v = []string{}
// // 		}
// // 	}
// // 	return
// // }

// Type returns the identifier for this numerable type
func (s *NSlice) Type() Type {
	return NSliceType
}

// // // Uniq removes all duplicates from the underlying slice
// // func (s *strSliceN) Uniq() *strSliceN {
// // 	hits := map[string]bool{}
// // 	for i := len(s.v) - 1; i >= 0; i-- {
// // 		if _, exists := hits[s.v[i]]; !exists {
// // 			hits[s.v[i]] = true
// // 		} else {
// // 			s.v = append(s.v[:i], s.v[i+1:]...)
// // 		}
// // 	}
// // 	return s
// // }

// // // YamlPair return the first and second entries as yaml types
// // func (s *strSliceN) YamlPair() (first string, second interface{}) {
// // 	if s.Len() > 0 {
// // 		first = s.v[0]
// // 	}
// // 	if s.Len() > 1 {
// // 		second = A(s.v[1]).YamlType()
// // 	} else {
// // 		second = nil
// // 	}
// // 	return
// // }

// // // YamlKeyVal return the first and second entries as KeyVal of yaml types
// // func (s *strSliceN) YamlKeyVal() KeyVal {
// // 	result := KeyVal{}
// // 	if s.Len() > 0 {
// // 		result.Key = A(s.v[0]).YamlType()
// // 	}
// // 	if s.Len() > 1 {
// // 		result.Val = A(s.v[1]).YamlType()
// // 	} else {
// // 		result.Val = ""
// // 	}
// // 	return result
// // }

// // check if the internal type is a slice type
// func (q *NSlice) isSliceType() bool {
// 	return q.k == reflect.Array || q.k == reflect.Slice
// }
