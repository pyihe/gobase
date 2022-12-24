package tree

import (
	"fmt"
	"testing"
)

func TestHuffmanTree_Encode(t *testing.T) {
	source := map[rune]int{
		'A': 27,
		'B': 8,
		'C': 15,
		'D': 15,
		'E': 30,
		'F': 5,
	}
	plain := "BADCADFEED"
	ht := NewHuffmanTree(source)
	// D A F B C E
	fmt.Println("huffman: ", ht)
	code := ht.Encode(plain)
	fmt.Println("code: ", code)
	dPlain := ht.Decode(code)
	fmt.Println("plain: ", dPlain)
}
