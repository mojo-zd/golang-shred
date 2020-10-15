#### 如何将golang源码转为汇编指令
go build -gcflags -S main.go

#### 如何节省内存
1. 数据结构层面
2. 定义数据类型(如果能用局部变量尽量使用局部变量)
3. 避免闭包中延长变量的生命周期

#### 编译原理


#### 逃逸分析
1. 闭包可能会导致变量从栈上逃逸到堆上，给GC造成压力
```
func myFunc() func() int {
    foo := 0
    return func () int {
        foo++
        return foo
    }
}

bar := myFunc()
val1 := bar() // 1
val2 := bar() // 2
// 由于闭包导致foo的生命周期被延长，从栈移到堆上了
```