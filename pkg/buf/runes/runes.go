// Package runes provides a simpler rune scanner for lexers.
//
// It reads an io.Reader entirely into memory to accomplish this so should only be used
// where the data can easily fit into memory.
package runes

import (
	"io"
	"io/ioutil"

	"github.com/phR0ze/n/pkg/buf"
	"github.com/phR0ze/n/pkg/errs"
	"github.com/pkg/errors"
)

// Scanner provides methods for working with documents as runes
type Scanner struct {
	src   io.Reader    // original source reader
	runes []rune       // runes to work with
	Pos   buf.Position // line, column and offset positioning in document
}

// Rune wraps a rune with with positioning
type Rune struct {
	Val rune         // actual rune
	Pos buf.Position // position of the rune as relates to the document
}

// Runes wraps a slice of rune with positioning
type Runes struct {
	Val []rune       // actual runes
	Pos buf.Position // position of the first rune as relates to the document
}

// NewScanner for scanning documents as runes
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{src: reader}
}

// read the source reader into a rune slice
func (s *Scanner) readAll() (err error) {
	if len(s.runes) != 0 {
		return
	}

	// Read all data into memory and conver to rune slice
	var data []byte
	if data, err = ioutil.ReadAll(s.src); err != nil {
		err = errors.Wrap(err, "failed to read all data from source reader")
		return
	}
	s.runes = []rune(string(data))
	return
}

// Size simply exposes the len of the rune slice
func (s *Scanner) Size() int {
	return len(s.runes)
}

// Read the rune at the current location and adjust positioning
func (s *Scanner) Read() (r Rune, err error) {
	r = Rune{Pos: s.Pos}
	if err = s.readAll(); err != nil {
		return
	}
	if len(s.runes) == 0 {
		return
	}
	if len(s.runes) <= s.Pos.Offset {
		err = errs.EOF
		return
	}

	// Read the current rune
	r.Val = s.runes[s.Pos.Offset]
	s.Pos = s.incPos(s.Pos)
	return
}

// Readline from the rune slice location up to and including the next newline adjusting positioning
func (s *Scanner) Readline() (line Runes, err error) {
	line = Runes{Val: []rune{}, Pos: s.Pos}
	if err = s.readAll(); err != nil {
		return
	}
	if len(s.runes) == 0 {
		return
	}
	if len(s.runes) <= s.Pos.Offset {
		err = errs.EOF
		return
	}

	for {
		// Read the current rune
		r := s.runes[s.Pos.Offset]
		line.Val = append(line.Val, r)
		s.Pos = s.incPos(s.Pos)

		// Break on newline or EOF
		if r == '\n' {
			break
		}
		if len(s.runes) <= s.Pos.Offset {
			err = errs.EOF
			break
		}
	}
	return
}

// Unread the rune at the current location and adjust positioning
func (s *Scanner) Unread() (r Rune, err error) {
	if err = s.readAll(); err != nil {
		return
	}
	if len(s.runes) == 0 {
		return
	}
	if s.Pos.Offset == 0 {
		err = errs.BOF
		return
	}

	// Read the previous rune
	s.Pos = s.decPos(s.Pos)
	r = Rune{Val: s.runes[s.Pos.Offset], Pos: s.Pos}
	return
}

// Unreadline from the current location back to the previous newline adjusting positioning
func (s *Scanner) Unreadline() (line Runes, err error) {
	line = Runes{Val: []rune{}, Pos: s.Pos}
	if err = s.readAll(); err != nil {
		return
	}
	if len(s.runes) == 0 {
		return
	}
	if s.Pos.Offset == 0 {
		err = errs.BOF
		return
	}

	first := true
	for {
		// Read previous line
		r := s.runes[s.Pos.Offset-1]
		if !first && r == '\n' {
			break
		}
		if first {
			first = false
		}

		line.Val = append([]rune{r}, line.Val...)
		s.Pos = s.decPos(s.Pos)
		line.Pos = s.Pos

		if s.Pos.Offset == 0 {
			err = errs.BOF
			return
		}
	}
	return
}

// Peek the rune at the current location don't adjust positioning
func (s *Scanner) Peek() (r Rune, err error) {
	r = Rune{Pos: s.Pos}
	if err = s.readAll(); err != nil {
		return
	}
	if len(s.runes) == 0 {
		return
	}
	if len(s.runes) <= s.Pos.Offset {
		err = errs.EOF
		return
	}

	// Read the current rune
	r.Val = s.runes[s.Pos.Offset]
	return
}

// PeekPrev the rune at the previous location don't adjust positioning
func (s *Scanner) PeekPrev() (r Rune, err error) {
	if err = s.readAll(); err != nil {
		return
	}
	if len(s.runes) == 0 {
		return
	}
	if s.Pos.Offset == 0 {
		err = errs.BOF
		return
	}

	// Read the previous rune
	pos := s.decPos(s.Pos)
	r = Rune{Val: s.runes[pos.Offset], Pos: pos}
	return
}

// increment the position based on the given rune
func (s *Scanner) incPos(pos buf.Position) buf.Position {
	val := pos
	r := s.runes[pos.Offset]
	if r == '\n' {
		val.Line++
		val.Col = 0
	} else {
		val.Col++
	}
	val.Offset++
	return val
}

// decrement the position based on the given rune; decrementing column needs to take into account
// the previous line to get the correct column position.
func (s *Scanner) decPos(pos buf.Position) buf.Position {
	val := pos
	val.Col--
	val.Offset--
	if val.Col < 0 {
		val.Col = 0
	}
	if val.Offset <= 0 {
		return buf.Position{}
	}

	r := s.runes[pos.Offset-1]
	if r == '\n' {
		val.Line--
		if val.Line < 0 {
			val.Line = 0
		}

		// Column needs to take into account the previous line
		i := pos.Offset - 1
		runes := []rune{'\n'}
		for {
			i--
			if i < 0 {
				break
			}
			r = s.runes[i]
			if r == '\n' {
				break
			}
			runes = append([]rune{r}, runes...)
		}
		val.Col = (pos.Offset + 1) - len(runes)
	}
	return val
}
