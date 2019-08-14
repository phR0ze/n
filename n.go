// Package n provides a set of types with convenience methods for Go akin to rapid
// development languages.
//
// n was created to reduce the friction I had adopting Go as my primary language of choice.
// It does this by reducing the coding verbosity Go normally requires. n types wrap various
// Go types to provide generic convenience methods reminiscent of C#'s Queryable interface,
// removing the need to implement the 'Contains' function, on basic list primitives for the
// millionth time. The intent is at a minimum to have support for YAML primitive scalars
// (Null, String, Integer, Boolean, Float, Timestamp), lists of the scalar types and maps
// of the scalar types with reflection based fallbacks for un-optimized types. I'll be
// using the terms 'n types' or 'queryable' interchangeably in the documentation and
// examples.
//
// Conventions used across n types and pkgs
//
// • In order to deal with Golang's decision to not support function overloading or special
// characters in their function names n makes use of a variety of prefix/suffix capital
// letters to indicate different function variations. The function/method that contains no
// suffix is referred to as the base function/method.
//
// • Function names suffixed with 'A' indicates the function is a variation to the function
// without the 'A' but either accepts a string as input or returns a string.
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
// • Documentation should be thorough and relied upon for guidance as, for a love of
// brevity, some functions are named with a single capital letter only. 'G' is being used
// to export the underlying Go type. 'O' is being used to indicate the interface{} type or
// to export the underlying Go type as an interface{}. 'S' is used to refer to slice types,
// 'M' refers to map types, 'A' refers to string types, 'I' ints types, 'R' rune types and
// combinations may be used to indicate complex types. The documentation will always call
// out what exactly they mean, but the function name may be cryptic until understood.
//
// Summary of Types
//
// • Char
// • FloatSlice
// • IntSlice
// • InterSlice
// • Object
// • RefSlice
// • Str
// • StringSlice
// • StringMap
package n

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// Misc convenience type/functions
//--------------------------------------------------------------------------------------------------

// O is an alias for interface{} used in lambda expresssions for brevity.
type O interface{}

// Break is a brevity helper for breaking out of lambda loops
var Break = errors.New("break")

// EitherStr returns the first string if not empty else the second
func EitherStr(first, second string) string {
	if first != "" {
		return first
	}
	return second
}

// ExB avoids Go's gastly 4 line monstrosity required to implement this providing
// instead a single clean line of code for lambdas.
func ExB(exp bool) bool {
	if exp {
		return true
	}
	return false
}

// Range creates slice of the given range of numbers inclusive
func Range(min, max int) []int {
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = min + i
	}
	return result
}

// SetOnFalseB only updates the result to the 'value' if the exp is false
func SetOnFalseB(result *bool, value, exp bool) bool {
	if result != nil {
		if !exp {
			*result = value
		}
		return *result
	}
	return false
}

// SetOnTrueB only updates the result to the 'value' if the exp is true
func SetOnTrueB(result *bool, value, exp bool) bool {
	if result != nil {
		if exp {
			*result = value
		}
		return *result
	}
	return false
}

// Load and From helper functions
//--------------------------------------------------------------------------------------------------

// LoadYaml reads in a yaml file and converts it to a *StringMap
func LoadYaml(filepath string) (m *StringMap) {
	m, _ = LoadYamlE(filepath)
	return m
}

// LoadYamlE reads in a yaml file and converts it to a *StringMap
func LoadYamlE(filepath string) (m *StringMap, err error) {
	m = NewStringMapV()

	// Read in the yaml file
	var bytes []byte
	if bytes, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read in the yaml file %s", filepath)
		return
	}

	// Unmarshal the yaml into a *StringMap
	m, err = ToStringMapE(bytes)

	return
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
// 					err = errors.Errorf("not a key value pair type")
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
// 					err = errors.Errorf("not a key value pair type")
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
// 					err = errors.Errorf("not a key value pair type")
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
// 				err = errors.Errorf("%v is not of type map[string]interface{}", x)
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
// 				err = errors.Errorf("%v is not of type map[string]string", x)
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
