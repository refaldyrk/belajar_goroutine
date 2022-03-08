package gorou

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunSynchronous(wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	fmt.Println("Synchronous Goroutine")
	time.Sleep(1 * time.Second)
}

func TestWg(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		go RunSynchronous(wg)
	}

	wg.Wait()
	fmt.Println("All Goroutines Finished")
}

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i <= 100; i++ {
		go func() {
			wg.Add(1)
			once.Do(OnlyOnce)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func TestPool(t *testing.T) {
	var pool sync.Pool
	var wg sync.WaitGroup

	pool.New = func() interface{} {
		return "New"
	}

	pool.Put("Efal")
	pool.Put("Sayang")
	pool.Put("Delia")

	for i := 0; i <= 10; i++ {
		wg.Add(1)

		go func() {
			data := pool.Get()
			fmt.Println("Data: ", data)
			pool.Put(data)
			wg.Done()
		}()
	}

	wg.Wait()
}

func AddMap(mp *sync.Map, key interface{}, value interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	mp.Store(key, value)
}

func TestMap(t *testing.T) {
	smap := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		go AddMap(smap, i, i, wg)
	}

	wg.Wait()
	smap.Range(func(key, value interface{}) bool {
		fmt.Println("Key: ", key, " Value: ", value)
		return true
	})
}

func WaitCond(wg *sync.WaitGroup, cond *sync.Cond, val int) {
	defer wg.Done()

	wg.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Selesai ...", val)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	cond := sync.Cond{L: &sync.Mutex{}}
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		go WaitCond(&wg, &cond, i)
	}
	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	wg.Wait()
	time.Sleep(time.Second)
}
