package n

import (
	"regexp"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

var (
	// ReGraphicalOnly is a regex to filter on graphical runes only
	ReGraphicalOnly = regexp.MustCompile(`[^[:graph:]]+`)
)

// Str wraps the Go []rune and implements the Slice interface providing
// convenience methods on par with other rapid development languages.
type Str []rune

// A is an alias to NewStr for brevity
func A(obj interface{}) *Str {
	return NewStr(obj)
}

// NewStr creates a new *Str which will never be nil
// Supports: Str *Str, string *string, []byte *[]byte, rune *rune, []rune *[]rune, []string *[]string ...
func NewStr(obj interface{}) *Str {
	new := Str{}
	o := Indirect(obj)
	switch x := o.(type) {
	case nil:
	case Str:
		new = x
	case string:
		new = Str(x)
	case byte:
		new = Str(string(x))
	case []byte:
		new = Str(string(x))
	case rune:
		new = append(new, x)
	case []rune:
		new = Str(x)
	case []string:
		for i := range x {
			new = append(new, []rune(x[i])...)
		}
	default:
		new = Str(ToString(o))
	}
	return &new
}

// NewStrV creates a new *Str from the given variadic elements. Returned *Str
// will never be nil.
// Supports: Str *Str, string *string, []byte *[]byte, rune *rune, []rune *[]rune, []string *[]string ...
func NewStrV(elems ...interface{}) *Str {
	new := Str{}
	for i := range elems {
		new.Append(elems[i])
	}
	return &new
}

// A exports the Str as a Go string
func (p *Str) A() string {
	if p == nil {
		return ""
	}
	return string(*p)
}

// All tests if this Slice contains all of the given variadic elements.
// Incompatible types will return false.
// Supports: Str *Str, string *string, rune []rune as a string, []byte as string
func (p *Str) All(elems ...interface{}) bool {
	if p == nil || len(*p) == 0 || len(elems) == 0 {
		return false
	}
	str := p.A()
	result := true
	for i := range elems {
		if !strings.Contains(str, NewStr(elems[i]).A()) {
			return false
		}
	}
	return result
}

// AllS tests if this Slice contains all of the given Slice's elements.
// Incompatible types will return false.
// Supports: []Str *[]Str, []string *[]string, StringSlice *StringSlice, []rune *[]rune as chars, []byte *[]byte as chars
func (p *Str) AllS(slice interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	str := p.A()
	result := true
	o := Indirect(slice)
	switch x := o.(type) {
	case []Str:
		SetOnTrue(&result, false, len(x) == 0)
		for i := range x {
			if !strings.Contains(str, string(x[i])) {
				return false
			}
		}
	case []string:
		SetOnTrue(&result, false, len(x) == 0)
		for i := range x {
			if !strings.Contains(str, x[i]) {
				return false
			}
		}
	case StringSlice:
		SetOnTrue(&result, false, len(x) == 0)
		for i := 0; i < x.Len(); i++ {
			if !strings.Contains(str, x.At(i).A()) {
				return false
			}
		}
	case []byte:
		SetOnTrue(&result, false, len(x) == 0)
		for i := range x {
			if !strings.ContainsRune(str, rune(x[i])) {
				return false
			}
		}
	case []rune:
		SetOnTrue(&result, false, len(x) == 0)
		for i := range x {
			if !strings.ContainsRune(str, x[i]) {
				return false
			}
		}
	default:
		result = false
	}
	return result
}

// Any tests if this Slice is not empty or optionally if it contains
// any of the given variadic elements. Incompatible types will return false.
// Supports: Str *Str, string *string, rune []rune as a string, []byte as string...
func (p *Str) Any(elems ...interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}

	// Not looking for anything
	if len(elems) == 0 {
		return true
	}

	// Looking for something specific returns false if incompatible type
	str := p.A()
	for i := range elems {
		if strings.Contains(str, NewStr(elems[i]).A()) {
			return true
		}
	}
	return false
}

