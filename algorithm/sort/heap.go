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
	maxHeap(data, count)

	for i := count - 1; i >= 0; i-- {
		data[0], data[i] = data[i], data[0]
		count--
		heap(data, 0, count)
	}
}

// 创建最大堆
func maxHeap(data []int, len int) {
	for i := len / 2; i >= 0; i-- {
		heap(data, i, len)
	}
}

func heap(data []int, index, len int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < len && data[left] > data[largest] {
		largest = left
	}
	if right < len && data[right] > data[largest] {
		largest = right
	}
	if largest != index {
		data[largest], data[index] = data[index], data[largest]
		heap(data, largest, len)
	}
}
