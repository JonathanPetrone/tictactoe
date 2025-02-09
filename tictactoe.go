package main

type Tictactoe struct {
	Board      [3][3]string
	PlayerTurn string
	GameStatus string
	TurnNumber int
}

func (t *Tictactoe) Init() {
	t.Board = [3][3]string{} // Reset the board
	t.PlayerTurn = "X"       // Start with player X
	t.GameStatus = "Game in progress..."
	t.TurnNumber = 1
}

func (t *Tictactoe) checkWin(player string) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if t.Board[i][0] == player && t.Board[i][1] == player && t.Board[i][2] == player {
			return true
		}
		if t.Board[0][i] == player && t.Board[1][i] == player && t.Board[2][i] == player {
			return true
		}
	}

	// Check diagonals
	if (t.Board[0][0] == player && t.Board[1][1] == player && t.Board[2][2] == player) ||
		(t.Board[0][2] == player && t.Board[1][1] == player && t.Board[2][0] == player) {
		return true
	}

	return false
}
