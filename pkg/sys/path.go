package sys

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/phR0ze/n/pkg/opt"
	"github.com/pkg/errors"
)

// Abs gets the absolute path, taking into account path expansion and protocols
func Abs(target string) (result string, err error) {

	// Check for empty string
	if target == "" {
		err = errors.Errorf("empty string is an invalid path")
		return
	}

	// Bail out early if the path is rooted
	if target[0] == '/' {
		result = target
		return
	}

	// Trim protocols and expand
	target = TrimProtocol(target)
	if result, err = Expand(target); err != nil {
		err = errors.Wrapf(err, "failed to expand the given path %s", target)
		return
	}

	// Get the absolute path
	if result, err = filepath.Abs(result); err != nil {
		err = errors.Wrapf(err, "failed to compute the absolute path for %s", result)
		return
	}
	return
}

// AllDirs returns a list of all dirs recursively for the given root path
// in a deterministic order. Follows links by default, but can be stopped
// by passing FollowOpt(false). Paths are distinct.
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

		// IsDir will ignore files and links
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
// by passing FollowOpt(false). Paths are distinct.
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
// in a deterministic order including the root path as first entry. Follows links
// by default, but can be stopped by passing FollowOpt(false). Paths are distinct.
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

// Dirs returns all directories from the given target path, sorted by filename
// Doesn't include the target itself only its children nor is this recursive.
func Dirs(target string) (result []string) {
	result = []string{}
	if target, err := Abs(target); err == nil {
		if IsDir(target) {
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

// Executable simply wraps os.Executable for convenience
func Executable() (string, error) {
	return os.Executable()
}

// Expand the path to include the home prefix if necessary
func Expand(target string) (path string, err error) {
	path = target

	// Nothing to do but not invalid
	if path == "" || !strings.Contains(path, "~") {
		return
	}

	// Invalid expansion requested
	if path[0] != '~' || len(path) < 2 || path[1] != '/' {
		path = ""
		err = errors.Errorf("failed to expand invalid path")
		return
	}

	// Get home directory
	var home string
	if home, err = UserHome(); err != nil {
		path = ""
		return
	}

	// Replace prefix with home directory
	path = filepath.Join(home, path[1:])

	// Invalid expansion requested
	if strings.Contains(path, "~") {
		path = ""
		err = errors.Errorf("invalid expansion requested")
		return
	}
	return
}

// Files returns all files from the given target path, sorted by filename.
// Doesn't include the target itself only its children nor is this recursive.
func Files(target string) (result []string) {
	result = []string{}
	if target, err := Abs(target); err == nil {
		if IsDir(target) {
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

// Glob wraps filepath.Glob but provides path expansion, recursion and error tracing.
// If no sources are found an empty string slice will be returned and a nil error.
// Enable recursion by passing in the option RecurseOpt(true).
func Glob(path string, opts ...*opt.Opt) (sources []string, err error) {
	recurse := getRecurseOpt(opts)

	// Path expansion
	if path, err = Abs(path); err != nil {
		return
	}

	// Handle globbing
	if sources, err = filepath.Glob(path); err != nil {
		err = errors.Wrapf(err, "failed to get glob for %s", path)
		return
	}

	// Execute the recursion if requested
	if recurse {
		for _, source := range sources {
			if IsDir(source) {
				var paths []string
				if paths, err = AllPaths(source); err != nil {
					return
				}
				sources = append(sources, paths[1:]...)
			}
		}
	}
	return
}

// Paths returns all directories/files from the given target path, sorted by filename.
// Doesn't include the target itself only its children nor is this recursive.
func Paths(target string) (result []string) {
	result = []string{}
	if target, err := Abs(target); err == nil {
		if IsDir(target) {
			if items, err := ioutil.ReadDir(target); err == nil {
				for _, item := range items {
					result = append(result, path.Join(target, item.Name()))
				}
			}
		}
	}
	return
}

// ReadDir reads the directory named by dirname and returns a list of directory entries
// sorted by filename. Similar to ioutil.ReadDir but using internal FileInfo types.
func ReadDir(dirname string) (infos []*FileInfo, err error) {
	if dirname, err = Abs(dirname); err != nil {
		return
	}

	var fr *os.File
	if fr, err = os.Open(dirname); err != nil {
		err = errors.Wrapf(err, "failed to open directory %s", dirname)
		return
	}
	defer fr.Close()

	// Read in os.FileInfo types as a sorted list
	var list []os.FileInfo
	if list, err = fr.Readdir(-1); err != nil {
		err = errors.Wrapf(err, "failed to read directory %s", dirname)
		return
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })

	// Create internal FileInfo types with path
	for _, info := range list {
		infos = append(infos, &FileInfo{Obj: info, Path: path.Join(dirname, info.Name())})
	}

	return
}

// ReadDirnames reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDirnames(dirname string) (names []string, err error) {
	if dirname, err = Abs(dirname); err != nil {
		return
	}

	// Open the directory for reading
	var fr *os.File
	if fr, err = os.Open(dirname); err != nil {
		err = errors.Wrapf(err, "failed to open directory %s", dirname)
		return
	}
	defer fr.Close()

	// Read the directory names as strings and sort
	if names, err = fr.Readdirnames(-1); err != nil {
		err = errors.Wrapf(err, "failed to read directory names for %s", dirname)
		return
	}
	sort.Strings(names)

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
	root := false
	x := strings.Split(target, "/")

	// Drop leading/trailing slashes
	if len(x) > 1 && x[len(x)-1] == "" {
		x = x[:len(x)-1]
	}
	if len(x) > 1 && x[0] == "" {
		root = true
		x = x[1:len(x)]
	}

	// Bail if there is nothing to do
	if len(x) < 2 {
		return target
	}

	// Slice and include root
	if i < 0 {
		i = len(x) + i
		if i < 0 {
			i = 0
		}
	}
	y := slice(x, i, j)
	if root && i == 0 {
		y = append([]string{""}, y...)
	}

	result = strings.Join(y, "/")
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

// TrimShared returns portion of target that isn't in shared
func TrimShared(target, shared string) string {
	results := []string{}

	targetParts := strings.Split(target, "/")
	sharedParts := strings.Split(shared, "/")
	secondLen := len(sharedParts)
	for i := range targetParts {
		if i < secondLen {
			if targetParts[i] != sharedParts[i] {
				results = targetParts[i:]
				break
			}
		} else {
			results = targetParts[i:]
			break
		}
	}

	return strings.Join(results, "/")
}

// WalkFunc works the same as the filepath.WalkFunc
type WalkFunc func(path string, info *FileInfo, err error) error

// Walk extends the filepath.Walk to allow for it to walk symlinks
// by default but can be turned off by passing in FollowOpt(false)
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
	targets := []string{}

	// First thing pass whatever we've got on to user walkFn so the user has
	// the ability to skip this path before processing is done on it
	if err = walkFn(root, info, err); err != nil {
		return
	}

	// Were done if it was a regular file or link and were not following
	if info.IsFile() || (info.IsSymlink() && !getFollowOpt(opts)) {
		return
	}

	// Links and directories are similar in that they have other paths to deal with
	if info.IsDir() {
		var names []string
		if names, err = ReadDirnames(root); err == nil {
			for _, name := range names {
				targets = append(targets, filepath.Join(root, name))
			}
		}
	} else {
		var target string
		if target, err = filepath.EvalSymlinks(root); err == nil {
			targets = append(targets, target)
		}
	}

	// Return errors to the user walkFn, return error from user means skip else continue
	if err != nil {
		if err = walkFn(root, info, err); err != nil {
			return
		}
	}

	// Recurse on target paths
	var targetInfo *FileInfo
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
