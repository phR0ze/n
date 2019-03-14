package n

import (
	"fmt"
)

// O is an alias for interface{} and provides a number of export methods
type O interface{}

// QObj provides a number of export methods
type QObj struct {
	v interface{} // underlying value
}

// Value implements the Queryable interface
func (q *QObj) Value() interface{} {
	return q.v
}

// A exports the interface as a string
func (q *QObj) A() (result string) {
	if q != nil {
		switch v := q.v.(type) {
		case string:
			result = v
		default:
			result = fmt.Sprintf("%v", v)
		}
	}
	return
}
