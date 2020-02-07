package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() (result string) {
	result = fmt.Sprintf("%d BTC", b)
	return
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() (result Bitcoin) {
	result = w.balance
	return
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientBalance = errors.New("Insufficient balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientBalance
	}

	w.balance -= amount
	return nil
}