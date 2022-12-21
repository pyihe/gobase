package tree

import (
	"container/list"
	"fmt"
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
	Data() Element      // 节点存储的数据
	Root() Node         // 返回根节点
	LeftChild() Node    // 左孩子
	RightChild() Node   // 右孩子
	LeftSibling() Node  // 左兄弟
	RightSibling() Node // 右兄弟
	Parent() Node       // 父节点
	Depth() int         // 返回自己所处的深度
	Color() Color       // 返回节点颜色
}

// Tree 树
type Tree interface {
	// Root 返回树的根节点
	Root() Node
	// Depth 返回树的深度
	Depth() int
	// Insert 插入新节点
	Insert(Element) Node
	// Remove 移除节点
	Remove(Element) bool
	// Find 查找节点
	Find(Element) Node
	// Update 更新节点
	Update(Element, Element) bool
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

/**********************************************************************************************************************/

// PreOrderTraverse 前序遍历:以当前节点为根节点，根——>左——>右
func PreOrderTraverse(root Node) (desc string) {
	s := &treeStack{
		List: list.New(),
	}
	p := root
	for p != nil || s.Len() > 0 {
		if p != nil {
			s.push(p)
			if desc == "" {
				desc = fmt.Sprintf("%v", p)
			} else {
				desc = fmt.Sprintf("%s->%v", desc, p)
			}
			p = p.LeftChild()
		} else {
			p = s.pop().(Node).RightChild()
		}
	}
	return
}

// InOrderTraverse 中序遍历:以当前节点为根节点，左——>根——>右
func InOrderTraverse(root Node) (desc string) {
	s := &treeStack{List: list.New()}
	p := root
	for p != nil || s.Len() > 0 {
		if p != nil {
			s.PushBack(p)
			p = p.LeftChild()
		} else {
			ele := s.Back()
			s.Remove(ele)
			p = ele.Value.(Node)
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
	s := &treeStack{List: list.New()}
	p := root

	var (
		topNode  Node
		lastNode Node
	)

	for p != nil || s.Len() > 0 {
		if p != nil {
			s.PushBack(p)
			p = p.LeftChild()
		} else {
			ele := s.Back().Value
			topNode = ele.(Node)
			if topNode.RightChild() == nil || topNode.RightChild() == lastNode {
				s.Remove(s.Back())
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
	treeList := list.New()
	treeList.PushBack(root)
	for treeList.Len() > 0 {
		ele := treeList.Front()
		p, ok := ele.Value.(Node)
		if !ok {
			break
		}
		if desc == "" {
			desc = fmt.Sprintf("%v", p)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p)
		}
		treeList.Remove(ele)

		if leftChild := p.LeftChild(); leftChild != nil {
			treeList.PushBack(leftChild)
		}
		if rightChild := p.RightChild(); rightChild != nil {
			treeList.PushBack(rightChild)
		}
	}
	return
}

// DFSTraverse 深度优先遍历(DFS), 从根节点开始向下访问每个子节点，直到最后一个节点或者没有节点可以访问了为止，
// 然后在向上返回至最近一个仍然有子节点未被访问的节点的子节点开始访问。算法实现利用栈的特性，先根节点入栈，然后出栈(遍历)，然后依次入栈右子树和左子树，继续出栈。
func DFSTraverse(root Node) (desc string) {
	s := &treeStack{List: list.New()}
	s.PushBack(root)

	for s.Len() > 0 {
		ele := s.Back()
		p, ok := ele.Value.(Node)
		if !ok {
			break
		}
		if desc == "" {
			desc = fmt.Sprintf("%v", p)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p)
		}

		s.Remove(ele)

		if rightChild := p.RightChild(); rightChild != nil {
			s.PushBack(rightChild)
		} else if leftChild := p.LeftChild(); leftChild != nil {
			s.PushBack(leftChild)
		}
	}
	return
}
