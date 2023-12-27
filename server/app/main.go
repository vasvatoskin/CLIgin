package main

import (
	"github.com/vasvatoskin/CLIgin/server/wsSrv"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	server := wsSrv.NewServer()

	http.HandleFunc("/ws", server.HandleWebSockets)
	http.ListenAndServe(":8080", nil)

	wg.Wait()
}
