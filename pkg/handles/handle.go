package handles

import (
	"bank/pkg/accounts"
	"bank/pkg/clients"
	"bank/pkg/scans"
	"bank/pkg/utils"
	"fmt"
	"math/rand"
	"time"
)

func HandleWithdrawal(account *accounts.CurrentAccount, typeAccount string) {
	utils.ClearTerminal()
	fmt.Println("")
	fmt.Println("//////////////")
	fmt.Println("//Withdrawal//")
	fmt.Println("//////////////")
	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| ", typeAccount, "  ")
	fmt.Println("╰-----------------")
	fmt.Println("")
	fmt.Println("Enter the desired withdrawal amount:")
	fmt.Println("")

	withdrawalAmount := scans.GetWithDrawalValue()
	response := account.Withdrawal(withdrawalAmount)

	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| ", response, "  ")
	fmt.Println("╰-----------------")
	fmt.Println("")

	utils.Log("Withdrawal", typeAccount, fmt.Sprintf("%.2f", withdrawalAmount), account.Holder.Name, "", response)

}

func HandleDeposit(account *accounts.CurrentAccount, typeAccount string) {
	utils.ClearTerminal()
	fmt.Println("")
	fmt.Println("//////////////")
	fmt.Println("///Deposits///")
	fmt.Println("//////////////")
	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| ", typeAccount, "  ")
	fmt.Println("╰-----------------")
	fmt.Println("")
	fmt.Println("Insert the value of deposit:")
	fmt.Println("")
	depositValue := scans.GetDepositValue()
	response, _ := account.Deposit(depositValue)

	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| ", response, "  ")
	fmt.Println("╰-----------------")
	fmt.Println("")

	utils.Log("Deposit", typeAccount, fmt.Sprintf("%.2f", depositValue), account.Holder.Name, "", response)
}

func HandleTransfer(account *accounts.CurrentAccount, typeAccount string) {
	utils.ClearTerminal()
	fmt.Println("")
	fmt.Println("//////////////")
	fmt.Println("///Transfer///")
	fmt.Println("//////////////")
	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| ", typeAccount, "  ")
	fmt.Println("╰-----------------")
	fmt.Println("")
	fmt.Println("transfer to (name):")
	fmt.Println("")
	transferFor := scans.GetTransferTo()

	rand.Seed(time.Now().UnixNano())

	agencyNumber := rand.Intn(1000)
	accountNumber := rand.Intn(1000)

	secondClient := clients.Holder{
		Name: transferFor,
	}
	secondAccount := accounts.CurrentAccount{
		Holder:        secondClient,
		AgencyNumber:  agencyNumber,
		AccountNumber: accountNumber,
	}

	fmt.Println("")
	fmt.Println("Transference value:")
	fmt.Println("")

	transferVal := scans.GetTransferValue()

	response, status := account.Transfer(transferVal, &secondAccount)

	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| ", response, "  ")
	fmt.Println("╰-----------------")
	fmt.Println("")

	if status {
		fmt.Println("")
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println(" Transfered to user:", secondClient.Name, " | value of:", transferVal, " ")
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("")
	}

	utils.Log("Transfer", typeAccount, fmt.Sprintf("%.2f", transferVal), account.Holder.Name, secondClient.Name, response)
}

func HandleBalance(account *accounts.CurrentAccount, userName string, typeAccount string) {
	utils.ClearTerminal()
	fmt.Println("")
	fmt.Println("//////////////")
	fmt.Println("///Balances///")
	fmt.Println("//////////////")
	fmt.Println("")
	fmt.Println("------------------------------")
	fmt.Println(" Account User:", userName, " ")
	fmt.Println("------------------------------")
	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| ", typeAccount, "  ")
	fmt.Println("╰-----------------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("╭-----------------")
	fmt.Println("| Total:", account.GetSale(), "R$")
	fmt.Println("╰-----------------")
	fmt.Println("")
}

func HandleMyAccount(account *accounts.CurrentAccount, userName string, cpf string, profession string, agencyNumber int, accountNumber int, typeAccount string) {
	utils.ClearTerminal()

	fmt.Println("")
	fmt.Println("/////////////////")
	fmt.Println("///My Account///")
	fmt.Println("/////////////////")
	fmt.Println("")

	fmt.Println("")
	fmt.Println("╭---------------------------------")
	fmt.Println(" My Personal Account -------------")
	fmt.Println("|")

	fmt.Println("╭-----------------")
	fmt.Println("| ", typeAccount, " ")
	fmt.Println("╰-----------------")

	fmt.Println("|")
	fmt.Println("| Name:", userName, "  ")
	fmt.Println("| CPF:", cpf, "  ")
	fmt.Println("| Profession:", profession, "  ")
	fmt.Println("| Agency Number:", agencyNumber, "  ")
	fmt.Println("| Account Number:", accountNumber, "  ")
	fmt.Println("╰---------------------------------")
	fmt.Println("╭-----------------")
	fmt.Println("| Total:", account.GetSale(), "R$")
	fmt.Println("╰-----------------")
	fmt.Println("")
}

func HandleShowLogs() {
	utils.ClearTerminal()

	fmt.Println("")
	fmt.Println("/////////////////")
	fmt.Println("///System Logs///")
	fmt.Println("/////////////////")
	fmt.Println("")
	fmt.Println("---------------------------------")
	fmt.Println("")

	utils.ReadLogs()

	fmt.Println("")
	fmt.Println("---------------------------------")
}
