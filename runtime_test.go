package gorou

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestCpu(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			wg.Done()
		}()
	}

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(-1))
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

}
