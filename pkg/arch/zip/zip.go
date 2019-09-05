// Package zip provides create and extract implementations
package zip

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/phR0ze/n/pkg/enc/bin"
	"github.com/phR0ze/n/pkg/sys"
	"github.com/pkg/errors"
)

var (
	// Zip files typically start with PK i.e. 50 4B
	gZipHeaderSig = []byte{0x50, 0x4B}
)

// Create a new zipfile at zipfile from the given srcPath directory
// Handles file globbing in the source path.
func Create(zipfile, glob string) (err error) {
	if zipfile, err = sys.Abs(zipfile); err != nil {
		return
	}
	if glob, err = sys.Abs(glob); err != nil {
		return
	}

	// Create the new file for writing to
	var fw *os.File
	if fw, err = os.Create(zipfile); err != nil {
		err = errors.Wrapf(err, "failed to create zipfile %s", zipfile)
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

	// Open zip writer
	zw := zip.NewWriter(fw)
	defer func() {
		if e := zw.Close(); e != nil {
			if err == nil {
				err = e
			}
			err = errors.Wrap(err, "failed to close zipfile writer")
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
	if err = addFiles(zw, infos, ""); err != nil {
		return
	}

	return
}

// AddFiles to the given zip writer recursively where infos are the paths to
// recurse on and base is the path the zip files should be based on in the zip
func addFiles(zw *zip.Writer, infos []*sys.FileInfo, base string) (err error) {
	for _, info := range infos {

		// Recurse on directory
		if info.IsDir() {
			var newInfos []*sys.FileInfo
			if newInfos, err = sys.ReadDir(info.Path); err != nil {
				err = errors.Wrapf(err, "failed to read directory %s to add files from", info.Path)
				return
			}
			newBase := path.Join(base, info.Name())
			if err = addFiles(zw, newInfos, newBase); err != nil {
				return
			}
		} else {

			// Open the target file for reading
			var fr *os.File
			if fr, err = os.Open(info.Path); err != nil {
				err = errors.Wrapf(err, "failed to open target file %s for zip", info.Path)
				return
			}

			// Add the files to the zip
			var fw io.Writer
			zipPath := path.Join(base, info.Name())
			if fw, err = zw.Create(zipPath); err != nil {
				err = errors.Wrapf(err, "failed to add target file %s to zip", info.Path)
				fr.Close()
				return
			}

			// Stream the data from the reader to the writer
			if _, err = io.Copy(fw, fr); err != nil {
				err = errors.Wrapf(err, "failed to copy data from reader to writer for zip target %s", info.Path)
				fr.Close()
				return
			}

			// Close the file reader/writers here as were in a loop
			fr.Close()
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

	// TrimPrefix from the zipfile if required
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
				err = errors.Wrapf(err, "failed to create file %s from zipfile", filePath)
				return
			}
			var fr io.ReadCloser
			if fr, err = file.Open(); err != nil {
				err = errors.Wrapf(err, "failed to open zipfile target %s for reading", info.Name())
				if e := fw.Close(); e != nil {
					err = errors.Wrap(err, "failed to close zipfile writer")
				}
				return
			}
			if _, err = io.Copy(fw, fr); err != nil {
				err = errors.Wrap(err, "failed to copy data from zip to disk")
				if e := fw.Close(); e != nil {
					err = errors.Wrap(err, "failed to close zipfile writer")
				}
				fr.Close()
				return
			}
			if err = fw.Close(); err != nil {
				err = errors.Wrap(err, "failed to close zipfile writer")
				fr.Close()
				return
			}
			fr.Close()

			// Set file mode to the original value
			if err = os.Chmod(filePath, info.Mode()); err != nil {
				err = errors.Wrapf(err, "failed to set file mode for %s", filePath)
				return
			}
		}

		// Set file access times to the original values
		if err = os.Chtimes(filePath, info.ModTime(), info.ModTime()); err != nil {
			err = errors.Wrapf(err, "failed to set file access times for %s", filePath)
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
		err = errors.Wrapf(err, "failed to open zipfile '%s' to detect zip data", path.Base(zipfile))
		return
	}
	defer func() {
		if e := rw.Close(); e != nil {
			if err == nil {
				err = e
			}
			err = errors.Wrap(err, "failed to close file writer")
		}
	}()

	// Detect begining of zipfile identified by PK.. i.e. 50 4B
	chunk := make([]byte, 1024)
	loc, roffset := int64(-1), int64(-1)
	for {
		l, e := rw.Read(chunk)
		if e != nil {
			if e != io.EOF {
				err = errors.Wrapf(e, "failed to read from zipfile '%s' for zip identification", path.Base(zipfile))
				return
			} else if roffset == -1 {
				err = errors.Errorf("unable to identify '%s' as a valid zipfile", path.Base(zipfile))
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
				err = errors.Wrapf(e, "failed to read from zipfile '%s' to shift data", path.Base(zipfile))
				return
			} else if l > 0 {
				data := chunk[0:l]
				if _, e := rw.WriteAt(data, woffset); e != nil {
					err = errors.Wrapf(e, "failed to write shifted data to zipfile '%s'", path.Base(zipfile))
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
