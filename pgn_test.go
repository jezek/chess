package chess

import (
	"github.com/andrewbackes/chess/board"
	"strings"
	"testing"
)

func TestPGNnullmoves(t *testing.T) {
	expected := `[Event ""]
[Site ""]
[Date ""]
[Round ""]
[White ""]
[Black ""]
[Result "1-0"]
[Setup "1"]
[FEN "rnbq1bnr/ppppkppp/8/4p2Q/4P3/8/PPPP1PPP/RNB1KBNR w KQ - 1 3"]

3. h5e5 1-0

`
	g, _ := FromFEN("rnbq1bnr/ppppkppp/8/4p2Q/4P3/8/PPPP1PPP/RNB1KBNR w KQ - 1 3")
	m, err := g.ParseMove("Qxe5#")
	if err != nil {
		t.Error("couldnt parse move")
	}
	g.MakeMove(m)
	if g.PGN() != expected {
		t.Log(g.PGN())
		t.Fail()
	}
}

func TestPGNoutput(t *testing.T) {
	expected := `[Event ""]
[Site ""]
[Date ""]
[Round ""]
[White ""]
[Black ""]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0

`
	g := NewGame()
	moves := []string{"e4", "e5", "Qh5", "Ke7", "Qxe5#"}
	for _, move := range moves {
		m, err := g.ParseMove(move)
		if err != nil {
			t.Error("couldnt parse move")
		}
		g.MakeMove(m)
	}
	if g.PGN() != expected {
		t.Fail()
	}
}

func TestReadOnePGN(t *testing.T) {
	input := `[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0
`
	games, err := ReadPGN(strings.NewReader(input))
	if err != nil || len(games) != 1 {
		t.Log(games)
		t.Log(err)
		t.Fail()
	}

}

func TestReadTwoPGN(t *testing.T) {
	input := `[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0

[Event "two"]
[Round "2"]
[Result "1/2-1/2"]

1. e2e4 e7e5 2. d1h5 e8e7 1/2-1/2
`
	games, err := ReadPGN(strings.NewReader(input))
	if err != nil || len(games) != 2 {
		t.Log(games)
		t.Log(err)
		t.Fail()
		t.Fail()
	}

}

func TestReadThreePGN(t *testing.T) {
	input := `[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0

[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0

[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0
`
	games, err := ReadPGN(strings.NewReader(input))
	if err != nil || len(games) != 3 {
		t.Log(games)
		t.Log(err)
		t.Fail()
		t.Fail()
	}
}

func TestReadPGN(t *testing.T) {
	input := `[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0
`
	games, _ := ReadPGN(strings.NewReader(input))
	games[0].Tags["Event"] = "one"
	games[0].Tags["Round"] = "1"
	games[0].Tags["Result"] = "1-0"

	moves := []string{"e2e4", "e7e5", "d1h5", "e8e7", "h5e5"}
	for i, m := range games[0].Moves {
		if m != moves[i] {
			t.Log(games[0].Moves)
			t.Fail()
		}
	}
}

func TestFromPGN(t *testing.T) {
	pgn := NewPGN()
	pgn.Tags["Event"] = "test"
	pgn.Moves = []string{"e2e4", "e7e5", "d1h5", "e8e7", "h5e5"}
	game, err := FromPGN(pgn)
	if err != nil {
		t.Error(err)
	}
	moves := game.MoveHistory()
	for i, m := range pgn.Moves {
		if board.Move(m) != moves[i] {
			t.Log(moves)
			t.Fail()
		}
	}
}
