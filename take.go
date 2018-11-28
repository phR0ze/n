package nub

// TakeFirst updates the underlying slice and returns the item and status
func (q *Queryable) TakeFirst() (int, bool) {
	// if len(slice.raw) > 0 {
	// 	item := slice.raw[0]
	// 	slice.raw = slice.raw[1:]
	// 	return item, true
	// }
	return 0, false
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
