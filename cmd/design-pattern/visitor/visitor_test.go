package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}
	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
}

func TestDec(t *testing.T) {
	dec(func(s string) {
		fmt.Println(s)
	})("hello world")
}

func dec(fn func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("start func")
		fn(s)
		fmt.Println("end func")
	}
}
