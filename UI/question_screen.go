package ui

import (
	"log"
	"quizapp/colors"
	"quizapp/quiz"

	"github.com/hajimehoshi/ebiten/v2"
)

type QuestionScreen struct {
	TB   *TextBox
	Opt1 *Button
	Opt2 *Button
	Opt3 *Button
	Opt4 *Button
}

func (qs *QuestionScreen) Draw(dst *ebiten.Image) {
	qs.TB.Draw(dst)
	qs.Opt1.Draw(dst)
	qs.Opt2.Draw(dst)
	qs.Opt3.Draw(dst)
	qs.Opt4.Draw(dst)
}

func (qs *QuestionScreen) Update() {
	qs.Opt1.Update()
	qs.Opt2.Update()
	qs.Opt3.Update()
	qs.Opt4.Update()
}

func NewQuestionScreen(q quiz.Question, scrW, scrH int) *QuestionScreen {
	w, h := 880, 300
	qs := QuestionScreen{
		TB:   NewTextBox(w, h, colors.Black, q.Body, scrW/2-w/2, 80),
		Opt1: NewButton(400, 80, colors.Blue, q.Answers[0].Body, FaceSourceRegular, 24, 72, 440, func() { log.Println("Click 1") }),
		Opt2: NewButton(400, 80, colors.Blue, q.Answers[1].Body, FaceSourceRegular, 24, 552, 440, func() { log.Println("Click 2") }),
		Opt3: NewButton(400, 80, colors.Blue, q.Answers[2].Body, FaceSourceRegular, 24, 72, 560, func() { log.Println("Click 3") }),
		Opt4: NewButton(400, 80, colors.Blue, q.Answers[3].Body, FaceSourceRegular, 24, 552, 560, func() { log.Println("Click 4") }),
	}

	return &qs
}
