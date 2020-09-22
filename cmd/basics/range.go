package basics

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

/**
 * 使用range遍历数组时 实际上是对原有数组元素做了拷贝 运行RangeYy可以得到结果
 */

// 验证range为副本拷贝 当需要使用value时 range中定义了一个变量存储当前value 当使用指针时每次循环都会覆盖这个变量 导致所有的指向都是同一个值
func Range(slice []int) map[int]*int {
	m := make(map[int]*int)
	for key, value := range slice {
		m[key] = &value
	}
	for i := 0; i < len(slice); i++ {
		m[len(slice)+i] = &slice[i]
	}
	return m
}

// 验证循环时增加元素  循环时并不会动态增加元素到数组  range时是对array做了数据拷贝
func RangAppend(slice []int) {
	for value := range slice {
		slice = append(slice, value)
	}
	log.Info().Interface("index", slice).Send()
}

func RangeYy(slice []int) {
	for i :=0; i<len(slice); i++ {
		fmt.Println("loop", &slice[i])
	}
	for _, value := range slice {
		fmt.Println("range", &value)
	}
}
