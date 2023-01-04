package tree

import (
	"fmt"
)

/*
	红黑树的特性:
	1. 每个节点要么是红色要么是黑色
	2. 根节点是黑色的
	3. 如果一个节点是红色的, 则它的两个子节点都是黑色的
	4. 对每个节点, 从该节点到其所有后代节点的简单路径上, 均包含相同数目的黑色节点
*/

// nil节点，即叶子节点，叶子节点均为黑色
var nilNode = &rbNode{
	color:  Black,
	height: 1,
}

type rbNode struct {
	color      Color   // 节点颜色
	height     int     //  节点高度
	element    Element // 节点存储的数据
	parent     *rbNode // 父节点
	leftChild  *rbNode // 左孩子
	rightChild *rbNode // 右孩子
}

func newRBNode(element Element) *rbNode {
	return &rbNode{
		color:      Black,
		height:     2,
		element:    element,
		parent:     nilNode,
		leftChild:  nilNode,
		rightChild: nilNode,
	}
}

func (node *rbNode) isZero() bool {
	if node.color != NoColor {
		return false
	}
	if node.height != 0 {
		return false
	}
	if node.element != nil {
		return false
	}
	if node.parent != nilNode {
		return false
	}
	if node.leftChild != nilNode {
		return false
	}
	if node.rightChild != nilNode {
		return false
	}
	return true
}

func isNil(node *rbNode) bool {
	if node == nil {
		return true
	}
	if node == nilNode {
		return true
	}
	return false
}

func (node *rbNode) reset() {
	if node == nil {
		return
	}
	*node = rbNode{
		color:      NoColor,
		height:     0,
		element:    nil,
		parent:     nilNode,
		leftChild:  nilNode,
		rightChild: nilNode,
	}
}

// String
func (node *rbNode) String() string {
	if node == nil {
		return "<nil>"
	}
	v, color := "", ""
	if node.element == nil {
		v = "<nil>"
	} else {
		v = fmt.Sprint(node.element.Value())
	}
	color = fmt.Sprint(node.color)

	return fmt.Sprintf("%v(%v)", v, color)
}

// Depth 返回自己所处的深度, 深度从根节点到自己所经历的节点数量
func (node *rbNode) Depth() int {
	if isNil(node) {
		return 0
	}
	switch {
	case node.parent == nilNode:
		return 1
	default:
		return 1 + node.parent.Depth()
	}
}

// Height 返回自己所处的高度, 高度为从叶子节点到自己所经历的节点数量
func (node *rbNode) Height() int {
	if node != nil {
		return node.height
	}
	return 0
}

// Data 节点存储的数据
func (node *rbNode) Data() Element {
	if isNil(node) {
		return nil
	}
	return node.element
}

// Root 返回根节点
func (node *rbNode) Root() Node {
	if node == nil {
		return nil
	}
	p := node
	for p.parent != nil && p.parent != nilNode {
		p = p.parent
	}
	return p
}

// LeftChild 左孩子
func (node *rbNode) LeftChild() Node {
	if isNil(node) || isNil(node.leftChild) {
		return nil
	}
	return node.leftChild
}

// RightChild 右孩子
func (node *rbNode) RightChild() Node {
	if isNil(node) || isNil(node.rightChild) {
		return nil
	}
	return node.rightChild
}

// LeftSibling 左兄弟
func (node *rbNode) LeftSibling() Node {
	if node == nil {
		return nil
	}
	if isNil(node.parent) {
		return nil
	}
	if leftChild := node.parent.leftChild; isNil(leftChild) || leftChild == node {
		return nil
	} else {
		return leftChild
	}
}

// RightSibling 右兄弟
func (node *rbNode) RightSibling() Node {
	if node == nil {
		return nil
	}
	if isNil(node.parent) {
		return nil
	}
	if rightChild := node.parent.rightChild; isNil(rightChild) || rightChild == node {
		return nil
	} else {
		return rightChild
	}
}

// Parent 父节点
func (node *rbNode) Parent() Node {
	if node == nil {
		return nil
	}
	if isNil(node.parent) {
		return nil
	}
	return node.parent
}

