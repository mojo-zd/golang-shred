### openstack组件
- horizon (控制台)
- nova (计算)
用于创建、调度、销毁云主机
- neutron (网络)
负责实现sdn(software define network)
- swift (对象存储)
- cinder (块存储)
- glance (镜像)
提供镜像服务，装机使用
- keystone (认证)
为访问openstack各个组件提供认证授权服务, 认证通过以后提供一个服务列表
- ceilometer (计费)
- heat 编排

#### 基本概念
- domain
包含project user
- group 
一个domain内的一组user,为了方便分配role
- project
是一组资源的集合
- user
- token
分为`unscoped token`、`domain-scoped token`、`project-scoped token`
- role
role可以简单分为admin、member, admin又包括`cloud admin`、`domain admin`、`project admin`

#### keystone


### openstack api list
- https://docs.openstack.org/zh_CN/api-quick-start/

### openstack identity link
- https://docs.openstack.org/api-ref/identity/v3/?expanded=password-authentication-with-unscoped-authorization-detail,password-authentication-with-scoped-authorization-detail#authentication-and-token-management

#### 获取token
curl -i -X POST http://xx.xx.xx.xx:35357/v3/auth/tokens -H "Content-type: application/json" -d '{"auth": {"identity": {"methods": ["password"],"password": {"user": {"name": "admin","domain": {"name": "Default"},"password": "lSvX2AukMMii5vpaBKlAEBzho18Vkbn8l7FRhwIC"}}}}}'|grep X-Subject-Token

#### 验证token
curl -i http://xx.xx.xx.xx:35357/v3/auth/tokens -H "Content-type: application/json" -H "X-Auth-Token: gAAAAABfxJIhzZ_abIauM_cVBkHU3R9HTWzmDew8DvIN_bmH5F7Mnhdun6opJDamH6B_Xm8uGv7oojbO8ahpdWiWrHsB6QOWL9-WdPhDfOQjJYu7XF8o_VpWEKZpoKcKNRXqxUBpn8vtECXHbBpUzBy8Gf_nzukTJg" -H "X-Subject-Token: gAAAAABfxJIhzZ_abIauM_cVBkHU3R9HTWzmDew8DvIN_bmH5F7Mnhdun6opJDamH6B_Xm8uGv7oojbO8ahpdWiWrHsB6QOWL9-WdPhDfOQjJYu7XF8o_VpWEKZpoKcKNRXqxUBpn8vtECXHbBpUzBy8Gf_nzukTJg"

#### container api (include magnum)
- https://docs.openstack.org/api-ref/container-infrastructure-management/

集群管理、镜像中心、监控、编排、helm jaeger istio集成