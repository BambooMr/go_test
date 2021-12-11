### 测试结果
```
redis-benchmark -d 5000 -t get,set

====== SET ======
  100000 requests completed in 1.13 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.50% <= 1 milliseconds
99.91% <= 2 milliseconds
99.95% <= 6 milliseconds
99.99% <= 7 milliseconds
100.00% <= 7 milliseconds
88573.96 requests per second

====== GET ======
  100000 requests completed in 1.08 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.80% <= 1 milliseconds
100.00% <= 1 milliseconds
92421.44 requests per second
```