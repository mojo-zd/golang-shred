package basics

import "fmt"

const (
	a = iota
	b = iota
)

func iotaDefined() {
	fmt.Println(a, b)
}

const (
	name = "name"
	c    = iota
	d    = iota
)

func iotaMulit() {
	fmt.Println(name, c, d)
}

type Direction int

const (
	EAST Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func direction() {
	fmt.Println(South)
}

type Math1 struct {
	X, y int
}
