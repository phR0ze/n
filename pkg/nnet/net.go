package nnet

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/phR0ze/n/pkg/sys"
)

// DownloadFile from the given URL to the given destination
// returning the full path to the resulting downloaded file
func DownloadFile(url, dst string, perms ...uint32) (result string, err error) {
	if result, err = sys.Abs(dst); err != nil {
		return
	}

	perm := uint32(0644)
	if len(perms) > 0 {
		perm = perms[0]
	}

	// Create the destination path if it doesn't exist
	sys.MkdirP(path.Dir(result))

	// Open destination truncating if it exists
	var f *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if f, err = os.OpenFile(result, flags, os.FileMode(perm)); err != nil {

	}
	defer f.Close()

	// Download the file to memory via GET
	var res *http.Response
	if res, err = http.Get(url); err != nil {
		return
	}
	defer res.Body.Close()

	// Write streamed http bits to the file
	_, err = io.Copy(f, res.Body)

	return
}
