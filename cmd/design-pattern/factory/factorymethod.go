package factory

type Factory interface {
	Create() Car
}

// 奥迪工厂
type audiFactory struct {
}

func (*audiFactory) Create() Car {
	return &Audi{}
}

func NewAudiFactory() Factory {
	return &audiFactory{}
}

// 比亚迪工厂
type bydFactory struct {
}

func (*bydFactory) Create() Car {
	return &Byd{}
}

func NewBydFactory() Factory {
	return &bydFactory{}
}
