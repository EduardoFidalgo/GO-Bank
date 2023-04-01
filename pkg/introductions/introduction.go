package introductions

import "fmt"

/*
* INTRODUCTION FUNCTIONS SECTION
**/

func Introduction() {
	var version float32 = 1.1
	fmt.Println("System version:", version)
	fmt.Println("")
	fmt.Println("/////////////////////////////////")
	fmt.Println("*/ Welcome to GO bank project /*")
	fmt.Println("////////////////////////////////")
}

func ShowMenu() {
	fmt.Println("1 - Withdrawal")
	fmt.Println("2 - Deposit")
	fmt.Println("3 - Transfer")
	fmt.Println("4 - Balance")
	fmt.Println("5 - My Account")
	fmt.Println("6 - System Logs (For Nerds)")
	fmt.Println("0 - Destroy")
	fmt.Println("")
}
