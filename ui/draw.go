package ui

import "github.com/gdamore/tcell/v2"

func Draw(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	row := y1
	col := x1
	s.SetContent(row, col, 0, nil, style)
}
