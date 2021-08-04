package main

import "fmt"

type ICar interface {
	Move()
}

type Car struct {
}

func (*Car) Move() {
	fmt.Println("car moving ...")
}

type FlyCar struct {
	ICar
}

func (f *FlyCar) Move() {
	f.ICar.Move()
	fmt.Println("fly car ...")
}

type WaterCar struct {
	ICar
}

func (w *WaterCar) Move() {
	w.ICar.Move()
	fmt.Println("water car ...")
}
