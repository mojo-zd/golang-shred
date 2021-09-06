工作中遇到的一些奇怪知识点
1. 如何实现一个kill tcp的功能
kill tcp需要通过fake一个tcp连接, 然后通过拦截之前的tcp获取seq,对连接进行rst操作。至于如何fake一个tcp可以通过FD获得，Linux系统中的一些皆为文件, 连接也不例外。只要找到文件句柄拿到引用就可以搞定
   
2. tcpdump的使用分析
对于一个http请求, 当进行三次握手以后。进行数据交换的过程中, client和server的[P.] (推送数据操作)和[.] (ack)操作会交替进行, 在请求中你会发现, http的请求是先返回http header在返回body的
   
3. docker的资源限制(cpu,memory)是怎样的
为什么我们对docker进行资源限定后,我们在docker的系统内看到的cpu核数是主机的核数。

4. 验证docker的cpu资源设定是否生效
e.g.
```
# 限制 docker的cpu的核数为2
docker run --rm -it --cpus=2 --memory=256m  --name test csighub.tencentyun.com/admin/tlinux2.2-bridge-tcloud-underlay bash
# yum install stress, 并执行下面的命令 打满cpu, 最多只能打满2核
stress --cpu 4 
```