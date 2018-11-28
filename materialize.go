package nub

// Collecting functions that return external Go types here

// Int materializes the result into an int
func (q *Queryable) Int() int {
	return q.O.(int)
}

// Ints materializes the results into a int slice
func (q *Queryable) Ints() (result []int) {
	q.Each(func(item interface{}) {
		result = append(result, item.(int))
	})
	return
}

// Str materializes the result into a string
func (q *Queryable) Str() string {
	return q.O.(string)
}

// StrMap materializes the result from the specific string to interface map
// called out by the given dot notation key.
func (q *Queryable) StrMap(target ...string) map[string]interface{} {
	//keys := Q(target).Join

	result := map[string]interface{}{}
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		pair := x.(*KeyVal)
		result[pair.Key.(string)] = pair.Val
	}
	return result
}

// StrMap returns a map of interface from the given map using the given key
// func (q *Queryable) StrMap(key string) *Queryable {
// 	result := NewStrMap()

// 	keys := Str(key).Split(".")
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

// Strs materializes the results into a string slice
func (q *Queryable) Strs() (result []string) {
	q.Each(func(item interface{}) {
		result = append(result, item.(string))
	})
	return
}