// Color 返回节点颜色
func (node *rbNode) Color() Color {
	if node == nil {
		return NoColor
	}
	return node.color
}

func (node *rbNode) insert(tree *RedBlackTree, element Element) {
	var (
		cmp  int
		z    = newRBNode(element)
		y    = nilNode
		head = node
		p    = node
	)
	for p != nilNode {
		y = p
		cmp = node.element.Compare(element)
		switch {
		case cmp > 0:
			p = p.leftChild
		case cmp < 0:
			p = p.rightChild
		case cmp == 0: // 如果元素已经存在, 直接返回头节点
			tree.root = head
			return
		}
	}
	switch {
	case y == nilNode: // 插入的是头节点, 不需要旋转
		tree.root = z
		return
	case cmp > 0: // z作为左孩子
		z.parent = y
		y.leftChild = z
	case cmp < 0: // z作为右孩子
		z.parent = y
		y.rightChild = z
	}

	z.color = Red

	// 旋转达到平衡
	tree.fixUpInsert(z)
	return
}

func (node *rbNode) find(element Element) *rbNode {
	p := node
loop:
	for p != nilNode {
		cmp := p.element.Compare(element)
		switch {
		case cmp > 0:
			p = p.leftChild
		case cmp < 0:
			p = p.rightChild
		case cmp == 0:
			break loop
		}
	}
	return p
}

func (node *rbNode) remove(tree *RedBlackTree, element Element) bool {
	var (
		p             = node.find(element) // 找到要删除的节点
		x             *rbNode
		y             = p
		originalColor = y.color
	)

	// 元素不存在
	if isNil(p) {
		return false
	}
	switch {
	case p.leftChild == nilNode:
		x = p.rightChild
		tree.replace(p, p.rightChild)
	case p.rightChild == nilNode:
		x = p.leftChild
		tree.replace(p, p.leftChild)
	default:
		y = getMinNode(p.rightChild).(*rbNode)
		originalColor = y.color
		x = y.rightChild
		if y.parent == p {
			x.parent = y
		} else {
			tree.replace(y, y.rightChild)
			y.rightChild = p.rightChild
			y.rightChild.parent = y
		}
		tree.replace(p, y)
		y.leftChild = p.leftChild
		y.leftChild.parent = y
		y.color = p.color
	}
	if originalColor == Black {
		tree.fixUpRemove(x)
	}
	return true
}

type RedBlackTree struct {
	root *rbNode
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		root: nilNode,
	}
}

func (tree *RedBlackTree) fixUpRemove(node *rbNode) {
	for node != tree.root && node.color == Black {
		if node == node.parent.leftChild {
			w := node.parent.rightChild
			if w.color == Red {
				w.color = Black
				node.parent.color = Red
				tree.leftRotate(node.parent)
				w = node.parent.rightChild
			}
			if w.leftChild.color == Black && w.rightChild.color == Black {
				w.color = Red
				node = node.parent
			} else if w.rightChild.color == Black {
				w.leftChild.color = Black
				w.color = Red
				tree.rightRotate(w)
				w = node.parent.rightChild

				w.color = node.parent.color
				node.parent.color = Black
				w.rightChild.color = Black
				tree.leftRotate(node.parent)
				node = tree.root
			} else {
				w.color = node.parent.color
				node.parent.color = Black
				w.rightChild.color = Black
				tree.leftRotate(node.parent)
				node = tree.root
			}
		} else {
			w := node.parent.leftChild
			if w.color == Red {
				w.color = Black
				node.parent.color = Red
				tree.rightRotate(node.parent)
				w = node.parent.leftChild
			}
			if w.rightChild.color == Black && w.leftChild.color == Black {
				w.color = Red
				node = node.parent
			} else if w.leftChild.color == Black {
				w.rightChild.color = Black
				w.color = Red
				tree.leftRotate(w)
				w = node.parent.leftChild

				w.color = node.parent.color
				node.parent.color = Black
				w.leftChild.color = Black
				tree.rightRotate(node.parent)
				node = tree.root
			} else {
				w.color = node.parent.color
				node.parent.color = Black
				w.leftChild.color = Black
				tree.rightRotate(node.parent)
				node = tree.root
			}
		}
	}
	node.color = Black
}

