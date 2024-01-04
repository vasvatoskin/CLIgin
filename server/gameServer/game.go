package gameServer

import "github.com/vasvatoskin/CLIgin/shared"

type Game struct {
	incomingChan chan shared.ClientMessage
	outgoingChan chan shared.ServerMessage
}

func New() *Game {
	return &Game{
		incomingChan: make(chan shared.ClientMessage),
		outgoingChan: make(chan shared.ServerMessage),
	}
}

func (g *Game) GameLogic() {
	obg := shared.Pixel{
		Point:   shared.Point{X: 0, Y: 0},
		Texture: 'X',
	}
	var msgC shared.ClientMessage
	var msgS shared.ServerMessage
	msgS.FScreen.Pixels = make([]shared.Pixel, 100000)
	for {
		select {
		case msgC = <-g.incomingChan:
			coun := 0
			for i := 0; i < 200; i++ {
				for j := 0; j < 100; j++ {
					msgS.FScreen.Pixels[coun] = shared.Pixel{
						Point:   shared.Point{X: i, Y: j},
						Texture: '_',
					}
					coun++
				}
			}
			obg.Y += msgC.Dy
			obg.X += msgC.Dx
			msgS.Type = shared.GameEventMessage
			obg.Texture = 'X'
			msgS.FScreen.Pixels[coun] = obg
		}
		g.outgoingChan <- msgS
	}
}

func (g *Game) GetOutgoingChannel() chan shared.ServerMessage { return g.outgoingChan }

func (g *Game) GetIncomingChannel() chan shared.ClientMessage { return g.incomingChan }
