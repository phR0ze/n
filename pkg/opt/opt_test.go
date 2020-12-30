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

func (s *TestStd) Dryrun(dryrun ...bool) bool {
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

	assert.Equal(t, strm.Dryrun(), false)
	strm.Dryrun(true)
	assert.Equal(t, strm.Dryrun(), true)

	assert.Equal(t, strm.Testing(), false)
	strm.Testing(true)
	assert.Equal(t, strm.Testing(), true)
}

func TestNew(t *testing.T) {
	t.Run("New creates a new empty Opts", func(t *testing.T) {
		opts := NewOpts()
		assert.NotNil(t, opts)
		assert.Equal(t, 0, opts.Len())
	})

	t.Run("New with an opt creates a new Opts with one Opt", func(t *testing.T) {
		opt := TestingOpt(true)
		opts := NewOpts(opt)
		assert.NotNil(t, opts)
		assert.Equal(t, 1, opts.Len())
		assert.Equal(t, opt, opts.Get(TestingOptKey))
	})

	t.Run("New with more than one opt creates a new Opts with the same", func(t *testing.T) {
		opt1 := DebugOpt(true)
		opt2 := TestingOpt(true)
		opts := NewOpts(opt1, opt2)
		assert.NotNil(t, opts)
		assert.Equal(t, 2, opts.Len())
		assert.Equal(t, opt1, opts.Get(DebugOptKey))
		assert.Equal(t, opt2, opts.Get(TestingOptKey))
	})
}

func TestAdd(t *testing.T) {

	t.Run("option is nil for function", func(t *testing.T) {
		opts := []*Opt{}
		assert.False(t, Add(&opts, nil))
	})

	t.Run("option is nil for method", func(t *testing.T) {
		opts := NewOpts()
		assert.False(t, opts.Add(nil))
	})

	t.Run("slice is nil for function", func(t *testing.T) {
		var opts *[]*Opt
		assert.False(t, Add(opts, nil))
	})

	t.Run("slice is nil for method", func(t *testing.T) {
		var opts Opts
		assert.False(t, opts.Add())
	})

	t.Run("Add a new option via the function", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opts := []*Opt{}
		assert.True(t, Add(&opts, opt1))
		assert.Equal(t, 1, len(opts))
		assert.Equal(t, opt1, Get(opts, TestingOptKey))
	})

	t.Run("Add two options via the function", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opt2 := DebugOpt(true)
		opts := []*Opt{}
		assert.True(t, Add(&opts, opt1))
		assert.True(t, Add(&opts, opt2))
		assert.Equal(t, 2, len(opts))
		assert.Equal(t, opt1, Get(opts, TestingOptKey))
		assert.Equal(t, opt2, Get(opts, DebugOptKey))
	})

	t.Run("Add a new option via the method", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opts := NewOpts()
		opts.Add(opt1)
		assert.Equal(t, 1, opts.Len())
		assert.Equal(t, opt1, opts.Get(TestingOptKey))
	})

	t.Run("Add two options via the methd", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opt2 := DebugOpt(true)
		opts := NewOpts()
		opts.Add(opt1)
		opts.Add(opt2)
		assert.Equal(t, 2, opts.Len())
		assert.Equal(t, opt1, opts.Get(TestingOptKey))
		assert.Equal(t, opt2, opts.Get(DebugOptKey))
	})
}

func TestCopy(t *testing.T) {

	t.Run("Ensure original isn't affected by copy changes function", func(t *testing.T) {
		opts1 := []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}
		opts2 := Copy(opts1)
		assert.Equal(t, []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}, opts1)
		assert.Equal(t, []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}, opts2)
		Remove(&opts2, "1")
		assert.Equal(t, []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}, opts1)
		assert.Equal(t, []*Opt{{"2", 2}, {"3", 3}}, opts2)
	})

	t.Run("Ensure original isn't affected by copy changes via method", func(t *testing.T) {
		opts1 := NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3})
		opts2 := NewOpts(opts1.Copy()...)
		assert.Equal(t, NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}), opts1)
		assert.Equal(t, NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}), opts2)
		opts2.Remove("1")
		assert.Equal(t, NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}), opts1)
		assert.Equal(t, NewOpts(&Opt{"2", 2}, &Opt{"3", 3}), opts2)
	})
}

