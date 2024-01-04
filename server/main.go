package main

import (
	"github.com/vasvatoskin/CLIgin/server/gameServer"
	"github.com/vasvatoskin/CLIgin/server/supervisorServer"
	"github.com/vasvatoskin/CLIgin/server/webSocketServer"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	server := webSocketServer.NewServer()
	game := gameServer.New()
	supervisor := supervisorServer.New(server, game, &wg)

	supervisor.GoroutinesStart()

	http.HandleFunc("/ws", server.HandleWebSockets)
	http.ListenAndServe(":8080", nil)

	wg.Wait()
}
