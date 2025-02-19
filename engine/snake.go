package engine

type Snake struct {
	x      []int
	y      []int
	width  int
	height int
	size   int
}

func SpawnSnake(startX int, startY int, width int, height int) *Snake {
	x := append(make([]int, 0), startX)
	y := append(make([]int, 0), startY)
	return &Snake{
		x,
		y,
		width,
		height,
		1,
	}
}

func (s *Snake) Move(x int, y int) {
	s.x = append([]int{x}, s.x[:len(s.x)-1]...)
	s.y = append([]int{y}, s.y[:len(s.y)-1]...)
}

func (s *Snake) Eat() {
	s.size++
}
