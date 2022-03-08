package gorou

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())

		case <-time.After(10 * time.Second):
			fmt.Println("Timeout")
			return
		default:
			fmt.Println("Menunggu...")
			time.Sleep(5 * time.Second)
		}
	}
}
