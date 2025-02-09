package main

import (
	"errors"
	"fmt"
)

type Tictactoe struct {
	Board [3][3]string
	Turn  int
}

func (t *Tictactoe) printBoard() {
	fmt.Printf("====TURN %d====\n", t.Turn)
	for i := 0; i < len(t.Board); i++ {
		for j := 0; j < len(t.Board[i]); j++ {
			if t.Board[i][j] == "" {
				fmt.Print("-")
			} else {
				fmt.Printf("%s", t.Board[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func (t *Tictactoe) makeMove(x, y int, playerSymbol string) error {
	if x < 0 || x >= 3 || y < 0 || y >= 3 { // Fix bounds check
		return errors.New("square is out of bounds, try again")
	}
	if t.Board[x][y] != "" { // Prevent overwriting existing moves
		return errors.New("square is already occupied, try again")
	}

	t.Board[x][y] = playerSymbol // Modify the board directly
	t.Turn++
	return nil
}

func (t *Tictactoe) checkWin(s string) bool {
	// check horizontal
	for i := 0; i < len(t.Board); i++ {
		if t.Board[i][0] == s && t.Board[i][1] == s && t.Board[i][2] == s {
			return true
		}
	}

	// check vertical
	for i := 0; i < len(t.Board); i++ {
		if t.Board[0][i] == s && t.Board[1][i] == s && t.Board[2][i] == s {
			return true
		}
	}

	// check diagonal
	if (t.Board[0][0] == s && t.Board[1][1] == s && t.Board[2][2] == s) || (t.Board[2][0] == s && t.Board[1][1] == s && t.Board[0][2] == s) {
		return true
	}

	return false
}

func (t *Tictactoe) Init() {
	t.Turn = 1
}

func main() {
	board := Tictactoe{}
	board.Init()

	board.printBoard()
	board.makeMove(2, 0, "X")
	board.printBoard()
	board.makeMove(1, 0, "O")
	board.printBoard()
	board.makeMove(2, 1, "X")
	board.printBoard()
	board.makeMove(2, 2, "0")
	board.printBoard()
	board.makeMove(1, 1, "X")
	board.printBoard()
	board.makeMove(0, 1, "0")
	board.printBoard()
	board.makeMove(0, 2, "X")
	board.printBoard()
	win := board.checkWin("X")
	if win {
		fmt.Println("Player X wins")
	}
}
