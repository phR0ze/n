// Package buf provides various specialized buffers
package buf

// Position provides line and column tracking for document
type Position struct {
	Line   int // current line as denoted by newline
	Col    int // column reset for every newline
	Offset int // Offset provides offset from begining of document
}
