// Package n provides many Go types with convenience functions reminiscent of Ruby or C#.
package n

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/pkg/errors"
)

// Slice provides a generic way to work with slice types providing convenience methods
// on par with other rapid development languages. 'this Slice' refers to the current slice
// instance being operated on.  'new Slice' refers to a copy of the slice based on a new
// underlying Array.
type Slice interface {
	Any(elems ...interface{}) bool                    // Any tests if this Slice is not empty or optionally if it contains any of the given variadic elements.
	AnyS(slice interface{}) bool                      // AnyS tests if this Slice contains any of the given Slice's elements.
	AnyW(sel func(O) bool) bool                       // AnyW tests if this Slice contains any that match the lambda selector.
	Append(elem interface{}) Slice                    // Append an element to the end of this Slice and returns a reference to this Slice.
	AppendV(elems ...interface{}) Slice               // AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
	At(i int) (elem *Object)                          // At returns the element at the given index location. Allows for negative notation.
	Clear() Slice                                     // Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
	Concat(slice interface{}) (new Slice)             // Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
	ConcatM(slice interface{}) Slice                  // ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
	Copy(indices ...int) (new Slice)                  // Copy returns a new Slice with the indicated range of elements copied from this Slice.
	Count(elem interface{}) (cnt int)                 // Count the number of elements in this Slice equal to the given element.
	CountW(sel func(O) bool) (cnt int)                // CountW counts the number of elements in this Slice that match the lambda selector.
	Drop(indices ...int) Slice                        // Drop modifies this Slice to delete the indicated range of elements and returns a referece to this Slice.
	DropAt(i int) Slice                               // DropAt modifies this Slice to delete the element at the given index location. Allows for negative notation.
	DropFirst() Slice                                 // DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
	DropFirstN(n int) Slice                           // DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
	DropLast() Slice                                  // DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
	DropLastN(n int) Slice                            // DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
	DropW(sel func(O) bool) Slice                     // DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
	Each(action func(O)) Slice                        // Each calls the given lambda once for each element in this Slice, passing in that element
	EachE(action func(O) error) (Slice, error)        // EachE calls the given lambda once for each element in this Slice, passing in that element
	EachI(action func(int, O)) Slice                  // EachI calls the given lambda once for each element in this Slice, passing in the index and element
	EachIE(action func(int, O) error) (Slice, error)  // EachIE calls the given lambda once for each element in this Slice, passing in the index and element
	EachR(action func(O)) Slice                       // EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
	EachRE(action func(O) error) (Slice, error)       // EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
	EachRI(action func(int, O)) Slice                 // EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
	EachRIE(action func(int, O) error) (Slice, error) // EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
	Empty() bool                                      // Empty tests if this Slice is empty.
	First() (elem *Object)                            // First returns the first element in this Slice as Object.
	FirstN(n int) Slice                               // FirstN returns the first n elements in this slice as a Slice reference to the original.
	Index(elem interface{}) (loc int)                 // Index returns the index of the first element in this Slice where element == elem
	Insert(i int, elem interface{}) Slice             // Insert modifies this Slice to insert the given element before the element with the given index.
	Join(separator ...string) (str *Object)           // Join converts each element into a string then joins them together using the given separator or comma by default.
	Last() (elem *Object)                             // Last returns the last element in this Slice as an Object.
	LastN(n int) Slice                                // LastN returns the last n elements in this Slice as a Slice reference to the original.
	Len() int                                         // Len returns the number of elements in this Slice.
	Less(i, j int) bool                               // Less returns true if the element indexed by i is less than the element indexed by j.
	Nil() bool                                        // Nil tests if this Slice is nil.
	O() interface{}                                   // O returns the underlying data structure as is.
	Pair() (first, second *Object)                    // Pair simply returns the first and second Slice elements as Objects.
	Pop() (elem *Object)                              // Pop modifies this Slice to remove the last element and returns the removed element as an Object.
	PopN(n int) (new Slice)                           // PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
	Prepend(elem interface{}) Slice                   // Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
	Reverse() (new Slice)                             // Reverse returns a new Slice with the order of the elements reversed.
	ReverseM() Slice                                  // ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
	Select(sel func(O) bool) (new Slice)              // Select creates a new slice with the elements that match the lambda selector.
	Set(i int, elem interface{}) Slice                // Set the element at the given index location to the given element. Allows for negative notation.
	SetE(i int, elem interface{}) (Slice, error)      // Set the element at the given index location to the given element. Allows for negative notation.
	Shift() (elem *Object)                            // Shift modifies this Slice to remove the first element and returns the removed element as an Object.
	ShiftN(n int) (new Slice)                         // ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
	Single() bool                                     // Single reports true if there is only one element in this Slice.
	Slice(indices ...int) Slice                       // Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
	Sort() (new Slice)                                // Sort returns a new Slice with sorted elements.
	SortM() Slice                                     // SortM modifies this Slice sorting the elements and returns a reference to this Slice.
	SortReverse() (new Slice)                         // SortReverse returns a new Slice sorting the elements in reverse.
	SortReverseM() Slice                              // SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
	String() string                                   // Returns a string representation of this Slice, implements the Stringer interface
	Swap(i, j int)                                    // Swap modifies this Slice swapping the indicated elements.
	Take(indices ...int) (new Slice)                  // Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
	TakeAt(i int) (elem *Object)                      // TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
	TakeW(sel func(O) bool) (new Slice)               // TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
	Union(slice interface{}) (new Slice)              // Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
	UnionM(slice interface{}) Slice                   // UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
	Uniq() (new Slice)                                // Uniq returns a new Slice with all non uniq elements removed while preserving element order.
	UniqM() Slice                                     // UniqM modifies this Slice to remove all non uniq elements while preserving element order.
}

