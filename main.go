package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "questions/problems.csv", "Path where the problems.csv file is present")
	help := flag.Bool("help", false, "Display help page")
	time := flag.Int("limit", 0, "Time limit for each question in seconds")
	flag.Parse()
	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}
	if *time > 0 {
		fmt.Println("Time limit for each question is ", *time)
	}
	data, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	problems := string(data)
	r := csv.NewReader(strings.NewReader(problems))
	score := 0
	numOfQuestions := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		numOfQuestions++
		var answer string
		fmt.Print(record[0], " ")
		fmt.Scanln(&answer)
		if answer == record[1] {
			score++
		}
	}
	fmt.Println("Your score is", score, "out of", numOfQuestions)
}
