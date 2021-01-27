package design_pattern

import (
	"strings"
	"testing"
)

func TestMapStrToStr(t *testing.T) {
	out := MapStrToStr([]string{"Hello", "World", ",", "Hello", "Mojo"}, func(s string) string {
		return strings.ToLower(s)
	})
	t.Log("map str to str", out)
	oo := MapStrToInt([]string{"hello", "world", "hello", "mojo"}, func(s string) int {
		return len(s)
	})
	t.Log("map str to int", oo)
	filter := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(s int) bool {
		return s%2 == 0
	})
	t.Log("filter", filter)
}

func TestMap(t *testing.T) {
	out := Map([]string{"Hello", "Today"}, func(s string) string {
		return strings.ToLower(s)
	})
	t.Log("generic map", out)
}
