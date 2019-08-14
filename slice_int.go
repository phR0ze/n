package n

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// IntSlice implements the Slice interface providing a generic way to work with slice types
// including convenience methods on par with rapid development languages.
type IntSlice []int

// NewIntSlice creates a new *IntSlice
func NewIntSlice(slice interface{}) *IntSlice {
	return ToIntSlice(slice)
}

// NewIntSliceV creates a new *IntSlice from the given variadic elements. Always returns
// at least a reference to an empty IntSlice.
func NewIntSliceV(elems ...interface{}) *IntSlice {
	return ToIntSlice(elems)
}

// A is an alias to String for brevity
func (p *IntSlice) A() string {
	return p.String()
}

// Any tests if this Slice is not empty or optionally if it contains
// any of the given variadic elements. Incompatible types will return false.
// Supports all possible int conversions
func (p *IntSlice) Any(elems ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}

	// Not looking for anything
	if len(elems) == 0 {
		return true
	}

	// Looking for something specific returns false if incompatible type
	return p.AnyS(elems)
}

// AnyS tests if this Slice contains any of the given Slice's elements.
// Incompatible types will return false.
// Supports IntSlice, *IntSlice, []int or *[]int
func (p *IntSlice) AnyS(slice interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	if elems, err := ToIntSliceE(slice); err == nil {
		for i := range *elems {
			for j := range *p {
				if (*p)[j] == (*elems)[i] {
					return true
				}
			}
		}
	}
	return false
}

// AnyW tests if this Slice contains any that match the lambda selector.
func (p *IntSlice) AnyW(sel func(O) bool) bool {
	return p.CountW(sel) != 0
}

// Append an element to the end of this Slice and returns a reference to this Slice.
func (p *IntSlice) Append(elem interface{}) Slice {
	if p == nil {
		p = NewIntSliceV()
	}
	if x, err := ToIntE(elem); err == nil {
		*p = append(*p, x)
	}
	return p
}

// AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
func (p *IntSlice) AppendV(elems ...interface{}) Slice {
	if p == nil {
		p = NewIntSliceV()
	}
	for _, elem := range elems {
		if x, err := ToIntE(elem); err == nil {
			*p = append(*p, x)
		}
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

// Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
func (p *IntSlice) Clear() Slice {
	if p == nil {
		p = NewIntSliceV()
	} else {
		p.Drop()
	}
	return p
}

// Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
// Supports IntSlice, *IntSlice, []int or *[]int
func (p *IntSlice) Concat(slice interface{}) (new Slice) {
	return p.Copy().ConcatM(slice)
}

// ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
// Supports IntSlice, *IntSlice, []int or *[]int
func (p *IntSlice) ConcatM(slice interface{}) Slice {
	if p == nil {
		p = NewIntSliceV()
	}
	if elems, err := ToIntSliceE(slice); err == nil {
		*p = append(*p, *elems...)
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
func (p *IntSlice) Copy(indices ...int) (new Slice) {
	if p == nil || len(*p) == 0 {
		return NewIntSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewIntSliceV()
	}

	// Copy elements over to new Slice
	x := make([]int, j-i, j-i)
	copy(x, (*p)[i:j])
	return NewIntSlice(x)
}

// Count the number of elements in this Slice equal to the given element.
func (p *IntSlice) Count(elem interface{}) (cnt int) {
	if y, ok := elem.(int); ok {
		cnt = p.CountW(func(x O) bool { return ExB(x.(int) == y) })
	}
	return
}

// CountW counts the number of elements in this Slice that match the lambda selector.
func (p *IntSlice) CountW(sel func(O) bool) (cnt int) {
	if p == nil || len(*p) == 0 {
		return
	}
	for i := range *p {
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

// DropAt modifies this Slice to delete the element at the given index location. Allows for negative notation.
// Returns a reference to this Slice.
func (p *IntSlice) DropAt(i int) Slice {
	return p.Drop(i, i)
}

// DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
func (p *IntSlice) DropFirst() Slice {
	return p.Drop(0, 0)
}

// DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
func (p *IntSlice) DropFirstN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(0, abs(n)-1)
}

// DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
func (p *IntSlice) DropLast() Slice {
	return p.Drop(-1, -1)
}

// DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
func (p *IntSlice) DropLastN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(absNeg(n), -1)
}

// DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// The slice is updated instantly when lambda expression is evaluated not after DropW completes.
func (p *IntSlice) DropW(sel func(O) bool) Slice {
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
func (p *IntSlice) Each(action func(O)) Slice {
	if p == nil {
		return p
	}
	for i := range *p {
		action((*p)[i])
	}
	return p
}

// EachE calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *IntSlice) EachE(action func(O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := range *p {
		if err = action((*p)[i]); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachI calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice
func (p *IntSlice) EachI(action func(int, O)) Slice {
	if p == nil {
		return p
	}
	for i := range *p {
		action(i, (*p)[i])
	}
	return p
}

// EachIE calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *IntSlice) EachIE(action func(int, O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := range *p {
		if err = action(i, (*p)[i]); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *IntSlice) EachR(action func(O)) Slice {
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
func (p *IntSlice) EachRE(action func(O) error) (Slice, error) {
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
func (p *IntSlice) EachRI(action func(int, O)) Slice {
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
func (p *IntSlice) EachRIE(action func(int, O) error) (Slice, error) {
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
func (p *IntSlice) Empty() bool {
	if p == nil || len(*p) == 0 {
		return true
	}
	return false
}

// First returns the first element in this Slice as Object.
// Object.Nil() == true will be returned when there are no elements in the slice.
func (p *IntSlice) First() (elem *Object) {
	return p.At(0)
}

// FirstN returns the first n elements in this slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *IntSlice) FirstN(n int) Slice {
	if n == 0 {
		return NewIntSliceV()
	}
	return p.Slice(0, abs(n)-1)
}

// G returns the underlying data structure as a builtin Go type
func (p *IntSlice) G() []int {
	return p.O().([]int)
}

// Generic returns true if the underlying implementation is a RefSlice
func (p *IntSlice) Generic() bool {
	return false
}

// Index returns the index of the first element in this Slice where element == elem
// Returns a -1 if the element was not not found.
func (p *IntSlice) Index(elem interface{}) (loc int) {
	loc = -1
	if p == nil || len(*p) == 0 {
		return
	}
	if x, err := ToIntE(elem); err == nil {
		for i := range *p {
			if (*p)[i] == x {
				return i
			}
		}
	}
	return
}

// Insert modifies this Slice to insert the given elements before the element(s) with the given index.
// Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// negative index is used, the given element will be inserted after that element, so using an index
// of -1 will insert the element at the end of the slice. If a Slice is given all elements will be
// inserted starting from the beging until the end. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *IntSlice) Insert(i int, obj interface{}) Slice {
	if p == nil || len(*p) == 0 {
		return p.ConcatM(obj)
	}

	// Insert the item before j if pos and after j if neg
	j := i
	if j = absIndex(len(*p), j); j == -1 {
		return p
	}
	if i < 0 {
		j++
	}
	if elems, err := ToIntSliceE(obj); err == nil {
		if j == 0 {
			*p = append(*elems, *p...)
		} else if j < len(*p) {
			*p = append(*p, *elems...)           // ensures enough space exists
			copy((*p)[j+len(*elems):], (*p)[j:]) // shifts right elements drop added
			copy((*p)[j:], *elems)               // set new in locations vacated
		} else {
			*p = append(*p, *elems...)
		}
	}
	return p
}

// Join converts each element into a string then joins them together using the given separator or comma by default.
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
	for i := range *p {
		builder.WriteString(ToString((*p)[i]))
		if i+1 < len(*p) {
			builder.WriteString(sep)
		}
	}
	str = &Object{builder.String()}
	return
}

// Last returns the last element in this Slice as an Object.
// Object.Nil() == true will be returned if there are no elements in the slice.
func (p *IntSlice) Last() (elem *Object) {
	return p.At(-1)
}

// LastN returns the last n elements in this Slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *IntSlice) LastN(n int) Slice {
	if n == 0 {
		return NewIntSliceV()
	}
	return p.Slice(absNeg(n), -1)
}

// Len returns the number of elements in this Slice
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

// Nil tests if this Slice is nil
func (p *IntSlice) Nil() bool {
	if p == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *IntSlice) O() interface{} {
	if p == nil {
		return []int{}
	}
	return []int(*p)
}

// Pair simply returns the first and second Slice elements as Objects
func (p *IntSlice) Pair() (first, second *Object) {
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
func (p *IntSlice) Pop() (elem *Object) {
	elem = p.Last()
	p.DropLast()
	return
}

// PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
func (p *IntSlice) PopN(n int) (new Slice) {
	if n == 0 {
		return NewIntSliceV()
	}
	new = p.Copy(absNeg(n), -1)
	p.DropLastN(n)
	return
}

// Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
func (p *IntSlice) Prepend(elem interface{}) Slice {
	return p.Insert(0, elem)
}

// Reverse returns a new Slice with the order of the elements reversed.
func (p *IntSlice) Reverse() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().ReverseM()
}

// ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
func (p *IntSlice) ReverseM() Slice {
	if p == nil || len(*p) == 0 {
		return p
	}
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		p.Swap(i, j)
	}
	return p
}

// S is an alias to ToStringSlice
func (p *IntSlice) S() (slice *StringSlice) {
	return ToStringSlice(p.O())
}

// SG is an alias to ToStringSliceG
func (p *IntSlice) SG() (slice []string) {
	return ToStringSliceG(p.O())
}

// Select creates a new slice with the elements that match the lambda selector.
func (p *IntSlice) Select(sel func(O) bool) (new Slice) {
	slice := NewIntSliceV()
	if p == nil || len(*p) == 0 {
		return slice
	}
	for i := range *p {
		if sel((*p)[i]) {
			*slice = append(*slice, (*p)[i])
		}
	}
	return slice
}

// Set the element(s) at the given index location to the given element(s). Allows for negative notation.
// Returns a reference to this Slice and swallows any errors.
func (p *IntSlice) Set(i int, elems interface{}) Slice {
	slice, _ := p.SetE(i, elems)
	return slice
}

// SetE the element(s) at the given index location to the given element(s). Allows for negative notation.
// Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
func (p *IntSlice) SetE(i int, elems interface{}) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	if i = absIndex(len(*p), i); i == -1 {
		err = errors.Errorf("slice assignment is out of bounds")
		return p, err
	}

	// Account for length of elems
	if x, err := ToIntSliceE(elems); err == nil {
		if len(*x) > 0 {
			copy((*p)[i:], *x)
		}
	} else {
		err = errors.Wrapf(err, "can't set type '%T' in '%T'", elems, p)
	}
	return p, err
}

// Shift modifies this Slice to remove the first element and returns the removed element as an Object.
func (p *IntSlice) Shift() (elem *Object) {
	elem = p.First()
	p.DropFirst()
	return
}

// ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
func (p *IntSlice) ShiftN(n int) (new Slice) {
	if n == 0 {
		return NewIntSliceV()
	}
	new = p.Copy(0, abs(n)-1)
	p.DropFirstN(n)
	return
}

// Single reports true if there is only one element in this Slice.
func (p *IntSlice) Single() bool {
	return p.Len() == 1
}

// Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
// Expects nothing, in which case everything is included, or two indices i and j, in which case an inclusive behavior
// is used such that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices will
// be moved within bounds.
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

	slice := IntSlice((*p)[i:j])
	return &slice
}

// Sort returns a new Slice with sorted elements.
func (p *IntSlice) Sort() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortM()
}

