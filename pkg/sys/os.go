// Package sys provides os level helper functions for interacting with the system
package sys

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/opt"
	"github.com/pkg/errors"
)

// Copy copies src to dst recursively, creating destination directories as needed.
// Handles globbing e.g. Copy("./*", "../")
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist.
// Doesn't follow links by default but can be turned on with &Opt{"follow", true}
func Copy(src, dst string, opts ...*opt.Opt) (err error) {
	var sources []string

	// Set following links to off by default
	defaultFollowOpt(&opts, false)

	// Get Abs src and dst roots
	var dstAbs, srcAbs string
	if dstAbs, err = Abs(dst); err != nil {
		return
	}
	if srcAbs, err = Abs(src); err != nil {
		return
	}

	// Handle globbing
	if sources, err = filepath.Glob(srcAbs); err != nil {
		err = errors.Wrapf(err, "failed to get glob for %s", srcAbs)
		return
	}

	// Clone given src as dst vs copy into dst
	clone := true
	if IsDir(dstAbs) {
		clone = false
	}

	// Copy all sources to dst
	for _, srcRoot := range sources {

		// Walk over file structure
		err = Walk(srcRoot, func(srcPath string, srcInfo *FileInfo, e error) error {
			if e != nil {
				return e
			}

			// Set proper dst path
			var dstPath string
			if clone {
				dstPath = path.Join(dstAbs, strings.TrimPrefix(srcPath, srcRoot))
			} else {
				dstPath = path.Join(dstAbs, strings.TrimPrefix(srcPath, path.Dir(srcRoot)))
			}

			// Create destination directories as needed
			if srcInfo.IsDir() {
				if e = os.MkdirAll(dstPath, srcInfo.Mode()); e != nil {
					return e
				}
			} else {
				CopyFile(srcPath, dstPath, newInfoOpt(srcInfo))
			}
			return nil
		}, newFollowOpt(false))
	}
	return
}

// CopyFile copies a single file from src to dsty, creating destination directories as needed.
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist.
// Doesn't follow links by default but can be turned on with &Opt{"follow", true}
func CopyFile(src, dst string, opts ...*opt.Opt) (err error) {
	var srcPath, dstPath string
	var srcInfo, srcDirInfo *FileInfo

	// Set following links to off by default
	defaultFollowOpt(&opts, false)

	// Check the source for issues
	if srcInfo = infoOpt(opts); srcInfo != nil {
		srcPath = srcInfo.path
	} else {
		if srcInfo, err = Lstat(src); err != nil {
			return
		}
		if srcPath, err = Abs(src); err != nil {
			return
		}
	}

	// Source dir permissions to use for destination directories
	if srcDirInfo, err = Lstat(path.Dir(src)); err != nil {
		return
	}

	// Error out if not a regular file or symlink
	if srcInfo.IsDir() {
		err = errors.Errorf("src target is not a regular file or a symlink to a file")
		return
	}

	// Get correct destination path
	dstInfo, e := os.Stat(dst)
	switch {

	// Doesn't exist so this is the new destination name
	case os.IsNotExist(e):
		if e = os.MkdirAll(path.Dir(dst), srcDirInfo.Mode()); e != nil {
			return e
		}
		if dstPath, err = Abs(path.Dir(dst)); err != nil {
			return
		}
		dstPath = path.Join(dstPath, path.Base(dst))

	// Destination exists and is either a file to overwrite or a dir to copy into
	case e == nil:
		if dstPath, err = Abs(dst); err != nil {
			return
		}
		if dstInfo.IsDir() {
			dstPath = path.Join(dstPath, path.Base(srcPath))
		}

	// unknown error case
	default:
		err = e
		return
	}

	// Handle links a bit differently
	if srcInfo.IsSymlink() {
		var target string
		if target, err = srcInfo.SymlinkTarget(); err != nil {
			return
		}
		if err = os.Symlink(target, dstPath); err != nil {
			return
		}
	} else {
		// Open srcPath for reading
		var fin *os.File
		if fin, err = os.Open(srcPath); err != nil {
			return
		}
		defer fin.Close()

		// Create dstPath for writing
		var fout *os.File
		if fout, err = os.Create(dstPath); err != nil {
			return
		}
		defer fout.Close()

		// Copy srcPath to dstPath
		if _, err = io.Copy(fout, fin); err != nil {
			return
		}

		// Sync to disk
		if err = fout.Sync(); err != nil {
			return
		}

		// Set permissions of dstPath same as srcPath
		err = os.Chmod(dstPath, srcInfo.Mode())
	}

	return
}

