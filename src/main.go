package main

import (
	"bank/pkg/accounts"
	"bank/pkg/clients"
	"bank/pkg/handles"
	"bank/pkg/introductions"
	"bank/pkg/scans"
	"bank/pkg/utils"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	WithdrawalOption     = 1
	DepositOption        = 2
	TransferOption       = 3
	BalanceOption        = 4
	PersonalAccounOption = 5
	showLogsOption       = 6
	ExitOption           = 0

	randMaxNumber = 1000

	CurrentAccountIdentification  = 1
	EconomicAccountIdentification = 2

	InvalidOptionMessage = "Invalid option. Please choose a valid option."
)

func main() {
	introductions.Introduction()

	fmt.Println("")
	userName, userCPF, userProfession, accountType := scans.GetUser()
	fmt.Println("")

	rand.Seed(time.Now().UnixNano())
	agencyNumber := rand.Intn(randMaxNumber)
	accountNumber := rand.Intn(randMaxNumber)

	firstClient := clients.Holder{
		Name:       userName,
		Cpf:        userCPF,
		Profession: userProfession,
	}

	economicAccount := accounts.CurrentAccount{
		Holder:        firstClient,
		AgencyNumber:  agencyNumber,
		AccountNumber: accountNumber,
	}

	currentAccount := accounts.CurrentAccount{
		Holder:        firstClient,
		AgencyNumber:  agencyNumber,
		AccountNumber: accountNumber,
	}

	fmt.Println("")
	fmt.Println("╭---------------------------------")
	fmt.Println(" Account Created -----------------")
	fmt.Println("")
	fmt.Println("| Name:", userName, "  ")
	fmt.Println("| CPF:", userCPF, "  ")
	fmt.Println("| Profession:", userProfession, "  ")
	fmt.Println("| Agency Number:", agencyNumber, "  ")
	fmt.Println("| Account Number:", accountNumber, "  ")
	fmt.Println("╰---------------------------------")
	fmt.Println("")

	if userName != "" {
		for {
			fmt.Println("")
			fmt.Println("Please, choose an option to continue.")
			fmt.Println("")

			introductions.ShowMenu()

			command := utils.ReceiveCommand()

			switch command {
			case WithdrawalOption:
				if accountType == CurrentAccountIdentification {
					handles.HandleWithdrawal(&economicAccount, "Current Account")
				} else if accountType == EconomicAccountIdentification {
					handles.HandleWithdrawal(&currentAccount, "Economic Account")
				} else {
					os.Exit(0)
				}
			case DepositOption:
				if accountType == CurrentAccountIdentification {
					handles.HandleDeposit(&economicAccount, "Current Account")
				} else if accountType == EconomicAccountIdentification {
					handles.HandleDeposit(&currentAccount, "Econimic Account")
				} else {
					os.Exit(0)
				}
			case TransferOption:
				if accountType == CurrentAccountIdentification {
					handles.HandleTransfer(&economicAccount, "Current Account")
				} else if accountType == EconomicAccountIdentification {
					handles.HandleTransfer(&currentAccount, "Economic Account")
				} else {
					os.Exit(0)
				}
			case BalanceOption:
				if accountType == CurrentAccountIdentification {
					handles.HandleBalance(&economicAccount, userName, "Current Account")
				} else if accountType == EconomicAccountIdentification {
					handles.HandleBalance(&currentAccount, userName, "Economic Account")
				} else {
					os.Exit(0)
				}
			case PersonalAccounOption:
				if accountType == CurrentAccountIdentification {
					handles.HandleMyAccount(&economicAccount, userName, userCPF, userProfession, agencyNumber, accountNumber, "Current Account")
				} else if accountType == EconomicAccountIdentification {
					handles.HandleMyAccount(&currentAccount, userName, userCPF, userProfession, agencyNumber, accountNumber, "Economic Account")
				} else {
					os.Exit(0)
				}
			case showLogsOption:
				handles.HandleShowLogs()
			case ExitOption:
				fmt.Println("")
				fmt.Println("Exiting.. 3, 2, 1...")
				fmt.Println("")
				os.Exit(0)
			default:
				fmt.Println("")
				fmt.Println("//////////////")
				fmt.Println("///NotFound///")
				fmt.Println("//////////////")
				fmt.Println("")
			}
		}
	}
}
