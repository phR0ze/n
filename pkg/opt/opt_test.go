package opt

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	// option is nil
	{
		opts := []*Opt{}
		assert.False(t, Add(&opts, nil))
	}

	// slice is nil
	{
		var opts *[]*Opt
		assert.False(t, Add(opts, nil))
	}

	// add a new option
	{
		opts := []*Opt{}
		assert.True(t, Add(&opts, TestingOpt(true)))
	}
}

// In
// -------------------------------------------------------------------------------------------------

func TestInOpt(t *testing.T) {

	// create
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stdin, GetInOpt(opts))
		assert.True(t, Add(&opts, InOpt(buf)))
		assert.Equal(t, buf, GetInOpt(opts))
	}

	// default
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stdin, GetInOpt(opts))
		DefaultInOpt(&opts, buf)
		assert.Equal(t, buf, GetInOpt(opts))
	}
}

// Out
// -------------------------------------------------------------------------------------------------

func TestOutOpt(t *testing.T) {

	// create
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stdout, GetOutOpt(opts))
		assert.True(t, Add(&opts, OutOpt(buf)))
		assert.Equal(t, buf, GetOutOpt(opts))
	}

	// default
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stdout, GetOutOpt(opts))
		DefaultOutOpt(&opts, buf)
		assert.Equal(t, buf, GetOutOpt(opts))
	}
}

// Err
// -------------------------------------------------------------------------------------------------

func TestErrOpt(t *testing.T) {

	// create
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stderr, GetErrOpt(opts))
		assert.True(t, Add(&opts, ErrOpt(buf)))
		assert.Equal(t, buf, GetErrOpt(opts))
	}

	// default
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stderr, GetErrOpt(opts))
		DefaultErrOpt(&opts, buf)
		assert.Equal(t, buf, GetErrOpt(opts))
	}
}

// Home
// -------------------------------------------------------------------------------------------------

func TestHomeOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.Equal(t, "", GetHomeOpt(opts))
		assert.True(t, Add(&opts, HomeOpt("foobar")))
		assert.Equal(t, "foobar", GetHomeOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.Equal(t, "", GetHomeOpt(opts))
		DefaultHomeOpt(&opts, "foobar")
		assert.Equal(t, "foobar", GetHomeOpt(opts))
	}
}

// Quiet
// -------------------------------------------------------------------------------------------------

func TestQuietOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetQuietOpt(opts))
		assert.True(t, Add(&opts, QuietOpt(true)))
		assert.True(t, GetQuietOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetQuietOpt(opts))
		DefaultQuietOpt(&opts, true)
		assert.True(t, GetQuietOpt(opts))
	}
}

// Debug
// -------------------------------------------------------------------------------------------------

func TestDebugOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetDebugOpt(opts))
		assert.True(t, Add(&opts, DebugOpt(true)))
		assert.True(t, GetDebugOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetDebugOpt(opts))
		DefaultDebugOpt(&opts, true)
		assert.True(t, GetDebugOpt(opts))
	}
}

// DryRun
// -------------------------------------------------------------------------------------------------

func TestDryRunOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetDryRunOpt(opts))
		assert.True(t, Add(&opts, DryRunOpt(true)))
		assert.True(t, GetDryRunOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetDryRunOpt(opts))
		DefaultDryRunOpt(&opts, true)
		assert.True(t, GetDryRunOpt(opts))
	}
}

// Testing
// -------------------------------------------------------------------------------------------------

func TestTestingOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetTestingOpt(opts))
		assert.True(t, Add(&opts, TestingOpt(true)))
		assert.True(t, GetTestingOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetTestingOpt(opts))
		DefaultTestingOpt(&opts, true)
		assert.True(t, GetTestingOpt(opts))
	}
}
