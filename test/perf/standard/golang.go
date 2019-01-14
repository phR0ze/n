package main

import (
	"github.com/phR0ze/n"
)

const size = 9999999

func main() {
	for _, x := range n.Range(0, size) {
		if x == size-1 {
			return
		}
	}
}
