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

#### 锁
Once.Do实现原理, 采用atomic进行原子操作
1. Mutex互斥锁 当一个goroutine获取锁后其他goroutine就只能等待锁被释放才能进行操作
2. RWMutex读写锁，多个goroutine可同时获得读锁，但此时不能被写入。写锁会阻止任何读写操作，被当前写独占

#### chan
不管是否带有缓冲都是进行FIFO的原则

#### 值与指针
golang中函数的参数传递其实是值传递, 是对原有数据的一份拷贝。如果参数传递的是一个指针，则传递的是这个指针的拷贝
> map, chan通过make以后其实返回的是*hmap、*hchan，所以在通过参数传递以后对map的修改会反应到外部
> slice中元素指向是的data的指针,但是永远修改不了len和cap，因为他们是属于值拷贝。如果要修改需要传入*slice

#### 工具
- golint静态代码检查工具
- golangci-lint定制化检查

public interface Channel<T>
{
    void Send(T data);

    T Receive();
}

public class Barrier
{
    public Barrier(int n)
    {

    }

    public void Wait()
    {
    
    }
}
