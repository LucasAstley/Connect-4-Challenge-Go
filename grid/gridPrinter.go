package grid

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"os"
	"os/exec"
	"runtime"
)

func PrintGrid(board *[7][7]string) {
	// Clear the terminal and print the grid line by line
	ClearTerminal()
	for j := 6; j >= 0; j-- {
		fmt.Printf("|")
		for i := 0; i <= 6; i++ {
			if board[i][j] == "X" {
				fmt.Printf(" %s |", aurora.Red(board[i][j])) // Print the "X" piece in red
			} else if board[i][j] == "O" {
				fmt.Printf(" %s |", aurora.Yellow(board[i][j])) // Print the "O" piece in yellow
			} else {
				fmt.Printf(" %s |", board[i][j])
			}
		}
		fmt.Println(" ")
		if j != 0 {
			fmt.Println("-----------------------------")
		}
	}
}

func ClearTerminal() {
	// Use clear command (windows and linux) to clear terminal
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
