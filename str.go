package n

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var (
	// ReGraphicalOnly is a regex to filter on graphical runes only
	ReGraphicalOnly = regexp.MustCompile(`[^[:graph:]]+`)
)

type strN struct {
	v string
}
type strSliceN struct {
	v []string
}

// A provides a new empty string nub
// handling both string and []byte
func A(str interface{}) (result *strN) {
	result = &strN{v: ""}
	switch x := str.(type) {
	case string:
		result.v = x
	case rune:
		result.v = string(x)
	case []byte:
		result.v = string(x)
	case int:
		result.v = strconv.Itoa(x)
	default:
		result.v = fmt.Sprintf("%v", x)
	}
	return
}

// B exports object as []byte
func (q *strN) B() []byte {
	return []byte(q.v)
}

// A exports object invoking deferred execution
func (q *strN) A() string {
	return string(q.v)
}

// Q creates a queryable from the string
func (q *strN) Q() *Queryable {
	return Q(string(q.v))
}

// At returns the rune at the given index location. Allows for negative notation
func (q *strN) At(i int) rune {
	if i < 0 {
		i = len(q.v) + i
	}
	if i >= 0 && i < len(q.v) {
		return rune(q.v[i])
	}
	panic(errors.New("Index out of string bounds"))
}

// Contains checks if the given target is contained in this string
func (q *strN) Contains(target string) bool {
	return strings.Contains(q.v, target)
}

// ContainsAny checks if any of the targets are contained in this string
func (q *strN) ContainsAny(targets ...string) bool {
	for i := range targets {
		if strings.Contains(q.v, targets[i]) {
			return true
		}
	}
	return false
}

