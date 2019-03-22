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

// Numerable interface methods
//--------------------------------------------------------------------------------------------------

// O returns the underlying data structure as is
func (n *NObj) O() interface{} {
	if n.Nil() {
		return nil
	}
	return n.o
}

// Nil tests if the numerable is nil
func (n *NObj) Nil() bool {
	if n == nil || n.o == nil {
		return true
	}
	return false
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

// ToUint casts an interface to a uint type.
func (n *NObj) ToUint() uint {
	v, _ := cast.ToUintE(n.o)
	return v
}

// ToUintE casts an interface to a uint type.
func (n *NObj) ToUintE() (uint, error) {
	return cast.ToUintE(n.o)
}

// ToUint8 casts an interface to a uint8 type.
func (n *NObj) ToUint8() uint8 {
	v, _ := cast.ToUint8E(n.o)
	return v
}

// ToUint8E casts an interface to a uint8 type.
func (n *NObj) ToUint8E() (uint8, error) {
	return cast.ToUint8E(n.o)
}

// ToUint16 casts an interface to a uint16 type.
func (n *NObj) ToUint16() uint16 {
	v, _ := cast.ToUint16E(n.o)
	return v
}

// ToUint16E casts an interface to a uint16 type.
func (n *NObj) ToUint16E() (uint16, error) {
	return cast.ToUint16E(n.o)
}

// ToUint32 casts an interface to a uint32 type.
func (n *NObj) ToUint32() uint32 {
	v, _ := cast.ToUint32E(n.o)
	return v
}

// ToUint32E casts an interface to a uint32 type.
func (n *NObj) ToUint32E() (uint32, error) {
	return cast.ToUint32E(n.o)
}

// ToUint64 casts an interface to a uint64 type.
func (n *NObj) ToUint64() uint64 {
	v, _ := cast.ToUint64E(n.o)
	return v
}

// ToUint64E casts an interface to a uint64 type.
func (n *NObj) ToUint64E() (uint64, error) {
	return cast.ToUint64E(n.o)
}

// String related
//--------------------------------------------------------------------------------------------------

// ToString casts an interface to a string type.
func (n *NObj) ToString() string {
	v, _ := cast.ToStringE(n.o)
	return v
}

// ToStringE casts an interface to a string type.
func (n *NObj) ToStringE() (string, error) {
	return cast.ToStringE(n.o)
}

// Map related
//--------------------------------------------------------------------------------------------------

// ToStringMapString casts an interface to a map[string]string type.
func (n *NObj) ToStringMapString() map[string]string {
	v, _ := cast.ToStringMapStringE(n.o)
	return v
}

// ToStringMapStringE casts an interface to a map[string]string type.
func (n *NObj) ToStringMapStringE() (map[string]string, error) {
	return cast.ToStringMapStringE(n.o)
}

// ToStringMapStringSlice casts an interface to a map[string][]string type.
func (n *NObj) ToStringMapStringSlice() map[string][]string {
	v, _ := cast.ToStringMapStringSliceE(n.o)
	return v
}

// ToStringMapStringSliceE casts an interface to a map[string][]string type.
func (n *NObj) ToStringMapStringSliceE() (map[string][]string, error) {
	return cast.ToStringMapStringSliceE(n.o)
}

// ToStringMapBool casts an interface to a map[string]bool type.
func (n *NObj) ToStringMapBool() map[string]bool {
	v, _ := cast.ToStringMapBoolE(n.o)
	return v
}

// ToStringMapBoolE casts an interface to a map[string]bool type.
func (n *NObj) ToStringMapBoolE() (map[string]bool, error) {
	return cast.ToStringMapBoolE(n.o)
}

// ToStringMapInt casts an interface to a map[string]int type.
func (n *NObj) ToStringMapInt() map[string]int {
	v, _ := cast.ToStringMapIntE(n.o)
	return v
}

// ToStringMapIntE casts an interface to a map[string]int type.
func (n *NObj) ToStringMapIntE() (map[string]int, error) {
	return cast.ToStringMapIntE(n.o)
}

// ToStringMapInt64 casts an interface to a map[string]int64 type.
func (n *NObj) ToStringMapInt64() map[string]int64 {
	v, _ := cast.ToStringMapInt64E(n.o)
	return v
}

// ToStringMapInt64E casts an interface to a map[string]int64 type.
func (n *NObj) ToStringMapInt64E() (map[string]int64, error) {
	return cast.ToStringMapInt64E(n.o)
}

// ToStringMap casts an interface to a map[string]interface{} type.
func (n *NObj) ToStringMap() map[string]interface{} {
	v, _ := cast.ToStringMapE(n.o)
	return v
}

// ToStringMapE casts an interface to a map[string]interface{} type.
func (n *NObj) ToStringMapE() (map[string]interface{}, error) {
	return cast.ToStringMapE(n.o)
}

// Slice related
//--------------------------------------------------------------------------------------------------

// ToSlice casts an interface to a []interface{} type.
func (n *NObj) ToSlice() []interface{} {
	v, _ := cast.ToSliceE(n.o)
	return v
}

// ToSliceE casts an interface to a []interface{} type.
func (n *NObj) ToSliceE() ([]interface{}, error) {
	return cast.ToSliceE(n.o)
}

// ToBoolSlice casts an interface to a []bool type.
func (n *NObj) ToBoolSlice() []bool {
	v, _ := cast.ToBoolSliceE(n.o)
	return v
}

// ToBoolSliceE casts an interface to a []bool type.
func (n *NObj) ToBoolSliceE() ([]bool, error) {
	return cast.ToBoolSliceE(n.o)
}

// ToStringSlice casts an interface to a []string type.
func (n *NObj) ToStringSlice() []string {
	v, _ := cast.ToStringSliceE(n.o)
	return v
}

// ToStringSliceE casts an interface to a []string type.
func (n *NObj) ToStringSliceE() ([]string, error) {
	return cast.ToStringSliceE(n.o)
}

// ToIntSlice casts an interface to a []int type.
func (n *NObj) ToIntSlice() []int {
	v, _ := cast.ToIntSliceE(n.o)
	return v
}

// ToIntSliceE casts an interface to a []int type.
func (n *NObj) ToIntSliceE() ([]int, error) {
	return cast.ToIntSliceE(n.o)
}

// ToDurationSlice casts an interface to a []time.Duration type.
func (n *NObj) ToDurationSlice() []time.Duration {
	v, _ := cast.ToDurationSliceE(n.o)
	return v
}

// ToDurationSliceE casts an interface to a []time.Duration type.
func (n *NObj) ToDurationSliceE() ([]time.Duration, error) {
	return cast.ToDurationSliceE(n.o)
}
