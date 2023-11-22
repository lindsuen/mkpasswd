package main

import (
	"testing"

	"github.com/lindsuen0/mkpasswd/src"
)

const (
	N uint64 = 999999
	l uint64 = 16
	n uint64 = 4
	c uint64 = 4
)

func BenchmarkMain(b *testing.B) {
	var i uint64
	for i = 0; i < N; i++ {
		src.GenerateRandomString(l, n, c)
	}
}
