package main

import (
	"fmt"
	"os"

	"github.com/rnemeth90/csv2json/convertor"
)

func main() {
	var csvFile *os.File
	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		csvFile = os.Stdin
	} else {
		csvFile, _ = os.Open(os.Args[1])
	}

	data, err := convertor.ReadAndParseCsv(csvFile)
	if err != nil {
		panic(fmt.Sprintf("error while handling csv file: %s\n", err))
	}

	json, err := convertor.CSVToJSON(data)
	if err != nil {
		panic(fmt.Sprintf("error while converting csv to json file: %s\n", err))
	}
	fmt.Println(json)
}

func usage() {
	fmt.Println("The usage:")
}
