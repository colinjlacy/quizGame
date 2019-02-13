package quizzer

import (
	"fmt"
	"time"
)

func Quiz(questions []string, seconds time.Duration) (answers []string, err error) {
	timer := time.NewTimer(seconds * time.Second)

	fmt.Println("Press enter to being...")
	_, _ = fmt.Scanf("%s\n", nil)

quizloop:
	for _, q := range questions {
		fmt.Println(q)
		answerCh := make(chan string)

		go func() {
			var text string
			_, err = fmt.Scanf("%s\n", &text)
			answerCh <- text
		}()

		select {
		case <-timer.C:
			fmt.Println("Timer expired")
			break quizloop
		case a := <-answerCh:
			answers = append(answers, a)
		}
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
