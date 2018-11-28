package nub

// Collecting functions that return external Go types here

// A materializes queryable into a string
func (q *Queryable) A() string {
	return q.ref.Interface().(string)
}

// I materializes queryable into an int
func (q *Queryable) I() int {
	return q.ref.Interface().(int)
}

// M materializes queryable into a map
func (q *Queryable) M() map[string]interface{} {
	result := map[string]interface{}{}
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		pair := x.(*KeyVal)
		result[pair.Key.(string)] = pair.Val
	}
	return result
}

// O materializes queryable into a interface{}
func (q *Queryable) O() interface{} {
	return q.ref.Interface()
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

// Strs materializes queryable into an string slice
func (q *Queryable) Strs() []string {
	result := []string{}
	next := q.Iter()
	for x, ok := next(); ok; x, ok = next() {
		result = append(result, x.(string))
	}
	return result
}
