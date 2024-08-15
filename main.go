package main

import (
	"aug/tetris/game"
	"aug/tetris/renderer"
	"fmt"
)

func main() {
	fmt.Println("Starting tetris")
	board := game.CreateBoard(10, 20)
	board.PopulateBoard(game.EmptyCell)
	renderer.StartGame(board)
}
