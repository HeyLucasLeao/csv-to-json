package main

import (
	"csv-to-json/config"
	pipe "csv-to-json/pipeline"
	"os"
	"sync"

	"github.com/dustin/go-humanize"
	"github.com/joho/godotenv"
)

var loggerError = config.NewErrorLogger()
var loggerInfo = config.NewInfoLogger()

func main() {
	var wg sync.WaitGroup
	err := godotenv.Load()

	if err != nil {
		loggerError.Fatal(err)
	}

	maxBytes, err := pipe.ConvInteger(os.Getenv("MAX_BYTES"))

	if err != nil {
		loggerError.Fatal(err)
	}

	loggerInfo.Printf("Partitioning into JSON files with a maximum size of %s", humanize.Bytes(uint64(maxBytes)))

	files := config.NewFile()

	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			pipe.WriteJson(file, int(maxBytes))
		}(file)
	}
	wg.Wait()

}
