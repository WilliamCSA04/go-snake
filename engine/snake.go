package engine

type Snake struct {
	x         []int
	y         []int
	lastTailX int
	lastTailY int
	width     int
	height    int
	size      int
}

func SpawnSnake(startX int, startY int, width int, height int) *Snake {
	x := append(make([]int, 0), startX)
	y := append(make([]int, 0), startY)
	return &Snake{
		x:         x,
		y:         y,
		lastTailX: startX,
		lastTailY: startY,
		width:     width,
		height:    height,
		size:      1,
	}
}

func (s *Snake) Move(x int, y int) {
	s.lastTailX = s.x[len(s.x)-1]
	s.lastTailY = s.y[len(s.y)-1]
	s.x = append([]int{x}, s.x[:len(s.x)-1]...)
	s.y = append([]int{y}, s.y[:len(s.y)-1]...)
}

func (s *Snake) Eat() {
	s.x = append(s.x, s.lastTailX)
	s.y = append(s.y, s.lastTailY)
}
