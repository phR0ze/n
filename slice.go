// Package nub provides helper interfaces and functions reminiscent ruby/C#
package nub

import (
	"errors"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// TODO: Refactor below here

//--------------------------------------------------------------------------------------------------
// IntSlice Nub
//--------------------------------------------------------------------------------------------------
type intSliceNub struct {
	raw []int
}

// NewIntSlice creates a new nub
func NewIntSlice() *intSliceNub {
	return &intSliceNub{raw: []int{}}
}

// IntSlice creates a new nub from the given int slice
func IntSlice(slice []int) *intSliceNub {
	if slice != nil {
		return &intSliceNub{raw: slice}
	}
	return &intSliceNub{raw: []int{}}
}

// Any checks if the slice has anything in it
func (slice *intSliceNub) Any() bool {
	return len(slice.raw) > 0
}

// Append items to the end of the slice and return slice
func (slice *intSliceNub) Append(items ...int) *intSliceNub {
	slice.raw = append(slice.raw, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *intSliceNub) At(i int) int {
	if i >= 0 && i < len(slice.raw) {
		return slice.raw[i]
	} else if i < 0 && i*-1 < len(slice.raw) {
		return slice.raw[len(slice.raw)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *intSliceNub) Clear() *intSliceNub {
	slice.raw = []int{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *intSliceNub) Contains(target int) bool {
	for i := range slice.raw {
		if slice.raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *intSliceNub) ContainsAny(targets []int) bool {
	if targets != nil && len(targets) > 0 {
		for i := range targets {
			for j := range slice.raw {
				if slice.raw[j] == targets[i] {
					return true
				}
			}
		}
	}
	return false
}

// Del deletes item using neg/pos index notation with status
func (slice *intSliceNub) Del(i int) bool {
	result := false
	if i < 0 {
		i = len(slice.raw) + i
	}
	if i >= 0 && i < len(slice.raw) {
		if i+1 < len(slice.raw) {
			slice.raw = append(slice.raw[:i], slice.raw[i+1:]...)
			result = true
		} else {
			slice.raw = slice.raw[:i]
			result = true
		}
	}
	return result
}

// Equals checks if the two slices are equal
func (slice *intSliceNub) Equals(other *intSliceNub) bool {
	return reflect.DeepEqual(slice, other)
}

// Join the underlying slice with the given delim
func (slice *intSliceNub) Join(delim string) *strNub {
	result := []string{}
	for i := range slice.raw {
		result = append(result, strconv.Itoa(slice.raw[i]))
	}
	return A(strings.Join(result, delim))
}

// Len is a pass through to the underlying slice
func (slice *intSliceNub) Len() int {
	return len(slice.raw)
}

// M materializes object invoking deferred execution
func (slice *intSliceNub) M() []int {
	return slice.raw
}

// Prepend items to the begining of the slice and return slice
func (slice *intSliceNub) Prepend(items ...int) *intSliceNub {
	items = append(items, slice.raw...)
	slice.raw = items
	return slice
}

// Sort the underlying slice
func (slice *intSliceNub) Sort() *intSliceNub {
	sort.Ints(slice.raw)
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *intSliceNub) TakeFirst() (int, bool) {
	if len(slice.raw) > 0 {
		item := slice.raw[0]
		slice.raw = slice.raw[1:]
		return item, true
	}
	return 0, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *intSliceNub) TakeFirstCnt(cnt int) (result *intSliceNub) {
	if cnt > 0 {
		if len(slice.raw) >= cnt {
			result = IntSlice(slice.raw[:cnt])
			slice.raw = slice.raw[cnt:]
		} else {
			result = IntSlice(slice.raw)
			slice.raw = []int{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *intSliceNub) TakeLast() (int, bool) {
	if len(slice.raw) > 0 {
		item := slice.raw[len(slice.raw)-1]
		slice.raw = slice.raw[:len(slice.raw)-1]
		return item, true
	}
	return 0, false
}

// TakeLastCnt updates the underlying slice and returns the items
func (slice *intSliceNub) TakeLastCnt(cnt int) (result *intSliceNub) {
	if cnt > 0 {
		if len(slice.raw) >= cnt {
			i := len(slice.raw) - cnt
			result = IntSlice(slice.raw[i:])
			slice.raw = slice.raw[:i]
		} else {
			result = IntSlice(slice.raw)
			slice.raw = []int{}
		}
	}
	return

}

// Uniq removes all duplicates from the underlying slice
func (slice *intSliceNub) Uniq() *intSliceNub {
	hits := map[int]bool{}
	for i := len(slice.raw) - 1; i >= 0; i-- {
		if _, exists := hits[slice.raw[i]]; !exists {
			hits[slice.raw[i]] = true
		} else {
			slice.raw = append(slice.raw[:i], slice.raw[i+1:]...)
		}
	}
	return slice
}

//--------------------------------------------------------------------------------------------------
// StrSlice Nub
//--------------------------------------------------------------------------------------------------
type strSliceNub struct {
	raw []string
}

// NewStrSlice creates a new nub
func NewStrSlice() *strSliceNub {
	return &strSliceNub{raw: []string{}}
}

// StrSlice creates a new nub from the given string slice
func StrSlice(slice []string) *strSliceNub {
	if slice != nil {
		return &strSliceNub{raw: slice}
	}
	return &strSliceNub{raw: []string{}}
}

// AnyContain checks if any items in this slice contain the target
func (slice *strSliceNub) AnyContain(target string) bool {
	for i := range slice.raw {
		if strings.Contains(slice.raw[i], target) {
			return true
		}
	}
	return false
}

// Any checks if the slice has anything in it
func (slice *strSliceNub) Any() bool {
	return len(slice.raw) > 0
}

// Append items to the end of the slice and return slice
func (slice *strSliceNub) Append(items ...string) *strSliceNub {
	slice.raw = append(slice.raw, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strSliceNub) At(i int) string {
	if i >= 0 && i < len(slice.raw) {
		return slice.raw[i]
	} else if i < 0 && i*-1 < len(slice.raw) {
		return slice.raw[len(slice.raw)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strSliceNub) Clear() *strSliceNub {
	slice.raw = []string{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *strSliceNub) Contains(target string) bool {
	for i := range slice.raw {
		if slice.raw[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strSliceNub) ContainsAny(targets []string) bool {
	if targets != nil && len(targets) > 0 {
		for i := range targets {
			for j := range slice.raw {
				if slice.raw[j] == targets[i] {
					return true
				}
			}
		}
	}
	return false
}

// Del deletes item using neg/pos index notation with status
func (slice *strSliceNub) Del(i int) bool {
	result := false
	if i < 0 {
		i = len(slice.raw) + i
	}
	if i >= 0 && i < len(slice.raw) {
		if i+1 < len(slice.raw) {
			slice.raw = append(slice.raw[:i], slice.raw[i+1:]...)
			result = true
		} else {
			slice.raw = slice.raw[:i]
			result = true
		}
	}
	return result
}

// Equals checks if the two slices are equal
func (slice *strSliceNub) Equals(other *strSliceNub) bool {
	return reflect.DeepEqual(slice, other)
}

// Join the underlying slice with the given delim
func (slice *strSliceNub) Join(delim string) *strNub {
	return A(strings.Join(slice.raw, delim))
}

// Len is a pass through to the underlying slice
func (slice *strSliceNub) Len() int {
	return len(slice.raw)
}

// M materializes object invoking deferred execution
func (slice *strSliceNub) M() []string {
	return slice.raw
}

// Prepend items to the begining of the slice and return slice
func (slice *strSliceNub) Prepend(items ...string) *strSliceNub {
	items = append(items, slice.raw...)
	slice.raw = items
	return slice
}

// Sort the underlying slice
func (slice *strSliceNub) Sort() *strSliceNub {
	sort.Strings(slice.raw)
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strSliceNub) TakeFirst() (string, bool) {
	if len(slice.raw) > 0 {
		item := slice.raw[0]
		slice.raw = slice.raw[1:]
		return item, true
	}
	return "", false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strSliceNub) TakeFirstCnt(cnt int) (result *strSliceNub) {
	if cnt > 0 {
		if len(slice.raw) >= cnt {
			result = StrSlice(slice.raw[:cnt])
			slice.raw = slice.raw[cnt:]
		} else {
			result = StrSlice(slice.raw)
			slice.raw = []string{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *strSliceNub) TakeLast() (string, bool) {
	if len(slice.raw) > 0 {
		item := slice.raw[len(slice.raw)-1]
		slice.raw = slice.raw[:len(slice.raw)-1]
		return item, true
	}
	return "", false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strSliceNub) TakeLastCnt(cnt int) (result *strSliceNub) {
	if cnt > 0 {
		if len(slice.raw) >= cnt {
			i := len(slice.raw) - cnt
			result = StrSlice(slice.raw[i:])
			slice.raw = slice.raw[:i]
		} else {
			result = StrSlice(slice.raw)
			slice.raw = []string{}
		}
	}
	return
}

// Uniq removes all duplicates from the underlying slice
func (slice *strSliceNub) Uniq() *strSliceNub {
	hits := map[string]bool{}
	for i := len(slice.raw) - 1; i >= 0; i-- {
		if _, exists := hits[slice.raw[i]]; !exists {
			hits[slice.raw[i]] = true
		} else {
			slice.raw = append(slice.raw[:i], slice.raw[i+1:]...)
		}
	}
	return slice
}

//--------------------------------------------------------------------------------------------------
// StrMapSlice Nub
//--------------------------------------------------------------------------------------------------
type strMapSliceNub struct {
	raw []map[string]interface{}
}

// NewStrMapSlice creates a new nub
func NewStrMapSlice() *strMapSliceNub {
	return &strMapSliceNub{raw: []map[string]interface{}{}}
}

// StrMapSlice creates a new nub from the given string map slice
func StrMapSlice(slice []map[string]interface{}) *strMapSliceNub {
	if slice != nil {
		return &strMapSliceNub{raw: slice}
	}
	return &strMapSliceNub{raw: []map[string]interface{}{}}
}

// Any checks if the slice has anything in it
func (slice *strMapSliceNub) Any() bool {
	return len(slice.raw) > 0
}

// Append items to the end of the slice and return slice
func (slice *strMapSliceNub) Append(items ...map[string]interface{}) *strMapSliceNub {
	slice.raw = append(slice.raw, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strMapSliceNub) At(i int) *strMapNub {
	if i >= 0 && i < len(slice.raw) {
		return StrMap(slice.raw[i])
	} else if i < 0 && i*-1 < len(slice.raw) {
		return StrMap(slice.raw[len(slice.raw)+i])
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strMapSliceNub) Clear() *strMapSliceNub {
	slice.raw = []map[string]interface{}{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *strMapSliceNub) Contains(key string) bool {
	for i := range slice.raw {
		if _, exists := slice.raw[i][key]; exists {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strMapSliceNub) ContainsAny(keys []string) bool {
	if keys != nil && len(keys) > 0 {
		for i := range keys {
			for j := range slice.raw {
				if _, exists := slice.raw[j][keys[i]]; exists {
					return true
				}
			}
		}
	}
	return false
}

// Del deletes item using neg/pos index notation with status
func (slice *strMapSliceNub) Del(i int) bool {
	result := false
	if i < 0 {
		i = len(slice.raw) + i
	}
	if i >= 0 && i < len(slice.raw) {
		if i+1 < len(slice.raw) {
			slice.raw = append(slice.raw[:i], slice.raw[i+1:]...)
			result = true
		} else {
			slice.raw = slice.raw[:i]
			result = true
		}
	}
	return result
}

// Equals checks if the two slices are equal
func (slice *strMapSliceNub) Equals(other *strMapSliceNub) bool {
	return reflect.DeepEqual(slice, other)
}

// Len is a pass through to the underlying slice
func (slice *strMapSliceNub) Len() int {
	return len(slice.raw)
}

// M materializes object invoking deferred execution
func (slice *strMapSliceNub) M() []map[string]interface{} {
	return slice.raw
}

// Prepend items to the begining of the slice and return slice
func (slice *strMapSliceNub) Prepend(items ...map[string]interface{}) *strMapSliceNub {
	items = append(items, slice.raw...)
	slice.raw = items
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strMapSliceNub) TakeFirst() (*strMapNub, bool) {
	if len(slice.raw) > 0 {
		item := StrMap(slice.raw[0])
		slice.raw = slice.raw[1:]
		return item, true
	}
	return nil, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strMapSliceNub) TakeFirstCnt(cnt int) (result *strMapSliceNub) {
	if cnt > 0 {
		if len(slice.raw) >= cnt {
			result = StrMapSlice(slice.raw[:cnt])
			slice.raw = slice.raw[cnt:]
		} else {
			result = StrMapSlice(slice.raw)
			slice.raw = []map[string]interface{}{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *strMapSliceNub) TakeLast() (*strMapNub, bool) {
	if len(slice.raw) > 0 {
		item := StrMap(slice.raw[len(slice.raw)-1])
		slice.raw = slice.raw[:len(slice.raw)-1]
		return item, true
	}
	return nil, false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strMapSliceNub) TakeLastCnt(cnt int) (result *strMapSliceNub) {
	if cnt > 0 {
		if len(slice.raw) >= cnt {
			i := len(slice.raw) - cnt
			result = StrMapSlice(slice.raw[i:])
			slice.raw = slice.raw[:i]
		} else {
			result = StrMapSlice(slice.raw)
			slice.raw = []map[string]interface{}{}
		}
	}
	return
}
