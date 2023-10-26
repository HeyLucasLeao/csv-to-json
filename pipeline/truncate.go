package pipe

import (
	"fmt"
	"os"
	"strings"
)

func TruncateFolder(path string) {
	splittedString := strings.Split(path, ".")[0]

	err := os.RemoveAll(splittedString)
	if err != nil {
		txt := fmt.Sprintf("🚨 error %s trying to TruncateFolder!", err.Error())
		panic(txt)
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
