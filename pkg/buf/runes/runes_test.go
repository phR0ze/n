package runes

import (
	"strings"
	"testing"

	"github.com/phR0ze/n/pkg/buf"
	"github.com/phR0ze/n/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestReadAll(t *testing.T) {
	scanner := NewScanner(strings.NewReader("foobar"))
	err := scanner.readAll()
	assert.Nil(t, err)
	assert.Equal(t, []rune("foobar"), scanner.runes)
}

func TestRead(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// read it
	r, err := scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)

	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)

	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// read off the edge, EOF position is one past end of content
	r, err = scanner.Read()
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
}

func TestReadline(t *testing.T) {
	scanner := NewScanner(strings.NewReader("foobar\nblah"))
	line, err := scanner.Readline()
	assert.Nil(t, err)
	assert.Equal(t, "foobar\n", string(line.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, line.Pos)
	assert.Equal(t, buf.Position{1, 0, 7}, scanner.Pos)

	line, err = scanner.Readline()
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, "blah", string(line.Val))
	assert.Equal(t, buf.Position{1, 0, 7}, line.Pos)
	assert.Equal(t, buf.Position{1, 4, 11}, scanner.Pos)

	// read off the edge
	line, err = scanner.Readline()
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, "", string(line.Val))
	assert.Equal(t, buf.Position{1, 4, 11}, line.Pos)
	assert.Equal(t, buf.Position{1, 4, 11}, scanner.Pos)
}

func TestUnread(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// unread before we read
	r, err := scanner.Unread()
	assert.Equal(t, errs.BOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)

	// read
	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))

	// unread
	r, err = scanner.Unread()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)

	r, err = scanner.Unread()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)

	r, err = scanner.Unread()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// unread off begining edge
	r, err = scanner.Unread()
	assert.Equal(t, errs.BOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)
}

func TestUnreadline(t *testing.T) {
	scanner := NewScanner(strings.NewReader("foobar\nblah"))
	line, err := scanner.Readline()
	line, err = scanner.Readline()

	// unreadline
	line, err = scanner.Unreadline()
	assert.Nil(t, err)
	assert.Equal(t, "blah", string(line.Val))
	assert.Equal(t, buf.Position{1, 0, 7}, line.Pos)
	assert.Equal(t, buf.Position{1, 0, 7}, scanner.Pos)

	line, err = scanner.Unreadline()
	assert.Equal(t, errs.BOF, err)
	assert.Equal(t, "foobar\n", string(line.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, line.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// unread off front
	line, err = scanner.Unreadline()
	assert.Equal(t, errs.BOF, err)
	assert.Equal(t, "", string(line.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, line.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)
}

func TestReadUnreadline(t *testing.T) {
	scanner := NewScanner(strings.NewReader("foobar\nblah"))
	line, err := scanner.Readline()
	line, err = scanner.Readline()

	// now unread it
	line, err = scanner.Unreadline()
	assert.Nil(t, err)
	assert.Equal(t, "blah", string(line.Val))
	assert.Equal(t, buf.Position{1, 0, 7}, line.Pos)
	assert.Equal(t, buf.Position{1, 0, 7}, scanner.Pos)

	// unread then re-read
	line, err = scanner.Unreadline()
	assert.Equal(t, errs.BOF, err)
	assert.Equal(t, "foobar\n", string(line.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, line.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	line, err = scanner.Readline()
	assert.Nil(t, err)
	assert.Equal(t, "foobar\n", string(line.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, line.Pos)
	assert.Equal(t, buf.Position{1, 0, 7}, scanner.Pos)

	line, err = scanner.Unreadline()
	assert.Equal(t, errs.BOF, err)
	assert.Equal(t, "foobar\n", string(line.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, line.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)
}

func TestPeek(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// peek with offset 1
	r, err := scanner.Peek(1)
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek with offset 0
	r, err = scanner.Peek(0)
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek with offset 2
	r, err = scanner.Peek(2)
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek off the edge
	r, err = scanner.Peek(3)
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek with offset 2
	r, err = scanner.Read()
	r, err = scanner.Read()
	r, err = scanner.Read()
	r, err = scanner.Peek(2)
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
}

func TestPeekPrev(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))
	r, err := scanner.Read()
	r, err = scanner.Read()
	r, err = scanner.Read()

	// peek prev with offset 1
	r, err = scanner.PeekPrev(1)
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// peek prev with offset 0
	r, err = scanner.PeekPrev(0)
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// peek prev with offset 2
	r, err = scanner.PeekPrev(2)
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// // peek off the edge
	// r, err = scanner.Peek(3)
	// assert.Equal(t, errs.EOF, err)
	// assert.Equal(t, int32(0), r.Val)
	// assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	// assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)
}

func TestPeekAndPrevPeek(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// peek, read then peekprev
	r, err := scanner.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)
	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)
	r, err = scanner.PeekPrev()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)

	// peek, read then peekprev
	r, err = scanner.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)
	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)
	r, err = scanner.PeekPrev()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)

	// peek, read then peekprev
	r, err = scanner.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)
	r, err = scanner.Read()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
	r, err = scanner.PeekPrev()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// peek then read off the edge, EOF position is one past end of content
	r, err = scanner.Peek()
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
	r, err = scanner.Read()
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
	r, err = scanner.PeekPrev()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
}
