// Package tar provides helper functions for tar archives
package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/pkg/errors"
)

// Create a new tar.gz file at tarfile from the given srcPath directory
func Create(tarfile, srcPath string) (err error) {
	var srcAbs string
	if srcAbs, err = sys.Abs(srcPath); err != nil {
		return
	}

	// Create the new file for writing to
	var fw *os.File
	if fw, err = os.Create(tarfile); err != nil {
		err = errors.Wrapf(err, "failed to create tarfile %s", tarfile)
		return
	}
	defer fw.Close()

	// Open gzip writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// Open tarball writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Add all files recursively
	if err = addFiles(tw, srcAbs, ""); err != nil {
		return
	}
	return
}

// AddFiles to the given zip writer recursively where root is the directory
// to recurse on and base is the path the zip files should be based on in the zip
func addFiles(tw *tar.Writer, root, base string) (err error) {
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
				err = errors.Wrapf(err, "failed to open target file %s for tar", target)
				return
			}
			defer fr.Close()

			// Add the files to the tar
			var header *tar.Header
			if header, err = tar.FileInfoHeader(info, ""); err != nil {
				err = errors.Wrapf(err, "failed to create target file header %s for tar", target)
				return
			}

			// Ensure target is a relative path
			header.Name = path.Join(base, info.Name())

			// Write header to tarball
			if err = tw.WriteHeader(header); err != nil {
				err = errors.Wrapf(err, "failed to write target file header %s for tar", target)
				return
			}

			// Stream the data from the reader to the writer
			if _, err = io.Copy(tw, fr); err != nil {
				err = errors.Wrapf(err, "failed to copy data from reader to writer for zip target %s", target)
				return
			}
		} else {
			newRoot := path.Join(root, info.Name())
			newBase := path.Join(base, info.Name())
			addFiles(tw, newRoot, newBase)
		}
	}

	return
}

// ExtractAll files into given destination directory
func ExtractAll(tarfile, dest string) (err error) {
	if _, err = sys.MkdirP(dest); err != nil {
		return
	}

	// Open tarball for use
	var fr *os.File
	if fr, err = os.Open(tarfile); err != nil {
		err = errors.Wrapf(err, "failed to open tarfile %s for reading", tarfile)
		return
	}
	defer fr.Close()

	// Open gzip reader
	var gr *gzip.Reader
	if gr, err = gzip.NewReader(fr); err != nil {
		err = errors.Wrapf(err, "failed to open gzip reader %s from", tarfile)
		return
	}
	defer gr.Close()

	// Extract all files from tarball
	dirCache := map[string]bool{}
	tr := tar.NewReader(gr)
	for {
		var header *tar.Header
		if header, err = tr.Next(); err == io.EOF {
			err = nil
			break
		} else if err != nil {
			err = errors.Wrapf(err, "failed to extract files from tarfile %s", tarfile)
			return
		}

		// Create directories
		if header.Typeflag == tar.TypeDir {
			dirPath := path.Join(dest, header.Name)
			if _, exist := dirCache[dirPath]; !exist {
				if _, err = sys.MkdirP(dirPath, uint32(header.Mode)); err != nil {
					err = errors.Wrapf(err, "failed to create directory %s", dirPath)
					return
				}
				dirCache[dirPath] = true
			}
			continue
		}

		// Write out files
		if header.Typeflag == tar.TypeReg {
			filePath := path.Join(dest, header.Name)

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
				err = errors.Wrapf(err, "failed to create target file %s from tarfile", filePath)
				return
			}
			io.Copy(fw, tr)
			fw.Close()

			// Set file mode to the original value
			if err = os.Chmod(filePath, os.FileMode(header.Mode)); err != nil {
				err = errors.Wrapf(err, "failed to set file mode for %s", filePath)
				return
			}
		}
	}
	return
}
