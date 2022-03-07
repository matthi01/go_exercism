// Package account provides bank account functions
package account

import (
	"math"
	"sync"
)

// Account is a representation of a bank account
type Account struct {
	mu      sync.Mutex
	balance int64
	closed  bool
}

// Open opens and returns a new account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, closed: false}
}

// Deposit will add an amount to the account balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return a.balance, false
	}
	if amount < 0 && int64(math.Abs(float64(amount))) > a.balance {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}

// Balance will return the current account balance
func (a *Account) Balance() (balance int64, ok bool) {
	return a.balance, !a.closed
}

// Close will close the current account
func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return a.balance, false
	}
	remaining := a.balance
	a.balance = 0
	a.closed = true
	return remaining, true
}
