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

// M is an alias to ToStringMapG for brevity
func (p *Object) M() map[string]interface{} {
	return p.ToStringMapG()
}

// S is an alias to ToStringSliceG for brevity
func (p *Object) S() []string {
	return p.ToStringSliceG()
}

// String returns a string representation of the Object, implements Stringer interface.
func (p *Object) String() string {
	if p == nil {
		return ""
	}
	return ToString(p.o)
}

// Bool
//--------------------------------------------------------------------------------------------------

// ToBool casts an interface to a bool type.
func (p *Object) ToBool() bool {
	if p == nil {
		return false
	}
	v, _ := ToBoolE(p.o)
	return v
}

// ToBoolE casts an interface to a bool type.
func (p *Object) ToBoolE() (bool, error) {
	if p == nil {
		return false, nil
	}
	return ToBoolE(p.o)
}

// Char
//--------------------------------------------------------------------------------------------------

// C is an alias to ToChar for brevity
func (p *Object) C() *Char {
	return p.ToChar()
}

// ToChar casts an interface to a *Char type.
func (p *Object) ToChar() *Char {
	if p == nil {
		val := Char(0)
		return &val
	}
	return ToChar(p.o)
}

// R is an alias to ToRune for brevity
func (p *Object) R() rune {
	return p.ToRune()
}

// ToRune casts an interface to a rune type.
func (p *Object) ToRune() rune {
	if p == nil {
		return rune(0)
	}
	return rune(*ToChar(p.o))
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
	if p == nil {
		return float32(0)
	}
	return ToFloat32(p.o)
}

// ToFloat32E casts an interface to a float32 type.
func (p *Object) ToFloat32E() (float32, error) {
	if p == nil {
		return float32(0), nil
	}
	return ToFloat32E(p.o)
}

// ToFloat64 converts an interface to a float64 type.
func (p *Object) ToFloat64() float64 {
	if p == nil {
		return float64(0)
	}
	return ToFloat64(p.o)
}

// ToFloat64E converts an interface to a float64 type.
func (p *Object) ToFloat64E() (float64, error) {
	if p == nil {
		return float64(0), nil
	}
	return ToFloat64E(p.o)
}

// Int related
//--------------------------------------------------------------------------------------------------

// ToInt casts an interface to an int type.
func (p *Object) ToInt() int {
	if p == nil {
		return 0
	}
	v, _ := ToIntE(p.o)
	return v
}

// ToIntE casts an interface to an int type.
func (p *Object) ToIntE() (int, error) {
	if p == nil {
		return 0, nil
	}
	return ToIntE(p.o)
}

// ToInt8 casts an interface to an int8 type.
func (p *Object) ToInt8() int8 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToInt8E(p.o)
	return v
}

// ToInt8E casts an interface to an int8 type.
func (p *Object) ToInt8E() (int8, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToInt8E(p.o)
}

// ToInt16 casts an interface to an int16 type.
func (p *Object) ToInt16() int16 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToInt16E(p.o)
	return v
}

// ToInt16E casts an interface to an int16 type.
func (p *Object) ToInt16E() (int16, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToInt16E(p.o)
}

// ToInt32 casts an interface to an int32 type.
func (p *Object) ToInt32() int32 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToInt32E(p.o)
	return v
}

// ToInt32E casts an interface to an int32 type.
func (p *Object) ToInt32E() (int32, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToInt32E(p.o)
}

// ToInt64 casts an interface to an int64 type.
func (p *Object) ToInt64() int64 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToInt64E(p.o)
	return v
}

// ToInt64E casts an interface to an int64 type.
func (p *Object) ToInt64E() (int64, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToInt64E(p.o)
}

// ToUint casts an interface to a uint type.
func (p *Object) ToUint() uint {
	if p == nil {
		return 0
	}
	v, _ := cast.ToUintE(p.o)
	return v
}

// ToUintE casts an interface to a uint type.
func (p *Object) ToUintE() (uint, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToUintE(p.o)
}

// ToUint8 casts an interface to a uint8 type.
func (p *Object) ToUint8() uint8 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToUint8E(p.o)
	return v
}

// ToUint8E casts an interface to a uint8 type.
func (p *Object) ToUint8E() (uint8, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToUint8E(p.o)
}

// ToUint16 casts an interface to a uint16 type.
func (p *Object) ToUint16() uint16 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToUint16E(p.o)
	return v
}

// ToUint16E casts an interface to a uint16 type.
func (p *Object) ToUint16E() (uint16, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToUint16E(p.o)
}

// ToUint32 casts an interface to a uint32 type.
func (p *Object) ToUint32() uint32 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToUint32E(p.o)
	return v
}

