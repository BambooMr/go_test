### 测试结果
```
redis-benchmark -d 100 -t get,set

====== SET ======
  100000 requests completed in 1.18 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.18% <= 1 milliseconds
99.97% <= 2 milliseconds
100.00% <= 2 milliseconds
85034.02 requests per second

====== GET ======
  100000 requests completed in 1.14 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.54% <= 1 milliseconds
100.00% <= 1 milliseconds
87796.30 requests per second
```