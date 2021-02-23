## etcd
etcd使用raft协议(一种共识算法)

### etcd节点角色
- follower 追随者
- candidate 候选人(临时状态) 
- leader 

1. follower节点
所有的follower节点都是被动选择的,他不会主动发送任何请求,只会响应leader、candidate的请求. 


### raft
raft协议分为三个子问题, 节点选主、日志复制、安全性
1. 节点选主
所有节点初始化时都为follower状态, 接下来所有节点会进入一个超大的for select从channel中获取待处理的事件
