package n

import (
	"fmt"
	"reflect"
)

// NSlice provides a generic way to work with slice types providing convenience methods
// on par with other rapid development languages.
//
// Implements the Numerable interface.
type NSlice struct {
	o   interface{} // slice
	len int         // slice length
}

// Slice creates a new NSlice by simply storing slice obj directly to avoid using reflection
// processing at a 10x overhead savings. Non slice obj are encapsulated in a new slice of
// that type using reflection, thus incurring the standard 10x overhead.
func Slice(obj interface{}) (n *NSlice) {
	v := reflect.ValueOf(obj)

	switch v.Kind() {

	// Simply store the slice
	case reflect.Slice:
		n = &NSlice{o: obj, len: v.Len()}

	// Seed new slice with array values
	case reflect.Array:
		n = newSlice(obj)

	// Encapsulate any non-slice in a slice of that type
	default:
		n = SliceV(obj)
	}
	return
}

// SliceV creates a new NSlice encapsulating the given variadic elements in a new slice of
// that type. Incurs the full 10x reflection overhead. For large slice params use the Slice
// func instead.
func SliceV(items ...interface{}) *NSlice {
	return newSlice(items)
}

// handles []interface{} and arrays everything else will return nil
func newSlice(items interface{}) (n *NSlice) {
	n = &NSlice{}

	v := reflect.ValueOf(items)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:

		// Return nil numerable if nothing given
		if v.Len() == 0 {
			return
		}
		elem := v.Index(0).Interface()
		if elem == nil {
			return
		}

		// Create new slice with type of the element
		typ := reflect.SliceOf(reflect.TypeOf(elem))
		slice := reflect.MakeSlice(typ, 0, 10)

		// Add the variadic elements to the new slice
		for i := 0; i < v.Len(); i++ {
			item := reflect.ValueOf(v.Index(i).Interface())
			slice = reflect.Append(slice, item)
		}

		n.o = slice.Interface()
		n.len = slice.Len()
	}
	return
}

// Numerable interface methods
//--------------------------------------------------------------------------------------------------

// O returns the underlying data structure as is
func (n *NSlice) O() interface{} {
	if n.Nil() {
		return nil
	}
	return n.o
}

// Any tests if the numerable is not empty
func (n *NSlice) Any() bool {
	return n.len > 0
}

// Len returns the number of elements in the numerable
func (n *NSlice) Len() int {
	if n.Nil() {
		return 0
	}
	return n.len
}

// Nil tests if the numerable is nil
func (n *NSlice) Nil() bool {
	if n == nil || n.o == nil {
		return true
	}
	return false
}

// Type returns the identifier for this numerable type
func (n *NSlice) Type() NType {
	return NSliceType
}

// Export methods
//--------------------------------------------------------------------------------------------------

// // A exports NSlice as a string slice
// func (q *NSlice) A() (result []string) {
// 	// //if v, ok := q.o.([]string); ok {
// 	// //	result = v
// 	// //} else {
// 	// for i := 0; i < len(q.o); i++ {
// 	// 	if x, ok := q.o[i].(string); ok {
// 	// 		result = append(result, x)
// 	// 	} else {
// 	// 		result = append(result, fmt.Sprint(q.o[i]))
// 	// 	}
// 	// }
// 	// //}
// 	return
// }

// NSlice methods
//--------------------------------------------------------------------------------------------------

