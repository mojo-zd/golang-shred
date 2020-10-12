#### ws分布式应用
通常情况下ws server具备多实例特性以实现高可用, 要想client能准确接受到所有的ws server消息需要处理分布式情况。需要借助第三方工具 eg: mq、redis。下面以mq为例

1、业务发送消息到mq(发布)

2、ws server从mq订阅消息消息

3、ws server把消息推送给ws client
