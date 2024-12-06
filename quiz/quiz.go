package quiz

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var ErrEmptyQuiz = fmt.Errorf("quiz does not contain any questions")

// Quiz represents a collection of questions.
type Quiz struct {
	Questions []question
}

// Size returns the number of questions in the quiz.
func (q *Quiz) Size() int {
	return len(q.Questions)
}

// Topic returns the topic of the first question in the quiz, which is also the
// topic of all quiz. If the quiz is empty, it returns an error.
func (q *Quiz) Topic() (string, error) {
	if q.Size() < 1 {
		return "", ErrEmptyQuiz
	}

	return q.Questions[0].Topic, nil
}

// load loads the quiz from a JSON file at the given path.
func (q *Quiz) load(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&q.Questions)
	if err != nil {
		log.Fatal(err)
	}
}

func New(jsonPath string) *Quiz {
	q := Quiz{}
	q.load(jsonPath)

	return &q
}
