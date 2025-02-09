package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMakeMoveHandler(t *testing.T) {
	tests := []struct {
		name           string
		game           Tictactoe
		x, y           string
		wantStatus     int
		wantBoard      [3][3]string
		wantStatusText string
	}{
		{
			name: "Valid move by X",
			game: Tictactoe{
				Board:      [3][3]string{},
				PlayerTurn: "X",
				GameStatus: "ongoing",
				TurnNumber: 0,
			},
			x: "1", y: "1",
			wantStatus: http.StatusOK,
			wantBoard: [3][3]string{
				{"", "", ""},
				{"", "X", ""},
				{"", "", ""},
			},
			wantStatusText: "It&#39;s O&#39;s turn",
		},
		{
			name: "Invalid move (out of bounds)",
			game: Tictactoe{
				Board:      [3][3]string{},
				PlayerTurn: "X",
				GameStatus: "ongoing",
			},
			x: "-1", y: "3",
			wantStatus:     http.StatusBadRequest,
			wantStatusText: "Invalid move",
		},
		{
			name: "Winning move for X",
			game: Tictactoe{
				Board: [3][3]string{
					{"X", "X", ""},
					{"O", "O", ""},
					{"", "", ""},
				},
				PlayerTurn: "X",
			},
			x: "0", y: "2",
			wantBoard: [3][3]string{
				{"X", "X", "X"}, // X wins in the top row
				{"O", "O", ""},
				{"", "", ""},
			},
			wantStatus:     http.StatusOK,
			wantStatusText: "Player X wins!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request with query parameters
			req := httptest.NewRequest("GET", "/move?x="+tt.x+"&y="+tt.y, nil)
			rr := httptest.NewRecorder()

			// Call handler
			tt.game.makeMoveHandler(rr, req)

			// Check HTTP status code
			if rr.Code != tt.wantStatus {
				t.Errorf("Expected status %d, got %d", tt.wantStatus, rr.Code)
			}

			// Verify board state if move was valid
			if rr.Code == http.StatusOK && tt.wantBoard != tt.game.Board {
				t.Errorf("Expected board:\n%v\nGot:\n%v", tt.wantBoard, tt.game.Board)
			}

			// Check response body (contains GameStatus)
			body := rr.Body.String()
			if !strings.Contains(body, tt.wantStatusText) {
				t.Errorf("Expected response to contain: %q, got: %q", tt.wantStatusText, body)
			}
		})
	}
}
