package proxy

import "testing"

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	if res := sub.Do(); res != "pre:real:after" {
		t.Fatal("proxy call failed")
	}
}
