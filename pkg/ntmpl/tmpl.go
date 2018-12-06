// Package ntmpl does simple template substitutions fast
package ntmpl

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/phR0ze/n/pkg/nerr"
	"github.com/valyala/bytebufferpool"
)

type tplEngine struct {
	data     string
	startTag string
	endTag   string

	texts          [][]byte
	tags           []string
	byteBufferPool bytebufferpool.Pool
}

// New parses the given data using the given startTag and endTag
// into components that can then be easily worked with
func New(data, startTag, endTag string) (*tplEngine, error) {
	var tpl tplEngine
	tpl.data = data
	tpl.startTag = startTag
	tpl.endTag = endTag
	tpl.texts = tpl.texts[:0]
	tpl.tags = tpl.tags[:0]

	// Check that we have valid tags
	if startTag == "" || endTag == "" {
		return nil, nerr.NewTmplTagsInvalid()
	}

	s := []byte(data)
	a := []byte(startTag)
	b := []byte(endTag)

	// Check that there is at least one tag
	tagsCount := bytes.Count(s, a)
	if tagsCount == 0 {
		return nil, errors.New("no template variables were found")
	}

	if tagsCount+1 > cap(tpl.texts) {
		tpl.texts = make([][]byte, 0, tagsCount+1)
	}
	if tagsCount > cap(tpl.tags) {
		tpl.tags = make([]string, 0, tagsCount)
	}

	// Now parse the data
	for {
		n := bytes.Index(s, a)
		if n < 0 {
			tpl.texts = append(tpl.texts, s)
			break
		}
		tpl.texts = append(tpl.texts, s[:n])

		s = s[n+len(a):]
		n = bytes.Index(s, b)
		if n < 0 {
			msg := fmt.Sprintf("Cannot find end tag=%q in the template=%q starting from %q", endTag, data, s)
			return nil, errors.New(msg)
		}

		// Fix bug in original code to remove wrapping spaces
		tag := strings.Trim(string(s[:n]), " ")
		tpl.tags = append(tpl.tags, tag)
		s = s[n+len(b):]
	}

	return &tpl, nil
}

// Process substitutes template tags (placeholders) with the corresponding
// values from the map m and returns the result.
func (tpl *tplEngine) Process(m map[string]string) (result string, err error) {
	w := tpl.byteBufferPool.Get()
	n := len(tpl.texts) - 1
	if n == -1 {
		if _, err = w.Write([]byte(tpl.data)); err != nil {
			return
		}
	} else {
		for i := 0; i < n; i++ {
			if _, err = w.Write(tpl.texts[i]); err != nil {
				return
			}
			v := m[tpl.tags[i]]
			if _, err = w.Write([]byte(v)); err != nil {
				return
			}
		}
		if _, err = w.Write(tpl.texts[n]); err != nil {
			return
		}
	}

	s := string(w.Bytes())
	w.Reset()
	tpl.byteBufferPool.Put(w)
	return s, err
}
