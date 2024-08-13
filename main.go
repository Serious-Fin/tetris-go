package main

import (
	"aug/tetris/game"
	"aug/tetris/renderer"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	gameBoard := game.CreateBoard(10, 20)
	gameBoard.PopulateBoard(game.EmptyCell)
	renderer.Main(gameBoard)
}
