package tree

import (
	"fmt"
)

/*
	二叉搜索(查找)树
*/

/**********************************************************************************************************************/

// BSTNode 二叉搜索平衡树节点
type BSTNode struct {
	initial    bool     // 必须通过New方法构建
	depth      int      // 节点所处深度
	element    Element  // 元素
	parent     *BSTNode // 父节点
	leftChild  *BSTNode // 左孩子
	rightChild *BSTNode // 右孩子
}

func NewBSTNode(e Element) *BSTNode {
	return &BSTNode{
		element: e,
		initial: true,
	}
}

func (node *BSTNode) assert() {
	if !node.initial {
		panic("node not init correctly")
	}
}

func (node *BSTNode) isZero() bool {
	if node.depth != 0 {
		return false
	}
	if node.element != nil || node.parent != nil {
		return false
	}
	if node.leftChild != nil || node.rightChild != nil {
		return false
	}
	return !node.initial
}

func (node *BSTNode) reset() {
	*node = BSTNode{
		initial: true,
	}
}

func (node *BSTNode) String() string {
	if node == nil {
		return "<nil>"
	}
	node.assert()
	return fmt.Sprintf("(value: %v, depth: %d)", node.element.Value(), node.depth)
}

func (node *BSTNode) Data() Element {
	if node == nil {
		return nil
	}
	node.assert()
	if node.element == nil {
		return nil
	}
	return node.element
}

// Root 返回自己所在树的根节点
func (node *BSTNode) Root() Node {
	if node == nil {
		return nil
	}

	node.assert()
	p := node
	for p.parent != nil {
		p = p.parent
	}
	return p
}

// LeftChild 返回自己的左子树
func (node *BSTNode) LeftChild() Node {
	if node == nil {
		return nil
	}

	node.assert()
	if node.leftChild == nil {
		return nil
	}
	return node.leftChild
}

// RightChild 返回自己的右子树
func (node *BSTNode) RightChild() Node {
	if node == nil {
		return nil
	}
	node.assert()
	if node.rightChild == nil {
		return nil
	}
	return node.rightChild
}

// RightSibling 返回节点的右兄弟, 可能是自己
func (node *BSTNode) RightSibling() Node {
	if node == nil {
		return nil
	}
	node.assert()

	parent := node.parent
	if parent != nil {
		return parent.rightChild
	}
	return nil
}

// LeftSibling 返回节点的左兄弟, 可能是自己
func (node *BSTNode) LeftSibling() Node {
	if node == nil {
		return nil
	}
	node.assert()
	parent := node.parent
	if parent != nil {
		return parent.leftChild
	}
	return nil
}

func (node *BSTNode) Parent() Node {
	if node == nil {
		return nil
	}
	node.assert()
	if node.parent == nil {
		return nil
	}
	return node.parent
}

func (node *BSTNode) Depth() int {
	if node == nil {
		return 0
	}
	node.assert()
	return node.depth
}

func (node *BSTNode) Color() Color {
	if node == nil {
		return NoColor
	}
	node.assert()
	return NoColor
}

func (node *BSTNode) getDepth() int {
	node.assert()

	switch {
	case node.leftChild == nil && node.rightChild == nil:
		return 1
	case node.leftChild != nil && node.rightChild == nil:
		return 1 + node.leftChild.getDepth()
	case node.leftChild == nil && node.rightChild != nil:
		return 1 + node.rightChild.getDepth()
	default:
		return 1 + maxInt(node.leftChild.getDepth(), node.rightChild.getDepth())
	}
}

// maxNode 获取node的最大子孙节点，包括自己
func (node *BSTNode) maxNode() *BSTNode {
	node.assert()
	p := node
	for p.rightChild != nil {
		p = p.rightChild
	}
	return p
}

// minNode 获取node的最小子孙节点，包括自己
func (node *BSTNode) minNode() *BSTNode {
	node.assert()
	p := node
	for p.leftChild != nil {
		p = p.leftChild
	}
	return p
}

