package nos

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// Abs gets the absolute path, taking into account homedir expansion
func Abs(target string) (result string, err error) {
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

// Dirs returns all named directories from the given target path
func Dirs(target string) (result []string) {
	result = []string{}
	if IsDir(target) {
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

// Paths returns a list of paths for the given root path in a deterministic order
func Paths(root string) (result []string) {
	result = []string{}
	filepath.Walk(root, func(p string, i os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if p != root && p != "." && p != ".." {
			result = append(result, p)
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
