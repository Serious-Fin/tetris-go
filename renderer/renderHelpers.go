package renderer

import (
	"aug/tetris/game"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora/v4"
)

func DisplayBoard(board *game.GameBoard) string {
	stringifiedBoard := ""

	for row := range board.Height {
		for col := range board.Width {
			cell, err := displayCell(board.Board[row][col])
			throwError(err)
			stringifiedBoard += cell
		}
		stringifiedBoard += "\n"
	}
	stringifiedBoard += "\n"

	return stringifiedBoard
}

func displayCell(cellValue int) (string, error) {
	switch cellValue {
	case game.EmptyCell:
		return aurora.Sprintf(aurora.BgWhite("  ")), nil

	case game.FilledCell:
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
