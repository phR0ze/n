package n

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
