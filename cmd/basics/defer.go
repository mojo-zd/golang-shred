package basics

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

/**
 *  defer类似于栈  先进后出   panic在defer之后进行
 */
func deferCall() {
	defer func() {
		log.Info().Interface("call", "打印前").Send()
	}()

	defer func() {
		log.Info().Interface("call", "打印中").Send()
	}()

	defer func() {
		log.Info().Interface("call", "打印后").Send()
	}()
	panic("触发异常")
}

type Person struct {
	age int
}

/**
 * 解释 https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466928&idx=1&sn=5d4f1c2c7802e3c7520b4e7c2fd05634&chksm=f247439fc530ca89da7b4446e6a5e3e87dc047d695ad1ad62916f2a7eccae244d7b4dfbca0de&scene=21#wechat_redirect
 */
func deferKnown(p *Person) {
	defer fmt.Println(p.age)
	defer func(person *Person) {
		fmt.Println(person.age)
	}(p)
	defer func() {
		fmt.Println(p.age)
	}()
	//p.age = 29
	p = &Person{29}
}

func deferRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
	}()

	go func() {
		panic("panic")
	}()
	panic("crash")
}
