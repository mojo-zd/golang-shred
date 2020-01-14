#### 效率提升
- string、bytes避免内存开销较大的高效装换方式。(二者转换底层会进行数据结构复制导致GC效率变低)
这两种类型转换是带内存分配与拷贝开销的，但有一种办法(trick)能够避免开销。利用了string和slice在runtime里结构只差一个Cap字段实现的

- string拼接. 小量小文本拼接使用+, 大量小文本使用strings.Join, 大量大文本使用`bytes.Buffe`