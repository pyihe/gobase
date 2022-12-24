package tree

import (
	"fmt"
	"testing"
)

func TestAVL_Insert(t *testing.T) {
	var (
		avl      Tree = NewAVL()
		testData      = []TestInt{1, 2, 3, 4, 5, 6, 7}
	)

	//4->2->1->3->6->5->7
	for _, data := range testData {
		avl.Insert(data)
	}

	fmt.Println(PreOrderTraverseRecursion(avl.Root()))

	fmt.Println(avl.Update(TestInt(4), TestInt(10)))
	fmt.Println(avl.Update(TestInt(3), TestInt(11)))

	fmt.Println(PreOrderTraverseRecursion(avl.Root()))
}
