package main

import (
	"csv-to-json/config"
	pipe "csv-to-json/pipeline"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	var wg sync.WaitGroup

	err := godotenv.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		return
	}

	maxRecords, err := pipe.ConvInteger(os.Getenv("MAX_RECORDS"))

	if err != nil {
		panic(err)
	}

	files := config.NewFile()

	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			pipe.WriteJson(file, maxRecords)
		}(file)
	}
	wg.Wait()

}
