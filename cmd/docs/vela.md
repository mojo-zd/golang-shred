## vela梳理
### vela cli
- `vela install` 安装`vela-core`组件

### Application
- `Application`包含了很多组件`components`
- `component` 层级结构 `workload` -> `trait`(属于`workload`属性)


## roadmap
- environment level only on namespace, cluster level will support

## demo
### install openfaas
`helm install openfaas openfaas/openfaas --set basic_auth=false --set operator.create=true --set functionNamespace=default`

### workflow
1. install dependency e.g. install openfaas crd
2. install workloaddefinition  e.g. `openfaas` (真实使用的是`workload`中的`template`)
3. install `application`  `vela-core`监听`application`并翻译`template`下发到k8s

### dashboard