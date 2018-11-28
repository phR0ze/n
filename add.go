package nub

import (
	"reflect"
)

// Append items to the end of the collection and return the queryable
// converting to a collection if necessary
func (q *Queryable) Append(obj ...interface{}) *Queryable {
	if q.TypeSingle() {
		*q = *S().Append(q.v.Interface())
	}

	// Append to slice type
	for i := 0; i < len(obj); i++ {
		item := reflect.ValueOf(obj[i])
		*q.v = reflect.Append(*q.v, item)
	}
	q.Iter = sliceIter(*q.v)

	return q
}
