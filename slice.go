package n

import (
	"github.com/pkg/errors"
)

// ISlice provides a generic way to work with slice types providing convenience methods
// on par with rapid development languages. 'this Slice' refers to the current slice
// instance being operated on.  'new Slice' refers to a copy of the slice based on a new
// underlying Array.
type ISlice interface {
	A() string                                         // A is an alias to String for brevity
	All(elems ...interface{}) bool                     // All tests if this Slice is not empty or optionally if it contains all of the given variadic elements.
	AllS(slice interface{}) bool                       // AnyS tests if this Slice contains all of the given Slice's elements.
	Any(elems ...interface{}) bool                     // Any tests if this Slice is not empty or optionally if it contains any of the given variadic elements.
	AnyS(slice interface{}) bool                       // AnyS tests if this Slice contains any of the given Slice's elements.
	AnyW(sel func(O) bool) bool                        // AnyW tests if this Slice contains any that match the lambda selector.
	Append(elem interface{}) ISlice                    // Append an element to the end of this Slice and returns a reference to this Slice.
	AppendV(elems ...interface{}) ISlice               // AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
	At(i int) (elem *Object)                           // At returns the element at the given index location. Allows for negative notation.
	Clear() ISlice                                     // Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
	Concat(slice interface{}) (new ISlice)             // Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
	ConcatM(slice interface{}) ISlice                  // ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
	Copy(indices ...int) (new ISlice)                  // Copy returns a new Slice with the indicated range of elements copied from this Slice.
	Count(elem interface{}) (cnt int)                  // Count the number of elements in this Slice equal to the given element.
	CountW(sel func(O) bool) (cnt int)                 // CountW counts the number of elements in this Slice that match the lambda selector.
	Drop(indices ...int) ISlice                        // Drop modifies this Slice to delete the indicated range of elements and returns a referece to this Slice.
	DropAt(i int) ISlice                               // DropAt modifies this Slice to delete the element at the given index location. Allows for negative notation.
	DropFirst() ISlice                                 // DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
	DropFirstN(n int) ISlice                           // DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
	DropLast() ISlice                                  // DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
	DropLastN(n int) ISlice                            // DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
	DropW(sel func(O) bool) ISlice                     // DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
	Each(action func(O)) ISlice                        // Each calls the given lambda once for each element in this Slice, passing in that element
	EachE(action func(O) error) (ISlice, error)        // EachE calls the given lambda once for each element in this Slice, passing in that element
	EachI(action func(int, O)) ISlice                  // EachI calls the given lambda once for each element in this Slice, passing in the index and element
	EachIE(action func(int, O) error) (ISlice, error)  // EachIE calls the given lambda once for each element in this Slice, passing in the index and element
	EachR(action func(O)) ISlice                       // EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
	EachRE(action func(O) error) (ISlice, error)       // EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
	EachRI(action func(int, O)) ISlice                 // EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
	EachRIE(action func(int, O) error) (ISlice, error) // EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
	Empty() bool                                       // Empty tests if this Slice is empty.
	First() (elem *Object)                             // First returns the first element in this Slice as Object.
	FirstN(n int) ISlice                               // FirstN returns the first n elements in this slice as a Slice reference to the original.
	InterSlice() bool                                  // Generic returns true if the underlying implementation uses reflection
	Index(elem interface{}) (loc int)                  // Index returns the index of the first element in this Slice where element == elem
	Insert(i int, elem interface{}) ISlice             // Insert modifies this Slice to insert the given element(s) before the element with the given index.
	Join(separator ...string) (str *Object)            // Join converts each element into a string then joins them together using the given separator or comma by default.
	Last() (elem *Object)                              // Last returns the last element in this Slice as an Object.
	LastN(n int) ISlice                                // LastN returns the last n elements in this Slice as a Slice reference to the original.
	Len() int                                          // Len returns the number of elements in this Slice.
	Less(i, j int) bool                                // Less returns true if the element indexed by i is less than the element indexed by j.
	Nil() bool                                         // Nil tests if this Slice is nil.
	Map(mod func(O) O) ISlice                          // Map creates a new slice with the modified elements from the lambda.
	O() interface{}                                    // O returns the underlying data structure as is.
	Pair() (first, second *Object)                     // Pair simply returns the first and second Slice elements as Objects.
	Pop() (elem *Object)                               // Pop modifies this Slice to remove the last element and returns the removed element as an Object.
	PopN(n int) (new ISlice)                           // PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
	Prepend(elem interface{}) ISlice                   // Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
	RefSlice() bool                                    // RefSlice returns true if the underlying implementation is a RefSlice
	Reverse() (new ISlice)                             // Reverse returns a new Slice with the order of the elements reversed.
	ReverseM() ISlice                                  // ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
	S() (slice *StringSlice)                           // S is an alias to ToStringSlice
	Select(sel func(O) bool) (new ISlice)              // Select creates a new slice with the elements that match the lambda selector.
	Set(i int, elems interface{}) ISlice               // Set the element(s) at the given index location to the given element(s). Allows for negative notation.
	SetE(i int, elems interface{}) (ISlice, error)     // SetE the element(s) at the given index location to the given element(s). Allows for negative notation.
	Shift() (elem *Object)                             // Shift modifies this Slice to remove the first element and returns the removed element as an Object.
	ShiftN(n int) (new ISlice)                         // ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
	Single() bool                                      // Single reports true if there is only one element in this Slice.
	Slice(indices ...int) ISlice                       // Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
	Sort() (new ISlice)                                // Sort returns a new Slice with sorted elements.
	SortM() ISlice                                     // SortM modifies this Slice sorting the elements and returns a reference to this Slice.
	SortReverse() (new ISlice)                         // SortReverse returns a new Slice sorting the elements in reverse.
	SortReverseM() ISlice                              // SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
	String() string                                    // String returns a string representation of this Slice, implements the Stringer interface
	Swap(i, j int)                                     // Swap modifies this Slice swapping the indicated elements.
	Take(indices ...int) (new ISlice)                  // Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
	TakeAt(i int) (elem *Object)                       // TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
	TakeW(sel func(O) bool) (new ISlice)               // TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
	ToInts() (slice []int)                             // ToInts converts the given slice into a native []int type
	ToIntSlice() (slice *IntSlice)                     // ToIntSlice converts the given slice into a *IntSlice
	ToInterSlice() (slice []interface{})               // ToInterSlice converts the given slice to a generic []interface{} slice
	ToStrs() (slice []string)                          // ToStrs converts the underlying slice into a []string slice
	ToStringSlice() (slice *StringSlice)               // ToStringSlice converts the underlying slice into a *StringSlice
	Union(slice interface{}) (new ISlice)              // Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
	UnionM(slice interface{}) ISlice                   // UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
	Uniq() (new ISlice)                                // Uniq returns a new Slice with all non uniq elements removed while preserving element order.
	UniqM() ISlice                                     // UniqM modifies this Slice to remove all non uniq elements while preserving element order.
}

