package game

const (
	EmptyCell     = 0
	TravelingCell = 1
	FilledCell    = 2
)

type GameBoard struct {
	Width  int
	Height int
	Board  [][]int
}

func CreateBoard(width, height int) *GameBoard {
	gameBoard := GameBoard{width, height, make([][]int, height)}
	for row := range height {
		gameBoard.Board[row] = make([]int, width)
	}
	return &gameBoard
}

func (b *GameBoard) PopulateBoard(value int) {
	for row := range b.Height {
		for col := range b.Width {
			b.Board[row][col] = value
		}
	}
}
