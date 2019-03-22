package n

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNObj_ToBool(t *testing.T) {

	// w/out error
	{
		n := &NObj{true}
		assert.IsType(t, true, n.ToBool())
	}

	// w/error
	{
		n := &NObj{true}
		b, e := n.ToBoolE()
		assert.Nil(t, e)
		assert.IsType(t, true, b)
	}
}

func TestNObj_ToTime(t *testing.T) {

	// w/out error
	{
		n := &NObj{time.Time{}}
		assert.IsType(t, time.Time{}, n.ToTime())
	}

	// w/error
	{
		n := &NObj{time.Time{}}
		obj, e := n.ToTimeE()
		assert.Nil(t, e)
		assert.IsType(t, time.Time{}, obj)
	}
}

func TestNObj_ToDuration(t *testing.T) {

	// w/out error
	{
		n := &NObj{time.Duration(0)}
		assert.IsType(t, time.Duration(0), n.ToDuration())
	}

	// w/error
	{
		n := &NObj{time.Duration(0)}
		obj, e := n.ToDurationE()
		assert.Nil(t, e)
		assert.IsType(t, time.Duration(0), obj)
	}
}

func TestNObj_ToFloat32(t *testing.T) {

	// w/out error
	{
		n := &NObj{float32(1.0)}
		assert.IsType(t, float32(1.0), n.ToFloat32())
	}

	// w/error
	{
		n := &NObj{float32(1.0)}
		obj, e := n.ToFloat32E()
		assert.Nil(t, e)
		assert.IsType(t, float32(1.0), obj)
	}
}

func TestNObj_ToFloat64(t *testing.T) {

	// w/out error
	{
		n := &NObj{float64(1.0)}
		assert.IsType(t, float64(1.0), n.ToFloat64())
	}

	// w/error
	{
		n := &NObj{float64(1.0)}
		obj, e := n.ToFloat64E()
		assert.Nil(t, e)
		assert.IsType(t, float64(1.0), obj)
	}
}

func TestNObj_ToInt(t *testing.T) {

	// w/out error
	{
		n := &NObj{1}
		assert.IsType(t, 1, n.ToInt())
	}

	// w/error
	{
		n := &NObj{1}
		obj, e := n.ToIntE()
		assert.Nil(t, e)
		assert.IsType(t, 1, obj)
	}
}

func TestNObj_ToInt8(t *testing.T) {

	// w/out error
	{
		n := &NObj{int8(1)}
		assert.IsType(t, int8(1), n.ToInt8())
	}

	// w/error
	{
		n := &NObj{int8(1)}
		obj, e := n.ToInt8E()
		assert.Nil(t, e)
		assert.IsType(t, int8(1), obj)
	}
}

func TestNObj_ToInt16(t *testing.T) {

	// w/out error
	{
		n := &NObj{int16(1)}
		assert.IsType(t, int16(1), n.ToInt16())
	}

	// w/error
	{
		n := &NObj{int16(1)}
		obj, e := n.ToInt16E()
		assert.Nil(t, e)
		assert.IsType(t, int16(1), obj)
	}
}

func TestNObj_ToInt32(t *testing.T) {

	// w/out error
	{
		n := &NObj{int32(1)}
		assert.IsType(t, int32(1), n.ToInt32())
	}

	// w/error
	{
		n := &NObj{int32(1)}
		obj, e := n.ToInt32E()
		assert.Nil(t, e)
		assert.IsType(t, int32(1), obj)
	}
}

func TestNObj_ToInt64(t *testing.T) {

	// w/out error
	{
		n := &NObj{int64(1)}
		assert.IsType(t, int64(1), n.ToInt64())
	}

	// w/error
	{
		n := &NObj{int64(1)}
		obj, e := n.ToInt64E()
		assert.Nil(t, e)
		assert.IsType(t, int64(1), obj)
	}
}

func TestNObj_ToUInt(t *testing.T) {

	// w/out error
	{
		n := &NObj{uint(1)}
		assert.IsType(t, uint(1), n.ToUint())
	}

	// w/error
	{
		n := &NObj{uint(1)}
		obj, e := n.ToUintE()
		assert.Nil(t, e)
		assert.IsType(t, uint(1), obj)
	}
}

func TestNObj_ToUint8(t *testing.T) {

	// w/out error
	{
		n := &NObj{uint8(1)}
		assert.IsType(t, uint8(1), n.ToUint8())
	}

	// w/error
	{
		n := &NObj{uint8(1)}
		obj, e := n.ToUint8E()
		assert.Nil(t, e)
		assert.IsType(t, uint8(1), obj)
	}
}

func TestNObj_ToUint16(t *testing.T) {

	// w/out error
	{
		n := &NObj{uint16(1)}
		assert.IsType(t, uint16(1), n.ToUint16())
	}

	// w/error
	{
		n := &NObj{uint16(1)}
		obj, e := n.ToUint16E()
		assert.Nil(t, e)
		assert.IsType(t, uint16(1), obj)
	}
}

