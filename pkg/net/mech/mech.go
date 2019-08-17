// Package mech provides some simple automation for working with web sites
package mech

import (
	"io"
	"net/http"
	"path"

	"github.com/phR0ze/n/pkg/net/agent"
	"github.com/phR0ze/n/pkg/sys"
	"github.com/pkg/errors"
)

// Mech provides some simple automation for working with web sites
type Mech struct {
	agent  string
	client *http.Client
}

// New creates a new Mech instance
func New() (mech *Mech) {
	mech = &Mech{
		agent:  agent.IPhoneIOS12,
		client: &http.Client{},
	}
	return
}

// Download wrapper for instance method
func Download(url, dst string, perms ...uint32) (filepath string, err error) {
	return New().Download(url, dst, perms...)
}

// Download from the given URL to the given destination
// returning the full path to the resulting downloaded file
func (mech *Mech) Download(url, dst string, perms ...uint32) (filepath string, err error) {
	if filepath, err = sys.Abs(dst); err != nil {
		return
	}

	// Create the destination path if it doesn't exist
	sys.MkdirP(path.Dir(filepath))

	// Make the request to get a stream reader for the target
	var reader io.ReadCloser
	if reader, err = mech.Stream(url); err != nil {
		return
	}
	defer reader.Close()

	// Stream down the file and write it to disk
	if err = sys.WriteStream(reader, filepath, perms...); err != nil {
		return
	}

	return
}

// Get wrapper for instance method
func Get(url string) (page *Page, err error) {
	return New().Get(url)
}

// Get the target html docuement as a *Page
// Caller is responsible for Closing the Page
func (mech *Mech) Get(url string) (page *Page, err error) {
	var stream io.ReadCloser
	if stream, err = mech.Stream(url); err != nil {
		return
	}
	page = &Page{
		stream: stream,
	}

	return
}

// GetLinks wrapper for instance method
func GetLinks(url string) (links []string, err error) {
	return New().GetLinks(url)
}

// GetLinks from the target html docuement
func (mech *Mech) GetLinks(url string) (links []string, err error) {
	var stream io.ReadCloser
	if stream, err = mech.Stream(url); err != nil {
		return
	}
	page := &Page{stream: stream}
	if links, err = page.Links(); err != nil {
		return
	}

	return
}

// Stream wrapper for instance method
func Stream(url string) (reader io.ReadCloser, err error) {
	return New().Stream(url)
}

// Stream the given url
// Caller is responsible for closing the reader when finished
func (mech *Mech) Stream(url string) (reader io.ReadCloser, err error) {

	// Create the request with the correct configuration
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		err = errors.Wrap(err, "failed to create http request")
		return
	}
	req.Header.Set("User-Agent", mech.agent)

	// Make the request to get a stream reader for the target
	var res *http.Response
	if res, err = mech.client.Do(req); err != nil {
		err = errors.Wrapf(err, "failed to GET url %s", url)
		return
	}
	reader = res.Body
	return
}
