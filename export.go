package n

import (
	"fmt"
	"strconv"
)

// Collecting functions that return external Go types here

// A exports queryable into a string
func (q *OldQueryable) A() (result string) {
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

// B exports queryable into a bool
func (q *OldQueryable) B() (result bool, err error) {
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

// I exports queryable into an int
func (q *OldQueryable) I() (result int, err error) {
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

// Ints exports queryable into an int slice
func (q *OldQueryable) Ints() (result []int, err error) {
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
			err = fmt.Errorf("queryable is not a slice type")
		}
	}
	return
}

// M exports queryable into a map
func (q *OldQueryable) M() (result map[string]interface{}, err error) {
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

// O exports queryable into a interface{}
func (q *OldQueryable) O() interface{} {
	if q == nil || q.Nil() {
		return nil
	}
	return q.v.Interface()
}

// Strs exports queryable into an string slice
func (q *OldQueryable) Strs() (result []string) {
	result = []string{}
	if q != nil && !q.Nil() && q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			result = append(result, fmt.Sprint(x))
		}
	}
	return
}

// AAMap exports queryable into an string to string map
func (q *OldQueryable) AAMap() (result map[string]string, err error) {
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

// ASAMap exports queryable into an string to []string map
func (q *OldQueryable) ASAMap() (result map[string][]string, err error) {
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

// S exports queryable into an interface{} slice
func (q *OldQueryable) S() []interface{} {
	result := []interface{}{}
	if q != nil && !q.Nil() && q.TypeSlice() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			result = append(result, x)
		}
	}
	return result
}

// SAMap exports queryable into an slice of string to interface{} map
func (q *OldQueryable) SAMap() (result []map[string]interface{}, err error) {
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

// SAAMap exports queryable into an slice of string to string map
func (q *OldQueryable) SAAMap() (result []map[string]string, err error) {
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
