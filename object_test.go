package n

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestObject_ToBool(t *testing.T) {

	// w/out error
	{
		n := &Object{true}
		assert.IsType(t, true, n.ToBool())
	}

	// w/error
	{
		n := &Object{true}
		b, e := n.ToBoolE()
		assert.Nil(t, e)
		assert.IsType(t, true, b)
	}
}

func TestObject_ToTime(t *testing.T) {

	// w/out error
	{
		n := &Object{time.Time{}}
		assert.IsType(t, time.Time{}, n.ToTime())
	}

	// w/error
	{
		n := &Object{time.Time{}}
		obj, e := n.ToTimeE()
		assert.Nil(t, e)
		assert.IsType(t, time.Time{}, obj)
	}
}

func TestObject_ToDuration(t *testing.T) {

	// w/out error
	{
		n := &Object{time.Duration(0)}
		assert.IsType(t, time.Duration(0), n.ToDuration())
	}

	// w/error
	{
		n := &Object{time.Duration(0)}
		obj, e := n.ToDurationE()
		assert.Nil(t, e)
		assert.IsType(t, time.Duration(0), obj)
	}
}

func TestObject_ToFloat32(t *testing.T) {

	// w/out error
	{
		n := &Object{float32(1.0)}
		assert.IsType(t, float32(1.0), n.ToFloat32())
	}

	// w/error
	{
		n := &Object{float32(1.0)}
		obj, e := n.ToFloat32E()
		assert.Nil(t, e)
		assert.IsType(t, float32(1.0), obj)
	}
}

func TestObject_ToFloat64(t *testing.T) {

	// w/out error
	{
		n := &Object{float64(1.0)}
		assert.IsType(t, float64(1.0), n.ToFloat64())
	}

	// w/error
	{
		n := &Object{float64(1.0)}
		obj, e := n.ToFloat64E()
		assert.Nil(t, e)
		assert.IsType(t, float64(1.0), obj)
	}
}

func TestObject_ToInt(t *testing.T) {

	// w/out error
	{
		n := &Object{1}
		assert.IsType(t, 1, n.ToInt())
	}

	// w/error
	{
		n := &Object{1}
		obj, e := n.ToIntE()
		assert.Nil(t, e)
		assert.IsType(t, 1, obj)
	}
}

func TestObject_ToInt8(t *testing.T) {

	// w/out error
	{
		n := &Object{int8(1)}
		assert.IsType(t, int8(1), n.ToInt8())
	}

	// w/error
	{
		n := &Object{int8(1)}
		obj, e := n.ToInt8E()
		assert.Nil(t, e)
		assert.IsType(t, int8(1), obj)
	}
}

func TestObject_ToInt16(t *testing.T) {

	// w/out error
	{
		n := &Object{int16(1)}
		assert.IsType(t, int16(1), n.ToInt16())
	}

	// w/error
	{
		n := &Object{int16(1)}
		obj, e := n.ToInt16E()
		assert.Nil(t, e)
		assert.IsType(t, int16(1), obj)
	}
}

func TestObject_ToInt32(t *testing.T) {

	// w/out error
	{
		n := &Object{int32(1)}
		assert.IsType(t, int32(1), n.ToInt32())
	}

	// w/error
	{
		n := &Object{int32(1)}
		obj, e := n.ToInt32E()
		assert.Nil(t, e)
		assert.IsType(t, int32(1), obj)
	}
}

func TestObject_ToInt64(t *testing.T) {

	// w/out error
	{
		n := &Object{int64(1)}
		assert.IsType(t, int64(1), n.ToInt64())
	}

	// w/error
	{
		n := &Object{int64(1)}
		obj, e := n.ToInt64E()
		assert.Nil(t, e)
		assert.IsType(t, int64(1), obj)
	}
}

func TestObject_ToUInt(t *testing.T) {

	// w/out error
	{
		n := &Object{uint(1)}
		assert.IsType(t, uint(1), n.ToUint())
	}

	// w/error
	{
		n := &Object{uint(1)}
		obj, e := n.ToUintE()
		assert.Nil(t, e)
		assert.IsType(t, uint(1), obj)
	}
}

func TestObject_ToUint8(t *testing.T) {

	// w/out error
	{
		n := &Object{uint8(1)}
		assert.IsType(t, uint8(1), n.ToUint8())
	}

	// w/error
	{
		n := &Object{uint8(1)}
		obj, e := n.ToUint8E()
		assert.Nil(t, e)
		assert.IsType(t, uint8(1), obj)
	}
}

func TestObject_ToUint16(t *testing.T) {

	// w/out error
	{
		n := &Object{uint16(1)}
		assert.IsType(t, uint16(1), n.ToUint16())
	}

	// w/error
	{
		n := &Object{uint16(1)}
		obj, e := n.ToUint16E()
		assert.Nil(t, e)
		assert.IsType(t, uint16(1), obj)
	}
}

