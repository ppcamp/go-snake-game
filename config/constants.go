package config

type PlaceType int

const (
	IsEmpty PlaceType = iota
	IsFood
	IsSnake
	IsSnakeHead
)
