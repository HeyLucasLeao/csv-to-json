package config

import (
	"fmt"
	"os"
	"strings"
)

func TruncateFolder(path string) {
	splittedString := strings.Split(path, ".")[0]

	err := os.RemoveAll(splittedString)
	if err != nil {
		txt := fmt.Sprintf("🚨Error %s trying to TruncateFolder!", err.Error())
		panic(txt)
	}
}
