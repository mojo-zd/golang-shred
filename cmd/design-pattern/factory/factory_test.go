package factory

import "testing"

func TestSampleFactory(t *testing.T) {
	factory := NewSimpleFactory()
	factory.CreateAudiCar().Running()
	factory.CreateBydCar().Running()
	t.Log("============")
	NewSimpleCarFactory("奥迪").Running()
	NewSimpleCarFactory("比亚迪").Running()
}

func TestFactoryMethod(t *testing.T) {
	NewAudiFactory().Create().Running()
	NewBydFactory().Create().Running()
}

func TestAbstractFactory(t *testing.T) {
	highFactory := NewHighFactory()
	highFactory.CreateSeat().Massage()
	highFactory.CreateTyre().Wear()
	t.Log("=========")
	lowFactory := NewLowFactory()
	lowFactory.CreateSeat().Massage()
	lowFactory.CreateTyre().Wear()
}
