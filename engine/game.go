package engine

import (
	"Snake/logger"

	"github.com/gdamore/tcell/v2"
)

func GameLoop(s tcell.Screen) {
	// Event loop
	logger.Info("Entering event loop")
	err := s.Init()
	if err != nil {
		logger.Error("Failed to initialize screen")
		panic(err)
	}
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()
		if ev != nil {
			logger.Info("Event received")
			switch ev := ev.(type) {
			case *tcell.EventKey:
				logger.Info("Key event received")
				switch ev.Key() {
				case tcell.KeyEscape:
					logger.Info("Escape key pressed")
					s.Fini()
					return
				}
			}
		}

	}
}
