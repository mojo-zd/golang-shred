package basics

// 接口的应用场景
// 强制用户通过构造函数实例化对象 使用接口可以生成对反的文档, 私有的对象方法注释并不会被生成
type Widget interface {
	// 获取唯一id
	ID() string
}

type widget struct {
	id string
}

// 返回一个Widget实例
func NewWidget() widget {
	return widget{
		id: "random string",
	}
}

func (w widget) ID() string {
	return w.id
}
