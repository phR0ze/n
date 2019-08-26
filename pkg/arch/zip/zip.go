// Package zip provides create and extract implementations
package zip

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/phR0ze/n/pkg/enc/bin"
	"github.com/phR0ze/n/pkg/sys"
	"github.com/pkg/errors"
)

var (
	// Zip files typically start with PK i.e. 50 4B
	gZipHeaderSig = []byte{0x50, 0x4B}
)

// Create a new zip file at zipfile from the given srcPath directory
func Create(zipfile, srcPath string) (err error) {
	if zipfile, err = sys.Abs(zipfile); err != nil {
		return
	}
	if srcPath, err = sys.Abs(srcPath); err != nil {
		return
	}

	// Create the new file for writing to
	var fw *os.File
	if fw, err = os.Create(zipfile); err != nil {
		err = errors.Wrapf(err, "failed to create zipfile %s", zipfile)
		return
	}
	defer fw.Close()

	// Open zip writer
	zw := zip.NewWriter(fw)
	defer zw.Close()

	// Add all files recursively
	if err = addFiles(zw, srcPath, ""); err != nil {
		return
	}
	return
}

// AddFiles to the given zip writer recursively where root is the directory
// to recurse on and base is the path the zip files should be based on in the zip
func addFiles(zw *zip.Writer, root, base string) (err error) {
	var infos []os.FileInfo
	if infos, err = ioutil.ReadDir(root); err != nil {
		err = errors.Wrapf(err, "failed to read directory %s to add files from", root)
		return
	}
	for _, info := range infos {
		if !info.IsDir() {
			target := path.Join(root, info.Name())

			// Open the target file for reading
			var fr *os.File
			if fr, err = os.Open(target); err != nil {
				err = errors.Wrapf(err, "failed to open target file %s for zip", target)
				return
			}
			defer fr.Close()

			// Add the files to the zip
			var fw io.Writer
			zipPath := path.Join(base, info.Name())
			if fw, err = zw.Create(zipPath); err != nil {
				err = errors.Wrapf(err, "failed to add target file %s to zip", target)
				return
			}

			// Stream the data from the reader to the writer
			if _, err = io.Copy(fw, fr); err != nil {
				err = errors.Wrapf(err, "failed to copy data from reader to writer for zip target %s", target)
				return
			}
		} else {
			newRoot := path.Join(root, info.Name())
			newBase := path.Join(base, info.Name())
			addFiles(zw, newRoot, newBase)
		}
	}

	return
}

// ExtractAll files into given destination directory
func ExtractAll(zipfile, dest string) (err error) {
	if zipfile, err = sys.Abs(zipfile); err != nil {
		return
	}
	if dest, err = sys.MkdirP(dest); err != nil {
		return
	}

	// TrimPrefix from the zip file if required
	if err = TrimPrefix(zipfile); err != nil {
		return
	}

	// Open zipfile for use
	var zr *zip.ReadCloser
	if zr, err = zip.OpenReader(zipfile); err != nil {
		err = errors.Wrapf(err, "failed to open zipfile %s for reading", zipfile)
		return
	}
	defer zr.Close()

	// Extract all files
	dirCache := map[string]bool{}
	for _, file := range zr.File {
		info := file.FileInfo()
		filePath := path.Join(dest, file.Name)

		// Create any directories with default mode
		dirPath := path.Dir(filePath)
		if info.IsDir() {
			dirPath = filePath
		}
		if _, exist := dirCache[dirPath]; !exist {
			if _, err = sys.MkdirP(dirPath); err != nil {
				err = errors.Wrapf(err, "failed to create directory %s", dirPath)
				return
			}
			dirCache[dirPath] = true
		}

		// Create file and write content to it
		if !info.IsDir() {
			var fw *os.File
			if fw, err = os.Create(filePath); err != nil {
				err = errors.Wrapf(err, "failed to create new file %s from zipfile", filePath)
				return
			}
			var fr io.ReadCloser
			if fr, err = file.Open(); err != nil {
				err = errors.Wrapf(err, "failed to open zip file target %s for reading", info.Name())
				return
			}
			if _, err = io.Copy(fw, fr); err != nil {
				err = errors.Wrap(err, "failed to copy data from zip to disk")
				return
			}
			fr.Close()
			fw.Close()

			// Set file mode to the original value
			if err = os.Chmod(filePath, info.Mode()); err != nil {
				err = errors.Wrapf(err, "failed to set original file mode for %s", filePath)
				return
			}
		}

		// Set file access times to the original values
		if err = os.Chtimes(filePath, info.ModTime(), info.ModTime()); err != nil {
			err = errors.Wrapf(err, "failed to set original file access times for %s", filePath)
			return
		}
	}
	return
}

// TrimPrefix simple drops the bytes up to the begining of the zipfile.
// Many custom zip files like the chromium crx extension file have had additional data prefixed
// to them that  needs to be stripped off before the zip can be processed.
func TrimPrefix(zipfile string) (err error) {
	if zipfile, err = sys.Abs(zipfile); err != nil {
		return
	}

	// Open the zipfile for reading
	var rw *os.File
	if rw, err = os.OpenFile(zipfile, os.O_RDWR, 0644); err != nil {
		err = errors.Wrapf(err, "failed to open zip file '%s' to detect zip data", path.Base(zipfile))
		return
	}
	defer rw.Close()

	// Detect begining of zip file identified by PK.. i.e. 50 4B
	chunk := make([]byte, 1024)
	loc, roffset := int64(-1), int64(-1)
	for {
		l, e := rw.Read(chunk)
		if e != nil {
			if e != io.EOF {
				err = errors.Wrapf(err, "failed to read from zipfile '%s' for zip identification", path.Base(zipfile))
				return
			} else if roffset == -1 {
				err = errors.Wrapf(err, "unable to identify '%s' as a valid zipfile", path.Base(zipfile))
				return
			}
			break
		}
		data := chunk[0:l]
		for i, b := range data {
			loc++
			if b == 0x50 {
				if i+1 < len(data) && bin.Equal(data[i:i+2], gZipHeaderSig) {
					roffset = loc
					break
				}
			}
		}
		if roffset != -1 {
			break
		}
	}

	// Now shift the data if required
	if roffset != 0 {
		woffset := int64(0)
		for {
			if l, e := rw.ReadAt(chunk, roffset); e != nil && e != io.EOF {
				err = errors.Wrapf(err, "failed to read from zipfile '%s' to shift data", path.Base(zipfile))
				return
			} else if l > 0 {
				data := chunk[0:l]
				if _, e := rw.WriteAt(data, woffset); e != nil {
					err = errors.Wrapf(err, "failed to write shifted data to zipfile '%s'", path.Base(zipfile))
					return
				}
				roffset += int64(l)
				woffset += int64(l)
				if e == io.EOF {
					break
				}
			} else {
				break
			}
		}

		// Drop trailing data
		if err = rw.Truncate(woffset); err != nil {
			err = errors.Wrapf(err, "failed to truncate zipfile '%s'", path.Base(zipfile))
			return
		}
	}

	return
}
