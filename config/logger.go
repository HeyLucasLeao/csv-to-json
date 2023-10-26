package config

import (
	"log"
	"os"
)

func CreateLogger() *log.Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(os.Stdout, "ERROR🚨: ", flags)
	return logger
}
