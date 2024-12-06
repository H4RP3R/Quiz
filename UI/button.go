package ui

import (
	"bytes"
	"image/color"
	"log"
	"quizapp/colors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/gobold"
)

var (
	uiImage      *ebiten.Image
	uiFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(gobold.TTF))
	if err != nil {
		log.Fatal(err)
	}
	uiFaceSource = s
}

type Button struct {
	Background *ebiten.Image
	Color      color.RGBA
	Text       string
	Width      int
	Heigh      int
	Coord

	Hover   bool
	BorderW int
}

func (b *Button) Draw(dst *ebiten.Image) {
	b.Background.Fill(b.Color)

	opText := &text.DrawOptions{}
	opText.GeoM.Translate(float64(b.Width)/2, float64(b.Heigh)/2)
	opText.ColorScale.ScaleWithColor(color.White)
	opText.PrimaryAlign = text.AlignCenter
	opText.SecondaryAlign = text.AlignCenter
	text.Draw(b.Background, b.Text, &text.GoTextFace{
		Source: uiFaceSource,
		Size:   48,
	}, opText)

	opImg := &ebiten.DrawImageOptions{}
	opImg.GeoM.Translate(float64(b.X), float64(b.Y))

	if b.Hover {
		vector.StrokeRect(b.Background, 1, 1, float32(b.Width)-1, float32(b.Heigh)-1, 3, colors.Orange, false)
	}

	dst.DrawImage(b.Background, opImg)

}

func (b *Button) Update() {
	x, y := ebiten.CursorPosition()

	if x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Heigh {
		b.Hover = true
	} else {
		b.Hover = false
	}

	b.click()
}

func (b *Button) click() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(ebiten.MouseButtonLeft)) && b.Hover {
		log.Println("Click!")
	}
}

func NewButton(w, h int, c color.RGBA, t string, posX, posY int) *Button {
	b := Button{
		Color:      c,
		Text:       t,
		Width:      w,
		Heigh:      h,
		Background: ebiten.NewImage(w, h),
		Coord:      Coord{posX, posY},
	}

	return &b
}