// NewSlice provides a generic way to work with Slice types. It does this by wrapping Go types
// directly for optimized types thus avoiding reflection processing overhead and making a plethora
// of Slice methods available. Non optimized types will fall back on reflection to generically
// handle the type incurring the full 10x reflection processing overhead.
//
// Optimized: []int
func NewSlice(obj interface{}) (new Slice) {
	switch x := obj.(type) {
	case []int:
		return NewIntSlice(x)
	case *[]int:
		return NewIntSlice(*x)
	default:
		//return NewRefectionSlice(obj)
		return nil
	}
}

// NewSliceV creates a new Slice encapsulating the given variadic elements in a new Slice of
// that type using type assertion for optimized types. Non optimized types will fall back
// on reflection to generically handle the type incurring the full 10x reflection processing
// overhead. In the case where nothing is given a new *RefSlice will be returned.
//
// Optimized: []int
func NewSliceV(elems ...interface{}) (new Slice) {

	// // Return NSlice.Nil if nothing given
	// if len(items) == 0 {
	// 	return
	// }

	// // Create new slice from the type of the first non Invalid element
	// var slice *reflect.Value
	// for i := 0; i < len(items); i++ {

	// 	// Create target slice from first Valid element
	// 	if slice == nil && reflect.ValueOf(items[i]).IsValid() {
	// 		typ := reflect.SliceOf(reflect.TypeOf(items[i]))
	// 		v := reflect.MakeSlice(typ, 0, 10)
	// 		slice = &v
	// 	}

	// 	// Append item to slice
	// 	if slice != nil {
	// 		item := reflect.ValueOf(items[i])
	// 		*slice = reflect.Append(*slice, item)
	// 	}
	// }
	// if slice != nil {
	// 	n.o = slice.Interface()
	// 	n.len = slice.Len()
	// }
	// return
	return nil
}

// NSlice provides a generic way to work with slice types providing convenience methods
// on par with other rapid development languages.
//
// Implements the Slice interface.
type NSlice struct {
	o   interface{} // underlying slice object
	len int         // slice length
}

