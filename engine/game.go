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
	snake := Spawn(5, 5, 10, 5)
	return &Game{
		screen: s,
		snake:  snake,
	}
}

func (g *Game) Screen() tcell.Screen {
	return g.screen
}

func Controller(ev tcell.Event) {
	switch event := ev.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlC:
			return
		}
	}
}

func (g *Game) GameLoop() {
	// Define square parameters
	startX, startY := g.snake.x, g.snake.y         // Top-left corner of the square
	width, height := g.snake.width, g.snake.height // Dimensions of the square

	// Style for the square
	style := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)

	// Draw the square
	ui.Draw(g.screen, startX, startY, width, height, style)
	for {
		switch ev := g.screen.PollEvent().(type) {
		case *tcell.EventKey:
			Controller(ev)
			return
		case *tcell.EventResize:
			g.screen.Sync() // Handle terminal resize
		}
	}
}
