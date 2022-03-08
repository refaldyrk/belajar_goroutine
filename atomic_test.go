package gorou

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	var angka int64 = 0

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&angka, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Angka:", angka)
}