// OldSlice creates a new NSlice by simply storing slice 'obj' directly to avoid using reflection
// processing at a 10x overhead savings. Non slice 'obj' are encapsulated in a new slice of
// that type using reflection, thus incurring the standard 10x overhead.
//
// Return value n *NSlice will never be nil but n.Nil() may be true as nil or empty []interface{}
// values are ignored to avoid internally using a []interface{}. The internal type will be
// set later with the given type when an n.AppendX method is called.
//
// Cost: ~0x - 10x
func OldSlice(obj interface{}) (n *NSlice) {
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
		n = OldSliceV(x...)

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

// OldSliceV creates a new NSlice encapsulating the given variadic elements in a new slice of
// that type. Thi call incurs the full 10x reflection overhead. For large slice params use
// the Slice() func instead.
//
// Cost: ~10x
func OldSliceV(items ...interface{}) (n *NSlice) {
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
func oldEmptySlice(items interface{}) (n *NSlice) {
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
func (p *NSlice) Any(obj ...interface{}) bool {

	// No elements
	if p.Nil() || p.len == 0 {
		return false
	}

	// Elements and not looking for anything
	if len(obj) == 0 {
		return true
	}

	// Looking for something specific
	ok := false
	var typ reflect.Type
	switch slice := p.o.(type) {
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
		v := reflect.ValueOf(p.o)
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
		panic(fmt.Sprintf("can't compare type '%v' with '%T' elements", typ, p.o))
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
func (p *NSlice) AnyS(obj interface{}) (result bool) {
	if p.Nil() {
		return
	}
	ok := false
	switch slice := p.o.(type) {
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
		v := reflect.ValueOf(p.o)
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
		panic(fmt.Sprintf("can't compare type '%T' with '%T' elements", obj, p.o))
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
func (p *NSlice) Append(item interface{}) *NSlice {
	if p.Nil() {
		if p == nil {
			p = oldEmptySlice(item)
		} else {
			*p = *(oldEmptySlice(item))
		}
	}
	ok := false
	switch slice := p.o.(type) {
	case []bool:
		var x bool
		if x, ok = item.(bool); ok {
			p.o = append(slice, x)
		}
	case []int:
		var x int
		if x, ok = item.(int); ok {
			p.o = append(slice, x)
		}
	case []string:
		var x string
		if x, ok = item.(string); ok {
			p.o = append(slice, x)
		}
	default:
		ok = true
		v := reflect.ValueOf(p.o)
		item := reflect.ValueOf(item)
		p.o = reflect.Append(v, item).Interface()
	}
	if !ok {
		panic(fmt.Sprintf("can't append type '%T' to '%T'", item, p.o))
	}
	p.len++
	return p
}

// AppendV appends the variadic items to the end of the NSlice and returns the NSlice for chaining.
// Avoids the 10x reflection overhead cost by type asserting common types. Types not optimized in
// this way incur the full 10x reflection overhead cost.
//
// Cost: ~6x - 10x
//
// Optimized types: bool, int, string
func (p *NSlice) AppendV(items ...interface{}) *NSlice {
	for _, item := range items {
		p.Append(item)
	}
	return p
}

// AppendS appends the given slice using variadic expansion and returns NSlice for chaining. Avoids
// the 10x reflection overhead cost by type asserting common types. Types not optimized in this
// way incur the full 10x reflection overhead cost. However when appending larger slices fewer times
// the cost reduces down to 2x.
//
// Cost: ~0x - 2x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) AppendS(items interface{}) *NSlice {
	if p.Nil() {
		if p == nil {
			p = oldEmptySlice(items)
		} else {
			*p = *(oldEmptySlice(items))
		}
	}
	ok := false
	switch slice := p.o.(type) {
	case []bool:
		var x []bool
		if x, ok = items.([]bool); ok {
			p.o = append(slice, x...)
			p.len += len(x)
		}
	case []int:
		var x []int
		if x, ok = items.([]int); ok {
			p.o = append(slice, x...)
			p.len += len(x)
		}
	case []string:
		var x []string
		if x, ok = items.([]string); ok {
			p.o = append(slice, x...)
			p.len += len(x)
		}
	default:
		ok = true
		v := reflect.ValueOf(p.o)
		x := reflect.ValueOf(items)
		p.o = reflect.AppendSlice(v, x).Interface()
		p.len += x.Len()
	}
	if !ok {
		panic(fmt.Sprintf("can't concat type '%T' with '%T'", items, p.o))
	}
	return p
}

// At returns the item at the given index location. Allows for negative notation.
// Cost for reflection in this case doesn't seem to add much.
//
// Cost: ~0x - 2x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) At(i int) (obj *Object) {
	obj = &Object{}
	if p.Nil() {
		return
	}
	if i = absIndex(p.len, i); i == -1 {
		return
	}

	switch slice := p.o.(type) {
	case []bool:
		obj.o = slice[i]
	case []int:
		obj.o = slice[i]
	case []string:
		obj.o = slice[i]
	default:
		obj.o = reflect.ValueOf(p.o).Index(i).Interface()
	}
	return
}

// Clear the underlying slice.
//
// Cost: constant
func (p *NSlice) Clear() *NSlice {
	if !p.Nil() {
		*p = *(oldEmptySlice(p.o))
	}
	return p
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
func (p *NSlice) Copy(indices ...int) (result *NSlice) {
	if p == nil {
		result = &NSlice{}
		return
	}
	result = oldEmptySlice(p.o)
	if p.len == 0 || len(indices) == 1 {
		return
	}

	// Get indices
	i, j := 0, p.len-1
	if len(indices) == 2 {
		i = indices[0]
		j = indices[1]
	}

	// Convert to postive notation
	if i < 0 {
		i = p.len + i
	}
	if j < 0 {
		j = p.len + j
	}

	// Start can't be past end else nothing to get
	if i > j {
		return
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= p.len {
		j = p.len - 1
	}

	// Go has an exclusive behavior by default and we want inclusive
	// so offsetting the end by one
	j++

	result.len = j - i
	switch slice := p.o.(type) {
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
		v := reflect.ValueOf(p.o)
		x := reflect.MakeSlice(v.Type(), result.len, result.len)
		reflect.Copy(x, v.Slice(i, j))
		result.o = x.Interface()
	}
	return
}

// Drop deletes the item at the given index location. Allows for negative notation.
// Returns the rest of the elements in the slice for chaining.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Drop(i int) *NSlice {
	if p.Nil() {
		return p
	}
	if i = absIndex(p.len, i); i == -1 {
		return p
	}

	// Delete the item
	switch slice := p.o.(type) {
	case []bool:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		p.o = slice
	case []int:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		p.o = slice
	case []string:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		p.o = slice
	default:
		v := reflect.ValueOf(p.o)
		if i+1 < v.Len() {
			v = reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len()))
		} else {
			v = v.Slice(0, i)
		}
		p.o = v.Interface()
	}
	p.len--
	return p
}

// DropFirst deletes the first element and returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) DropFirst() *NSlice {
	n := 1
	if p.Nil() {
		return p
	} else if p.len >= n {
		switch slice := p.o.(type) {
		case []bool:
			slice = slice[n:]
			p.o = slice
		case []int:
			slice = slice[n:]
			p.o = slice
		case []string:
			slice = slice[n:]
			p.o = slice
		default:
			v := reflect.ValueOf(p.o)
			p.o = v.Slice(n, v.Len()).Interface()
		}
		p.len -= n
	} else {
		*p = *(oldEmptySlice(p.o))
	}
	return p
}

// DropFirstN deletes first n elements and returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) DropFirstN(n int) *NSlice {
	if n == 0 {
		return p
	}
	if p.Nil() {
		return p
	} else if p.len >= n {
		switch slice := p.o.(type) {
		case []bool:
			slice = slice[n:]
			p.o = slice
		case []int:
			slice = slice[n:]
			p.o = slice
		case []string:
			slice = slice[n:]
			p.o = slice
		default:
			v := reflect.ValueOf(p.o)
			p.o = v.Slice(n, v.Len()).Interface()
		}
		p.len -= n
	} else {
		*p = *(oldEmptySlice(p.o))
	}
	return p
}

// DropLast deletes last element returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) DropLast() *NSlice {
	n := 1
	if p.Nil() {
		return p
	} else if p.len >= n {
		switch slice := p.o.(type) {
		case []bool:
			slice = slice[:len(slice)-n]
			p.o = slice
		case []int:
			slice = slice[:len(slice)-n]
			p.o = slice
		case []string:
			slice = slice[:len(slice)-n]
			p.o = slice
		default:
			v := reflect.ValueOf(p.o)
			p.o = v.Slice(0, v.Len()-n).Interface()
		}
		p.len -= n
	} else {
		*p = *(oldEmptySlice(p.o))
	}
	return p
}

// DropLastN deletes last n elements and returns the rest of the elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) DropLastN(n int) *NSlice {
	if n == 0 {
		return p
	}
	if p.Nil() {
		return p
	} else if p.len >= n {
		switch slice := p.o.(type) {
		case []bool:
			slice = slice[:len(slice)-n]
			p.o = slice
		case []int:
			slice = slice[:len(slice)-n]
			p.o = slice
		case []string:
			slice = slice[:len(slice)-n]
			p.o = slice
		default:
			v := reflect.ValueOf(p.o)
			p.o = v.Slice(0, v.Len()-n).Interface()
		}
		p.len -= n
	} else {
		*p = *(oldEmptySlice(p.o))
	}
	return p
}

