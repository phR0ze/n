package n

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// InterSlice implements the Slice interface providing a generic way to work with slice types
// including convenience methods on par with rapid development languages. This type incurs
// some reflection overhead characteristics but not all and differs in one important way from
// the RefSlice type. The given slice will be converted to a slice of types interface and
// left that way while RefSlice keeps the internal types typed as they were originally.
type InterSlice []interface{}

// NewInterSlice uses reflection to encapsulate the given Go slice type inside a new *InterSlice.
// Expects a Go slice type to be provided and will create an empty *InterSlice if nothing valid
// is given.
func NewInterSlice(slice interface{}) (new *InterSlice) {
	return ToInterSlice(slice)
}

// NewInterSliceV creates a new *InterSlice from the given variadic elements. Always returns
// at least a reference to an empty InterSlice.
func NewInterSliceV(elems ...interface{}) (new *InterSlice) {
	if len(elems) == 0 {
		new = &InterSlice{}
	} else {
		s := InterSlice(elems)
		new = &s
	}
	return
}

// A is an alias to String for brevity
func (p *InterSlice) A() string {
	return p.String()
}

// All tests if this Slice is not empty or optionally if it contains
// all of the given variadic elements. Incompatible types will return false.
func (p *InterSlice) All(elems ...interface{}) bool {

	// No elements
	if p.Nil() || p.Len() == 0 {
		return false
	}

	// Not looking for anything
	if len(elems) == 0 {
		return true
	}

	// Looking for something specific returns false if incompatible type
	return p.AllS(elems)
}

