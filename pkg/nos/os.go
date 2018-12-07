package nos

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

// Copy copies the src to the dst.
// If the src is a file the dst must be a target directory.
// If the src is a dir the dst will be created as a clone of src
func Copy(src, dst string) error {
	return filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			var dstPath string
			dstPath, err = destPath(src, dst)
			fmt.Println(src)
			fmt.Println(dstPath)
			//if err = os.MkdirAll(dst, info.Mode()); err != nil {
			//	return err
			//}
		} else {
			//CopyFile(srcPath, dst)
		}
		return nil
	})
}

// CopyFile a single file from src to dst.
// if dst doesn't exist treat this as a move/rename.
// if dst exists and is a dir copy to the dir
func CopyFile(src, dst string) (err error) {
	var srcInfo os.FileInfo
	var srcPath, dstPath string

	// Get absolute path of source
	if srcPath, err = filepath.Abs(src); err != nil {
		return
	}

	// Check the source for issues
	if srcInfo, err = os.Stat(src); err != nil {
		return
	}

	// Get correct destination path
	if dstPath, err = destPath(srcPath, dst); err != nil {
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

// Exists checks if the given path exists
func Exists(src string) bool {
	if _, err := os.Stat(src); err == nil {
		return true
	}
	return false
}

// IsDir checks if the given path is a directory
func IsDir(src string) bool {
	if info, err := os.Stat(src); err == nil {
		return info.IsDir()
	}
	return false
}

// IsFile checks if the given path is a file
func IsFile(src string) bool {
	if info, err := os.Stat(src); err == nil {
		return !info.IsDir()
	}
	return false
}

// MD5 computes the md5 has of the given file
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

// MkdirP create the target directory and any parent directories needed
func MkdirP(target string, mode ...os.FileMode) error {
	if mode == nil || len(mode) == 0 {
		return os.MkdirAll(target, 0755)
	}
	return os.MkdirAll(target, mode[0])
}

// Get correct destination path.
// If it exists then it is the destination + src base
// If it doesn't exist and parent does then it is the new destination
func destPath(src, dst string) (result string, err error) {
	_, e := os.Stat(dst)
	switch {

	// Doesn't exist but maybe the parent does and this is the new dst name
	case os.IsNotExist(e):
		if _, err = os.Stat(path.Dir(dst)); err != nil {
			return
		}
		if result, err = filepath.Abs(path.Dir(dst)); err != nil {
			return
		}
		result = path.Join(result, path.Base(dst))

	// dst is a valid directory so copy to it using src name
	case e == nil:
		if result, err = filepath.Abs(dst); err != nil {
			return
		}
		result = path.Join(result, path.Base(src))

	// unknown error case
	default:
		err = e
		return
	}

	return
}
