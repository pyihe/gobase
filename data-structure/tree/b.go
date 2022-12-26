package tree

import (
	"fmt"
	"math"
	"sort"
)

type bTreeDataList []*bTreeData

func (l bTreeDataList) String() (desc string) {
	for _, data := range l {
		if desc == "" {
			desc = fmt.Sprintf("(%d, %v)", data.key, data.value)
		} else {
			desc = fmt.Sprintf("%s->(%d, %v)", desc, data.key, data.value)
		}
	}
	return
}

func (l bTreeDataList) Len() int {
	return len(l)
}

func (l bTreeDataList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l bTreeDataList) Less(i, j int) bool {
	return l[i].key < l[i].key
}

// B-Tree节点存储的数据
// 其中key用于大小比较，data为实际的数据信息
type bTreeData struct {
	key   int
	value interface{}
}

func newBTreeData(key int, value interface{}) *bTreeData {
	return &bTreeData{
		key:   key,
		value: value,
	}
}

/**********************************************************************************************************************/

type bNode struct {
	parent   *bNode
	data     []*bTreeData
	children []*bNode
}

func newBNode(data []*bTreeData, parent *bNode) *bNode {
	return &bNode{
		parent:   parent,
		data:     data,
		children: nil,
	}
}

func (node *bNode) String() (desc string) {
	if node == nil {
		return "<nil>"
	}
	// 先打印自己的数据
	desc = fmt.Sprintf("%v", node.data)
	// 然后打印孩子节点的信息
	for _, child := range node.children {
		if desc == "" {
			desc = fmt.Sprintf("%v", child.String())
		} else {
			desc = fmt.Sprintf("%s->%v", desc, child.String())
		}
	}
	return
}

// 节点存储的数据达到上限，需要分裂
func (node *bNode) split(m int) (root *bNode) {
	dataLen := len(node.data)
	if dataLen < m {
		root = node
		return
	}
	mid := 0
	if dataLen%2 != 0 {
		mid = dataLen / 2
	} else {
		mid = dataLen/2 - 1
	}

	// promoteData 为需要晋升到上一层的数据
	promoteData, parent := node.data[mid], node.parent

	// 分裂操作步骤为:
	// 1. 将promoteKey从当前节点中剔除，KeyData分成大小两部分，生成两个新的节点，并重新分配孩子结点
	smallerKey := node.data[:mid]
	biggerKey := node.data[mid+1:]
	smallerNode, biggerNode := newBNode(smallerKey, parent), newBNode(biggerKey, parent)
	// 重新分配孩子节点，以promoteKey为中心，小的放在smallerNode里面，其他的放在biggerNode里面，并绑定父子结点关系
	for i := range node.children {
		child := node.children[i]
		switch {
		case i <= mid:
			child.parent = smallerNode
			smallerNode.children = append(smallerNode.children, child)
		default:
			child.parent = biggerNode
			biggerNode.children = append(biggerNode.children, child)
		}
	}
	// 如果父节点为空, 则需要新生成一个父节点，并将promoteKey晋升到该父节点中去
	if parent == nil {
		parent = newBNode([]*bTreeData{promoteData}, nil)
		parent.children = append(parent.children, smallerNode, biggerNode)
		smallerNode.parent = parent
		biggerNode.parent = parent
	} else {
		parent.removeChild(node)
		parent.addData(promoteData)
		parent.insertChild(smallerNode)
		parent.insertChild(biggerNode)
	}

	root = parent.split(m)
	return
}

func (node *bNode) insertChild(child *bNode) {
	for i, c := range node.children {
		if c.data[0].key > child.data[len(child.data)-1].key {
			childList := make([]*bNode, len(node.children)+1)
			copy(childList[:i], node.children[:i])
			childList[i] = child
			copy(childList[i+1:], node.children[i:])
			node.children = childList
			return
		}
	}
	node.children = append(node.children, child)
}

func (node *bNode) removeChild(child *bNode) {
	for i, v := range node.children {
		if v == child {
			copy(node.children[i:], node.children[i+1:])
			node.children[len(node.children)-1] = nil
			node.children = node.children[:len(node.children)-1]
			break
		}
	}
}

