// Package nub provides helper interfaces and functions reminiscent of C#'s IEnumerable methods
package nub

import (
	"strings"
)

//--------------------------------------------------------------------------------------------------
// IntSlice Nub
//--------------------------------------------------------------------------------------------------
type intSliceNub struct {
	Raw []int
}

// IntSlice creates a new nub from the given int slice
func IntSlice(slice []int) *intSliceNub {
	return &intSliceNub{Raw: slice}
}

// Contains checks if the given target is contained in this slice
func (slice *intSliceNub) Contains(target int) bool {
	for i := range slice.Raw {
		if slice.Raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *intSliceNub) ContainsAny(targets []int) bool {
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
func (slice *intSliceNub) Distinct() *intSliceNub {
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
func (slice *intSliceNub) Len() int {
	return len(slice.Raw)
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *intSliceNub) TakeFirst() (int, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[0]
		slice.Raw = slice.Raw[1:]
		return item, true
	}
	return 0, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *intSliceNub) TakeFirstCnt(cnt int) (result *intSliceNub) {
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
func (slice *intSliceNub) TakeLast() (int, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[len(slice.Raw)-1]
		slice.Raw = slice.Raw[:len(slice.Raw)-1]
		return item, true
	}
	return 0, false
}

// TakeLastCnt updates the underlying slice and returns the items
func (slice *intSliceNub) TakeLastCnt(cnt int) (result *intSliceNub) {
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
type strSliceNub struct {
	Raw []string
}

// StrSlice creates a new nub from the given string slice
func StrSlice(slice []string) *strSliceNub {
	return &strSliceNub{Raw: slice}
}

// AnyContain checks if any items in this slice contain the target
func (slice *strSliceNub) AnyContain(target string) bool {
	for i := range slice.Raw {
		if strings.Contains(slice.Raw[i], target) {
			return true
		}
	}
	return false
}

// Contains checks if the given target is contained in this slice
func (slice *strSliceNub) Contains(target string) bool {
	for i := range slice.Raw {
		if slice.Raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strSliceNub) ContainsAny(targets []string) bool {
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
func (slice *strSliceNub) Distinct() *strSliceNub {
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
func (slice *strSliceNub) Join(delim string) *strNub {
	return Str(strings.Join(slice.Raw, delim))
}

// Len is a pass through to the underlying slice
func (slice *strSliceNub) Len() int {
	return len(slice.Raw)
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strSliceNub) TakeFirst() (string, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[0]
		slice.Raw = slice.Raw[1:]
		return item, true
	}
	return "", false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strSliceNub) TakeFirstCnt(cnt int) (result *strSliceNub) {
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
func (slice *strSliceNub) TakeLast() (string, bool) {
	if len(slice.Raw) > 0 {
		item := slice.Raw[len(slice.Raw)-1]
		slice.Raw = slice.Raw[:len(slice.Raw)-1]
		return item, true
	}
	return "", false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strSliceNub) TakeLastCnt(cnt int) (result *strSliceNub) {
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
