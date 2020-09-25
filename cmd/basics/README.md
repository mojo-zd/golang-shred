#### PGM
M 代表内核线程
P 代表一个go代码片段所需资源

#### 切片
- 切片扩容
切片在进行append时，如果<1024按照2倍扩容。反之按照1/4扩容

- 切片在函数间传递
切片发生复制时，底层数组不会被复制，数组大小也不会有影响

- 切片拷贝
copy(dst, src) 如果dst为nil则拷贝为0，反之以最小的容量进行拷贝(eg: dst cap为3, src cap为10 拷贝以后则为3)

- 子切片
```
slice := make([]int, 2, 10)
s1 := slice[4:] // 这种情况会出错，获取子切片时 如果没有指定j的值 j=len(slice) 又因为4>2所以会出错
```

#### map
- map拷贝
range每个元素到新的Map

- map函数间传递
map发生了复制，指向的底层数据不变，在函数内部修改会影响到源map的值

#### interface
示例如下
```
type Shape interface {
	Area() float32
}

// 接口嵌套 任何实现了Shape、Object定义的方法则也实现了类型Math
type Math interface {
	Shape
	Object
}

type Object interface {
	Perimeter() float32
}

type Circle struct {
	radius float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}
```

- 接口断言
```
c := Circle{2}
var s Shape = c
r, ok := s.(Circle)
if ok {
    r.Perimeter()
}
```

- 类型选择
```
func selectorType(i interface) {
    switch i.(type) {
        case Shape:

        case Object:

        case Circle:
    }
}
```

- 指针接受者、值接受者
值接受者可以使用值调用或者指针调用
```
// 值类型变量既可以使用值调用也可以使用指针调用
c := Circle{2}
var (
    s Shape = c
    sp Shape = &c
)

s.Area()
sp.Area()

// 必须使用指针类型才可以使用
square := &Square{4}
var sq Shape = square
sq.Area()
```

#### defer
defer详细使用请参考下面几篇解释
https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466918&idx=2&sn=151a8135f22563b7b97bf01ff480497b&chksm=f2474389c530ca9f3dc2ae1124e4e5ed3db4c45096924265bccfcb8908a829b9207b0dd26047&scene=21#wechat_redirect
https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466928&idx=1&sn=5d4f1c2c7802e3c7520b4e7c2fd05634&chksm=f247439fc530ca89da7b4446e6a5e3e87dc047d695ad1ad62916f2a7eccae244d7b4dfbca0de&scene=21#wechat_redirect
https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466923&idx=1&sn=cfca1b9031db3b5c8e6cb52e854fdb45&chksm=f2474384c530ca92722af02b7b36c6d4a7545ddd57b65fd3b48b61535da8417f89e12d79e5eb&scene=21#wechat_redirect

- recover调用时机
recover() 必须在 defer() 函数中直接调用才有效
```
func main() {
    defer func() {
        recover()
    }()
    panic(1)
}
```

#### chan
- select chan
```
ch1 := make(chan string)
ch2 := make(chan string)
select {
    case v1 := <-ch1:
        fmt.Println("from", v1)
    case v2 := <-ch2:
        fmt.Println("from", v2)
    case <- time.After(time.Second*2): // 超时处理
        fmt.Println("timeout")
    default: // 如果指定时间没有执行则执行下面的逻辑
        fmt.Println("no value receive")
}
```

- chan关闭判断
警告: 
1. 使用for range循环信道数据发送完毕以后必须关闭信道 否则会发生死锁
2. 有方向的chan不能被关闭

```
v, ok := <- ch
if !ok { //如果ok为false则表示该chan已经被关闭了 不能进行任何操作
}
```

- chan缓冲信道和信道容量
1. 无缓冲的信道会阻塞当前协程, 当缓冲信道没有被使用完时不会发生阻塞
2. 如果缓冲使用完毕则会阻塞当前协程, 如果没有其他的协程来读取数据则会发生死锁
3. 当信道中还有数据时进行close(ch)操作，还是可以读取数据的

- chan有单向信道类别(用于信道作为参数传递时, 提高程序的安全性)
```
rch := make(<-chan string) //单向发送信道
sch := make(chan<- string) //单向接受信道
```

- 给nil信道发送数据或者从nil信道中获取数据都会造成协程阻塞

#### 类型别名定义
```
type User struct{}
type User1 User // 基于User定义一个新的类型
type User2 = User // User2为User的别名
```

#### 锁
- 当锁被当做匿名字段时相关的方法必须使用指针接受者，否则锁将失效
解决方式: 
1. 要么使用指针接受者
2. 要么锁定义为指针类型