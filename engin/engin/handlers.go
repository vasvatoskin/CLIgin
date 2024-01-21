package engin

import "github.com/gdamore/tcell/v2"

func (e *Engin) handleMouseEvent(event *tcell.EventMouse, s tcell.Screen) {
	if event.Buttons() == tcell.Button1 {
		// Обработка нажатия левой кнопки мыши
		x, y := event.Position()
		defStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorGold)
		s.SetContent(x, y, '\u28FF', []rune{'o', 'A'}, defStyle)
		// Далее ваш код обработки события
	}
}

func (e *Engin) handleKeyEvent(event *tcell.EventKey, s tcell.Screen) {
	switch event.Rune() {
	case 'c':
		s.Clear()
	}

}