// Each calls the given function once for each element in the numerable, passing that element in
// as a parameter. Returns a reference to the numerable
//
// Cost: ~0
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Each(action func(O)) *NSlice {
	if p.Nil() {
		return p
	}
	switch slice := p.o.(type) {
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
		v := reflect.ValueOf(p.o)
		for i := 0; i < v.Len(); i++ {
			action(v.Index(i).Interface())
		}
	}
	return p
}

// EachE calls the given function once for each element in the numerable, passing that element in
// as a parameter. Returns a reference to the numerable and any error from the user function.
//
// Cost: ~0
//
// Optimized types: []bool, []int, []string
func (p *NSlice) EachE(action func(O) error) (*NSlice, error) {
	var err error
	if p.Nil() {
		return p, err
	}
	switch slice := p.o.(type) {
	case []bool:
		for i := 0; i < len(slice); i++ {
			if err = action(slice[i]); err != nil {
				return p, err
			}
		}
	case []int:
		for i := 0; i < len(slice); i++ {
			if err = action(slice[i]); err != nil {
				return p, err
			}
		}
	case []string:
		for i := 0; i < len(slice); i++ {
			if err = action(slice[i]); err != nil {
				return p, err
			}
		}
	default:
		v := reflect.ValueOf(p.o)
		for i := 0; i < v.Len(); i++ {
			if err = action(v.Index(i).Interface()); err != nil {
				return p, err
			}
		}
	}
	return p, err
}

