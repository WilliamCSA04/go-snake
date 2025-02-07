package main

import (
	"Snake/engine"
	"Snake/logger"
	"Snake/ui"

	"github.com/gdamore/tcell/v2"
)

func main() {
	logger.Info("Starting Snake")
	screen := ui.NewScreen()

	logger.Info("Setting styles")
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screenStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	logger.Info("Defining styles")
	screen.SetStyle(screenStyle)

	logger.Info("Drawing")
	ui.Draw(screen, 1, 1, 42, 7, style)

	logger.Info("Declaring quit function")
	quit := func() {
		logger.Error("Quitting")
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	engine.GameLoop(screen)
}
