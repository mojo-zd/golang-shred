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

	for i := 0; i < len(arr)-1; i++ {
		var j int
		value := arr[i+1]
		for j = i + 1; j >= 0 && value < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = value
	}
}
