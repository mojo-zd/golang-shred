package basics

import "fmt"

type Notify interface {
	notify()
}

type Dog struct {
}

func (d Dog) notify() {
	fmt.Println("implement notify")
}

func sendNotify(notify Notify) {
	notify.notify()
}
