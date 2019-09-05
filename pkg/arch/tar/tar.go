// Package tar provides helper functions for tar archives
package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/pkg/errors"
)

// Create a new tar.gz file at tarfile from the given srcPath directory.
// Handles file globbing in the source path.
func Create(tarfile, glob string) (err error) {
	if tarfile, err = sys.Abs(tarfile); err != nil {
		return
	}
	if glob, err = sys.Abs(glob); err != nil {
		return
	}

	// Create the new file for writing to
	var fw *os.File
	if fw, err = os.Create(tarfile); err != nil {
		err = errors.Wrapf(err, "failed to create tarfile %s", tarfile)
		return
	}
	defer func() {
		if e := fw.Close(); e != nil {
			if err == nil {
				err = e
			}
			err = errors.Wrap(err, "failed to close file writer")
		}
	}()

	// Open gzip writer
	gw := gzip.NewWriter(fw)
	defer func() {
		if e := gw.Close(); e != nil {
			if err == nil {
				err = e
			}
			err = errors.Wrap(err, "failed to close gzip writer")
		}
	}()

	// Open tarball writer
	tw := tar.NewWriter(gw)
	defer func() {
		if e := tw.Close(); e != nil {
			if err == nil {
				err = e
			}
			err = errors.Wrap(err, "failed to close tarball writer")
		}
	}()

	// Handle globbing
	var sources []string
	if sources, err = filepath.Glob(glob); err != nil {
		err = errors.Wrapf(err, "failed to get glob for %s", glob)
		return
	}

	// Fail no sources were found
	if len(sources) == 0 {
		err = errors.Errorf("failed to get any sources for %s", glob)
		return
	}

	// Get the source file infos
	infos := []*sys.FileInfo{}
	for _, source := range sources {
		var info *sys.FileInfo
		if info, err = sys.Lstat(source); err != nil {
			return
		}
		infos = append(infos, info)
	}

	// Add all files recursively
	if err = addFiles(tw, infos, ""); err != nil {
		return
	}

	return
}

// AddFiles to the given tar writer recursively where infos are the paths to
// recurse on and base is the path the tar files should be based on in the tar
func addFiles(tw *tar.Writer, infos []*sys.FileInfo, base string) (err error) {
	for _, info := range infos {

		// Recurse on directory
		if info.IsDir() {
			var newInfos []*sys.FileInfo
			if newInfos, err = sys.ReadDir(info.Path); err != nil {
				err = errors.Wrapf(err, "failed to read directory %s to add files from", info.Path)
				return
			}
			newBase := path.Join(base, info.Name())
			if err = addFiles(tw, newInfos, newBase); err != nil {
				return
			}
		} else {

			// Open the file for reading
			var fr *os.File
			if fr, err = os.Open(info.Path); err != nil {
				err = errors.Wrapf(err, "failed to open target file %s for tarball", info.Path)
				return
			}

			// Add the files to the tar
			var header *tar.Header
			if header, err = tar.FileInfoHeader(info.Val, ""); err != nil {
				err = errors.Wrapf(err, "failed to create target file header %s for tarball", info.Path)
				fr.Close()
				return
			}

			// Ensure target is a relative path
			header.Name = path.Join(base, info.Name())

			// Write header to tarball
			if err = tw.WriteHeader(header); err != nil {
				err = errors.Wrapf(err, "failed to write target file header %s for tarball", info.Path)
				fr.Close()
				return
			}

			// Stream the data from the reader to the writer
			if _, err = io.Copy(tw, fr); err != nil {
				err = errors.Wrapf(err, "failed to copy data from reader to writer for tar target %s", info.Path)
				fr.Close()
				return
			}

			// Close reader on success
			fr.Close()
		}
	}

	return
}

// ExtractAll files into given destination directory
func ExtractAll(tarfile, dest string) (err error) {
	if tarfile, err = sys.Abs(tarfile); err != nil {
		return
	}
	if dest, err = sys.MkdirP(dest); err != nil {
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
		err = errors.Wrapf(err, "failed to open gzip reader from %s", tarfile)
		return
	}
	defer gr.Close()

	// Extract all files from tarball
	dirCache := map[string]bool{}
	tr := tar.NewReader(gr)
	for {
		var info *tar.Header
		if info, err = tr.Next(); err == io.EOF {
			err = nil
			break
		} else if err != nil {
			err = errors.Wrapf(err, "failed to extract files from tarfile %s", tarfile)
			return
		}
		filePath := path.Join(dest, info.Name)

		// Create any directories with default mode
		dirPath := path.Dir(filePath)
		if info.Typeflag == tar.TypeDir {
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
		if info.Typeflag == tar.TypeReg {

			// Create file and write content to it
			var fw *os.File
			if fw, err = os.Create(filePath); err != nil {
				err = errors.Wrapf(err, "failed to create file %s from tarfile", filePath)
				return
			}
			if _, err = io.Copy(fw, tr); err != nil {
				err = errors.Wrap(err, "failed to copy data from tar to disk")
				if e := fw.Close(); e != nil {
					err = errors.Wrap(err, "failed to close file")
				}
				return
			}
			if err = fw.Close(); err != nil {
				err = errors.Wrap(err, "failed to close file")
				return
			}

			// Set file mode to the original value
			if err = os.Chmod(filePath, os.FileMode(info.Mode)); err != nil {
				err = errors.Wrapf(err, "failed to set file mode for %s", filePath)
				return
			}
		}

		// Set file access times to the original values
		if err = os.Chtimes(filePath, info.AccessTime, info.ModTime); err != nil {
			err = errors.Wrapf(err, "failed to set file access times for %s", filePath)
			return
		}
	}
	return
}
