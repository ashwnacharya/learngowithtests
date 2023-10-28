package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &wallet.balance)
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

var ErrInsufficientFunds = errors.New("can't withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error {

	if wallet.balance < amount {
		return ErrInsufficientFunds
	}

	wallet.balance -= amount
	return nil
}