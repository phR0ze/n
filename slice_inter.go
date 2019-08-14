package n

import (
	"reflect"
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
	new = &InterSlice{}
	v := reflect.ValueOf(slice)
	k := v.Kind()
	x, interfaceSliceType := slice.([]interface{})
	switch {

	// Return the NSlice.Nil
	case k == reflect.Invalid:

	// Slice of distinct types or or arrays must be converted
	case k == reflect.Array || k == reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			*new = append(*new, v.Index(i).Interface())
		}

	// []interface slice can be used directly
	case interfaceSliceType:
		s := InterSlice(x)
		new = &s

	// Append single items
	default:
		*new = append(*new, slice)
	}
	return
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
func (p *InterSlice) Append(elem interface{}) Slice {
	if p == nil {
		p = NewInterSliceV()
	}
	*p = append(*p, elem)
	return p
}

// AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
func (p *InterSlice) AppendV(elems ...interface{}) Slice {
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
func (p *InterSlice) Clear() Slice {
	if p == nil {
		p = NewInterSliceV()
	} else {
		p.Drop()
	}
	return p
}

// Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
// Supports InterSlice, *InterSlice, []int or *[]int
func (p *InterSlice) Concat(slice interface{}) (new Slice) {
	return p.Copy().ConcatM(slice)
}

// ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) ConcatM(slice interface{}) Slice {
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
func (p *InterSlice) Copy(indices ...int) (new Slice) {
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
func (p *InterSlice) Drop(indices ...int) Slice {
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
func (p *InterSlice) DropAt(i int) Slice {
	return p.Drop(i, i)
}

// DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
func (p *InterSlice) DropFirst() Slice {
	return p.Drop(0, 0)
}

// DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
func (p *InterSlice) DropFirstN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(0, abs(n)-1)
}

// DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
func (p *InterSlice) DropLast() Slice {
	return p.Drop(-1, -1)
}

// DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
func (p *InterSlice) DropLastN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(absNeg(n), -1)
}

// DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// The slice is updated instantly when lambda expression is evaluated not after DropW completes.
func (p *InterSlice) DropW(sel func(O) bool) Slice {
	// l := p.Len()
	// if p.Nil() || l == 0 {
	// 	return p
	// }

	// for i := 0; i < l; i++ {
	// 	if sel(p.v.Index(i).Interface()) {
	// 		p.DropAt(i)
	// 		l--
	// 		i--
	// 	}
	// }
	return p
}

// Each calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *InterSlice) Each(action func(O)) Slice {
	// if p.Nil() {
	// 	return p
	// }
	// for i := 0; i < p.Len(); i++ {
	// 	action(p.v.Index(i).Interface())
	// }
	return p
}

// EachE calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *InterSlice) EachE(action func(O) error) (slice Slice, err error) {
	// slice = p
	// if p.Nil() {
	// 	return
	// }
	// for i := 0; i < p.Len(); i++ {
	// 	if err = action(p.v.Index(i).Interface()); err != nil {
	// 		return
	// 	}
	// }
	return
}

// EachI calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice
func (p *InterSlice) EachI(action func(int, O)) Slice {
	// if p.Nil() {
	// 	return p
	// }
	// for i := 0; i < p.Len(); i++ {
	// 	action(i, p.v.Index(i).Interface())
	// }
	return p
}

// EachIE calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *InterSlice) EachIE(action func(int, O) error) (slice Slice, err error) {
	// slice = p
	// if p.Nil() {
	// 	return
	// }
	// for i := 0; i < p.Len(); i++ {
	// 	if err = action(i, p.v.Index(i).Interface()); err != nil {
	// 		return
	// 	}
	// }
	return
}

// EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *InterSlice) EachR(action func(O)) Slice {
	// if p.Nil() {
	// 	return p
	// }
	// for i := p.Len() - 1; i >= 0; i-- {
	// 	action(p.v.Index(i).Interface())
	// }
	return p
}

// EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *InterSlice) EachRE(action func(O) error) (slice Slice, err error) {
	// slice = p
	// if p.Nil() {
	// 	return
	// }
	// for i := p.Len() - 1; i >= 0; i-- {
	// 	if err = action(p.v.Index(i).Interface()); err != nil {
	// 		return
	// 	}
	// }
	return
}

// EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *InterSlice) EachRI(action func(int, O)) Slice {
	// if p.Nil() {
	// 	return p
	// }
	// for i := p.Len() - 1; i >= 0; i-- {
	// 	action(i, p.v.Index(i).Interface())
	// }
	return p
}

// EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *InterSlice) EachRIE(action func(int, O) error) (slice Slice, err error) {
	// slice = p
	// if p.Nil() {
	// 	return
	// }
	// for i := p.Len() - 1; i >= 0; i-- {
	// 	if err = action(i, p.v.Index(i).Interface()); err != nil {
	// 		return
	// 	}
	// }
	return
}

// Empty tests if this Slice is empty.
func (p *InterSlice) Empty() bool {
	if p.Nil() || p.Len() == 0 {
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
func (p *InterSlice) FirstN(n int) Slice {
	if p.Nil() {
		return NewInterSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
	}
	return p.Slice(0, abs(n)-1)
}

// G returns the underlying Go type as is
func (p *InterSlice) G() []interface{} {
	if p.Nil() {
		return []interface{}{}
	}
	return []interface{}(*p)
}

// Generic returns true if the underlying implementation is a InterSlice
func (p *InterSlice) Generic() bool {
	return true
}

// Index returns the index of the first element in this Slice where element == elem
// Returns a -1 if the element was not not found.
func (p *InterSlice) Index(elem interface{}) (loc int) {
	// loc = -1
	// l := p.Len()
	// if p.Nil() || l == 0 {
	// 	return
	// }
	// for i := 0; i < l; i++ {
	// 	if elem == p.v.Index(i).Interface() {
	// 		return i
	// 	}
	// }
	return
}

// Insert modifies this Slice to insert the given element before the element with the given index.
// Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// negative index is used, the given element will be inserted after that element, so using an index
// of -1 will insert the element at the end of the slice. If a Slice is given all elements will be
// inserted starting from the beging until the end. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *InterSlice) Insert(i int, elem interface{}) Slice {
	// l := p.Len()
	// if p.Nil() || l == 0 {
	// 	return p.Append(elem)
	// }

	// // Insert the item before j if pos and after j if neg
	// j := i
	// if j = absIndex(l, j); j == -1 {
	// 	return p
	// }
	// if i < 0 {
	// 	j++
	// }
	// x := reflect.ValueOf(elem)
	// if p.v.Type().Elem() != x.Type() {
	// 	panic(fmt.Sprintf("can't insert type '%v' into '%v'", x.Type(), p.v.Type()))
	// } else {
	// 	if j == 0 {
	// 		*p.v = reflect.Append(*p.v, x)
	// 		reflect.Copy(p.v.Slice(1, p.v.Len()), p.v.Slice(0, p.v.Len()-1))
	// 		p.v.Index(0).Set(x)
	// 	} else if j < l {
	// 		*p.v = reflect.Append(*p.v, x)
	// 		reflect.Copy(p.v.Slice(j+1, p.v.Len()), p.v.Slice(j, p.v.Len()))
	// 		p.v.Index(j).Set(x)
	// 	} else {
	// 		*p.v = reflect.Append(*p.v, x)
	// 	}
	// }
	return p
}

// InsertS modifies this Slice to insert the given elements before the element with the given index.
// Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// negative index is used, the given element will be inserted after that element, so using an index
// of -1 will insert the element at the end of the slice. If a Slice is given all elements will be
// inserted starting from the beging until the end. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *InterSlice) InsertS(i int, slice interface{}) Slice {
	l := p.Len()
	if p.Nil() || l == 0 {
		return p.ConcatM(slice)
	}

	// // Insert the item before j if pos and after j if neg
	j := i
	if j = absIndex(l, j); j == -1 {
		return p
	}
	if i < 0 {
		j++
	}
	// x := reflect.ValueOf(elem)
	// if p.v.Type().Elem() != x.Type() {
	// 	panic(fmt.Sprintf("can't insert type '%v' into '%v'", x.Type(), p.v.Type()))
	// } else {
	// 	if j == 0 {
	// 		*p.v = reflect.Append(*p.v, x)
	// 		reflect.Copy(p.v.Slice(1, p.v.Len()), p.v.Slice(0, p.v.Len()-1))
	// 		p.v.Index(0).Set(x)
	// 	} else if j < l {
	// 		*p.v = reflect.Append(*p.v, x)
	// 		reflect.Copy(p.v.Slice(j+1, p.v.Len()), p.v.Slice(j, p.v.Len()))
	// 		p.v.Index(j).Set(x)
	// 	} else {
	// 		*p.v = reflect.Append(*p.v, x)
	// 	}
	// }
	return p
}

// Join converts each element into a string then joins them together using the given separator or comma by default.
func (p *InterSlice) Join(separator ...string) (str *Object) {
	// l := p.Len()
	// if p.Nil() || l == 0 {
	// 	str = &Object{""}
	// 	return
	// }
	// sep := ","
	// if len(separator) > 0 {
	// 	sep = separator[0]
	// }

	// var builder strings.Builder
	// for i := 0; i < l; i++ {
	// 	builder.WriteString(Obj(p.v.Index(i).Interface()).ToString())
	// 	if i+1 < l {
	// 		builder.WriteString(sep)
	// 	}
	// }
	// str = &Object{builder.String()}
	return
}

