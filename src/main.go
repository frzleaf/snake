package main

import (
	"game"
)

var (
	BOARD_HEIGHT = 80
	BOARD_WIDTH = 24
)

func main() {
	funnyGame := game.CliGame{}
	funnyGame.InitDefault()
	funnyGame.Snakes[0].Eat(10)
	funnyGame.Start()
}