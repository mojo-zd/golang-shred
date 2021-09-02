### tcpdump 入门
|  参数分类  | 常用关键字  |
|  ----  | ----  |
| 类型  | host、port |
| 方向  | src、dst |
| 协议  | ip、tcp、udp |
| 逻辑  | and、or、not |
| 其他  | -i 指定网卡eth0、eth1、tunl0、any，<br> -s 指定包长度,-s 0表示抓取完整包,<br> -c 指定抓包的数量, <br> -n 以数字方式显示ip、port <br> -A 以ASCII打印数据包 |

### 常见标志位
|  标志位  | 说明  |
|  ----  | ----  |
| SYN  | 发起连接 |
| ACK  | 确认连接 |
| PSH  | 可以直接将数据传给应用程序 |
| FIN  | 关闭连接 |
| RST  | 重置(异常关闭)连接 |

### 三次握手
![avatar](imgs/tcpshark.png)

### 关闭连接的四次挥手
![avatar](imgs/tcpendshark.png)

### 一次完整的交互实例
![avatar](imgs/iexample.png)

### RST
![avatar](imgs/rst.png)

### 关于tunl包
![avatar](imgs/tunl.png)

### 确定使用的是哪一个CGI
![avatar](imgs/witch.png)

### 同样的方法抓取referer
![avatar](imgs/referer.png)

### 关于QHTTP的QVIA字段
![avatar](imgs/qvia.png)

### 抓取RST
![avatar](imgs/rst_with.png)

### 抓取403、500等异常请求
![avatar](imgs/with_http_status.png)

### 统计返回码次数
![avatar](imgs/stat_code.png)

### 统计包请求来源
![avatar](imgs/package_from.png)


### 统计包请求域名
![avatar](imgs/stat_domain.png)