package accounts

import (
	"bank/pkg/clients"
)

type EconomicAccount struct {
	Holder                                 clients.Holder
	AgencyNumber, Operation, AccountNumber int
	sale                                   float64
}

func (c *EconomicAccount) Withdrawal(withdrawalValue float64) string {
	validWithdrawal := withdrawalValue > 0 && withdrawalValue <= c.sale

	if validWithdrawal {
		c.sale -= float64(withdrawalValue)
		return "Withdrawal made successfully!"
	} else {
		return "Insufficient funds"
	}
}

func (c *EconomicAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.sale += depositValue
		return "Deposit made successfully!", c.sale
	} else {
		return "Deposit amount must be greater than 0", c.sale
	}
}

func (c *EconomicAccount) GetSale() float64 {
	return c.sale
}
