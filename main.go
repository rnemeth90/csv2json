package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type file struct {
	path string
}

func main() {
	f := file{
		path: "./questions.csv",
	}
	result, _ := readCsvFile(f)
	fmt.Println(result)
}

func readCsvFile(f file) ([][]string, error) {
	file, err := os.Open(f.path)
	check(err)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	check(err)

	return records, nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
