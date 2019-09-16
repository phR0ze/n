package n

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Map provides a generic way to work with map types providing convenience methods
// on par with rapid development languages. 'this Map' refers to the current map
// instance being operated on.  'new Map' refers to a copy of the map.
type Map interface {
	Any(keys ...interface{}) bool // Any tests if this Map is not empty or optionally if it contains any of the given variadic keys.
	// AnyS(slice interface{}) bool                      // AnyS tests if this Map contains any of the given Slice's elements.
	// AnyW(sel func(O) bool) bool                       // AnyW tests if this Map contains any that match the lambda selector.
	// Append(elem interface{}) Slice                    // Append an element to the end of this Map and returns a reference to this Map.
	// AppendV(elems ...interface{}) Slice               // AppendV appends the variadic elements to the end of this Map and returns a reference to this Map.
	Clear() Map // Clear modifies this Map to clear out all key-value pairs and returns a reference to this Map.
	// Concat(slice interface{}) (new Slice)             // Concat returns a new Slice by appending the given Slice to this Map using variadic expansion.
	// ConcatM(slice interface{}) Slice                  // ConcatM modifies this Map by appending the given Slice using variadic expansion and returns a reference to this Map.
	Copy(keys ...interface{}) (new Map) // Copy returns a new Map with the indicated key-value pairs copied from this Map or all if not given.
	// Count(elem interface{}) (cnt int)                 // Count the number of elements in this Map equal to the given element.
	// CountW(sel func(O) bool) (cnt int)                // CountW counts the number of elements in this Map that match the lambda selector.
	Delete(key interface{}) (val *Object) // Delete modifies this Map to delete the indicated key-value pair and returns the value from the Map.
	DeleteM(key interface{}) Map          // DeleteM modifies this Map to delete the indicated key-value pair and returns a reference to this Map rather than the key-value pair.
	//DeleteS(keys interface{}) (obj *Object) // DeleteS modifies this Map to delete the indicated key-value pairs and returns the values from the Map as a Slice.
	Exists(key interface{}) bool // Exists checks if the given key exists in this Map.
	// DeleteW(sel func(O) bool) Slice                     // DropW modifies this Map to delete the elements that match the lambda selector and returns a reference to this Map.
	// Each(action func(O)) Slice                        // Each calls the given lambda once for each element in this Map, passing in that element
	// EachE(action func(O) error) (ISlice, error)        // EachE calls the given lambda once for each element in this Map, passing in that element
	// EachI(action func(int, O)) Slice                  // EachI calls the given lambda once for each element in this Map, passing in the index and element
	// EachIE(action func(int, O) error) (ISlice, error)  // EachIE calls the given lambda once for each element in this Map, passing in the index and element
	// EachR(action func(O)) Slice                       // EachR calls the given lambda once for each element in this Map in reverse, passing in that element
	// EachRE(action func(O) error) (ISlice, error)       // EachRE calls the given lambda once for each element in this Map in reverse, passing in that element
	// EachRI(action func(int, O)) Slice                 // EachRI calls the given lambda once for each element in this Map in reverse, passing in that element
	// EachRIE(action func(int, O) error) (ISlice, error) // EachRIE calls the given lambda once for each element in this Map in reverse, passing in that element
	// Empty() bool                                      // Empty tests if this Map is empty.
	Generic() bool                                          // Generic returns true if the underlying implementation uses reflection
	Get(key interface{}) (val *Object)                      // Get returns the value at the given key location. Returns empty *Object if not found.
	Inject(key string, val interface{}) Map                 // Inject sets the value for the given key location, using jq type selectors. Returns a reference to this Map.
	InjectE(key string, val interface{}) (m Map, err error) // InjectE sets the value for the given key location, using jq type selectors. Returns a reference to this Map.
	// Join(separator ...string) (str *Object)           // Join converts each element into a string then joins them together using the given separator or comma by default.
	Keys() ISlice                        // Keys returns all the keys in this Map as a Slice of the key type.
	Len() int                            // Len returns the number of elements in this Map.
	M() (m *StringMap)                   // M is an alias to ToStringMap
	MG() (m map[string]interface{})      // MG is an alias to ToStringMapG
	Merge(m Map, location ...string) Map // Merge modifies this Map by overriding its values at location with the given map where they both exist and returns a reference to this Map.
	// Less(i, j int) bool                               // Less returns true if the element indexed by i is less than the element indexed by j.
	// Nil() bool                                        // Nil tests if this Map is nil.
	O() interface{} // O returns the underlying data structure as is.
	// Pair() (first, second *Object)                    // Pair simply returns the first and second Slice elements as Objects.
	// Pop() (elem *Object)                              // Pop modifies this Map to remove the last element and returns the removed element as an Object.
	// PopN(n int) (new Map)                           // PopN modifies this Map to remove the last n elements and returns the removed elements as a new Map.
	// Prepend(elem interface{}) Slice                   // Prepend modifies this Map to add the given element at the begining and returns a reference to this Map.
	Query(key string) (val *Object)             // Query returns the value at the given key location, using jq type selectors. Returns empty *Object if not found.
	QueryE(key string) (val *Object, err error) // Query returns the value at the given key location, using jq type selectors. Returns empty *Object if not found.
	Remove(key string) Map                      // Remove modifies this map to remove the value at the given key location, using jq type selectors. Returns a reference to this Map
	RemoveE(key string) (m Map, err error)      // RemoveE modifies this map to remove the value at the given key location, using jq type selectors. Returns a reference to this Map
	// Reverse() (new Map)                             // Reverse returns a new Map with the order of the elements reversed.
	// ReverseM() Slice                                  // ReverseM modifies this Map reversing the order of the elements and returns a reference to this Map.
	// Select(sel func(O) bool) (new Map)              // Select creates a new Map with the elements that match the lambda selector.
	Set(key, val interface{}) bool // Set the value for the given key to the given val. Returns true if the key did not yet exists in this Map.
	SetM(key, val interface{}) Map // SetM the value for the given key to the given val creating map if necessary. Returns a reference to this Map.
	// Shift() (elem *Object)                            // Shift modifies this Map to remove the first element and returns the removed element as an Object.
	// ShiftN(n int) (new Map)                         // ShiftN modifies this Map to remove the first n elements and returns the removed elements as a new Map.
	// Single() bool                                     // Single reports true if there is only one element in this Map.
	// Slice(indices ...int) Slice                       // Slice returns a range of elements from this Map as a Slice reference to the original. Allows for negative notation.
	// Sort() (new Map)                                // Sort returns a new Map with sorted elements.
	// SortM() Slice                                     // SortM modifies this Map sorting the elements and returns a reference to this Map.
	// SortReverse() (new Map)                         // SortReverse returns a new Map sorting the elements in reverse.
	// SortReverseM() Slice                              // SortReverseM modifies this Map sorting the elements in reverse and returns a reference to this Map.
	// String() string                                   // Returns a string representation of this Map, implements the Stringer interface
	// Swap(i, j int)                                    // Swap modifies this Map swapping the indicated elements.
	ToStringMap() (m *StringMap)              // ToStringMap converts the map to a *StringMap
	ToStringMapG() (m map[string]interface{}) // ToStringMapG converts the map to a Golang map[string]interface{}
	// Take(indices ...int) (new Map)                  // Take modifies this Map removing the indicated range of elements from this Map and returning them as a new Map.
	// TakeAt(i int) (elem *Object)                      // TakeAt modifies this Map removing the elemement at the given index location and returns the removed element as an Object.
	// TakeW(sel func(O) bool) (new Map)               // TakeW modifies this Map removing the elements that match the lambda selector and returns them as a new Map.
	// Union(slice interface{}) (new Map)              // Union returns a new Map by joining uniq elements from this Map with uniq elements from the given Slice while preserving order.
	// UnionM(slice interface{}) Slice                   // UnionM modifies this Map by joining uniq elements from this Map with uniq elements from the given Slice while preserving order.
	// Uniq() (new Map)                                // Uniq returns a new Map with all non uniq elements removed while preserving element order.
	// UniqM() Slice                                     // UniqM modifies this Map to remove all non uniq elements while preserving element order.
	YAML() (data string)                   // YAML converts the Map into a YAML string
	YAMLE() (data string, err error)       // YAMLE converts the Map into a YAML string
	WriteJSON(filename string) (err error) // WriteJSON converts the Map into a map[string]interface{} then calls json.WriteJSON on it to write it out to disk.
	WriteYAML(filename string) (err error) // WriteYAML converts the Map into a map[string]interface{} then calls yaml.WriteYAML on it to write it out to disk.
}

