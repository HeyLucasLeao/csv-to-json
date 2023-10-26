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
		panic("🚨 error loading .env file")
	}

	maxBytes, err := pipe.ConvInteger(os.Getenv("MAX_BYTES"))

	if err != nil {
		txt := fmt.Sprintf("🚨 error %s trying to create maxBytes!", err.Error())
		panic(txt)
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
