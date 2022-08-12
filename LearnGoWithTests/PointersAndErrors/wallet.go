package main

import "fmt"

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	fmt.Printf("address in Deposit is %v \n", &(w.balance))
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	fmt.Printf("address in Balance is %v \n", &(w.balance))
	return w.balance
}

func main() {}
