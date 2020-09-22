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
