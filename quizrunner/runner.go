// Runs the quiz and prints result
package quizrunner

import (
	"errors"
	"fmt"
	"quiz/client"
)

type QuizRunner struct {
	Reader client.QuizReader
}

func (qr *QuizRunner) Start() error {
	questions, err := qr.Reader.FetchQuestions()

	if err != nil {
		return err
	}

	if len(questions) == 0 {
		return errors.New("No questions found")
	}

	var totalMarks = 0

	for _, q := range questions {
		var ans string
		fmt.Print("\n" + q.Quesion)

		if _, err := fmt.Scanln(&ans); err != nil {
			return err
		}

		if ans == q.Answer {
			totalMarks = totalMarks + int(q.Weight)
			fmt.Println("[Correct]")
		} else {
			fmt.Println("[Incorrect]")
		}
	}

	fmt.Println("\n\nResult: ", totalMarks)
	return nil
}
