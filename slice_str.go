package n

// import (
// 	"sort"
// 	"strings"

// 	"github.com/pkg/errors"
// )

// // StrSlice implements the Slice interface providing a generic way to work with slice types
// // including convenience methods on par with rapid development languages.
// type StrSlice []Str

// // NewStrSlice creates a new *StrSlice
// func NewStrSlice(slice []Str) *StrSlice {
// 	new := StrSlice(slice)
// 	return &new
// }

// // NewStrSliceV creates a new *StrSlice from the given variadic elements. Always returns
// // at least a reference to an empty StrSlice.
// func NewStrSliceV(elems ...Str) *StrSlice {
// 	var new StrSlice
// 	if len(elems) == 0 {
// 		new = StrSlice([]Str{})
// 	} else {
// 		new = StrSlice(elems)
// 	}
// 	return &new
// }

// // Any tests if this Slice is not empty or optionally if it contains
// // any of the given variadic elements. Incompatible types will return false.
// func (p *StrSlice) Any(elems ...interface{}) bool {
// 	if p == nil || len(*p) == 0 {
// 		return false
// 	}

// 	// Not looking for anything
// 	if len(elems) == 0 {
// 		return true
// 	}

// 	// Looking for something specific returns false if incompatible type
// 	for i := range elems {
// 		if x, ok := elems[i].(Str); ok {
// 			for j := range *p {
// 				if (*p)[j] == x {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// // AnyS tests if this Slice contains any of the given Slice's elements.
// // Incompatible types will return false.
// // Supports StrSlice, *StrSlice, []Str or *[]Str
// func (p *StrSlice) AnyS(slice interface{}) bool {
// 	if p == nil || len(*p) == 0 {
// 		return false
// 	}
// 	var elems []Str
// 	switch x := slice.(type) {
// 	case []Str:
// 		elems = x
// 	case *[]Str:
// 		if x != nil {
// 			elems = *x
// 		}
// 	case StrSlice:
// 		elems = x
// 	case *StrSlice:
// 		if x != nil {
// 			elems = (*x)
// 		}
// 	}
// 	for i := range elems {
// 		for j := range *p {
// 			if (*p)[j] == elems[i] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // AnyW tests if this Slice contains any that match the lambda selector.
// func (p *StrSlice) AnyW(sel func(O) bool) bool {
// 	return p.CountW(sel) != 0
// }

// // Append an element to the end of this Slice and returns a reference to this Slice.
// func (p *StrSlice) Append(elem interface{}) ISlice {
// 	if p == nil {
// 		p = NewStrSliceV()
// 	}
// 	if x, ok := elem.(Str); ok {
// 		*p = append(*p, x)
// 	}
// 	return p
// }

// // AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
// func (p *StrSlice) AppendV(elems ...interface{}) ISlice {
// 	if p == nil {
// 		p = NewStrSliceV()
// 	}
// 	for _, elem := range elems {
// 		p.Append(elem)
// 	}
// 	return p
// }

// // At returns the element at the given index location. Allows for negative notation.
// func (p *StrSlice) At(i int) (elem *Object) {
// 	elem = &Object{}
// 	if p == nil {
// 		return
// 	}
// 	if i = absIndex(len(*p), i); i == -1 {
// 		return
// 	}
// 	elem.o = (*p)[i]
// 	return
// }

// // Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
// func (p *StrSlice) Clear() ISlice {
// 	if p == nil {
// 		p = NewStrSliceV()
// 	} else {
// 		p.Drop()
// 	}
// 	return p
// }

// // Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
// // Supports StrSlice, *StrSlice, []Str or *[]Str
// func (p *StrSlice) Concat(slice interface{}) (new ISlice) {
// 	return p.Copy().ConcatM(slice)
// }

// // ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
// // Supports StrSlice, *StrSlice, []Str or *[]Str
// func (p *StrSlice) ConcatM(slice interface{}) ISlice {
// 	if p == nil {
// 		p = NewStrSliceV()
// 	}
// 	switch x := slice.(type) {
// 	case []Str:
// 		*p = append(*p, x...)
// 	case *[]Str:
// 		if x != nil {
// 			*p = append(*p, (*x)...)
// 		}
// 	case StrSlice:
// 		*p = append(*p, x...)
// 	case *StrSlice:
// 		if x != nil {
// 			*p = append(*p, (*x)...)
// 		}
// 	}
// 	return p
// }

