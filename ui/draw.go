package ui

import "github.com/gdamore/tcell/v2"

func Draw(s tcell.Screen, startX int, startY int, width int, height int, style tcell.Style) {
	for y := startY; y < startY+height; y++ {
		for x := startX; x < startX+width; x++ {
			s.SetContent(x, y, ' ', nil, style)
		}
	}
}
