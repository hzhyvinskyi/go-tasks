package main

import (
	"encoding/csv"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestReadCSVFile(t *testing.T) {
	file, err := os.Open("problems.csv")
	if err != nil {
		t.Errorf("Failed to open file. Error: %s", err)
	}

	r := csv.NewReader(file)

	_, err = r.ReadAll()
	if err != nil {
		t.Errorf("Failed to read CSV file. Error: %s", err)
	}
}

func TestParseLines(t *testing.T) {
	lines := [][]string{
		[]string{"5+5", "10"},
		[]string{"1+3", "4"},
		[]string{"2+7", "9"},
	}
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	got := problems
	want := []problem{
		problem{"5+5", "10"},
		problem{"1+3", "4"},
		problem{"2+7", "9"},
	}
	
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
