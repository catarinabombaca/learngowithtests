package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (btc Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", btc)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= value
	return nil
}
