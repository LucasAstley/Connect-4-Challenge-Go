package players

import (
	"Connect-4-Challenge-Go/grid"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
)

func AskToPlayAgain() bool {
	// Ask the player to restart the program
	fmt.Println("Play again ? (y/n)")
	input := ""
	fmt.Scan(&input)
	if input == "y" || input == "Y" || input == "yes" || input == "YES" || input == "Yes" {
		return true
	} else {
		grid.ClearTerminal()
		fmt.Println(aurora.BgCyan("Thanks for playing :)"))
		return false
	}
}
