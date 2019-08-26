// Package json provides helper functions for working with json
package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/phR0ze/go-errors"
	"github.com/phR0ze/n/pkg/sys"
)

// Marshal wraps the json.Marshal
func Marshal(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

// ReadJSON reads the target file and returns a map[string]interface{} data
// structure representing the json read in.
func ReadJSON(filepath string) (obj map[string]interface{}, err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
		return
	}

	// Convert data structure into a json string
	if err = json.Unmarshal(data, &obj); err != nil {
		err = errors.Wrapf(err, "failed to unmarshal object %T", obj)
	}
	return
}

// Unmarshal wraps the json.Unmarshal
func Unmarshal(y []byte, o interface{}) error {
	return json.Unmarshal(y, o)
}

// WriteJSON converts the given obj interface{} into json then writes to disk
// with default permissions. Expects obj to be a structure that encoding/json understands
func WriteJSON(filepath string, obj interface{}, indent ...int) (err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Ensure we don't have a string
	switch obj.(type) {
	case string, []byte:
		err = errors.Errorf("invalid data structure to marshal - %T", obj)
		return
	}

	// Set indent level
	i := 2
	if len(indent) > 0 {
		i = indent[0]
	}

	// Convert data structure into a json string
	var data []byte
	if i == 0 {
		if data, err = json.Marshal(obj); err != nil {
			err = errors.Wrapf(err, "failed to marshal object %T", obj)
			return
		}
	} else {
		if data, err = json.MarshalIndent(obj, "", strings.Repeat(" ", i)); err != nil {
			err = errors.Wrapf(err, "failed to marshal object %T", obj)
			return
		}
	}

	// Use default permissions for file
	if err = ioutil.WriteFile(filepath, data, os.FileMode(0644)); err != nil {
		err = errors.Wrapf(err, "failed to write out json data to file %s", filepath)
	}
	return
}
