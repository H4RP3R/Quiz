package quiz

import (
	"encoding/json"
	"log"
	"os"
)

type quiz struct {
	Questions []question
}

func (q *quiz) Size() int {
	return len(q.Questions)
}

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
