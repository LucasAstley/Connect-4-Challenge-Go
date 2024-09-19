# Ymmersion - Connect 4 Challenge <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/>

This is the latest ymmersion exercice, this challenge consists of creating a console connect 4 in Go.

## Instructions

```console
$ go run . 
|   |   |   |   |   |   |   |
-----------------------------
|   |   |   |   |   |   |   |
-----------------------------
|   |   |   |   |   |   |   |
-----------------------------
|   |   | X |   |   |   |   |
-----------------------------
|   |   | O | X |   |   |   |
-----------------------------
|   | X | X | O | O |   |   |
-----------------------------
| O | X | O | X | O |   |   |

Player X's turn. Choose column (0-7):
```

- Check for wins after each move: The game should verify if there are four connected pieces in any direction—vertical, horizontal, or diagonal.
- Continue the game until a player wins or all spaces are filled, resulting in a draw.
- Comment your code to explain the implementation of game mechanics and logic decisions. Be careful to put relevant comments!
- Structure your code to handle game state management, user input, and display output cleanly and efficiently.

## Bonus tasks

- Include error handling for invalid inputs such as entering non-numeric values, choosing columns that are out of bounds, or selecting columns that are already full.
- Enhance the game by adding colors to the discs (X and O).
- Allow the players to play multiple rounds without restarting the program.
- Improve the terminal UI to make it more user-friendly, e.g., by clearing the screen before displaying the board.

## My structure

```console
├── grid
│   ├── gridCreator.go
│   ├── gridPrinter.go
│   ├── isGridFull.go
│   └── winChecker.go
├── players
│   ├── nextPlayer.go
│   ├── pieceInsert.go
│   └── playAgain.go
├── main.go
├── technical_structure_and_logic.txt
└── README.md
```

As explained in `technical_structure_and_logic` my playing grid is represented by a bi-dimensional array :

![schema](https://i.imgur.com/Llq6nET.png)

I also use [aurora](https://github.com/logrusorgru/aurora) package to print elements with color in the console. 

## Demo

| Gameplay                                     | Full column error                                    | Bad input error                                     |
|----------------------------------------------|------------------------------------------------------|-----------------------------------------------------|
| ![demo-gif](https://i.imgur.com/Nxbo1IN.gif) | ![column-error-gif](https://i.imgur.com/FNbWuZs.gif) | ![input-error-gif](https://i.imgur.com/6NN40Qe.gif) |
