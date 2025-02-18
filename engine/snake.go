package engine

type Snake struct {
	Coords
	width  int
	height int
}

type Coords struct {
	x int
	y int
}

func Spawn(x int, y int, width int, height int) *Snake {
	return &Snake{
		Coords{
			x, y,
		},
		width,
		height,
	}
}
