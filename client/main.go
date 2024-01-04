package main

import (
	"github.com/vasvatoskin/CLIgin/client/gameClient"
	"github.com/vasvatoskin/CLIgin/client/supervisorClient"
	"github.com/vasvatoskin/CLIgin/client/webSocketClient"
	"github.com/vasvatoskin/CLIgin/logger"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	logFile, err := logger.InitLogFile("logfile_client.txt")
	if err != nil {
		os.Exit(1)
		return
	}
	defer logFile.Close()

	wg := sync.WaitGroup{}

	client := webSocketClient.New()
	client.Connect("ws://localhost:18181/ws")

	game, err := gameClient.New()
	if err != nil {
		log.Fatal("Not create Game")
		os.Exit(1)
		return
	}

	supervisor := supervisorClient.New(client, game, &wg)

	supervisor.GoroutinesStart()
	time.Sleep(2 * time.Second)
	wg.Wait()
	log.Println("Game Closed!")
}
