package n

import (
	"sort"
	"strings"

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
		p.Drop()
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
	if p == nil || len(*p) == 0 {
		other = NewIntSliceV()
		return
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		other = NewIntSliceV()
		return
	}

	x := make([]int, j-i, j-i)
	copy(x, (*p)[i:j])
	other = NewIntSlice(x)
	return
}

// Drop deletes a range of elements and returns the rest of the elements in the slice.
// Expects nothing, in which case everything is dropped, or two indices i and j, in which case
// positive and negative notation is supported and uses an inclusive behavior such that
// DropAt(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices
// will be moved within bounds.
func (p *IntSlice) Drop(indices ...int) Slice {
	if p == nil || len(*p) == 0 {
		return p
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return p
	}

	// Execute
	n := j - i
	if i+n < len(*p) {
		*p = append((*p)[:i], (*p)[i+n:]...)
	} else {
		*p = (*p)[:i]
	}
	return p
}

// DropAt deletes the element at the given index location. Allows for negative notation.
// Returns the rest of the elements in the slice for chaining.
func (p *IntSlice) DropAt(i int) Slice {
	return p.Drop(i, i)
}

// DropFirst deletes the first element and returns the rest of the elements in the slice.
func (p *IntSlice) DropFirst() Slice {
	return p.Drop(0, 0)
}

// DropFirstN deletes first n elements and returns the rest of the elements in the slice.
func (p *IntSlice) DropFirstN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(0, abs(n)-1)
}

// DropLast deletes last element returns the rest of the elements in the slice.
func (p *IntSlice) DropLast() Slice {
	return p.Drop(-1, -1)
}

// DropLastN deletes last n elements and returns the rest of the elements in the slice.
func (p *IntSlice) DropLastN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(absNeg(n), -1)
}

// DropWhere deletes the elements where the lambda returns true. Returns the Slice for chaining.
// The slice is updated instantly when lambda expression is evaluated not after DropWhere is called.
func (p *IntSlice) DropWhere(sel func(O) bool) Slice {
	if p == nil || len(*p) == 0 {
		return p
	}

	l := len(*p)
	for i := 0; i < l; i++ {
		if sel((*p)[i]) {
			p.DropAt(i)
			l--
			i--
		}
	}
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
	return p.Slice(0, abs(n)-1)
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

// Join converts each element into a string then joins them together using the given separator or comma
// if the separator is not given.
func (p *IntSlice) Join(separator ...string) (str *Object) {
	if p == nil || len(*p) == 0 {
		str = &Object{""}
		return
	}
	sep := ","
	if len(separator) > 0 {
		sep = separator[0]
	}

	var builder strings.Builder
	for i := 0; i < len(*p); i++ {
		builder.WriteString((&Object{(*p)[i]}).ToString())
		if i+1 < len(*p) {
			builder.WriteString(sep)
		}
	}
	str = &Object{builder.String()}
	return
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
	return p.Slice(absNeg(n), -1)
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

// Map projects the slice into a new form by executing the lambda against all elements.
func (p *IntSlice) Map(sel func(O) O) (other Slice) {
	return p
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

// Reverse reverses the order of the elements in the slice and returns a reference for chaining.
func (p *IntSlice) Reverse() Slice {
	if p == nil || len(*p) == 0 {
		return p
	}
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		p.Swap(i, j)
	}
	return p
}

// Select creates a new slice with the elements that match the lambda expression.
func (p *IntSlice) Select(sel func(O) bool) (other Slice) {
	new := NewIntSliceV()
	if p == nil || len(*p) == 0 {
		return new
	}

	for i := 0; i < len(*p); i++ {
		if sel((*p)[i]) {
			*new = append(*new, (*p)[i])
		}
	}
	return new
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
// Expects nothing, in which case everything is included, or two indices i and j, in which case
// an inclusive behavior is used such that Slice(0, -1) includes index -1 as opposed to Go's
// exclusive behavior. Out of bounds indices will be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. NewIntSliceV(1,2,3).Slice(0, -1) == [1,2,3] && NewIntSliceV(1,2,3).Slice(1,2) == [2,3]
func (p *IntSlice) Slice(indices ...int) Slice {
	if p == nil || len(*p) == 0 {
		return NewIntSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewIntSliceV()
	}

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

// SortReverse sorts the underlying slice in reverse and return a pointer for chaining.
func (p *IntSlice) SortReverse() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(sort.Reverse(p))
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

// Take deletes a range of elements and returns them as a new slice.
// Expects nothing, in which case everything is taken, or two indices i and j, in which case
// positive and negative notation is supported and uses an inclusive behavior such that
// DropAt(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices
// will be moved within bounds.
func (p *IntSlice) Take(indices ...int) (other Slice) {
	other = p.Copy(indices...)
	p.Drop(indices...)
	return
}

// TakeAt deletes the elemement at the given index location and returns it as an Object.
// Allows for negative notation.
func (p *IntSlice) TakeAt(i int) (elem *Object) {
	elem = p.At(i)
	p.DropAt(i)
	return
}

// TakeFirst deletes the first element and returns it as an Object.
func (p *IntSlice) TakeFirst() (elem *Object) {
	elem = p.First()
	p.DropFirst()
	return
}

// TakeFirstN deletes the first n elements and returns them as a new slice.
func (p *IntSlice) TakeFirstN(n int) (other Slice) {
	if n == 0 {
		return NewIntSliceV()
	}
	other = p.Copy(0, abs(n)-1)
	p.DropFirstN(n)
	return
}

// TakeLast deletes the last element and returns it as an Object.
func (p *IntSlice) TakeLast() (elem *Object) {
	elem = p.Last()
	p.DropLast()
	return
}

// TakeLastN deletes the last n elements and returns them as a new slice.
func (p *IntSlice) TakeLastN(n int) (other Slice) {
	if n == 0 {
		return NewIntSliceV()
	}
	other = p.Copy(absNeg(n), -1)
	p.DropLastN(n)
	return
}
