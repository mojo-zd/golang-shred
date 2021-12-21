`golang`性能分析比较重要，一般的做法可以在应用程序中引入以下两个包

1.`net/http/pprof`,用于提供http访问方式

2.`runtime/pprof`

以1为例: http接口提供了`heap、allocs、profile、goroutine`等
```
http://localhost:8100/debug/pprof
```
使用方式:
```
go tool pprof http://localhost:8100/debug/pprof/profile?seconds=20
Fetching profile over HTTP from http://localhost:8100/debug/pprof/profile?seconds=5
Saved profile in /Users/mojo/pprof/pprof.metric-agent.samples.cpu.008.pb.gz
File: metric-agent
Type: cpu
Time: Dec 21, 2021 at 11:28am (CST)
Duration: 5s, Total samples = 0
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) svg //输出svg格式的文件
```
