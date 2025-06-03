package utils

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func InitLogger() {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening the file:", err)
	}

	InfoLog = log.New(logFile, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(logFile, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}
