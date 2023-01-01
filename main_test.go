package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	binName string = "csv2json"
	fileName string = "example.csv"
	expected string = `
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
)


func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go","build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	resultCode := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)

	os.Exit(resultCode)
}

func TestCsv2JsonCli(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir,binName)

	t.Run("TestConvertCli", func(t *testing.T){
		cmd := exec.Command(cmdPath, fileName)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		require.JSONEq(t, expected, string(out))
	})
}
