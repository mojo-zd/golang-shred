package memento

// 负责人: 负责管理备忘录对象
type CareTaker struct {
	Memento
	//list.List //存储所有变更到栈
}
