package basics

import "github.com/rs/zerolog/log"

/**
 * copy(dst, src) 切片拷贝时候 如果dst为nil则不能拷贝任何元素  如果dst不为空 则最多拷贝dst cap大小的元素到dst数组
 */

// 初始化一个slice 指定len、cap
// 一个slice进行append时 如果原始cap小于1024 cap=old*2 反之按照1/4倍进行扩展
func gSlice() {
	s := make([]int, 2, 5)
	s[0] = 1
	log.Info().Interface("len", len(s)).Interface("cap", cap(s)).Interface("object", s).Send()
}

// 当新切片加上以后超出老切片容量会新建一个数组
// 当新切片加上以后未超出老切片的容量
// 对子切片的操作会影响原切片或原数组的变更
func gSliceAppend() {
	s := make([]int, 2, 3)
	s = append(s, []int{1, 2, 3, 4, 5}...)
	m := s[:3]
	m[0] = 1
	log.Info().Interface("len", len(s)).Interface("cap", cap(s)).Interface("data", s).Send()
}

// 通过数组或切片获取子切片 子切片的len为j-i cap为k-i
func gSliceCut() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	m := a[3:5]
	log.Info().Interface("len", (5-3) == len(m)).Interface("cap", cap(a)-3 == cap(m)).Send()
}

func gSliceCopy() {
	a := make([]int, 4)
	b := []int{1, 2, 3, 4}
	n := copy(a, b)
	log.Info().Interface("len", n).Interface("a", a).Interface("b", b).Send()
}

type S struct {
	Name string
}

// range 做了拷贝和原数组不是指向的同一地址不能直接修改  需要使用索引的形式修改
func gSliceCite(slice []S) {
	for index := range slice {
		slice[index].Name = "kimo"
	}
}
