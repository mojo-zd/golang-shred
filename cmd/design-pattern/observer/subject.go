package observer

import (
	"errors"
	"fmt"
)

type Subject interface {
	Notify() error
	Registry(observers ...*Observer)
	Remove(observer *Observer)
	Msg() string
}

type SubjectBasic struct {
	observers map[string]*Observer
	Subject
}

func (basic *SubjectBasic) Notify() error {
	if basic.observers == nil {
		fmt.Println("no observer founded")
		return errors.New("no observer founded")
	}

	if basic.Subject == nil {
		fmt.Println("subject not set")
		return errors.New("subject not set")
	}

	for _, obs := range basic.observers {
		obs.Update(basic.Subject)
	}
	return nil
}

func (basic *SubjectBasic) Registry(observers ...*Observer) {
	for _, obs := range observers {
		if _, ok := basic.observers[obs.UUID]; !ok {
			basic.observers[obs.UUID] = obs
			fmt.Println("registry observer[", obs.UUID, "]success...")
			continue
		}
		fmt.Println("observer[", obs.UUID, "] has been registry")
	}

}

func (basic *SubjectBasic) Remove(observer *Observer) {
	delete(basic.observers, observer.UUID)
}

type subjectMsg struct {
	SubjectBasic
	Context string
}

func NewSubjectMsg() Subject {
	sub := &subjectMsg{}
	sub.SubjectBasic = SubjectBasic{make(map[string]*Observer), sub}
	return sub
}

func (sub *subjectMsg) SetContext(ctx string) {
	sub.Context = ctx
}

func (sub *subjectMsg) Msg() string {
	return "i'm message subject[" + sub.Context + "]"
}

type subjectMail struct {
	SubjectBasic
	Context string
}

func NewSubjectMail() Subject {
	sub := &subjectMail{}
	sub.SubjectBasic = SubjectBasic{make(map[string]*Observer), sub}
	return sub
}

func (sub *subjectMail) SetContext(ctx string) {
	sub.Context = ctx
}

func (sub *subjectMail) Msg() string {
	return "i'm mail subject[" + sub.Context + "]"
}
