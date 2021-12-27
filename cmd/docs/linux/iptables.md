iptables用于控制linux下的网络访问策略。 基础用法：

INPUT链操作

```
// 添加一条INPUT规则, 拒绝通过127.0.0.1访问
iptables -A INPUT -p tcp -s 127.0.0.1 -j DROP

// 拒绝访问127.0.0.1:8888
iptables -A INPUT -p tcp -s 127.0.0.1 --sport 8888 -j DROP

// 查看INPUT规则
iptables -nvL INPUT --line-number

// 删除一条规则
iptables -D INPUT {line-number}

iptables -D INPUT -p tcp -s 127.0.0.1 --sport 8888 -j DROP
```

添加自定义链

```
// 在INPUT链中添加自定义链
iptables -N KUBE-ROUTER-INPUT
iptables -A INPUT   -m comment --comment "kube-router netpol" -j KUBE-ROUTER-INPUT

// 在自定义链中添加规则
iptables -A KUBE-ROUTER-INPUT -p tcp -s 127.0.0.1 --sport 8888 -j DROP
```