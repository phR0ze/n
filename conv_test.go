package n

import (
	"errors"
	"fmt"
	"html/template"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestObj struct {
	o interface{}
}

// DeReference
//--------------------------------------------------------------------------------------------------
func ExampleDeReference() {
	slice := DeReference((*[]int)(nil))
	fmt.Println(slice)
	// Output: []
}

func TestDeReference(t *testing.T) {
	// nil
	{
		assert.Equal(t, nil, DeReference(nil))
	}

	// bool
	//-----------------------------------------
	{
		var test bool
		assert.Equal(t, true, DeReference(true))
		assert.Equal(t, false, DeReference(&test))
		assert.Equal(t, false, DeReference((*bool)(nil)))

		// []bool
		assert.Equal(t, []bool{true}, DeReference([]bool{true}))
		assert.Equal(t, []bool{true}, DeReference(&[]bool{true}))
		assert.Equal(t, []*bool{&test}, DeReference(&[]*bool{&test}))
	}

	// byte
	//-----------------------------------------
	{
		var test byte
		assert.Equal(t, byte(3), DeReference(byte(3)))
		assert.Equal(t, byte(0), DeReference(&test))
		assert.Equal(t, byte(0), DeReference((*byte)(nil)))

		// []byte
		assert.Equal(t, "test", string(DeReference([]byte{0x74, 0x65, 0x73, 0x74}).([]byte)))
		assert.Equal(t, "test", string(DeReference(&[]byte{0x74, 0x65, 0x73, 0x74}).([]byte)))
		assert.Equal(t, "", string(DeReference((*[]byte)(nil)).([]byte)))
		assert.Equal(t, []*byte{&test}, DeReference(&[]*byte{&test}))
	}

	// float32
	//-----------------------------------------
	{
		var test float32
		assert.Equal(t, float32(7.22), DeReference(float32(7.22)))
		assert.Equal(t, float32(0), DeReference(&test))
		assert.Equal(t, float32(0), DeReference((*float32)(nil)))

		// []float32
		assert.Equal(t, []float32{1.2}, DeReference([]float32{1.2}))
		assert.Equal(t, []float32{1.2}, DeReference(&[]float32{1.2}))
		assert.Equal(t, []*float32{&test}, DeReference(&[]*float32{&test}))
	}

	// float64
	//-----------------------------------------
	{
		var test float64
		assert.Equal(t, float64(7.22), DeReference(float64(7.22)))
		assert.Equal(t, float64(0), DeReference(&test))
		assert.Equal(t, float64(0), DeReference((*float64)(nil)))

		// []float64
		assert.Equal(t, []float64{1.2}, DeReference([]float64{1.2}))
		assert.Equal(t, []float64{1.2}, DeReference(&[]float64{1.2}))
		assert.Equal(t, []*float64{&test}, DeReference(&[]*float64{&test}))
	}

	// FloatSlice
	//-----------------------------------------
	{
		assert.Equal(t, FloatSlice{1.2}, DeReference(FloatSlice{1.2}))
		assert.Equal(t, FloatSlice{1.2}, DeReference(&FloatSlice{1.2}))
		assert.Equal(t, FloatSlice{}, DeReference((*FloatSlice)(nil)))

		test := FloatSlice{}
		assert.Equal(t, []FloatSlice{{1.2}}, DeReference([]FloatSlice{{1.2}}))
		assert.Equal(t, []FloatSlice{{1.2}}, DeReference(&[]FloatSlice{{1.2}}))
		assert.Equal(t, []*FloatSlice{&test}, DeReference(&[]*FloatSlice{&test}))
	}

	// int
	//-----------------------------------------
	{
		var test int
		assert.Equal(t, 1, DeReference(1))
		assert.Equal(t, 0, DeReference(&test))
		assert.Equal(t, 0, DeReference((*int)(nil)))

		// []int
		assert.Equal(t, []int{1, 2, 3}, DeReference([]int{1, 2, 3}))
		assert.Equal(t, []int{1, 2, 3}, DeReference(&[]int{1, 2, 3}))
		assert.Equal(t, []int{}, DeReference((*[]int)(nil)))
		assert.Equal(t, []*int{}, DeReference((*[]*int)(nil)))
		assert.Equal(t, []*int{&test}, DeReference(&[]*int{&test}))
	}

	// interface
	//-----------------------------------------
	{
		assert.Equal(t, []interface{}{1, 2, 3}, DeReference([]interface{}{1, 2, 3}))
		assert.Equal(t, []interface{}{1, 2, 3}, DeReference(&[]interface{}{1, 2, 3}))
		assert.Equal(t, []interface{}{}, DeReference((*[]interface{})(nil)))
		assert.Equal(t, []*interface{}{}, DeReference((*[]*interface{})(nil)))
	}

	// int8
	//-----------------------------------------
	{
		var test int8
		assert.Equal(t, int8(3), DeReference(int8(3)))
		assert.Equal(t, int8(0), DeReference(&test))
		assert.Equal(t, int8(0), DeReference((*int8)(nil)))

		// []int8
		assert.Equal(t, []int8{1, 2, 3}, DeReference([]int8{1, 2, 3}))
		assert.Equal(t, []int8{1, 2, 3}, DeReference(&[]int8{1, 2, 3}))
		assert.Equal(t, []int8{}, DeReference((*[]int8)(nil)))
		assert.Equal(t, []*int8{}, DeReference((*[]*int8)(nil)))
		assert.Equal(t, []*int8{&test}, DeReference(&[]*int8{&test}))
	}

	// int16
	//-----------------------------------------
	{
		var test int16
		assert.Equal(t, int16(3), DeReference(int16(3)))
		assert.Equal(t, int16(0), DeReference(&test))
		assert.Equal(t, int16(0), DeReference((*int16)(nil)))

		// []int16
		assert.Equal(t, []int16{1, 2, 3}, DeReference([]int16{1, 2, 3}))
		assert.Equal(t, []int16{1, 2, 3}, DeReference(&[]int16{1, 2, 3}))
		assert.Equal(t, []int16{}, DeReference((*[]int16)(nil)))
		assert.Equal(t, []*int16{}, DeReference((*[]*int16)(nil)))
		assert.Equal(t, []*int16{&test}, DeReference(&[]*int16{&test}))
	}

	// int32
	//-----------------------------------------
	{
		var test int32
		assert.Equal(t, int32(3), DeReference(int32(3)))
		assert.Equal(t, int32(0), DeReference(&test))
		assert.Equal(t, int32(0), DeReference((*int32)(nil)))

		// []int32
		assert.Equal(t, []int32{1, 2, 3}, DeReference([]int32{1, 2, 3}))
		assert.Equal(t, []int32{1, 2, 3}, DeReference(&[]int32{1, 2, 3}))
		assert.Equal(t, []int32{}, DeReference((*[]int32)(nil)))
		assert.Equal(t, []*int32{}, DeReference((*[]*int32)(nil)))
		assert.Equal(t, []*int32{&test}, DeReference(&[]*int32{&test}))
	}

	// int64
	//-----------------------------------------
	{
		var test int64
		assert.Equal(t, int64(3), DeReference(int64(3)))
		assert.Equal(t, int64(0), DeReference(&test))
		assert.Equal(t, int64(0), DeReference((*int64)(nil)))

		// []int64
		assert.Equal(t, []int64{1, 2, 3}, DeReference([]int64{1, 2, 3}))
		assert.Equal(t, []int64{1, 2, 3}, DeReference(&[]int64{1, 2, 3}))
		assert.Equal(t, []int64{}, DeReference((*[]int64)(nil)))
		assert.Equal(t, []*int64{}, DeReference((*[]*int64)(nil)))
		assert.Equal(t, []*int64{&test}, DeReference(&[]*int64{&test}))
	}

	// IntSlice
	//-----------------------------------------
	{
		assert.Equal(t, IntSlice{2}, DeReference(IntSlice{2}))
		assert.Equal(t, IntSlice{2}, DeReference(&IntSlice{2}))
		assert.Equal(t, IntSlice{}, DeReference((*IntSlice)(nil)))

		test := IntSlice{}
		assert.Equal(t, []IntSlice{{2}}, DeReference([]IntSlice{{2}}))
		assert.Equal(t, []IntSlice{{2}}, DeReference(&[]IntSlice{{2}}))
		assert.Equal(t, []*IntSlice{&test}, DeReference(&[]*IntSlice{&test}))
	}

	// InterSlice
	//-----------------------------------------
	{
		assert.Equal(t, InterSlice{2}, DeReference(InterSlice{2}))
		assert.Equal(t, InterSlice{2}, DeReference(&InterSlice{2}))
		assert.Equal(t, InterSlice{}, DeReference((*InterSlice)(nil)))

		test := InterSlice{}
		assert.Equal(t, []InterSlice{{2}}, DeReference([]InterSlice{{2}}))
		assert.Equal(t, []InterSlice{{2}}, DeReference(&[]InterSlice{{2}}))
		assert.Equal(t, []*InterSlice{&test}, DeReference(&[]*InterSlice{&test}))
	}

	// Object
	//-----------------------------------------
	{
		assert.Equal(t, Object{3}, DeReference(Object{3}))
		assert.Equal(t, Object{3}, DeReference(&Object{3}))
		assert.Equal(t, Object{}, DeReference((*Object)(nil)))

		// []Object
		assert.Equal(t, []Object{{1}, {2}, {3}}, DeReference([]Object{{1}, {2}, {3}}))
		assert.Equal(t, []Object{{1}, {2}, {3}}, DeReference(&[]Object{{1}, {2}, {3}}))
		assert.Equal(t, []Object{}, DeReference((*[]Object)(nil)))
		assert.Equal(t, []*Object{}, DeReference((*[]*Object)(nil)))
		assert.Equal(t, []*Object{&Object{}}, DeReference(&[]*Object{&Object{}}))
	}

	// rune
	//-----------------------------------------
	{
		var test rune
		assert.Equal(t, 'r', DeReference('r'))
		assert.Equal(t, rune(0), DeReference(&test))
		assert.Equal(t, rune(0), DeReference((*rune)(nil)))

		// []rune
		assert.Equal(t, []rune{'t', 'e', 's', 't'}, DeReference([]rune{'t', 'e', 's', 't'}))
		assert.Equal(t, []rune{'t', 'e', 's', 't'}, DeReference(&[]rune{'t', 'e', 's', 't'}))
		assert.Equal(t, []rune{}, DeReference((*[]rune)(nil)))
		assert.Equal(t, []*rune{}, DeReference((*[]*rune)(nil)))
		assert.Equal(t, []*rune{&test}, DeReference(&[]*rune{&test}))
	}

	// Str
	//-----------------------------------------
	{
		assert.Equal(t, *A("test"), DeReference(*A("test")))
		assert.Equal(t, *A("test"), DeReference(A("test")))
		assert.Equal(t, *A(""), DeReference((*Str)(nil)))

		// []Str
		assert.Equal(t, []Str{{1}, {2}, {3}}, DeReference([]Str{{1}, {2}, {3}}))
		assert.Equal(t, []Str{{1}, {2}, {3}}, DeReference(&[]Str{{1}, {2}, {3}}))
		assert.Equal(t, []Str{}, DeReference((*[]Str)(nil)))
		assert.Equal(t, []*Str{}, DeReference((*[]*Str)(nil)))
		assert.Equal(t, []*Str{&Str{}}, DeReference(&[]*Str{&Str{}}))
	}

	// string
	//-----------------------------------------
	{
		var test string
		assert.Equal(t, "test", DeReference("test"))
		assert.Equal(t, "", DeReference(&test))
		assert.Equal(t, "", DeReference((*string)(nil)))

		// []string
		assert.Equal(t, []string{"test"}, DeReference([]string{"test"}))
		assert.Equal(t, []string{"test"}, DeReference(&[]string{"test"}))
		assert.Equal(t, []string{}, DeReference((*[]string)(nil)))
		assert.Equal(t, []*string{}, DeReference((*[]*string)(nil)))
	}

	// StringSlice
	//-----------------------------------------
	{
		assert.Equal(t, StringSlice{"2"}, DeReference(StringSlice{"2"}))
		assert.Equal(t, StringSlice{"2"}, DeReference(&StringSlice{"2"}))
		assert.Equal(t, StringSlice{}, DeReference((*StringSlice)(nil)))

		test := StringSlice{}
		assert.Equal(t, []StringSlice{{"2"}}, DeReference([]StringSlice{{"2"}}))
		assert.Equal(t, []StringSlice{{"2"}}, DeReference(&[]StringSlice{{"2"}}))
		assert.Equal(t, []*StringSlice{&test}, DeReference(&[]*StringSlice{&test}))
	}

	// StringMap
	//-----------------------------------------
	{
		assert.Equal(t, *M().Add("2", "two"), DeReference(*M().Add("2", "two")))
		assert.Equal(t, *M().Add("2", "two"), DeReference(M().Add("2", "two")))
		assert.Equal(t, *M(), DeReference((*StringMap)(nil)))

		assert.Equal(t, []*StringMap{M().Add("2", "two")}, DeReference([]*StringMap{M().Add("2", "two")}))
		assert.Equal(t, []*StringMap{M().Add("2", "two")}, DeReference(&[]*StringMap{M().Add("2", "two")}))
	}

	// template.CSS
	//-----------------------------------------
	{
		var test template.CSS
		assert.Equal(t, "test", string(DeReference(template.CSS("test")).(template.CSS)))
		assert.Equal(t, "", string(DeReference(&test).(template.CSS)))
		assert.Equal(t, "", string(DeReference((*template.CSS)(nil)).(template.CSS)))

		// []template.CSS
		assert.Equal(t, []template.CSS{"test"}, DeReference([]template.CSS{"test"}))
		assert.Equal(t, []template.CSS{"test"}, DeReference(&[]template.CSS{"test"}))
		assert.Equal(t, []template.CSS{}, DeReference((*[]template.CSS)(nil)))
		assert.Equal(t, []*template.CSS{}, DeReference((*[]*template.CSS)(nil)))
	}

	// template.HTML
	//-----------------------------------------
	{
		var test template.HTML
		assert.Equal(t, "test", string(DeReference(template.HTML("test")).(template.HTML)))
		assert.Equal(t, "", string(DeReference(&test).(template.HTML)))
		assert.Equal(t, "", string(DeReference((*template.HTML)(nil)).(template.HTML)))

		// []template.HTML
		assert.Equal(t, []template.HTML{"test"}, DeReference([]template.HTML{"test"}))
		assert.Equal(t, []template.HTML{"test"}, DeReference(&[]template.HTML{"test"}))
		assert.Equal(t, []template.HTML{}, DeReference((*[]template.HTML)(nil)))
		assert.Equal(t, []*template.HTML{}, DeReference((*[]*template.HTML)(nil)))
	}

	// template.HTMLAttr
	//-----------------------------------------
	{
		var test template.HTMLAttr
		assert.Equal(t, "test", string(DeReference(template.HTMLAttr("test")).(template.HTMLAttr)))
		assert.Equal(t, "", string(DeReference(&test).(template.HTMLAttr)))
		assert.Equal(t, "", string(DeReference((*template.HTMLAttr)(nil)).(template.HTMLAttr)))

		// []template.HTMLAttr
		assert.Equal(t, []template.HTMLAttr{"test"}, DeReference([]template.HTMLAttr{"test"}))
		assert.Equal(t, []template.HTMLAttr{"test"}, DeReference(&[]template.HTMLAttr{"test"}))
		assert.Equal(t, []template.HTMLAttr{}, DeReference((*[]template.HTMLAttr)(nil)))
		assert.Equal(t, []*template.HTMLAttr{}, DeReference((*[]*template.HTMLAttr)(nil)))
	}

	// template.JS
	//-----------------------------------------
	{
		var test template.JS
		assert.Equal(t, "test", string(DeReference(template.JS("test")).(template.JS)))
		assert.Equal(t, "", string(DeReference(&test).(template.JS)))
		assert.Equal(t, "", string(DeReference((*template.JS)(nil)).(template.JS)))

		// []template.JS
		assert.Equal(t, []template.JS{"test"}, DeReference([]template.JS{"test"}))
		assert.Equal(t, []template.JS{"test"}, DeReference(&[]template.JS{"test"}))
		assert.Equal(t, []template.JS{}, DeReference((*[]template.JS)(nil)))
		assert.Equal(t, []*template.JS{}, DeReference((*[]*template.JS)(nil)))
	}

	// template.URL
	//-----------------------------------------
	{
		var test template.URL
		assert.Equal(t, "test", string(DeReference(template.URL("test")).(template.URL)))
		assert.Equal(t, "", string(DeReference(&test).(template.URL)))
		assert.Equal(t, "", string(DeReference((*template.URL)(nil)).(template.URL)))

		// []template.URL
		assert.Equal(t, []template.URL{"test"}, DeReference([]template.URL{"test"}))
		assert.Equal(t, []template.URL{"test"}, DeReference(&[]template.URL{"test"}))
		assert.Equal(t, []template.URL{}, DeReference((*[]template.URL)(nil)))
		assert.Equal(t, []*template.URL{}, DeReference((*[]*template.URL)(nil)))
	}

	// uint
	//-----------------------------------------
	{
		var test uint
		assert.Equal(t, uint(1), DeReference(uint(1)))
		assert.Equal(t, uint(0), DeReference(&test))
		assert.Equal(t, uint(0), DeReference((*uint)(nil)))

		// []uint
		assert.Equal(t, []uint{1, 2, 3}, DeReference([]uint{1, 2, 3}))
		assert.Equal(t, []uint{1, 2, 3}, DeReference(&[]uint{1, 2, 3}))
		assert.Equal(t, []uint{}, DeReference((*[]uint)(nil)))
		assert.Equal(t, []*uint{}, DeReference((*[]*uint)(nil)))
		assert.Equal(t, []*uint{&test}, DeReference(&[]*uint{&test}))
	}

	// uint8
	//-----------------------------------------
	{
		var test uint8
		assert.Equal(t, uint8(3), DeReference(uint8(3)))
		assert.Equal(t, uint8(0), DeReference(&test))
		assert.Equal(t, uint8(0), DeReference((*uint8)(nil)))

		// []uint8
		assert.Equal(t, []uint8{1, 2, 3}, DeReference([]uint8{1, 2, 3}))
		assert.Equal(t, []uint8{1, 2, 3}, DeReference(&[]uint8{1, 2, 3}))
		assert.Equal(t, []uint8{}, DeReference((*[]uint8)(nil)))
		assert.Equal(t, []*uint8{}, DeReference((*[]*uint8)(nil)))
		assert.Equal(t, []*uint8{&test}, DeReference(&[]*uint8{&test}))
	}

	// uint16
	//-----------------------------------------
	{
		var test uint16
		assert.Equal(t, uint16(3), DeReference(uint16(3)))
		assert.Equal(t, uint16(0), DeReference(&test))
		assert.Equal(t, uint16(0), DeReference((*uint16)(nil)))

		// []uint16
		assert.Equal(t, []uint16{1, 2, 3}, DeReference([]uint16{1, 2, 3}))
		assert.Equal(t, []uint16{1, 2, 3}, DeReference(&[]uint16{1, 2, 3}))
		assert.Equal(t, []uint16{}, DeReference((*[]uint16)(nil)))
		assert.Equal(t, []*uint16{}, DeReference((*[]*uint16)(nil)))
		assert.Equal(t, []*uint16{&test}, DeReference(&[]*uint16{&test}))
	}

	// uint32
	//-----------------------------------------
	{
		var test uint32
		assert.Equal(t, uint32(3), DeReference(uint32(3)))
		assert.Equal(t, uint32(0), DeReference(&test))
		assert.Equal(t, uint32(0), DeReference((*uint32)(nil)))

		// []uint32
		assert.Equal(t, []uint32{1, 2, 3}, DeReference([]uint32{1, 2, 3}))
		assert.Equal(t, []uint32{1, 2, 3}, DeReference(&[]uint32{1, 2, 3}))
		assert.Equal(t, []uint32{}, DeReference((*[]uint32)(nil)))
		assert.Equal(t, []*uint32{}, DeReference((*[]*uint32)(nil)))
		assert.Equal(t, []*uint32{&test}, DeReference(&[]*uint32{&test}))
	}

	// uint64
	//-----------------------------------------
	{
		var test uint64
		assert.Equal(t, uint64(3), DeReference(uint64(3)))
		assert.Equal(t, uint64(0), DeReference(&test))
		assert.Equal(t, uint64(0), DeReference((*uint64)(nil)))

		// []uint64
		assert.Equal(t, []uint64{1, 2, 3}, DeReference([]uint64{1, 2, 3}))
		assert.Equal(t, []uint64{1, 2, 3}, DeReference(&[]uint64{1, 2, 3}))
		assert.Equal(t, []uint64{}, DeReference((*[]uint64)(nil)))
		assert.Equal(t, []*uint64{}, DeReference((*[]*uint64)(nil)))
		assert.Equal(t, []*uint64{&test}, DeReference(&[]*uint64{&test}))
	}
}

// Reference
//--------------------------------------------------------------------------------------------------
func ExampleReference() {
	slice := Reference([]int{1, 2, 3})
	fmt.Println(slice)
	// Output: &[1 2 3]
}

func TestReference(t *testing.T) {
	// nil
	{
		assert.Equal(t, nil, Reference(nil))
	}

	// bool
	//-----------------------------------------
	{
		trueVal := true
		falseVal := false
		assert.Equal(t, &trueVal, Reference(true))
		assert.Equal(t, &falseVal, Reference(&falseVal))
		assert.Equal(t, (*bool)(nil), Reference((*bool)(nil)))

		// []bool
		val1 := []bool{true}
		val2 := []*bool{&trueVal}
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, &val2, Reference(&[]*bool{&trueVal}))
	}

	// byte
	//-----------------------------------------
	{
		var val1 byte
		val2 := byte(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*byte)(nil), Reference((*byte)(nil)))

		// []byte
		val3 := []byte{0x74, 0x65, 0x73, 0x74}
		val4 := []*byte{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]byte)(nil), Reference((*[]byte)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// float32
	//-----------------------------------------
	{
		var val1 float32
		val2 := float32(7.22)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*float32)(nil), Reference((*float32)(nil)))

		// []float32
		val3 := []float32{1.2}
		val4 := []*float32{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, &val4, Reference(&val4))
	}

	// float64
	//-----------------------------------------
	{
		var val1 float64
		val2 := float64(7.22)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*float64)(nil), Reference((*float64)(nil)))

		// []float64
		val3 := []float64{1.2}
		val4 := []*float64{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, &val4, Reference(&val4))
	}

	// int
	//-----------------------------------------
	{
		var val1 int
		val2 := int(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*int)(nil), Reference((*int)(nil)))

		// []int
		val3 := []int{1, 2, 3}
		val4 := []*int{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]int)(nil), Reference((*[]int)(nil)))
		assert.Equal(t, (*[]*int)(nil), Reference((*[]*int)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// interface
	//-----------------------------------------
	{
		val1 := []interface{}{1, 2, 3}
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*[]interface{})(nil), Reference((*[]interface{})(nil)))
		assert.Equal(t, (*[]*interface{})(nil), Reference((*[]*interface{})(nil)))
	}

	// int8
	//-----------------------------------------
	{
		var val1 int8
		val2 := int8(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*int8)(nil), Reference((*int8)(nil)))

		// []int8
		val3 := []int8{1, 2, 3}
		val4 := []*int8{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]int8)(nil), Reference((*[]int8)(nil)))
		assert.Equal(t, (*[]*int8)(nil), Reference((*[]*int8)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// int16
	//-----------------------------------------
	{
		var val1 int16
		val2 := int16(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*int16)(nil), Reference((*int16)(nil)))

		// []int16
		val3 := []int16{1, 2, 3}
		val4 := []*int16{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]int16)(nil), Reference((*[]int16)(nil)))
		assert.Equal(t, (*[]*int16)(nil), Reference((*[]*int16)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// int32
	//-----------------------------------------
	{
		var val1 int32
		val2 := int32(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*int32)(nil), Reference((*int32)(nil)))

		// []int32
		val3 := []int32{1, 2, 3}
		val4 := []*int32{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]int32)(nil), Reference((*[]int32)(nil)))
		assert.Equal(t, (*[]*int32)(nil), Reference((*[]*int32)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// int64
	//-----------------------------------------
	{
		var val1 int64
		val2 := int64(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*int64)(nil), Reference((*int64)(nil)))

		// []int64
		val3 := []int64{1, 2, 3}
		val4 := []*int64{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]int64)(nil), Reference((*[]int64)(nil)))
		assert.Equal(t, (*[]*int64)(nil), Reference((*[]*int64)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// Object
	//-----------------------------------------
	{
		val1 := Object{3}
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*Object)(nil), Reference((*Object)(nil)))

		// []Object
		val2 := []Object{{1}, {2}, {3}}
		val3 := []*Object{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]Object)(nil), Reference((*[]Object)(nil)))
		assert.Equal(t, (*[]*Object)(nil), Reference((*[]*Object)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// rune
	//-----------------------------------------
	{
		val1 := rune('r')
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*rune)(nil), Reference((*rune)(nil)))

		// []rune
		val2 := []rune{'t', 'e', 's', 't'}
		val3 := []*rune{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]rune)(nil), Reference((*[]rune)(nil)))
		assert.Equal(t, (*[]*rune)(nil), Reference((*[]*rune)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// Str
	//-----------------------------------------
	{
		val1 := *NewStr("test")
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*Str)(nil), Reference((*Str)(nil)))

		// []Str
		val2 := []Str{{1}, {2}, {3}}
		val3 := []*Str{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]Str)(nil), Reference((*[]Str)(nil)))
		assert.Equal(t, (*[]*Str)(nil), Reference((*[]*Str)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// string
	//-----------------------------------------
	{
		val1 := "test"
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*string)(nil), Reference((*string)(nil)))

		// []string
		val2 := []string{"1", "2", "3"}
		val3 := []*string{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]string)(nil), Reference((*[]string)(nil)))
		assert.Equal(t, (*[]*string)(nil), Reference((*[]*string)(nil)))
		assert.Equal(t, &val3, Reference(val3))

		// map[interface]interface{}
		{
			map2 := map[interface{}]interface{}{"1": 1}
			assert.Equal(t, &map2, Reference(map2))
			assert.Equal(t, &map2, Reference(&map2))
			assert.Equal(t, (*map[interface{}]interface{})(nil), Reference((*map[interface{}]interface{})(nil)))
			assert.Equal(t, (*map[interface{}]*interface{})(nil), Reference((*map[interface{}]*interface{})(nil)))
		}

		// map[string]interface{}
		{
			map2 := map[string]interface{}{"1": 1}
			assert.Equal(t, &map2, Reference(map2))
			assert.Equal(t, &map2, Reference(&map2))
			assert.Equal(t, (*map[string]interface{})(nil), Reference((*map[string]interface{})(nil)))
			assert.Equal(t, (*map[string]*interface{})(nil), Reference((*map[string]*interface{})(nil)))
		}

		// map[string]string
		{
			map2 := map[string]string{"1": "one"}
			assert.Equal(t, &map2, Reference(map2))
			assert.Equal(t, &map2, Reference(&map2))
			assert.Equal(t, (*map[string]string)(nil), Reference((*map[string]string)(nil)))
			assert.Equal(t, (*map[string]*string)(nil), Reference((*map[string]*string)(nil)))
		}

		// map[string]float32
		{
			map2 := map[string]float32{"1": 1.0}
			assert.Equal(t, &map2, Reference(map2))
			assert.Equal(t, &map2, Reference(&map2))
			assert.Equal(t, (*map[string]float32)(nil), Reference((*map[string]float32)(nil)))
			assert.Equal(t, (*map[string]*float32)(nil), Reference((*map[string]*float32)(nil)))
		}

		// map[string]float64
		{
			map2 := map[string]float64{"1": 1.0}
			assert.Equal(t, &map2, Reference(map2))
			assert.Equal(t, &map2, Reference(&map2))
			assert.Equal(t, (*map[string]float64)(nil), Reference((*map[string]float64)(nil)))
			assert.Equal(t, (*map[string]*float64)(nil), Reference((*map[string]*float64)(nil)))
		}

		// map[string]int
		{
			map2 := map[string]int{"1": 1}
			assert.Equal(t, &map2, Reference(map2))
			assert.Equal(t, &map2, Reference(&map2))
			assert.Equal(t, (*map[string]int)(nil), Reference((*map[string]int)(nil)))
			assert.Equal(t, (*map[string]*int)(nil), Reference((*map[string]*int)(nil)))
		}

		// map[string]int64
		{
			map2 := map[string]int64{"1": 1}
			assert.Equal(t, &map2, Reference(map2))
			assert.Equal(t, &map2, Reference(&map2))
			assert.Equal(t, (*map[string]int64)(nil), Reference((*map[string]int64)(nil)))
			assert.Equal(t, (*map[string]*int64)(nil), Reference((*map[string]*int64)(nil)))
		}
	}

	// template.CSS
	//-----------------------------------------
	{
		val1 := template.CSS("test")
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*template.CSS)(nil), Reference((*template.CSS)(nil)))

		// []template.CSS
		val2 := []template.CSS{"1", "2", "3"}
		val3 := []*template.CSS{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]template.CSS)(nil), Reference((*[]template.CSS)(nil)))
		assert.Equal(t, (*[]*template.CSS)(nil), Reference((*[]*template.CSS)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// template.HTML
	//-----------------------------------------
	{
		val1 := template.HTML("test")
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*template.HTML)(nil), Reference((*template.HTML)(nil)))

		// []template.HTML
		val2 := []template.HTML{"1", "2", "3"}
		val3 := []*template.HTML{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]template.HTML)(nil), Reference((*[]template.HTML)(nil)))
		assert.Equal(t, (*[]*template.HTML)(nil), Reference((*[]*template.HTML)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// template.HTMLAttr
	//-----------------------------------------
	{
		val1 := template.HTMLAttr("test")
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*template.HTMLAttr)(nil), Reference((*template.HTMLAttr)(nil)))

		// []template.HTMLAttr
		val2 := []template.HTMLAttr{"1", "2", "3"}
		val3 := []*template.HTMLAttr{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]template.HTMLAttr)(nil), Reference((*[]template.HTMLAttr)(nil)))
		assert.Equal(t, (*[]*template.HTMLAttr)(nil), Reference((*[]*template.HTMLAttr)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// template.JS
	//-----------------------------------------
	{
		val1 := template.JS("test")
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*template.JS)(nil), Reference((*template.JS)(nil)))

		// []template.JS
		val2 := []template.JS{"1", "2", "3"}
		val3 := []*template.JS{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]template.JS)(nil), Reference((*[]template.JS)(nil)))
		assert.Equal(t, (*[]*template.JS)(nil), Reference((*[]*template.JS)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// template.URL
	//-----------------------------------------
	{
		val1 := template.URL("test")
		assert.Equal(t, &val1, Reference(val1))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*template.URL)(nil), Reference((*template.URL)(nil)))

		// []template.URL
		val2 := []template.URL{"1", "2", "3"}
		val3 := []*template.URL{&val1}
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val2, Reference(&val2))
		assert.Equal(t, (*[]template.URL)(nil), Reference((*[]template.URL)(nil)))
		assert.Equal(t, (*[]*template.URL)(nil), Reference((*[]*template.URL)(nil)))
		assert.Equal(t, &val3, Reference(val3))
	}

	// uint
	//-----------------------------------------
	{
		var val1 uint
		val2 := uint(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*uint)(nil), Reference((*uint)(nil)))

		// []uint
		val3 := []uint{1, 2, 3}
		val4 := []*uint{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]uint)(nil), Reference((*[]uint)(nil)))
		assert.Equal(t, (*[]*uint)(nil), Reference((*[]*uint)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// uint8
	//-----------------------------------------
	{
		var val1 uint8
		val2 := uint8(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*uint8)(nil), Reference((*uint8)(nil)))

		// []uint8
		val3 := []uint8{1, 2, 3}
		val4 := []*uint8{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]uint8)(nil), Reference((*[]uint8)(nil)))
		assert.Equal(t, (*[]*uint8)(nil), Reference((*[]*uint8)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// uint16
	//-----------------------------------------
	{
		var val1 uint16
		val2 := uint16(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*uint16)(nil), Reference((*uint16)(nil)))

		// []uint16
		val3 := []uint16{1, 2, 3}
		val4 := []*uint16{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]uint16)(nil), Reference((*[]uint16)(nil)))
		assert.Equal(t, (*[]*uint16)(nil), Reference((*[]*uint16)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// uint32
	//-----------------------------------------
	{
		var val1 uint32
		val2 := uint32(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*uint32)(nil), Reference((*uint32)(nil)))

		// []uint32
		val3 := []uint32{1, 2, 3}
		val4 := []*uint32{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]uint32)(nil), Reference((*[]uint32)(nil)))
		assert.Equal(t, (*[]*uint32)(nil), Reference((*[]*uint32)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}

	// uint64
	//-----------------------------------------
	{
		var val1 uint64
		val2 := uint64(3)
		assert.Equal(t, &val2, Reference(val2))
		assert.Equal(t, &val1, Reference(&val1))
		assert.Equal(t, (*uint64)(nil), Reference((*uint64)(nil)))

		// []uint64
		val3 := []uint64{1, 2, 3}
		val4 := []*uint64{&val1}
		assert.Equal(t, &val3, Reference(val3))
		assert.Equal(t, &val3, Reference(&val3))
		assert.Equal(t, (*[]uint64)(nil), Reference((*[]uint64)(nil)))
		assert.Equal(t, (*[]*uint64)(nil), Reference((*[]*uint64)(nil)))
		assert.Equal(t, &val4, Reference(val4))
	}
}

