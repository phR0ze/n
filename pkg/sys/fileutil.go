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
	"strings"

	"github.com/phR0ze/n/pkg/opt"
	"github.com/pkg/errors"
)

// Chmod wraps os.Chmod but provides path expansion, globbing, recursion and error tracing.
// For each resulting path if the file is a symbolic link, it changes the mode of the link's
// target. Recursively apply chmod to all files and directories by passing in RecurseOpt(true)
// Apply chmod to only directories or files with OnlyDirsOpt(true) and/or OnlyFilesOpt(true)
func Chmod(path string, mode os.FileMode, opts ...*opt.Opt) (err error) {
	recurse := getRecurseOpt(opts)
	onlyDirs := getOnlyDirsOpt(opts)
	onlyFiles := getOnlyFilesOpt(opts)

	// Path expansion
	if path, err = Abs(path); err != nil {
		return
	}

	// Handle globbing
	var sources []string
	if sources, err = filepath.Glob(path); err != nil {
		err = errors.Wrapf(err, "failed to get glob for %s", path)
		return
	}

	// Fail if no sources were found
	if len(sources) == 0 {
		err = errors.Errorf("failed to get any sources for %s", path)
		return
	}

	// Execute the chmod for all sources
	for _, source := range sources {

		// Only check if dir if we are required to
		isDir := false
		if onlyDirs || onlyFiles || recurse {
			isDir = IsDir(source)
		}

		// Only get old mode if we have to
		var oldMode os.FileMode
		if recurse {
			oldMode = Mode(source)
		}

		// We have to be careful of the order of applying permissions we'll get into a
		// scenario where we are revoking read/execute on a dir before we get to the
		// bottom of the stack. Like wise when adding permissions we have to
		// do it on the way in or we won't be able to read to get there.
		if (!onlyDirs && !onlyFiles) || (onlyDirs && isDir) || (onlyFiles && !isDir) {

			// Chmod on the way in if not recursing or recursing and adding permissions
			if !recurse || !isDir || (recurse && !revokingMode(oldMode, mode)) {
				if err = os.Chmod(source, mode); err != nil {
					err = errors.Wrapf(err, "failed to add permissions with chmod %s", path)
					return
				}
			}
		}

		// Handle recursion only one dir at a time as permissions need set first
		// incase we are adding read/execute permissions as we go.
		if recurse && isDir {
			for _, path := range Paths(source) {
				if err = Chmod(path, mode, opts...); err != nil {
					return
				}
			}
		}

		// Chmod on the way out if recursing and revoking permissions
		if (!onlyDirs && !onlyFiles) || (onlyDirs && isDir) || (onlyFiles && !isDir) {
			if recurse && isDir && revokingMode(oldMode, mode) {
				if err = os.Chmod(source, mode); err != nil {
					err = errors.Wrapf(err, "failed to revoke permissions with chmod %s", path)
					return
				}
			}
		}
	}

	return
}

// determine if the mode change is revoking permissions or adding permissions.
// only taking into account read/execute on the directory
func revokingMode(old, new os.FileMode) bool {
	return old&0500 > new&0500 || old&0050 > new&0050 || old&0005 > new&0005
}

// Chown wraps os.Chown but provides path expansion, globbing, recursion and error tracing.
// For each resulting path change the numeric uid and gid. If the file is a symbolic link,
// it changes the uid and gid of the link's target. A uid or gid of -1 means to not change
// that value. Recursively apply chown to all files and directories by passing in RecurseOpt(true)
func Chown(path string, uid, gid int, opts ...*opt.Opt) (err error) {

	// Glob the path
	var paths []string
	if paths, err = Glob(path, opts...); err != nil {
		return
	}

	// Fail if no sources were found
	if len(paths) == 0 {
		err = errors.Errorf("failed to get any sources for %s", path)
		return
	}

	// Execute the chown for all sources
	for _, path := range paths {
		if err = os.Chown(path, uid, gid); err != nil {
			err = errors.Wrapf(err, "failed to chown %s", path)
			return
		}
	}
	return
}

// Copy copies src to dst recursively, creating destination directories as needed.
// Handles globbing e.g. Copy("./*", "../")
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist.
// Doesn't follow links by default but can be turned by passing in FollowOpt(true)
func Copy(src, dst string, opts ...*opt.Opt) (err error) {
	clone := true
	var sources []string

	// Trim trailing slashes
	src = strings.TrimSuffix(src, "/")

	// Set following links to false by default
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

	// Fail no sources were found
	if len(sources) == 0 {
		err = errors.Errorf("failed to get any sources for %s", srcAbs)
		return
	}

	// Clone given src as dst vs copy into dst
	if IsDir(dstAbs) || len(sources) > 1 {
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

			// Handle individual copies
			switch {

			// Create destination directories as needed
			case srcInfo.IsDir():
				if e = os.MkdirAll(dstPath, srcInfo.Mode()); e != nil {
					return e
				}

			// Copy dir links
			case srcInfo.IsSymlinkDir():
				var target string
				if target, e = srcInfo.SymlinkTarget(); e != nil {
					return e
				}
				if e = os.Symlink(target, dstPath); e != nil {
					return e
				}

			// Copy file
			default:
				CopyFile(srcPath, dstPath, InfoOpt(srcInfo))
			}
			return nil
		}, opts...)
	}
	return
}

