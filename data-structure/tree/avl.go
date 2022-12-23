package tree

import (
	"fmt"

	"github.com/pyihe/gobase/pkg"
)

type avlNode struct {
	height     int      // 节点所处高度
	element    Element  // 节点存储的数据
	parent     *avlNode // 父节点
	leftChild  *avlNode // 左孩子
	rightChild *avlNode // 右孩子
}

func newAVLNode(ele Element) *avlNode {
	return &avlNode{
		height:  1,
		element: ele,
	}
}

func (node *avlNode) isZero() bool {
	if node.height != 0 {
		return false
	}
	if node.element != nil || node.parent != nil {
		return false
	}
	if node.leftChild != nil || node.rightChild != nil {
		return false
	}
	return true
}

func (node *avlNode) String() string { // String
	if node == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", node.element.Value())
}
func (node *avlNode) Depth() int { // 返回自己所处的深度, 深度从根节点到自己所经历的节点数量
	if node == nil {
		return 0
	}
	switch {
	case node.parent == nil:
		return 1
	default:
		return 1 + node.parent.Depth()
	}
}
func (node *avlNode) Height() int { // 返回自己所处的高度, 高度为从叶子节点到自己所经历的节点数量
	if node == nil {
		return 0
	}
	return node.height
}
func (node *avlNode) Data() Element { // 节点存储的数据
	if node == nil {
		return nil
	}
	if node.element == nil {
		return nil
	}
	return node.element
}
func (node *avlNode) Root() Node { // 返回根节点
	p := node
	for p != nil {
		p = p.parent
	}
	return p
}
func (node *avlNode) LeftChild() Node { // 左孩子
	if node == nil || node.leftChild == nil {
		return nil
	}
	return node.leftChild
}
func (node *avlNode) RightChild() Node { // 右孩子
	if node == nil || node.rightChild == nil {
		return nil
	}
	return node.rightChild
}
func (node *avlNode) LeftSibling() Node { // 左兄弟
	if node == nil {
		return nil
	}
	parent := node.parent
	if parent == nil || parent.leftChild == nil || node == parent.leftChild {
		return nil
	}
	return parent.leftChild
}
func (node *avlNode) RightSibling() Node { // 右兄弟
	if node == nil {
		return nil
	}
	parent := node.parent
	// 如果自己就是父节点的右孩子，则右兄弟为nil
	if parent == nil || parent.rightChild == nil || node == parent.rightChild {
		return nil
	}
	return parent.rightChild
}
func (node *avlNode) Parent() Node { // 父节点
	if node == nil || node.parent == nil {
		return nil
	}
	return node.parent
}
func (node *avlNode) Color() Color { // 返回节点颜色
	if node == nil {
		return NoColor
	}
	return NoColor
}

func (node *avlNode) insert(element Element) (*avlNode, bool) {
	var (
		created    = false
		cmp        = 0
		p          = node
		rookieNode = newAVLNode(element)
	)

loop:
	for p != nil {
		cmp = p.element.Compare(element)
		// 每往下走一层，新节点的高度+1
		rookieNode.height += 1
		switch {
		case cmp > 0:
			if p.leftChild == nil {
				rookieNode.parent = p
				p.leftChild = rookieNode
				created = true
				break loop
			}
			p = p.leftChild
		case cmp < 0:
			if p.rightChild == nil {
				rookieNode.parent = p
				p.rightChild = rookieNode
				created = true
				break loop
			}
			p = p.rightChild
		case cmp == 0:
			rookieNode = p
			break loop
		}
	}
	// 如果没有创造新节点，则直接返回，计算高度
	if !created {
		goto end
	}
	// 创造了新节点, 需要重新计算插入路径上的每个节点的高度
	if rookieNode.parent != nil {
		// 父节点左右孩子都不为空，证明插入的叶子节点没有影响父节点高度，不需要计算高度，也不需要重新平衡
		if rookieNode.parent.leftChild != nil && rookieNode.parent.rightChild != nil {
			created = false
			goto end
		}
		// 如果新节点的父节点只有一个孩子节点，则父节点的高度受影响，需要更新整个插入路径上每个节点的高度
		p = rookieNode.parent
		for p != nil {
			p.height = pkg.MaxInt(getHeight(p.leftChild), getHeight(p.rightChild))
			p = p.parent
		}
	}
end:
	return rookieNode, created
}

