package quizzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Quiz(questions []string) (answers []string, err error) {
	reader := bufio.NewReader(os.Stdin)
	for _, q := range questions {
		fmt.Println(q)
		text, err := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if err != nil {
			return []string{}, err
		}
		answers = append(answers, text)
	}
	return
}

func Grade(correctAnswers []string, userAnswers []string) {
	numberCorrect := 0
	numberQuestions := len(correctAnswers)
	for i, a := range userAnswers {
		if a == correctAnswers[i] {
			numberCorrect++
		}
	}
	fmt.Printf("You answered %d questions correct out of %d", numberCorrect, numberQuestions)
}