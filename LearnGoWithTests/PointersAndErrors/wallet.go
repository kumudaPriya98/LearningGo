package main

import "fmt"
import "errors"

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g Bitcoin", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {}
