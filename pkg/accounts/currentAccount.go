package accounts

import (
	"bank/pkg/clients"
)

type CurrentAccount struct {
	Holder        clients.Holder
	AgencyNumber  int
	AccountNumber int
	sale          float64
}

func (c *CurrentAccount) Withdrawal(withdrawalValue float64) string {
	validWithdrawal := withdrawalValue > 0 && withdrawalValue <= c.sale

	if validWithdrawal {
		c.sale -= float64(withdrawalValue)
		return "Withdrawal made successfully!"
	} else {
		return "Failed, insufficient funds!"
	}
}

func (c *CurrentAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.sale += depositValue
		return "Deposit made successfully!", c.sale
	} else {
		return "Deposit amount must be greater than R$00.00", c.sale
	}
}

func (c *CurrentAccount) Transfer(transferValue float64, destinationAccount *CurrentAccount) (string, bool) {
	if transferValue < c.sale && transferValue > 0 {
		c.sale -= transferValue
		destinationAccount.Deposit(transferValue)
		return "Transfer Concluded!", true
	} else {
		return "Failed, insufficient funds!", false
	}
}

func (c *CurrentAccount) GetSale() float64 {
	return c.sale
}