// NewMap provides a generic way to work with Map types. It does this by wrapping Go types
// directly for optimized types thus avoiding reflection processing overhead and making a plethora
// of Map methods available. Non optimized types will fall back on reflection to generically
// handle the type incurring the full 10x reflection processing overhead. Defaults to StringMap
// type if nothing is given.
//
// Optimized: map[string]interface{}
func NewMap(obj interface{}) (new Map) {
	o := Reference(obj)
	switch x := o.(type) {

	// StringMap
	// ---------------------------------------------------------------------------------------------
	case []byte, *[]byte, string, *string, *StringMap, *map[string]interface{}, *map[string]string, *map[string]float32,
		*map[string]float64, *map[string]int, *map[string]int64:
		new, _ = ToStringMapE(x)

	// RefMap
	// ---------------------------------------------------------------------------------------------
	default:
		panic("RefMap not yet implemented")
	}
	return
}

// MergeStringMap b into a at location and returns the new modified a, b takes higher precedence and will override a.
// Only merges map types by key recursively, does not attempt to merge lists.
func MergeStringMap(a, b map[string]interface{}, location ...string) map[string]interface{} {
	a2 := a

	// 1. Handle location if given
	key := ""
	if len(location) > 0 {
		key = location[0]
		var val interface{}
		val = a

		// Process keys from left to right
		keys, err := KeysFromSelector(key)
		if err == nil {
			for ko := keys.Shift(); !ko.Nil(); ko = keys.Shift() {
				key := ko.ToString()
				m := ToStringMap(val)

				// Set a new map as the value if not a map
				if v, ok := (*m)[key]; ok {
					if !ToStringMap(v).Any() {
						m.Set(key, map[string]interface{}{})
					}
				} else {
					m.Set(key, map[string]interface{}{})
				}
				val = (*m)[key]
			}
		}
		a2 = ToStringMap(val).G()
	}

	// 2. Merge at location
	switch {
	case a2 == nil && b == nil:
		return map[string]interface{}{}
	case a2 == nil:
		return b
	case b == nil:
		return a2
	}

	for k, v := range b {
		var av, bv interface{}

		// Ensure b value is Go type
		if val, ok := v.(*StringMap); ok {
			bv = val.G()
		} else {
			bv = v
		}

		// a doesn't have the key so just set b's value
		if val, exists := a2[k]; !exists {
			a2[k] = bv
		} else {
			if _val, ok := val.(*StringMap); ok {
				av = _val.G()
			} else {
				av = val
			}

			if bc, ok := bv.(map[string]interface{}); ok {
				if ac, ok := av.(map[string]interface{}); ok {
					// a and b both contain the key and are both submaps so recurse
					a2[k] = MergeStringMap(ac, bc)
				} else {
					// a is not a map so just override with b
					a2[k] = bv
				}
			} else {
				// b is not a map so just override a, no need to recurse
				a2[k] = bv
			}
		}
	}

	return a
}

