#### 为什么使用prometheus
1. 简单、高效
2. 动态服务发现
3. 可扩展(多个prometheus组成联邦)
4. 强大的查询语言
5. 强大的数据模型
6. 易于集成、社区活跃(支持其他监控系统的集成)

#### prometheus核心组件
- prometheus server
负责监控数据的获取、存储、查询

注意: 大规模集群下可以通过联邦和功能分区的方式进行扩展

#### prometheus数据类型
- counter(只增不减)
- gauge(仪表盘) 用于反映系统的当前状态
- Histogram(直方图)  // 统计样本
- Summary(摘要)  // 分析数据分布情况

#### prometheus运行原理
采集所有样本以时间序列的方式存储在内存数据库中，并定期刷新到磁盘。样本的组成如下:
1. metrics name 以及描述当前样本的labelset
2. timestamp  一个精确到毫秒的时间戳
3. value  一个float64的当前样本值

#### 时间序列类别
1. 瞬时向量
2. 区间向量
3. 位移操作 offset
4. 聚合函数 eg:sum  返回的时间序列的特征不是唯一的情况下可以做聚合操作
5. 标量(只有值 没有时序) 可以使用内置函数scalar讲瞬时向量转换为标量

#### 集合运算
1. and(并且)
2. or
3. unless
vector1 and vector2  包含vector1中完全匹配vector2的向量
vector1 or vector2  包含vector1的所有数据以及vector2中没有匹配到的数据
vector1 unless vector2  包含vector1中没有匹配到vector2的所有元素

#### 内置函数
- irate更精确
- rate会陷入长尾问题
- label_replace
- label_join

#### 告警
- 分组 避免接受大量重复的通知,将同一分组的告警合并为一个通知
- 抑制 当某一告警发出后，可以停止重复发送由此告警引发的其它告警的机制
- 静默 提供了一个简单的机制可以快速根据标签对告警进行静默处理。如果接收到的告警符合静默的配置，Alertmanager则不会发送告警通知