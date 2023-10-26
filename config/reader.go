package config

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NewFolder(path string) error {
	splittedString := strings.Split(path, ".")[0]

	err := os.MkdirAll(splittedString, os.ModePerm)

	if err != nil {
		panic("🚨Error trying to create a new folder!")
	}

	return nil
}

func NewJSON(folder string, p int) *os.File {
	jsonfile := fmt.Sprintf("data/"+folder+"/"+"part-%d.json", p)

	f, err := os.OpenFile(jsonfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic("🚨Error trying to open a new file!")
	}

	// Write an empty array to the file
	_, err = f.WriteString("[")

	if err != nil {
		panic("🚨Error writing '[' in the JSON!")
	}

	return f
}

func NewSize(f *os.File) int64 {
	// Get the current size of the file
	fileInfo, err := f.Stat()

	if err != nil {
		panic("🚨Couldn't generate fileInfo from os.File!")
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
		panic("🚨Error unexpected searching from NewFile!")
	}

	return files
}

func NewCSV(path string) *csv.Reader {
	f, err := os.Open(path)

	if err != nil {
		panic("🚨Error trying to open NewCSV!")
	}

	fr := csv.NewReader(f)
	fr.Comma = []rune(os.Getenv("CSV_COMMA"))[0]

	return fr
}
