package nub

// Any checks if the queryable has anything in it
func (q *Queryable) Any() bool {
	if q.Iter != nil {
		return q.ref.Len() > 0
	}
	return q.O != nil
}

// AnyWhere checka if any match the given lambda
func (q *Queryable) AnyWhere(lambda func(interface{}) bool) bool {
	if !q.Singular() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if lambda(x) {
				return true
			}
		}
	}
	return false
}
