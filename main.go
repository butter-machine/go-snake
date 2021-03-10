package main

import (
	"snake/base"
	"snake/field"
	"snake/snake"
	"snake/views"
)

func main() {
	snakeInitialCoordinates := make([]base.Coordinate, 4)
	for i := 0; i < 4; i++ {
		snakeInitialCoordinates[i].X = 0 + i
		snakeInitialCoordinates[i].Y = 0
	}
	snake_ := snake.Snake{Direction: 2, Coordinates: snakeInitialCoordinates}
	gameField := field.NewField(15, 15, &snake_)

	view := views.TerminalView{
		View: views.View{
			SnakeElem: "■",
			AppleElem: "●",
			FieldElem: ".",
			Field:     &gameField,
			Score:     0,
		},
	}

	view.Run()
}
