# Concurrent data structure implementations and benchmarks in Golang

Implementations and benchmarks of concurrent data structures in Golang using
only the `sync.Mutex` concurrency control primitive (not even channels).

System: `cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz`

## Concurrent counter

Benchmark of incrementing counter to `N`. `PROC` goroutines are spun up, each incrementing to `N / PROC`

### Basic counter (no locks)
[Code](./counter/basic.go)
Lack of concurrency control causes incorrect counts.

Run benchmarks with
```
go test -bench=BenchmarkBasicCounter ./counter
```

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e5 | 209 us/op   |
| 2   | 1e5 | 495 us/op   |
| 4   | 1e5 | 442 us/op   |
| 8   | 1e5 | 437 us/op   |
| 16  | 1e5 | 565 us/op   |

N = 1e6
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e6 | 2054 us/op  |
| 2   | 1e6 | 4877 us/op  |
| 4   | 1e6 | 4266 us/op  |
| 8   | 1e6 | 4358 us/op  |
| 16  | 1e6 | 4482 us/op  |

### Lock counter
[Code](./counter/lock_counter.go)
Coarse-grained concurrency control results in accurate counts.

Run benchmarks with
```
go test -bench=BenchmarkLockCounter ./counter
```

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e5 | 1908  us/op  |
| 2   | 1e5 | 6105  us/op  |
| 4   | 1e5 | 5839  us/op  |
| 8   | 1e5 | 11964 us/op |
| 16  | 1e5 | 12809 us/op |

N = 1e6
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e6 | 22    ms/op |
| 2   | 1e6 | 44.6  ms/op |
| 4   | 1e6 | 82.5  ms/op |
| 8   | 1e6 | 106.8 ms/op |
| 16  | 1e6 | 113.3 ms/op |
