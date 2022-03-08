package gorou

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserAccount struct {
	mutex   sync.Mutex
	Name    string
	Balance int
}

func (us *UserAccount) Lock() {
	us.mutex.Lock()
}

func (us *UserAccount) Unlock() {
	us.mutex.Unlock()
}

func (us *UserAccount) Change(amount int) {
	us.Balance += amount
}

func Transfer(from *UserAccount, to *UserAccount, amount int) {
	from.Lock()
	fmt.Println("Process 1...", from.Name)
	from.Change(-amount)

	time.Sleep(1 * time.Second)

	to.Lock()
	fmt.Println("Process 2...", to.Name)
	to.Change(amount)

	time.Sleep(1 * time.Second)

	from.Unlock()
	to.Unlock()
}

func TestTransferWithMutexAndDeadlock(t *testing.T) {
	acc1 := UserAccount{Name: "John", Balance: 10000}
	acc2 := UserAccount{Name: "Mary", Balance: 5000}

	go Transfer(&acc1, &acc2, 1000)
	go Transfer(&acc2, &acc1, 1000)

	time.Sleep(10 * time.Second)
	fmt.Println("Name 1: ", acc1.Name, "Balance: ", acc1.Balance)
	fmt.Println("Name 2: ", acc2.Name, "Balance: ", acc2.Balance)
}
