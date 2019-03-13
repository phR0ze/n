package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQStr_Type(t *testing.T) {
	assert.Equal(t, QStrType, A("test").Type())
}

func TestQStr_A(t *testing.T) {
	// QStr
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

func TestQStr_B(t *testing.T) {
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

func TestQStr_Q(t *testing.T) {
	assert.Equal(t, QStrType, A("test").Q().Type())
}

func TestQStr_At(t *testing.T) {
	q := A("test")
	assert.Equal(t, 't', q.At(0))
	assert.Equal(t, 'e', q.At(1))
	assert.Equal(t, 's', q.At(2))
	assert.Equal(t, 't', q.At(3))
	assert.Equal(t, 't', q.At(-1))
	assert.Equal(t, 's', q.At(-2))
	assert.Equal(t, 'e', q.At(-3))
	assert.Equal(t, 't', q.At(-4))
	defer func() {
		err := recover()
		assert.Equal(t, "Index out of QStr bounds", err)
	}()
	assert.Equal(t, 't', q.At(5))
}

// func TestStrContains(t *testing.T) {
// 	assert.True(t, A("test").Contains("tes"))
// 	assert.False(t, A("test").Contains("bob"))
// }

// func TestStrContainsAny(t *testing.T) {
// 	assert.True(t, A("test").ContainsAny("tes"))
// 	assert.True(t, A("test").ContainsAny("f", "t"))
// 	assert.False(t, A("test").ContainsAny("f", "b"))
// }

// func TestStrHasAnyPrefix(t *testing.T) {
// 	assert.True(t, A("test").HasAnyPrefix("tes"))
// 	assert.True(t, A("test").HasAnyPrefix("bob", "tes"))
// 	assert.False(t, A("test").HasAnyPrefix("bob"))
// }

// func TestStrHasAnySuffix(t *testing.T) {
// 	assert.True(t, A("test").HasAnySuffix("est"))
// 	assert.True(t, A("test").HasAnySuffix("bob", "est"))
// 	assert.False(t, A("test").HasAnySuffix("bob"))
// }

// func TestStrHasPrefix(t *testing.T) {
// 	assert.True(t, A("test").HasPrefix("tes"))
// }

// func TestStrHasSuffix(t *testing.T) {
// 	assert.True(t, A("test").HasSuffix("est"))
// }

// func TestStrSplit(t *testing.T) {
// 	assert.Equal(t, []string{"1", "2"}, A("1.2").Split(".").S())
// }

// func TestStrSplitOn(t *testing.T) {
// 	{
// 		first, second := A("").SplitOn(":")
// 		assert.Equal(t, "", first)
// 		assert.Equal(t, "", second)
// 	}
// 	{
// 		first, second := A("foo").SplitOn(":")
// 		assert.Equal(t, "foo", first)
// 		assert.Equal(t, "", second)
// 	}
// 	{
// 		first, second := A("foo:").SplitOn(":")
// 		assert.Equal(t, "foo:", first)
// 		assert.Equal(t, "", second)
// 	}
// 	{
// 		first, second := A(":foo").SplitOn(":")
// 		assert.Equal(t, ":", first)
// 		assert.Equal(t, "foo", second)
// 	}
// 	{
// 		first, second := A("foo: bar").SplitOn(":")
// 		assert.Equal(t, "foo:", first)
// 		assert.Equal(t, " bar", second)
// 	}
// 	{
// 		first, second := A("foo: bar:frodo").SplitOn(":")
// 		assert.Equal(t, "foo:", first)
// 		assert.Equal(t, " bar:frodo", second)
// 	}
// }

// func TestStrSpaceLeft(t *testing.T) {
// 	assert.Equal(t, "", A("").SpaceLeft())
// 	assert.Equal(t, "", A("bob").SpaceLeft())
// 	assert.Equal(t, "  ", A("  bob").SpaceLeft())
// 	assert.Equal(t, "    ", A("    bob").SpaceLeft())
// 	assert.Equal(t, "\n", A("\nbob").SpaceLeft())
// 	assert.Equal(t, "\t", A("\tbob").SpaceLeft())
// }

// func TestToASCII(t *testing.T) {
// 	assert.Equal(t, "2 gspu data gspm data", A("2�gspu�data�gspm�data").ToASCII().A())
// }

// func TestStrTrimPrefix(t *testing.T) {
// 	assert.Equal(t, "test]", A("[test]").TrimPrefix("[").A())
// }

// func TestStrTrimSpace(t *testing.T) {
// 	{
// 		//Left
// 		assert.Equal(t, "bob", A("bob").TrimSpaceLeft().A())
// 		assert.Equal(t, "bob", A("  bob").TrimSpaceLeft().A())
// 		assert.Equal(t, "bob  ", A("  bob  ").TrimSpaceLeft().A())
// 		assert.Equal(t, 3, A("  bob").TrimSpaceLeft().Len())
// 	}
// 	{
// 		// Right
// 		assert.Equal(t, "bob", A("bob").TrimSpaceRight().A())
// 		assert.Equal(t, "bob", A("bob  ").TrimSpaceRight().A())
// 		assert.Equal(t, "  bob", A("  bob  ").TrimSpaceRight().A())
// 		assert.Equal(t, 3, A("bob  ").TrimSpaceRight().Len())
// 	}
// }

// func TestStrTrimSuffix(t *testing.T) {
// 	assert.Equal(t, "[test", A("[test]").TrimSuffix("]").A())
// }

// func TestYamlType(t *testing.T) {
// 	{
// 		// string
// 		assert.Equal(t, "test", A("\"test\"").YamlType())
// 		assert.Equal(t, "test", A("'test'").YamlType())
// 		assert.Equal(t, "1", A("\"1\"").YamlType())
// 		assert.Equal(t, "1", A("'1'").YamlType())
// 	}
// 	{
// 		// int
// 		assert.Equal(t, 1.0, A("1").YamlType())
// 		assert.Equal(t, 0.0, A("0").YamlType())
// 		assert.Equal(t, 25.0, A("25").YamlType())
// 	}
// 	{
// 		// bool
// 		assert.Equal(t, true, A("true").YamlType())
// 		assert.Equal(t, false, A("false").YamlType())
// 	}
// 	{
// 		// default
// 		assert.Equal(t, "True", A("True").YamlType())
// 	}
// }
