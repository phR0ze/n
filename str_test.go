package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A
//--------------------------------------------------------------------------------------------------
func ExampleA() {
	str := A("test")
	fmt.Println(str)
	// Output: test
}

func TestStr_A(t *testing.T) {
	// Str
	{
		assert.Equal(t, "test", A(A("test")).A())
	}

	// string
	{
		assert.Equal(t, "test", A("test").A())
	}

	// runes
	{
		assert.Equal(t, "b", A('b').A())
		assert.Equal(t, "test", A([]rune("test")).A())
	}

	// bytes
	{
		assert.Equal(t, "test", A([]byte{0x74, 0x65, 0x73, 0x74}).A())
	}

	// ints
	{
		assert.Equal(t, "10", A(10).A())
	}
}

// NewStr
//--------------------------------------------------------------------------------------------------
func ExampleNewStr() {
	str := NewStr("test")
	fmt.Println(str)
	// Output: test
}

func TestStr_NewStr(t *testing.T) {
	assert.Equal(t, "test", NewStr("test").A())
}

// All
//--------------------------------------------------------------------------------------------------
func ExampleStr_All() {
	fmt.Println(A("foobar").All([]string{"foo"}))
	// Output: true
}

func TestStr_All(t *testing.T) {
	assert.True(t, A("test").All([]string{"tes", "est"}))
	assert.False(t, A("test").All([]string{"bob", "est"}))
}

// AllV
//--------------------------------------------------------------------------------------------------
func ExampleStr_AllV() {
	fmt.Println(A("foobar").AllV("foo"))
	// Output: true
}

func TestStr_AllV(t *testing.T) {
	assert.True(t, A("test").AllV("tes", "est"))
	assert.False(t, A("test").AllV("bob", "est"))
}

// At
//--------------------------------------------------------------------------------------------------
func TestStr_At(t *testing.T) {
	q := A("test")
	assert.Equal(t, 't', q.At(0))
	assert.Equal(t, 'e', q.At(1))
	assert.Equal(t, 's', q.At(2))
	assert.Equal(t, 't', q.At(3))
	assert.Equal(t, 't', q.At(-1))
	assert.Equal(t, 's', q.At(-2))
	assert.Equal(t, 'e', q.At(-3))
	assert.Equal(t, 't', q.At(-4))
	assert.Equal(t, int32(0), q.At(5))
}

// AtE
//--------------------------------------------------------------------------------------------------
func TestStr_AtE(t *testing.T) {
	q := A("test")
	r, err := q.AtE(0)
	assert.Nil(t, err)
	assert.Equal(t, 't', r)

	r, err = q.AtE(1)
	assert.Equal(t, 'e', r)
	assert.Nil(t, err)

	r, err = q.AtE(2)
	assert.Equal(t, 's', r)
	assert.Nil(t, err)

	r, err = q.AtE(5)
	assert.Equal(t, int32(0), r)
	assert.Equal(t, "index out of Str bounds", err.Error())

	// nil
	{
		r, err := (*Str)(nil).AtE(2)
		assert.Equal(t, int32(0), r)
		assert.Equal(t, "Str is nil", err.Error())
	}
}

func TestStr_B(t *testing.T) {
	// string
	{
		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A("test").B())
	}

	// runes
	{
		assert.Equal(t, []byte{0x74}, A('t').B())
		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A([]rune("test")).B())
	}

	// bytes
	{
		assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, A([]byte("test")).B())
	}

	// ints
	{
		assert.Equal(t, []byte{0x31, 0x30}, A(10).B())
	}
}

func TestStr_Contains(t *testing.T) {
	assert.True(t, A("test").Contains("tes"))
	assert.False(t, A("test").Contains("bob"))
}

// func TestStr_ContainsAny(t *testing.T) {
// 	assert.True(t, A("test").ContainsAny("tes"))
// 	assert.True(t, A("test").ContainsAny("f", "t"))
// 	assert.False(t, A("test").ContainsAny("f", "b"))
// }

// func TestStr_Empty(t *testing.T) {
// 	var empty *Str
// 	assert.True(t, A("").Empty())
// 	assert.True(t, empty.Empty())
// 	assert.True(t, A("  ").Empty())
// 	assert.True(t, A("\n").Empty())
// 	assert.True(t, A("\t").Empty())
// }

// func TestStr_HasAnyPrefix(t *testing.T) {
// 	assert.True(t, A("test").HasAnyPrefix("tes"))
// 	assert.True(t, A("test").HasAnyPrefix("bob", "tes"))
// 	assert.False(t, A("test").HasAnyPrefix("bob"))
// }

// func TestStr_HasAnySuffix(t *testing.T) {
// 	assert.True(t, A("test").HasAnySuffix("est"))
// 	assert.True(t, A("test").HasAnySuffix("bob", "est"))
// 	assert.False(t, A("test").HasAnySuffix("bob"))
// }

// func TestStr_HasPrefix(t *testing.T) {
// 	assert.True(t, A("test").HasPrefix("tes"))
// }

// func TestStr_HasSuffix(t *testing.T) {
// 	assert.True(t, A("test").HasSuffix("est"))
// }

// func TestStr_Len(t *testing.T) {
// 	assert.Equal(t, 0, A("").Len())
// 	assert.Equal(t, 4, A("test").Len())
// }

