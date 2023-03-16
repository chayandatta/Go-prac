package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Package flag implements command-line flag parsing.
	csvFileName := flag.String("csv", "problems.csv", "read the csv for question & answers")
	flag.Parse()

	// csvFileName will be the pointer to the string
	// we need to make sure we're are using actual value, not pointer
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file %s \n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Filed to parse the csv")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem no #%d: %s = \n", i+1, p.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.ans {
			correct++
		}
	}
	fmt.Printf("You have scored %d out of %d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}
	return res
}

// expect the problem to be like in this format
type problem struct {
	ques string
	ans  string
}

func exit(msg string) {
	fmt.Println(msg)
	// Typically, an exit code of 0 means that the program executes successfully.
	// Any other numerical value between 1 and 125 (golang) shows the program encountered an error.
	os.Exit(1)
}
