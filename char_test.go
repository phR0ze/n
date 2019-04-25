package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Equal
//--------------------------------------------------------------------------------------------------
func ExampleChar_Equal() {
	fmt.Println(NewChar('3').Equal('3'))
	// Output: true
}

func TestChar_Equal(t *testing.T) {
	// neg
	{
		// assert.Equal(t, true, NewChar('-3').Equal('3'))
		// assert.Equal(t, true, NewChar('3').Equal("3"))
		// assert.Equal(t, true, NewChar('3').Equal(3))
		// assert.Equal(t, false, NewChar('3').Equal('4'))
		// assert.Equal(t, false, NewChar('3').Equal("4"))
		// assert.Equal(t, false, NewChar('3').Equal(4))
	}

	// pos
	{
		assert.Equal(t, true, NewChar('3').Equal('3'))
		assert.Equal(t, true, NewChar('3').Equal("3"))
		assert.Equal(t, true, NewChar('3').Equal(3))
		assert.Equal(t, false, NewChar('3').Equal('4'))
		assert.Equal(t, false, NewChar('3').Equal("4"))
		assert.Equal(t, false, NewChar('3').Equal(4))
	}
}

// Less
//--------------------------------------------------------------------------------------------------
func ExampleChar_Less() {
	fmt.Println(NewChar('3').Less('2'))
	// Output: false
}

func TestChar_Less(t *testing.T) {
	assert.Equal(t, true, NewChar('3').Less('4'))
	assert.Equal(t, true, NewChar('3').Less("4"))
	assert.Equal(t, true, NewChar('3').Less(4))
	assert.Equal(t, false, NewChar('3').Less('2'))
	assert.Equal(t, false, NewChar('3').Less("2"))
	assert.Equal(t, false, NewChar('3').Less(2))
}

// String
//--------------------------------------------------------------------------------------------------
func ExampleChar_String() {
	char := NewChar('3')
	fmt.Println(char)
	// Output: 3
}

func TestChar_String(t *testing.T) {

	// nil
	{
		var char *Char
		assert.Equal(t, "", char.String())
		assert.Equal(t, "", NewChar("").String())

		c := Char(0)
		assert.Equal(t, "", c.String())
	}

	// bytes
	{
		assert.Equal(t, "t", NewChar(byte(0x74)).String())
		assert.Equal(t, "t", NewChar([]byte{0x74}).String())
	}

	// ints
	{
		assert.Equal(t, "2", NewChar(2).String())
		assert.Equal(t, "4", NewChar(int64(4)).String())
		assert.Equal(t, 't', NewChar(uint8(0x74)).O())
	}

	// rune
	{
		assert.Equal(t, "1", NewChar('1').String())
	}

	// string
	{
		assert.Equal(t, "1", NewChar("1").String())
	}
}
