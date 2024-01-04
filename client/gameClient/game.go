package gameClient

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vasvatoskin/CLIgin/shared"
	"log"
)

type Game struct {
	screen       tcell.Screen
	width        int
	height       int
	id           uint64
	defStyle     tcell.Style
	incomingChan chan shared.ServerMessage
	outgoingChan chan shared.ClientMessage
	close        chan bool
}

func New() (*Game, error) {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Println("Error creating gameClient:", err)
		return nil, err
	}

	if err := screen.Init(); err != nil {
		log.Println("Error initializing gameClient:", err)
		return nil, err
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	screen.SetStyle(defStyle)
	screen.Clear()

	w, h := screen.Size()
	return &Game{
		screen:       screen,
		width:        w,
		height:       h,
		defStyle:     defStyle,
		incomingChan: make(chan shared.ServerMessage),
		outgoingChan: make(chan shared.ClientMessage),
		close:        make(chan bool),
	}, nil
}

func (g *Game) Close() {
	g.screen.Fini()
	close(g.close)
}

func (g *Game) EventHandler() {

	msg := shared.ClientMessage{}
	msg.Type = shared.GameEventMessage
	var vector shared.Vector
	for {
		if ev := g.screen.PollEvent(); ev != nil {
			switch ev := ev.(type) {

			case *tcell.EventResize:
				g.screen.Sync()
				g.width, g.height = g.screen.Size()

			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					msg.Type = shared.DisconnectMessage
				} else if ev.Key() == tcell.KeyUp || ev.Rune() == 'w' {
					vector.SetDy(-1)
				} else if ev.Key() == tcell.KeyDown || ev.Rune() == 's' {
					vector.SetDy(1)
				} else if ev.Key() == tcell.KeyLeft || ev.Rune() == 'a' {
					vector.SetDx(-1)
				} else if ev.Key() == tcell.KeyRight || ev.Rune() == 'd' {
					vector.SetDx(1)
				}
				msg.Vector = vector
				g.outgoingChan <- msg
			}

		} else {
			select {

			case <-g.close:
				log.Println("EventHandler closed")
				return
			}
		}
	}
}

func (g *Game) Render() {
	var msg shared.ServerMessage
	g.screen.Clear()
	for {
		select {

		case msg = <-g.incomingChan:
			switch msg.Type {
			case shared.GameEventMessage:
				g.screen.Clear()
				for _, pixel := range msg.FScreen.Pixels {
					g.screen.SetContent(pixel.X, pixel.Y, pixel.Texture, nil, g.defStyle)
				}
			}
			g.screen.Show()

		case <-g.close:
			log.Println("Render closed")
			return
		}
	}
}

func (g *Game) GetOutgoingChannel() chan shared.ClientMessage {
	return g.outgoingChan
}

func (g *Game) GetIncomingChannel() chan shared.ServerMessage {
	return g.incomingChan
}
