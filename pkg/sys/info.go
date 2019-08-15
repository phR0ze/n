package sys

import (
	"os"
	"path/filepath"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// FileInfo wraps the os.FileInfo interface and provide additional helper functions
type FileInfo struct {
	v os.FileInfo

	path string // path used to create the info file
}

// Lstat wraps os.Lstate to give back a FileInfo
// Resolves home dir and relative dir pathing into absolute paths
func Lstat(src string) (result *FileInfo, err error) {
	result = &FileInfo{path: src}
	if result.v, err = os.Lstat(src); err != nil {
		result = nil
		err = errors.Wrapf(err, "failed to execute Lstate against %s", src)
	}
	return
}

// Name implements os.FileInfo and returns the base name of the file
func (info *FileInfo) Name() string {
	return info.v.Name()
}

// Path returns the absolute path for this file
func (info *FileInfo) Path() string {
	return info.path
}

// AbsPath returns the absolute path for this file
func (info *FileInfo) AbsPath() (path string, err error) {
	if path, err = homedir.Expand(info.path); err != nil {
		return
	}
	path, err = filepath.Abs(path)
	return
}

// Size implements os.FileInfo and returns the size of file in bytes
func (info *FileInfo) Size() int64 {
	return info.v.Size()
}

// Size returns the size of the file in bytes
func Size(src string) (size int64) {
	if info, err := os.Lstat(src); err == nil {
		if !info.IsDir() {
			size = info.Size()
		}
	}
	return
}

// Mode implements os.FileInfo and returns bits of the file
func (info *FileInfo) Mode() os.FileMode {
	return info.v.Mode()
}

// ModeTime implements os.FileInfo and is the modification time of the file
func (info *FileInfo) ModeTime() time.Time {
	return info.v.ModTime()
}

// Sys implements os.FileInfo and provides access to the underlying data source
func (info *FileInfo) Sys() interface{} {
	return info.v.Sys()
}

// IsDir returns true if the info is a directory
func (info *FileInfo) IsDir() bool {
	return info.v.IsDir()
}

// IsDir returns true if the given path is a directory
func IsDir(src string) bool {
	if info, err := os.Lstat(src); err == nil {
		return info.IsDir()
	}
	return false
}

// IsFile returns true if the info is a file
func (info *FileInfo) IsFile() bool {
	return !info.v.IsDir() && info.v.Mode()&os.ModeSymlink == 0
}

// IsFile returns true if the given path is a file
func IsFile(src string) bool {
	if info, err := os.Lstat(src); err == nil {
		return !info.IsDir() && info.Mode()&os.ModeSymlink == 0
	}
	return false
}

// IsSymlink returns true if the info is a symlink
func (info *FileInfo) IsSymlink() bool {
	return info.v.Mode()&os.ModeSymlink != 0
}

// IsSymlink returns true if the given path is a symlink
func IsSymlink(src string) bool {
	if info, err := Lstat(src); err == nil {
		return info.IsSymlink()
	}
	return false
}

// IsSymlinkDir returns true if the symlink's target is a directory
func (info *FileInfo) IsSymlinkDir() bool {
	if info.v.Mode()&os.ModeSymlink != 0 {
		if target, err := filepath.EvalSymlinks(info.path); err == nil {
			if subinfo, err := os.Lstat(target); err == nil {
				if subinfo.IsDir() {
					return true
				}
			}
		}
	}
	return false
}

// IsSymlinkDir returns true if the given symlink's target is a directory
func IsSymlinkDir(src string) bool {
	if info, err := Lstat(src); err == nil {
		return info.IsSymlinkDir()
	}
	return false
}

// IsSymlinkFile returns true if the symlink's target is a file
func (info *FileInfo) IsSymlinkFile() bool {
	if info.v.Mode()&os.ModeSymlink != 0 {
		if target, err := filepath.EvalSymlinks(info.path); err == nil {
			if subinfo, err := os.Lstat(target); err == nil {
				if !subinfo.IsDir() {
					return true
				}
			}
		}
	}
	return false
}

// IsSymlinkFile returns true if the given symlink's target is a directory
func IsSymlinkFile(src string) bool {
	if info, err := Lstat(src); err == nil {
		return info.IsSymlinkFile()
	}
	return false
}

// SymlinkTarget follows the symlink to get the path for the target
func (info *FileInfo) SymlinkTarget() (target string, err error) {
	if info.v.Mode()&os.ModeSymlink == 0 {
		err = errors.Errorf("not a symlink")
		return
	}
	if target, err = os.Readlink(info.path); err != nil {
		return
	}
	return
}

// SymlinkTarget follows the symlink to get the path for the target
func SymlinkTarget(src string) (target string, err error) {
	var info *FileInfo
	if info, err = Lstat(src); err != nil {
		return
	}
	return info.SymlinkTarget()
}
