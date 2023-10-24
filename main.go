package main

import (
	"csv-to-json/config"
	pipe "csv-to-json/pipeline"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		return
	}

	j := config.NewJSON()
	fr := config.NewCSV()
	defer j.WriteString("]")

	// Read the first row
	header, err := fr.Read()

	if err != nil {
		panic(err)
	}

	for {

		dataJson, err := pipe.ConvJson(fr, header)

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		_, err = j.WriteString(fmt.Sprintf("%s,\n", dataJson))

		if err != nil {
			panic(err)
		}
	}

	size, err := config.NewSize(j)

	if err != nil {
		panic(err)
	}

	config.TruncateComma(j, size)
}
