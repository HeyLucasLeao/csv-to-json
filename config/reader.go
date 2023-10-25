package config

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func NewFolder() error {
	splittedString := strings.Split(os.Getenv("CSV_FILENAME"), ".")[0]
	folderPath := fmt.Sprintf("./data/" + splittedString)

	err := os.MkdirAll(folderPath, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

func TruncateFolder() error {
	splittedString := strings.Split(os.Getenv("CSV_FILENAME"), ".")[0]
	folderPath := fmt.Sprintf("./data/" + splittedString)

	err := os.RemoveAll(folderPath)
	if err != nil {
		panic(err)
	}

	return nil
}

func NewJSON(p int) *os.File {
	splittedString := strings.Split(os.Getenv("CSV_FILENAME"), ".")[0]
	fileName := fmt.Sprintf(splittedString+"-%d.json", p)
	folderPath := fmt.Sprintf("./data/" + splittedString)

	filePath := fmt.Sprintf(folderPath+"/%s", fileName)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err := os.Truncate(filePath, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	if err != nil {
		panic(err)
	}

	// Write an empty array to the file
	_, err = f.WriteString("[")

	if err != nil {
		panic(err)
	}

	return f
}

func NewCSV() *csv.Reader {
	filepath := fmt.Sprintf("./data/%s", os.Getenv("CSV_FILENAME"))
	f, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	fr := csv.NewReader(f)
	fr.Comma = []rune(os.Getenv("CSV_COMMA"))[0]

	if err != nil {
		panic(err)
	}

	return fr
}

func NewSize(f *os.File) (int64, error) {
	// Get the current size of the file
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return fileInfo.Size(), nil
}

func TruncateComma(f *os.File, s int64) error {
	// Truncate the file to a new size
	err := f.Truncate(s - 2)

	if err != nil {
		return err
	}

	return nil
}
