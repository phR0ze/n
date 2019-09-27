package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Bob string
	foo string
	i   int
}

func TestNew(t *testing.T) {

	// pointer
	{
		st := New(&Foo{})
		assert.NotNil(t, st)
	}

	// non pointer
	{
		st := New(Foo{})
		assert.NotNil(t, st)
	}
}

func TestNew_Panic(t *testing.T) {
	defer func() {
		if msg := recover(); msg != nil {
			assert.Equal(t, "structs.New requires a non nil struct type not a string", msg)
		}
	}()
	New("test")
}

func TestInitStrs(t *testing.T) {
	foo := &Foo{}
	assert.Equal(t, "", foo.Bob)
	assert.Equal(t, "", foo.foo)
	ref := Init(foo)
	assert.Equal(t, "Bob", foo.Bob)
	assert.Equal(t, "foo", foo.foo)

	assert.Equal(t, &Foo{Bob: "Bob", foo: "foo"}, ref.(*Foo))
}

func TestInit_Nil(t *testing.T) {
	defer func() {
		if msg := recover(); msg != nil {
			assert.Equal(t, "structs.New requires a non nil struct type not a invalid", msg)
		}
	}()
	Init((*Foo)(nil))
}

func TestFields(t *testing.T) {
	st := New(Foo{})
	assert.NotNil(t, st)
	fields := st.Fields()
	assert.Len(t, fields, 3)
	assert.Equal(t, "Bob", fields[0].Name)
	assert.Equal(t, "foo", fields[1].Name)
	assert.Equal(t, "i", fields[2].Name)
}

func TestSetFieldByIndex(t *testing.T) {

	// int64 - not set as not same type
	{
		foo := &Foo{}
		st := New(foo)
		assert.NotNil(t, st)
		assert.Equal(t, 0, foo.i)
		st.SetFieldByIndex(2, int64(1))
		assert.Equal(t, 0, foo.i)
	}

	// int - set as same type
	{
		foo := &Foo{}
		st := New(foo)
		assert.NotNil(t, st)
		assert.Equal(t, 0, foo.i)
		st.SetFieldByIndex(2, 1)
		assert.Equal(t, 1, foo.i)
	}
}