// CopyFile copies a single file from src to dsty, creating destination directories as needed.
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist.
// Supports passing in the FileInfo object directly with FollowOpt(true)
// Returns the destination path for copied file
func CopyFile(src, dst string, opts ...*opt.Opt) (result string, err error) {
	var srcPath, dstPath string
	var srcInfo, srcDirInfo *FileInfo

	// Set following links to false by default
	defaultFollowOpt(&opts, false)

	// Check the source for issues
	if srcInfo = getInfoOpt(opts); srcInfo != nil {
		if srcPath, err = Abs(srcInfo.Path); err != nil {
			return
		}
	} else {
		if srcPath, err = Abs(src); err != nil {
			return
		}
		if srcInfo, err = Lstat(srcPath); err != nil {
			return
		}
	}

	// Source dir permissions to use for destination directories
	if srcDirInfo, err = Lstat(path.Dir(srcPath)); err != nil {
		return
	}

	// Error out if not a regular file or symlink
	if srcInfo.IsDir() || srcInfo.IsSymlinkDir() {
		err = errors.Errorf("src target is not a regular file or a symlink to a file")
		return
	}

	// Get correct destination path
	if dstPath, err = Abs(dst); err != nil {
		return
	}
	dstInfo, e := os.Stat(dstPath)
	switch {

	// Doesn't exist so this is the new destination name, ensure all paths exist
	case os.IsNotExist(e):
		if err = os.MkdirAll(path.Dir(dstPath), srcDirInfo.Mode()); err != nil {
			return
		}

	// Destination exists and is either a file to overwrite or a dir to copy into
	case e == nil:
		if dstInfo.IsDir() {
			dstPath = path.Join(dstPath, path.Base(srcPath))
		}

	// unknown error case
	default:
		err = errors.Wrapf(e, "failed to Stat destination %s", dst)
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
		var fr *os.File
		if fr, err = os.Open(srcPath); err != nil {
			err = errors.Wrapf(err, "failed to open file %s for reading", srcPath)
			return
		}
		defer fr.Close()

		// Create dstPath for writing
		var fw *os.File
		if fw, err = os.Create(dstPath); err != nil {
			err = errors.Wrapf(err, "failed to create file %s", dstPath)
			return
		}

		// Copy srcPath to dstPath
		if _, err = io.Copy(fw, fr); err != nil {
			err = errors.Wrapf(err, "failed to copy data to file %s", dstPath)
			if e := fw.Close(); e != nil {
				err = errors.Wrapf(err, "failed to close file %s", dstPath)
			}
			return
		}

		// Sync to disk
		if err = fw.Sync(); err != nil {
			err = errors.Wrapf(err, "failed to sync data to file %s", dstPath)
			if e := fw.Close(); e != nil {
				err = errors.Wrapf(err, "failed to close file %s", dstPath)
			}
			return
		}

		// Close file for writing
		if err = fw.Close(); err != nil {
			err = errors.Wrapf(err, "failed to close file %s", dstPath)
			return
		}

		// Set permissions of dstPath same as srcPath
		if err = os.Chmod(dstPath, srcInfo.Mode()); err != nil {
			err = errors.Wrapf(err, "failed to chmod file %s", dstPath)
			return
		}
	}

	result = dstPath
	return
}

// Exists return true if the given path exists
func Exists(src string) bool {
	if target, err := Abs(src); err == nil {
		if _, err := os.Stat(target); err != nil {

			// If we got permission denied then the file probably exists
			if strings.HasSuffix(err.Error(), ": permission denied") {
				return true
			}

		} else {
			return true
		}
	}
	return false
}

// MkdirP creates the target directory and any parent directories needed
// and returns the ABS path of the created directory
func MkdirP(dirname string, perms ...uint32) (dir string, err error) {
	if dir, err = Abs(dirname); err != nil {
		return
	}

	// Get/set default permission
	perm := os.FileMode(0755)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}

	// Create directory
	if err = os.MkdirAll(dir, perm); err != nil {
		err = errors.Wrapf(err, "failed creating directories for %s", dir)
		return
	}

	return
}

