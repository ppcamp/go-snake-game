package main

import "go-spyder-game/pkg/game"

func main() {
	game := game.NewGame(30)
	game.Run()
}