// ToUint32E casts an interface to a uint32 type.
func (p *Object) ToUint32E() (uint32, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToUint32E(p.o)
}

// ToUint64 casts an interface to a uint64 type.
func (p *Object) ToUint64() uint64 {
	if p == nil {
		return 0
	}
	v, _ := cast.ToUint64E(p.o)
	return v
}

// ToUint64E casts an interface to a uint64 type.
func (p *Object) ToUint64E() (uint64, error) {
	if p == nil {
		return 0, nil
	}
	return cast.ToUint64E(p.o)
}

// String related
//--------------------------------------------------------------------------------------------------

// ToStr converts object into a *Str
func (p *Object) ToStr() *Str {
	if p == nil {
		return ToStr(p)
	}
	return ToStr(p.o)
}

// ToString casts an interface to a string type.
func (p *Object) ToString() string {
	if p == nil {
		return ""
	}
	return p.String()
}

// ToStringE casts an interface to a string type.
func (p *Object) ToStringE() (string, error) {
	if p == nil {
		return "", nil
	}
	return p.ToString(), nil
}

// Map related
//--------------------------------------------------------------------------------------------------

// ToStringMap converts an interface to a *StringMap type.
func (p *Object) ToStringMap() *StringMap {
	if p == nil {
		return NewStringMapV()
	}
	return ToStringMap(p.o)
}

// ToStringMapE converts an interface to a *StringMap type.
func (p *Object) ToStringMapE() (*StringMap, error) {
	if p == nil {
		return NewStringMapV(), nil
	}
	return ToStringMapE(p.o)
}

// ToStringMapG converts an interface to a map[string]interface{} type.
func (p *Object) ToStringMapG() map[string]interface{} {
	if p == nil {
		return map[string]interface{}{}
	}
	v, _ := ToStringMapE(p.o)
	return v.G()
}

// ToStringMapGE converts an interface to a map[string]interface{} type.
func (p *Object) ToStringMapGE() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if p == nil {
		return result, nil
	}
	v, err := ToStringMapE(p.o)
	if err != nil {
		return result, err
	}
	return v.G(), err
}

// ToStringMapString casts an interface to a map[string]string type.
func (p *Object) ToStringMapString() map[string]string {
	if p == nil {
		return map[string]string{}
	}
	v, _ := cast.ToStringMapStringE(p.o)
	return v
}

// ToStringMapStringE casts an interface to a map[string]string type.
func (p *Object) ToStringMapStringE() (map[string]string, error) {
	if p == nil {
		return map[string]string{}, nil
	}
	return cast.ToStringMapStringE(p.o)
}

// ToStringMapStringSlice casts an interface to a map[string][]string type.
func (p *Object) ToStringMapStringSlice() map[string][]string {
	if p == nil {
		return map[string][]string{}
	}
	v, _ := cast.ToStringMapStringSliceE(p.o)
	return v
}

// ToStringMapStringSliceE casts an interface to a map[string][]string type.
func (p *Object) ToStringMapStringSliceE() (map[string][]string, error) {
	if p == nil {
		return map[string][]string{}, nil
	}
	return cast.ToStringMapStringSliceE(p.o)
}

// ToStringMapBool casts an interface to a map[string]bool type.
func (p *Object) ToStringMapBool() map[string]bool {
	if p == nil {
		return map[string]bool{}
	}
	v, _ := cast.ToStringMapBoolE(p.o)
	return v
}

// ToStringMapBoolE casts an interface to a map[string]bool type.
func (p *Object) ToStringMapBoolE() (map[string]bool, error) {
	if p == nil {
		return map[string]bool{}, nil
	}
	return cast.ToStringMapBoolE(p.o)
}

// ToStringMapInt casts an interface to a map[string]int type.
func (p *Object) ToStringMapInt() map[string]int {
	if p == nil {
		return map[string]int{}
	}
	v, _ := cast.ToStringMapIntE(p.o)
	return v
}

// ToStringMapIntE casts an interface to a map[string]int type.
func (p *Object) ToStringMapIntE() (map[string]int, error) {
	if p == nil {
		return map[string]int{}, nil
	}
	return cast.ToStringMapIntE(p.o)
}

// ToStringMapInt64 casts an interface to a map[string]int64 type.
func (p *Object) ToStringMapInt64() map[string]int64 {
	if p == nil {
		return map[string]int64{}
	}
	v, _ := cast.ToStringMapInt64E(p.o)
	return v
}

