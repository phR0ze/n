package nos

import (
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// Path helper class
type pathN struct {
	v string
}

// Path creats a new path nub from the given string
func Path(path string) *pathN {
	return &pathN{v: strings.TrimSpace(path)}
}

// Abs gets the absolute path, taking into account homedir expansion
func (p *pathN) Abs() (result string, err error) {
	if result, err = homedir.Expand(p.v); err == nil {
		result, err = filepath.Abs(result)
	}
	return
}

// Slice provides a ruby like slice function for path nubs
func (p *pathN) Slice(i, j int) (result string) {
	x := strings.Split(p.v, "/")

	// Convert to positive notation to simplify logic
	if i < 0 {
		i = len(x) + i
	}

	// Offset indices to include root
	if p.v != "" && rune(p.v[0]) == rune('/') {
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

// Home returns the absolute home directory for the current user
func Home() (result string, err error) {
	if result, err = homedir.Dir(); err == nil {
		result, err = filepath.Abs(result)
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
		result = append(result, p)
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
