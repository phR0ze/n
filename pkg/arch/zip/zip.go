// Package zip provides create and extract implementations
package zip

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/pkg/errors"
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

		// Create directories
		if info.IsDir() {
			if _, exist := dirCache[filePath]; !exist {
				if _, err = sys.MkdirP(filePath, uint32(info.Mode())); err != nil {
					err = errors.Wrapf(err, "failed to create directory %s", filePath)
					return
				}
				dirCache[filePath] = true
			}
			continue
		}

		// Create any directories with default permissions that don't exist
		dirPath := path.Dir(filePath)
		if _, exist := dirCache[dirPath]; !exist {
			if _, err = sys.MkdirP(path.Dir(filePath)); err != nil {
				err = errors.Wrapf(err, "failed to create directory %s", path.Dir(filePath))
				return
			}
			dirCache[dirPath] = true
		}

		// Create file and write content to it
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
		if err = os.Chmod(filePath, os.FileMode(info.Mode())); err != nil {
			err = errors.Wrapf(err, "failed to set file mode for %s", filePath)
			return
		}
	}
	return
}
