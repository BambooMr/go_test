### 测试结果
```
redis-benchmark -d 200 -t get,set

====== SET ======
  100000 requests completed in 1.05 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.65% <= 1 milliseconds
100.00% <= 1 milliseconds
95057.03 requests per second

====== GET ======
  100000 requests completed in 1.04 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.84% <= 1 milliseconds
100.00% <= 1 milliseconds
96432.02 requests per second
```