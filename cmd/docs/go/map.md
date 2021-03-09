map使用key、value存储, 通过key来做hash运算。为了解决hash冲突, 可以采用`开放寻址`、`拉链法`
golang的map底层使用hash表实现, 一个hash表可以有多个hash表节点,即bucket,而每个bucket则保存了一个或一组键值对
```
type hmap struct {
 count
 B
 buckets
}
```