// Append an item to the end of the NSlice and returns the NSlice for chaining. Avoids the 10x
// reflection overhead cost by type asserting common types. Types not optimized in this way incur
// the full 10x reflection overhead cost.
//
// Optimized types: bool, int, string
func (n *NSlice) Append(item interface{}) *NSlice {
	if n.Nil() {
		new := SliceV(item)
		if !new.Nil() {
			*n = *new
		}
	} else {
		ok := false
		switch slice := n.o.(type) {
		case []bool:
			var x bool
			if x, ok = item.(bool); ok {
				n.o = append(slice, x)
			}
		case []int:
			var x int
			if x, ok = item.(int); ok {
				n.o = append(slice, x)
			}
		case []string:
			var x string
			if x, ok = item.(string); ok {
				n.o = append(slice, x)
			}
		default:
			ok = true
			v := reflect.ValueOf(n.o)
			item := reflect.ValueOf(item)
			n.o = reflect.Append(v, item).Interface()
		}
		if !ok {
			panic(fmt.Sprintf("can't insert type '%T' into '%T'", item, n.o))
		}
		n.len++
	}
	return n
}

// AppendV appends an item to the end of the NSlice and returns the NSlice for chaining. Avoids
// the 10x reflection overhead cost by type asserting common types. Types not optimized in this
// way incur the full 10x reflection overhead cost.
//
// Optimized types: bool, int, string
func (n *NSlice) AppendV(items ...interface{}) *NSlice {
	for _, item := range items {
		n.Append(item)
	}
	return n
}

// // AppendS appends the slice using variadic expansion and returns NSlice for chaining.  Avoids
// // the 10x reflection overhead cost by type asserting common types. Types not optimized in this
// // way incur the full 10x reflection overhead cost.
// //
// // Optimized types: []string
// func (q *NSlice) AppendS(items interface{}) *NSlice {
// 	if q.Nil() {
// 		*q = *(Slice(items))
// 	} else {
// 		ok := false
// 		switch slice := q.v.Interface().(type) {
// 		case []string:
// 			var x []string
// 			if x, ok = items.([]string); ok {
// 				slice = append(slice, x...)
// 			}
// 		default:
// 			x := reflect.ValueOf(items)
// 			*q.v = reflect.AppendSlice(*q.v, x)
// 		}
// 		if !ok {
// 			panic(fmt.Sprintf("can't concat type '%T' with '%T'", items, q.v.Interface()))
// 		}
// 	}
// 	return q
// }

// // At returns the item at the given index location. Allows for negative notation
// func (q *NSlice) At(i int) interface{} {
// 	if i < 0 {
// 		i = q.v.Len() + i
// 	}
// 	if i >= 0 && i < q.v.Len() {
// 		return q.v.Index(i).Interface()
// 	}
// 	panic("index out of slice bounds")
// }

// // AtQ returns a QObj for the item at the given index location. Allows for negative notation
// func (q *NSlice) AtQ(i int) *QObj {
// 	if i < 0 {
// 		i = q.v.Len() + i
// 	}
// 	if i >= 0 && i < q.v.Len() {
// 		return &QObj{v: q.v.Index(i).Interface()}
// 	}
// 	panic("index out of slice bounds")
// }

// // // Clear the underlying slice
// // func (s *strSliceN) Clear() *strSliceN {
// // 	s.v = []string{}
// // 	return s
// // }

// // // AnyContain checks if any items in this slice contain the target
// // func (q *NSlice) AnyContain(target string) bool {
// // 	for i := range q.v {
// // 		if strings.Contains(q.v[i], target) {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // // Contains checks if the given target is contained in this slice
// // func (s *strSliceN) Contains(target string) bool {
// // 	for i := range s.v {
// // 		if s.v[i] == target {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // // ContainsAny checks if any of the targets are contained in this slice
// // func (s *strSliceN) ContainsAny(targets []string) bool {
// // 	if targets != nil && len(targets) > 0 {
// // 		for i := range targets {
// // 			for j := range s.v {
// // 				if s.v[j] == targets[i] {
// // 					return true
// // 				}
// // 			}
// // 		}
// // 	}
// // 	return false
// // }

// // // Del deletes item using neg/pos index notation with status
// // func (s *strSliceN) Del(i int) bool {
// // 	result := false
// // 	if i < 0 {
// // 		i = len(s.v) + i
// // 	}
// // 	if i >= 0 && i < len(s.v) {
// // 		if i+1 < len(s.v) {
// // 			s.v = append(s.v[:i], s.v[i+1:]...)
// // 			result = true
// // 		} else {
// // 			s.v = s.v[:i]
// // 			result = true
// // 		}
// // 	}
// // 	return result
// // }

