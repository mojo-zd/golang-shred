package iterator

// 定义迭代器
type MyIterator interface {
	First() interface{}
	Last() interface{}
	HasNext() bool
	GetCurrent() interface{}
}

// myIteratorImp 定义迭代器的实现
type myIteratorImp struct {
	cursor int
	data   []interface{}
}

func (mi *myIteratorImp) First() interface{} {
	return mi.data[0]
}

func (mi *myIteratorImp) Last() interface{} {
	return mi.data[mi.cursor-1]
}

func (mi *myIteratorImp) HasNext() bool {
	return mi.cursor < len(mi.data)
}

func (mi *myIteratorImp) Next() interface{} {
	mi.cursor++
	if mi.cursor < len(mi.data) {
		return nil
	}
	return mi.data[mi.cursor]
}

type MyIteratorAggregate struct {
	container []interface{}
	iterator  MyIterator
}

func (mia *MyIteratorAggregate) Add(obj interface{}) {
	mia.container = append(mia.container, obj)
}

func (mia *MyIteratorAggregate) Remove(index int) {
	mia.container = append(mia.container[:index], mia.container[index+1:])
}
