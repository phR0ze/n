package nub

import (
	"fmt"
	"reflect"
)

// Collecting functions that return external Go types here

// A materializes queryable into a string
func (q *Queryable) A() string {
	return q.v.Interface().(string)
}

// I materializes queryable into an int
func (q *Queryable) I() int {
	return q.v.Interface().(int)
}

// Ints materializes queryable into an int slice
func (q *Queryable) Ints() []int {
	result := []int{}
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		result = append(result, x.(int))
	}
	return result
}

// M materializes queryable into a map
func (q *Queryable) M() (result map[string]interface{}) {
	if v, ok := q.O().(map[string]interface{}); ok {
		result = v
	} else {
		result = map[string]interface{}{}
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			pair := x.(KeyVal)
			result[pair.Key.(string)] = pair.Val
		}
	}
	return result
}

// O materializes queryable into a interface{}
func (q *Queryable) O() interface{} {
	return q.v.Interface()
}

// Strs materializes queryable into an string slice
func (q *Queryable) Strs() []string {
	result := []string{}
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		result = append(result, fmt.Sprint(x))
	}
	return result
}

// CastToTypeOf casts the obj to the type of the typof
func CastToTypeOf(typof interface{}, obj interface{}) *reflect.Value {
	panic("TODO: experimenting with reflection")
	typ := reflect.TypeOf(typof)
	switch typ.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map:
		targetType := typ.Elem()
		originType := reflect.TypeOf(obj)
		fmt.Println(targetType)
		fmt.Println(originType)
	default:
	}

	return nil
}
