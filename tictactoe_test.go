package main

import (
	"testing"
)

func newBoard(rows [3]string) [3][3]string {
	var board [3][3]string
	for i, row := range rows {
		for j, cell := range row {
			board[i][j] = string(cell) // Convert rune to string
		}
	}
	return board
}

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name     string
		input    Tictactoe
		expected bool
	}{
		{
			name: "Check X win (row)",
			input: Tictactoe{
				Board: [3][3]string{
					{"X", "X", "X"},
					{"O", "", "O"},
					{"", "O", ""},
				},
				PlayerTurn: "X",
				GameStatus: "ongoing",
				TurnNumber: 7,
			},
			expected: true,
		},
		{
			name: "Check O win (column)",
			input: Tictactoe{
				Board: [3][3]string{
					{"X", "O", "X"},
					{"X", "O", ""},
					{"", "O", "X"},
				},
				PlayerTurn: "O",
				GameStatus: "ongoing",
				TurnNumber: 8,
			},
			expected: true,
		},
		{
			name: "No winner yet",
			input: Tictactoe{
				Board: [3][3]string{
					{"X", "O", "X"},
					{"O", "X", "O"},
					{"O", "X", ""},
				},
				PlayerTurn: "X",
				GameStatus: "ongoing",
				TurnNumber: 9,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.checkWin(tt.input.PlayerTurn)
			if actual != tt.expected {
				t.Errorf("CheckWin(%v) = %v; want %v", tt.input.Board, actual, tt.expected)
			}
		})
	}
}
