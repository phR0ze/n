// Package nub provides helper interfaces and functions reminiscent of C#'s IEnumerable methods
package nub

//--------------------------------------------------------------------------------------------------
// IntSlice implementation
//--------------------------------------------------------------------------------------------------

// IntSliceImpl ...
type IntSliceImpl struct {
	Raw []int
}

// IntSlice ...
func IntSlice(slice []int) *IntSliceImpl {
	return &IntSliceImpl{Raw: slice}
}

// Contains provides a reusable to check if the given target exists in the enumerable
func (slice *IntSliceImpl) Contains(target int) bool {
	for i := range slice.Raw {
		if slice.Raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny provides a reusable to check if the given target exists in the enumerable
func (slice *IntSliceImpl) ContainsAny(targets []int) bool {
	for i := range targets {
		for j := range slice.Raw {
			if slice.Raw[j] == targets[i] {
				return true
			}
		}
	}
	return false
}

// Distinct removes all duplicates from the underlying slice and returns it
func (slice *IntSliceImpl) Distinct() []int {
	result := []int{}
	hits := map[int]bool{}
	for i := range slice.Raw {
		if _, exists := hits[slice.Raw[i]]; !exists {
			hits[slice.Raw[i]] = true
			result = append(result, slice.Raw[i])
		}
	}
	return result
}

//--------------------------------------------------------------------------------------------------
// StrSlice implementation
//--------------------------------------------------------------------------------------------------

// StrSliceImpl ...
type StrSliceImpl struct {
	Raw []string
}

// StrSlice ...
func StrSlice(slice []string) *StrSliceImpl {
	return &StrSliceImpl{Raw: slice}
}

// Contains provides a reusable to check if the given target exists in the enumerable
func (slice *StrSliceImpl) Contains(target string) bool {
	for i := range slice.Raw {
		if slice.Raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny provides a reusable to check if the given target exists in the enumerable
func (slice *StrSliceImpl) ContainsAny(targets []string) bool {
	for i := range targets {
		for j := range slice.Raw {
			if slice.Raw[j] == targets[i] {
				return true
			}
		}
	}
	return false
}

// Distinct removes all duplicates from the underlying slice and returns it
func (slice *StrSliceImpl) Distinct() []string {
	result := []string{}
	hits := map[string]bool{}
	for i := range slice.Raw {
		if _, exists := hits[slice.Raw[i]]; !exists {
			hits[slice.Raw[i]] = true
			result = append(result, slice.Raw[i])
		}
	}
	return result
}
