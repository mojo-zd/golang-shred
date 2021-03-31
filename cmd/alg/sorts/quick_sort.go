package sorts

// 快速排序
// 1.选定中心轴
// 2.将大于中心轴的数字放在右边
// 3.将小于中心轴的数字放在左边
// 4.分别将左右子序重复前三步操作
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 选定基础数据进行比较
	split := arr[0]
	// 把数据分为三段
	left := make([]int, 0, 0)
	right := make([]int, 0, 0)
	mid := []int{split}
	for i := 1; i < len(arr); i++ {
		if arr[i] < split {
			left = append(left, arr[i])
		} else if arr[i] > split {
			right = append(right, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	left, right = QuickSort(left), QuickSort(right)
	return append(append(left, mid...), right...)
}
