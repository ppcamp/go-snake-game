package game

import (
	"fmt"
	"go-spyder-game/pkg/config"
	"go-spyder-game/pkg/models"
	"go-spyder-game/pkg/utils"
	"time"

	"github.com/nsf/termbox-go"
)

type Game struct {
	snake      Snake
	food       models.Point
	boundaries models.Point
	Stop       bool
}

func NewGame(size int) *Game {
	boundaries := models.Point{
		X: size,
		Y: size,
	}

	return &Game{
		snake:      NewSnake(boundaries),
		food:       GenFoodLocation(boundaries),
		boundaries: boundaries,
	}
}

func (g *Game) getColisionType() config.PlaceType {
	head := g.snake.Head()
	food := g.food
	nextMove := g.snake.FutureMove()

	if head.X == food.X && head.Y == food.Y {
		return config.IsFood
	} else if g.snake.At(nextMove.X, nextMove.Y) == config.IsSnake {
		return config.IsSnake
	}

	return config.IsEmpty
}

func (g *Game) isFood(x, y int) bool {
	return g.food.X == x && g.food.Y == y
}

func (g *Game) score() int {
	return (len(g.snake.path) - 2) * 10
}

func (g *Game) Run() {
	go g.WatchKeyboard()

	for {

		fmt.Println("\nGame")
		fmt.Printf("Board: %dx%d\n", g.boundaries.X, g.boundaries.Y)
		fmt.Printf("Score: %d\n", g.score())

		if g.Stop {
			break
		}

		for y := -1; y <= g.boundaries.Y; y++ {
			if y == -1 || y == g.boundaries.Y {
				continue
			}

			for x := -1; x <= g.boundaries.X; x++ {
				if x == -1 {
					fmt.Print("+")
				} else if x == g.boundaries.X {
					fmt.Print("+\n")
				} else {
					print := "█"

					switch g.getColisionType() {
					case config.IsSnake:
						g.Stop = true
					case config.IsFood:
						g.snake.Eat()
						g.food = GenFoodLocation(g.boundaries)
					default:
					}

					switch g.snake.At(x, y) {
					case config.IsSnakeHead:
						print = "*"
					case config.IsSnake:
						print = "*"
					default:
						print = "█"
					}

					if g.isFood(x, y) {
						print = "~"
					}

					fmt.Print(print)
				}
			}
		}

		time.Sleep(50 * time.Millisecond)
		g.snake.Move()

		g.snake.UnlockMovement()

		utils.ClearConsole()
	}
}

func (g *Game) WatchKeyboard() {
	if err := termbox.Init(); err != nil {
		panic("fail to startup the termbox library")
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				g.snake.SetDirection(config.DirectionLeft)
			case termbox.KeyArrowDown:
				g.snake.SetDirection(config.DirectionUp)
			case termbox.KeyArrowRight:
				g.snake.SetDirection(config.DirectionRight)
			case termbox.KeyArrowUp:
				g.snake.SetDirection(config.DirectionDown)
			case termbox.KeyEsc:
				g.Stop = true
			default:
				g.snake.SetDirection(g.snake.direction)
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		if g.Stop {
			break
		}
	}
}
