package main

import (
	"github.com/vasvatoskin/CLIgin/engin/engin"
	"github.com/vasvatoskin/CLIgin/internal/logger"
	"log"
	"os"
)

func main() {
	logFile, err := logger.InitLogFile("logfile_ehgin.txt")
	if err != nil {
		os.Exit(1)
		return
	}
	defer logFile.Close()

	engin, err := engin.New()
	if err != nil {
		log.Fatal("Not create Engin")
		os.Exit(1)
		return
	}

	if err := engin.Run(); err != nil {
		panic(err)
	}
}
