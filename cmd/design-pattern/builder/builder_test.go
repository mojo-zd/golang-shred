package builder

import "testing"

func TestBuilder(t *testing.T) {
	builder := new(CarBuilder)
	director := new(Director).SetBuilder(builder)
	car := director.Generate()
	t.Log("car seat level[" + car.Seat.Level + "], " +
		"engine[" + car.Engine.Level + "], " +
		"wheel[" + car.Wheel.Name + "]")
}
