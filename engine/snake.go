package engine

type Snake struct {
	x,
	y,
	width int
	height int
	size   int
}

func SpawnSnake(x int, y int, width int, height int) *Snake {
	return &Snake{
		x,
		y,
		width,
		height,
		1,
	}
}

func (s *Food) Move(x int, y int) {
	s.x = x
	s.y = y
}

func (s *Snake) Eat() {
	s.size++
}
