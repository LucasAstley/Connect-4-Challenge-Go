package main

import (
	"Connect-4-Challenge-Go/grid"
	"Connect-4-Challenge-Go/players"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
)

func main() {
	var playGrid = grid.GridCreator()
	var player = ""
	grid.PrintGrid(&playGrid)
	for !grid.IsGridFull(&playGrid) {
		players.NextPlayer(&player)
		x, y := players.PieceInsert(&playGrid, player)
		grid.PrintGrid(&playGrid)
		if grid.WinCheck(&playGrid, x, y) {
			fmt.Printf("\x1b[2J") // Clear the console
			grid.PrintGrid(&playGrid)
			fmt.Println("")
			if player == "X" {
				fmt.Printf("🎉 Congratulations %s, you won the game!\n", aurora.BgRed(player))
			} else if player == "O" {
				fmt.Printf("🎉 Congratulations %s, you won the game!\n", aurora.BgYellow(player))
			}
			fmt.Println("")
			if players.AskToPlayAgain() {
				main()
				return
			}
			return
		}
	}
	fmt.Println("")
	fmt.Println(aurora.BgRed("❌  Oh snap!"), aurora.BrightRed("Grid is full!"))
	if players.AskToPlayAgain() {
		main()
		return
	}
	return
}
