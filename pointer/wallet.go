package main

import "fmt"
import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
    String() string
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return InsufficientFundsError
    }

    w.balance -= amount
    return nil
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}