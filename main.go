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

var currentQuiz *quiz.Quiz

type Game struct {
	title       string
	startButton *ui.Button
	status      GameStatus
}

func NewGame() *Game {
	g := &Game{}
	g.startButton = ui.NewButton(180, 60, colors.Blue, "СТАРТ", screenWidth/2-90, 480, func() { g.status = Quiz })
	topic, err := currentQuiz.Topic()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: fix numbers endings
	g.title = fmt.Sprintf("Викторина по теме %q (%d вопросов)", topic, currentQuiz.Size())
	g.status = MainMenu

	return g
}

func (g *Game) Update() error {
	switch g.status {
	case MainMenu:
		g.startButton.Update()
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
