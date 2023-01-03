package sort

/*
	快速排序: 算法实际上是使用分治法策略进行排序，原理是将数组分成较大和较小的两个序列，然后递归地排序两个序列
	1. 从数组中选出一个基准元素(pivot)
	2. 将比基准小的元素放在基准左边，反之放在右边，如果相同，则左右都可以。
	3. 然后递归地对基准值左右的两个序列进行排序

	基准值的选择对排序性能有决定性影响。
*/

func Quick(array []int, p, r int) {
	if p < r {
		q := partition(array, 0, len(array)-1)
		Quick(array, p, q-1)
		Quick(array, q+1, r)
	}
}

func partition(array []int, p, r int) int {
	x := array[r]
	i := p - 1
	for j := p; j < r; j++ {
		if array[j] < x {
			i += 1
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[r] = array[r], array[i+1]
	return i + 1
}
