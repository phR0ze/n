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
	o   interface{} // underlying slice object
	len int         // slice length
}

// Slice creates a new NSlice by simply storing slice obj directly to avoid using reflection
// processing at a 10x overhead savings. Non slice obj are encapsulated in a new slice of
// that type using reflection, thus incurring the standard 10x overhead.
//
// Return value n *NSlice will never be nil but n.Nil() may be true as nil or empty []interface{}
// values are ignored to avoid internally using a []interface{}. The internal type will be
// set later with the given type when an n.AppendX method is called.
//
// Cost: ~1x - 10x
func Slice(obj interface{}) (n *NSlice) {
	n = &NSlice{}
	v := reflect.ValueOf(obj)

	k := v.Kind()
	x, interfaceSliceType := obj.([]interface{})
	switch {

	// Return the NSlice.Nil
	case k == reflect.Invalid:

	// Iterate over array and append
	case k == reflect.Array:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			n.Append(item)
		}

	// Convert []interface to slice of elem type
	case interfaceSliceType:
		n = SliceV(x...)

	// Slice of distinct type can be used directly
	case k == reflect.Slice:
		n.o = obj
		n.len = v.Len()

	// Append single items
	default:
		n.Append(obj)
	}
	return
}

// SliceV creates a new NSlice encapsulating the given variadic elements in a new slice of
// that type. Incurs the full 10x reflection overhead. For large slice params use the Slice
// func instead.
//
// Cost: ~10x
func SliceV(items ...interface{}) (n *NSlice) {
	n = &NSlice{}

	// Return NSlice.Nil if nothing given
	if len(items) == 0 {
		return
	}

	// Create new slice from the type of the first non Invalid element
	var slice *reflect.Value
	for i := 0; i < len(items); i++ {

		// Create target slice from first Valid element
		if slice == nil && reflect.ValueOf(items[i]).IsValid() {
			typ := reflect.SliceOf(reflect.TypeOf(items[i]))
			v := reflect.MakeSlice(typ, 0, 10)
			slice = &v
		}

		// Append item to slice
		if slice != nil {
			item := reflect.ValueOf(items[i])
			*slice = reflect.Append(*slice, item)
		}
	}
	if slice != nil {
		n.o = slice.Interface()
		n.len = slice.Len()
	}
	return
}

// create a new empty slice of the given type or element type if a slice/array
// want to return a new *NSlice so that we can use this in the AppendX functions
// to defer creating an underlying slice type until we have an actual type to work with.
func newEmptySlice(items interface{}) (n *NSlice) {
	n = &NSlice{}
	v := reflect.ValueOf(items)
	typ := reflect.TypeOf([]interface{}{})

	k := v.Kind()
	switch k {

	// Use a new generic slice for nils
	case reflect.Invalid:

	// Use the element type of slice/arrays
	case reflect.Slice, reflect.Array:

		// Use slice type if not generic
		if _, ok := items.([]interface{}); !ok {
			typ = reflect.SliceOf(v.Type().Elem())
		} else {
			// For generics try to find actual element type
			if v.Len() != 0 {
				elem := v.Index(0).Interface()
				if elem != nil {
					typ = reflect.SliceOf(reflect.TypeOf(elem))
				}
			}
		}
	default:
		typ = reflect.SliceOf(v.Type())
	}

	// Create new slice with type of the element
	slice := reflect.MakeSlice(typ, 0, 10)
	n.o = slice.Interface()
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

// Contains checks if the given obj is contained in this slice
func (n *NSlice) Contains(obj interface{}) (result bool) {
	if n.Nil() {
		return
	}
	ok := false
	switch slice := n.o.(type) {
	case []bool:
		var x bool
		if x, ok = obj.(bool); ok {
			for i := range slice {
				if slice[i] == x {
					return true
				}
			}
		}
	case []int:
		var x int
		if x, ok = obj.(int); ok {
			for i := range slice {
				if slice[i] == x {
					return true
				}
			}
		}
	case []string:
		var x string
		if x, ok = obj.(string); ok {
			for i := range slice {
				if slice[i] == x {
					return true
				}
			}
		}
	default:
		v := reflect.ValueOf(n.o)
		x := reflect.ValueOf(obj)
		if v.Type().Elem() == x.Type() {
			ok = true
			for i := 0; i < v.Len(); i++ {
				if v.Index(i).Interface() == obj {
					return true
				}
			}
		}
	}
	if !ok {
		panic(fmt.Sprintf("can't compare type '%T' with '%T' elements", obj, n.o))
	}
	return
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
func (n *NSlice) Type() Type {
	return NSliceType
}

// Export methods
//--------------------------------------------------------------------------------------------------

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
// Cost: ~4x - 10x
//
// Optimized types: bool, int, string
func (n *NSlice) Append(item interface{}) *NSlice {
	if n.Nil() {
		*n = *(newEmptySlice(item))
	}
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
	return n
}

// AppendV appends an item to the end of the NSlice and returns the NSlice for chaining. Avoids
// the 10x reflection overhead cost by type asserting common types. Types not optimized in this
// way incur the full 10x reflection overhead cost.
//
// Cost: ~6x - 10x
//
// Optimized types: bool, int, string
func (n *NSlice) AppendV(items ...interface{}) *NSlice {
	for _, item := range items {
		n.Append(item)
	}
	return n
}

// AppendS appends the given slice using variadic expansion and returns NSlice for chaining. Avoids
// the 10x reflection overhead cost by type asserting common types. Types not optimized in this
// way incur the full 10x reflection overhead cost. However when appending larger slices fewer times
// the cost reduces down to 2x.
//
// Cost: ~1x - 2x
//
// Optimized types: []bool, []int, []string
func (n *NSlice) AppendS(items interface{}) *NSlice {
	if n.Nil() {
		*n = *(newEmptySlice(items))
	}
	ok := false
	switch slice := n.o.(type) {
	case []bool:
		var x []bool
		if x, ok = items.([]bool); ok {
			n.o = append(slice, x...)
			n.len += len(x)
		}
	case []int:
		var x []int
		if x, ok = items.([]int); ok {
			n.o = append(slice, x...)
			n.len += len(x)
		}
	case []string:
		var x []string
		if x, ok = items.([]string); ok {
			n.o = append(slice, x...)
			n.len += len(x)
		}
	default:
		ok = true
		v := reflect.ValueOf(n.o)
		x := reflect.ValueOf(items)
		n.o = reflect.AppendSlice(v, x).Interface()
		n.len += x.Len()
	}
	if !ok {
		panic(fmt.Sprintf("can't concat type '%T' with '%T'", items, n.o))
	}
	return n
}

// At returns the item at the given index location. Allows for negative notation.
// Cost even for reflection in this case doesn't seem to to add much.
//
// Cost: ~20% - 2x
func (n *NSlice) At(i int) *NObj {
	if i < 0 {
		i = n.len + i
	}
	if i >= 0 && i < n.len {
		switch slice := n.o.(type) {
		case []bool:
			return &NObj{slice[i]}
		case []int:
			return &NObj{slice[i]}
		case []string:
			return &NObj{slice[i]}
		default:
			return &NObj{reflect.ValueOf(n.o).Index(i).Interface()}
		}
	}
	panic("index out of slice bounds")
}

// Clear the underlying slice.
//
// Cost: constant
func (n *NSlice) Clear() *NSlice {
	if !n.Nil() {
		*n = *(newEmptySlice(n.o))
	}
	return n
}

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
