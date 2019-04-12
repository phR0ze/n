// Package n provides many Go types with convenience functions reminiscent of Ruby or C#.
//
// n was created to reduce the friction I had adopting Go as my primary language of choice by
// reducing coding verbosity required by Go via code reuse. The n types wrap various Go
// types to provide this functionality.
//
// Conventions used across n types and pkgs
//
// • In order to deal with Golang's decision to not support function overloading or special
// characters in their function names n makes use of a variety of prefix/suffix capital
// letters to indicate different function varieties. The function/method that contains no
// suffix is known as the base function/method.
//
// • Function names suffixed with 'E' indicates the function is a variation to the function
// without the 'E' but returns an Error while the base function does not.
//
// • Function names suffixed with 'M' indicates the function is a variation to the function
// without the 'M' but modifies the n type directly rather than a copy.
//
// • Function names suffixed with 'R' indicates the function is a variation to the function
// without the 'R' but reverses the order of operations.
//
// • Function names suffixed with 'S' indicates the function is a variation to the function
// without the 'S' but either accepts a Slice as input or returns a Slice.
//
// • Function names suffixed with 'V' indicates the function is a variation to the function
// • Function names suffixed with 'V' indicates the function is a variation to the function
// without the 'V' but accepts variadic input.
//
// • Function names suffixed with 'W' indicates the function is a variation to the function
// without the 'W' but accepts a lambda expression as input.
//
// • Documentation should be thorough and relied upon for guidance as, for a love of brevity,
// some functions use single capital letters only to indicate types. 'O' is being used to
// indicate the interface{} type or to export the underlying Go type as an interface{}. 'S' is
// used to refer to slice types, 'M' refers to map types, 'A' refers to string types, 'I' ints
// types and combinations may be used to indicate complex types. The documentation will always
// call out what exactly they mean, but the function name may be cryptic until understood.
package n

import (
	"errors"
	"reflect"
)

// Lambda convenience type/functions
//--------------------------------------------------------------------------------------------------

// O is an alias for interface{} used in lambda expresssions for brevity.
type O interface{}

// ErrBreak is a brevity helper for breaking out of lambda loops
var ErrBreak = errors.New("break")

// ExB avoids Go's gastly 4 line monstrosity required to implement this providing
// instead a single clean line of code for lambdas.
func ExB(exp bool) bool {
	if exp {
		return true
	}
	return false
}

// Reflection convenience type/functions
//--------------------------------------------------------------------------------------------------

// Indirect dereferences the reflect.Value recursively until its a non-pointer type
func Indirect(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}
	return Indirect(v.Elem())
}

// Misc convenience type/functions
//--------------------------------------------------------------------------------------------------

// Range creates slice of the given range of numbers inclusive
func Range(min, max int) []int {
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = min + i
	}
	return result
}

// SetValueOrDefault sets the value of the given string 'target' to the
// given 'value' if value is not empty else 'defaulty' and returns the value as well
func SetValueOrDefault(target *string, value, defaulty string) string {
	if value != "" {
		*target = value
	} else {
		*target = defaulty
	}

	return *target
}

// SetValueIfEmpty sets the value of the given string to the other if not empty
// and returns the target as well.
func SetValueIfEmpty(target *string, value string) string {
	if *target == "" {
		*target = value
	}
	return *target
}

// ValueOrDefault returns the value or defaulty if the value is empty
func ValueOrDefault(value, defaulty string) string {
	if value != "" {
		return value
	}
	return defaulty
}

// MergeMap b into a and returns the new modified a
// b takes higher precedence and will override a
func MergeMap(a, b map[string]interface{}) map[string]interface{} {
	switch {
	case (a == nil || len(a) == 0) && (b == nil || len(b) == 0):
		return map[string]interface{}{}
	case a == nil || len(a) == 0:
		return b
	case b == nil || len(b) == 0:
		return a
	}

	for k, bv := range b {
		if av, exists := a[k]; !exists {
			a[k] = bv
		} else if bc, ok := bv.(map[string]interface{}); ok {
			if ac, ok := av.(map[string]interface{}); ok {
				a[k] = MergeMap(ac, bc)
			} else {
				a[k] = bv
			}
		} else {
			a[k] = bv
		}
	}

	return a
}

