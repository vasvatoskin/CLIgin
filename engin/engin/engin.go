package engin

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/vasvatoskin/CLIgin/internal/window"
)

type Engin struct {
	screen tcell.Screen
	window.Size
}

func New() (*Engin, error) {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Println("Error creating Engin:", err)
		return nil, err
	}

	if err := screen.Init(); err != nil {
		log.Println("Error initializing Engin:", err)
		return nil, err
	}

	screen.EnableMouse()
	screen.Clear()

	w, h := screen.Size()
	return &Engin{
		screen: screen,
		Size:   window.Size{Width: w, Height: h},
	}, nil
}
