package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

// exit prints error message and exits programm with zero code status
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// readCSVFile reads CSV file from the supplied string and returns slice of string slices
func readCSVFile(csvFile *string) [][]string {
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open %s\n", *csvFile))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided CSV file.")
	}

	return lines
}

// parseLines parses slice of string slices into the slice of structs
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// initQuiz initializes quiz-game in the standard i/o
func initQuiz(problems []problem) (score int) {
	score = 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanln(&answer)
		if answer == p.a {
			score++
		}
	}

	return
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file in question/answer format")
	flag.Parse()
	csvFileData := readCSVFile(csvFile)
	problems := parseLines(csvFileData)
	score := initQuiz(problems)
	fmt.Printf("You scored %d out of %d\n", score, len(problems))
}
