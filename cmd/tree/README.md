问题1: slice何时传递指，何时传递指针
slice在进行参数传递时，传入的是原slice的拷贝，但是新slice元素指向的地址并没有变。

1. 如果在函数中对某个slice进行append操作, 
并不会反映到原slice，因为slice在进行append时cap不足会重新生成一个新的slice(小于1024容量时扩容2倍，大于1024时扩容25%),
此时如果想追加新的元素并反映到原slice需要传递指针

2. 如果只是修改slice中的某个元素可以直接值传递

问题2: 虽然效率有点低，我还是写以下
参考outOfOrder

问题3：二叉树构建+深度遍历

> https://github.com/mojo-zd