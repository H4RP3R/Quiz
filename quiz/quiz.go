package quiz

import (
	"encoding/json"
	"log"
	"os"
)

// Quiz represents a collection of questions.
type quiz struct {
	Questions []question
}

// Size returns the number of questions in the quiz.
func (q *quiz) Size() int {
	return len(q.Questions)
}

// load loads the quiz from a JSON file at the given path.
func (q *quiz) load(path string) {
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

func New(jsonPath string) *quiz {
	q := quiz{}
	q.load(jsonPath)

	return &q
}
