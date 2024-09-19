package grid

func WinCheck(board *[7][7]string, x, y int) bool {
	// Check if there are four connected pieces in any direction
	for i := 0; i <= 6; i++ { // Check on same column
		counter := 0
		for j := 1; j <= 6; j++ {
			if board[i][j-1] == board[i][j] && board[i][j] != " " {
				counter++
			} else {
				counter = 0
			}
			if counter == 3 {
				return true
			}
		}
	}
	for i := 0; i <= 6; i++ { // Check on same line
		counter := 0
		for j := 1; j <= 6; j++ {
			if board[j-1][i] == board[j][i] && board[j][i] != " " {
				counter++
			} else {
				counter = 0
			}
			if counter == 3 {
				return true
			}
		}
	}
	tempLine := x
	counter := 0
	for k := y; k <= 6; k++ { // Check on top right diagonal ↗
		if board[tempLine][k] == board[x][y] && k != y {
			counter++
		} else {
			counter = 0
		}
		if counter == 3 {
			return true
		}
		if tempLine < 6 {
			tempLine += 1
		} else {
			k = 7
		}
	}
	tempLine = x
	counter = 0
	for k := y; k >= 0; k-- { // Check on bottom right diagonal ↘
		if board[tempLine][k] == board[x][y] && k != y {
			counter++
		} else {
			counter = 0
		}
		if counter == 3 {
			return true
		}
		if tempLine < 6 {
			tempLine += 1
		} else {
			k = -1
		}
	}
	tempLine = x
	counter = 0
	for k := y; k >= 0; k-- { // Check on bottom left diagonal ↙
		if board[tempLine][k] == board[x][y] && k != y {
			counter++
		} else {
			counter = 0
		}
		if counter == 3 {
			return true
		}
		if tempLine > 0 {
			tempLine -= 1
		} else {
			k = -1
		}
	}
	tempLine = x
	counter = 0
	for k := y; k <= 6; k++ { // Check on top left diagonal ↖
		if board[tempLine][k] == board[x][y] && k != y {
			counter++
		} else {
			counter = 0
		}
		if counter == 3 {
			return true
		}
		if tempLine > 0 {
			tempLine -= 1
		} else {
			k = 7
		}
	}
	return false
}
