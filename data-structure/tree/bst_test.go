package tree

import (
	"fmt"
	"testing"
)

type TestInt int

func (ti TestInt) Value() interface{} {
	return ti
}

func (ti TestInt) Compare(ce Element) int {
	cv := ce.Value().(TestInt)
	if ti > cv {
		return 1
	}
	if ti < cv {
		return -1
	}
	return 0
}

func TestBST_Remove(t *testing.T) {
	var (
		bst      Tree = NewBST()
		testData      = []TestInt{2, 1, 8, 5, 3, 4, 9, 6, 10, 7}
	)

	for _, data := range testData {
		bst.Insert(data)
	}
	fmt.Println("bst init: ", PreOrderTraverseRecursion(bst.Root()))
	fmt.Println(bst.Root(), bst.Depth())
	bst.Remove(TestInt(2))
	fmt.Println(bst.Root(), bst.Depth())
	fmt.Println("after remove 2:", PreOrderTraverseRecursion(bst.Root()))
	bst.Remove(TestInt(10))
	fmt.Println(bst.Root(), bst.Depth())
	fmt.Println("after remove 10:", PreOrderTraverseRecursion(bst.Root()))
	fmt.Println(bst.Find(TestInt(1)), bst.Find(TestInt(9)))
	bst.Update(TestInt(5), TestInt(11))
	fmt.Println("after update:", PreOrderTraverseRecursion(bst.Root()))
	fmt.Println(bst.Root(), bst.Depth())
}
