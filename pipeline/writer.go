package pipe

import (
	"csv-to-json/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func NewFolder(path string) error {
	splittedString := strings.Split(path, ".")[0]

	err := os.MkdirAll(splittedString, os.ModePerm)

	if err != nil {
		txt := fmt.Sprintf("🚨Error %s trying to create a new folder!", err.Error())
		panic(txt)
	}

	return nil
}

func NewJSON(folder string, p int) *os.File {
	jsonfile := fmt.Sprintf("data/"+folder+"/"+"part-%d.json", p)

	f, err := os.OpenFile(jsonfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		txt := fmt.Sprintf("🚨Error %s trying to open a new file!", err.Error())
		panic(txt)
	}

	// Write an empty array to the file
	_, err = f.WriteString("[")

	if err != nil {
		txt := fmt.Sprintf("🚨Error %s writing '[' in the JSON!", err.Error())
		panic(txt)
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
		txt := fmt.Sprintf("🚨Error %s when trying to read row in CSV!", err.Error())
		panic(txt)
	}

	for {

		row, err := fr.Read()

		if err == io.EOF {

			CloseJson(j)

			break
		}

		if err != nil {
			txt := fmt.Sprintf("🚨Error %s when trying to read row in CSV!", err.Error())
			panic(txt)
		}

		dataJson, err := ConvJson(row, header)

		if err != nil {
			txt := fmt.Sprintf("🚨Error %s trying to Marshal a new row!", err.Error())
			panic(txt)
		}

		bytes, err := j.WriteString(fmt.Sprintf("%s,\n", dataJson))

		if err != nil {
			txt := fmt.Sprintf("🚨Error %s trying to write a JSON row!", err.Error())
			panic(txt)
		}

		sf.Bytes += bytes

		if sf.Bytes > maxBytes {
			sf.Bytes = 0
			sf.Partitions++

			CloseJson(j)

			j = NewJSON(folder, sf.Partitions)
		}
	}
	fmt.Printf("%s done! Created %d partitions.\n", folder, sf.Partitions+1)
}
