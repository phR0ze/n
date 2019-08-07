package sys

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/phR0ze/n/pkg/opt"
	"github.com/pkg/errors"
)

// Abs gets the absolute path, taking into account homedir expansion
func Abs(target string) (result string, err error) {
	if target == "" {
		err = errors.Errorf("empty string is an invalid path")
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

	// Drop the trailing slash if it exists
	if len(src) > 1 && rune(src[len(src)-1]) == '/' {
		src = src[:len(src)-1]
	}

	// Call base function
	base := filepath.Base(src)
	if base != "." {
		result = base
	}
	return
}

// Dir wraps the filepath.Dir and trims off tailing slashes
func Dir(src string) (result string) {

	// Drop the trailing slash if it exists
	if len(src) > 1 && rune(src[len(src)-1]) == '/' {
		src = src[:len(src)-1]
	}

	// Call base function
	base := filepath.Dir(src)
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

// Dirs returns all directories from the given target path, sorted by filename
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

// Files returns all files from the given target path, sorted by filename
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

// Paths returns all directories/files from the given target path, sorted by filename
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

// AllDirs returns a list of all dirs recursively for the given root path
// in a deterministic order. Follows links by default, but can be stopped
// with &Opt{"follow", false}. Paths are distinct.
func AllDirs(root string, opts ...*opt.Opt) (result []string, err error) {
	distinct := map[string]bool{}
	if root, err = Abs(root); err != nil {
		return
	}

	// Set following links by default
	defaultFollowOpt(&opts, true)

	err = Walk(root, func(p string, i *FileInfo, e error) error {
		if e != nil {
			return e
		}

		// IsDir will ignore files
		if p != root && p != "." && p != ".." && i.IsDir() {
			absPath, e := Abs(p)
			if e != nil {
				return e
			}

			// Ensure file paths are distinct
			if _, exists := distinct[absPath]; !exists {
				distinct[absPath] = true
				result = append(result, absPath)
			}
		}
		return nil
	}, opts...)
	return
}

// AllFiles returns a list of all files recursively for the given root path
// in a deterministic order. Follows links by default, but can be stopped
// with &Opt{"follow", false}. Paths are distinct.
func AllFiles(root string, opts ...*opt.Opt) (result []string, err error) {
	distinct := map[string]bool{}
	if root, err = Abs(root); err != nil {
		return
	}

	// Set following links by default
	defaultFollowOpt(&opts, true)

	err = Walk(root, func(p string, i *FileInfo, e error) error {
		if e != nil {
			return e
		}
		// IsFile will ignore both directories and links to files/dirs when not following.
		// When followOpt is set then Walk will return followed files.
		if p != root && p != "." && p != ".." && i.IsFile() {
			absPath, e := Abs(p)
			if e != nil {
				return e
			}

			// Ensure file paths are distinct
			if _, exists := distinct[absPath]; !exists {
				distinct[absPath] = true
				result = append(result, absPath)
			}
		}
		return nil
	}, opts...)
	return
}

// AllPaths returns a list of all paths recursively for the given root path
// in a deterministic order including the root path as first entry.
// Follows links by default, but can be stopped with &Opt{"follow", false}.
// Paths are distinct.
func AllPaths(root string, opts ...*opt.Opt) (result []string, err error) {
	distinct := map[string]bool{}
	if root, err = Abs(root); err != nil {
		return
	}

	// Set following links by default
	defaultFollowOpt(&opts, true)

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

			// Ensure paths are distinct
			if _, exists := distinct[absPath]; !exists {
				distinct[absPath] = true
				result = append(result, absPath)
			}
		}
		return nil
	}, opts...)
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

// SlicePath provides a ruby like slice function for path manipulation
func SlicePath(target string, i, j int) (result string) {

	// Drop the trailing slash if it exists
	if len(target) > 1 && rune(target[len(target)-1]) == '/' {
		target = target[:len(target)-1]
	}

	x := strings.Split(target, "/")

	// Convert to positive notation to simplify logic
	if i < 0 {
		i = len(x) + i
	}

	// Offset indices to include root
	if target != "" && rune(target[0]) == '/' {
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
// by default but can be turned off with &Opt{"follow", false}
func Walk(root string, walkFn WalkFunc, opts ...*opt.Opt) (err error) {

	// Set following links by default
	defaultFollowOpt(&opts, true)

	var info *FileInfo
	if info, err = Lstat(root); err != nil {
		err = walkFn(root, nil, err)
	} else {
		err = walk(root, info, walkFn, opts)
	}
	if err == filepath.SkipDir {
		err = nil
	}
	return
}

// walk supports the public Walk function to allow for recursively walking a tree
// and following links unlike the filepath.Walk which doesn't follow links.
func walk(root string, info *FileInfo, walkFn WalkFunc, opts []*opt.Opt) (err error) {
	target := root
	targetInfo := info
	targets := []string{}

	// First thing pass whatever we've got on to user walkFn
	if err = walkFn(root, info, err); err != nil {
		return
	}

	// Were done if it was a regular file or link and were not following
	if info.IsFile() || (info.IsSymlink() && !followOpt(opts)) {
		return
	}

	// Links and directories are similar in that they have other paths to deal with
	if info.IsDir() {
		var names []string
		if names, err = ReadDir(root); err == nil {
			for _, name := range names {
				target = filepath.Join(root, name)
				targets = append(targets, target)
			}
		}
	} else {
		if target, err = filepath.EvalSymlinks(root); err == nil {
			targets = append(targets, target)
		}
	}

	// Return errors to the user walkFn, return error from user means skip else continue
	if err != nil {
		if err = walkFn(target, targetInfo, err); err != nil {
			return
		}
	}

	// Recurse on target paths
	for _, target := range targets {
		if targetInfo, err = Lstat(target); err != nil {
			// Return errors to the user walkFn
			if err = walkFn(target, targetInfo, err); err != nil {
				return
			}
		} else {
			// No error so recurse on the path
			if err = walk(target, targetInfo, walkFn, opts); err != nil {
				return
			}
		}
	}
	return
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDir(dirname string) (names []string, err error) {
	var f *os.File
	if f, err = os.Open(dirname); err != nil {
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
