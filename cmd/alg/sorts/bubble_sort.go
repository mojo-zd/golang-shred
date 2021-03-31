package sorts

// 冒泡排序1 每次循环把最大的沉到最后
// 把N个数进行N-1轮排序,升序排序中大的往下沉小的往上浮
func Bubble1Sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 冒泡排序2  每次循环把最大的数下沉到最后
func Bubble2Sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