// AnyS tests if this Slice contains any of the given Slice's elements.
// Incompatible types will return false.
// Supports: []Str *[]Str, []string *[]string, StringSlice *StringSlice, []rune *[]rune as chars, []byte *[]byte as chars
func (p *Str) AnyS(slice interface{}) bool {
	if p == nil || len(*p) == 0 {
		return false
	}
	str := p.A()
	o := Indirect(slice)
	switch x := o.(type) {
	case []Str:
		for i := range x {
			if strings.Contains(str, string(x[i])) {
				return true
			}
		}
	case []string:
		for i := range x {
			if strings.Contains(str, x[i]) {
				return true
			}
		}
	case StringSlice:
		for i := 0; i < x.Len(); i++ {
			if strings.Contains(str, x.At(i).A()) {
				return true
			}
		}
	case []byte:
		for i := range x {
			if strings.ContainsRune(str, rune(x[i])) {
				return true
			}
		}
	case []rune:
		for i := range x {
			if strings.ContainsRune(str, x[i]) {
				return true
			}
		}
	}
	return false
}

// AnyW tests if this Slice contains any that match the lambda selector.
func (p *Str) AnyW(sel func(O) bool) bool {
	return p.CountW(sel) != 0
}

// Append an element to the end of this Slice and returns a reference to this Slice.
// Supports: Str *Str, string *string, []byte *[]byte, rune *rune, []rune *[]rune, []string *[]string ...
func (p *Str) Append(elem interface{}) Slice {
	if p == nil {
		p = NewStrV()
	}
	*p = append(*p, (*NewStr(elem))...)
	return p
}

// AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
// Supports: Str *Str, string *string, []byte *[]byte, rune *rune, []rune *[]rune, []string *[]string ...
func (p *Str) AppendV(elems ...interface{}) Slice {
	if p == nil {
		p = NewStrV()
	}
	for _, elem := range elems {
		*p = append(*p, (*NewStr(elem))...)
	}
	return p
}

// Ascii converts the string to pure ASCII
func (p *Str) Ascii() *Str {
	if p == nil {
		return A("")
	}
	return A(ReGraphicalOnly.ReplaceAllString(string(*p), " "))
}

// AsciiA converts the string to pure ASCII
func (p *Str) AsciiA() string {
	if p == nil {
		return ""
	}
	return ReGraphicalOnly.ReplaceAllString(string(*p), " ")
}

// AsciiOnly checks to see if this is an ASCII only string
func (p *Str) AsciiOnly() bool {
	if p == nil {
		return true
	}
	return len(*p) == len(ReGraphicalOnly.ReplaceAllString(string(*p), ""))
}

// At returns the element at the given index location. Allows for negative notation.
func (p *Str) At(i int) (elem *Object) {
	elem = &Object{}
	if p == nil {
		return
	}
	if i = absIndex(len(*p), i); i == -1 {
		return
	}
	elem.o = Char((*p)[i])
	return
}

// B exports the Str as a Go []byte
func (p *Str) B() []byte {
	if p == nil {
		return []byte{}
	}
	return []byte(string(*p))
}

// C exports the Str as a Char
func (p *Str) C() *Char {
	return NewChar(p)
}

// Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
func (p *Str) Clear() Slice {
	if p == nil {
		p = NewStrV()
	} else {
		p.Drop()
	}
	return p
}

// Concat returns a new Slice by appending the given Slice to this Slice.
// Supports Str, *Str, []string or *[]string
func (p *Str) Concat(slice interface{}) (new Slice) {
	return p.Copy().ConcatM(slice)
}

// ConcatA calls Concat and returns it as a string for brevity
func (p *Str) ConcatA(str interface{}) string {
	return p.Concat(str).A()
}

// ConcatM modifies this Slice by appending the given and returns a reference to this Slice.
// Supports: Str *Str, string *string, []byte *[]byte, rune *rune, []rune *[]rune, []string *[]string ...
func (p *Str) ConcatM(slice interface{}) Slice {
	if p == nil {
		p = NewStrV()
	}
	*p = append(*p, (*NewStr(slice))...)
	return p
}

// Contains checks if the given str is contained in this Str.
//
// Pass through for strings.Contains
func (p *Str) Contains(str string) bool {
	if p == nil {
		return false
	}
	return strings.Contains(string(*p), str)
}

