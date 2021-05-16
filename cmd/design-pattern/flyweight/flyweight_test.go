package flyweight

import "testing"

func TestFlyWeight(t *testing.T) {
	chess1 := NewFlyWeightFactory().GetChess("黑色")
	chess2 := NewFlyWeightFactory().GetChess("黑色")
	t.Log("chess1 addr:", chess1, "chess2 addr:", chess2)
	chess1.Display(Coordinate{X: 10, Y: 20})
	chess2.Display(Coordinate{X: 14, Y: 50})
}
