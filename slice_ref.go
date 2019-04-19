package n

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

// RefSlice implements the Slice interface providing a generic way to work with slice types
// including convenience methods on par with other rapid development languages. This type
// incurs the typical 10x reflection overhead costs. For high performance use the Slice
// implementation matching the type your working with or implement a new type that satisfies
// the Slice interface.
type RefSlice struct {
	k reflect.Kind
	v *reflect.Value
}

// NewRefSlice uses reflection to encapsulate the given Go slice type inside a new *RefSlice.
// Expects a Go slice type to be provided and will create an empty *RefSlice if nothing valid
// is given.
func NewRefSlice(slice interface{}) (new *RefSlice) {
	new = &RefSlice{}
	v := reflect.ValueOf(slice)
	k := v.Kind()
	x, interfaceSliceType := slice.([]interface{})
	switch {

	// Return the NSlice.Nil
	case k == reflect.Invalid:

	// Iterate over array and append
	case k == reflect.Array:
		new = newEmptySlice(slice)
		for i := 0; i < v.Len(); i++ {
			*new.v = reflect.Append(*new.v, v.Index(i))
		}

	// Convert []interface to slice of elem type
	case interfaceSliceType:
		new = NewRefSliceV(x...)

	// Slice of distinct type can be used directly
	case k == reflect.Slice:
		new.v = &v
		new.k = k

	// Append single items
	default:
		new = newEmptySlice(slice)
		*new.v = reflect.Append(*new.v, v)
	}
	return
}

// NewRefSliceV creates a new *RefSlice from the given variadic elements. Always returns
// at least a reference to an empty RefSlice.
func NewRefSliceV(elems ...interface{}) (new *RefSlice) {
	new = &RefSlice{}

	// Return RefSlice.Nil if nothing given
	if len(elems) == 0 {
		return
	}

	// Create new slice from the type of the first non Invalid element
	var slice *reflect.Value
	for i := 0; i < len(elems); i++ {

		// Create target slice from first Valid element
		if slice == nil && reflect.ValueOf(elems[i]).IsValid() {
			typ := reflect.SliceOf(reflect.TypeOf(elems[i]))
			v := reflect.MakeSlice(typ, 0, 10)
			slice = &v
		}

		// Append element to slice
		if slice != nil {
			elem := reflect.ValueOf(elems[i])
			*slice = reflect.Append(*slice, elem)
		}
	}
	if slice != nil {
		new.v = slice
		new.k = slice.Kind()
	}
	return
}

