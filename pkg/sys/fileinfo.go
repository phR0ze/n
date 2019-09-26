package sys

import (
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

// FileInfo wraps the os.FileInfo interface and provide additional helper functions
type FileInfo struct {
	Path string      // absolute path to the file set when created
	Obj  os.FileInfo // handle on the actual OS object to use where needed
}

// Lstat wraps os.Lstate to give back a FileInfo
// Resolves home dir and relative dir pathing into absolute paths
func Lstat(src string) (result *FileInfo, err error) {
	if src, err = Abs(src); err != nil {
		return
	}
	result = &FileInfo{Path: src}
	if result.Obj, err = os.Lstat(src); err != nil {
		result = nil
		err = errors.Wrapf(err, "failed to execute Lstat against %s", src)
	}
	return
}

// Name implements os.FileInfo and returns the base name of the file
func (info *FileInfo) Name() string {
	return info.Obj.Name()
}

// Size implements os.FileInfo and returns the size of file in bytes
func (info *FileInfo) Size() int64 {
	return info.Obj.Size()
}

// Size returns the size of the file/dir in bytes
func Size(src string) (size int64) {
	if info, err := Lstat(src); err == nil {
		size = info.Size()
	}
	return
}

// Mode implements os.FileInfo and returns bits of the file
func (info *FileInfo) Mode() os.FileMode {
	return info.Obj.Mode()
}

// Mode implements os.FileInfo and returns bits of the file
func Mode(src string) (mode os.FileMode) {
	if info, err := Lstat(src); err == nil {
		mode = info.Mode()
	}
	return
}

// ModTime implements os.FileInfo and is the modification time of the file
func (info *FileInfo) ModTime() time.Time {
	return info.Obj.ModTime()
}

// Sys implements os.FileInfo and provides access to the underlying data source
func (info *FileInfo) Sys() interface{} {
	return info.Obj.Sys()
}

// IsDir returns true if the info is a directory
func (info *FileInfo) IsDir() bool {
	return info.Obj.IsDir()
}

// IsDir returns true if the given path is a directory
func IsDir(src string) bool {
	if info, err := Lstat(src); err == nil {
		return info.IsDir()
	}
	return false
}

// IsFile returns true if the info is a file
func (info *FileInfo) IsFile() bool {
	return !info.Obj.IsDir() && info.Obj.Mode()&os.ModeSymlink == 0
}

// IsFile returns true if the given path is a file
func IsFile(src string) bool {
	if info, err := Lstat(src); err == nil {
		return !info.IsDir() && info.Mode()&os.ModeSymlink == 0
	}
	return false
}

// IsSymlink returns true if the info is a symlink
func (info *FileInfo) IsSymlink() bool {
	return info.Obj.Mode()&os.ModeSymlink != 0
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
	if info.Obj.Mode()&os.ModeSymlink != 0 {
		if target, err := filepath.EvalSymlinks(info.Path); err == nil {
			if subinfo, err := Lstat(target); err == nil {
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
	if info.Obj.Mode()&os.ModeSymlink != 0 {
		if target, err := filepath.EvalSymlinks(info.Path); err == nil {
			if subinfo, err := Lstat(target); err == nil {
				if !subinfo.IsDir() {
					return true
				}
			}
		}
	}
	return false
}

// IsSymlinkFile returns true if the given symlink's target is a file
func IsSymlinkFile(src string) bool {
	if info, err := Lstat(src); err == nil {
		return info.IsSymlinkFile()
	}
	return false
}

// SymlinkTarget follows the symlink to get the path for the target
func (info *FileInfo) SymlinkTarget() (target string, err error) {
	if info.Obj.Mode()&os.ModeSymlink == 0 {
		err = errors.Errorf("not a symlink")
		return
	}
	if target, err = os.Readlink(info.Path); err != nil {
		err = errors.Errorf("failed to read the link target")
		return
	}
	return
}

// SymlinkTarget follows the symlink to get the path for the target.
// Will get the path regardless if the target actually exists.
func SymlinkTarget(src string) (target string, err error) {
	var info *FileInfo
	if info, err = Lstat(src); err != nil {
		return
	}
	return info.SymlinkTarget()
}

// SymlinkTargetExists returns true if the symlink's target exists
func (info *FileInfo) SymlinkTargetExists() bool {
	if info.Obj.Mode()&os.ModeSymlink != 0 {
		if _, err := filepath.EvalSymlinks(info.Path); err == nil {
			return true
		}
	}
	return false
}

// SymlinkTargetExists returns true if the symlink's target exists
func SymlinkTargetExists(src string) bool {
	info, err := Lstat(src)
	if err != nil {
		return false
	}
	return info.SymlinkTargetExists()
}
