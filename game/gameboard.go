package game

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

const (
	EmptyCell  = 0
	FilledCell = 1
	CellO      = 2
	CellI      = 3
	CellS      = 4
	CellZ      = 5
	CellL      = 6
	CellJ      = 7
	CellT      = 8
)

type GameBoard struct {
	Width  int
	Height int
	Board  [][]int
}

func CreateBoard(width, height int) *GameBoard {
	board := GameBoard{width, height, make([][]int, height)}
	for row := range height {
		board.Board[row] = make([]int, width)
	}
	return &board
}

func (b *GameBoard) PopulateBoard(value int) {
	for row := range b.Height {
		for col := range b.Width {
			b.Board[row][col] = value
		}
	}
}

func (b *GameBoard) StringifyBoard() string {
	stringifiedBoard := ""

	for row := range b.Height {
		stringifiedBoard += b.stringifyRow(row)
	}

	for range b.Width + 2 {
		stringifiedBoard += aurora.Sprintf(aurora.BgBlue("  "))
	}
	stringifiedBoard += "\n"

	return stringifiedBoard
}

func (b *GameBoard) stringifyRow(row int) string {
	stringifiedRow := aurora.Sprintf(aurora.BgBlue("  "))
	for col := range b.Width {
		cell, err := stringifyCell(b.Board[row][col])
		throwError(err)
		stringifiedRow += cell
	}
	stringifiedRow += aurora.Sprintf(aurora.BgBlue("  "))
	return stringifiedRow + "\n"
}

func stringifyCell(cellValue int) (string, error) {
	switch cellValue {
	case EmptyCell:
		return aurora.Sprintf(aurora.BgBlack("  ")), nil

	case FilledCell:
		return aurora.Sprintf(aurora.BgWhite("  ")), nil

	case CellO:
		return aurora.Sprintf(aurora.BgYellow("  ")), nil

	case CellI:
		return aurora.Sprintf(aurora.BgCyan("  ")), nil

	case CellS:
		return aurora.Sprintf(aurora.BgRed("  ")), nil

	case CellZ:
		return aurora.Sprintf(aurora.BgGreen("  ")), nil

	case CellL:
		return aurora.Sprintf(aurora.BgBrightRed("  ")), nil

	case CellJ:
		return aurora.Sprintf(aurora.BgBrightMagenta("  ")), nil

	case CellT:
		return aurora.Sprintf(aurora.BgMagenta("  ")), nil
	}

	return "", fmt.Errorf("\"%d\" is not a valid board cell value", cellValue)
}

func throwError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (b *GameBoard) DrawFigureAs(figure *Figure, state int) {
	for _, offset := range figure.Geometries[figure.GeometryIndex].Points {
		if b.PointInRenderBounds(figure.MiddlePos.Row+offset.Row, figure.MiddlePos.Col+offset.Col) {
			b.Board[figure.MiddlePos.Row+offset.Row][figure.MiddlePos.Col+offset.Col] = state
		}
	}
}

func (b *GameBoard) PointInRenderBounds(row, col int) bool {
	return row >= 0 && row < b.Height && col >= 0 && col < b.Width
}