// Exists return true if the given path exists
func Exists(src string) bool {
	if target, err := Abs(src); err == nil {
		if _, err := os.Stat(target); err == nil {
			return true
		}
	}
	return false
}

// ExtractString reads the filepath data then compiles the given regular
// expression exp and applies it to the data and returns the results.
// Match will be empty if no matches were found. Use (?m) to have ^ $ apply
// to each line in the string. Use (?s) to have . span lines.
func ExtractString(filepath string, exp string) (match string, err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
		return
	}

	// Compile the regular expression
	var rx *regexp.Regexp
	if rx, err = regexp.Compile(exp); err != nil {
		err = errors.Wrapf(err, "failed to compile regex '%s'", exp)
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
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
		return
	}

	// Compile the regular expression
	var rx *regexp.Regexp
	if rx, err = regexp.Compile(exp); err != nil {
		err = errors.Wrapf(err, "failed to compile regex '%s'", exp)
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
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
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
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	// Read in the file data
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
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

// IsDarwin returns true if the OS is OSX
func IsDarwin() (result bool) {
	if runtime.GOOS == "darwin" {
		result = true
	}
	return
}

// IsLinux returns true if the OS is Linux
func IsLinux() (result bool) {
	if runtime.GOOS == "linux" {
		result = true
	}
	return
}

// MD5 returns the md5 of the given file
func MD5(target string) (result string, err error) {
	if target, err = Abs(target); err != nil {
		return
	}
	if !Exists(target) {
		return "", os.ErrNotExist
	}

	// Open target file for reading
	var f *os.File
	if f, err = os.Open(target); err != nil {
		err = errors.Wrapf(err, "failed to open target file %s", target)
		return
	}
	defer f.Close()

	// Create a new md5 hash and copy in file bits
	hash := md5.New()
	if _, err = io.Copy(hash, f); err != nil {
		err = errors.Wrapf(err, "failed to copy file data into hash from %s", target)
		return
	}

	// Compute 32 byte hash
	result = hex.EncodeToString(hash.Sum(nil))

	return
}

// MkdirP creates the target directory and any parent directories needed
// and returns the ABS path of the created directory
func MkdirP(target string, perms ...uint32) (dir string, err error) {
	if dir, err = Abs(target); err != nil {
		return
	}

	// Get/set default permission
	perm := os.FileMode(0755)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}

	// Create directory
	if err = os.MkdirAll(dir, perm); err != nil {
		err = errors.Wrapf(err, "failed to create directories for %s", dir)
		return
	}

	return
}

// Move the src path to the dst path. If the dst already exists and is not a directory
// src will replace it. If there is an error it will be of type *LinkError. Wraps
// os.Rename but fixes the issue where dst name is required
func Move(src, dst string) (err error) {

	// Add src base name to dst directory to fix golang oversight
	if IsDir(dst) {
		dst = path.Join(dst, path.Base(src))
	}
	if err = os.Rename(src, dst); err != nil {
		err = errors.Wrapf(err, "failed to rename file %s", src)
		return
	}
	return
}

// Pwd returns the current working directory
func Pwd() (pwd string) {
	pwd, _ = os.Getwd()
	return
}

// ReadBytes returns the entire file as []byte
func ReadBytes(filepath string) (result []byte, err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	if result, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
		return
	}
	return
}

// ReadLines returns a new slice of string representing lines
func ReadLines(filepath string) (result []string, err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}

// ReadString returns the entire file as a string
func ReadString(filepath string) (result string, err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(filepath); err != nil {
		err = errors.Wrapf(err, "failed to read the file %s", filepath)
		return
	}
	result = string(data)
	return
}

