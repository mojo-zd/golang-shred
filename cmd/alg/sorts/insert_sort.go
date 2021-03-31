package sorts

// 插入排序 边赋值边遍历
// 1. 每次取出i+1的位置一次和前一个数进行比较, 如果参考数据arr[i+1]大于arr[j]则进行位置交换
// 2. 通过1计算得到j的位置数据轮空,把arr[i+1]值赋给arr[j]
func InsertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		value := arr[i+1]
		var j int
		for j = i + 1; j >= 0 && value < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = value
	}
}
