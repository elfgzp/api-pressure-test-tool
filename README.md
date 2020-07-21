# api-pressure-test-tool
api-pressure-test-tool 压测工具（架构师训练营作业）

## 使用方法

```bash
$ go run cmd/main.go -u https://www.baidu.com -t 100 -c 10 -p 0.95
...
并发序号 3 第 11 次请求，耗时 42 ms 
并发序号 6 第 11 次请求，耗时 42 ms 
并发序号 4 第 11 次请求，耗时 47 ms 
并发序号 8 第 12 次请求，耗时 41 ms 
并发序号 2 第 12 次请求，耗时 49 ms 
并发序号 0 第 12 次请求，耗时 49 ms 
平均请求耗时 58 ms 
百分位 0.950000 耗时 186 ms 

```
