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

	maxRecords, err := pipe.ConvInteger(os.Getenv("MAX_RECORDS"))

	if err != nil {
		panic(err)
	}

	err = config.TruncateFolder()

	if err != nil {
		panic(err)
	}

	err = config.NewFolder()

	if err != nil {
		panic(err)
	}

	sf := config.StatFile{}
	j := config.NewJSON(sf.Partitions)
	fr := config.NewCSV()

	// Read the first row
	header, err := fr.Read()

	if err != nil {
		panic(err)
	}

	fmt.Print("Writting...\n")
	for {

		dataJson, err := pipe.ConvJson(fr, header)

		if err == io.EOF {

			err := pipe.CloseJson(j)

			if err != nil {
				panic(err)
			}

			break
		}

		if err != nil {
			panic(err)
		}

		bytes, err := j.WriteString(fmt.Sprintf("%s,\n", dataJson))

		if err != nil {
			panic(err)
		}

		sf.Bytes += bytes
		sf.Records++

		if sf.Records > maxRecords {
			sf.Bytes = 0
			sf.Records = 0
			sf.Partitions++

			err := pipe.CloseJson(j)

			if err != nil {
				panic(err)
			}

			j = config.NewJSON(sf.Partitions)
		}
	}

	fmt.Printf("Done! Created %d partitions\n", sf.Partitions+1)
}
