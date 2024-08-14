package game

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

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

func (b *GameBoard) StringifyBoard() string {
	stringifiedBoard := ""

	for row := range b.Height {
		for col := range b.Width {
			cell, err := stringifyCell(b.Board[row][col])
			throwError(err)
			stringifiedBoard += cell
		}
		stringifiedBoard += "\n"
	}
	stringifiedBoard += "\n"

	return stringifiedBoard
}

func stringifyCell(cellValue int) (string, error) {
	switch cellValue {
	case EmptyCell:
		return aurora.Sprintf(aurora.BgWhite("  ")), nil

	case TravelingCell:
		return aurora.Sprintf(aurora.BgGreen("  ")), nil

	case FilledCell:
		return aurora.Sprintf(aurora.BgBlack("  ")), nil
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
