package simpleselect

import "sort"

// Sort 选择排序:每次选择出未排序切片里最大或者最小的数放入已排好序的数组里
// 平均时间复杂度：O(n^2)
// 最坏时间复杂度：O(n^2)
// 最优时间复杂度：O(n^2)
// 空间复杂度: O(1)
// 稳定性: NO
func Sort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data.Less(minIndex, j) {
				minIndex = j
			}
		}
		if minIndex != i {
			data.Swap(i, minIndex)
		}
	}
}
