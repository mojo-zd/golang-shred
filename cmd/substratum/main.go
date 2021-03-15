package main

import (
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	energize()
	isStop()
	optimize()
	goroutine()
}

// for range slice 并把value赋值给新的map
// 观察为什么得不到期望的结果
// 原因: for range语法糖 循环之前做了值拷贝 每次循环都是对同一个变量赋值
// 矫正: 1.为v赋值临时变量 2.采用index方式赋值
func energize() {
	slice := []int{1, 2, 3, 4}
	m := map[int]*int{}
	for i, v := range slice {
		m[i] = &v
	}
	log.Info().Interface("map", m).Send()
}

// 下面的循环为停止吗
// 原因: for range循环之前对slice做了拷贝, 因此长度是固定的
func isStop() {
	slice := []int{1, 2, 3, 4}
	for i := range slice {
		slice = append(slice, i)
	}
	log.Info().Msg("stop")
}

// 思考: 采用for range方式循环大数组时是否需要对循环做优化
// 原因: 循环前拷贝的内容对内存是极大的浪费
// 矫正: 1.使用地址的方式进行遍历 2.对数组做切片引用
func optimize() {
	arr := [102400]int{1, 1, 1}
	for i, v := range arr {
		_ = i
		_ = v
	}
}

// 思考: map如何实现的无序性
// 原因: map的内部实现是采用的链式hash表
// https://github.com/golang/go/blob/0bd3853512ea0dcb252ce02113d3929db03d6aa6/src/runtime/map.go#L826

// 思考: 对map遍历时 新增的元素会遍历到吗
// 答案: 可能会 原因如上

// 思考: 下面的遍历中启动goroutine有问题吗
// 原因: 拷贝
// 矫正: 1.传值到goroutine 2.使用局部变量
func goroutine() {
	arr := []int{1, 2, 3, 4}
	for _, v := range arr {
		go func() {
			log.Info().Int("value", v).Send()
		}()
	}

	time.Sleep(time.Second)
}
