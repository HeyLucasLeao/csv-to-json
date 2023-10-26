package config

import (
	"os"
	"strings"
)

func TruncateFolder(path string) {
	splittedString := strings.Split(path, ".")[0]

	err := os.RemoveAll(splittedString)
	if err != nil {
		panic("🚨Error trying to TruncateFolder!")
	}
}
