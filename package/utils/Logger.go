package utils

import (
	"log"
	"os"
)

// Logger represents a logger instance.
var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an informational message.
func Info(message string) {
	Logger.Println("INFO: " + message)
}

// Error logs an error message.
func Error(err error) {
	Logger.Println("ERROR: " + err.Error())
}
