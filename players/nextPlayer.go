package players

func NextPlayer(actualPlayer *string) {
	// Change to next player
	if *actualPlayer == "X" {
		*actualPlayer = "O"
	} else if *actualPlayer == "O" {
		*actualPlayer = "X"
	} else {
		*actualPlayer = "X"
	}
}