// ReadYaml reads the target file and returns a map[string]interface{} data
// structure representing the yaml read in.
func ReadYaml(filepath string) (obj map[string]interface{}, err error) {
	if filepath, err = Abs(filepath); err != nil {
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
		err = errors.Wrapf(err, "failed to marshal object %T", obj)
	}
	return
}

// Remove the given target file or empty directory. If there is an
// error it will be of type *PathError
func Remove(target string) error {
	return os.Remove(target)
}

// RemoveAll removes the target path and any children it contains.
// It removes everything it can but returns the first error it encounters.
// If the target path does not exist nil is returned
func RemoveAll(target string) error {
	return os.RemoveAll(target)
}

// Symlink creates newname as a symbolic link to oldname. If there is an error,
// it will be of type *LinkError. newname is created as ???
func Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

// Touch creates an empty text file similar to the linux touch command
func Touch(filepath string) (path string, err error) {
	if path, err = Abs(filepath); err != nil {
		return
	}

	var f *os.File
	if f, err = os.Create(path); !os.IsExist(err) {
		err = errors.Wrapf(err, "failed create file %s", filepath)
		return
	}
	if err == nil {
		defer f.Close()
	}
	return
}

// WriteBytes is a pass through to ioutil.WriteBytes with default permissions
func WriteBytes(filepath string, data []byte, perms ...uint32) (err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}
	if err = ioutil.WriteFile(filepath, data, perm); err != nil {
		err = errors.Wrapf(err, "failed write string to file %s", filepath)
		return
	}
	return
}

// WriteLines is a pass through to ioutil.WriteFile with default permissions
func WriteLines(filepath string, lines []string, perms ...uint32) (err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}

	var writer *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if writer, err = os.OpenFile(filepath, flags, perm); err != nil {
		err = errors.Wrapf(err, "failed to open file %s for writing", filepath)
		return
	}
	defer writer.Close()

	for i := range lines {
		if _, err = writer.WriteString(lines[i]); err != nil {
			err = errors.Wrapf(err, "failed to write string to %s", filepath)
			return
		}
		if _, err = writer.WriteString("\n"); err != nil {
			err = errors.Wrapf(err, "failed to write newline to %s", filepath)
			return
		}
	}
	err = writer.Sync()

	return
}

// WriteStream reads from the io.Reader and writes to the given file using io.Copy
// thus never filling memory i.e. streaming.  dest will be overwritten if it exists.
func WriteStream(reader io.Reader, filepath string, perms ...uint32) (err error) {
	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}

	var writer *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if writer, err = os.OpenFile(filepath, flags, perm); err != nil {
		err = errors.Wrapf(err, "failed to open file %s for writing", filepath)
		return
	}
	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		err = errors.Wrapf(err, "failed to copy stream data to file %s", filepath)
		return
	}
	if err = writer.Sync(); err != nil {
		err = errors.Wrapf(err, "failed to sync stream to file %s", filepath)
		return
	}

	return
}

// WriteString is a pass through to ioutil.WriteFile with default permissions
func WriteString(filepath string, data string, perms ...uint32) (err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}
	if err = ioutil.WriteFile(filepath, []byte(data), perm); err != nil {
		err = errors.Wrapf(err, "failed write string to file %s", filepath)
		return
	}
	return
}

// WriteYaml converts the given obj interface{} into yaml then writes to disk
// with default permissions. Expects obj to be a structure that github.com/ghodss/yaml understands
func WriteYaml(filepath string, obj interface{}, perms ...uint32) (err error) {
	if filepath, err = Abs(filepath); err != nil {
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

func slice(x []string, i, j int) (result []string) {

	// Convert to postive notation
	if i < 0 {
		i = len(x) + i
	}
	if j < 0 {
		j = len(x) + j
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= len(x) {
		j = len(x) - 1
	}

	// Specifically offsetting j to get an inclusive behavior out of Go
	j++

	// Only operate when indexes are within bounds
	// allow j to be len of s as that is how we include last item
	if i >= 0 && i < len(x) && j >= 0 && j <= len(x) {
		result = x[i:j]
	} else {
		result = []string{}
	}
	return
}
