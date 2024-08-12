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
// # Conventions used across n types and pkgs
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
// without the 'S' but either accepts a ISlice as input or returns a ISlice.
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
// # Summary of Types
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
	"encoding/json"
	"os"

	yaml "github.com/phR0ze/yaml/v2"
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

// EitherInt returns the first int if not zero else the second
func EitherInt(first, second int) int {
	if first != 0 {
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

// SetOnEmpty updates the given result string to the given value if it is empty
func SetOnEmpty(result *string, value string) string {
	if result != nil {
		if *result == "" {
			*result = value
		}
		return *result
	}
	return value
}

// SetOnTrueA updates the given result string to the given value if the exp is true
func SetOnTrueA(result *string, value string, exp bool) string {
	if result != nil {
		if *result == "" {
			*result = value
		}
		return *result
	}
	return value
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

// LoadJSON reads in a json file and converts it to a *StringMap
func LoadJSON(filepath string) (m *StringMap) {
	m, _ = LoadJSONE(filepath)
	return m
}

// LoadJSONE reads in a json file and converts it to a *StringMap
func LoadJSONE(filepath string) (m *StringMap, err error) {

	// Read in the yaml file
	var data []byte
	if data, err = os.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read in the yaml file %s", filepath)
		return
	}

	// Unmarshal the json into a *StringMap
	buff := map[string]interface{}{}
	if err = json.Unmarshal(data, &buff); err != nil {
		err = errors.Wrapf(err, "failed to unmarshal json file %s into a *StringMap", filepath)
		return
	}
	m = MV(buff)

	return
}

// LoadYAML reads in a yaml file and converts it to a *StringMap
func LoadYAML(filepath string) (m *StringMap) {
	m, _ = LoadYAMLE(filepath)
	return m
}

// LoadYAMLE reads in a yaml file and converts it to a *StringMap
func LoadYAMLE(filepath string) (m *StringMap, err error) {
	m = NewStringMapV()

	// Read in the yaml file
	var data []byte
	if data, err = os.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read in the yaml file %s", filepath)
		return
	}

	// Unmarshal the yaml into a *StringMap
	m, err = ToStringMapE(data)

	return
}

// YAMLCont checks if the given value is a valid YAML container
func YAMLCont(obj interface{}) bool {
	o := DeReference(obj)
	switch o.(type) {
	case map[string]interface{}, StringMap, yaml.MapSlice:
		return true
	case []interface{}, []string, []int:
		return true
	default:
		return false
	}
}

// YAMLMap checks if the given value is map compatible
func YAMLMap(obj interface{}) bool {
	o := DeReference(obj)
	switch o.(type) {
	case map[string]interface{}, StringMap, yaml.MapSlice:
		return true
	default:
		return false
	}
}

// YAMLArray checks if the given value is array compatible
func YAMLArray(obj interface{}) bool {
	o := DeReference(obj)
	switch o.(type) {
	case []interface{}, []string, []int:
		return true
	default:
		return false
	}
}
