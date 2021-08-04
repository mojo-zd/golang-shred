package main

import "testing"

func TestCar(t *testing.T) {
	var car ICar = &Car{}
	car = &WaterCar{car}
	car = &FlyCar{car}
	car.Move()
}
