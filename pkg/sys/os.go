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
	"runtime"
	"strings"

	"github.com/phR0ze/n/pkg/opt"
)

// Copy copies src to dst recursively.
// Handles globbing e.g. Copy("./*", "../")
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist, but it's parent directory does.
// Follows links by default but can be turned off with &Opt{"links", false}
func Copy(src, dst string, opts ...*opt.Opt) (err error) {
	var sources []string

	// Set following links by default
	if links := opt.Find(opts, "links"); links == nil {
		opts = append(opts, &opt.Opt{Key: "links", Val: true})
	}

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
		return err
	}

	// Copy all sources to dst
	for _, srcRoot := range sources {

		// Clone given src as dst vs copy into dst
		clone := true
		if IsDir(dstAbs) {
			clone = false
		}

		// Walk over file structure
		err = Walk(srcRoot, func(srcPath string, info *FileInfo, e error) error {
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

			// Copy to destination
			if info.IsDir() {
				if e = os.MkdirAll(dstPath, info.Mode()); e != nil {
					return e
				}
			} else {
				CopyFile(srcPath, dstPath)
			}
			return nil
		})
	}
	return
}

// CopyFile copies a single file from src to dst.
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist, but it's parent directory does.
func CopyFile(src, dst string) (err error) {
	var srcInfo os.FileInfo
	var srcPath, dstPath string

	// Get absolute path of source
	if srcPath, err = Abs(src); err != nil {
		return
	}

	// Check the source for issues
	if srcInfo, err = os.Stat(src); err != nil {
		return
	}

	// Get correct destination path
	_, e := os.Stat(dst)
	switch {

	// Doesn't exist but maybe the parent does and this is the new dst name
	case os.IsNotExist(e):
		if _, err = os.Stat(path.Dir(dst)); err != nil {
			return
		}
		if dstPath, err = Abs(path.Dir(dst)); err != nil {
			return
		}
		dstPath = path.Join(dstPath, path.Base(dst))

	// dst is a valid directory so copy to it using src name
	case e == nil:
		if dstPath, err = Abs(dst); err != nil {
			return
		}
		dstPath = path.Join(dstPath, path.Base(srcPath))

	// unknown error case
	default:
		err = e
		return
	}

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
		return
	}
	defer f.Close()

	// Create a new md5 hash and copy in file bits
	hash := md5.New()
	if _, err = io.Copy(hash, f); err != nil {
		return
	}

	// Compute 32 byte hash
	result = hex.EncodeToString(hash.Sum(nil))

	return
}

// MkdirP creates the target directory and any parent directories needed
func MkdirP(target string, mode ...os.FileMode) (err error) {
	if target, err = Abs(target); err != nil {
		return
	}

	if mode == nil || len(mode) == 0 {
		return os.MkdirAll(target, 0755)
	}
	return os.MkdirAll(target, mode[0])
}

// Move the src path to the dst path. If the dst already exists and is not a directory
// src will replace it. If there is an error it will be of type *LinkError. Wraps
// os.Rename but fixes the issue where dst name is required
func Move(src, dst string) (err error) {

	// Add src base name to dst directory to fix golang oversight
	if IsDir(dst) {
		dst = path.Join(dst, path.Base(src))
	}
	return os.Rename(src, dst)
}

// Pwd returns the current working directory
func Pwd() (pwd string) {
	pwd, _ = os.Getwd()
	return
}

// ReadLines returns a new slice of string representing lines
func ReadLines(target string) (result []string, err error) {
	if target, err = Abs(target); err != nil {
		return
	}

	var fileBytes []byte
	if fileBytes, err = ioutil.ReadFile(target); err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(fileBytes))
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
	}
	return
}

// Remove the given target file or empty directory. If there is an
// error it will be of type *PathError
func Remove(target string) error {
	return os.Remove(target)
}

// RemoveAll remotes the target path and any children it contains.
// It removes everything it can but returns the first error it encounters.
// If the target path does not exist nil is returned
func RemoveAll(target string) error {
	return os.RemoveAll(target)
}

// Size returns the size of the given path file or directory
func Size(src string) (result int64) {
	if target, err := Abs(src); err == nil {
		if info, err := os.Stat(target); err == nil {
			if !info.IsDir() {
				result = info.Size()
			}
		}
	}
	return
}

// Touch creates an empty text file similar to the linux touch command
func Touch(target string) (err error) {
	if target, err = Abs(target); err != nil {
		return
	}

	var f *os.File
	if f, err = os.Create(target); !os.IsExist(err) {
		return
	}
	if err == nil {
		defer f.Close()
	}
	return
}

// WriteFile is a pass through to ioutil.WriteFile with default permissions
func WriteFile(target string, data []byte, perms ...uint32) (err error) {
	if target, err = Abs(target); err != nil {
		return
	}

	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}
	err = ioutil.WriteFile(target, data, perm)
	return
}

// WriteStream reads from the io.Reader and writes to the given file using io.Copy
// thus never filling memory i.e. streaming.  dest will be overwritten if it exists.
func WriteStream(reader io.Reader, dest string, perms ...uint32) (err error) {
	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}

	var writer *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if writer, err = os.OpenFile(dest, flags, perm); err != nil {
		return
	}
	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return
	}
	err = writer.Sync()

	return
}

// WriteLines is a pass through to ioutil.WriteFile with default permissions
func WriteLines(target string, lines []string, perms ...uint32) (err error) {
	if target, err = Abs(target); err != nil {
		return
	}

	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}

	var writer *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if writer, err = os.OpenFile(target, flags, perm); err != nil {
		return
	}
	defer writer.Close()

	for i := range lines {
		if _, err = writer.WriteString(lines[i]); err != nil {
			return
		}
		if _, err = writer.WriteString("\n"); err != nil {
			return
		}
	}
	err = writer.Sync()

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