// // // Drop deletes first n elements and returns the modified slice
// // func (s *strSliceN) Drop(cnt int) *strSliceN {
// // 	if cnt > 0 {
// // 		if len(s.v) >= cnt {
// // 			s.v = s.v[cnt:]
// // 		} else {
// // 			s.v = []string{}
// // 		}
// // 	}
// // 	return s
// // }

// // // Each iterates over the numerable and executes the given action
// // func (s *strSliceN) Each(action func(O)) {
// // 	for i := range s.v {
// // 		action(s.v[i])
// // 	}
// // }

// // // Equals checks if the two slices are equal
// // func (s *strSliceN) Equals(other *strSliceN) bool {
// // 	return reflect.DeepEqual(s, other)
// // }

// // // First returns the first item or nil
// // func (q *NSlice) First() (result interface{}) {
// // 	if !q.Nil() && q.v.Len() > 0 {
// // 		result = q.v.Index(0).Interface()
// // 	}
// // 	return
// // }

// // // Join the underlying slice with the given delim
// // func (s *strSliceN) Join(delim string) *strN {
// // 	return A(strings.Join(s.v, delim))
// // }

// // // Last returns the last item as a nub type
// // func (s *strSliceN) Last() (result *strN) {
// // 	if len(s.v) > 0 {
// // 		result = A(s.At(-1))
// // 	} else {
// // 		result = A("")
// // 	}
// // 	return
// // }

// // // Map manipulates the slice into a new form
// // func (s *strSliceN) Map(sel func(string) O) (result *Numerable) {
// // 	for i := range s.v {
// // 		obj := sel(s.v[i])

// // 		// Drill into numerables
// // 		if s, ok := obj.(*Numerable); ok {
// // 			obj = s.v.Interface()
// // 		}

// // 		// Create new slice of the return type of sel
// // 		if result == nil {
// // 			typ := reflect.TypeOf(obj)
// // 			result = Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
// // 		}
// // 		result.Append(obj)
// // 	}
// // 	if result == nil {
// // 		result = Q([]interface{}{})
// // 	}
// // 	return
// // }

// // // MapF manipulates the numerable data into a new form then flattens
// // func (s *strSliceN) MapF(sel func(string) O) (result *Numerable) {
// // 	result = s.Map(sel).Flatten()
// // 	return
// // }

// // // Pair simply returns the first and second slice items
// // func (s *strSliceN) Pair() (first, second string) {
// // 	if s.Len() > 0 {
// // 		first = s.v[0]
// // 	}
// // 	if s.Len() > 1 {
// // 		second = s.v[1]
// // 	}
// // 	return
// // }

// // // Prepend items to the begining of the slice and return slice
// // func (s *strSliceN) Prepend(items ...string) *strSliceN {
// // 	items = append(items, s.v...)
// // 	s.v = items
// // 	return s
// // }

// // // Single simple report true if there is only one item
// // func (s *strSliceN) Single() (result bool) {
// // 	return s.Len() == 1
// // }

// // // Slice provides a python like slice function for slice nubs.
// // // Has an inclusive behavior such that Slice(0, -1) includes index -1
// // // e.g. [1,2,3][0:-1] eq [1,2,3] and [1,2,3][1:2] eq [2,3]
// // // returns entire slice if indices are out of bounds
// // func (s *strSliceN) Slice(i, j int) (result *strSliceN) {

// // 	// Convert to postive notation
// // 	if i < 0 {
// // 		i = s.Len() + i
// // 	}
// // 	if j < 0 {
// // 		j = s.Len() + j
// // 	}

