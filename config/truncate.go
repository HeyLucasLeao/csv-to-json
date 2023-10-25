package config

import (
	"fmt"
	"os"
	"strings"
)

func TruncateFolder() error {
	splittedString := strings.Split(os.Getenv("CSV_FILENAME"), ".")[0]
	folderPath := fmt.Sprintf("./data/" + splittedString)

	err := os.RemoveAll(folderPath)
	if err != nil {
		panic(err)
	}

	return nil
}
