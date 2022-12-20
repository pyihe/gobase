package tree

import (
	"container/list"
	"fmt"
)

/*
	单纯的二叉树没有实际意义，需要赋予除了二叉以外的其他特性，如平衡二叉树、二叉搜索树等
*/

// 二叉搜索树特性：
// 1. 二叉树
// 2. 平衡树
// 3. 搜索树

// Element 树节点存储的元素
type Element interface {
	Value() interface{}  // 元素值
	Compare(Element) int // 两个元素相比，大于返回>0, 小于返回<0, 等于返回=0
}

/**********************************************************************************************************************/

// BinarySearchNode 二叉搜索平衡树节点
type BinarySearchNode struct {
	N       int               // 节点存储Element的个数
	Height  int               // 节点所处深度
	Element Element           // 元素
	Parent  *BinarySearchNode // 父节点
	Left    *BinarySearchNode // 左孩子
	Right   *BinarySearchNode // 右孩子
	initial bool              // 必须通过New方法构建
}

func NewBinarySearchNode(e Element) *BinarySearchNode {
	return &BinarySearchNode{
		N:       1,
		Element: e,
		initial: true,
	}
}

func (node *BinarySearchNode) assert() {
	if !node.initial {
		panic("node not init correctly")
	}
}

func (node *BinarySearchNode) reset() {
	*node = BinarySearchNode{}
}

func (node *BinarySearchNode) String() string {
	return fmt.Sprintf("(value: %v, height: %d)", node.Element.Value(), node.Height)
}

// Root 返回某个节点的根节点
func (node *BinarySearchNode) Root() *BinarySearchNode {
	node.assert()

	p := node
	for p.Parent != nil {
		p = p.Parent
	}
	return p
}

// RightSibling 返回节点的右兄弟, 可能是自己
func (node *BinarySearchNode) RightSibling() *BinarySearchNode {
	node.assert()

	parent := node.Parent
	if parent == nil {
		return nil
	}
	return parent.Right
}

// LeftSibling 返回节点的左兄弟, 可能是自己
func (node *BinarySearchNode) LeftSibling() *BinarySearchNode {
	node.assert()

	parent := node.Parent
	if parent == nil {
		return nil
	}
	return parent.Left
}

// maxNode 获取node的最大子孙节点，包括自己
func (node *BinarySearchNode) maxNode() *BinarySearchNode {
	p := node
	for p.Right != nil {
		p = p.Right
	}
	return p
}

// minNode 获取node的最小子孙节点，包括自己
func (node *BinarySearchNode) minNode() *BinarySearchNode {
	p := node
	for p.Left != nil {
		p = p.Left
	}
	return p
}

func (node *BinarySearchNode) insert(element Element) *BinarySearchNode {
	var (
		cmp        = 0
		p          = node
		rookieNode = NewBinarySearchNode(element)
	)

loop:
	for p != nil {
		cmp = p.Element.Compare(element)
		switch {
		case cmp > 0:
			if p.Left == nil {
				rookieNode.Height = p.Height + 1
				rookieNode.Parent = p
				p.Left = rookieNode
				break loop
			}
			p = p.Left
		case cmp < 0:
			if p.Right == nil {
				rookieNode.Height = p.Height + 1
				rookieNode.Parent = p
				p.Right = rookieNode
				break loop
			}
			p = p.Right
		case cmp == 0:
			p.N += 1
			rookieNode = p
			break loop
		}
	}
	return rookieNode
}

func (node *BinarySearchNode) remove(element Element) bool {
	p := node
	for p != nil {
		cmp := p.Element.Compare(element)
		switch {
		case cmp > 0:
			if p.Left == nil {
				return false
			}
			p = p.Left
		case cmp < 0:
			if p.Right == nil {
				return false
			}
			p = p.Right
		case cmp == 0:
			break
		}
	}

	p.N -= 1
	if p.N > 0 {
		return true
	}

	// N为0时节点需要删除
	switch {
	case p.Right != nil:
		mNode := p.Right.minNode()
		p.Element = mNode.Element
		p.N = mNode.N
		mNode.N = 1
		p.Right.remove(mNode.Element)

	case p.Left != nil:
		mNode := p.Left.maxNode()
		p.Element = mNode.Element
		p.N = mNode.N
		mNode.N = 1
		p.Left.remove(mNode.Element)

	default:
		if p.Parent == nil {
			// 头节点
			p.Parent.reset()
		} else {
			if p.Parent.Left == p {
				p.Parent.Left = nil
			} else {
				p.Parent.Right = nil
			}
		}
	}

	return true
}