func TestNObj_ToUint32(t *testing.T) {

	// w/out error
	{
		n := &NObj{uint32(1)}
		assert.IsType(t, uint32(1), n.ToUint32())
	}

	// w/error
	{
		n := &NObj{uint32(1)}
		obj, e := n.ToUint32E()
		assert.Nil(t, e)
		assert.IsType(t, uint32(1), obj)
	}
}

func TestNObj_ToUint64(t *testing.T) {

	// w/out error
	{
		n := &NObj{uint64(1)}
		assert.IsType(t, uint64(1), n.ToUint64())
	}

	// w/error
	{
		n := &NObj{uint64(1)}
		obj, e := n.ToUint64E()
		assert.Nil(t, e)
		assert.IsType(t, uint64(1), obj)
	}
}

func TestNObj_ToString(t *testing.T) {

	// w/out error
	{
		n := &NObj{""}
		assert.IsType(t, "", n.ToString())
	}

	// w/error
	{
		n := &NObj{""}
		obj, e := n.ToStringE()
		assert.Nil(t, e)
		assert.IsType(t, "", obj)
	}
}

func TestNObj_ToStringMapString(t *testing.T) {

	// w/out error
	{
		n := &NObj{map[string]string{}}
		assert.IsType(t, map[string]string{}, n.ToStringMapString())
	}

	// w/error
	{
		n := &NObj{map[string]string{}}
		obj, e := n.ToStringMapStringE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]string{}, obj)
	}
}

func TestNObj_ToStringMapStringSlice(t *testing.T) {

	// w/out error
	{
		n := &NObj{map[string][]string{}}
		assert.IsType(t, map[string][]string{}, n.ToStringMapStringSlice())
	}

	// w/error
	{
		n := &NObj{map[string][]string{}}
		obj, e := n.ToStringMapStringSliceE()
		assert.Nil(t, e)
		assert.IsType(t, map[string][]string{}, obj)
	}
}

func TestNObj_ToStringMapBool(t *testing.T) {

	// w/out error
	{
		n := &NObj{map[string]bool{}}
		assert.IsType(t, map[string]bool{}, n.ToStringMapBool())
	}

	// w/error
	{
		n := &NObj{map[string]bool{}}
		obj, e := n.ToStringMapBoolE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]bool{}, obj)
	}
}

func TestNObj_ToStringMapInt(t *testing.T) {

	// w/out error
	{
		n := &NObj{map[string]int{}}
		assert.IsType(t, map[string]int{}, n.ToStringMapInt())
	}

	// w/error
	{
		n := &NObj{map[string]int{}}
		obj, e := n.ToStringMapIntE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]int{}, obj)
	}
}

func TestNObj_ToStringMapInt64(t *testing.T) {

	// w/out error
	{
		n := &NObj{map[string]int64{}}
		assert.IsType(t, map[string]int64{}, n.ToStringMapInt64())
	}

	// w/error
	{
		n := &NObj{map[string]int64{}}
		obj, e := n.ToStringMapInt64E()
		assert.Nil(t, e)
		assert.IsType(t, map[string]int64{}, obj)
	}
}

func TestNObj_ToStringMap(t *testing.T) {

	// w/out error
	{
		n := &NObj{map[string]interface{}{}}
		assert.IsType(t, map[string]interface{}{}, n.ToStringMap())
	}

	// w/error
	{
		n := &NObj{map[string]interface{}{}}
		obj, e := n.ToStringMapE()
		assert.Nil(t, e)
		assert.IsType(t, map[string]interface{}{}, obj)
	}
}

func TestNObj_ToSlice(t *testing.T) {

	// w/out error
	{
		n := &NObj{[]interface{}{}}
		assert.IsType(t, []interface{}{}, n.ToSlice())
	}

	// w/error
	{
		n := &NObj{[]interface{}{}}
		obj, e := n.ToSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []interface{}{}, obj)
	}
}

func TestNObj_ToBoolSlice(t *testing.T) {

	// w/out error
	{
		n := &NObj{[]bool{}}
		assert.IsType(t, []bool{}, n.ToBoolSlice())
	}

	// w/error
	{
		n := &NObj{[]bool{}}
		obj, e := n.ToBoolSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []bool{}, obj)
	}
}

func TestNObj_ToStringSlice(t *testing.T) {

	// w/out error
	{
		n := &NObj{[]string{}}
		assert.IsType(t, []string{}, n.ToStringSlice())
	}

	// w/error
	{
		n := &NObj{[]string{}}
		obj, e := n.ToStringSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []string{}, obj)
	}
}

func TestNObj_ToIntSlice(t *testing.T) {

	// w/out error
	{
		n := &NObj{[]int{}}
		assert.IsType(t, []int{}, n.ToIntSlice())
	}

	// w/error
	{
		n := &NObj{[]int{}}
		obj, e := n.ToIntSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []int{}, obj)
	}
}

func TestNObj_ToDurationSlice(t *testing.T) {

	// w/out error
	{
		n := &NObj{[]time.Duration{}}
		assert.IsType(t, []time.Duration{}, n.ToDurationSlice())
	}

	// w/error
	{
		n := &NObj{[]time.Duration{}}
		obj, e := n.ToDurationSliceE()
		assert.Nil(t, e)
		assert.IsType(t, []time.Duration{}, obj)
	}
}
