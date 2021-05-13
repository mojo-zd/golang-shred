package memento

// 源发器类
type Emp struct {
	Name   string
	Age    int
	Salary float64
}

func (e *Emp) Memento() Memento {
	return Memento{Name: e.Name, Age: e.Age, Salary: e.Salary}
}

func (e *Emp) Recovery(memento Memento) {
	e.Name = memento.Name
	e.Age = memento.Age
	e.Salary = memento.Salary
}
