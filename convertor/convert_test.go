package convertor

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var sliceControl = [][]string{
	[]string{"Year", "Score", "Title"},
	[]string{"1968", "86", "Greetings"},
	[]string{"1970", "17", "Bloody Mama"},
	[]string{"1970", "73", "Hi,Mom!"},
	[]string{"1971", "40", "Born to Win"},
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
		fileName, _ := os.Open("UnitTestExample.csv")
		want := sliceControl
		got, _ := ReadAndParseCsv(fileName)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("[Test] CSVToJSON()", func(t *testing.T) {
		want := jsonControl
		got, _ := CSVToJSON(sliceControl)

		require.JSONEq(t, want, got)
	})
}

func BenchmarkConvertor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CSVToJSON(sliceControl)
	}
}
