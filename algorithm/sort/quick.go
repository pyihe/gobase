package sort

/*
	快速排序: 算法实际上是使用分治法策略进行排序，原理是将数组分成较大和较小的两个序列，然后递归地排序两个序列
	1. 从数组中选出一个基准元素(pivot)
	2. 将比基准小的元素放在基准左边，反之放在右边，如果相同，则左右都可以。
	3. 然后递归地对基准值左右的两个序列进行排序

	基准值的选择对排序性能有决定性影响。
*/

//获取基准位置
func getPivot(src []int, start, end int) int {
	if len(src) <= end || start < 0 {
		panic("invalid start or end value.")
	}
	index := start - 1 //用于存放最后返回的pivot
	pivot := end       //选取最后一个元素为pivot

	for i := start; i < end; i++ {
		if src[i] <= src[pivot] { //将所有比基准值小的移动到前面，这样大的都在后面
			index++
			src[index], src[i] = src[i], src[index]
		}
	}
	src[index+1], src[end] = src[end], src[index+1]
	return index + 1
}

func quickSort(data []int, start, end int) {
	if start >= end {
		return
	}
	//获取基准位置
	pivot := getPivot(data, start, end)
	//递归对基准位置左右两边的元素排序
	quickSort(data, start, pivot-1)
	quickSort(data, pivot+1, end)
}

//原地分割版本
func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

//快速排序(直接采用中间元素作为基准值，此方法需要额外的存储空间，相对而言，在空间复杂度上不可取)
func QuickSort2(data []int) []int {
	count := len(data)
	if count <= 0 {
		return nil
	}

	//选择中间的数作为参考
	keyIndex := count / 2
	key := data[keyIndex]

	//分成左右两部分，左边放比key小的值，右边放比key大的值
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < count; i++ {
		if i == keyIndex {
			continue
		}
		if data[i] < key {
			left = append(left, data[i])
		} else {
			right = append(right, data[i])
		}
	}

	left = QuickSort2(left)
	right = QuickSort2(right)

	//最后将得到的两组数组合起来
	var result []int
	result = append(result, left...)
	result = append(result, key)
	result = append(result, right...)
	return result
}