func TestObject_ToUint32(t *testing.T) {

	// w/out error
	{
		n := &Object{uint32(1)}
		assert.IsType(t, uint32(1), n.ToUint32())
	}

	// w/error
	{
		n := &Object{uint32(1)}
		obj, e := n.ToUint32E()
		assert.Nil(t, e)
		assert.IsType(t, uint32(1), obj)
	}
}

func TestObject_ToUint64(t *testing.T) {

	// w/out error
	{
		n := &Object{uint64(1)}
		assert.IsType(t, uint64(1), n.ToUint64())
	}

	// w/error
	{
		n := &Object{uint64(1)}
		obj, e := n.ToUint64E()
		assert.Nil(t, e)
		assert.IsType(t, uint64(1), obj)
	}
}

func TestObject_ToString(t *testing.T) {

	// w/out error
	{
		n := &Object{""}
		assert.IsType(t, "", n.ToString())
	}

	// w/error
	{
		n := &Object{""}
		obj, e := n.ToStringE()
		assert.Nil(t, e)
		assert.IsType(t, "", obj)
	}
}

func TestObject_ToStringMapString(t *testing.T) {

	// w/out error
	{
		n := &Object{map[string]string{}}
		assert.IsType(t, map[string]string{}, n.ToStringMapString())
	}

	// w/error
	{
		n := &Object{map[string]string{}}
		obj, e := n.ToStringMapStringE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]string{}, obj)
	}
}

func TestObject_ToStringMapStringSlice(t *testing.T) {

	// w/out error
	{
		n := &Object{map[string][]string{}}
		assert.IsType(t, map[string][]string{}, n.ToStringMapStringSlice())
	}

	// w/error
	{
		n := &Object{map[string][]string{}}
		obj, e := n.ToStringMapStringSliceE()
		assert.Nil(t, e)
		assert.IsType(t, map[string][]string{}, obj)
	}
}

func TestObject_ToStringMapBool(t *testing.T) {

	// w/out error
	{
		n := &Object{map[string]bool{}}
		assert.IsType(t, map[string]bool{}, n.ToStringMapBool())
	}

	// w/error
	{
		n := &Object{map[string]bool{}}
		obj, e := n.ToStringMapBoolE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]bool{}, obj)
	}
}

func TestObject_ToStringMapInt(t *testing.T) {

	// w/out error
	{
		n := &Object{map[string]int{}}
		assert.IsType(t, map[string]int{}, n.ToStringMapInt())
	}

	// w/error
	{
		n := &Object{map[string]int{}}
		obj, e := n.ToStringMapIntE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]int{}, obj)
	}
}

func TestObject_ToStringMapInt64(t *testing.T) {

	// w/out error
	{
		n := &Object{map[string]int64{}}
		assert.IsType(t, map[string]int64{}, n.ToStringMapInt64())
	}

	// w/error
	{
		n := &Object{map[string]int64{}}
		obj, e := n.ToStringMapInt64E()
		assert.Nil(t, e)
		assert.IsType(t, map[string]int64{}, obj)
	}
}

func TestObject_ToStringMap(t *testing.T) {

	// w/out error
	{
		n := &Object{map[string]interface{}{}}
		assert.IsType(t, map[string]interface{}{}, n.ToStringMap())
	}

	// w/error
	{
		n := &Object{map[string]interface{}{}}
		obj, e := n.ToStringMapE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]interface{}{}, obj)
	}
}

func TestObject_ToSlice(t *testing.T) {

	// w/out error
	{
		n := &Object{[]interface{}{}}
		assert.IsType(t, []interface{}{}, n.ToSlice())
	}

	// w/error
	{
		n := &Object{[]interface{}{}}
		obj, e := n.ToSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []interface{}{}, obj)
	}
}

func TestObject_ToBoolSlice(t *testing.T) {

	// w/out error
	{
		n := &Object{[]bool{}}
		assert.IsType(t, []bool{}, n.ToBoolSlice())
	}

	// w/error
	{
		n := &Object{[]bool{}}
		obj, e := n.ToBoolSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []bool{}, obj)
	}
}

func TestObject_ToStringSlice(t *testing.T) {

	// w/out error
	{
		n := &Object{[]string{}}
		assert.IsType(t, []string{}, n.ToStringSlice())
	}

	// w/error
	{
		n := &Object{[]string{}}
		obj, e := n.ToStringSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []string{}, obj)
	}
}

func TestObject_ToIntSlice(t *testing.T) {

	// w/out error
	{
		n := &Object{[]int{}}
		assert.IsType(t, []int{}, n.ToIntSlice())
	}

	// w/error
	{
		n := &Object{[]int{}}
		obj, e := n.ToIntSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []int{}, obj)
	}
}

func TestObject_ToDurationSlice(t *testing.T) {

	// w/out error
	{
		n := &Object{[]time.Duration{}}
		assert.IsType(t, []time.Duration{}, n.ToDurationSlice())
	}

	// w/error
	{
		n := &Object{[]time.Duration{}}
		obj, e := n.ToDurationSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []time.Duration{}, obj)
	}
}
