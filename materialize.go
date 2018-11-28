package nub

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

// Strs materializes the results into a string slice
func (q *Queryable) Strs() (result []string) {
	q.Each(func(item interface{}) {
		result = append(result, item.(string))
	})
	return
}
