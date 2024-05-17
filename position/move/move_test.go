package move

import (
	"testing"
	"time"

	"github.com/andrewbackes/chess/piece"
	"github.com/andrewbackes/chess/position/square"
)

func TestMovesIs(t *testing.T) {
	testCases := []struct {
		Name  string
		Moves Moves
		Move  Move
		Want  bool
	}{
		{"Isn't-In-Nil", nil, Move{}, false},
		{"Is-In-One", Moves{Move{}}, Move{}, true},
		{"Isn't-In-One", Moves{Move{}}, Move{square.A1, square.A2, piece.Pawn, time.Duration(0)}, false},
		{"Is-In-Two", Moves{Move{square.A1, square.A2, piece.Pawn, time.Duration(0)}, Move{}}, Move{}, true},
		{"Isn't-In-Two", Moves{Move{square.A1, square.A2, piece.Pawn, time.Duration(0)}, Move{}}, Move{square.A1, square.A3, piece.Pawn, time.Duration(0)}, false},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if is := tc.Moves.Is(tc.Move); is != tc.Want {
				t.Errorf("%#v.Move(%#v) == %v, want %v", tc.Moves, tc.Move, is, tc.Want)
			}
		})
	}
}
