package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mkimura54/csv2json/logic"
)

const (
	PROGRAM_NAME = "csv2json"
	VERSION      = "v.0.2"

	VALID_ARGS_MIN_LEN = 2
)

func main() {
	version := flag.Bool("v", false, "バージョン情報を表示する")
	autoFormat := flag.Bool("a", false, "JSONの自動整形を行う")
	considerType := flag.Bool("t", false, "型を判定して適切に変換する")
	outputBlank := flag.Bool("b", false, "ブランクデータ部も出力する")
	flag.Parse()

	if *version {
		fmt.Println(PROGRAM_NAME + " " + VERSION)
		return
	}

	csvFilePath := ""
	if len(os.Args) >= VALID_ARGS_MIN_LEN {
		csvFilePath = os.Args[len(os.Args)-1]
		if !strings.Contains(csvFilePath, ".csv") {
			csvFilePath = ""
		}
	}
	if csvFilePath == "" {
		fmt.Println("Please specify the file path as an argument.")
		return
	}

	op := logic.Option{
		IsAutoFormat:   *autoFormat,
		IsConsiderType: *considerType,
		IsOutputBlank:  *outputBlank,
	}
	json, err := logic.ConvertToJson(csvFilePath, op)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		return
	}
	fmt.Println(json)
}
