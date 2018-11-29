package n

import "reflect"

// TakeFirst remove an return the first item.
func (q *Queryable) TakeFirst() (item interface{}, ok bool) {
	if !q.TypeSingle() && q.v.Len() > 0 {
		if q.Kind == reflect.Array || q.Kind == reflect.Slice {

			// Make a new slice minus the first one
			n := reflect.MakeSlice(q.v.Type(), 0, q.v.Len()-1)
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

// TakeFirstCnt remove cnt items and return them
func (q *Queryable) TakeFirstCnt(cnt int) (result interface{}) {
	result = []interface{}{}
	if !q.TypeSingle() && q.v.Len() > 0 && cnt > 0 {
		if q.Kind == reflect.Array || q.Kind == reflect.Slice {

			// This slice is larger that asked for
			if q.Len() >= cnt {

				// Copy out the first cnt
				// items := reflect.MakeSlice(q.v.Type(), 0, cnt)
				// for i := 0; i < cnt; i++ {

				// }

				// // Make a new slice minus the first cnt
				// n := reflect.MakeSlice(q.v.Type(), 0, q.v.Cap())
				// for i := 1; i < q.v.Len(); i++ {
				// 	n = reflect.Append(n, q.v.Index(i))
				// }

				// // Capture item, status and update queryable
				// item = q.v.Index(0).Interface()
				// ok = true
				// *q.v = n
			} else {
				result = q.O()
				*q.v = reflect.MakeSlice(q.v.Type(), 0, q.v.Cap())
			}
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