// ContainsAny checks if any of the given chars exist in this Str.
//
// Pass through for strings.ContainsAny
func (p *Str) ContainsAny(chars string) bool {
	if p == nil {
		return false
	}
	return strings.ContainsAny(string(*p), chars)
}

// ContainsRune checks if the given rune exists in this Str.
//
// Pass through for strings.ContainsRune
func (p *Str) ContainsRune(r rune) bool {
	if p == nil {
		return false
	}
	return strings.ContainsRune(string(*p), r)
}

// Copy returns a new Slice with the indicated range of elements copied from this Slice.
// Expects nothing, in which case everything is copied, or two indices i and j, in which
// case positive and negative notation is supported and uses an inclusive behavior such
// that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of
// bounds indices will be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
func (p *Str) Copy(indices ...int) (new Slice) {
	if p == nil || len(*p) == 0 {
		return NewStrV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewStrV()
	}

	// Copy elements over to new Slice
	x := make([]rune, j-i, j-i)
	copy(x, []rune(*p)[i:j])
	return NewStr(x)
}

// CopyA calls Copy and returns it as a string for brevity
func (p *Str) CopyA(indices ...int) (new string) {
	return p.Copy(indices...).A()
}

// Count counts the number of non-overlapping instances of substr in this string.
// Supports: Str, *Str, string, *string, rune, []rune as a string, []byte as string.
//
// Pass through for strings.Count
func (p *Str) Count(elem interface{}) (cnt int) {
	if p == nil || len(*p) == 0 || elem == nil {
		return
	}
	return strings.Count(string(*p), string(*NewStr(elem)))
}

// CountW counts the number of elements in this Slice that match the lambda selector.
func (p *Str) CountW(sel func(O) bool) (cnt int) {
	if p == nil || len(*p) == 0 {
		return
	}
	for i := range *p {
		if sel(Char((*p)[i])) {
			cnt++
		}
	}
	return
}

// Drop modifies this Slice to delete the indicated range of elements and returns a referece to this Slice.
// Expects nothing, in which case everything is dropped, or two indices i and j, in which case positive and
// negative notation is supported and uses an inclusive behavior such that DropAt(0, -1) includes index -1
// as opposed to Go's exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *Str) Drop(indices ...int) Slice {
	if p == nil || len(*p) == 0 {
		return NewStrV()
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
func (p *Str) DropAt(i int) Slice {
	return p.Drop(i, i)
}

// DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
func (p *Str) DropFirst() Slice {
	return p.Drop(0, 0)
}

// DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
func (p *Str) DropFirstN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(0, abs(n)-1)
}

// DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
func (p *Str) DropLast() Slice {
	return p.Drop(-1, -1)
}

// DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
func (p *Str) DropLastN(n int) Slice {
	if n == 0 {
		return p
	}
	return p.Drop(absNeg(n), -1)
}

// DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// The slice is updated instantly when lambda expression is evaluated not after DropW completes.
func (p *Str) DropW(sel func(O) bool) Slice {
	if p == nil || len(*p) == 0 {
		return p
	}
	l := len(*p)
	for i := 0; i < l; i++ {
		if sel(Char((*p)[i])) {
			p.DropAt(i)
			l--
			i--
		}
	}
	return p
}

// Each calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *Str) Each(action func(O)) Slice {
	if p == nil {
		return p
	}
	for i := range *p {
		action(Char((*p)[i]))
	}
	return p
}

// EachE calls the given lambda once for each element in this Slice, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *Str) EachE(action func(O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := range *p {
		if err = action(Char((*p)[i])); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachI calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice
func (p *Str) EachI(action func(int, O)) Slice {
	if p == nil {
		return p
	}
	for i := range *p {
		action(i, Char((*p)[i]))
	}
	return p
}

// EachIE calls the given lambda once for each element in this Slice, passing in the index and element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *Str) EachIE(action func(int, O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := range *p {
		if err = action(i, Char((*p)[i])); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *Str) EachR(action func(O)) Slice {
	if p == nil {
		return p
	}
	for i := len(*p) - 1; i >= 0; i-- {
		action(Char((*p)[i]))
	}
	return p
}

// EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *Str) EachRE(action func(O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := len(*p) - 1; i >= 0; i-- {
		if err = action(Char((*p)[i])); err != nil {
			return p, err
		}
	}
	return p, err
}

// EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice
func (p *Str) EachRI(action func(int, O)) Slice {
	if p == nil {
		return p
	}
	for i := len(*p) - 1; i >= 0; i-- {
		action(i, Char((*p)[i]))
	}
	return p
}

// EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
// as a parameter. Returns a reference to this Slice and any error from the lambda.
func (p *Str) EachRIE(action func(int, O) error) (Slice, error) {
	var err error
	if p == nil {
		return p, err
	}
	for i := len(*p) - 1; i >= 0; i-- {
		if err = action(i, Char((*p)[i])); err != nil {
			return p, err
		}
	}
	return p, err
}

// Empty tests if this Slice is empty.
func (p *Str) Empty() bool {
	if p == nil || len(*p) == 0 {
		return true
	}
	return false
}

// First returns the first element in this Slice as Object.
// Object.Nil() == true will be returned when there are no elements in the slice.
func (p *Str) First() (elem *Object) {
	return p.At(0)
}

// FirstN returns the first n elements in this slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *Str) FirstN(n int) Slice {
	if n == 0 {
		return NewStrV()
	}
	return p.Slice(0, abs(n)-1)
}

// G returns the underlying data structure as a builtin Go type
func (p *Str) G() string {
	return p.O().(string)
}

// Generic returns true if the underlying implementation is a RefSlice
func (p *Str) Generic() bool {
	return false
}