// 平衡以root为根节点的子树, 并返回平衡后该子树新的根节点
func balance(root *avlNode) (*avlNode, bool) {
	const bFactor = 2

	diff := getHeight(root.leftChild) - getHeight(root.rightChild)
	// 如果高度差小于平衡因子，则不需要平衡，直接返回原节点
	if pkg.AbsInt(diff) < bFactor {
		return root, false
	}

	// 需要平衡

	return nil, true
}

// LL类型旋转：右单旋转
// 1. root的左孩子上升为父节点
// 2. root下降为其左孩子的右孩子
// 3. root的左孙子成为其左兄弟，root的右孙子变为其左孩子
func rightRotate(root *avlNode) *avlNode {
	left := root.leftChild
	root.leftChild = left.rightChild
	left.rightChild = root

	// 更新参与旋转的节点的高度
	root.height = pkg.MaxInt(getHeight(root.leftChild), getHeight(root.rightChild)) + 1
	left.height = pkg.MaxInt(getHeight(left.leftChild), getHeight(left.rightChild)) + 1
	return left
}

// RR类型旋转: 左单旋转
// 1. root的右孩子上升为父节点
// 2. root下降为其右孩子的左孩子
// 3. root的右孙子成为其右兄弟，左孙子成为其右孩子
func leftRotate(root *avlNode) *avlNode {
	right := root.rightChild
	root.rightChild = right.leftChild
	right.leftChild = root

	// 更新参与旋转的节点的高度
	// 这里之所以能够这样来更新高度，是因为AVL的平衡因子为2，子树之间的高度差为2时即需要旋转调整平衡
	// 所以对于单旋转来说，其形状必然只有一种情况
	root.height = pkg.MaxInt(getHeight(root.leftChild), getHeight(root.rightChild)) + 1
	right.height = pkg.MaxInt(getHeight(right.leftChild), getHeight(right.rightChild)) + 1
	return right
}
func getHeight(node Node) int {
	return node.Height()
}

type AVL struct {
	root *avlNode
}

func (avl *AVL) Root() Node {
	if avl == nil || avl.root == nil {
		return nil
	}
	return avl.root
}

func (avl *AVL) Depth() int {
	if avl == nil || avl.root == nil {
		return 0
	}
	return avl.root.Height()
}

func (avl *AVL) Insert(element Element) Node {
	if avl == nil || element == nil {
		return nil
	}

	var (
		balanced    bool
		needBalance bool
		rookieNode  *avlNode
		grandFather *avlNode
		newRoot     *avlNode
	)

	// 插入的是第一个节点
	if avl.root == nil || avl.root.isZero() {
		if avl.root == nil {
			avl.root = newAVLNode(element)
		} else {
			avl.root.element = element
			avl.root.height = 1
		}
		rookieNode = avl.root
	} else {
		rookieNode, needBalance = avl.root.insert(element)
	}

	if !needBalance || rookieNode.parent == nil {
		goto end
	}

	// 因为AVL的平衡因子为2，所以对于新插入的节点
	// 发生不平衡的节点必然只能是其爷爷节点
	// 所以只需要平衡爷爷节点即可
	grandFather = rookieNode.parent.parent
	if grandFather == nil {
		goto end
	}
	newRoot, balanced = balance(grandFather)
	if !balanced {
		goto end
	}
	if grandFather.parent == nil { // 如果爷爷节点是根节点，则需要更新AVL树的根节点为平衡后返回的根节点
		avl.root = newRoot
	} else {
		// 否则只需要根新爷爷节点的父节点与孩子节点的关系即可
		if grandFather.parent.leftChild == grandFather {
			grandFather.parent.leftChild = newRoot
		} else {
			grandFather.parent.rightChild = newRoot
		}
	}
end:
	return rookieNode
}

func (avl *AVL) Remove(element Element) bool {
	return false
}

func (avl *AVL) Update(old, new Element) bool {
	return false
}

func (avl *AVL) Find(element Element) Node {
	return nil
}
