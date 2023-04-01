package scans

import (
	"fmt"
)

/*
* INTERACT FUNCTIONS SECTION
**/

func GetUser() (string, string, string, int) {
	var userName, cpf, profession string
	var accountType int

	fmt.Print("Insert your first name: ")
	fmt.Scan(&userName)
	fmt.Print("Insert your CPF: ")
	fmt.Scan(&cpf)
	fmt.Print("Insert your profession: ")
	fmt.Scan(&profession)

	fmt.Println("Select your account type: ")
	fmt.Println("1 - Economic ")
	fmt.Println("2 - Current ")
	fmt.Scan(&accountType)

	return userName, cpf, profession, accountType
}

func GetDepositValue() float64 {
	var depositValue float64
	fmt.Scan(&depositValue)
	fmt.Printf("\n-------------------------------------\n Deposit Value: %.2f R$ \n-------------------------------------\n\n", depositValue)
	return depositValue
}

func GetWithDrawalValue() float64 {
	var withdrawalValue float64
	fmt.Scan(&withdrawalValue)
	fmt.Printf("\n-------------------------------------------\n Withdrawal Value: %.2f R$ \n-------------------------------------------\n\n", withdrawalValue)
	return withdrawalValue
}

func GetTransferTo() string {
	var transferTo string
	fmt.Scan(&transferTo)
	fmt.Printf("\n-----------------------------------\n Transferring to: %s \n-----------------------------------\n\n", transferTo)
	return transferTo
}

func GetTransferValue() float64 {
	var transferValue float64
	fmt.Scan(&transferValue)
	fmt.Printf("\n----------------------------------------------------\n Inserted transfer value: %.2f R$ \n----------------------------------------------------\n\n", transferValue)
	return transferValue
}