// MD5 returns the md5 of the given file
func MD5(filename string) (result string, err error) {
	if filename, err = Abs(filename); err != nil {
		return
	}
	if !Exists(filename) {
		return "", os.ErrNotExist
	}

	// Open file for reading
	var fr *os.File
	if fr, err = os.Open(filename); err != nil {
		err = errors.Wrapf(err, "failed opening target file %s", filename)
		return
	}
	defer fr.Close()

	// Create a new md5 hash and copy in file bits
	hash := md5.New()
	if _, err = io.Copy(hash, fr); err != nil {
		err = errors.Wrapf(err, "failed copying file data into hash from %s", filename)
		return
	}

	// Compute 32 byte hash
	result = hex.EncodeToString(hash.Sum(nil))

	return
}

// Move the src path to the dst path. If the dst already exists and is not a directory
// src will replace it. If there is an error it will be of type *LinkError. Wraps
// os.Rename but fixes the issue where dst name is required. Returns the new location
func Move(src, dst string) (result string, err error) {

	// Add src base name to dst directory to fix golang oversight
	if IsDir(dst) {
		dst = path.Join(dst, path.Base(src))
	}
	if err = os.Rename(src, dst); err != nil {
		err = errors.Wrapf(err, "failed renaming file %s", src)
		return
	}
	result = dst
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
		err = errors.Wrapf(err, "failed reading the file %s", filepath)
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
		err = errors.Wrapf(err, "failed reading the file %s", filepath)
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}

// ReadLinesP returns a new slice of string representing lines
func ReadLinesP(reader io.Reader) (result []string) {
	scanner := bufio.NewScanner(reader)
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
		err = errors.Wrapf(err, "failed reading the file %s", filepath)
		return
	}
	result = string(data)
	return
}

// Remove the given target file or empty directory. If there is an
// error it will be of type *PathError
func Remove(target string) (err error) {
	if err = os.Remove(target); err != nil {
		err = errors.Wrapf(err, "failed removing %s", target)
		return
	}
	return
}

// RemoveAll removes the target path and any children it contains. It will
// retry the operation after attempting to correct file permissions on failure.
// If the target path does not exist nil is returned
func RemoveAll(target string) (err error) {
	if err = os.RemoveAll(target); err != nil {
		Chmod(target, 0777, RecurseOpt(true))
		if err = os.RemoveAll(target); err != nil {
			err = errors.Wrapf(err, "failed removing %s", target)
		}
	}
	return
}

// Symlink creates newname as a symbolic link to link. If there is an error,
// it will be of type *LinkError.
func Symlink(src, link string) error {
	return os.Symlink(src, link)
}

// Touch creates an empty text file similar to the linux touch command
func Touch(filepath string) (path string, err error) {
	if path, err = Abs(filepath); err != nil {
		return
	}

	var fw *os.File
	if fw, err = os.Create(path); err != nil {
		err = errors.Wrapf(err, "failed creating/truncating file %s", filepath)
		return
	}

	// Ignoring close in the error case above is ok as the file pointer will be nil
	if err = fw.Close(); err != nil {
		err = errors.Wrapf(err, "failed closing file %s", filepath)
		return
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
		err = errors.Wrapf(err, "failed writing bytes to file %s", filepath)
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
	if err = ioutil.WriteFile(filepath, []byte(strings.Join(lines, "\n")), perm); err != nil {
		err = errors.Wrapf(err, "failed writing lines to file %s", filepath)
		return
	}
	return
}

// WriteStream reads from the io.Reader and writes to the given file using io.Copy
// thus never filling memory i.e. streaming.  dest will be overwritten if it exists.
func WriteStream(reader io.Reader, filepath string, perms ...uint32) (err error) {
	if filepath, err = Abs(filepath); err != nil {
		return
	}

	perm := os.FileMode(0644)
	if len(perms) > 0 {
		perm = os.FileMode(perms[0])
	}

	var fw *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if fw, err = os.OpenFile(filepath, flags, perm); err != nil {
		err = errors.Wrapf(err, "failed opening file %s for writing", filepath)
		return
	}

	if _, err = io.Copy(fw, reader); err != nil {
		err = errors.Wrap(err, "failed copying stream data")
		if e := fw.Close(); e != nil {
			err = errors.Wrapf(err, "failed to close file %s", filepath)
		}
		return
	}
	if err = fw.Sync(); err != nil {
		err = errors.Wrapf(err, "failed syncing stream to file %s", filepath)
		if e := fw.Close(); e != nil {
			err = errors.Wrapf(err, "failed to close file %s", filepath)
		}
		return
	}

	if err = fw.Close(); err != nil {
		err = errors.Wrapf(err, "failed to close file %s", filepath)
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
		err = errors.Wrapf(err, "failed writing string to file %s", filepath)
		return
	}
	return
}