// create a new empty slice of the given type or element type if a slice/array
// want to return a new Slice so that we can use this in the AppendX functions
// to defer creating an underlying slice type until we have an actual type to work with.
func newEmptySlice(elems interface{}) (new *RefSlice) {
	new = &RefSlice{}
	v := reflect.ValueOf(elems)
	typ := reflect.TypeOf([]interface{}{})

	k := v.Kind()
	switch k {

	// Use a new generic slice for nils
	case reflect.Invalid:

	// Use the element type of slice/arrays
	case reflect.Slice, reflect.Array:

		// Use slice type if not generic
		if _, ok := elems.([]interface{}); !ok {
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
	new.v = &slice
	new.k = k
	return
}

// A is an alias to String for brevity
func (p *RefSlice) A() string {
	return p.String()
}

// Any tests if this Slice is not empty or optionally if it contains
// any of the given variadic elements. Incompatible types will return false.
func (p *RefSlice) Any(elems ...interface{}) bool {

	// No elements
	if p.Nil() || p.Len() == 0 {
		return false
	}

	// Not looking for anything
	if len(elems) == 0 {
		return true
	}

	// Looking for something specific returns false if incompatible type
	for i := 0; i < len(elems); i++ {
		x := reflect.ValueOf(elems[i])
		if p.v.Type().Elem() != x.Type() {
			break
		} else {
			for j := 0; j < p.v.Len(); j++ {
				if p.v.Index(j).Interface() == x.Interface() {
					return true
				}
			}
		}
	}
	return false
}

// AnyS tests if this Slice contains any of the given Slice's elements.
// Incompatible types will return false.
// Supports RefSlice, *RefSlice, Slice and Go slice types
func (p *RefSlice) AnyS(slice interface{}) bool {

	// No elements
	if p.Nil() || p.Len() == 0 {
		return false
	}

	// Handle supported types
	var v reflect.Value
	if x, ok := slice.(RefSlice); ok {
		if !x.Nil() {
			v = *(x.v)
		}
	} else if x, ok := slice.(*RefSlice); ok {
		if !x.Nil() {
			v = *(x.v)
		}
	} else if x, ok := slice.(Slice); ok {
		if !x.Nil() {
			v = reflect.ValueOf(x.O())
		}
	} else {
		v = reflect.ValueOf(slice)
	}
	if !v.IsValid() {
		return false
	}

	if p.v.Type() == v.Type() {
		for i := 0; i < v.Len(); i++ {
			for j := 0; j < p.v.Len(); j++ {
				if p.v.Index(j).Interface() == v.Index(i).Interface() {
					return true
				}
			}
		}
	}
	return false
}

// AnyW tests if this Slice contains any that match the lambda selector.
func (p *RefSlice) AnyW(sel func(O) bool) bool {
	return p.CountW(sel) != 0
}

// Append an element to the end of this Slice and returns a reference to this Slice.
func (p *RefSlice) Append(elem interface{}) Slice {
	if p.Nil() {
		if p == nil {
			p = newEmptySlice(elem)
		} else {
			*p = *(newEmptySlice(elem))
		}
	}
	x := reflect.ValueOf(elem)
	if p.v.Type().Elem() != x.Type() {
		panic(fmt.Sprintf("can't append type '%v' to '%v'", x.Type(), p.v.Type()))
	} else {
		*p.v = reflect.Append(*p.v, x)
	}
	return p
}

// AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
func (p *RefSlice) AppendV(elems ...interface{}) Slice {
	if p.Nil() {
		if p == nil {
			p = newEmptySlice(elems)
		} else {
			*p = *(newEmptySlice(elems))
		}
	}
	for _, elem := range elems {
		p.Append(elem)
	}
	return p
}

// At returns the element at the given index location. Allows for negative notation.
func (p *RefSlice) At(i int) (elem *Object) {
	elem = &Object{}
	if p.Nil() {
		return
	}
	if i = absIndex(p.Len(), i); i == -1 {
		return
	}
	elem.o = p.v.Index(i).Interface()
	return
}

// Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
func (p *RefSlice) Clear() Slice {
	if p.Nil() {
		if p == nil {
			p = NewRefSliceV()
		} else {
			*p = *NewRefSliceV()
		}
	} else {
		p.Drop()
	}
	return p
}

// Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
// Supports RefSlice, *RefSlice, []int or *[]int
func (p *RefSlice) Concat(slice interface{}) (new Slice) {
	return p.Copy().ConcatM(slice)
}

// ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
// Supports RefSlice, *RefSlice, Slice and Go slice types
func (p *RefSlice) ConcatM(slice interface{}) Slice {

	// Handle supported types
	var v reflect.Value
	if x, ok := slice.(RefSlice); ok {
		if !x.Nil() {
			v = *(x.v)
		}
	} else if x, ok := slice.(*RefSlice); ok {
		if !x.Nil() {
			v = *(x.v)
		}
	} else if x, ok := slice.(Slice); ok {
		if !x.Nil() {
			v = reflect.ValueOf(x.O())
		}
	} else {
		v = reflect.ValueOf(slice)
	}
	if !v.IsValid() {
		return p
	}

	// Nothing in this slice so return new slice from given
	if p.Nil() {
		if p == nil {
			p = newEmptySlice(v.Interface())
		} else {
			*p = *(newEmptySlice(v.Interface()))
		}
	}

	// Concat the two slices
	if p.v.Type() != v.Type() {
		panic(fmt.Sprintf("can't concat type '%v' with '%v'", v.Type(), p.v.Type()))
	} else {
		*p.v = reflect.AppendSlice(*p.v, v)
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
func (p *RefSlice) Copy(indices ...int) (new Slice) {
	if p.Nil() {
		return NewRefSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(p.Len(), indices...)
	if err != nil {
		return newEmptySlice(p.O())
	}

	// Copy elements over to new Slice
	x := reflect.MakeSlice(p.v.Type(), j-i, j-i)
	reflect.Copy(x, p.v.Slice(i, j))
	new = &RefSlice{v: &x, k: x.Kind()}
	return
}

// Count the number of elements in this Slice equal to the given element.
func (p *RefSlice) Count(elem interface{}) (cnt int) {
	l := p.Len()
	if p.Nil() || l == 0 {
		return
	}
	x := reflect.ValueOf(elem)
	if p.v.Type().Elem() == x.Type() {
		for i := 0; i < l; i++ {
			if p.v.Index(i).Interface() == x.Interface() {
				cnt++
			}
		}
	}
	return
}

// CountW counts the number of elements in this Slice that match the lambda selector.
func (p *RefSlice) CountW(sel func(O) bool) (cnt int) {
	l := p.Len()
	if p.Nil() || l == 0 {
		return
	}
	for i := 0; i < l; i++ {
		if sel(p.v.Index(i).Interface()) {
			cnt++
		}
	}
	return
}

// Drop modifies this Slice to delete the indicated range of elements and returns a referece to this Slice.
// Expects nothing, in which case everything is dropped, or two indices i and j, in which case positive and
// negative notation is supported and uses an inclusive behavior such that DropAt(0, -1) includes index -1
// as opposed to Go's exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *RefSlice) Drop(indices ...int) Slice {
	l := p.Len()
	if p == nil || l == 0 {
		return p
	}

	// Handle index manipulation
	i, j, err := absIndices(l, indices...)
	if err != nil {
		return p
	}

	// Execute
	n := j - i
	if i+n < l {
		*p.v = reflect.AppendSlice(p.v.Slice(0, i), p.v.Slice(i+n, p.v.Len()))
	} else {
		*p.v = p.v.Slice(0, i)
	}
	return p
}

// DropAt modifies this Slice to delete the element at the given index location. Allows for negative notation.
// Returns a reference to this Slice.
func (p *RefSlice) DropAt(i int) Slice {
	return p.Drop(i, i)
}

// DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
func (p *RefSlice) DropFirst() Slice {
	return p.Drop(0, 0)
}

// DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
func (p *RefSlice) DropFirstN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(0, abs(n)-1)
}

// DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
func (p *RefSlice) DropLast() Slice {
	return p.Drop(-1, -1)
}

// DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
func (p *RefSlice) DropLastN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(absNeg(n), -1)
}

// DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// The slice is updated instantly when lambda expression is evaluated not after DropW completes.
func (p *RefSlice) DropW(sel func(O) bool) Slice {
	l := p.Len()
	if p.Nil() || l == 0 {
		return p
	}

	for i := 0; i < l; i++ {
		if sel(p.v.Index(i).Interface()) {
			p.DropAt(i)
			l--
			i--
		}
	}
	return p
}

// Each calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *RefSlice) Each(action func(O)) Slice {
	if p.Nil() {
		return p
	}
	for i := 0; i < p.Len(); i++ {
		action(p.v.Index(i).Interface())
	}
	return p
}

// EachE calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *RefSlice) EachE(action func(O) error) (Slice, error) {
	var err error
	if p.Nil() {
		return p, err
	}
	for i := 0; i < p.Len(); i++ {
		if err = action(p.v.Index(i).Interface()); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachI calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice
func (p *RefSlice) EachI(action func(int, O)) Slice {
	if p.Nil() {
		return p
	}
	for i := 0; i < p.Len(); i++ {
		action(i, p.v.Index(i).Interface())
	}
	return p
}

// EachIE calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *RefSlice) EachIE(action func(int, O) error) (Slice, error) {
	var err error
	if p.Nil() {
		return p, err
	}
	for i := 0; i < p.Len(); i++ {
		if err = action(i, p.v.Index(i).Interface()); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *RefSlice) EachR(action func(O)) Slice {
	if p.Nil() {
		return p
	}
	for i := p.Len() - 1; i >= 0; i-- {
		action(p.v.Index(i).Interface())
	}
	return p
}

// EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *RefSlice) EachRE(action func(O) error) (Slice, error) {
	var err error
	if p.Nil() {
		return p, err
	}
	for i := p.Len() - 1; i >= 0; i-- {
		if err = action(p.v.Index(i).Interface()); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *RefSlice) EachRI(action func(int, O)) Slice {
	if p.Nil() {
		return p
	}
	for i := p.Len() - 1; i >= 0; i-- {
		action(i, p.v.Index(i).Interface())
	}
	return p
}

// EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *RefSlice) EachRIE(action func(int, O) error) (Slice, error) {
	var err error
	if p.Nil() {
		return p, err
	}
	for i := p.Len() - 1; i >= 0; i-- {
		if err = action(i, p.v.Index(i).Interface()); err != nil {
			return p, err
		}
	}
	return p, err
}

// Empty tests if this Slice is empty.
func (p *RefSlice) Empty() bool {
	if p.Nil() || p.Len() == 0 {
		return true
	}
	return false
}

// First returns the first element in this Slice as Object.
// Object.Nil() == true will be returned when there are no elements in the slice.
func (p *RefSlice) First() (elem *Object) {
	return p.At(0)
}

// FirstN returns the first n elements in this slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *RefSlice) FirstN(n int) Slice {
	if p.Nil() {
		return NewRefSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
	}
	return p.Slice(0, abs(n)-1)
}

