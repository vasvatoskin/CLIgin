package engin

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vasvatoskin/CLIgin/internal/shared"
	"github.com/vasvatoskin/CLIgin/internal/window"
)

func (e *Engin) Run() error {
	start := shared.Point{0, 0}
	style := window.WindowStyle{
		Min:       window.Size{5, 5},
		Occupancy: 30,
		Style:     tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorBlack),
	}
	win := window.New(start, style, window.Size(e.Size))
	win.Draw(e.screen)
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			e.screen.Sync()
			win.Actual.CalculateActual(window.Size(e.Size), style)
			win.Draw(e.screen)
			e.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEsc {
				e.screen.Fini()
				return nil
			}
			e.handleKeyEvent(ev)
		case *tcell.EventMouse:
			e.handleMouseEvent(ev)
		}

		e.screen.Show()
	}
}
