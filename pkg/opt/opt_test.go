package opt

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStd struct {
	in      io.Reader
	out     io.Writer
	err     io.Writer
	home    string
	quiet   bool
	debug   bool
	dryrun  bool
	testing bool
}

func (s *TestStd) In(in ...io.Reader) io.Reader {
	if s != nil {
		if len(in) > 0 {
			s.in = in[0]
		}
		return s.in
	}
	return nil
}

func (s *TestStd) Out(out ...io.Writer) io.Writer {
	if s != nil {
		if len(out) > 0 {
			s.out = out[0]
		}
		return s.out
	}
	return nil
}

func (s *TestStd) Err(err ...io.Writer) io.Writer {
	if s != nil {
		if len(err) > 0 {
			s.err = err[0]
		}
		return s.err
	}
	return nil
}

func (s *TestStd) Home(home ...string) string {
	if s != nil {
		if len(home) > 0 {
			s.home = home[0]
		}
		return s.home
	}
	return ""
}

func (s *TestStd) Quiet(quiet ...bool) bool {
	if s != nil {
		if len(quiet) > 0 {
			s.quiet = quiet[0]
		}
		return s.quiet
	}
	return false
}

func (s *TestStd) Debug(debug ...bool) bool {
	if s != nil {
		if len(debug) > 0 {
			s.debug = debug[0]
		}
		return s.debug
	}
	return false
}

func (s *TestStd) DryRun(dryrun ...bool) bool {
	if s != nil {
		if len(dryrun) > 0 {
			s.dryrun = dryrun[0]
		}
		return s.dryrun
	}
	return false
}

func (s *TestStd) Testing(testing ...bool) bool {
	if s != nil {
		if len(testing) > 0 {
			s.testing = testing[0]
		}
		return s.testing
	}
	return false
}

func TestStdStreamInterface(t *testing.T) {
	var strm Std
	strm = &TestStd{}

	assert.Equal(t, strm.In(), nil)
	strm.In(os.Stdin)
	assert.Equal(t, strm.In(), os.Stdin)

	assert.Equal(t, strm.Out(), nil)
	strm.Out(os.Stdout)
	assert.Equal(t, strm.Out(), os.Stdout)

	assert.Equal(t, strm.Err(), nil)
	strm.Err(os.Stderr)
	assert.Equal(t, strm.Err(), os.Stderr)

	assert.Equal(t, strm.Home(), "")
	strm.Home("foo")
	assert.Equal(t, strm.Home(), "foo")

	assert.Equal(t, strm.Quiet(), false)
	strm.Quiet(true)
	assert.Equal(t, strm.Quiet(), true)

	assert.Equal(t, strm.Debug(), false)
	strm.Debug(true)
	assert.Equal(t, strm.Debug(), true)

	assert.Equal(t, strm.DryRun(), false)
	strm.DryRun(true)
	assert.Equal(t, strm.DryRun(), true)

	assert.Equal(t, strm.Testing(), false)
	strm.Testing(true)
	assert.Equal(t, strm.Testing(), true)
}

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

func TestCopy(t *testing.T) {

	// Ensure original isn't affected by copy changes
	{
		opts1 := []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}
		opts2 := Copy(opts1)
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}, opts1)
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}, opts2)
		Remove(&opts2, "1")
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}, opts1)
		assert.Equal(t, []*Opt{&Opt{"2", 2}, &Opt{"3", 3}}, opts2)
	}
}

func TestDefault(t *testing.T) {

	// option is nil
	{
		opts := []*Opt{}
		assert.False(t, Default(&opts, nil))
	}

	// slice is nil
	{
		var opts *[]*Opt
		assert.False(t, Default(opts, nil))
	}

	// add a new option
	{
		opts := []*Opt{}
		assert.True(t, Default(&opts, TestingOpt(true)))
	}
}

func TestOverwrite(t *testing.T) {

	// happy
	{
		opts := []*Opt{}
		result := Overwrite(&opts, &Opt{"1", 1})
		assert.Equal(t, &Opt{"1", 1}, result)
		result = Overwrite(&opts, &Opt{"2", 2})
		assert.Equal(t, &Opt{"2", 2}, result)
		result = Overwrite(&opts, &Opt{"3", 3})
		assert.Equal(t, &Opt{"3", 3}, result)
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}, opts)

		result = Overwrite(&opts, &Opt{"2", 5})
		assert.Equal(t, &Opt{"2", 5}, result)
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"2", 5}, &Opt{"3", 3}}, opts)
	}
}

