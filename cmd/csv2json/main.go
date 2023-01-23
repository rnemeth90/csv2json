package main

import (
	"fmt"
	"os"

	"github.com/rnemeth90/csv2json"
)

func main() {
	var csvFile *os.File
	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		csvFile = os.Stdin
	} else {
		if len(os.Args) < 2 {
			usage()
			os.Exit(1)
		}
		csvFile, _ = os.Open(os.Args[1])
	}

	data, err := csv2json.ReadAndParseCsv(csvFile)
	if err != nil {
		panic(fmt.Sprintf("error while reading csv file: %s\n", err))
	}

	json, err := csv2json.CSVToJSON(data)
	if err != nil {
		panic(fmt.Sprintf("error while converting csv to json file: %s\n", err))
	}
	fmt.Println(json)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s file.csv\n\n", os.Args[0])
	fmt.Println("*** You can also pass in CSV via stdin")
}