func (node *BinarySearchNode) update(old, element Element) bool {
	target := node.find(old)
	if target == nil {
		return false
	}
	target.Element = element
	return true
}

func (node *BinarySearchNode) find(element Element) *BinarySearchNode {
	p := node
loop:
	for p != nil {
		cmp := p.Element.Compare(element)
		switch {
		case cmp > 0:
			if p.Left == nil {
				return nil
			}
			p = p.Left
		case cmp < 0:
			if p.Right == nil {
				return nil
			}
			p = p.Right
		case cmp == 0:
			break loop
		}
	}
	return p
}

func balance(node *BinarySearchNode) *BinarySearchNode {
	if getHeight(node.Right)-getHeight(node.Left) == 2 {
		if getHeight(node.Right.Right) > getHeight(node.Right.Left) {
			node = node.leftRotate()
		} else {
			node = node.rightLeftRotate()
		}
	} else if getHeight(node.Left)-getHeight(node.Right) == 2 {
		if getHeight(node.Left.Left) > getHeight(node.Left.Right) {
			node = node.rightRotate()
		} else {
			node = node.leftRightRotate()
		}
	}
	return node
}

// 向左旋转，返回旋转后的头节点
func (node *BinarySearchNode) leftRotate() *BinarySearchNode {
	temp := node.Right
	node.Right = temp.Left
	temp.Left = node

	node.Height = maxInt(getHeight(node.Left), getHeight(node.Right)) + 1
	temp.Height = maxInt(getHeight(temp.Left), getHeight(temp.Right)) + 1
	return temp
}

// 向右旋转
func (node *BinarySearchNode) rightRotate() *BinarySearchNode {
	temp := node.Left
	node.Left = temp.Right
	temp.Right = node

	node.Height = maxInt(getHeight(node.Left), getHeight(node.Right)) + 1
	temp.Height = maxInt(getHeight(temp.Left), getHeight(temp.Right)) + 1
	return temp
}

// 右孩子右旋转，自己右旋转
func (node *BinarySearchNode) rightLeftRotate() *BinarySearchNode {
	node.Right = node.Right.rightRotate()
	return node.leftRotate()
}

// 左孩子左旋转，自己右旋转
func (node *BinarySearchNode) leftRightRotate() *BinarySearchNode {
	node.Left = node.Left.leftRotate()
	return node.rightRotate()
}

func getHeight(node *BinarySearchNode) int {
	if node != nil {
		return node.Height
	}
	return 0
}

func maxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

/**********************************************************************************************************************/

// BinarySearchTree 二叉搜索平衡树
type BinarySearchTree struct {
	root *BinarySearchNode // 根节点
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Root 返回树的根节点
func (tree *BinarySearchTree) Root() *BinarySearchNode {
	return tree.root
}

// Height 返回树的深度
func (tree *BinarySearchTree) Height() int {
	if tree.root != nil {
		return tree.root.Height
	}
	return 0
}

// Insert 插入节点
func (tree *BinarySearchTree) Insert(element Element) (node *BinarySearchNode) {
	if element == nil {
		return nil
	}

	if tree.root == nil {
		tree.root = NewBinarySearchNode(element)
		node = tree.root
	} else {
		node = tree.root.insert(element)
	}

	balance(tree.root)

	return
}

// Remove 移除与element相等的节点
func (tree *BinarySearchTree) Remove(element Element) (ok bool) {
	if tree.root == nil {
		return false
	}
	if ok = tree.root.remove(element); ok {
		balance(tree.root)
	}
	return
}

// Find 查找与element相等的节点
func (tree *BinarySearchTree) Find(element Element) *BinarySearchNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.find(element)
}

// Update 将old处的值更新为element
func (tree *BinarySearchTree) Update(old, element Element) (ok bool) {
	if tree.root == nil {
		return false
	}
	if ok = tree.root.update(old, element); ok {
		balance(tree.root)
	}
	return
}

type treeStack struct {
	*list.List
}

func (s *treeStack) pop() interface{} {
	if s == nil || s.Len() <= 0 {
		return nil
	}
	value := s.Back()
	s.Remove(value)
	return value.Value
}

