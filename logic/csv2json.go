package logic

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
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

	reader := csv.NewReader(file)
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
