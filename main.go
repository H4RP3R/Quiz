package main

import (
	"fmt"
	"image/color"
	"log"
	ui "quizapp/UI"
	"quizapp/colors"
	"quizapp/quiz"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameStatus string

const (
	MainMenu  GameStatus = "main menu"
	Quiz      GameStatus = "quiz"
	Statistic GameStatus = "statistic"
)

var (
	screenWidth  int = 1024
	screenHeight int = 768
)

var (
	currentQuiz *quiz.Quiz
	questionNum int = 0
)

type Game struct {
	title           string
	startButton     *ui.Button
	startOverButton *ui.Button
	status          GameStatus
	qScreens        []ui.QuestionScreen

	correctAnswers int
	resultStr      string
}

func (g *Game) reset() {
	g.status = MainMenu
	g.correctAnswers = 0
	questionNum = 0
}

func NewGame() *Game {
	g := Game{}
	g.startButton = ui.NewButton(180, 60, colors.Blue, "СТАРТ", ui.FaceSourceBold, 48, screenWidth/2-90, 480, func() { g.status = Quiz })
	g.startOverButton = ui.NewButton(240, 60, colors.Blue, "ПОВТОРИТЬ", ui.FaceSourceBold, 36, screenWidth/2-90, 480, func() { g.reset() })
	topic, err := currentQuiz.Topic()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: fix numbers endings
	g.title = fmt.Sprintf("Викторина по теме %q (%d вопросов)", topic, currentQuiz.Size())
	g.status = MainMenu

	for _, q := range currentQuiz.Questions {
		qScreen := ui.NewQuestionScreen(q, screenWidth, screenHeight)
		g.qScreens = append(g.qScreens, *qScreen)
	}

	g.correctAnswers = 0
	g.resultStr = "Правильных ответов: %d из %d"

	// Listen for question counter signals and process user answers. The sender
	// code ensures that the question counter signal is sent before the option
	// value, so we can safely receive from OptionChan after receiving the signal.
	go func() {
		for {
			// Wait for the next question counter signal.
			<-ui.QuestionCounterChan
			if option, ok := <-ui.OptionChan; ok {
				// Check if the user's answer is correct.
				if option == currentQuiz.Questions[questionNum].RightAnsw {
					g.correctAnswers++
				}
			}
			// Move on to the next question.
			questionNum++
			if questionNum == currentQuiz.Size() {
				g.status = Statistic
			}
		}
	}()

	return &g
}

func (g *Game) Update() error {
	switch g.status {
	case MainMenu:
		g.startButton.Update()
	case Quiz:
		if questionNum < currentQuiz.Size() {
			g.qScreens[questionNum].Update()
		}
	case Statistic:
		g.startOverButton.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colors.Teal)

	infoStr := fmt.Sprintf("TPS: %.0f, FPS: %.0f\nStatus: %s", ebiten.ActualTPS(), ebiten.ActualFPS(), g.status)
	ebitenutil.DebugPrint(screen, infoStr)

	switch g.status {
	case MainMenu:
		opText := &text.DrawOptions{}
		opText.GeoM.Translate(float64(screenWidth)/2, 250)
		opText.ColorScale.ScaleWithColor(color.White)
		opText.PrimaryAlign = text.AlignCenter
		opText.SecondaryAlign = text.AlignCenter
		text.Draw(screen, g.title, &text.GoTextFace{
			Source: ui.FaceSourceRegular,
			Size:   36,
		}, opText)
		g.startButton.Draw(screen)
	case Quiz:
		if questionNum < currentQuiz.Size() {
			g.qScreens[questionNum].Draw(screen)
		}
	case Statistic:
		opText := &text.DrawOptions{}
		opText.GeoM.Translate(float64(screenWidth)/2, 250)
		opText.ColorScale.ScaleWithColor(color.White)
		opText.PrimaryAlign = text.AlignCenter
		opText.SecondaryAlign = text.AlignCenter
		text.Draw(screen, fmt.Sprintf(g.resultStr, g.correctAnswers, currentQuiz.Size()),
			&text.GoTextFace{
				Source: ui.FaceSourceRegular,
				Size:   36,
			}, opText)
		g.startOverButton.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	currentQuiz = quiz.New("quiz/test_questions.json")
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Quiz")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
