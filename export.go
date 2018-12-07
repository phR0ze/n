package n

import (
	"fmt"
	"reflect"
	"strconv"
)

// Collecting functions that return external Go types here

// A exports queryable into a string
func (q *Queryable) A() (result string) {
	if !q.Nil() {
		switch v := q.v.Interface().(type) {
		case string:
			result = v
		default:
			result = fmt.Sprintf("%v", v)
		}
	}
	return
}

// B exports queryable into a bool
func (q *Queryable) B() (result bool) {
	if !q.Nil() {
		switch v := q.v.Interface().(type) {
		case bool:
			result = v
		case int:
			result = v != 0
		case string:
			result, _ = strconv.ParseBool(v)
		}
	}
	return
}

// I exports queryable into an int
func (q *Queryable) I() (result int) {
	if !q.Nil() {
		switch v := q.v.Interface().(type) {
		case int:
			result = v
		case bool:
			if v {
				result = 1
			}
		case string:
			result, _ = strconv.Atoi(v)
		}
	}
	return

}

// Ints exports queryable into an int slice
func (q *Queryable) Ints() (result []int) {
	result = []int{}
	if q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			result = append(result, x.(int))
		}
	}
	return result
}

// M exports queryable into a map
func (q *Queryable) M() (result map[string]interface{}) {
	result = map[string]interface{}{}
	if !q.Nil() {
		if v, ok := q.O().(map[string]interface{}); ok {
			result = v
		} else {
			next := q.Iter()
			for x, ok := next(); ok; x, ok = next() {
				if pair, ok := x.(KeyVal); ok {
					result[fmt.Sprint(pair.Key)] = pair.Val
				}
			}
		}
	}
	return result
}

// O exports queryable into a interface{}
func (q *Queryable) O() interface{} {
	if q.Nil() {
		return nil
	}
	return q.v.Interface()
}

// Strs exports queryable into an string slice
func (q *Queryable) Strs() []string {
	result := []string{}
	if q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			result = append(result, fmt.Sprint(x))
		}
	}
	return result
}

// AAMap exports queryable into an string to string map
func (q *Queryable) AAMap() (result map[string]string) {
	result = map[string]string{}
	if !q.Nil() {
		if v, ok := q.O().(map[string]string); ok {
			result = v
		} else {
			next := q.Iter()
			for x, ok := next(); ok; x, ok = next() {
				if pair, ok := x.(KeyVal); ok {
					result[fmt.Sprint(pair.Key)] = fmt.Sprint(pair.Val)
				}
			}
		}
	}
	return result
}

// S exports queryable into an interface{} slice
func (q *Queryable) S() []interface{} {
	result := []interface{}{}
	if q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			result = append(result, x)
		}
	}
	return result
}

// SAMap exports queryable into an slice of string to interface{} map
func (q *Queryable) SAMap() (result []map[string]interface{}) {
	result = []map[string]interface{}{}
	if q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if m, ok := x.(map[string]interface{}); ok {
				result = append(result, m)
			}
		}
	}
	return
}

// SAAMap exports queryable into an slice of string to string map
func (q *Queryable) SAAMap() (result []map[string]string) {
	result = []map[string]string{}
	if q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			m := map[string]string{}
			if md, ok := x.(map[string]string); ok {
				m = md
			} else if md, ok := x.(map[string]interface{}); ok {
				for k, v := range md {
					m[fmt.Sprint(k)] = fmt.Sprint(v)
				}
			}
			result = append(result, m)
		}
	}
	return
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
