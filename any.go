package nub

// Any checks if the queryable has anything in it
func (q *Queryable) Any() bool {
	if q.Iter != nil {
		return q.ref.Len() > 0
	}
	return q.ref.Interface() != nil
}

// AnyWhere checka if any match the given lambda
func (q *Queryable) AnyWhere(lambda func(interface{}) bool) bool {
	if !q.TypeSingle() {
		next := q.Iter()
		for x, ok := next(); ok; x, ok = next() {
			if lambda(x) {
				return true
			}
		}
	} else if lambda(q.ref.Interface()) {
		return true
	}
	return false
}
