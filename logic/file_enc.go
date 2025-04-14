package logic

import (
	"os"
	"unicode/utf8"
)

func isUTF8(csvFilePath string) (bool, error) {
	dat, err := os.ReadFile(csvFilePath)
	if err != nil {
		return false, err
	}

	return utf8.Valid(dat), nil
}
