package futil

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

func TestExtractString(t *testing.T) {
	clearTmpDir()

	// Write out test data
	err := sys.WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// invalid regex
	{
		match, err := ExtractString(tmpfile, `^\K`)
		assert.Equal(t, "", match)
		assert.Equal(t, "failed compiling regex '^\\K': error parsing regexp: invalid escape sequence: `\\K`", err.Error())
	}

	// attempt to read a write only file
	{
		assert.Nil(t, os.Chmod(tmpfile, 0222))
		match, err := ExtractString(tmpfile, "")
		assert.Equal(t, "", match)
		assert.True(t, strings.HasPrefix(err.Error(), "failed reading the file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))
		assert.Nil(t, os.Chmod(tmpfile, 0644))
	}

	// filepath is empty
	{
		match, err := ExtractString("", "")
		assert.Equal(t, "", match)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// No match
	{
		match, err := ExtractString(tmpfile, `foobar`)
		assert.Nil(t, err)
		assert.Equal(t, "", match)
	}

	// Single Match - whole string
	{
		match, err := ExtractString(tmpfile, `(?m)^(pkgver=.*)`)
		assert.Nil(t, err)
		assert.Equal(t, "pkgver=76.0.3809.100", match)
	}

	// Single Match
	{
		match, err := ExtractString(tmpfile, `(?m)^pkgver=(.*)`)
		assert.Nil(t, err)
		assert.Equal(t, "76.0.3809.100", match)
	}

	// Multiple Match - only sees the first one
	{
		match, err := ExtractString(tmpfile, `.*pkgver=(.*)`)
		assert.Nil(t, err)
		assert.Equal(t, "76.0.3809.100", match)
	}
}

func TestExtractStrings(t *testing.T) {
	clearTmpDir()

	// Write out test data
	err := sys.WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// invalid regex
	{
		match, err := ExtractStrings(tmpfile, `^\K`)
		assert.Equal(t, ([]string)(nil), match)
		assert.Equal(t, "failed compiling regex '^\\K': error parsing regexp: invalid escape sequence: `\\K`", err.Error())
	}

	// attempt to read a write only file
	{
		assert.Nil(t, os.Chmod(tmpfile, 0222))
		match, err := ExtractStrings(tmpfile, "")
		assert.Equal(t, ([]string)(nil), match)
		assert.True(t, strings.HasPrefix(err.Error(), "failed reading the file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))
		assert.Nil(t, os.Chmod(tmpfile, 0644))
	}

	// filepath is empty
	{
		match, err := ExtractStrings("", "")
		assert.Equal(t, ([]string)(nil), match)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// No match
	{
		matches, err := ExtractStrings(tmpfile, `foobar`)
		assert.Nil(t, err)
		assert.Nil(t, matches)
	}

	// Single Match substring
	{
		matches, err := ExtractStrings(tmpfile, `(?m)^pkgver=(.*)`)
		assert.Nil(t, err)
		assert.Equal(t, []string{"76.0.3809.100"}, matches)
	}

	// Single Match wholestring
	{
		matches, err := ExtractStrings(tmpfile, `(?m)^(pkgver=.*)`)
		assert.Nil(t, err)
		assert.Equal(t, []string{"pkgver=76.0.3809.100"}, matches)
	}

	// Multiple Match substrings
	{
		matches, err := ExtractStrings(tmpfile, `.*pkgver=(.*)`)
		assert.Nil(t, err)
		assert.Equal(t, []string{"76.0.3809.100", "foo"}, matches)
	}

	// Multiple Match whole strings
	{
		matches, err := ExtractStrings(tmpfile, `.*(pkgver=.*)`)
		assert.Nil(t, err)
		assert.Equal(t, []string{"pkgver=76.0.3809.100", "pkgver=foo"}, matches)
	}
}

func TestExtractStringP(t *testing.T) {
	clearTmpDir()

	// Write out test data
	err := sys.WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// attempt to read a write only file
	{
		rx, err := regexp.Compile(`foobar`)
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(tmpfile, 0222))
		match, err := ExtractStringP(tmpfile, rx)
		assert.Equal(t, "", match)
		assert.True(t, strings.HasPrefix(err.Error(), "failed reading the file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))
		assert.Nil(t, os.Chmod(tmpfile, 0644))
	}

	// filepath is empty
	{
		rx, err := regexp.Compile(`foobar`)
		assert.Nil(t, err)
		match, err := ExtractStringP("", rx)
		assert.Equal(t, "", match)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// No match
	{
		rx, err := regexp.Compile(`foobar`)
		assert.Nil(t, err)
		match, err := ExtractStringP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, "", match)
	}

	// Single Match - whole string
	{
		rx, err := regexp.Compile(`(?m)^(pkgver=.*)`)
		assert.Nil(t, err)
		match, err := ExtractStringP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, "pkgver=76.0.3809.100", match)
	}

	// Single Match
	{
		rx, err := regexp.Compile(`(?m)^pkgver=(.*)`)
		assert.Nil(t, err)
		match, err := ExtractStringP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, "76.0.3809.100", match)
	}

	// Multiple Match - only sees the first one
	{
		rx, err := regexp.Compile(`.*pkgver=(.*)`)
		assert.Nil(t, err)
		match, err := ExtractStringP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, "76.0.3809.100", match)
	}
}

func TestExtractStringsP(t *testing.T) {
	clearTmpDir()

	// Write out test data
	err := sys.WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// attempt to read a write only file
	{
		rx, err := regexp.Compile(`foobar`)
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(tmpfile, 0222))
		match, err := ExtractStringsP(tmpfile, rx)
		assert.Equal(t, ([]string)(nil), match)
		assert.True(t, strings.HasPrefix(err.Error(), "failed reading the file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))
		assert.Nil(t, os.Chmod(tmpfile, 0644))
	}

	// filepath is empty
	{
		rx, err := regexp.Compile(`foobar`)
		assert.Nil(t, err)
		match, err := ExtractStringsP("", rx)
		assert.Equal(t, ([]string)(nil), match)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// No match
	{
		rx, err := regexp.Compile(`foobar`)
		assert.Nil(t, err)
		matches, err := ExtractStringsP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Nil(t, matches)
	}

	// Single Match substring
	{
		rx, err := regexp.Compile(`(?m)^pkgver=(.*)`)
		assert.Nil(t, err)
		matches, err := ExtractStringsP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, []string{"76.0.3809.100"}, matches)
	}

	// Single Match wholestring
	{
		rx, err := regexp.Compile(`(?m)^(pkgver=.*)`)
		assert.Nil(t, err)
		matches, err := ExtractStringsP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, []string{"pkgver=76.0.3809.100"}, matches)
	}

	// Multiple Match substrings
	{
		rx, err := regexp.Compile(`.*pkgver=(.*)`)
		assert.Nil(t, err)
		matches, err := ExtractStringsP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, []string{"76.0.3809.100", "foo"}, matches)
	}

	// Multiple Match whole strings
	{
		rx, err := regexp.Compile(`.*(pkgver=.*)`)
		assert.Nil(t, err)
		matches, err := ExtractStringsP(tmpfile, rx)
		assert.Nil(t, err)
		assert.Equal(t, []string{"pkgver=76.0.3809.100", "pkgver=foo"}, matches)
	}
}

func clearTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