// 返回node的左右兄弟节点（相邻的左右两个节点）
func (node *bNode) siblingNode() []*bNode {
	// 没有父节点也就没有兄弟节点
	// 或者父节点的孩子节点只有一个
	if node.parent == nil || len(node.parent.children) <= 1 {
		return nil
	}
	siblingNodes := make([]*bNode, 0, 2)
	for i, child := range node.parent.children {
		if child != node {
			continue
		}
		if i > 0 {
			siblingNodes = append(siblingNodes, node.parent.children[i-1])
		}
		if i < len(node.parent.children)-1 {
			siblingNodes = append(siblingNodes, node.parent.children[i+1])
		}
	}
	return siblingNodes
}

// 获取节点中某个关键字的左右子树
func (node *bNode) getChildren(key int) []*bNode {
	if len(node.children) == 0 {
		return nil
	}

	children := make([]*bNode, 0, 2)
	for i, v := range node.data {
		if v.key != key {
			continue
		}
		switch {
		case i == 0:
			children = append(children, node.children[0], node.children[1])
		case i == len(node.data)-1:
			n := len(node.children)
			children = append(children, node.children[n-1], node.children[n-2])
		default:
			children = append(children, node.children[i], node.children[i+1])
		}
		break
	}
	return children
}

// 获取node在父节点中对应的分界数据
func (node *bNode) getParentKey() *bTreeData {
	if node.parent == nil {
		return nil
	}
	dataLen := len(node.parent.data)
	if dataLen == 0 {
		return nil
	}
	for _, data := range node.parent.data {
		if data.key > node.data[len(node.data)-1].key {
			return data
		}
	}
	return node.parent.data[dataLen-1]
}

func (node *bNode) getData(key int) *bTreeData {
	for _, v := range node.data {
		if v.key == key {
			return v
		}
	}
	return nil
}

func (node *bNode) deleteData(key int) {
	n := len(node.data)
	for i, v := range node.data {
		if v.key == key {
			copy(node.data[i:], node.data[i+1:])
			node.data[n-1] = nil
			node.data = node.data[:n-1]
			break
		}
	}
}

func (node *bNode) addData(data *bTreeData) {
	n := len(node.data)
	switch {
	case n == 0: // 如果节点还没有数据
		node.data = []*bTreeData{data}

	case data.key < node.data[0].key: // 如果插入的数据最小
		dataList := make([]*bTreeData, n+1)
		dataList[0] = data
		copy(dataList[1:], node.data)
		node.data = dataList

	case data.key > node.data[n-1].key: // 如果插入的数据最大
		node.data = append(node.data, data)

	default:
		for i := 0; i < n; i++ {
			src := node.data[i]
			if src.key == data.key {
				src.value = data.value
				return
			}

			if src.key > data.key {
				dataList := make([]*bTreeData, n+1)
				copy(dataList[:i], node.data[:i])
				dataList[i] = data
				copy(dataList[i+1:], node.data[i:])
				node.data = dataList
				return
			}
		}
	}
}

func (node *bNode) remove(m int, key int) bool {
	/*
		目标结点为关键字targetKey所在的结点
		相邻关键字：对于不在终端结点上的关键字，其相邻关键字为该关键字左子树中最大的关键字或者右子树中最小的关键字

		删除操作总体分两种：
		1. targetKey在叶子结点上
			a) 目标结点内的关键字数量大于math.Ceil(m/2)-1，这时删除不会破坏B树的性质，可以直接删除
			b) 目标结点内的关键字数量等于math.Ceil(m/2)-1，并且其左右兄弟结点中存在关键字数量大于math.Ceil(m/2)-1的结点，
			   则删除后需要向兄弟结点借关键字(将兄弟结点中的某个关键字提升到父结点，将父结点中的某个关键字下沉到目标结点)
			c) 目标结点内的关键字数量等于math.Ceil(m/2)-1，而兄弟结点中不存在关键字数量大于math.Ceil(m/2)-1的结点，则需要进行结点
			   合并从父结点中取一个关键字与兄弟结点合并，并将取出的结点从父结点中删除，同时更新父子关系(如果需要)

		2. targetKey不在叶子结点上
			a) targetKey存在关键字数量大于math.Ceil(m/2)-1结点的左子树或者右子树，在对应子树上找到该关键字的相邻关键字，并交换相邻关键字与目标
			   关键字，然后在替换后的位置上删除targetKey（此时已经转换成了在叶子结点上删除关键字）
			b) targetKey左右子树的关键字数量均等于math.Ceil(m/2)-1，则将这两个左右子树结点进行合并，然后删除targetKey，并调整关系
	*/
	targetNode := node.find(key)
	if targetNode == nil {
		return false
	}
	ceil := int(math.Ceil(float64(m)/float64(2))) - 1
	switch {
	case len(targetNode.children) == 0: // 如果targetKey在叶子结点中，则直接移除
		return targetNode.removeAtLeaf(ceil, key)
	default: // 否则中间删除
		return targetNode.removeAtMid(ceil, key)
	}
}

