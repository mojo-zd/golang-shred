linux中`namespace`用于做资源隔离, 隔离的资源类型包含以下6类

- net
- ipc (InterProcess Communication,进程间的通信机制。通过消息队列、共享存储、信号量)
- pid
- uts
- user
- mnt

查看linux中的ns可以使用`lsns`查看。

```
lsns // 查看所有ns
lsns -t net // 查看所有net ns
lsns -t net -p 32574 // 查看指定进程ns

nsenter -t 32574 -n ip a // 查看进程的所在net ns中的网卡信息 
```
