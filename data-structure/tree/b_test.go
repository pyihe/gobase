package tree

import (
	"fmt"
	"testing"
)

func TestBTree_Insert(t *testing.T) {
	bt := NewBTree(4)
	testData := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	for _, r := range testData {
		bt.Insert(int(r), string(r))
	}

	// fmt.Println(bt.root)
	fmt.Println(bt.Find('Z'))
	fmt.Println(bt.Find('H'))
	fmt.Println(bt.Find('L'))

	fmt.Println(bt.Remove('Z'))
	fmt.Println(bt.root)
}
