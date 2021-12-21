- 什么是挂载点

linux中的磁盘文件系统的入口目录，查看挂载点的方式:

```
// 查看所有挂载点, 也可以使用mount -l查看
df
Filesystem     1K-blocks     Used Available Use% Mounted on
overlay        206291912 34604180 163199240  18% /
tmpfs              65536        0     65536   0% /dev
tmpfs          162123940        0 162123940   0% /sys/fs/cgroup
/dev/vdb        30832548    45324  30770840   1% /data
/dev/vda1      206291912 34604180 163199240  18% /etc/hosts
shm                65536        0     65536   0% /dev/shm
tmpfs          162123940       12 162123928   1% /run/secrets/kubernetes.io/serviceaccount
tmpfs          162123940        0 162123940   0% /proc/acpi
tmpfs          162123940        0 162123940   0% /proc/scsi
tmpfs          162123940        0 162123940   0% /sys/firmware

// 查看指定目录所属挂载点
[root@computingnode-0 /data/files]# pwd
/data/files
[root@computingnode-0 /data/files]# df .
Filesystem     1K-blocks  Used Available Use% Mounted on
/dev/vdb        30832548 45324  30770840   1% /data
```

- 查看分区

```
// 查看分区列表
[root@computingnode-0 /data/files]# lsblk
NAME   MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT
vdd    251:48   0   100G  0 disk
vdb    251:16   0    30G  0 disk /data
sr0     11:0    1 133.2M  0 rom
vdc    251:32   0   100G  0 disk
vda    251:0    0   200G  0 disk
`-vda1 251:1    0   200G  0 part /etc/hosts
```

- 查看磁盘io

```
// 如果%util接近100%,表明i/o请求太多, i/o系统已经满负荷，磁盘可能存在瓶颈。
[root@computingnode-0 /data/files]# iostat -x 1 10 
Linux 4.14.105-19-0018_virtblk_test_v1 (computingnode-0) 	12/21/21 	_x86_64_	(80 CPU)

avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.09    0.00    0.14    0.00    0.00   99.77

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
loop0             0.00     0.00    0.00    0.00     0.00     0.00    50.33     0.00    1.67    1.67    0.00   0.28   0.00
scd0              0.00     0.00    0.00    0.00     0.03     0.00    80.27     0.00    2.20    2.20    0.00   1.38   0.00
vda               0.00     5.08    0.01    4.93     0.11   101.08    41.00     0.01    1.55    1.27    1.55   0.03   0.01
vdb               0.00     0.02    0.00    0.01     0.00     0.20    65.12     0.00    0.37    0.49    0.37   0.03   0.00
vdc               0.00     0.01    0.00    0.00     0.00     0.39   442.03     0.00    2.10    0.62    2.13   0.39   0.00
vdd               0.00     0.01    0.00    0.00     0.00     0.39   441.87     0.00    2.03    0.44    2.06   0.38   0.00


// 查看指定设备io
iostat -x 1 100 vdc
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.11    0.00    0.11    0.00    0.00   99.78

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vdc               0.00     0.00    0.00    0.00     0.00     0.00     0.00     0.00    0.00    0.00    0.00   0.00   0.00

avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.06    0.00    0.09    0.01    0.00   99.85

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vdc               0.00     0.00    0.00    0.00     0.00     0.00     0.00     0.00    0.00    0.00    0.00   0.00   0.00
```

- io测试

```
// 写入性能测试
time dd if=/dev/zero of=test.file bs=1G count=2 oflag=direct

// 读取性能测试
dd if=test.file of=/dev/null  iflag=direct
```