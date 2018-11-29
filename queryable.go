package nub

import (
	"bytes"
	"errors"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
)

// Queryable provides chainable deferred execution
// and is the heart of the algorithm abstraction layer
type Queryable struct {
	v    *reflect.Value
	Kind reflect.Kind
	Iter func() Iterator
}

// Iterator provides a closure to capture the index and reset it
type Iterator func() (item interface{}, ok bool)

// KeyVal similar to C# for iterator over maps
type KeyVal struct {
	Key interface{}
	Val interface{}
}

func strIter(ref reflect.Value) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		return func() (item interface{}, ok bool) {
			if ok = i < len; ok {
				item = ref.Index(i).Interface()
				i++
			}
			return
		}
	}
}

func mapIter(ref reflect.Value) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		keys := ref.MapKeys()
		return func() (item interface{}, ok bool) {
			if ok = i < len; ok {
				item = KeyVal{
					Key: keys[i].Interface(),
					Val: ref.MapIndex(keys[i]).Interface(),
				}
				i++
			}
			return
		}
	}
}

func sliceIter(ref reflect.Value) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		return func() (item interface{}, ok bool) {
			if ok = i < len; ok {
				item = ref.Index(i).Interface()
				i++
			}
			return
		}
	}
}

// LoadYAML a yaml/json file as a str map
// returns nil on failure of any kind
func LoadYAML(filepath string) *Queryable {
	if yamlFile, err := ioutil.ReadFile(filepath); err == nil {
		data := map[string]interface{}{}
		yaml.Unmarshal(yamlFile, &data)
		return Q(data)
	}
	return nil
}

// N provides a new empty Queryable slice
func N() *Queryable {
	v := reflect.ValueOf([]interface{}{})
	return &Queryable{v: &v, Kind: v.Kind(), Iter: sliceIter(v)}
}

// Q provides origination for the Queryable abstraction layer
func Q(obj interface{}) *Queryable {
	if obj == nil {
		return N()
	}

	v := reflect.ValueOf(obj)
	q := &Queryable{v: &v, Kind: v.Kind()}
	switch q.Kind {

	// Slice types
	case reflect.Array, reflect.Slice:
		q.Iter = sliceIter(v)

	// Handle map types
	case reflect.Map:
		q.Iter = mapIter(v)

	// Handle string types
	case reflect.String:
		q.Iter = strIter(v)

	// Chan types
	case reflect.Chan:
		panic("TODO: handle reflect.Chan")
	}

	return q
}

// Any checks if the queryable has anything in it
func (q *Queryable) Any() bool {
	if q.Iter != nil {
		return q.v.Len() > 0
	}
	return q.v.Interface() != nil
}

// AnyWhere checka if any match the given lambda
func (q *Queryable) AnyWhere(lambda func(interface{}) bool) bool {
	if !q.TypeSingle() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if lambda(x) {
				return true
			}
		}
	} else if lambda(q.v.Interface()) {
		return true
	}
	return false
}

// Append items to the end of the collection and return the queryable
// converting to a collection if necessary
func (q *Queryable) Append(obj ...interface{}) *Queryable {
	if q.TypeSingle() {
		*q = *N().Append(q.v.Interface())
	}

	// Append to slice type
	for i := 0; i < len(obj); i++ {
		item := reflect.ValueOf(obj[i])
		*q.v = reflect.Append(*q.v, item)
	}
	q.Iter = sliceIter(*q.v)

	return q
}

// At returns the item at the given index location. Allows for negative notation
func (q *Queryable) At(i int) *Queryable {
	if q.TypeIter() {
		if i < 0 {
			i = q.v.Len() + i
		}
		if i >= 0 && i < q.v.Len() {
			if str, ok := q.v.Interface().(string); ok {
				return Q(string(str[i]))
			}
			return Q(q.v.Index(i).Interface())
		}
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the queryable collection
func (q *Queryable) Clear() *Queryable {
	*q = *N()
	return q
}

// Contains checks if all of the given obj are found.
// When obj is a string and this is a string check will fall back on strings.Contains.
// When obj is a string and this is a string slice, slice will be checked for obj.
// When obj is a non-interable and this is non-iterable a direct check is made.
// When obj is a non-interable and this is slice, slice will be checked for obj.
// When obj is a slice of string and this is a string each string check using strings.Contains.
// When obj is a slice and this is a slice each item will be checked in the slice.
// When obj is a slice and this is a map each item will be checked in the map as a key.
func (q *Queryable) Contains(obj interface{}) bool {
	other := Q(obj)
	if !q.Any() || !other.Any() {
		return false
	}

	// Non iterable type
	if q.TypeSingle() {

		// Both strings - pass through to stings.Contains
		if q.TypeStr() && other.TypeStr() {
			return strings.Contains(q.v.Interface().(string), obj.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Set([]interface{}{obj})
		}
		next := other.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if str, ok := x.(string); ok {
				if !strings.Contains(q.v.Interface().(string), str) {
					return false
				}
			} else {
				if q.v.Interface() != x {
					return false
				}
			}
		}
	} else {
		switch q.Kind {
		case reflect.Array, reflect.Slice:
			if !other.TypeSingle() {
				next := other.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if !q.Contains(x) {
						return false
					}
				}
			} else {
				next := q.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if x == obj {
						return true
					}
				}
				return false
			}
		case reflect.Map:
			keys := Q(q.v.MapKeys()).Map(func(x interface{}) interface{} {
				return x.(reflect.Value).Interface()
			})
			if !other.TypeSingle() {
				next := other.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if !keys.Contains(x) {
						return false
					}
				}
			} else {
				if !keys.Contains(obj) {
					return false
				}
			}
		default:
			panic("TODO: implement Contains")
		}
	}
	return true
}