// Last returns the last element in this Slice as an Object.
// Object.Nil() == true will be returned if there are no elements in the slice.
func (p *InterSlice) Last() (elem *Object) {
	return p.At(-1)
}

// LastN returns the last n elements in this Slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *InterSlice) LastN(n int) Slice {
	if p.Nil() {
		return NewInterSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
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
	// l := p.Len()
	// if p.Nil() || l < 2 || i < 0 || j < 0 || i >= l || j >= l {
	// 	return false
	// }

	// // Handle supported types
	// slice := NewSlice(p.v.Interface())
	// if !slice.Generic() {
	// 	return slice.Less(i, j)
	// }

	// panic(fmt.Sprintf("unsupported comparable type '%v'", p.v.Type()))
	return false
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
func (p *InterSlice) Pop() (elem *Object) {
	elem = p.Last()
	p.DropLast()
	return
}

// PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
func (p *InterSlice) PopN(n int) (new Slice) {
	if p.Nil() {
		return NewInterSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
	}
	new = p.Copy(absNeg(n), -1)
	p.DropLastN(n)
	return
}

// Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
func (p *InterSlice) Prepend(elem interface{}) Slice {
	return p.Insert(0, elem)
}

// Reverse returns a new Slice with the order of the elements reversed.
func (p *InterSlice) Reverse() (new Slice) {
	if p.Nil() || p.Len() < 2 {
		return p.Copy()
	}
	return p.Copy().ReverseM()
}

// ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
func (p *InterSlice) ReverseM() Slice {
	l := p.Len()
	if p.Nil() || l == 0 {
		return p
	}
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		p.Swap(i, j)
	}
	return p
}

// S is an alias to ToStringSlice
func (p *InterSlice) S() (slice *StringSlice) {
	return ToStringSlice(p.O())
}

// SG is an alias to ToStringSliceG
func (p *InterSlice) SG() (slice []string) {
	return ToStringSliceG(p.O())
}

// Select creates a new slice with the elements that match the lambda selector.
func (p *InterSlice) Select(sel func(O) bool) (new Slice) {
	// l := p.Len()
	// slice := NewInterSliceV()
	// if p.Nil() || l == 0 {
	// 	return slice
	// }
	// for i := 0; i < l; i++ {
	// 	obj := p.v.Index(i).Interface()
	// 	if sel(obj) {
	// 		slice.Append(obj)
	// 	}
	// }
	return
}

// Set the element at the given index location to the given element. Allows for negative notation.
// Returns a reference to this Slice and swallows any errors.
func (p *InterSlice) Set(i int, elem interface{}) Slice {
	slice, _ := p.SetE(i, elem)
	return slice
}

// SetE the element at the given index location to the given element. Allows for negative notation.
// Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
func (p *InterSlice) SetE(i int, elem interface{}) (slice Slice, err error) {
	// slice = p
	// if p.Nil() {
	// 	return
	// }
	// if i = absIndex(p.Len(), i); i == -1 {
	// 	err = errors.Errorf("slice assignment is out of bounds")
	// 	return
	// }

	// x := reflect.ValueOf(elem)
	// if p.v.Type().Elem() != x.Type() {
	// 	err = errors.Errorf("can't set type '%v' in '%v'", x.Type(), p.v.Type())
	// } else {
	// 	p.v.Index(i).Set(x)
	// }
	return
}

// Shift modifies this Slice to remove the first element and returns the removed element as an Object.
func (p *InterSlice) Shift() (elem *Object) {
	elem = p.First()
	p.DropFirst()
	return
}

// ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
func (p *InterSlice) ShiftN(n int) (new Slice) {
	if p.Nil() {
		return NewInterSliceV()
	}
	if n == 0 {
		return newEmptySlice(p.O())
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
func (p *InterSlice) Slice(indices ...int) Slice {
	// if p.Nil() {
	// 	return NewInterSliceV()
	// }

	// // Handle index manipulation
	// i, j, err := absIndices(p.Len(), indices...)
	// if err != nil {
	// 	return newEmptySlice(p.O())
	// }

	// return NewInterSlice(p.v.Slice(i, j).Interface())
	return nil
}

// Sort returns a new Slice with sorted elements.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) Sort() (new Slice) {
	if p.Nil() || p.Len() < 2 {
		return p.Copy()
	}
	return p.Copy().SortM()
}

// SortM modifies this Slice sorting the elements and returns a reference to this Slice.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) SortM() Slice {
	// if p.Nil() || p.Len() < 2 {
	// 	return p
	// }

	// // Handle supported types
	// slice := NewSlice(p.v.Interface())
	// if !slice.Generic() {
	// 	slice.SortM()
	// 	*p = *NewInterSlice(slice.O())
	// } else {
	// 	panic(fmt.Sprintf("unsupported comparable type '%v'", p.v.Type()))
	// }

	return p
}