// Generic returns true if the underlying implementation is a RefSlice
func (p *RefSlice) Generic() bool {
	return true
}

// Index returns the index of the first element in this Slice where element == elem
// Returns a -1 if the element was not not found.
func (p *RefSlice) Index(elem interface{}) (loc int) {
	loc = -1
	l := p.Len()
	if p.Nil() || l == 0 {
		return
	}
	for i := 0; i < l; i++ {
		if elem == p.v.Index(i).Interface() {
			return i
		}
	}
	return
}

// Insert modifies this Slice to insert the given element(s) before the element with the given index.
// Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// negative index is used, the given element will be inserted after that element, so using an index
// of -1 will insert the element at the end of the slice. If a Slice is given all elements will be
// inserted starting from the beging until the end. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *RefSlice) Insert(i int, elem interface{}) Slice {
	l := p.Len()
	if p.Nil() || l == 0 {
		return p.Append(elem)
	}
	j := i
	if j = absIndex(l, j); j == -1 {
		return p
	}
	if i < 0 {
		j++
	}

	// Insert the item before j if pos and after j if neg
	x := reflect.ValueOf(elem)
	if p.v.Type().Elem() != x.Type() {
		panic(fmt.Sprintf("can't insert type '%v' into '%v'", x.Type(), p.v.Type()))
	} else {
		if j == 0 {
			*p.v = reflect.Append(*p.v, x)
			reflect.Copy(p.v.Slice(1, p.v.Len()), p.v.Slice(0, p.v.Len()-1))
			p.v.Index(0).Set(x)
		} else if j < l {
			*p.v = reflect.Append(*p.v, x)
			reflect.Copy(p.v.Slice(j+1, p.v.Len()), p.v.Slice(j, p.v.Len()))
			p.v.Index(j).Set(x)
		} else {
			*p.v = reflect.Append(*p.v, x)
		}
	}
	return p
}

