# Concurrent data structure implementations and benchmarks in Golang

Implementations and benchmarks of concurrent data structures in Golang using
only the `sync.Mutex` concurrency control primitive (not even channels).

System: `cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz`

## Concurrent counter

Benchmark of incrementing counter to `N`. `PROC` goroutines are spun up, each incrementing to `N`

### Basic counter (no locks)
[Code](./counter/basic.go)
Lack of concurrency control causes incorrect counts.

Run benchmarks with
```
go test -v -bench=BenchmarkBasicCounter ./counter
```

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e5 | 319  us/op  |
| 2   | 1e5 | 1088 us/op  |
| 4   | 1e5 | 1740 us/op  |
| 8   | 1e5 | 3410 us/op  |
| 16  | 1e5 | 7594 us/op  |

N = 1e6
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e6 | 1.9  ms/op  |
| 2   | 1e6 | 10.9 ms/op  |
| 4   | 1e6 | 19.0 ms/op  |
| 8   | 1e6 | 36.1 ms/op  |
| 16  | 1e6 | 65.4 ms/op  |

### Lock counter
[Code](./counter/lock_counter.go)
Coarse-grained concurrency control results in accurate counts.

Run benchmarks with
```
go test -v -bench=BenchmarkLockCounter ./counter
```

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e5 | 2     ms/op |
| 2   | 1e5 | 7.8   ms/op |
| 4   | 1e5 | 25.9  ms/op |
| 8   | 1e5 | 76.4  ms/op |
| 16  | 1e5 | 192.7 ms/op |

N = 1e6
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e6 | 19   ms/op  |
| 2   | 1e6 | 74   ms/op  |
| 4   | 1e6 | 296  ms/op  |
| 8   | 1e6 | 754  ms/op  |
| 16  | 1e6 | 1795 ms/op  |

### Approx counter
[Code](./counter/approx_counter.go)
Maintains multiple local counters and periodically combines them into a global
counter.

Run benchmarks with
```
go test -v -bench=BenchmarkApproxCounter ./counter
```

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e5 | 4     ms/op |
| 2   | 1e5 | 12.1  ms/op |
| 4   | 1e5 | 26.8  ms/op |
| 8   | 1e5 | 70.1  ms/op |
| 16  | 1e5 | 135.3 ms/op |

N = 1e6
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e6 | 36   ms/op  |
| 2   | 1e6 | 109  ms/op  |
| 4   | 1e6 | 264  ms/op  |
| 8   | 1e6 | 637  ms/op  |
| 16  | 1e6 | 1359 ms/op  |

## Concurrent linked list

### Basic linked list (no locks)

Run benchmarks with
```
go test -v -bench=BenchmarkBasicLinkedList ./linkedlist
```

N = 1000
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1000 | 0.8  ms/op |
| 2   | 1000 | 1.7  ms/op |
| 4   | 1000 | 4.9  ms/op |
| 8   | 1000 | 14.9 ms/op |
| 16  | 1000 | 68.5 ms/op |

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e4 | 72.3  ms/op |
| 2   | 1e4 | 201.5 ms/op |
| 4   | 1e4 | 372.9 ms/op |
| 8   | 1e4 | 2440  ms/op |
| 16  | 1e4 | 4732  ms/op |

### Lock linked list

Run benchmarks with
```
go test -v -bench=BenchmarkLockLinkedList ./linkedlist
```

N = 1000
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1000 | 0.7  ms/op |
| 2   | 1000 | 3.0  ms/op |
| 4   | 1000 | 16.2 ms/op |
| 8   | 1000 | 58.9 ms/op |
| 16  | 1000 | 230 ms/op  |

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e4 | 78.1  ms/op |
| 2   | 1e4 | 342.3 ms/op |
| 4   | 1e4 | 1253 ms/op  |
| 8   | 1e4 | 6502  ms/op |
| 16  | 1e4 | 21185  ms/op |

### Couple lock linked list
Hand-over-hand locking with per-node locks

Run benchmarks with
```
go test -v -bench=BenchmarkCoupleLockLinkedList ./linkedlist
```

N = 1000
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1000 | 6.6  ms/op |
| 2   | 1000 | 25.5 ms/op |
| 4   | 1000 | 101  ms/op |
| 8   | 1000 | 255  ms/op |
| 16  | 1000 | 1132 ms/op |

N = 1e5
| CPU | N   | Performance |
|-----|-----|-------------|
| 1   | 1e4 | 1041  ms/op |
| 2   | 1e4 | 1706  ms/op |
| 4   | 1e4 | 7072  ms/op |
| 8   | 1e4 | 26399 ms/op |
| 16  | 1e4 | ???   ms/op |

## Concurrent queues

### Locked queue

Run benchmarks with
```
go test -v -bench=BenchmarkLockQueue ./queue
```

N = 1000
| CPU | N    | Performance | Push + Pop |
|-----|------|-------------|------------|
| 1   | 1000 | 0.3  ms/op  | 51  us/op  |
| 2   | 1000 | 1.2  ms/op  | 51  us/op  |
| 4   | 1000 | 4.1  ms/op  | 122 us/op  |
| 8   | 1000 | 10.1 ms/op  | 356 us/op  |
| 16  | 1000 | 20.1 ms/op  | 740 us/op  |

N = 1e5
| CPU | N   | Push        | Push + Pop  |
|-----|-----|-------------|-------------|
| 1   | 1e4 | 3.8  ms/op  | 345   us/op |
| 2   | 1e4 | 12.5 ms/op  | 826   us/op |
| 4   | 1e4 | 47.5 ms/op  | 1423  us/op |
| 8   | 1e4 | 90.2 ms/op  | 3487  us/op |
| 16  | 1e4 | 226  ms/op  | 10258 us/op |

