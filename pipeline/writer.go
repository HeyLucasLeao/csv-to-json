package pipe

import (
	"csv-to-json/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func truncateComma(f *os.File, s int64) error {
	// Truncate the file to a new size
	err := f.Truncate(s - 2)

	if err != nil {
		return err
	}

	return nil
}

func CloseJson(f *os.File) {
	size := config.NewSize(f)
	defer f.Close()
	defer f.WriteString("]")

	truncateComma(f, size)
}

func WriteJson(path string, maxBytes int) {

	config.TruncateFolder(path)

	config.NewFolder(path)

	folder := strings.Split(filepath.Base(path), ".")[0]
	fr := config.NewCSV(path)
	sf := config.StatFile{}
	j := config.NewJSON(folder, sf.Partitions)

	// Read the first row
	header, err := fr.Read()

	if err != nil {
		panic("🚨Error trying to read row in CSV!")
	}

	for {

		row, err := fr.Read()

		if err != nil {
			panic("🚨Error trying to read row in CSV!")
		}

		if err == io.EOF {

			CloseJson(j)

			break
		}

		dataJson, err := ConvJson(row, header)

		if err != nil {
			panic("🚨Error trying to Marshal a new row!")
		}

		bytes, err := j.WriteString(fmt.Sprintf("%s,\n", dataJson))

		if err != nil {
			panic("🚨Error trying to write a JSON row!")
		}

		sf.Bytes += bytes

		if sf.Bytes > maxBytes {
			sf.Bytes = 0
			sf.Partitions++

			CloseJson(j)

			j = config.NewJSON(folder, sf.Partitions)
		}
	}
	fmt.Printf("%s done! Created %d partitions.\n", folder, sf.Partitions+1)
}
