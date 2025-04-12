package logic

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToJson(cells [][]string, option Option) string {
	csv := analyzeCsvLines(cells)
	return createAllJson(csv, option)
}

func analyzeCsvLines(cells [][]string) CSV {
	result := CSV{}
	for i, line := range cells {
		if i == 0 {
			result.Header = line
		} else {
			result.Body = append(result.Body, line)
		}
	}
	return result
}

func createAllJson(csv CSV, option Option) string {
	json := "["
	for _, body := range csv.Body {
		json += createDataJson(csv.Header, body, option)
	}
	json = strings.TrimRight(json, ",")
	json += "]"

	return json
}

func createDataJson(header, body []string, option Option) string {
	json := "{"
	for i, b := range body {
		if i >= len(header) {
			continue
		}
		name := header[i]
		if !option.IsOutputBlank && b == "" {
			continue
		}
		if option.IsConsiderType {
			json += convertType(name, b)
		} else {
			json += fmt.Sprintf("\"%s\": \"%s\",", name, b)
		}
	}
	json = strings.TrimRight(json, ",")
	json += "},"

	return json
}

func convertType(name, value string) string {
	b, ok := getBool(value)
	if ok {
		return fmt.Sprintf("\"%s\": %t,", name, b) // 論理値
	}

	n, ok := getNumber(value)
	if ok {
		return fmt.Sprintf("\"%s\": %d,", name, n) // 整数
	}

	d, ok := getDecimal(value)
	if ok {
		return fmt.Sprintf("\"%s\": %g,", name, d) // 小数
	}

	return fmt.Sprintf("\"%s\": \"%s\",", name, value) // 文字列
}

func getBool(value string) (bool, bool) {
	if isTrue(value) {
		return true, true
	} else if isFalse(value) {
		return false, true
	}

	return false, false
}

func isTrue(value string) bool {
	return strings.ToLower(value) == "true"
}

func isFalse(value string) bool {
	return strings.ToLower(value) == "false"
}

func getNumber(value string) (int64, bool) {
	if strings.Contains(value, ".") {
		return 0, false
	}

	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, false
	}
	return n, true
}

func getDecimal(value string) (float64, bool) {
	if !strings.Contains(value, ".") {
		return 0, false
	}

	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return n, true
}
