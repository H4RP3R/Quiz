package quiz

type question struct {
	Topic     string   `json:"topic"`
	ID        int      `json:"id"`
	Body      string   `json:"body"`
	Answers   []answer `json:"answers"`
	RightAnsw int      `json:"right_answer"`
}
