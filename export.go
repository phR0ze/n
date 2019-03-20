package n

import (
	"fmt"
	"strconv"
)

// Collecting functions that return external Go types here

// A exports numerable into a string
func (q *OldNumerable) A() (result string) {
	if q != nil && !q.Nil() {
		switch v := q.v.Interface().(type) {
		case string:
			result = v
		default:
			result = fmt.Sprintf("%v", v)
		}
	}
	return
}

// B exports numerable into a bool
func (q *OldNumerable) B() (result bool, err error) {
	if q != nil && !q.Nil() {
		switch v := q.v.Interface().(type) {
		case bool:
			result = v
		case int:
			result = v != 0
		case string:
			result, err = strconv.ParseBool(v)
		}
	}
	return
}

// I exports numerable into an int
func (q *OldNumerable) I() (result int, err error) {
	if q != nil && !q.Nil() {
		switch v := q.v.Interface().(type) {
		case int:
			result = v
		case bool:
			if v {
				result = 1
			}
		case string:
			result, err = strconv.Atoi(v)
		}
	}
	return
}

// Ints exports numerable into an int slice
func (q *OldNumerable) Ints() (result []int, err error) {
	result = []int{}
	if q != nil && !q.Nil() {
		if q.TypeSlice() {
			next := q.Iter()
			for x, ok := next(); ok; x, ok = next() {
				if i, ok := x.(int); ok {
					result = append(result, i)
				} else {
					err = fmt.Errorf("%v is not an int type", x)
				}
			}
		} else {
			err = fmt.Errorf("numerable is not a slice type")
		}
	}
	return
}

// M exports numerable into a map
func (q *OldNumerable) M() (result map[string]interface{}, err error) {
	result = map[string]interface{}{}
	if q != nil && !q.Nil() {
		if v, ok := q.O().(map[string]interface{}); ok {
			result = v
		} else {
			next := q.Iter()
			for x, ok := next(); ok; x, ok = next() {
				if pair, ok := x.(KeyVal); ok {
					result[fmt.Sprint(pair.Key)] = pair.Val
				} else {
					err = fmt.Errorf("not a key value pair type")
					return
				}
			}
		}
	}
	return
}

// O exports numerable into a interface{}
func (q *OldNumerable) O() interface{} {
	if q == nil || q.Nil() {
		return nil
	}
	return q.v.Interface()
}

// Strs exports numerable into an string slice
func (q *OldNumerable) Strs() (result []string) {
	result = []string{}
	if q != nil && !q.Nil() && q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			result = append(result, fmt.Sprint(x))
		}
	}
	return
}

// AAMap exports numerable into an string to string map
func (q *OldNumerable) AAMap() (result map[string]string, err error) {
	result = map[string]string{}
	if q != nil && !q.Nil() {
		if v, ok := q.O().(map[string]string); ok {
			result = v
		} else {
			next := q.Iter()
			for x, ok := next(); ok; x, ok = next() {
				if pair, ok := x.(KeyVal); ok {
					result[fmt.Sprint(pair.Key)] = fmt.Sprint(pair.Val)
				} else {
					err = fmt.Errorf("not a key value pair type")
					return
				}
			}
		}
	}
	return
}

// ASAMap exports numerable into an string to []string map
func (q *OldNumerable) ASAMap() (result map[string][]string, err error) {
	result = map[string][]string{}
	if q != nil && !q.Nil() {
		if v, ok := q.O().(map[string][]string); ok {
			result = v
		} else {
			next := q.Iter()
			for x, ok := next(); ok; x, ok = next() {
				if pair, ok := x.(KeyVal); ok {
					key := fmt.Sprint(pair.Key)
					if slice, ok := pair.Val.([]string); ok {
						result[key] = slice
					} else {
						result[key] = []string{}
						nexty := Q(pair.Val).Iter()
						for y, ok := nexty(); ok; y, ok = nexty() {
							result[key] = append(result[key], fmt.Sprint(y))
						}
					}
				} else {
					err = fmt.Errorf("not a key value pair type")
					return
				}
			}
		}
	}
	return
}

// S exports numerable into an interface{} slice
func (q *OldNumerable) S() []interface{} {
	result := []interface{}{}
	if q != nil && !q.Nil() && q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			result = append(result, x)
		}
	}
	return result
}

// SAMap exports numerable into an slice of string to interface{} map
func (q *OldNumerable) SAMap() (result []map[string]interface{}, err error) {
	result = []map[string]interface{}{}
	if q != nil && !q.Nil() && q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			m := map[string]interface{}{}
			switch x := x.(type) {
			case map[string]interface{}:
				m = x
			case map[interface{}]interface{}:
				for k, v := range x {
					m[fmt.Sprint(k)] = v
				}
			default:
				err = fmt.Errorf("%v is not of type map[string]interface{}", x)
				return
			}
			result = append(result, m)
		}
	}
	return
}

// SAAMap exports numerable into an slice of string to string map
func (q *OldNumerable) SAAMap() (result []map[string]string, err error) {
	result = []map[string]string{}
	if q != nil && !q.Nil() && q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			m := map[string]string{}
			switch x := x.(type) {
			case map[string]string:
				m = x
			case map[string]interface{}:
				for k, v := range x {
					m[fmt.Sprint(k)] = fmt.Sprint(v)
				}
			case map[interface{}]string:
				for k, v := range x {
					m[fmt.Sprint(k)] = fmt.Sprint(v)
				}
			case map[interface{}]interface{}:
				for k, v := range x {
					m[fmt.Sprint(k)] = fmt.Sprint(v)
				}
			default:
				err = fmt.Errorf("%v is not of type map[string]string", x)
				return
			}
			result = append(result, m)
		}
	}
	return
}
