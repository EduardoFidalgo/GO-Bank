package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ReceiveCommand() int {
	var receivedCommand int
	fmt.Scan(&receivedCommand)
	return receivedCommand
}

func Log(action string, accountType string, val string, user, transferFor string, response string) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	var transferUser string
	if transferFor != "" {
		transferUser = "\n - Transfer for: " + transferFor
	} else {
		transferUser = ""
	}

	var txt string = time.Now().Format("02/01/2006 15:04:05") + " - " + action + "\n - AccountType: " + accountType + "\n - Action Value: " + val + "\n - User: " + user + transferUser + "\n - Response: " + response + "\n \n"

	archive.WriteString(txt)
	archive.Close()
}

func ReadLogs() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(archive))
}
