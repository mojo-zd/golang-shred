package factory

import "fmt"

type Seat interface {
	Massage()
}

type highSeat struct {
}

func (*highSeat) Massage() {
	fmt.Println("自动按摩功能")
}

type lowSeat struct {
}

func (*lowSeat) Massage() {
	fmt.Println("不带按摩功能")
}

type Tyre interface {
	Wear() // 轮胎磨损
}

type highTyre struct {
}

func (*highTyre) Wear() {
	fmt.Println("磨损消耗较小")
}

type lowTyre struct {
}

func (*lowTyre) Wear() {
	fmt.Println("磨损消耗较大")
}

type CarFactory interface {
	CreateSeat() Seat
	CreateTyre() Tyre
}

type highFactory struct {
}

func NewHighFactory() CarFactory {
	return &highFactory{}
}

func (*highFactory) CreateSeat() Seat {
	return &highSeat{}
}

func (*highFactory) CreateTyre() Tyre {
	return &highTyre{}
}

type lowFactory struct {
}

func NewLowFactory() CarFactory {
	return &lowFactory{}
}
func (*lowFactory) CreateSeat() Seat {
	return &lowSeat{}
}

func (*lowFactory) CreateTyre() Tyre {
	return &lowTyre{}
}
