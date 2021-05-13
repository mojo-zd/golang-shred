package observer

import "fmt"

type Observer struct {
	UUID string
}

func (obs *Observer) Update(subject Subject) {
	fmt.Println("observer["+obs.UUID+"] receive msg", subject.Msg())
}
