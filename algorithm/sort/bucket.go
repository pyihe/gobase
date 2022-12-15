package sort

/*
	桶排序
	工作的原理是将数组分到有限数量的桶里。每个桶再个别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排序）

	桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。为了使桶排序更加高效，我们需要做到这两点：
		1. 在额外空间充足的情况下，尽量增大桶的数量
		2. 使用的映射函数能够将输入的 N 个数据均匀的分配到 K 个桶中
	同时，对于桶中元素的排序，选择何种比较排序算法对于性能的影响至关重要。

	1. 设置一个定量的数组当作空桶子。
	2. 寻访序列，并且把项目一个一个放到对应的桶子去。
	3. 对每个不是空的桶子进行排序。
	4. 从不是空的桶子里把项目再放回原来的序列中。
*/

func BucketSort(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}
	//首先初始化桶的数量
	//获取最大最小值
	max, min := getMaxMin(data)
	var factor = 10                      //决定桶数量大小的因子
	var bucketCnt = (max-min)/factor + 1 //保证最少一个桶,这里的10可以根据实际情况做调整

	//初始化桶
	var bucket = make([][]int, bucketCnt)
	//将每个元素放进对应的桶里
	for _, v := range data {
		i := getBucketIndex(v, min, factor)
		bucket[i] = append(bucket[i], v)
	}
	//然后对每个有元素的桶中的元素进行排序(排序方法任意，如插入排序), 同时将排好序的元素放回原数组中
	var index int
	for i := range bucket {
		src := bucket[i]
		if len(src) > 0 {
			InsertSort(src)
			for j := range src {
				data[index] = src[j]
				index++
			}
		}
	}
}

//这里可以看出目标值越大，对应桶的索引也越大，所以后面的桶存放的元素一定比前面的桶存放的元素大。
//将每个桶排好序后，直接合并每个桶的元素即为最终的排序结果
func getBucketIndex(target, min, step int) int {
	return (target - min) / step
}

func getMaxMin(data []int) (max int, min int) {
	max, min = data[0], data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}
