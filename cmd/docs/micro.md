#### 什么是微服务
微服务是一些协同工作小而自治的服务。避免进行过多的暴露,暴露过多肯能引起消费放与服务内部产生高度耦合,从而降低服务的自治性。
微服务的黄金法则: 是否能够修改一个微服务对其进行部署而不影响其他任何服务

#### 概念
1) 内聚性: 将相关代码放到一起
2) 单一职责原则

#### 微服务优势
1) 技术异构(各个模块使用不同的技术 eg: 开发语言,存储等等、尝试在单个服务中使用新技术)
2) 弹性(舱壁: 系统中的某一个组件不可用并没有导致连级故障, 系统的其他组件还以正常工作)
3) 扩展(仅对单个需要进行扩展的服务进行扩展,无需对整个系统进行扩展)
4) 简化部署
5) 可组合性
6) 共享库

#### 微服务设计理论及技术
1) 分布式事务
2) CAP理论
3) 服务容错(不能让某个服务异常导致整个系统崩掉) 被正常处理的请求、错误请求、被访问的服务宕机

#### 代码治理
1) 范例
2) 裁剪代码模板

#### 技术债务
因为某些特殊的需求而忽略了一些约束,导致短期内带来利益长期来看需要付出代价。

#### 什么样的服务是好服务
1) 松耦合
2) 高内聚

#### 界限上下文
一个界限上下文包括两个部分, 每个上下文都有明确规定需要暴露那些接口给其他上下文
1) 需要与外部通信
2) 不需要与外部通信

#### 服务集成准则
微服务之间通信方式比较多样化, SOAP、XML-RPC、REST、Protocol Buffer
1) 避免破坏性修改  有时候对API的修改会导致消费方也会方式改变
2) 保证API技术无关性
3) 使服务易于消费方使用
4) 隐藏内部实现

#### 服务集成方式
1) 共享数据库(不推荐)
2) 同步或异步通信, 最好设置最大重试次数否则可能会导致程序一直崩溃
   1. 基于事件的异步协作方式 eg: mq
   2. 基于请求/相应的异步协作方式 eg: http、rpc
   
#### 版本管理
语义化版本管理 MAJOR.MINOR.PATCH
MAJOR 包含向后不兼容的修改
MINOR 有新功能增加但向后兼容
PATCH 代表已有功能的修复


### 微服务治理
#### 基本层面
- 并发控制、限流
- 服务注册与发现
- 负载均衡
- 超时控制
- 熔断设计
#### 高阶功能
- 请求认证
- 链路追踪
- 日志
- 可观测性
   
### 扩展
- 服务配置中心