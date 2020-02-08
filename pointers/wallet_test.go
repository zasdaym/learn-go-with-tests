package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientBalance)
	})
}