package snake

import (
	"math"
	"snake/base"
)

type Snake struct {
	Coordinates []base.Coordinate
	Direction   int
}

func (s Snake) IsOppositeDirection(direction int) bool {
	isOpposite := direction-s.Direction == int(math.Abs(2))
	return isOpposite
}

func (s Snake) Move() {
	// 1 - up, 2 - right, 3 - down, 4 - left
	if !s.IsOppositeDirection(s.Direction) {
		for i := len(s.Coordinates) - 1; i >= 1; i-- {
			s.Coordinates[i] = s.Coordinates[i-1]
		}

		switch s.Direction {
		case 1:
			s.Coordinates[0].Y -= 1
		case 2:
			s.Coordinates[0].X += 1
		case 3:
			s.Coordinates[0].Y += 1
		case 4:
			s.Coordinates[0].X -= 1
		}
	}
}

func (s Snake) TryEat(apple base.Coordinate) bool {
	if s.Coordinates[0].X == apple.X && s.Coordinates[0].Y == apple.Y {
		return true
	}
	return false
}
