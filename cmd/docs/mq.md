#### RabbitMq
RabbitMq实现了高级消息队列协议的服务
- 为什么选择RabbitMq
1. 实现了高级消息队列协议
2. 可靠性，支持消息持久化
3. Erlang支持高并发
4. 集群简单
5. 社区活跃

- 角色
1. 生产者
2. 消费者

- 消息发送原理
App ----tcp连接---- RabbitMq   连接完成以后应用程序和RabbitMq之间就建立了一个信道(channel)
AMQP都是通过channel发送出去的

- 交换器
1. direct
2. headers 可以匹配AMQP的header而不是路由键
3. fanout
4. topic

- ACK机制
应用程序接受了消息如果未手动确认，消息队列的状态会从Ready变为Unacked。对于未确认的APP Rabbit将不会发送更多的消息

- 如何解决RabbitMq消息消费乱序问题
1. 将任务做成hash固定消费者

- 消息重复消费
在执行超时或者断开连接以后会导致重复消费
1. 消费者最好支持幂等性操作
2. 任务做成hash固定消费者 用map记录消费情况