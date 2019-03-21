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
