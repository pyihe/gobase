package kmp

// SimpleSearch 朴素的模式匹配算法
// s 被匹配的串(非空)
// t 需要匹配的串(非空)
// pos 开始匹配的位置
func SimpleSearch(s, t string, pos int) (index int) {
	n := len(t)
	switch {
	case pos < 0 || pos >= len(s)-n || n > len(s):
		return -1
	case n == 0:
		return 0
	case n == 1:
		if i := indexByte(s[pos:], t[0]); i != -1 {
			return pos + i
		} else {
			return i
		}
	case n == len(s):
		if s == t {
			return 0
		}
		return -1
	}
	i, j := pos, 0
	for i < len(s) && j < n {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j >= n {
		index = i - n
	} else {
		index = -1
	}
	return
}

func indexByte(s string, b byte) int {
	for i, e := range s {
		if rune(b) == e {
			return i
		}
	}
	return -1
}

// Search KMP字符串搜索算法
// s 被搜索字符串
// t 需要搜索的字符串
// pos 开始搜索的位置
func Search(s, t string, pos int) (index int) {
	n := len(t)
	switch {
	case pos < 0 || pos >= len(s)-n || n > len(s):
		return -1
	case n == 0:
		return 0
	case n == 1:
		if i := indexByte(s[pos:], t[0]); i != -1 {
			return pos + i
		} else {
			return i
		}
	case n == len(s):
		if s == t {
			return 0
		}
		return -1
	}

	i, j := pos, 0
	next := nextVal(t)
	for i < len(s) && j < len(t) {
		switch {
		case s[i] == t[j]:
			i++
			j++
		default:
			if j == 0 {
				i++
			} else {
				j = next[j-1]
			}
		}
	}

	if j == len(t) {
		index = i - len(t)
	} else {
		index = -1
	}
	return
}

func nextVal(t string) []int {
	i, j, next := 1, 0, make([]int, len(t))
	for i < len(t) {
		switch {
		case t[i] == t[j]:
			j++
			next[i] = j
			i++
		default:
			if j == 0 {
				next[i] = 0
				i++
			} else {
				j = next[j-1]
			}
		}
	}
	return next
}
