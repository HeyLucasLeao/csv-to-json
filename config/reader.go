package config

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

var loggerError = NewErrorLogger()

func NewSize(f *os.File) int64 {
	// Get the current size of the file
	fileInfo, err := f.Stat()

	if err != nil {
		loggerError.Fatal(err)
	}

	return fileInfo.Size()
}

func NewFile() []string {
	root := "data"
	pattern := os.Getenv("CSV_FILENAME")
	files := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			m, err := filepath.Match(pattern, f.Name())

			if err != nil {
				return err
			}

			if m {
				files = append(files, path)
			}
		}

		return nil
	},
	)

	if err != nil {
		loggerError.Fatal(err)
	}

	if len(files) <= 0 {
		loggerError.Fatal(err)
	}

	return files
}

func NewCSV(path string) *csv.Reader {
	f, err := os.Open(path)

	if err != nil {
		loggerError.Fatal(err)
	}

	fr := csv.NewReader(f)
	fr.Comma = []rune(os.Getenv("CSV_COMMA"))[0]

	return fr
}
