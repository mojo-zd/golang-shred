package basics

import (
	"github.com/rs/zerolog/log"
	"math"
)

type Shape interface {
	Area() float32
}

// 接口嵌套 任何实现了Shape、Object定义的方法则也实现了类型Math
type Math interface {
	Shape
	Object
}

type Object interface {
	Perimeter() float32
}

type Circle struct {
	radius float32
}

type Square struct {
	side float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func interfaceAssert() {
	c := Circle{2}
	var (
		s Shape  = c
		o Object = c
	)

	// 各自只能调用各自的方法
	s.Area()
	o.Perimeter()
	//	增加断言处理 断言成功可以调用circle的所有方法
	r, ok := s.(Circle)
	if ok {
		r.Perimeter()
		r.Area()
		log.Info().Msg("interfaceAssert success")
	}
}

// 类型选择器  circle实现了 object、shape所以第一个满足条件就终止
func switchType(i interface{}) {
	switch i.(type) {
	case Object:
		log.Info().Msg("object selector")
	case Shape:
		log.Info().Msg("shape selector")
	case Circle:
		log.Info().Msg("circle selector")
	}
}

func interfaceEmbed() {
	c := Circle{2}
	var m Math = c
	log.Info().Interface("area", m.Area()).Interface("perimeter", m.Perimeter()).Send()
}

// 值调用、指针调用
func valueAndPointerCall() {
	c := Circle{2}
	// 注意:必须使用指针类型才可以调用
	square := &Square{4}
	var (
		s  Shape = c
		sp Shape = &c
		sq Shape = square
	)

	log.Info().Interface("value call", s.Area()).
		Interface("pointer call", sp.Area()).
		Interface("only pointer call", sq.Area()).Send()
}