// Index returns the index of the first element in this Slice where element == elem
// Returns a -1 if the element was not not found.
func (p *Str) Index(elem interface{}) (loc int) {
	loc = -1
	if p == nil || len(*p) == 0 {
		return
	}
	for i := range *p {
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
func (p *Str) Insert(i int, elem interface{}) Slice {
	if p == nil || len(*p) == 0 {
		return p.Append(elem)
	}

	// // Insert the item before j if pos and after j if neg
	// j := i
	// if j = absIndex(len(*p), j); j == -1 {
	// 	return p
	// }
	// if i < 0 {
	// 	j++
	// }
	// if x, ok := elem.(string); ok {
	// 	if j == 0 {
	// 		*p = append([]string{x}, (*p)...)
	// 	} else if j < len(*p) {
	// 		*p = append(*p, x)
	// 		copy((*p)[j+1:], (*p)[j:])
	// 		(*p)[j] = x
	// 	} else {
	// 		*p = append(*p, x)
	// 	}
	// }
	return p
}

// InsertS modifies this Slice to insert the given elements before the element with the given index.
// Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// negative index is used, the given element will be inserted after that element, so using an index
// of -1 will insert the element at the end of the slice. Slice is returned for chaining. Invalid
// index locations will not change the slice.
func (p *Str) InsertS(i int, slice interface{}) Slice {
	if p == nil || len(*p) == 0 {
		return p.ConcatM(slice)
	}

	// // Insert the item before j if pos and after j if neg
	// j := i
	// if j = absIndex(len(*p), j); j == -1 {
	// 	return p
	// }
	// if i < 0 {
	// 	j++
	// }
	// if x, ok := elem.(string); ok {
	// 	if j == 0 {
	// 		*p = append([]string{x}, (*p)...)
	// 	} else if j < len(*p) {
	// 		*p = append(*p, x)
	// 		copy((*p)[j+1:], (*p)[j:])
	// 		(*p)[j] = x
	// 	} else {
	// 		*p = append(*p, x)
	// 	}
	// }
	return p
}

// Join converts each element into a string then joins them together using the given separator or comma by default.
func (p *Str) Join(separator ...string) (str *Object) {
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
func (p *Str) Last() (elem *Object) {
	return p.At(-1)
}

// LastN returns the last n elements in this Slice as a Slice reference to the original.
// Best effort is used such that as many as can be will be returned up until the request is satisfied.
func (p *Str) LastN(n int) Slice {
	if n == 0 {
		return NewStrV()
	}
	return p.Slice(absNeg(n), -1)
}

// Len returns the number of elements in this Slice
func (p *Str) Len() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// Less returns true if the element indexed by i is less than the element indexed by j.
func (p *Str) Less(i, j int) bool {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return false
	}
	return (*p)[i] < (*p)[j]
}

// Nil tests if this Slice is nil
func (p *Str) Nil() bool {
	if p == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *Str) O() interface{} {
	return string(*p)
}

// Pair simply returns the first and second Slice elements as Objects
func (p *Str) Pair() (first, second *Object) {
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
func (p *Str) Pop() (elem *Object) {
	elem = p.Last()
	p.DropLast()
	return
}

// PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
func (p *Str) PopN(n int) (new Slice) {
	if n == 0 {
		return NewStrV()
	}
	new = p.Copy(absNeg(n), -1)
	p.DropLastN(n)
	return
}

// Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
func (p *Str) Prepend(elem interface{}) Slice {
	return p.Insert(0, elem)
}

// R exports the Str as a rune
func (p *Str) R() rune {
	return NewChar(p).R()
}

// Reverse returns a new Slice with the order of the elements reversed.
func (p *Str) Reverse() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().ReverseM()
}

// ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
func (p *Str) ReverseM() Slice {
	if p == nil || len(*p) == 0 {
		return p
	}
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		p.Swap(i, j)
	}
	return p
}

// Select creates a new slice with the elements that match the lambda selector.
func (p *Str) Select(sel func(O) bool) (new Slice) {
	slice := NewStrV()
	if p == nil || len(*p) == 0 {
		return slice
	}
	// for i := 0; i < len(*p); i++ {
	// 	if sel((*p)[i]) {
	// 		*slice = append(*slice, (*p)[i])
	// 	}
	// }
	return slice
}

// Set the element at the given index location to the given element. Allows for negative notation.
// Returns a reference to this Slice and swallows any errors.
func (p *Str) Set(i int, elem interface{}) Slice {
	slice, _ := p.SetE(i, elem)
	return slice
}

// SetE the element at the given index location to the given element. Allows for negative notation.
// Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
func (p *Str) SetE(i int, elem interface{}) (Slice, error) {
	var err error
	if p == nil {
		return NewStrV(), err
	}
	if i = absIndex(len(*p), i); i == -1 {
		err = errors.Errorf("slice assignment is out of bounds")
		return p, err
	}

	y := NewStr(elem)
	if len(*y) > 0 {
		(*p)[i] = (*y)[0]
	}
	return p, err
}

// Shift modifies this Slice to remove the first element and returns the removed element as an Object.
func (p *Str) Shift() (elem *Object) {
	elem = p.First()
	p.DropFirst()
	return
}

// ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
func (p *Str) ShiftN(n int) (new Slice) {
	if n == 0 {
		return NewStrV()
	}
	new = p.Copy(0, abs(n)-1)
	p.DropFirstN(n)
	return
}

// Single reports true if there is only one element in this Slice.
func (p *Str) Single() bool {
	return p.Len() == 1
}

// Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
// Expects nothing, in which case everything is included, or two indices i and j, in which case an inclusive behavior
// is used such that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices will
// be moved within bounds.
//
// An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
//
// e.g. NewStrV(1,2,3).Slice(0, -1) == [1,2,3] && NewStrV(1,2,3).Slice(1,2) == [2,3]
func (p *Str) Slice(indices ...int) Slice {
	if p == nil || len(*p) == 0 {
		return NewStrV()
	}

	// Handle index manipulation
	i, j, err := absIndices(len(*p), indices...)
	if err != nil {
		return NewStrV()
	}

	return NewStr((*p)[i:j])
}

// Sort returns a new Slice with sorted elements.
func (p *Str) Sort() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortM()
}

// SortM modifies this Slice sorting the elements and returns a reference to this Slice.
func (p *Str) SortM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(p)
	return p
}

// SortReverse returns a new Slice sorting the elements in reverse.
func (p *Str) SortReverse() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	return p.Copy().SortReverseM()
}

// SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
func (p *Str) SortReverseM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	sort.Sort(sort.Reverse(p))
	return p
}

// String returns a string representation of this Slice, implements the Stringer interface
func (p *Str) String() string {
	if p == nil {
		return ""
	}
	return string(*p)
}

// Swap modifies this Slice swapping the indicated elements.
func (p *Str) Swap(i, j int) {
	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
		return
	}
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

// Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// exclusive behavior. Out of bounds indices will be moved within bounds.
func (p *Str) Take(indices ...int) (new Slice) {
	new = p.Copy(indices...)
	p.Drop(indices...)
	return
}

// TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
// Allows for negative notation.
func (p *Str) TakeAt(i int) (elem *Object) {
	elem = p.At(i)
	p.DropAt(i)
	return
}

// TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
func (p *Str) TakeW(sel func(O) bool) (new Slice) {
	slice := NewStrV()
	if p == nil || len(*p) == 0 {
		return slice
	}
	l := len(*p)
	for i := 0; i < l; i++ {
		if sel(Char((*p)[i])) {
			*slice = append(*slice, (*p)[i])
			p.DropAt(i)
			l--
			i--
		}
	}
	return slice
}

// Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports Str, *Str, []string or *[]string
func (p *Str) Union(slice interface{}) (new Slice) {
	return p.Copy().UnionM(slice)
}

// UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// Supports Str, *Str, []string or *[]string
func (p *Str) UnionM(slice interface{}) Slice {
	return p.ConcatM(slice).UniqM()
}

// Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
func (p *Str) Uniq() (new Slice) {
	if p == nil || len(*p) < 2 {
		return p.Copy()
	}
	m := NewRuneMapBool()
	slice := NewStrV()
	for i := range *p {
		if ok := m.Set((*p)[i], true); ok {
			slice.Append((*p)[i])
		}
	}
	return slice
}

// UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
func (p *Str) UniqM() Slice {
	if p == nil || len(*p) < 2 {
		return p
	}
	m := NewStringMapBool()
	l := len(*p)
	for i := range *p {
		if ok := m.Set((*p)[i], true); !ok {
			p.DropAt(i)
			l--
			i--
		}
	}
	return p
}

// // // // Empty returns true if the pointer is nil, string is empty or whitespace only
// // // func (q *QStr) Empty() bool {
// // // 	return q == nil || q.TrimSpace().v == ""
// // // }

// // // // HasAnyPrefix checks if the string has any of the given prefixes
// // // func (q *QStr) HasAnyPrefix(prefixes ...string) bool {
// // // 	for i := range prefixes {
// // // 		if strings.HasPrefix(q.v, prefixes[i]) {
// // // 			return true
// // // 		}
// // // 	}
// // // 	return false
// // // }

// // // // HasAnySuffix checks if the string has any of the given suffixes
// // // func (q *QStr) HasAnySuffix(suffixes ...string) bool {
// // // 	for i := range suffixes {
// // // 		if strings.HasSuffix(q.v, suffixes[i]) {
// // // 			return true
// // // 		}
// // // 	}
// // // 	return false
// // // }

// // // // HasPrefix checks if the string has the given prefix
// // // func (q *QStr) HasPrefix(prefix string) bool {
// // // 	return strings.HasPrefix(q.v, prefix)
// // // }

// // // // HasSuffix checks if the string has the given suffix
// // // func (q *QStr) HasSuffix(suffix string) bool {
// // // 	return strings.HasSuffix(q.v, suffix)
// // // }

// // // // Len returns the length of the string
// // // func (q *QStr) Len() int {
// // // 	return len(q.v)
// // // }

// // // // Nil tests if the numerable is nil.
// // // // Implements Numerable interface
// // // func (q *QStr) Nil() bool {
// // // 	if q == nil {
// // // 		return true
// // // 	}
// // // 	return false
// // // }

// // // O returns the underlying data structure as is
// // func (p *Str) O() interface{} {
// // 	if p == nil {
// // 		return ""
// // 	}
// // 	return string(*p)
// // }