// SortReverse returns a new Slice sorting the elements in reverse.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) SortReverse() (new Slice) {
	if p.Nil() || p.Len() < 2 {
		return p.Copy()
	}
	return p.Copy().SortReverseM()
}

// SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
// Supports optimized Slice types or Go types that can be converted into an optimized Slice type.
func (p *InterSlice) SortReverseM() Slice {
	// if p.Nil() || p.Len() < 2 {
	// 	return p
	// }

	// // Handle supported types
	// slice := NewSlice(p.v.Interface())
	// if !slice.Generic() {
	// 	slice.SortReverseM()
	// 	*p = *NewInterSlice(slice.O())
	// } else {
	// 	panic(fmt.Sprintf("unsupported comparable type '%v'", p.v.Type()))
	// }

	return p
}

// Returns a string representation of this Slice, implements the Stringer interface
func (p *InterSlice) String() string {
	// l := p.Len()
	// var builder strings.Builder
	// builder.WriteString("[")
	// for i := 0; i < l; i++ {
	// 	builder.WriteString(fmt.Sprintf("%d", p.v.Index(i).Interface()))
	// 	if i+1 < l {
	// 		builder.WriteString(" ")
	// 	}
	// }
	// builder.WriteString("]")
	//return builder.String()
	return ""
}

// Swap modifies this Slice swapping the indicated elements.
func (p *InterSlice) Swap(i, j int) {
	// l := p.Len()
	// if p.Nil() || l < 2 || i < 0 || j < 0 || i >= l || j >= l {
	// 	return
	// }
	// reflect.Swapper(p.v.Interface())(i, j)
}

// Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *InterSlice) Take(indices ...int) (new Slice) {
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
func (p *InterSlice) TakeW(sel func(O) bool) (new Slice) {
	// l := p.Len()
	// slice := NewInterSliceV()
	// if p.Nil() || l == 0 {
	// 	return slice
	// }
	// for i := 0; i < l; i++ {
	// 	obj := p.v.Index(i).Interface()
	// 	if sel(obj) {
	// 		slice.Append(obj)
	// 		p.DropAt(i)
	// 		l--
	// 		i--
	// 	}
	// }
	return
}

// ToInterSlice converts the given slice to a generic []interface{} slice
func (p *InterSlice) ToInterSlice() (slice []interface{}) {
	return p.G()
}

// ToStringSlice converts the underlying slice into a *StringSlice
func (p *InterSlice) ToStringSlice() (slice *StringSlice) {
	return ToStringSlice(p.O())
}

// ToStringSliceG converts the underlying slice into a []string slice
func (p *InterSlice) ToStringSliceG() (slice []string) {
	return ToStringSliceG(p.O())
}

// Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) Union(slice interface{}) (new Slice) {
	return p.Copy().UnionM(slice)
}

// UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports InterSlice, *InterSlice, Slice and Go slice types
func (p *InterSlice) UnionM(slice interface{}) Slice {
	return p.ConcatM(slice).UniqM()
}

// Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
func (p *InterSlice) Uniq() (new Slice) {
	// l := p.Len()
	// if p.Nil() || l < 2 {
	// 	return p.Copy()
	// }
	// slice := NewInterSliceV()
	// v := reflect.ValueOf(true)
	// typ := reflect.MapOf(p.v.Type().Elem(), v.Type())
	// m := reflect.MakeMap(typ)
	// for i := 0; i < l; i++ {
	// 	k := p.v.Index(i)
	// 	if ok := m.MapIndex(k); ok == (reflect.Value{}) {
	// 		m.SetMapIndex(k, v)
	// 		slice.Append(k.Interface())
	// 	}
	// }
	return
}

// UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
func (p *InterSlice) UniqM() Slice {
	// l := p.Len()
	// if p.Nil() || l < 2 {
	// 	return p
	// }
	// v := reflect.ValueOf(true)
	// typ := reflect.MapOf(p.v.Type().Elem(), v.Type())
	// m := reflect.MakeMap(typ)
	// for i := 0; i < l; i++ {
	// 	k := p.v.Index(i)
	// 	if ok := m.MapIndex(k); ok == (reflect.Value{}) {
	// 		m.SetMapIndex(k, v)
	// 	} else {
	// 		p.DropAt(i)
	// 		l--
	// 		i--
	// 	}
	// }
	return p
}