// 进栈
func (s *treeStack) push(d interface{}) {
	if s == nil {
		return
	}
	s.PushBack(d)
}

// 获取栈顶元素
func (s *treeStack) top() interface{} {
	if s == nil {
		return nil
	}
	return s.Back().Value
}

// PreOrderTraverse 前序遍历:以当前节点为根节点，根——>左——>右
func (tree *BinarySearchTree) PreOrderTraverse() (desc string) {
	s := &treeStack{
		List: list.New(),
	}
	p := tree.root
	for p != nil || s.Len() > 0 {
		if p != nil {
			s.push(p)
			if desc == "" {
				desc = fmt.Sprintf("%v", p)
			} else {
				desc = fmt.Sprintf("%s->%v", desc, p)
			}
			p = p.Left
		} else {
			p = s.pop().(*BinarySearchNode).Right
		}
	}
	return
}

// InOrderTraverse 中序遍历:以当前节点为根节点，左——>根——>右
func (tree *BinarySearchTree) InOrderTraverse() (desc string) {
	s := &treeStack{List: list.New()}
	p := tree.root
	for p != nil || s.Len() > 0 {
		if p != nil {
			s.PushBack(p)
			p = p.Left
		} else {
			ele := s.Back()
			s.Remove(ele)
			p = ele.Value.(*BinarySearchNode)
			if desc == "" {
				desc = fmt.Sprintf("%v", p)
			} else {
				desc = fmt.Sprintf("%s->%v", desc, p)
			}
			p = p.Right
		}
	}
	return
}

// PostOrderTraverse  后序遍历：以当前节点为根节点，左——>右——>根
func (tree *BinarySearchTree) PostOrderTraverse() (desc string) {
	s := &treeStack{List: list.New()}
	p := tree.root

	var (
		topNode  *BinarySearchNode
		lastNode *BinarySearchNode
	)

	for p != nil || s.Len() > 0 {
		if p != nil {
			s.PushBack(p)
			p = p.Left
		} else {
			ele := s.Back().Value
			topNode = ele.(*BinarySearchNode)
			if topNode.Right == nil || topNode.Right == lastNode {
				s.Remove(s.Back())
				lastNode = topNode
				if desc == "" {
					desc = fmt.Sprintf("%v", topNode)
				} else {
					desc = fmt.Sprintf("%s->%v", desc, topNode)
				}
			} else {
				p = topNode.Right
			}
		}
	}
	return
}

// BFSTraverse 广度优先遍历(BFS), 即层次遍历, 从根节点开始从左向右每一层遍历。
// 这里利用的队列，将根节点入列，当队列中元素大于0时，挨个出列，每出列一个元素，同时将该元素的左右节点依次入列，直到队列为空
func (tree *BinarySearchTree) BFSTraverse() (desc string) {
	treeList := list.New()
	treeList.PushBack(tree.root)
	for treeList.Len() > 0 {
		ele := treeList.Front()
		p, ok := ele.Value.(*BinarySearchNode)
		if !ok {
			break
		}
		if desc == "" {
			desc = fmt.Sprintf("%v", p)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p)
		}
		treeList.Remove(ele)

		if p.Left != nil {
			treeList.PushBack(p.Left)
		}
		if p.Right != nil {
			treeList.PushBack(p.Right)
		}
	}
	return
}

// DFSTraverse 深度优先遍历(DFS), 从根节点开始向下访问每个子节点，直到最后一个节点或者没有节点可以访问了为止，
// 然后在向上返回至最近一个仍然有子节点未被访问的节点的子节点开始访问。算法实现利用栈的特性，先根节点入栈，然后出栈(遍历)，然后依次入栈右子树和左子树，继续出栈。
func (tree *BinarySearchTree) DFSTraverse() (desc string) {
	s := &treeStack{List: list.New()}
	s.PushBack(tree.root)

	for s.Len() > 0 {
		ele := s.Back()
		p, ok := ele.Value.(*BinarySearchNode)
		if !ok {
			break
		}
		if desc == "" {
			desc = fmt.Sprintf("%v", p)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p)
		}

		s.Remove(ele)
		if p.Right != nil {
			s.PushBack(p.Right)
		} else {
			s.PushBack(p.Left)
		}
	}
	return
}
