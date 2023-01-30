package sort

/*
	选择排序
	1. 在未排序的数组中找到最大或者最小的元素
	2. 将最大或者最小的元素放到已排序好的数组的末尾
	3. 重复步骤1、2
*/

// Select 选择排序:每次选择出未排序切片里最大或者最小的数放入已排好序的数组里
func Select(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}

	var minIndex int
	for i := 0; i < count-1; i++ {
		minIndex = i
		for j := i + 1; j < count; j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
	return
}
