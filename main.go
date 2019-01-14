package main

import (
	"flag"
	"github.com/colinjlacy/quizGame/quizzer"
	"github.com/colinjlacy/quizGame/reader"
	"github.com/colinjlacy/quizGame/splitter"
	"log"
)

func main() {
	file := flag.String("f", "./test-set.csv", "the source file for quiz questions")
	flag.Parse()

	err := reader.CheckFile(*file)
	if err != nil {
		log.Fatalf("could not access specified file: %s", err)
	}
	records, err := reader.ReadFile(*file)
	if err != nil {
		log.Fatalf("could not read specified file: %s", err)
	}
	questions, correctAnswers := splitter.Split(records)
	userAnswers, err := quizzer.Quiz(questions)
	quizzer.Grade(correctAnswers, userAnswers)
}
