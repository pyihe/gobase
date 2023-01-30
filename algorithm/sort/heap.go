package sort

/*
	堆排序
	利用堆这种数据结构进行排序。堆是一个近似完全二叉树的结构，并同时满足堆的性质：即子节点的键值或索引总是小于（或者大于）它的父节点。
	大顶堆: 每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；
	小顶堆: 每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列；
	堆排序: 移除堆顶元素，然后递归做最大堆(或最小堆)运算

	1. 利用原始数据创建一个堆
	2. 做最大堆(或者最小堆)调整
	3. 移除堆顶元素，重复步骤2，直到堆的尺寸为1
*/

func Heap(data []int) {
	count := len(data)
	if count < 1 {
		return
	}
	// 构建最大堆
	// maxHeap(data, count)
	for i := 0; i < count; i++ {
		buildHeap(data[i:])
	}
}

func buildHeap(src []int) {
	// 表示n叉堆的数组中，叶子节点位于len(src)/n往后的元素中
	// 从第一个非叶子节点开始，比较节点和自己孩子的大小
	// 最大堆选出最大的值作为父节点，最小堆选择最小的值作为父节点
	n := len(src)
	// 二叉堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(src, i)
	}
}

func heapify(src []int, rootIndex int) {
	leftIndex, rightIndex, largestIndex := 2*rootIndex+1, 2*rootIndex+2, rootIndex
	if leftIndex < len(src) && src[leftIndex] > src[rootIndex] {
		largestIndex = leftIndex
	}
	if rightIndex < len(src) && src[rightIndex] > src[rootIndex] {
		largestIndex = rightIndex
	}
	if largestIndex != rootIndex {
		src[rootIndex], src[largestIndex] = src[largestIndex], src[rootIndex]
		heapify(src, rootIndex)
	}
}
