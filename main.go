package main

import (
	"csv-to-json/config"
	pipe "csv-to-json/pipeline"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	var wg sync.WaitGroup
	loggerError := config.NewErrorLogger()
	err := godotenv.Load()

	if err != nil {
		loggerError.Fatal(err)
	}

	maxBytes, err := pipe.ConvInteger(os.Getenv("MAX_BYTES"))

	if err != nil {
		loggerError.Fatal(err)
	}

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
