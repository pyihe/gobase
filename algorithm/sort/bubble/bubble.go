package bubble

import "sort"

// Sort 冒泡排序，每次找出最小的值放在最前面
// 平均时间复杂度：O(n^2)
// 最坏时间复杂度：O(n^2)
// 最优时间复杂度：O(n)
// 空间复杂度：O(1)
// 稳定性：YES
func Sort(data sort.Interface) {
	// isChange 是否有发生了交换动作, 如果没有发生交换，证明数据已经有序了, 不需要在进行遍历比较
	isChange, n := false, data.Len()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}
