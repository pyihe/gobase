package tree

import (
	"fmt"

	"github.com/pyihe/gobase/data-structure/list"
	"github.com/pyihe/gobase/data-structure/stack"
)

const (
	NoColor Color = iota // 没有颜色
	Red                  // 红黑树：红色
	Black                // 红黑树：黑色
)

// Color 节点颜色
type Color uint8

// Element 树节点存储的元素
type Element interface {
	Value() interface{}  // 元素值
	Compare(Element) int // 两个元素相比，大于返回>0, 小于返回<0, 等于返回=0
}

// Node 树节点
type Node interface {
	String() string     // String
	Depth() int         // 返回自己所处的深度, 深度从根节点到自己所经历的节点数量
	Height() int        // 返回自己所处的高度, 高度为从叶子节点到自己所经历的节点数量
	Data() Element      // 节点存储的数据
	Root() Node         // 返回根节点
	LeftChild() Node    // 左孩子
	RightChild() Node   // 右孩子
	LeftSibling() Node  // 左兄弟
	RightSibling() Node // 右兄弟
	Parent() Node       // 父节点
	Color() Color       // 返回节点颜色
}

// Tree 树
type Tree interface {
	// Root 返回树的根节点
	Root() Node
	// Depth 返回树的深度(高度)
	Depth() int
	// Insert 插入新节点
	Insert(Element)
	// Remove 移除节点
	Remove(Element) bool
	// Find 查找节点
	Find(Element) Node
	// Update 更新节点
	Update(Element, Element) bool
}

/**********************************************************************************************************************/

func getMinNode(node interface{}) (mNode interface{}) {
	switch node.(type) {
	case *avlNode:
		aNode := node.(*avlNode)
		for aNode.leftChild != nil {
			aNode = aNode.leftChild
		}
		mNode = aNode
	case *bstNode:
		bNode := node.(*bstNode)
		for bNode.leftChild != nil {
			bNode = bNode.leftChild
		}
		mNode = bNode
	}
	return
}

func getMaxNode(node interface{}) (mNode interface{}) {
	switch node.(type) {
	case *avlNode:
		aNode := node.(*avlNode)
		for aNode.rightChild != nil {
			aNode = aNode.rightChild
		}
		mNode = aNode
	case *bstNode:
		bNode := node.(*bstNode)
		for bNode.rightChild != nil {
			bNode = bNode.rightChild
		}
		mNode = bNode
	}
	return
}

/**********************************************************************************************************************/

// PreOrderTraverseRecursion 前序遍历(递归)
// 遍历顺序: 以当前节点为根节点，根——>左——>右
func PreOrderTraverseRecursion(root Node) (desc string) {
	if root == nil {
		return
	}
	desc = fmt.Sprintf("%v", root)
	if leftDesc := PreOrderTraverseRecursion(root.LeftChild()); leftDesc != "" {
		desc = fmt.Sprintf("%s->%s", desc, leftDesc)
	}
	if rightDesc := PreOrderTraverseRecursion(root.RightChild()); rightDesc != "" {
		desc = fmt.Sprintf("%s->%s", desc, rightDesc)
	}
	return
}

// InOrderTraverseRecursion 中序遍历
// 遍历顺序: 以当前节点为根节点，左->根->右
func InOrderTraverseRecursion(root Node) (desc string) {
	if root == nil {
		return
	}
	desc = InOrderTraverseRecursion(root.LeftChild())
	if midDesc := fmt.Sprintf("%v", root); midDesc != "" {
		if desc != "" {
			desc = fmt.Sprintf("%s->%s", desc, midDesc)
		} else {
			desc = midDesc
		}
	}
	if rightDesc := InOrderTraverseRecursion(root.RightChild()); rightDesc != "" {
		if desc != "" {
			desc = fmt.Sprintf("%s->%s", desc, rightDesc)
		} else {
			desc = rightDesc
		}
	}
	return
}

// PostOrderTraverseRecursion 后序遍历
// 遍历顺序: 以当前节点为父节点，左->右->根
func PostOrderTraverseRecursion(root Node) (desc string) {
	if root == nil {
		return
	}
	desc = PostOrderTraverseRecursion(root.LeftChild())
	if rightDesc := PostOrderTraverseRecursion(root.RightChild()); rightDesc != "" {
		if desc != "" {
			desc = fmt.Sprintf("%s->%s", desc, rightDesc)
		} else {
			desc = rightDesc
		}
	}
	if rootDesc := fmt.Sprintf("%v", root); rootDesc != "" {
		if desc != "" {
			desc = fmt.Sprintf("%s->%s", desc, rootDesc)
		} else {
			desc = rootDesc
		}
	}
	return
}

