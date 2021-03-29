# 协议栈
## 链接层(mac层)
负责在以太网、WiFi这样的底层网络发送原始数据包，工作在网卡的层次，使用MAC地址标记网络设备，所以也叫MAC层

## 网际层(网络互联层)
IP协议就在这一层，IP协议定义了IP地址的概念，所以就可以在链接层的基础上使用IP地址代替MAC地址。IP地址可以把许许多多的设备链接到一个大的局域网，在网络里面寻找设备只需要把IP地址再翻译成MAC地址即可

## 传输层
这个层次协议负责保证数据在两个IP节点之间可靠传输，是TCP、UDP协议工作的层次
TCP为有状态协议，需要先建立连接再发送数据保证数据不重复不丢失。而UDP则比较简单，它是无状态的，不用事先建立连接可以任意发送数据，但不保证数据一定会发送到对方。两个协议的另一个重要的区别在于数据形式上。TCP的数据是连续的'字节流'，有先后顺序。UDP是分散的数据包，发送有序，接受乱序

## 应用层
Telnet、SSH、FTP、SMTP等等

# OSI模型

## 物理层
网络的物理形式，电缆、光缆、网卡、集线器等

## 数据链路层 
类似TCP/IP中的链接层

## 网络层
类似TCP/IP中的网际层

## 传输层
类似TCP/IP中的传输层

## 会话层
维护网络连接状态，即保持会话和同步

## 表示层
把数据转化为合适、可理解的语法和语义

## 应用层
面向具体的应用传输数据

# 四层负载均衡 七层负载均衡

## 负载均衡
1. 二层负载均衡 虚拟mac地址接受请求，再转发到真实mac地址
2. 三层负载均衡 虚拟ip地址接受请求, 再转发到真实ip地址
3. 四层负载均衡 虚拟ip+端口接受请求, 再转发到真实服务器地址
4. 七层负载均衡 虚拟URL或主机名接受请求, 再转发到真实服务器地址 (建立在四层之上)
## 四层负载均衡
ip+端口找到真实主机(通过mac地址或者相关算法)

## 七层负载均衡
利用URL等协议实现


## TCP
### 什么是协议栈
是一种特定的沟通的模式

### TCP如何保证包的顺序性
### TCP的拥塞控制
### 网断了,TCP怎么处理
### TCP的粘包与分包

> 网卡是将数字型号转换为光电信号

### 查看端口占用
- netstat
```
-t (tcp) 仅显示tcp相关选项
-u (udp)仅显示udp相关选项
-n 拒绝显示别名，能显示数字的全部转化为数字
-l 仅列出在Listen(监听)的服务状态
-p 显示建立相关链接的程序名
```

### 创建网络命名空间
```
ip netns add eden //创建命名空间
ip netns exec eden bash //进入网络命名空间
ip netns exec eden bash -rcfile <(echo "PS1=\"eden>\"")
```

### ip命令
#### 两个命名空间通讯
```
ip link //查看链路层信息
ip link set lo(device) down //关闭指定接口
// 连接两个netns
ip link add xi type veth peer name pan // 为两个ns创建共享工具
ip link set xi netns xi //把共享工具的一边搭建到xi的netns
ip link set pan netns pan //把共享工具的另一边搭建到pan的netns
ip netns exec xi ip addr add dev xi 192.168.188.96/24 //固定范围
ip netns exec pan ip addr add dev pan 192.168.188.96/24 //固定范围
ip netns exec xi ip link set xi up //启用接口
ip netns exec pan ip link set pan up //启用接口
```

#### 多个命名空间通讯
```
ip link add wangpo type bridge // 搭建桥
ip link add wp2xmq type veth peer name xmq2wp // 在桥上搭建梯子
ip link set xmq2wp netns xi // xi到wang的梯子放到xi的netns
ip link set wp2xmq master wangpo
ip netns exec xi ip addr add dev xi 192.168.188.96/24 //固定范围
ip netns exec xi ip link set xi up //启用接口
ip link set wp2xmq up // 启用接口
...
```

## 网络设备
### hub(集线器)
集线器工作在物理设备上(广播)

### Network bridge
网桥工作在数据链接层(通过mac地址寻址),不会进行广播。端口映射了多个mac地址,并未解决完全的广播风暴(学习过程中优化).

### Switch(交换机)
交换机工作在L2(数据链接层),一个网口对应一个mac地址.完全解决了广播风暴

### DHCP Server 
动态主机配置协议, 给客户机动态分配IP、网关、DNS

### NAT Device
工作在L3(网络层)