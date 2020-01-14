package design_pattern

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
	monitorBoard *MonitorBoard
}

func NewStartCommand(mb *MonitorBoard) Command {
	return &StartCommand{
		monitorBoard: mb,
	}
}

func (sc *StartCommand) Execute() {
	sc.monitorBoard.Start()
}

type RebootCommand struct {
	monitorBoard *MonitorBoard
}

func NewRebootCommand(mb *MonitorBoard) Command {
	return &RebootCommand{
		monitorBoard: mb,
	}
}

func (rc *RebootCommand) Execute() {
	rc.monitorBoard.Reboot()
}

type MonitorBoard struct {
}

func (mb *MonitorBoard) Start() {
	fmt.Println("execute start...")
}

func (mb *MonitorBoard) Reboot() {
	fmt.Println("execute reboot")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
