package single

import "testing"

func TestSingleton(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 != s2 {
		t.Fatal("instance is not equal")
	}
	t.Log("s1 == s2")
}
