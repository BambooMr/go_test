### 测试结果
```
redis-benchmark -d 1000 -t get,set

====== SET ======
  100000 requests completed in 1.17 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.19% <= 1 milliseconds
99.96% <= 2 milliseconds
99.98% <= 3 milliseconds
100.00% <= 3 milliseconds
85689.80 requests per second

====== GET ======
  100000 requests completed in 1.06 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.81% <= 1 milliseconds
100.00% <= 1 milliseconds
94517.96 requests per second
```