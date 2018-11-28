package nub

import "reflect"

// TakeFirst remove an return the first item.
func (q *Queryable) TakeFirst() (item interface{}, ok bool) {
	switch q.v.Kind() {

	// Append to slice type
	case reflect.Array, reflect.Slice:
		if q.v.Len() > 0 {

			// Create new slice minus first
			n := reflect.MakeSlice(q.v.Type(), 0, q.v.Cap())
			for i := 1; i < q.v.Len(); i++ {
				n = reflect.Append(n, q.v.Index(i))
			}

			// Capture item, status and update queryable
			item = q.v.Index(0).Interface()
			ok = true
			*q.v = n
			q.Iter = sliceIter(*q.v)
		}
	}
	return
}

// // TakeFirstCnt updates the underlying slice and returns the items
// func (slice *intSliceNub) TakeFirstCnt(cnt int) (result *intSliceNub) {
// 	if cnt > 0 {
// 		if len(slice.raw) >= cnt {
// 			result = IntSlice(slice.raw[:cnt])
// 			slice.raw = slice.raw[cnt:]
// 		} else {
// 			result = IntSlice(slice.raw)
// 			slice.raw = []int{}
// 		}
// 	}
// 	return
// }

// // TakeLast updates the underlying slice and returns the item and status
// func (slice *intSliceNub) TakeLast() (int, bool) {
// 	if len(slice.raw) > 0 {
// 		item := slice.raw[len(slice.raw)-1]
// 		slice.raw = slice.raw[:len(slice.raw)-1]
// 		return item, true
// 	}
// 	return 0, false
// }

// // TakeLastCnt updates the underlying slice and returns the items
// func (slice *intSliceNub) TakeLastCnt(cnt int) (result *intSliceNub) {
// 	if cnt > 0 {
// 		if len(slice.raw) >= cnt {
// 			i := len(slice.raw) - cnt
// 			result = IntSlice(slice.raw[i:])
// 			slice.raw = slice.raw[:i]
// 		} else {
// 			result = IntSlice(slice.raw)
// 			slice.raw = []int{}
// 		}
// 	}
// 	return

// }