// // M exports numerable into a map
// func (q *OldNumerable) M() (result map[string]interface{}, err error) {
// 	result = map[string]interface{}{}
// 	if q != nil && !q.Nil() {
// 		if v, ok := q.O().(map[string]interface{}); ok {
// 			result = v
// 		} else {
// 			next := q.Iter()
// 			for x, ok := next(); ok; x, ok = next() {
// 				if pair, ok := x.(KeyVal); ok {
// 					result[fmt.Sprint(pair.Key)] = pair.Val
// 				} else {
// 					err = fmt.Errorf("not a key value pair type")
// 					return
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// // O exports numerable into a interface{}
// func (q *OldNumerable) O() interface{} {
// 	if q == nil || q.Nil() {
// 		return nil
// 	}
// 	return q.v.Interface()
// }

// // Strs exports numerable into an string slice
// func (q *OldNumerable) Strs() (result []string) {
// 	result = []string{}
// 	if q != nil && !q.Nil() && q.TypeSlice() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			result = append(result, fmt.Sprint(x))
// 		}
// 	}
// 	return
// }

// // AAMap exports numerable into an string to string map
// func (q *OldNumerable) AAMap() (result map[string]string, err error) {
// 	result = map[string]string{}
// 	if q != nil && !q.Nil() {
// 		if v, ok := q.O().(map[string]string); ok {
// 			result = v
// 		} else {
// 			next := q.Iter()
// 			for x, ok := next(); ok; x, ok = next() {
// 				if pair, ok := x.(KeyVal); ok {
// 					result[fmt.Sprint(pair.Key)] = fmt.Sprint(pair.Val)
// 				} else {
// 					err = fmt.Errorf("not a key value pair type")
// 					return
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// // ASAMap exports numerable into an string to []string map
// func (q *OldNumerable) ASAMap() (result map[string][]string, err error) {
// 	result = map[string][]string{}
// 	if q != nil && !q.Nil() {
// 		if v, ok := q.O().(map[string][]string); ok {
// 			result = v
// 		} else {
// 			next := q.Iter()
// 			for x, ok := next(); ok; x, ok = next() {
// 				if pair, ok := x.(KeyVal); ok {
// 					key := fmt.Sprint(pair.Key)
// 					if slice, ok := pair.Val.([]string); ok {
// 						result[key] = slice
// 					} else {
// 						result[key] = []string{}
// 						nexty := Q(pair.Val).Iter()
// 						for y, ok := nexty(); ok; y, ok = nexty() {
// 							result[key] = append(result[key], fmt.Sprint(y))
// 						}
// 					}
// 				} else {
// 					err = fmt.Errorf("not a key value pair type")
// 					return
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// // S exports numerable into an interface{} slice
// func (q *OldNumerable) S() []interface{} {
// 	result := []interface{}{}
// 	if q != nil && !q.Nil() && q.TypeSlice() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			result = append(result, x)
// 		}
// 	}
// 	return result
// }

// // SAMap exports numerable into an slice of string to interface{} map
// func (q *OldNumerable) SAMap() (result []map[string]interface{}, err error) {
// 	result = []map[string]interface{}{}
// 	if q != nil && !q.Nil() && q.TypeSlice() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			m := map[string]interface{}{}
// 			switch x := x.(type) {
// 			case map[string]interface{}:
// 				m = x
// 			case map[interface{}]interface{}:
// 				for k, v := range x {
// 					m[fmt.Sprint(k)] = v
// 				}
// 			default:
// 				err = fmt.Errorf("%v is not of type map[string]interface{}", x)
// 				return
// 			}
// 			result = append(result, m)
// 		}
// 	}
// 	return
// }

// // SAAMap exports numerable into an slice of string to string map
// func (q *OldNumerable) SAAMap() (result []map[string]string, err error) {
// 	result = []map[string]string{}
// 	if q != nil && !q.Nil() && q.TypeSlice() {
// 		next := q.Iter()
// 		for x, ok := next(); ok; x, ok = next() {
// 			m := map[string]string{}
// 			switch x := x.(type) {
// 			case map[string]string:
// 				m = x
// 			case map[string]interface{}:
// 				for k, v := range x {
// 					m[fmt.Sprint(k)] = fmt.Sprint(v)
// 				}
// 			case map[interface{}]string:
// 				for k, v := range x {
// 					m[fmt.Sprint(k)] = fmt.Sprint(v)
// 				}
// 			case map[interface{}]interface{}:
// 				for k, v := range x {
// 					m[fmt.Sprint(k)] = fmt.Sprint(v)
// 				}
// 			default:
// 				err = fmt.Errorf("%v is not of type map[string]string", x)
// 				return
// 			}
// 			result = append(result, m)
// 		}
// 	}
// 	return
// }

// // When obj is a slice and this is a map each item will be checked in the map as a key.
// func (q *OldNumerable) Contains(obj interface{}) bool {
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
// func (q *OldNumerable) ContainsAny(obj interface{}) bool {
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
