package mech

import (
	"io"

	"github.com/phR0ze/goquery"
	"github.com/pkg/errors"
)

// Page encapsulates an HTML document and provides helper methods
type Page struct {
	doc    *goquery.Document // goquery document
	stream io.ReadCloser     // initial stream
}

// Close implementes the Closer interface
func (page *Page) Close() error {
	return page.stream.Close()
}

// Find exposes the underlying goquery Find function
func (page *Page) Find(selector string) (goquery *goquery.Selection, err error) {
	if err = page.Parse(); err != nil {
		return
	}
	goquery = page.doc.Find(selector)
	return
}

// Links returns all links for the given page
func (page *Page) Links() (links []string, err error) {
	if err = page.Parse(); err != nil {
		return
	}

	page.doc.Find("a").Each(func(i int, elem *goquery.Selection) {
		if href, ok := elem.Attr("href"); ok {
			links = append(links, href)
		}
	})
	return
}

// Parse the page as a goquery document if not already done. This method
// is idempotent and will only parse the document if not already done so.
func (page *Page) Parse() (err error) {
	if page.doc == nil {
		if page.doc, err = goquery.NewDocumentFromReader(page.stream); err != nil {
			err = errors.Wrap(err, "failed to create goquery from Page stream")
			return
		}
	}
	return
}
