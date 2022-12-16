package kmp

import (
	"fmt"
	"testing"
)

var testdata = map[string]string{
	"hello golang!":   "go ",
	"i'am gopher":     " go",
	"aaaabbbdddd":     "aaaac",
	"cccddddzzz":      "cccdddc",
	"aaaaabbbbbsssss": "abbbbbs",
	"0000111111":      "000001",
}

func TestSimpleMatch(t *testing.T) {
	for s, t := range testdata {
		fmt.Printf("[%s], [%s], %d\n", s, t, Search(s, t, 0))
	}
}
