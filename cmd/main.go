package main

import (
	"Snake/engine"
	"Snake/logger"
	"Snake/ui"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Initialize the tcell s
	g := engine.NewGame()

	s := g.Screen()

	if err := s.Init(); err != nil {
		logger.Error("Failed to initialize screen")
	}

	defer s.Fini()

	// Define square parameters
	startX, startY := 5, 5 // Top-left corner of the square
	width, height := 10, 5 // Dimensions of the square

	// Style for the square
	style := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)

	// Draw the square
	ui.Draw(s, startX, startY, width, height, style)

	// Show the square
	g.GameLoop()

	// Wait for user input before exiting

}