// Join converts each element into a string then joins them together using the given separator or comma by default.
func (p *RefSlice) Join(separator ...string) (str *Object) {
	l := p.Len()
	if p.Nil() || l == 0 {
		str = &Object{""}
		return
	}
	sep := ","
	if len(separator) > 0 {
		sep = separator[0]
	}

	var builder strings.Builder
	for i := 0; i < l; i++ {
		builder.WriteString(Obj(p.v.Index(i).Interface()).ToString())
		if i+1 < l {
			builder.WriteString(sep)
		}
	}
	str = &Object{builder.String()}
	return
}

// Last returns the last element in this Slice as an Object.
// Object.Nil() == true will be returned if there are no elements in the slice.
func (p *RefSlice) Last() (elem *Object) {
	return p.At(-1)
}

// LastN returns the last n elements in this Slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *RefSlice) LastN(n int) Slice {
	if p.Nil() {
		return NewRefSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
	}
	return p.Slice(absNeg(n), -1)
}

// Len returns the number of elements in this Slice
func (p *RefSlice) Len() int {
	if p.Nil() {
		return 0
	}
	return p.v.Len()
}

// Less returns true if the element indexed by i is less than the element indexed by j.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *RefSlice) Less(i, j int) bool {
	l := p.Len()
	if p.Nil() || l < 2 || i < 0 || j < 0 || i >= l || j >= l {
		return false
	}

	// Handle supported types
	slice := NewSlice(p.v.Interface())
	if !slice.Generic() {
		return slice.Less(i, j)
	}

	panic(fmt.Sprintf("unsupported comparable type '%v'", p.v.Type()))
}

