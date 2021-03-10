package views

import (
	"fmt"
	tm "github.com/buger/goterm"
	"time"
)

type TerminalView struct {
	View
	IView
}

func (tv *TerminalView) GetFieldString() string {
	fieldString := ""
	for y := 0; y < tv.Field.MaxY; y++ {
		for x := 0; x < tv.Field.MaxX; x++ {
			currentElem := ""
			for _, coordinate := range tv.Field.Snake.Coordinates {
				if x == coordinate.X && y == coordinate.Y {
					currentElem = tv.SnakeElem + " "
				}
			}

			if tv.Field.Apple.X == x && tv.Field.Apple.Y == y {
				currentElem = tv.AppleElem + " "
			}

			if currentElem == "" {
				currentElem = tv.FieldElem + " "
			}
			fieldString += currentElem
		}
		fieldString += "\n"
	}

	return fieldString
}

func (tv *TerminalView) Draw() {
	tm.MoveCursor(1, 1)
	tm.Flush()
	tm.Println(tv.GetFieldString())
	time.Sleep(tv.Field.RefreshRate)
}

func (tv *TerminalView) Init() {
	tm.Clear()
}

func (tv *TerminalView) Destroy() {
	tm.MoveCursor(1, tv.Field.MaxY+1)
	tm.Flush()
	fmt.Printf("Game over. Score: %d\n", tv.Score)
}

func (tv *TerminalView) Run() {
	go tv.Field.ListenKeys()
	tv.Field.SpawnApple()
	tv.Init()

	for {
		tv.Field.Snake.Move()

		appleEaten := tv.Field.Snake.TryEat(*tv.Field.Apple)
		if appleEaten {
			tv.Field.SpawnApple()
			tv.Score++
		}
		tv.Draw()

		if tv.Field.IsBump() {
			break
		}
	}

	tv.Destroy()
}
