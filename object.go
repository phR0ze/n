package n

import (
	"time"

	"github.com/phR0ze/cast"
)

// Object is a wrapper around an interface{} value wproviding a number of export methods
// for casting and converting to other types via the excellent cast package.
type Object struct {
	o interface{} // value
}

// Obj creates a new Object from the given obj appending the Object methods
func Obj(obj interface{}) *Object {
	return &Object{obj}
}

// Object interface methods
//--------------------------------------------------------------------------------------------------

// O returns the underlying data structure as is
func (p *Object) O() interface{} {
	if p.Nil() {
		return nil
	}
	return p.o
}

// Nil tests if the numerable is nil
func (p *Object) Nil() bool {
	if p == nil || p.o == nil {
		return true
	}
	return false
}

// A is an alias to String for brevity
func (p *Object) A() string {
	return p.String()
}

// String returns a string representation of the Object, implements Stringer interface.
func (p *Object) String() string {
	if p == nil {
		return ""
	}
	return ToString(p.o)
}

// Bool related
//--------------------------------------------------------------------------------------------------

// ToBool casts an interface to a bool type.
func (p *Object) ToBool() bool {
	v, _ := ToBoolE(p.o)
	return v
}

// ToBoolE casts an interface to a bool type.
func (p *Object) ToBoolE() (bool, error) {
	return ToBoolE(p.o)
}

// Time related
//--------------------------------------------------------------------------------------------------

// ToTime casts an interface to a time.Time type.
func (p *Object) ToTime() time.Time {
	v, _ := cast.ToTimeE(p.o)
	return v
}

// ToTimeE casts an interface to a time.Time type.
func (p *Object) ToTimeE() (time.Time, error) {
	return cast.ToTimeE(p.o)
}

// ToDuration casts an interface to a time.Duration type.
func (p *Object) ToDuration() time.Duration {
	v, _ := cast.ToDurationE(p.o)
	return v
}

// ToDurationE casts an interface to a time.Duration type.
func (p *Object) ToDurationE() (time.Duration, error) {
	return cast.ToDurationE(p.o)
}

// Float related
//--------------------------------------------------------------------------------------------------

// ToFloat32 casts an interface to a float32 type.
func (p *Object) ToFloat32() float32 {
	v, _ := cast.ToFloat32E(p.o)
	return v
}

// ToFloat32E casts an interface to a float32 type.
func (p *Object) ToFloat32E() (float32, error) {
	return cast.ToFloat32E(p.o)
}

// ToFloat64 casts an interface to a float64 type.
func (p *Object) ToFloat64() float64 {
	v, _ := cast.ToFloat64E(p.o)
	return v
}

// ToFloat64E casts an interface to a float64 type.
func (p *Object) ToFloat64E() (float64, error) {
	return cast.ToFloat64E(p.o)
}

// Int related
//--------------------------------------------------------------------------------------------------

// ToInt casts an interface to an int type.
func (p *Object) ToInt() int {
	v, _ := ToIntE(p.o)
	return v
}

// ToIntE casts an interface to an int type.
func (p *Object) ToIntE() (int, error) {
	return ToIntE(p.o)
}

// ToInt8 casts an interface to an int8 type.
func (p *Object) ToInt8() int8 {
	v, _ := cast.ToInt8E(p.o)
	return v
}

// ToInt8E casts an interface to an int8 type.
func (p *Object) ToInt8E() (int8, error) {
	return cast.ToInt8E(p.o)
}

// ToInt16 casts an interface to an int16 type.
func (p *Object) ToInt16() int16 {
	v, _ := cast.ToInt16E(p.o)
	return v
}

// ToInt16E casts an interface to an int16 type.
func (p *Object) ToInt16E() (int16, error) {
	return cast.ToInt16E(p.o)
}

// ToInt32 casts an interface to an int32 type.
func (p *Object) ToInt32() int32 {
	v, _ := cast.ToInt32E(p.o)
	return v
}

// ToInt32E casts an interface to an int32 type.
func (p *Object) ToInt32E() (int32, error) {
	return cast.ToInt32E(p.o)
}

// ToInt64 casts an interface to an int64 type.
func (p *Object) ToInt64() int64 {
	v, _ := cast.ToInt64E(p.o)
	return v
}

// ToInt64E casts an interface to an int64 type.
func (p *Object) ToInt64E() (int64, error) {
	return cast.ToInt64E(p.o)
}

// ToUint casts an interface to a uint type.
func (p *Object) ToUint() uint {
	v, _ := cast.ToUintE(p.o)
	return v
}

// ToUintE casts an interface to a uint type.
func (p *Object) ToUintE() (uint, error) {
	return cast.ToUintE(p.o)
}

// ToUint8 casts an interface to a uint8 type.
func (p *Object) ToUint8() uint8 {
	v, _ := cast.ToUint8E(p.o)
	return v
}

// ToUint8E casts an interface to a uint8 type.
func (p *Object) ToUint8E() (uint8, error) {
	return cast.ToUint8E(p.o)
}

// ToUint16 casts an interface to a uint16 type.
func (p *Object) ToUint16() uint16 {
	v, _ := cast.ToUint16E(p.o)
	return v
}

// ToUint16E casts an interface to a uint16 type.
func (p *Object) ToUint16E() (uint16, error) {
	return cast.ToUint16E(p.o)
}

