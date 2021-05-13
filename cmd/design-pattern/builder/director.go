package builder

// 指挥者
type Director struct {
	Builder Builder
}

func (d *Director) SetBuilder(builder Builder) *Director {
	d.Builder = builder
	return d
}

func (d *Director) Generate() *Car {
	d.Builder.BuildEngine().BuildSeat().BuildWheel()
	return d.Builder.Instance()
}
