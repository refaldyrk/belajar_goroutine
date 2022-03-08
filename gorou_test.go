package gorou

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorldGorou() {
	fmt.Println("Hello World!")
}

func TestGorou(t *testing.T) {
	go HelloWorldGorou()
	fmt.Println("Ini Bukan Dari Goroutine")

	time.Sleep(1 * time.Second)
}

func DisplayNum(value int) {
	fmt.Println("Value:", value)
}

func TestManyGorou(t *testing.T) {
	for i := 0; i < 10; i++ {
		go DisplayNum(i)
	}

	time.Sleep(10 * time.Second)
}
