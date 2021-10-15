package main

import "go-spyder-game/game"

func main() {
	game := game.NewGame(30)
	game.Run()
}
