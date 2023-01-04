package tree

import (
	"fmt"
	"testing"
)

func TestRedBlackTree_Insert(t *testing.T) {
	var (
		rbt      Tree = NewRedBlackTree()
		testData      = []TestInt{1, 2, 3, 4, 5, 6, 7}
	)

	for _, data := range testData {
		rbt.Insert(data)
	}

	fmt.Println(PreOrderTraverseRecursion(rbt.Root()))

	fmt.Println(rbt.Remove(TestInt(1)))
	fmt.Println(PreOrderTraverseRecursion(rbt.Root()))
}
