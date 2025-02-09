package main

type Tictactoe struct {
	Board      [3][3]string
	PlayerTurn string
	GameStatus string
}

func (t *Tictactoe) Init() {
	t.Board = [3][3]string{} // Reset the board
	t.PlayerTurn = "X"       // Start with player X
	t.GameStatus = "Game in progress..."
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
