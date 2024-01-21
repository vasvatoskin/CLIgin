package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vasvatoskin/CLIgin/engin/engin"
	"github.com/vasvatoskin/CLIgin/internal/logger"
	"os"
)

func main() {
	logFile, err := logger.InitLogFile("logfile_ehgin.txt")
	if err != nil {
		os.Exit(1)
		return
	}
	defer logFile.Close()
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	app := &engin.Engin{}

	screen.EnableMouse()

	if err := app.Run(screen); err != nil {
		panic(err)
	}
}
