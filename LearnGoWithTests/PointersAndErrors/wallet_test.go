package main

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	wallet := Wallet{Bitcoin(0.0)}
	wallet.Deposit(Bitcoin(10.0))

	fmt.Printf("address in Test is %v \n", &wallet.balance)
	got := wallet.Balance()
	want := Bitcoin(12.0)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
