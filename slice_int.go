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

// At returns the element at the given index location. Allows for negative notation.
func (p *IntSlice) At(i int) (obj *Object) {
	obj = &Object{}
	if p == nil {
		return
	}
	if i = absIndex(len(*p), i); i == -1 {
		return
	}
	obj.o = (*p)[i]
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
func (p *IntSlice) Copy(indices ...int) (result Slice) {
	if p == nil || len(*p) == 0 || len(indices) == 1 {
		result = NewIntSliceV()
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
	result = NewIntSlice(x)
	return
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
