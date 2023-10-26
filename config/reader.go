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
	f := []string{}
	files, err := os.ReadDir(root)

	if err != nil {
		loggerError.Fatal(err)
	}

	for _, file := range files {
		m, err := filepath.Match(pattern, file.Name())

		if err != nil {
			loggerError.Fatal(err)
		}

		if m {
			path := filepath.Join(root, file.Name())
			f = append(f, path)
		}
	}

	if len(f) < 1 {
		loggerError.Fatal("file not found.")
	}

	return f
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
