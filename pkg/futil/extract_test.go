package futil

import (
	"regexp"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

func TestExtractString(t *testing.T) {
	cleanTmpDir()

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

	// No match
	match, err := ExtractString(tmpfile, `foobar`)
	assert.Nil(t, err)
	assert.Equal(t, "", match)

	// Single Match - whole string
	match, err = ExtractString(tmpfile, `(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	assert.Equal(t, "pkgver=76.0.3809.100", match)

	// Single Match
	match, err = ExtractString(tmpfile, `(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)

	// Multiple Match - only sees the first one
	match, err = ExtractString(tmpfile, `.*pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)
}

func TestExtractStrings(t *testing.T) {
	cleanTmpDir()

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

	// No match
	matches, err := ExtractStrings(tmpfile, `foobar`)
	assert.Nil(t, err)
	assert.Nil(t, matches)

	// Single Match substring
	matches, err = ExtractStrings(tmpfile, `(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100"}, matches)

	// Single Match wholestring
	matches, err = ExtractStrings(tmpfile, `(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100"}, matches)

	// Multiple Match substrings
	matches, err = ExtractStrings(tmpfile, `.*pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100", "foo"}, matches)

	// Multiple Match whole strings
	matches, err = ExtractStrings(tmpfile, `.*(pkgver=.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100", "pkgver=foo"}, matches)
}

func TestExtractStringP(t *testing.T) {
	cleanTmpDir()

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

	// No match
	rx, err := regexp.Compile(`foobar`)
	assert.Nil(t, err)
	match, err := ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "", match)

	// Single Match - whole string
	rx, err = regexp.Compile(`(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	match, err = ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "pkgver=76.0.3809.100", match)

	// Single Match
	rx, err = regexp.Compile(`(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	match, err = ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)

	// Multiple Match - only sees the first one
	rx, err = regexp.Compile(`.*pkgver=(.*)`)
	assert.Nil(t, err)
	match, err = ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)
}

func TestExtractStringsP(t *testing.T) {
	cleanTmpDir()

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

	// No match
	rx, err := regexp.Compile(`foobar`)
	assert.Nil(t, err)
	matches, err := ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Nil(t, matches)

	// Single Match substring
	rx, err = regexp.Compile(`(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100"}, matches)

	// Single Match wholestring
	rx, err = regexp.Compile(`(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100"}, matches)

	// Multiple Match substrings
	rx, err = regexp.Compile(`.*pkgver=(.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100", "foo"}, matches)

	// Multiple Match whole strings
	rx, err = regexp.Compile(`.*(pkgver=.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100", "pkgver=foo"}, matches)
}

func cleanTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
