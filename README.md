# Concurrent data structure implementations and benchmarks in Golang

Implementations and benchmarks of concurrent data structures in Golang using
only the `sync.Mutex` concurrency control primitive (not even channels).

System: `cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz`

## Concurrent counter

Benchmark of incrementing counter to `N`. `PROC` goroutines are spun up, each incrementing to `N / PROC`

### Basic counter (no locks)
[Code](./counter/basic.go)
Lack of concurrency control causes incorrect counts.

Run benchmark with
```
go test -bench=. ./counter
```

N = 1e6
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e6 | 2054 us/op |
| 2   | 1e6 | 4877 us/op |
| 4   | 1e6 | 4266 us/op |
| 8   | 1e6 | 4358 us/op |
| 16  | 1e6 | 4482 us/op |

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e5 | 209 us/op |
| 2   | 1e5 | 495 us/op |
| 4   | 1e5 | 442 us/op |
| 8   | 1e5 | 437 us/op |
| 16  | 1e5 | 565 us/op |
