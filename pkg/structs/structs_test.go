package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Bob string
	foo string
	bar int
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
	Init(foo)
	assert.Equal(t, "Bob", foo.Bob)
	assert.Equal(t, "foo", foo.foo)
}

func TestInit_Nil(t *testing.T) {
	defer func() {
		if msg := recover(); msg != nil {
			assert.Equal(t, "structs.New requires a non nil struct type not a invalid", msg)
		}
	}()
	Init((*Foo)(nil))
}

func TestInit_Panic(t *testing.T) {
	defer func() {
		if msg := recover(); msg != nil {
			assert.Equal(t, "structs.InitStrs requires a struct pointer type not a struct", msg)
		}
	}()
	Init(Foo{})
}

func TestFields(t *testing.T) {
	st := New(Foo{})
	assert.NotNil(t, st)
	fields := st.Fields()
	assert.Len(t, fields, 3)
	assert.Equal(t, "Bob", fields[0].Name)
	assert.Equal(t, "foo", fields[1].Name)
	assert.Equal(t, "bar", fields[2].Name)
}
