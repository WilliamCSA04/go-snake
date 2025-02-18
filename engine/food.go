package engine

type Food struct {
	Coords
	width  int
	height int
}

func SpawnFood(x int, y int, width int, height int) *Food {
	return &Food{
		Coords{
			x, y,
		},
		width,
		height,
	}
}

func (s *Snake) Move(x int, y int) {
	s.x = x
	s.y = y
}