func TestRemove(t *testing.T) {
	// nil opt in slice
	{
		// removing the first should force check on middle before iterating past
		opts := []*Opt{&Opt{"1", 1}, (*Opt)(nil), &Opt{"3", 3}}
		Remove(&opts, "1")
		assert.Equal(t, []*Opt{(*Opt)(nil), &Opt{"3", 3}}, opts)

		// Remove end of slice
		Remove(&opts, "3")
		assert.Equal(t, []*Opt{(*Opt)(nil)}, opts)
	}

	// nil opt
	{
		opts := []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}
		Remove(&opts, "")
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}, opts)
	}

	// nil opts
	{
		Remove((*[]*Opt)(nil), "2")
	}

	// remove the middle
	{
		opts := []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}
		Remove(&opts, "2")
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"3", 3}}, opts)
	}

	// remove the end
	{
		opts := []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}
		Remove(&opts, "3")
		assert.Equal(t, []*Opt{&Opt{"1", 1}, &Opt{"2", 2}}, opts)
	}

	// remove the begining
	{
		opts := []*Opt{&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}}
		Remove(&opts, "1")
		assert.Equal(t, []*Opt{&Opt{"2", 2}, &Opt{"3", 3}}, opts)
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
		assert.True(t, Default(&opts, InOpt(buf)))
		assert.Equal(t, buf, GetInOpt(opts))
	}

	// default
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stdin, GetInOpt(opts))
		assert.False(t, InOptExists(opts))
		assert.Equal(t, buf, DefaultInOpt(opts, buf))
		assert.Equal(t, os.Stdin, GetInOpt(opts))
		assert.True(t, Default(&opts, InOpt(buf)))
		assert.Equal(t, buf, GetInOpt(opts))
		assert.True(t, InOptExists(opts))
		assert.Equal(t, buf, DefaultInOpt(opts, os.Stdin))
	}

	// overwrite
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, InOpt(buf)))
		assert.Equal(t, buf, GetInOpt(opts))
		assert.Equal(t, os.Stdin, OverwriteInOpt(&opts, os.Stdin))
		assert.Equal(t, os.Stdin, GetInOpt(opts))
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
		assert.True(t, Default(&opts, OutOpt(buf)))
		assert.Equal(t, buf, GetOutOpt(opts))
	}

	// default
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stdout, GetOutOpt(opts))
		assert.False(t, OutOptExists(opts))
		assert.Equal(t, buf, DefaultOutOpt(opts, buf))
		assert.Equal(t, os.Stdout, GetOutOpt(opts))
		assert.True(t, Default(&opts, OutOpt(buf)))
		assert.Equal(t, buf, GetOutOpt(opts))
		assert.True(t, OutOptExists(opts))
		assert.Equal(t, buf, DefaultOutOpt(opts, os.Stdout))
	}

	// overwrite
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, OutOpt(buf)))
		assert.Equal(t, buf, GetOutOpt(opts))
		assert.Equal(t, os.Stdout, OverwriteOutOpt(&opts, os.Stdout))
		assert.Equal(t, os.Stdout, GetOutOpt(opts))
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
		assert.True(t, Default(&opts, ErrOpt(buf)))
		assert.Equal(t, buf, GetErrOpt(opts))
	}

	// default
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, os.Stderr, GetErrOpt(opts))
		assert.False(t, ErrOptExists(opts))
		assert.Equal(t, buf, DefaultErrOpt(opts, buf))
		assert.Equal(t, os.Stderr, GetErrOpt(opts))
		assert.True(t, Default(&opts, ErrOpt(buf)))
		assert.Equal(t, buf, GetErrOpt(opts))
		assert.True(t, ErrOptExists(opts))
		assert.Equal(t, buf, DefaultErrOpt(opts, os.Stderr))
	}

	// overwrite
	{
		buf := &bytes.Buffer{}
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, ErrOpt(buf)))
		assert.Equal(t, buf, GetErrOpt(opts))
		assert.Equal(t, os.Stderr, OverwriteErrOpt(&opts, os.Stderr))
		assert.Equal(t, os.Stderr, GetErrOpt(opts))
	}
}

// Home
// -------------------------------------------------------------------------------------------------

func TestHomeOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.Equal(t, "", GetHomeOpt(opts))
		assert.True(t, Default(&opts, HomeOpt("foobar")))
		assert.Equal(t, "foobar", GetHomeOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.Equal(t, "", GetHomeOpt(opts))
		assert.False(t, HomeOptExists(opts))
		assert.Equal(t, "foobar", DefaultHomeOpt(opts, "foobar"))
		assert.Equal(t, "", GetHomeOpt(opts))
		assert.True(t, Default(&opts, HomeOpt("foobar")))
		assert.Equal(t, "foobar", GetHomeOpt(opts))
		assert.True(t, HomeOptExists(opts))
		assert.Equal(t, "foobar", DefaultHomeOpt(opts, "blah"))
	}

	// overwrite
	{
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, HomeOpt("foo")))
		assert.Equal(t, "foo", GetHomeOpt(opts))
		assert.Equal(t, "bar", OverwriteHomeOpt(&opts, "bar"))
		assert.Equal(t, "bar", GetHomeOpt(opts))
	}
}

// Quiet
// -------------------------------------------------------------------------------------------------

func TestQuietOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetQuietOpt(opts))
		assert.True(t, Default(&opts, QuietOpt(true)))
		assert.True(t, GetQuietOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetQuietOpt(opts))
		assert.False(t, QuietOptExists(opts))
		assert.True(t, DefaultQuietOpt(opts, true))
		assert.False(t, GetQuietOpt(opts))
		assert.True(t, Default(&opts, QuietOpt(true)))
		assert.True(t, GetQuietOpt(opts))
		assert.True(t, QuietOptExists(opts))
		assert.True(t, DefaultQuietOpt(opts, false))
	}

	// overwrite
	{
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, QuietOpt(true)))
		assert.Equal(t, true, GetQuietOpt(opts))
		assert.Equal(t, false, OverwriteQuietOpt(&opts, false))
		assert.Equal(t, false, GetQuietOpt(opts))
	}
}

// Debug
// -------------------------------------------------------------------------------------------------

func TestDebugOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetDebugOpt(opts))
		assert.True(t, Default(&opts, DebugOpt(true)))
		assert.True(t, GetDebugOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetDebugOpt(opts))
		assert.False(t, DebugOptExists(opts))
		assert.True(t, DefaultDebugOpt(opts, true))
		assert.False(t, GetDebugOpt(opts))
		assert.True(t, Default(&opts, DebugOpt(true)))
		assert.True(t, GetDebugOpt(opts))
		assert.True(t, DebugOptExists(opts))
		assert.True(t, DefaultDebugOpt(opts, false))
	}

	// overwrite
	{
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, DebugOpt(true)))
		assert.Equal(t, true, GetDebugOpt(opts))
		assert.Equal(t, false, OverwriteDebugOpt(&opts, false))
		assert.Equal(t, false, GetDebugOpt(opts))
	}
}

// DryRun
// -------------------------------------------------------------------------------------------------

func TestDryRunOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetDryRunOpt(opts))
		assert.True(t, Default(&opts, DryRunOpt(true)))
		assert.True(t, GetDryRunOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetDryRunOpt(opts))
		assert.False(t, DryRunOptExists(opts))
		assert.True(t, DefaultDryRunOpt(opts, true))
		assert.False(t, GetDryRunOpt(opts))
		assert.True(t, Default(&opts, DryRunOpt(true)))
		assert.True(t, GetDryRunOpt(opts))
		assert.True(t, DryRunOptExists(opts))
		assert.True(t, DefaultDryRunOpt(opts, false))
	}

	// overwrite
	{
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, DryRunOpt(true)))
		assert.Equal(t, true, GetDryRunOpt(opts))
		assert.Equal(t, false, OverwriteDryRunOpt(&opts, false))
		assert.Equal(t, false, GetDryRunOpt(opts))
	}
}

// Testing
// -------------------------------------------------------------------------------------------------

func TestTestingOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetTestingOpt(opts))
		assert.True(t, Default(&opts, TestingOpt(true)))
		assert.True(t, GetTestingOpt(opts))
	}

	// default/exists
	{
		opts := []*Opt{}
		assert.False(t, GetTestingOpt(opts))
		assert.False(t, TestingOptExists(opts))
		assert.True(t, DefaultTestingOpt(opts, true))
		assert.False(t, GetTestingOpt(opts))
		assert.True(t, Default(&opts, TestingOpt(true)))
		assert.True(t, GetTestingOpt(opts))
		assert.True(t, TestingOptExists(opts))
		assert.True(t, DefaultTestingOpt(opts, false))
	}

	// overwrite
	{
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, TestingOpt(true)))
		assert.Equal(t, true, GetTestingOpt(opts))
		assert.Equal(t, false, OverwriteTestingOpt(&opts, false))
		assert.Equal(t, false, GetTestingOpt(opts))
	}
}
