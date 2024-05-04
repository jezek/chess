package board

type Side uint

const (
	ShortSide, kingSide = Side(0), Side(0)
	LongSide, queenSide = Side(1), Side(1)
)

const SIDE_COUNT = 2

// Sides is used to range through the sides of the board.
var Sides = [SIDE_COUNT]Side{ShortSide, LongSide}