// ToBool
//--------------------------------------------------------------------------------------------------
func ExampleToBool() {
	fmt.Println(B(1))
	// Output: true
}

func TestToBool(t *testing.T) {

	// invalid
	{
		assert.Equal(t, false, B(nil))
		assert.Equal(t, false, B(&Object{}))
	}

	// bool
	{
		var test bool
		assert.Equal(t, true, B(true))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*bool)(nil)))
	}

	// byte
	{
		var test byte
		assert.Equal(t, true, B(byte(3)))
		assert.Equal(t, false, B(byte(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*byte)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, true, B(1))
		assert.Equal(t, false, B(0))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, true, B(int8(3)))
		assert.Equal(t, false, B(int8(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, true, B(int16(3)))
		assert.Equal(t, false, B(int16(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, true, B(int32(3)))
		assert.Equal(t, false, B(int32(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, true, B(int64(3)))
		assert.Equal(t, false, B(int64(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*int64)(nil)))
	}

	// int
	{
		assert.Equal(t, true, B(Object{1}))
		assert.Equal(t, false, B(Object{0}))
		assert.Equal(t, true, B(&Object{1}))
		assert.Equal(t, false, B((*Object)(nil)))
	}

	// Str
	{
		assert.Equal(t, true, B(NewStr("1")))
		assert.Equal(t, true, B(NewStr("true")))
		assert.Equal(t, true, B(NewStr("TRUE")))
		assert.Equal(t, false, B(NewStr("0")))
		assert.Equal(t, false, B(NewStr("false")))
		assert.Equal(t, false, B(NewStr("FALSE")))
		assert.Equal(t, false, B(NewStr("")))
	}

	// string
	{
		assert.Equal(t, true, B("1"))
		assert.Equal(t, true, B("true"))
		assert.Equal(t, true, B("TRUE"))
		assert.Equal(t, false, B("0"))
		assert.Equal(t, false, B("false"))
		assert.Equal(t, false, B("FALSE"))
		assert.Equal(t, false, B(""))
	}

	// uint
	{
		var test uint
		assert.Equal(t, true, B(uint(1)))
		assert.Equal(t, false, B(uint(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, true, B(uint8(3)))
		assert.Equal(t, false, B(uint8(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, true, B(uint16(3)))
		assert.Equal(t, false, B(uint16(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, true, B(uint32(3)))
		assert.Equal(t, false, B(uint32(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, true, B(uint64(3)))
		assert.Equal(t, false, B(uint64(0)))
		assert.Equal(t, false, B(&test))
		assert.Equal(t, false, B((*uint64)(nil)))
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
		assert.Equal(t, true, B("1"))
		assert.Equal(t, true, B("true"))
		assert.Equal(t, true, B("TRUE"))
		assert.Equal(t, false, B("0"))
		assert.Equal(t, false, B("false"))
		assert.Equal(t, false, B("FALSE"))
		assert.Equal(t, false, B(""))
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

// ToChar
//--------------------------------------------------------------------------------------------------
func ExampleToChar() {
	val := ToChar("v")
	fmt.Println(val)
	// Output: v
}

func TestToChar(t *testing.T) {

	// nil
	{
		assert.Equal(t, NewCharV(), ToChar(nil))
	}

	// bool
	{
		var test bool
		assert.Equal(t, '1', ToChar(true).O())
		assert.Equal(t, '0', ToChar(false).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*bool)(nil)))
	}

	// []byte
	{
		assert.Equal(t, 't', ToChar([]byte{0x74, 0x65, 0x73, 0x74}).O())
		assert.Equal(t, 'e', ToChar(&[]byte{0x65, 0x73, 0x74}).O())
		assert.Equal(t, NewCharV(), ToChar((*[]byte)(nil)))
	}

	// float32
	{
		var test float32
		assert.Equal(t, '7', ToChar(float32(7.22)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*float32)(nil)))
	}

	// float64
	{
		var test float64
		assert.Equal(t, '7', ToChar(float64(7.22)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*float64)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, '3', ToChar(3).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, '3', ToChar(int8(3)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, '3', ToChar(int16(3)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*int16)(nil)))
	}

	// int32 a.k.a rune
	{
		test := '3'
		assert.Equal(t, '3', ToChar('3').O())
		assert.Equal(t, '3', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, '3', ToChar(int64(3)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*int64)(nil)))
	}

	// Object
	{
		var val1 Object
		val2 := Object{3}
		assert.Equal(t, '3', ToChar(val2).G())
		assert.Equal(t, '3', ToChar(&val2).G())
		assert.Equal(t, NewCharV(), ToChar(val1))
		assert.Equal(t, NewCharV(), ToChar((*Object)(nil)))
	}

	// []rune
	{
		val1 := []rune{'1'}
		assert.Equal(t, '1', ToChar(val1).O())
		assert.Equal(t, '1', ToChar(&val1).O())
		assert.Equal(t, NewCharV(), ToChar((*[]rune)(nil)))
	}

	// Str
	{
		val1 := *NewStr("1")
		assert.Equal(t, '1', ToChar(val1).O())
		assert.Equal(t, '1', ToChar(&val1).O())
		assert.Equal(t, NewCharV(), ToChar((*[]Str)(nil)))
	}

	// string
	{
		val1 := "1"
		assert.Equal(t, '1', ToChar(val1).O())
		assert.Equal(t, '1', ToChar(&val1).O())
		assert.Equal(t, NewCharV(), ToChar(""))
		assert.Equal(t, NewCharV(), ToChar((*[]string)(nil)))

		// neg
		assert.Equal(t, '-', ToChar("-1").O())
		assert.Equal(t, '-', ToChar(-1).O())
	}

	// uint
	{
		var test uint
		assert.Equal(t, '3', ToChar(3).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*uint)(nil)))
	}

	// uint8 i.e. byte handle differently
	{
		var test uint8
		assert.Equal(t, 't', ToChar(uint8(0x74)).O())
		assert.Equal(t, rune(0), ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, '3', ToChar(uint16(3)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, '3', ToChar(uint32(3)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, '3', ToChar(uint64(3)).O())
		assert.Equal(t, '0', ToChar(&test).O())
		assert.Equal(t, NewCharV(), ToChar((*uint64)(nil)))
	}
}

// ToFloat32
//--------------------------------------------------------------------------------------------------
func ExampleToFloat32() {
	fmt.Println(ToFloat32("3.1"))
	// Output: 3.1
}

func TestToFloat32(t *testing.T) {

	// invalid
	{
		assert.Equal(t, float32(0), ToFloat32(nil))
		assert.Equal(t, float32(0), ToFloat32(&Object{}))
	}

	// bool
	{
		var test bool
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(1), ToFloat32(true))
		assert.Equal(t, float32(0), ToFloat32((*bool)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(3))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(int8(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(int16(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(int32(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(int64(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*int64)(nil)))
	}

	// string
	{
		var test string
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32("3"))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*string)(nil)))
		assert.Equal(t, float32(0), ToFloat32("0"))
		assert.Equal(t, float32(1), ToFloat32("true"))
		assert.Equal(t, float32(1), ToFloat32("TRUE"))
		assert.Equal(t, float32(0), ToFloat32("false"))
		assert.Equal(t, float32(0), ToFloat32("FALSE"))
		assert.Equal(t, float32(0), ToFloat32("bob"))
	}

	// uint
	{
		var test uint
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(uint(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(uint8(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(uint16(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(uint32(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, float32(0), ToFloat32(&test))
		assert.Equal(t, float32(3), ToFloat32(uint64(3)))
		assert.Equal(t, float32(0), ToFloat32(0))
		assert.Equal(t, float32(0), ToFloat32((*uint64)(nil)))
	}
}

// ToFloat32E
//--------------------------------------------------------------------------------------------------
func ExampleToFloat32E() {
	fmt.Println(ToFloat32E("1.1"))
	// Output: 1.1 <nil>
}

func TestToFloat32E(t *testing.T) {

	// invalid
	{
		val, err := ToFloat32E(nil)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to int", err.Error())
		assert.Equal(t, float32(0), val)
	}

	// bool
	{
		val, err := ToFloat32E(true)
		assert.Nil(t, err)
		assert.Equal(t, float32(1), val)

		var test bool
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// int
	{
		val, err := ToFloat32E(3)
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(0)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test int
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// int8
	{
		val, err := ToFloat32E(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test int8
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// int16
	{
		val, err := ToFloat32E(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test int16
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// int32
	{
		val, err := ToFloat32E(int32(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(int32(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test int32
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// int64
	{
		val, err := ToFloat32E(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test int64
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// Object
	{
		val, err := ToFloat32E(Object{3})
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(Object{0})
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*Object)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// Str
	{
		val, err := ToFloat32E(NewStr("2"))
		assert.Nil(t, err)
		assert.Equal(t, float32(2), val)

		var test Str
		val, err = ToFloat32E(&test)
		assert.NotNil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*Str)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// string
	{
		val, err := ToFloat32E("true")
		assert.Nil(t, err)
		assert.Equal(t, float32(1), val)

		val, err = ToFloat32E("TRUE")
		assert.Nil(t, err)
		assert.Equal(t, float32(1), val)

		val, err = ToFloat32E("false")
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E("FALSE")
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E("bob")
		assert.NotNil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E("3")
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E("0")
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test string
		val, err = ToFloat32E(&test)
		assert.NotNil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*string)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// uint
	{
		val, err := ToFloat32E(uint(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(0)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test uint
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// uint8
	{
		val, err := ToFloat32E(uint8(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(uint8(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test uint8
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// uint16
	{
		val, err := ToFloat32E(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test uint16
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// uint32
	{
		val, err := ToFloat32E(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test uint32
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}

	// uint64
	{
		val, err := ToFloat32E(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, float32(3), val)

		val, err = ToFloat32E(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		var test uint64
		val, err = ToFloat32E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)

		val, err = ToFloat32E((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float32(0), val)
	}
}

// ToFloat64
//--------------------------------------------------------------------------------------------------
func ExampleToFloat64() {
	fmt.Println(ToFloat64("3.1"))
	// Output: 3.1
}

func TestToFloat64(t *testing.T) {

	// invalid
	{
		assert.Equal(t, float64(0), ToFloat64(nil))
		assert.Equal(t, float64(0), ToFloat64(&Object{}))
	}

	// bool
	{
		var test bool
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(1), ToFloat64(true))
		assert.Equal(t, float64(0), ToFloat64((*bool)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(3))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(int8(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(int16(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(int32(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(int64(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*int64)(nil)))
	}

	// string
	{
		var test string
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64("3"))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*string)(nil)))
		assert.Equal(t, float64(0), ToFloat64("0"))
		assert.Equal(t, float64(1), ToFloat64("true"))
		assert.Equal(t, float64(1), ToFloat64("TRUE"))
		assert.Equal(t, float64(0), ToFloat64("false"))
		assert.Equal(t, float64(0), ToFloat64("FALSE"))
		assert.Equal(t, float64(0), ToFloat64("bob"))
	}

	// uint
	{
		var test uint
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(uint(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(uint8(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(uint16(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(uint32(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, float64(0), ToFloat64(&test))
		assert.Equal(t, float64(3), ToFloat64(uint64(3)))
		assert.Equal(t, float64(0), ToFloat64(0))
		assert.Equal(t, float64(0), ToFloat64((*uint64)(nil)))
	}
}

// ToFloat64E
//--------------------------------------------------------------------------------------------------
func ExampleToFloat64E() {
	fmt.Println(ToFloat64E("1.1"))
	// Output: 1.1 <nil>
}

func TestToFloat64E(t *testing.T) {

	// invalid
	{
		val, err := ToFloat64E(nil)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to int", err.Error())
		assert.Equal(t, float64(0), val)
	}

	// bool
	{
		val, err := ToFloat64E(true)
		assert.Nil(t, err)
		assert.Equal(t, float64(1), val)

		var test bool
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// int
	{
		val, err := ToFloat64E(3)
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(0)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test int
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// int8
	{
		val, err := ToFloat64E(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test int8
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// int16
	{
		val, err := ToFloat64E(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test int16
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// int32
	{
		val, err := ToFloat64E(int32(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(int32(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test int32
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// int64
	{
		val, err := ToFloat64E(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test int64
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// Object
	{
		val, err := ToFloat64E(Object{3})
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(Object{0})
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*Object)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// Str
	{
		val, err := ToFloat64E(NewStr("2"))
		assert.Nil(t, err)
		assert.Equal(t, float64(2), val)

		var test Str
		val, err = ToFloat64E(&test)
		assert.NotNil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*Str)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// string
	{
		val, err := ToFloat64E("true")
		assert.Nil(t, err)
		assert.Equal(t, float64(1), val)

		val, err = ToFloat64E("TRUE")
		assert.Nil(t, err)
		assert.Equal(t, float64(1), val)

		val, err = ToFloat64E("false")
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E("FALSE")
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E("bob")
		assert.NotNil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E("3")
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E("0")
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test string
		val, err = ToFloat64E(&test)
		assert.NotNil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*string)(nil))
		assert.NotNil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// uint
	{
		val, err := ToFloat64E(uint(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(0)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test uint
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// uint8
	{
		val, err := ToFloat64E(uint8(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(uint8(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test uint8
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// uint16
	{
		val, err := ToFloat64E(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test uint16
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// uint32
	{
		val, err := ToFloat64E(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test uint32
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
	}

	// uint64
	{
		val, err := ToFloat64E(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, float64(3), val)

		val, err = ToFloat64E(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		var test uint64
		val, err = ToFloat64E(&test)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)

		val, err = ToFloat64E((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, float64(0), val)
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

// ToInterSlice
//--------------------------------------------------------------------------------------------------
func ExampleToInterSlice() {
	fmt.Println(ToInterSlice("3").O())
	// Output: [3]
}

func TestToInterSlice(t *testing.T) {

	// Unknown object force the reflection path
	{
		type Bob struct{ Name string }
		assert.Equal(t, &InterSlice{Bob{"1"}, Bob{"2"}, Bob{"3"}}, ToInterSlice([]Bob{{"1"}, {"2"}, {"3"}}))
	}

	// invalid
	{
		assert.Equal(t, &InterSlice{}, ToInterSlice(nil))
		assert.Equal(t, &InterSlice{&Object{}}, ToInterSlice(&Object{}))
	}

	// bool
	{
		var test bool
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{true}, ToInterSlice(true))
		assert.Equal(t, &InterSlice{(*bool)(nil)}, ToInterSlice((*bool)(nil)))
		assert.Equal(t, &InterSlice{true, false, true}, ToInterSlice([]bool{true, false, true}))
	}

	// float32
	{
		var test float32
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{float32(1.2)}, ToInterSlice(float32(1.2)))
		assert.Equal(t, &InterSlice{(*float32)(nil)}, ToInterSlice((*float32)(nil)))
		assert.Equal(t, &InterSlice{float32(1), float32(2), float32(3)}, ToInterSlice([]float32{1, 2, 3}))
	}

	// float64
	{
		var test float64
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{float64(1.2)}, ToInterSlice(float64(1.2)))
		assert.Equal(t, &InterSlice{(*float64)(nil)}, ToInterSlice((*float64)(nil)))
		assert.Equal(t, &InterSlice{float64(1), float64(2), float64(3)}, ToInterSlice([]float64{1, 2, 3}))
	}

	// FloatSlice
	{
		assert.Equal(t, &InterSlice{}, ToInterSlice(&FloatSlice{}))
		assert.Equal(t, &InterSlice{float64(1.2)}, ToInterSlice(&FloatSlice{1.2}))
		assert.Equal(t, &InterSlice{}, ToInterSlice((*FloatSlice)(nil)))
		assert.Equal(t, &InterSlice{float64(1), float64(2), float64(3)}, ToInterSlice(FloatSlice{1, 2, 3}))
	}

	// []interface{}
	{
		assert.Equal(t, &InterSlice{1}, ToInterSlice([]interface{}{1}))
		assert.Equal(t, &InterSlice{"1"}, ToInterSlice([]interface{}{"1"}))
	}

	// int
	{
		var test int
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{3}, ToInterSlice(3))
		assert.Equal(t, &InterSlice{(*int)(nil)}, ToInterSlice((*int)(nil)))
		assert.Equal(t, &InterSlice{1, 2, 3}, ToInterSlice([]int{1, 2, 3}))
	}

	// int8
	{
		var test int8
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{int8(3)}, ToInterSlice(int8(3)))
		assert.Equal(t, &InterSlice{(*int8)(nil)}, ToInterSlice((*int8)(nil)))
		assert.Equal(t, &InterSlice{int8(1), int8(2), int8(3)}, ToInterSlice([]int8{1, 2, 3}))
	}

	// int16
	{
		var test int16
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{int16(3)}, ToInterSlice(int16(3)))
		assert.Equal(t, &InterSlice{(*int16)(nil)}, ToInterSlice((*int16)(nil)))
		assert.Equal(t, &InterSlice{int16(1), int16(2), int16(3)}, ToInterSlice([]int16{1, 2, 3}))
	}

	// int32
	{
		var test int32
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{int32(3)}, ToInterSlice(int32(3)))
		assert.Equal(t, &InterSlice{(*int32)(nil)}, ToInterSlice((*int32)(nil)))
		assert.Equal(t, &InterSlice{int32(1), int32(2), int32(3)}, ToInterSlice([]int32{1, 2, 3}))
	}

	// int64
	{
		var test int64
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{int64(3)}, ToInterSlice(int64(3)))
		assert.Equal(t, &InterSlice{(*int64)(nil)}, ToInterSlice((*int64)(nil)))
		assert.Equal(t, &InterSlice{int64(1), int64(2), int64(3)}, ToInterSlice([]int64{1, 2, 3}))
	}

	// IntSlice
	{
		assert.Equal(t, &InterSlice{}, ToInterSlice(&IntSlice{}))
		assert.Equal(t, &InterSlice{2}, ToInterSlice(&IntSlice{2}))
		assert.Equal(t, &InterSlice{}, ToInterSlice((*IntSlice)(nil)))
		assert.Equal(t, &InterSlice{1, 2, 3}, ToInterSlice(IntSlice{1, 2, 3}))
	}

	// InterSlice
	{
		assert.Equal(t, &InterSlice{1}, ToInterSlice(&InterSlice{1}))
		assert.Equal(t, &InterSlice{"1"}, ToInterSlice(&InterSlice{"1"}))
	}

	// string
	{
		var test string
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{"3"}, ToInterSlice("3"))
		assert.Equal(t, &InterSlice{(*string)(nil)}, ToInterSlice((*string)(nil)))
		assert.Equal(t, &InterSlice{"0"}, ToInterSlice("0"))
		assert.Equal(t, &InterSlice{"true"}, ToInterSlice("true"))
		assert.Equal(t, &InterSlice{"TRUE"}, ToInterSlice("TRUE"))
		assert.Equal(t, &InterSlice{"bob"}, ToInterSlice("bob"))
		assert.Equal(t, &InterSlice{"1", "2", "3"}, ToInterSlice([]string{"1", "2", "3"}))
	}

	// // StringMap
	// {
	// 	assert.Equal(t, &InterSlice{map[string]interface{}{}}, ToInterSlice(&StringMap{}))
	// 	assert.Equal(t, &InterSlice{map[string]interface{}{"2": "two"}}, ToInterSlice(&StringMap{"2": "two"}))
	// 	assert.Equal(t, &InterSlice{map[string]interface{}{}}, ToInterSlice((*StringMap)(nil)))
	// 	assert.Equal(t, &InterSlice{map[string]interface{}{"1": "one", "2": "two"}}, ToInterSlice(&StringMap{"1": "one", "2": "two"}))
	// }

	// StringSlice
	{
		assert.Equal(t, &InterSlice{}, ToInterSlice(&StringSlice{}))
		assert.Equal(t, &InterSlice{"2"}, ToInterSlice(&StringSlice{"2"}))
		assert.Equal(t, &InterSlice{}, ToInterSlice((*StringSlice)(nil)))
		assert.Equal(t, &InterSlice{"1", "2", "3"}, ToInterSlice(StringSlice{"1", "2", "3"}))
	}

	// uint
	{
		var test uint
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{uint(3)}, ToInterSlice(uint(3)))
		assert.Equal(t, &InterSlice{(*uint)(nil)}, ToInterSlice((*uint)(nil)))
		assert.Equal(t, &InterSlice{uint(1), uint(2), uint(3)}, ToInterSlice([]uint{1, 2, 3}))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{uint8(3)}, ToInterSlice(uint8(3)))
		assert.Equal(t, &InterSlice{(*uint8)(nil)}, ToInterSlice((*uint8)(nil)))
		assert.Equal(t, &InterSlice{uint8(1), uint8(2), uint8(3)}, ToInterSlice([]uint8{1, 2, 3}))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{uint16(3)}, ToInterSlice(uint16(3)))
		assert.Equal(t, &InterSlice{(*uint16)(nil)}, ToInterSlice((*uint16)(nil)))
		assert.Equal(t, &InterSlice{uint16(1), uint16(2), uint16(3)}, ToInterSlice([]uint16{1, 2, 3}))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{uint32(3)}, ToInterSlice(uint32(3)))
		assert.Equal(t, &InterSlice{(*uint32)(nil)}, ToInterSlice((*uint32)(nil)))
		assert.Equal(t, &InterSlice{uint32(1), uint32(2), uint32(3)}, ToInterSlice([]uint32{1, 2, 3}))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, &InterSlice{&test}, ToInterSlice(&test))
		assert.Equal(t, &InterSlice{uint64(3)}, ToInterSlice(uint64(3)))
		assert.Equal(t, &InterSlice{(*uint64)(nil)}, ToInterSlice((*uint64)(nil)))
		assert.Equal(t, &InterSlice{uint64(1), uint64(2), uint64(3)}, ToInterSlice([]uint64{1, 2, 3}))
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

// ToStr
//--------------------------------------------------------------------------------------------------
func ExampleToStr() {
	val := ToStr(true)
	fmt.Println(val)
	// Output: true
}

func TestToStr(t *testing.T) {

	// bool
	{
		assert.Equal(t, "", ToStr(nil).G())
	}

	// bool
	{
		var test bool
		assert.Equal(t, "true", ToStr(true).G())
		assert.Equal(t, "false", ToStr(false).G())
		assert.Equal(t, "false", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*bool)(nil)).G())
	}

	// []byte
	{
		assert.Equal(t, "test", ToStr([]byte{0x74, 0x65, 0x73, 0x74}).G())
		assert.Equal(t, "test", ToStr(&[]byte{0x74, 0x65, 0x73, 0x74}).G())
		assert.Equal(t, "", ToStr((*[]byte)(nil)).G())
	}

	// error
	{
		assert.Equal(t, "test", ToStr(errors.New("test")).G())
	}

	// float32
	{
		var test float32
		assert.Equal(t, "7.22", ToStr(float32(7.22)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*float32)(nil)).G())
	}

	// float64
	{
		var test float64
		assert.Equal(t, "7.22", ToStr(float64(7.22)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*float64)(nil)).G())
	}

	// int
	{
		var test int
		assert.Equal(t, "3", ToStr(3).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*int)(nil)).G())
	}

	// int8
	{
		var test int8
		assert.Equal(t, "3", ToStr(int8(3)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*int8)(nil)).G())
	}

	// int16
	{
		var test int16
		assert.Equal(t, "3", ToStr(int16(3)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*int16)(nil)).G())
	}

	// int32 i.e. rune
	{
		var test int32
		assert.Equal(t, "t", ToStr(int32('t')).G())
		assert.Equal(t, "", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*int32)(nil)).G())
	}

	// int64
	{
		var test int64
		assert.Equal(t, "3", ToStr(int64(3)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*int64)(nil)).G())
	}

	// ints
	{
		assert.Equal(t, "7", ToStr(int(7)).G())
		assert.Equal(t, "7", ToStr(int8(7)).G())
		assert.Equal(t, "7", ToStr(int16(7)).G())
		assert.Equal(t, "7", ToStr(uint32(7)).G())
		assert.Equal(t, "7", ToStr(int64(7)).G())
	}

	// nil
	{
		assert.Equal(t, "", ToStr(nil).G())
	}

	// template.CSS
	{
		var test template.CSS
		assert.Equal(t, "test", ToStr(template.CSS("test")).G())
		assert.Equal(t, "", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*template.CSS)(nil)).G())
	}

	// template.HTML
	{
		var test template.HTML
		assert.Equal(t, "test", ToStr(template.HTML("test")).G())
		assert.Equal(t, "", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*template.HTML)(nil)).G())
	}

	// template.HTMLAttr
	{
		var test template.HTMLAttr
		assert.Equal(t, "test", ToStr(template.HTMLAttr("test")).G())
		assert.Equal(t, "", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*template.HTMLAttr)(nil)).G())
	}

	// template.JS
	{
		var test template.JS
		assert.Equal(t, "test", ToStr(template.JS("test")).G())
		assert.Equal(t, "", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*template.JS)(nil)).G())
	}

	// template.URL
	{
		var test template.URL
		assert.Equal(t, "test", ToStr(template.URL("test")).G())
		assert.Equal(t, "", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*template.URL)(nil)).G())
	}

	// uint
	{
		var test uint
		assert.Equal(t, "3", ToStr(uint(3)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*uint)(nil)).G())
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, "3", ToStr(uint8('3')).G())
		assert.Equal(t, "", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*uint8)(nil)).G())
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, "3", ToStr(uint16(3)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*uint16)(nil)).G())
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, "3", ToStr(uint32(3)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*uint32)(nil)).G())
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, "3", ToStr(uint64(3)).G())
		assert.Equal(t, "0", ToStr(&test).G())
		assert.Equal(t, "", ToStr((*uint64)(nil)).G())
	}

	// uints
	{
		assert.Equal(t, "7", ToStr(uint(7)).G())
		assert.Equal(t, "7", ToStr(uint8('7')).G())
		assert.Equal(t, "7", ToStr(uint16(7)).G())
		assert.Equal(t, "7", ToStr(uint32(7)).G())
		assert.Equal(t, "7", ToStr(uint64(7)).G())
	}
}

// ToString
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
		assert.Equal(t, "", ToString((*bool)(nil)))
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
		assert.Equal(t, "", ToString((*float32)(nil)))
	}

	// float64
	{
		var test float64
		assert.Equal(t, "7.22", ToString(float64(7.22)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*float64)(nil)))
	}

	// int
	{
		var test int
		assert.Equal(t, "3", ToString(3))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*int)(nil)))
	}

	// int8
	{
		var test int8
		assert.Equal(t, "3", ToString(int8(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*int8)(nil)))
	}

	// int16
	{
		var test int16
		assert.Equal(t, "3", ToString(int16(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*int16)(nil)))
	}

	// int32
	{
		var test int32
		assert.Equal(t, "3", ToString(int32('3')))
		assert.Equal(t, "", ToString(&test))
		assert.Equal(t, "", ToString((*int32)(nil)))
	}

	// int64
	{
		var test int64
		assert.Equal(t, "3", ToString(int64(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*int64)(nil)))
	}

	// ints
	{
		assert.Equal(t, "7", ToString(int(7)))
		assert.Equal(t, "7", ToString(int8(7)))
		assert.Equal(t, "7", ToString(int16(7)))
		assert.Equal(t, "7", ToString(int32('7')))
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
		assert.Equal(t, "", ToString((*uint)(nil)))
	}

	// uint8
	{
		var test uint8
		assert.Equal(t, "3", ToString(uint8('3')))
		assert.Equal(t, "", ToString(&test))
		assert.Equal(t, "", ToString((*uint8)(nil)))
	}

	// uint16
	{
		var test uint16
		assert.Equal(t, "3", ToString(uint16(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*uint16)(nil)))
	}

	// uint32
	{
		var test uint32
		assert.Equal(t, "3", ToString(uint32(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*uint32)(nil)))
	}

	// uint64
	{
		var test uint64
		assert.Equal(t, "3", ToString(uint64(3)))
		assert.Equal(t, "0", ToString(&test))
		assert.Equal(t, "", ToString((*uint64)(nil)))
	}

	// uints
	{
		assert.Equal(t, "7", ToString(uint(7)))
		assert.Equal(t, "7", ToString(uint8('7')))
		assert.Equal(t, "7", ToString(uint16(7)))
		assert.Equal(t, "7", ToString(uint32(7)))
		assert.Equal(t, "7", ToString(uint64(7)))
	}
}

// ToStringMapE
//--------------------------------------------------------------------------------------------------
func ExampleToStringMapE() {
	fmt.Println(ToStringMapE(map[interface{}]interface{}{"1": "one"}))
	// Output: &map[1:one] <nil>
}

func TestToStringMapE(t *testing.T) {

	// map[string]uint64
	{
		val, err := ToStringMapE(map[string]uint64{"1": uint64(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", uint64(1)), val)
	}

	// map[string]uint32
	{
		val, err := ToStringMapE(map[string]uint32{"1": uint32(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", uint32(1)), val)
	}

	// map[string]uint16
	{
		val, err := ToStringMapE(map[string]uint16{"1": uint16(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", uint16(1)), val)
	}

	// map[string]uint8
	{
		val, err := ToStringMapE(map[string]uint8{"1": uint8(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", uint8(1)), val)
	}

	// map[string]uint
	{
		val, err := ToStringMapE(map[string]uint{"1": uint(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", uint(1)), val)
	}

	// map[string]int64
	{
		val, err := ToStringMapE(map[string]int64{"1": int64(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", int64(1)), val)
	}

	// map[string]int32
	{
		val, err := ToStringMapE(map[string]int32{"1": int32(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", int32(1)), val)
	}

	// map[string]int16
	{
		val, err := ToStringMapE(map[string]int16{"1": int16(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", int16(1)), val)
	}

	// map[string]int8
	{
		val, err := ToStringMapE(map[string]int8{"1": int8(1)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", int8(1)), val)
	}

	// map[string]int
	{
		val, err := ToStringMapE(map[string]int{"1": 1})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", 1), val)
	}

	// map[string]float64
	{
		val, err := ToStringMapE(map[string]float64{"1": float64(2.0)})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", float64(2.0)), val)
	}

	// map[string]bool
	{
		val, err := ToStringMapE(map[string]bool{"1": true})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", true), val)
	}

	// map[string]string
	{
		val, err := ToStringMapE(map[string]string{"1": "one"})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", "one"), val)
		assert.Equal(t, map[string]interface{}{"1": "one"}, val.G())

		val, err = ToStringMapE(&map[string]string{"1": "one", "2": "two"})
		assert.Nil(t, err)
		assert.Equal(t, M().Add("1", "one").Add("2", "two"), val)
		assert.Equal(t, map[string]interface{}{"1": "one", "2": "two"}, val.G())
	}

	// 	// map[string]interface{}
	// 	{
	// 		val, err := ToStringMapE(map[interface{}]interface{}{"1": "one"})
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("1": "one"}, val)
	// 		assert.Equal(t, map[string]interface{}{"1": "one"}, val.G())

	// 		val, err = ToStringMapE(&map[interface{}]interface{}{"1": "one"})
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("1": "one"}, val)
	// 		assert.Equal(t, map[string]interface{}{"1": "one"}, val.G())

	// 		val, err = ToStringMapE(map[string]interface{}{"1": "one"})
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("1": "one"}, val)
	// 		assert.Equal(t, map[string]interface{}{"1": "one"}, val.G())

	// 		val, err = ToStringMapE(&map[interface{}]interface{}{"1": "one"})
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("1": "one"}, val)
	// 		assert.Equal(t, map[string]interface{}{"1": "one"}, val.G())
	// 	}

	// 	// invalid
	// 	{
	// 		val, err := ToStringMapE(nil)
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add(}, val)

	// 		val, err = ToStringMapE(&TestObj{})
	// 		assert.Equal(t, "unable to convert type *n.TestObj to a StringMap", err.Error())
	// 		assert.Equal(t, M().Add(}, val)
	// 	}

	// 	// Object
	// 	{
	// 		val, err := ToStringMapE(Object{map[string]interface{}{"foo": "bar"}})
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("foo": "bar"}, val)

	// 		val, err = ToStringMapE(Object{map[string]interface{}{"foo1": "bar1", "foo2": "bar2"}})
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("foo1": "bar1", "foo2": "bar2"}, val)

	// 		val, err = ToStringMapE(&Object{})
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add(}, val)

	// 		val, err = ToStringMapE((*Object)(nil))
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add(}, val)
	// 	}

	// 	// []byte
	// 	{
	// 		yml := `foo:
	//   - 1
	//   - 2
	//   - 3
	// `
	// 		val, err := ToStringMapE([]byte(yml))
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("foo": []interface{}{int(1), int(2), int(3)}}, val)
	// 		assert.Equal(t, map[string]interface{}{"foo": []interface{}{int(1), int(2), int(3)}}, val.G())
	// 	}

	// 	// object
	// 	{

	// 	}

	// 	// string
	// 	{
	// 		yml := `foo:
	//   - 1
	//   - 2
	//   - 3
	// `
	// 		val, err := ToStringMapE(yml)
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("foo": []interface{}{int(1), int(2), int(3)}}, val)
	// 		assert.Equal(t, map[string]interface{}{"foo": []interface{}{int(1), int(2), int(3)}}, val.G())

	// 		// string map
	// 		yml = `foo1: bar1
	// foo2: bar2
	// foo3: bar3
	// `
	// 		val, err = ToStringMapE(yml)
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("foo1": "bar1", "foo2": "bar2", "foo3": "bar3"}, val)
	// 		assert.Equal(t, map[string]interface{}{"foo1": "bar1", "foo2": "bar2", "foo3": "bar3"}, val.G())

	// 		// string map nested
	// 		yml = `foo:
	//   - name: foo1
	//     val:
	//       bar1: 1
	//   - name: foo2
	//     val:
	//       bar2: 1
	// `
	// 		val, err = ToStringMapE(yml)
	// 		assert.Nil(t, err)
	// 		assert.Equal(t, M().Add("foo": []interface{}{map[string]interface{}{"name": "foo1", "val": map[string]interface{}{"bar1": int(1)}}, map[string]interface{}{"name": "foo2", "val": map[string]interface{}{"bar2": int(1)}}}}, val)
	// 		assert.Equal(t, map[string]interface{}{"foo": []interface{}{map[string]interface{}{"name": "foo1", "val": map[string]interface{}{"bar1": int(1)}}, map[string]interface{}{"name": "foo2", "val": map[string]interface{}{"bar2": int(1)}}}}, val.G())
	// 	}
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
		val, err := ToStringSliceE(int32('3'))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(int32('0'))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test int32
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{""}, val)

		val, err = ToStringSliceE((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{""}, val)
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

	// ISlice
	{
		val, err := ToStringSliceE(Slice([]string{"1", "2"}))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"1", "2"}, val)
		assert.Equal(t, []string{"1", "2"}, val.G())
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

	// // StringSlice
	// {
	// 	val, err := ToStringSliceE("true")

	// }

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
		val, err := ToStringSliceE(uint8('3'))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"3"}, val)

		val, err = ToStringSliceE(uint8('0'))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{"0"}, val)

		var test uint8
		val, err = ToStringSliceE(&test)
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{""}, val)

		val, err = ToStringSliceE((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, &StringSlice{""}, val)
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

// ToStrsE
//--------------------------------------------------------------------------------------------------
func ExampleToStrsE() {
	fmt.Println(ToStrsE("1"))
	// Output: [1] <nil>
}

func TestToStrsE(t *testing.T) {

	// invalid
	{
		val, err := ToStrsE(nil)
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStrsE(&TestObj{})
		assert.Equal(t, "unable to convert type n.TestObj to []int", err.Error())
		assert.Equal(t, []string{}, val)
	}

	// bool
	{
		val, err := ToStrsE(true)
		assert.Nil(t, err)
		assert.Equal(t, []string{"true"}, val)

		var test bool
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"false"}, val)

		val, err = ToStrsE((*bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"false"}, val)
	}

	// []bool
	{
		val, err := ToStrsE([]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, []string{"true", "false"}, val)

		val, err = ToStrsE(&[]bool{true, false})
		assert.Nil(t, err)
		assert.Equal(t, []string{"true", "false"}, val)

		one, two := true, false
		val, err = ToStrsE(&[]*bool{&one, &two})
		assert.Nil(t, err)
		assert.Equal(t, []string{"true", "false"}, val)

		val, err = ToStrsE((*[]bool)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// interface
	//-----------------------------------------
	{
		val, err := ToStrsE([]interface{}{"1", "2", "3"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "2", "3"}, val)

		val, err = ToStrsE(&[]interface{}{"1", "2", "3"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"1", "2", "3"}, val)
	}

	// int
	{
		val, err := ToStrsE(3)
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(0)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*int)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// int8
	{
		val, err := ToStrsE(int8(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(int8(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int8
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*int8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// int16
	{
		val, err := ToStrsE(int16(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(int16(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int16
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*int16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// int32
	{
		val, err := ToStrsE(int32('3'))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(int32('0'))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int32
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{""}, val)

		val, err = ToStrsE((*int32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{""}, val)
	}

	// int64
	{
		val, err := ToStrsE(int64(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(int64(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test int64
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*int64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// Object
	{
		val, err := ToStrsE(Object{"3"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(Object{"0"})
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE(&Object{})
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStrsE((*Object)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// Str
	{
		val, err := ToStrsE(NewStr("2"))
		assert.Nil(t, err)
		assert.Equal(t, []string{"2"}, val)

		val, err = ToStrsE(NewStr("bob"))
		assert.Nil(t, err)
		assert.Equal(t, []string{"bob"}, val)

		var test Str
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStrsE((*Str)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// string
	{
		val, err := ToStrsE("true")
		assert.Nil(t, err)
		assert.Equal(t, []string{"true"}, val)

		val, err = ToStrsE("TRUE")
		assert.Nil(t, err)
		assert.Equal(t, []string{"TRUE"}, val)

		val, err = ToStrsE("false")
		assert.Nil(t, err)
		assert.Equal(t, []string{"false"}, val)

		val, err = ToStrsE("FALSE")
		assert.Nil(t, err)
		assert.Equal(t, []string{"FALSE"}, val)

		val, err = ToStrsE("bob")
		assert.Nil(t, err)
		assert.Equal(t, []string{"bob"}, val)

		val, err = ToStrsE("3")
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE("0")
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test string
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)

		val, err = ToStrsE((*string)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{}, val)
	}

	// uint
	{
		val, err := ToStrsE(uint(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(0)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*uint)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// uint8
	{
		val, err := ToStrsE(uint8('3'))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(uint8('0'))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint8
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{""}, val)

		val, err = ToStrsE((*uint8)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{""}, val)
	}

	// uint16
	{
		val, err := ToStrsE(uint16(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(uint16(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint16
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*uint16)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// uint32
	{
		val, err := ToStrsE(uint32(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(uint32(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint32
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*uint32)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}

	// uint64
	{
		val, err := ToStrsE(uint64(3))
		assert.Nil(t, err)
		assert.Equal(t, []string{"3"}, val)

		val, err = ToStrsE(uint64(0))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		var test uint64
		val, err = ToStrsE(&test)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)

		val, err = ToStrsE((*uint64)(nil))
		assert.Nil(t, err)
		assert.Equal(t, []string{"0"}, val)
	}
}

// ToTime
//--------------------------------------------------------------------------------------------------
func ExampleToTime() {
	fmt.Println(ToTime("2008-01-10"))
	// Output: 2008-01-10 00:00:00 +0000 UTC
}

func TestTime(t *testing.T) {

	// 	time.RFC3339,  // "2006-01-02T15:04:05Z07:00" // ISO8601
	// 	time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	// 	time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST"
	// 	time.RFC822Z,  // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	// 	time.RFC822,   // "02 Jan 06 15:04 MST"
	// 	time.RFC850,   // "Monday, 02-Jan-06 15:04:05 MST"
	// 	time.ANSIC,    // "Mon Jan _2 15:04:05 2006"
	// 	time.UnixDate, // "Mon Jan _2 15:04:05 MST 2006"

	// 	time.RubyDate, // "Mon Jan 02 15:04:05 -0700 2006"
	{
	}

	// US: Month Day, Year
	{
		assert.Equal(t, time.Date(2008, 10, 2, 0, 0, 0, 0, time.UTC), ToTime("October 2, 2008"))
	}

	// 	"02 Jan 2006", // Day Month Year
	{
		assert.Equal(t, time.Date(2008, 12, 10, 0, 0, 0, 0, time.UTC), ToTime("10 Dec 2008"))
	}

	// 	"2006-01-02",  // Year-Month-Day
	{
		assert.Equal(t, time.Date(2008, 01, 10, 0, 0, 0, 0, time.UTC), ToTime("2008-01-10"))
	}

	// 	time.Kitchen,  // "3:04PM"
	{
		assert.Equal(t, time.Date(0, 1, 1, 23, 4, 0, 0, time.UTC), ToTime("11:04PM"))
	}

	// 	time.Stamp,      // "Jan _2 15:04:05"
	{
		assert.Equal(t, time.Date(0, 12, 1, 11, 4, 5, 0, time.UTC), ToTime("Dec 1 11:04:05"))
	}

	// 	time.StampMilli, // "Jan _2 15:04:05.000"
	{
		assert.Equal(t, time.Date(0, 12, 1, 11, 4, 5, 0, time.UTC), ToTime("Dec 1 11:04:05.000"))
	}

	// 	time.StampMicro, // "Jan _2 15:04:05.000000"
	{
		assert.Equal(t, time.Date(0, 12, 1, 11, 4, 5, 0, time.UTC), ToTime("Dec 1 11:04:05.000000"))
	}

	// 	time.StampNano,  // "Jan _2 15:04:05.000000000"
	{
		assert.Equal(t, time.Date(0, 12, 1, 11, 4, 5, 0, time.UTC), ToTime("Dec 1 11:04:05.000000000"))
	}

	// ints and uints
	{

		assert.Equal(t, time.Date(2019, 9, 18, 11, 58, 56, 0, time.UTC), ToTime("1568807936"))
		assert.Equal(t, time.Date(2019, 9, 18, 11, 58, 56, 0, time.UTC), ToTime(int(1568807936)))
		assert.Equal(t, time.Date(2019, 9, 18, 11, 58, 56, 0, time.UTC), ToTime(int32(1568807936)))
		assert.Equal(t, time.Date(2019, 9, 18, 11, 58, 56, 0, time.UTC), ToTime(int64(1568807936)))
		assert.Equal(t, time.Date(2019, 9, 18, 11, 58, 56, 0, time.UTC), ToTime(uint(1568807936)))
		assert.Equal(t, time.Date(2019, 9, 18, 11, 58, 56, 0, time.UTC), ToTime(uint32(1568807936)))
		assert.Equal(t, time.Date(2019, 9, 18, 11, 58, 56, 0, time.UTC), ToTime(uint64(1568807936)))
	}

	// time pointer
	{
		result := time.Date(2008, 01, 10, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, time.Date(2008, 01, 10, 0, 0, 0, 0, time.UTC), ToTime(&result))
	}

	// nil time
	{
		assert.Equal(t, time.Time{}, ToTime(nil))
	}
}
