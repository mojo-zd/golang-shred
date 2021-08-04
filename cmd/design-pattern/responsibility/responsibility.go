package responsibility

import (
	"fmt"
	"strconv"
)

/**
责任链模式, 以请假为例
*/

type LeaveRequest struct {
	Name   string // 请假人姓名
	Day    int    // 请假天数
	Reason string // 请假原因
}

type Leave interface {
	HandlerRequest(request LeaveRequest)
}

// Director 主任审批
type Director struct {
	Next Leave
}

func (d *Director) HandlerRequest(request LeaveRequest) {
	if request.Day < 3 {
		fmt.Println("主任审批:"+request.Name, ", 请假"+strconv.Itoa(request.Day)+"天,原因:"+request.Reason)
		return
	}

	if d.Next != nil {
		d.Next.HandlerRequest(request)
	}

}

type ViceManager struct {
	Next Leave
}

func (v *ViceManager) HandlerRequest(request LeaveRequest) {
	if request.Day < 5 {
		fmt.Println("副经理审批:"+request.Name, ", 请假"+strconv.Itoa(request.Day)+"天,原因:"+request.Reason)
	}

	if v.Next != nil {
		v.Next.HandlerRequest(request)
		return
	}
}

// 经理审批
type Manager struct {
	Next Leave
}

func (m *Manager) HandlerRequest(request LeaveRequest) {
	if request.Day < 10 {
		fmt.Println("经理审批:"+request.Name, ", 请假"+strconv.Itoa(request.Day)+"天,原因:"+request.Reason)
		return
	}

	if m.Next != nil {
		m.Next.HandlerRequest(request)
		return
	}
}
