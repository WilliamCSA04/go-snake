package ui

import (
	"Snake/logger"

	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

func NewScreen() tcell.Screen {
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

	return screen
}

func ScreenStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
}
