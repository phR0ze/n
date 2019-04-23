package n

import (
	"errors"
	"fmt"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestObj struct {
	o interface{}
}

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
	//-----------------------------------------
	{
		var test bool
		assert.Equal(t, true, Indirect(true))
		assert.Equal(t, false, Indirect(&test))
		assert.Equal(t, false, Indirect((*bool)(nil)))

		// []bool
		assert.Equal(t, []bool{true}, Indirect([]bool{true}))
		assert.Equal(t, []bool{true}, Indirect(&[]bool{true}))
		assert.Equal(t, []*bool{&test}, Indirect(&[]*bool{&test}))
	}

	// byte
	//-----------------------------------------
	{
		var test byte
		assert.Equal(t, byte(3), Indirect(byte(3)))
		assert.Equal(t, byte(0), Indirect(&test))
		assert.Equal(t, byte(0), Indirect((*byte)(nil)))

		// []byte
		assert.Equal(t, "test", string(Indirect([]byte{0x74, 0x65, 0x73, 0x74}).([]byte)))
		assert.Equal(t, "test", string(Indirect(&[]byte{0x74, 0x65, 0x73, 0x74}).([]byte)))
		assert.Equal(t, "", string(Indirect((*[]byte)(nil)).([]byte)))
		assert.Equal(t, []*byte{&test}, Indirect(&[]*byte{&test}))
	}

	// float32
	//-----------------------------------------
	{
		var test float32
		assert.Equal(t, float32(7.22), Indirect(float32(7.22)))
		assert.Equal(t, float32(0), Indirect(&test))
		assert.Equal(t, float32(0), Indirect((*float32)(nil)))

		// []float32
		assert.Equal(t, []float32{1.2}, Indirect([]float32{1.2}))
		assert.Equal(t, []float32{1.2}, Indirect(&[]float32{1.2}))
		assert.Equal(t, []*float32{&test}, Indirect(&[]*float32{&test}))
	}

	// float64
	//-----------------------------------------
	{
		var test float64
		assert.Equal(t, float64(7.22), Indirect(float64(7.22)))
		assert.Equal(t, float64(0), Indirect(&test))
		assert.Equal(t, float64(0), Indirect((*float64)(nil)))

		// []float64
		assert.Equal(t, []float64{1.2}, Indirect([]float64{1.2}))
		assert.Equal(t, []float64{1.2}, Indirect(&[]float64{1.2}))
		assert.Equal(t, []*float64{&test}, Indirect(&[]*float64{&test}))
	}

	// int
	//-----------------------------------------
	{
		var test int
		assert.Equal(t, 1, Indirect(1))
		assert.Equal(t, 0, Indirect(&test))
		assert.Equal(t, 0, Indirect((*int)(nil)))

		// []int
		assert.Equal(t, []int{1, 2, 3}, Indirect([]int{1, 2, 3}))
		assert.Equal(t, []int{1, 2, 3}, Indirect(&[]int{1, 2, 3}))
		assert.Equal(t, []int{}, Indirect((*[]int)(nil)))
		assert.Equal(t, []*int{}, Indirect((*[]*int)(nil)))
		assert.Equal(t, []*int{&test}, Indirect(&[]*int{&test}))
	}

	// interface
	//-----------------------------------------
	{
		assert.Equal(t, []interface{}{1, 2, 3}, Indirect([]interface{}{1, 2, 3}))
		assert.Equal(t, []interface{}{1, 2, 3}, Indirect(&[]interface{}{1, 2, 3}))
		assert.Equal(t, []interface{}{}, Indirect((*[]interface{})(nil)))
		assert.Equal(t, []*interface{}{}, Indirect((*[]*interface{})(nil)))
	}

	// int8
	//-----------------------------------------
	{
		var test int8
		assert.Equal(t, int8(3), Indirect(int8(3)))
		assert.Equal(t, int8(0), Indirect(&test))
		assert.Equal(t, int8(0), Indirect((*int8)(nil)))

		// []int8
		assert.Equal(t, []int8{1, 2, 3}, Indirect([]int8{1, 2, 3}))
		assert.Equal(t, []int8{1, 2, 3}, Indirect(&[]int8{1, 2, 3}))
		assert.Equal(t, []int8{}, Indirect((*[]int8)(nil)))
		assert.Equal(t, []*int8{}, Indirect((*[]*int8)(nil)))
		assert.Equal(t, []*int8{&test}, Indirect(&[]*int8{&test}))
	}

	// int16
	//-----------------------------------------
	{
		var test int16
		assert.Equal(t, int16(3), Indirect(int16(3)))
		assert.Equal(t, int16(0), Indirect(&test))
		assert.Equal(t, int16(0), Indirect((*int16)(nil)))

		// []int16
		assert.Equal(t, []int16{1, 2, 3}, Indirect([]int16{1, 2, 3}))
		assert.Equal(t, []int16{1, 2, 3}, Indirect(&[]int16{1, 2, 3}))
		assert.Equal(t, []int16{}, Indirect((*[]int16)(nil)))
		assert.Equal(t, []*int16{}, Indirect((*[]*int16)(nil)))
		assert.Equal(t, []*int16{&test}, Indirect(&[]*int16{&test}))
	}

	// int32
	//-----------------------------------------
	{
		var test int32
		assert.Equal(t, int32(3), Indirect(int32(3)))
		assert.Equal(t, int32(0), Indirect(&test))
		assert.Equal(t, int32(0), Indirect((*int32)(nil)))

		// []int32
		assert.Equal(t, []int32{1, 2, 3}, Indirect([]int32{1, 2, 3}))
		assert.Equal(t, []int32{1, 2, 3}, Indirect(&[]int32{1, 2, 3}))
		assert.Equal(t, []int32{}, Indirect((*[]int32)(nil)))
		assert.Equal(t, []*int32{}, Indirect((*[]*int32)(nil)))
		assert.Equal(t, []*int32{&test}, Indirect(&[]*int32{&test}))
	}

	// int64
	//-----------------------------------------
	{
		var test int64
		assert.Equal(t, int64(3), Indirect(int64(3)))
		assert.Equal(t, int64(0), Indirect(&test))
		assert.Equal(t, int64(0), Indirect((*int64)(nil)))

		// []int64
		assert.Equal(t, []int64{1, 2, 3}, Indirect([]int64{1, 2, 3}))
		assert.Equal(t, []int64{1, 2, 3}, Indirect(&[]int64{1, 2, 3}))
		assert.Equal(t, []int64{}, Indirect((*[]int64)(nil)))
		assert.Equal(t, []*int64{}, Indirect((*[]*int64)(nil)))
		assert.Equal(t, []*int64{&test}, Indirect(&[]*int64{&test}))
	}

	// Object
	//-----------------------------------------
	{
		assert.Equal(t, Object{3}, Indirect(Object{3}))
		assert.Equal(t, Object{3}, Indirect(&Object{3}))
		assert.Equal(t, Object{}, Indirect((*Object)(nil)))

		// []Object
		assert.Equal(t, []Object{{1}, {2}, {3}}, Indirect([]Object{{1}, {2}, {3}}))
		assert.Equal(t, []Object{{1}, {2}, {3}}, Indirect(&[]Object{{1}, {2}, {3}}))
		assert.Equal(t, []Object{}, Indirect((*[]Object)(nil)))
		assert.Equal(t, []*Object{}, Indirect((*[]*Object)(nil)))
		assert.Equal(t, []*Object{&Object{}}, Indirect(&[]*Object{&Object{}}))
	}

	// rune
	//-----------------------------------------
	{
		var test rune
		assert.Equal(t, 'r', Indirect('r'))
		assert.Equal(t, rune(0), Indirect(&test))
		assert.Equal(t, rune(0), Indirect((*rune)(nil)))

		// []rune
		assert.Equal(t, []rune{'t', 'e', 's', 't'}, Indirect([]rune{'t', 'e', 's', 't'}))
		assert.Equal(t, []rune{'t', 'e', 's', 't'}, Indirect(&[]rune{'t', 'e', 's', 't'}))
		assert.Equal(t, []rune{}, Indirect((*[]rune)(nil)))
		assert.Equal(t, []*rune{}, Indirect((*[]*rune)(nil)))
		assert.Equal(t, []*rune{&test}, Indirect(&[]*rune{&test}))
	}

	// Str
	//-----------------------------------------
	{
		assert.Equal(t, *A("test"), Indirect(*A("test")))
		assert.Equal(t, *A("test"), Indirect(A("test")))
		assert.Equal(t, *A(""), Indirect((*Str)(nil)))

		// []Str
		assert.Equal(t, []Str{{1}, {2}, {3}}, Indirect([]Str{{1}, {2}, {3}}))
		assert.Equal(t, []Str{{1}, {2}, {3}}, Indirect(&[]Str{{1}, {2}, {3}}))
		assert.Equal(t, []Str{}, Indirect((*[]Str)(nil)))
		assert.Equal(t, []*Str{}, Indirect((*[]*Str)(nil)))
		assert.Equal(t, []*Str{&Str{}}, Indirect(&[]*Str{&Str{}}))
	}

	// string
	//-----------------------------------------
	{
		var test string
		assert.Equal(t, "test", Indirect("test"))
		assert.Equal(t, "", Indirect(&test))
		assert.Equal(t, "", Indirect((*string)(nil)))

		// []string
		assert.Equal(t, []string{"test"}, Indirect([]string{"test"}))
		assert.Equal(t, []string{"test"}, Indirect(&[]string{"test"}))
		assert.Equal(t, []string{}, Indirect((*[]string)(nil)))
		assert.Equal(t, []*string{}, Indirect((*[]*string)(nil)))
	}

	// template.CSS
	//-----------------------------------------
	{
		var test template.CSS
		assert.Equal(t, "test", string(Indirect(template.CSS("test")).(template.CSS)))
		assert.Equal(t, "", string(Indirect(&test).(template.CSS)))
		assert.Equal(t, "", string(Indirect((*template.CSS)(nil)).(template.CSS)))

		// []template.CSS
		assert.Equal(t, []template.CSS{"test"}, Indirect([]template.CSS{"test"}))
		assert.Equal(t, []template.CSS{"test"}, Indirect(&[]template.CSS{"test"}))
		assert.Equal(t, []template.CSS{}, Indirect((*[]template.CSS)(nil)))
		assert.Equal(t, []*template.CSS{}, Indirect((*[]*template.CSS)(nil)))
	}

	// template.HTML
	//-----------------------------------------
	{
		var test template.HTML
		assert.Equal(t, "test", string(Indirect(template.HTML("test")).(template.HTML)))
		assert.Equal(t, "", string(Indirect(&test).(template.HTML)))
		assert.Equal(t, "", string(Indirect((*template.HTML)(nil)).(template.HTML)))

		// []template.HTML
		assert.Equal(t, []template.HTML{"test"}, Indirect([]template.HTML{"test"}))
		assert.Equal(t, []template.HTML{"test"}, Indirect(&[]template.HTML{"test"}))
		assert.Equal(t, []template.HTML{}, Indirect((*[]template.HTML)(nil)))
		assert.Equal(t, []*template.HTML{}, Indirect((*[]*template.HTML)(nil)))
	}

	// template.HTMLAttr
	//-----------------------------------------
	{
		var test template.HTMLAttr
		assert.Equal(t, "test", string(Indirect(template.HTMLAttr("test")).(template.HTMLAttr)))
		assert.Equal(t, "", string(Indirect(&test).(template.HTMLAttr)))
		assert.Equal(t, "", string(Indirect((*template.HTMLAttr)(nil)).(template.HTMLAttr)))

		// []template.HTMLAttr
		assert.Equal(t, []template.HTMLAttr{"test"}, Indirect([]template.HTMLAttr{"test"}))
		assert.Equal(t, []template.HTMLAttr{"test"}, Indirect(&[]template.HTMLAttr{"test"}))
		assert.Equal(t, []template.HTMLAttr{}, Indirect((*[]template.HTMLAttr)(nil)))
		assert.Equal(t, []*template.HTMLAttr{}, Indirect((*[]*template.HTMLAttr)(nil)))
	}

	// template.JS
	//-----------------------------------------
	{
		var test template.JS
		assert.Equal(t, "test", string(Indirect(template.JS("test")).(template.JS)))
		assert.Equal(t, "", string(Indirect(&test).(template.JS)))
		assert.Equal(t, "", string(Indirect((*template.JS)(nil)).(template.JS)))

		// []template.JS
		assert.Equal(t, []template.JS{"test"}, Indirect([]template.JS{"test"}))
		assert.Equal(t, []template.JS{"test"}, Indirect(&[]template.JS{"test"}))
		assert.Equal(t, []template.JS{}, Indirect((*[]template.JS)(nil)))
		assert.Equal(t, []*template.JS{}, Indirect((*[]*template.JS)(nil)))
	}

	// template.URL
	//-----------------------------------------
	{
		var test template.URL
		assert.Equal(t, "test", string(Indirect(template.URL("test")).(template.URL)))
		assert.Equal(t, "", string(Indirect(&test).(template.URL)))
		assert.Equal(t, "", string(Indirect((*template.URL)(nil)).(template.URL)))

		// []template.URL
		assert.Equal(t, []template.URL{"test"}, Indirect([]template.URL{"test"}))
		assert.Equal(t, []template.URL{"test"}, Indirect(&[]template.URL{"test"}))
		assert.Equal(t, []template.URL{}, Indirect((*[]template.URL)(nil)))
		assert.Equal(t, []*template.URL{}, Indirect((*[]*template.URL)(nil)))
	}

	// uint
	//-----------------------------------------
	{
		var test uint
		assert.Equal(t, uint(1), Indirect(uint(1)))
		assert.Equal(t, uint(0), Indirect(&test))
		assert.Equal(t, uint(0), Indirect((*uint)(nil)))

		// []uint
		assert.Equal(t, []uint{1, 2, 3}, Indirect([]uint{1, 2, 3}))
		assert.Equal(t, []uint{1, 2, 3}, Indirect(&[]uint{1, 2, 3}))
		assert.Equal(t, []uint{}, Indirect((*[]uint)(nil)))
		assert.Equal(t, []*uint{}, Indirect((*[]*uint)(nil)))
		assert.Equal(t, []*uint{&test}, Indirect(&[]*uint{&test}))
	}

	// uint8
	//-----------------------------------------
	{
		var test uint8
		assert.Equal(t, uint8(3), Indirect(uint8(3)))
		assert.Equal(t, uint8(0), Indirect(&test))
		assert.Equal(t, uint8(0), Indirect((*uint8)(nil)))

		// []uint8
		assert.Equal(t, []uint8{1, 2, 3}, Indirect([]uint8{1, 2, 3}))
		assert.Equal(t, []uint8{1, 2, 3}, Indirect(&[]uint8{1, 2, 3}))
		assert.Equal(t, []uint8{}, Indirect((*[]uint8)(nil)))
		assert.Equal(t, []*uint8{}, Indirect((*[]*uint8)(nil)))
		assert.Equal(t, []*uint8{&test}, Indirect(&[]*uint8{&test}))
	}

	// uint16
	//-----------------------------------------
	{
		var test uint16
		assert.Equal(t, uint16(3), Indirect(uint16(3)))
		assert.Equal(t, uint16(0), Indirect(&test))
		assert.Equal(t, uint16(0), Indirect((*uint16)(nil)))

		// []uint16
		assert.Equal(t, []uint16{1, 2, 3}, Indirect([]uint16{1, 2, 3}))
		assert.Equal(t, []uint16{1, 2, 3}, Indirect(&[]uint16{1, 2, 3}))
		assert.Equal(t, []uint16{}, Indirect((*[]uint16)(nil)))
		assert.Equal(t, []*uint16{}, Indirect((*[]*uint16)(nil)))
		assert.Equal(t, []*uint16{&test}, Indirect(&[]*uint16{&test}))
	}

	// uint32
	//-----------------------------------------
	{
		var test uint32
		assert.Equal(t, uint32(3), Indirect(uint32(3)))
		assert.Equal(t, uint32(0), Indirect(&test))
		assert.Equal(t, uint32(0), Indirect((*uint32)(nil)))

		// []uint32
		assert.Equal(t, []uint32{1, 2, 3}, Indirect([]uint32{1, 2, 3}))
		assert.Equal(t, []uint32{1, 2, 3}, Indirect(&[]uint32{1, 2, 3}))
		assert.Equal(t, []uint32{}, Indirect((*[]uint32)(nil)))
		assert.Equal(t, []*uint32{}, Indirect((*[]*uint32)(nil)))
		assert.Equal(t, []*uint32{&test}, Indirect(&[]*uint32{&test}))
	}

	// uint64
	//-----------------------------------------
	{
		var test uint64
		assert.Equal(t, uint64(3), Indirect(uint64(3)))
		assert.Equal(t, uint64(0), Indirect(&test))
		assert.Equal(t, uint64(0), Indirect((*uint64)(nil)))

		// []uint64
		assert.Equal(t, []uint64{1, 2, 3}, Indirect([]uint64{1, 2, 3}))
		assert.Equal(t, []uint64{1, 2, 3}, Indirect(&[]uint64{1, 2, 3}))
		assert.Equal(t, []uint64{}, Indirect((*[]uint64)(nil)))
		assert.Equal(t, []*uint64{}, Indirect((*[]*uint64)(nil)))
		assert.Equal(t, []*uint64{&test}, Indirect(&[]*uint64{&test}))
	}
}

