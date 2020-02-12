package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represents a number of Bitcoins
type Bitcoin int

func (b Bitcoin) String() (result string) {
	result = fmt.Sprintf("%d BTC", b)
	return
}

// Wallet stores the number of Bitcoin someone owns
type Wallet struct {
	balance Bitcoin
}

// Balance returns the number of Bitcoin of a wallet
func (w *Wallet) Balance() (result Bitcoin) {
	result = w.balance
	return
}

// Deposit will add some Bitcoin to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientBalance means a wallet does not have enough Bitcoint to perform a transaction
var ErrInsufficientBalance = errors.New("Insufficient balance")

// Withdraw takes some Bitcoin from a wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientBalance
	}

	w.balance -= amount
	return nil
}
