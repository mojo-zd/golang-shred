package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	emp := Emp{Name: "mojo", Age: 30, Salary: 1000}
	t.Log(fmt.Sprintf("第一次打印: 名字[%s], 年龄[%d], 薪资[%f]", emp.Name, emp.Age, emp.Salary))

	taker := CareTaker{}
	taker.Memento = emp.Memento() //做一次备忘

	emp.Name = "mojo new"
	emp.Age = 31
	emp.Salary = 2000
	t.Log(fmt.Sprintf("第二次打印: 名字[%s], 年龄[%d], 薪资[%f]", emp.Name, emp.Age, emp.Salary))

	t.Log("恢复数据")
	emp.Recovery(taker.Memento)
	t.Log(fmt.Sprintf("第三次打印: 名字[%s], 年龄[%d], 薪资[%f]", emp.Name, emp.Age, emp.Salary))
}
