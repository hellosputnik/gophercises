package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

type Quiz struct {
	score            int
	numberOfProblems int
}

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fin := bufio.NewScanner(file)

	quiz := Quiz{}
	for fin.Scan() {
		quiz.numberOfProblems++

		line := strings.Split(fin.Text(), ",")
		problem := Problem{question: line[0], answer: line[1]}

		fmt.Printf("Problem #%d: %s = ", quiz.numberOfProblems, problem.question)

		var input string
		fmt.Scan(&input)

		if input == problem.answer {
			quiz.score++
		}
	}

	fmt.Printf("You scored %d out of %d.", quiz.score, quiz.numberOfProblems)
}
