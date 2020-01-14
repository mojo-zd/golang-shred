package design_pattern

// Target 适配器的目标接口
type Target interface {
	Request() string
}

//NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecRequest() string
}

type adapteeImpl struct {
}

func (ai *adapteeImpl) SpecRequest() string {
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{Adaptee: adaptee}
}

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecRequest()
}
