package pipe

import (
	"csv-to-json/config"
	"os"
)

func truncateComma(f *os.File, s int64) error {
	// Truncate the file to a new size
	err := f.Truncate(s - 2)

	if err != nil {
		return err
	}

	return nil
}

func CloseJson(f *os.File) error {
	size, err := config.NewSize(f)
	defer f.Close()
	defer f.WriteString("]")

	if err != nil {
		return err
	}

	truncateComma(f, size)

	return nil
}
