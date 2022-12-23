package tree

import (
	"fmt"
	"testing"
)

var (
	traverseBST Tree = NewBST()
)

func init() {
	testData := []TestInt{
		2, 1, 8, 5, 3, 4, 9, 6, 10, 7,
	}
	for _, data := range testData {
		traverseBST.Insert(data)
	}
}

func TestPreOrderTraverse(t *testing.T) {
	fmt.Println(PreOrderTraverseRecursion(traverseBST.Root()))
	fmt.Println(PreOrderTraverse(traverseBST.Root()))
	fmt.Println()

	// output:
	// 2->1->8->5->3->4->6->7->9->10
	// 2->1->8->5->3->4->6->7->9->10
}

func TestInOrderTraverse(t *testing.T) {
	fmt.Println(InOrderTraverseRecursion(traverseBST.Root()))
	fmt.Println(InOrderTraverse(traverseBST.Root()))
	fmt.Println()
	// output
	// 1->2->3->4->5->6->7->8->9->10
	// 1->2->3->4->5->6->7->8->9->10
}

func TestPostOrderTraverse(t *testing.T) {
	fmt.Println(PostOrderTraverseRecursion(traverseBST.Root()))
	fmt.Println(PostOrderTraverse(traverseBST.Root()))
	fmt.Println()
	// output
	// 1->4->3->7->6->5->10->9->8->2
	// 1->4->3->7->6->5->10->9->8->2
}

func TestBFSTraverse(t *testing.T) {
	fmt.Println(BFSTraverse(traverseBST.Root()))
	fmt.Println()
	// output
	// 2->1->8->5->9->3->6->10->4->7
}

func TestDFSTraverse(t *testing.T) {
	fmt.Println(DFSTraverse(traverseBST.Root()))
	fmt.Println()
	// output
	// 	2->1->8->5->3->4->6->7->9->10
}
