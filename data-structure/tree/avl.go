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

func (node *avlNode) insert1(element Element) *avlNode {
	cmp := node.element.Compare(element)
	switch {
	case cmp > 0:
		if node.leftChild == nil {
			node.leftChild = newAVLNode(element)
		} else {
			node.leftChild = node.leftChild.insert1(element)
		}
	case cmp < 0:
		if node.rightChild == nil {
			node.rightChild = newAVLNode(element)
		} else {
			node.rightChild = node.rightChild.insert1(element)
		}
	case cmp == 0:
		return node
	}

	node.height = pkg.MaxInt(getHeight(node.leftChild), getHeight(node.rightChild)) + 1

	newRoot, balanced := balance(node)
	if balanced {
		return newRoot
	} else {
		return node
	}
}

// 平衡以root为根节点的子树, 并返回平衡后该子树新的根节点
func balance(root *avlNode) (newRoot *avlNode, balanced bool) {
	const bFactor = 2

	diff := getHeight(root.leftChild) - getHeight(root.rightChild)
	// 如果高度差小于平衡因子，则不需要平衡，直接返回原节点
	if pkg.AbsInt(diff) < bFactor {
		newRoot, balanced = root, false
		return
	}

	switch {
	case diff > 0: // 左子树比右子树高
		leftTree := root.leftChild
		if getHeight(leftTree.leftChild) > getHeight(leftTree.rightChild) {
			newRoot, balanced = rightRotate(root), true
		} else {
			newRoot, balanced = leftRightRotate(root), true
		}
	case diff < 0: // 右子树比左子树高
		rightTree := root.rightChild
		if getHeight(rightTree.rightChild) > getHeight(rightTree.leftChild) {
			newRoot, balanced = leftRotate(root), true
		} else {
			newRoot, balanced = rightLeftRotate(root), true
		}
	}
	return
}

// LL类型旋转：右单旋转
// 1. root的左孩子上升为父节点
// 2. root下降为其左孩子的右孩子
// 3. root的左孙子成为其左兄弟，root的右孙子变为其左孩子
// 触发场景: (假设不平衡子树的根结点为root, 新插入的节点为N)
// N为root的左孩子的左孩子
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
// 触发场景: (假设不平衡子树的根结点为root, 新插入的节点为N)
// N为root的右孩子的右孩子
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

// LR类型旋转: 先左旋转然后右旋转
// 1. 先对root的左孩子进行左旋转
// 2. 然后对root进行右旋转
// 触发场景: (假设不平衡子树的根结点为root, 新插入的节点为N)
// N为root左孩子的右孩子
func leftRightRotate(root *avlNode) *avlNode {
	root.leftChild = leftRotate(root.leftChild)
	return rightRotate(root)
}

// RL类型旋转: 先右旋转然后左旋转
// 1. 先对root的有孩子进行右旋转
// 2. 然后对root进行左旋转
// 触发场景: (假设不平衡子树的根结点为root, 新插入的节点为N)
// N为root的右孩子的左孩子
func rightLeftRotate(root *avlNode) *avlNode {
	root.rightChild = rightRotate(root.rightChild)
	return leftRotate(root)
}
func getHeight(node Node) int {
	return node.Height()
}

// AVL 二叉平衡树
type AVL struct {
	root *avlNode
}

func NewAVL() *AVL {
	return &AVL{}
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

func (avl *AVL) Insert(element Element) {
	if avl == nil || element == nil {
		return
	}

	// 插入的是第一个节点
	if avl.root == nil || avl.root.isZero() {
		if avl.root == nil {
			avl.root = newAVLNode(element)
		} else {
			avl.root.element = element
			avl.root.height = 1
		}
	} else {
		avl.root = avl.root.insert1(element)
	}
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