// Empty tests if the numerable is empty.
//
// Cost: ~0x
func (p *NSlice) Empty() bool {
	if p.Nil() || p.len == 0 {
		return true
	}
	return false
}

// First returns the first element in the slice as Object which will be Object.Nil true if
// there are no elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) First() (obj *Object) {
	return p.At(0)
}

// FirstN returns the first n elements in the slice as a NSlice. Best effort is used such
// that as many as can be will be returned up until the request is satisfied.
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) FirstN(n int) (result *NSlice) {
	j := n - 1
	if n < 0 {
		j = (n * -1) - 1
	}
	return p.Slice(0, j)
}

// Insert the given value before the element with the given index. Negative indices count
// backwards from the end of the array, where -1 is the last element. If a negative index
// is used, the given values will be inserted after that element, so using an index of -1
// will insert the values at the end of the slice. Slice is returned for chaining. Invalid
// index locations will not change the slice.
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Insert(i int, obj interface{}) *NSlice {
	if p.Nil() {
		p.Append(obj)
		return p
	}
	j := i
	if j = absIndex(p.len, j); j == -1 {
		return p
	}
	if i < 0 {
		j++
	}

	// Insert the item before j if pos and after j if neg
	var ok bool
	switch slice := p.o.(type) {
	case []bool:
		var x bool
		if x, ok = obj.(bool); ok {
			if j == 0 {
				slice = append([]bool{x}, slice...)
			} else if j < len(slice) {
				slice = append(slice, x)
				copy(slice[j+1:], slice[j:])
				slice[j] = x
			} else {
				slice = append(slice, x)
			}
			p.o = slice
		}
	case []int:
		var x int
		if x, ok = obj.(int); ok {
			if j == 0 {
				slice = append([]int{x}, slice...)
			} else if j < len(slice) {
				slice = append(slice, x)
				copy(slice[j+1:], slice[j:])
				slice[j] = x
			} else {
				slice = append(slice, x)
			}
			p.o = slice
		}
	case []string:
		var x string
		if x, ok = obj.(string); ok {
			if j == 0 {
				slice = append([]string{x}, slice...)
			} else if j < len(slice) {
				slice = append(slice, x)
				copy(slice[j+1:], slice[j:])
				slice[j] = x
			} else {
				slice = append(slice, x)
			}
			p.o = slice
		}
	default:
		ok = true
		v := reflect.ValueOf(p.o)
		x := reflect.ValueOf(obj)
		if j == 0 {
			new := reflect.MakeSlice(reflect.SliceOf(x.Type()), 1, 1)
			new.Index(0).Set(x)
			v = reflect.AppendSlice(new, v.Slice(0, v.Len()))
		} else if j < p.len {
			v = reflect.Append(v, x)
			reflect.Copy(v.Slice(j+1, v.Len()), v.Slice(j, v.Len()))
			v.Index(j).Set(x)
		} else {
			v = reflect.Append(v, x)
		}
		p.o = v.Interface()
	}
	if !ok {
		panic(fmt.Sprintf("can't insert type '%T' into '%T'", obj, p.o))
	}
	p.len++
	return p
}