// Empty return true if the string is nothing or just whitespace
func (q *strN) Empty() bool {
	return q.TrimSpace().v == ""
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (q *strN) HasAnyPrefix(prefixes ...string) bool {
	for i := range prefixes {
		if strings.HasPrefix(q.v, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (q *strN) HasAnySuffix(suffixes ...string) bool {
	for i := range suffixes {
		if strings.HasSuffix(q.v, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (q *strN) HasPrefix(prefix string) bool {
	return strings.HasPrefix(q.v, prefix)
}

// HasSuffix checks if the string has the given suffix
func (q *strN) HasSuffix(suffix string) bool {
	return strings.HasSuffix(q.v, suffix)
}

// Len returns the length of the string
func (q *strN) Len() int {
	return len(q.v)
}

// Replace allows for chaining and default to all instances
func (q *strN) Replace(new, old string, ns ...int) *strN {
	n := -1
	if len(ns) > 0 {
		n = ns[0]
	}
	return A(strings.Replace(q.v, new, old, n))
}

// SpaceLeft returns leading whitespace
func (q *strN) SpaceLeft() string {
	spaces := []rune{}
	for _, r := range q.v {
		if unicode.IsSpace(r) {
			spaces = append(spaces, r)
		} else {
			break
		}
	}
	return string(spaces)
}

// Split creates a new nub from the split string
func (q *strN) Split(delim string) *strSliceN {
	return S(strings.Split(q.v, delim)...)
}

// SplitOn splits the string on the first occurance of the delim.
// The delim is included in the first component.
func (q *strN) SplitOn(delim string) (first, second string) {
	if q.v != "" {
		s := q.Split(delim)
		if s.Len() > 0 {
			first = s.First().A()
			if strings.Contains(q.v, delim) {
				first += delim
			}
		}
		if s.Len() > 1 {
			second = s.Slice(1, -1).Join(delim).A()
		}
	}
	return
}

// ToASCII with given string
func (q *strN) ToASCII() *strN {
	return A(ReGraphicalOnly.ReplaceAllString(q.v, " "))
}

// TrimPrefix trims the given prefix off the string
func (q *strN) TrimPrefix(prefix string) *strN {
	return A(strings.TrimPrefix(q.v, prefix))
}

// TrimSpace pass through to strings.TrimSpace
func (q *strN) TrimSpace() *strN {
	return A(strings.TrimSpace(q.v))
}

// TrimSpaceLeft trims leading whitespace
func (q *strN) TrimSpaceLeft() *strN {
	return A(strings.TrimLeftFunc(q.v, unicode.IsSpace))
}

// TrimSpaceRight trims trailing whitespace
func (q *strN) TrimSpaceRight() *strN {
	return A(strings.TrimRightFunc(q.v, unicode.IsSpace))
}

// TrimSuffix trims the given suffix off the string
func (q *strN) TrimSuffix(suffix string) *strN {
	return A(strings.TrimSuffix(q.v, suffix))
}

// YamlType converts the given string into a type expected in Yaml.
// Quotes signifies a string.
// No quotes signifies an int.
// true or false signifies a bool.
func (q *strN) YamlType() interface{} {
	if q.HasAnyPrefix("\"", "'") && q.HasAnySuffix("\"", "'") {
		return q.v[1 : len(q.v)-1]
	} else if q.v == "true" || q.v == "false" {
		if b, err := strconv.ParseBool(q.v); err == nil {
			return b
		}
	} else if f, err := strconv.ParseFloat(q.v, 32); err == nil {
		return f
	}
	return q.v
}

// String Slice
//--------------------------------------------------------------------------------------------------

// S provides a new empty string slice
func S(v ...string) *strSliceN {
	if v == nil {
		v = []string{}
	}
	return &strSliceN{v: v}
}

// S convert the slice into an string slice
func (s *strSliceN) S() []string {
	return s.v
}

// Any checks if the slice has anything in it
func (s *strSliceN) Any() bool {
	return len(s.v) > 0
}

// AnyContain checks if any items in this slice contain the target
func (s *strSliceN) AnyContain(target string) bool {
	for i := range s.v {
		if strings.Contains(s.v[i], target) {
			return true
		}
	}
	return false
}

// Append items to the end of the slice and return slice
func (s *strSliceN) Append(items ...string) *strSliceN {
	s.v = append(s.v, items...)
	return s
}

// At returns the item at the given index location. Allows for negative notation
func (s *strSliceN) At(i int) string {
	if i < 0 {
		i = len(s.v) + i
	}
	if i >= 0 && i < len(s.v) {
		return s.v[i]
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the underlying slice
func (s *strSliceN) Clear() *strSliceN {
	s.v = []string{}
	return s
}

// Contains checks if the given target is contained in this slice
func (s *strSliceN) Contains(target string) bool {
	for i := range s.v {
		if s.v[i] == target {
			return true
		}
	}
	return false
}

// ContainsAny checks if any of the targets are contained in this slice
func (s *strSliceN) ContainsAny(targets []string) bool {
	if targets != nil && len(targets) > 0 {
		for i := range targets {
			for j := range s.v {
				if s.v[j] == targets[i] {
					return true
				}
			}
		}
	}
	return false
}

// Del deletes item using neg/pos index notation with status
func (s *strSliceN) Del(i int) bool {
	result := false
	if i < 0 {
		i = len(s.v) + i
	}
	if i >= 0 && i < len(s.v) {
		if i+1 < len(s.v) {
			s.v = append(s.v[:i], s.v[i+1:]...)
			result = true
		} else {
			s.v = s.v[:i]
			result = true
		}
	}
	return result
}

// Drop deletes first n elements and returns the modified slice
func (s *strSliceN) Drop(cnt int) *strSliceN {
	if cnt > 0 {
		if len(s.v) >= cnt {
			s.v = s.v[cnt:]
		} else {
			s.v = []string{}
		}
	}
	return s
}

// Each iterates over the queryable and executes the given action
func (s *strSliceN) Each(action func(O)) {
	for i := range s.v {
		action(s.v[i])
	}
}

// Equals checks if the two slices are equal
func (s *strSliceN) Equals(other *strSliceN) bool {
	return reflect.DeepEqual(s, other)
}

// First returns the first time as a nub type
func (s *strSliceN) First() (result *strN) {
	if len(s.v) > 0 {
		result = A(s.v[0])
	} else {
		result = A("")
	}
	return
}

// Join the underlying slice with the given delim
func (s *strSliceN) Join(delim string) *strN {
	return A(strings.Join(s.v, delim))
}

// Last returns the last item as a nub type
func (s *strSliceN) Last() (result *strN) {
	if len(s.v) > 0 {
		result = A(s.At(-1))
	} else {
		result = A("")
	}
	return
}

// Len is a pass through to the underlying slice
func (s *strSliceN) Len() int {
	return len(s.v)
}

// Map manipulates the slice into a new form
func (s *strSliceN) Map(sel func(string) O) (result *Queryable) {
	for i := range s.v {
		obj := sel(s.v[i])

		// Drill into queryables
		if s, ok := obj.(*Queryable); ok {
			obj = s.v.Interface()
		}

		// Create new slice of the return type of sel
		if result == nil {
			typ := reflect.TypeOf(obj)
			result = Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
		}
		result.Append(obj)
	}
	if result == nil {
		result = Q([]interface{}{})
	}
	return
}

// MapF manipulates the queryable data into a new form then flattens
func (s *strSliceN) MapF(sel func(string) O) (result *Queryable) {
	result = s.Map(sel).Flatten()
	return
}

// Pair simply returns the first and second slice items
func (s *strSliceN) Pair() (first, second string) {
	if s.Len() > 0 {
		first = s.v[0]
	}
	if s.Len() > 1 {
		second = s.v[1]
	}
	return
}

// Prepend items to the begining of the slice and return slice
func (s *strSliceN) Prepend(items ...string) *strSliceN {
	items = append(items, s.v...)
	s.v = items
	return s
}

// Single simple report true if there is only one item
func (s *strSliceN) Single() (result bool) {
	return s.Len() == 1
}

// Slice provides a python like slice function for slice nubs.
// Has an inclusive behavior such that Slice(0, -1) includes index -1
// e.g. [1,2,3][0:-1] eq [1,2,3] and [1,2,3][1:2] eq [2,3]
// returns entire slice if indices are out of bounds
func (s *strSliceN) Slice(i, j int) (result *strSliceN) {

	// Convert to postive notation
	if i < 0 {
		i = s.Len() + i
	}
	if j < 0 {
		j = s.Len() + j
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= s.Len() {
		j = s.Len() - 1
	}

	// Specifically offsetting j to get an inclusive behavior out of Go
	j++

	// Only operate when indexes are within bounds
	// allow j to be len of s as that is how we include last item
	if i >= 0 && i < s.Len() && j >= 0 && j <= s.Len() {
		result = S(s.v[i:j]...)
	} else {
		result = S()
	}
	return
}

// Sort the underlying slice
func (s *strSliceN) Sort() *strSliceN {
	sort.Strings(s.v)
	return s
}

// TakeFirst updates the underlying slice and returns the item and status
func (s *strSliceN) TakeFirst() (string, bool) {
	if len(s.v) > 0 {
		item := s.v[0]
		s.v = s.v[1:]
		return item, true
	}
	return "", false
}

// TakeFirstCnt updates the underlying slice and returns the items
func (s *strSliceN) TakeFirstCnt(cnt int) (result *strSliceN) {
	if cnt > 0 {
		if len(s.v) >= cnt {
			result = S(s.v[:cnt]...)
			s.v = s.v[cnt:]
		} else {
			result = S(s.v...)
			s.v = []string{}
		}
	}
	return
}

// TakeLast updates the underlying slice and returns the item and status
func (s *strSliceN) TakeLast() (string, bool) {
	if len(s.v) > 0 {
		item := s.v[len(s.v)-1]
		s.v = s.v[:len(s.v)-1]
		return item, true
	}
	return "", false
}

// TakeLastCnt updates the underlying slice and returns a new nub
func (s *strSliceN) TakeLastCnt(cnt int) (result *strSliceN) {
	if cnt > 0 {
		if len(s.v) >= cnt {
			i := len(s.v) - cnt
			result = S(s.v[i:]...)
			s.v = s.v[:i]
		} else {
			result = S(s.v...)
			s.v = []string{}
		}
	}
	return
}

// Uniq removes all duplicates from the underlying slice
func (s *strSliceN) Uniq() *strSliceN {
	hits := map[string]bool{}
	for i := len(s.v) - 1; i >= 0; i-- {
		if _, exists := hits[s.v[i]]; !exists {
			hits[s.v[i]] = true
		} else {
			s.v = append(s.v[:i], s.v[i+1:]...)
		}
	}
	return s
}

// YamlPair return the first and second entries as yaml types
func (s *strSliceN) YamlPair() (first string, second interface{}) {
	if s.Len() > 0 {
		first = s.v[0]
	}
	if s.Len() > 1 {
		second = A(s.v[1]).YamlType()
	} else {
		second = nil
	}
	return
}

// YamlKeyVal return the first and second entries as KeyVal of yaml types
func (s *strSliceN) YamlKeyVal() KeyVal {
	result := KeyVal{}
	if s.Len() > 0 {
		result.Key = A(s.v[0]).YamlType()
	}
	if s.Len() > 1 {
		result.Val = A(s.v[1]).YamlType()
	} else {
		result.Val = ""
	}
	return result
}
