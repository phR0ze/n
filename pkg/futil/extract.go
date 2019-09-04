// Package futil provides helper functions for interacting with files
package futil

import (
	"io/ioutil"
	"regexp"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/pkg/errors"
)

// ExtractString reads the filepath data then compiles the given regular
// expression exp and applies it to the data and returns the results.
// Match will be empty if no matches were found. Use (?m) to have ^ $ apply
// to each line in the string. Use (?s) to have . span lines.
func ExtractString(filepath string, exp string) (match string, err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed reading the file %s", filepath)
		return
	}

	// Compile the regular expression
	var rx *regexp.Regexp
	if rx, err = regexp.Compile(exp); err != nil {
		err = errors.Wrapf(err, "failed compiling regex '%s'", exp)
		return
	}

	// Apply the regular expression to the data
	if results := rx.FindStringSubmatch(string(data)); len(results) > 1 {
		match = results[1]
	}

	return
}

// ExtractStrings reads the filepath data then compiles the given regular
// expression exp and applies it to the data and returns the results.
// Matches will be nil if no matches were found. Use (?m) to have ^ $ apply
// to each line in the string. Use (?s) to have . span lines.
func ExtractStrings(filepath string, exp string) (matches []string, err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed reading the file %s", filepath)
		return
	}

	// Compile the regular expression
	var rx *regexp.Regexp
	if rx, err = regexp.Compile(exp); err != nil {
		err = errors.Wrapf(err, "failed compiling regex '%s'", exp)
		return
	}

	// Apply the regular expression to the data
	for _, x := range rx.FindAllStringSubmatch(string(data), -1) {
		if len(x) > 1 {
			matches = append(matches, x[1])
		}
	}

	return
}

// ExtractStringP reads the filepath data then applies the given regular
// expression to the data and returns the results. See ExtractString
func ExtractStringP(filepath string, exp *regexp.Regexp) (match string, err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed reading the file %s", filepath)
		return
	}

	// Apply the regular expression to the data
	if results := exp.FindStringSubmatch(string(data)); len(results) > 1 {
		match = results[1]
	}

	return
}

// ExtractStringsP reads the filepath data then applies the given regular
// expression to the data and returns the results. See ExtractStrings
func ExtractStringsP(filepath string, exp *regexp.Regexp) (matches []string, err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed reading the file %s", filepath)
		return
	}

	// Apply the regular expression to the data
	for _, x := range exp.FindAllStringSubmatch(string(data), -1) {
		if len(x) > 1 {
			matches = append(matches, x[1])
		}
	}

	return
}
