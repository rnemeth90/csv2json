package main

import (
	"fmt"
	"os"

	"github.com/rnemeth90/csv2json/convertor"
)

func main() {
	file := os.Args[1]

	data, err := convertor.ReadAndParseCsv(file)
	if err != nil {
		panic(fmt.Sprintf("error while handling csv file: %s\n", err))
	}

	json, err := convertor.CSVToJSON(data)
	if err != nil {
		panic(fmt.Sprintf("error while converting csv to json file: %s\n", err))
	}
	fmt.Println(json)
}