// ToUint32 casts an interface to a uint32 type.
func (p *Object) ToUint32() uint32 {
	v, _ := cast.ToUint32E(p.o)
	return v
}

// ToUint32E casts an interface to a uint32 type.
func (p *Object) ToUint32E() (uint32, error) {
	return cast.ToUint32E(p.o)
}

// ToUint64 casts an interface to a uint64 type.
func (p *Object) ToUint64() uint64 {
	v, _ := cast.ToUint64E(p.o)
	return v
}

// ToUint64E casts an interface to a uint64 type.
func (p *Object) ToUint64E() (uint64, error) {
	return cast.ToUint64E(p.o)
}

// String related
//--------------------------------------------------------------------------------------------------

// ToString casts an interface to a string type.
func (p *Object) ToString() string {
	return p.String()
}

// ToStringE casts an interface to a string type.
func (p *Object) ToStringE() (string, error) {
	return cast.ToStringE(p.o)
}

// Map related
//--------------------------------------------------------------------------------------------------

// ToStringMapString casts an interface to a map[string]string type.
func (p *Object) ToStringMapString() map[string]string {
	v, _ := cast.ToStringMapStringE(p.o)
	return v
}

// ToStringMapStringE casts an interface to a map[string]string type.
func (p *Object) ToStringMapStringE() (map[string]string, error) {
	return cast.ToStringMapStringE(p.o)
}

// ToStringMapStringSlice casts an interface to a map[string][]string type.
func (p *Object) ToStringMapStringSlice() map[string][]string {
	v, _ := cast.ToStringMapStringSliceE(p.o)
	return v
}

// ToStringMapStringSliceE casts an interface to a map[string][]string type.
func (p *Object) ToStringMapStringSliceE() (map[string][]string, error) {
	return cast.ToStringMapStringSliceE(p.o)
}

// ToStringMapBool casts an interface to a map[string]bool type.
func (p *Object) ToStringMapBool() map[string]bool {
	v, _ := cast.ToStringMapBoolE(p.o)
	return v
}

// ToStringMapBoolE casts an interface to a map[string]bool type.
func (p *Object) ToStringMapBoolE() (map[string]bool, error) {
	return cast.ToStringMapBoolE(p.o)
}

// ToStringMapInt casts an interface to a map[string]int type.
func (p *Object) ToStringMapInt() map[string]int {
	v, _ := cast.ToStringMapIntE(p.o)
	return v
}

// ToStringMapIntE casts an interface to a map[string]int type.
func (p *Object) ToStringMapIntE() (map[string]int, error) {
	return cast.ToStringMapIntE(p.o)
}

// ToStringMapInt64 casts an interface to a map[string]int64 type.
func (p *Object) ToStringMapInt64() map[string]int64 {
	v, _ := cast.ToStringMapInt64E(p.o)
	return v
}

// ToStringMapInt64E casts an interface to a map[string]int64 type.
func (p *Object) ToStringMapInt64E() (map[string]int64, error) {
	return cast.ToStringMapInt64E(p.o)
}

// ToStringMap casts an interface to a map[string]interface{} type.
func (p *Object) ToStringMap() map[string]interface{} {
	v, _ := cast.ToStringMapE(p.o)
	return v
}

// ToStringMapE casts an interface to a map[string]interface{} type.
func (p *Object) ToStringMapE() (map[string]interface{}, error) {
	return cast.ToStringMapE(p.o)
}

// Slice related
//--------------------------------------------------------------------------------------------------

// ToSlice casts an interface to a []interface{} type.
func (p *Object) ToSlice() []interface{} {
	v, _ := cast.ToSliceE(p.o)
	return v
}

// ToSliceE casts an interface to a []interface{} type.
func (p *Object) ToSliceE() ([]interface{}, error) {
	return cast.ToSliceE(p.o)
}

// ToBoolSlice casts an interface to a []bool type.
func (p *Object) ToBoolSlice() []bool {
	v, _ := cast.ToBoolSliceE(p.o)
	return v
}

// ToBoolSliceE casts an interface to a []bool type.
func (p *Object) ToBoolSliceE() ([]bool, error) {
	return cast.ToBoolSliceE(p.o)
}

// ToStringSlice casts an interface to a []string type.
func (p *Object) ToStringSlice() []string {
	v, _ := cast.ToStringSliceE(p.o)
	return v
}

// ToStringSliceE casts an interface to a []string type.
func (p *Object) ToStringSliceE() ([]string, error) {
	return cast.ToStringSliceE(p.o)
}

// ToIntSlice casts an interface to a []int type.
func (p *Object) ToIntSlice() []int {
	v, _ := cast.ToIntSliceE(p.o)
	return v
}

// ToIntSliceE casts an interface to a []int type.
func (p *Object) ToIntSliceE() ([]int, error) {
	return cast.ToIntSliceE(p.o)
}

// ToDurationSlice casts an interface to a []time.Duration type.
func (p *Object) ToDurationSlice() []time.Duration {
	v, _ := cast.ToDurationSliceE(p.o)
	return v
}

// ToDurationSliceE casts an interface to a []time.Duration type.
func (p *Object) ToDurationSliceE() ([]time.Duration, error) {
	return cast.ToDurationSliceE(p.o)
}
