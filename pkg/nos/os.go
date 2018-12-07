package nos

import (
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
		if !info.IsDir() {
			CopyFile(srcPath, dst)
		} else {
			fmt.Println("Dir: ", srcPath)
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
	_, e := os.Stat(dst)
	switch {

	// Doesn't exist but maybe the parent does and this is the new file name
	case os.IsNotExist(e):
		if _, err = os.Stat(path.Dir(dst)); err != nil {
			return
		}
		if dstPath, err = filepath.Abs(path.Dir(dst)); err != nil {
			return
		}
		dstPath = path.Join(dstPath, path.Base(dst))

	// dst is a valid directory so copy to it using src name
	case e == nil:
		if dstPath, err = filepath.Abs(dst); err != nil {
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
