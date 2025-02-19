package engine

import (
	"Snake/ui"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	START     = iota
	ON_GOING  = iota
	GAME_OVER = iota
)

type Direction int

type Game struct {
	screen    tcell.Screen
	snake     *Snake
	food      *Food
	state     int
	direction Direction
}

const (
	UP    Direction = iota
	DOWN  Direction = iota
	LEFT  Direction = iota
	RIGHT Direction = iota
)

var (
	width, height int
)

func NewGame() *Game {
	width, height = 2, 1
	s := ui.NewScreen()
	snake := SpawnSnake(6, 6, 2, 1)
	f := SpawnFood(0, 0, 2, 1)
	return &Game{
		screen:    s,
		snake:     snake,
		food:      f,
		state:     START,
		direction: RIGHT,
	}
}

func (g *Game) Screen() tcell.Screen {
	return g.screen
}

func (g *Game) Controller(ev tcell.Event) bool {
	switch event := ev.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlC:
			return false
		case tcell.KeyLeft:
			if g.direction == RIGHT {
				return true
			}
			g.direction = LEFT
		case tcell.KeyRight:
			if g.direction == LEFT {
				return true
			}
			g.direction = RIGHT
		case tcell.KeyUp:
			if g.direction == DOWN {
				return true
			}
			g.direction = UP
		case tcell.KeyDown:
			if g.direction == UP {
				return true
			}
			g.direction = DOWN
		}
	}
	return true
}

func (g *Game) CanSnakeMove(x, y int) bool {
	for i := 0; i < len(g.snake.x); i++ {
		if g.snake.x[i] == x && g.snake.y[i] == y {
			g.state = GAME_OVER
			return false
		}
	}
	sw, sh := g.screen.Size()
	if x < 0 || x >= sw || y < 0 || y >= sh {
		g.state = GAME_OVER
		return false
	}
	return true
}

func (g *Game) Update(x, y int) {
	if !g.CanSnakeMove(x, y) {
		return
	}
	g.screen.Clear()

	// Style for the food
	foodStyle := tcell.StyleDefault.Background(tcell.ColorYellow).Foreground(tcell.ColorWhite)
	if g.food.x == x && g.food.y == y {
		g.snake.Eat()
		sw, sh := g.screen.Size()
		newPositionX := rand.Intn(sw)
		if newPositionX%2 != 0 {
			newPositionX++
		}
		newPositionY := rand.Intn(sh)
		ui.Draw(g.screen, newPositionX, newPositionY, width, height, foodStyle)
		g.food.Move(newPositionX, newPositionY)

	} else {
		ui.Draw(g.screen, g.food.x, g.food.y, width, height, foodStyle)
	}
	// Style for the snake
	g.snake.Move(x, y)
	style := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
	for i := 0; i < len(g.snake.x); i++ {
		ui.Draw(g.screen, g.snake.x[i], g.snake.y[i], width, height, style)
	}

}

func (g *Game) GameLoop() {
	style := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
	ui.Draw(g.screen, g.snake.x[0], g.snake.y[0], width, height, style)
	foodStyle := tcell.StyleDefault.Background(tcell.ColorYellow).Foreground(tcell.ColorWhite)
	ui.Draw(g.screen, g.food.x, g.food.y, width, height, foodStyle)
	g.state = ON_GOING

	evCh := make(chan tcell.Event)

	// Start a goroutine to poll events
	go func() {
		for {
			evCh <- g.screen.PollEvent()
		}
	}()

	for {

		select {
		case ev := <-evCh:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if !g.Controller(ev) {
					return
				}
			default:
				g.screen.Sync()
			}
		default:

		}

		if g.state == GAME_OVER {
			return
		}
		switch g.direction {
		case RIGHT:
			g.Update(g.snake.x[0]+2, g.snake.y[0])
		case LEFT:
			g.Update(g.snake.x[0]-2, g.snake.y[0])
		case UP:
			g.Update(g.snake.x[0], g.snake.y[0]-1)
		case DOWN:
			g.Update(g.snake.x[0], g.snake.y[0]+1)
		}
		g.screen.Sync()
		time.Sleep(time.Second / 8)
	}
}
