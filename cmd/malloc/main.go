package main

//#cgo LDFLAGS:
//char* allocMemory();
import "C"
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	loop := 10
	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		lp, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err == nil {
			loop = int(lp)
		}
	}
	for i := 1; i <= loop; i++ {
		fmt.Println("Allocating memory...")
		C.allocMemory()
		time.Sleep(time.Second * 10)
	}
}
