package n

import (
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// StringSlice implements the Slice interface providing a generic way to work with slice types
// including convenience methods on par with other rapid development languages.
type StringSlice []string

// NewStringSlice creates a new *StringSlice
func NewStringSlice(slice []string) *StringSlice {
	new := StringSlice(slice)
	return &new
}

// NewStringSliceV creates a new *StringSlice from the given variadic elements. Always returns
// at least a reference to an empty StringSlice.
func NewStringSliceV(elems ...string) *StringSlice {
	var new StringSlice
	if len(elems) == 0 {
		new = StringSlice([]string{})
	} else {
		new = StringSlice(elems)
	}
	return &new
}

// Any tests if this Slice is not empty or optionally if it contains
// any of the given variadic elements. Incompatible types will return false.
func (p *StringSlice) Any(elems ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}

	// Not looking for anything
	if len(elems) == 0 {
		return true
	}

	// Looking for something specific returns false if incompatible type
	for i := range elems {
		if x, ok := elems[i].(string); ok {
			for j := range *p {
				if (*p)[j] == x {
					return true
				}
			}
		}
	}
	return false
}

// AnyS tests if this Slice contains any of the given Slice's elements.
// Incompatible types will return false.
// Supports StringSlice, *StringSlice, []string or *[]string
func (p *StringSlice) AnyS(slice interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	var elems []string
	switch x := slice.(type) {
	case []string:
		elems = x
	case *[]string:
		if x != nil {
			elems = *x
		}
	case StringSlice:
		elems = x
	case *StringSlice:
		if x != nil {
			elems = (*x)
		}
	}
	for i := range elems {
		for j := range *p {
			if (*p)[j] == elems[i] {
				return true
			}
		}
	}
	return false
}

// AnyW tests if this Slice contains any that match the lambda selector.
func (p *StringSlice) AnyW(sel func(O) bool) bool {
	return p.CountW(sel) != 0
}

// Append an element to the end of this Slice and returns a reference to this Slice.
func (p *StringSlice) Append(elem interface{}) Slice {
	if p == nil {
		p = NewStringSliceV()
	}
	if x, ok := elem.(string); ok {
		*p = append(*p, x)
	}
	return p
}

// AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
func (p *StringSlice) AppendV(elems ...interface{}) Slice {
	if p == nil {
		p = NewStringSliceV()
	}
	for _, elem := range elems {
		p.Append(elem)
	}
	return p
}

// At returns the element at the given index location. Allows for negative notation.
func (p *StringSlice) At(i int) (elem *Object) {
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

// Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
func (p *StringSlice) Clear() Slice {
	if p == nil {
		p = NewStringSliceV()
	} else {
		p.Drop()
	}
	return p
}

// Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
// Supports StringSlice, *StringSlice, []string or *[]string
func (p *StringSlice) Concat(slice interface{}) (new Slice) {
	return p.Copy().ConcatM(slice)
}

// ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
// Supports StringSlice, *StringSlice, []string or *[]string
func (p *StringSlice) ConcatM(slice interface{}) Slice {
	if p == nil {
		p = NewStringSliceV()
	}
	switch x := slice.(type) {
	case []string:
		*p = append(*p, x...)
	case *[]string:
		if x != nil {
			*p = append(*p, (*x)...)
		}
	case StringSlice:
		*p = append(*p, x...)
	case *StringSlice:
		if x != nil {
			*p = append(*p, (*x)...)
		}
	}
	return p
}

// Copy returns a new Slice with the indicated range of elements copied from this Slice.
// Expects nothing, in which case everything is copied, or two indices i and j, in which
// case positive and negative notation is supported and uses an inclusive behavior such
// that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of
// bounds indices will be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
func (p *StringSlice) Copy(indices ...int) (new Slice) {
	if p == nil || len(*p) == 0 {
		return NewStringSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewStringSliceV()
	}

	// Copy elements over to new Slice
	x := make([]string, j-i, j-i)
	copy(x, (*p)[i:j])
	return NewStringSlice(x)
}

// Count the number of elements in this Slice equal to the given element.
func (p *StringSlice) Count(elem interface{}) (cnt int) {
	if y, ok := elem.(string); ok {
		cnt = p.CountW(func(x O) bool { return ExB(x.(string) == y) })
	}
	return
}

// CountW counts the number of elements in this Slice that match the lambda selector.
func (p *StringSlice) CountW(sel func(O) bool) (cnt int) {
	if p == nil || len(*p) == 0 {
		return
	}
	for i := 0; i < len(*p); i++ {
		if sel((*p)[i]) {
			cnt++
		}
	}
	return
}

// Drop modifies this Slice to delete the indicated range of elements and returns a referece to this Slice.
// Expects nothing, in which case everything is dropped, or two indices i and j, in which case positive and
// negative notation is supported and uses an inclusive behavior such that DropAt(0, -1) includes index -1
// as opposed to Go's exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *StringSlice) Drop(indices ...int) Slice {
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

// DropAt modifies this Slice to delete the element at the given index location. Allows for negative notation.
// Returns a reference to this Slice.
func (p *StringSlice) DropAt(i int) Slice {
	return p.Drop(i, i)
}

// DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
func (p *StringSlice) DropFirst() Slice {
	return p.Drop(0, 0)
}

// DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
func (p *StringSlice) DropFirstN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(0, abs(n)-1)
}

// DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
func (p *StringSlice) DropLast() Slice {
	return p.Drop(-1, -1)
}

// DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
func (p *StringSlice) DropLastN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(absNeg(n), -1)
}

// DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// The slice is updated instantly when lambda expression is evaluated not after DropW completes.
func (p *StringSlice) DropW(sel func(O) bool) Slice {
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

// Each calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *StringSlice) Each(action func(O)) Slice {
	if p == nil {
		return p
	}
	for i := 0; i < len(*p); i++ {
		action((*p)[i])
	}
	return p
}

// EachE calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *StringSlice) EachE(action func(O) error) (Slice, error) {
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

// EachI calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice
func (p *StringSlice) EachI(action func(int, O)) Slice {
	if p == nil {
		return p
	}
	for i := 0; i < len(*p); i++ {
		action(i, (*p)[i])
	}
	return p
}

// EachIE calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *StringSlice) EachIE(action func(int, O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := 0; i < len(*p); i++ {
		if err = action(i, (*p)[i]); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *StringSlice) EachR(action func(O)) Slice {
	if p == nil {
		return p
	}
	for i := len(*p) - 1; i >= 0; i-- {
		action((*p)[i])
	}
	return p
}

// EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *StringSlice) EachRE(action func(O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := len(*p) - 1; i >= 0; i-- {
		if err = action((*p)[i]); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *StringSlice) EachRI(action func(int, O)) Slice {
	if p == nil {
		return p
	}
	for i := len(*p) - 1; i >= 0; i-- {
		action(i, (*p)[i])
	}
	return p
}

// EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *StringSlice) EachRIE(action func(int, O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := len(*p) - 1; i >= 0; i-- {
		if err = action(i, (*p)[i]); err != nil {
			return p, err
		}
	}
	return p, err
}

// Empty tests if this Slice is empty.
func (p *StringSlice) Empty() bool {
	if p == nil || len(*p) == 0 {
		return true
	}
	return false
}

// First returns the first element in this Slice as Object.
// Object.Nil() == true will be returned when there are no elements in the slice.
func (p *StringSlice) First() (elem *Object) {
	return p.At(0)
}

// FirstN returns the first n elements in this slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *StringSlice) FirstN(n int) Slice {
	if n == 0 {
		return NewStringSliceV()
	}
	return p.Slice(0, abs(n)-1)
}

// Generic returns true if the underlying implementation is a RefSlice
func (p *StringSlice) Generic() bool {
	return false
}

