package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0.0)}
		wallet.Deposit(Bitcoin(10.0))
		assertBalance(t, wallet.Balance(), Bitcoin(10.0))
	})

	t.Run("Withdraw Insufficient Amount", func(t *testing.T) {
		startingBalance := Bitcoin(20.0)
		wallet := Wallet{balance: Bitcoin(20.0)}
		err := wallet.Withdraw(Bitcoin(100.0))
		
		assertBalance(t, wallet.Balance(), startingBalance)

		if err == nil {
			t.Errorf("Wanted an Error but didnt get one")
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.0)}
		wallet.Withdraw(Bitcoin(10.0))
		assertBalance(t, wallet.Balance(), Bitcoin(10.0))
	})

}

func TestWithDraw(t *testing.T) {

}