// ToBool
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

	// byte
	{
		var test byte
		assert.Equal(t, true, ToBool(byte(3)))
		assert.Equal(t, false, ToBool(byte(0)))
		assert.Equal(t, false, ToBool(&test))
		assert.Equal(t, false, ToBool((*byte)(nil)))
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

	// int
	{
		assert.Equal(t, true, ToBool(Object{1}))
		assert.Equal(t, false, ToBool(Object{0}))
		assert.Equal(t, true, ToBool(&Object{1}))
		assert.Equal(t, false, ToBool((*Object)(nil)))
	}

	// Str
	{
		assert.Equal(t, true, ToBool(NewStr("1")))
		assert.Equal(t, true, ToBool(NewStr("true")))
		assert.Equal(t, true, ToBool(NewStr("TRUE")))
		assert.Equal(t, false, ToBool(NewStr("0")))
		assert.Equal(t, false, ToBool(NewStr("false")))
		assert.Equal(t, false, ToBool(NewStr("FALSE")))
		assert.Equal(t, false, ToBool(NewStr("")))
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

// ToBoolE
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

		val, err = ToBoolE(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to bool", err.Error())
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

	// Object
	{
		val, err := ToBoolE(Object{1})
		assert.Nil(t, err)
		assert.Equal(t, true, val)

		val, err = ToBoolE(Object{0})
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, false, val)

		val, err = ToBoolE((*Object)(nil))
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

// ToInt
//--------------------------------------------------------------------------------------------------
func ExampleToInt() {
	fmt.Println(ToInt("3"))
	// Output: 3
}

func TestToInt(t *testing.T) {

	// invalid
	{
		assert.Equal(t, 0, ToInt(nil))
		assert.Equal(t, 0, ToInt(&Object{}))
	}

	// bool
	{
		var test bool
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 1, ToInt(true))
		assert.Equal(t, 0, ToInt((*bool)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(3))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(int8(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(int16(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(int32(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(int64(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*int64)(nil)))
	}

	// string
	{
		var test string
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt("3"))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*string)(nil)))
		assert.Equal(t, 0, ToInt("0"))
		assert.Equal(t, 1, ToInt("true"))
		assert.Equal(t, 1, ToInt("TRUE"))
		assert.Equal(t, 0, ToInt("false"))
		assert.Equal(t, 0, ToInt("FALSE"))
		assert.Equal(t, 0, ToInt("bob"))
	}

	// uint
	{
		var test uint
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(uint(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(uint8(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(uint16(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(uint32(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, 0, ToInt(&test))
		assert.Equal(t, 3, ToInt(uint64(3)))
		assert.Equal(t, 0, ToInt(0))
		assert.Equal(t, 0, ToInt((*uint64)(nil)))
	}
}

// ToIntE
//--------------------------------------------------------------------------------------------------
func ExampleToIntE() {
	fmt.Println(ToIntE("1"))
	// Output: 1 <nil>
}

func TestToIntE(t *testing.T) {

	// invalid
	{
		val, err := ToIntE(nil)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to int", err.Error())
		assert.Equal(t, 0, val)
	}

	// bool
	{
		val, err := ToIntE(true)
		assert.Nil(t, err)
		assert.Equal(t, 1, val)

		var test bool
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// int
	{
		val, err := ToIntE(3)
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(0)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test int
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// int8
	{
		val, err := ToIntE(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test int8
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// int16
	{
		val, err := ToIntE(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test int16
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// int32
	{
		val, err := ToIntE(int32(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(int32(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test int32
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// int64
	{
		val, err := ToIntE(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test int64
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// Object
	{
		val, err := ToIntE(Object{3})
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(Object{0})
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*Object)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// Str
	{
		val, err := ToIntE(NewStr("2"))
		assert.Nil(t, err)
		assert.Equal(t, 2, val)

		var test Str
		val, err = ToIntE(&test)
		assert.NotNil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*Str)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, 0, val)
	}

	// string
	{
		val, err := ToIntE("true")
		assert.Nil(t, err)
		assert.Equal(t, 1, val)

		val, err = ToIntE("TRUE")
		assert.Nil(t, err)
		assert.Equal(t, 1, val)

		val, err = ToIntE("false")
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE("FALSE")
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE("bob")
		assert.NotNil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE("3")
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE("0")
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test string
		val, err = ToIntE(&test)
		assert.NotNil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*string)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, 0, val)
	}

	// uint
	{
		val, err := ToIntE(uint(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(0)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test uint
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// uint8
	{
		val, err := ToIntE(uint8(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(uint8(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test uint8
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// uint16
	{
		val, err := ToIntE(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test uint16
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// uint32
	{
		val, err := ToIntE(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test uint32
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}

	// uint64
	{
		val, err := ToIntE(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, 3, val)

		val, err = ToIntE(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		var test uint64
		val, err = ToIntE(&test)
		assert.Nil(t, err)
		assert.Equal(t, 0, val)

		val, err = ToIntE((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	}
}

// ToIntSliceE
//--------------------------------------------------------------------------------------------------
func ExampleToIntSliceE() {
	fmt.Println(ToIntSliceE("1"))
	// Output: [1] <nil>
}

func TestToIntSliceE(t *testing.T) {

	// invalid
	{
		val, err := ToIntSliceE(nil)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{}, val)

		val, err = ToIntSliceE(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to []int", err.Error())
		assert.Equal(t, &IntSlice{}, val)
	}

	// bool
	{
		val, err := ToIntSliceE(true)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1}, val)

		var test bool
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// []bool
	{
		val, err := ToIntSliceE([]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1, 0}, val)

		val, err = ToIntSliceE(&[]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1, 0}, val)

		one, two := true, false
		val, err = ToIntSliceE(&[]*bool{&one, &two})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1, 0}, val)

		val, err = ToIntSliceE((*[]bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{}, val)
	}

	// interface
	//-----------------------------------------
	{
		val, err := ToIntSliceE([]interface{}{1, 2, 3})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1, 2, 3}, val)

		val, err = ToIntSliceE(&[]interface{}{1, 2, 3})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1, 2, 3}, val)
	}

	// int
	{
		val, err := ToIntSliceE(3)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(0)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test int
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// int8
	{
		val, err := ToIntSliceE(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test int8
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// int16
	{
		val, err := ToIntSliceE(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test int16
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// int32
	{
		val, err := ToIntSliceE(int32(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(int32(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test int32
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// int64
	{
		val, err := ToIntSliceE(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test int64
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// Object
	{
		val, err := ToIntSliceE(Object{3})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(Object{0})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*Object)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// Str
	{
		val, err := ToIntSliceE(NewStr("2"))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{2}, val)

		var test Str
		val, err = ToIntSliceE(&test)
		assert.NotNil(t, err)
		assert.Equal(t, &IntSlice{}, val)

		val, err = ToIntSliceE((*Str)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, &IntSlice{}, val)
	}

	// string
	{
		val, err := ToIntSliceE("true")
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1}, val)

		val, err = ToIntSliceE("TRUE")
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{1}, val)

		val, err = ToIntSliceE("false")
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE("FALSE")
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE("bob")
		assert.NotNil(t, err)
		assert.Equal(t, &IntSlice{}, val)

		val, err = ToIntSliceE("3")
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE("0")
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test string
		val, err = ToIntSliceE(&test)
		assert.NotNil(t, err)
		assert.Equal(t, &IntSlice{}, val)

		val, err = ToIntSliceE((*string)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, &IntSlice{}, val)
	}

	// uint
	{
		val, err := ToIntSliceE(uint(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(0)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test uint
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// uint8
	{
		val, err := ToIntSliceE(uint8(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(uint8(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test uint8
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// uint16
	{
		val, err := ToIntSliceE(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test uint16
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// uint32
	{
		val, err := ToIntSliceE(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test uint32
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
	}

	// uint64
	{
		val, err := ToIntSliceE(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{3}, val)

		val, err = ToIntSliceE(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		var test uint64
		val, err = ToIntSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)

		val, err = ToIntSliceE((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &IntSlice{0}, val)
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
		assert.Equal(t, "", ToString(nil))
	}

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

// ToStringSliceE
//--------------------------------------------------------------------------------------------------
func ExampleToStringSliceE() {
	fmt.Println(ToStringSliceE("1"))
	// Output: [1] <nil>
}

func TestToStringSliceE(t *testing.T) {

	// invalid
	{
		val, err := ToStringSliceE(nil)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)

		val, err = ToStringSliceE(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to []int", err.Error())
		assert.Equal(t, &StringSlice{}, val)
	}

	// bool
	{
		val, err := ToStringSliceE(true)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"true"}, val)

		var test bool
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"false"}, val)

		val, err = ToStringSliceE((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"false"}, val)
	}

	// []bool
	{
		val, err := ToStringSliceE([]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"true", "false"}, val)

		val, err = ToStringSliceE(&[]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"true", "false"}, val)

		one, two := true, false
		val, err = ToStringSliceE(&[]*bool{&one, &two})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"true", "false"}, val)

		val, err = ToStringSliceE((*[]bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)
	}

	// interface
	//-----------------------------------------
	{
		val, err := ToStringSliceE([]interface{}{"1", "2", "3"})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"1", "2", "3"}, val)

		val, err = ToStringSliceE(&[]interface{}{"1", "2", "3"})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"1", "2", "3"}, val)
	}

	// int
	{
		val, err := ToStringSliceE(3)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(0)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test int
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// int8
	{
		val, err := ToStringSliceE(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test int8
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// int16
	{
		val, err := ToStringSliceE(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test int16
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// int32
	{
		val, err := ToStringSliceE(int32(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(int32(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test int32
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// int64
	{
		val, err := ToStringSliceE(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test int64
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// Object
	{
		val, err := ToStringSliceE(Object{"3"})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(Object{"0"})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)

		val, err = ToStringSliceE((*Object)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)
	}

	// Str
	{
		val, err := ToStringSliceE(NewStr("2"))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"2"}, val)

		val, err = ToStringSliceE(NewStr("bob"))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"bob"}, val)

		var test Str
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)

		val, err = ToStringSliceE((*Str)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)
	}

	// string
	{
		val, err := ToStringSliceE("true")
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"true"}, val)

		val, err = ToStringSliceE("TRUE")
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"TRUE"}, val)

		val, err = ToStringSliceE("false")
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"false"}, val)

		val, err = ToStringSliceE("FALSE")
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"FALSE"}, val)

		val, err = ToStringSliceE("bob")
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"bob"}, val)

		val, err = ToStringSliceE("3")
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE("0")
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test string
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)

		val, err = ToStringSliceE((*string)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{}, val)
	}

	// uint
	{
		val, err := ToStringSliceE(uint(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(0)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test uint
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// uint8
	{
		val, err := ToStringSliceE(uint8(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(uint8(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test uint8
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// uint16
	{
		val, err := ToStringSliceE(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test uint16
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// uint32
	{
		val, err := ToStringSliceE(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test uint32
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}

	// uint64
	{
		val, err := ToStringSliceE(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test uint64
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		val, err = ToStringSliceE((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)
	}
}

// ToStringSliceGE
//--------------------------------------------------------------------------------------------------
func ExampleToStringSliceGE() {
	fmt.Println(ToStringSliceGE("1"))
	// Output: [1] <nil>
}

func TestToStringSliceGE(t *testing.T) {

	// invalid
	{
		val, err := ToStringSliceGE(nil)
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStringSliceGE(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to []int", err.Error())
		assert.Equal(t, []string{}, val)
	}

	// bool
	{
		val, err := ToStringSliceGE(true)
		assert.Nil(t, err)
		assert.Equal(t, []string{"true"}, val)

		var test bool
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"false"}, val)

		val, err = ToStringSliceGE((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"false"}, val)
	}

	// []bool
	{
		val, err := ToStringSliceGE([]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, []string{"true", "false"}, val)

		val, err = ToStringSliceGE(&[]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, []string{"true", "false"}, val)

		one, two := true, false
		val, err = ToStringSliceGE(&[]*bool{&one, &two})
		assert.Nil(t, err)
		assert.Equal(t, []string{"true", "false"}, val)

		val, err = ToStringSliceGE((*[]bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// interface
	//-----------------------------------------
	{
		val, err := ToStringSliceGE([]interface{}{"1", "2", "3"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "2", "3"}, val)

		val, err = ToStringSliceGE(&[]interface{}{"1", "2", "3"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "2", "3"}, val)
	}

	// int
	{
		val, err := ToStringSliceGE(3)
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(0)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// int8
	{
		val, err := ToStringSliceGE(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int8
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// int16
	{
		val, err := ToStringSliceGE(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int16
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// int32
	{
		val, err := ToStringSliceGE(int32(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(int32(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int32
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// int64
	{
		val, err := ToStringSliceGE(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int64
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// Object
	{
		val, err := ToStringSliceGE(Object{"3"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(Object{"0"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStringSliceGE((*Object)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// Str
	{
		val, err := ToStringSliceGE(NewStr("2"))
		assert.Nil(t, err)
		assert.Equal(t, []string{"2"}, val)

		val, err = ToStringSliceGE(NewStr("bob"))
		assert.Nil(t, err)
		assert.Equal(t, []string{"bob"}, val)

		var test Str
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStringSliceGE((*Str)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// string
	{
		val, err := ToStringSliceGE("true")
		assert.Nil(t, err)
		assert.Equal(t, []string{"true"}, val)

		val, err = ToStringSliceGE("TRUE")
		assert.Nil(t, err)
		assert.Equal(t, []string{"TRUE"}, val)

		val, err = ToStringSliceGE("false")
		assert.Nil(t, err)
		assert.Equal(t, []string{"false"}, val)

		val, err = ToStringSliceGE("FALSE")
		assert.Nil(t, err)
		assert.Equal(t, []string{"FALSE"}, val)

		val, err = ToStringSliceGE("bob")
		assert.Nil(t, err)
		assert.Equal(t, []string{"bob"}, val)

		val, err = ToStringSliceGE("3")
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE("0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test string
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStringSliceGE((*string)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// uint
	{
		val, err := ToStringSliceGE(uint(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(0)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// uint8
	{
		val, err := ToStringSliceGE(uint8(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(uint8(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint8
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// uint16
	{
		val, err := ToStringSliceGE(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint16
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// uint32
	{
		val, err := ToStringSliceGE(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint32
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// uint64
	{
		val, err := ToStringSliceGE(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStringSliceGE(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint64
		val, err = ToStringSliceGE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStringSliceGE((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}
}