// Index returns the index of the first element in this Slice where element == elem
// Returns a -1 if the element was not not found.
func (p *StringSlice) Index(elem interface{}) (loc int) {
	loc = -1
	if p == nil || len(*p) == 0 {
		return
	}
	for i := 0; i < len(*p); i++ {
		if elem == (*p)[i] {
			return i
		}
	}
	return
}

// Insert modifies this Slice to insert the given element before the element with the given index.
// Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// negative index is used, the given element will be inserted after that element, so using an index
// of -1 will insert the element at the end of the slice. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *StringSlice) Insert(i int, elem interface{}) Slice {
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
	if x, ok := elem.(string); ok {
		if j == 0 {
			*p = append([]string{x}, (*p)...)
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

// Join converts each element into a string then joins them together using the given separator or comma by default.
func (p *StringSlice) Join(separator ...string) (str *Object) {
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
		builder.WriteString(Obj((*p)[i]).ToString())
		if i+1 < len(*p) {
			builder.WriteString(sep)
		}
	}
	str = &Object{builder.String()}
	return
}

// Last returns the last element in this Slice as an Object.
// Object.Nil() == true will be returned if there are no elements in the slice.
func (p *StringSlice) Last() (elem *Object) {
	return p.At(-1)
}

// LastN returns the last n elements in this Slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *StringSlice) LastN(n int) Slice {
	if n == 0 {
		return NewStringSliceV()
	}
	return p.Slice(absNeg(n), -1)
}

// Len returns the number of elements in this Slice
func (p *StringSlice) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// Less returns true if the element indexed by i is less than the element indexed by j.
func (p *StringSlice) Less(i, j int) bool {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return false
	}
	return (*p)[i] < (*p)[j]
}

// Nil tests if this Slice is nil
func (p *StringSlice) Nil() bool {
	if p == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *StringSlice) O() interface{} {
	return []string(*p)
}

// Pair simply returns the first and second Slice elements as Objects
func (p *StringSlice) Pair() (first, second *Object) {
	first, second = &Object{}, &Object{}
	if p == nil {
		return
	}
	if len(*p) > 0 {
		first = p.At(0)
	}
	if len(*p) > 1 {
		second = p.At(1)
	}
	return
}

// Pop modifies this Slice to remove the last element and returns the removed element as an Object.
func (p *StringSlice) Pop() (elem *Object) {
	elem = p.Last()
	p.DropLast()
	return
}

// PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
func (p *StringSlice) PopN(n int) (new Slice) {
	if n == 0 {
		return NewStringSliceV()
	}
	new = p.Copy(absNeg(n), -1)
	p.DropLastN(n)
	return
}

// Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
func (p *StringSlice) Prepend(elem interface{}) Slice {
	return p.Insert(0, elem)
}

// Reverse returns a new Slice with the order of the elements reversed.
func (p *StringSlice) Reverse() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().ReverseM()
}

// ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
func (p *StringSlice) ReverseM() Slice {
	if p == nil || len(*p) == 0 {
		return p
	}
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		p.Swap(i, j)
	}
	return p
}

// Select creates a new slice with the elements that match the lambda selector.
func (p *StringSlice) Select(sel func(O) bool) (new Slice) {
	slice := NewStringSliceV()
	if p == nil || len(*p) == 0 {
		return slice
	}
	for i := 0; i < len(*p); i++ {
		if sel((*p)[i]) {
			*slice = append(*slice, (*p)[i])
		}
	}
	return slice
}

// Set the element at the given index location to the given element. Allows for negative notation.
// Returns a reference to this Slice and swallows any errors.
func (p *StringSlice) Set(i int, elem interface{}) Slice {
	slice, _ := p.SetE(i, elem)
	return slice
}

