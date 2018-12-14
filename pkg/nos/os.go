package nos

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
	"strings"
)

// Copy copies src to dst recursively.
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist, but it's parent directory does.
func Copy(src, dst string) error {
	var e error
	var srcRoot, dstRoot string
	if srcRoot, e = Abs(src); e != nil {
		return e
	}
	if dstRoot, e = Abs(dst); e != nil {
		return e
	}

	// Clone given src as dst vs copy into dst
	clone := true
	if IsDir(dstRoot) {
		clone = false
	}

	// Walk over file structure
	return filepath.Walk(srcRoot, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Set proper dst path
		var dstPath string
		if clone {
			dstPath = path.Join(dstRoot, strings.TrimPrefix(srcPath, srcRoot))
		} else {
			dstPath = path.Join(dstRoot, strings.TrimPrefix(srcPath, path.Dir(srcRoot)))
		}

		// Copy to destination
		if info.IsDir() {
			if err = os.MkdirAll(dstPath, info.Mode()); err != nil {
				return err
			}
		} else {
			CopyFile(srcPath, dstPath)
		}
		return nil
	})
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

	// Set permissions of dstPath same as srcPath
	err = os.Chmod(dstPath, srcInfo.Mode())

	return
}

// Exists return true if the given path exists
func Exists(src string) bool {
	if _, err := os.Stat(src); err == nil {
		return true
	}
	return false
}

// IsDir returns true if the given path is a directory
func IsDir(src string) bool {
	if info, err := os.Stat(src); err == nil {
		return info.IsDir()
	}
	return false
}

// IsFile returns true if the given path is a file
func IsFile(src string) bool {
	if info, err := os.Stat(src); err == nil {
		return !info.IsDir()
	}
	return false
}

// MD5 returns the md5 of the given file
func MD5(target string) (result string, err error) {
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
func MkdirP(target string, mode ...os.FileMode) error {
	if mode == nil || len(mode) == 0 {
		return os.MkdirAll(target, 0755)
	}
	return os.MkdirAll(target, mode[0])
}

// ReadLines returns a new slice of string representing lines
func ReadLines(target string) (result []string, err error) {
	var fileBytes []byte
	if fileBytes, err = ioutil.ReadFile(target); err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(fileBytes))
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
	}
	return
}

// Touch creates an empty text file similar to the linux touch command
func Touch(target string) (err error) {
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
	perm := uint32(0644)
	if len(perms) > 0 {
		perm = perms[0]
	}
	err = ioutil.WriteFile(target, data, os.FileMode(perm))
	return
}

// WriteLines is a pass through to ioutil.WriteFile with default permissions
func WriteLines(target string, lines []string, perms ...uint32) (err error) {
	perm := uint32(0644)
	if len(perms) > 0 {
		perm = perms[0]
	}

	var f *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if f, err = os.OpenFile(target, flags, os.FileMode(perm)); err != nil {
		return
	}
	defer f.Close()

	for i := range lines {
		if _, err = f.WriteString(lines[i]); err != nil {
			return
		}
		if _, err = f.WriteString("\n"); err != nil {
			return
		}
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
