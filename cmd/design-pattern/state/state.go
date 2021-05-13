package state

import "fmt"

type StateType string

const (
	Free    StateType = "free"
	Booked  StateType = "booked"
	Checked StateType = "checked"
)

type State interface {
	Handler() error
}

type FreeState struct {
}

func (*FreeState) Handler() error {
	fmt.Println("空闲处理器!!!")
	return nil
}

type BookedState struct {
}

func (*BookedState) Handler() error {
	fmt.Println("预定处理器!!!")
	return nil
}

type CheckedState struct {
}

func (*CheckedState) Handler() error {
	fmt.Println("已入住处理器!!!")
	return nil
}

type Context interface {
	Handler(state StateType)
}

type hotelContext struct {
	State
	store map[StateType]State
}

func NewHotelContext() Context {
	return &hotelContext{store: map[StateType]State{
		Free:    &FreeState{},
		Booked:  &BookedState{},
		Checked: &CheckedState{},
	}}
}

func (ctx *hotelContext) Handler(state StateType) {
	ctx.store[state].Handler()
}
