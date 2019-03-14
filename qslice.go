package n

import (
	"fmt"
	"reflect"
)

// QSlice implementes the Queryable Interface and supports generic slice types.
// It provides a plethora of convenience methods to work with slice types.
type QSlice struct {
	v    *reflect.Value // underlying value
	kind reflect.Kind   // kind of the underlying value
}

// S instantiates a new QSlice optionally seeding it with the given
// with the given obj which may be a scalar value or a slice value.
// Scalar values will be encapsulated in a new slice of that type.
func S(obj interface{}) *QSlice {
	v := reflect.ValueOf(obj)
	q := &QSlice{v: &v, kind: v.Kind()}

	//Switch:
	switch q.kind {
	case reflect.Array, reflect.Slice:
		if q.v.IsNil() {
			*q.v = reflect.MakeSlice(q.v.Type(), 0, 10)
		}
	case reflect.Ptr:
		// pv := reflect.Indirect(reflect.ValueOf(obj))
		// pq := &OldQueryable{v: &pv, Kind: pv.Kind()}
		// switch nq.Kind {
		// case reflect.Array, reflect.Slice, reflect.Map, reflect.String, reflect.Chan:
		// 	v = nv
		// 	q = nq
		// 	goto Switch
		// }

	// Create a new slice of the given obj type and add the obj
	default:
		*q.v = reflect.MakeSlice(q.v.Type(), 0, 10)
		*q.v = reflect.Append(*q.v, *q.v)
	}
	return q
}

// Type returns the identifier for this queryable type.
// Implements the queryable interface.
func (q *QSlice) Type() QType {
	return QSliceType
}

// O returns the underlying data structure
// Implements the queryable interface.
func (q *QSlice) O() interface{} {
	if q.Nil() {
		return nil
	}
	return q.v.Interface()
}

// S exports QSlice into an string slice
func (q *QSlice) S() (result []string) {
	if !q.Nil() && q.isSliceType() {
		if v, ok := q.O().([]string); ok {
			result = v
		} else {
			for i := 0; i < q.v.Len(); i++ {
				item := q.v.Index(i).Interface()
				result = append(result, fmt.Sprint(item))
			}
		}
	}
	return
}

// // Any checks if the slice has anything in it
// func (s *strSliceN) Any() bool {
// 	return len(s.v) > 0
// }

// // AnyContain checks if any items in this slice contain the target
// func (s *strSliceN) AnyContain(target string) bool {
// 	for i := range s.v {
// 		if strings.Contains(s.v[i], target) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Append items to the end of the slice and return slice
// func (s *strSliceN) Append(items ...string) *strSliceN {
// 	s.v = append(s.v, items...)
// 	return s
// }

// // At returns the item at the given index location. Allows for negative notation
// func (s *strSliceN) At(i int) string {
// 	if i < 0 {
// 		i = len(s.v) + i
// 	}
// 	if i >= 0 && i < len(s.v) {
// 		return s.v[i]
// 	}
// 	panic(errors.New("Index out of slice bounds"))
// }

// // Clear the underlying slice
// func (s *strSliceN) Clear() *strSliceN {
// 	s.v = []string{}
// 	return s
// }

// // Contains checks if the given target is contained in this slice
// func (s *strSliceN) Contains(target string) bool {
// 	for i := range s.v {
// 		if s.v[i] == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// // ContainsAny checks if any of the targets are contained in this slice
// func (s *strSliceN) ContainsAny(targets []string) bool {
// 	if targets != nil && len(targets) > 0 {
// 		for i := range targets {
// 			for j := range s.v {
// 				if s.v[j] == targets[i] {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// // Del deletes item using neg/pos index notation with status
// func (s *strSliceN) Del(i int) bool {
// 	result := false
// 	if i < 0 {
// 		i = len(s.v) + i
// 	}
// 	if i >= 0 && i < len(s.v) {
// 		if i+1 < len(s.v) {
// 			s.v = append(s.v[:i], s.v[i+1:]...)
// 			result = true
// 		} else {
// 			s.v = s.v[:i]
// 			result = true
// 		}
// 	}
// 	return result
// }

// // Drop deletes first n elements and returns the modified slice
// func (s *strSliceN) Drop(cnt int) *strSliceN {
// 	if cnt > 0 {
// 		if len(s.v) >= cnt {
// 			s.v = s.v[cnt:]
// 		} else {
// 			s.v = []string{}
// 		}
// 	}
// 	return s
// }

// // Each iterates over the queryable and executes the given action
// func (s *strSliceN) Each(action func(O)) {
// 	for i := range s.v {
// 		action(s.v[i])
// 	}
// }

// // Equals checks if the two slices are equal
// func (s *strSliceN) Equals(other *strSliceN) bool {
// 	return reflect.DeepEqual(s, other)
// }

// First returns the first item or nil
func (q *QSlice) First() (result interface{}) {
	if !q.Nil() && q.v.Len() > 0 {
		result = q.v.Index(0).Interface()
	}
	return
}

// // Join the underlying slice with the given delim
// func (s *strSliceN) Join(delim string) *strN {
// 	return A(strings.Join(s.v, delim))
// }

// // Last returns the last item as a nub type
// func (s *strSliceN) Last() (result *strN) {
// 	if len(s.v) > 0 {
// 		result = A(s.At(-1))
// 	} else {
// 		result = A("")
// 	}
// 	return
// }

// Len of the collection type
func (q *QSlice) Len() int {
	if q.Nil() {
		return 0
	}
	return q.v.Len()
}

// // Map manipulates the slice into a new form
// func (s *strSliceN) Map(sel func(string) O) (result *Queryable) {
// 	for i := range s.v {
// 		obj := sel(s.v[i])

