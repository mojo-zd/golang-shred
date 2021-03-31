package search

// 二分法查找
// 通过mid把数据分为三段
// 1. target和mid进行比较, 如果相等则返回mid
// 2. 如果target大于mid, left向右移动 left = mid+1
// 3. 如果target小于mid, right向左移动 right = mid-1
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if target == arr[mid] {
			return mid
		}

		if target > arr[mid] {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return -1
}