// // // // Replace wraps strings.Replace and allows for chaining and defaults
// // // func (q *QStr) Replace(old, new string, ns ...int) *QStr {
// // // 	n := -1
// // // 	if len(ns) > 0 {
// // // 		n = ns[0]
// // // 	}
// // // 	return A(strings.Replace(q.v, old, new, n))
// // // }

// // // // SpaceLeft returns leading whitespace
// // // func (q *QStr) SpaceLeft() *QStr {
// // // 	spaces := []rune{}
// // // 	for _, r := range q.v {
// // // 		if unicode.IsSpace(r) {
// // // 			spaces = append(spaces, r)
// // // 		} else {
// // // 			break
// // // 		}
// // // 	}
// // // 	return A(spaces)
// // // }

// // // // SpaceRight returns trailing whitespace
// // // func (q *QStr) SpaceRight() *QStr {
// // // 	spaces := []rune{}
// // // 	for i := len(q.v) - 1; i > 0; i-- {
// // // 		if unicode.IsSpace(rune(q.v[i])) {
// // // 			spaces = append(spaces, rune(q.v[i]))
// // // 		} else {
// // // 			break
// // // 		}
// // // 	}
// // // 	return A(spaces)
// // // }

// // // // // Split creates a new NSlice from the split string.
// // // // // Optional 'delim' defaults to space allows for changing the split delimiter.
// // // // func (q *QStr) Split(delim ...string) *NSlice {
// // // // 	_delim := " "
// // // // 	if len(delim) > 0 {
// // // // 		_delim = delim[0]
// // // // 	}
// // // // 	return S(strings.Split(q.v, _delim))
// // // // }

// // // // SplitOn splits the string on the first occurance of the delim.
// // // // The delim is included in the first component.
// // // func (q *QStr) SplitOn(delim string) (first, second string) {
// // // 	// if q.v != "" {
// // // 	// 	s := q.Split(delim)
// // // 	// 	if s.Len() > 0 {
// // // 	// 		first = s.First().A()
// // // 	// 		if strings.Contains(q.v, delim) {
// // // 	// 			first += delim
// // // 	// 		}
// // // 	// 	}
// // // 	// 	if s.Len() > 1 {
// // // 	// 		second = s.Slice(1, -1).Join(delim).A()
// // // 	// 	}
// // // 	// }
// // // 	return
// // // }

// // // // // TrimPrefix trims the given prefix off the string
// // // // func (q *strN) TrimPrefix(prefix string) *strN {
// // // // 	return A(strings.TrimPrefix(q.v, prefix))
// // // // }

// // // // TrimSpace pass through to strings.TrimSpace
// // // func (q *QStr) TrimSpace() *QStr {
// // // 	return A(strings.TrimSpace(q.v))
// // // }

// // // // // TrimSpaceLeft trims leading whitespace
// // // // func (q *strN) TrimSpaceLeft() *strN {
// // // // 	return A(strings.TrimLeftFunc(q.v, unicode.IsSpace))
// // // // }

// // // // // TrimSpaceRight trims trailing whitespace
// // // // func (q *strN) TrimSpaceRight() *strN {
// // // // 	return A(strings.TrimRightFunc(q.v, unicode.IsSpace))
// // // // }

// // // // // TrimSuffix trims the given suffix off the string
// // // // func (q *strN) TrimSuffix(suffix string) *strN {
// // // // 	return A(strings.TrimSuffix(q.v, suffix))
// // // // }

// // // // // YamlType converts the given string into a type expected in Yaml.
// // // // // Quotes signifies a string.
// // // // // No quotes signifies an int.
// // // // // true or false signifies a bool.
// // // // func (q *strN) YamlType() interface{} {
// // // // 	if q.HasAnyPrefix("\"", "'") && q.HasAnySuffix("\"", "'") {
// // // // 		return q.v[1 : len(q.v)-1]
// // // // 	} else if q.v == "true" || q.v == "false" {
// // // // 		if b, err := strconv.ParseBool(q.v); err == nil {
// // // // 			return b
// // // // 		}
// // // // 	} else if f, err := strconv.ParseFloat(q.v, 32); err == nil {
// // // // 		return f
// // // // 	}
// // // // 	return q.v
// // // // }
