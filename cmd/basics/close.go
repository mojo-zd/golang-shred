package basics

import "fmt"

func closeFunc(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func closeInnerFunc() func() int {
	i := 1
	s := func() int {
		return i
	}
	i = 3
	return s
}

func closeFuncArr(arr []int) []func() int {
	var xx []func() int
	for i := 0; i < len(arr); i++ {
		xx = append(xx, func() int {
			fmt.Println(i)
			return i
		})
	}
	return xx
}

func closeFuncArrUp(arr []int) []func(m int) int {
	var xx []func(i int) int
	for i := 0; i < len(arr); i++ {
		xx = append(xx, func(x int) int {
			return x
		})
	}
	return xx
}

func myFunc()  func() int{
	foo := 0
	return func() int {
		foo ++
		return foo
	}
}

func extFunc(n int) func() {
	n++
	return func() {
		fmt.Println(n)
	}
}
