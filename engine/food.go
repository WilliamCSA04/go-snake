package engine

type Food struct {
	x      int
	y      int
	width  int
	height int
}

func SpawnFood(x int, y int, width int, height int) *Food {
	return &Food{
		x,
		y,
		width,
		height,
	}
}

func (s *Food) Move(x int, y int) {
	s.x = x
	s.y = y
}
