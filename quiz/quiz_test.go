package quiz

import (
	"errors"
	"testing"
)

func TestQuizNewNoQuestionsFile(t *testing.T) {
	var invalidPath = ""
	wantErr := ErrQuestionsJson
	_, err := New(invalidPath)
	if !errors.Is(err, wantErr) {
		t.Errorf("want error: %v, got error: %v", wantErr, err)
	}
}

func TestQuizNew(t *testing.T) {
	var validPath = "test_questions.json"
	_, err := New(validPath)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestQuizTopicEmptyQuiz(t *testing.T) {
	wantErr := ErrEmptyQuiz
	wantTopic := ""
	q := Quiz{Questions: []Question{}}
	topic, err := q.Topic()
	if !errors.Is(err, ErrEmptyQuiz) {
		t.Errorf("want error: %v, got error: %v", wantErr, err)
	}
	if topic != wantTopic {
		t.Errorf("want topic: %v, got topic: %v", wantTopic, topic)
	}
}

func TestQuizTopic(t *testing.T) {
	wantTopic := "History"
	q, err := New("test_questions.json")
	topic, err := q.Topic()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if topic != wantTopic {
		t.Errorf("want topic: %v, got topic: %v", wantTopic, topic)
	}
}

func TestQuizSize(t *testing.T) {
	var jsonPath = "test_questions.json"
	wantSize := 2
	q, err := New(jsonPath)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	gotSize := q.Size()
	if gotSize != wantSize {
		t.Errorf("quiz size: want %d, got %d", wantSize, gotSize)
	}
}
