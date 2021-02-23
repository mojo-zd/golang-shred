- headless service 和 service的区别
1. headless没有负载效果
2. 客户端自主选择ep

- finalizers
定义一种删除策略。某些kubernetes对象是其他一些对象的属主  eg: ReplicaSet是一组pod的属主,具有属主的对象被称为附属
联级删除策略:
1. Foreground 删除所有附属以后才会删除属主
2. Background 先删除属主, 后台执行删除附属

- admission webhook
对象持久化之前对api server的请求进行拦截,在请求进行身份认证或授权之后放行通过. validating进行校验操作, mutating进行数据修改操作

- kubernetes pause容器
让pod中的多个容器共享某些资源, pause为了解决pod中的网络共享问题,其他的容器通过join namespace的方式加入到infra container(pause)的network namespace中,
pod中其他业务容器共享了pause容器的网络栈、存储卷。
1. PID命名空间: pod中不同的应用程序可以看到其他应用程序的进程ID, 开启init进程
2. 网络命名空间: pod中的多个容器能够访问同一个IP和端口范围
3. IPC命名空间: pod中的多个容器能够使用systemv ipc或posix消息队列进行通信
4. UTS命名空间: pod中的多个容器共享一个主机名
5. 共享存储: pod中的多个容器可以共享pod级别定义的存储卷


- operator
生成工具
KUDO (Kubernetes 通用声明式 Operator)  https://kudo.dev/
kubebuilder https://book.kubebuilder.io/
Metacontroller https://metacontroller.app/

- pod创建流程
kubectl ---> api-server-->认证授权-->创建deploy的etcd记录-->deploy controller监听-->创建rs--> rs controller监听-->创建pod--> 
schedule监听选择合适的主机--> kubelet监听获取信息进行资源分配-->kubeproxy进行ipvs下发。
kubelet收到消息后通过CNI创建网络、CRI创建容器、CSI创建存储卷的挂载, 然后等待进程启动...

- List-Watch机制
1. List
2. Watch
informer监听机制 https://juejin.im/post/6844903631787917319

组件:
Reflect 监听资源变化 把信息放入delta queue
Informer 注册回调函数，获取队列信息交给indexer组件做缓存处理。回调函数需要处理数据时直接从缓存中获取
Indexer 接受消息缓存数据到本地

- pod回调方式
pod回调包含PostStart(容器创建之后)、PreStop(容器停止之前)
1. exec
2. http

#### kubernetes高可用
- 堆控制平面 etcd和控制平面在同一个集群中
- 使用外部集群方式  etcd和控制平面分开部署


### 控制平面
主要包含组件
- kube-apiserver
- kube-scheduler
- etcd
- kube-controller-manager (理论上讲控制器都是单独的进程,为了降低复杂度被编译到一个可执行文件中并在一个进程中执行), 
包括node-controller、replicas-controller、endpoints-controller、service account & token controller

### Node组件
- kubelet
- kube-proxy
- CRI

### 插件
https://kubernetes.io/zh/docs/concepts/cluster-administration/addons/

#### kubernetes扩展方式
1. crd
2. controller
3. scheduler extender
4. scheduler framework、aggregated APIServer

- schedule extender
kube-scheduler先执行内置filter，然后调度再通过http调用extender注册的webhook，将调度的pod和node信息发送给extender，根据返回的filter结果作为最终结果
弊端:
性能较差，无法支持高吞吐量（序列化反序列化，http调用）
只能在Filter、Prioritize、Bind之后使用
需要自己进行缓存处理

- multiple schedulers
和default scheduler平级
弊端:
研发成本较高

- scheduler framework
为了Kube-scheduler扩展性更好、代码更简洁, 社区推出了scheduler framework


#### 依赖启动
- initContainer方式
1. 通过initContainer定义exec或者curl方法检查第三方组件是否成功
2. 被依赖的组件需要配置readness检查

- 业务逻辑处理方式
不能连接直接crash业务容器

#### 健康检查
readinessProbe 通过检查pod状态将ready设置为true|false，为false时候service关联的ep将会被移除，防止流量路由到不健康的实例
livenessProbe 检查容器应用是否正常，不正常就按照重启策略进行操作 

#### pod阶段
pod阶段: Pending、Running、Succeeded、Failed、Unknown
容器状态: Waiting、Running、Terminated

#### CNI
cni分为三种Overlay、router、underlay

#### CSI
CSI的发展经历了in-tree、flexVolume、csi
1. in-tree即为讲插件打包进行kubernetes可执行文件  授权、可执行文件偏大、重新编译要处理整个可执行文件、插件越多crash风险越高
2. flexVolume 通过外部脚本集成外部存储  驱动需要通过宿主机的跟文件系统访问脚本、如果插件需要依赖第三方包需要在所有节点解决依赖和兼容问题
3. CSI分为一下三个部分
1) Identity Service 用于返回插件信息
2) Controller Service 实现Volume的CRUD
3）Node Service 运行在所有的节点, 用于实现把所有的volume挂载到当前节点的指定目录

- 分布式存储部署
1. 设计一套针对云原生的分布式存储系统 e.g. Longhorn、OpenEBS
2. 通过适配器的方式 适配原有的分布式存储 e.g. Rook

- Longhorn
Longhorn 存储的xxx.img文件采用raw(linux Sparse稀疏文件)格式文件,会受到磁盘大小和性能限制, 提供了内置Web ui

#### 工具
1. kubectl-debug

#### 调试小妙招
1. 查看deploy中所有的image
```
for i in $(kubectl get deploy | awk 'NR>1 {print $1}'); do kubectl get deploy $i --template '{{range $idx, $c := .spec.template.spec.containers}}{{$c.image}}
{{end}}'; done
```
2. 查看主机上的块设备
```
lsblk -S
```

#### sts
- sts使用场景
1. 稳定的唯一网络标识(eg: pod-1, pod-2)
2. 稳定的持久化存储

- sts特性
1. 采用partition方式可实现灰度部署
2. 执行podManagementPolicy="Parallel"可实现并行删除sts pod

### 网络组件
#### flannel
flannel支持多种网络模型, vxlan、udp、hostgw、ipip， vxlan和udp的区别在于vxlan是内核封包而udp是flanneld用户态进程封包，
所以udp的性能会差一些。hostgw是一种主机网关模式容器到另一个主机上容器的网关设置成所在主机的网卡地址，和calico非常相似，
只不过calico是通过BGP声明的，而hostgw是通过中心的etcd分发，所以hostgw是直连模式不需要经过overlay封包和拆包，性能比较高,
但是hostgw必须是在一个二层网络，毕竟吓一跳的路由必须在邻居表中，否则无法通行。