// func TestStr_Replace(t *testing.T) {
// 	assert.Equal(t, "tfoo", A("test").Replace("est", "foo").A())
// 	assert.Equal(t, "foost", A("test").Replace("te", "foo").A())
// 	assert.Equal(t, "foostfoo", A("testte").Replace("te", "foo").A())
// }

// func TestStr_SpaceLeft(t *testing.T) {
// 	assert.Equal(t, "", A("").SpaceLeft().A())
// 	assert.Equal(t, "  ", A("  bob").SpaceLeft().A())
// 	assert.Equal(t, "\n", A("\nbob").SpaceLeft().A())
// 	assert.Equal(t, " \t ", A(" \t bob").SpaceLeft().A())
// }

// func TestStr_SpaceRight(t *testing.T) {
// 	assert.Equal(t, "", A("").SpaceRight().A())
// 	assert.Equal(t, "  ", A("bob  ").SpaceRight().A())
// 	assert.Equal(t, "\n", A("bob\n").SpaceRight().A())
// 	assert.Equal(t, " \t ", A("bob \t ").SpaceRight().A())
// }

// // func TestStr_Split(t *testing.T) {
// // 	assert.Equal(t, []string{""}, A("").Split(".").O())
// // 	assert.Equal(t, []string{"1", "2"}, A("1.2").Split(".").O())
// // 	assert.Equal(t, []string{"1", "2"}, A("1.2").Split(".").S())
// // }

// // func TestStr_SplitOn(t *testing.T) {
// // 	{
// // 		first, second := A("").SplitOn(":")
// // 		assert.Equal(t, "", first)
// // 		assert.Equal(t, "", second)
// // 	}
// // 	{
// // 		first, second := A("foo").SplitOn(":")
// // 		assert.Equal(t, "foo", first)
// // 		assert.Equal(t, "", second)
// // 	}
// // 	{
// // 		first, second := A("foo:").SplitOn(":")
// // 		assert.Equal(t, "foo:", first)
// // 		assert.Equal(t, "", second)
// // 	}
// // 	{
// // 		first, second := A(":foo").SplitOn(":")
// // 		assert.Equal(t, ":", first)
// // 		assert.Equal(t, "foo", second)
// // 	}
// // 	{
// // 		first, second := A("foo: bar").SplitOn(":")
// // 		assert.Equal(t, "foo:", first)
// // 		assert.Equal(t, " bar", second)
// // 	}
// // 	{
// // 		first, second := A("foo: bar:frodo").SplitOn(":")
// // 		assert.Equal(t, "foo:", first)
// // 		assert.Equal(t, " bar:frodo", second)
// // 	}
// // }

// // func TestStrSpaceLeft(t *testing.T) {
// // 	assert.Equal(t, "", A("").SpaceLeft())
// // 	assert.Equal(t, "", A("bob").SpaceLeft())
// // 	assert.Equal(t, "  ", A("  bob").SpaceLeft())
// // 	assert.Equal(t, "    ", A("    bob").SpaceLeft())
// // 	assert.Equal(t, "\n", A("\nbob").SpaceLeft())
// // 	assert.Equal(t, "\t", A("\tbob").SpaceLeft())
// // }

// // func TestToASCII(t *testing.T) {
// // 	assert.Equal(t, "2 gspu data gspm data", A("2�gspu�data�gspm�data").ToASCII().A())
// // }

// // func TestStrTrimPrefix(t *testing.T) {
// // 	assert.Equal(t, "test]", A("[test]").TrimPrefix("[").A())
// // }

// // func TestStrTrimSpace(t *testing.T) {
// // 	{
// // 		//Left
// // 		assert.Equal(t, "bob", A("bob").TrimSpaceLeft().A())
// // 		assert.Equal(t, "bob", A("  bob").TrimSpaceLeft().A())
// // 		assert.Equal(t, "bob  ", A("  bob  ").TrimSpaceLeft().A())
// // 		assert.Equal(t, 3, A("  bob").TrimSpaceLeft().Len())
// // 	}
// // 	{
// // 		// Right
// // 		assert.Equal(t, "bob", A("bob").TrimSpaceRight().A())
// // 		assert.Equal(t, "bob", A("bob  ").TrimSpaceRight().A())
// // 		assert.Equal(t, "  bob", A("  bob  ").TrimSpaceRight().A())
// // 		assert.Equal(t, 3, A("bob  ").TrimSpaceRight().Len())
// // 	}
// // }

// // func TestStrTrimSuffix(t *testing.T) {
// // 	assert.Equal(t, "[test", A("[test]").TrimSuffix("]").A())
// // }

// // func TestYamlType(t *testing.T) {
// // 	{
// // 		// string
// // 		assert.Equal(t, "test", A("\"test\"").YamlType())
// // 		assert.Equal(t, "test", A("'test'").YamlType())
// // 		assert.Equal(t, "1", A("\"1\"").YamlType())
// // 		assert.Equal(t, "1", A("'1'").YamlType())
// // 	}
// // 	{
// // 		// int
// // 		assert.Equal(t, 1.0, A("1").YamlType())
// // 		assert.Equal(t, 0.0, A("0").YamlType())
// // 		assert.Equal(t, 25.0, A("25").YamlType())
// // 	}
// // 	{
// // 		// bool
// // 		assert.Equal(t, true, A("true").YamlType())
// // 		assert.Equal(t, false, A("false").YamlType())
// // 	}
// // 	{
// // 		// default
// // 		assert.Equal(t, "True", A("True").YamlType())
// // 	}
// // }
