package design_pattern

import "fmt"

type Observer interface {
	Update(*Sub)
}

type Sub struct {
	observers []Observer
	context   string
}

func NewSub() *Sub {
	return &Sub{
		observers: make([]Observer, 0),
	}
}

func (s *Sub) Attach(o ...Observer) {
	s.observers = append(s.observers, o...)
}

func (s *Sub) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Sub) UpdateContent(content string) {
	s.context = content
	s.notify()
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Sub) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
