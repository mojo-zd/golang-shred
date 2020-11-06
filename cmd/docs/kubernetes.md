- headless service 和 service的区别
1. headless没有负载效果
2. 客户端自主选择ep

- finalizers
定义一种删除策略。某些kubernetes对象是其他一些对象的属主  eg: ReplicaSet是一组pod的属主,具有属主的对象被称为附属

- admission webhook
对象持久化之前对api server的请求进行拦截,在请求进行身份认证或授权之后放行通过. validating进行校验操作, mutating进行数据修改操作

- kubernetes pause容器
让pod中的多个容器共享某些资源, pause为了解决pod中的网络共享问题,其他的容器通过join namespace的方式加入到infra container(pause)的network namespace中
1. 在pod中担任Linux命名空间共享的基础
2. 启用pid命名空间，开启init进程

- operator
生成工具
KUDO (Kubernetes 通用声明式 Operator)  https://kudo.dev/
kubebuilder https://book.kubebuilder.io/
Metacontroller https://metacontroller.app/

- pod创建流程
kubectl ---> api-server-->认证授权-->创建deploy的etcd记录-->deploy controller监听-->创建rs--> rs controller监听-->创建pod--> 
schedule监听选择合适的主机--> kubelet监听获取信息进行资源分配-->kubeproxy进行ipvs下发

- List-Watch机制
1. List
2. Watch
informer监听机制 https://juejin.im/post/6844903631787917319

- pod回调方式
pod回调包含PostStart(容器创建之后)、PreStop(容器停止之前)
1. exec
2. http

#### kubernetes扩展方式
1. crd
2. controller
3. schedule extender
4. scheduler framework、aggregated APIServer

- schedule extender
kube-scheduler先执行内置filter，然后调度再通过http调用extender注册的webhook，将调度的pod和node信息发送给extender，根据返回的filter结果作为最终结果
弊端:
性能较差，无法支持高吞吐量（序列化反序列化，http调用）

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
readinessProbe 通过检查pod状态讲ready设置为true|false，为false时候service关联的ep将会被移除，放置流量路由到不健康的实例
livenessProbe 检查容器应用是否正常，不正常就按照重启策略进行操作 

#### pod阶段
pod阶段: Pending、Running、Succeeded、Failed、Unknown
容器状态: Waiting、Running、Terminated

#### CNI
cni分为三种Overlay、router、underlay

#### 工具
1. kubectl-debug

#### 调试小妙招
for i in $(kubectl get deploy | awk 'NR>1 {print $1}'); do kubectl get deploy $i --template '{{range $idx, $c := .spec.template.spec.containers}}{{$c.image}}
{{end}}'; done

#### sts
- sts使用场景
1. 稳定的唯一网络标识(eg: pod-1, pod-2)
2. 稳定的持久化存储

- sts特性
1. 采用partition方式可实现灰度部署
2. 执行podManagementPolicy="Parallel"可实现并行删除sts pod