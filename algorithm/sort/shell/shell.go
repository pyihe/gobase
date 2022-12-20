package shell

import "sort"

// Sort 排序
// 稳定性：NO
func Sort(data sort.Interface) {
	gap, n := 1, data.Len()
	for gap < n/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < n; i++ {
			j := i
			for j >= gap && data.Less(j-gap, j) {
				data.Swap(j, j-gap)
				j = j - gap
			}
		}
		gap = gap / 3
	}
}