// SetE the element at the given index location to the given element. Allows for negative notation.
// Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
func (p *StringSlice) SetE(i int, elem interface{}) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	if i = absIndex(len(*p), i); i == -1 {
		err = errors.Errorf("slice assignment is out of bounds")
		return p, err
	}

	if x, ok := elem.(string); ok {
		(*p)[i] = x
	} else {
		err = errors.Errorf("can't set type '%T' in '%T'", elem, p)
	}
	return p, err
}

// Shift modifies this Slice to remove the first element and returns the removed element as an Object.
func (p *StringSlice) Shift() (elem *Object) {
	elem = p.First()
	p.DropFirst()
	return
}

// ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
func (p *StringSlice) ShiftN(n int) (new Slice) {
	if n == 0 {
		return NewStringSliceV()
	}
	new = p.Copy(0, abs(n)-1)
	p.DropFirstN(n)
	return
}

// Single reports true if there is only one element in this Slice.
func (p *StringSlice) Single() bool {
	return len(*p) == 1
}

// Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
// Expects nothing, in which case everything is included, or two indices i and j, in which case an inclusive behavior
// is used such that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices will
// be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. NewStringSliceV(1,2,3).Slice(0, -1) == [1,2,3] && NewStringSliceV(1,2,3).Slice(1,2) == [2,3]
func (p *StringSlice) Slice(indices ...int) Slice {
	if p == nil || len(*p) == 0 {
		return NewStringSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewStringSliceV()
	}

	return NewStringSlice((*p)[i:j])
}

// Sort returns a new Slice with sorted elements.
func (p *StringSlice) Sort() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortM()
}

// SortM modifies this Slice sorting the elements and returns a reference to this Slice.
func (p *StringSlice) SortM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(p)
	return p
}

// SortReverse returns a new Slice sorting the elements in reverse.
func (p *StringSlice) SortReverse() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortReverseM()
}

// SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
func (p *StringSlice) SortReverseM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(sort.Reverse(p))
	return p
}

// String returns a string representation of this Slice, implements the Stringer interface
func (p *StringSlice) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	if p != nil {
		for i := 0; i < len(*p); i++ {
			builder.WriteString((*p)[i])
			if i+1 < len(*p) {
				builder.WriteString(" ")
			}
		}
	}
	builder.WriteString("]")
	return builder.String()
}

// Swap modifies this Slice swapping the indicated elements.
func (p *StringSlice) Swap(i, j int) {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return
	}
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

// Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *StringSlice) Take(indices ...int) (new Slice) {
	new = p.Copy(indices...)
	p.Drop(indices...)
	return
}

// TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
// Allows for negative notation.
func (p *StringSlice) TakeAt(i int) (elem *Object) {
	elem = p.At(i)
	p.DropAt(i)
	return
}

// TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
func (p *StringSlice) TakeW(sel func(O) bool) (new Slice) {
	slice := NewStringSliceV()
	if p == nil || len(*p) == 0 {
		return slice
	}
	l := len(*p)
	for i := 0; i < l; i++ {
		if sel((*p)[i]) {
			*slice = append(*slice, (*p)[i])
			p.DropAt(i)
			l--
			i--
		}
	}
	return slice
}

// Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports StringSlice, *StringSlice, []string or *[]string
func (p *StringSlice) Union(slice interface{}) (new Slice) {
	return p.Copy().UnionM(slice)
}

// UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports StringSlice, *StringSlice, []string or *[]string
func (p *StringSlice) UnionM(slice interface{}) Slice {
	return p.ConcatM(slice).UniqM()
}

// Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
func (p *StringSlice) Uniq() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	m := NewStringMapBool()
	slice := NewStringSliceV()
	for i := 0; i < len(*p); i++ {
		if ok := m.Set((*p)[i], true); ok {
			slice.Append((*p)[i])
		}
	}
	return slice
}

// UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
func (p *StringSlice) UniqM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	m := NewStringMapBool()
	l := len(*p)
	for i := 0; i < l; i++ {
		if ok := m.Set((*p)[i], true); !ok {
			p.DropAt(i)
			l--
			i--
		}
	}
	return p
}
