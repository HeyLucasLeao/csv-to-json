package pipe

import (
	"csv-to-json/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var loggerError = config.NewErrorLogger()
var loggerInfo = config.NewInfoLogger()

func NewFolder(path string) error {
	splittedString := strings.Split(path, ".")[0]

	err := os.MkdirAll(splittedString, os.ModePerm)

	if err != nil {
		loggerError.Fatal(err)
	}

	return nil
}

func NewJSON(folder string, p int) *os.File {
	jsonfile := fmt.Sprintf("data/"+folder+"/"+"part-%d.json", p)

	f, err := os.OpenFile(jsonfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		loggerError.Fatal(err)
	}

	// Write an empty array to the file
	_, err = f.WriteString("[")

	if err != nil {
		loggerError.Fatal(err)
	}

	return f
}

func CloseJson(f *os.File) {
	size := config.NewSize(f)
	defer f.Close()
	defer f.WriteString("]")

	TruncateComma(f, size)
}

func WriteJson(path string, maxBytes int) {

	TruncateFolder(path)

	NewFolder(path)

	folder := strings.Split(filepath.Base(path), ".")[0]
	fr := config.NewCSV(path)
	sf := config.StatFile{}
	j := NewJSON(folder, sf.Partitions)

	// Read the first row
	header, err := fr.Read()

	if err != nil {
		loggerError.Fatal(err)
	}

	for {

		row, err := fr.Read()

		if err == io.EOF {

			CloseJson(j)

			break
		}

		if err != nil {
			loggerError.Fatal(err)
		}

		dataJson, err := ConvJson(row, header)

		if err != nil {
			loggerError.Fatal(err)
		}

		bytes, err := j.WriteString(fmt.Sprintf("%s,\n", dataJson))

		if err != nil {
			loggerError.Fatal(err)
		}

		sf.Bytes += bytes

		if sf.Bytes > maxBytes {
			sf.Bytes = 0
			sf.Partitions++

			CloseJson(j)

			j = NewJSON(folder, sf.Partitions)
		}
	}
	loggerInfo.Printf("%s done! Created %d partitions.\n", folder, sf.Partitions+1)
}
