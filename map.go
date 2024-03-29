package n

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// IMap provides a generic way to work with map types providing convenience methods
// on par with rapid development languages. 'this IMap' refers to the current map
// instance being operated on.  'new IMap' refers to a copy of the map.
type IMap interface {
	Any(keys ...interface{}) bool // Any tests if this Map is not empty or optionally if it contains any of the given variadic keys.
	// AnyS(slice interface{}) bool                      // AnyS tests if this Map contains any of the given Slice's elements.
	// AnyW(sel func(O) bool) bool                       // AnyW tests if this Map contains any that match the lambda selector.
	// Append(elem interface{}) Slice                    // Append an element to the end of this Map and returns a reference to this Map.
	// AppendV(elems ...interface{}) Slice               // AppendV appends the variadic elements to the end of this Map and returns a reference to this Map.
	Clear() IMap // Clear modifies this Map to clear out all key-value pairs and returns a reference to this Map.
	// Concat(slice interface{}) (new Slice)             // Concat returns a new Slice by appending the given Slice to this Map using variadic expansion.
	// ConcatM(slice interface{}) Slice                  // ConcatM modifies this Map by appending the given Slice using variadic expansion and returns a reference to this Map.
	Copy(keys ...interface{}) (new IMap) // Copy returns a new Map with the indicated key-value pairs copied from this Map or all if not given.
	// Count(elem interface{}) (cnt int)                 // Count the number of elements in this Map equal to the given element.
	// CountW(sel func(O) bool) (cnt int)                // CountW counts the number of elements in this Map that match the lambda selector.
	Delete(key interface{}) (val *Object) // Delete modifies this Map to delete the indicated key-value pair and returns the value from the Map.
	DeleteM(key interface{}) IMap         // DeleteM modifies this Map to delete the indicated key-value pair and returns a reference to this Map rather than the key-value pair.
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
	Generic() bool                                                // Generic returns true if the underlying implementation uses reflection
	Get(key interface{}) (val *Object)                            // Get returns the value at the given key location. Returns empty *Object if not found.
	Update(selector string, val interface{}) IMap                 // Update sets the value for the given key location, using jq type selectors. Returns a reference to this Map.
	UpdateE(selector string, val interface{}) (m IMap, err error) // UpdateE sets the value for the given key location, using jq type selectors. Returns a reference to this Map.
	// Join(separator ...string) (str *Object)           // Join converts each element into a string then joins them together using the given separator or comma by default.
	Keys() ISlice                          // Keys returns all the keys in this Map as a Slice of the key type.
	Len() int                              // Len returns the number of elements in this Map.
	M() (m *StringMap)                     // M is an alias to ToStringMap
	MG() (m map[string]interface{})        // MG is an alias to ToStringMapG
	Merge(m IMap, location ...string) IMap // Merge modifies this Map by overriding its values at location with the given map where they both exist and returns a reference to this Map.
	// Less(i, j int) bool                               // Less returns true if the element indexed by i is less than the element indexed by j.
	// Nil() bool                                        // Nil tests if this Map is nil.
	O() interface{} // O returns the underlying data structure as is.
	// Pair() (first, second *Object)                    // Pair simply returns the first and second Slice elements as Objects.
	// Pop() (elem *Object)                              // Pop modifies this Map to remove the last element and returns the removed element as an Object.
	// PopN(n int) (new Map)                           // PopN modifies this Map to remove the last n elements and returns the removed elements as a new Map.
	// Prepend(elem interface{}) Slice                   // Prepend modifies this Map to add the given element at the begining and returns a reference to this Map.
	Query(selector string, params ...interface{}) (val *Object)             // Query returns the value at the given selector location, using jq type selectors. Returns empty *Object if not found.
	QueryE(selector string, params ...interface{}) (val *Object, err error) // Query returns the value at the given selector location, using jq type selectors. Returns empty *Object if not found.
	Remove(selector string, params ...interface{}) IMap                     // Remove modifies this map to remove the value at the given selector location, using jq type selectors. Returns a reference to this Map
	RemoveE(selector string, params ...interface{}) (m IMap, err error)     // RemoveE modifies this map to remove the value at the given selector location, using jq type selectors. Returns a reference to this Map
	// Reverse() (new Map)                             // Reverse returns a new Map with the order of the elements reversed.
	// ReverseM() Slice                                  // ReverseM modifies this Map reversing the order of the elements and returns a reference to this Map.
	// Select(sel func(O) bool) (new Map)              // Select creates a new Map with the elements that match the lambda selector.
	Set(selector, val interface{}) bool  // Set the value for the given key to the given val. Returns true if the selector did not yet exists in this Map.
	SetM(selector, val interface{}) IMap // SetM the value for the given selector to the given val creating map if necessary. Returns a reference to this Map.
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

// Map provides a generic way to work with Map types. It does this by wrapping Go types
// directly for optimized types thus avoiding reflection processing overhead and making a plethora
// of Map methods available. Non-optimized types will fall back on reflection to generically
// handle the type incurring the full 10x reflection processing overhead.
//
// Optimized: map[string]interface{}
func Map(obj interface{}) (new *StringMap) {
	o := Reference(obj)
	switch x := o.(type) {

	// StringMap
	// ---------------------------------------------------------------------------------------------
	case []byte, *[]byte, string, *string, *StringMap, *map[string]interface{}, *map[string]string,
		*map[string]bool, *map[string]float32, *map[string]float64,
		*map[string]int, *map[string]int8, *map[string]int16, *map[string]int32, *map[string]int64,
		*map[string]uint, *map[string]uint8, *map[string]uint16, *map[string]uint32, *map[string]uint64:
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
func MergeStringMap(a, b map[string]interface{}, selector ...string) map[string]interface{} {
	return ToStringMap(a).MergeG(ToStringMap(b), selector...)
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
//   - `selector` supports dot notation similar to https://stedolan.github.io/jq/manual/#Basicfilters with some caveats
//   - `params` are the string interpolation paramaters similar to fmt.Sprintf()
func KeysFromSelector(selector string, params ...interface{}) (keys *StringSlice, err error) {
	keys = NewStringSliceV()

	var quotes *StringSlice
	if quotes, err = A(fmt.Sprintf(selector, params...)).SplitQuotes(); err != nil {
		return
	}
	for i := 0; i < quotes.Len(); i++ {
		quote := quotes.At(i).ToStr()

		// Split quotes into keys
		// 1. a single dot notation string that needs split
		// 2. a single quoted key to leave intact
		var qKeys *StringSlice
		if quote.First().A() != `"` {
			qKeys = A(quote).SplitEscape(".", "\\")
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
