package main

import (
	"fmt"
	"log"
	"quizapp/quiz"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	screenWidth  int = 1024
	screenHeight int = 768
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	infoStr := fmt.Sprintf("TPS: %.0f, FPS: %.0f", ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, infoStr)
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

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