// ToStringMapInt64E casts an interface to a map[string]int64 type.
func (p *Object) ToStringMapInt64E() (map[string]int64, error) {
	if p == nil {
		return map[string]int64{}, nil
	}
	return cast.ToStringMapInt64E(p.o)
}

// Slice related
//--------------------------------------------------------------------------------------------------

// ToSlice casts an interface to a []interface{} type.
func (p *Object) ToSlice() []interface{} {
	if p == nil {
		return []interface{}{}
	}
	v, _ := cast.ToSliceE(p.o)
	return v
}

// ToSliceE casts an interface to a []interface{} type.
func (p *Object) ToSliceE() ([]interface{}, error) {
	if p == nil {
		return []interface{}{}, nil
	}
	return cast.ToSliceE(p.o)
}

// ToBoolSlice casts an interface to a []bool type.
func (p *Object) ToBoolSlice() []bool {
	if p == nil {
		return []bool{}
	}
	v, _ := cast.ToBoolSliceE(p.o)
	return v
}

// ToBoolSliceE casts an interface to a []bool type.
func (p *Object) ToBoolSliceE() ([]bool, error) {
	if p == nil {
		return []bool{}, nil
	}
	return cast.ToBoolSliceE(p.o)
}

// ToMapSlice converts an interface to a *MapSlice type.
func (p *Object) ToMapSlice() *MapSlice {
	if p == nil {
		return NewMapSliceV()
	}
	return ToMapSlice(p.o)
}

// ToMapSliceE converts an interface to a *MapSlice type.
func (p *Object) ToMapSliceE() (*MapSlice, error) {
	if p == nil {
		return NewMapSliceV(), nil
	}
	return ToMapSliceE(p.o)
}

// ToMapSliceG converts an interface to a []map[string]interface{} type.
func (p *Object) ToMapSliceG() []map[string]interface{} {
	if p == nil {
		return NewMapSliceV().G()
	}
	return ToMapSlice(p.o).G()
}

// ToMapSliceGE converts an interface to a []map[string]interface{} type.
func (p *Object) ToMapSliceGE() ([]map[string]interface{}, error) {
	if p == nil {
		return NewMapSliceV().G(), nil
	}
	m, err := ToMapSliceE(p.o)
	if err != nil {
		return nil, err
	}
	return m.G(), nil
}

// ToStringSlice converts an interface to a *StringSlice type.
func (p *Object) ToStringSlice() *StringSlice {
	if p == nil {
		return NewStringSliceV()
	}
	return ToStringSlice(p.o)
}

// ToStringSliceE converts an interface to a *StringSlice type.
func (p *Object) ToStringSliceE() (*StringSlice, error) {
	if p == nil {
		return NewStringSliceV(), nil
	}
	return ToStringSliceE(p.o)
}

// ToStringSliceG casts an interface to a []string type.
func (p *Object) ToStringSliceG() []string {
	return p.ToStringSlice().G()
}

// ToStringSliceGE casts an interface to a []string type.
func (p *Object) ToStringSliceGE() ([]string, error) {
	result, err := p.ToStringSliceE()
	if err != nil {
		return []string{}, err
	}
	return result.G(), nil
}

// ToIntSlice converts an interface to a *IntSlice type.
func (p *Object) ToIntSlice() *IntSlice {
	if p == nil {
		return NewIntSliceV()
	}
	v, _ := ToIntSliceE(p.o)
	return v
}

// ToIntSliceE converts an interface to a *IntSlice type.
func (p *Object) ToIntSliceE() (*IntSlice, error) {
	if p == nil {
		return NewIntSliceV(), nil
	}
	return ToIntSliceE(p.o)
}

// ToIntSliceG converts an interface to a []int type.
func (p *Object) ToIntSliceG() []int {
	result := []int{}
	if p == nil {
		return result
	}
	v, err := ToIntSliceE(p.o)
	if err != nil {
		return result
	}
	return v.G()
}

// ToIntSliceGE converts an interface to a []int type.
func (p *Object) ToIntSliceGE() ([]int, error) {
	result := []int{}
	if p == nil {
		return result, nil
	}
	v, err := ToIntSliceE(p.o)
	if err != nil {
		return result, err
	}
	return v.G(), err
}

// ToDurationSlice casts an interface to a []time.Duration type.
func (p *Object) ToDurationSlice() []time.Duration {
	if p == nil {
		return []time.Duration{}
	}
	v, _ := cast.ToDurationSliceE(p.o)
	return v
}

// ToDurationSliceE casts an interface to a []time.Duration type.
func (p *Object) ToDurationSliceE() ([]time.Duration, error) {
	if p == nil {
		return []time.Duration{}, nil
	}
	return cast.ToDurationSliceE(p.o)
}
