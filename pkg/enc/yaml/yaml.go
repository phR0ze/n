// Package yaml provides helper functions for working with yaml
package yaml

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/go-errors"
	"github.com/phR0ze/n/pkg/sys"
)

// Marshal wraps the ghodss/yaml.Marshal
func Marshal(o interface{}) ([]byte, error) {
	return yaml.Marshal(o)
}

// ReadYAML reads the target file and returns a map[string]interface{} data
// structure representing the yaml read in.
func ReadYAML(filepath string) (obj map[string]interface{}, err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
		return
	}

	// Convert data structure into a yaml string
	obj = map[string]interface{}{}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = errors.Wrapf(err, "failed to unmarshal object %T", obj)
	}
	return
}

// Unmarshal wraps the ghodss/yaml.Unmarshal
func Unmarshal(y []byte, o interface{}) error {
	return yaml.Unmarshal(y, o)
}

// WriteYAML converts the given obj interface{} into yaml then writes to disk
// with default permissions. Expects obj to be a structure that github.com/ghodss/yaml understands
func WriteYAML(filepath string, obj interface{}, perms ...uint32) (err error) {
	if filepath, err = sys.Abs(filepath); err != nil {
		return
	}

	// Ensure we don't have a string
	switch obj.(type) {
	case string, []byte:
		err = errors.Errorf("invalid data structure to marshal - %T", obj)
		return
	}

	// Convert data structure into a yaml string
	var data []byte
	if data, err = yaml.Marshal(obj); err != nil {
		err = errors.Wrapf(err, "failed to marshal object %T", obj)
		return
	}

	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}
	if err = ioutil.WriteFile(filepath, data, perm); err != nil {
		err = errors.Wrapf(err, "failed to write out yaml data to file %s", filepath)
	}
	return
}
