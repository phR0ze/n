// Package nub provides helper interfaces and functions reminiscent of C#'s IEnumerable methods
package nub

import (
	"strings"
)

//--------------------------------------------------------------------------------------------------
// IntSlice Nub
//--------------------------------------------------------------------------------------------------

// IntSliceNub provides missing functionality and elegant chaining
type IntSliceNub struct {
	Raw []int
}

// IntSlice creates a new nub
func IntSlice(slice []int) *IntSliceNub {
	return &IntSliceNub{Raw: slice}
}

// Contains provides a reusable check for the given target
func (slice *IntSliceNub) Contains(target int) bool {
	for i := range slice.Raw {
		if slice.Raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny provides a reusable check for the given targets
func (slice *IntSliceNub) ContainsAny(targets []int) bool {
	for i := range targets {
		for j := range slice.Raw {
			if slice.Raw[j] == targets[i] {
				return true
			}
		}
	}
	return false
}

// Distinct removes all duplicates from the underlying slice
func (slice *IntSliceNub) Distinct() *IntSliceNub {
	hits := map[int]bool{}
	for i := len(slice.Raw) - 1; i >= 0; i-- {
		if _, exists := hits[slice.Raw[i]]; !exists {
			hits[slice.Raw[i]] = true
		} else {
			slice.Raw = append(slice.Raw[:i], slice.Raw[i+1:]...)
		}
	}
	return slice
}

// Len is a pass through to the underlying slice
func (slice *IntSliceNub) Len() int {
	return len(slice.Raw)
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *IntSliceNub) TakeFirst() (int, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[0]
		slice.Raw = slice.Raw[1:]
		return item, true
	}
	return 0, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *IntSliceNub) TakeFirstCnt(cnt int) (result *IntSliceNub) {
	if cnt > 0 {
		if len(slice.Raw) >= cnt {
			result = IntSlice(slice.Raw[:cnt])
			slice.Raw = slice.Raw[cnt:]
		} else {
			result = IntSlice(slice.Raw)
			slice.Raw = []int{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *IntSliceNub) TakeLast() (int, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[len(slice.Raw)-1]
		slice.Raw = slice.Raw[:len(slice.Raw)-1]
		return item, true
	}
	return 0, false
}

// TakeLastCnt updates the underlying slice and returns the items
func (slice *IntSliceNub) TakeLastCnt(cnt int) (result *IntSliceNub) {
	if cnt > 0 {
		if len(slice.Raw) >= cnt {
			i := len(slice.Raw) - cnt
			result = IntSlice(slice.Raw[i:])
			slice.Raw = slice.Raw[:i]
		} else {
			result = IntSlice(slice.Raw)
			slice.Raw = []int{}
		}
	}
	return

}

//--------------------------------------------------------------------------------------------------
// StrSlice Nub
//--------------------------------------------------------------------------------------------------

// StrSliceNub provides missing functionality and elegant chaining
type StrSliceNub struct {
	Raw []string
}

// StrSlice creates a new nub
func StrSlice(slice []string) *StrSliceNub {
	return &StrSliceNub{Raw: slice}
}

// Contains provides a reusable check for the given target
func (slice *StrSliceNub) Contains(target string) bool {
	for i := range slice.Raw {
		if slice.Raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny provides a reusable check for the given targets
func (slice *StrSliceNub) ContainsAny(targets []string) bool {
	for i := range targets {
		for j := range slice.Raw {
			if slice.Raw[j] == targets[i] {
				return true
			}
		}
	}
	return false
}

// Distinct removes all duplicates from the underlying slice
func (slice *StrSliceNub) Distinct() *StrSliceNub {
	hits := map[string]bool{}
	for i := len(slice.Raw) - 1; i >= 0; i-- {
		if _, exists := hits[slice.Raw[i]]; !exists {
			hits[slice.Raw[i]] = true
		} else {
			slice.Raw = append(slice.Raw[:i], slice.Raw[i+1:]...)
		}
	}
	return slice
}

// Join the underlying slice with the given delim
func (slice *StrSliceNub) Join(delim string) *StrNub {
	return Str(strings.Join(slice.Raw, delim))
}

// Len is a pass through to the underlying slice
func (slice *StrSliceNub) Len() int {
	return len(slice.Raw)
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *StrSliceNub) TakeFirst() (string, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[0]
		slice.Raw = slice.Raw[1:]
		return item, true
	}
	return "", false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *StrSliceNub) TakeFirstCnt(cnt int) (result *StrSliceNub) {
	if cnt > 0 {
		if len(slice.Raw) >= cnt {
			result = StrSlice(slice.Raw[:cnt])
			slice.Raw = slice.Raw[cnt:]
		} else {
			result = StrSlice(slice.Raw)
			slice.Raw = []string{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *StrSliceNub) TakeLast() (string, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[len(slice.Raw)-1]
		slice.Raw = slice.Raw[:len(slice.Raw)-1]
		return item, true
	}
	return "", false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *StrSliceNub) TakeLastCnt(cnt int) (result *StrSliceNub) {
	if cnt > 0 {
		if len(slice.Raw) >= cnt {
			i := len(slice.Raw) - cnt
			result = StrSlice(slice.Raw[i:])
			slice.Raw = slice.Raw[:i]
		} else {
			result = StrSlice(slice.Raw)
			slice.Raw = []string{}
		}
	}
	return
}
