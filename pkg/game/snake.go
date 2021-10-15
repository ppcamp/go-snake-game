package game

import (
	"go-spyder-game/pkg/config"
	"go-spyder-game/pkg/models"
	"go-spyder-game/pkg/utils"
)

type Snake struct {
	path      []models.Point
	direction config.Direction

	boudaries models.Point
	table     [][]config.PlaceType
	blocked   bool
}

func NewSnake(boudaries models.Point) Snake {
	head := models.Point{
		X: utils.Median(boudaries.X) - 12,
		Y: utils.Median(boudaries.Y),
	}
	tail := models.Point{
		X: head.X + 1,
		Y: head.Y,
	}

	table := make([][]config.PlaceType, boudaries.X)

	for i := range table {
		table[i] = make([]config.PlaceType, boudaries.Y)
	}

	table[head.X][head.Y] = config.IsSnakeHead
	table[tail.X][tail.Y] = config.IsSnake

	return Snake{
		path:      []models.Point{head, tail},
		direction: config.DirectionLeft,
		boudaries: boudaries,
		table:     table,
	}
}

func (s *Snake) grow(point models.Point) {

	oldHead := s.Head()

	// mark in view table
	s.table[point.X][point.Y] = config.IsSnakeHead
	s.table[oldHead.X][oldHead.Y] = config.IsSnake

	s.path = append([]models.Point{point}, s.path...)
}

func (s *Snake) pop() {
	tail := s.Tail()

	s.table[tail.X][tail.Y] = config.IsEmpty
	s.path = s.path[:len(s.path)-1]
}

func (s *Snake) SetDirection(direction config.Direction) {
	if (s.direction == config.DirectionDown && direction == config.DirectionUp) ||
		(s.direction == config.DirectionUp && direction == config.DirectionDown) ||
		(s.direction == config.DirectionRight && direction == config.DirectionLeft) ||
		(s.direction == config.DirectionLeft && direction == config.DirectionRight) ||
		s.blocked {
		return
	}
	s.blocked = true
	s.direction = direction
}

func (s *Snake) Path() []models.Point {
	return s.path
}

func (s *Snake) Head() models.Point {
	return s.path[0]
}

func (s *Snake) Tail() models.Point {
	return s.path[len(s.path)-1]
}

func (s *Snake) movementPoint() models.Point {
	head := s.Head()
	point := models.Point{}

	switch s.direction {
	case config.DirectionUp:
		point = models.Point{X: head.X, Y: head.Y + 1}
	case config.DirectionDown:
		point = models.Point{X: head.X, Y: head.Y - 1}
	case config.DirectionLeft:
		point = models.Point{X: head.X - 1, Y: head.Y}
	case config.DirectionRight:
		point = models.Point{X: head.X + 1, Y: head.Y}
	default:
		panic("must pass a direction")
	}

	// don't skip from the screen
	if point.X == s.boudaries.X-1 {
		point.X = 0
	} else if point.Y == s.boudaries.Y-1 {
		point.Y = 0
	} else if point.X == -1 {
		point.X = s.boudaries.X - 1
	} else if point.Y == -1 {
		point.Y = s.boudaries.Y - 1
	}

	return point
}

func (s *Snake) Move() {
	switch s.direction {
	case config.DirectionUp:
		s.grow(s.movementPoint())
		s.pop()
	case config.DirectionDown:
		s.grow(s.movementPoint())
		s.pop()
	case config.DirectionLeft:
		s.grow(s.movementPoint())
		s.pop()
	case config.DirectionRight:
		s.grow(s.movementPoint())
		s.pop()
	default:
		panic("must pass a direction")
	}
}

func (s *Snake) Eat() {
	switch s.direction {
	case config.DirectionUp:
		s.grow(s.movementPoint())
	case config.DirectionDown:
		s.grow(s.movementPoint())
	case config.DirectionLeft:
		s.grow(s.movementPoint())
	case config.DirectionRight:
		s.grow(s.movementPoint())
	default:
		panic("must pass a direction")
	}

	s.Move()
}

func (s *Snake) At(x, y int) config.PlaceType {
	return s.table[x][y]
}

func (s *Snake) UnlockMovement() {
	s.blocked = false
}

func (s *Snake) FutureMove() models.Point {
	return s.movementPoint()
}
