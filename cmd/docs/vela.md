## vela梳理
### vela cli
- `vela install` 安装`vela-core`组件

### Application
- `Application`包含了很多组件`components`
- `component` 层级结构 `workload` -> `trait`(属于`workload`属性)

## roadmap
- environment level only on namespace, cluster level will support
- storage only support local

## demo
### install openfaas
`helm install openfaas openfaas/openfaas --set basic_auth=false --set operator.create=true --set functionNamespace=default`

### workflow
1. install dependency e.g. install openfaas crd
2. install workloaddefinition  e.g. `openfaas` (真实使用的是`workload`中的`template`)
3. install `application`  `vela-core`监听`application`并翻译`template`下发到k8s

### dashboard
1. start apiserver
2. support a dashboard ui


### trait flow
1. create `TraitDefinition` crd schema
2. create `TraitDefinition` instance build-in trait  e.g. `metrics`、`route` (this trait includes it's own crd)
3. load workload witch you can load instance with `metrics` trait's `workloadRef` attr (`workloadRef` spec the `kind`、`appVersion`、`name`)

#### trait 作用
1. `trait` 定义模板
2. 通过`trait` `spec.definitionRef` 查找 `trait`的实现   e.g. `routes.standard.oam.dev`
3. 通过`routes` 中的`ownerReferences`可以查找到`ApplicationConfiguration` 
