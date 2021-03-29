package alg

func quickSort(nums []int) {

}

// 冒泡排序
func bubblingSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		// 每过一遍把最大的数放到最前面
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func InsertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for ; j >= 0; j-- {
			// 如果前一个比后一个大则换位置
			if arr[j] > arr[i] {
				arr[j] = arr[i]
			}
		}
		arr[j] = arr[i]
	}
}
