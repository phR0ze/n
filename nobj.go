package n

import (
	"time"

	"github.com/phR0ze/cast"
)

// NObj is a wrapper around an interface{} value wproviding a number of export methods
// for casting and converting to other types via the excellent cast package.
type NObj struct {
	o interface{} // value
}

// Bool related
//--------------------------------------------------------------------------------------------------

// ToBool casts an interface to a bool type.
func (n *NObj) ToBool() bool {
	v, _ := cast.ToBoolE(n.o)
	return v
}

// ToBoolE casts an interface to a bool type.
func (n *NObj) ToBoolE() (bool, error) {
	return cast.ToBoolE(n.o)
}

// Time related
//--------------------------------------------------------------------------------------------------

// ToTime casts an interface to a time.Time type.
func (n *NObj) ToTime() time.Time {
	v, _ := cast.ToTimeE(n.o)
	return v
}

// ToTimeE casts an interface to a time.Time type.
func (n *NObj) ToTimeE() (time.Time, error) {
	return cast.ToTimeE(n.o)
}

// ToDuration casts an interface to a time.Duration type.
func (n *NObj) ToDuration() time.Duration {
	v, _ := cast.ToDurationE(n.o)
	return v
}

// ToDurationE casts an interface to a time.Duration type.
func (n *NObj) ToDurationE() (time.Duration, error) {
	return cast.ToDurationE(n.o)
}

// Float related
//--------------------------------------------------------------------------------------------------

// ToFloat32 casts an interface to a float32 type.
func (n *NObj) ToFloat32() float32 {
	v, _ := cast.ToFloat32E(n.o)
	return v
}

// ToFloat32E casts an interface to a float32 type.
func (n *NObj) ToFloat32E() (float32, error) {
	return cast.ToFloat32E(n.o)
}

// ToFloat64 casts an interface to a float64 type.
func (n *NObj) ToFloat64() float64 {
	v, _ := cast.ToFloat64E(n.o)
	return v
}

// ToFloat64E casts an interface to a float64 type.
func (n *NObj) ToFloat64E() (float64, error) {
	return cast.ToFloat64E(n.o)
}

// Int related
//--------------------------------------------------------------------------------------------------

// ToInt casts an interface to an int type.
func (n *NObj) ToInt() int {
	v, _ := cast.ToIntE(n.o)
	return v
}

// ToIntE casts an interface to an int type.
func (n *NObj) ToIntE() (int, error) {
	return cast.ToIntE(n.o)
}

// ToInt8 casts an interface to an int8 type.
func (n *NObj) ToInt8() int8 {
	v, _ := cast.ToInt8E(n.o)
	return v
}

// ToInt8E casts an interface to an int8 type.
func (n *NObj) ToInt8E() (int8, error) {
	return cast.ToInt8E(n.o)
}

// ToInt16 casts an interface to an int16 type.
func (n *NObj) ToInt16() int16 {
	v, _ := cast.ToInt16E(n.o)
	return v
}

// ToInt16E casts an interface to an int16 type.
func (n *NObj) ToInt16E() (int16, error) {
	return cast.ToInt16E(n.o)
}

// ToInt32 casts an interface to an int32 type.
func (n *NObj) ToInt32() int32 {
	v, _ := cast.ToInt32E(n.o)
	return v
}

// ToInt32E casts an interface to an int32 type.
func (n *NObj) ToInt32E() (int32, error) {
	return cast.ToInt32E(n.o)
}

// ToInt64 casts an interface to an int64 type.
func (n *NObj) ToInt64() int64 {
	v, _ := cast.ToInt64E(n.o)
	return v
}

// ToInt64E casts an interface to an int64 type.
func (n *NObj) ToInt64E() (int64, error) {
	return cast.ToInt64E(n.o)
}

// // ToUint casts an interface to a uint type.
// func ToUint(i interface{}) uint {
// 	v, _ := ToUintE(i)
// 	return v
// }

// // ToUint64 casts an interface to a uint64 type.
// func ToUint64(i interface{}) uint64 {
// 	v, _ := ToUint64E(i)
// 	return v
// }

// // ToUint32 casts an interface to a uint32 type.
// func ToUint32(i interface{}) uint32 {
// 	v, _ := ToUint32E(i)
// 	return v
// }

// // ToUint16 casts an interface to a uint16 type.
// func ToUint16(i interface{}) uint16 {
// 	v, _ := ToUint16E(i)
// 	return v
// }

// // ToUint8 casts an interface to a uint8 type.
// func ToUint8(i interface{}) uint8 {
// 	v, _ := ToUint8E(i)
// 	return v
// }

// // ToString casts an interface to a string type.
// func ToString(i interface{}) string {
// 	v, _ := ToStringE(i)
// 	return v
// }

// // ToStringMapString casts an interface to a map[string]string type.
// func ToStringMapString(i interface{}) map[string]string {
// 	v, _ := ToStringMapStringE(i)
// 	return v
// }

// // ToStringMapStringSlice casts an interface to a map[string][]string type.
// func ToStringMapStringSlice(i interface{}) map[string][]string {
// 	v, _ := ToStringMapStringSliceE(i)
// 	return v
// }

// // ToStringMapBool casts an interface to a map[string]bool type.
// func ToStringMapBool(i interface{}) map[string]bool {
// 	v, _ := ToStringMapBoolE(i)
// 	return v
// }

// // ToStringMapInt casts an interface to a map[string]int type.
// func ToStringMapInt(i interface{}) map[string]int {
// 	v, _ := ToStringMapIntE(i)
// 	return v
// }

// // ToStringMapInt64 casts an interface to a map[string]int64 type.
// func ToStringMapInt64(i interface{}) map[string]int64 {
// 	v, _ := ToStringMapInt64E(i)
// 	return v
// }

// // ToStringMap casts an interface to a map[string]interface{} type.
// func ToStringMap(i interface{}) map[string]interface{} {
// 	v, _ := ToStringMapE(i)
// 	return v
// }

// // ToSlice casts an interface to a []interface{} type.
// func ToSlice(i interface{}) []interface{} {
// 	v, _ := ToSliceE(i)
// 	return v
// }

// // ToBoolSlice casts an interface to a []bool type.
// func ToBoolSlice(i interface{}) []bool {
// 	v, _ := ToBoolSliceE(i)
// 	return v
// }

// // ToStringSlice casts an interface to a []string type.
// func ToStringSlice(i interface{}) []string {
// 	v, _ := ToStringSliceE(i)
// 	return v
// }

// // ToIntSlice casts an interface to a []int type.
// func ToIntSlice(i interface{}) []int {
// 	v, _ := ToIntSliceE(i)
// 	return v
// }

// // ToDurationSlice casts an interface to a []time.Duration type.
// func ToDurationSlice(i interface{}) []time.Duration {
// 	v, _ := ToDurationSliceE(i)
// 	return v
// }
