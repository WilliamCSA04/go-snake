package engine

import (
	"Snake/ui"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen tcell.Screen
	snake  *Snake
}

func NewGame() *Game {
	s := ui.NewScreen()
	snake := SpawnSnake(5, 5, 10, 5)
	return &Game{
		screen: s,
		snake:  snake,
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
			g.Update(g.snake.Coords.x-1, g.snake.Coords.y)
		case tcell.KeyRight:
			g.Update(g.snake.Coords.x+1, g.snake.Coords.y)
		case tcell.KeyUp:
			g.Update(g.snake.Coords.x, g.snake.Coords.y-1)
		case tcell.KeyDown:
			g.Update(g.snake.Coords.x, g.snake.Coords.y+1)
		}
	}
	return true
}

func (g *Game) Update(x, y int) {
	// Style for the square
	style := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
	width, height := 10, 5
	// Draw the square
	ui.Draw(g.screen, x, y, width, height, style)
	g.snake.Move(x, y)
}

func (g *Game) GameLoop() {

	g.Update(g.snake.Coords.x, g.snake.Coords.y)

	for {
		switch ev := g.screen.PollEvent().(type) {
		case *tcell.EventKey:
			if !g.Controller(ev) {
				return
			}
		default:
			g.screen.Sync() // Handle terminal resize
		}
		g.screen.Sync()
	}
}
