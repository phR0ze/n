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
