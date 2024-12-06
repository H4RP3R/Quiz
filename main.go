package main

import (
	"fmt"
	"log"
	ui "quizapp/UI"
	"quizapp/colors"
	"quizapp/quiz"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

type Game struct {
	startButton *ui.Button
	status      GameStatus
}

func NewGame() *Game {
	g := &Game{}
	g.startButton = ui.NewButton(180, 60, colors.Blue, "СТАРТ", screenWidth/2-90, 480, func() { g.status = Quiz })
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
		g.startButton.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Quiz")

	q := quiz.New("quiz/test_questions.json")

	for _, v := range q.Questions {
		fmt.Println(v)
	}
	fmt.Println(q.Size())

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
