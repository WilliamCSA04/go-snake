package main

import (
	"Snake/engine"
	"Snake/logger"
)

func main() {
	// Initialize the tcell s
	g := engine.NewGame()

	s := g.Screen()

	if err := s.Init(); err != nil {
		logger.Error("Failed to initialize screen")
	}

	defer s.Fini()

	// Show the square
	g.GameLoop()

	// Wait for user input before exiting

}