// 		// Drill into queryables
// 		if s, ok := obj.(*Queryable); ok {
// 			obj = s.v.Interface()
// 		}

// 		// Create new slice of the return type of sel
// 		if result == nil {
// 			typ := reflect.TypeOf(obj)
// 			result = Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
// 		}
// 		result.Append(obj)
// 	}
// 	if result == nil {
// 		result = Q([]interface{}{})
// 	}
// 	return
// }

// // MapF manipulates the queryable data into a new form then flattens
// func (s *strSliceN) MapF(sel func(string) O) (result *Queryable) {
// 	result = s.Map(sel).Flatten()
// 	return
// }

// Nil tests if the queryable is nil.
// Implements Queryable interface
func (q *QSlice) Nil() bool {
	if q == nil || q.v == nil || q.kind == reflect.Invalid {
		return true
	}
	return false
}

// // Pair simply returns the first and second slice items
// func (s *strSliceN) Pair() (first, second string) {
// 	if s.Len() > 0 {
// 		first = s.v[0]
// 	}
// 	if s.Len() > 1 {
// 		second = s.v[1]
// 	}
// 	return
// }

// // Prepend items to the begining of the slice and return slice
// func (s *strSliceN) Prepend(items ...string) *strSliceN {
// 	items = append(items, s.v...)
// 	s.v = items
// 	return s
// }

// // Single simple report true if there is only one item
// func (s *strSliceN) Single() (result bool) {
// 	return s.Len() == 1
// }

// // Slice provides a python like slice function for slice nubs.
// // Has an inclusive behavior such that Slice(0, -1) includes index -1
// // e.g. [1,2,3][0:-1] eq [1,2,3] and [1,2,3][1:2] eq [2,3]
// // returns entire slice if indices are out of bounds
// func (s *strSliceN) Slice(i, j int) (result *strSliceN) {

// 	// Convert to postive notation
// 	if i < 0 {
// 		i = s.Len() + i
// 	}
// 	if j < 0 {
// 		j = s.Len() + j
// 	}

// 	// Move start/end within bounds
// 	if i < 0 {
// 		i = 0
// 	}
// 	if j >= s.Len() {
// 		j = s.Len() - 1
// 	}

// 	// Specifically offsetting j to get an inclusive behavior out of Go
// 	j++

// 	// Only operate when indexes are within bounds
// 	// allow j to be len of s as that is how we include last item
// 	if i >= 0 && i < s.Len() && j >= 0 && j <= s.Len() {
// 		result = S(s.v[i:j]...)
// 	} else {
// 		result = S()
// 	}
// 	return
// }

// // Sort the underlying slice
// func (s *strSliceN) Sort() *strSliceN {
// 	sort.Strings(s.v)
// 	return s
// }

// // TakeFirst updates the underlying slice and returns the item and status
// func (s *strSliceN) TakeFirst() (string, bool) {
// 	if len(s.v) > 0 {
// 		item := s.v[0]
// 		s.v = s.v[1:]
// 		return item, true
// 	}
// 	return "", false
// }

// // TakeFirstCnt updates the underlying slice and returns the items
// func (s *strSliceN) TakeFirstCnt(cnt int) (result *strSliceN) {
// 	if cnt > 0 {
// 		if len(s.v) >= cnt {
// 			result = S(s.v[:cnt]...)
// 			s.v = s.v[cnt:]
// 		} else {
// 			result = S(s.v...)
// 			s.v = []string{}
// 		}
// 	}
// 	return
// }

// // TakeLast updates the underlying slice and returns the item and status
// func (s *strSliceN) TakeLast() (string, bool) {
// 	if len(s.v) > 0 {
// 		item := s.v[len(s.v)-1]
// 		s.v = s.v[:len(s.v)-1]
// 		return item, true
// 	}
// 	return "", false
// }

// // TakeLastCnt updates the underlying slice and returns a new nub
// func (s *strSliceN) TakeLastCnt(cnt int) (result *strSliceN) {
// 	if cnt > 0 {
// 		if len(s.v) >= cnt {
// 			i := len(s.v) - cnt
// 			result = S(s.v[i:]...)
// 			s.v = s.v[:i]
// 		} else {
// 			result = S(s.v...)
// 			s.v = []string{}
// 		}
// 	}
// 	return
// }

// // Uniq removes all duplicates from the underlying slice
// func (s *strSliceN) Uniq() *strSliceN {
// 	hits := map[string]bool{}
// 	for i := len(s.v) - 1; i >= 0; i-- {
// 		if _, exists := hits[s.v[i]]; !exists {
// 			hits[s.v[i]] = true
// 		} else {
// 			s.v = append(s.v[:i], s.v[i+1:]...)
// 		}
// 	}
// 	return s
// }

// // YamlPair return the first and second entries as yaml types
// func (s *strSliceN) YamlPair() (first string, second interface{}) {
// 	if s.Len() > 0 {
// 		first = s.v[0]
// 	}
// 	if s.Len() > 1 {
// 		second = A(s.v[1]).YamlType()
// 	} else {
// 		second = nil
// 	}
// 	return
// }

// // YamlKeyVal return the first and second entries as KeyVal of yaml types
// func (s *strSliceN) YamlKeyVal() KeyVal {
// 	result := KeyVal{}
// 	if s.Len() > 0 {
// 		result.Key = A(s.v[0]).YamlType()
// 	}
// 	if s.Len() > 1 {
// 		result.Val = A(s.v[1]).YamlType()
// 	} else {
// 		result.Val = ""
// 	}
// 	return result
// }

func (q *QSlice) isSliceType() bool {
	return q.kind == reflect.Array || q.kind == reflect.Slice
}