// SortM modifies this Slice sorting the elements and returns a reference to this Slice.
func (p *IntSlice) SortM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(p)
	return p
}

// SortReverse returns a new Slice sorting the elements in reverse.
func (p *IntSlice) SortReverse() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortReverseM()
}

// SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
func (p *IntSlice) SortReverseM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(sort.Reverse(p))
	return p
}

// Returns a string representation of this Slice, implements the Stringer interface
func (p *IntSlice) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	if p != nil {
		for i := range *p {
			builder.WriteString(fmt.Sprintf("%d", (*p)[i]))
			if i+1 < len(*p) {
				builder.WriteString(" ")
			}
		}
	}
	builder.WriteString("]")
	return builder.String()
}

// Swap modifies this Slice swapping the indicated elements.
func (p *IntSlice) Swap(i, j int) {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return
	}
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

// Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *IntSlice) Take(indices ...int) (new Slice) {
	new = p.Copy(indices...)
	p.Drop(indices...)
	return
}

// TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
// Allows for negative notation.
func (p *IntSlice) TakeAt(i int) (elem *Object) {
	elem = p.At(i)
	p.DropAt(i)
	return
}

// TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
func (p *IntSlice) TakeW(sel func(O) bool) (new Slice) {
	slice := NewIntSliceV()
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

// ToInterSlice converts the given slice to a generic []interface{} slice
func (p *IntSlice) ToInterSlice() (slice []interface{}) {
	return ToInterSlice(p.O()).G()
}

// ToStringSlice converts the underlying slice into a *StringSlice
func (p *IntSlice) ToStringSlice() (slice *StringSlice) {
	return ToStringSlice(p.O())
}

// ToStringSliceG converts the underlying slice into a []string slice
func (p *IntSlice) ToStringSliceG() (slice []string) {
	return ToStringSliceG(p.O())
}

// Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports IntSlice, *IntSlice, []int or *[]int
func (p *IntSlice) Union(slice interface{}) (new Slice) {
	return p.Copy().UnionM(slice)
}

// UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports IntSlice, *IntSlice, []int or *[]int
func (p *IntSlice) UnionM(slice interface{}) Slice {
	return p.ConcatM(slice).UniqM()
}

// Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
func (p *IntSlice) Uniq() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	m := NewIntMapBool()
	slice := NewIntSliceV()
	for i := range *p {
		if ok := m.Set((*p)[i], true); ok {
			slice.Append((*p)[i])
		}
	}
	return slice
}

// UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
func (p *IntSlice) UniqM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	m := NewIntMapBool()
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
