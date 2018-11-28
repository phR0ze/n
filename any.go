package nub

// Any checks if the queryable has anything in it
func (q *Queryable) Any() bool {
	if !q.Singular() {
		return q.ref.Len() > 0
	}
	return q.O != nil
}

// AnyWhere checka if any match the given lambda
func (q *Queryable) AnyWhere(lambda func(item interface{}) bool) bool {
	if !q.Singular() {
		//next := q.Iter
	}
	return false
}
