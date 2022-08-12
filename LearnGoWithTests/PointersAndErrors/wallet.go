package main

import "fmt"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g Bitcoin", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address in Deposit is %v \n", &(w.balance))
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address in Balance is %v \n", &(w.balance))
	return w.balance
}

func main() {}
