package sort

/*
	冒泡排序, 算法步骤(升序):
	1. 每次比较相邻的两个元素，如果第一个比第二个大，则交换位置
	2. 从头到尾比较完称之为一趟，每趟比较完成后，该趟最大的元素都将被移动到所有参与比较的元素最后的位置，所以下一趟比较的时候，已经放在最后位置的元素不需要再参与比较。
	3. 继续对变少的元素进行下一趟的比较，直到所有元素有序。
*/

func Bubble(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	// 设置一个标志，如果内层循环没有数据交换，证明数据已经有序，不需要再进行遍历
	swap := true
	for i := 0; i < n && swap; i++ {
		swap = false
		for j := n - 2; j >= i; j-- {
			if data[j] > data[j+1] {
				data[j+1], data[j] = data[j], data[j+1]
				swap = true
			}
		}
	}
}
