package signalhandler

import (
	"fmt"
	"testing"
	"time"
)

func TestSignal(t *testing.T)  {
	stop := make(chan struct{})
	//close(stop)
	//fmt.Println("=== had close chan ==")


	go func() {
		<- stop
		fmt.Println("========")
		<- stop
		fmt.Println("=== 退出 ===")
	}()
	ticker := time.NewTicker(time.Second * 5)
	for  {
		select {
		case <- ticker.C:
			stop <- struct{}{}
			fmt.Println("====")
		}
	}
	//time.Sleep(time.Hour)
}