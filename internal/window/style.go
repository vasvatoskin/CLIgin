package window

import "github.com/gdamore/tcell/v2"

type WindowStyle struct {
	Min       Size
	Occupancy int
	tcell.Style
}
