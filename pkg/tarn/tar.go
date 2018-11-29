package tarn

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
)

// ExtractAll files into given destination directory
func ExtractAll(tarball, dest string) error {

	// Create destination directory if it doesn't exist
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		os.MkdirAll(dest, 0744)
	}

	// Open tarball for use
	f, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer f.Close()

	// Open gzip reader
	gzReader, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	// Extract all files from tarball
	dirCache := map[string]bool{}
	tarReader := tar.NewReader(gzReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if header.Typeflag == tar.TypeReg {
			filePath := path.Join(dest, header.Name)
			dirPath := path.Join(dest, path.Dir(header.Name))

			// Create directories as needed
			if _, exists := dirCache[dirPath]; !exists {
				os.MkdirAll(dirPath, os.FileMode(header.Mode))
				dirCache[dirPath] = true
			}

			// Write out file
			ef, err := os.Create(filePath)
			if err != nil {
				return err
			}
			io.Copy(ef, tarReader)
			ef.Close()
		}
	}
	return nil
}
