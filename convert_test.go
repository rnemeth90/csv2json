package csv2json_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/rnemeth90/csv2json"
	"github.com/stretchr/testify/require"
)

var sliceControl = [][]string{
	{"Year", "Score", "Title"},
	{"1968", "86", "Greetings"},
	{"1970", "17", "Bloody Mama"},
	{"1970", "73", "Hi,Mom!"},
	{"1971", "40", "Born to Win"},
}

var jsonControl = `
[
  {
    "Score": "86",
    "Title": "Greetings",
    "Year": "1968"
  },
  {
    "Score": "17",
    "Title": "Bloody Mama",
    "Year": "1970"
  },
  {
    "Score": "73",
    "Title": "Hi,Mom!",
    "Year": "1970"
  },
  {
    "Score": "40",
    "Title": "Born to Win",
    "Year": "1971"
  }
]
`

func TestConvertor(t *testing.T) {
	t.Run("[Test] ReadAndParseCSV()", func(t *testing.T) {
		fileName, _ := os.Open("./testData/UnitTestExample.csv")
		want := sliceControl
		got, _ := csv2json.ReadAndParseCsv(fileName)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("[Test] CSVToJSON()", func(t *testing.T) {
		want := jsonControl
		got, _ := csv2json.CSVToJSON(sliceControl)

		require.JSONEq(t, want, got)
	})
}

func BenchmarkConvertor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		csv2json.CSVToJSON(sliceControl)
	}
}
