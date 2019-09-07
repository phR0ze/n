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
		assert.Equal(t, buf, DefaultInOpt((*[]*Opt)(nil), buf))
		assert.Equal(t, buf, DefaultInOpt(&opts, buf))
		assert.Equal(t, os.Stdin, GetInOpt(opts))
		assert.True(t, Add(&opts, InOpt(buf)))
		assert.Equal(t, buf, GetInOpt(opts))
		assert.Equal(t, buf, DefaultInOpt(&opts, os.Stdin))
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
		assert.Equal(t, buf, DefaultOutOpt((*[]*Opt)(nil), buf))
		assert.Equal(t, buf, DefaultOutOpt(&opts, buf))
		assert.Equal(t, os.Stdout, GetOutOpt(opts))
		assert.True(t, Add(&opts, OutOpt(buf)))
		assert.Equal(t, buf, GetOutOpt(opts))
		assert.Equal(t, buf, DefaultOutOpt(&opts, os.Stdout))
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
		assert.Equal(t, buf, DefaultErrOpt((*[]*Opt)(nil), buf))
		assert.Equal(t, buf, DefaultErrOpt(&opts, buf))
		assert.Equal(t, os.Stderr, GetErrOpt(opts))
		assert.True(t, Add(&opts, ErrOpt(buf)))
		assert.Equal(t, buf, GetErrOpt(opts))
		assert.Equal(t, buf, DefaultErrOpt(&opts, os.Stderr))
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
		assert.Equal(t, "foobar", DefaultHomeOpt((*[]*Opt)(nil), "foobar"))
		assert.Equal(t, "foobar", DefaultHomeOpt(&opts, "foobar"))
		assert.Equal(t, "", GetHomeOpt(opts))
		assert.True(t, Add(&opts, HomeOpt("foobar")))
		assert.Equal(t, "foobar", GetHomeOpt(opts))
		assert.Equal(t, "foobar", DefaultHomeOpt(&opts, "blah"))
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
		assert.True(t, DefaultQuietOpt((*[]*Opt)(nil), true))
		assert.True(t, DefaultQuietOpt(&opts, true))
		assert.False(t, GetQuietOpt(opts))
		assert.True(t, Add(&opts, QuietOpt(true)))
		assert.True(t, GetQuietOpt(opts))
		assert.True(t, DefaultQuietOpt(&opts, false))
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
		assert.True(t, DefaultDebugOpt((*[]*Opt)(nil), true))
		assert.True(t, DefaultDebugOpt(&opts, true))
		assert.False(t, GetDebugOpt(opts))
		assert.True(t, Add(&opts, DebugOpt(true)))
		assert.True(t, GetDebugOpt(opts))
		assert.True(t, DefaultDebugOpt(&opts, false))
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
		assert.True(t, DefaultDryRunOpt((*[]*Opt)(nil), true))
		assert.True(t, DefaultDryRunOpt(&opts, true))
		assert.False(t, GetDryRunOpt(opts))
		assert.True(t, Add(&opts, DryRunOpt(true)))
		assert.True(t, GetDryRunOpt(opts))
		assert.True(t, DefaultDryRunOpt(&opts, false))
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
		assert.True(t, DefaultTestingOpt((*[]*Opt)(nil), true))
		assert.True(t, DefaultTestingOpt(&opts, true))
		assert.False(t, GetTestingOpt(opts))
		assert.True(t, Add(&opts, TestingOpt(true)))
		assert.True(t, GetTestingOpt(opts))
		assert.True(t, DefaultTestingOpt(&opts, false))
	}
}