// Slice provides a generic way to work with Slice types. It does this by wrapping Go types
// directly for optimized types thus avoiding reflection processing overhead and making a plethora
// of Slice methods available. Non optimized types will fall back on reflection to generically
// handle the type incurring the full 10x reflection processing overhead.
//
// Optimized: []int, []string, StrSlice
func Slice(obj interface{}) (new ISlice) {
	ref := Reference(obj)
	switch o := ref.(type) {

	// []interface{}
	// ---------------------------------------------------------------------------------------------
	case *[]interface{}, *InterSlice:
		x := ToInterSlice(o)
		if x != nil && len(*x) > 0 {
			item := Reference((*x)[0])
			switch item.(type) {

			// FloatSlice
			// ---------------------------------------------------------------------------------------------
			case *float32, *float64:
				new = ToFloatSlice(*x)

			// IntSlice
			// ---------------------------------------------------------------------------------------------
			case *int, *int8, *int16, *int64, *uint, *uint16, *uint32, *uint64:
				new = ToIntSlice(*x)

			// StrSlice
			// ---------------------------------------------------------------------------------------------
			case *string, *[]byte, *[]rune, *Str:
				new = ToStringSlice(*x)

			// MapSlice
			//----------------------------------------------------------------------------------------------
			case *map[string]interface{}, *map[string]string:
				new = ToMapSlice(*x)

			// RefSlice
			// ---------------------------------------------------------------------------------------------
			default:
				new = NewRefSlice(obj)
			}
		} else {
			new = x
			return
		}

	// FloatSlice
	// ---------------------------------------------------------------------------------------------
	case *float32, *float64, *[]float32, *[]float64, *[]*float32, *[]*float64:
		new = ToFloatSlice(o)

	// IntSlice
	// ---------------------------------------------------------------------------------------------
	case *int, *int8, *int16, *int64, *uint, *uint8, *uint16, *uint32, *uint64,
		*[]int, *[]int8, *[]int16, *[]int64, *[]uint, *[]uint16, *[]uint32, *[]uint64,
		*[]*int, *[]*int8, *[]*int16, *[]*int64, *[]*uint, *[]*uint16, *[]*uint32, *[]*uint64:
		new = ToIntSlice(o)

	// StrSlice
	// ---------------------------------------------------------------------------------------------
	case *string, *Str, *[]string, *[][]byte, *[][]rune, *[]Str,
		*[]*string, *[]*[]byte, *[]*[]rune, *[]*Str:
		new = ToStringSlice(o)

	// Str
	// ---------------------------------------------------------------------------------------------
	case *Char, *rune, *[]Char, *[]rune, *[]byte,
		*[]*Char, *[]*rune, *[]*byte:
		new = ToStr(o)

	// MapSlice
	//----------------------------------------------------------------------------------------------
	case *MapSlice, *map[string]interface{}, *map[string]string, *[]map[string]interface{}, *[]map[string]string:
		new = ToMapSlice(o)

	// RefSlice
	// ---------------------------------------------------------------------------------------------
	default:
		new = NewRefSlice(obj)
	}
	return
}

// NewSliceV creates a new Slice encapsulating the given variadic elements in a new Slice of
// that type using type assertion for optimized types. Non optimized types will fall back
// on reflection to generically handle the type incurring the full 10x reflection processing
// overhead. In the case where nothing is given a new *RefSlice will be returned.
//
// Optimized: []int, []string, Str
func NewSliceV(elems ...interface{}) (new ISlice) {
	if len(elems) == 0 {
		new = NewRefSliceV(elems...)
	} else {
		switch Reference(elems[0]).(type) {

		// FloatSlice
		// ---------------------------------------------------------------------------------------------
		case *float32, *float64, *[]float32, *[]float64, *[]*float32, *[]*float64:
			new, _ = ToFloatSliceE(elems)

		// IntSlice
		// -----------------------------------------------------------------------------------------
		case *int, *int8, *int16, *int64, *uint, *uint16, *uint32, *uint64:
			new, _ = ToIntSliceE(elems)

		// StringSlice
		// -----------------------------------------------------------------------------------------
		case *string, *[]byte, *[]rune, *Str:
			new, _ = ToStringSliceE(elems)

		// Str
		// -----------------------------------------------------------------------------------------
		case *Char, *rune, *byte:
			new = ToStr(elems)

		// RefSlice
		// -----------------------------------------------------------------------------------------
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
