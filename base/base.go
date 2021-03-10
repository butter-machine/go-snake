package base

type Coordinate struct {
	X int
	Y int
}

func Contains(s []int, elem int) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}

	return false
}
