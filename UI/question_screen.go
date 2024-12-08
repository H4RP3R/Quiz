package ui

import (
	"quizapp/colors"
	"quizapp/quiz"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	QuestionCounterChan chan struct{}
	OptionChan          chan int
)

func init() {
	QuestionCounterChan = make(chan struct{}, 0)
	OptionChan = make(chan int, 0)
}

// QuestionScreen displays a question and its options.
type QuestionScreen struct {
	// TB is the textbox that displays the question.
	TB *TextBox
	// Opt1, Opt2, ... are the buttons representing the question options.
	Opt1 *Button
	Opt2 *Button
	Opt3 *Button
	Opt4 *Button
}

// Draw renders the question screen on the destination image.
func (qs *QuestionScreen) Draw(dst *ebiten.Image) {
	qs.TB.Draw(dst)
	qs.Opt1.Draw(dst)
	qs.Opt2.Draw(dst)
	qs.Opt3.Draw(dst)
	qs.Opt4.Draw(dst)
}

// Update updates the state of the question screen based on user input.
func (qs *QuestionScreen) Update() {
	qs.Opt1.Update()
	qs.Opt2.Update()
	qs.Opt3.Update()
	qs.Opt4.Update()
}

func NewQuestionScreen(q quiz.Question, scrW, scrH int) *QuestionScreen {
	w, h := 880, 300
	qs := QuestionScreen{
		TB: NewTextBox(w, h, colors.Black, q.Body, scrW/2-w/2, 80),
		Opt1: NewButton(400, 80, colors.Blue, q.Answers[0].Body, FaceSourceRegular, 24, 72, 440,
			// Send question counter signal and option value in a strict order to ensure
			// that the receiver code receives the signal before trying to receive the option value.
			func() { QuestionCounterChan <- struct{}{}; OptionChan <- 1 }),
		Opt2: NewButton(400, 80, colors.Blue, q.Answers[1].Body, FaceSourceRegular, 24, 552, 440,
			func() { QuestionCounterChan <- struct{}{}; OptionChan <- 2 }),
		Opt3: NewButton(400, 80, colors.Blue, q.Answers[2].Body, FaceSourceRegular, 24, 72, 560,
			func() { QuestionCounterChan <- struct{}{}; OptionChan <- 3 }),
		Opt4: NewButton(400, 80, colors.Blue, q.Answers[3].Body, FaceSourceRegular, 24, 552, 560,
			func() { QuestionCounterChan <- struct{}{}; OptionChan <- 4 }),
	}

	return &qs
}
