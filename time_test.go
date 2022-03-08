package gorou

import (
	"fmt"
	"sync"
	"testing"
	time "time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	ch := <-timer.C
	fmt.Println(ch)
}

func TestAfter(t *testing.T) {
	timer := time.After(5 * time.Second)
	fmt.Println(time.Now())

	ch := <-timer
	fmt.Println(ch)
}

func TestAfterF(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("Start")
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("From After Func")
		wg.Done()
	})
	fmt.Println("End")

	wg.Wait()
}
