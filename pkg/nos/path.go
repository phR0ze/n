package nos

import (
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
