package grid

func GridCreator() [7][7]string {
	// Create the 7x7 bi-dimensional grid
	var board [7][7]string
	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}
	return board
}