// AllS tests if this Slice contains all of the given Slice's elements.
// Incompatible types will return false.
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) AllS(slice interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	elems := ToInterSlice(slice)
	for i := range *elems {
		found := false
		for j := range *p {
			if (*p)[j] == (*elems)[i] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Any tests if this Slice is not empty or optionally if it contains
// any of the given variadic elements. Incompatible types will return false.
func (p *InterSlice) Any(elems ...interface{}) bool {

	// No elements
	if p.Nil() || p.Len() == 0 {
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
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) AnyS(slice interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	elems := ToInterSlice(slice)
	for i := range *elems {
		for j := range *p {
			if (*p)[j] == (*elems)[i] {
				return true
			}
		}
	}
	return false
}

// AnyW tests if this Slice contains any that match the lambda selector.
func (p *InterSlice) AnyW(sel func(O) bool) bool {
	return p.CountW(sel) != 0
}

// Append an element to the end of this Slice and returns a reference to this Slice.
func (p *InterSlice) Append(elem interface{}) ISlice {
	if p == nil {
		p = NewInterSliceV()
	}
	*p = append(*p, elem)
	return p
}

// AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
func (p *InterSlice) AppendV(elems ...interface{}) ISlice {
	if p == nil {
		p = NewInterSliceV()
	}
	*p = append(*p, elems...)
	return p
}

// At returns the element at the given index location. Allows for negative notation.
func (p *InterSlice) At(i int) (elem *Object) {
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
func (p *InterSlice) Clear() ISlice {
	if p == nil {
		p = NewInterSliceV()
	} else {
		p.Drop()
	}
	return p
}

// Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
// Supports InterSlice, *InterSlice, []int or *[]int
func (p *InterSlice) Concat(slice interface{}) (new ISlice) {
	return p.Copy().ConcatM(slice)
}

// ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) ConcatM(slice interface{}) ISlice {
	if p == nil {
		p = NewInterSliceV()
	}
	elems := ToInterSlice(slice)
	*p = append(*p, *elems...)
	return p
}

// Copy returns a new Slice with the indicated range of elements copied from this Slice.
// Expects nothing, in which case everything is copied, or two indices i and j, in which
// case positive and negative notation is supported and uses an inclusive behavior such
// that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of
// bounds indices will be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
func (p *InterSlice) Copy(indices ...int) (new ISlice) {
	if p == nil || len(*p) == 0 {
		return NewInterSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewInterSliceV()
	}

	// Copy elements over to new Slice
	x := make([]interface{}, j-i, j-i)
	copy(x, (*p)[i:j])
	return NewInterSlice(x)
}

// Count the number of elements in this Slice equal to the given element.
func (p *InterSlice) Count(elem interface{}) (cnt int) {
	cnt = p.CountW(func(x O) bool { return ExB(x == elem) })
	return
}

// CountW counts the number of elements in this Slice that match the lambda selector.
func (p *InterSlice) CountW(sel func(O) bool) (cnt int) {
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
func (p *InterSlice) Drop(indices ...int) ISlice {
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
func (p *InterSlice) DropAt(i int) ISlice {
	return p.Drop(i, i)
}

// DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
func (p *InterSlice) DropFirst() ISlice {
	return p.Drop(0, 0)
}

// DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
func (p *InterSlice) DropFirstN(n int) ISlice {
	if n == 0 {
		return p
	}
	return p.Drop(0, abs(n)-1)
}

// DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
func (p *InterSlice) DropLast() ISlice {
	return p.Drop(-1, -1)
}

// DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
func (p *InterSlice) DropLastN(n int) ISlice {
	if n == 0 {
		return p
	}
	return p.Drop(absNeg(n), -1)
}

// DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// The slice is updated instantly when lambda expression is evaluated not after DropW completes.
func (p *InterSlice) DropW(sel func(O) bool) ISlice {
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
func (p *InterSlice) Each(action func(O)) ISlice {
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
func (p *InterSlice) EachE(action func(O) error) (ISlice, error) {
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
func (p *InterSlice) EachI(action func(int, O)) ISlice {
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
func (p *InterSlice) EachIE(action func(int, O) error) (ISlice, error) {
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
func (p *InterSlice) EachR(action func(O)) ISlice {
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
func (p *InterSlice) EachRE(action func(O) error) (ISlice, error) {
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
func (p *InterSlice) EachRI(action func(int, O)) ISlice {
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
func (p *InterSlice) EachRIE(action func(int, O) error) (ISlice, error) {
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
func (p *InterSlice) Empty() bool {
	if p == nil || len(*p) == 0 {
		return true
	}
	return false
}

// First returns the first element in this Slice as Object.
// Object.Nil() == true will be returned when there are no elements in the slice.
func (p *InterSlice) First() (elem *Object) {
	return p.At(0)
}

// FirstN returns the first n elements in this slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *InterSlice) FirstN(n int) ISlice {
	if n == 0 {
		return NewInterSliceV()
	}
	return p.Slice(0, abs(n)-1)
}

// G returns the underlying Go type as is
func (p *InterSlice) G() []interface{} {
	if p == nil {
		return []interface{}{}
	}
	return []interface{}(*p)
}

// Index returns the index of the first element in this Slice where element == elem
// Returns a -1 if the element was not not found.
func (p *InterSlice) Index(elem interface{}) (loc int) {
	loc = -1
	if p == nil || len(*p) == 0 {
		return
	}
	for i := range *p {
		if (*p)[i] == elem {
			return i
		}
	}
	return
}

// Insert modifies this Slice to insert the given element before the element with the given index.
// Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// negative index is used, the given element will be inserted after that element, so using an index
// of -1 will insert the element at the end of the slice. If a Slice is given all elements will be
// inserted starting from the beging until the end. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *InterSlice) Insert(i int, obj interface{}) ISlice {
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
	elems := ToInterSlice(obj)
	if j == 0 {
		*p = append(*elems, *p...)
	} else if j < len(*p) {
		*p = append(*p, *elems...)           // ensures enough space exists
		copy((*p)[j+len(*elems):], (*p)[j:]) // shifts right elements drop added
		copy((*p)[j:], *elems)               // set new in locations vacated
	} else {
		*p = append(*p, *elems...)
	}
	return p
}

// InterSlice returns true if the underlying implementation is a RefSlice
func (p *InterSlice) InterSlice() bool {
	return true
}

// Join converts each element into a string then joins them together using the given separator or comma by default.
func (p *InterSlice) Join(separator ...string) (str *Object) {
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
func (p *InterSlice) Last() (elem *Object) {
	return p.At(-1)
}

// LastN returns the last n elements in this Slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *InterSlice) LastN(n int) ISlice {
	if n == 0 {
		return NewInterSliceV()
	}
	return p.Slice(absNeg(n), -1)
}

// Len returns the number of elements in this Slice
func (p *InterSlice) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// Less returns true if the element indexed by i is less than the element indexed by j.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) Less(i, j int) bool {
	l := p.Len()
	if p.Nil() || l < 2 || i < 0 || j < 0 || i >= l || j >= l {
		return false
	}

	// Handle supported types
	slice := Slice(*p)
	if !slice.RefSlice() {
		return slice.Less(i, j)
	}

	panic(fmt.Sprintf("unsupported comparable type '%T'", *p))
}

// Map creates a new slice with the modified elements from the lambda.
func (p *InterSlice) Map(mod func(O) O) ISlice {
	var slice ISlice
	if p == nil || len(*p) == 0 {
		return NewInterSliceV()
	}
	for i := range *p {
		v := mod((*p)[i])
		if slice == nil {
			slice = Slice(v)
		} else {
			slice.Append(v)
		}
	}
	return slice
}

// Nil tests if this Slice is nil
func (p *InterSlice) Nil() bool {
	if p == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *InterSlice) O() interface{} {
	if p.Nil() {
		return []interface{}{}
	}
	return []interface{}(*p)
}

// Pair simply returns the first and second Slice elements as Objects
func (p *InterSlice) Pair() (first, second *Object) {
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
func (p *InterSlice) Pop() (elem *Object) {
	elem = p.Last()
	p.DropLast()
	return
}

// PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
func (p *InterSlice) PopN(n int) (new ISlice) {
	if n == 0 {
		return NewInterSliceV()
	}
	new = p.Copy(absNeg(n), -1)
	p.DropLastN(n)
	return
}

// Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
func (p *InterSlice) Prepend(elem interface{}) ISlice {
	return p.Insert(0, elem)
}

// RefSlice returns true if the underlying implementation is a RefSlice
func (p *InterSlice) RefSlice() bool {
	return false
}

// Reverse returns a new Slice with the order of the elements reversed.
func (p *InterSlice) Reverse() (new ISlice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().ReverseM()
}

// ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
func (p *InterSlice) ReverseM() ISlice {
	if p == nil || len(*p) == 0 {
		return p
	}
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		p.Swap(i, j)
	}
	return p
}

// S is an alias to ToStringSlice
func (p *InterSlice) S() (slice *StringSlice) {
	return ToStringSlice(p.O())
}

// Select creates a new slice with the elements that match the lambda selector.
func (p *InterSlice) Select(sel func(O) bool) (new ISlice) {
	slice := NewInterSliceV()
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

// Set the element at the given index location to the given element. Allows for negative notation.
// Returns a reference to this Slice and swallows any errors.
func (p *InterSlice) Set(i int, elem interface{}) ISlice {
	slice, _ := p.SetE(i, elem)
	return slice
}

// SetE the element at the given index location to the given element. Allows for negative notation.
// Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
func (p *InterSlice) SetE(i int, elems interface{}) (ISlice, error) {
	var err error
	if p == nil {
		return p, err
	}
	if i = absIndex(len(*p), i); i == -1 {
		err = errors.Errorf("slice assignment is out of bounds")
		return p, err
	}

	// Account for length of elems
	x := ToInterSlice(elems)
	if len(*x) > 0 {
		copy((*p)[i:], *x)
	}
	return p, err
}

// Shift modifies this Slice to remove the first element and returns the removed element as an Object.
func (p *InterSlice) Shift() (elem *Object) {
	elem = p.First()
	p.DropFirst()
	return
}

// ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
func (p *InterSlice) ShiftN(n int) (new ISlice) {
	if n == 0 {
		return NewInterSliceV()
	}
	new = p.Copy(0, abs(n)-1)
	p.DropFirstN(n)
	return
}

// Single reports true if there is only one element in this Slice.
func (p *InterSlice) Single() bool {
	return p.Len() == 1
}

// Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
// Expects nothing, in which case everything is included, or two indices i and j, in which case an inclusive behavior
// is used such that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices will
// be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. NewInterSliceV(1,2,3).Slice(0, -1) == [1,2,3] && NewInterSliceV(1,2,3).Slice(1,2) == [2,3]
func (p *InterSlice) Slice(indices ...int) ISlice {
	if p == nil || len(*p) == 0 {
		return NewInterSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewInterSliceV()
	}

	slice := InterSlice((*p)[i:j])
	return &slice
}

// Sort returns a new Slice with sorted elements.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) Sort() (new ISlice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortM()
}

// SortM modifies this Slice sorting the elements and returns a reference to this Slice.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) SortM() ISlice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(p)
	return p
}

// SortReverse returns a new Slice sorting the elements in reverse.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) SortReverse() (new ISlice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortReverseM()
}

// SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) SortReverseM() ISlice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(sort.Reverse(p))
	return p
}

// Returns a string representation of this Slice, implements the Stringer interface
func (p *InterSlice) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	if p != nil {
		for i := range *p {
			builder.WriteString(ToString((*p)[i]))
			if i+1 < len(*p) {
				builder.WriteString(" ")
			}
		}
	}
	builder.WriteString("]")
	return builder.String()
}

// Swap modifies this Slice swapping the indicated elements.
func (p *InterSlice) Swap(i, j int) {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return
	}
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

// Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *InterSlice) Take(indices ...int) (new ISlice) {
	new = p.Copy(indices...)
	p.Drop(indices...)
	return
}

// TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
// Allows for negative notation.
func (p *InterSlice) TakeAt(i int) (elem *Object) {
	elem = p.At(i)
	p.DropAt(i)
	return
}

// TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
func (p *InterSlice) TakeW(sel func(O) bool) (new ISlice) {
	slice := NewInterSliceV()
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

// ToInts converts the underlying slice into a []int
func (p *InterSlice) ToInts() (slice []int) {
	return ToIntSlice(p.O()).G()
}

// ToIntSlice converts the underlying slice into a *IntSlice
func (p *InterSlice) ToIntSlice() (slice *IntSlice) {
	return ToIntSlice(p.O())
}

// ToInterSlice converts the given slice to a generic []interface{} slice
func (p *InterSlice) ToInterSlice() (slice []interface{}) {
	if p == nil {
		return []interface{}{}
	}
	return []interface{}(*p)
}

// ToStringSlice converts the underlying slice into a *StringSlice
func (p *InterSlice) ToStringSlice() (slice *StringSlice) {
	return ToStringSlice(p.O())
}

// ToStrs converts the underlying slice into a []string slice
func (p *InterSlice) ToStrs() (slice []string) {
	return ToStrs(p.O())
}

// Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) Union(slice interface{}) (new ISlice) {
	return p.Copy().UnionM(slice)
}

// UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) UnionM(slice interface{}) ISlice {
	return p.ConcatM(slice).UniqM()
}

// Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
func (p *InterSlice) Uniq() (new ISlice) {
	panic("NOT IMPLEMENTED")
}

// UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
func (p *InterSlice) UniqM() ISlice {
	panic("NOT IMPLEMENTED")
}
