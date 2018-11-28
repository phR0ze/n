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
	Iter func() Iterator

	ref *reflect.Value
}

// Iterator provides a closure to capture the index and reset it
type Iterator func() (item interface{}, ok bool)

// KeyVal similar to C# for iterator over maps
type KeyVal struct {
	Key interface{}
	Val interface{}
}

// Load YAML/JSON from file into queryable
func Load(filepath string) *Queryable {
	if yamlFile, err := ioutil.ReadFile(filepath); err == nil {
		data := map[string]interface{}{}
		yaml.Unmarshal(yamlFile, &data)
		return Q(data)
	}
	return M()
}

// A provides a new empty Queryable string
func A() *Queryable {
	ref := reflect.ValueOf(string(""))
	return &Queryable{ref: &ref, Iter: strIter(ref)}
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

// M provides a new empty Queryable map
func M() *Queryable {
	ref := reflect.ValueOf(map[interface{}]interface{}{})
	return &Queryable{ref: &ref, Iter: mapIter(ref)}
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

// S provides a new empty Queryable slice
func S() *Queryable {
	ref := reflect.ValueOf([]interface{}{})
	return &Queryable{ref: &ref, Iter: sliceIter(ref)}
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

// Q provides origination for the Queryable abstraction layer
func Q(obj interface{}) *Queryable {
	if obj == nil {
		return S()
	}

	ref := reflect.ValueOf(obj)
	result := &Queryable{ref: &ref}
	switch ref.Kind() {

	// Slice types
	case reflect.Array, reflect.Slice:
		result.Iter = sliceIter(ref)

	// Handle map types
	case reflect.Map:
		result.Iter = mapIter(ref)

	// Handle string types
	case reflect.String:
		result.Iter = strIter(ref)

	// Chan types
	case reflect.Chan:
		panic("TODO: handle reflect.Chan")
	}

	return result
}

// At returns the item at the given index location. Allows for negative notation
func (q *Queryable) At(i int) *Queryable {
	if q.TypeIter() {
		if i < 0 {
			i = q.ref.Len() + i
		}
		if i >= 0 && i < q.ref.Len() {
			if str, ok := q.ref.Interface().(string); ok {
				return Q(string(str[i]))
			}
			return Q(q.ref.Index(i).Interface())
		}
	}
	panic(errors.New("Index out of slice bounds"))
}

// Clear the queryable collection
func (q *Queryable) Clear() *Queryable {
	*q = *S()
	return q
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

// Get an item by key which can be dot delimited
func (q *Queryable) Get(key string) *Queryable {
	//keys := Q(key).Split(".")
	// 	if k, ok := keys.TakeFirst(); ok {
	// 		if entry, exists := m.raw[k]; exists {
	// 			if v, ok := entry.(map[string]interface{}); ok {
	// 				result.raw = v
	// 				if keys.Len() != 0 {
	// 					result = result.StrMap(keys.Join(".").M())
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return result
	// }
	return nil
}

// Join slice items as string with given delimeter
func (q *Queryable) Join(delim string) *Queryable {
	var joined bytes.Buffer
	if q.TypeStr() {
		joined.WriteString(q.ref.Interface().(string))
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
		return q.ref.Len()
	}
	return 1
}

// Set provides a way to set underlying object Queryable is operating on
func (q *Queryable) Set(obj interface{}) *Queryable {
	other := Q(obj)
	q.Iter = other.Iter
	q.ref = other.ref
	return q
}

// Split the string into a slice on delimiter
func (q *Queryable) Split(delim string) *Queryable {
	if q.TypeStr() {
		return Q(strings.Split(q.ref.Interface().(string), delim))
	}
	return A()
}

// TypeIter checks if the queryable is iterable
func (q *Queryable) TypeIter() bool {
	if q.Iter != nil {
		return true
	}
	return false
}

// TypeStr checks if the queryable is encapsulating a string
func (q *Queryable) TypeStr() bool {
	if _, ok := q.ref.Interface().(string); ok {
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

// // Where acts as a filter narrowing in on specific items
// func (q *Queryable) Where() *Queryable {
// 	return nil
// }
