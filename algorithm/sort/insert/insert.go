package insert

// Sort 插入排序:从第一个元素开始，该元素可以认为已经被排序，取出下一个元素，
//
//	在已经排序的元素序列中从后向前扫描如果该元素（已排序）大于新元素，
//	将该元素移到下一位置，重复步骤3，直到找到已排序的元素小于或者等于新
//	元素的位置，将新元素插入到下一位置中，重复步骤2
//
// 平均时间复杂度：O(n^2)
// 最优时间复杂度：O(n)
// 最坏时间复杂度：O(n^2)
// 空间复杂度: O(1)
// 稳定性：YES
func Sort(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		key, j := data[i], i-1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
