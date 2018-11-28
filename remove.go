package nub

import "reflect"

// TakeFirst remove an return the first item
func (q *Queryable) TakeFirst() (item interface{}, ok bool) {
	switch q.ref.Kind() {

	// Append to slice type
	case reflect.Array, reflect.Slice:
		if q.ref.Len() > 0 {
			// ref := reflect.MakeSlice(q.ref.Type(), 0, q.ref.Cap())
			// for i := 0; i < q.ref.Len(); i++ {

			// }
			// //*q.ref = reflect.Append(*q.ref, ref.Index(i))
			// //}
			// q.Iter = sliceIter(*q.ref)
		}

	// Append to map type
	case reflect.Map:
		panic("TODO: implement append for TakeFirst")

	// Not a collection type create a new queryable
	default:
		//*q = *S().Append(q.ref.Interface())
	}
	return nil, false
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
