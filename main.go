package main

import (
	"flag"
	"github.com/colinjlacy/quizGame/quizzer"
	"github.com/colinjlacy/quizGame/reader"
	"github.com/colinjlacy/quizGame/splitter"
	"log"
	"time"
)

func main() {
	file := flag.String("f", "./test-set.csv", "the source file for quiz questions")
	timer := flag.Int("t", 30, "the time limit for the quiz")
	flag.Parse()

	err := reader.CheckFile(*file)
	if err != nil {
		log.Fatalf("could not access specified file: %s", err)
	}
	dur := time.Duration(*timer)
	records, err := reader.ReadFile(*file)
	if err != nil {
		log.Fatalf("could not read specified file: %s", err)
	}
	questions, correctAnswers := splitter.Split(records)
	userAnswers, err := quizzer.Quiz(questions, dur)
	quizzer.Grade(correctAnswers, userAnswers)
}
