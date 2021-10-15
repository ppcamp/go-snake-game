package game

import (
	"go-spyder-game/models"
	"math/rand"
)

func GenFoodLocation(boundaries models.Point) models.Point {
	return models.Point{
		X: rand.Intn(boundaries.X - 1),
		Y: rand.Intn(boundaries.Y - 1),
	}
}
