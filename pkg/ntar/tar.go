package ntar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
	"strings"

	"github.com/phR0ze/n/pkg/nos"
)

// Create a 'dst' tarball from the given 'src' directory
func Create(dst string, src string) (err error) {
	var root string
	if root, err = nos.Abs(src); err != nil {
		return
	}

	// Create the new file for writing to
	var fw *os.File
	if fw, err = os.Create(dst); err != nil {
		return
	}
	defer fw.Close()

	// Open gzip writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// Open tarball writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Walk the given root paths and add them to the tarball
	var paths []string
	if paths, err = nos.AbsPaths(root); err != nil {
		return
	}
	for _, x := range paths {
		if err = addPath(tw, root, x); err != nil {
			return
		}
	}
	return
}

// AddPath to the given tarball
func addPath(tw *tar.Writer, root, target string) (err error) {

	// Open the file for reading
	var f *os.File
	if f, err = os.Open(target); err != nil {
		return
	}
	defer f.Close()

	// Get the file information
	var info os.FileInfo
	if info, err = f.Stat(); err != nil {
		return
	}

	// Create target header
	var header *tar.Header
	if header, err = tar.FileInfoHeader(info, ""); err != nil {
		return
	}

	// Ensure target is a relative path
	header.Name = strings.TrimPrefix(target, path.Dir(root)+"/")

	// Write header to tarball
	if err = tw.WriteHeader(header); err != nil {
		return
	}

	// Write the target to the tarball
	if !info.IsDir() {
		if _, err = io.Copy(tw, f); err != nil {
			return
		}
	}

	return
}

// ExtractAll files into given destination directory
func ExtractAll(tarball, dest string) (err error) {

	// Create destination directory if it doesn't exist
	if !nos.Exists(dest) {
		nos.MkdirP(dest)
	}

	// Open tarball for use
	var fr *os.File
	if fr, err = os.Open(tarball); err != nil {
		return
	}
	defer fr.Close()

	// Open gzip reader
	var gr *gzip.Reader
	if gr, err = gzip.NewReader(fr); err != nil {
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
			return
		}

		// Create directories
		if header.Typeflag == tar.TypeDir {
			dirPath := path.Join(dest, header.Name)
			if _, exists := dirCache[dirPath]; !exists {
				nos.MkdirP(dirPath, os.FileMode(header.Mode))
				dirCache[dirPath] = true
			}
		}

		// Write out files
		if header.Typeflag == tar.TypeReg {
			filePath := path.Join(dest, header.Name)

			// Create file and write content to it
			var fw *os.File
			if fw, err = os.Create(filePath); err != nil {
				return
			}
			io.Copy(fw, tr)
			fw.Close()

			// Set file attributes
			if err = os.Chmod(filePath, os.FileMode(header.Mode)); err != nil {
				return
			}
		}
	}
	return
}
