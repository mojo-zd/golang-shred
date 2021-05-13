## 服务注册发现
- 需要有权限限制

## API网关
- 提供统一访问入口

## DNS Server选择

## MEC
(Multi-access Edge Compute)多接入边缘计算

## 边沿计算
- openness
- edge gallery


### openness
- 一个边缘应用程序管理系统, 使服务提供商和企业能够在任何网络边缘上构建、部署和操作自己的
边缘应用程序

- OpenNESS 还是一个边缘网络的服务平台，支持在多云环境中跨越多种类型的网络平台以及多种类
型的访问技术。让用户能够轻松的编排和管理运行在边缘上的应用程序，并为客户端（UE）和边缘应用程序提供端到端（End to End）的网络连接服务



### 架构参考
- 腾讯云TMEC https://cloud.tencent.com/developer/article/1653624


每个边缘设备就是一个边缘基础环境，对外提供服务: 
- 部署基础业务支撑应用
- 基础环境
- 平台能力(网关、服务发现、镜像中心) ? 基础平台能力如何提供

### 5G引流
将终端发送给核心网的数据流量依据MEC业务的要求路由到MEC并分发给MEC业务处理(通过PCF、NEF).

### 网络分层
- 接入网 负责把数据收集上来(e.g. 端设备到基站)
- 承载网 负责把数据送来送去(基站机房双向)
- 核心网 管理中枢,负责管理数据对数据进行分拣

### MEP
MEC应用的集成部署、网络开放等中间件能力，可托管5G网络能力、业务能力等MEC服务

### QA
- MEP定义 https://carrier.huawei.com/~/media/CNBGV2/download/program/5G-MEC-IP-Network-White-Paper-cn-v2.pdf
提供对MEC节点部署的服务的服务注册发现、DNS、负载均衡、流量管理、API能力开放和安全功能
- 边缘侧提供基础层、平台层、业务层。云侧主要只是管理边缘云(监控、告警放在哪里?)
- MEPM 针对的监控、配置、性能等管理以及应用规则和需求进行管理,
  MEP既然已经对MEC相关组件进行管理为什么还要MEPM
- MEC编排(类似kubesphere的k8s系统?)
- 边缘部分是一套完整的系统(基于k8s 包括独立的镜像仓库、存储挂载?)
- openness是边缘管理系统(部署在中心云上的?)
- pacific 租户怎么定义的(如何做到网络隔离)?应用到边缘还是核心网络
- upf安全保障(upf下沉到边缘)