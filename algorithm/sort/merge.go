package sort

/*
	归并排序, 采用分支法实现， 有两种实现方法: 递归和迭代
	1. 递归的将数组平均分割成两半
	2. 在保持元素顺序的同时将上一步得到的子序列合并到一起
*/

// MergeByRecursion 递归实现归并排序
func MergeByRecursion(data []int) []int {
	count := len(data)
	if count <= 1 {
		return data
	}
	// 将data平均分为两部分
	left := data[:count/2]
	right := data[count/2:]
	left = MergeByRecursion(left)
	right = MergeByRecursion(right)

	// 递归完后将左右两部分归并
	var leftIndex, rightIndex int
	var result = make([]int, count)
	for i := 0; i < count; i++ {
		if leftIndex < len(left) && rightIndex < len(right) {
			if left[leftIndex] < right[rightIndex] {
				result[i] = left[leftIndex]
				leftIndex++
			} else {
				result[i] = right[rightIndex]
				rightIndex++
			}
		} else if leftIndex < len(left) {
			result[i] = left[leftIndex]
			leftIndex++
		} else if rightIndex < len(right) {
			result[i] = right[rightIndex]
			rightIndex++
		}
	}
	return result
}

// MergeByIter 迭代实现归并排序
func MergeByIter(data []int) []int {
	count := len(data)
	if count <= 1 {
		return data
	}
	middle := count / 2
	left := MergeByIter(data[:middle])
	right := MergeByIter(data[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	// 将左右两边按照大小合并到一起
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
