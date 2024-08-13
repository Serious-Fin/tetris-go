package main

import (
	"aug/tetris/game"
	"aug/tetris/renderer"
	"fmt"
)

func main() {
	fmt.Println("Starting tetris")
	gameBoard := game.CreateBoard(10, 20)
	gameBoard.PopulateBoard(game.EmptyCell)
	renderer.StartGame(gameBoard)
}
