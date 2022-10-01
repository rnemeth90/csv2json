// csv2json converts csv files to json and prints to stdout
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("[ERR] Missing parameter, provide a file name")
	}

	fileName := os.Args[1]
	result, err := readCsvFile(fileName)
	check(err)
	json := createOutput(result)
	fmt.Println(string(json))
}

func readCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	check(err)

	return records, nil
}

func createOutput(records [][]string) []byte {

	if len(records) < 1 {
		log.Fatal("[ERR] the file may be empty or the number of values in each line may be different")
	}

	parsedJson := parseJson(records)
	rawMessage := json.RawMessage(parsedJson)
	formattedJson, _ := json.MarshalIndent(rawMessage, "", "  ")
	return formattedJson
}

func parseJson(records [][]string) string {
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
	return buffer.String()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
