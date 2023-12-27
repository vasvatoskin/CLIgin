package supvSrv

import (
	"github.com/vasvatoskin/CLIgin/server/gmSrv"
	"github.com/vasvatoskin/CLIgin/server/wsSrv"
	"sync"
)

type Supervisor struct {
	server *wsSrv.Server
	game   *gmSrv.Game
	wg     *sync.WaitGroup
}

func New(s *wsSrv.Server, g *gmSrv.Game, wg *sync.WaitGroup) *Supervisor {
	return &Supervisor{
		server: s,
		game:   g,
		wg:     wg,
	}
}

func (s *Supervisor) goroutinesStart() {
	gs := []func(){
		func() {
			s.server.HandleBroadcast()
		},
		func() {
			s.game.GameLogic()
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
