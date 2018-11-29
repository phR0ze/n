// package n provides helper interfaces and functions reminiscent ruby/C#
package n

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
type strSliceN struct {
	v []string
}

// S provides a new empty Queryable slice
func S(v ...string) *strSliceN {
	if v == nil {
		v = []string{}
	}
	return &strSliceN{v: v}
}

// S convert the slice into an string slice
func (slice *strSliceN) S() []string {
	return slice.v
}

// Any checks if the slice has anything in it
func (slice *strSliceN) Any() bool {
	return len(slice.v) > 0
}

// AnyContain checks if any items in this slice contain the target
func (slice *strSliceN) AnyContain(target string) bool {
	for i := range slice.v {
		if strings.Contains(slice.v[i], target) {
			return true
		}
	}
	return false
}

// Append items to the end of the slice and return slice
func (slice *strSliceN) Append(items ...string) *strSliceN {
	slice.v = append(slice.v, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strSliceN) At(i int) string {
	if i >= 0 && i < len(slice.v) {
		return slice.v[i]
	} else if i < 0 && i*-1 < len(slice.v) {
		return slice.v[len(slice.v)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strSliceN) Clear() *strSliceN {
	slice.v = []string{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *strSliceN) Contains(target string) bool {
	for i := range slice.v {
		if slice.v[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strSliceN) ContainsAny(targets []string) bool {
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
func (slice *strSliceN) Del(i int) bool {
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
func (slice *strSliceN) Equals(other *strSliceN) bool {
	return reflect.DeepEqual(slice, other)
}

// Join the underlying slice with the given delim
func (slice *strSliceN) Join(delim string) *strN {
	return A(strings.Join(slice.v, delim))
}

// Len is a pass through to the underlying slice
func (slice *strSliceN) Len() int {
	return len(slice.v)
}

// Pair simply returns the first and second slice items
func (slice *strSliceN) Pair() (first string, second string) {
	if slice.Len() > 0 {
		first = slice.v[0]
	}
	if slice.Len() > 1 {
		second = slice.v[1]
	}
	return
}

// Prepend items to the begining of the slice and return slice
func (slice *strSliceN) Prepend(items ...string) *strSliceN {
	items = append(items, slice.v...)
	slice.v = items
	return slice
}

// Sort the underlying slice
func (slice *strSliceN) Sort() *strSliceN {
	sort.Strings(slice.v)
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strSliceN) TakeFirst() (string, bool) {
	if len(slice.v) > 0 {
		item := slice.v[0]
		slice.v = slice.v[1:]
		return item, true
	}
	return "", false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strSliceN) TakeFirstCnt(cnt int) (result *strSliceN) {
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
func (slice *strSliceN) TakeLast() (string, bool) {
	if len(slice.v) > 0 {
		item := slice.v[len(slice.v)-1]
		slice.v = slice.v[:len(slice.v)-1]
		return item, true
	}
	return "", false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strSliceN) TakeLastCnt(cnt int) (result *strSliceN) {
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
func (slice *strSliceN) Uniq() *strSliceN {
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
type intSliceN struct {
	v []int
}

// NewIntSlice creates a new nub
func NewIntSlice() *intSliceN {
	return &intSliceN{v: []int{}}
}

// IntSlice creates a new nub from the given int slice
func IntSlice(slice []int) *intSliceN {
	if slice != nil {
		return &intSliceN{v: slice}
	}
	return &intSliceN{v: []int{}}
}

// S convert the slice into an int slice
func (slice *intSliceN) S() []int {
	return slice.v
}

// Any checks if the slice has anything in it
func (slice *intSliceN) Any() bool {
	return len(slice.v) > 0
}

// Append items to the end of the slice and return slice
func (slice *intSliceN) Append(items ...int) *intSliceN {
	slice.v = append(slice.v, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *intSliceN) At(i int) int {
	if i >= 0 && i < len(slice.v) {
		return slice.v[i]
	} else if i < 0 && i*-1 < len(slice.v) {
		return slice.v[len(slice.v)+i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *intSliceN) Clear() *intSliceN {
	slice.v = []int{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *intSliceN) Contains(target int) bool {
	for i := range slice.v {
		if slice.v[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *intSliceN) ContainsAny(targets []int) bool {
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
func (slice *intSliceN) Del(i int) bool {
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
func (slice *intSliceN) Equals(other *intSliceN) bool {
	return reflect.DeepEqual(slice, other)
}

// Join the underlying slice with the given delim
func (slice *intSliceN) Join(delim string) *strN {
	result := []string{}
	for i := range slice.v {
		result = append(result, strconv.Itoa(slice.v[i]))
	}
	return A(strings.Join(result, delim))
}

// Len is a pass through to the underlying slice
func (slice *intSliceN) Len() int {
	return len(slice.v)
}

// Prepend items to the begining of the slice and return slice
func (slice *intSliceN) Prepend(items ...int) *intSliceN {
	items = append(items, slice.v...)
	slice.v = items
	return slice
}

// Sort the underlying slice
func (slice *intSliceN) Sort() *intSliceN {
	sort.Ints(slice.v)
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *intSliceN) TakeFirst() (int, bool) {
	if len(slice.v) > 0 {
		item := slice.v[0]
		slice.v = slice.v[1:]
		return item, true
	}
	return 0, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *intSliceN) TakeFirstCnt(cnt int) (result *intSliceN) {
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
func (slice *intSliceN) TakeLast() (int, bool) {
	if len(slice.v) > 0 {
		item := slice.v[len(slice.v)-1]
		slice.v = slice.v[:len(slice.v)-1]
		return item, true
	}
	return 0, false
}

// TakeLastCnt updates the underlying slice and returns the items
func (slice *intSliceN) TakeLastCnt(cnt int) (result *intSliceN) {
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
func (slice *intSliceN) Uniq() *intSliceN {
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
type strMapSliceN struct {
	v []map[string]interface{}
}

// NewStrMapSlice creates a new nub
func NewStrMapSlice() *strMapSliceN {
	return &strMapSliceN{v: []map[string]interface{}{}}
}

// StrMapSlice creates a new nub from the given string map slice
func StrMapSlice(slice []map[string]interface{}) *strMapSliceN {
	if slice != nil {
		return &strMapSliceN{v: slice}
	}
	return &strMapSliceN{v: []map[string]interface{}{}}
}

// S convert the slice into an string slice
func (slice *strMapSliceN) S() []map[string]interface{} {
	return slice.v
}

// Any checks if the slice has anything in it
func (slice *strMapSliceN) Any() bool {
	return len(slice.v) > 0
}

// Append items to the end of the slice and return slice
func (slice *strMapSliceN) Append(items ...map[string]interface{}) *strMapSliceN {
	slice.v = append(slice.v, items...)
	return slice
}

// At returns the item at the given index location. Allows for negative notation
func (slice *strMapSliceN) At(i int) *strMapN {
	if i >= 0 && i < len(slice.v) {
		return M(slice.v[i])
	} else if i < 0 && i*-1 < len(slice.v) {
		return M(slice.v[len(slice.v)+i])
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (slice *strMapSliceN) Clear() *strMapSliceN {
	slice.v = []map[string]interface{}{}
	return slice
}

// Contains checks if the given target is contained in this slice
func (slice *strMapSliceN) Contains(key string) bool {
	for i := range slice.v {
		if _, exists := slice.v[i][key]; exists {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (slice *strMapSliceN) ContainsAny(keys []string) bool {
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
func (slice *strMapSliceN) Del(i int) bool {
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
func (slice *strMapSliceN) Equals(other *strMapSliceN) bool {
	return reflect.DeepEqual(slice, other)
}

// Len is a pass through to the underlying slice
func (slice *strMapSliceN) Len() int {
	return len(slice.v)
}

// Prepend items to the begining of the slice and return slice
func (slice *strMapSliceN) Prepend(items ...map[string]interface{}) *strMapSliceN {
	items = append(items, slice.v...)
	slice.v = items
	return slice
}

// TakeFirst updates the underlying slice and returns the item and status
func (slice *strMapSliceN) TakeFirst() (*strMapN, bool) {
	if len(slice.v) > 0 {
		item := M(slice.v[0])
		slice.v = slice.v[1:]
		return item, true
	}
	return nil, false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (slice *strMapSliceN) TakeFirstCnt(cnt int) (result *strMapSliceN) {
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
func (slice *strMapSliceN) TakeLast() (*strMapN, bool) {
	if len(slice.v) > 0 {
		item := M(slice.v[len(slice.v)-1])
		slice.v = slice.v[:len(slice.v)-1]
		return item, true
	}
	return nil, false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (slice *strMapSliceN) TakeLastCnt(cnt int) (result *strMapSliceN) {
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
func (slice *strSliceN) YAMLPair() (first string, second interface{}) {
	if slice.Len() > 0 {
		first = slice.v[0]
	}
	if slice.Len() > 1 {
		second = A(slice.v[1]).YAMLType()
	}
	return
}