// Last returns the last element in the slice as Object which will be Object.Nil true if
// there are no elements in the slice.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Last() *Object {
	return p.At(-1)
}

// LastN returns the last n elements in the slice as a NSlice. Best effort is used such
// that as many as can be will be returned up until the request is satisfied.
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) LastN(n int) *NSlice {
	i := n * -1
	if n < 0 {
		i = n
	}
	return p.Slice(i, -1)
}

// // Join the underlying slice with the given delim
// func (p *strSliceN) Join(delim string) *strN {
// 	return A(strings.Join(s.v, delim))
// }

// Len returns the number of elements in the numerable
// Implements the sort.Interface
func (p *NSlice) Len() int {
	if p.Nil() {
		return 0
	}
	return p.len
}

// Less returns true if the object indexed by i is less than the object indexed by j.
// Implements the sort.Interface
//
// Cost: ~0x - 30x
//
// Optimized types: []bool, []int, []string, []Object
func (p *NSlice) Less(i, j int) (result bool) {
	if p.Nil() || p.len < 2 || i < 0 || j < 0 || i >= p.len || j >= p.len {
		return
	}

	switch slice := p.o.(type) {
	case []bool:
		result = (slice[i] == false && slice[j] == true)
	case []int:
		result = slice[i] < slice[j]
	case []string:
		result = slice[i] < slice[j]
	case []Object:
		result = slice[i].Less(slice[j])
	default:
		v := reflect.ValueOf(p.o)
		x, y := Indirect(v.Index(i)), Indirect(v.Index(j))
		if comparable, ok := x.Interface().(Comparable); ok {
			if comparable.Less(y.Interface()) {
				result = true
			}
		}
	}
	return
}

// // // Map manipulates the slice into a new form
// // func (p *strSliceN) Map(sel func(string) O) (result *Numerable) {
// // 	for i := range s.v {
// // 		obj := sel(p.v[i])

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
// // func (p *strSliceN) MapF(sel func(string) O) (result *Numerable) {
// // 	result = s.Map(sel).Flatten()
// // 	return
// // }

