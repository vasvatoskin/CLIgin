package gameServer

import (
	shared2 "github.com/vasvatoskin/CLIgin/internal/shared"
)

type Game struct {
	incomingChan chan shared2.ClientMessage
	outgoingChan chan shared2.ServerMessage
}

func New() *Game {
	return &Game{
		incomingChan: make(chan shared2.ClientMessage),
		outgoingChan: make(chan shared2.ServerMessage),
	}
}

func (g *Game) GameLogic() {
	obg := shared2.Pixel{
		Point:  shared2.Point{X: 0, Y: 0},
		Symbol: 'X',
	}
	var msgC shared2.ClientMessage
	var msgS shared2.ServerMessage
	msgS.FScreen.Pixels = make([]shared2.Pixel, 100000)
	for {
		select {
		case msgC = <-g.incomingChan:
			coun := 0
			for i := 0; i < 200; i++ {
				for j := 0; j < 100; j++ {
					msgS.FScreen.Pixels[coun] = shared2.Pixel{
						Point:  shared2.Point{X: i, Y: j},
						Symbol: '_',
					}
					coun++
				}
			}
			obg.Y += msgC.Dy
			obg.X += msgC.Dx
			msgS.Type = shared2.GameEventMessage
			obg.Symbol = 'X'
			msgS.FScreen.Pixels[coun] = obg
		}
		g.outgoingChan <- msgS
	}
}

func (g *Game) GetOutgoingChannel() chan shared2.ServerMessage { return g.outgoingChan }

func (g *Game) GetIncomingChannel() chan shared2.ClientMessage { return g.incomingChan }
