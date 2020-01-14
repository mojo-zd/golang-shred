package design_pattern

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c, 10)
	c = WrapMulDecorator(c, 8)
	res := c.Calc()
	t.Logf("res %d\n", res)
}
