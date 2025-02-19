package engine

import (
	"Snake/ui"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

const (
	START     = iota
	ON_GOING  = iota
	GAME_OVER = iota
)

type Game struct {
	screen tcell.Screen
	snake  *Snake
	food   *Food
	state  int
}

var (
	width, height int
)

func NewGame() *Game {
	width, height = 2, 1
	s := ui.NewScreen()
	snake := SpawnSnake(6, 6, 2, 1)
	f := SpawnFood(0, 0, 2, 1)
	return &Game{
		screen: s,
		snake:  snake,
		food:   f,
		state:  START,
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
			g.Update(g.snake.x[0]-2, g.snake.y[0])
		case tcell.KeyRight:
			g.Update(g.snake.x[0]+2, g.snake.y[0])
		case tcell.KeyUp:
			g.Update(g.snake.x[0], g.snake.y[0]-1)
		case tcell.KeyDown:
			g.Update(g.snake.x[0], g.snake.y[0]+1)
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
	for {
		switch ev := g.screen.PollEvent().(type) {
		case *tcell.EventKey:
			if !g.Controller(ev) {
				return
			}
		default:
			g.screen.Sync() // Handle terminal resize
		}
		if g.state == GAME_OVER {
			return
		}
		g.screen.Sync()
	}
}
