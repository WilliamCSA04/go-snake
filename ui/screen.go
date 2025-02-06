package ui

import (
	"Snake/logger"

	"github.com/gdamore/tcell/v2"
)

type Screen struct {
	screen *tcell.Screen
}

var screen *Screen

func NewScreen() *Screen {
	logger.Info("Creating screen")
	if screen != nil {
		logger.Info("Screen already created")
		return screen
	}
	screen, err := tcell.NewScreen()

	if err != nil {
		logger.Error("Failed to create screen")
		panic(err)
	}

	return &Screen{screen: &screen}
}