// Nil tests if the numerable is nil
func (p *NSlice) Nil() bool {
	if p == nil || p.o == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *NSlice) O() interface{} {
	if p.Nil() {
		return nil
	}
	return p.o
}

// Pair simply returns the first and second slice items as Object
func (p *NSlice) Pair() (first, second *Object) {
	first, second = &Object{}, &Object{}
	if p.len > 0 {
		first = p.At(0)
	}
	if p.len > 1 {
		second = p.At(1)
	}
	return
}

// Prepend the given value at the begining of the slice.
//
// Cost: ~0x - 10x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Prepend(obj interface{}) *NSlice {
	return p.Insert(0, obj)
}

// Set the item at the given index location to the given item. Allows for negative notation.
// Returns the slice for chaining. Cost for reflection in this case doesn't seem to add much.
//
// Cost: ~1x - 10x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Set(i int, obj interface{}) *NSlice {
	if p.Nil() {
		return p
	}
	if i = absIndex(p.len, i); i == -1 {
		panic("slice assignment is out of bounds")
	}

	var ok bool
	switch slice := p.o.(type) {
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
		v := reflect.ValueOf(p.o)
		item := v.Index(i)
		item.Set(reflect.ValueOf(obj))
	}
	if !ok {
		panic(fmt.Sprintf("can't set type '%T' in '%T'", obj, p.o))
	}
	return p
}

// Single simply reports true if there is only one item
func (p *NSlice) Single() bool {
	return p.len == 1
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
func (p *NSlice) Slice(i, j int) (result *NSlice) {
	if p == nil {
		result = &NSlice{}
		return
	}
	result = oldEmptySlice(p.o)
	if p.len == 0 {
		return
	}

	// Convert to postive notation
	if i < 0 {
		i = p.len + i
	}
	if j < 0 {
		j = p.len + j
	}

	// Start can't be past end else nothing to get
	if i > j {
		return
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= p.len {
		j = p.len - 1
	}

	// Go has an exclusive behavior by default and we want inclusive
	// so offsetting the end by one
	j++

	switch slice := p.o.(type) {
	case []bool:
		result.o = slice[i:j]
	case []int:
		result.o = slice[i:j]
	case []string:
		result.o = slice[i:j]
	default:
		v := reflect.ValueOf(p.o)
		result.o = v.Slice(i, j).Interface()
	}
	result.len = j - i
	return
}

// Sort the underlying slice and return a pointer for chaining.
// Reflection cost is indirectly doubled, for non-optimized types, as we are
// swapping two items repeatedly during the quick sort.
//
// Cost: ~0x - 20x
//
// Optimized types: []bool, []int, []string, sort.Interface
func (p *NSlice) Sort() *NSlice {
	if p.Nil() || p.len < 2 {
		return p
	}
	switch slice := p.o.(type) {
	case []bool:
		//sort.Sort(nutil.BoolSlice(slice))
	case []int:
		sort.Sort(sort.IntSlice(slice))
	case []string:
		sort.Sort(sort.StringSlice(slice))
	default:
		if sorter, ok := p.o.(sort.Interface); ok {
			sort.Sort(sorter)
		}
	}
	return p
}

// Swap elements in the underlying slice. Implements the sort.Interface.
// Takes advantage of underlying slice's sort.Interface implementations if they exist.
// Reflection cost is doubled, for non-optimized types, as two items being reflected over.
//
// Cost: ~0x - 20x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Swap(i, j int) {
	if p.Nil() || p.len < 2 || i < 0 || j < 0 || i >= p.len || j >= p.len {
		return
	}
	if x, ok := p.o.([]bool); ok {
		x[i], x[j] = x[j], x[i]
	} else if x, ok := p.o.([]int); ok {
		x[i], x[j] = x[j], x[i]
	} else if x, ok := p.o.([]string); ok {
		x[i], x[j] = x[j], x[i]
	} else if x, ok := p.o.(sort.Interface); ok {
		x.Swap(i, j)
	} else {
		v := reflect.ValueOf(p.o)
		x, y := v.Index(i).Interface(), v.Index(j).Interface()
		v.Index(i).Set(reflect.ValueOf(y))
		v.Index(j).Set(reflect.ValueOf(x))
		p.o = v.Interface()
	}
}

// Take deletes the item at the given index location and returns it as an *Object which
// will be Object.Nil() true if it didn't exist. Allows for negative notation.
//
// Cost: ~0x - 3x
//
// Optimized types: []bool, []int, []string
func (p *NSlice) Take(i int) (obj *Object) {

	// Get the item and check out-of-bounds
	obj = p.At(i)
	if obj.Nil() {
		return
	}
	i = absIndex(p.len, i) // don't need bounds check as At call handles this

	// Delete the item
	switch slice := p.o.(type) {
	case []bool:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		p.o = slice
	case []int:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		p.o = slice
	case []string:
		if i+1 < len(slice) {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			slice = slice[:i]
		}
		p.o = slice
	default:
		v := reflect.ValueOf(p.o)
		if i+1 < v.Len() {
			v = reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len()))
		} else {
			v = v.Slice(0, i)
		}
		p.o = v.Interface()
	}
	p.len--
	return
}

