### kubernetes中的设计模式
- out of tree
模式:
```go
// 定义Option类型
type Option func(registry runtime.Registry) error

func Apply(options ...Option) {
    // 使用的地方申明registry
    var registry runtime.Registry
	for _, option := range Options {
        option(registry)	
    }   
}
```
- visitor
- option
