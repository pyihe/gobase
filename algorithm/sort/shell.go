package sort

/*
	希尔排序，是对插入排序更高效对一种改进。
	希尔排序基于插入排序的以下两点改进:
		1. 插入排序在对几乎已经排好序的数据操作时，效率高，即可以达到线性排序的效率
		2. 但插入排序一般来说是低效的，因为插入排序每次只能将数据移动一位

	希尔排序通过将比较的全部元素分为几个区域来提升插入排序的性能。这样可以让一个元素可以一次性地朝最终位置前进一大步。
	然后算法再取越来越小的步长进行排序，算法的最后一步就是普通的插入排序，但是到了这步，需排序的数据几乎是已排好的了（此时插入排序较快）。

	步长的选择是希尔排序的重要部分。只要最终步长为1任何步长序列都可以工作。算法最开始以一定的步长进行排序。然后会继续以一定步长进行排序，
	最终算法以步长为1进行排序。当步长为1时，算法变为普通插入排序，这就保证了数据一定会被排序。
*/

func Shell(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}
	// 最开始选取count/2为步长
	step := count / 2
	var key, pos int
	for step > 0 {
		for i := step; i < count; i++ {
			key = data[i]
			pos = i
			for pos >= step && data[pos-step] > key {
				data[pos] = data[pos-step]
				pos -= step
			}
			data[pos] = key
		}
		step /= 2
	}
}