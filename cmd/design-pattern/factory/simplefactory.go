package factory

import "fmt"

type Car interface {
	Running()
}

type Audi struct {
}

func (*Audi) Running() {
	fmt.Println("奥迪跑起来了")
}

type Byd struct {
}

func (*Byd) Running() {
	fmt.Println("比亚迪跑起来了")
}

type simpleCarFactory struct {
}

// 方法一： 静态方法
func NewSimpleFactory() *simpleCarFactory {
	return &simpleCarFactory{}
}

func (*simpleCarFactory) CreateAudiCar() Car {
	return &Audi{}
}

func (*simpleCarFactory) CreateBydCar() Car {
	return &Byd{}
}

// 方法二: 根据传入type 返回对应的Car实现
func NewSimpleCarFactory(typ string) Car {
	if typ == "奥迪" {
		return &Audi{}
	} else if typ == "比亚迪" {
		return &Byd{}
	}
	return nil
}
