### 测试结果
```
redis-benchmark -d 10 -t get,set
====== SET ======
  100000 requests completed in 1.02 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.75% <= 1 milliseconds
99.90% <= 2 milliseconds
99.96% <= 3 milliseconds
100.00% <= 3 milliseconds
97751.71 requests per second

====== GET ======
  100000 requests completed in 1.00 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 1 milliseconds
99601.60 requests per second
```