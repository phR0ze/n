package nub

import (
	"reflect"
)

// Append items to the end of the collection and return the queryable
// converting to a collection if necessary
func (q *Queryable) Append(obj ...interface{}) *Queryable {
	if q.TypeSingle() {
		*q = *S().Append(q.ref.Interface())
	}

	// Append to slice type
	ref := reflect.ValueOf(obj)
	for i := 0; i < ref.Len(); i++ {
		*q.ref = reflect.Append(*q.ref, ref.Index(i))
	}
	q.Iter = sliceIter(*q.ref)

	return q
}
