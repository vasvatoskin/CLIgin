package engin

import "github.com/gdamore/tcell/v2"

func (e *Engin) handleMouseEvent(event *tcell.EventMouse) {
	if event.Buttons() == tcell.Button1 {
		x, y := event.Position()
		defStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorGold)
		e.screen.SetContent(x, y, '\u28FF', []rune{'o', 'A'}, defStyle)
	}
}

func (e *Engin) handleKeyEvent(event *tcell.EventKey) {
	switch event.Rune() {
	case 'c':
		e.screen.Clear()
	}

}
