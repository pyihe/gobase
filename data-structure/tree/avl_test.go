package tree

import (
	"fmt"
	"testing"
)

func TestAVL_Insert(t *testing.T) {
	var (
		avl      Tree = NewAVL()
		testData      = []TestInt{1, 2, 3, 4, 5, 6, 7, 7, 4, 6, 5}
	)

	//4->2->1->3->6->5->7
	for _, data := range testData {
		avl.Insert(data)
	}

	fmt.Println(PreOrderTraverseRecursion(avl.Root()))
}
