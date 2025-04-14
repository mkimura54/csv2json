package logic

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func ConvertToJson(csvFilePath string, option Option) (string, error) {
	cells, err := readCsvFile(csvFilePath)
	if err != nil {
		return "", err
	}

	jt := convertToJson(cells, option)
	if option.IsAutoFormat {
		return adjustJson(jt)
	} else {
		return jt, nil
	}
}

func readCsvFile(csvFilePath string) ([][]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reader *csv.Reader
	utf8, err := isUTF8(csvFilePath)
	if err != nil {
		return nil, err
	}
	if utf8 {
		reader = csv.NewReader(file)
	} else {
		r := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())
		reader = csv.NewReader(r)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func adjustJson(src string) (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(src), "", "  ")
	if err != nil {
		return "", err
	}
	return fmt.Sprint(buf.String()), nil
}
