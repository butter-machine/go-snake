package views

import (
	"snake/field"
)

type View struct {
	SnakeElem string
	AppleElem string
	FieldElem string
	Field     *field.Field
	Score     int
}

type IView interface {
	Draw()
	Run()
	Init()
	Destroy()
}
