package n

import (
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

// Len returns the number of elements in the slice
func (p *IntSlice) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
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
