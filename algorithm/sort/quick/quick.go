package quick

import "sort"

// Sort 快速排序
// 稳定性：NO
func Sort(data sort.Interface) {
	sort.Sort(data)
}
