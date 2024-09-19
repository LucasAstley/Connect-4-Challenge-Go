package grid

func IsGridFull(board *[7][7]string) bool {
	// Check if there is a " " string in the grid
	for i := 0; i <= 6; i++ {
		for j := 0; j <= 6; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}