func TestDefault(t *testing.T) {

	t.Run("option is nil for function", func(t *testing.T) {
		opts := []*Opt{}
		assert.False(t, Default(&opts, nil))
	})

	t.Run("option is nil for method", func(t *testing.T) {
		opts := NewOpts()
		assert.False(t, opts.Default(nil))
	})

	t.Run("slice is nil for function", func(t *testing.T) {
		var opts *[]*Opt
		assert.False(t, Default(opts, nil))
	})

	t.Run("slice is nil for method", func(t *testing.T) {
		var opts Opts
		assert.False(t, opts.Default())
	})

	t.Run("Default a new option via the function", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opts := []*Opt{}
		assert.True(t, Default(&opts, opt1))
		assert.Equal(t, 1, len(opts))
		assert.Equal(t, opt1, Get(opts, TestingOptKey))
	})

	t.Run("Default two options via the function", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opt2 := DebugOpt(true)
		opts := []*Opt{}
		assert.True(t, Default(&opts, opt1))
		assert.True(t, Default(&opts, opt2))
		assert.Equal(t, 2, len(opts))
		assert.Equal(t, opt1, Get(opts, TestingOptKey))
		assert.Equal(t, opt2, Get(opts, DebugOptKey))
	})

	t.Run("Default a new option via the method", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opts := NewOpts()
		opts.Default(opt1)
		assert.Equal(t, 1, opts.Len())
		assert.Equal(t, opt1, opts.Get(TestingOptKey))
	})

	t.Run("Default two options via the methd", func(t *testing.T) {
		opt1 := TestingOpt(true)
		opt2 := DebugOpt(true)
		opts := NewOpts()
		opts.Default(opt1)
		opts.Default(opt2)
		assert.Equal(t, 2, opts.Len())
		assert.Equal(t, opt1, opts.Get(TestingOptKey))
		assert.Equal(t, opt2, opts.Get(DebugOptKey))
	})
}

func TestOverwrite(t *testing.T) {

	t.Run("Overwrite an option in the list via function", func(t *testing.T) {
		opts := []*Opt{}
		result := Overwrite(&opts, &Opt{"1", 1})
		assert.Equal(t, &Opt{"1", 1}, result)
		result = Overwrite(&opts, &Opt{"2", 2})
		assert.Equal(t, &Opt{"2", 2}, result)
		result = Overwrite(&opts, &Opt{"3", 3})
		assert.Equal(t, &Opt{"3", 3}, result)
		assert.Equal(t, []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}, opts)

		result = Overwrite(&opts, &Opt{"2", 5})
		assert.Equal(t, &Opt{"2", 5}, result)
		assert.Equal(t, []*Opt{{"1", 1}, {"2", 5}, {"3", 3}}, opts)
	})

	t.Run("Overwrite an option in the list via function", func(t *testing.T) {
		opts := NewOpts()
		result := opts.Overwrite(&Opt{"1", 1})
		assert.Equal(t, &Opt{"1", 1}, result)
		result = opts.Overwrite(&Opt{"2", 2})
		assert.Equal(t, &Opt{"2", 2}, result)
		result = opts.Overwrite(&Opt{"3", 3})
		assert.Equal(t, &Opt{"3", 3}, result)
		assert.Equal(t, NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3}), opts)

		result = opts.Overwrite(&Opt{"2", 5})
		assert.Equal(t, &Opt{"2", 5}, result)
		assert.Equal(t, NewOpts(&Opt{"1", 1}, &Opt{"2", 5}, &Opt{"3", 3}), opts)
	})
}

