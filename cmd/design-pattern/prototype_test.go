package design_pattern

import "testing"

var pm *PrototypeManager

func TestPrototype(t *testing.T) {
	t1 := pm.Get("t1")
	t2 := t1.Clone()
	if t1 == t2 {
		t.Fatal("t1 clone not work!")
	}
}

func init() {
	pm = NewPrototypeManager()
	pm.Set("t1", &Type1{})
}