// Nil tests if this Slice is nil
func (p *RefSlice) Nil() bool {
	if p == nil || p.v == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *RefSlice) O() interface{} {
	if p.Nil() {
		return nil
	}
	return p.v.Interface()
}

// Pair simply returns the first and second Slice elements as Objects
func (p *RefSlice) Pair() (first, second *Object) {
	l := p.Len()
	first, second = &Object{}, &Object{}
	if l > 0 {
		first = p.At(0)
	}
	if l > 1 {
		second = p.At(1)
	}
	return
}

// Pop modifies this Slice to remove the last element and returns the removed element as an Object.
func (p *RefSlice) Pop() (elem *Object) {
	elem = p.Last()
	p.DropLast()
	return
}

// PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
func (p *RefSlice) PopN(n int) (new Slice) {
	if p.Nil() {
		return NewRefSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
	}
	new = p.Copy(absNeg(n), -1)
	p.DropLastN(n)
	return
}

// Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
func (p *RefSlice) Prepend(elem interface{}) Slice {
	return p.Insert(0, elem)
}

// Reverse returns a new Slice with the order of the elements reversed.
func (p *RefSlice) Reverse() (new Slice) {
	if p.Nil() || p.Len() < 2 {
		return p.Copy()
	}
	return p.Copy().ReverseM()
}

// ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
func (p *RefSlice) ReverseM() Slice {
	l := p.Len()
	if p.Nil() || l == 0 {
		return p
	}
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		p.Swap(i, j)
	}
	return p
}

// Select creates a new slice with the elements that match the lambda selector.
func (p *RefSlice) Select(sel func(O) bool) (new Slice) {
	l := p.Len()
	slice := NewRefSliceV()
	if p.Nil() || l == 0 {
		return slice
	}
	for i := 0; i < l; i++ {
		obj := p.v.Index(i).Interface()
		if sel(obj) {
			slice.Append(obj)
		}
	}
	return slice
}

// Set the element at the given index location to the given element. Allows for negative notation.
// Returns a reference to this Slice and swallows any errors.
func (p *RefSlice) Set(i int, elem interface{}) Slice {
	slice, _ := p.SetE(i, elem)
	return slice
}

// SetE the element at the given index location to the given element. Allows for negative notation.
// Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
func (p *RefSlice) SetE(i int, elem interface{}) (Slice, error) {
	var err error
	if p.Nil() {
		return p, err
	}
	if i = absIndex(p.Len(), i); i == -1 {
		err = errors.Errorf("slice assignment is out of bounds")
		return p, err
	}

	x := reflect.ValueOf(elem)
	if p.v.Type().Elem() != x.Type() {
		err = errors.Errorf("can't set type '%v' in '%v'", x.Type(), p.v.Type())
	} else {
		p.v.Index(i).Set(x)
	}
	return p, err
}

// Shift modifies this Slice to remove the first element and returns the removed element as an Object.
func (p *RefSlice) Shift() (elem *Object) {
	elem = p.First()
	p.DropFirst()
	return
}

// ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
func (p *RefSlice) ShiftN(n int) (new Slice) {
	if p.Nil() {
		return NewRefSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
	}
	new = p.Copy(0, abs(n)-1)
	p.DropFirstN(n)
	return
}

// Single reports true if there is only one element in this Slice.
func (p *RefSlice) Single() bool {
	return p.Len() == 1
}

// Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
// Expects nothing, in which case everything is included, or two indices i and j, in which case an inclusive behavior
// is used such that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices will
// be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. NewRefSliceV(1,2,3).Slice(0, -1) == [1,2,3] && NewRefSliceV(1,2,3).Slice(1,2) == [2,3]
func (p *RefSlice) Slice(indices ...int) Slice {
	if p.Nil() {
		return NewRefSliceV()
	}

	// Handle index manipulation
	i, j, err := absIndices(p.Len(), indices...)
	if err != nil {
		return newEmptySlice(p.O())
	}

	return NewRefSlice(p.v.Slice(i, j).Interface())
}

