package merge

// Sort 归并排序
// 稳定性：YES
func Sort(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}
	// 分
	left := data[:n/2]
	lSize := len(left)
	if lSize > 1 {
		left = Sort(left)
	}

	right := data[n/2:]
	rSize := len(right)
	if rSize > 1 {
		right = Sort(right)
	}
	// 治
	i, j, result := 0, 0, make([]int, 0, n)
	for i < lSize || j < rSize {
		switch {
		case i < lSize && j < rSize:
			if left[i] <= right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		case i < lSize:
			result = append(result, left[i:]...)
			i += lSize - i
		case j < rSize:
			result = append(result, right[j:]...)
			j += rSize - j
		}
	}
	return result
}