func (node *bNode) removeAtMid(ceil int, key int) bool {
	data := node.getData(key)
	if data == nil {
		return false
	}

	children := node.getChildren(key)
	if len(children) == 0 {
		return false
	}

	for _, c := range children {
		if len(c.data) <= ceil {
			continue
		}
		// 如果是右子树
		if c.data[0].key > key {
			minNode := getMinNode(node).(*bNode)
			minData := minNode.data[0]
			node.deleteData(key)
			node.addData(minData)
			minNode.deleteData(minData.key)
			minNode.addData(data)
			return minNode.removeAtLeaf(ceil, key)
		}
		// 如果是左子树
		if c.data[len(c.data)-1].key < key {
			maxNode := getMaxNode(node).(*bNode)
			maxData := maxNode.data[len(maxNode.data)-1]
			node.deleteData(key)
			node.addData(maxData)
			maxNode.deleteData(maxData.key)
			maxNode.addData(data)
			return maxNode.removeAtLeaf(ceil, key)
		}
	}
	node.deleteData(key)
	var (
		newNode  = newBNode(nil, node)
		dataList []*bTreeData
	)

	for _, child := range children {
		dataList = append(dataList, child.data...)
		for _, grandChild := range child.children {
			newNode.insertChild(grandChild)
		}
		node.removeChild(child)
	}
	sort.Sort(bTreeDataList(dataList))

	newNode.data = dataList
	node.insertChild(newNode)
	return true
}

func (node *bNode) removeAtLeaf(ceil int, key int) bool {
	dataLen := len(node.data)
	// 目标结点内的关键字数量大于math.Ceil(m/2)-1，这时删除不会破坏B树的性质，可以直接删除
	if dataLen > ceil {
		node.deleteData(key)
		return true
	}
	if dataLen != ceil {
		return false
	}

	// 目标结点内的关键字数量等于math.Ceil(m/2)-1，并且其左右兄弟结点中存在关键字数量大于math.Ceil(m/2)-1的结点，
	// 则删除后需要向兄弟结点借关键字(将兄弟结点中的某个关键字提升到父结点，将父结点中的某个关键字下沉到目标结点)
	siblingNodes := node.siblingNode()
	if len(siblingNodes) == 0 {
		return false
	}

	parentData := node.getParentKey()
	if parentData == nil {
		return false
	}

	for _, sibling := range siblingNodes {
		if len(sibling.data) <= ceil {
			continue
		}

		var data *bTreeData

		// 如果存在关键字数量大于ceil的兄弟结点
		// 如果是parentKey的右子树
		if sibling.data[0].key > parentData.key {
			data = sibling.data[0]
		}

		// 如果是parentKey的左子树
		if sibling.data[len(sibling.data)-1].key < parentData.key {
			data = sibling.data[len(sibling.data)-1]
		}
		if data != nil {
			node.deleteData(key)                   // 删除目标关键值
			node.addData(parentData)               // 将父结点中的关键值放进目标结点
			node.parent.deleteData(parentData.key) // 从父结点中移除parentKey
			node.parent.addData(data)              // 从兄弟结点借关键值放进父结点
			sibling.deleteData(data.key)           // 从兄弟结点中移除关键值
			return true
		}
	}

	// 目标结点内的关键字数量等于math.Ceil(m/2)-1，而兄弟结点中不存在关键字数量大于math.Ceil(m/2)-1的结点，则需要进行结点
	// 合并: 从父结点中取一个关键字与兄弟结点合并，并将取出的结点从父结点中删除，同时更新父子关系(如果需要)
	node.deleteData(key)                     // 从当前结点中移除targetKe
	node.parent.deleteData(parentData.key)   // 从父结点中移除parentKey
	node.parent.removeChild(node)            // 从父结点中移除当前结点
	node.parent.removeChild(siblingNodes[0]) // 从父结点中移除被合并的兄弟结点

	// 利用parentKey、兄弟结点中的key以及本结点剩余的key重新组成一个结点
	newNode := newBNode([]*bTreeData{parentData}, node.parent)
	newNode.data = append(newNode.data, node.data...)
	newNode.data = append(newNode.data, siblingNodes[0].data...)
	sort.Sort(bTreeDataList(newNode.data))
	// 将新结点插入父结点中去
	node.parent.insertChild(newNode)
	return true
}

