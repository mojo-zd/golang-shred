#### helm chart相关信息获取
首先添加为host执行添加repo命令`helm repo add bitnami https://charts.bitnami.com/bitnami`, 下面以`bitnami/apache`为例
1. helm说明获取
```
helm show readme bitnami/apache
```

2. 获取所有chart version
```
helm search repo bitnami/apache -l
```

3. 获取指定chart的value.yaml
```
helm show values bitnami/apache --version 7.6.0
```

4. 获取chart信息 e.g. Home url、Maintainers、Related
```
helm show chart bitnami/apache --version 0.3.8
```

5. 查看历史版本
```
helm history repo
```

6. 安装信息
通过release的chart.Value获取

7. deploy资源信息
直接调用kubernetes api

8. 部署时如果指定的名称跟chart名称相同则部署的deploy为部署名称 否则为填写的名称+chart名称 (helm本身的行为)

9. 具体需要获取那些资源
通过`helm get manifest repo` 获取到资源列表  过滤条件为 `app.kubernetes.io/instance={releaseName}`

10. 一些额外的信息可以存放在chart的annotation中

#### helm 内置函数
helm中除了包含go模板内置函数还包含sprig函数
1. default e.g. {{ default value  give_value}}
2. if/else
3. define // 定义新的命名模板
4. template // 导入模板  `{{ template "xx.xxx" .}}` 后面的`.`标识作用域
5. with // 指定范围 e.g. 使用`.Values.corese`的`k8s`字段可以用`.Values.corese.k8s` 
如果使用`with`则可以`{{ with .Values.corese }}` `{{ .k8s }}`
6. range // 循环块
7. block // 生命一个特殊的可填写的模块区域
8. 空格控制 `{{- ` 空格左移, ` -}}`表示删除右边的空格,这种需要慎用可能会删除右侧的换行符导致两行合并为一行
9. title // 把首字母变为大写
10. 变量申明 // 申明: `{{- $releaseName := .Release.Name -}}` 使用: `{{ $releaseName }}`
11. now // 获取当前时间
12. include // 引入模板文件 可以通过 `indent count`控制空格
13. index .arr 2
 
#### 提取helm公共包
1. 缓存index文件
2. 缓存chart.tar.gz
3. 提供helm相关操作入口
4. add repo更新index file