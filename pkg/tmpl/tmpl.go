// Package tmpl does simple template substitutions fast
package tmpl

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/phR0ze/n/pkg/errs"
	"github.com/valyala/bytebufferpool"
)

// Engine provides encapsulation and class methods for templating
type Engine struct {
	data     string
	startTag string
	endTag   string

	texts          [][]byte
	tags           []string
	byteBufferPool bytebufferpool.Pool
}

// Load a yaml file from disk and processes any templating returning a string
func Load(filepath string, startTag, endTag string, vars map[string]string) (result string, err error) {
	var data []byte
	if data, err = ioutil.ReadFile(filepath); err == nil {
		var tpl *Engine
		if tpl, err = New(string(data), startTag, endTag); err == nil {
			if result, err = tpl.Process(vars); err == nil {
				return
			}
		} else if err == errs.TmplVarsNotFoundError {
			err = nil
			result = string(data)
		}
	}
	return
}

// New parses the given data using the given startTag and endTag
// into components that can then be easily worked with
func New(data, startTag, endTag string) (*Engine, error) {
	var tpl Engine
	tpl.data = data
	tpl.startTag = startTag
	tpl.endTag = endTag
	tpl.texts = tpl.texts[:0]
	tpl.tags = tpl.tags[:0]

	// Check that we have valid tags
	if startTag == "" || endTag == "" {
		return nil, errs.TmplTagsInvalidError
	}

	s := []byte(data)
	a := []byte(startTag)
	b := []byte(endTag)

	// Done if there are no template variables
	tagsCount := bytes.Count(s, a)
	if tagsCount == 0 {
		return nil, errs.TmplVarsNotFoundError
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
			return nil, errs.NewTmplEndTagNotFoundError(endTag, s)
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
func (tpl *Engine) Process(m map[string]string) (result string, err error) {
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
