package sys

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// Abs gets the absolute path, taking into account homedir expansion
func Abs(target string) (result string, err error) {
	if target == "" {
		err = fmt.Errorf("Empty string is an invalid path")
		return
	}
	target = TrimProtocol(target)
	if result, err = homedir.Expand(target); err == nil {
		result, err = filepath.Abs(result)
	}
	return
}

// Home returns the absolute home directory for the current user
func Home() (result string, err error) {
	if result, err = homedir.Dir(); err == nil {
		result, err = filepath.Abs(result)
	}
	return
}

// Dirs returns all directories from the given target path
func Dirs(target string) (result []string) {
	result = []string{}
	if target != "" && IsDir(target) {
		if target, err := Abs(target); err == nil {
			if items, err := ioutil.ReadDir(target); err == nil {
				for _, item := range items {
					if item.IsDir() {
						result = append(result, path.Join(target, item.Name()))
					}
				}
			}
		}
	}
	return
}

// Files returns all directories from the given target path
func Files(target string) (result []string) {
	result = []string{}
	if target != "" && IsDir(target) {
		if target, err := Abs(target); err == nil {
			if items, err := ioutil.ReadDir(target); err == nil {
				for _, item := range items {
					if !item.IsDir() {
						result = append(result, path.Join(target, item.Name()))
					}
				}
			}
		}
	}
	return
}

// Paths returns all directories and files from the given target path
func Paths(target string) (result []string) {
	result = []string{}
	if target != "" && IsDir(target) {
		if target, err := Abs(target); err == nil {
			if items, err := ioutil.ReadDir(target); err == nil {
				for _, item := range items {
					result = append(result, path.Join(target, item.Name()))
				}
			}
		}
	}
	return
}

// AllPaths returns a list of all paths recursively for the given root path
// in a deterministic order including the root path as first entry
func AllPaths(root string) (result []string, err error) {
	if root, err = Abs(root); err != nil {
		return
	}
	result = []string{root}
	err = filepath.Walk(root, func(p string, i os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if p != root && p != "." && p != ".." {
			absPath, e := Abs(p)
			if e != nil {
				return e
			}
			result = append(result, absPath)
		}
		return nil
	})
	return
}

// SharedDir returns the dir portion that two paths share
func SharedDir(first, second string) (result string) {
	sharedParts := []string{}

	firstParts := strings.Split(first, "/")
	secondParts := strings.Split(second, "/")
	secondLen := len(secondParts)
	for i := range firstParts {
		if i < secondLen {
			if firstParts[i] == secondParts[i] {
				sharedParts = append(sharedParts, firstParts[i])
			}
		} else {
			break
		}
	}

	return strings.Join(sharedParts, "/")
}

// SlicePath provides a ruby like slice function for path nubs
func SlicePath(target string, i, j int) (result string) {
	x := strings.Split(target, "/")

	// Convert to positive notation to simplify logic
	if i < 0 {
		i = len(x) + i
	}

	// Offset indices to include root
	if target != "" && rune(target[0]) == rune('/') {
		if i == 1 {
			i--
		} else if i == 0 && j >= 0 {
			j++
		} else if j > 0 {
			i, j = i+1, j+1
		}
	}

	result = strings.Join(slice(x, i, j), "/")

	return
}

// TrimExt removes the extension from the given target path
func TrimExt(target string) string {
	ext := path.Ext(target)
	target = strings.TrimSuffix(target, ext)
	return target
}

// TrimProtocol removes well known protocol prefixes
func TrimProtocol(target string) string {
	target = strings.TrimPrefix(target, "file://")
	target = strings.TrimPrefix(target, "ftp://")
	target = strings.TrimPrefix(target, "http://")
	target = strings.TrimPrefix(target, "https://")
	return target
}
