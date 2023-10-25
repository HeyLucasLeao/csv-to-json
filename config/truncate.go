package config

import (
	"os"
	"strings"
)

func TruncateFolder(path string) error {
	splittedString := strings.Split(path, ".")[0]

	err := os.RemoveAll(splittedString)
	if err != nil {
		panic(err)
	}

	return nil
}
