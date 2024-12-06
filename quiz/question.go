package quiz

// Question represents a single quiz question.
//
// Fields:
//
//	Topic: The topic or category of the question.
//	ID: A unique identifier for the question.
//	Body: The text of the question itself.
//	Answers: A list of possible answer variants.
//	RightAnsw: The index of the correct answer in the Answers slice.
type question struct {
	Topic     string   `json:"topic"`
	ID        int      `json:"id"`
	Body      string   `json:"body"`
	Answers   []answer `json:"answers"`
	RightAnsw int      `json:"right_answer"`
}
