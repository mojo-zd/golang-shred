## istio
- 数据平面   
由一组智能代理（Envoy）组成，被部署为 sidecar。这些代理负责协调和控制微服务之间的所有网络通信。
他们还收集和报告所有网格流量的遥测数据。
- 控制平面   
管理并配置代理来进行流量路由

### Envoy
由 Envoy 代理启用的一些 Istio 的功能和任务包括:
- 流量控制功能：通过丰富的 HTTP、gRPC、WebSocket 和 TCP 流量路由规则来执行细粒度的流量控制
- 网络弹性特性：重试设置、故障转移、熔断器和故障注入
- 安全性和身份验证特性：执行安全性策略以及通过配置 API 定义的访问控制和速率限制
- 基于 WebAssembly 的可插拔扩展模型，允许通过自定义策略实施和生成网格流量的遥测 

### Pilot
Pilot 为 Envoy sidecar 提供服务发现、用于智能路由的流量管理功能（例如，A/B 测试、金丝雀发布等）
以及弹性功能（超时、重试、熔断器等）

### Citadel
Citadel 通过内置的身份和证书管理，可以支持强大的服务到服务以及最终用户的身份验证

### Galley
Galley 是 Istio 的配置验证、提取、处理和分发组件。它负责将其余的 Istio 组件与从底层平台
（例如 Kubernetes）获取用户配置的细节隔离开来

### Mixer
负责策略控制和遥测数据收集

### 流量控制

### 熔断
在熔断器中，设置一个对服务中的单个主机调用的限制，例如并发连接的数量或对该主机调用失败的次数。
一旦限制被触发，熔断器就会“跳闸”并停止连接到该主机

### 故障注入
在配置了网络，包括故障恢复策略之后，可以使用 Istio 的故障注入机制来为整个应用程序测试故障恢
复能力.故障注入是一种将错误引入系统以确保系统能够承受并从错误条件中恢复的测试方法。使用故障注
入特别有用，能确保故障恢复策略不至于不兼容或者太严格，这会导致关键服务不可用。

### 可观察性
Istio 为网格内所有的服务通信生成详细的遥测数据。这种遥测技术提供了服务行为的可观察性.
- 指标  
Istio 基于 4 个监控的黄金标识（延迟、流量、错误、饱和）生成了一系列服务指标
- 分布式追踪  
Istio 为每个服务生成分布式追踪 span，运维人员可以理解网格内服务的依赖和调用流程
- 访问日志  
当流量流入网格中的服务时，Istio 可以生成每个请求的完整记录，包括源和目标的元数据。此信息使
运维人员能够将服务行为的审查控制到单个工作负载实例的级别