func (node *bNode) insert(m int, data *bTreeData) *bNode {
	switch {
	case len(node.children) == 0: // 如果node是叶子结点，则直接插入到data中，然后根据data的长度来判断是否需要分裂
		node.addData(data)
		return node.split(m)
	default: // 如果node不是叶子节点，则需要先找到插入的位置
		var leaf *bNode
		for i, v := range node.data {
			if data.key < v.key {
				leaf = node.children[i]
				break
			}
		}
		if leaf == nil {
			leaf = node.children[len(node.children)-1]
		}
		return leaf.insert(m, data)
	}
}

func (node *bNode) find(key int) *bNode {
	for i, v := range node.data {
		if v.key == key {
			return node
		}
		if v.key > key && len(node.children) == 0 {
			return node.children[i].find(key)
		}
	}
	// 如果node中所有key都比targetKey小，则在node的最后一个孩子结点中搜索，否则表示没有targetKey对应的结点
	n := len(node.children)
	if n == 0 {
		return nil
	}
	return node.children[n-1].find(key)
}

/**********************************************************************************************************************/

/*
	B树（B-tree）是一种自平衡的树，能够保持数据有序。这种数据结构能够让查找数据、顺序访问、插入数据及删除的动作，都在对数时间内完成。
	B树，概括来说是一个一般化的二叉查找树（binary search tree）一个节点可以拥有2个以上的子节点。与自平衡二叉查找树不同，B树适用于读写相
	对大的数据块的存储系统，例如磁盘。B树减少定位记录时所经历的中间过程，从而加快存取速度。B树这种数据结构可以用来描述外部存储。这种数据结构
	常被应用在数据库和文件系统的实现上。

	描述一颗B树时需要指定它的阶数，阶数表示了一个结点最多有多少个孩子结点，一般用字母m表示阶数。当m取2时，就是我们常见的二叉搜索树。
	一颗m阶的B树有如下性质：

	1. 每个结点最多有m个孩子结点（子树），最多有m-1个关键字。
	2. 如果根结点不是叶子结点，则根结点至少有2个子树，至少有一个关键字
	3. 除根结点外的所有非叶子结点至少有math.Ceil(m/2)个孩子结点，至少有math.Ceil(m/2)-1个关键字
	4. 所有叶子结点位于同一层
	5. 每个结点中的关键字都按照从小到大排序
*/

type BTree struct {
	m    int
	root *bNode
}

func NewBTree(m int) *BTree {
	return &BTree{
		m: m,
	}
}

func (bt *BTree) Insert(key int, value interface{}) {
	if bt == nil {
		return
	}
	data := newBTreeData(key, value)
	if bt.root == nil {
		bt.root = newBNode([]*bTreeData{data}, nil)
	}
	root := bt.root.insert(bt.m, data)
	for root.parent != nil {
		root = root.parent
	}
	bt.root = root
}

func (bt *BTree) Find(key int) (value interface{}, ok bool) {
	if bt == nil || bt.root == nil {
		return
	}
	node := bt.root.find(key)
	if node == nil {
		return
	}
	for _, data := range node.data {
		if data.key == key {
			ok = true
			value = data.value
			break
		}
	}
	return
}

func (bt *BTree) Remove(key int) bool {
	if bt == nil || bt.root == nil {
		return false
	}
	return bt.root.remove(bt.m, key)
}
