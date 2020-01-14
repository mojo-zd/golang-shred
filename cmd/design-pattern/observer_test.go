package design_pattern

import "testing"

func TestSub(t *testing.T)  {
	sub := NewSub()
	reader1 := NewReader("kimo")
	reader2 := NewReader("kinder")
	reader3 := NewReader("mojo")
	sub.Attach(reader1, reader2, reader3)
	sub.UpdateContent("shopping")
}