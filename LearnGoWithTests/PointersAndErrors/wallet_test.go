package main

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	wallet := Wallet{0.0}
	wallet.Deposit(10.0)

	fmt.Printf("address in Test is %v \n", &wallet.balance)
	got := wallet.Balance()
	want := 10.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
