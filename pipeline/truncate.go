package pipe

import (
	"os"
	"strings"
)

func TruncateFolder(path string) {
	splittedString := strings.Split(path, ".")[0]

	err := os.RemoveAll(splittedString)
	if err != nil {
		logger.Fatal(err)
	}
}

func TruncateComma(f *os.File, s int64) error {
	// Truncate the file to a new size
	err := f.Truncate(s - 2)

	if err != nil {
		return err
	}

	return nil
}
