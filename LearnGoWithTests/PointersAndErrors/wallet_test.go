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

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("Expected Error, but did not get any!!")
		}

		if got.Error() != want {
			t.Errorf("got %s want %s", got.Error(), want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("got %s expected nil", got.Error())
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

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet.Balance(), startingBalance)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.0)}
		err := wallet.Withdraw(Bitcoin(10.0))

		assertNoError(t, err)
		assertBalance(t, wallet.Balance(), Bitcoin(10.0))
	})

}

func TestWithDraw(t *testing.T) {

}
