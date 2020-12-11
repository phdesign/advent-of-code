package wallet

import (
    "fmt"
    "errors"
)

var InsufficientFundsError = errors.New("insufficient funds for withdrawl")

type Bitcoin float64

func (b Bitcoin) String() string {
    return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
    if amount > w.balance {
        return InsufficientFundsError
    }
    w.balance -= amount
    return nil
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

type Stringer interface {
    String() string
}

