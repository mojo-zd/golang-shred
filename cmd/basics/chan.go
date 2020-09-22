package basics

import (
	"fmt"
	"time"
)

func service1(ch chan string) {
	ch <- "service1"
}

func service2(ch chan string) {
	ch <- "service2"
}

func selectChan() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go service1(ch1)
	go service2(ch2)

	select {
	case service := <-ch1:
		fmt.Println("from", service)
	case service := <-ch2:
		fmt.Println("from", service)
	case <-time.After(time.Second * 2): // 添加超时机制
		fmt.Println("no case execute")
	}
}

func chanBuffer() {
	c := make(chan string, 3)
	c <- "a"
	c <- "b"
	fmt.Println("len", len(c), ", cap", cap(c))
	c <- "c"
	fmt.Println("start block")
	c <- "d"
	fmt.Println("blocked")
}

// 只要chan中有数据 就算close了chan也能从中获取到数据
func chanClose() {
	c := make(chan string, 3)
	c <- "a"
	c <- "b"
	close(c)
	for {
		if v, ok := <-c; ok {
			fmt.Println("receive", v)
		} else {
			fmt.Println("break for loop")
			break
		}
	}
}

// 单向chan(接受)
func chanSingleRec(c chan<- int) {
	c <- 1
}

// 单向chan(发送)
func chanSingleSend(c <-chan int) {
	<-c
}
