package sort

/*
	基数排序: 排序对象为正整数
	基数排序用于比较整数形式的数据。其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较。

	将所有待排序数值（正整数）统一为同样的数位长度，数位较短的数前面补零。然后，从最低位开始，依次进行一次排序。
	这样从最低位排序一直到最高位排序完成以后，数列就变成一个有序序列

	三种利用桶机制进行排序的排序方法:
		1. 基数排序：根据键值的每位数字来分配桶；
		2. 计数排序：每个桶只存储单一键值；
		3. 桶排序：每个桶存储一定范围的数值；

	利用桶排序的原理，依次对每个元素对每一位进行桶排序
*/

func Radix(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}
	// 找出原始数据中的最大值的位数，这里是10进制
	max, _ := getMaxMin(data)
	var maxBitCnt int
	for max > 0 {
		maxBitCnt++
		max /= 10
	}

	var bucket [][]int               // 0-9对应每一位上的数值
	mod := 10                        // 用于获取每个元素的每一位的值
	radix := 1                       // 配合mod用于获取桶的索引
	for i := 0; i < maxBitCnt; i++ { // 操作每一位，个十百千...
		// 每次操作前需要将桶清空，否则桶中还包含上一次处理的结果
		bucket = make([][]int, 10)
		for j := 0; j < length; j++ { // 获取每个元素的每一位
			bucketIndex := (data[j] % mod) / radix
			bucket[bucketIndex] = append(bucket[bucketIndex], data[j])
		}
		pos := 0
		for j := 0; j < len(bucket); j++ {
			if len(bucket[j]) > 0 {
				// 对桶内的元素进行排序
				for k := 0; k < len(bucket[j]); k++ {
					data[pos] = bucket[j][k]
					pos++
				}
			}
		}
		radix *= 10
		mod *= 10
	}
}