func (node *BSTNode) insert(element Element) *BSTNode {
	node.assert()

	var (
		cmp        = 0
		p          = node
		rookieNode = NewBSTNode(element)
	)

loop:
	for p != nil {
		cmp = p.element.Compare(element)
		switch {
		case cmp > 0:
			if p.leftChild == nil {
				rookieNode.depth = p.depth + 1
				rookieNode.parent = p
				p.leftChild = rookieNode
				break loop
			}
			p = p.leftChild
		case cmp < 0:
			if p.rightChild == nil {
				rookieNode.depth = p.depth + 1
				rookieNode.parent = p
				p.rightChild = rookieNode
				break loop
			}
			p = p.rightChild
		case cmp == 0:
			rookieNode = p
			break loop
		}
	}
	return rookieNode
}

func (node *BSTNode) remove(element Element) bool {
	node.assert()
	p := node
loop:
	for p != nil {
		cmp := p.element.Compare(element)
		switch {
		case cmp > 0:
			if p.leftChild == nil {
				return false
			}
			p = p.leftChild
		case cmp < 0:
			if p.rightChild == nil {
				return false
			}
			p = p.rightChild
		case cmp == 0:
			break loop
		}
	}

	// N为0时节点需要删除
	switch {
	case p.rightChild != nil: // 被删除节点存在右子树，则将右子树的最小节点位置提升到p处，然后从右子树中删除该最小节点
		mNode := p.rightChild.minNode()
		p.element = mNode.element
		p.rightChild.remove(mNode.element)

	case p.leftChild != nil: // 被删除节点只存在左子树，则将左子树中的最大节点提升至p处，然后从左子树中删除该最大节点
		mNode := p.leftChild.maxNode()
		p.element = mNode.element
		p.leftChild.remove(mNode.element)

	default: // 被删除节点没有孩子节点，直接删除该节点
		if p.parent == nil {
			// 头节点
			p.reset()
		} else {
			if p.parent.leftChild == p {
				p.parent.leftChild = nil
			} else {
				p.parent.rightChild = nil
			}
		}
	}
	return true
}

func (node *BSTNode) update(old, element Element) bool {
	node.assert()
	// 1. 旧值不存在，返回删除失败
	oNode := node.find(old)
	if oNode == nil {
		return false
	}
	// 2. 如果新的值已经在树中存在，直接将old删除即可
	eNode := node.find(element)
	if eNode != nil {
		return node.remove(old)
	}
	// 3. 如果新值在树中不存在，则先删除旧值
	if !node.remove(old) {
		return false
	}

	// 4. 最后插入新值
	return node.insert(element) != nil
}

func (node *BSTNode) find(element Element) *BSTNode {
	node.assert()
	p := node
loop:
	for p != nil {
		cmp := p.element.Compare(element)
		switch {
		case cmp > 0:
			if p.leftChild == nil {
				return nil
			}
			p = p.leftChild
		case cmp < 0:
			if p.rightChild == nil {
				return nil
			}
			p = p.rightChild
		case cmp == 0:
			break loop
		}
	}
	return p
}

/**********************************************************************************************************************/

// BST 二叉搜索平衡树
type BST struct {
	root *BSTNode // 根节点
}

func NewBST() *BST {
	return &BST{}
}

// Root 返回树的根节点
func (tree *BST) Root() Node {
	if tree == nil {
		return nil
	}
	if tree.root == nil {
		return nil
	}
	return tree.root
}

// Depth 返回树的深度
func (tree *BST) Depth() int {
	if tree == nil || tree.root == nil {
		return 0
	}
	return tree.root.getDepth()
}

// Insert 插入节点
func (tree *BST) Insert(element Element) (node Node) {
	if tree == nil || element == nil {
		return nil
	}

	if tree.root == nil || tree.root.isZero() {
		if tree.root == nil {
			tree.root = NewBSTNode(element)
		} else {
			tree.root.element = element
		}
		tree.root.depth = 1
		node = tree.root
	} else {
		node = tree.root.insert(element)
	}
	return
}

// Remove 移除与element相等的节点
func (tree *BST) Remove(element Element) bool {
	if tree == nil || tree.root == nil {
		return false
	}
	return tree.root.remove(element)
}

// Find 查找与element相等的节点
func (tree *BST) Find(element Element) Node {
	if tree == nil || tree.root == nil {
		return nil
	}
	return tree.root.find(element)
}

// Update 将old处的值更新为element
func (tree *BST) Update(old, element Element) bool {
	if tree == nil || tree.root == nil {
		return false
	}
	return tree.root.update(old, element)
}