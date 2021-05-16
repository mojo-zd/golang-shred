package flyweight

import "fmt"

// ChessFlyWeight 定义享元模式抽象类
type ChessFlyWeight interface {
	SetColor(color string)         // 设置棋子颜色,属于内部状态
	Color() string                 // 获取棋子颜色
	Display(coordinate Coordinate) // 显示棋子位置 + 颜色
}

// Coordinate 定义棋子坐标, 享元模式外部状态会随着环境变化而变化
type Coordinate struct {
	X, Y int // X、Y坐标
}

// ConcreteChessFlyWeight 享元模式具体实现
type ConcreteChessFlyWeight struct {
	color string
}

func (fw *ConcreteChessFlyWeight) SetColor(color string) {
	fw.color = color
}

func (fw *ConcreteChessFlyWeight) Color() string {
	return fw.color
}

func (fw *ConcreteChessFlyWeight) Display(coordinate Coordinate) {
	fmt.Println("棋子颜色:", fw.color, ",棋子位置:", coordinate.X, "------", coordinate.Y)
}

type FlyWeightFactory struct {
	weights map[string]ChessFlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{weights: make(map[string]ChessFlyWeight)}
}

func (ff *FlyWeightFactory) GetChess(color string) ChessFlyWeight {
	if ff.weights == nil {
		ff.weights = make(map[string]ChessFlyWeight)
	}

	if _, ok := ff.weights[color]; !ok {
		ff.weights[color] = &ConcreteChessFlyWeight{color: color}
	}

	return ff.weights[color]
}