// // 	// Move start/end within bounds
// // 	if i < 0 {
// // 		i = 0
// // 	}
// // 	if j >= s.Len() {
// // 		j = s.Len() - 1
// // 	}

// // 	// Specifically offsetting j to get an inclusive behavior out of Go
// // 	j++

// // 	// Only operate when indexes are within bounds
// // 	// allow j to be len of s as that is how we include last item
// // 	if i >= 0 && i < s.Len() && j >= 0 && j <= s.Len() {
// // 		result = S(s.v[i:j]...)
// // 	} else {
// // 		result = S()
// // 	}
// // 	return
// // }

// // // Sort the underlying slice
// // func (s *strSliceN) Sort() *strSliceN {
// // 	sort.Strings(s.v)
// // 	return s
// // }

// // // TakeFirst updates the underlying slice and returns the item and status
// // func (s *strSliceN) TakeFirst() (string, bool) {
// // 	if len(s.v) > 0 {
// // 		item := s.v[0]
// // 		s.v = s.v[1:]
// // 		return item, true
// // 	}
// // 	return "", false
// // }

// // // TakeFirstCnt updates the underlying slice and returns the items
// // func (s *strSliceN) TakeFirstCnt(cnt int) (result *strSliceN) {
// // 	if cnt > 0 {
// // 		if len(s.v) >= cnt {
// // 			result = S(s.v[:cnt]...)
// // 			s.v = s.v[cnt:]
// // 		} else {
// // 			result = S(s.v...)
// // 			s.v = []string{}
// // 		}
// // 	}
// // 	return
// // }

// // // TakeLast updates the underlying slice and returns the item and status
// // func (s *strSliceN) TakeLast() (string, bool) {
// // 	if len(s.v) > 0 {
// // 		item := s.v[len(s.v)-1]
// // 		s.v = s.v[:len(s.v)-1]
// // 		return item, true
// // 	}
// // 	return "", false
// // }

// // // TakeLastCnt updates the underlying slice and returns a new nub
// // func (s *strSliceN) TakeLastCnt(cnt int) (result *strSliceN) {
// // 	if cnt > 0 {
// // 		if len(s.v) >= cnt {
// // 			i := len(s.v) - cnt
// // 			result = S(s.v[i:]...)
// // 			s.v = s.v[:i]
// // 		} else {
// // 			result = S(s.v...)
// // 			s.v = []string{}
// // 		}
// // 	}
// // 	return
// // }

// // // Uniq removes all duplicates from the underlying slice
// // func (s *strSliceN) Uniq() *strSliceN {
// // 	hits := map[string]bool{}
// // 	for i := len(s.v) - 1; i >= 0; i-- {
// // 		if _, exists := hits[s.v[i]]; !exists {
// // 			hits[s.v[i]] = true
// // 		} else {
// // 			s.v = append(s.v[:i], s.v[i+1:]...)
// // 		}
// // 	}
// // 	return s
// // }

// // // YamlPair return the first and second entries as yaml types
// // func (s *strSliceN) YamlPair() (first string, second interface{}) {
// // 	if s.Len() > 0 {
// // 		first = s.v[0]
// // 	}
// // 	if s.Len() > 1 {
// // 		second = A(s.v[1]).YamlType()
// // 	} else {
// // 		second = nil
// // 	}
// // 	return
// // }

// // // YamlKeyVal return the first and second entries as KeyVal of yaml types
// // func (s *strSliceN) YamlKeyVal() KeyVal {
// // 	result := KeyVal{}
// // 	if s.Len() > 0 {
// // 		result.Key = A(s.v[0]).YamlType()
// // 	}
// // 	if s.Len() > 1 {
// // 		result.Val = A(s.v[1]).YamlType()
// // 	} else {
// // 		result.Val = ""
// // 	}
// // 	return result
// // }

// // check if the internal type is a slice type
// func (q *NSlice) isSliceType() bool {
// 	return q.k == reflect.Array || q.k == reflect.Slice
// }