// ContainsAny checks if any of the given obj is found.
// ContainsAny behaves much like Contains only it allows for matching any not all.
func (q *Queryable) ContainsAny(obj interface{}) bool {
	other := Q(obj)

	// Non iterable type
	if q.TypeSingle() {

		// Both strings - pass through to stings.Contains
		if q.TypeStr() && other.TypeStr() {
			return strings.Contains(q.v.Interface().(string), obj.(string))
		}

		// Other is non iterable, convert to iterable
		if other.TypeSingle() {
			other.Set([]interface{}{obj})
		}
		next := other.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if q.v.Interface() == x {
				return true
			}
		}
	} else {
		switch q.Kind {
		case reflect.Array, reflect.Slice:
			if !other.TypeSingle() {
				next := q.Iter()
				for x, ok := next(); ok; x, ok = next() {
					nexty := other.Iter()
					for y, oky := nexty(); oky; y, oky = nexty() {
						if x == y {
							return true
						}
					}
				}
			} else {
				next := q.Iter()
				for x, ok := next(); ok; x, ok = next() {
					if x == obj {
						return true
					}
				}
			}
		case reflect.Map:
			if !other.TypeSingle() {
				for _, key := range q.v.MapKeys() {
					next := other.Iter()
					for x, ok := next(); ok; x, ok = next() {
						if key.Interface() == x {
							return true
						}
					}
				}
			} else {
				for _, key := range q.v.MapKeys() {
					if key.Interface() == obj {
						return true
					}
				}
			}
		default:
			panic("TODO: implement Contains")
		}
	}
	return false
}

// Each iterates over the queryable and executes the given action
func (q *Queryable) Each(action func(interface{})) {
	if q.TypeIter() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			action(x)
		}
	}
}

// Join slice items as string with given delimeter
func (q *Queryable) Join(delim string) *Queryable {
	var joined bytes.Buffer
	if q.TypeStr() {
		joined.WriteString(q.v.Interface().(string))
	} else if q.TypeIter() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			switch x := x.(type) {
			case string:
				joined.WriteString(x)
				joined.WriteString(delim)
			case int:
				joined.WriteString(strconv.Itoa(x))
				joined.WriteString(delim)
			}
		}
	}
	return Q(strings.TrimSuffix(joined.String(), delim))
}

// Len of the collection type including string
func (q *Queryable) Len() int {
	if q.TypeIter() {
		return q.v.Len()
	}
	return 1
}

// Map manipulates the queryable data into a new form
func (q *Queryable) Map(sel func(interface{}) interface{}) *Queryable {
	result := N()
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		result.Append(sel(x))
	}
	return result
}

// Set provides a way to set underlying object Queryable is operating on
func (q *Queryable) Set(obj interface{}) *Queryable {
	other := Q(obj)
	q.Kind = other.Kind
	q.Iter = other.Iter
	q.v = other.v
	return q
}

// Split the string into a slice on delimiter
func (q *Queryable) Split(delim string) *strSliceNub {
	if q.TypeStr() {
		return A(q.v.Interface().(string)).Split(delim)
	}
	return S()
}

// TypeIter checks if the queryable is iterable
func (q *Queryable) TypeIter() bool {
	if q.Iter != nil {
		return true
	}
	return false
}

// TypeMap checks if the queryable is reflect.Map
func (q *Queryable) TypeMap() bool {
	return q.Kind == reflect.Map
}

// TypeSlice checks if the queryable is reflect.Array or reflect.Slice
func (q *Queryable) TypeSlice() bool {
	return q.Kind == reflect.Array || q.Kind == reflect.Slice
}

// TypeStr checks if the queryable is encapsulating a string
func (q *Queryable) TypeStr() bool {
	if _, ok := q.v.Interface().(string); ok {
		return true
	}
	return false
}

// TypeSingle checks if the queryable is ecapuslating a string or is not iterable
func (q *Queryable) TypeSingle() bool {
	if !q.TypeIter() || q.TypeStr() {
		return true
	}
	return false
}

// YAML gets data by key which can be dot delimited
func (q *Queryable) YAML(key string) (result *Queryable) {
	keys := A(key).Split(".")
	if key, ok := keys.TakeFirst(); ok {
		switch x := q.v.Interface().(type) {
		case map[string]interface{}:
			if !A(key).ContainsAny(":", "[", "]") {
				if v, ok := x[key]; ok {
					result = Q(v)
				}
			}
		case []interface{}:
			k, v := A(key).TrimPrefix("[").TrimSuffix("]").Split(":").YAMLPair()
			for i := range x {
				if m, ok := x[i].(map[string]interface{}); ok {
					if entry, ok := m[k]; ok {
						if v == entry {
							return Q(m)
						}
					}
				}
			}
		}
		if keys.Len() != 0 && result.Any() {
			result = result.YAML(keys.Join(".").A())
		}
	}
	if result == nil {
		result = N()
	}
	return
}