// Sort returns a new Slice with sorted elements.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *RefSlice) Sort() (new Slice) {
	if p.Nil() || p.Len() < 2 {
		return p.Copy()
	}
	return p.Copy().SortM()
}

// SortM modifies this Slice sorting the elements and returns a reference to this Slice.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *RefSlice) SortM() Slice {
	if p.Nil() || p.Len() < 2 {
		return p
	}

	// Handle supported types
	slice := NewSlice(p.v.Interface())
	if !slice.Generic() {
		slice.SortM()
		*p = *NewRefSlice(slice.O())
	} else {
		panic(fmt.Sprintf("unsupported comparable type '%v'", p.v.Type()))
	}

	return p
}

// SortReverse returns a new Slice sorting the elements in reverse.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *RefSlice) SortReverse() (new Slice) {
	if p.Nil() || p.Len() < 2 {
		return p.Copy()
	}
	return p.Copy().SortReverseM()
}

// SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *RefSlice) SortReverseM() Slice {
	if p.Nil() || p.Len() < 2 {
		return p
	}

	// Handle supported types
	slice := NewSlice(p.v.Interface())
	if !slice.Generic() {
		slice.SortReverseM()
		*p = *NewRefSlice(slice.O())
	} else {
		panic(fmt.Sprintf("unsupported comparable type '%v'", p.v.Type()))
	}

	return p
}

// Returns a string representation of this Slice, implements the Stringer interface
func (p *RefSlice) String() string {
	l := p.Len()
	var builder strings.Builder
	builder.WriteString("[")
	for i := 0; i < l; i++ {
		builder.WriteString(fmt.Sprintf("%d", p.v.Index(i).Interface()))
		if i+1 < l {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

// Swap modifies this Slice swapping the indicated elements.
func (p *RefSlice) Swap(i, j int) {
	l := p.Len()
	if p.Nil() || l < 2 || i < 0 || j < 0 || i >= l || j >= l {
		return
	}
	reflect.Swapper(p.v.Interface())(i, j)
}

// Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *RefSlice) Take(indices ...int) (new Slice) {
	new = p.Copy(indices...)
	p.Drop(indices...)
	return
}

// TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
// Allows for negative notation.
func (p *RefSlice) TakeAt(i int) (elem *Object) {
	elem = p.At(i)
	p.DropAt(i)
	return
}

// TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
func (p *RefSlice) TakeW(sel func(O) bool) (new Slice) {
	l := p.Len()
	slice := NewRefSliceV()
	if p.Nil() || l == 0 {
		return slice
	}
	for i := 0; i < l; i++ {
		obj := p.v.Index(i).Interface()
		if sel(obj) {
			slice.Append(obj)
			p.DropAt(i)
			l--
			i--
		}
	}
	return slice
}

// Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports RefSlice, *RefSlice, Slice and Go slice types
func (p *RefSlice) Union(slice interface{}) (new Slice) {
	return p.Copy().UnionM(slice)
}

// UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports RefSlice, *RefSlice, Slice and Go slice types
func (p *RefSlice) UnionM(slice interface{}) Slice {
	return p.ConcatM(slice).UniqM()
}

// Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
func (p *RefSlice) Uniq() (new Slice) {
	l := p.Len()
	if p.Nil() || l < 2 {
		return p.Copy()
	}
	slice := NewRefSliceV()
	v := reflect.ValueOf(true)
	typ := reflect.MapOf(p.v.Type().Elem(), v.Type())
	m := reflect.MakeMap(typ)
	for i := 0; i < l; i++ {
		k := p.v.Index(i)
		if ok := m.MapIndex(k); ok == (reflect.Value{}) {
			m.SetMapIndex(k, v)
			slice.Append(k.Interface())
		}
	}
	return slice
}

// UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
func (p *RefSlice) UniqM() Slice {
	l := p.Len()
	if p.Nil() || l < 2 {
		return p
	}
	v := reflect.ValueOf(true)
	typ := reflect.MapOf(p.v.Type().Elem(), v.Type())
	m := reflect.MakeMap(typ)
	for i := 0; i < l; i++ {
		k := p.v.Index(i)
		if ok := m.MapIndex(k); ok == (reflect.Value{}) {
			m.SetMapIndex(k, v)
		} else {
			p.DropAt(i)
			l--
			i--
		}
	}
	return p
}
