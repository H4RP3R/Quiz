package main

import (
	"fmt"
	"image/color"
	"log"
	ui "quizapp/UI"
	"quizapp/quiz"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	screenWidth  int = 1024
	screenHeight int = 768
)

var (
	Teal color.RGBA = color.RGBA{35, 186, 155, 255}
	Blue color.RGBA = color.RGBA{52, 73, 94, 255}
)

type Game struct {
	startButton *ui.Button
}

func NewGame() *Game {
	g := &Game{}
	g.startButton = ui.NewButton(180, 60, Blue, "СТАРТ", screenWidth/2-90, 480)

	return g
}

func (g *Game) Update() error {
	g.startButton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(Teal)

	infoStr := fmt.Sprintf("TPS: %.0f, FPS: %.0f", ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, infoStr)

	g.startButton.Draw(screen)
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