// // Copy returns a new Slice with the indicated range of elements copied from this Slice.
// // Expects nothing, in which case everything is copied, or two indices i and j, in which
// // case positive and negative notation is supported and uses an inclusive behavior such
// // that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of
// // bounds indices will be moved within bounds.
// //
// // An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
// func (p *StrSlice) Copy(indices ...int) (new ISlice) {
// 	if p == nil || len(*p) == 0 {
// 		return NewStrSliceV()
// 	}

// 	// Handle index manipulation
// 	i, j, err := absIndices(len(*p), indices...)
// 	if err != nil {
// 		return NewStrSliceV()
// 	}

// 	// Copy elements over to new Slice
// 	x := make([]Str, j-i, j-i)
// 	copy(x, (*p)[i:j])
// 	return NewStrSlice(x)
// }

// // Count the number of elements in this Slice equal to the given element.
// func (p *StrSlice) Count(elem interface{}) (cnt int) {
// 	if y, ok := elem.(Str); ok {
// 		cnt = p.CountW(func(x O) bool { return ExB(x.(Str) == y) })
// 	}
// 	return
// }

// // CountW counts the number of elements in this Slice that match the lambda selector.
// func (p *StrSlice) CountW(sel func(O) bool) (cnt int) {
// 	if p == nil || len(*p) == 0 {
// 		return
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if sel((*p)[i]) {
// 			cnt++
// 		}
// 	}
// 	return
// }

// // Drop modifies this Slice to delete the indicated range of elements and returns a referece to this Slice.
// // Expects nothing, in which case everything is dropped, or two indices i and j, in which case positive and
// // negative notation is supported and uses an inclusive behavior such that DropAt(0, -1) includes index -1
// // as opposed to Go's exclusive behavior. Out of bounds indices will be moved within bounds.
// func (p *StrSlice) Drop(indices ...int) ISlice {
// 	if p == nil || len(*p) == 0 {
// 		return p
// 	}

// 	// Handle index manipulation
// 	i, j, err := absIndices(len(*p), indices...)
// 	if err != nil {
// 		return p
// 	}

// 	// Execute
// 	n := j - i
// 	if i+n < len(*p) {
// 		*p = append((*p)[:i], (*p)[i+n:]...)
// 	} else {
// 		*p = (*p)[:i]
// 	}
// 	return p
// }

// // DropAt modifies this Slice to delete the element at the given index location. Allows for negative notation.
// // Returns a reference to this Slice.
// func (p *StrSlice) DropAt(i int) ISlice {
// 	return p.Drop(i, i)
// }

// // DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
// func (p *StrSlice) DropFirst() ISlice {
// 	return p.Drop(0, 0)
// }

// // DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
// func (p *StrSlice) DropFirstN(n int) ISlice {
// 	if n == 0 {
// 		return p
// 	}
// 	return p.Drop(0, abs(n)-1)
// }

// // DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
// func (p *StrSlice) DropLast() ISlice {
// 	return p.Drop(-1, -1)
// }

// // DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
// func (p *StrSlice) DropLastN(n int) ISlice {
// 	if n == 0 {
// 		return p
// 	}
// 	return p.Drop(absNeg(n), -1)
// }

// // DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// // The slice is updated instantly when lambda expression is evaluated not after DropW completes.
// func (p *StrSlice) DropW(sel func(O) bool) ISlice {
// 	if p == nil || len(*p) == 0 {
// 		return p
// 	}
// 	l := len(*p)
// 	for i := 0; i < l; i++ {
// 		if sel((*p)[i]) {
// 			p.DropAt(i)
// 			l--
// 			i--
// 		}
// 	}
// 	return p
// }

// // Each calls the given lambda once for each element in this Slice, passing in that element
// // as a parameter. Returns a reference to this Slice
// func (p *StrSlice) Each(action func(O)) ISlice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		action((*p)[i])
// 	}
// 	return p
// }

// // EachE calls the given lambda once for each element in this Slice, passing in that element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *StrSlice) EachE(action func(O) error) (ISlice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if err = action((*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // EachI calls the given lambda once for each element in this Slice, passing in the index and element
// // as a parameter. Returns a reference to this Slice
// func (p *StrSlice) EachI(action func(int, O)) ISlice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		action(i, (*p)[i])
// 	}
// 	return p
// }

// // EachIE calls the given lambda once for each element in this Slice, passing in the index and element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *StrSlice) EachIE(action func(int, O) error) (ISlice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if err = action(i, (*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice
// func (p *StrSlice) EachR(action func(O)) ISlice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		action((*p)[i])
// 	}
// 	return p
// }

// // EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *StrSlice) EachRE(action func(O) error) (ISlice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		if err = action((*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice
// func (p *StrSlice) EachRI(action func(int, O)) ISlice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		action(i, (*p)[i])
// 	}
// 	return p
// }

// // EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *StrSlice) EachRIE(action func(int, O) error) (ISlice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		if err = action(i, (*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // Empty tests if this Slice is empty.
// func (p *StrSlice) Empty() bool {
// 	if p == nil || len(*p) == 0 {
// 		return true
// 	}
// 	return false
// }

// // First returns the first element in this Slice as Object.
// // Object.Nil() == true will be returned when there are no elements in the slice.
// func (p *StrSlice) First() (elem *Object) {
// 	return p.At(0)
// }

// // FirstN returns the first n elements in this slice as a Slice reference to the original.
// // Best effort is used such that as many as can be will be returned up until the request is satisfied.
// func (p *StrSlice) FirstN(n int) ISlice {
// 	if n == 0 {
// 		return NewStrSliceV()
// 	}
// 	return p.Slice(0, abs(n)-1)
// }

// // Generic returns true if the underlying implementation is a RefSlice
// func (p *StrSlice) Generic() bool {
// 	return false
// }

// // Index returns the index of the first element in this Slice where element == elem
// // Returns a -1 if the element was not not found.
// func (p *StrSlice) Index(elem interface{}) (loc int) {
// 	loc = -1
// 	if p == nil || len(*p) == 0 {
// 		return
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if elem == (*p)[i] {
// 			return i
// 		}
// 	}
// 	return
// }

// // Insert modifies this Slice to insert the given element before the element with the given index.
// // Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// // negative index is used, the given element will be inserted after that element, so using an index
// // of -1 will insert the element at the end of the slice. Slice is returned for chaining. Invalid
// // index locations will not change the slice.
// func (p *StrSlice) Insert(i int, elem interface{}) ISlice {
// 	if p == nil || len(*p) == 0 {
// 		return p.Append(elem)
// 	}
// 	j := i
// 	if j = absIndex(len(*p), j); j == -1 {
// 		return p
// 	}
// 	if i < 0 {
// 		j++
// 	}

// 	// Insert the item before j if pos and after j if neg
// 	if x, ok := elem.(Str); ok {
// 		if j == 0 {
// 			*p = append([]Str{x}, (*p)...)
// 		} else if j < len(*p) {
// 			*p = append(*p, x)
// 			copy((*p)[j+1:], (*p)[j:])
// 			(*p)[j] = x
// 		} else {
// 			*p = append(*p, x)
// 		}
// 	}
// 	return p
// }

// // Join converts each element into a string then joins them together using the given separator or comma by default.
// func (p *StrSlice) Join(separator ...string) (str *Object) {
// 	if p == nil || len(*p) == 0 {
// 		str = &Object{""}
// 		return
// 	}
// 	sep := ","
// 	if len(separator) > 0 {
// 		sep = separator[0]
// 	}

// 	var builder strings.Builder
// 	for i := 0; i < len(*p); i++ {
// 		builder.WriteString(Obj((*p)[i]).ToString())
// 		if i+1 < len(*p) {
// 			builder.WriteString(sep)
// 		}
// 	}
// 	str = &Object{builder.String()}
// 	return
// }

// // Last returns the last element in this Slice as an Object.
// // Object.Nil() == true will be returned if there are no elements in the slice.
// func (p *StrSlice) Last() (elem *Object) {
// 	return p.At(-1)
// }

// // LastN returns the last n elements in this Slice as a Slice reference to the original.
// // Best effort is used such that as many as can be will be returned up until the request is satisfied.
// func (p *StrSlice) LastN(n int) ISlice {
// 	if n == 0 {
// 		return NewStrSliceV()
// 	}
// 	return p.Slice(absNeg(n), -1)
// }

// // Len returns the number of elements in this Slice
// func (p *StrSlice) Len() int {
// 	if p == nil {
// 		return 0
// 	}
// 	return len(*p)
// }

// // Less returns true if the element indexed by i is less than the element indexed by j.
// func (p *StrSlice) Less(i, j int) bool {
// 	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
// 		return false
// 	}
// 	return (*p)[i] < (*p)[j]
// }

// // Nil tests if this Slice is nil
// func (p *StrSlice) Nil() bool {
// 	if p == nil {
// 		return true
// 	}
// 	return false
// }

// // O returns the underlying data structure as is
// func (p *StrSlice) O() interface{} {
// 	return []Str(*p)
// }

// // Pair simply returns the first and second Slice elements as Objects
// func (p *StrSlice) Pair() (first, second *Object) {
// 	first, second = &Object{}, &Object{}
// 	if p == nil {
// 		return
// 	}
// 	if len(*p) > 0 {
// 		first = p.At(0)
// 	}
// 	if len(*p) > 1 {
// 		second = p.At(1)
// 	}
// 	return
// }

// // Pop modifies this Slice to remove the last element and returns the removed element as an Object.
// func (p *StrSlice) Pop() (elem *Object) {
// 	elem = p.Last()
// 	p.DropLast()
// 	return
// }

// // PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
// func (p *StrSlice) PopN(n int) (new ISlice) {
// 	if n == 0 {
// 		return NewStrSliceV()
// 	}
// 	new = p.Copy(absNeg(n), -1)
// 	p.DropLastN(n)
// 	return
// }

// // Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
// func (p *StrSlice) Prepend(elem interface{}) ISlice {
// 	return p.Insert(0, elem)
// }

// // Reverse returns a new Slice with the order of the elements reversed.
// func (p *StrSlice) Reverse() (new ISlice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	return p.Copy().ReverseM()
// }

// // ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
// func (p *StrSlice) ReverseM() ISlice {
// 	if p == nil || len(*p) == 0 {
// 		return p
// 	}
// 	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
// 		p.Swap(i, j)
// 	}
// 	return p
// }

// // Select creates a new slice with the elements that match the lambda selector.
// func (p *StrSlice) Select(sel func(O) bool) (new ISlice) {
// 	slice := NewStrSliceV()
// 	if p == nil || len(*p) == 0 {
// 		return slice
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if sel((*p)[i]) {
// 			*slice = append(*slice, (*p)[i])
// 		}
// 	}
// 	return slice
// }

// // Set the element at the given index location to the given element. Allows for negative notation.
// // Returns a reference to this Slice and swallows any errors.
// func (p *StrSlice) Set(i int, elem interface{}) ISlice {
// 	slice, _ := p.SetE(i, elem)
// 	return slice
// }

// // SetE the element at the given index location to the given element. Allows for negative notation.
// // Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
// func (p *StrSlice) SetE(i int, elem interface{}) (ISlice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	if i = absIndex(len(*p), i); i == -1 {
// 		err = errors.Errorf("slice assignment is out of bounds")
// 		return p, err
// 	}

// 	if x, ok := elem.(Str); ok {
// 		(*p)[i] = x
// 	} else {
// 		err = errors.Errorf("can't set type '%T' in '%T'", elem, p)
// 	}
// 	return p, err
// }

// // Shift modifies this Slice to remove the first element and returns the removed element as an Object.
// func (p *StrSlice) Shift() (elem *Object) {
// 	elem = p.First()
// 	p.DropFirst()
// 	return
// }

// // ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
// func (p *StrSlice) ShiftN(n int) (new ISlice) {
// 	if n == 0 {
// 		return NewStrSliceV()
// 	}
// 	new = p.Copy(0, abs(n)-1)
// 	p.DropFirstN(n)
// 	return
// }

// // Single reports true if there is only one element in this Slice.
// func (p *StrSlice) Single() bool {
// 	return len(*p) == 1
// }

// // Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
// // Expects nothing, in which case everything is included, or two indices i and j, in which case an inclusive behavior
// // is used such that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices will
// // be moved within bounds.
// //
// // An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
// //
// // e.g. NewStrSliceV(1,2,3).Slice(0, -1) == [1,2,3] && NewStrSliceV(1,2,3).Slice(1,2) == [2,3]
// func (p *StrSlice) Slice(indices ...int) ISlice {
// 	if p == nil || len(*p) == 0 {
// 		return NewStrSliceV()
// 	}

// 	// Handle index manipulation
// 	i, j, err := absIndices(len(*p), indices...)
// 	if err != nil {
// 		return NewStrSliceV()
// 	}

// 	return NewStrSlice((*p)[i:j])
// }

// // Sort returns a new Slice with sorted elements.
// func (p *StrSlice) Sort() (new ISlice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	return p.Copy().SortM()
// }

// // SortM modifies this Slice sorting the elements and returns a reference to this Slice.
// func (p *StrSlice) SortM() ISlice {
// 	if p == nil || len(*p) < 2 {
// 		return p
// 	}
// 	sort.Sort(p)
// 	return p
// }

// // SortReverse returns a new Slice sorting the elements in reverse.
// func (p *StrSlice) SortReverse() (new ISlice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	return p.Copy().SortReverseM()
// }

// // SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
// func (p *StrSlice) SortReverseM() ISlice {
// 	if p == nil || len(*p) < 2 {
// 		return p
// 	}
// 	sort.Sort(sort.Reverse(p))
// 	return p
// }

// // String returns a string representation of this Slice, implements the Stringer interface
// func (p *StrSlice) String() string {
// 	var builder strings.Builder
// 	builder.WriteString("[")
// 	if p != nil {
// 		for i := 0; i < len(*p); i++ {
// 			builder.WriteString(string((*p)[i]))
// 			if i+1 < len(*p) {
// 				builder.WriteString(" ")
// 			}
// 		}
// 	}
// 	builder.WriteString("]")
// 	return builder.String()
// }

// // Swap modifies this Slice swapping the indicated elements.
// func (p *StrSlice) Swap(i, j int) {
// 	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
// 		return
// 	}
// 	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
// }

// // Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// // Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// // notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// // exclusive behavior. Out of bounds indices will be moved within bounds.
// func (p *StrSlice) Take(indices ...int) (new ISlice) {
// 	new = p.Copy(indices...)
// 	p.Drop(indices...)
// 	return
// }

// // TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
// // Allows for negative notation.
// func (p *StrSlice) TakeAt(i int) (elem *Object) {
// 	elem = p.At(i)
// 	p.DropAt(i)
// 	return
// }

// // TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
// func (p *StrSlice) TakeW(sel func(O) bool) (new ISlice) {
// 	slice := NewStrSliceV()
// 	if p == nil || len(*p) == 0 {
// 		return slice
// 	}
// 	l := len(*p)
// 	for i := 0; i < l; i++ {
// 		if sel((*p)[i]) {
// 			*slice = append(*slice, (*p)[i])
// 			p.DropAt(i)
// 			l--
// 			i--
// 		}
// 	}
// 	return slice
// }

// // Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// // Supports StrSlice, *StrSlice, []Str or *[]Str
// func (p *StrSlice) Union(slice interface{}) (new ISlice) {
// 	return p.Copy().UnionM(slice)
// }

// // UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// // Supports StrSlice, *StrSlice, []Str or *[]Str
// func (p *StrSlice) UnionM(slice interface{}) ISlice {
// 	return p.ConcatM(slice).UniqM()
// }

// // Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// // Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
// func (p *StrSlice) Uniq() (new ISlice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	m := NewStringMapBool()
// 	slice := NewStrSliceV()
// 	for i := 0; i < len(*p); i++ {
// 		if ok := m.Set((*p)[i], true); ok {
// 			slice.Append((*p)[i])
// 		}
// 	}
// 	return slice
// }

// // UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// // Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
// func (p *StrSlice) UniqM() ISlice {
// 	if p == nil || len(*p) < 2 {
// 		return p
// 	}
// 	m := NewStringMapBool()
// 	l := len(*p)
// 	for i := 0; i < l; i++ {
// 		if ok := m.Set((*p)[i], true); !ok {
// 			p.DropAt(i)
// 			l--
// 			i--
// 		}
// 	}
// 	return p
// }
