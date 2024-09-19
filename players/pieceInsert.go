package players

import (
	"Connect-4-Challenge-Go/grid"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"time"
)

func PieceInsert(board *[7][7]string, player string) (x, y int) {
	// Insert the piece in the selected column
	input := 0
	fmt.Println("")
	fmt.Printf("Player %s's turn. Choose column (1-7):\n", player)
	fmt.Scan(&input)
	grid.ClearTerminal()
	grid.PrintGrid(board)
	if input < 1 || input > 7 {
		fmt.Println("")
		fmt.Println(aurora.BgRed("❌  Oh snap!"), aurora.BrightRed("Bad input"))
		PieceInsert(board, player)
		return
	}
	if board[input-1][6] != " " {
		fmt.Println("")
		fmt.Println(aurora.BgRed("❌  Oh snap!"), aurora.Sprintf(aurora.BrightRed("Column %v is full"), aurora.BrightWhite(input)))
		PieceInsert(board, player)
		return
	} else {
		for i := 6; i >= 0; i-- {
			if board[input-1][i] != " " {
				board[input-1][i+1] = player
				return input - 1, i
			} else if i == 0 {
				board[input-1][0] = player
				return input - 1, 0
			} else {
				board[input-1][i] = player
				grid.PrintGrid(board)
				fmt.Println("")
				fmt.Printf("Player %s's turn. Choose column (1-7):\n", player)
				board[input-1][i] = " "
				time.Sleep(55 * time.Millisecond) // Delay for falling animation
			}
		}
	}
	return
}
