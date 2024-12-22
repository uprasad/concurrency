package counter_test

import (
	"concurrency/counter"
	"fmt"
	"sync"
	"testing"
)

var (
	parallelisms = []int{1, 2, 4, 8, 16}
)

func doIncs(cnt counter.Counter, incs int) {
	for i := 0; i < incs; i++ {
		cnt.Inc()
	}
}

func benchmarkCounter(b *testing.B, cnt counter.Counter, parallelism int, limit int) {
	for i := 0; i < b.N; i++ {
		cnt.Reset()
		// Kick off goroutines that increment in parallel
		var wg sync.WaitGroup
		wg.Add(parallelism)
		for j := 0; j < parallelism; j++ {
			go func() {
				defer wg.Done()
				doIncs(cnt, limit)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkBasicCounter_100_000(b *testing.B) {
	for _, p := range parallelisms {
		b.Run(fmt.Sprintf("%d", p), func(b *testing.B) {
			cnt := counter.NewBasicCounter()
			benchmarkCounter(b, cnt, p, 100_000)
			fmt.Printf("expected count: %d, got %d\n", 100_000*p, cnt.Get())
		})
	}
}

func BenchmarkBasicCounter_1_000_000(b *testing.B) {
	for _, p := range parallelisms {
		b.Run(fmt.Sprintf("%d", p), func(b *testing.B) {
			cnt := counter.NewBasicCounter()
			benchmarkCounter(b, cnt, p, 1_000_000)
			fmt.Printf("expected count: %d, got %d\n", 1_000_000*p, cnt.Get())
		})
	}
}

func BenchmarkLockCounter_100_000(b *testing.B) {
	for _, p := range parallelisms {
		b.Run(fmt.Sprintf("%d", p), func(b *testing.B) {
			cnt := counter.NewLockCounter()
			benchmarkCounter(b, cnt, p, 100_000)
			fmt.Printf("expected count: %d, got %d\n", 100_000*p, cnt.Get())
		})
	}
}

func BenchmarkLockCounter_1_000_000(b *testing.B) {
	for _, p := range parallelisms {
		b.Run(fmt.Sprintf("%d", p), func(b *testing.B) {
			cnt := counter.NewLockCounter()
			benchmarkCounter(b, cnt, p, 1_000_000)
			fmt.Printf("expected count: %d, got %d\n", 1_000_000(*p, cnt.Get())
		})
	}
}
