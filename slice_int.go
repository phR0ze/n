package n

import (
	"sort"

	"github.com/pkg/errors"
)

// IntSlice implements the Slice interface providing a generic way to work with slice types
// including convenience methods on par with other rapid development languages.
type IntSlice []int

// NewIntSlice creates a new IntSlice
func NewIntSlice(slice []int) *IntSlice {
	new := IntSlice(slice)
	return &new
}

// NewIntSliceV creates a new IntSlice from the given variadic elements. Always returns
// at least an empty slice.
func NewIntSliceV(elems ...int) *IntSlice {
	var new IntSlice
	if len(elems) == 0 {
		new = IntSlice([]int{})
	} else {
		new = IntSlice(elems)
	}
	return &new
}

// Any tests if the slice is not empty or optionally if it contains
// any of the given Variadic elements. Incompatible types will return false.
func (p *IntSlice) Any(elems ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}

	// Elements and not looking for anything
	if len(elems) == 0 {
		return true
	}

	// Looking for something specific returns false if incompatible type
	for i := range elems {
		if x, ok := elems[i].(int); ok {
			for j := range *p {
				if (*p)[j] == x {
					return true
				}
			}
		}
	}
	return false
}

// AnyS tests if the slice contains any of the other slice's elements.
// Incompatible types will return false.
func (p *IntSlice) AnyS(other interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	if elems, ok := other.([]int); ok {
		for i := range elems {
			for j := range *p {
				if (*p)[j] == elems[i] {
					return true
				}
			}
		}
	}
	return false
}

// Append an element to the end of the Slice and returns the Slice for chaining
func (p *IntSlice) Append(elem interface{}) Slice {
	if p == nil {
		p = NewIntSliceV()
	}
	if x, ok := elem.(int); ok {
		*p = append(*p, x)
	}
	return p
}

// AppendS appends the other slice using variadic expansion and returns Slice for chaining
func (p *IntSlice) AppendS(other interface{}) Slice {
	if p == nil {
		p = NewIntSliceV()
	}
	if x, ok := other.([]int); ok {
		*p = append(*p, x...)
	}
	return p
}

// AppendV appends the variadic elements to the end of the Slice and returns the Slice for chaining
func (p *IntSlice) AppendV(elems ...interface{}) Slice {
	if p == nil {
		p = NewIntSliceV()
	}
	for _, elem := range elems {
		p.Append(elem)
	}
	return p
}

// At returns the element at the given index location. Allows for negative notation.
func (p *IntSlice) At(i int) (elem *Object) {
	elem = &Object{}
	if p == nil {
		return
	}
	if i = absIndex(len(*p), i); i == -1 {
		return
	}
	elem.o = (*p)[i]
	return
}

// Clear the underlying slice, returns Slice for chaining.
func (p *IntSlice) Clear() Slice {
	if p == nil {
		p = NewIntSliceV()
	} else {
		*p = *NewIntSliceV()
	}
	return p
}

