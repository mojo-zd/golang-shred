package builder

type Car struct {
	*Seat
	*Engine
	*Wheel
}

// 座椅
type Seat struct {
	Level string // 座椅级别 豪华、普通
}

// 轮胎
type Wheel struct {
	Name string // 米其林、马牌
}

// 引擎
type Engine struct {
	Level string // 高级、普通
}

type Builder interface {
	BuildSeat() Builder // 建造座椅
	BuildEngine() Builder
	BuildWheel() Builder
	Instance() *Car
}

type CarBuilder struct {
	Car Car
}

func (cb *CarBuilder) BuildSeat() Builder {
	cb.Car.Seat = &Seat{Level: "豪华座椅"}
	return cb
}

func (cb *CarBuilder) BuildEngine() Builder {
	cb.Car.Engine = &Engine{Level: "兰博基尼豪华版发动机"}
	return cb
}

func (cb *CarBuilder) BuildWheel() Builder {
	cb.Car.Wheel = &Wheel{Name: "马牌轮胎"}
	return cb
}

func (cb *CarBuilder) Instance() *Car {
	return &cb.Car
}
