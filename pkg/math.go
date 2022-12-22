package pkg

// MaxInt 返回一组整型中的最大值
func MaxInt(is ...int) int {
	max := is[0]
	for i := 1; i < len(is); i++ {
		if is[i] > max {
			max = is[i]
		}
	}
	return max
}

// AbsInt 返回整型数值i的绝对值
func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
