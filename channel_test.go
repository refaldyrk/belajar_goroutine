package gorou

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello World!"
		fmt.Println("Selesai ...")
	}()

	fmt.Println(<-ch)
}

func GiveMeChanPlease(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "I Love You!"
	fmt.Println("Selesai...")
}

func TestChanAsParam(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go GiveMeChanPlease(ch)
	fmt.Println(<-ch)
}

func OnlyIn(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Love You!"
	fmt.Println("Selesai ...")
}

func OnlyOut(ch <-chan string) {
	fmt.Println(<-ch)
}

func TestChanInOut(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go OnlyIn(ch)
	go OnlyOut(ch)

	time.Sleep(3 * time.Second)
}

type User struct {
	Name  string
	Age   int
	Email string
}

func SendData(ch chan<- User) {
	time.Sleep(2 * time.Second)
	ch <- User{Name: "Budi", Age: 20, Email: "budi@email.com"}
	fmt.Println("Selesai ...")
}

func ReceiveData(ch <-chan User) {
	data := <-ch
	fmt.Println("Data:", data)
}

func TestStructChan(t *testing.T) {
	ch := make(chan User)
	defer close(ch)

	go SendData(ch)
	go ReceiveData(ch)

	time.Sleep(3 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Value:", v)
	}

	fmt.Println("Selesai ...")
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go GiveMeChanPlease(ch1)
	go GiveMeChanPlease(ch2)

	counter := 0
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("Value 1:", v1)
			counter++
		case v2 := <-ch2:
			fmt.Println("Value 2:", v2)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChan(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	defer close(ch1)
	defer close(ch2)

	go GiveMeChanPlease(ch1)
	go OnlyIn(ch2)

	counter := 0
	for {
		select {
		case d := <-ch1:
			fmt.Println("Value 1:", d)
			counter++
		case d := <-ch2:
			fmt.Println("Value 2:", d)
			counter++
		default:
			fmt.Println("Waiting ...")
			time.Sleep(2 * time.Second)

		}
		if counter == 2 {
			break
		}
	}
}