func TestRemove(t *testing.T) {
	t.Run("nil opt in slice via function", func(t *testing.T) {
		// removing the first should force check on middle before iterating past
		opts := []*Opt{{"1", 1}, (*Opt)(nil), {"3", 3}}
		Remove(&opts, "1")
		assert.Equal(t, []*Opt{(*Opt)(nil), {"3", 3}}, opts)

		// Remove end of slice
		Remove(&opts, "3")
		assert.Equal(t, []*Opt{(*Opt)(nil)}, opts)
	})

	t.Run("nil opt via function", func(t *testing.T) {
		opts := []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}
		Remove(&opts, "")
		assert.Equal(t, []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}, opts)
	})

	t.Run("nil opts via function", func(t *testing.T) {
		Remove((*[]*Opt)(nil), "2")
	})

	t.Run("remove middle via function", func(t *testing.T) {
		opts := []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}
		Remove(&opts, "2")
		assert.Equal(t, []*Opt{{"1", 1}, {"3", 3}}, opts)
	})

	t.Run("remove end via function", func(t *testing.T) {
		opts := []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}
		Remove(&opts, "3")
		assert.Equal(t, []*Opt{{"1", 1}, {"2", 2}}, opts)
	})

	t.Run("remove begining via function", func(t *testing.T) {
		opts := []*Opt{{"1", 1}, {"2", 2}, {"3", 3}}
		Remove(&opts, "1")
		assert.Equal(t, []*Opt{{"2", 2}, {"3", 3}}, opts)
	})

	t.Run("remove middle via method", func(t *testing.T) {
		opts := NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3})
		opts.Remove("2")
		assert.Equal(t, NewOpts(&Opt{"1", 1}, &Opt{"3", 3}), opts)
	})

	t.Run("remove end via method", func(t *testing.T) {
		opts := NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3})
		opts.Remove("3")
		assert.Equal(t, NewOpts(&Opt{"1", 1}, &Opt{"2", 2}), opts)
	})

	t.Run("remove begining via method", func(t *testing.T) {
		opts := NewOpts(&Opt{"1", 1}, &Opt{"2", 2}, &Opt{"3", 3})
		opts.Remove("1")
		assert.Equal(t, NewOpts(&Opt{"2", 2}, &Opt{"3", 3}), opts)
	})
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

// Dryrun
// -------------------------------------------------------------------------------------------------

func TestDryrunOpt(t *testing.T) {

	// create
	{
		opts := []*Opt{}
		assert.False(t, GetDryrunOpt(opts))
		assert.True(t, Default(&opts, DryrunOpt(true)))
		assert.True(t, GetDryrunOpt(opts))
	}

	// default
	{
		opts := []*Opt{}
		assert.False(t, GetDryrunOpt(opts))
		assert.False(t, DryrunOptExists(opts))
		assert.True(t, DefaultDryrunOpt(opts, true))
		assert.False(t, GetDryrunOpt(opts))
		assert.True(t, Default(&opts, DryrunOpt(true)))
		assert.True(t, GetDryrunOpt(opts))
		assert.True(t, DryrunOptExists(opts))
		assert.True(t, DefaultDryrunOpt(opts, false))
	}

	// overwrite
	{
		opts := []*Opt{}
		assert.Equal(t, true, Default(&opts, DryrunOpt(true)))
		assert.Equal(t, true, GetDryrunOpt(opts))
		assert.Equal(t, false, OverwriteDryrunOpt(&opts, false))
		assert.Equal(t, false, GetDryrunOpt(opts))
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

	// gets generically
	{
		opts := NewOpts()
		assert.Equal(t, false, GetBool(*opts, "testing"))
		assert.Equal(t, false, opts.GetBool("testing"))
		assert.True(t, Add((*[]*Opt)(opts), TestingOpt(true)))
		assert.Equal(t, true, GetBool(*opts, "testing"))
		assert.Equal(t, true, opts.GetBool("testing"))

		assert.Equal(t, "", GetString(*opts, "generic"))
		assert.Equal(t, "", opts.GetString("generic"))
		assert.Equal(t, &Opt{Key: "generic", Val: "foo"}, New("generic", "foo"))
		assert.True(t, Add((*[]*Opt)(opts), New("generic", "foo")))
		assert.Equal(t, "foo", GetString(*opts, "generic"))
		assert.Equal(t, "foo", opts.GetString("generic"))
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
