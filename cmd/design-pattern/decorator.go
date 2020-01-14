package design_pattern

type Component interface {
	Calc() int
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		c, num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