// Copy performs a deep copy such that modifications to the copy will not affect
// the original. Expects nothing, in which case everything is copied, or two
// indices i and j, in which case positive and negative notation is supported and
// uses an inclusive behavior such that Slice(0, -1) includes index -1 as opposed
// to Go's exclusive  behavior. Out of bounds indices will be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
func (p *IntSlice) Copy(indices ...int) (other Slice) {
	if p == nil || len(*p) == 0 || len(indices) == 1 {
		other = NewIntSliceV()
		return
	}

	// Get indices
	i, j := 0, len(*p)-1
	if len(indices) == 2 {
		i = indices[0]
		j = indices[1]
	}

	// Convert to postive notation
	if i < 0 {
		i = len(*p) + i
	}
	if j < 0 {
		j = len(*p) + j
	}

	// Start can't be past end else nothing to get
	if i > j {
		other = NewIntSliceV()
		return
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= len(*p) {
		j = len(*p) - 1
	}

	// Go has an exclusive behavior by default and we want inclusive
	// so offsetting the end by one
	j++
	x := make([]int, j-i, j-i)
	copy(x, (*p)[i:j])
	other = NewIntSlice(x)
	return
}

// Drop deletes the element at the given index location. Allows for negative notation.
// Returns the rest of the elements in the slice for chaining.
func (p *IntSlice) Drop(i int) Slice {
	if p == nil {
		return p
	}
	if i = absIndex(len(*p), i); i == -1 {
		return p
	}

	if i+1 < len(*p) {
		*p = append((*p)[:i], (*p)[i+1:]...)
	} else {
		*p = (*p)[:i]
	}
	return p
}

// DropFirst deletes the first element and returns the rest of the elements in the slice.
func (p *IntSlice) DropFirst() Slice {
	return p.DropFirstN(1)
}

// DropFirstN deletes first n elements and returns the rest of the elements in the slice.
func (p *IntSlice) DropFirstN(n int) Slice {
	if p == nil || n == 0 {
		return p
	}
	if len(*p) <= n {
		*p = *NewIntSliceV()
		return p
	}
	*p = (*p)[n:]
	return p
}

// DropLast deletes last element returns the rest of the elements in the slice.
func (p *IntSlice) DropLast() Slice {
	return p.DropLastN(1)
}

// DropLastN deletes last n elements and returns the rest of the elements in the slice.
func (p *IntSlice) DropLastN(n int) Slice {
	if p == nil || n == 0 {
		return p
	}
	if len(*p) <= n {
		*p = *NewIntSliceV()
		return p
	}
	*p = (*p)[:len(*p)-n]
	return p
}

// Each calls the given function once for each element in the slice, passing that element in
// as a parameter. Returns a reference to the slice
func (p *IntSlice) Each(action func(O)) Slice {
	if p == nil {
		return p
	}
	for i := 0; i < len(*p); i++ {
		action((*p)[i])
	}
	return p
}

// EachE calls the given function once for each element in the slice, passing that element in
// as a parameter. Returns a reference to the slice and any error from the user function.
func (p *IntSlice) EachE(action func(O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := 0; i < len(*p); i++ {
		if err = action((*p)[i]); err != nil {
			return p, err
		}
	}
	return p, err
}

// Empty tests if the slice is empty.
func (p *IntSlice) Empty() bool {
	if p == nil || len(*p) == 0 {
		return true
	}
	return false
}

// First returns the first element in the slice as Object which will be Object.Nil true if
// there are no elements in the slice.
func (p *IntSlice) First() (elem *Object) {
	elem = p.At(0)
	return
}

// FirstN returns the first n elements in the slice as a Slice. Best effort is used such
// that as many as can be will be returned up until the request is satisfied.
func (p *IntSlice) FirstN(n int) Slice {
	j := n - 1
	if n < 0 {
		j = (n * -1) - 1
	}
	return p.Slice(0, j)
}

// Insert the given element before the element with the given index. Negative indices count
// backwards from the end of the slice, where -1 is the last element. If a negative index
// is used, the given element will be inserted after that element, so using an index of -1
// will insert the element at the end of the slice. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *IntSlice) Insert(i int, elem interface{}) Slice {
	if p == nil || len(*p) == 0 {
		return p.Append(elem)
	}
	j := i
	if j = absIndex(len(*p), j); j == -1 {
		return p
	}
	if i < 0 {
		j++
	}

	// Insert the item before j if pos and after j if neg
	if x, ok := elem.(int); ok {
		if j == 0 {
			*p = append([]int{x}, (*p)...)
		} else if j < len(*p) {
			*p = append(*p, x)
			copy((*p)[j+1:], (*p)[j:])
			(*p)[j] = x
		} else {
			*p = append(*p, x)
		}
	}
	return p
}

// Last returns the last element in the slice as Object which will be Object.Nil true if
// there are no elements in the slice.
func (p *IntSlice) Last() (elem *Object) {
	elem = p.At(-1)
	return
}

// LastN returns the last n elements in the slice as a NSlice. Best effort is used such
// that as many as can be will be returned up until the request is satisfied.
func (p *IntSlice) LastN(n int) Slice {
	i := n * -1
	if n < 0 {
		i = n
	}
	return p.Slice(i, -1)
}

// Len returns the number of elements in the slice
func (p *IntSlice) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// Less returns true if the element indexed by i is less than the element indexed by j.
func (p *IntSlice) Less(i, j int) bool {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return false
	}
	return (*p)[i] < (*p)[j]
}

// Nil tests if the slice is nil
func (p *IntSlice) Nil() bool {
	if p == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *IntSlice) O() interface{} {
	return []int(*p)
}

// Pair simply returns the first and second slice elements as Object
func (p *IntSlice) Pair() (first, second *Object) {
	first, second = &Object{}, &Object{}
	if len(*p) > 0 {
		first = p.At(0)
	}
	if len(*p) > 1 {
		second = p.At(1)
	}
	return
}

// Prepend the given element at the begining of the slice.
func (p *IntSlice) Prepend(elem interface{}) Slice {
	return p.Insert(0, elem)
}

// Set the element at the given index location to the given element. Allows for negative notation.
// Returns the slice for chaining and swallows any errors if out of bounds or elem is the wrong type
func (p *IntSlice) Set(i int, elem interface{}) Slice {
	slice, _ := p.SetE(i, elem)
	return slice
}

// SetE the element at the given index location to the given element. Allows for negative notation.
// Returns the slice for chaining and an error if out of bounds or elem is the wrong type
func (p *IntSlice) SetE(i int, elem interface{}) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	if i = absIndex(len(*p), i); i == -1 {
		err = errors.Errorf("slice assignment is out of bounds")
		return p, err
	}

	if x, ok := elem.(int); ok {
		(*p)[i] = x
	} else {
		err = errors.Errorf("can't set type '%T' in '%T'", elem, p)
	}
	return p, err
}

// Single simply reports true if there is only one element in the slice
func (p *IntSlice) Single() bool {
	return len(*p) == 1
}

// Slice provides a Ruby like slice function for Slice allowing for positive and negative notation.
// Slice uses an inclusive behavior such that Slice(0, -1) includes index -1 as opposed to Go's exclusive
// behavior. Out of bounds indices will be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. NewIntSliceV(1,2,3).Slice(0, -1) == [1,2,3] && NewIntSliceV(1,2,3).Slice(1,2) == [2,3]
func (p *IntSlice) Slice(i, j int) Slice {
	if p == nil || len(*p) == 0 {
		return NewIntSliceV()
	}

	// Convert to postive notation
	if i < 0 {
		i = len(*p) + i
	}
	if j < 0 {
		j = len(*p) + j
	}

	// Start can't be past end else nothing to get
	if i > j {
		return NewIntSliceV()
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= len(*p) {
		j = len(*p) - 1
	}

	// Go has an exclusive behavior by default and we want inclusive
	// so offsetting the end by one
	j++

	return NewIntSlice((*p)[i:j])
}

// Sort the underlying slice and return a pointer for chaining.
func (p *IntSlice) Sort() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(p)
	return p
}

// Swap elements in the underlying slice. Implements the sort.Interface.
// Takes advantage of underlying slice's sort.Interface implementations if they exist.
func (p *IntSlice) Swap(i, j int) {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return
	}
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}
