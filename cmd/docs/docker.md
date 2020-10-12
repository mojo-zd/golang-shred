docker实现主要依赖于linux的namespace、UnionFS、cgroups,可以实现网络、进程、挂载点、文件系统的隔离

- cmd、entrypoint区别
cmd命令可以被覆盖参数, 组合使用 entrypoint指定命令指定参数

- add、copy区别
copy仅仅是拷贝，add具备一些其他功能 eg:tar提取

- 什么是容器
提供运行应用程序所需的环境, 和其他容器共享内核，在主机操作系统的用户空间独立运行的进程

- 容器的运行状态
运行、已暂停、重新启动、已退出

- dockershim
