package basics

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

// 什么情况下两个指针类型是相等的
// 在两个指针变量指向同一个地址的时候或者都为nil的时候
func pointer() {
	i := 10
	var p1, p2 *int
	log.Info().Interface("no init", p1 == p2).Send()
	p1 = &i
	p2 = &i
	log.Info().Interface("equal", p1 == p2).Send()
}

// new 出来的对象为指针类型
func gNew() {
	a := new(int)
	fmt.Println("打印值", *a)
	fmt.Println("打印地址", a)
}

// make出来的为非指针
func gMake() {
	a := make([]int, 0)
	fmt.Println("a", &a)
}
