package gorou

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	rwMutex sync.RWMutex
	balance int
}

func (acc *BankAccount) AddAmount(amount int) {
	acc.rwMutex.Lock()
	defer acc.rwMutex.Unlock()
	acc.balance += amount
}

func (acc *BankAccount) GetBalance() int {
	acc.rwMutex.RLock()
	balance := acc.balance
	acc.rwMutex.RUnlock()
	return balance
}

func TestRwMutex(t *testing.T) {
	rek := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				rek.AddAmount(1)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance: ", rek.GetBalance())
}
