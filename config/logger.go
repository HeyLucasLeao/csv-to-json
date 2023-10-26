package config

import (
	"log"
	"os"
)

func NewErrorLogger() *log.Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(os.Stdout, "ERROR🚨: ", flags)
	return logger
}

func NewInfoLogger() *log.Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(os.Stdout, "INFO✅: ", flags)
	return logger
}
