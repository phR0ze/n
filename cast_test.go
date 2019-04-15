package n

import (
	"errors"
	"fmt"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Indirect
//--------------------------------------------------------------------------------------------------
func ExampleIndirect() {
	slice := Indirect((*[]int)(nil))
	fmt.Println(slice)
	// Output: []
}

func TestIndirect(t *testing.T) {
	// nil
	{
		assert.Equal(t, nil, Indirect(nil))
	}

	// bool
	{
		var test bool
		assert.Equal(t, true, Indirect(true))
		assert.Equal(t, false, Indirect(&test))
		assert.Equal(t, false, Indirect((*bool)(nil)))
	}

	// []byte
	{
		assert.Equal(t, "test", string(Indirect([]byte{0x74, 0x65, 0x73, 0x74}).([]byte)))
		assert.Equal(t, "test", string(Indirect(&[]byte{0x74, 0x65, 0x73, 0x74}).([]byte)))
		assert.Equal(t, "", string(Indirect((*[]byte)(nil)).([]byte)))
	}

	// float32
	{
		var test float32
		assert.Equal(t, float32(7.22), Indirect(float32(7.22)))
		assert.Equal(t, float32(0), Indirect(&test))
		assert.Equal(t, float32(0), Indirect((*float32)(nil)))
	}

	// float64
	{
		var test float64
		assert.Equal(t, float64(7.22), Indirect(float64(7.22)))
		assert.Equal(t, float64(0), Indirect(&test))
		assert.Equal(t, float64(0), Indirect((*float64)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, 1, Indirect(1))
		assert.Equal(t, 0, Indirect(&test))
		assert.Equal(t, 0, Indirect((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, int8(3), Indirect(int8(3)))
		assert.Equal(t, int8(0), Indirect(&test))
		assert.Equal(t, int8(0), Indirect((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, int16(3), Indirect(int16(3)))
		assert.Equal(t, int16(0), Indirect(&test))
		assert.Equal(t, int16(0), Indirect((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, int32(3), Indirect(int32(3)))
		assert.Equal(t, int32(0), Indirect(&test))
		assert.Equal(t, int32(0), Indirect((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, int64(3), Indirect(int64(3)))
		assert.Equal(t, int64(0), Indirect(&test))
		assert.Equal(t, int64(0), Indirect((*int64)(nil)))
	}

	// []int
	{
		assert.Equal(t, []int{1, 2, 3}, Indirect([]int{1, 2, 3}))
		assert.Equal(t, []int{1, 2, 3}, Indirect(&[]int{1, 2, 3}))
		assert.Equal(t, []int{}, Indirect((*[]int)(nil)))
	}

	// []int8
	{
		assert.Equal(t, []int8{1, 2, 3}, Indirect([]int8{1, 2, 3}))
		assert.Equal(t, []int8{1, 2, 3}, Indirect(&[]int8{1, 2, 3}))
		assert.Equal(t, []int8{}, Indirect((*[]int8)(nil)))
	}

	// []int16
	{
		assert.Equal(t, []int16{1, 2, 3}, Indirect([]int16{1, 2, 3}))
		assert.Equal(t, []int16{1, 2, 3}, Indirect(&[]int16{1, 2, 3}))
		assert.Equal(t, []int16{}, Indirect((*[]int16)(nil)))
	}

	// []int32
	{
		assert.Equal(t, []int32{1, 2, 3}, Indirect([]int32{1, 2, 3}))
		assert.Equal(t, []int32{1, 2, 3}, Indirect(&[]int32{1, 2, 3}))
		assert.Equal(t, []int32{}, Indirect((*[]int32)(nil)))
	}

	// []int64
	{
		assert.Equal(t, []int64{1, 2, 3}, Indirect([]int64{1, 2, 3}))
		assert.Equal(t, []int64{1, 2, 3}, Indirect(&[]int64{1, 2, 3}))
		assert.Equal(t, []int64{}, Indirect((*[]int64)(nil)))
	}

	// rune
	{
		var test rune
		assert.Equal(t, 'r', Indirect('r'))
		assert.Equal(t, rune(0), Indirect(&test))
		assert.Equal(t, rune(0), Indirect((*rune)(nil)))
	}

	// []rune
	{
		assert.Equal(t, []rune{'t', 'e', 's', 't'}, Indirect([]rune{'t', 'e', 's', 't'}))
		assert.Equal(t, []rune{'t', 'e', 's', 't'}, Indirect(&[]rune{'t', 'e', 's', 't'}))
		assert.Equal(t, []rune{}, Indirect((*[]rune)(nil)))
	}

	// Str
	{
		assert.Equal(t, *A("test"), Indirect(*A("test")))
		assert.Equal(t, *A("test"), Indirect(A("test")))
		assert.Equal(t, *A(""), Indirect((*Str)(nil)))
	}

	// string
	{
		var test string
		assert.Equal(t, "test", Indirect("test"))
		assert.Equal(t, "", Indirect(&test))
		assert.Equal(t, "", Indirect((*string)(nil)))
	}

	// []string
	{
		assert.Equal(t, []string{"test"}, Indirect([]string{"test"}))
		assert.Equal(t, []string{"test"}, Indirect(&[]string{"test"}))
		assert.Equal(t, []string{}, Indirect((*[]string)(nil)))
	}

	// template.CSS
	{
		var test template.CSS
		assert.Equal(t, "test", string(Indirect(template.CSS("test")).(template.CSS)))
		assert.Equal(t, "", string(Indirect(&test).(template.CSS)))
		assert.Equal(t, "", string(Indirect((*template.CSS)(nil)).(template.CSS)))
	}

	// template.HTML
	{
		var test template.HTML
		assert.Equal(t, "test", string(Indirect(template.HTML("test")).(template.HTML)))
		assert.Equal(t, "", string(Indirect(&test).(template.HTML)))
		assert.Equal(t, "", string(Indirect((*template.HTML)(nil)).(template.HTML)))
	}

	// template.HTMLAttr
	{
		var test template.HTMLAttr
		assert.Equal(t, "test", string(Indirect(template.HTMLAttr("test")).(template.HTMLAttr)))
		assert.Equal(t, "", string(Indirect(&test).(template.HTMLAttr)))
		assert.Equal(t, "", string(Indirect((*template.HTMLAttr)(nil)).(template.HTMLAttr)))
	}

	// template.JS
	{
		var test template.JS
		assert.Equal(t, "test", string(Indirect(template.JS("test")).(template.JS)))
		assert.Equal(t, "", string(Indirect(&test).(template.JS)))
		assert.Equal(t, "", string(Indirect((*template.JS)(nil)).(template.JS)))
	}

	// template.URL
	{
		var test template.URL
		assert.Equal(t, "test", string(Indirect(template.URL("test")).(template.URL)))
		assert.Equal(t, "", string(Indirect(&test).(template.URL)))
		assert.Equal(t, "", string(Indirect((*template.URL)(nil)).(template.URL)))
	}

	// uint
	{
		var test uint
		assert.Equal(t, uint(1), Indirect(uint(1)))
		assert.Equal(t, uint(0), Indirect(&test))
		assert.Equal(t, uint(0), Indirect((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, uint8(3), Indirect(uint8(3)))
		assert.Equal(t, uint8(0), Indirect(&test))
		assert.Equal(t, uint8(0), Indirect((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, uint16(3), Indirect(uint16(3)))
		assert.Equal(t, uint16(0), Indirect(&test))
		assert.Equal(t, uint16(0), Indirect((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, uint32(3), Indirect(uint32(3)))
		assert.Equal(t, uint32(0), Indirect(&test))
		assert.Equal(t, uint32(0), Indirect((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, uint64(3), Indirect(uint64(3)))
		assert.Equal(t, uint64(0), Indirect(&test))
		assert.Equal(t, uint64(0), Indirect((*uint64)(nil)))
	}

	// []uint
	{
		assert.Equal(t, []uint{1, 2, 3}, Indirect([]uint{1, 2, 3}))
		assert.Equal(t, []uint{1, 2, 3}, Indirect(&[]uint{1, 2, 3}))
		assert.Equal(t, []uint{}, Indirect((*[]uint)(nil)))
	}

	// []uint8
	{
		assert.Equal(t, []uint8{1, 2, 3}, Indirect([]uint8{1, 2, 3}))
		assert.Equal(t, []uint8{1, 2, 3}, Indirect(&[]uint8{1, 2, 3}))
		assert.Equal(t, []uint8{}, Indirect((*[]uint8)(nil)))
	}

	// []uint16
	{
		assert.Equal(t, []uint16{1, 2, 3}, Indirect([]uint16{1, 2, 3}))
		assert.Equal(t, []uint16{1, 2, 3}, Indirect(&[]uint16{1, 2, 3}))
		assert.Equal(t, []uint16{}, Indirect((*[]uint16)(nil)))
	}

	// []uint32
	{
		assert.Equal(t, []uint32{1, 2, 3}, Indirect([]uint32{1, 2, 3}))
		assert.Equal(t, []uint32{1, 2, 3}, Indirect(&[]uint32{1, 2, 3}))
		assert.Equal(t, []uint32{}, Indirect((*[]uint32)(nil)))
	}

	// []uint64
	{
		assert.Equal(t, []uint64{1, 2, 3}, Indirect([]uint64{1, 2, 3}))
		assert.Equal(t, []uint64{1, 2, 3}, Indirect(&[]uint64{1, 2, 3}))
		assert.Equal(t, []uint64{}, Indirect((*[]uint64)(nil)))
	}
}

// Bool
//--------------------------------------------------------------------------------------------------
func ExampleToBool() {
	fmt.Println(ToBool(1))
	// Output: true
}

func TestToBool(t *testing.T) {

	// invalid
	{
		assert.Equal(t, false, ToBool(nil))
		assert.Equal(t, false, ToBool(&Object{}))
	}

	// bool
	{
		var test bool
		assert.Equal(t, true, ToBool(true))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*bool)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, true, ToBool(1))
		assert.Equal(t, false, ToBool(0))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, true, ToBool(int8(3)))
		assert.Equal(t, false, ToBool(int8(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, true, ToBool(int16(3)))
		assert.Equal(t, false, ToBool(int16(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, true, ToBool(int32(3)))
		assert.Equal(t, false, ToBool(int32(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, true, ToBool(int64(3)))
		assert.Equal(t, false, ToBool(int64(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*int64)(nil)))
	}

	// string
	{
		assert.Equal(t, true, ToBool("1"))
		assert.Equal(t, true, ToBool("true"))
		assert.Equal(t, true, ToBool("TRUE"))
		assert.Equal(t, false, ToBool("0"))
		assert.Equal(t, false, ToBool("false"))
		assert.Equal(t, false, ToBool("FALSE"))
		assert.Equal(t, false, ToBool(""))
	}

	// uint
	{
		var test uint
		assert.Equal(t, true, ToBool(uint(1)))
		assert.Equal(t, false, ToBool(uint(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, true, ToBool(uint8(3)))
		assert.Equal(t, false, ToBool(uint8(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, true, ToBool(uint16(3)))
		assert.Equal(t, false, ToBool(uint16(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, true, ToBool(uint32(3)))
		assert.Equal(t, false, ToBool(uint32(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, true, ToBool(uint64(3)))
		assert.Equal(t, false, ToBool(uint64(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*uint64)(nil)))
	}

}

// BoolE
//--------------------------------------------------------------------------------------------------
func ExampleToBoolE() {
	fmt.Println(ToBoolE(1))
	// Output: true <nil>
}

func TestToBoolE(t *testing.T) {

	// invalid
	{
		val, err := ToBoolE(nil)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&Object{})
		assert.Equal(t, "unable to cast type *n.Object to bool", err.Error())
		assert.Equal(t, false, val)
	}

	// bool
	{
		var test bool
		val, err := ToBoolE(true)
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// int
	{
		var test int
		val, err := ToBoolE(1)
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(0)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// int8
	{
		var test int8
		val, err := ToBoolE(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// int16
	{
		var test int16
		val, err := ToBoolE(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// int32
	{
		var test int32
		val, err := ToBoolE(int32(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(int32(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// int64
	{
		var test int64
		val, err := ToBoolE(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// string
	{
		assert.Equal(t, true, ToBool("1"))
		assert.Equal(t, true, ToBool("true"))
		assert.Equal(t, true, ToBool("TRUE"))
		assert.Equal(t, false, ToBool("0"))
		assert.Equal(t, false, ToBool("false"))
		assert.Equal(t, false, ToBool("FALSE"))
		assert.Equal(t, false, ToBool(""))
	}

	// uint
	{
		var test uint
		val, err := ToBoolE(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(0)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// uint8
	{
		var test uint8
		val, err := ToBoolE(uint8(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(uint8(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// uint16
	{
		var test uint16
		val, err := ToBoolE(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// uint32
	{
		var test uint32
		val, err := ToBoolE(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}

	// uint64
	{
		var test uint64
		val, err := ToBoolE(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&test)
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, false, val)
	}
}

// String
//--------------------------------------------------------------------------------------------------
func ExampleToString() {
	val := ToString(true)
	fmt.Println(val)
	// Output: true
}

func TestToString(t *testing.T) {

	// bool
	{
		var test bool
		assert.Equal(t, "true", ToString(true))
		assert.Equal(t, "false", ToString(false))
		assert.Equal(t, "false", ToString(&test))
		assert.Equal(t, "false", ToString((*bool)(nil)))
	}

	// []byte
	{
		assert.Equal(t, "test", ToString([]byte{0x74, 0x65, 0x73, 0x74}))
		assert.Equal(t, "test", ToString(&[]byte{0x74, 0x65, 0x73, 0x74}))
		assert.Equal(t, "", ToString((*[]byte)(nil)))
	}

	// error
	{
		assert.Equal(t, "test", ToString(errors.New("test")))
	}

	// float32
	{
		var test float32
		assert.Equal(t, "7.22", ToString(float32(7.22)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*float32)(nil)))
	}

	// float64
	{
		var test float64
		assert.Equal(t, "7.22", ToString(float64(7.22)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*float64)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, "3", ToString(3))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, "3", ToString(int8(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, "3", ToString(int16(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, "3", ToString(int32(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, "3", ToString(int64(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*int64)(nil)))
	}

	// ints
	{
		assert.Equal(t, "7", ToString(int(7)))
		assert.Equal(t, "7", ToString(int8(7)))
		assert.Equal(t, "7", ToString(int16(7)))
		assert.Equal(t, "7", ToString(int32(7)))
		assert.Equal(t, "7", ToString(int64(7)))
	}

	// nil
	{
		assert.Equal(t, "", ToString(nil))
	}

	// template.CSS
	{
		var test template.CSS
		assert.Equal(t, "test", ToString(template.CSS("test")))
		assert.Equal(t, "", ToString(&test))
		assert.Equal(t, "", ToString((*template.CSS)(nil)))
	}

	// template.HTML
	{
		var test template.HTML
		assert.Equal(t, "test", ToString(template.HTML("test")))
		assert.Equal(t, "", ToString(&test))
		assert.Equal(t, "", ToString((*template.HTML)(nil)))
	}

	// template.HTMLAttr
	{
		var test template.HTMLAttr
		assert.Equal(t, "test", ToString(template.HTMLAttr("test")))
		assert.Equal(t, "", ToString(&test))
		assert.Equal(t, "", ToString((*template.HTMLAttr)(nil)))
	}

	// template.JS
	{
		var test template.JS
		assert.Equal(t, "test", ToString(template.JS("test")))
		assert.Equal(t, "", ToString(&test))
		assert.Equal(t, "", ToString((*template.JS)(nil)))
	}

	// template.URL
	{
		var test template.URL
		assert.Equal(t, "test", ToString(template.URL("test")))
		assert.Equal(t, "", ToString(&test))
		assert.Equal(t, "", ToString((*template.URL)(nil)))
	}

	// uint
	{
		var test uint
		assert.Equal(t, "3", ToString(uint(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, "3", ToString(uint8(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, "3", ToString(uint16(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, "3", ToString(uint32(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, "3", ToString(uint64(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "0", ToString((*uint64)(nil)))
	}

	// uints
	{
		assert.Equal(t, "7", ToString(uint(7)))
		assert.Equal(t, "7", ToString(uint8(7)))
		assert.Equal(t, "7", ToString(uint16(7)))
		assert.Equal(t, "7", ToString(uint32(7)))
		assert.Equal(t, "7", ToString(uint64(7)))
	}
}