// // // Uniq removes all duplicates from the underlying slice
// // func (p *strSliceN) Uniq() *strSliceN {
// // 	hits := map[string]bool{}
// // 	for i := len(p.v) - 1; i >= 0; i-- {
// // 		if _, exists := hits[s.v[i]]; !exists {
// // 			hits[s.v[i]] = true
// // 		} else {
// // 			s.v = append(p.v[:i], s.v[i+1:]...)
// // 		}
// // 	}
// // 	return p
// // }

// // // YamlPair return the first and second entries as yaml types
// // func (p *strSliceN) YamlPair() (first string, second interface{}) {
// // 	if p.len() > 0 {
// // 		first = s.v[0]
// // 	}
// // 	if p.len() > 1 {
// // 		second = A(p.v[1]).YamlType()
// // 	} else {
// // 		second = nil
// // 	}
// // 	return
// // }

// // // YamlKeyVal return the first and second entries as KeyVal of yaml types
// // func (p *strSliceN) YamlKeyVal() KeyVal {
// // 	result := KeyVal{}
// // 	if p.len() > 0 {
// // 		result.Key = A(p.v[0]).YamlType()
// // 	}
// // 	if p.len() > 1 {
// // 		result.Val = A(p.v[1]).YamlType()
// // 	} else {
// // 		result.Val = ""
// // 	}
// // 	return result
// // }

// // check if the internal type is a slice type
// func (q *NSlice) isSliceType() bool {
// 	return q.k == reflect.Array || q.k == reflect.Slice
// }

// simply pass positive values through and convert negative to positive
func abs(i int) (abs int) {
	if i < 0 {
		abs = i * -1
	} else {
		abs = i
	}
	return
}

// simply pass negative values through and convert positive to negative
func absNeg(i int) (abs int) {
	if i < 0 {
		abs = i
	} else {
		abs = i * -1
	}
	return
}

// get the absolute value for the pos/neg index.
// return of -1 indicates out of bounds
func absIndex(len, i int) (abs int) {
	if i < 0 {
		abs = len + i
	} else {
		abs = i
	}
	if abs < 0 || abs >= len {
		abs = -1
	}
	return
}

// convert to positive notation, move them within bounds
// returns an error if mutually exclusive
func absIndices(l int, indices ...int) (i int, j int, err error) {

	// Get indices must be either none or two
	i, j = 0, -1
	if len(indices) == 2 {
		i = indices[0]
		j = indices[1]
	} else if len(indices) == 1 {
		err = errors.Errorf("only one index given")
		return
	}

	// Convert to postive notation
	if i < 0 {
		i = l + i
	}
	if j < 0 {
		j = l + j
	}

	// Start can't be past end else invalid
	if i > j {
		err = errors.Errorf("indices are mutually exclusive")
		return
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= l {
		j = l - 1
	}

	// Go has an exclusive behavior by default and we want inclusive
	// so offsetting the end by one
	j++

	return
}
