package main

import (
	"fmt"
	"javic/cmd/javic"
	"log"
	"os"
)

func main() {
	InitLogFile()

	// NOTE: Functional Code Starts Here:

	if len(os.Args) < 2 {
		fmt.Println("Usage: javic <file>")
	}

	fileName := os.Args[1]

	javic.Javic(fileName)
}

// NOTE: LOGGING
var logFileWritten bool

type logWriter struct{}

func (lw *logWriter) Write(p []byte) (n int, err error) {
	logFileWritten = true
	return os.Stdout.Write(p)
}

func InitLogFile() {
	// HACK: New Log File Location and Writing
	logFile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	defer func() {
		logFile.Close()
		fileInfo, err := os.Stat("debug.log")
		if err == nil && fileInfo.Size() == 0 {
			os.Remove("debug.log")
		} else if logFileWritten {
			fmt.Println("Errors stored in debug.log file")
		}
	}()

	logFileWritten = false
	log.SetOutput(&logWriter{})
}
