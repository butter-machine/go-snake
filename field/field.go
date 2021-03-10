package field

import (
	"github.com/eiannone/keyboard"
	"math/rand"
	"snake/base"
	"snake/snake"
	"time"
)

type Field struct {
	MaxX             int
	MaxY             int
	RefreshDelimiter int
	RefreshRate      time.Duration
	Apple            *base.Coordinate
	Snake            *snake.Snake
}

func NewField(maxX int, maxY int, snake *snake.Snake) Field {
	rand.Seed(time.Now().Unix())
	return Field{
		MaxX:             maxX,
		MaxY:             maxY,
		RefreshDelimiter: 5,
		RefreshRate:      time.Second / 5,
		Apple:            &base.Coordinate{X: 0, Y: 0},
		Snake:            snake,
	}
}

func (f *Field) SpawnApple() {
	snakeX := make([]int, len(f.Snake.Coordinates))
	snakeY := make([]int, len(f.Snake.Coordinates))
	for index, coordinate := range f.Snake.Coordinates {
		snakeX[index] = coordinate.X
		snakeY[index] = coordinate.Y
	}

	appleX := 0
	for base.Contains(snakeX, appleX) {
		appleX = rand.Intn(f.MaxX)
	}

	appleY := 0
	for base.Contains(snakeY, appleY) {
		appleY = rand.Intn(f.MaxY)
	}

	f.Apple = &base.Coordinate{X: appleX, Y: appleY}
}

func (f *Field) IsBump() bool {
	return f.Snake.Coordinates[0].X >= f.MaxX ||
		f.Snake.Coordinates[0].X < 0 ||
		f.Snake.Coordinates[0].Y >= f.MaxY ||
		f.Snake.Coordinates[0].Y < 0
}

func (f *Field) ListenKeys() {
	up := 119
	right := 100
	down := 115
	left := 97

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case rune(up):
			f.Snake.Direction = 1
		case rune(right):
			f.Snake.Direction = 2
		case rune(down):
			f.Snake.Direction = 3
		case rune(left):
			f.Snake.Direction = 4
		}
	}
}
