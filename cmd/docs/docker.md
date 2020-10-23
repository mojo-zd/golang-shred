docker实现主要依赖于linux的namespace、UnionFS、cgroups,可以实现网络、进程、挂载点、文件系统的隔离
#### cgroup限制资源使用(control groups), cgroups为每种可控制的资源定义一种子系统
1. 子系统 cpu、cpuacct、cpuset、memory、blkio、devices、net_cls、freezer
2. 容器利用/proc文件系统提供/proc/{pid}/root实现了为隔离后的容器提供单独的根文件系统
3. chroot 一个系统调用，将根目录改变到一个新的文件夹下并运行shell (为了安全可以作为一个沙箱使用)

#### namespaces隔离 
1. network
2. pid

- cmd、entrypoint区别
cmd命令可以被覆盖参数, 组合使用 entrypoint指定命令指定参数

- add、copy区别
copy仅仅是拷贝，add具备一些其他功能 eg:tar提取

- 什么是容器
提供运行应用程序所需的环境, 和其他容器共享内核，在主机操作系统的用户空间独立运行的进程

- 容器的运行状态
运行、已暂停、重新启动、已退出

- dockershim
(kubelet --> dockershim)(dockershim内嵌在kubelet中) --> dockershim接受请求转化为docker deamon能懂的请求(请求创建容器) -->
docker daemon运行了一个守护进程containerd，并通知containerd创建容器-->containerd shim真正执行容器的创建