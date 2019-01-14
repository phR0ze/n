package main

import (
	"github.com/phR0ze/n"
)

const size = 9999999

func main() {
	n.Q(n.Range(0, size)).Find(func(x n.O) bool {
		return x.(int) == size-1
	})
}
