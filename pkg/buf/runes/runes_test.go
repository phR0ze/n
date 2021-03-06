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
	assert.Equal(t, "f", string(scanner.Read()))
	assert.Equal(t, "\n", string(scanner.Read()))
	assert.Equal(t, "o", string(scanner.Read()))
	assert.Equal(t, int32(0), scanner.Read())
	assert.Equal(t, int32(0), scanner.Read())
}

func TestReadE(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// read it
	r, err := scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)

	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)

	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// read off the edge, EOF position is one past end of content
	r, err = scanner.ReadE()
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

func TestSize(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))
	assert.Equal(t, 0, scanner.Size())
	scanner.readAll()
	assert.Equal(t, 3, scanner.Size())
}

func TestUnread(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))
	assert.Equal(t, int32(0), scanner.Unread())
	assert.Equal(t, "f", string(scanner.Read()))
	assert.Equal(t, "\n", string(scanner.Read()))
	assert.Equal(t, "o", string(scanner.Read()))
	assert.Equal(t, "o", string(scanner.Unread()))
	assert.Equal(t, "\n", string(scanner.Unread()))
	assert.Equal(t, "f", string(scanner.Unread()))
	assert.Equal(t, int32(0), scanner.Unread())
}

func TestUnreadE(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// unread before we read
	r, err := scanner.UnreadE()
	assert.Equal(t, errs.BOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)

	// read
	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))

	// unread
	r, err = scanner.UnreadE()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)

	r, err = scanner.UnreadE()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)

	r, err = scanner.UnreadE()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// unread off begining edge
	r, err = scanner.UnreadE()
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

func TestPeekE(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// peek with offset 1
	r, err := scanner.PeekE(1)
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek with offset 0
	r, err = scanner.PeekE(0)
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek with offset 2
	r, err = scanner.PeekE(2)
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek off the edge
	r, err = scanner.PeekE(3)
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)

	// peek with offset 2
	r, err = scanner.ReadE()
	r, err = scanner.ReadE()
	r, err = scanner.ReadE()
	r, err = scanner.PeekE(2)
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
}

func TestPeekPrevE(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))
	r, err := scanner.ReadE()
	r, err = scanner.ReadE()
	r, err = scanner.ReadE()

	// peek prev with offset 1
	r, err = scanner.PeekPrevE(1)
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// peek prev with offset 0
	r, err = scanner.PeekPrevE(0)
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// peek prev with offset 2
	r, err = scanner.PeekPrevE(2)
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
}

func TestPeekAndPrev(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))
	assert.Equal(t, int32(0), scanner.Unread())
	assert.Equal(t, "f", string(scanner.Peek()))
	assert.Equal(t, "f", string(scanner.Read()))
	assert.Equal(t, "f", string(scanner.PeekPrev()))

	assert.Equal(t, "\n", string(scanner.Peek()))
	assert.Equal(t, "o", string(scanner.Peek(1)))
	assert.Equal(t, "\n", string(scanner.Read()))
	assert.Equal(t, "\n", string(scanner.PeekPrev()))

	assert.Equal(t, "o", string(scanner.Peek()))
	assert.Equal(t, "o", string(scanner.Read()))
	assert.Equal(t, "o", string(scanner.PeekPrev()))
	assert.Equal(t, "o", string(scanner.Unread()))
	assert.Equal(t, "o", string(scanner.Peek()))

	assert.Equal(t, "f", string(scanner.PeekPrev(2)))
}

func TestPeekAndPrevPeekE(t *testing.T) {
	scanner := NewScanner(strings.NewReader("f\no"))

	// peek, read then peekprev
	r, err := scanner.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 0, 0}, scanner.Pos)
	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)
	r, err = scanner.PeekPrevE()
	assert.Nil(t, err)
	assert.Equal(t, "f", string(r.Val))
	assert.Equal(t, buf.Position{0, 0, 0}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)

	// peek, read then peekprev
	r, err = scanner.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{0, 1, 1}, scanner.Pos)
	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)
	r, err = scanner.PeekPrevE()
	assert.Nil(t, err)
	assert.Equal(t, "\n", string(r.Val))
	assert.Equal(t, buf.Position{0, 1, 1}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)

	// peek, read then peekprev
	r, err = scanner.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 0, 2}, scanner.Pos)
	r, err = scanner.ReadE()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
	r, err = scanner.PeekPrevE()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)

	// peek then read off the edge, EOF position is one past end of content
	r, err = scanner.PeekE()
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
	r, err = scanner.ReadE()
	assert.Equal(t, errs.EOF, err)
	assert.Equal(t, int32(0), r.Val)
	assert.Equal(t, buf.Position{1, 1, 3}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
	r, err = scanner.PeekPrevE()
	assert.Nil(t, err)
	assert.Equal(t, "o", string(r.Val))
	assert.Equal(t, buf.Position{1, 0, 2}, r.Pos)
	assert.Equal(t, buf.Position{1, 1, 3}, scanner.Pos)
}
