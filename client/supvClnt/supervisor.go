package supvClnt

import (
	"github.com/vasvatoskin/CLIgin/client/gmCint"
	"github.com/vasvatoskin/CLIgin/client/wsClnt"
	"github.com/vasvatoskin/CLIgin/shared"
	"sync"
)

type Supervisor struct {
	client *wsClnt.Client
	game   *gmCint.Game
	wg     *sync.WaitGroup
}

func New(c *wsClnt.Client, g *gmCint.Game, wg *sync.WaitGroup) *Supervisor {
	return &Supervisor{
		client: c,
		game:   g,
		wg:     wg,
	}
}

func (s *Supervisor) GoroutinesStart() {

	gs := []func(){
		func() {
			s.client.ReceiveServerMsg()
		},
		func() {
			s.client.SendServerMsg()
		},
		func() {
			s.game.EventHandler()
		},
		func() {
			s.game.Render()
		},
		func() {
			s.Router()
		},
	}

	for _, gr := range gs {
		s.wg.Add(1)
		go func(g func()) {
			defer s.wg.Done()
			g()
		}(gr)
	}
}

func (s *Supervisor) Router() {
	for {
		select {
		case msg := <-s.client.GetIncomingChannel():
			if msg.Type == shared.DisconnectMessage {
				s.client.Disconnect()
				s.game.Close()
				return
			}
			s.game.GetIncomingChannel() <- msg
		case msg := <-s.game.GetOutgoingChannel():
			if msg.Type == shared.DisconnectMessage {
				s.client.Disconnect()
				s.game.Close()
				return
			}
			s.client.GetOutgoingChannel() <- msg

		}

	}
}
