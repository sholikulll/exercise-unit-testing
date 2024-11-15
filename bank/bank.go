package bank

import "errors"

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount cannot be less than equal zero")
	}

	a.Balance += amount

	return nil
}
