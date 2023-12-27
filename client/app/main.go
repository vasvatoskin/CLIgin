package main

import (
	"github.com/vasvatoskin/CLIgin/client/gmCint"
	"github.com/vasvatoskin/CLIgin/client/supvClnt"
	"github.com/vasvatoskin/CLIgin/client/wsClnt"
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

	client := wsClnt.New()
	client.Connect("ws://localhost:8080/ws")

	game, err := gmCint.New()
	if err != nil {
		log.Fatal("Not create Game")
		os.Exit(1)
		return
	}

	supervisor := supvClnt.New(client, game, &wg)

	supervisor.GoroutinesStart()
	time.Sleep(2 * time.Second)
	wg.Wait()
	log.Println("Game Closed!")
}
