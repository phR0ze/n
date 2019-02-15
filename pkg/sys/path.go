package sys

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
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

// Base wraps the filepath.Base but doesn't default to . when empty
func Base(src string) (result string) {
	base := filepath.Base(src)
	if base != "." {
		result = base
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

// AllFiles returns a list of all files recursively for the given root path
func AllFiles(root string) (result []string, err error) {
	result = []string{}
	if root, err = Abs(root); err != nil {
		return
	}
	err = Walk(root, func(p string, i *FileInfo, e error) error {
		if e != nil {
			return e
		}
		if p != root && p != "." && p != ".." && !i.IsDir() && !i.IsSymlinkDir() {
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

// AllPaths returns a list of all paths recursively for the given root path
// in a deterministic order including the root path as first entry
func AllPaths(root string) (result []string, err error) {
	if root, err = Abs(root); err != nil {
		return
	}
	result = []string{root}
	err = Walk(root, func(p string, i *FileInfo, e error) error {
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

// WalkFunc works the same as the filepath.WalkFunc
type WalkFunc func(path string, info *FileInfo, err error) error

// Walk extends the filepath.Walk to allow for it to walk symlinks
func Walk(root string, walkFn WalkFunc) (err error) {
	var info *FileInfo
	if info, err = Lstat(root); err != nil {
		err = walkFn(root, nil, err)
	} else {
		err = walk(root, info, walkFn)
	}
	if err == filepath.SkipDir {
		err = nil
	}
	return
}

// walk supports the public Walk function to allow for recursively walking a tree
// and following links unlike the filepath.Walk which doesn't follow links.
func walk(root string, info *FileInfo, walkFn WalkFunc) (err error) {
	if err = symlinkRecurse(root, info, walkFn); err != nil {
		return
	}

	// Call user walkFn if we have a file
	if !info.IsDir() {
		return walkFn(root, info, err)
	}

	// Recurse on directories
	var names []string
	names, err = SortedPaths(root)
	if e := walkFn(root, info, err); e != nil || err != nil {
		err = e
		return
	}
	for _, name := range names {
		target := filepath.Join(root, name)
		var targetInfo *FileInfo
		if targetInfo, err = Lstat(target); err != nil {
			if err = walkFn(target, targetInfo, err); err != nil && err != filepath.SkipDir {
				return
			}
		} else {
			if err = walk(target, targetInfo, walkFn); err != nil {
				if !targetInfo.IsDir() && (targetInfo.Mode()&os.ModeSymlink == 0) || err != filepath.SkipDir {
					return
				}
			}
		}
	}
	return
}

// recurse on symlinks for walk
func symlinkRecurse(root string, info *FileInfo, walkFn WalkFunc) (err error) {
	if info.IsSymlink() {

		// Evaluate the symlink to get the symlink's target
		var target string
		if target, err = filepath.EvalSymlinks(root); err != nil {
			return
		}

		// Ensure that the target exists
		if info, err = Lstat(target); err != nil {
			return
		}

		// Recurse on links to get to their target
		if err = walk(target, info, walkFn); err != nil && err != filepath.SkipDir {
			return
		}
	}

	return
}

// SortedPaths returns a list of the given directory's path names sorted
func SortedPaths(dir string) (names []string, err error) {
	var f *os.File
	if f, err = os.Open(dir); err != nil {
		return
	}
	names, err = f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return
	}
	sort.Strings(names)
	return
}
