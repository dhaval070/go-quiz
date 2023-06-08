// Fetch quiz questions from csv data source [ fields: question, correct answer, weight ]
package client

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"quiz/model"
	"strconv"
)

type QuizReader interface {
	FetchQuestions() ([]model.QuizQuestion, error)
}

type FileReader struct {
	Path string
}

func (fr *FileReader) FetchQuestions() ([]model.QuizQuestion, error) {
	fh, err := os.Open(fr.Path)

	if err != nil {
		return nil, err
	}

	var result = []model.QuizQuestion{}

	csvreader := csv.NewReader(fh)

	for {
		fields, err := csvreader.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		weight, err := strconv.ParseInt(fields[2], 10, 8)

		if err != nil {
			return nil, err
		}

		result = append(result, model.QuizQuestion{
			Quesion: fields[0],
			Answer:  fields[1],
			Weight:  int8(weight),
		})
	}

	return result, nil
}
