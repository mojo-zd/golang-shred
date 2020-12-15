docker实现主要依赖于linux的namespace、UnionFS、cgroups,可以实现网络、进程、挂载点、文件系统的隔离
#### cgroup限制资源使用(control groups), cgroups为每种可控制的资源定义一种子系统
1. 子系统 cpu、cpuacct、cpuset、memory、blkio、devices、net_cls、freezer
2. 容器利用/proc文件系统提供/proc/{pid}/root实现了为隔离后的容器提供单独的根文件系统
3. chroot 一个系统调用，将根目录改变到一个新的文件夹下并运行shell (为了安全可以作为一个沙箱使用)

#### namespaces隔离 
1. network
2. pid

- cmd、entrypoint区别
cmd为默认执行方式，如果不指定entrypoint则会默认执行cmd，如果指定则以entrypoint为准(可以动态指定参数)。
cmd执行方式:
1. CMD ["executable", "param1", "param2"]
必须指定可执行进程
2. CMD ["param1","param2"] (as default parameters to ENTRYPOINT)
3. CMD command param1 param2 (shell form)
默认在 "/bin/sh -c"下执行
注: 一个dockerfile只能有一个cmd, 如果有多个只有最后一个生效

entrypoint执行方式:
如果run命令后有参数，所有内容都会作为entrypoint的参数传入。如果run后没有参数，只有cmd后面有，那么cmd的内容将会作为entrypoint的默认参数。--entrypoint可以覆盖dockerfile中默认的entrypoint
1. ENTRYPOINT ["executable", "param1", "param2"] (exec form, preferred)
这种方式默认是在shell下执行的
2. ENTRYPOINT command param1 param2 (shell form)
任何cmd或run参数都不会被传入到entrypoint。

- add、copy区别
copy仅仅是拷贝，add具备一些其他功能 eg:tar提取

- 什么是容器
提供运行应用程序所需的环境, 和其他容器共享内核，在主机操作系统的用户空间独立运行的进程

- 容器的运行状态
运行、已暂停、重新启动、已退出

- dockershim
(kubelet --> dockershim)(dockershim内嵌在kubelet中) --> dockershim接受请求转化为docker deamon能懂的请求(请求创建容器) -->
docker daemon运行了一个守护进程containerd，并通知containerd创建容器-->containerd shim真正执行容器的创建
参考: https://blog.kelu.org/tech/2020/10/09/the-diff-between-docker-containerd-runc-docker-shim.html