package engine

import (
	"github.com/gdamore/tcell/v2"
)

func Controller(ev tcell.Event) {
	switch event := ev.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlC:
			return
		}
	}
}

func GameLoop(s tcell.Screen) {

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			Controller(ev)
			return
		case *tcell.EventResize:
			s.Sync() // Handle terminal resize
		}
	}
}
