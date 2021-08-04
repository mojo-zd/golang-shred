package mediator

import "fmt"

// 中介者对象
type Mediator interface {
	Registry(dname string)
}

type Department interface {
	SelfAction() //本部门执行
	OutAction()  //其他部门执行
}

type Mediate struct {
	// 定义锁防止并发操作
	register map[string]Department
}

func (m *Mediate) Registry(dname string, department Department) {
	if m.register == nil {
		m.register = make(map[string]Department)
	}
	m.register[dname] = department
}

type Deployment struct {
}

func (*Deployment) SelfAction() {
	fmt.Println("开发...")
}

func (*Deployment) OutAction() {
	fmt.Println("跑客户...")
}
