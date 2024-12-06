package quiz

// Answer represents a single answer variant for a question.
//
// Fields:
//
//	Num: The numerical index of the answer (e.g. 1, 2, 3, etc.).
//	Body: The text of the answer itself.
type answer struct {
	Num  int    `json:"number"`
	Body string `json:"body"`
}