// IdxFromSelector splits the given array index selector into individual components.
// The selector param is a jq like array selector [], []; size is the size of the target array.
// Getting a i==-1 and nil err indicates full array slice.
func IdxFromSelector(selector string, size int) (i int, k, v string, err error) {
	i = -1

	sel := A(selector)
	if sel.First().A() == "[" && sel.Last().A() == "]" {

		// Trim off the indexer/selector brackets and check the indexer
		idx := sel.TrimPrefix("[").TrimSuffix("]").A()
		if idx != "" {
			pieces := strings.Split(idx, "==")
			i, err = strconv.Atoi(idx)
			switch {

			// Select by key==value, e.g. .[k==v]
			case len(pieces) == 2:
				k, v = pieces[0], pieces[1]
				err = nil

			// Index in if the value is a valid integer, e.g. .[2], .[-1]
			case err == nil:
				if i = absIndex(size, i); i == -1 {
					err = errors.Errorf("invalid array index %v", idx)
				}
			default:
				err = errors.Errorf("invalid array index selector %v", idx)
			}
		}
	}

	// Set a consistent return value for errors
	if err != nil || k != "" || v != "" {
		i = -1
	}
	return
}

// KeysFromSelector splits the given key selectors into individual keys
func KeysFromSelector(selector string) (keys *StringSlice, err error) {
	keys = NewStringSliceV()

	var quotes *StringSlice
	if quotes, err = A(selector).SplitQuotes(); err != nil {
		return
	}
	for i := 0; i < quotes.Len(); i++ {
		quote := quotes.At(i).ToStr()

		// Split quotes into keys
		// 1. a single dot notation string that needs split
		// 2. a single quoted key to leave intact
		var qKeys *StringSlice
		if quote.First().A() != `"` {
			qKeys = A(quote).Split(".")
		} else {
			qKeys = ToStringSlice(quote.TrimPrefix(`"`).TrimSuffix(`"`))
		}

		// Process keys from left to right
		for k := qKeys.Shift(); !k.Nil(); k = qKeys.Shift() {
			key := k.ToStr()

			// Skip empty keys e.g. ".key" => ["", "key"]
			if k.A() == "" {
				continue
			}

			keys.Append(key)
		}
	}
	return
}
