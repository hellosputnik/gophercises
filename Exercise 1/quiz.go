package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

type Quiz struct {
	problems []Problem
	score    int
}

type Settings struct {
	filename  *string
	timeLimit *int
}

func main() {
	quiz := Quiz{}
	settings := Settings{}

	settings.filename = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	settings.timeLimit = flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*settings.filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fin := bufio.NewScanner(file)

	for fin.Scan() {
		line := strings.Split(fin.Text(), ",")
		problem := Problem{question: line[0], answer: line[1]}

		quiz.problems = append(quiz.problems, problem)
	}

	timer := time.NewTimer(time.Second * time.Duration(*settings.timeLimit))
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\nYou scored %d out of %d.", quiz.score, len(quiz.problems))
	}()

	for i, problem := range quiz.problems {
		fmt.Printf("Problem #%d: %s = ", (i + 1), problem.question)

		var input string
		fmt.Scan(&input)

		if input == problem.answer {
			quiz.score++
		}
	}
	fmt.Printf("You scored %d out of %d.", quiz.score, len(quiz.problems))
}
