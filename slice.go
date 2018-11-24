// Package nub provides helper interfaces and functions reminiscent of C#'s IEnumerable methods
package nub

import (
	"errors"
	"strings"
)

//--------------------------------------------------------------------------------------------------
// IntSlice Nub
//--------------------------------------------------------------------------------------------------
type intSliceNub struct {
	Raw []int
}

// NewIntSlice creates a new nub
func NewIntSlice() *intSliceNub {
	return &intSliceNub{Raw: []int{}}
}

// IntSlice creates a new nub from the given int slice
func IntSlice(slice []int) *intSliceNub {
	if slice != nil {
		return &intSliceNub{Raw: slice}
	}
	return &intSliceNub{Raw: []int{}}
}

// Any checks if the slice has anything in it
func (slice *intSliceNub) Any() bool {
	return len(slice.Raw) > 0
}

// Append items to the end of the slice and return new slice
func (slice *intSliceNub) Append(items ...int) *intSliceNub {
	slice.Raw = append(slice.Raw, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *intSliceNub) At(i int) int {
	if i >= 0 && i < len(slice.Raw) {
		return slice.Raw[i]
	} else if i < 0 && i*-1 < len(slice.Raw) {
		return slice.Raw[len(slice.Raw)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *intSliceNub) Clear() *intSliceNub {
	slice.Raw = []int{}
	return slice
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
	if targets != nil && len(targets) > 0 {
		for i := range targets {
			for j := range slice.Raw {
				if slice.Raw[j] == targets[i] {
					return true
				}
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

// Prepend items to the begining of the slice and return new slice
func (slice *intSliceNub) Prepend(items ...int) *intSliceNub {
	items = append(items, slice.Raw...)
	slice.Raw = items
	return slice
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

// NewStrSlice creates a new nub
func NewStrSlice() *strSliceNub {
	return &strSliceNub{Raw: []string{}}
}

// StrSlice creates a new nub from the given string slice
func StrSlice(slice []string) *strSliceNub {
	if slice != nil {
		return &strSliceNub{Raw: slice}
	}
	return &strSliceNub{Raw: []string{}}
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

// Any checks if the slice has anything in it
func (slice *strSliceNub) Any() bool {
	return len(slice.Raw) > 0
}

// Append items to the end of the slice and return new slice
func (slice *strSliceNub) Append(items ...string) *strSliceNub {
	slice.Raw = append(slice.Raw, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strSliceNub) At(i int) string {
	if i >= 0 && i < len(slice.Raw) {
		return slice.Raw[i]
	} else if i < 0 && i*-1 < len(slice.Raw) {
		return slice.Raw[len(slice.Raw)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strSliceNub) Clear() *strSliceNub {
	slice.Raw = []string{}
	return slice
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
	if targets != nil && len(targets) > 0 {
		for i := range targets {
			for j := range slice.Raw {
				if slice.Raw[j] == targets[i] {
					return true
				}
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

// Prepend items to the begining of the slice and return new slice
func (slice *strSliceNub) Prepend(items ...string) *strSliceNub {
	items = append(items, slice.Raw...)
	slice.Raw = items
	return slice
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

//--------------------------------------------------------------------------------------------------
// StrMapSlice Nub
//--------------------------------------------------------------------------------------------------
type strMapSliceNub struct {
	Raw []map[string]interface{}
}

// NewStrMapSlice creates a new nub
func NewStrMapSlice() *strMapSliceNub {
	return &strMapSliceNub{Raw: []map[string]interface{}{}}
}

// StrMapSlice creates a new nub from the given string map slice
func StrMapSlice(slice []map[string]interface{}) *strMapSliceNub {
	if slice != nil {
		return &strMapSliceNub{Raw: slice}
	}
	return &strMapSliceNub{Raw: []map[string]interface{}{}}
}

// Any checks if the slice has anything in it
func (slice *strMapSliceNub) Any() bool {
	return len(slice.Raw) > 0
}

// Append items to the end of the slice and return new slice
func (slice *strMapSliceNub) Append(items ...map[string]interface{}) *strMapSliceNub {
	slice.Raw = append(slice.Raw, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strMapSliceNub) At(i int) *strMapNub {
	if i >= 0 && i < len(slice.Raw) {
		return StrMap(slice.Raw[i])
	} else if i < 0 && i*-1 < len(slice.Raw) {
		return StrMap(slice.Raw[len(slice.Raw)+i])
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strMapSliceNub) Clear() *strMapSliceNub {
	slice.Raw = []map[string]interface{}{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *strMapSliceNub) Contains(key string) bool {
	for i := range slice.Raw {
		if _, exists := slice.Raw[i][key]; exists {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strMapSliceNub) ContainsAny(keys []string) bool {
	if keys != nil && len(keys) > 0 {
		for i := range keys {
			for j := range slice.Raw {
				if _, exists := slice.Raw[j][keys[i]]; exists {
					return true
				}
			}
		}
	}
	return false
}

// Len is a pass through to the underlying slice
func (slice *strMapSliceNub) Len() int {
	return len(slice.Raw)
}

// Prepend items to the begining of the slice and return new slice
func (slice *strMapSliceNub) Prepend(items ...map[string]interface{}) *strMapSliceNub {
	items = append(items, slice.Raw...)
	slice.Raw = items
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strMapSliceNub) TakeFirst() (*strMapNub, bool) {
	if len(slice.Raw) > 0 {
		item := StrMap(slice.Raw[0])
		slice.Raw = slice.Raw[1:]
		return item, true
	}
	return nil, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strMapSliceNub) TakeFirstCnt(cnt int) (result *strMapSliceNub) {
	if cnt > 0 {
		if len(slice.Raw) >= cnt {
			result = StrMapSlice(slice.Raw[:cnt])
			slice.Raw = slice.Raw[cnt:]
		} else {
			result = StrMapSlice(slice.Raw)
			slice.Raw = []map[string]interface{}{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *strMapSliceNub) TakeLast() (*strMapNub, bool) {
	if len(slice.Raw) > 0 {
		item := StrMap(slice.Raw[len(slice.Raw)-1])
		slice.Raw = slice.Raw[:len(slice.Raw)-1]
		return item, true
	}
	return nil, false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strMapSliceNub) TakeLastCnt(cnt int) (result *strMapSliceNub) {
	if cnt > 0 {
		if len(slice.Raw) >= cnt {
			i := len(slice.Raw) - cnt
			result = StrMapSlice(slice.Raw[i:])
			slice.Raw = slice.Raw[:i]
		} else {
			result = StrMapSlice(slice.Raw)
			slice.Raw = []map[string]interface{}{}
		}
	}
	return
}
