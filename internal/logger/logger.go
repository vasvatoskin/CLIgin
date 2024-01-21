package logger

import (
	"fmt"
	"log"
	"os"
)

func InitLogFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating logfile:", err)
		return nil, err
	}

	log.SetOutput(file)
	log.Println("Logger START!")
	return file, nil
}
