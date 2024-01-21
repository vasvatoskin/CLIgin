package engin

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

type Engin struct {
	screen tcell.Screen
	width  int
	height int
}

func New() (*Engin, error) {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Println("Error creating gameClient:", err)
		return nil, err
	}

	if err := screen.Init(); err != nil {
		log.Println("Error initializing gameClient:", err)
		return nil, err
	}

	screen.Clear()

	w, h := screen.Size()
	return &Engin{
		screen: screen,
		width:  w,
		height: h,
	}, nil
}

func (e *Engin) Run(screen tcell.Screen) error {
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEsc {
				return nil
			}
			e.handleKeyEvent(ev, screen)
		case *tcell.EventMouse:
			e.handleMouseEvent(ev, screen)
		}

		screen.Show()
	}
}