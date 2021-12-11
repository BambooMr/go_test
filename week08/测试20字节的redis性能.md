### 测试结果
```
redis-benchmark -d 20 -t get,set

====== SET ======
  100000 requests completed in 1.10 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

99.57% <= 1 milliseconds
99.96% <= 2 milliseconds
100.00% <= 3 milliseconds
91157.70 requests per second

====== GET ======
  100000 requests completed in 1.10 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

99.75% <= 1 milliseconds
100.00% <= 1 milliseconds
90661.83 requests per second
```