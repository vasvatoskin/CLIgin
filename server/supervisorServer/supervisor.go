package supervisorServer

import (
	"fmt"
	"github.com/vasvatoskin/CLIgin/server/gameServer"
	"github.com/vasvatoskin/CLIgin/server/webSocketServer"
	"sync"
)

type Supervisor struct {
	server *webSocketServer.Server
	game   *gameServer.Game
	wg     *sync.WaitGroup
}

func New(s *webSocketServer.Server, g *gameServer.Game, wg *sync.WaitGroup) *Supervisor {
	return &Supervisor{
		server: s,
		game:   g,
		wg:     wg,
	}
}

func (s *Supervisor) GoroutinesStart() {
	gs := []func(){
		func() {
			s.server.HandleBroadcast()
		},
		func() {
			s.game.GameLogic()
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
		case msg := <-s.server.GetIncomingChannel():
			fmt.Println("Client MSG: ", msg)
			s.game.GetIncomingChannel() <- msg
		case msg := <-s.game.GetOutgoingChannel():
			//fmt.Println("Server MSG: ", msg)
			s.server.GetOutgoingChannel() <- msg
		}

	}
}
