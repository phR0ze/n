package n

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
	// EachE(action func(O) error) (Slice, error)        // EachE calls the given lambda once for each element in this Map, passing in that element
	// EachI(action func(int, O)) Slice                  // EachI calls the given lambda once for each element in this Map, passing in the index and element
	// EachIE(action func(int, O) error) (Slice, error)  // EachIE calls the given lambda once for each element in this Map, passing in the index and element
	// EachR(action func(O)) Slice                       // EachR calls the given lambda once for each element in this Map in reverse, passing in that element
	// EachRE(action func(O) error) (Slice, error)       // EachRE calls the given lambda once for each element in this Map in reverse, passing in that element
	// EachRI(action func(int, O)) Slice                 // EachRI calls the given lambda once for each element in this Map in reverse, passing in that element
	// EachRIE(action func(int, O) error) (Slice, error) // EachRIE calls the given lambda once for each element in this Map in reverse, passing in that element
	// Empty() bool                                      // Empty tests if this Map is empty.
	// First() (elem *Object)                            // First returns the first element in this Map as Object.
	// FirstN(n int) Slice                               // FirstN returns the first n elements in this Map as a Slice reference to the original.
	Generic() bool                     // Generic returns true if the underlying implementation uses reflection
	Get(key interface{}) (val *Object) // Get returns the value at the given key location. Returns empty *Object if not found.
	// Index(elem interface{}) (loc int)                 // Index returns the index of the first element in this Map where element == elem
	// Insert(i int, elem interface{}) Slice             // Insert modifies this Map to insert the given element before the element with the given index.
	// Join(separator ...string) (str *Object)           // Join converts each element into a string then joins them together using the given separator or comma by default.
	// Last() (elem *Object)                             // Last returns the last element in this Map as an Object.
	// LastN(n int) Slice                                // LastN returns the last n elements in this Map as a Slice reference to the original.
	Keys() Slice                         // Keys returns all the keys in this Map as a Slice of the key type.
	Len() int                            // Len returns the number of elements in this Map.
	Merge(m Map, location ...string) Map // Merge modifies this Map by overriding its values at location with the given map where they both exist and returns a reference to this Map.
	// Less(i, j int) bool                               // Less returns true if the element indexed by i is less than the element indexed by j.
	// Nil() bool                                        // Nil tests if this Map is nil.
	O() interface{} // O returns the underlying data structure as is.
	// Pair() (first, second *Object)                    // Pair simply returns the first and second Slice elements as Objects.
	// Pop() (elem *Object)                              // Pop modifies this Map to remove the last element and returns the removed element as an Object.
	// PopN(n int) (new Map)                           // PopN modifies this Map to remove the last n elements and returns the removed elements as a new Map.
	// Prepend(elem interface{}) Slice                   // Prepend modifies this Map to add the given element at the begining and returns a reference to this Map.
	Query(key string) (val *Object)             // Query returns the value at the given key location, using a jq type selectors. Returns empty *Object if not found.
	QueryE(key string) (val *Object, err error) // Query returns the value at the given key location, using a jq type selectors. Returns empty *Object if not found.
	Remove(key string) Map                      // Remove modifies this map to remove the value at the given key location, using a jq type selectors. Returns a reference to this Map
	RemoveE(key string) (m Map, err error)      // RemoveE modifies this map to remove the value at the given key location, using a jq type selectors. Returns a reference to this Map
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
	ToStringMapG() (m map[string]interface{}) // ToStringMapG converts the map to a google type map[string]interface{}
	ToYaml() (m *StringMap)                   // ToYaml is an alias to ToStringMap
	ToYamlG() (m map[string]interface{})      // ToYamlG is an alias to ToStringMapG
	// Take(indices ...int) (new Map)                  // Take modifies this Map removing the indicated range of elements from this Map and returning them as a new Map.
	// TakeAt(i int) (elem *Object)                      // TakeAt modifies this Map removing the elemement at the given index location and returns the removed element as an Object.
	// TakeW(sel func(O) bool) (new Map)               // TakeW modifies this Map removing the elements that match the lambda selector and returns them as a new Map.
	// Union(slice interface{}) (new Map)              // Union returns a new Map by joining uniq elements from this Map with uniq elements from the given Slice while preserving order.
	// UnionM(slice interface{}) Slice                   // UnionM modifies this Map by joining uniq elements from this Map with uniq elements from the given Slice while preserving order.
	// Uniq() (new Map)                                // Uniq returns a new Map with all non uniq elements removed while preserving element order.
	// UniqM() Slice                                     // UniqM modifies this Map to remove all non uniq elements while preserving element order.
	WriteYaml(filename string) (err error) // WriteYaml converts the *StringMap into a map[string]interface{} then calls sys.WriteYaml on it to write it out to disk.
}

