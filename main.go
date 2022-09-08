package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type file struct {
	path string
}

func main() {
	f := file{
		path: "./questions.csv",
	}
	result, _ := readCsvFile(f)
	json := parseToJson(result)
	fmt.Println(string(json))
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

func parseToJson(records [][]string) []byte {

	if len(records) < 1 {
		log.Fatal("Something wrong, the file maybe empty or length of the lines are not the same")
	}

	headers := make([]string, 0)
	for _, v := range records[0] {
		headers = append(headers, v)
	}

	// Remove the header
	records = records[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range records {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headers[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}
		//end of object of the array
		buffer.WriteString("}")

		if i < len(records)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(`]`)

	rawMessage := json.RawMessage(buffer.String())
	x, _ := json.MarshalIndent(rawMessage, "", "  ")
	return x
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
