package sort

/*
	插入排序(升序)
	1. 每次从原始数组中找一个元素出来
	2. 将找出来的元素排序构建成一个有序的数组
	3. 在已经构建好的有序数组中找到当前元素需要插入的位置，并插入到有序数组中
	4. (场景与玩扑克牌摸牌时将摸的牌放进自己对手中形成有序组合相似)
*/

/*
 1. 从第一个元素开始，该元素可以认为已经被排序
 2. 取出剩余数组中的第一个元素
 3. 在已经排序的元素序列中从后向前扫描如果该元素（已排序）大于新元素，将该元素移到下一位置，直到找到已排序的元素小于或者等于新元素的位置，将新元素插入到该位置的后面
 3. 重复步骤2
*/

func Insert(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}

	var key, pos int
	for i := 1; i < count; i++ {
		key = data[i]
		pos = i

		// 此处循环为了将比key大的数往后移动
		for pos >= 1 && data[pos-1] > key {
			data[pos] = data[pos-1]
			pos--
		}
		data[pos] = key
	}

	return
}