// M is an alias for NewMap
func M(m interface{}) (new Map) {
	return NewMap(m)
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

// KeysFromSelector splits the given key selectors into individual keys
func KeysFromSelector(key string) (keys *StringSlice, err error) {
	keys = NewStringSliceV()

	var quotes *StringSlice
	if quotes, err = A(key).SplitQuotes(); err != nil {
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

// // Any checks if the numerable has anything in it
// func (q *Numerable) Any() bool {
// 	if q.v == nil {
// 		return false
// 	}
// 	if q.Iter != nil {
// 		return q.v.Len() > 0
// 	}
// 	return q.v.Interface() != nil
// }

// // AnyWhere check if any match the given lambda
// func (q *Numerable) AnyWhere(lambda func(O) bool) bool {
// 	if !q.TypeSingle() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			if lambda(x) {
// 				return true
// 			}
// 		}
// 	} else if q.Nil() {
// 		return false
// 	} else if lambda(q.v.Interface()) {
// 		return true
// 	}
// 	return false
// }

// // Append modifies the underlying type, converting it to a slice as needed,
// // then appending the given items to the underlying collection.
// // Returns the numerable for chaining.
// func (q *Numerable) Append(items ...interface{}) *Numerable {
// 	if q.TypeMap() {
// 		panic("Append doesn't support map types")
// 	}
// 	q.toSlice(items...)
// 	for i := 0; i < len(items); i++ {
// 		item := reflect.ValueOf(items[i])
// 		*q.v = reflect.Append(*q.v, item)
// 	}
// 	q.Iter = sliceIter(*q.v)
// 	return q
// }

// // At returns the item at the given index location. Allows for negative notation
// func (q *Numerable) At(i int) *Numerable {
// 	if q.TypeIter() {
// 		if i < 0 {
// 			i = q.v.Len() + i
// 		}
// 		if i >= 0 && i < q.v.Len() {
// 			if str, ok := q.v.Interface().(string); ok {
// 				return Q(string(str[i]))
// 			}
// 			return Q(q.v.Index(i).Interface())
// 		}
// 	}
// 	return Nil()
// }

// // Clear the underlying collection in the numerable
// func (q *Numerable) Clear() *Numerable {
// 	switch q.Kind {
// 	case reflect.Array, reflect.Slice:
// 		*q.v = reflect.MakeSlice(q.v.Type(), 0, 10)
// 		q.Iter = sliceIter(*q.v)
// 	case reflect.Map:
// 		*q.v = reflect.MakeMap(q.v.Type())
// 		q.Iter = mapIter(*q.v)
// 	case reflect.String:
// 		*q.v = reflect.ValueOf("")
// 		q.Iter = strIter(*q.v)
// 	default:
// 		panic("unhandled type")
// 	}
// 	return q
// }

// // Contains checks if all of the given obj are found.
// // When obj is a string and this is a string check will fall back on strings.Contains.
// // When obj is a string and this is a string slice, slice will be checked for obj.
// // When obj is a non-interable and this is non-iterable a direct check is made.
// // When obj is a non-interable and this is slice, slice will be checked for obj.
// // When obj is a slice of string and this is a string each string check using strings.Contains.
// // When obj is a slice and this is a slice each item will be checked in the slice.
// // When obj is a slice and this is a map each item will be checked in the map as a key.
// func (q *Numerable) Contains(obj interface{}) bool {
// 	other := Q(obj)
// 	if !q.Any() || !other.Any() {
// 		return false
// 	}

// 	// Non iterable type
// 	if q.TypeSingle() {

// 		// Both strings - pass through to stings.Contains
// 		if q.TypeStr() && other.TypeStr() {
// 			return strings.Contains(q.v.Interface().(string), obj.(string))
// 		}

// 		// Other is non iterable, convert to iterable
// 		if other.TypeSingle() {
// 			other.Copy([]interface{}{obj})
// 		}
// 		next := other.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			if str, ok := x.(string); ok {
// 				if !strings.Contains(q.v.Interface().(string), str) {
// 					return false
// 				}
// 			} else {
// 				if q.v.Interface() != x {
// 					return false
// 				}
// 			}
// 		}
// 	} else {
// 		switch q.Kind {
// 		case reflect.Array, reflect.Slice:
// 			if !other.TypeSingle() {
// 				next := other.Iter()
// 				for x, ok := next(); ok; x, ok = next() {
// 					if !q.Contains(x) {
// 						return false
// 					}
// 				}
// 			} else {
// 				next := q.Iter()
// 				for x, ok := next(); ok; x, ok = next() {
// 					if x == obj {
// 						return true
// 					}
// 				}
// 				return false
// 			}
// 		case reflect.Map:
// 			keys := Q(q.v.MapKeys()).Map(func(x O) O {
// 				return x.(reflect.Value).Interface()
// 			})
// 			if !other.TypeSingle() {
// 				next := other.Iter()
// 				for x, ok := next(); ok; x, ok = next() {
// 					if !keys.Contains(x) {
// 						return false
// 					}
// 				}
// 			} else {
// 				if !keys.Contains(obj) {
// 					return false
// 				}
// 			}
// 		default:
// 			panic("TODO: implement Contains")
// 		}
// 	}
// 	return true
// }

// // ContainsAny checks if any of the given obj is found.
// // ContainsAny behaves much like Contains only it allows for matching any not all.
// func (q *Numerable) ContainsAny(obj interface{}) bool {
// 	other := Q(obj)
// 	if q.Nil() {
// 		return false
// 	}

// 	// Non iterable type
// 	if q.TypeSingle() {

// 		// Both strings - pass through to stings.Contains
// 		if q.TypeStr() && other.TypeStr() {
// 			return strings.Contains(q.v.Interface().(string), obj.(string))
// 		}

// 		// Other is non iterable, convert to iterable
// 		if other.TypeSingle() {
// 			other.Copy([]interface{}{obj})
// 		}
// 		next := other.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			if q.v.Interface() == x {
// 				return true
// 			}
// 		}
// 	} else {
// 		switch q.Kind {
// 		case reflect.Array, reflect.Slice:
// 			if !other.TypeSingle() {
// 				next := q.Iter()
// 				for x, ok := next(); ok; x, ok = next() {
// 					nexty := other.Iter()
// 					for y, oky := nexty(); oky; y, oky = nexty() {
// 						if x == y {
// 							return true
// 						}
// 					}
// 				}
// 			} else {
// 				next := q.Iter()
// 				for x, ok := next(); ok; x, ok = next() {
// 					if x == obj {
// 						return true
// 					}
// 				}
// 			}
// 		case reflect.Map:
// 			if !other.TypeSingle() {
// 				for _, key := range q.v.MapKeys() {
// 					next := other.Iter()
// 					for x, ok := next(); ok; x, ok = next() {
// 						if key.Interface() == x {
// 							return true
// 						}
// 					}
// 				}
// 			} else {
// 				for _, key := range q.v.MapKeys() {
// 					if key.Interface() == obj {
// 						return true
// 					}
// 				}
// 			}
// 		default:
// 			panic("TODO: implement Contains")
// 		}
// 	}
// 	return false
// }

// // Copy given obj into this one and reset types
// func (q *Numerable) Copy(obj interface{}) *Numerable {
// 	var other *Numerable
// 	if x, ok := obj.(*Numerable); ok {
// 		other = x
// 	} else {
// 		other = Q(obj)
// 	}
// 	q.Kind = other.Kind
// 	q.Iter = other.Iter
// 	q.v = other.v
// 	return q
// }

// // Delete all items that match the given item for slices or the key value
// // pair for maps or matching rune for strings. Returns true if something was deleted.
// func (q *Numerable) Delete(obj interface{}) (ok bool) {
// 	switch q.Kind {
// 	case reflect.Array, reflect.Slice:
// 		//*q.v = reflect.MakeSlice(q.v.Type(), 0, 10)
// 		//q.Iter = sliceIter(*q.v)
// 	case reflect.Map:
// 		key := reflect.ValueOf(obj)
// 		if val := q.v.MapIndex(key); val != (reflect.Value{}) {
// 			ok = true
// 			q.v.SetMapIndex(reflect.ValueOf(obj), reflect.Value{})
// 		}
// 	case reflect.String:
// 		//*q.v = reflect.ValueOf("")
// 		//q.Iter = strIter(*q.v)
// 	default:
// 		panic("unhandled type")
// 	}
// 	return
// }

// // DeleteAt deletes the item at the given index location. Allows for negative notation.
// // Returns the deleted element Numerable or Nil Numerable if missing.
// func (q *Numerable) DeleteAt(i int) (item *Numerable) {
// 	if q.TypeIter() && !q.TypeMap() {
// 		if i < 0 {
// 			i = q.v.Len() + i
// 		}
// 		if i >= 0 && i < q.v.Len() {
// 			switch x := q.v.Interface().(type) {

// 			// for strings delete at the rune level
// 			case string:
// 				item = Q(string(x[i]))
// 				if i+1 < len(x) {
// 					*q.v = reflect.ValueOf(string(append([]rune(x[:i]), []rune(x[i+1:])...)))
// 				} else {
// 					*q.v = reflect.ValueOf(x[:i])
// 				}

// 			// delete object from iterable
// 			default:
// 				item = Q(q.v.Index(i).Interface())
// 				if i+1 < q.v.Len() {
// 					*q.v = reflect.AppendSlice(q.v.Slice(0, i), q.v.Slice(i+1, q.v.Len()))
// 				} else {
// 					*q.v = q.v.Slice(0, i)
// 				}
// 			}

// 			q.Iter = sliceIter(*q.v)
// 			return item
// 		}
// 	}
// 	if item == nil {
// 		item = Nil()
// 	}
// 	return
// }

// // Each iterates over the numerable and executes the given action
// func (q *Numerable) Each(action func(O)) {
// 	if q.TypeIter() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			action(x)
// 		}
// 	}
// }

// // EachE iterates over the numerable and executes the given action
// // Abort early and return error if non nil
// func (q *Numerable) EachE(action func(O) error) (err error) {
// 	if q.TypeIter() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			if err = action(x); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return
// }

// // Find returns a new numerable containing the first item which matches the given lambda.
// // Returns nil if not found.
// func (q *Numerable) Find(lambda func(O) bool) (result *Numerable) {
// 	next := q.Iter()
// 	for x, ok := next(); ok; x, ok = next() {
// 		if lambda(x) {
// 			result = Q(x)
// 			return
// 		}
// 	}
// 	return
// }

// // First returns the first item as numerable
// // returns a nil numerable when index out of bounds
// func (q *Numerable) First() (result *Numerable) {
// 	if q.Len() > 0 {
// 		return q.At(0)
// 	}
// 	return Nil()
// }

// // Flatten returns a new slice that is one-dimensional flattening.
// // That is, for every item that is a slice, extract its items into the new slice.
// func (q *Numerable) Flatten() (result *Numerable) {
// 	if q.TypeSlice() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {

// 			// Create new slice of the inner type
// 			if result == nil {
// 				result = Q(reflect.MakeSlice(reflect.TypeOf(x), 0, 10).Interface())
// 			}

// 			// Add sub slice's elements to new slice
// 			Q(x).Each(func(y O) {
// 				result.Append(y)
// 			})
// 		}
// 	} else {
// 		panic("TODO: implement Flatten() for maps")
// 	}
// 	if result == nil {
// 		result = q
// 	}
// 	return
// }

// // Insert the item at the given index, negative notation supported
// func (q *Numerable) Insert(i int, items ...interface{}) *Numerable {
// 	q.toSlice(items...)
// 	if i < 0 {
// 		i = q.v.Len() + i
// 	}
// 	if i >= 0 && i < q.v.Len() && q.v.Len() > 0 && len(items) > 0 {

// 		// Create a new slice
// 		typ := q.v.Index(0).Type()
// 		slice := reflect.MakeSlice(reflect.SliceOf(typ), 0, 10)

// 		// Append those before i
// 		for _, j := range Range(0, i-1) {
// 			slice = reflect.Append(slice, q.v.Index(j))
// 		}

// 		// Append new items
// 		for j := 0; j < len(items); j++ {
// 			slice = reflect.Append(slice, reflect.ValueOf(items[j]))
// 		}

// 		// Append those after
// 		for _, j := range Range(i, q.Len()-1) {
// 			slice = reflect.Append(slice, q.v.Index(j))
// 		}

// 		*q = *Q(slice.Interface())
// 		q.Iter = sliceIter(*q.v)
// 	} else {
// 		q.Append(items...)
// 	}
// 	return q
// }

// // Join slice items as string with given delimeter
// func (q *Numerable) Join(delim string) *Numerable {
// 	var joined bytes.Buffer
// 	if q.TypeStr() {
// 		joined.WriteString(q.v.Interface().(string))
// 	} else if q.TypeIter() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			switch x := x.(type) {
// 			case string:
// 				joined.WriteString(x)
// 				joined.WriteString(delim)
// 			case int:
// 				joined.WriteString(strconv.Itoa(x))
// 				joined.WriteString(delim)
// 			}
// 		}
// 	}
// 	return Q(strings.TrimSuffix(joined.String(), delim))
// }

// // Last returns the last item as numerable
// // returns a nil numerable when index out of bounds
// func (q *Numerable) Last() (result *Numerable) {
// 	if q.Len() > 0 {
// 		return q.At(-1)
// 	}
// 	return Nil()
// }

// // Len of the collection type including string
// func (q *Numerable) Len() int {
// 	if q.TypeIter() {
// 		return q.v.Len()
// 	} else if q.Nil() {
// 		return 0
// 	}
// 	return 1
// }

// // Map manipulates the numerable data into a new form
// func (q *Numerable) Map(sel func(O) O) (result *Numerable) {
// 	next := q.Iter()
// 	for x, ok := next(); ok; x, ok = next() {
// 		obj := sel(x)

// 		// Drill into numerables
// 		if s, ok := obj.(*Numerable); ok {
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

// // MapF manipulates the numerable data into a new form then flattens
// func (q *Numerable) MapF(sel func(O) O) (result *Numerable) {
// 	result = q.Map(sel).Flatten()
// 	return
// }

// // MapMany manipulates the numerable data from two sources in a cross join
// func (q *Numerable) MapMany(sel func(O) O) (result *Numerable) {
// 	// next := q.Iter()
// 	// for x, ok := next(); ok; x, ok = next() {
// 	// 	s := sel(x)

// 	// 	// Create new slice of the return type of sel
// 	// 	if result == nil {
// 	// 		typ := reflect.TypeOf(s)
// 	// 		result = Q(reflect.MakeSlice(reflect.SliceOf(typ), 0, 10).Interface())
// 	// 	}
// 	// 	result.Append(s)
// 	// }
// 	// return result
// 	return
// }

// // Nil tests if the numerable is a nil numerable
// func (q *Numerable) Nil() bool {
// 	if q.v == nil || q.Kind == reflect.Invalid {
// 		return true
// 	}
// 	return false
// }

// // Select returns a new numerable containing all items which match the given lambda
// func (q *Numerable) Select(lambda func(O) bool) (result *Numerable) {
// 	result = q.newSlice()
// 	next := q.Iter()
// 	for x, ok := next(); ok; x, ok = next() {
// 		if lambda(x) {
// 			result.Append(x)
// 		}
// 	}
// 	return result
// }

// // Set the item at the given index to the given item
// func (q *Numerable) Set(i int, item interface{}) *Numerable {
// 	if q.TypeIter() && !q.TypeStr() {
// 		if i < 0 {
// 			i = q.v.Len() + i
// 		}
// 		if i >= 0 && i < q.v.Len() {
// 			v := reflect.ValueOf(item)
// 			q.v.Index(i).Set(v)
// 		}
// 	}
// 	return q
// }

// // Split the string into a slice on delimiter
// func (q *Numerable) Split(delim string) *strSliceN {
// 	if q.TypeStr() {
// 		return A(q.v.Interface().(string)).Split(delim)
// 	}
// 	return S()
// }

// // TypeIter checks if the numerable is iterable
// func (q *Numerable) TypeIter() bool {
// 	if q.Iter != nil {
// 		return true
// 	}
// 	return false
// }

// // TypeMap checks if the numerable is reflect.Map
// func (q *Numerable) TypeMap() bool {
// 	return q.Kind == reflect.Map
// }

// // TypeSlice checks if the numerable is reflect.Array or reflect.Slice
// func (q *Numerable) TypeSlice() bool {
// 	return q.Kind == reflect.Array || q.Kind == reflect.Slice
// }

// // TypeStr checks if the numerable is encapsulating a string
// func (q *Numerable) TypeStr() bool {
// 	return q.Kind == reflect.String
// }

// // TypeSingle checks if the numerable is ecapuslating a string or is not iterable
// func (q *Numerable) TypeSingle() bool {
// 	if !q.TypeIter() || q.TypeStr() || q.Nil() {
// 		return true
// 	}
// 	return false
// }

// // Convert the single type into a slice type
// func (q *Numerable) toSlice(items ...interface{}) {
// 	if q.TypeSingle() {
// 		nq := q.newSlice(items...)
// 		if !q.Nil() {
// 			*nq.v = reflect.Append(*nq.v, *q.v)
// 		}
// 		*q = *nq
// 	}
// }

// // Create a new slice of the inner type
// func (q *Numerable) newSlice(items ...interface{}) *Numerable {
// 	var typ reflect.Type
// 	switch {
// 	case len(items) > 0:
// 		typ = reflect.SliceOf(reflect.TypeOf(items[0]))
// 	case q.Nil():
// 		typ = reflect.TypeOf([]interface{}{})
// 	case q.TypeSingle():
// 		typ = reflect.SliceOf(q.v.Type())
// 	case q.TypeMap():
// 		typ = reflect.SliceOf(reflect.TypeOf(KeyVal{}))
// 	default:
// 		if q.Any() {
// 			typ = reflect.SliceOf(q.v.Index(0).Type())
// 		} else {
// 			typ = q.v.Type()
// 		}
// 	}
// 	return Q(reflect.MakeSlice(typ, 0, 10).Interface())
// }
