// Package n provides many Go types with convenience functions reminiscent of Ruby or C#.
package n

import (
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
	Generic() bool                                    // Generic returns true if the underlying implementation is a RefSlice
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
	case []string:
		return NewStringSlice(x)
	case *[]string:
		return NewStringSlice(*x)
	default:
		return NewRefSlice(obj)
	}
}

// NewSliceV creates a new Slice encapsulating the given variadic elements in a new Slice of
// that type using type assertion for optimized types. Non optimized types will fall back
// on reflection to generically handle the type incurring the full 10x reflection processing
// overhead. In the case where nothing is given a new *RefSlice will be returned.
//
// Optimized: []int
func NewSliceV(elems ...interface{}) (new Slice) {
	if len(elems) == 0 {
		new = NewRefSliceV(elems...)
	} else {
		switch elems[0].(type) {
		case int:
			var slice []int
			for _, x := range elems {
				slice = append(slice, x.(int))
			}
			new = NewIntSlice(slice)
		case *int:
			var slice []int
			for _, x := range elems {
				slice = append(slice, *(x.(*int)))
			}
			new = NewIntSlice(slice)
		case string:
			var slice []string
			for _, x := range elems {
				slice = append(slice, x.(string))
			}
			new = NewStringSlice(slice)
		case *string:
			var slice []string
			for _, x := range elems {
				slice = append(slice, *(x.(*string)))
			}
			new = NewStringSlice(slice)

		default:
			new = NewRefSliceV(elems...)
		}
	}
	return
}

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
