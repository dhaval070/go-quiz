package main

import (
	"quiz/client"
	"quiz/quizrunner"
)

func main() {
	reader := &client.FileReader{
		Path: "quiz.csv",
	}

	app := quizrunner.QuizRunner{
		Reader: reader,
	}

	err := app.Start()

	if err != nil {
		panic(err)
	}

}
