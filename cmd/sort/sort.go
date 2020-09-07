package main

import "log"

// 对以下数据进行排序操作

func main() {
	arr := []int{8, 2, 4, 19, 30, 3}
	log.Println(bubblingSort(arr))
}

func bubblingSort(sort []int) []int {
	for i := 0; i < len(sort)-1; i++ {
		for j := i + 1; j < len(sort); j++ {
			if sort[i] > sort[j] {
				si := sort[i]
				sort[i] = sort[j]
				sort[j] = si
			}
		}
	}
	return sort
}