func (tree *RedBlackTree) fixUpInsert(node *rbNode) {
	var (
		y *rbNode
	)
	for node.parent.color == Red {
		if node.parent == node.parent.parent.leftChild {
			y = node.parent.parent.rightChild
			// 父节点以及父节点的兄弟节点都是红色
			if y.color == Red {
				node.parent.color = Black
				y.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else if node == node.parent.rightChild {
				node = node.parent
				tree.leftRotate(node)
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.rightRotate(node.parent.parent)
			} else {
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.rightRotate(node.parent.parent)
			}
		} else {
			y = node.parent.parent.leftChild
			if y.color == Red {
				node.parent.color = Black
				y.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else if node == node.parent.leftChild {
				node = node.parent
				tree.rightRotate(node)
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.leftRotate(node.parent.parent)
			} else {
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.leftRotate(node.parent.parent)
			}
		}
	}
	tree.root.color = Black
}

func (tree *RedBlackTree) leftRotate(node *rbNode) {
	right := node.rightChild
	// 将node右孩子的左孩子提升为自己的右孩子
	node.rightChild = right.leftChild
	// 将node右孩子的左孩子的父节点变更为node
	if right.leftChild != nilNode {
		right.leftChild.parent = node
	}
	// 用node的右孩子替换自己的位置
	right.parent = node.parent
	// 变更node父节点和node右孩子的关系
	switch {
	case node.parent == nilNode: // node父节点为nilNode，证明node是头节点
		tree.root = right
	case node == node.parent.leftChild: // node为父节点的左孩子，则将父节点的左孩子变更为right
		node.parent.leftChild = right
	default: // node为父节点的右孩子，则将父节点的右孩子变更为right
		node.parent.rightChild = right
	}
	// 变更right和node关系：将node变为right的左孩子
	right.leftChild = node
	node.parent = right
}

func (tree *RedBlackTree) rightRotate(node *rbNode) {
	left := node.leftChild
	node.leftChild = left.rightChild
	if left.rightChild != nilNode {
		left.rightChild.parent = node
	}
	left.parent = node.parent
	switch {
	case node.parent == nilNode:
		tree.root = left
	case node == node.parent.rightChild:
		node.parent.rightChild = left
	default:
		node.parent.leftChild = left
	}
	left.rightChild = node
	node.parent = left
}

func (tree *RedBlackTree) replace(n1, n2 *rbNode) {
	switch {
	case n1.parent == nilNode:
		tree.root = n2
	case n1.parent.leftChild == n1:
		n1.parent.leftChild = n2
	case n1.parent.rightChild == n1:
		n1.parent.rightChild = n2
	}
	n2.parent = n1.parent
}

func (tree *RedBlackTree) Root() Node {
	if tree == nil || tree.root == nil {
		return nil
	}
	return tree.root
}

func (tree *RedBlackTree) Depth() int {
	if tree == nil || tree.root == nil {
		return 0
	}
	return tree.root.Height()
}

func (tree *RedBlackTree) Insert(element Element) {
	if tree == nil || element == nil {
		return
	}
	tree.root.insert(tree, element)
}

func (tree *RedBlackTree) Remove(element Element) bool {
	if tree == nil || element == nil {
		return false
	}
	if isNil(tree.root) {
		return false
	}
	return tree.root.remove(tree, element)
}

func (tree *RedBlackTree) Find(element Element) Node {
	if tree == nil || element == nil {
		return nil
	}
	if isNil(tree.root) {
		return nil
	}
	node := tree.root.find(element)
	if isNil(node) {
		return nil
	}
	return node
}

func (tree *RedBlackTree) Update(old, element Element) bool {
	if tree == nil || element == nil || old == nil {
		return false
	}
	if isNil(tree.root) {
		return false
	}
	if !tree.root.remove(tree, old) {
		return false
	}
	tree.root.insert(tree, element)
	return true
}
