## `go test`
写测试时需要在测试前或者在测试后进行额外的设置(setup)或拆卸(teardown), 有时测试还需要控制
在主线程上运行的代码。为了满足这些需求, testing包提供了`TestMain`函数. e.g.
```
func TestMain(m *testing.M)
```

### `httptest`
go内置了`httptest`包用于测试http请求