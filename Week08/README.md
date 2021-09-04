# Go进阶训练营第8周作业

## 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。 

### 命令:

redis-benchmark -d 10 -t get,set

### SET：

| -    | 执行次数和耗时                            | 每秒请求次数                  |
| ---- | ----------------------------------------- | ----------------------------- |
| 10   | 100000 requests completed in 0.85 seconds | 117096.02 requests per second |
| 20   | 100000 requests completed in 0.76 seconds | 132100.39 requests per second |
| 50   | 100000 requests completed in 0.80 seconds | 124378.11 requests per second |
| 100  | 100000 requests completed in 0.81 seconds | 124069.48 requests per second |
| 200  | 100000 requests completed in 0.80 seconds | 124843.95 requests per second |
| 1k   | 100000 requests completed in 0.80 seconds | 125156.45 requests per second |
| 5k   | 100000 requests completed in 0.92 seconds | 108577.63 requests per second |

### GET：

| -    | 执行次数和耗时                            | 每秒请求次数                  |
| ---- | ----------------------------------------- | ----------------------------- |
| 10   | 100000 requests completed in 0.82seconds  | 121654.50 requests per second |
| 20   | 100000 requests completed in 0.78 seconds | 127551.02 requests per second |
| 50   | 100000 requests completed in 0.88 seconds | 114285.71 requests per second |
| 100  | 100000 requests completed in 0.88 seconds | 113507.38 requests per second |
| 200  | 100000 requests completed in 0.98 seconds | 102354.15 requests per second |
| 1k   | 100000 requests completed in 0.95 seconds | 104821.80 requests per second |
| 5k   | 100000 requests completed in 1.02 seconds | 98039.22  requests per second |

# 2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

## 原理：

写入不同数量不同长度的value, 分析内存占用, 导出结果到csv文件 

## 结论：

相同长度的value在写入数量越多情况下，平均每个value占用内存更多

