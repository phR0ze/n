// Package nub provides helper interfaces and functions reminiscent ruby/C#
package nub

import (
	"errors"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//--------------------------------------------------------------------------------------------------
// StrSlice Nub
//--------------------------------------------------------------------------------------------------
type strSliceNub struct {
	v []string
}

// S provides a new empty Queryable slice
func S(v ...string) *strSliceNub {
	if v == nil {
		v = []string{}
	}
	return &strSliceNub{v: v}
}

// S convert the slice into an string slice
func (slice *strSliceNub) S() []string {
	return slice.v
}

// Any checks if the slice has anything in it
func (slice *strSliceNub) Any() bool {
	return len(slice.v) > 0
}

// AnyContain checks if any items in this slice contain the target
func (slice *strSliceNub) AnyContain(target string) bool {
	for i := range slice.v {
		if strings.Contains(slice.v[i], target) {
			return true
		}
	}
	return false
}

// Append items to the end of the slice and return slice
func (slice *strSliceNub) Append(items ...string) *strSliceNub {
	slice.v = append(slice.v, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strSliceNub) At(i int) string {
	if i >= 0 && i < len(slice.v) {
		return slice.v[i]
	} else if i < 0 && i*-1 < len(slice.v) {
		return slice.v[len(slice.v)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strSliceNub) Clear() *strSliceNub {
	slice.v = []string{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *strSliceNub) Contains(target string) bool {
	for i := range slice.v {
		if slice.v[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strSliceNub) ContainsAny(targets []string) bool {
	if targets != nil && len(targets) > 0 {
		for i := range targets {
			for j := range slice.v {
				if slice.v[j] == targets[i] {
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
		i = len(slice.v) + i
	}
	if i >= 0 && i < len(slice.v) {
		if i+1 < len(slice.v) {
			slice.v = append(slice.v[:i], slice.v[i+1:]...)
			result = true
		} else {
			slice.v = slice.v[:i]
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
	return A(strings.Join(slice.v, delim))
}

// Len is a pass through to the underlying slice
func (slice *strSliceNub) Len() int {
	return len(slice.v)
}

// Pair simply returns the first and second slice items
func (slice *strSliceNub) Pair() (first string, second string) {
	if slice.Len() > 0 {
		first = slice.v[0]
	}
	if slice.Len() > 1 {
		second = slice.v[1]
	}
	return
}

// Prepend items to the begining of the slice and return slice
func (slice *strSliceNub) Prepend(items ...string) *strSliceNub {
	items = append(items, slice.v...)
	slice.v = items
	return slice
}

// Sort the underlying slice
func (slice *strSliceNub) Sort() *strSliceNub {
	sort.Strings(slice.v)
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strSliceNub) TakeFirst() (string, bool) {
	if len(slice.v) > 0 {
		item := slice.v[0]
		slice.v = slice.v[1:]
		return item, true
	}
	return "", false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strSliceNub) TakeFirstCnt(cnt int) (result *strSliceNub) {
	if cnt > 0 {
		if len(slice.v) >= cnt {
			result = S(slice.v[:cnt]...)
			slice.v = slice.v[cnt:]
		} else {
			result = S(slice.v...)
			slice.v = []string{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *strSliceNub) TakeLast() (string, bool) {
	if len(slice.v) > 0 {
		item := slice.v[len(slice.v)-1]
		slice.v = slice.v[:len(slice.v)-1]
		return item, true
	}
	return "", false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strSliceNub) TakeLastCnt(cnt int) (result *strSliceNub) {
	if cnt > 0 {
		if len(slice.v) >= cnt {
			i := len(slice.v) - cnt
			result = S(slice.v[i:]...)
			slice.v = slice.v[:i]
		} else {
			result = S(slice.v...)
			slice.v = []string{}
		}
	}
	return
}

// Uniq removes all duplicates from the underlying slice
func (slice *strSliceNub) Uniq() *strSliceNub {
	hits := map[string]bool{}
	for i := len(slice.v) - 1; i >= 0; i-- {
		if _, exists := hits[slice.v[i]]; !exists {
			hits[slice.v[i]] = true
		} else {
			slice.v = append(slice.v[:i], slice.v[i+1:]...)
		}
	}
	return slice
}

//--------------------------------------------------------------------------------------------------
// IntSlice Nub
//--------------------------------------------------------------------------------------------------
type intSliceNub struct {
	v []int
}

// NewIntSlice creates a new nub
func NewIntSlice() *intSliceNub {
	return &intSliceNub{v: []int{}}
}

// IntSlice creates a new nub from the given int slice
func IntSlice(slice []int) *intSliceNub {
	if slice != nil {
		return &intSliceNub{v: slice}
	}
	return &intSliceNub{v: []int{}}
}

// S convert the slice into an int slice
func (slice *intSliceNub) S() []int {
	return slice.v
}

// Any checks if the slice has anything in it
func (slice *intSliceNub) Any() bool {
	return len(slice.v) > 0
}

// Append items to the end of the slice and return slice
func (slice *intSliceNub) Append(items ...int) *intSliceNub {
	slice.v = append(slice.v, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *intSliceNub) At(i int) int {
	if i >= 0 && i < len(slice.v) {
		return slice.v[i]
	} else if i < 0 && i*-1 < len(slice.v) {
		return slice.v[len(slice.v)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *intSliceNub) Clear() *intSliceNub {
	slice.v = []int{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *intSliceNub) Contains(target int) bool {
	for i := range slice.v {
		if slice.v[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *intSliceNub) ContainsAny(targets []int) bool {
	if targets != nil && len(targets) > 0 {
		for i := range targets {
			for j := range slice.v {
				if slice.v[j] == targets[i] {
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
		i = len(slice.v) + i
	}
	if i >= 0 && i < len(slice.v) {
		if i+1 < len(slice.v) {
			slice.v = append(slice.v[:i], slice.v[i+1:]...)
			result = true
		} else {
			slice.v = slice.v[:i]
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
	for i := range slice.v {
		result = append(result, strconv.Itoa(slice.v[i]))
	}
	return A(strings.Join(result, delim))
}

// Len is a pass through to the underlying slice
func (slice *intSliceNub) Len() int {
	return len(slice.v)
}

// Prepend items to the begining of the slice and return slice
func (slice *intSliceNub) Prepend(items ...int) *intSliceNub {
	items = append(items, slice.v...)
	slice.v = items
	return slice
}

// Sort the underlying slice
func (slice *intSliceNub) Sort() *intSliceNub {
	sort.Ints(slice.v)
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *intSliceNub) TakeFirst() (int, bool) {
	if len(slice.v) > 0 {
		item := slice.v[0]
		slice.v = slice.v[1:]
		return item, true
	}
	return 0, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *intSliceNub) TakeFirstCnt(cnt int) (result *intSliceNub) {
	if cnt > 0 {
		if len(slice.v) >= cnt {
			result = IntSlice(slice.v[:cnt])
			slice.v = slice.v[cnt:]
		} else {
			result = IntSlice(slice.v)
			slice.v = []int{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *intSliceNub) TakeLast() (int, bool) {
	if len(slice.v) > 0 {
		item := slice.v[len(slice.v)-1]
		slice.v = slice.v[:len(slice.v)-1]
		return item, true
	}
	return 0, false
}

// TakeLastCnt updates the underlying slice and returns the items
func (slice *intSliceNub) TakeLastCnt(cnt int) (result *intSliceNub) {
	if cnt > 0 {
		if len(slice.v) >= cnt {
			i := len(slice.v) - cnt
			result = IntSlice(slice.v[i:])
			slice.v = slice.v[:i]
		} else {
			result = IntSlice(slice.v)
			slice.v = []int{}
		}
	}
	return

}

// Uniq removes all duplicates from the underlying slice
func (slice *intSliceNub) Uniq() *intSliceNub {
	hits := map[int]bool{}
	for i := len(slice.v) - 1; i >= 0; i-- {
		if _, exists := hits[slice.v[i]]; !exists {
			hits[slice.v[i]] = true
		} else {
			slice.v = append(slice.v[:i], slice.v[i+1:]...)
		}
	}
	return slice
}

//--------------------------------------------------------------------------------------------------
// StrMapSlice Nub
//--------------------------------------------------------------------------------------------------
type strMapSliceNub struct {
	v []map[string]interface{}
}

// NewStrMapSlice creates a new nub
func NewStrMapSlice() *strMapSliceNub {
	return &strMapSliceNub{v: []map[string]interface{}{}}
}

// StrMapSlice creates a new nub from the given string map slice
func StrMapSlice(slice []map[string]interface{}) *strMapSliceNub {
	if slice != nil {
		return &strMapSliceNub{v: slice}
	}
	return &strMapSliceNub{v: []map[string]interface{}{}}
}

// S convert the slice into an string slice
func (slice *strMapSliceNub) S() []map[string]interface{} {
	return slice.v
}

// Any checks if the slice has anything in it
func (slice *strMapSliceNub) Any() bool {
	return len(slice.v) > 0
}

// Append items to the end of the slice and return slice
func (slice *strMapSliceNub) Append(items ...map[string]interface{}) *strMapSliceNub {
	slice.v = append(slice.v, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strMapSliceNub) At(i int) *strMapNub {
	if i >= 0 && i < len(slice.v) {
		return StrMap(slice.v[i])
	} else if i < 0 && i*-1 < len(slice.v) {
		return StrMap(slice.v[len(slice.v)+i])
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strMapSliceNub) Clear() *strMapSliceNub {
	slice.v = []map[string]interface{}{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *strMapSliceNub) Contains(key string) bool {
	for i := range slice.v {
		if _, exists := slice.v[i][key]; exists {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strMapSliceNub) ContainsAny(keys []string) bool {
	if keys != nil && len(keys) > 0 {
		for i := range keys {
			for j := range slice.v {
				if _, exists := slice.v[j][keys[i]]; exists {
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
		i = len(slice.v) + i
	}
	if i >= 0 && i < len(slice.v) {
		if i+1 < len(slice.v) {
			slice.v = append(slice.v[:i], slice.v[i+1:]...)
			result = true
		} else {
			slice.v = slice.v[:i]
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
	return len(slice.v)
}

// Prepend items to the begining of the slice and return slice
func (slice *strMapSliceNub) Prepend(items ...map[string]interface{}) *strMapSliceNub {
	items = append(items, slice.v...)
	slice.v = items
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strMapSliceNub) TakeFirst() (*strMapNub, bool) {
	if len(slice.v) > 0 {
		item := StrMap(slice.v[0])
		slice.v = slice.v[1:]
		return item, true
	}
	return nil, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strMapSliceNub) TakeFirstCnt(cnt int) (result *strMapSliceNub) {
	if cnt > 0 {
		if len(slice.v) >= cnt {
			result = StrMapSlice(slice.v[:cnt])
			slice.v = slice.v[cnt:]
		} else {
			result = StrMapSlice(slice.v)
			slice.v = []map[string]interface{}{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (slice *strMapSliceNub) TakeLast() (*strMapNub, bool) {
	if len(slice.v) > 0 {
		item := StrMap(slice.v[len(slice.v)-1])
		slice.v = slice.v[:len(slice.v)-1]
		return item, true
	}
	return nil, false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strMapSliceNub) TakeLastCnt(cnt int) (result *strMapSliceNub) {
	if cnt > 0 {
		if len(slice.v) >= cnt {
			i := len(slice.v) - cnt
			result = StrMapSlice(slice.v[i:])
			slice.v = slice.v[:i]
		} else {
			result = StrMapSlice(slice.v)
			slice.v = []map[string]interface{}{}
		}
	}
	return
}

// YAMLPair return the first and second entries as yaml types
func (slice *strSliceNub) YAMLPair() (first interface{}, second interface{}) {
	if slice.Len() > 0 {
		first = A(slice.v[0]).YAMLType()
	}
	if slice.Len() > 1 {
		second = A(slice.v[1]).YAMLType()
	}
	return
}