/**********************************************************************************************************************/

// PreOrderTraverse 前序遍历:以当前节点为根节点，根——>左——>右
func PreOrderTraverse(root Node) (desc string) {
	s := stack.NewLinkStack()
	p := root
	for p != nil || s.Len() > 0 {
		if p != nil {
			s.Push(p)
			if desc == "" {
				desc = fmt.Sprintf("%v", p)
			} else {
				desc = fmt.Sprintf("%s->%v", desc, p)
			}
			p = p.LeftChild()
		} else {
			v, _ := s.Pop()
			p = v.(Node).RightChild()
		}
	}
	return
}

// InOrderTraverse 中序遍历:以当前节点为根节点，左——>根——>右
func InOrderTraverse(root Node) (desc string) {
	s := list.NewDoubleLink()
	p := root
	for p != nil || s.Len() > 0 {
		if p != nil {
			s.Insert(s.Len(), p, 1)
			p = p.LeftChild()
		} else {
			ele := s.RemoveByLocate(s.Len() - 1)
			p = ele.Value().(Node)
			if desc == "" {
				desc = fmt.Sprintf("%v", p)
			} else {
				desc = fmt.Sprintf("%s->%v", desc, p)
			}
			p = p.RightChild()
		}
	}
	return
}

// PostOrderTraverse  后序遍历：以当前节点为根节点，左——>右——>根
func PostOrderTraverse(root Node) (desc string) {
	s := list.NewDoubleLink()
	p := root

	var (
		topNode  Node
		lastNode Node
	)

	for p != nil || s.Len() > 0 {
		if p != nil {
			s.Insert(s.Len(), p, 1)
			p = p.LeftChild()
		} else {
			ele := s.Get(s.Len() - 1)
			topNode = ele.Value().(Node)
			if topNode.RightChild() == nil || topNode.RightChild() == lastNode {
				s.Remove(ele)
				lastNode = topNode
				if desc == "" {
					desc = fmt.Sprintf("%v", topNode)
				} else {
					desc = fmt.Sprintf("%s->%v", desc, topNode)
				}
			} else {
				p = topNode.RightChild()
			}
		}
	}
	return
}

// BFSTraverse 广度优先遍历(BFS), 即层次遍历, 从根节点开始从左向右每一层遍历。
// 这里利用的队列，将根节点入列，当队列中元素大于0时，挨个出列，每出列一个元素，同时将该元素的左右节点依次入列，直到队列为空
func BFSTraverse(root Node) (desc string) {
	treeList := list.NewDoubleLink()
	treeList.Insert(treeList.Len(), root, 1)
	for treeList.Len() > 0 {
		ele := treeList.RemoveByLocate(treeList.Len() - 1)
		p, ok := ele.Value().(Node)
		if !ok {
			break
		}
		if desc == "" {
			desc = fmt.Sprintf("%v", p)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p)
		}

		if leftChild := p.LeftChild(); leftChild != nil {
			treeList.Insert(treeList.Len(), leftChild, 1)
		}
		if rightChild := p.RightChild(); rightChild != nil {
			treeList.Insert(treeList.Len(), rightChild, 1)
		}
	}
	return
}

// DFSTraverse 深度优先遍历(DFS), 从根节点开始向下访问每个子节点，直到最后一个节点或者没有节点可以访问了为止，
// 然后在向上返回至最近一个仍然有子节点未被访问的节点的子节点开始访问。算法实现利用栈的特性，先根节点入栈，然后出栈(遍历)，然后依次入栈右子树和左子树，继续出栈。
func DFSTraverse(root Node) (desc string) {
	s := list.NewDoubleLink()
	s.Insert(s.Len(), root, 1)

	for s.Len() > 0 {
		ele := s.RemoveByLocate(s.Len() - 1)
		p, ok := ele.Value().(Node)
		if !ok {
			break
		}
		if desc == "" {
			desc = fmt.Sprintf("%v", p)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p)
		}
		if rightChild := p.RightChild(); rightChild != nil {
			s.Insert(s.Len(), rightChild, 1)
		}
		if leftChild := p.LeftChild(); leftChild != nil {
			s.Insert(s.Len(), leftChild, 1)
		}
	}
	return
}
