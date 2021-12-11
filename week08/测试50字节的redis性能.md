### 测试结果
```
redis-benchmark -d 50 -t get,set

====== SET ======
  100000 requests completed in 1.04 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

99.68% <= 1 milliseconds
100.00% <= 1 milliseconds
96339.12 requests per second

====== GET ======
  100000 requests completed in 1.02 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

99.93% <= 1 milliseconds
100.00% <= 1 milliseconds
97943.19 requests per second
```