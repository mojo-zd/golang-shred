package basics

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestIoReader(t *testing.T) {
	out, err := ReadFrom(os.Stdin, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
	time.Sleep(time.Minute)
}

type AA struct {
	Name string
}

func TestSeekEnd(t *testing.T) {
	//aa := AA{Name: "mojo"}
	fmt.Printf("%t", false)
}
