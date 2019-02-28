package kata

import "errors"

type Account struct {
	Balance float64
}

type ATM interface {
	GetAccount() (Account, error)
	GetBalance() float64
	Deposit() error
}

func (a *Account) GetAccount() (*Account, error) {
	return a, nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) Deposit(amount float64) {
	newBalance := a.Balance + amount
	a.Balance = newBalance
}

func (a *Account) Withdraw(amount float64) error {
	if hasEnoughBalance(amount, a.Balance) {
		return errors.New("not enough money to withdraw")
	}
	newBalance := a.Balance - amount
	a.Balance = newBalance
	return nil
}

func hasEnoughBalance(amount float64, balance float64) bool {
	return amount > balance
}
